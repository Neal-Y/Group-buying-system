package product

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/middleware"
	"shopping-cart/repository"
	"shopping-cart/service"
)

type Product struct {
	productService service.ProductService
}

func NewProductController(r *gin.RouterGroup) *Product {
	productRepo := repository.NewProductRepository()

	productService := service.NewProductService(productRepo)

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
	adminRoute.Use(middleware.JWTAuthMiddleware())
	adminRoute.POST("", h.CreateProduct)
	adminRoute.GET("", h.GetAllProducts)
	adminRoute.PATCH("/:id", h.UpdateProduct)
	adminRoute.DELETE("/:id", h.DeleteProduct)
}
