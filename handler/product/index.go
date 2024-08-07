package product

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/constant"
	"shopping-cart/middleware"
	"shopping-cart/service"
)

type Product struct {
	productService service.ProductService
}

func NewProductController(r *gin.RouterGroup, productService service.ProductService) *Product {
	h := &Product{
		productService: productService,
	}

	newRoute(h, r)
	adminRoute(h, r)

	return h
}

func newRoute(h *Product, r *gin.RouterGroup) {
	r.GET("/products/:id", h.GetProduct)
}

func adminRoute(h *Product, r *gin.RouterGroup) {
	adminRoute := r.Group("/admin/products")
	adminRoute.Use(middleware.JWTAuthMiddleware(constant.AdminType))
	adminRoute.POST("", h.CreateProduct)
	adminRoute.PATCH("/:id", h.UpdateProduct)
	adminRoute.DELETE("/:id", h.DeleteProduct)
	adminRoute.GET("/search", h.SearchProducts)
	adminRoute.GET("/:id", h.GetByID)
}
