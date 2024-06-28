package datatransfer

import "shopping-cart/model/database"

type OrderRequest struct {
	UserID       int                    `json:"user_id"`
	Note         string                 `json:"note"`
	OrderDetails []database.OrderDetail `json:"order_details"`
}
