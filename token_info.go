package pleo_api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

type tokenInfoRequest struct {
	Token string `json:"token"`
}

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

func (client *Client) TokenInfo(ctx context.Context, token *oauth2.Token) (*TokenInfo, string, error) {
	introspectUrl := client.config.Endpoint.TokenURL + "/introspect"
	body := tokenInfoRequest{
		Token: token.AccessToken,
	}

	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, "", err
	}

	request, err := http.NewRequest("POST", introspectUrl, bytes.NewReader(marshal))
	if err != nil {
		return nil, "", err
	}

	request.SetBasicAuth(url.QueryEscape(client.config.ClientID), url.QueryEscape(client.config.ClientSecret))
	response, err := http.DefaultClient.Do(request.WithContext(ctx))
	if err != nil {
		return nil, "", err
	}

	defer response.Body.Close()

	all, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, "", err
	}

	var info TokenInfo
	err = json.Unmarshal(all, &info)
	if err != nil {
		return nil, string(all), err
	}

	return &info, string(all), nil
}
