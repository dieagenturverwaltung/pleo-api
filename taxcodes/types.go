package taxcodes

type PageOrder string

const (
	PageOrderAsc            PageOrder = "ASC"
	PageOrderAscNullsFirst  PageOrder = "ASC_NULLS_FIRST"
	PageOrderAscNullsLast   PageOrder = "ASC_NULLS_LAST"
	PageOrderDesc           PageOrder = "DESC"
	PageOrderDescNullsFirst PageOrder = "DESC_NULLS_FIRST"
	PageOrderDescNullsLast  PageOrder = "DESC_NULLS_LAST"
)

type TaxCodeType string

const (
	TaxCodeTypeInclusive TaxCodeType = "inclusive"
	TaxCodeTypeExclusive TaxCodeType = "exclusive"
	TaxCodeTypeReverse   TaxCodeType = "reverse"
)

type TaxCodeCreateUpdateModel struct {
	Archived           bool        `json:"archived"`
	Code               string      `json:"code,omitempty"`
	CompanyID          string      `json:"companyId"`
	IngoingTaxAccount  string      `json:"ingoingTaxAccount,omitempty"`
	Name               string      `json:"name"`
	OutgoingTaxAccount string      `json:"outgoingTaxAccount,omitempty"`
	Rate               float64     `json:"rate"`
	Type               TaxCodeType `json:"type"`
}

type TaxCodeSearchRequest struct {
	Archived   *bool       `json:"archived,omitempty"`
	TaxCodeIDs []string    `json:"taxCodeIds,omitempty"`
	Type       TaxCodeType `json:"type,omitempty"`
}

type TaxCodeModel struct {
	AccountingIntegrationSystem string      `json:"accountingIntegrationSystem,omitempty"`
	Archived                    bool        `json:"archived"`
	Code                        string      `json:"code,omitempty"`
	CompanyID                   string      `json:"companyId"`
	CreatedAt                   string      `json:"createdAt"`
	ID                          string      `json:"id"`
	IngoingTaxAccount           string      `json:"ingoingTaxAccount,omitempty"`
	Name                        string      `json:"name"`
	OutgoingTaxAccount          string      `json:"outgoingTaxAccount,omitempty"`
	Rate                        float64     `json:"rate"`
	Type                        TaxCodeType `json:"type"`
	UpdatedAt                   string      `json:"updatedAt"`
}
