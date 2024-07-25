package pleo_api

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

func (client *Client) AuthURL(redirectURL, state string) (url string, verifier string) {
	cfg := &oauth2.Config{
		ClientID:     client.config.ClientID,
		ClientSecret: client.config.ClientSecret,
		Endpoint:     client.config.Endpoint,
		RedirectURL:  redirectURL,
		Scopes:       client.config.Scopes,
	}

	verifier = oauth2.GenerateVerifier()
	url = cfg.AuthCodeURL(state, oauth2.S256ChallengeOption(verifier))
	return url, verifier
}

func (client *Client) AuthRedirectHandler(ctx context.Context, redirectURL, verifier string, r *http.Request) (*oauth2.Token, string, error) {
	code := r.FormValue("code")
	state := r.FormValue("state")

	cfg := &oauth2.Config{
		ClientID:     client.config.ClientID,
		ClientSecret: client.config.ClientSecret,
		Endpoint:     client.config.Endpoint,
		RedirectURL:  redirectURL,
		Scopes:       client.config.Scopes,
	}

	exchange, err := cfg.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	return exchange, state, err
}
