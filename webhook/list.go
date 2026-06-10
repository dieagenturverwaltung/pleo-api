package webhook

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type ListExec struct {
	config *shared.Config
	ctx    context.Context

	status     *string
	eventTypes []string
	pagingInfo shared.PagingInfo
}

func (c *Client) List(ctx context.Context) *ListExec {
	return &ListExec{ctx: ctx, config: c.config}
}

func (e *ListExec) WithContext(ctx context.Context) *ListExec {
	e.ctx = ctx
	return e
}

func (e *ListExec) WithStatus(status string) *ListExec {
	e.status = &status
	return e
}

func (e *ListExec) WithEventTypes(eventTypes []string) *ListExec {
	e.eventTypes = eventTypes
	return e
}

func (e *ListExec) WithPagingInfo(pagingInfo shared.PagingInfo) *ListExec {
	e.pagingInfo = pagingInfo
	return e
}

func (e *ListExec) Execute() (*shared.PageResponse[Info], error) {
	queryParams := make(url.Values)
	e.pagingInfo.Apply(queryParams)
	if e.status != nil {
		queryParams.Add("status", *e.status)
	}

	for _, eventType := range e.eventTypes {
		queryParams.Add("event_types", eventType)
	}

	var out shared.PageResponse[Info]
	_, _, err := e.config.SendRequest(e.ctx, "GET", basePath+"?"+queryParams.Encode(), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
