package controllers

import (
	"net/http"

	"github.com/iragsraghu/go-book-management/config"
	"github.com/iragsraghu/go-book-management/models"

	"github.com/gin-gonic/gin"
)

// CreateBook godoc
// @Summary      Create a new book
// @Description  Adds a new book to the database
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      models.Book  true  "Book Data"
// @Success      200   {object}  models.Book
// @Failure      400   {object}  map[string]interface{}
// @Router       /books [post]
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&book)
	c.JSON(http.StatusOK, book)
}

// GetBooks godoc
// @Summary      List all books
// @Description  Get all books from the database
// @Tags         books
// @Produce      json
// @Success      200   {array}   models.Book
// @Router       /books [get]
func GetBooks(c *gin.Context) {
	var books []models.Book
	config.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

// GetBook godoc
// @Summary      Get a single book
// @Description  Get details of a book by ID
// @Tags         books
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  models.Book
// @Failure      404  {object}  map[string]interface{}
// @Router       /books/{id} [get]
func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// UpdateBook godoc
// @Summary      Update an existing book
// @Description  Update a book's details by ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Book ID"
// @Param        book  body      models.Book true  "Updated Book Data"
// @Success      200   {object}  models.Book
// @Failure      400   {object}  map[string]interface{}
// @Failure      404   {object}  map[string]interface{}
// @Router       /books/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

// DeleteBook godoc
// @Summary      Delete a book
// @Description  Delete a book by ID
// @Tags         books
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]interface{}
// @Router       /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	config.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
