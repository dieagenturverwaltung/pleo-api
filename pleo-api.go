package pleo_api

import "golang.org/x/oauth2"

const (
	tokenURL = "https://auth.pleo.io/oauth/token"
	authURL  = "https://auth.pleo.io/oauth/authorize"

	stagingTokenURL = "https://auth.staging.pleo.io/oauth/token"
	stagingAuthURL  = "https://auth.staging.pleo.io/oauth/authorize"

	stagingBaseURL = "https://openapi.staging.pleo.io/v1"
	baseURL        = "https://openapi.pleo.io/v1"
)

type Client struct {
	config  *oauth2.Config
	baseURL string
}

func New(clientID, clientSecret string, staging bool, scopes ...string) *Client {
	var endpoint oauth2.Endpoint
	baseUrl := baseURL
	if staging {
		baseUrl = stagingBaseURL
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
		config:  conf,
		baseURL: baseUrl,
	}
}
