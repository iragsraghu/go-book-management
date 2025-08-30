package routes

import (
	"github.com/iragsraghu/go-book-management/controllers"

	"github.com/gin-gonic/gin"
)

func BookRoutes(r *gin.Engine) {
	r.POST("/books", controllers.CreateBook)
	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
}
