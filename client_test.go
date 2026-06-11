package pleo_api

import (
	"context"
	"embed"
	"encoding/json"
	"log"
	"os"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

//go:embed .token.json
var testingFS embed.FS
var tokenConfig tokenConfigData

type tokenConfigData struct {
	Token        *oauth2.Token `json:"token"`
	ClientID     string        `json:"client_id"`
	ClientSecret string        `json:"client_secret"`
}

func init() {
	bytes, err := testingFS.ReadFile(".token.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &tokenConfig)
	if err != nil {
		panic(err)
	}
}

func onTokenChange(newToken *oauth2.Token, err error) {
	if err != nil {
		panic(err)
	}

	tokenConfig.Token = newToken
	marshal, err := json.Marshal(tokenConfig)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(".token.json", marshal, 0644)
	if err != nil {
		panic(err)
	}
}

func client() (*HttpClient, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	cfg := &HttpConfiguration{
		Token:         tokenConfig.Token,
		OnTokenUpdate: onTokenChange,
		Debug:         true,
		Logger:        log.Printf,
	}

	return New(tokenConfig.ClientID, tokenConfig.ClientSecret, true, AllScopes...).Http(ctx, cfg), cancel
}

func TestTagGroup(t *testing.T) {
	client, cancelClient := client()
	defer cancelClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	t.Run("List", func(t *testing.T) {
		companies, err := client.Companies.Search(ctx).Execute()
		if err != nil {
			t.Fatal(err)
		}

		if len(companies.Data) == 0 {
			t.Fatal("no companies found")
		}

		_, err = client.Tags.TagGroupsApi.GetTagGroups(ctx).WithCompanyID(companies.Data[0].ID).Execute()
		if err != nil {
			t.Fatal(err)
		}
	})
}
