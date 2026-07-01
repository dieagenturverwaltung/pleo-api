package shared

import (
	"net/http"
	"strconv"
)

type ResponseExtra struct {
	RateLimit *RateLimitInfo `json:"rateLimit,omitempty"`
}

type RateLimitInfo struct {
	Limit             *int `json:"limit,omitempty"`
	Remaining         *int `json:"remaining,omitempty"`
	ResetSeconds      *int `json:"resetSeconds,omitempty"`
	RetryAfterSeconds *int `json:"retryAfterSeconds,omitempty"`
}

func ResponseExtraFromResponse(response *http.Response) *ResponseExtra {
	if response == nil {
		return nil
	}

	rateLimit := RateLimitInfo{
		Limit:             headerInt(response.Header, "ratelimit-limit"),
		Remaining:         headerInt(response.Header, "ratelimit-remaining"),
		ResetSeconds:      headerInt(response.Header, "ratelimit-reset"),
		RetryAfterSeconds: headerInt(response.Header, "Retry-After"),
	}
	if rateLimit.Limit == nil &&
		rateLimit.Remaining == nil &&
		rateLimit.ResetSeconds == nil &&
		rateLimit.RetryAfterSeconds == nil {
		return nil
	}

	return &ResponseExtra{RateLimit: &rateLimit}
}

func headerInt(headers http.Header, name string) *int {
	value := headers.Get(name)
	if value == "" {
		return nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return nil
	}

	return &parsed
}
