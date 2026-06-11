package pleo_api

import (
	"github.com/dieagenturverwaltung/pleo-api/shared"
	"golang.org/x/oauth2"
)

const (
	tokenURL = "https://auth.pleo.io/oauth/token"
	authURL  = "https://auth.pleo.io/oauth/authorize"

	stagingTokenURL = "https://auth.staging.pleo.io/oauth/token"
	stagingAuthURL  = "https://auth.staging.pleo.io/oauth/authorize"
)

type Client struct {
	config               *oauth2.Config
	openApiConfiguration shared.Config
}

func New(clientID, clientSecret string, staging bool, scopes ...string) *Client {
	var endpoint oauth2.Endpoint
	var openApiConfiguration shared.Config
	if staging {
		openApiConfiguration = *shared.NewStagingConfig()
		endpoint = oauth2.Endpoint{
			AuthURL:  stagingAuthURL,
			TokenURL: stagingTokenURL,
		}
	} else {
		openApiConfiguration = *shared.NewConfig()
		endpoint = oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		}

	}
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint:     endpoint,
	}

	return &Client{
		config:               conf,
		openApiConfiguration: openApiConfiguration,
	}
}
