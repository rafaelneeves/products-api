package main

import (
	"net/http"
	"products-api/config"
	"products-api/internal/controllers"
	"products-api/internal/routes"
)

func main() {
	db := config.ConnectDatabase()
	defer db.Close()

	productController := controllers.NewProductController(db)

	router := routes.SetupRoutes(productController)
	http.ListenAndServe(":8080", router)
}
