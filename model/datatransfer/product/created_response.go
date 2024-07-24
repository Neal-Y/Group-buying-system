package product

import "shopping-cart/model/database"

type CreatedResponse struct {
	Product *database.Product
	Url     string
}
