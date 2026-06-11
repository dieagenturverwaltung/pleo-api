package exportapi

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type GetExportJobsExec struct {
	config *shared.Config
	ctx    context.Context

	companyID  *string
	statuses   []ExportJobStatus
	pagingInfo shared.PagingInfo
}

func (c *Client) GetExportJobs(ctx context.Context) *GetExportJobsExec {
	return &GetExportJobsExec{ctx: ctx, config: c.config}
}

func (e *GetExportJobsExec) WithContext(ctx context.Context) *GetExportJobsExec {
	e.ctx = ctx
	return e
}

func (e *GetExportJobsExec) WithCompanyID(companyID string) *GetExportJobsExec {
	e.companyID = &companyID
	return e
}

func (e *GetExportJobsExec) WithStatuses(statuses ...ExportJobStatus) *GetExportJobsExec {
	e.statuses = append(e.statuses, statuses...)
	return e
}

func (e *GetExportJobsExec) WithPagingInfo(pagingInfo shared.PagingInfo) *GetExportJobsExec {
	e.pagingInfo = pagingInfo
	return e
}

func (e *GetExportJobsExec) Execute() (*shared.CursorPageResponse[ExportJob], error) {
	queryParams := make(url.Values)
	shared.AddQueryCompanyID(queryParams, e.companyID)
	shared.AddQueryStrings(queryParams, "statuses", e.statuses)
	e.pagingInfo.Apply(queryParams)

	var out shared.CursorPageResponse[ExportJob]
	_, _, err := e.config.SendRequest(e.ctx, "GET", shared.URLWithQuery(basePath+"/export-jobs", queryParams), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

type CreateExportJobExec struct {
	config *shared.Config
	ctx    context.Context
	body   CreateExportJobRequest
}

func (c *Client) CreateExportJob(ctx context.Context) *CreateExportJobExec {
	return &CreateExportJobExec{ctx: ctx, config: c.config}
}

func (e *CreateExportJobExec) WithContext(ctx context.Context) *CreateExportJobExec {
	e.ctx = ctx
	return e
}

func (e *CreateExportJobExec) WithBody(body CreateExportJobRequest) *CreateExportJobExec {
	e.body = body
	return e
}

func (e *CreateExportJobExec) WithAccountingEntryIDs(ids ...string) *CreateExportJobExec {
	e.body.AccountingEntryIDs = append(e.body.AccountingEntryIDs, ids...)
	return e
}

func (e *CreateExportJobExec) WithCompanyID(companyID string) *CreateExportJobExec {
	e.body.CompanyID = companyID
	return e
}

func (e *CreateExportJobExec) WithEmployeeID(employeeID string) *CreateExportJobExec {
	e.body.EmployeeID = &employeeID
	return e
}

func (e *CreateExportJobExec) WithIsInteractive(isInteractive bool) *CreateExportJobExec {
	e.body.IsInteractive = &isInteractive
	return e
}

func (e *CreateExportJobExec) WithVendorBasedBookkeeping(enabled bool) *CreateExportJobExec {
	e.body.Options = &Options{VendorBasedBookkeeping: &enabled}
	return e
}

func (e *CreateExportJobExec) Execute() (*ExportJob, error) {
	var out shared.Response[ExportJob]
	_, _, err := e.config.SendRequest(e.ctx, "POST", basePath+"/export-jobs", e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type GetExportJobExec struct {
	config *shared.Config
	ctx    context.Context
	jobID  string
}

func (c *Client) GetExportJob(ctx context.Context, jobID string) *GetExportJobExec {
	return &GetExportJobExec{ctx: ctx, config: c.config, jobID: jobID}
}

func (e *GetExportJobExec) WithContext(ctx context.Context) *GetExportJobExec {
	e.ctx = ctx
	return e
}

func (e *GetExportJobExec) Execute() (*ExportJob, error) {
	var out shared.Response[ExportJob]
	_, _, err := e.config.SendRequest(e.ctx, "GET", basePath+"/export-jobs/"+url.PathEscape(e.jobID), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type CreateExportJobEventExec struct {
	config *shared.Config
	ctx    context.Context
	body   CreateExportJobEventRequest
}

func (c *Client) CreateExportJobEvent(ctx context.Context) *CreateExportJobEventExec {
	return &CreateExportJobEventExec{ctx: ctx, config: c.config}
}

func (e *CreateExportJobEventExec) WithContext(ctx context.Context) *CreateExportJobEventExec {
	e.ctx = ctx
	return e
}

func (e *CreateExportJobEventExec) WithBody(body CreateExportJobEventRequest) *CreateExportJobEventExec {
	e.body = body
	return e
}

func (e *CreateExportJobEventExec) WithJobID(jobID string) *CreateExportJobEventExec {
	e.body.JobID = jobID
	return e
}

func (e *CreateExportJobEventExec) WithEvent(event ExportJobEventType) *CreateExportJobEventExec {
	e.body.Event = event
	return e
}

func (e *CreateExportJobEventExec) WithFailure(reason string, reasonType ExportJobFailureReasonType) *CreateExportJobEventExec {
	e.body.FailureReason = reason
	e.body.FailureReasonType = &reasonType
	return e
}

func (e *CreateExportJobEventExec) Execute() error {
	_, _, err := e.config.SendRequest(e.ctx, "POST", basePath+"/export-job-events", e.body, nil)
	return err
}
