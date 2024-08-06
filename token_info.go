package pleo_api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

// TokenInfo
//
//	{
//		"active": true,
//		"sub": "RESOURCE_ID",
//		"token_type": "at+JWT",
//		"exp": ACCESS_TOKEN_LIFETIME,
//		"client_id": "CLIENT_ID",
//		"iat": TIME_TOKEN_ISSUED,
//		"aud": "https://external.pleo.io",
//		"iss": "pleo.production",
//		"jti": "TOKEN_ID"
//	}
type TokenInfo struct {
	Active    bool   `json:"active"`
	Sub       string `json:"sub"`
	TokenType string `json:"token_type"`
	Esp       int    `json:"exp"`
	ClientId  string `json:"client_id"`
	Iat       int    `json:"iat"`
	Aud       string `json:"aud"`
	Iss       string `json:"iss"`
	Jti       string `json:"jti"`
}

func (client *Client) TokenInfo(ctx context.Context, token *oauth2.Token) (*TokenInfo, error) {
	introspectUrl := client.config.Endpoint.TokenURL + "/introspect"
	var form = make(url.Values)
	form.Set("token", token.AccessToken)

	request, err := http.NewRequest("POST", introspectUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth(url.QueryEscape(client.config.ClientID), url.QueryEscape(client.config.ClientSecret))
	response, err := http.DefaultClient.Do(request.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var info TokenInfo
	err = json.NewDecoder(response.Body).Decode(&info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
