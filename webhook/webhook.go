package webhook

import (
	"github.com/dieagenturverwaltung/pleo-api/shared"
)

const basePath = "/v1/subscriptions"

type Client struct {
	config *shared.Config
}

func New(config *shared.Config) *Client {
	return &Client{config: config}
}
