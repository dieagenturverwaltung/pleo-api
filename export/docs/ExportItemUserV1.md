# ExportItemUserV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** | This would usually represent the employee code | [optional] 
**Id** | **NullableString** | This is the Pleo internal identifier of the user | 
**Name** | **string** | Full name of the user | 

## Methods

### NewExportItemUserV1

`func NewExportItemUserV1(id NullableString, name string, ) *ExportItemUserV1`

NewExportItemUserV1 instantiates a new ExportItemUserV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportItemUserV1WithDefaults

`func NewExportItemUserV1WithDefaults() *ExportItemUserV1`

NewExportItemUserV1WithDefaults instantiates a new ExportItemUserV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ExportItemUserV1) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExportItemUserV1) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExportItemUserV1) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ExportItemUserV1) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ExportItemUserV1) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ExportItemUserV1) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetId

`func (o *ExportItemUserV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExportItemUserV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExportItemUserV1) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ExportItemUserV1) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExportItemUserV1) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ExportItemUserV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExportItemUserV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExportItemUserV1) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


