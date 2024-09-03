package es

import (
	"context"
	"fmt"
	esapi7 "github.com/elastic/go-elasticsearch/v7/esapi"
)

func (es *Client) CreateIndex(indexName string) error {
	indexGetReq := esapi7.IndicesGetRequest{
		Index: []string{indexName},
	}
	resp, err := indexGetReq.Do(context.Background(), es.EsClient)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)
	if resp.StatusCode == 200 {
		return nil
	}
	req := esapi7.IndicesCreateRequest{
		Index: indexName,
	}

	resp, err = req.Do(context.Background(), es.EsClient)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)
	return nil
}

func (es *Client) DeleteIndex(indexName string) error {
	req := esapi7.IndicesDeleteRequest{
		Index: []string{indexName},
	}
	resp, err := req.Do(context.Background(), es.EsClient)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", resp)
	return nil
}

func (es *Client) GetIndex(indexName string) error {
	req := esapi7.IndicesGetRequest{
		Index: []string{indexName},
	}
	resp, err := req.Do(context.Background(), es.EsClient)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", resp)
	return nil
}
