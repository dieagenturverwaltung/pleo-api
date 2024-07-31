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

// AggregatedTagSearchRequest struct for AggregatedTagSearchRequest
type AggregatedTagSearchRequest struct {
	// Useful for fetching a list of tags given a specific list of IDs.
	TagIds []string `json:"tagIds,omitempty"`
	// Free text search for tags by name
	TextSearch *string `json:"textSearch,omitempty"`
}

// NewAggregatedTagSearchRequest instantiates a new AggregatedTagSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregatedTagSearchRequest() *AggregatedTagSearchRequest {
	this := AggregatedTagSearchRequest{}
	return &this
}

// NewAggregatedTagSearchRequestWithDefaults instantiates a new AggregatedTagSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregatedTagSearchRequestWithDefaults() *AggregatedTagSearchRequest {
	this := AggregatedTagSearchRequest{}
	return &this
}

// GetTagIds returns the TagIds field value if set, zero value otherwise.
func (o *AggregatedTagSearchRequest) GetTagIds() []string {
	if o == nil || o.TagIds == nil {
		var ret []string
		return ret
	}
	return o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedTagSearchRequest) GetTagIdsOk() ([]string, bool) {
	if o == nil || o.TagIds == nil {
		return nil, false
	}
	return o.TagIds, true
}

// HasTagIds returns a boolean if a field has been set.
func (o *AggregatedTagSearchRequest) HasTagIds() bool {
	if o != nil && o.TagIds != nil {
		return true
	}

	return false
}

// SetTagIds gets a reference to the given []string and assigns it to the TagIds field.
func (o *AggregatedTagSearchRequest) SetTagIds(v []string) {
	o.TagIds = v
}

// GetTextSearch returns the TextSearch field value if set, zero value otherwise.
func (o *AggregatedTagSearchRequest) GetTextSearch() string {
	if o == nil || o.TextSearch == nil {
		var ret string
		return ret
	}
	return *o.TextSearch
}

// GetTextSearchOk returns a tuple with the TextSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedTagSearchRequest) GetTextSearchOk() (*string, bool) {
	if o == nil || o.TextSearch == nil {
		return nil, false
	}
	return o.TextSearch, true
}

// HasTextSearch returns a boolean if a field has been set.
func (o *AggregatedTagSearchRequest) HasTextSearch() bool {
	if o != nil && o.TextSearch != nil {
		return true
	}

	return false
}

// SetTextSearch gets a reference to the given string and assigns it to the TextSearch field.
func (o *AggregatedTagSearchRequest) SetTextSearch(v string) {
	o.TextSearch = &v
}

func (o AggregatedTagSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TagIds != nil {
		toSerialize["tagIds"] = o.TagIds
	}
	if o.TextSearch != nil {
		toSerialize["textSearch"] = o.TextSearch
	}
	return json.Marshal(toSerialize)
}

type NullableAggregatedTagSearchRequest struct {
	value *AggregatedTagSearchRequest
	isSet bool
}

func (v NullableAggregatedTagSearchRequest) Get() *AggregatedTagSearchRequest {
	return v.value
}

func (v *NullableAggregatedTagSearchRequest) Set(val *AggregatedTagSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregatedTagSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregatedTagSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregatedTagSearchRequest(val *AggregatedTagSearchRequest) *NullableAggregatedTagSearchRequest {
	return &NullableAggregatedTagSearchRequest{value: val, isSet: true}
}

func (v NullableAggregatedTagSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregatedTagSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
