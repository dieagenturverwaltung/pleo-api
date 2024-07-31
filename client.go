package pleo_api

import (
	"context"
	"sync"

	"github.com/dieagenturverwaltung/pleo-api/export"
	"github.com/dieagenturverwaltung/pleo-api/tags"
	"github.com/dieagenturverwaltung/pleo-api/tax_codes"
	"golang.org/x/oauth2"
)

type tokenSourceWrapper struct {
	mu           sync.Mutex
	currentToken *oauth2.Token
	source       oauth2.TokenSource
	onUpdate     func(updatedToken *oauth2.Token)
}

func (w *tokenSourceWrapper) Token() (*oauth2.Token, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	token, err := w.source.Token()
	if err != nil {
		return nil, err
	}

	if w.currentToken.AccessToken != token.AccessToken {
		w.currentToken = token
		if w.onUpdate != nil {
			go w.onUpdate(token)
		}
	}

	return token, nil
}

type HttpClient struct {
	Export   *export.APIClient
	Tags     *tags.APIClient
	TaxCodes *tax_codes.APIClient
}

type HttpConfiguration struct {
	Token         *oauth2.Token
	OnTokenUpdate func(token *oauth2.Token)
	Logger        func(string, ...any)
	Debug         bool
}

// Http creates a new API client with the provided token.
func (client *Client) Http(ctx context.Context, cfg *HttpConfiguration) *HttpClient {
	source := client.config.TokenSource(ctx, cfg.Token)

	newClient := oauth2.NewClient(ctx, &tokenSourceWrapper{
		currentToken: cfg.Token,
		source:       source,
		onUpdate:     cfg.OnTokenUpdate,
	})

	config := client.openApiConfiguration
	config.HTTPClient = newClient
	config.Logger = cfg.Logger
	config.Debug = cfg.Debug

	return &HttpClient{
		Export:   export.NewAPIClient(&config),
		Tags:     tags.NewAPIClient(&config),
		TaxCodes: tax_codes.NewAPIClient(&config),
	}
}
