package webhook

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type UpdateExec struct {
	config *shared.Config
	ctx    context.Context
	id     string
	body   createBody
}

func (c *Client) Update(ctx context.Context, id string) *UpdateExec {
	body := createBody{
		Status: StatusActive,
	}
	return &UpdateExec{ctx: ctx, config: c.config, id: id, body: body}
}

func (e *UpdateExec) WithContext(ctx context.Context) *UpdateExec {
	e.ctx = ctx
	return e
}

func (e *UpdateExec) WithEndpointUrl(endpointURL string) *UpdateExec {
	e.body.EndpointUrl = endpointURL
	return e
}

func (e *UpdateExec) WithEventTypes(eventTypes ...Event) *UpdateExec {
	e.body.EventTypes = append(e.body.EventTypes, eventTypes...)
	return e
}

func (e *UpdateExec) WithCustomHeaders(headers map[string]any) *UpdateExec {
	e.body.CustomHeaders = headers
	return e
}

func (e *UpdateExec) WithStatus(status Status) *UpdateExec {
	e.body.Status = status
	return e
}

func (e *UpdateExec) WithTokenAuth(token string) *UpdateExec {
	e.body.EndpointAuth = &EndpointAuth{
		Credentials: &EndpointAuthCredentials{
			Token: token,
		},
		Type: "NONE",
	}
	return e
}

func (e *UpdateExec) WithBasicAuth(username, password string) *UpdateExec {
	e.body.EndpointAuth = &EndpointAuth{
		Credentials: &EndpointAuthCredentials{
			Username: username,
			Password: password,
		},
		Type: "BASIC",
	}
	return e
}

func (e *UpdateExec) Execute() (*Info, error) {
	var out shared.Response[Info]
	_, _, err := e.config.SendRequest(e.ctx, "PUT", basePath+"/"+url.PathEscape(e.id), e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}
