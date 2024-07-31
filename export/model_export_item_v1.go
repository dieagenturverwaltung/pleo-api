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

// ExportItemV1 struct for ExportItemV1
type ExportItemV1 struct {
	Links LinksResponse `json:"_links"`
	// This is the Pleo internal identifier of the export item
	AccountingEntryId string `json:"accountingEntryId"`
	// Accounting entry broken down in entry lines. There will always be an accounting entry line. If the accounting entry has been split into separate lines, then each line would represent the details of that accounting entry line, else there would only be one line present, representing the whole entry.
	AccountingEntryLines  []ExportItemLineV1                `json:"accountingEntryLines"`
	AdditionalInformation ExportItemAdditionalInformationV1 `json:"additionalInformation"`
	Amount                ExportItemAmount                  `json:"amount"`
	// Pleo company identifier this export item belongs to
	CompanyId string `json:"companyId"`
	// Date the accounting entry should be bookkept
	Date time.Time `json:"date"`
	// Files that have been attached to this accounting entry
	Files []ExportItemFile `json:"files,omitempty"`
	// Additional comments potentially describing the accounting entry
	Note     NullableString     `json:"note,omitempty"`
	Supplier NullableSupplierV1 `json:"supplier,omitempty"`
	Team     NullableTeam       `json:"team,omitempty"`
	// Team code is an identifier assigned to an expense to categorize it or associate it with a specific team
	TeamCode NullableString `json:"teamCode,omitempty"`
	// This is the Pleo internal identifier of the accounting entry this export item is associated with
	Type string           `json:"type"`
	User ExportItemUserV1 `json:"user"`
}

// NewExportItemV1 instantiates a new ExportItemV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportItemV1(links LinksResponse, accountingEntryId string, accountingEntryLines []ExportItemLineV1, additionalInformation ExportItemAdditionalInformationV1, amount ExportItemAmount, companyId string, date time.Time, type_ string, user ExportItemUserV1) *ExportItemV1 {
	this := ExportItemV1{}
	this.Links = links
	this.AccountingEntryId = accountingEntryId
	this.AccountingEntryLines = accountingEntryLines
	this.AdditionalInformation = additionalInformation
	this.Amount = amount
	this.CompanyId = companyId
	this.Date = date
	this.Type = type_
	this.User = user
	return &this
}

// NewExportItemV1WithDefaults instantiates a new ExportItemV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportItemV1WithDefaults() *ExportItemV1 {
	this := ExportItemV1{}
	return &this
}

// GetLinks returns the Links field value
func (o *ExportItemV1) GetLinks() LinksResponse {
	if o == nil {
		var ret LinksResponse
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ExportItemV1) GetLinksOk() (*LinksResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ExportItemV1) SetLinks(v LinksResponse) {
	o.Links = v
}

// GetAccountingEntryId returns the AccountingEntryId field value
func (o *ExportItemV1) GetAccountingEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountingEntryId
}

// GetAccountingEntryIdOk returns a tuple with the AccountingEntryId field value
// and a boolean to check if the value has been set.
func (o *ExportItemV1) GetAccountingEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountingEntryId, true
}

// SetAccountingEntryId sets field value
func (o *ExportItemV1) SetAccountingEntryId(v string) {
	o.AccountingEntryId = v
}

// GetAccountingEntryLines returns the AccountingEntryLines field value
func (o *ExportItemV1) GetAccountingEntryLines() []ExportItemLineV1 {
	if o == nil {
		var ret []ExportItemLineV1
		return ret
	}

	return o.AccountingEntryLines
}

// GetAccountingEntryLinesOk returns a tuple with the AccountingEntryLines field value
// and a boolean to check if the value has been set.
func (o *ExportItemV1) GetAccountingEntryLinesOk() ([]ExportItemLineV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountingEntryLines, true
}

// SetAccountingEntryLines sets field value
func (o *ExportItemV1) SetAccountingEntryLines(v []ExportItemLineV1) {
	o.AccountingEntryLines = v
}

// GetAdditionalInformation returns the AdditionalInformation field value
func (o *ExportItemV1) GetAdditionalInformation() ExportItemAdditionalInformationV1 {
	if o == nil {
		var ret ExportItemAdditionalInformationV1
		return ret
	}

	return o.AdditionalInformation
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value
// and a boolean to check if the value has been set.
func (o *ExportItemV1) GetAdditionalInformationOk() (*ExportItemAdditionalInformationV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdditionalInformation, true
}

// SetAdditionalInformation sets field value
func (o *ExportItemV1) SetAdditionalInformation(v ExportItemAdditionalInformationV1) {
	o.AdditionalInformation = v
}

// GetAmount returns the Amount field value
func (o *ExportItemV1) GetAmount() ExportItemAmount {
	if o == nil {
		var ret ExportItemAmount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ExportItemV1) GetAmountOk() (*ExportItemAmount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ExportItemV1) SetAmount(v ExportItemAmount) {
	o.Amount = v
}

// GetCompanyId returns the CompanyId field value
func (o *ExportItemV1) GetCompanyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *ExportItemV1) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *ExportItemV1) SetCompanyId(v string) {
	o.CompanyId = v
}

// GetDate returns the Date field value
func (o *ExportItemV1) GetDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *ExportItemV1) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *ExportItemV1) SetDate(v time.Time) {
	o.Date = v
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportItemV1) GetFiles() []ExportItemFile {
	if o == nil {
		var ret []ExportItemFile
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportItemV1) GetFilesOk() ([]ExportItemFile, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *ExportItemV1) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []ExportItemFile and assigns it to the Files field.
func (o *ExportItemV1) SetFiles(v []ExportItemFile) {
	o.Files = v
}

// GetNote returns the Note field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportItemV1) GetNote() string {
	if o == nil || o.Note.Get() == nil {
		var ret string
		return ret
	}
	return *o.Note.Get()
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportItemV1) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Note.Get(), o.Note.IsSet()
}

// HasNote returns a boolean if a field has been set.
func (o *ExportItemV1) HasNote() bool {
	if o != nil && o.Note.IsSet() {
		return true
	}

	return false
}

// SetNote gets a reference to the given NullableString and assigns it to the Note field.
func (o *ExportItemV1) SetNote(v string) {
	o.Note.Set(&v)
}

// SetNoteNil sets the value for Note to be an explicit nil
func (o *ExportItemV1) SetNoteNil() {
	o.Note.Set(nil)
}

// UnsetNote ensures that no value is present for Note, not even an explicit nil
func (o *ExportItemV1) UnsetNote() {
	o.Note.Unset()
}

// GetSupplier returns the Supplier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportItemV1) GetSupplier() SupplierV1 {
	if o == nil || o.Supplier.Get() == nil {
		var ret SupplierV1
		return ret
	}
	return *o.Supplier.Get()
}

// GetSupplierOk returns a tuple with the Supplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportItemV1) GetSupplierOk() (*SupplierV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Supplier.Get(), o.Supplier.IsSet()
}

// HasSupplier returns a boolean if a field has been set.
func (o *ExportItemV1) HasSupplier() bool {
	if o != nil && o.Supplier.IsSet() {
		return true
	}

	return false
}

// SetSupplier gets a reference to the given NullableSupplierV1 and assigns it to the Supplier field.
func (o *ExportItemV1) SetSupplier(v SupplierV1) {
	o.Supplier.Set(&v)
}

// SetSupplierNil sets the value for Supplier to be an explicit nil
func (o *ExportItemV1) SetSupplierNil() {
	o.Supplier.Set(nil)
}

// UnsetSupplier ensures that no value is present for Supplier, not even an explicit nil
func (o *ExportItemV1) UnsetSupplier() {
	o.Supplier.Unset()
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportItemV1) GetTeam() Team {
	if o == nil || o.Team.Get() == nil {
		var ret Team
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportItemV1) GetTeamOk() (*Team, bool) {
	if o == nil {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *ExportItemV1) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableTeam and assigns it to the Team field.
func (o *ExportItemV1) SetTeam(v Team) {
	o.Team.Set(&v)
}

// SetTeamNil sets the value for Team to be an explicit nil
func (o *ExportItemV1) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *ExportItemV1) UnsetTeam() {
	o.Team.Unset()
}

// GetTeamCode returns the TeamCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportItemV1) GetTeamCode() string {
	if o == nil || o.TeamCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.TeamCode.Get()
}

// GetTeamCodeOk returns a tuple with the TeamCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportItemV1) GetTeamCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TeamCode.Get(), o.TeamCode.IsSet()
}

// HasTeamCode returns a boolean if a field has been set.
func (o *ExportItemV1) HasTeamCode() bool {
	if o != nil && o.TeamCode.IsSet() {
		return true
	}

	return false
}

// SetTeamCode gets a reference to the given NullableString and assigns it to the TeamCode field.
func (o *ExportItemV1) SetTeamCode(v string) {
	o.TeamCode.Set(&v)
}

// SetTeamCodeNil sets the value for TeamCode to be an explicit nil
func (o *ExportItemV1) SetTeamCodeNil() {
	o.TeamCode.Set(nil)
}

// UnsetTeamCode ensures that no value is present for TeamCode, not even an explicit nil
func (o *ExportItemV1) UnsetTeamCode() {
	o.TeamCode.Unset()
}

// GetType returns the Type field value
func (o *ExportItemV1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExportItemV1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExportItemV1) SetType(v string) {
	o.Type = v
}

// GetUser returns the User field value
func (o *ExportItemV1) GetUser() ExportItemUserV1 {
	if o == nil {
		var ret ExportItemUserV1
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ExportItemV1) GetUserOk() (*ExportItemUserV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ExportItemV1) SetUser(v ExportItemUserV1) {
	o.User = v
}

func (o ExportItemV1) MarshalJSON() ([]byte, error) {
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
		toSerialize["additionalInformation"] = o.AdditionalInformation
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
	if o.Note.IsSet() {
		toSerialize["note"] = o.Note.Get()
	}
	if o.Supplier.IsSet() {
		toSerialize["supplier"] = o.Supplier.Get()
	}
	if o.Team.IsSet() {
		toSerialize["team"] = o.Team.Get()
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

type NullableExportItemV1 struct {
	value *ExportItemV1
	isSet bool
}

func (v NullableExportItemV1) Get() *ExportItemV1 {
	return v.value
}

func (v *NullableExportItemV1) Set(val *ExportItemV1) {
	v.value = val
	v.isSet = true
}

func (v NullableExportItemV1) IsSet() bool {
	return v.isSet
}

func (v *NullableExportItemV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportItemV1(val *ExportItemV1) *NullableExportItemV1 {
	return &NullableExportItemV1{value: val, isSet: true}
}

func (v NullableExportItemV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportItemV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
