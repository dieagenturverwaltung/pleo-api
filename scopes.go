package pleo_api

const (
	ScopeExportItemsRead = "export-items:read"
	ScopeExportJobsRead  = "export-jobs:read"
	ScopeExportJobsWrite = "export-jobs:write"
	ScopeTagGroupsRead   = "tag-groups:read"
	ScopeTagGroupsWrite  = "tag-groups:write"
	ScopeTaxCodesRead    = "tax-codes:read"
	ScopeTaxCodesWrite   = "tax-codes:write"
)

var AllScopes = []string{
	ScopeExportItemsRead,
	ScopeExportJobsRead,
	ScopeExportJobsWrite,
	ScopeTagGroupsRead,
	ScopeTagGroupsWrite,
	ScopeTaxCodesRead,
	ScopeTaxCodesWrite,
}
