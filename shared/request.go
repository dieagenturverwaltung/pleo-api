package shared

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func (c *Config) SendRequest(ctx context.Context, method string, url string, body any, output any) ([]byte, *http.Response, error) {
	var bodyReady io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, nil, err
		}

		bodyReady = bytes.NewReader(bodyBytes)
	}

	if !strings.HasPrefix(url, "/") && !strings.HasSuffix(c.BaseURL, "/") {
		url = "/" + url
	}

	url = c.BaseURL + url

	request, err := http.NewRequestWithContext(ctx, method, url, bodyReady)
	if err != nil {
		if c.Debug {
			c.Logger("Failed to create request: %v: %s %s", err, method, url)
		}

		return nil, nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		request.Header.Set("User-Agent", c.UserAgent)
	}

	if c.Debug {
		c.Logger("Sending request: %s %s", method, url)
	}

	response, err := c.HttpClient.Do(request)
	if err != nil {
		if c.Debug {
			c.Logger("Failed to send request: %v: %s %s", err, method, url)
		}

		return nil, nil, err
	}

	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		if c.Debug {
			c.Logger("Failed to send request: %s %s: %s", method, url, response.Status)
		}

		return nil, response, errors.New(response.Status)
	}

	if c.Debug {
		c.Logger("Received response body: %s %s %d", method, url, response.StatusCode)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		if c.Debug {
			c.Logger("Failed to read response body: %v: %s %s", err, method, url)
		}

		return nil, response, err
	}

	if output != nil && len(responseData) > 0 {
		err = json.Unmarshal(responseData, output)
		if err != nil {
			if c.Debug {
				c.Logger("Failed to unmarshal response body: %v: %s %s", err, method, url)
			}

			return responseData, response, err
		}
	}

	return responseData, response, nil
}
