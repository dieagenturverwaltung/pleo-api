# InstallationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | The unique identifier of the application this installation is for | [readonly] 
**CreatedAt** | **time.Time** | The date and time the installation was created | [readonly] 
**ErrorCode** | Pointer to **string** | The text to indicate why installation is inactive | [optional] [readonly] 
**Id** | **string** | The unique identifier of the installation | [readonly] 
**Metadata** | **map[string]map[string]interface{}** | The metadata of the installation | 
**Resource** | **string** | The resource that the installation is for | [readonly] 
**Status** | **string** | The status of the installation | 
**UpdatedAt** | **time.Time** | The date and time the installation was last updated | [readonly] 

## Methods

### NewInstallationResponse

`func NewInstallationResponse(applicationId string, createdAt time.Time, id string, metadata map[string]map[string]interface{}, resource string, status string, updatedAt time.Time, ) *InstallationResponse`

NewInstallationResponse instantiates a new InstallationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallationResponseWithDefaults

`func NewInstallationResponseWithDefaults() *InstallationResponse`

NewInstallationResponseWithDefaults instantiates a new InstallationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *InstallationResponse) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *InstallationResponse) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *InstallationResponse) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetCreatedAt

`func (o *InstallationResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstallationResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstallationResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetErrorCode

`func (o *InstallationResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *InstallationResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *InstallationResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *InstallationResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetId

`func (o *InstallationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstallationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstallationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *InstallationResponse) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InstallationResponse) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InstallationResponse) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetResource

`func (o *InstallationResponse) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *InstallationResponse) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *InstallationResponse) SetResource(v string)`

SetResource sets Resource field to given value.


### GetStatus

`func (o *InstallationResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstallationResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstallationResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *InstallationResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InstallationResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InstallationResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


