package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	indexApi "github.com/blevesearch/bleve_index_api"
	"github.com/matheuschimelli/relaxasearch/utils"
)

type DocCountResult struct {
	Status string `json:"status"`
	Count  uint64 `json:"count"`
}

func DocCount(indexName string) (DocCountResult, error) {

	if indexName == "" {
		message := "error 11: indexName cannot be blank"
		return DocCountResult{Status: "error", Count: 0}, errors.New(message)
	}

	indexName = utils.Normalize(indexName)

	index := IndexByName(indexName)

	if index == nil {
		message := fmt.Sprintf("error 12: cannot find index %s on indexName param on ShowDoc on document_core.", indexName)
		return DocCountResult{Status: "error", Count: 0}, errors.New(message)
	}

	docCount, err := index.DocCount()

	if err != nil {
		if index == nil {
			message := fmt.Sprintf("error 13: error counting docs on index %s. aditional error: %s", indexName, err)
			return DocCountResult{Status: "error", Count: 0}, errors.New(message)
		}
	}
	result := DocCountResult{
		Status: "ok",
		Count:  docCount,
	}
	return result, nil
}

func UpsertDoc(indexName string, docId string, docData []byte) (OperationResult, error) {
	if indexName == "" {
		message := "error 14: indexName cannot be blank on upsertDoc"
		return OperationResult{Success: false, Message: message}, errors.New(message)
	}

	indexName = utils.Normalize(indexName)

	index := IndexByName(indexName)
	if index == nil {
		message := fmt.Sprintf("error 15: cannot find index %s on indexName param on UpsertDoc on document_core.", indexName)
		return OperationResult{Success: false, Message: message}, errors.New(message)
	}

	if docId == "" {
		message := "error 16: document id cannot be blank"
		return OperationResult{Success: false, Message: message}, errors.New(message)
	}

	var doc interface{}
	err := json.Unmarshal(docData, &doc)
	if err != nil {
		message := "error 17: error parsing json document data"
		return OperationResult{Success: false, Message: message}, errors.New(message)
	}

	err = index.Index(docId, doc)
	if err != nil {
		message := "error 18: error indexing document"
		return OperationResult{Success: false, Message: message}, errors.New(message)
	}

	lastUpsert, err := ShowDoc(indexName, docId)

	if err != nil {
		message := "error 18: error indexing document"
		return OperationResult{Success: false, Message: message}, errors.New(message)
	}
	return OperationResult{
		Success: true,
		Message: "Document indexed",
		Data:    lastUpsert,
	}, nil
}

func ShowDoc(indexName string, docId string) (OperationResult, error) {
	if indexName == "" {
		message := "error 14: indexName cannot be blank on upsertDoc"
		return OperationResult{Success: false, Message: message}, errors.New(message)
	}

	indexName = utils.Normalize(indexName)

	index := IndexByName(indexName)

	if index == nil {
		message := fmt.Sprintf("error 15: cannot find index %s on indexName param on UpsertDoc on document_core.", indexName)
		return OperationResult{Success: false, Message: message}, errors.New(message)
	}

	if docId == "" {
		message := "error 16: document id cannot be blank"
		return OperationResult{Success: false, Message: message}, errors.New(message)
	}

	doc, err := index.Document(docId)

	if err != nil {
		message := "error 16: document id cannot be blank"
		return OperationResult{Success: false, Message: message}, errors.New(message)
	}

	if doc == nil {
		message := "error 16: document id cannot be blank"
		return OperationResult{Success: false, Message: message}, errors.New(message)
	}

	rv := struct {
		ID     string                 `json:"id"`
		Fields map[string]interface{} `json:"fields"`
	}{
		ID:     docId,
		Fields: map[string]interface{}{},
	}

	doc.VisitFields(func(field indexApi.Field) {
		var newval interface{}
		switch field := field.(type) {
		case indexApi.TextField:
			newval = field.Text()
		case indexApi.NumericField:
			n, err := field.Number()
			if err == nil {
				newval = n
			}
		case indexApi.DateTimeField:
			d, err := field.DateTime()
			if err == nil {
				newval = d.Format(time.RFC3339Nano)
			}
		}
		existing, existed := rv.Fields[field.Name()]
		if existed {
			switch existing := existing.(type) {
			case []interface{}:
				rv.Fields[field.Name()] = append(existing, newval)
			case interface{}:
				arr := make([]interface{}, 2)
				arr[0] = existing
				arr[1] = newval
				rv.Fields[field.Name()] = arr
			}
		} else {
			rv.Fields[field.Name()] = newval
		}
	})

	return OperationResult{
		Success: true,
		Data:    rv,
	}, nil
}
