package routes

import (
	"awesomeProject/models"
	"awesomeProject/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Настройка маршрутов для работы с книгами
func SetupBookRoutes(r *gin.Engine, bookService *services.BookService) {
	r.POST("/books", func(c *gin.Context) {
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		if err := bookService.CreateBook(&book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
			return
		}
		c.JSON(http.StatusCreated, book)
	})

	r.GET("/books", func(c *gin.Context) {
		books, err := bookService.GetAllBooks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
			return
		}
		c.JSON(http.StatusOK, books)
	})

	r.GET("/books/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		book, err := bookService.GetBookByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusOK, book)
	})

	r.PUT("/books/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		if err := bookService.UpdateBook(uint(id), &book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
			return
		}
		c.JSON(http.StatusOK, book)
	})

	r.DELETE("/books/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := bookService.DeleteBook(uint(id)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
	})
}
