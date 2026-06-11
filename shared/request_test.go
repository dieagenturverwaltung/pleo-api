package shared

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/oauth2"
)

type failingTokenSource struct {
	err error
}

func (s failingTokenSource) Token() (*oauth2.Token, error) {
	return nil, s.err
}

func TestSendRequestLogsFailedStatusResponseBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte(`{"message":"invalid request"}`))
	}))
	defer server.Close()

	var logs []string
	config := &Config{
		BaseURL:    server.URL,
		HttpClient: server.Client(),
		Debug:      true,
		Logger: func(format string, args ...any) {
			logs = append(logs, fmt.Sprintf(format, args...))
		},
	}

	_, response, err := config.SendRequest(context.Background(), http.MethodGet, "/resource", nil, nil)
	if err == nil {
		t.Fatal("expected error")
	}
	if response == nil {
		t.Fatal("expected response")
	}
	if response.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("expected status %d, got %d", http.StatusUnprocessableEntity, response.StatusCode)
	}
	if !containsLog(logs, `Response body: GET `+server.URL+`/resource: {"message":"invalid request"}`) {
		t.Fatalf("expected response body log, got %v", logs)
	}
}

func TestSendRequestLogsOAuthRetrieveErrorBody(t *testing.T) {
	retrieveErr := &oauth2.RetrieveError{
		Response: &http.Response{
			Status:     "400 Bad Request",
			StatusCode: http.StatusBadRequest,
		},
		Body: []byte(`{"error":true,"message":"invalid_refresh_token"}`),
	}

	var logs []string
	config := &Config{
		BaseURL: "https://example.test",
		HttpClient: oauth2.NewClient(context.Background(), failingTokenSource{
			err: retrieveErr,
		}),
		Debug: true,
		Logger: func(format string, args ...any) {
			logs = append(logs, fmt.Sprintf(format, args...))
		},
	}

	_, response, err := config.SendRequest(context.Background(), http.MethodGet, "/resource", nil, nil)
	if err == nil {
		t.Fatal("expected error")
	}
	if response != nil {
		t.Fatalf("expected no API response, got %v", response)
	}
	if !containsLog(logs, `Response body: GET https://example.test/resource: {"error":true,"message":"invalid_refresh_token"}`) {
		t.Fatalf("expected oauth response body log, got %v", logs)
	}
}

func containsLog(logs []string, want string) bool {
	for _, log := range logs {
		if strings.Contains(log, want) {
			return true
		}
	}

	return false
}
