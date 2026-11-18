package accountingentries

import (
	"context"
	"net/url"

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

type GetAccountingEntryExec struct {
	config                       *shared.Config
	ctx                          context.Context
	accountingEntryID            string
	includeDeleted               *bool
	includePerDiemRelatedEntries *bool
}

func (c *Client) GetAccountingEntry(ctx context.Context, accountingEntryID string) *GetAccountingEntryExec {
	return &GetAccountingEntryExec{ctx: ctx, config: c.config, accountingEntryID: accountingEntryID}
}

func (e *GetAccountingEntryExec) WithContext(ctx context.Context) *GetAccountingEntryExec {
	e.ctx = ctx
	return e
}

func (e *GetAccountingEntryExec) WithIncludeDeleted(includeDeleted bool) *GetAccountingEntryExec {
	e.includeDeleted = &includeDeleted
	return e
}

func (e *GetAccountingEntryExec) WithIncludePerDiemRelatedEntries(include bool) *GetAccountingEntryExec {
	e.includePerDiemRelatedEntries = &include
	return e
}

func (e *GetAccountingEntryExec) Execute() (*AccountingEntry, *shared.ResponseExtra, error) {
	queryParams := make(url.Values)
	shared.AddQueryBool(queryParams, "include_deleted", e.includeDeleted)
	shared.AddQueryBool(queryParams, "include_per_diem_related_entries", e.includePerDiemRelatedEntries)

	path := basePath + "/" + url.PathEscape(e.accountingEntryID)
	var out shared.Response[AccountingEntry]
	_, response, err := e.config.SendRequest(e.ctx, "GET", shared.URLWithQuery(path, queryParams), nil, &out)
	extra := shared.ResponseExtraFromResponse(response)
	if err != nil {
		return nil, extra, err
	}

	return &out.Data, extra, nil
}

type SearchAccountingEntriesExec struct {
	config         *shared.Config
	ctx            context.Context
	companyID      *string
	organizationID *string
	pagingInfo     shared.PagingInfo
	body           SearchAccountingEntriesRequest
}

func (c *Client) SearchAccountingEntries(ctx context.Context) *SearchAccountingEntriesExec {
	return &SearchAccountingEntriesExec{ctx: ctx, config: c.config}
}

func (e *SearchAccountingEntriesExec) WithContext(ctx context.Context) *SearchAccountingEntriesExec {
	e.ctx = ctx
	return e
}

func (e *SearchAccountingEntriesExec) WithCompanyID(companyID string) *SearchAccountingEntriesExec {
	e.companyID = &companyID
	return e
}

func (e *SearchAccountingEntriesExec) WithOrganizationID(organizationID string) *SearchAccountingEntriesExec {
	e.organizationID = &organizationID
	return e
}

func (e *SearchAccountingEntriesExec) WithPagingInfo(pagingInfo shared.PagingInfo) *SearchAccountingEntriesExec {
	e.pagingInfo = pagingInfo
	return e
}

func (e *SearchAccountingEntriesExec) WithBody(body SearchAccountingEntriesRequest) *SearchAccountingEntriesExec {
	e.body = body
	return e
}

func (e *SearchAccountingEntriesExec) WithIncludeDeleted(includeDeleted bool) *SearchAccountingEntriesExec {
	e.body.IncludeDeleted = includeDeleted
	return e
}

func (e *SearchAccountingEntriesExec) WithAccountIDs(ids ...string) *SearchAccountingEntriesExec {
	e.body.AccountIDs = append(e.body.AccountIDs, ids...)
	return e
}

func (e *SearchAccountingEntriesExec) WithBookkeepingDateEnd(date string) *SearchAccountingEntriesExec {
	e.body.BookkeepingDateEnd = date
	return e
}

func (e *SearchAccountingEntriesExec) WithBookkeepingDateStart(date string) *SearchAccountingEntriesExec {
	e.body.BookkeepingDateStart = date
	return e
}

func (e *SearchAccountingEntriesExec) WithEmployeeIDs(ids ...string) *SearchAccountingEntriesExec {
	e.body.EmployeeIDs = append(e.body.EmployeeIDs, ids...)
	return e
}

func (e *SearchAccountingEntriesExec) WithExportStatus(status ...ExportStatus) *SearchAccountingEntriesExec {
	e.body.ExportStatus = append(e.body.ExportStatus, status...)
	return e
}

func (e *SearchAccountingEntriesExec) WithFamilies(families ...Family) *SearchAccountingEntriesExec {
	e.body.Families = append(e.body.Families, families...)
	return e
}

func (e *SearchAccountingEntriesExec) WithForeignExtensionIDs(ids ...string) *SearchAccountingEntriesExec {
	e.body.ForeignExtensionIDs = append(e.body.ForeignExtensionIDs, ids...)
	return e
}

func (e *SearchAccountingEntriesExec) WithHasAccount(hasAccount bool) *SearchAccountingEntriesExec {
	e.body.HasAccount = &hasAccount
	return e
}

func (e *SearchAccountingEntriesExec) WithHasBookkeepingDate(hasBookkeepingDate bool) *SearchAccountingEntriesExec {
	e.body.HasBookkeepingDate = &hasBookkeepingDate
	return e
}

func (e *SearchAccountingEntriesExec) WithHasNote(hasNote bool) *SearchAccountingEntriesExec {
	e.body.HasNote = &hasNote
	return e
}

func (e *SearchAccountingEntriesExec) WithPerformedAtEnd(date string) *SearchAccountingEntriesExec {
	e.body.PerformedAtEnd = date
	return e
}

func (e *SearchAccountingEntriesExec) WithPerformedAtStart(date string) *SearchAccountingEntriesExec {
	e.body.PerformedAtStart = date
	return e
}

func (e *SearchAccountingEntriesExec) WithReviewStatus(status ...ReviewStatus) *SearchAccountingEntriesExec {
	e.body.ReviewStatus = append(e.body.ReviewStatus, status...)
	return e
}

func (e *SearchAccountingEntriesExec) WithSettled(settled bool) *SearchAccountingEntriesExec {
	e.body.Settled = &settled
	return e
}

func (e *SearchAccountingEntriesExec) WithSettledAtEnd(date string) *SearchAccountingEntriesExec {
	e.body.SettledAtEnd = date
	return e
}

func (e *SearchAccountingEntriesExec) WithSettledAtStart(date string) *SearchAccountingEntriesExec {
	e.body.SettledAtStart = date
	return e
}

func (e *SearchAccountingEntriesExec) WithStatus(status ...Status) *SearchAccountingEntriesExec {
	e.body.Status = append(e.body.Status, status...)
	return e
}

func (e *SearchAccountingEntriesExec) WithSubFamilies(subFamilies ...SubFamily) *SearchAccountingEntriesExec {
	e.body.SubFamilies = append(e.body.SubFamilies, subFamilies...)
	return e
}

func (e *SearchAccountingEntriesExec) WithTagGroupIDsUnset(ids ...string) *SearchAccountingEntriesExec {
	e.body.TagGroupIDsUnset = append(e.body.TagGroupIDsUnset, ids...)
	return e
}

func (e *SearchAccountingEntriesExec) WithTagRowIDs(ids ...string) *SearchAccountingEntriesExec {
	e.body.TagRowIDs = append(e.body.TagRowIDs, ids...)
	return e
}

func (e *SearchAccountingEntriesExec) WithTaxCodeIDs(ids ...string) *SearchAccountingEntriesExec {
	e.body.TaxCodeIDs = append(e.body.TaxCodeIDs, ids...)
	return e
}

func (e *SearchAccountingEntriesExec) WithTeamIDs(ids ...string) *SearchAccountingEntriesExec {
	e.body.TeamIDs = append(e.body.TeamIDs, ids...)
	return e
}

func (e *SearchAccountingEntriesExec) WithTransactionValueCurrencies(currencies ...string) *SearchAccountingEntriesExec {
	e.body.TransactionValueCurrencies = append(e.body.TransactionValueCurrencies, currencies...)
	return e
}

func (e *SearchAccountingEntriesExec) Execute() (*shared.CursorPageResponse[AccountingEntry], *shared.ResponseExtra, error) {
	queryParams := make(url.Values)
	shared.AddQueryCompanyID(queryParams, e.companyID)
	shared.AddQueryString(queryParams, "organization_id", e.organizationID)
	e.pagingInfo.Apply(queryParams)

	var out shared.CursorPageResponse[AccountingEntry]
	_, response, err := e.config.SendRequest(e.ctx, "POST", shared.URLWithQuery(basePath+":search", queryParams), e.body, &out)
	extra := shared.ResponseExtraFromResponse(response)
	if err != nil {
		return nil, extra, err
	}

	return &out, extra, nil
}
