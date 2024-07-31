# CreateExportJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountingEntryIds** | **[]string** | List of accounting entry IDs the job is being created for | 
**CompanyId** | **string** | Company ID the export job is being created for | 
**EmployeeId** | Pointer to **NullableString** | This is the Pleo unique identifier of the user that initiated the export job | [optional] 
**IsInteractive** | Pointer to **NullableBool** | Non-Interactive jobs are jobs that have not been initiated by a user. These jobs are usually triggered in the background and required no user interaction. Whereas interactive jobs are the opposite. This flag should be set to true if the Job is interactive and if not set to false. By default this is set to true, i.e. by default jobs are deemed to be interactive, i.e. initiated by a user. | [optional] [default to true]

## Methods

### NewCreateExportJob

`func NewCreateExportJob(accountingEntryIds []string, companyId string, ) *CreateExportJob`

NewCreateExportJob instantiates a new CreateExportJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExportJobWithDefaults

`func NewCreateExportJobWithDefaults() *CreateExportJob`

NewCreateExportJobWithDefaults instantiates a new CreateExportJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountingEntryIds

`func (o *CreateExportJob) GetAccountingEntryIds() []string`

GetAccountingEntryIds returns the AccountingEntryIds field if non-nil, zero value otherwise.

### GetAccountingEntryIdsOk

`func (o *CreateExportJob) GetAccountingEntryIdsOk() (*[]string, bool)`

GetAccountingEntryIdsOk returns a tuple with the AccountingEntryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntryIds

`func (o *CreateExportJob) SetAccountingEntryIds(v []string)`

SetAccountingEntryIds sets AccountingEntryIds field to given value.


### GetCompanyId

`func (o *CreateExportJob) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CreateExportJob) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CreateExportJob) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetEmployeeId

`func (o *CreateExportJob) GetEmployeeId() string`

GetEmployeeId returns the EmployeeId field if non-nil, zero value otherwise.

### GetEmployeeIdOk

`func (o *CreateExportJob) GetEmployeeIdOk() (*string, bool)`

GetEmployeeIdOk returns a tuple with the EmployeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId

`func (o *CreateExportJob) SetEmployeeId(v string)`

SetEmployeeId sets EmployeeId field to given value.

### HasEmployeeId

`func (o *CreateExportJob) HasEmployeeId() bool`

HasEmployeeId returns a boolean if a field has been set.

### SetEmployeeIdNil

`func (o *CreateExportJob) SetEmployeeIdNil(b bool)`

 SetEmployeeIdNil sets the value for EmployeeId to be an explicit nil

### UnsetEmployeeId
`func (o *CreateExportJob) UnsetEmployeeId()`

UnsetEmployeeId ensures that no value is present for EmployeeId, not even an explicit nil
### GetIsInteractive

`func (o *CreateExportJob) GetIsInteractive() bool`

GetIsInteractive returns the IsInteractive field if non-nil, zero value otherwise.

### GetIsInteractiveOk

`func (o *CreateExportJob) GetIsInteractiveOk() (*bool, bool)`

GetIsInteractiveOk returns a tuple with the IsInteractive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInteractive

`func (o *CreateExportJob) SetIsInteractive(v bool)`

SetIsInteractive sets IsInteractive field to given value.

### HasIsInteractive

`func (o *CreateExportJob) HasIsInteractive() bool`

HasIsInteractive returns a boolean if a field has been set.

### SetIsInteractiveNil

`func (o *CreateExportJob) SetIsInteractiveNil(b bool)`

 SetIsInteractiveNil sets the value for IsInteractive to be an explicit nil

### UnsetIsInteractive
`func (o *CreateExportJob) UnsetIsInteractive()`

UnsetIsInteractive ensures that no value is present for IsInteractive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


