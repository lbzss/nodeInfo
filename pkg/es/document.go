package es

import (
	"context"
	"encoding/json"
	"fmt"
	esapi7 "github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/google/uuid"
	"strings"
)

func (es *Client) CreateDocument(indexName string, data interface{}) error {
	dataString, _ := json.Marshal(data)
	req := esapi7.CreateRequest{
		Index:      indexName,
		Body:       strings.NewReader(string(dataString)),
		DocumentID: uuid.New().String(),
	}
	resp, err := req.Do(context.Background(), es.EsClient)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", resp)
	return nil
}

func (es *Client) DeleteDocument(indexName string) error {
	req := esapi7.DeleteRequest{
		Index: indexName,
	}
	resp, err := req.Do(context.Background(), es.EsClient)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", resp)
	return nil
}

func (es *Client) GetDocument(indexName string) error {
	req := esapi7.SearchRequest{
		Index: []string{indexName},
	}
	resp, err := req.Do(context.Background(), es.EsClient)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)
	return nil
}

func (es *Client) UpdateDocument(indexName string, documentId string, data interface{}) error {
	dataString, _ := json.Marshal(data)
	req := esapi7.UpdateRequest{
		Index:      indexName,
		DocumentID: documentId,
		Body:       strings.NewReader(string(dataString)),
	}
	resp, err := req.Do(context.Background(), es.EsClient)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", resp)
	return nil
}
