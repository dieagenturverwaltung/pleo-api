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

// TagModel struct for TagModel
type TagModel struct {
	// This tag is not used anymore
	Archived bool `json:"archived"`
	// External identifier of the Tag
	Code string `json:"code"`
	// Creation date and time
	CreatedAt shared.Time `json:"createdAt"`
	// Unique identifier of the Tag Group this tag belongs to
	GroupId string `json:"groupId"`
	// Unique identifier of Tag
	Id string `json:"id"`
	// User readable name that is used for the possible value within a tag group on an expense
	Name *string `json:"name,omitempty"`
	// Date and time of the last update
	UpdatedAt shared.Time `json:"updatedAt"`
}

// NewTagModel instantiates a new TagModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagModel(archived bool, code string, createdAt shared.Time, groupId string, id string, updatedAt shared.Time) *TagModel {
	this := TagModel{}
	this.Archived = archived
	this.Code = code
	this.CreatedAt = createdAt
	this.GroupId = groupId
	this.Id = id
	this.UpdatedAt = updatedAt
	return &this
}

// NewTagModelWithDefaults instantiates a new TagModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagModelWithDefaults() *TagModel {
	this := TagModel{}
	return &this
}

// GetArchived returns the Archived field value
func (o *TagModel) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *TagModel) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *TagModel) SetArchived(v bool) {
	o.Archived = v
}

// GetCode returns the Code field value
func (o *TagModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *TagModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *TagModel) SetCode(v string) {
	o.Code = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TagModel) GetCreatedAt() shared.Time {
	if o == nil {
		var ret shared.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TagModel) GetCreatedAtOk() (*shared.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TagModel) SetCreatedAt(v shared.Time) {
	o.CreatedAt = v
}

// GetGroupId returns the GroupId field value
func (o *TagModel) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *TagModel) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *TagModel) SetGroupId(v string) {
	o.GroupId = v
}

// GetId returns the Id field value
func (o *TagModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TagModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TagModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TagModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TagModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TagModel) SetName(v string) {
	o.Name = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TagModel) GetUpdatedAt() shared.Time {
	if o == nil {
		var ret shared.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TagModel) GetUpdatedAtOk() (*shared.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TagModel) SetUpdatedAt(v shared.Time) {
	o.UpdatedAt = v
}

func (o TagModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["archived"] = o.Archived
	}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["groupId"] = o.GroupId
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableTagModel struct {
	value *TagModel
	isSet bool
}

func (v NullableTagModel) Get() *TagModel {
	return v.value
}

func (v *NullableTagModel) Set(val *TagModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTagModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTagModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagModel(val *TagModel) *NullableTagModel {
	return &NullableTagModel{value: val, isSet: true}
}

func (v NullableTagModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
