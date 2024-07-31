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

// ExportItemFile Files that have been attached to this accounting entry
type ExportItemFile struct {
	// Size of the file in bytes
	Size int32 `json:"size"`
	// Type of file, usually represented similar to mime type
	Type string `json:"type"`
	// URL to retrieve the file attachment from. This URL is active for a limited time (24 Hours) from the time the export item is fetched.
	Url string `json:"url"`
}

// NewExportItemFile instantiates a new ExportItemFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportItemFile(size int32, type_ string, url string) *ExportItemFile {
	this := ExportItemFile{}
	this.Size = size
	this.Type = type_
	this.Url = url
	return &this
}

// NewExportItemFileWithDefaults instantiates a new ExportItemFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportItemFileWithDefaults() *ExportItemFile {
	this := ExportItemFile{}
	return &this
}

// GetSize returns the Size field value
func (o *ExportItemFile) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ExportItemFile) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ExportItemFile) SetSize(v int32) {
	o.Size = v
}

// GetType returns the Type field value
func (o *ExportItemFile) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExportItemFile) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExportItemFile) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *ExportItemFile) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ExportItemFile) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ExportItemFile) SetUrl(v string) {
	o.Url = v
}

func (o ExportItemFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableExportItemFile struct {
	value *ExportItemFile
	isSet bool
}

func (v NullableExportItemFile) Get() *ExportItemFile {
	return v.value
}

func (v *NullableExportItemFile) Set(val *ExportItemFile) {
	v.value = val
	v.isSet = true
}

func (v NullableExportItemFile) IsSet() bool {
	return v.isSet
}

func (v *NullableExportItemFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportItemFile(val *ExportItemFile) *NullableExportItemFile {
	return &NullableExportItemFile{value: val, isSet: true}
}

func (v NullableExportItemFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportItemFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
