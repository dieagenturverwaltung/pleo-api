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

// DataResponseTagGroupModel struct for DataResponseTagGroupModel
type DataResponseTagGroupModel struct {
	Data TagGroupModel `json:"data"`
}

// NewDataResponseTagGroupModel instantiates a new DataResponseTagGroupModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataResponseTagGroupModel(data TagGroupModel) *DataResponseTagGroupModel {
	this := DataResponseTagGroupModel{}
	this.Data = data
	return &this
}

// NewDataResponseTagGroupModelWithDefaults instantiates a new DataResponseTagGroupModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataResponseTagGroupModelWithDefaults() *DataResponseTagGroupModel {
	this := DataResponseTagGroupModel{}
	return &this
}

// GetData returns the Data field value
func (o *DataResponseTagGroupModel) GetData() TagGroupModel {
	if o == nil {
		var ret TagGroupModel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DataResponseTagGroupModel) GetDataOk() (*TagGroupModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DataResponseTagGroupModel) SetData(v TagGroupModel) {
	o.Data = v
}

func (o DataResponseTagGroupModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableDataResponseTagGroupModel struct {
	value *DataResponseTagGroupModel
	isSet bool
}

func (v NullableDataResponseTagGroupModel) Get() *DataResponseTagGroupModel {
	return v.value
}

func (v *NullableDataResponseTagGroupModel) Set(val *DataResponseTagGroupModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDataResponseTagGroupModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDataResponseTagGroupModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataResponseTagGroupModel(val *DataResponseTagGroupModel) *NullableDataResponseTagGroupModel {
	return &NullableDataResponseTagGroupModel{value: val, isSet: true}
}

func (v NullableDataResponseTagGroupModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataResponseTagGroupModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
