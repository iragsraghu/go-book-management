// @title           Book Management API
// @version         1.0
// @description     A simple REST API using Gin + GORM + MySQL.
// @host            localhost:8080
// @BasePath        /
package main

import (
	_ "github.com/iragsraghu/go-book-management/docs" // swagger docs

	"github.com/gin-gonic/gin"
	"github.com/iragsraghu/go-book-management/config"
	"github.com/iragsraghu/go-book-management/models"
	"github.com/iragsraghu/go-book-management/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin + GORM Example API
// @version 1.0
// @description This is a simple CRUD API with Gin and GORM
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Connect DB
	config.ConnectDatabase()

	// Auto Migrate Model
	config.DB.AutoMigrate(&models.Book{})

	// Routes
	routes.BookRoutes(r)

	// Run server
	r.Run(":8080")
}
