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

	"github.com/dieagenturverwaltung/pleo-api/shared"
)

// TagGroupModel struct for TagGroupModel
type TagGroupModel struct {
	// This tag group is not used anymore
	Archived bool `json:"archived"`
	// External identifier of the Tag group / Dimension used for mapping to accounting system
	Code string `json:"code"`
	// Unique identifier of the company the Tag Group belongs to
	CompanyId string `json:"companyId"`
	// Creation date and time
	CreatedAt shared.Time `json:"createdAt"`
	// Unique identifier of Tag Group (generated on creation)
	Id string `json:"id"`
	// Place for API users to store flexible data
	Metadata map[string]string `json:"metadata"`
	// User readable name of Tag Group
	Name string `json:"name"`
	// Date and time of the last update
	UpdatedAt shared.Time `json:"updatedAt"`
}

// NewTagGroupModel instantiates a new TagGroupModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagGroupModel(archived bool, code string, companyId string, createdAt shared.Time, id string, metadata map[string]string, name string, updatedAt shared.Time) *TagGroupModel {
	this := TagGroupModel{}
	this.Archived = archived
	this.Code = code
	this.CompanyId = companyId
	this.CreatedAt = createdAt
	this.Id = id
	this.Metadata = metadata
	this.Name = name
	this.UpdatedAt = updatedAt
	return &this
}

// NewTagGroupModelWithDefaults instantiates a new TagGroupModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagGroupModelWithDefaults() *TagGroupModel {
	this := TagGroupModel{}
	return &this
}

// GetArchived returns the Archived field value
func (o *TagGroupModel) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *TagGroupModel) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *TagGroupModel) SetArchived(v bool) {
	o.Archived = v
}

// GetCode returns the Code field value
func (o *TagGroupModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *TagGroupModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *TagGroupModel) SetCode(v string) {
	o.Code = v
}

// GetCompanyId returns the CompanyId field value
func (o *TagGroupModel) GetCompanyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *TagGroupModel) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *TagGroupModel) SetCompanyId(v string) {
	o.CompanyId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TagGroupModel) GetCreatedAt() shared.Time {
	if o == nil {
		var ret shared.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TagGroupModel) GetCreatedAtOk() (*shared.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TagGroupModel) SetCreatedAt(v shared.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *TagGroupModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TagGroupModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TagGroupModel) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value
func (o *TagGroupModel) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *TagGroupModel) GetMetadataOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *TagGroupModel) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetName returns the Name field value
func (o *TagGroupModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagGroupModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagGroupModel) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TagGroupModel) GetUpdatedAt() shared.Time {
	if o == nil {
		var ret shared.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TagGroupModel) GetUpdatedAtOk() (*shared.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TagGroupModel) SetUpdatedAt(v shared.Time) {
	o.UpdatedAt = v
}

func (o TagGroupModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["archived"] = o.Archived
	}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["companyId"] = o.CompanyId
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableTagGroupModel struct {
	value *TagGroupModel
	isSet bool
}

func (v NullableTagGroupModel) Get() *TagGroupModel {
	return v.value
}

func (v *NullableTagGroupModel) Set(val *TagGroupModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTagGroupModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTagGroupModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagGroupModel(val *TagGroupModel) *NullableTagGroupModel {
	return &NullableTagGroupModel{value: val, isSet: true}
}

func (v NullableTagGroupModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagGroupModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
