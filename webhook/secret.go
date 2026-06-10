package webhook

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type SecretExec struct {
	config *shared.Config
	ctx    context.Context
	id     string
}

func (c *Client) Secret(ctx context.Context, id string) *SecretExec {
	return &SecretExec{ctx: ctx, config: c.config, id: id}
}

func (e *SecretExec) WithContext(ctx context.Context) *SecretExec {
	e.ctx = ctx
	return e
}

func (e *SecretExec) Execute() (*Secret, error) {
	var out shared.Response[Secret]
	_, _, err := e.config.SendRequest(e.ctx, "GET", basePath+"/"+url.PathEscape(e.id)+"/secret", nil, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}
