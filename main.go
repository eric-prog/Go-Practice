package main

import (
	"errors"
	"net/http" // http.StatusOK (200)

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"` // uppercase because we need them to be viewed publically and be able to read. Serialized and convert to lowercase
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Eric Sheen", Quantity: 12},
	{ID: "2", Title: "In Search of Lost Time", Author: "Eric Sheen", Quantity: 12},
}

func getBooks(c *gin.Context) { // gin context is all the information about request and allows to return response
	c.IndentedJSON(http.StatusOK, books) // nicely formatted indented json. Sending books data. Books struct get serialized we return JSON objects
}

func createBook(c *gin.Context) { // stores all data in c and get query parameters from c
	var newBook book // contains all data

	if err := c.BindJSON(&newBook); err != nil { // binding request data to newBook
		return // return if error
	} // pointer to newBook and check if error

	books = append(books, newBook)              // add to array
	c.IndentedJSON(http.StatusCreated, newBook) // return book created and status msg
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func bookById(c *gin.Context) {
	id := c.Param("id") // "/books/2"
	book, error := getBookById(id)

	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"}) // H allows us to write quick custom json
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()           // router comes from gin and allows us to handle routes
	router.GET("/books", getBooks)    // /books when u got to localhost its calls getBooks
	router.POST("/books", createBook) // /books when u got to localhost its calls getBooks
	router.GET("/books/:id", bookById)
	router.Run("localhost:8080")
}
