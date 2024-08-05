package pleo_api

import (
	"context"
	"strings"
	"sync"

	"github.com/dieagenturverwaltung/pleo-api/export"
	"github.com/dieagenturverwaltung/pleo-api/marketplace"
	"github.com/dieagenturverwaltung/pleo-api/tags"
	"github.com/dieagenturverwaltung/pleo-api/tax_codes"
	"golang.org/x/oauth2"
)

type pleoError string

func (e pleoError) Error() string {
	return string(e)
}

const ErrRefreshTokenExpired = pleoError("refresh token expired")

type tokenSourceWrapper struct {
	mu           sync.Mutex
	currentToken *oauth2.Token
	source       oauth2.TokenSource
	onUpdate     func(updatedToken *oauth2.Token, refreshError error)
}

func (w *tokenSourceWrapper) Token() (*oauth2.Token, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	token, err := w.source.Token()
	if err != nil {
		if w.onUpdate != nil {
			if strings.Contains(err.Error(), "invalid_refresh_token") {
				go w.onUpdate(nil, ErrRefreshTokenExpired)
			} else {
				go w.onUpdate(nil, err)
			}
		}
		return nil, err
	}

	if w.currentToken.AccessToken != token.AccessToken {
		w.currentToken = token
		if w.onUpdate != nil {
			go w.onUpdate(token, nil)
		}
	}

	return token, nil
}

type HttpClient struct {
	Export      *export.APIClient
	Tags        *tags.APIClient
	TaxCodes    *tax_codes.APIClient
	Marketplace *marketplace.APIClient
}

type HttpConfiguration struct {
	Token         *oauth2.Token
	CompanyID     *string
	OnTokenUpdate func(token *oauth2.Token, refreshError error)
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
	config.CompanyID = cfg.CompanyID

	return &HttpClient{
		Export:      export.NewAPIClient(&config),
		Tags:        tags.NewAPIClient(&config),
		TaxCodes:    tax_codes.NewAPIClient(&config),
		Marketplace: marketplace.NewAPIClient(&config),
	}
}
