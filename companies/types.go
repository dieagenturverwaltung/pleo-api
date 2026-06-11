package companies

type PageOrder string

const (
	PageOrderAsc            PageOrder = "ASC"
	PageOrderAscNullsFirst  PageOrder = "ASC_NULLS_FIRST"
	PageOrderAscNullsLast   PageOrder = "ASC_NULLS_LAST"
	PageOrderDesc           PageOrder = "DESC"
	PageOrderDescNullsFirst PageOrder = "DESC_NULLS_FIRST"
	PageOrderDescNullsLast  PageOrder = "DESC_NULLS_LAST"
)

type Company struct {
	Address            Address `json:"address"`
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	OrganizationID     string  `json:"organizationId,omitempty"`
	RegistrationNumber string  `json:"registrationNumber,omitempty"`
}

type Address struct {
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	Country      string `json:"country,omitempty"`
	Locality     string `json:"locality,omitempty"`
	PostalCode   string `json:"postalCode,omitempty"`
	Region       string `json:"region,omitempty"`
}
