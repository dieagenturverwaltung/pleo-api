# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** | Team code, usually the unique identifier of the team in the accounting system | [optional] 
**Id** | **string** | This is the Pleo internal identifier of the team | 
**Name** | **string** | This is the name of the team | 

## Methods

### NewTeam

`func NewTeam(id string, name string, ) *Team`

NewTeam instantiates a new Team object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamWithDefaults

`func NewTeamWithDefaults() *Team`

NewTeamWithDefaults instantiates a new Team object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Team) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Team) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Team) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Team) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Team) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Team) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetId

`func (o *Team) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Team) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Team) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Team) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Team) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Team) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


