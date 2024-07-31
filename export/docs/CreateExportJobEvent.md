# CreateExportJobEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** |  | 
**FailureReason** | Pointer to **string** | Reason why the job failed in the case of a failure | [optional] 
**FailureReasonType** | Pointer to **string** | The classification for the failure from a list of described failure reason types | [optional] 
**JobId** | **string** | Identifier of the job the event is to be processed for | 

## Methods

### NewCreateExportJobEvent

`func NewCreateExportJobEvent(event string, jobId string, ) *CreateExportJobEvent`

NewCreateExportJobEvent instantiates a new CreateExportJobEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExportJobEventWithDefaults

`func NewCreateExportJobEventWithDefaults() *CreateExportJobEvent`

NewCreateExportJobEventWithDefaults instantiates a new CreateExportJobEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *CreateExportJobEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CreateExportJobEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CreateExportJobEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetFailureReason

`func (o *CreateExportJobEvent) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *CreateExportJobEvent) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *CreateExportJobEvent) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *CreateExportJobEvent) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetFailureReasonType

`func (o *CreateExportJobEvent) GetFailureReasonType() string`

GetFailureReasonType returns the FailureReasonType field if non-nil, zero value otherwise.

### GetFailureReasonTypeOk

`func (o *CreateExportJobEvent) GetFailureReasonTypeOk() (*string, bool)`

GetFailureReasonTypeOk returns a tuple with the FailureReasonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReasonType

`func (o *CreateExportJobEvent) SetFailureReasonType(v string)`

SetFailureReasonType sets FailureReasonType field to given value.

### HasFailureReasonType

`func (o *CreateExportJobEvent) HasFailureReasonType() bool`

HasFailureReasonType returns a boolean if a field has been set.

### GetJobId

`func (o *CreateExportJobEvent) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *CreateExportJobEvent) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *CreateExportJobEvent) SetJobId(v string)`

SetJobId sets JobId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


