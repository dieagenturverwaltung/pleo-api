package taxcodes

import "github.com/dieagenturverwaltung/pleo-api/shared"

const basePath = "/v0/tax-codes"

type Client struct {
	config *shared.Config
}

func New(config *shared.Config) *Client {
	return &Client{config: config}
}
