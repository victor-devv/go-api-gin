package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"errors"
)

type book struct {
	ID		string	`json:"id"`
	Title 	string	`json:"title"`
	Author 	string	`json:"author"`
	Quantity int	`json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Book 1", Author: "John Doe", Quantity: 2},
	{ID: "2", Title: "Book 2", Author: "Jane Doe", Quantity: 5},
	{ID: "3", Title: "Book 3", Author: "Cyril Autumn", Quantity: 6},
}

//*gin.Context is all of the info about a request and allows you to return a response
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}