package webhook

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type GetExec struct {
	config *shared.Config
	ctx    context.Context
	id     string
}

func (c *Client) Get(ctx context.Context, id string) *GetExec {
	return &GetExec{ctx: ctx, config: c.config, id: id}
}

func (e *GetExec) WithContext(ctx context.Context) *GetExec {
	e.ctx = ctx
	return e
}

func (e *GetExec) Execute() (*Info, error) {
	var out shared.Response[Info]
	_, _, err := e.config.SendRequest(e.ctx, "GET", basePath+"/"+url.PathEscape(e.id), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}
