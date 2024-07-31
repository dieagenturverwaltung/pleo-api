# ExportJobItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountingEntryId** | **string** | The accounting entry identifier being exported | 
**ExportJobId** | **string** | The export job identifier for which this item belongs to | 
**ExportedAt** | Pointer to **NullableTime** | Date and time the item was exported | [optional] 
**ExternalId** | Pointer to **NullableString** | This is an external identifier of the corresponding accounting entry in the accounting system after export | [optional] 
**ExternalUrl** | Pointer to **NullableString** | This is the external URL pointing to the accounting entry resource in the accounting system after the export | [optional] 
**FailureReason** | Pointer to **NullableString** | Reason why the export of this item failed in the case of a failure | [optional] 
**FailureReasonType** | Pointer to **NullableString** | The classification for the failure from a list of described failure reason types | [optional] 
**Status** | **string** | Status of the export job Item after being processed. The status of the export job item will remain null until the item an attempt has been made to process the export job item or the job is marked as failed or expired, in which case the status will then be abandoned. The status of an export job item that has been processed will either be failed or successful. | 

## Methods

### NewExportJobItem

`func NewExportJobItem(accountingEntryId string, exportJobId string, status string, ) *ExportJobItem`

NewExportJobItem instantiates a new ExportJobItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportJobItemWithDefaults

`func NewExportJobItemWithDefaults() *ExportJobItem`

NewExportJobItemWithDefaults instantiates a new ExportJobItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountingEntryId

`func (o *ExportJobItem) GetAccountingEntryId() string`

GetAccountingEntryId returns the AccountingEntryId field if non-nil, zero value otherwise.

### GetAccountingEntryIdOk

`func (o *ExportJobItem) GetAccountingEntryIdOk() (*string, bool)`

GetAccountingEntryIdOk returns a tuple with the AccountingEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntryId

`func (o *ExportJobItem) SetAccountingEntryId(v string)`

SetAccountingEntryId sets AccountingEntryId field to given value.


### GetExportJobId

`func (o *ExportJobItem) GetExportJobId() string`

GetExportJobId returns the ExportJobId field if non-nil, zero value otherwise.

### GetExportJobIdOk

`func (o *ExportJobItem) GetExportJobIdOk() (*string, bool)`

GetExportJobIdOk returns a tuple with the ExportJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportJobId

`func (o *ExportJobItem) SetExportJobId(v string)`

SetExportJobId sets ExportJobId field to given value.


### GetExportedAt

`func (o *ExportJobItem) GetExportedAt() time.Time`

GetExportedAt returns the ExportedAt field if non-nil, zero value otherwise.

### GetExportedAtOk

`func (o *ExportJobItem) GetExportedAtOk() (*time.Time, bool)`

GetExportedAtOk returns a tuple with the ExportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedAt

`func (o *ExportJobItem) SetExportedAt(v time.Time)`

SetExportedAt sets ExportedAt field to given value.

### HasExportedAt

`func (o *ExportJobItem) HasExportedAt() bool`

HasExportedAt returns a boolean if a field has been set.

### SetExportedAtNil

`func (o *ExportJobItem) SetExportedAtNil(b bool)`

 SetExportedAtNil sets the value for ExportedAt to be an explicit nil

### UnsetExportedAt
`func (o *ExportJobItem) UnsetExportedAt()`

UnsetExportedAt ensures that no value is present for ExportedAt, not even an explicit nil
### GetExternalId

`func (o *ExportJobItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ExportJobItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ExportJobItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ExportJobItem) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ExportJobItem) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ExportJobItem) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalUrl

`func (o *ExportJobItem) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *ExportJobItem) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *ExportJobItem) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *ExportJobItem) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.

### SetExternalUrlNil

`func (o *ExportJobItem) SetExternalUrlNil(b bool)`

 SetExternalUrlNil sets the value for ExternalUrl to be an explicit nil

### UnsetExternalUrl
`func (o *ExportJobItem) UnsetExternalUrl()`

UnsetExternalUrl ensures that no value is present for ExternalUrl, not even an explicit nil
### GetFailureReason

`func (o *ExportJobItem) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ExportJobItem) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ExportJobItem) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *ExportJobItem) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *ExportJobItem) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *ExportJobItem) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetFailureReasonType

`func (o *ExportJobItem) GetFailureReasonType() string`

GetFailureReasonType returns the FailureReasonType field if non-nil, zero value otherwise.

### GetFailureReasonTypeOk

`func (o *ExportJobItem) GetFailureReasonTypeOk() (*string, bool)`

GetFailureReasonTypeOk returns a tuple with the FailureReasonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReasonType

`func (o *ExportJobItem) SetFailureReasonType(v string)`

SetFailureReasonType sets FailureReasonType field to given value.

### HasFailureReasonType

`func (o *ExportJobItem) HasFailureReasonType() bool`

HasFailureReasonType returns a boolean if a field has been set.

### SetFailureReasonTypeNil

`func (o *ExportJobItem) SetFailureReasonTypeNil(b bool)`

 SetFailureReasonTypeNil sets the value for FailureReasonType to be an explicit nil

### UnsetFailureReasonType
`func (o *ExportJobItem) UnsetFailureReasonType()`

UnsetFailureReasonType ensures that no value is present for FailureReasonType, not even an explicit nil
### GetStatus

`func (o *ExportJobItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExportJobItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExportJobItem) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


