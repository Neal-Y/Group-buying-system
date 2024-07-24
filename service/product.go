package service

import (
	"errors"
	"shopping-cart/builder"
	"shopping-cart/model/database"
	"shopping-cart/model/datatransfer/product"
	"shopping-cart/repository"
	"shopping-cart/util"
)

type ProductService interface {
	UpdateProduct(id int, productDto *product.Update) error
	CreateProduct(productDto *product.Payload) (*database.Product, error)
	DeleteProduct(id int) error
	FindByID(id int) (*database.Product, error)
	SearchProducts(params util.SearchContainer) ([]database.Product, int64, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		productRepo: repo,
	}
}

func (s *productService) UpdateProduct(id int, productDto *product.Update) error {
	product, err := s.productRepo.InternalFindByID(id)
	if err != nil {
		return err
	}

	product = builder.NewProductBuilder().
		SetID(product.ID).
		SetName(productDto.Name).
		SetPicture(productDto.Picture).
		SetPrice(productDto.Price).
		SetStock(productDto.Stock).
		SetDescription(productDto.Description).
		SetExpirationTime(productDto.ExpirationTime).
		SetSupplier(productDto.Supplier).
		Build()

	return s.productRepo.Update(product)
}

func (s *productService) CreateProduct(productDto *product.Payload) (*database.Product, error) {
	var check database.Product
	err := s.productRepo.FindByName(productDto.Name, &check)
	if err == nil {
		return nil, errors.New("product name already exists")
	}

	product := builder.NewProductBuilder().
		SetName(productDto.Name).
		SetPicture(productDto.Picture).
		SetPrice(productDto.Price).
		SetStock(productDto.Stock).
		SetDescription(productDto.Description).
		SetExpirationTime(productDto.ExpirationTime).
		SetSupplier(productDto.Supplier).
		Build()

	err = s.productRepo.Create(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) DeleteProduct(id int) error {
	product, err := s.productRepo.InternalFindByID(id)
	if err != nil {
		return err
	}

	return s.productRepo.SoftDelete(product)
}

func (s *productService) FindByID(id int) (*database.Product, error) {
	return s.productRepo.FindByID(id)
}

func (s *productService) SearchProducts(params util.SearchContainer) ([]database.Product, int64, error) {
	return s.productRepo.SearchProducts(params.Keyword, params.StartDate, params.EndDate, params.Offset, params.Limit)
}
