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

// AggregatedTagGroupDimensionModel List of all the dimensions associated with this tag group
type AggregatedTagGroupDimensionModel struct {
	// External identifier of the Tag group Dimension used for mapping to accounting system
	Code string `json:"code"`
	// Value of the given dimension for the given tag
	DisplayOrder int32 `json:"displayOrder"`
	// Unique identifier of Tag Group Dimension
	Id string `json:"id"`
	// User readable name of Tag Group Dimension
	Name string `json:"name"`
	// Determines if this dimension is displayed in the UI
	Visible bool `json:"visible"`
}

// NewAggregatedTagGroupDimensionModel instantiates a new AggregatedTagGroupDimensionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregatedTagGroupDimensionModel(code string, displayOrder int32, id string, name string, visible bool) *AggregatedTagGroupDimensionModel {
	this := AggregatedTagGroupDimensionModel{}
	this.Code = code
	this.DisplayOrder = displayOrder
	this.Id = id
	this.Name = name
	this.Visible = visible
	return &this
}

// NewAggregatedTagGroupDimensionModelWithDefaults instantiates a new AggregatedTagGroupDimensionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregatedTagGroupDimensionModelWithDefaults() *AggregatedTagGroupDimensionModel {
	this := AggregatedTagGroupDimensionModel{}
	return &this
}

// GetCode returns the Code field value
func (o *AggregatedTagGroupDimensionModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AggregatedTagGroupDimensionModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AggregatedTagGroupDimensionModel) SetCode(v string) {
	o.Code = v
}

// GetDisplayOrder returns the DisplayOrder field value
func (o *AggregatedTagGroupDimensionModel) GetDisplayOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value
// and a boolean to check if the value has been set.
func (o *AggregatedTagGroupDimensionModel) GetDisplayOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayOrder, true
}

// SetDisplayOrder sets field value
func (o *AggregatedTagGroupDimensionModel) SetDisplayOrder(v int32) {
	o.DisplayOrder = v
}

// GetId returns the Id field value
func (o *AggregatedTagGroupDimensionModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AggregatedTagGroupDimensionModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AggregatedTagGroupDimensionModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AggregatedTagGroupDimensionModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AggregatedTagGroupDimensionModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AggregatedTagGroupDimensionModel) SetName(v string) {
	o.Name = v
}

// GetVisible returns the Visible field value
func (o *AggregatedTagGroupDimensionModel) GetVisible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value
// and a boolean to check if the value has been set.
func (o *AggregatedTagGroupDimensionModel) GetVisibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visible, true
}

// SetVisible sets field value
func (o *AggregatedTagGroupDimensionModel) SetVisible(v bool) {
	o.Visible = v
}

func (o AggregatedTagGroupDimensionModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["visible"] = o.Visible
	}
	return json.Marshal(toSerialize)
}

type NullableAggregatedTagGroupDimensionModel struct {
	value *AggregatedTagGroupDimensionModel
	isSet bool
}

func (v NullableAggregatedTagGroupDimensionModel) Get() *AggregatedTagGroupDimensionModel {
	return v.value
}

func (v *NullableAggregatedTagGroupDimensionModel) Set(val *AggregatedTagGroupDimensionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregatedTagGroupDimensionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregatedTagGroupDimensionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregatedTagGroupDimensionModel(val *AggregatedTagGroupDimensionModel) *NullableAggregatedTagGroupDimensionModel {
	return &NullableAggregatedTagGroupDimensionModel{value: val, isSet: true}
}

func (v NullableAggregatedTagGroupDimensionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregatedTagGroupDimensionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}