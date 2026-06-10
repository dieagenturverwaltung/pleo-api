package exportapi

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type GetExportItemsExec struct {
	config *shared.Config
	ctx    context.Context

	jobID      *string
	pagingInfo shared.PagingInfo
}

func (c *Client) GetExportItems(ctx context.Context) *GetExportItemsExec {
	return &GetExportItemsExec{ctx: ctx, config: c.config}
}

func (e *GetExportItemsExec) WithContext(ctx context.Context) *GetExportItemsExec {
	e.ctx = ctx
	return e
}

func (e *GetExportItemsExec) WithJobID(jobID string) *GetExportItemsExec {
	e.jobID = &jobID
	return e
}

func (e *GetExportItemsExec) WithPagingInfo(pagingInfo shared.PagingInfo) *GetExportItemsExec {
	e.pagingInfo = pagingInfo
	return e
}

func (e *GetExportItemsExec) Execute() (*shared.CursorPageResponse[ExportItem], error) {
	queryParams := make(url.Values)
	shared.AddQueryString(queryParams, "job_id", e.jobID)
	e.pagingInfo.Apply(queryParams)

	var out shared.CursorPageResponse[ExportItem]
	_, _, err := e.config.SendRequest(e.ctx, "GET", shared.URLWithQuery(basePath+"/export-items", queryParams), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

type GetExportJobItemByAccountingEntryIDExec struct {
	config *shared.Config
	ctx    context.Context

	companyID         *string
	accountingEntryID *string
}

func (c *Client) GetExportJobItemByAccountingEntryID(ctx context.Context) *GetExportJobItemByAccountingEntryIDExec {
	return &GetExportJobItemByAccountingEntryIDExec{ctx: ctx, config: c.config}
}

func (e *GetExportJobItemByAccountingEntryIDExec) WithContext(ctx context.Context) *GetExportJobItemByAccountingEntryIDExec {
	e.ctx = ctx
	return e
}

func (e *GetExportJobItemByAccountingEntryIDExec) WithCompanyID(companyID string) *GetExportJobItemByAccountingEntryIDExec {
	e.companyID = &companyID
	return e
}

func (e *GetExportJobItemByAccountingEntryIDExec) WithAccountingEntryID(accountingEntryID string) *GetExportJobItemByAccountingEntryIDExec {
	e.accountingEntryID = &accountingEntryID
	return e
}

func (e *GetExportJobItemByAccountingEntryIDExec) Execute() (*ExportJobItem, error) {
	queryParams := make(url.Values)
	shared.AddQueryCompanyID(queryParams, e.companyID, e.config)
	shared.AddQueryString(queryParams, "accounting_entry_id", e.accountingEntryID)

	var out shared.Response[ExportJobItem]
	_, _, err := e.config.SendRequest(e.ctx, "GET", shared.URLWithQuery(basePath+"/export-job-items", queryParams), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type SearchExportJobItemsExec struct {
	config *shared.Config
	ctx    context.Context

	companyID  *string
	pagingInfo shared.PagingInfo
	body       SearchExportJobItemsRequest
}

func (c *Client) SearchExportJobItems(ctx context.Context) *SearchExportJobItemsExec {
	return &SearchExportJobItemsExec{ctx: ctx, config: c.config}
}

func (e *SearchExportJobItemsExec) WithContext(ctx context.Context) *SearchExportJobItemsExec {
	e.ctx = ctx
	return e
}

func (e *SearchExportJobItemsExec) WithCompanyID(companyID string) *SearchExportJobItemsExec {
	e.companyID = &companyID
	return e
}

func (e *SearchExportJobItemsExec) WithPagingInfo(pagingInfo shared.PagingInfo) *SearchExportJobItemsExec {
	e.pagingInfo = pagingInfo
	return e
}

func (e *SearchExportJobItemsExec) WithBody(body SearchExportJobItemsRequest) *SearchExportJobItemsExec {
	e.body = body
	return e
}

func (e *SearchExportJobItemsExec) WithAccountingEntryIDs(ids ...string) *SearchExportJobItemsExec {
	e.body.AccountingEntryIDs = append(e.body.AccountingEntryIDs, ids...)
	return e
}

func (e *SearchExportJobItemsExec) WithStatus(status ExportJobItemStatus) *SearchExportJobItemsExec {
	e.body.Status = &status
	return e
}

func (e *SearchExportJobItemsExec) Execute() (*shared.CursorPageResponse[ExportJobItem], error) {
	queryParams := make(url.Values)
	shared.AddQueryCompanyID(queryParams, e.companyID, e.config)
	e.pagingInfo.Apply(queryParams)

	var out shared.CursorPageResponse[ExportJobItem]
	_, _, err := e.config.SendRequest(e.ctx, "POST", shared.URLWithQuery(basePath+"/export-job-items:search", queryParams), e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

type GetExportJobItemsExec struct {
	config *shared.Config
	ctx    context.Context

	jobID      string
	status     *ExportJobItemStatus
	pagingInfo shared.PagingInfo
}

func (c *Client) GetExportJobItems(ctx context.Context, jobID string) *GetExportJobItemsExec {
	return &GetExportJobItemsExec{ctx: ctx, config: c.config, jobID: jobID}
}

func (e *GetExportJobItemsExec) WithContext(ctx context.Context) *GetExportJobItemsExec {
	e.ctx = ctx
	return e
}

func (e *GetExportJobItemsExec) WithStatus(status ExportJobItemStatus) *GetExportJobItemsExec {
	e.status = &status
	return e
}

func (e *GetExportJobItemsExec) WithPagingInfo(pagingInfo shared.PagingInfo) *GetExportJobItemsExec {
	e.pagingInfo = pagingInfo
	return e
}

func (e *GetExportJobItemsExec) Execute() (*shared.CursorPageResponse[ExportJobItem], error) {
	queryParams := make(url.Values)
	if e.status != nil {
		queryParams.Add("status", string(*e.status))
	}
	e.pagingInfo.Apply(queryParams)

	path := basePath + "/export-jobs/" + url.PathEscape(e.jobID) + "/items"
	var out shared.CursorPageResponse[ExportJobItem]
	_, _, err := e.config.SendRequest(e.ctx, "GET", shared.URLWithQuery(path, queryParams), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

type UpdateExportJobItemsExec struct {
	config *shared.Config
	ctx    context.Context
	jobID  string
	body   []UpdateExportJobItem
}

func (c *Client) UpdateExportJobItems(ctx context.Context, jobID string) *UpdateExportJobItemsExec {
	return &UpdateExportJobItemsExec{ctx: ctx, config: c.config, jobID: jobID}
}

func (e *UpdateExportJobItemsExec) WithContext(ctx context.Context) *UpdateExportJobItemsExec {
	e.ctx = ctx
	return e
}

func (e *UpdateExportJobItemsExec) WithBody(body []UpdateExportJobItem) *UpdateExportJobItemsExec {
	e.body = body
	return e
}

func (e *UpdateExportJobItemsExec) WithItems(items ...UpdateExportJobItem) *UpdateExportJobItemsExec {
	e.body = append(e.body, items...)
	return e
}

func (e *UpdateExportJobItemsExec) Execute() (*ExportJobItemUpdate, error) {
	path := basePath + "/export-jobs/" + url.PathEscape(e.jobID) + "/items"
	var out ExportJobItemUpdate
	_, _, err := e.config.SendRequest(e.ctx, "PUT", path, e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
