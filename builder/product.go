package builder

import (
	"shopping-cart/model/database"
	"time"
)

type ProductBuilder struct {
	product *database.Product
}

func NewProductBuilder() *ProductBuilder {
	return &ProductBuilder{product: &database.Product{}}
}

func (b *ProductBuilder) SetID(id int) *ProductBuilder {
	b.product.ID = id
	return b
}

func (b *ProductBuilder) SetName(name string) *ProductBuilder {
	b.product.Name = name
	return b
}

func (b *ProductBuilder) SetPicture(picture string) *ProductBuilder {
	b.product.Picture = picture
	return b
}

func (b *ProductBuilder) SetPrice(price float64) *ProductBuilder {
	b.product.Price = price
	return b
}

func (b *ProductBuilder) SetStock(stock int) *ProductBuilder {
	b.product.Stock = stock
	return b
}

func (b *ProductBuilder) SetDescription(description string) *ProductBuilder {
	b.product.Description = description
	return b
}

func (b *ProductBuilder) SetExpirationTime(expirationTime time.Time) *ProductBuilder {
	b.product.ExpirationTime = expirationTime
	return b
}

func (b *ProductBuilder) SetSupplier(supplier string) *ProductBuilder {
	b.product.Supplier = supplier
	return b
}

func (b *ProductBuilder) Build() *database.Product {
	return b.product
}
