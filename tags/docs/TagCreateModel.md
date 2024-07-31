# TagCreateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | This tag is not used anymore | 
**Code** | **string** | External identifier of the Tag | 
**Name** | Pointer to **string** | User readable name that is used for the possible value within a tag group on an expense | [optional] 

## Methods

### NewTagCreateModel

`func NewTagCreateModel(archived bool, code string, ) *TagCreateModel`

NewTagCreateModel instantiates a new TagCreateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagCreateModelWithDefaults

`func NewTagCreateModelWithDefaults() *TagCreateModel`

NewTagCreateModelWithDefaults instantiates a new TagCreateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *TagCreateModel) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *TagCreateModel) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *TagCreateModel) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetCode

`func (o *TagCreateModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TagCreateModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TagCreateModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *TagCreateModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagCreateModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagCreateModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TagCreateModel) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


