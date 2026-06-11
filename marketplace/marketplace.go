package marketplace

import "github.com/dieagenturverwaltung/pleo-api/shared"

const basePath = "/v1"

type Client struct {
	config *shared.Config
}

func New(config *shared.Config) *Client {
	return &Client{config: config}
}
