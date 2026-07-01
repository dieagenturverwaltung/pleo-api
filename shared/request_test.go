package shared

import (
	"context"
	"errors"
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

func TestSendRequestReturnsRateLimitExtra(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("ratelimit-limit", "600")
		w.Header().Set("ratelimit-remaining", "599")
		w.Header().Set("ratelimit-reset", "42")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	config := &Config{
		BaseURL:    server.URL,
		HttpClient: server.Client(),
	}

	_, response, err := config.SendRequest(context.Background(), http.MethodGet, "/resource", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	extra := ResponseExtraFromResponse(response)
	if extra == nil || extra.RateLimit == nil {
		t.Fatal("expected rate limit extra")
	}
	if got := intValue(extra.RateLimit.Limit); got != 600 {
		t.Fatalf("expected limit 600, got %d", got)
	}
	if got := intValue(extra.RateLimit.Remaining); got != 599 {
		t.Fatalf("expected remaining 599, got %d", got)
	}
	if got := intValue(extra.RateLimit.ResetSeconds); got != 42 {
		t.Fatalf("expected reset 42, got %d", got)
	}
}

func TestSendRequestIncludesRateLimitExtraOnTooManyRequests(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("ratelimit-limit", "600")
		w.Header().Set("ratelimit-remaining", "0")
		w.Header().Set("ratelimit-reset", "12")
		w.Header().Set("Retry-After", "12")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"message":"too many requests"}`))
	}))
	defer server.Close()

	config := &Config{
		BaseURL:    server.URL,
		HttpClient: server.Client(),
	}

	_, response, err := config.SendRequest(context.Background(), http.MethodGet, "/resource", nil, nil)
	if err == nil {
		t.Fatal("expected error")
	}

	extra := ResponseExtraFromResponse(response)
	if extra == nil || extra.RateLimit == nil {
		t.Fatal("expected rate limit extra")
	}
	if got := intValue(extra.RateLimit.Remaining); got != 0 {
		t.Fatalf("expected remaining 0, got %d", got)
	}
	if got := intValue(extra.RateLimit.RetryAfterSeconds); got != 12 {
		t.Fatalf("expected retry after 12, got %d", got)
	}

	var requestErr *RequestError
	if !errors.As(err, &requestErr) {
		t.Fatalf("expected RequestError, got %T", err)
	}
	if requestErr.Extra == nil || requestErr.Extra.RateLimit == nil {
		t.Fatal("expected request error rate limit extra")
	}
	if got := intValue(requestErr.Extra.RateLimit.ResetSeconds); got != 12 {
		t.Fatalf("expected request error reset 12, got %d", got)
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

func intValue(value *int) int {
	if value == nil {
		return -1
	}

	return *value
}
