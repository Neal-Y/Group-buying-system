package repository

import (
	"gorm.io/gorm"
	"shopping-cart/infrastructure"
	"shopping-cart/model/database"
	"time"
)

type ProductRepository interface {
	FindByID(id int) (*database.Product, error)
	InternalFindByID(id int) (*database.Product, error)
	Update(product *database.Product) error
	Create(product *database.Product) error
	FindAll() ([]database.Product, error)
	FindByName(name string, product *database.Product) error
	BatchUpdate(products []*database.Product) error
	FindByIDs(ids []int) ([]*database.Product, error)
	SoftDelete(product *database.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		db: infrastructure.Db,
	}
}

func (r *productRepository) Create(product *database.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindByID(id int) (*database.Product, error) {
	var product database.Product
	err := r.db.Where("is_sold_out = ?", false).First(&product, id).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) InternalFindByID(id int) (*database.Product, error) {
	var product database.Product
	err := r.db.First(&product, id).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) Update(product *database.Product) error {
	return r.db.Updates(product).Error
}

func (r *productRepository) FindAll() ([]database.Product, error) {
	var products []database.Product
	err := r.db.Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) FindByName(name string, product *database.Product) error {
	return r.db.Where("name = ? AND is_sold_out = ?", name, false).First(product).Error
}

func (r *productRepository) SoftDelete(product *database.Product) error {
	now := time.Now()
	product.IsSoldOut = true
	product.SoldOutAt = &now
	return r.db.Save(product).Error
}

func (r *productRepository) BatchUpdate(products []*database.Product) error {
	for _, product := range products {
		err := r.db.Updates(product).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *productRepository) FindByIDs(ids []int) ([]*database.Product, error) {
	var products []*database.Product
	err := r.db.Where("id IN (?) AND is_sold_out = ?", ids, false).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
