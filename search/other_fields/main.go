package main

import (
	"fmt"

	"github.com/blevesearch/bleve"
)

func main() {
	index, err := Connect("example.bleve")
	if err != nil {
		fmt.Println("bleve.Open: err: ", err)
		return
	}

	posts := []*Post{
		&Post{
			ID:    "111",
			Title: "Hello AAA",
			Body:  "XXXXXXX",
		},
		&Post{
			ID:    "222",
			Title: "Hello BBB",
			Body:  "YYYYYYY",
		},
	}

	for _, p := range posts {
		index.Index(p.ID, p)
	}

	search := bleve.NewSearchRequest(bleve.NewMatchQuery("AAA"))
	search.Fields = []string{"title"}
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)
	fmt.Println(searchResults.Hits[0].ID)
	fmt.Println(searchResults.Hits[0].Fields)
}

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

//
// Connect
//
func Connect(indexPath string) (bleve.Index, error) {
	index, err := bleve.Open(indexPath)
	if err != nil {
		index, err = bleve.New(indexPath, bleve.NewIndexMapping())
		if err != nil {
			return nil, err
		}
	}
	return index, nil
}
