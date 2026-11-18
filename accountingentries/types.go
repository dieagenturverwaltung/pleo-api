package accountingentries

import "encoding/json"

type PageOrder string

const (
	PageOrderAsc            PageOrder = "ASC"
	PageOrderAscNullsFirst  PageOrder = "ASC_NULLS_FIRST"
	PageOrderAscNullsLast   PageOrder = "ASC_NULLS_LAST"
	PageOrderDesc           PageOrder = "DESC"
	PageOrderDescNullsFirst PageOrder = "DESC_NULLS_FIRST"
	PageOrderDescNullsLast  PageOrder = "DESC_NULLS_LAST"
)

type Family string

const (
	FamilyBalanceAmendment   Family = "BALANCE_AMENDMENT"
	FamilyBillInvoice        Family = "BILL_INVOICE"
	FamilyBillInvoicePayment Family = "BILL_INVOICE_PAYMENT"
	FamilyCardPurchase       Family = "CARD_PURCHASE"
	FamilyCashback           Family = "CASHBACK"
	FamilyCreditNote         Family = "CREDIT_NOTE"
	FamilyMileage            Family = "MILEAGE"
	FamilyOutOfPocket        Family = "OUT_OF_POCKET"
	FamilyOverdraft          Family = "OVERDRAFT"
	FamilyPerDiem            Family = "PER_DIEM"
	FamilyPleoInvoice        Family = "PLEO_INVOICE"
	FamilyReimbursement      Family = "REIMBURSEMENT"
	FamilyWallet             Family = "WALLET"
)

// SubFamily is optional, and Pleo may add values beyond the constants below.
type SubFamily string

const (
	SubFamilyATM                           SubFamily = "ATM"
	SubFamilyInStore                       SubFamily = "IN_STORE"
	SubFamilyOnline                        SubFamily = "ONLINE"
	SubFamilyPleoInvoice                   SubFamily = "PLEO_INVOICE"
	SubFamilyPleoInvoiceRefund             SubFamily = "PLEO_INVOICE_REFUND"
	SubFamilyReimbursementToCompany        SubFamily = "REIMBURSEMENT_TO_COMPANY"
	SubFamilyReimbursementToEmployee       SubFamily = "REIMBURSEMENT_TO_EMPLOYEE"
	SubFamilyReimbursementToEmployeeRefund SubFamily = "REIMBURSEMENT_TO_EMPLOYEE_REFUND"
	SubFamilyLoad                          SubFamily = "LOAD"
	SubFamilyUnload                        SubFamily = "UNLOAD"
	SubFamilyWithdrawal                    SubFamily = "WITHDRAWAL"
	SubFamilyFee                           SubFamily = "FEE"
	SubFamilyInterest                      SubFamily = "INTEREST"
	SubFamilyBillInvoicePaymentReturned    SubFamily = "BILL_INVOICE_PAYMENT_RETURNED"
	SubFamilyBillInvoicePaymentReceived    SubFamily = "BILL_INVOICE_PAYMENT_RECEIVED"
)

type Status string

const (
	StatusDraft               Status = "DRAFT"
	StatusPending             Status = "PENDING"
	StatusCompleted           Status = "COMPLETED"
	StatusCompletedExternally Status = "COMPLETED_EXTERNALLY"
	StatusCancelled           Status = "CANCELLED"
	StatusRejected            Status = "REJECTED"
	StatusError               Status = "ERROR"
)

type ExportStatus string

const (
	ExportStatusCreated     ExportStatus = "CREATED"
	ExportStatusExported    ExportStatus = "EXPORTED"
	ExportStatusExporting   ExportStatus = "EXPORTING"
	ExportStatusNotCreated  ExportStatus = "NOT_CREATED"
	ExportStatusNotExported ExportStatus = "NOT_EXPORTED"
	ExportStatusQueued      ExportStatus = "QUEUED"
)

type ReviewStatus string

const (
	ReviewStatusNotOK                     ReviewStatus = "NOT_OK"
	ReviewStatusNotRequired               ReviewStatus = "NOT_REQUIRED"
	ReviewStatusNotSet                    ReviewStatus = "NOT_SET"
	ReviewStatusOK                        ReviewStatus = "OK"
	ReviewStatusWaitingForExpenseOwner    ReviewStatus = "WAITING_FOR_EXPENSE_OWNER"
	ReviewStatusWaitingForReviewer        ReviewStatus = "WAITING_FOR_REVIEWER"
	ReviewStatusWaitingForCompanyReviewer ReviewStatus = "WAITING_FOR_COMPANY_REVIEWER"
	ReviewStatusWaitingForTeamReviewer    ReviewStatus = "WAITING_FOR_TEAM_REVIEWER"
	ReviewStatusWaitingForTagReviewer     ReviewStatus = "WAITING_FOR_TAG_REVIEWER"
)

type ReceiptFileType string

const (
	ReceiptFileTypeEmbed     ReceiptFileType = "EMBED"
	ReceiptFileTypeOriginal  ReceiptFileType = "ORIGINAL"
	ReceiptFileTypeThumbnail ReceiptFileType = "THUMBNAIL"
	ReceiptFileTypePreview   ReceiptFileType = "PREVIEW"
)

type ReceiptMIMEType string

const (
	ReceiptMIMETypeUnknown    ReceiptMIMEType = "unknown"
	ReceiptMIMETypePDF        ReceiptMIMEType = "application/pdf"
	ReceiptMIMETypePostScript ReceiptMIMEType = "application/postscript"
	ReceiptMIMETypeXML        ReceiptMIMEType = "application/xml"
	ReceiptMIMETypeJPEG       ReceiptMIMEType = "image/jpeg"
	ReceiptMIMETypePNG        ReceiptMIMEType = "image/png"
	ReceiptMIMETypeHEIC       ReceiptMIMEType = "image/heic"
	ReceiptMIMETypeHEIF       ReceiptMIMEType = "image/heif"
	ReceiptMIMETypeGIF        ReceiptMIMEType = "image/gif"
	ReceiptMIMETypeTextXML    ReceiptMIMEType = "text/xml"
)

type ReceiptSource string

const (
	ReceiptSourceAppShare             ReceiptSource = "APP_SHARE"
	ReceiptSourceBolt                 ReceiptSource = "BOLT"
	ReceiptSourceCamera               ReceiptSource = "CAMERA"
	ReceiptSourceCashback             ReceiptSource = "CASHBACK"
	ReceiptSourceEmailFetching        ReceiptSource = "EMAIL_FETCHING"
	ReceiptSourceEmailFetchingBatched ReceiptSource = "EMAIL_FETCHING_BATCHED"
	ReceiptSourceEmailForwarding      ReceiptSource = "EMAIL_FORWARDING"
	ReceiptSourceGeneratedReceipt     ReceiptSource = "GENERATED_RECEIPT"
	ReceiptSourceLostReceipt          ReceiptSource = "LOST_RECEIPT"
	ReceiptSourceRoger                ReceiptSource = "ROGER"
	ReceiptSourceStorebox             ReceiptSource = "STOREBOX"
	ReceiptSourceStripe               ReceiptSource = "STRIPE"
	ReceiptSourceTravelperk           ReceiptSource = "TRAVELPERK"
	ReceiptSourceUser                 ReceiptSource = "USER"
)

type SearchAccountingEntriesRequest struct {
	IncludeDeleted             bool           `json:"includeDeleted"`
	AccountIDs                 []string       `json:"accountIds,omitempty"`
	BookkeepingDateEnd         string         `json:"bookkeepingDateEnd,omitempty"`
	BookkeepingDateStart       string         `json:"bookkeepingDateStart,omitempty"`
	EmployeeIDs                []string       `json:"employeeIds,omitempty"`
	ExportStatus               []ExportStatus `json:"exportStatus,omitempty"`
	Families                   []Family       `json:"families,omitempty"`
	ForeignExtensionIDs        []string       `json:"foreignExtensionIds,omitempty"`
	HasAccount                 *bool          `json:"hasAccount,omitempty"`
	HasBookkeepingDate         *bool          `json:"hasBookkeepingDate,omitempty"`
	HasNote                    *bool          `json:"hasNote,omitempty"`
	PerformedAtEnd             string         `json:"performedAtEnd,omitempty"`
	PerformedAtStart           string         `json:"performedAtStart,omitempty"`
	ReviewStatus               []ReviewStatus `json:"reviewStatus,omitempty"`
	Settled                    *bool          `json:"settled,omitempty"`
	SettledAtEnd               string         `json:"settledAtEnd,omitempty"`
	SettledAtStart             string         `json:"settledAtStart,omitempty"`
	Status                     []Status       `json:"status,omitempty"`
	SubFamilies                []SubFamily    `json:"subFamilies,omitempty"`
	TagGroupIDsUnset           []string       `json:"tagGroupIdsUnset,omitempty"`
	TagRowIDs                  []string       `json:"tagRowIds,omitempty"`
	TaxCodeIDs                 []string       `json:"taxCodeIds,omitempty"`
	TeamIDs                    []string       `json:"teamIds,omitempty"`
	TransactionValueCurrencies []string       `json:"transactionValueCurrencies,omitempty"`
}

type SearchReceiptsRequest struct {
	IncludeDeleted     bool              `json:"includeDeleted"`
	AccountingEntryIDs []string          `json:"accountingEntryIds,omitempty"`
	EmployeeIDs        []string          `json:"employeeIds,omitempty"`
	FileType           *ReceiptFileType  `json:"fileType,omitempty"`
	MIMEType           []ReceiptMIMEType `json:"mimeType,omitempty"`
	Source             []ReceiptSource   `json:"source,omitempty"`
}

type AccountingEntry struct {
	Attendees                []Attendee        `json:"attendees"`
	CompanyID                string            `json:"companyId"`
	CreatedAt                string            `json:"createdAt"`
	ExportStatus             ExportStatus      `json:"exportStatus"`
	Family                   Family            `json:"family"`
	ID                       string            `json:"id"`
	PerformedAt              string            `json:"performedAt"`
	ReceiptIDs               []string          `json:"receiptIds"`
	ReviewStatus             ReviewStatus      `json:"reviewStatus"`
	SplitItems               []SplitItem       `json:"splitItems"`
	Status                   Status            `json:"status"`
	Tags                     []Tag             `json:"tags"`
	TransactionValue         Money             `json:"transactionValue"`
	UpdatedAt                string            `json:"updatedAt"`
	AccountCode              string            `json:"accountCode,omitempty"`
	AccountID                string            `json:"accountId,omitempty"`
	BookkeepingDate          string            `json:"bookkeepingDate,omitempty"`
	DeletedAt                string            `json:"deletedAt,omitempty"`
	EmployeeID               string            `json:"employeeId,omitempty"`
	ForeignExtensionID       string            `json:"foreignExtensionId,omitempty"`
	Merchant                 *Merchant         `json:"merchant,omitempty"`
	Note                     string            `json:"note,omitempty"`
	RelatedAccountingEntries []json.RawMessage `json:"relatedAccountingEntries,omitempty"`
	SettledAt                string            `json:"settledAt,omitempty"`
	SubFamily                *SubFamily        `json:"subFamily,omitempty"`
	Supplier                 *Supplier         `json:"supplier,omitempty"`
	TaxCodeID                string            `json:"taxCodeId,omitempty"`
	TeamID                   string            `json:"teamId,omitempty"`
	TotalBillValue           *Money            `json:"totalBillValue,omitempty"`
}

// Money holds a monetary value in minor currency units, without floating-point rounding.
type Money struct {
	Currency string `json:"currency"`
	Minors   int64  `json:"minors"`
}

type Attendee struct {
	AttendableID   string `json:"attendableId"`
	AttendableType string `json:"attendableType"`
	ID             string `json:"id"`
}

type SplitItem struct {
	CreatedAt        string `json:"createdAt"`
	ID               string `json:"id"`
	Tags             []Tag  `json:"tags"`
	UpdatedAt        string `json:"updatedAt"`
	AccountCode      string `json:"accountCode,omitempty"`
	AccountID        string `json:"accountId,omitempty"`
	TaxCodeID        string `json:"taxCodeId,omitempty"`
	TotalBillValue   *Money `json:"totalBillValue,omitempty"`
	TransactionValue *Money `json:"transactionValue,omitempty"`
}

type Tag struct {
	TagGroupID string `json:"tagGroupId"`
	TagID      string `json:"tagId"`
}

type Merchant struct {
	Address              map[string]any `json:"address"`
	ExternalMerchantID   string         `json:"externalMerchantId,omitempty"`
	MerchantCategoryCode string         `json:"merchantCategoryCode,omitempty"`
	Name                 string         `json:"name,omitempty"`
}

type Supplier struct {
	CIF                   string `json:"cif,omitempty"`
	DocumentNumber        string `json:"documentNumber,omitempty"`
	DocumentType          string `json:"documentType,omitempty"`
	ExternalAccountNumber string `json:"externalAccountNumber,omitempty"`
	ExternalSupplierID    string `json:"externalSupplierId,omitempty"`
}

// Receipt includes a time-bound URL to the requested representation of the file.
type Receipt struct {
	AccountingEntryID string          `json:"accountingEntryId"`
	CreatedAt         string          `json:"createdAt"`
	FileType          ReceiptFileType `json:"fileType"`
	ID                string          `json:"id"`
	MIMEType          ReceiptMIMEType `json:"mimeType"`
	SizeInBytes       int64           `json:"sizeInBytes"`
	Source            ReceiptSource   `json:"source"`
	URL               string          `json:"url"`
	OCRDocumentID     string          `json:"ocrDocumentId,omitempty"`
}
