package main

import (
	// "fmt"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Titile   string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Titile: "In search", Author: "David", Quantity: 2},
	{ID: "2", Titile: "At Home", Author: "Robert", Quantity: 3},
	{ID: "3", Titile: "In Company", Author: "Wiliam", Quantity: 5},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")

}

func createBooks(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBooks)
	router.Run("localhost:8080")
}
