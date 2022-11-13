package bleve

import (
	"errors"
	"fmt"

	"github.com/blevesearch/bleve/v2"
	index "github.com/blevesearch/bleve_index_api"
)

type IndexStats struct {
	Name     string `json:"name"`
	DocCount uint64 `json:"docCount"`
	Success  bool   `json:"success"`
	//Stats    interface{}
}

func NewIndex(name string) (bool, error) {
	if name == "" {
		return false, errors.New("Database fileName param cannot be empty")
	}

	indexName := fmt.Sprintf("./%s.bleve", name)
	mapping := bleve.NewIndexMapping()

	_, err := bleve.New(indexName, mapping)
	if err != nil {
		fmt.Println(err)
		return false, errors.New("Database fileName param cannot be empty")
	}
	return true, nil
}

func Index(indexName string, id string, data interface{}) (index.Document, error) {
	index, err := bleve.Open(indexName)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("IndexName doest not exists")
	}

	// index some data
	index.Index(id, data)

	lastInsert, err := index.Document(id)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Last doc cold not be inserted")
	}

	return lastInsert, nil
}

func GetIndex(indexName string) (IndexStats, error) {
	index, err := bleve.Open(indexName)

	if err != nil {
		fmt.Println(err)
		return IndexStats{Success: false}, errors.New("IndexName doest not exists")
	}

	docCount, err := index.DocCount()
	if err != nil {
		fmt.Println(err)
		return IndexStats{Success: false}, errors.New("IndexName doest not exists")
	}

	data := IndexStats{
		Name:     index.Name(),
		DocCount: docCount,
		Success:  true,
	}

	return data, nil
}

func Search(indexName string, queryString string) (*bleve.SearchResult, error) {

	index, err := bleve.Open(indexName)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("IndexName doest not exists")
	}

	// search for some text
	query := bleve.NewMatchQuery(queryString)
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)

	return searchResults, nil
}
