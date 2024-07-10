package service

import (
	"errors"
	"shopping-cart/builder"
	"shopping-cart/model/database"
	"shopping-cart/model/datatransfer/product"
	"shopping-cart/repository"
)

type ProductService interface {
	UpdateProduct(id int, productDto *product.Update) (*database.Product, error)
	CreateProduct(productDto *product.Payload) (*database.Product, error)
	DeleteProduct(id int) error
	FindAllProducts() ([]database.Product, error)
	FindByID(id int) (*database.Product, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		productRepo: repo,
	}
}

func (s *productService) UpdateProduct(id int, productDto *product.Update) (*database.Product, error) {
	product, err := s.productRepo.InternalFindByID(id)
	if err != nil {
		return nil, err
	}

	product = builder.NewProductBuilder().
		SetName(productDto.Name).
		SetPicture(productDto.Picture).
		SetPrice(productDto.Price).
		SetStock(productDto.Stock).
		SetDescription(productDto.Description).
		SetExpirationTime(productDto.ExpirationTime).
		Build()

	err = s.productRepo.Update(product)

	if err != nil {
		return nil, err
	}

	return product, nil
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

func (s *productService) FindAllProducts() ([]database.Product, error) {
	return s.productRepo.FindAll()
}

func (s *productService) FindByID(id int) (*database.Product, error) {
	return s.productRepo.FindByID(id)
}
