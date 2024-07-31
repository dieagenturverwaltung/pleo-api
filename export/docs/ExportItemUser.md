# ExportItemUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** | This would usually represent the employee code | [optional] 
**Name** | **string** | Full name of the user | 

## Methods

### NewExportItemUser

`func NewExportItemUser(name string, ) *ExportItemUser`

NewExportItemUser instantiates a new ExportItemUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportItemUserWithDefaults

`func NewExportItemUserWithDefaults() *ExportItemUser`

NewExportItemUserWithDefaults instantiates a new ExportItemUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ExportItemUser) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExportItemUser) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExportItemUser) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ExportItemUser) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ExportItemUser) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ExportItemUser) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *ExportItemUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExportItemUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExportItemUser) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


