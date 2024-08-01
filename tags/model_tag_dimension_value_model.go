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
	"time"
)

// TagDimensionValueModel struct for TagDimensionValueModel
type TagDimensionValueModel struct {
	// Creation date and time
	CreatedAt time.Time `json:"createdAt"`
	// Unique identifier of the Tag Group Dimension this dimension value is associated with
	DimensionId string `json:"dimensionId"`
	// Unique identifier of the Tag this dimension value belongs to
	TagId string `json:"tagId"`
	// Date and time of the last update
	UpdatedAt time.Time `json:"updatedAt"`
	// Value of the given dimension for the given tag
	Value string `json:"value"`
}

// NewTagDimensionValueModel instantiates a new TagDimensionValueModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagDimensionValueModel(createdAt time.Time, dimensionId string, tagId string, updatedAt time.Time, value string) *TagDimensionValueModel {
	this := TagDimensionValueModel{}
	this.CreatedAt = createdAt
	this.DimensionId = dimensionId
	this.TagId = tagId
	this.UpdatedAt = updatedAt
	this.Value = value
	return &this
}

// NewTagDimensionValueModelWithDefaults instantiates a new TagDimensionValueModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagDimensionValueModelWithDefaults() *TagDimensionValueModel {
	this := TagDimensionValueModel{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *TagDimensionValueModel) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TagDimensionValueModel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TagDimensionValueModel) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDimensionId returns the DimensionId field value
func (o *TagDimensionValueModel) GetDimensionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DimensionId
}

// GetDimensionIdOk returns a tuple with the DimensionId field value
// and a boolean to check if the value has been set.
func (o *TagDimensionValueModel) GetDimensionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DimensionId, true
}

// SetDimensionId sets field value
func (o *TagDimensionValueModel) SetDimensionId(v string) {
	o.DimensionId = v
}

// GetTagId returns the TagId field value
func (o *TagDimensionValueModel) GetTagId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TagId
}

// GetTagIdOk returns a tuple with the TagId field value
// and a boolean to check if the value has been set.
func (o *TagDimensionValueModel) GetTagIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagId, true
}

// SetTagId sets field value
func (o *TagDimensionValueModel) SetTagId(v string) {
	o.TagId = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TagDimensionValueModel) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TagDimensionValueModel) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TagDimensionValueModel) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetValue returns the Value field value
func (o *TagDimensionValueModel) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TagDimensionValueModel) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TagDimensionValueModel) SetValue(v string) {
	o.Value = v
}

func (o TagDimensionValueModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["dimensionId"] = o.DimensionId
	}
	if true {
		toSerialize["tagId"] = o.TagId
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableTagDimensionValueModel struct {
	value *TagDimensionValueModel
	isSet bool
}

func (v NullableTagDimensionValueModel) Get() *TagDimensionValueModel {
	return v.value
}

func (v *NullableTagDimensionValueModel) Set(val *TagDimensionValueModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTagDimensionValueModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTagDimensionValueModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagDimensionValueModel(val *TagDimensionValueModel) *NullableTagDimensionValueModel {
	return &NullableTagDimensionValueModel{value: val, isSet: true}
}

func (v NullableTagDimensionValueModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagDimensionValueModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}