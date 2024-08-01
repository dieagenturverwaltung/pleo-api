# CreateClientInstallationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **string** | The error code to indicate why installation is INACTIVE. Can only be set when status is INACTIVE. | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata associated with the installation. This can be used to store additional information about the installation, for example  external references, but it is not used by any Pleo system | [optional] 
**Status** | Pointer to **string** | The status of the installation. If it is ACTIVATED, it means that the application is installed and operational. | [optional] [default to "ACTIVATED"]

## Methods

### NewCreateClientInstallationRequest

`func NewCreateClientInstallationRequest() *CreateClientInstallationRequest`

NewCreateClientInstallationRequest instantiates a new CreateClientInstallationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClientInstallationRequestWithDefaults

`func NewCreateClientInstallationRequestWithDefaults() *CreateClientInstallationRequest`

NewCreateClientInstallationRequestWithDefaults instantiates a new CreateClientInstallationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *CreateClientInstallationRequest) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CreateClientInstallationRequest) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CreateClientInstallationRequest) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *CreateClientInstallationRequest) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateClientInstallationRequest) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateClientInstallationRequest) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateClientInstallationRequest) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateClientInstallationRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStatus

`func (o *CreateClientInstallationRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateClientInstallationRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateClientInstallationRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateClientInstallationRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


