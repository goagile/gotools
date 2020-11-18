package controller

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/goagile/gotools/gin/simple/book"
	"github.com/goagile/gotools/gin/simple/db"
)

// CreateBook - create new book
func CreateBook(c *gin.Context) {
	var r CreateBookRequest
	if err := c.BindJSON(&r); err != nil {
		log.Println("CreateBook BindJSON", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "fail to create book",
		})
		return
	}
	b := newBook(r)
	db.Save(b)
	c.JSON(http.StatusCreated, gin.H{"data": b})
}

// CreateBookRequest - request to create book
type CreateBookRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// newBook - return new Book
func newBook(r CreateBookRequest) *book.Book {
	b := new(book.Book)
	b.ID = db.NextID()
	b.Author = r.Author
	b.Title = r.Title
	return b
}

// FindBook one book by ID
func FindBook(c *gin.Context) {
	id, err := book.IDFromString(c.Param("id"))
	if err != nil {
		log.Println("FindBook IDFromString", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("not found book by id %q", id),
		})
		return
	}
	b := db.Find(id)
	c.JSON(http.StatusOK, gin.H{"data": b})
}

// UpdateBook - update existing book
func UpdateBook(c *gin.Context) {
	var r UpdateBookRequest
	c.BindJSON(&r)
	id, err := book.IDFromString(c.Param("id"))
	if err != nil {
		log.Println("UpdateBook IDFromString", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("not found book by id %q to update", id),
		})
		return
	}
	b := db.Find(id)
	b = merge(r, b)
	db.Save(b)
	c.JSON(http.StatusOK, gin.H{"data": b})
}

// merge book from request and found book
func merge(r UpdateBookRequest, b *book.Book) *book.Book {
	if "" != r.Author {
		b.Author = r.Author
	}
	if "" != r.Title {
		b.Title = r.Title
	}
	return b
}

// UpdateBookRequest - request to update book
type UpdateBookRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// FindBooks - return all books
func FindBooks(c *gin.Context) {
	bs := db.FindAll()
	sort.Sort(book.ByID{bs})
	c.JSON(http.StatusOK, gin.H{"data": bs})
}
