package product

import (
	"github.com/gin-gonic/gin"
	middleware "shopping-cart/middlerware"
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

	r.Use(middleware.JWTAuthMiddleware())
	{
		adminRoute(h, r)
	}

	return h
}

func newRoute(h *Product, r *gin.RouterGroup) {
	r.GET("/products/:id", h.GetProduct)
}

func adminRoute(h *Product, r *gin.RouterGroup) {
	r.POST("admin/products", h.CreateProduct)
	r.GET("admin/products", h.GetAllProducts)
	r.PATCH("admin/products/:id", h.UpdateProduct)
	r.DELETE("admin/products/:id", h.DeleteProduct)
}
