package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goagile/gotools/gin/simple/db"
)

func main() {
	r := gin.Default()
	r.POST("/books", createBook) // CREATE	(C)
	// r.PUT("/books/:id", updateBook) // UPDATE	(U)
	r.GET("/books", findBooks) // LIST		(L)
	r.Run(":8080")
}

// createBook - create new book
// {
// 		"title": "New Russian Book",
// 		"autor": "Ivan Petrov"
// }
func createBook(c *gin.Context) {
	var book *db.Book
	c.BindJSON(&book)
	db.Save(book)
	c.JSON(http.StatusCreated, gin.H{
		"data": book,
	})
}

// findBooks - return all books
func findBooks(c *gin.Context) {
	books := db.FindAll()
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

// updateBook - update existing book
// func updateBook(c *gin.Context) {
// 	idstr := c.Query("id")
// 	id, _ := strconv.ParseInt(idstr, 10, 64)
// 	var book *db.Book
// 	c.BindJSON(&book)
// }
