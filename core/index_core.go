package core

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/matheuschimelli/relaxasearch/utils"
)

//
//Code errors list
// 01 - CreateIndex Fail - indexName param is blank or empty
// 02 - CreateIndex Fail - Bleve error. Check return error
// 03 - DeleteIndex Fail - indexName param is blank or empty
// 04 - DeleteIndex Fail - no such index was found with provided index name. Nothing changed.
// 05 - DeleteIndex Fail - cannot close index.
// 06 - ShowIndex   Fail - cannot find a index with provided indexName param.
//

type IndexStats struct {
	Name     string `json:"name"`
	DocCount uint64 `json:"docCount"`
	Success  bool   `json:"success"`
	//Stats    interface{}
}

type OperationResult struct {
	Success bool
	Message string
	Data    interface{}
}

type IndexDetails struct {
	Success bool                 `json:"success"`
	Status  string               `json:"status"`
	Name    string               `json:"name"`
	Mapping mapping.IndexMapping `json:"mapping"`
}

type Indexes struct {
	Status  string   `json:"status"`
	Indexes []string `json:"indexes"`
}

var indexNameMapping map[string]bleve.Index
var indexNameMappingLock sync.RWMutex

type IndexService struct {
}

func NewIndexService() *IndexService {
	return &IndexService{}
}

func (i *IndexService) ListIndexes() Indexes {
	indexNames := IndexNames()
	indexesList := Indexes{
		Status:  "ok",
		Indexes: indexNames,
	}
	return indexesList
}

func (i *IndexService) InitIndex(dataDir string) {
	dirEntries, err := ioutil.ReadDir(dataDir)
	if err != nil {
		log.Fatalf("error reading data dir: %v", err)
	}

	for _, dirInfo := range dirEntries {
		indexPath := dataDir + string(os.PathSeparator) + dirInfo.Name()

		// skip single files in data dir since a valid index is a directory that
		// contains multiple files
		if !dirInfo.IsDir() {
			log.Printf("not registering %s, skipping", indexPath)
			continue
		}

		i, err := bleve.Open(indexPath)
		if err != nil {
			log.Printf("error opening index %s: %v", indexPath, err)
		} else {
			log.Printf("registered index: %s", dirInfo.Name())
			RegisterIndexName(dirInfo.Name(), i)
			// set correct name in stats
			i.SetName(dirInfo.Name())
		}
	}
}

func (i *IndexService) CreateIndex(path string, indexName string) (bleve.Index, error) {
	if indexName == "" {
		return nil, errors.New("error 01: index name cannot be blank")
	}

	//
	// indexName will be normalized, which means that it will removes
	// blank spaces, special caracters.
	//

	indexName = utils.Normalize(indexName)

	indexMapping := bleve.NewIndexMapping()
	newIndex, err := bleve.New(i.indexPath(indexName), indexMapping)

	if err != nil {
		message := fmt.Sprintf("error 02: Index cannot be created: %s", err)
		return nil, errors.New(message)
	}

	newIndex.SetName(indexName)
	RegisterIndexName(indexName, newIndex)

	return newIndex, nil
}

func (i *IndexService) DeleteIndex(path string, indexName string) (OperationResult, error) {
	if indexName == "" {
		return OperationResult{Success: false, Message: "IndexName param cannot be blank"},
			errors.New("error 03: index name cannot be blank")
	}

	indexToDelete := UnregisterIndexByName(indexName)

	if indexToDelete == nil {
		message := "error 04: no such index was found with provided index name. Nothing changed"
		return OperationResult{Success: false, Message: message}, errors.New(message)
	}

	err := indexToDelete.Close()
	if err != nil {
		message := fmt.Sprintf("error 05: cannot close index. Error: %s", err)
		return OperationResult{Success: false, Message: message}, errors.New(message)
	}

	err = os.RemoveAll(i.indexPath(indexName))
	if err != nil {
		message := fmt.Sprintf("error 06: cannot delete index. Error: %s", err)
		return OperationResult{Success: false, Message: message}, errors.New(message)
	}
	return OperationResult{
		Success: true,
		Message: "Index deleted",
	}, nil
}

func (i *IndexService) ShowIndex(indexName string) (interface{}, error) {
	if indexName == "" {
		return OperationResult{Success: false, Message: "IndexName param cannot be blank"},
			errors.New("error 03: index name cannot be blank")
	}

	index := IndexByName(indexName)
	if index == nil {
		message := "error 06: index not found"
		return OperationResult{Success: false, Message: message}, errors.New(message)
	}

	indexDetails := IndexDetails{
		Success: true,
		Status:  "ok",
		Name:    indexName,
		Mapping: index.Mapping(),
	}

	return indexDetails, nil
}

//
// utils
//

func RegisterIndexName(name string, idx bleve.Index) {
	indexNameMappingLock.Lock()
	defer indexNameMappingLock.Unlock()

	if indexNameMapping == nil {
		indexNameMapping = make(map[string]bleve.Index)
	}
	indexNameMapping[name] = idx
}

func UnregisterIndexByName(name string) bleve.Index {
	indexNameMappingLock.Lock()
	defer indexNameMappingLock.Unlock()

	if indexNameMapping == nil {
		return nil
	}
	rv := indexNameMapping[name]
	if rv != nil {
		delete(indexNameMapping, name)
	}
	return rv
}

func (i *IndexService) indexPath(name string) string {
	return "relaxasearchData" + string(os.PathSeparator) + name
}

func IndexByName(name string) bleve.Index {
	indexNameMappingLock.RLock()
	defer indexNameMappingLock.RUnlock()

	return indexNameMapping[name]
}

func IndexNames() []string {
	indexNameMappingLock.RLock()
	defer indexNameMappingLock.RUnlock()

	rv := make([]string, len(indexNameMapping))
	count := 0
	for k := range indexNameMapping {
		rv[count] = k
		count++
	}
	return rv
}
