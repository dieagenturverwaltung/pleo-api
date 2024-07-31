/*
Export API

Export OpenAPI definitions

API version: 20.0.0
Contact: apiteam@pleo.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package export

import (
	"encoding/json"
	"time"
)

// ExportItem struct for ExportItem
type ExportItem struct {
	Links LinksResponse `json:"_links"`
	// This is the Pleo internal identifier of the export item
	AccountingEntryId string `json:"accountingEntryId"`
	// Accounting entry broken down in entry lines. There will always be an accounting entry line. If the accounting entry has been split into separate lines, then each line would represent the details of that accounting entry line, else there would only be one line present, representing the whole entry.
	AccountingEntryLines []ExportItemLine `json:"accountingEntryLines"`
	Amount               ExportItemAmount `json:"amount"`
	// Pleo company identifier this export item belongs to
	CompanyId string `json:"companyId"`
	// Date the accounting entry should be bookkept
	Date time.Time `json:"date"`
	// Files that have been attached to this accounting entry
	Files    []ExportItemFile   `json:"files,omitempty"`
	Metadata ExportItemMetadata `json:"metadata"`
	// Additional comments potentially describing the accounting entry
	Note     NullableString   `json:"note,omitempty"`
	Supplier NullableSupplier `json:"supplier,omitempty"`
	// Team code is an identifier assigned to an expense to categorize it or associate it with a specific team
	TeamCode NullableString `json:"teamCode,omitempty"`
	// This is the Pleo internal identifier of the accounting entry this export item is associated with
	Type string         `json:"type"`
	User ExportItemUser `json:"user"`
}

// NewExportItem instantiates a new ExportItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportItem(links LinksResponse, accountingEntryId string, accountingEntryLines []ExportItemLine, amount ExportItemAmount, companyId string, date time.Time, metadata ExportItemMetadata, type_ string, user ExportItemUser) *ExportItem {
	this := ExportItem{}
	this.Links = links
	this.AccountingEntryId = accountingEntryId
	this.AccountingEntryLines = accountingEntryLines
	this.Amount = amount
	this.CompanyId = companyId
	this.Date = date
	this.Metadata = metadata
	this.Type = type_
	this.User = user
	return &this
}

// NewExportItemWithDefaults instantiates a new ExportItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportItemWithDefaults() *ExportItem {
	this := ExportItem{}
	return &this
}

// GetLinks returns the Links field value
func (o *ExportItem) GetLinks() LinksResponse {
	if o == nil {
		var ret LinksResponse
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ExportItem) GetLinksOk() (*LinksResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ExportItem) SetLinks(v LinksResponse) {
	o.Links = v
}

// GetAccountingEntryId returns the AccountingEntryId field value
func (o *ExportItem) GetAccountingEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountingEntryId
}

// GetAccountingEntryIdOk returns a tuple with the AccountingEntryId field value
// and a boolean to check if the value has been set.
func (o *ExportItem) GetAccountingEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountingEntryId, true
}

// SetAccountingEntryId sets field value
func (o *ExportItem) SetAccountingEntryId(v string) {
	o.AccountingEntryId = v
}

// GetAccountingEntryLines returns the AccountingEntryLines field value
func (o *ExportItem) GetAccountingEntryLines() []ExportItemLine {
	if o == nil {
		var ret []ExportItemLine
		return ret
	}

	return o.AccountingEntryLines
}

// GetAccountingEntryLinesOk returns a tuple with the AccountingEntryLines field value
// and a boolean to check if the value has been set.
func (o *ExportItem) GetAccountingEntryLinesOk() ([]ExportItemLine, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountingEntryLines, true
}

// SetAccountingEntryLines sets field value
func (o *ExportItem) SetAccountingEntryLines(v []ExportItemLine) {
	o.AccountingEntryLines = v
}

// GetAmount returns the Amount field value
func (o *ExportItem) GetAmount() ExportItemAmount {
	if o == nil {
		var ret ExportItemAmount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ExportItem) GetAmountOk() (*ExportItemAmount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ExportItem) SetAmount(v ExportItemAmount) {
	o.Amount = v
}

// GetCompanyId returns the CompanyId field value
func (o *ExportItem) GetCompanyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *ExportItem) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *ExportItem) SetCompanyId(v string) {
	o.CompanyId = v
}

// GetDate returns the Date field value
func (o *ExportItem) GetDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *ExportItem) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *ExportItem) SetDate(v time.Time) {
	o.Date = v
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportItem) GetFiles() []ExportItemFile {
	if o == nil {
		var ret []ExportItemFile
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportItem) GetFilesOk() ([]ExportItemFile, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *ExportItem) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []ExportItemFile and assigns it to the Files field.
func (o *ExportItem) SetFiles(v []ExportItemFile) {
	o.Files = v
}

// GetMetadata returns the Metadata field value
func (o *ExportItem) GetMetadata() ExportItemMetadata {
	if o == nil {
		var ret ExportItemMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ExportItem) GetMetadataOk() (*ExportItemMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ExportItem) SetMetadata(v ExportItemMetadata) {
	o.Metadata = v
}

// GetNote returns the Note field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportItem) GetNote() string {
	if o == nil || o.Note.Get() == nil {
		var ret string
		return ret
	}
	return *o.Note.Get()
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportItem) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Note.Get(), o.Note.IsSet()
}

// HasNote returns a boolean if a field has been set.
func (o *ExportItem) HasNote() bool {
	if o != nil && o.Note.IsSet() {
		return true
	}

	return false
}

// SetNote gets a reference to the given NullableString and assigns it to the Note field.
func (o *ExportItem) SetNote(v string) {
	o.Note.Set(&v)
}

// SetNoteNil sets the value for Note to be an explicit nil
func (o *ExportItem) SetNoteNil() {
	o.Note.Set(nil)
}

// UnsetNote ensures that no value is present for Note, not even an explicit nil
func (o *ExportItem) UnsetNote() {
	o.Note.Unset()
}

// GetSupplier returns the Supplier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportItem) GetSupplier() Supplier {
	if o == nil || o.Supplier.Get() == nil {
		var ret Supplier
		return ret
	}
	return *o.Supplier.Get()
}

// GetSupplierOk returns a tuple with the Supplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportItem) GetSupplierOk() (*Supplier, bool) {
	if o == nil {
		return nil, false
	}
	return o.Supplier.Get(), o.Supplier.IsSet()
}

// HasSupplier returns a boolean if a field has been set.
func (o *ExportItem) HasSupplier() bool {
	if o != nil && o.Supplier.IsSet() {
		return true
	}

	return false
}

// SetSupplier gets a reference to the given NullableSupplier and assigns it to the Supplier field.
func (o *ExportItem) SetSupplier(v Supplier) {
	o.Supplier.Set(&v)
}

// SetSupplierNil sets the value for Supplier to be an explicit nil
func (o *ExportItem) SetSupplierNil() {
	o.Supplier.Set(nil)
}

// UnsetSupplier ensures that no value is present for Supplier, not even an explicit nil
func (o *ExportItem) UnsetSupplier() {
	o.Supplier.Unset()
}

// GetTeamCode returns the TeamCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportItem) GetTeamCode() string {
	if o == nil || o.TeamCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.TeamCode.Get()
}

// GetTeamCodeOk returns a tuple with the TeamCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportItem) GetTeamCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TeamCode.Get(), o.TeamCode.IsSet()
}

// HasTeamCode returns a boolean if a field has been set.
func (o *ExportItem) HasTeamCode() bool {
	if o != nil && o.TeamCode.IsSet() {
		return true
	}

	return false
}

// SetTeamCode gets a reference to the given NullableString and assigns it to the TeamCode field.
func (o *ExportItem) SetTeamCode(v string) {
	o.TeamCode.Set(&v)
}

// SetTeamCodeNil sets the value for TeamCode to be an explicit nil
func (o *ExportItem) SetTeamCodeNil() {
	o.TeamCode.Set(nil)
}

// UnsetTeamCode ensures that no value is present for TeamCode, not even an explicit nil
func (o *ExportItem) UnsetTeamCode() {
	o.TeamCode.Unset()
}

// GetType returns the Type field value
func (o *ExportItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExportItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExportItem) SetType(v string) {
	o.Type = v
}

// GetUser returns the User field value
func (o *ExportItem) GetUser() ExportItemUser {
	if o == nil {
		var ret ExportItemUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ExportItem) GetUserOk() (*ExportItemUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ExportItem) SetUser(v ExportItemUser) {
	o.User = v
}

func (o ExportItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_links"] = o.Links
	}
	if true {
		toSerialize["accountingEntryId"] = o.AccountingEntryId
	}
	if true {
		toSerialize["accountingEntryLines"] = o.AccountingEntryLines
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["companyId"] = o.CompanyId
	}
	if true {
		toSerialize["date"] = o.Date
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Note.IsSet() {
		toSerialize["note"] = o.Note.Get()
	}
	if o.Supplier.IsSet() {
		toSerialize["supplier"] = o.Supplier.Get()
	}
	if o.TeamCode.IsSet() {
		toSerialize["teamCode"] = o.TeamCode.Get()
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableExportItem struct {
	value *ExportItem
	isSet bool
}

func (v NullableExportItem) Get() *ExportItem {
	return v.value
}

func (v *NullableExportItem) Set(val *ExportItem) {
	v.value = val
	v.isSet = true
}

func (v NullableExportItem) IsSet() bool {
	return v.isSet
}

func (v *NullableExportItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportItem(val *ExportItem) *NullableExportItem {
	return &NullableExportItem{value: val, isSet: true}
}

func (v NullableExportItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
