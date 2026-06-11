package exportapi

type PageOrder string

const (
	PageOrderAsc            PageOrder = "ASC"
	PageOrderAscNullsFirst  PageOrder = "ASC_NULLS_FIRST"
	PageOrderAscNullsLast   PageOrder = "ASC_NULLS_LAST"
	PageOrderDesc           PageOrder = "DESC"
	PageOrderDescNullsFirst PageOrder = "DESC_NULLS_FIRST"
	PageOrderDescNullsLast  PageOrder = "DESC_NULLS_LAST"
)

type ExportJobStatus string

const (
	ExportJobStatusPending             ExportJobStatus = "pending"
	ExportJobStatusInProgress          ExportJobStatus = "in_progress"
	ExportJobStatusFailed              ExportJobStatus = "failed"
	ExportJobStatusCompleted           ExportJobStatus = "completed"
	ExportJobStatusCompletedWithErrors ExportJobStatus = "completed_with_errors"
)

type ExportJobEventType string

const (
	ExportJobEventStarted             ExportJobEventType = "started"
	ExportJobEventFailed              ExportJobEventType = "failed"
	ExportJobEventCompleted           ExportJobEventType = "completed"
	ExportJobEventCompletedWithErrors ExportJobEventType = "completed_with_errors"
)

type ExportJobFailureReasonType string

const (
	ExportJobFailureInvalidConfiguration           ExportJobFailureReasonType = "invalid_configuration"
	ExportJobFailureMissingConfiguration           ExportJobFailureReasonType = "missing_configuration"
	ExportJobFailureAuthenticationFailure          ExportJobFailureReasonType = "authentication_failure"
	ExportJobFailureAccountingSystemAuthentication ExportJobFailureReasonType = "accounting_system_authentication_failure"
	ExportJobFailurePleoRateLimit                  ExportJobFailureReasonType = "pleo_rate_limit"
	ExportJobFailureAccountingSystemRateLimit      ExportJobFailureReasonType = "accounting_system_rate_limit"
	ExportJobFailureServiceUnreachable             ExportJobFailureReasonType = "service_unreachable"
	ExportJobFailureAccountingSystemUnreachable    ExportJobFailureReasonType = "accounting_system_unreachable"
	ExportJobFailureValidationFailure              ExportJobFailureReasonType = "validation_failure"
	ExportJobFailureAuthorizationFailure           ExportJobFailureReasonType = "authorization_failure"
	ExportJobFailureIntegrationUnusable            ExportJobFailureReasonType = "integration_unusable"
	ExportJobFailureJobExpired                     ExportJobFailureReasonType = "job_expired"
	ExportJobFailureServiceTimeout                 ExportJobFailureReasonType = "service_timeout"
	ExportJobFailureAccountingSystemTimeout        ExportJobFailureReasonType = "accounting_system_timeout"
	ExportJobFailureCanceledByUser                 ExportJobFailureReasonType = "canceled_by_user"
)

type ExportJobItemStatus string

const (
	ExportJobItemStatusPending    ExportJobItemStatus = "pending"
	ExportJobItemStatusInProgress ExportJobItemStatus = "in_progress"
	ExportJobItemStatusFailed     ExportJobItemStatus = "failed"
	ExportJobItemStatusSuccessful ExportJobItemStatus = "successful"
	ExportJobItemStatusAbandoned  ExportJobItemStatus = "abandoned"
)

type ExportJobItemFailureReasonType string

const (
	ExportJobItemFailureReceiptUpload            ExportJobItemFailureReasonType = "receipt_upload_failure"
	ExportJobItemFailureReceiptDownload          ExportJobItemFailureReasonType = "receipt_download_failure"
	ExportJobItemFailureReceiptFileSizeLimit     ExportJobItemFailureReasonType = "receipt_file_size_limit_exceeded"
	ExportJobItemFailureReceiptConversion        ExportJobItemFailureReasonType = "receipt_conversion_failure"
	ExportJobItemFailureUnexpected               ExportJobItemFailureReasonType = "unexpected_failure"
	ExportJobItemFailureInvalidConfiguration     ExportJobItemFailureReasonType = "invalid_configuration"
	ExportJobItemFailureMissingConfiguration     ExportJobItemFailureReasonType = "missing_configuration"
	ExportJobItemFailureAccountingAuthentication ExportJobItemFailureReasonType = "accounting_system_authentication_failure"
	ExportJobItemFailureAccountingRateLimit      ExportJobItemFailureReasonType = "accounting_system_rate_limit"
	ExportJobItemFailureAccountingUnreachable    ExportJobItemFailureReasonType = "accounting_system_unreachable"
	ExportJobItemFailureValidation               ExportJobItemFailureReasonType = "validation_failure"
	ExportJobItemFailureAccountingValidation     ExportJobItemFailureReasonType = "accounting_system_validation_failure"
	ExportJobItemFailureAuthorization            ExportJobItemFailureReasonType = "authorization_failure"
	ExportJobItemFailureAccountingTimeout        ExportJobItemFailureReasonType = "accounting_system_timeout"
	ExportJobItemFailureVendorUnknown            ExportJobItemFailureReasonType = "vendor_unknown"
	ExportJobItemFailureAccountUnknown           ExportJobItemFailureReasonType = "account_unknown"
	ExportJobItemFailureTaxCodeUnknown           ExportJobItemFailureReasonType = "tax_code_unknown"
	ExportJobItemFailureTagUnknown               ExportJobItemFailureReasonType = "tag_unknown"
)

type CreateExportJobRequest struct {
	AccountingEntryIDs []string `json:"accountingEntryIds"`
	CompanyID          string   `json:"companyId"`
	EmployeeID         *string  `json:"employeeId,omitempty"`
	IsInteractive      *bool    `json:"isInteractive,omitempty"`
	Options            *Options `json:"options,omitempty"`
}

type Options struct {
	VendorBasedBookkeeping *bool `json:"vendorBasedBookkeeping,omitempty"`
}

type CreateExportJobEventRequest struct {
	Event             ExportJobEventType          `json:"event"`
	FailureReason     string                      `json:"failureReason,omitempty"`
	FailureReasonType *ExportJobFailureReasonType `json:"failureReasonType,omitempty"`
	JobID             string                      `json:"jobId"`
}

type SearchExportJobItemsRequest struct {
	AccountingEntryIDs []string             `json:"accountingEntryIds"`
	Status             *ExportJobItemStatus `json:"status,omitempty"`
}

type UpdateExportJobItem struct {
	AccountingEntryID string                          `json:"accountingEntryId"`
	ExportedAt        string                          `json:"exportedAt,omitempty"`
	ExternalID        string                          `json:"externalId,omitempty"`
	ExternalURL       string                          `json:"externalUrl,omitempty"`
	FailureReason     string                          `json:"failureReason,omitempty"`
	FailureReasonType *ExportJobItemFailureReasonType `json:"failureReasonType,omitempty"`
	Status            ExportJobItemStatus             `json:"status"`
}

type ExportJob struct {
	CompanyID              string `json:"companyId"`
	CompletedAt            string `json:"completedAt"`
	CreatedAt              string `json:"createdAt"`
	CreatedBy              string `json:"createdBy"`
	ExpiredAt              string `json:"expiredAt"`
	ExpiresIn              int    `json:"expiresIn"`
	FailureReason          string `json:"failureReason"`
	FailureReasonType      string `json:"failureReasonType"`
	ID                     string `json:"id"`
	IsInteractive          bool   `json:"isInteractive"`
	LastUpdatedAt          string `json:"lastUpdatedAt"`
	NumberOfItems          int    `json:"numberOfItems"`
	StartedAt              string `json:"startedAt"`
	Status                 string `json:"status"`
	VendorBasedBookkeeping bool   `json:"vendorBasedBookkeeping"`
}

type ExportJobItem struct {
	AccountingEntryID string `json:"accountingEntryId"`
	ExportJobID       string `json:"exportJobId"`
	ExportedAt        string `json:"exportedAt"`
	ExternalID        string `json:"externalId"`
	ExternalURL       string `json:"externalUrl"`
	FailureReason     string `json:"failureReason"`
	FailureReasonType string `json:"failureReasonType"`
	Status            string `json:"status"`
}

type ExportJobItemUpdate struct {
	Data   []ExportJobItem            `json:"data"`
	Errors []ExportJobItemUpdateError `json:"errors"`
}

type ExportJobItemUpdateError struct {
	AccountingEntryID string `json:"accountingEntryId"`
	Message           string `json:"message"`
	Type              string `json:"type"`
}

type ExceptionInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

type ExportItem struct {
	Links                 LinksResponse                   `json:"_links"`
	AccountingEntryID     string                          `json:"accountingEntryId"`
	AccountingEntryLines  []ExportItemLine                `json:"accountingEntryLines"`
	AdditionalInformation ExportItemAdditionalInformation `json:"additionalInformation"`
	Amount                ExportItemAmount                `json:"amount"`
	Bookkeeping           *Bookkeeping                    `json:"bookkeeping,omitempty"`
	CompanyID             string                          `json:"companyId"`
	ContraAccount         *ContraAccount                  `json:"contraAccount,omitempty"`
	Date                  string                          `json:"date"`
	Files                 []ExportItemFile                `json:"files,omitempty"`
	Note                  string                          `json:"note,omitempty"`
	ServicePeriod         *ServicePeriod                  `json:"servicePeriod,omitempty"`
	SubType               string                          `json:"subType,omitempty"`
	Supplier              *Supplier                       `json:"supplier,omitempty"`
	Team                  *Team                           `json:"team,omitempty"`
	Type                  string                          `json:"type"`
	User                  *ExportItemUser                 `json:"user,omitempty"`
	Vendor                *Vendor                         `json:"vendor,omitempty"`
}

type LinksResponse struct {
	API    map[string]string `json:"api,omitempty"`
	Mobile map[string]string `json:"mobile,omitempty"`
	Web    map[string]string `json:"web,omitempty"`
}

type ExportItemLine struct {
	Account               *Account         `json:"account,omitempty"`
	AccountingEntryLineID string           `json:"accountingEntryLineId"`
	LineAmount            ExportItemAmount `json:"lineAmount"`
	NetAmount             ExportItemAmount `json:"netAmount"`
	Tags                  []Tag            `json:"tags,omitempty"`
	Tax                   *Tax             `json:"tax,omitempty"`
}

type ExportItemAdditionalInformation struct {
	Attendees          []string            `json:"attendees,omitempty"`
	InvoiceInformation *InvoiceInformation `json:"invoiceInformation,omitempty"`
	ReconciledEntries  []string            `json:"reconciledEntries,omitempty"`
	ReconciliationID   string              `json:"reconciliationId"`
}

type InvoiceInformation struct {
	DueDate             string               `json:"dueDate,omitempty"`
	InvoiceDate         string               `json:"invoiceDate"`
	InvoiceNumber       string               `json:"invoiceNumber"`
	PaymentDate         string               `json:"paymentDate,omitempty"`
	Status              string               `json:"status"`
	SupplierBankAccount *SupplierBankAccount `json:"supplierBankAccount,omitempty"`
}

type ExportItemAmount struct {
	InSupplierCurrency Money `json:"inSupplierCurrency"`
	InWalletCurrency   Money `json:"inWalletCurrency"`
}

type Money struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}

type ExportItemFile struct {
	Size int    `json:"size"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type ServicePeriod struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Account struct {
	Code       string `json:"code,omitempty"`
	ID         string `json:"id"`
	Identifier string `json:"identifier,omitempty"`
	Name       string `json:"name"`
}

type ContraAccount struct {
	Code       string `json:"code,omitempty"`
	ID         string `json:"id"`
	Identifier string `json:"identifier,omitempty"`
	Name       string `json:"name"`
}

type Bookkeeping struct {
	Method string `json:"method"`
}

type Supplier struct {
	Account       string `json:"account,omitempty"`
	CategoryCode  string `json:"categoryCode,omitempty"`
	Code          string `json:"code,omitempty"`
	Country       string `json:"country,omitempty"`
	Name          string `json:"name,omitempty"`
	TaxIdentifier string `json:"taxIdentifier,omitempty"`
}

type SupplierBankAccount struct {
	AccountNumber string `json:"accountNumber,omitempty"`
	BankCode      string `json:"bankCode,omitempty"`
	BankName      string `json:"bankName"`
	BIC           string `json:"bic,omitempty"`
	Country       string `json:"country"`
	IBAN          string `json:"iban,omitempty"`
}

type Team struct {
	Code string `json:"code,omitempty"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	Code      string `json:"code"`
	GroupCode string `json:"groupCode"`
	ID        string `json:"id"`
}

type Tax struct {
	Amount             ExportItemAmount `json:"amount"`
	Code               string           `json:"code,omitempty"`
	ID                 string           `json:"id"`
	IngoingTaxAccount  string           `json:"ingoingTaxAccount,omitempty"`
	OutgoingTaxAccount string           `json:"outgoingTaxAccount,omitempty"`
	Rate               float64          `json:"rate"`
	Type               string           `json:"type"`
}

type ExportItemUser struct {
	Code string `json:"code,omitempty"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Vendor struct {
	Code                  string `json:"code,omitempty"`
	Country               string `json:"country,omitempty"`
	DefaultCurrency       string `json:"defaultCurrency"`
	ExternalID            string `json:"externalId,omitempty"`
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	RegistrationNumber    string `json:"registrationNumber,omitempty"`
	TaxRegistrationNumber string `json:"taxRegistrationNumber,omitempty"`
}
