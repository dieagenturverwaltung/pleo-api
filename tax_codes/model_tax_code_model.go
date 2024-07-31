/*
Tax Code API

Tax Codes OpenAPI definitions

API version: 42.14.0
Contact: team-expense-core@pleo.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tax_codes

import (
	"encoding/json"
	"time"
)

// TaxCodeModel struct for TaxCodeModel
type TaxCodeModel struct {
	// Accounting integration system
	AccountingIntegrationSystem NullableString `json:"accountingIntegrationSystem,omitempty"`
	// Boolean flag used to archive or unarchive an account (when set to true, account is not visible or usable on the platform)
	Archived bool `json:"archived"`
	// The accounting system's internal identifier of the tax code
	Code NullableString `json:"code,omitempty"`
	// The Pleo unique identifier of the company the tax code belongs to
	CompanyId string `json:"companyId"`
	// Date and time the tax code was created
	CreatedAt time.Time `json:"createdAt"`
	// The unique identifier generated by Pleo for the tax code
	Id string `json:"id"`
	// Ingoing tax account usually used to account for reverse VAT
	IngoingTaxAccount NullableString `json:"ingoingTaxAccount,omitempty"`
	// Name of the tax code
	Name string `json:"name"`
	// Outgoing tax account usually used to account for reverse VAT
	OutgoingTaxAccount NullableString `json:"outgoingTaxAccount,omitempty"`
	// Percentage rate applied for this tax code (This is represented in decimals and not the percentage. e.g. 20% tax rate would be 0.20)
	Rate float32 `json:"rate"`
	// Classification of this tax code
	Type string `json:"type"`
	// Date and time the tax code was last updated
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewTaxCodeModel instantiates a new TaxCodeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxCodeModel(archived bool, companyId string, createdAt time.Time, id string, name string, rate float32, type_ string, updatedAt time.Time) *TaxCodeModel {
	this := TaxCodeModel{}
	this.Archived = archived
	this.CompanyId = companyId
	this.CreatedAt = createdAt
	this.Id = id
	this.Name = name
	this.Rate = rate
	this.Type = type_
	this.UpdatedAt = updatedAt
	return &this
}

// NewTaxCodeModelWithDefaults instantiates a new TaxCodeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxCodeModelWithDefaults() *TaxCodeModel {
	this := TaxCodeModel{}
	return &this
}

// GetAccountingIntegrationSystem returns the AccountingIntegrationSystem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxCodeModel) GetAccountingIntegrationSystem() string {
	if o == nil || o.AccountingIntegrationSystem.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountingIntegrationSystem.Get()
}

// GetAccountingIntegrationSystemOk returns a tuple with the AccountingIntegrationSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxCodeModel) GetAccountingIntegrationSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountingIntegrationSystem.Get(), o.AccountingIntegrationSystem.IsSet()
}

// HasAccountingIntegrationSystem returns a boolean if a field has been set.
func (o *TaxCodeModel) HasAccountingIntegrationSystem() bool {
	if o != nil && o.AccountingIntegrationSystem.IsSet() {
		return true
	}

	return false
}

// SetAccountingIntegrationSystem gets a reference to the given NullableString and assigns it to the AccountingIntegrationSystem field.
func (o *TaxCodeModel) SetAccountingIntegrationSystem(v string) {
	o.AccountingIntegrationSystem.Set(&v)
}

// SetAccountingIntegrationSystemNil sets the value for AccountingIntegrationSystem to be an explicit nil
func (o *TaxCodeModel) SetAccountingIntegrationSystemNil() {
	o.AccountingIntegrationSystem.Set(nil)
}

// UnsetAccountingIntegrationSystem ensures that no value is present for AccountingIntegrationSystem, not even an explicit nil
func (o *TaxCodeModel) UnsetAccountingIntegrationSystem() {
	o.AccountingIntegrationSystem.Unset()
}

// GetArchived returns the Archived field value
func (o *TaxCodeModel) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *TaxCodeModel) SetArchived(v bool) {
	o.Archived = v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxCodeModel) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxCodeModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *TaxCodeModel) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *TaxCodeModel) SetCode(v string) {
	o.Code.Set(&v)
}

// SetCodeNil sets the value for Code to be an explicit nil
func (o *TaxCodeModel) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *TaxCodeModel) UnsetCode() {
	o.Code.Unset()
}

// GetCompanyId returns the CompanyId field value
func (o *TaxCodeModel) GetCompanyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *TaxCodeModel) SetCompanyId(v string) {
	o.CompanyId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TaxCodeModel) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TaxCodeModel) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *TaxCodeModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaxCodeModel) SetId(v string) {
	o.Id = v
}

// GetIngoingTaxAccount returns the IngoingTaxAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxCodeModel) GetIngoingTaxAccount() string {
	if o == nil || o.IngoingTaxAccount.Get() == nil {
		var ret string
		return ret
	}
	return *o.IngoingTaxAccount.Get()
}

// GetIngoingTaxAccountOk returns a tuple with the IngoingTaxAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxCodeModel) GetIngoingTaxAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IngoingTaxAccount.Get(), o.IngoingTaxAccount.IsSet()
}

// HasIngoingTaxAccount returns a boolean if a field has been set.
func (o *TaxCodeModel) HasIngoingTaxAccount() bool {
	if o != nil && o.IngoingTaxAccount.IsSet() {
		return true
	}

	return false
}

// SetIngoingTaxAccount gets a reference to the given NullableString and assigns it to the IngoingTaxAccount field.
func (o *TaxCodeModel) SetIngoingTaxAccount(v string) {
	o.IngoingTaxAccount.Set(&v)
}

// SetIngoingTaxAccountNil sets the value for IngoingTaxAccount to be an explicit nil
func (o *TaxCodeModel) SetIngoingTaxAccountNil() {
	o.IngoingTaxAccount.Set(nil)
}

// UnsetIngoingTaxAccount ensures that no value is present for IngoingTaxAccount, not even an explicit nil
func (o *TaxCodeModel) UnsetIngoingTaxAccount() {
	o.IngoingTaxAccount.Unset()
}

// GetName returns the Name field value
func (o *TaxCodeModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TaxCodeModel) SetName(v string) {
	o.Name = v
}

// GetOutgoingTaxAccount returns the OutgoingTaxAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxCodeModel) GetOutgoingTaxAccount() string {
	if o == nil || o.OutgoingTaxAccount.Get() == nil {
		var ret string
		return ret
	}
	return *o.OutgoingTaxAccount.Get()
}

// GetOutgoingTaxAccountOk returns a tuple with the OutgoingTaxAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxCodeModel) GetOutgoingTaxAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutgoingTaxAccount.Get(), o.OutgoingTaxAccount.IsSet()
}

// HasOutgoingTaxAccount returns a boolean if a field has been set.
func (o *TaxCodeModel) HasOutgoingTaxAccount() bool {
	if o != nil && o.OutgoingTaxAccount.IsSet() {
		return true
	}

	return false
}

// SetOutgoingTaxAccount gets a reference to the given NullableString and assigns it to the OutgoingTaxAccount field.
func (o *TaxCodeModel) SetOutgoingTaxAccount(v string) {
	o.OutgoingTaxAccount.Set(&v)
}

// SetOutgoingTaxAccountNil sets the value for OutgoingTaxAccount to be an explicit nil
func (o *TaxCodeModel) SetOutgoingTaxAccountNil() {
	o.OutgoingTaxAccount.Set(nil)
}

// UnsetOutgoingTaxAccount ensures that no value is present for OutgoingTaxAccount, not even an explicit nil
func (o *TaxCodeModel) UnsetOutgoingTaxAccount() {
	o.OutgoingTaxAccount.Unset()
}

// GetRate returns the Rate field value
func (o *TaxCodeModel) GetRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value
func (o *TaxCodeModel) SetRate(v float32) {
	o.Rate = v
}

// GetType returns the Type field value
func (o *TaxCodeModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaxCodeModel) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TaxCodeModel) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TaxCodeModel) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TaxCodeModel) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o TaxCodeModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountingIntegrationSystem.IsSet() {
		toSerialize["accountingIntegrationSystem"] = o.AccountingIntegrationSystem.Get()
	}
	if true {
		toSerialize["archived"] = o.Archived
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if true {
		toSerialize["companyId"] = o.CompanyId
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.IngoingTaxAccount.IsSet() {
		toSerialize["ingoingTaxAccount"] = o.IngoingTaxAccount.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.OutgoingTaxAccount.IsSet() {
		toSerialize["outgoingTaxAccount"] = o.OutgoingTaxAccount.Get()
	}
	if true {
		toSerialize["rate"] = o.Rate
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableTaxCodeModel struct {
	value *TaxCodeModel
	isSet bool
}

func (v NullableTaxCodeModel) Get() *TaxCodeModel {
	return v.value
}

func (v *NullableTaxCodeModel) Set(val *TaxCodeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCodeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCodeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCodeModel(val *TaxCodeModel) *NullableTaxCodeModel {
	return &NullableTaxCodeModel{value: val, isSet: true}
}

func (v NullableTaxCodeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCodeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
