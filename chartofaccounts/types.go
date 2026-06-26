package chartofaccounts

type PageOrder string

const (
	PageOrderAsc            PageOrder = "ASC"
	PageOrderAscNullsFirst  PageOrder = "ASC_NULLS_FIRST"
	PageOrderAscNullsLast   PageOrder = "ASC_NULLS_LAST"
	PageOrderDesc           PageOrder = "DESC"
	PageOrderDescNullsFirst PageOrder = "DESC_NULLS_FIRST"
	PageOrderDescNullsLast  PageOrder = "DESC_NULLS_LAST"
)

type BookkeepingMethod string

const (
	BookkeepingMethodNone            BookkeepingMethod = "NONE"
	BookkeepingMethodJournal         BookkeepingMethod = "JOURNAL"
	BookkeepingMethodAccountsPayable BookkeepingMethod = "ACCOUNTS_PAYABLE"
)

type AccountCreateRequest struct {
	Archived          bool           `json:"archived"`
	Code              string         `json:"code,omitempty"`
	CompanyID         string         `json:"companyId"`
	ExternalID        string         `json:"externalId"`
	Metadata          map[string]any `json:"metadata,omitempty"`
	Name              string         `json:"name"`
	TaxCodeExternalID string         `json:"taxCodeExternalId,omitempty"`
}

type AccountUpdateRequest struct {
	Archived          bool           `json:"archived"`
	Code              string         `json:"code,omitempty"`
	Metadata          map[string]any `json:"metadata,omitempty"`
	Name              string         `json:"name"`
	TaxCodeExternalID string         `json:"taxCodeExternalId,omitempty"`
}

type AccountSearchRequest struct {
	Archived                         *bool    `json:"archived,omitempty"`
	Code                             string   `json:"code,omitempty"`
	Codes                            []string `json:"codes,omitempty"`
	CompanyID                        string   `json:"companyId"`
	ExcludeIfAssignedToCategory      bool     `json:"excludeIfAssignedToCategory"`
	ExcludeIfAssignedToContraAccount bool     `json:"excludeIfAssignedToContraAccount"`
	ExternalID                       string   `json:"externalId,omitempty"`
	IDs                              []string `json:"ids,omitempty"`
	Name                             string   `json:"name,omitempty"`
}

type AccountBatchCreateRequest struct {
	CompanyID string                    `json:"companyId"`
	Items     []AccountBatchRequestItem `json:"items"`
}

type AccountBatchRequestItem struct {
	Archived          bool           `json:"archived"`
	Code              string         `json:"code,omitempty"`
	ExternalID        string         `json:"externalId"`
	Metadata          map[string]any `json:"metadata,omitempty"`
	Name              string         `json:"name"`
	TaxCodeExternalID string         `json:"taxCodeExternalId,omitempty"`
}

type Account struct {
	Archived          bool           `json:"archived"`
	Code              string         `json:"code,omitempty"`
	CompanyID         string         `json:"companyId"`
	ExternalID        string         `json:"externalId"`
	ID                string         `json:"id"`
	Metadata          map[string]any `json:"metadata,omitempty"`
	Name              string         `json:"name"`
	TaxCodeExternalID string         `json:"taxCodeExternalId,omitempty"`
}

type AccountBatchCreateResponse struct {
	Created []Account                `json:"created"`
	Failed  []AccountBatchFailedItem `json:"failed"`
}

type AccountBatchFailedItem struct {
	Reasons []string                `json:"reasons"`
	Request AccountBatchRequestItem `json:"request"`
}
