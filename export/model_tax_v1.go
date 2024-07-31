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
)

// TaxV1 Tax details applicable to this accounting entry line. The tax details would usually either have the code or the identifier or both.
type TaxV1 struct {
	Amount ExportItemAmount `json:"amount"`
	// Tax code
	Code NullableString `json:"code,omitempty"`
	// This is the Pleo internal identifier of the tax code
	Id string `json:"id"`
	// Tax rate. This is represented in decimals and not the percentage. e.g. 20% tax rate would be 0.20
	Rate float32 `json:"rate"`
	// Tax type
	Type string `json:"type"`
}

// NewTaxV1 instantiates a new TaxV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxV1(amount ExportItemAmount, id string, rate float32, type_ string) *TaxV1 {
	this := TaxV1{}
	this.Amount = amount
	this.Id = id
	this.Rate = rate
	this.Type = type_
	return &this
}

// NewTaxV1WithDefaults instantiates a new TaxV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxV1WithDefaults() *TaxV1 {
	this := TaxV1{}
	return &this
}

// GetAmount returns the Amount field value
func (o *TaxV1) GetAmount() ExportItemAmount {
	if o == nil {
		var ret ExportItemAmount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TaxV1) GetAmountOk() (*ExportItemAmount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TaxV1) SetAmount(v ExportItemAmount) {
	o.Amount = v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxV1) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxV1) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *TaxV1) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *TaxV1) SetCode(v string) {
	o.Code.Set(&v)
}

// SetCodeNil sets the value for Code to be an explicit nil
func (o *TaxV1) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *TaxV1) UnsetCode() {
	o.Code.Unset()
}

// GetId returns the Id field value
func (o *TaxV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TaxV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaxV1) SetId(v string) {
	o.Id = v
}

// GetRate returns the Rate field value
func (o *TaxV1) GetRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *TaxV1) GetRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value
func (o *TaxV1) SetRate(v float32) {
	o.Rate = v
}

// GetType returns the Type field value
func (o *TaxV1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TaxV1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaxV1) SetType(v string) {
	o.Type = v
}

func (o TaxV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["rate"] = o.Rate
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTaxV1 struct {
	value *TaxV1
	isSet bool
}

func (v NullableTaxV1) Get() *TaxV1 {
	return v.value
}

func (v *NullableTaxV1) Set(val *TaxV1) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxV1) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxV1(val *TaxV1) *NullableTaxV1 {
	return &NullableTaxV1{value: val, isSet: true}
}

func (v NullableTaxV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
