// Package accountingentries provides Pleo accounting entry and receipt operations.
// See https://developers.pleo.io/reference/accounting-entries/accounting-entry-overview.
package accountingentries

import "github.com/dieagenturverwaltung/pleo-api/shared"

const basePath = "/v1/accounting-entries"

type Client struct {
	config *shared.Config
}

func New(config *shared.Config) *Client {
	return &Client{config: config}
}
