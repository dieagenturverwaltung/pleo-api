package shared

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
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
			if response != nil {
				c.logFailedResponseBody(method, url, response)
			} else {
				c.logRequestErrorBody(method, url, err)
			}
		}

		return nil, response, err
	}

	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		if c.Debug {
			c.Logger("Failed to send request: %s %s: %s", method, url, response.Status)
			c.logFailedResponseBody(method, url, response)
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

func (c *Config) logFailedResponseBody(method, url string, response *http.Response) {
	if response.Body == nil {
		return
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		c.Logger("Failed to read response body: %v: %s %s", err, method, url)
		return
	}

	c.Logger("Response body: %s %s: %s", method, url, string(responseData))

	response.Body.Close()
	response.Body = io.NopCloser(bytes.NewBuffer(responseData))
}

func (c *Config) logRequestErrorBody(method, url string, err error) {
	var retrieveErr *oauth2.RetrieveError
	if !errors.As(err, &retrieveErr) || len(retrieveErr.Body) == 0 {
		return
	}

	c.Logger("Response body: %s %s: %s", method, url, string(retrieveErr.Body))
}
