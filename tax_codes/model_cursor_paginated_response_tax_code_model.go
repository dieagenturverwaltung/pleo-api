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

// CursorPaginatedResponseTaxCodeModel struct for CursorPaginatedResponseTaxCodeModel
type CursorPaginatedResponseTaxCodeModel struct {
	Data       []TaxCodeModel `json:"data"`
	Pagination CursorPageInfo `json:"pagination"`
}

// NewCursorPaginatedResponseTaxCodeModel instantiates a new CursorPaginatedResponseTaxCodeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCursorPaginatedResponseTaxCodeModel(data []TaxCodeModel, pagination CursorPageInfo) *CursorPaginatedResponseTaxCodeModel {
	this := CursorPaginatedResponseTaxCodeModel{}
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewCursorPaginatedResponseTaxCodeModelWithDefaults instantiates a new CursorPaginatedResponseTaxCodeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCursorPaginatedResponseTaxCodeModelWithDefaults() *CursorPaginatedResponseTaxCodeModel {
	this := CursorPaginatedResponseTaxCodeModel{}
	return &this
}

// GetData returns the Data field value
func (o *CursorPaginatedResponseTaxCodeModel) GetData() []TaxCodeModel {
	if o == nil {
		var ret []TaxCodeModel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CursorPaginatedResponseTaxCodeModel) GetDataOk() ([]TaxCodeModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CursorPaginatedResponseTaxCodeModel) SetData(v []TaxCodeModel) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *CursorPaginatedResponseTaxCodeModel) GetPagination() CursorPageInfo {
	if o == nil {
		var ret CursorPageInfo
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *CursorPaginatedResponseTaxCodeModel) GetPaginationOk() (*CursorPageInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *CursorPaginatedResponseTaxCodeModel) SetPagination(v CursorPageInfo) {
	o.Pagination = v
}

func (o CursorPaginatedResponseTaxCodeModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableCursorPaginatedResponseTaxCodeModel struct {
	value *CursorPaginatedResponseTaxCodeModel
	isSet bool
}

func (v NullableCursorPaginatedResponseTaxCodeModel) Get() *CursorPaginatedResponseTaxCodeModel {
	return v.value
}

func (v *NullableCursorPaginatedResponseTaxCodeModel) Set(val *CursorPaginatedResponseTaxCodeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCursorPaginatedResponseTaxCodeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCursorPaginatedResponseTaxCodeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCursorPaginatedResponseTaxCodeModel(val *CursorPaginatedResponseTaxCodeModel) *NullableCursorPaginatedResponseTaxCodeModel {
	return &NullableCursorPaginatedResponseTaxCodeModel{value: val, isSet: true}
}

func (v NullableCursorPaginatedResponseTaxCodeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCursorPaginatedResponseTaxCodeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
