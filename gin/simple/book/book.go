package book

import "strconv"

// New - book constructor
func New(id int64, title, author string) *Book {
	b := new(Book)
	b.Title = title
	b.Author = author
	return b
}

// Book - model entity
type Book struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// ByID books sorter
type ByID struct {
	Books []*Book
}

func (s ByID) Less(i, j int) bool {
	return (s.Books[i].Title < s.Books[j].Title)
}

func (s ByID) Swap(i, j int) {
	s.Books[i], s.Books[j] = s.Books[j], s.Books[i]
}

func (s ByID) Len() int {
	return len(s.Books)
}

// IDFromString - create book ID from string
func IDFromString(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
