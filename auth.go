package pleo_api

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

func (client *Client) AuthURL(redirectURL, state string) string {
	cfg := &oauth2.Config{
		ClientID:     client.config.ClientID,
		ClientSecret: client.config.ClientSecret,
		Endpoint:     client.config.Endpoint,
		RedirectURL:  redirectURL,
		Scopes:       client.config.Scopes,
	}

	return cfg.AuthCodeURL(state)
}

func (client *Client) AuthRedirectHandler(ctx context.Context, r *http.Request) (*oauth2.Token, string, error) {
	code := r.FormValue("code")
	state := r.FormValue("state")

	exchange, err := client.config.Exchange(ctx, code)
	return exchange, state, err
}
