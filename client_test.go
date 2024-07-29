package pleo_api

import (
	"context"
	"embed"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/dieagenturverwaltung/pleo-api/swagger"
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

func onTokenChange(newToken *oauth2.Token) {
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

func client() (*swagger.APIClient, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)

	return New(tokenConfig.ClientID, tokenConfig.ClientSecret, true, AllScopes...).Http(ctx, tokenConfig.Token, onTokenChange), cancel
}

func TestTagGroup(t *testing.T) {
	client, cancelClient := client()
	defer cancelClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	t.Run("List", func(t *testing.T) {
		_, _, err := client.TagGroupsApi.GetTagGroups(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}
	})
}
