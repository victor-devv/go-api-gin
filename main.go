package main

import (
	"errors"
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

//*gin.Context has all of the info (query params, payload, etc) about a request and allows you to return a response
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookByID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("Book not found")
}

func createBook(c *gin.Context) {
	//get data of book to be created
	var newBook book

	//Bind json from payload to book struct
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing query parameter"})
		return
	}
	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available!"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.Run("localhost:8080")
}