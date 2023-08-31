package handlers

import (
	"net/http"

	"goswagger/models"

	"github.com/gin-gonic/gin"
)

// GetBooks responds with the list of all books as JSON.
// GetBooks             godoc
// @Summary      Get books array
// @Description  Responds with the list of all books as JSON.
// @Tags         books
// @Produce      json
// @Success      200  {array}  models.Book
// @Router       /books [get]
func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, []models.Book{})
}

// PostBook takes a book JSON and store in DB.
// PostBook             godoc
// @Summary      Store a new book
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         books
// @Produce      json
// @Param        book  body      models.Book  true  "Book JSON"
// @Success      200   {object}  models.Book
// @Router       /books [post]
func PostBook(c *gin.Context) {
	var newBook models.Book

	c.JSON(http.StatusCreated, newBook)
}

// GetBookByISBN locates the book whose ISBN value matches the isbn
// GetBookByISBN                godoc
// @Summary      Get single book by isbn
// @Description  Returns the book whose ISBN value matches the isbn.
// @Tags         books
// @Produce      json
// @Param        isbn  path      string  true  "search book by isbn"
// @Success      200  {object}  models.Book
// @Router       /books/{isbn} [get]
func GetBookByISBN(c *gin.Context) {
	// isbn := c.Param("isbn")

	c.JSON(http.StatusOK, models.Book{})
}
