package pleo_api

import (
	"context"
	"sync"

	"github.com/dieagenturverwaltung/pleo-api/api"
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

func (client *Client) Http(ctx context.Context, token *oauth2.Token, updateTokenFunc func(updatedToken *oauth2.Token)) *api.Api {
	source := client.config.TokenSource(ctx, token)

	newClient := oauth2.NewClient(ctx, &tokenSourceWrapper{
		currentToken: token,
		source:       source,
		onUpdate:     updateTokenFunc,
	})
	return api.New(newClient)
}
