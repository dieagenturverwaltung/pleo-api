package taxcodes

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type CreateTaxCodeExec struct {
	config *shared.Config
	ctx    context.Context
	body   TaxCodeCreateUpdateModel
}

func (c *Client) CreateTaxCode(ctx context.Context) *CreateTaxCodeExec {
	return &CreateTaxCodeExec{ctx: ctx, config: c.config}
}

func (e *CreateTaxCodeExec) WithContext(ctx context.Context) *CreateTaxCodeExec {
	e.ctx = ctx
	return e
}

func (e *CreateTaxCodeExec) WithBody(body TaxCodeCreateUpdateModel) *CreateTaxCodeExec {
	e.body = body
	return e
}

func (e *CreateTaxCodeExec) WithArchived(archived bool) *CreateTaxCodeExec {
	e.body.Archived = archived
	return e
}

func (e *CreateTaxCodeExec) WithCode(code string) *CreateTaxCodeExec {
	e.body.Code = code
	return e
}

func (e *CreateTaxCodeExec) WithCompanyID(companyID string) *CreateTaxCodeExec {
	e.body.CompanyID = companyID
	return e
}

func (e *CreateTaxCodeExec) WithIngoingTaxAccount(account string) *CreateTaxCodeExec {
	e.body.IngoingTaxAccount = account
	return e
}

func (e *CreateTaxCodeExec) WithName(name string) *CreateTaxCodeExec {
	e.body.Name = name
	return e
}

func (e *CreateTaxCodeExec) WithOutgoingTaxAccount(account string) *CreateTaxCodeExec {
	e.body.OutgoingTaxAccount = account
	return e
}

func (e *CreateTaxCodeExec) WithRate(rate float64) *CreateTaxCodeExec {
	e.body.Rate = rate
	return e
}

func (e *CreateTaxCodeExec) WithType(taxCodeType TaxCodeType) *CreateTaxCodeExec {
	e.body.Type = taxCodeType
	return e
}

func (e *CreateTaxCodeExec) Execute() (*TaxCodeModel, error) {
	var out shared.Response[TaxCodeModel]
	_, _, err := e.config.SendRequest(e.ctx, "POST", basePath, e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type GetTaxCodeExec struct {
	config    *shared.Config
	ctx       context.Context
	taxCodeID string
}

func (c *Client) GetTaxCode(ctx context.Context, taxCodeID string) *GetTaxCodeExec {
	return &GetTaxCodeExec{ctx: ctx, config: c.config, taxCodeID: taxCodeID}
}

func (e *GetTaxCodeExec) WithContext(ctx context.Context) *GetTaxCodeExec {
	e.ctx = ctx
	return e
}

func (e *GetTaxCodeExec) Execute() (*TaxCodeModel, error) {
	var out shared.Response[TaxCodeModel]
	_, _, err := e.config.SendRequest(e.ctx, "GET", basePath+"/"+url.PathEscape(e.taxCodeID), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type UpdateTaxCodeExec struct {
	config    *shared.Config
	ctx       context.Context
	taxCodeID string
	body      TaxCodeCreateUpdateModel
}

func (c *Client) UpdateTaxCode(ctx context.Context, taxCodeID string) *UpdateTaxCodeExec {
	return &UpdateTaxCodeExec{ctx: ctx, config: c.config, taxCodeID: taxCodeID}
}

func (e *UpdateTaxCodeExec) WithContext(ctx context.Context) *UpdateTaxCodeExec {
	e.ctx = ctx
	return e
}

func (e *UpdateTaxCodeExec) WithBody(body TaxCodeCreateUpdateModel) *UpdateTaxCodeExec {
	e.body = body
	return e
}

func (e *UpdateTaxCodeExec) WithArchived(archived bool) *UpdateTaxCodeExec {
	e.body.Archived = archived
	return e
}

func (e *UpdateTaxCodeExec) WithCode(code string) *UpdateTaxCodeExec {
	e.body.Code = code
	return e
}

func (e *UpdateTaxCodeExec) WithCompanyID(companyID string) *UpdateTaxCodeExec {
	e.body.CompanyID = companyID
	return e
}

func (e *UpdateTaxCodeExec) WithIngoingTaxAccount(account string) *UpdateTaxCodeExec {
	e.body.IngoingTaxAccount = account
	return e
}

func (e *UpdateTaxCodeExec) WithName(name string) *UpdateTaxCodeExec {
	e.body.Name = name
	return e
}

func (e *UpdateTaxCodeExec) WithOutgoingTaxAccount(account string) *UpdateTaxCodeExec {
	e.body.OutgoingTaxAccount = account
	return e
}

func (e *UpdateTaxCodeExec) WithRate(rate float64) *UpdateTaxCodeExec {
	e.body.Rate = rate
	return e
}

func (e *UpdateTaxCodeExec) WithType(taxCodeType TaxCodeType) *UpdateTaxCodeExec {
	e.body.Type = taxCodeType
	return e
}

func (e *UpdateTaxCodeExec) Execute() (*TaxCodeModel, error) {
	var out shared.Response[TaxCodeModel]
	_, _, err := e.config.SendRequest(e.ctx, "PUT", basePath+"/"+url.PathEscape(e.taxCodeID), e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type DeleteTaxCodeExec struct {
	config    *shared.Config
	ctx       context.Context
	taxCodeID string
}

func (c *Client) DeleteTaxCode(ctx context.Context, taxCodeID string) *DeleteTaxCodeExec {
	return &DeleteTaxCodeExec{ctx: ctx, config: c.config, taxCodeID: taxCodeID}
}

func (e *DeleteTaxCodeExec) WithContext(ctx context.Context) *DeleteTaxCodeExec {
	e.ctx = ctx
	return e
}

func (e *DeleteTaxCodeExec) Execute() error {
	_, _, err := e.config.SendRequest(e.ctx, "DELETE", basePath+"/"+url.PathEscape(e.taxCodeID), nil, nil)
	return err
}

type GetTaxCodesExec struct {
	config *shared.Config
	ctx    context.Context

	companyID  *string
	pagingInfo shared.PagingInfo
	body       TaxCodeSearchRequest
}

func (c *Client) GetTaxCodes(ctx context.Context) *GetTaxCodesExec {
	return &GetTaxCodesExec{ctx: ctx, config: c.config}
}

func (e *GetTaxCodesExec) WithContext(ctx context.Context) *GetTaxCodesExec {
	e.ctx = ctx
	return e
}

func (e *GetTaxCodesExec) WithCompanyID(companyID string) *GetTaxCodesExec {
	e.companyID = &companyID
	return e
}

func (e *GetTaxCodesExec) WithPagingInfo(pagingInfo shared.PagingInfo) *GetTaxCodesExec {
	e.pagingInfo = pagingInfo
	return e
}

func (e *GetTaxCodesExec) WithBody(body TaxCodeSearchRequest) *GetTaxCodesExec {
	e.body = body
	return e
}

func (e *GetTaxCodesExec) WithArchived(archived bool) *GetTaxCodesExec {
	e.body.Archived = &archived
	return e
}

func (e *GetTaxCodesExec) WithTaxCodeIDs(taxCodeIDs ...string) *GetTaxCodesExec {
	e.body.TaxCodeIDs = append(e.body.TaxCodeIDs, taxCodeIDs...)
	return e
}

func (e *GetTaxCodesExec) WithType(taxCodeType TaxCodeType) *GetTaxCodesExec {
	e.body.Type = taxCodeType
	return e
}

func (e *GetTaxCodesExec) Execute() (*shared.CursorPageResponse[TaxCodeModel], error) {
	queryParams := make(url.Values)
	shared.AddQueryCompanyID(queryParams, e.companyID)
	e.pagingInfo.Apply(queryParams)

	var out shared.CursorPageResponse[TaxCodeModel]
	_, _, err := e.config.SendRequest(e.ctx, "POST", shared.URLWithQuery(basePath+":search", queryParams), e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
