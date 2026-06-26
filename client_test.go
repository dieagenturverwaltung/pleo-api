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
	CompanyID    string        `json:"company_id"`
	ClientID     string        `json:"client_id"`
	ClientSecret string        `json:"client_secret"`
	Staging      bool          `json:"staging"`
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

	var scopes []string
	if tokenConfig.Staging {
		scopes = AllScopes
	} else {
		scopes = AllScopesProd
	}

	return New(tokenConfig.ClientID, tokenConfig.ClientSecret, tokenConfig.Staging, scopes...).Http(ctx, cfg), cancel
}

func TestActivateInstallation(t *testing.T) {
	client, cancelClient := client()
	defer cancelClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	t.Run("ActivateInstallation", func(t *testing.T) {
		execute, err := client.Marketplace.ActivateMyInstallation(ctx).Execute()
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("Installation activated: %v", execute)
	})
}

func TestTagGroup(t *testing.T) {
	client, cancelClient := client()
	defer cancelClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	t.Run("List", func(t *testing.T) {
		_, err := client.Tags.TagGroupsApi.GetTagGroups(ctx).WithCompanyID(tokenConfig.CompanyID).Execute()
		if err != nil {
			t.Fatal(err)
		}
	})
}
