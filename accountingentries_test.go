package pleo_api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/oauth2"
)

type accountingEntriesTransport struct {
	t *testing.T
}

func (transport accountingEntriesTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if r.Method != http.MethodGet || r.URL.Path != "/v1/accounting-entries/entry-id" {
		transport.t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
	}
	if got := r.Header.Get("Authorization"); got != "Bearer test-token" {
		transport.t.Errorf("authorization = %q, want Bearer test-token", got)
	}
	w := httptest.NewRecorder()
	_, _ = w.Write([]byte(`{"data":{"id":"entry-id","family":"CARD_PURCHASE"}}`))
	return w.Result(), nil
}

func TestHttpAccountingEntries(t *testing.T) {
	api := New("client-id", "client-secret", true, ScopeAccountingEntriesRead)
	api.openApiConfiguration.BaseURL = "https://example.test"
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{Transport: accountingEntriesTransport{t: t}})
	client := api.Http(ctx, &HttpConfiguration{Token: &oauth2.Token{AccessToken: "test-token"}})
	if client.AccountingEntries == nil {
		t.Fatal("accounting entries client is not initialized")
	}
	entry, _, err := client.AccountingEntries.GetAccountingEntry(ctx, "entry-id").Execute()
	if err != nil {
		t.Fatal(err)
	}
	if entry.ID != "entry-id" {
		t.Fatalf("unexpected entry: %+v", entry)
	}
}
