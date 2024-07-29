package repository

import (
	"gorm.io/gorm"
	"log"
	"shopping-cart/infrastructure"
	"shopping-cart/model/database"
	"time"
)

type ProductRepository interface {
	FindByID(id int) (*database.Product, error)
	InternalFindByID(id int) (*database.Product, error)
	Update(product *database.Product) error
	Create(product *database.Product) error
	FindByName(name string, product *database.Product) error
	BatchUpdate(products []*database.Product) error
	FindByIDs(ids []int) ([]*database.Product, error)
	SoftDelete(product *database.Product) error
	SearchProducts(keyword string, startDate, endDate time.Time, offset int, limit int) ([]database.Product, int64, error)
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
	err := r.db.Where("id = ? AND is_sold_out = ?", id, false).First(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) InternalFindByID(id int) (*database.Product, error) {
	var product database.Product
	err := r.db.Where("id = ?", id).First(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) Update(product *database.Product) error {
	return r.db.Updates(product).Error
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
		err := r.db.Save(product).Error
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

func (r *productRepository) SearchProducts(keyword string, startDate, endDate time.Time, offset int, limit int) ([]database.Product, int64, error) {
	var products []database.Product
	var count int64
	log.Printf("keyword: %s, startDate: %s, endDate: %s, offset: %d, limit: %d", keyword, startDate, endDate, offset, limit)
	query := r.db.Model(&database.Product{})

	if keyword != "" {
		query = query.Where("name LIKE ? OR description LIKE ? OR supplier LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	err := query.Count(&count).
		Offset(offset).
		Limit(limit).
		Find(&products).Error

	if err != nil {
		return nil, 0, err
	}
	return products, count, nil
}
