package service

import (
	"errors"
	"fmt"
	"shopping-cart/builder"
	"shopping-cart/model/database"
	"shopping-cart/model/datatransfer/order"
	"shopping-cart/repository"
	"shopping-cart/util"
	"time"
)

type OrderService interface {
	CreateOrder(orderRequest *order.Request) (*database.Order, error)
	UpdateOrderStatusAndNote(id int, orderRequest *order.StatusRequest) (*database.Order, error)
	DeleteOrder(id int) error
	ListHistoryOrdersByDisplayNameAndProductID(displayName string, productID int) ([]database.Order, error)
	SearchOrders(params util.SearchContainer) ([]database.Order, int64, error)
	GetRevenueByTimePeriod(startDate, endDate time.Time) (float64, error)
}

type orderService struct {
	orderRepo           repository.OrderRepository
	productRepo         repository.ProductRepository
	userRepo            repository.UserRepository
	notificationService NotificationService
	notificationCache   *util.NotificationCache
}

func NewOrderService(orderRepo repository.OrderRepository, productRepo repository.ProductRepository, userRepo repository.UserRepository, notificationService NotificationService, notificationCache *util.NotificationCache) OrderService {
	return &orderService{
		orderRepo:           orderRepo,
		productRepo:         productRepo,
		userRepo:            userRepo,
		notificationService: notificationService,
		notificationCache:   notificationCache,
	}
}

func validateOrderRequest(s *orderService, orderRequest *order.Request) (float64, map[int]*database.Product, error) {
	totalPrice := 0.0
	productMap := make(map[int]*database.Product)

	productIDs := make([]int, 0, len(orderRequest.OrderDetails))
	for _, detail := range orderRequest.OrderDetails {
		if detail.Quantity <= 0 {
			return 0, nil, errors.New("Quantity must be greater than zero")
		}
		productIDs = append(productIDs, detail.ProductID)
	}

	products, err := s.productRepo.FindByIDs(productIDs)
	if err != nil {
		return 0, nil, err
	}

	for _, product := range products {
		productMap[product.ID] = product
	}

	for i, detail := range orderRequest.OrderDetails {
		product, exists := productMap[detail.ProductID]
		if !exists {
			return 0, nil, errors.New("Product not found or already sold out")
		}

		if product.Stock < detail.Quantity {
			return 0, nil, errors.New(fmt.Sprintf("%s庫存不足，只剩下%d個", product.Name, product.Stock))
		}

		if time.Now().After(product.ExpirationTime) {
			return 0, nil, errors.New("Product " + product.Name + " is expired")
		}

		if s.notificationCache.Get(product.ID) == 0 {
			s.notificationCache.Set(product.ID, product.Stock)
		}

		threshold := int(0.5 * float64(s.notificationCache.Get(product.ID)))

		if product.Stock-detail.Quantity < threshold {
			s.notificationService.Notify(product.ID, product.Name, product.Stock-detail.Quantity)
			s.notificationCache.Set(product.ID, product.Stock-detail.Quantity)
		}

		orderRequest.OrderDetails[i].Price = product.Price
		totalPrice += float64(detail.Quantity) * product.Price
	}

	return totalPrice, productMap, nil
}

func (s *orderService) CreateOrder(orderRequest *order.Request) (*database.Order, error) {
	totalPrice, productMap, err := validateOrderRequest(s, orderRequest)
	if err != nil {
		return nil, err
	}

	order := builder.NewOrderBuilder().
		SetUserID(orderRequest.UserID).
		SetTotalPrice(totalPrice).
		SetNote(orderRequest.Note).
		SetStatus("pending").
		SetOrderDetails(orderRequest.OrderDetails).
		Build()

	tx := s.orderRepo.BeginTransaction()

	err = s.orderRepo.Create(order)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var productsToUpdate []*database.Product
	for _, detail := range order.OrderDetails {
		product := productMap[detail.ProductID]
		product.Stock -= detail.Quantity
		productsToUpdate = append(productsToUpdate, product)
	}

	err = s.productRepo.BatchUpdate(productsToUpdate)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return order, nil
}

func (s *orderService) UpdateOrderStatusAndNote(id int, orderRequest *order.StatusRequest) (*database.Order, error) {
	order, err := s.orderRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("Order not found")
	}

	if orderRequest.Status != "" {
		order.Status = orderRequest.Status
	}

	if orderRequest.Note != "" {
		order.Note = orderRequest.Note
	}

	err = s.orderRepo.Update(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *orderService) DeleteOrder(id int) error {
	order, err := s.orderRepo.FindByID(id)
	if err != nil {
		return errors.New("Order not found")
	}

	tx := s.orderRepo.BeginTransaction()

	for _, detail := range order.OrderDetails {
		product, err := s.productRepo.InternalFindByID(detail.ProductID)
		if err != nil {
			tx.Rollback()
			return err
		}
		product.Stock += detail.Quantity
		err = s.productRepo.Update(product)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = s.orderRepo.SoftDelete(order)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *orderService) ListHistoryOrdersByDisplayNameAndProductID(displayName string, productID int) ([]database.Order, error) {
	user, err := s.userRepo.FindByDisplayName(displayName)
	if err != nil {
		return nil, err
	}

	return s.orderRepo.FindByUserIDAndProductID(user.ID, productID)
}

func (s *orderService) SearchOrders(params util.SearchContainer) ([]database.Order, int64, error) {
	return s.orderRepo.SearchOrders(params.Keyword, params.StartDate, params.EndDate, params.Offset, params.Limit)
}

func (s *orderService) GetRevenueByTimePeriod(startDate, endDate time.Time) (float64, error) {
	return s.orderRepo.GetRevenueByTimePeriod(startDate, endDate)
}
