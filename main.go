package main

import (
	"log"

	"github.com/gabrieldiaspereira/echoGoApi/config"
	"github.com/gabrieldiaspereira/echoGoApi/domain/product/controllers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	db := config.InitDB()

	route := echo.New()

	productController := controllers.NewProductController(db)
	route.POST("product", productController.Create)
	route.GET("product", productController.GetAll)
	route.GET("product/:id_product", productController.GetById)
	route.PUT("product/:id_product", productController.Update)
	route.DELETE("product/:id_product", productController.Delete)

	route.Start(":8080")
}
