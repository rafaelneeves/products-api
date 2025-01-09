package routes

import (
	"net/http"
	"products-api/internal/controllers"
	"products-api/internal/middleware" // Importa o middleware
)

func SetupRoutes(productController *controllers.ProductController) *http.ServeMux {
	router := http.NewServeMux()

	// Middleware de autenticação
	authMiddleware := middleware.BasicAuthMiddleware("admin", "1234")

	router.HandleFunc("/hello", greet)

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
