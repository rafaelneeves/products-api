package routes

import (
	"net/http"
	"products-api/internal/controllers"
	"products-api/internal/middleware"
)

func SetupRoutes(productController *controllers.ProductController) *http.ServeMux {
	router := http.NewServeMux()

	authMiddleware := middleware.JWTMiddleware

	router.HandleFunc("/hello", greet)
	router.HandleFunc("/login", controllers.Login)

	router.Handle("/products", authMiddleware(http.HandlerFunc(productController.GetProductsAll)))
	router.Handle("/product/", authMiddleware(http.HandlerFunc(productController.GetProductByID)))
	router.Handle("/product-create", authMiddleware(http.HandlerFunc(productController.CreateProduct)))
	router.Handle("/product-update", authMiddleware(http.HandlerFunc(productController.UpdateProduct)))
	router.Handle("/product-delete/", authMiddleware(http.HandlerFunc(productController.DeleteProduct)))

	return router
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
}
