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
)

// TaxCodeSearchRequest struct for TaxCodeSearchRequest
type TaxCodeSearchRequest struct {
	// If set to true will only return tax codes that have been archived, if set to false will only return tax codes that have not been archived. By default,both archived and non-archived tax codes will be returned.
	Archived *bool `json:"archived,omitempty"`
	// Useful for fetching a list of tax codes given a specific list of IDs.
	TaxCodeIds []string `json:"taxCodeIds,omitempty"`
	// Classification of the tax codes to fetch
	Type *string `json:"type,omitempty"`
}

// NewTaxCodeSearchRequest instantiates a new TaxCodeSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxCodeSearchRequest() *TaxCodeSearchRequest {
	this := TaxCodeSearchRequest{}
	return &this
}

// NewTaxCodeSearchRequestWithDefaults instantiates a new TaxCodeSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxCodeSearchRequestWithDefaults() *TaxCodeSearchRequest {
	this := TaxCodeSearchRequest{}
	return &this
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *TaxCodeSearchRequest) GetArchived() bool {
	if o == nil || o.Archived == nil {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCodeSearchRequest) GetArchivedOk() (*bool, bool) {
	if o == nil || o.Archived == nil {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *TaxCodeSearchRequest) HasArchived() bool {
	if o != nil && o.Archived != nil {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *TaxCodeSearchRequest) SetArchived(v bool) {
	o.Archived = &v
}

// GetTaxCodeIds returns the TaxCodeIds field value if set, zero value otherwise.
func (o *TaxCodeSearchRequest) GetTaxCodeIds() []string {
	if o == nil || o.TaxCodeIds == nil {
		var ret []string
		return ret
	}
	return o.TaxCodeIds
}

// GetTaxCodeIdsOk returns a tuple with the TaxCodeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCodeSearchRequest) GetTaxCodeIdsOk() ([]string, bool) {
	if o == nil || o.TaxCodeIds == nil {
		return nil, false
	}
	return o.TaxCodeIds, true
}

// HasTaxCodeIds returns a boolean if a field has been set.
func (o *TaxCodeSearchRequest) HasTaxCodeIds() bool {
	if o != nil && o.TaxCodeIds != nil {
		return true
	}

	return false
}

// SetTaxCodeIds gets a reference to the given []string and assigns it to the TaxCodeIds field.
func (o *TaxCodeSearchRequest) SetTaxCodeIds(v []string) {
	o.TaxCodeIds = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TaxCodeSearchRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCodeSearchRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TaxCodeSearchRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TaxCodeSearchRequest) SetType(v string) {
	o.Type = &v
}

func (o TaxCodeSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Archived != nil {
		toSerialize["archived"] = o.Archived
	}
	if o.TaxCodeIds != nil {
		toSerialize["taxCodeIds"] = o.TaxCodeIds
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTaxCodeSearchRequest struct {
	value *TaxCodeSearchRequest
	isSet bool
}

func (v NullableTaxCodeSearchRequest) Get() *TaxCodeSearchRequest {
	return v.value
}

func (v *NullableTaxCodeSearchRequest) Set(val *TaxCodeSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCodeSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCodeSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCodeSearchRequest(val *TaxCodeSearchRequest) *NullableTaxCodeSearchRequest {
	return &NullableTaxCodeSearchRequest{value: val, isSet: true}
}

func (v NullableTaxCodeSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCodeSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
