# UpdateInstallationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]map[string]interface{}** | A means to store any metadata about the installation. Could for example be settings for the application | [optional] 
**Status** | **string** | The status of the installation. If it is ACTIVATED, it means that the application is installed and operational. | [default to "ACTIVATED"]

## Methods

### NewUpdateInstallationRequest

`func NewUpdateInstallationRequest(status string, ) *UpdateInstallationRequest`

NewUpdateInstallationRequest instantiates a new UpdateInstallationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstallationRequestWithDefaults

`func NewUpdateInstallationRequestWithDefaults() *UpdateInstallationRequest`

NewUpdateInstallationRequestWithDefaults instantiates a new UpdateInstallationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *UpdateInstallationRequest) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateInstallationRequest) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateInstallationRequest) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateInstallationRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateInstallationRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateInstallationRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateInstallationRequest) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


