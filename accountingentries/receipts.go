package accountingentries

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type GetReceiptsExec struct {
	config            *shared.Config
	ctx               context.Context
	accountingEntryID string
	source            []ReceiptSource
	fileType          ReceiptFileType
	pagingInfo        shared.PagingInfo
}

func (c *Client) GetReceipts(ctx context.Context, accountingEntryID string) *GetReceiptsExec {
	return &GetReceiptsExec{ctx: ctx, config: c.config, accountingEntryID: accountingEntryID}
}

func (e *GetReceiptsExec) WithContext(ctx context.Context) *GetReceiptsExec {
	e.ctx = ctx
	return e
}

func (e *GetReceiptsExec) WithSource(source ...ReceiptSource) *GetReceiptsExec {
	e.source = append(e.source, source...)
	return e
}

func (e *GetReceiptsExec) WithFileType(fileType ReceiptFileType) *GetReceiptsExec {
	e.fileType = fileType
	return e
}

func (e *GetReceiptsExec) WithPagingInfo(pagingInfo shared.PagingInfo) *GetReceiptsExec {
	e.pagingInfo = pagingInfo
	return e
}

func (e *GetReceiptsExec) Execute() (*shared.CursorPageResponse[Receipt], *shared.ResponseExtra, error) {
	queryParams := make(url.Values)
	shared.AddQueryStrings(queryParams, "source", e.source)
	shared.AddQueryStringValue(queryParams, "file_type", string(e.fileType))
	e.pagingInfo.Apply(queryParams)

	path := basePath + "/" + url.PathEscape(e.accountingEntryID) + "/receipts"
	var out shared.CursorPageResponse[Receipt]
	_, response, err := e.config.SendRequest(e.ctx, "GET", shared.URLWithQuery(path, queryParams), nil, &out)
	extra := shared.ResponseExtraFromResponse(response)
	if err != nil {
		return nil, extra, err
	}

	return &out, extra, nil
}

type GetReceiptExec struct {
	config            *shared.Config
	ctx               context.Context
	accountingEntryID string
	receiptID         string
	fileType          ReceiptFileType
}

// GetReceipt retrieves receipt metadata and a time-bound URL to the file.
func (c *Client) GetReceipt(ctx context.Context, accountingEntryID, receiptID string) *GetReceiptExec {
	return &GetReceiptExec{ctx: ctx, config: c.config, accountingEntryID: accountingEntryID, receiptID: receiptID}
}

func (e *GetReceiptExec) WithContext(ctx context.Context) *GetReceiptExec {
	e.ctx = ctx
	return e
}

func (e *GetReceiptExec) WithFileType(fileType ReceiptFileType) *GetReceiptExec {
	e.fileType = fileType
	return e
}

func (e *GetReceiptExec) Execute() (*Receipt, *shared.ResponseExtra, error) {
	queryParams := make(url.Values)
	shared.AddQueryStringValue(queryParams, "file_type", string(e.fileType))

	path := basePath + "/" + url.PathEscape(e.accountingEntryID) + "/receipts/" + url.PathEscape(e.receiptID)
	var out shared.Response[Receipt]
	_, response, err := e.config.SendRequest(e.ctx, "GET", shared.URLWithQuery(path, queryParams), nil, &out)
	extra := shared.ResponseExtraFromResponse(response)
	if err != nil {
		return nil, extra, err
	}

	return &out.Data, extra, nil
}

type SearchReceiptsExec struct {
	config         *shared.Config
	ctx            context.Context
	companyID      *string
	partnerID      *string
	organizationID *string
	pagingInfo     shared.PagingInfo
	body           SearchReceiptsRequest
}

func (c *Client) SearchReceipts(ctx context.Context) *SearchReceiptsExec {
	return &SearchReceiptsExec{ctx: ctx, config: c.config}
}

func (e *SearchReceiptsExec) WithContext(ctx context.Context) *SearchReceiptsExec {
	e.ctx = ctx
	return e
}

func (e *SearchReceiptsExec) WithCompanyID(companyID string) *SearchReceiptsExec {
	e.companyID = &companyID
	return e
}

func (e *SearchReceiptsExec) WithPartnerID(partnerID string) *SearchReceiptsExec {
	e.partnerID = &partnerID
	return e
}

func (e *SearchReceiptsExec) WithOrganizationID(organizationID string) *SearchReceiptsExec {
	e.organizationID = &organizationID
	return e
}

func (e *SearchReceiptsExec) WithPagingInfo(pagingInfo shared.PagingInfo) *SearchReceiptsExec {
	e.pagingInfo = pagingInfo
	return e
}

func (e *SearchReceiptsExec) WithBody(body SearchReceiptsRequest) *SearchReceiptsExec {
	e.body = body
	return e
}

func (e *SearchReceiptsExec) WithIncludeDeleted(includeDeleted bool) *SearchReceiptsExec {
	e.body.IncludeDeleted = includeDeleted
	return e
}

func (e *SearchReceiptsExec) WithAccountingEntryIDs(ids ...string) *SearchReceiptsExec {
	e.body.AccountingEntryIDs = append(e.body.AccountingEntryIDs, ids...)
	return e
}

func (e *SearchReceiptsExec) WithEmployeeIDs(ids ...string) *SearchReceiptsExec {
	e.body.EmployeeIDs = append(e.body.EmployeeIDs, ids...)
	return e
}

func (e *SearchReceiptsExec) WithFileType(fileType ReceiptFileType) *SearchReceiptsExec {
	e.body.FileType = &fileType
	return e
}

func (e *SearchReceiptsExec) WithMIMEType(mimeType ...ReceiptMIMEType) *SearchReceiptsExec {
	e.body.MIMEType = append(e.body.MIMEType, mimeType...)
	return e
}

func (e *SearchReceiptsExec) WithSource(source ...ReceiptSource) *SearchReceiptsExec {
	e.body.Source = append(e.body.Source, source...)
	return e
}

func (e *SearchReceiptsExec) Execute() (*shared.CursorPageResponse[Receipt], *shared.ResponseExtra, error) {
	queryParams := make(url.Values)
	shared.AddQueryCompanyID(queryParams, e.companyID)
	shared.AddQueryString(queryParams, "partner_id", e.partnerID)
	shared.AddQueryString(queryParams, "organization_id", e.organizationID)
	e.pagingInfo.Apply(queryParams)

	var out shared.CursorPageResponse[Receipt]
	_, response, err := e.config.SendRequest(e.ctx, "POST", shared.URLWithQuery(basePath+"/receipts:search", queryParams), e.body, &out)
	extra := shared.ResponseExtraFromResponse(response)
	if err != nil {
		return nil, extra, err
	}

	return &out, extra, nil
}
