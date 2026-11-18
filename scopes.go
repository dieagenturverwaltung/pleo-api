package pleo_api

const (
	ScopeAccountingEntriesRead = "accounting-entries:read"
	ScopeCompaniesRead         = "companies:read"
	ScopeExportItemsRead       = "export-items:read"
	ScopeExportJobsRead        = "export-jobs:read"
	ScopeExportJobsWrite       = "export-jobs:write"
	ScopeTagGroupsRead         = "tag-groups:read"
	ScopeTagGroupsWrite        = "tag-groups:write"
	ScopeTagGroupRead          = "tag-group:read"
	ScopeTagGroupWrite         = "tag-group:write"
	ScopeTaxCodesRead          = "tax-codes:read"
	ScopeTaxCodesWrite         = "tax-codes:write"
	ScopeAccountsRead          = "accounts:read"
	ScopeAccountsWrite         = "accounts:write"
)

var AllScopes = []string{
	ScopeAccountingEntriesRead,
	ScopeExportItemsRead,
	ScopeExportJobsRead,
	ScopeExportJobsWrite,
	ScopeTagGroupsRead,
	ScopeTagGroupsWrite,
	ScopeTaxCodesRead,
	ScopeTaxCodesWrite,
	ScopeAccountsRead,
	ScopeAccountsWrite,
}

var AllScopesProd = []string{
	ScopeAccountingEntriesRead,
	ScopeExportItemsRead,
	ScopeExportJobsRead,
	ScopeExportJobsWrite,
	ScopeTagGroupRead,
	ScopeTagGroupWrite,
	ScopeTaxCodesRead,
	ScopeTaxCodesWrite,
	ScopeAccountsRead,
	ScopeAccountsWrite,
}
