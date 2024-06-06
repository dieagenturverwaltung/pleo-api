package pleo_api

import "golang.org/x/oauth2"

const (
	tokenURL = "https://auth.pleo.io/oauth/token"
	authURL  = "https://auth.pleo.io/oauth/authorize"

	stagingTokenURL = "https://auth.staging.pleo.io/oauth/token"
	stagingAuthURL  = "https://auth.staging.pleo.io/oauth/authorize"
)

type Client struct {
	config *oauth2.Config
}

func New(clientID, clientSecret string, staging bool, scopes ...string) *Client {
	var endpoint oauth2.Endpoint
	if staging {
		endpoint = oauth2.Endpoint{
			AuthURL:  stagingAuthURL,
			TokenURL: stagingTokenURL,
		}
	} else {
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
		config: conf,
	}
}
