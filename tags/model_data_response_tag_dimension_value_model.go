/*
Tags API

Tags OpenAPI definitions

API version: 31.38.0
Contact: team-expense-core@pleo.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tags

import (
	"encoding/json"
)

// DataResponseTagDimensionValueModel struct for DataResponseTagDimensionValueModel
type DataResponseTagDimensionValueModel struct {
	Data TagDimensionValueModel `json:"data"`
}

// NewDataResponseTagDimensionValueModel instantiates a new DataResponseTagDimensionValueModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataResponseTagDimensionValueModel(data TagDimensionValueModel) *DataResponseTagDimensionValueModel {
	this := DataResponseTagDimensionValueModel{}
	this.Data = data
	return &this
}

// NewDataResponseTagDimensionValueModelWithDefaults instantiates a new DataResponseTagDimensionValueModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataResponseTagDimensionValueModelWithDefaults() *DataResponseTagDimensionValueModel {
	this := DataResponseTagDimensionValueModel{}
	return &this
}

// GetData returns the Data field value
func (o *DataResponseTagDimensionValueModel) GetData() TagDimensionValueModel {
	if o == nil {
		var ret TagDimensionValueModel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DataResponseTagDimensionValueModel) GetDataOk() (*TagDimensionValueModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DataResponseTagDimensionValueModel) SetData(v TagDimensionValueModel) {
	o.Data = v
}

func (o DataResponseTagDimensionValueModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableDataResponseTagDimensionValueModel struct {
	value *DataResponseTagDimensionValueModel
	isSet bool
}

func (v NullableDataResponseTagDimensionValueModel) Get() *DataResponseTagDimensionValueModel {
	return v.value
}

func (v *NullableDataResponseTagDimensionValueModel) Set(val *DataResponseTagDimensionValueModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDataResponseTagDimensionValueModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDataResponseTagDimensionValueModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataResponseTagDimensionValueModel(val *DataResponseTagDimensionValueModel) *NullableDataResponseTagDimensionValueModel {
	return &NullableDataResponseTagDimensionValueModel{value: val, isSet: true}
}

func (v NullableDataResponseTagDimensionValueModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataResponseTagDimensionValueModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
