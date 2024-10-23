package main

import (
	"context"
	"rest-playground/common/app"
	"rest-playground/common/postgresql"
	"rest-playground/controller/handler"

	"rest-playground/repository"
	"rest-playground/service"

	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	e := echo.New()

	configurationManager := app.NewConfigurationManager()

	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgreSqlConfig)

	productRepository := repository.NewProductRepository(dbPool)

	productService := service.NewProductService(productRepository)

	productController := handler.NewProductController(productService)

	productController.RegisterRoutes(e)

	e.Start("localhost:8080")
}