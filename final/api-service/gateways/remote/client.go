package remote

import (
	"exam-api/domain"
	"net/http"
)

// Implement the following client to connect to the remote storage server

type Client struct {
	client http.Client
}

func NewClient(client http.Client) *Client {
	return &Client{client: client}
}

func (c *Client) Save(product domain.Product) (string, bool, error) {
	panic("Implement me")
}

func (c *Client) Get(id string) (domain.Product, bool, error) {
	panic("Implement me")

}

func (c *Client) Update(id string, diff domain.Product) (bool, error) {
	panic("Implement me")

}

func (c *Client) Delete(id string) (bool, error) {
	panic("Implement me")

}
