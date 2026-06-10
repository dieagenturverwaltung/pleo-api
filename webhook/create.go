package webhook

import (
	"context"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type Event string

const (
	EventJobCreated    Event = "v1.export.job-created"
	EventVendorCreated Event = "v1.vendor.created"
)

type Status string

const (
	StatusActive   Status = "ACTIVE"
	StatusInactive Status = "INACTIVE"
)

type CreateExec struct {
	config *shared.Config
	ctx    context.Context
	body   createBody
}

type createBody struct {
	EndpointUrl   string         `json:"endpointUrl"`
	EventTypes    []Event        `json:"eventTypes"`
	CustomHeaders map[string]any `json:"customHeaders,omitempty"`
	EndpointAuth  *EndpointAuth  `json:"endpointAuth,omitempty"`
	Status        Status         `json:"status,omitempty"`
}

func (c *Client) Create(ctx context.Context) *CreateExec {
	body := createBody{
		Status: StatusActive,
	}
	return &CreateExec{ctx: ctx, config: c.config, body: body}
}

func (ce *CreateExec) WithContext(ctx context.Context) *CreateExec {
	ce.ctx = ctx
	return ce
}

func (ce *CreateExec) WithEndpointUrl(url string) *CreateExec {
	ce.body.EndpointUrl = url
	return ce
}

func (ce *CreateExec) WithEventTypes(eventTypes ...Event) *CreateExec {
	ce.body.EventTypes = append(ce.body.EventTypes, eventTypes...)
	return ce
}

func (ce *CreateExec) WithCustomHeaders(headers map[string]any) *CreateExec {
	ce.body.CustomHeaders = headers
	return ce
}

func (ce *CreateExec) WithTokenAuth(token string) *CreateExec {
	ce.body.EndpointAuth = &EndpointAuth{
		Credentials: &EndpointAuthCredentials{
			Token: token,
		},
		Type: "NONE",
	}
	return ce
}

func (ce *CreateExec) WithBasicAuth(username, password string) *CreateExec {
	ce.body.EndpointAuth = &EndpointAuth{
		Credentials: &EndpointAuthCredentials{
			Username: username,
			Password: password,
		},
		Type: "BASIC",
	}
	return ce
}

func (ce *CreateExec) Execute() (*Info, error) {
	var response shared.Response[Info]
	_, _, err := ce.config.SendRequest(ce.ctx, "POST", basePath, ce.body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}
