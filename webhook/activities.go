package webhook

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type ActivitiesExec struct {
	config *shared.Config
	ctx    context.Context
	id     string

	operationTypes []ActivityOperationType
	pagingInfo     shared.PagingInfo
}

func (c *Client) Activities(ctx context.Context, id string) *ActivitiesExec {
	return &ActivitiesExec{ctx: ctx, config: c.config, id: id}
}

func (e *ActivitiesExec) WithContext(ctx context.Context) *ActivitiesExec {
	e.ctx = ctx
	return e
}

func (e *ActivitiesExec) WithOperationTypes(operationTypes ...ActivityOperationType) *ActivitiesExec {
	e.operationTypes = append(e.operationTypes, operationTypes...)
	return e
}

func (e *ActivitiesExec) WithPagingInfo(pagingInfo shared.PagingInfo) *ActivitiesExec {
	e.pagingInfo = pagingInfo
	return e
}

func (e *ActivitiesExec) Execute() (*shared.CursorPageResponse[Activity], error) {
	queryParams := make(url.Values)
	shared.AddQueryStrings(queryParams, "operation_types", e.operationTypes)
	e.pagingInfo.Apply(queryParams)

	path := basePath + "/" + url.PathEscape(e.id) + "/activities"
	var out shared.CursorPageResponse[Activity]
	_, _, err := e.config.SendRequest(e.ctx, "GET", shared.URLWithQuery(path, queryParams), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
