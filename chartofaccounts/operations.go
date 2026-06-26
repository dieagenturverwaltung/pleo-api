package chartofaccounts

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type CreateAccountExec struct {
	config *shared.Config
	ctx    context.Context
	body   AccountCreateRequest
}

func (c *Client) CreateAccount(ctx context.Context) *CreateAccountExec {
	return &CreateAccountExec{ctx: ctx, config: c.config}
}

func (e *CreateAccountExec) WithContext(ctx context.Context) *CreateAccountExec {
	e.ctx = ctx
	return e
}

func (e *CreateAccountExec) WithBody(body AccountCreateRequest) *CreateAccountExec {
	e.body = body
	return e
}

func (e *CreateAccountExec) WithArchived(archived bool) *CreateAccountExec {
	e.body.Archived = archived
	return e
}

func (e *CreateAccountExec) WithCode(code string) *CreateAccountExec {
	e.body.Code = code
	return e
}

func (e *CreateAccountExec) WithCompanyID(companyID string) *CreateAccountExec {
	e.body.CompanyID = companyID
	return e
}

func (e *CreateAccountExec) WithExternalID(externalID string) *CreateAccountExec {
	e.body.ExternalID = externalID
	return e
}

func (e *CreateAccountExec) WithMetadata(metadata map[string]any) *CreateAccountExec {
	e.body.Metadata = metadata
	return e
}

func (e *CreateAccountExec) WithName(name string) *CreateAccountExec {
	e.body.Name = name
	return e
}

func (e *CreateAccountExec) WithTaxCodeExternalID(taxCodeExternalID string) *CreateAccountExec {
	e.body.TaxCodeExternalID = taxCodeExternalID
	return e
}

func (e *CreateAccountExec) Execute() (*Account, error) {
	var out shared.Response[Account]
	_, _, err := e.config.SendRequest(e.ctx, "POST", basePath, e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type CreateAccountsBatchExec struct {
	config *shared.Config
	ctx    context.Context
	body   AccountBatchCreateRequest
}

func (c *Client) CreateAccountsBatch(ctx context.Context) *CreateAccountsBatchExec {
	return &CreateAccountsBatchExec{ctx: ctx, config: c.config}
}

func (e *CreateAccountsBatchExec) WithContext(ctx context.Context) *CreateAccountsBatchExec {
	e.ctx = ctx
	return e
}

func (e *CreateAccountsBatchExec) WithBody(body AccountBatchCreateRequest) *CreateAccountsBatchExec {
	e.body = body
	return e
}

func (e *CreateAccountsBatchExec) WithCompanyID(companyID string) *CreateAccountsBatchExec {
	e.body.CompanyID = companyID
	return e
}

func (e *CreateAccountsBatchExec) WithItems(items ...AccountBatchRequestItem) *CreateAccountsBatchExec {
	e.body.Items = append(e.body.Items, items...)
	return e
}

func (e *CreateAccountsBatchExec) Execute() (*AccountBatchCreateResponse, error) {
	var out shared.Response[AccountBatchCreateResponse]
	_, _, err := e.config.SendRequest(e.ctx, "POST", basePath+"/batch", e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type GetAccountExec struct {
	config    *shared.Config
	ctx       context.Context
	accountID string
}

func (c *Client) GetAccount(ctx context.Context, accountID string) *GetAccountExec {
	return &GetAccountExec{ctx: ctx, config: c.config, accountID: accountID}
}

func (e *GetAccountExec) WithContext(ctx context.Context) *GetAccountExec {
	e.ctx = ctx
	return e
}

func (e *GetAccountExec) Execute() (*Account, error) {
	var out shared.Response[Account]
	_, _, err := e.config.SendRequest(e.ctx, "GET", basePath+"/"+url.PathEscape(e.accountID), nil, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type UpdateAccountExec struct {
	config    *shared.Config
	ctx       context.Context
	accountID string
	body      AccountUpdateRequest
}

func (c *Client) UpdateAccount(ctx context.Context, accountID string) *UpdateAccountExec {
	return &UpdateAccountExec{ctx: ctx, config: c.config, accountID: accountID}
}

func (e *UpdateAccountExec) WithContext(ctx context.Context) *UpdateAccountExec {
	e.ctx = ctx
	return e
}

func (e *UpdateAccountExec) WithBody(body AccountUpdateRequest) *UpdateAccountExec {
	e.body = body
	return e
}

func (e *UpdateAccountExec) WithArchived(archived bool) *UpdateAccountExec {
	e.body.Archived = archived
	return e
}

func (e *UpdateAccountExec) WithCode(code string) *UpdateAccountExec {
	e.body.Code = code
	return e
}

func (e *UpdateAccountExec) WithMetadata(metadata map[string]any) *UpdateAccountExec {
	e.body.Metadata = metadata
	return e
}

func (e *UpdateAccountExec) WithName(name string) *UpdateAccountExec {
	e.body.Name = name
	return e
}

func (e *UpdateAccountExec) WithTaxCodeExternalID(taxCodeExternalID string) *UpdateAccountExec {
	e.body.TaxCodeExternalID = taxCodeExternalID
	return e
}

func (e *UpdateAccountExec) Execute() (*Account, error) {
	var out shared.Response[Account]
	_, _, err := e.config.SendRequest(e.ctx, "PUT", basePath+"/"+url.PathEscape(e.accountID), e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out.Data, nil
}

type DeleteAccountExec struct {
	config    *shared.Config
	ctx       context.Context
	accountID string
}

func (c *Client) DeleteAccount(ctx context.Context, accountID string) *DeleteAccountExec {
	return &DeleteAccountExec{ctx: ctx, config: c.config, accountID: accountID}
}

func (e *DeleteAccountExec) WithContext(ctx context.Context) *DeleteAccountExec {
	e.ctx = ctx
	return e
}

func (e *DeleteAccountExec) Execute() error {
	_, _, err := e.config.SendRequest(e.ctx, "DELETE", basePath+"/"+url.PathEscape(e.accountID), nil, nil)
	return err
}

type SearchAccountsExec struct {
	config *shared.Config
	ctx    context.Context

	pagingInfo shared.PagingInfo
	body       AccountSearchRequest
}

func (c *Client) SearchAccounts(ctx context.Context) *SearchAccountsExec {
	return &SearchAccountsExec{ctx: ctx, config: c.config}
}

func (e *SearchAccountsExec) WithContext(ctx context.Context) *SearchAccountsExec {
	e.ctx = ctx
	return e
}

func (e *SearchAccountsExec) WithPagingInfo(pagingInfo shared.PagingInfo) *SearchAccountsExec {
	e.pagingInfo = pagingInfo
	return e
}

func (e *SearchAccountsExec) WithBody(body AccountSearchRequest) *SearchAccountsExec {
	e.body = body
	return e
}

func (e *SearchAccountsExec) WithArchived(archived bool) *SearchAccountsExec {
	e.body.Archived = &archived
	return e
}

func (e *SearchAccountsExec) WithCode(code string) *SearchAccountsExec {
	e.body.Code = code
	return e
}

func (e *SearchAccountsExec) WithCodes(codes ...string) *SearchAccountsExec {
	e.body.Codes = append(e.body.Codes, codes...)
	return e
}

func (e *SearchAccountsExec) WithCompanyID(companyID string) *SearchAccountsExec {
	e.body.CompanyID = companyID
	return e
}

func (e *SearchAccountsExec) WithExcludeIfAssignedToCategory(exclude bool) *SearchAccountsExec {
	e.body.ExcludeIfAssignedToCategory = exclude
	return e
}

func (e *SearchAccountsExec) WithExcludeIfAssignedToContraAccount(exclude bool) *SearchAccountsExec {
	e.body.ExcludeIfAssignedToContraAccount = exclude
	return e
}

func (e *SearchAccountsExec) WithExternalID(externalID string) *SearchAccountsExec {
	e.body.ExternalID = externalID
	return e
}

func (e *SearchAccountsExec) WithIDs(ids ...string) *SearchAccountsExec {
	e.body.IDs = append(e.body.IDs, ids...)
	return e
}

func (e *SearchAccountsExec) WithName(name string) *SearchAccountsExec {
	e.body.Name = name
	return e
}

func (e *SearchAccountsExec) Execute() (*shared.CursorPageResponse[Account], error) {
	queryParams := make(url.Values)
	e.pagingInfo.Apply(queryParams)

	var out shared.CursorPageResponse[Account]
	_, _, err := e.config.SendRequest(e.ctx, "POST", shared.URLWithQuery(basePath+":search", queryParams), e.body, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
