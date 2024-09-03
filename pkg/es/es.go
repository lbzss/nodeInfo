package es

import (
	es7 "github.com/elastic/go-elasticsearch/v7"
)

type Client struct {
	EsClient *es7.Client
}

func NewEsClient(address []string, username, password string) (*Client, error) {
	client, err := es7.NewClient(es7.Config{
		Addresses: address,
		Username:  username,
		Password:  password,
	})
	if err != nil {
		return nil, err
	}

	return &Client{EsClient: client}, nil
}
