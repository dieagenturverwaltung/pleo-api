package webhook

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type DeleteExec struct {
	config *shared.Config
	ctx    context.Context
	id     string
}

func (c *Client) Delete(ctx context.Context, id string) *DeleteExec {
	return &DeleteExec{ctx: ctx, config: c.config, id: id}
}

func (e *DeleteExec) WithContext(ctx context.Context) *DeleteExec {
	e.ctx = ctx
	return e
}

func (e *DeleteExec) Execute() error {
	_, _, err := e.config.SendRequest(e.ctx, "DELETE", basePath+"/"+url.PathEscape(e.id), nil, nil)
	return err
}
