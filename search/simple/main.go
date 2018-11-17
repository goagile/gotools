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
			ID:   "123",
			Text: "Hello world",
		},
		&Post{
			ID:   "3435",
			Text: "Ola-la",
		},
	}
	for _, p := range posts {
		index.Index(p.ID, p)
	}

	search := bleve.NewSearchRequest(bleve.NewMatchQuery("la"))
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults.Hits[0].ID)
}

type Post struct {
	ID   string
	Text string
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
