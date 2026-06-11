package companies

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type SearchExec struct {
	config *shared.Config
	ctx    context.Context

	organizationID *string
	pagingInfo     shared.PagingInfo
}

func (c *Client) Search(ctx context.Context) *SearchExec {
	return &SearchExec{ctx: ctx, config: c.config}
}

func (e *SearchExec) WithContext(ctx context.Context) *SearchExec {
	e.ctx = ctx
	return e
}

func (e *SearchExec) WithOrganizationID(organizationID string) *SearchExec {
	e.organizationID = &organizationID
	return e
}

func (e *SearchExec) WithPagingInfo(pagingInfo shared.PagingInfo) *SearchExec {
	e.pagingInfo = pagingInfo
	return e
}

func (e *SearchExec) Execute() (*shared.CursorPageResponse[Company], error) {
	queryParams := make(url.Values)
	shared.AddQueryString(queryParams, "organization_id", e.organizationID)
	e.pagingInfo.Apply(queryParams)

	var out shared.CursorPageResponse[Company]
	_, _, err := e.config.SendRequest(e.ctx, "GET", shared.URLWithQuery(basePath, queryParams), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

type GetExec struct {
	config    *shared.Config
	ctx       context.Context
	companyID string
}

func (c *Client) Get(ctx context.Context, companyID string) *GetExec {
	return &GetExec{ctx: ctx, config: c.config, companyID: companyID}
}

func (e *GetExec) WithContext(ctx context.Context) *GetExec {
	e.ctx = ctx
	return e
}

func (e *GetExec) Execute() (*Company, error) {
	var out shared.Response[Company]
	_, _, err := e.config.SendRequest(e.ctx, "GET", basePath+"/"+url.PathEscape(e.companyID), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}
