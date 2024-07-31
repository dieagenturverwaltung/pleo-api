# UpdateExportJobItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountingEntryId** | **string** | ID of the accounting entry being updated | 
**ExportedAt** | Pointer to **time.Time** | Date and time of the export attempt | [optional] 
**ExternalId** | Pointer to **string** | The accounting system identifier of the entry after export | [optional] 
**ExternalUrl** | Pointer to **string** | URL to access the resource of the entry in the accounting system | [optional] 
**FailureReason** | Pointer to **string** | Detailed message explaining the failure | [optional] 
**FailureReasonType** | Pointer to **string** | If the export of this accounting entry failed, specify the failure reason type | [optional] 
**Status** | **string** | Status of the export Item after being processed | 

## Methods

### NewUpdateExportJobItem

`func NewUpdateExportJobItem(accountingEntryId string, status string, ) *UpdateExportJobItem`

NewUpdateExportJobItem instantiates a new UpdateExportJobItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExportJobItemWithDefaults

`func NewUpdateExportJobItemWithDefaults() *UpdateExportJobItem`

NewUpdateExportJobItemWithDefaults instantiates a new UpdateExportJobItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountingEntryId

`func (o *UpdateExportJobItem) GetAccountingEntryId() string`

GetAccountingEntryId returns the AccountingEntryId field if non-nil, zero value otherwise.

### GetAccountingEntryIdOk

`func (o *UpdateExportJobItem) GetAccountingEntryIdOk() (*string, bool)`

GetAccountingEntryIdOk returns a tuple with the AccountingEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntryId

`func (o *UpdateExportJobItem) SetAccountingEntryId(v string)`

SetAccountingEntryId sets AccountingEntryId field to given value.


### GetExportedAt

`func (o *UpdateExportJobItem) GetExportedAt() time.Time`

GetExportedAt returns the ExportedAt field if non-nil, zero value otherwise.

### GetExportedAtOk

`func (o *UpdateExportJobItem) GetExportedAtOk() (*time.Time, bool)`

GetExportedAtOk returns a tuple with the ExportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedAt

`func (o *UpdateExportJobItem) SetExportedAt(v time.Time)`

SetExportedAt sets ExportedAt field to given value.

### HasExportedAt

`func (o *UpdateExportJobItem) HasExportedAt() bool`

HasExportedAt returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateExportJobItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateExportJobItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateExportJobItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateExportJobItem) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalUrl

`func (o *UpdateExportJobItem) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *UpdateExportJobItem) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *UpdateExportJobItem) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *UpdateExportJobItem) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.

### GetFailureReason

`func (o *UpdateExportJobItem) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *UpdateExportJobItem) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *UpdateExportJobItem) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *UpdateExportJobItem) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetFailureReasonType

`func (o *UpdateExportJobItem) GetFailureReasonType() string`

GetFailureReasonType returns the FailureReasonType field if non-nil, zero value otherwise.

### GetFailureReasonTypeOk

`func (o *UpdateExportJobItem) GetFailureReasonTypeOk() (*string, bool)`

GetFailureReasonTypeOk returns a tuple with the FailureReasonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReasonType

`func (o *UpdateExportJobItem) SetFailureReasonType(v string)`

SetFailureReasonType sets FailureReasonType field to given value.

### HasFailureReasonType

`func (o *UpdateExportJobItem) HasFailureReasonType() bool`

HasFailureReasonType returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateExportJobItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateExportJobItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateExportJobItem) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


