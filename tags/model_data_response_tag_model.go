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

// DataResponseTagModel struct for DataResponseTagModel
type DataResponseTagModel struct {
	Data TagModel `json:"data"`
}

// NewDataResponseTagModel instantiates a new DataResponseTagModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataResponseTagModel(data TagModel) *DataResponseTagModel {
	this := DataResponseTagModel{}
	this.Data = data
	return &this
}

// NewDataResponseTagModelWithDefaults instantiates a new DataResponseTagModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataResponseTagModelWithDefaults() *DataResponseTagModel {
	this := DataResponseTagModel{}
	return &this
}

// GetData returns the Data field value
func (o *DataResponseTagModel) GetData() TagModel {
	if o == nil {
		var ret TagModel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DataResponseTagModel) GetDataOk() (*TagModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DataResponseTagModel) SetData(v TagModel) {
	o.Data = v
}

func (o DataResponseTagModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableDataResponseTagModel struct {
	value *DataResponseTagModel
	isSet bool
}

func (v NullableDataResponseTagModel) Get() *DataResponseTagModel {
	return v.value
}

func (v *NullableDataResponseTagModel) Set(val *DataResponseTagModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDataResponseTagModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDataResponseTagModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataResponseTagModel(val *DataResponseTagModel) *NullableDataResponseTagModel {
	return &NullableDataResponseTagModel{value: val, isSet: true}
}

func (v NullableDataResponseTagModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataResponseTagModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
