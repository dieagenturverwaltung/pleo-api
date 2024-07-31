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
	"time"
)

// ExportJob struct for ExportJob
type ExportJob struct {
	// The Pleo unique identifier of the company the export job belongs to
	CompanyId string `json:"companyId"`
	// Date and time the job was completed
	CompletedAt NullableTime `json:"completedAt,omitempty"`
	// Date and time the job was created. When the export button is clicked on the Pleo UI, i.e. an export is initiated by the user, the job is created and the createdAt date is set.
	CreatedAt time.Time `json:"createdAt"`
	// This is the Pleo unique identifier of the user that initiated the export job
	CreatedBy NullableString `json:"createdBy,omitempty"`
	// Date and time the job expired
	ExpiredAt NullableTime `json:"expiredAt,omitempty"`
	// This is the amount of time in seconds the job will expire relative to the last time an action was taken on the job. The last time an action was taken on the job is reflected by the \"lastUpdatedAt\" field
	ExpiresIn int32 `json:"expiresIn"`
	// Reason why the job failed in the case of a failure
	FailureReason NullableString `json:"failureReason,omitempty"`
	// The classification for the failure from a list of described failure reason types
	FailureReasonType NullableString `json:"failureReasonType,omitempty"`
	// The unique identifier generated by Pleo for the export job
	Id string `json:"id"`
	// Indicates whether the export job was initiated by a user or by the system
	IsInteractive bool `json:"isInteractive"`
	// Last time the job was updated or action on the job was taken
	LastUpdatedAt NullableTime `json:"lastUpdatedAt,omitempty"`
	// Number of accounting entries that were selected for processing
	NumberOfItems int32 `json:"numberOfItems"`
	// Date and time the job was started. When the export-job-events endpoint is called to start the export Job, the startedAt date is set.
	StartedAt NullableTime `json:"startedAt,omitempty"`
	// Specifies the current execution state of the export job. Some here are the definitions of the values that are supported for this field
	Status string `json:"status"`
}

// NewExportJob instantiates a new ExportJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportJob(companyId string, createdAt time.Time, expiresIn int32, id string, isInteractive bool, numberOfItems int32, status string) *ExportJob {
	this := ExportJob{}
	this.CompanyId = companyId
	this.CreatedAt = createdAt
	this.ExpiresIn = expiresIn
	this.Id = id
	this.IsInteractive = isInteractive
	this.NumberOfItems = numberOfItems
	this.Status = status
	return &this
}

// NewExportJobWithDefaults instantiates a new ExportJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportJobWithDefaults() *ExportJob {
	this := ExportJob{}
	return &this
}

// GetCompanyId returns the CompanyId field value
func (o *ExportJob) GetCompanyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *ExportJob) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *ExportJob) SetCompanyId(v string) {
	o.CompanyId = v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportJob) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportJob) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *ExportJob) HasCompletedAt() bool {
	if o != nil && o.CompletedAt.IsSet() {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given NullableTime and assigns it to the CompletedAt field.
func (o *ExportJob) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}

// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil
func (o *ExportJob) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
func (o *ExportJob) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *ExportJob) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ExportJob) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ExportJob) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportJob) GetCreatedBy() string {
	if o == nil || o.CreatedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportJob) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ExportJob) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *ExportJob) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *ExportJob) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *ExportJob) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportJob) GetExpiredAt() time.Time {
	if o == nil || o.ExpiredAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt.Get()
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportJob) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiredAt.Get(), o.ExpiredAt.IsSet()
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *ExportJob) HasExpiredAt() bool {
	if o != nil && o.ExpiredAt.IsSet() {
		return true
	}

	return false
}

// SetExpiredAt gets a reference to the given NullableTime and assigns it to the ExpiredAt field.
func (o *ExportJob) SetExpiredAt(v time.Time) {
	o.ExpiredAt.Set(&v)
}

// SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil
func (o *ExportJob) SetExpiredAtNil() {
	o.ExpiredAt.Set(nil)
}

// UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
func (o *ExportJob) UnsetExpiredAt() {
	o.ExpiredAt.Unset()
}

// GetExpiresIn returns the ExpiresIn field value
func (o *ExportJob) GetExpiresIn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value
// and a boolean to check if the value has been set.
func (o *ExportJob) GetExpiresInOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresIn, true
}

// SetExpiresIn sets field value
func (o *ExportJob) SetExpiresIn(v int32) {
	o.ExpiresIn = v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportJob) GetFailureReason() string {
	if o == nil || o.FailureReason.Get() == nil {
		var ret string
		return ret
	}
	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportJob) GetFailureReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// HasFailureReason returns a boolean if a field has been set.
func (o *ExportJob) HasFailureReason() bool {
	if o != nil && o.FailureReason.IsSet() {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given NullableString and assigns it to the FailureReason field.
func (o *ExportJob) SetFailureReason(v string) {
	o.FailureReason.Set(&v)
}

// SetFailureReasonNil sets the value for FailureReason to be an explicit nil
func (o *ExportJob) SetFailureReasonNil() {
	o.FailureReason.Set(nil)
}

// UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
func (o *ExportJob) UnsetFailureReason() {
	o.FailureReason.Unset()
}

// GetFailureReasonType returns the FailureReasonType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportJob) GetFailureReasonType() string {
	if o == nil || o.FailureReasonType.Get() == nil {
		var ret string
		return ret
	}
	return *o.FailureReasonType.Get()
}

// GetFailureReasonTypeOk returns a tuple with the FailureReasonType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportJob) GetFailureReasonTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureReasonType.Get(), o.FailureReasonType.IsSet()
}

// HasFailureReasonType returns a boolean if a field has been set.
func (o *ExportJob) HasFailureReasonType() bool {
	if o != nil && o.FailureReasonType.IsSet() {
		return true
	}

	return false
}

// SetFailureReasonType gets a reference to the given NullableString and assigns it to the FailureReasonType field.
func (o *ExportJob) SetFailureReasonType(v string) {
	o.FailureReasonType.Set(&v)
}

// SetFailureReasonTypeNil sets the value for FailureReasonType to be an explicit nil
func (o *ExportJob) SetFailureReasonTypeNil() {
	o.FailureReasonType.Set(nil)
}

// UnsetFailureReasonType ensures that no value is present for FailureReasonType, not even an explicit nil
func (o *ExportJob) UnsetFailureReasonType() {
	o.FailureReasonType.Unset()
}

// GetId returns the Id field value
func (o *ExportJob) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExportJob) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExportJob) SetId(v string) {
	o.Id = v
}

// GetIsInteractive returns the IsInteractive field value
func (o *ExportJob) GetIsInteractive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsInteractive
}

// GetIsInteractiveOk returns a tuple with the IsInteractive field value
// and a boolean to check if the value has been set.
func (o *ExportJob) GetIsInteractiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsInteractive, true
}

// SetIsInteractive sets field value
func (o *ExportJob) SetIsInteractive(v bool) {
	o.IsInteractive = v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportJob) GetLastUpdatedAt() time.Time {
	if o == nil || o.LastUpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedAt.Get()
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportJob) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdatedAt.Get(), o.LastUpdatedAt.IsSet()
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *ExportJob) HasLastUpdatedAt() bool {
	if o != nil && o.LastUpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetLastUpdatedAt gets a reference to the given NullableTime and assigns it to the LastUpdatedAt field.
func (o *ExportJob) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt.Set(&v)
}

// SetLastUpdatedAtNil sets the value for LastUpdatedAt to be an explicit nil
func (o *ExportJob) SetLastUpdatedAtNil() {
	o.LastUpdatedAt.Set(nil)
}

// UnsetLastUpdatedAt ensures that no value is present for LastUpdatedAt, not even an explicit nil
func (o *ExportJob) UnsetLastUpdatedAt() {
	o.LastUpdatedAt.Unset()
}

// GetNumberOfItems returns the NumberOfItems field value
func (o *ExportJob) GetNumberOfItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfItems
}

// GetNumberOfItemsOk returns a tuple with the NumberOfItems field value
// and a boolean to check if the value has been set.
func (o *ExportJob) GetNumberOfItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfItems, true
}

// SetNumberOfItems sets field value
func (o *ExportJob) SetNumberOfItems(v int32) {
	o.NumberOfItems = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportJob) GetStartedAt() time.Time {
	if o == nil || o.StartedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportJob) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// HasStartedAt returns a boolean if a field has been set.
func (o *ExportJob) HasStartedAt() bool {
	if o != nil && o.StartedAt.IsSet() {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given NullableTime and assigns it to the StartedAt field.
func (o *ExportJob) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}

// SetStartedAtNil sets the value for StartedAt to be an explicit nil
func (o *ExportJob) SetStartedAtNil() {
	o.StartedAt.Set(nil)
}

// UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
func (o *ExportJob) UnsetStartedAt() {
	o.StartedAt.Unset()
}

// GetStatus returns the Status field value
func (o *ExportJob) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExportJob) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ExportJob) SetStatus(v string) {
	o.Status = v
}

func (o ExportJob) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.CompletedAt.IsSet() {
		toSerialize["completedAt"] = o.CompletedAt.Get()
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.ExpiredAt.IsSet() {
		toSerialize["expiredAt"] = o.ExpiredAt.Get()
	}
	if true {
		toSerialize["expiresIn"] = o.ExpiresIn
	}
	if o.FailureReason.IsSet() {
		toSerialize["failureReason"] = o.FailureReason.Get()
	}
	if o.FailureReasonType.IsSet() {
		toSerialize["failureReasonType"] = o.FailureReasonType.Get()
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["isInteractive"] = o.IsInteractive
	}
	if o.LastUpdatedAt.IsSet() {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt.Get()
	}
	if true {
		toSerialize["numberOfItems"] = o.NumberOfItems
	}
	if o.StartedAt.IsSet() {
		toSerialize["startedAt"] = o.StartedAt.Get()
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableExportJob struct {
	value *ExportJob
	isSet bool
}

func (v NullableExportJob) Get() *ExportJob {
	return v.value
}

func (v *NullableExportJob) Set(val *ExportJob) {
	v.value = val
	v.isSet = true
}

func (v NullableExportJob) IsSet() bool {
	return v.isSet
}

func (v *NullableExportJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportJob(val *ExportJob) *NullableExportJob {
	return &NullableExportJob{value: val, isSet: true}
}

func (v NullableExportJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
