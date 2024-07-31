# TagUpdateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | This tag is not used anymore | 
**Code** | **string** | External identifier of the Tag | 
**Name** | Pointer to **string** | User readable name that is used for the possible value within a tag group on an expense | [optional] 

## Methods

### NewTagUpdateModel

`func NewTagUpdateModel(archived bool, code string, ) *TagUpdateModel`

NewTagUpdateModel instantiates a new TagUpdateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagUpdateModelWithDefaults

`func NewTagUpdateModelWithDefaults() *TagUpdateModel`

NewTagUpdateModelWithDefaults instantiates a new TagUpdateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *TagUpdateModel) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *TagUpdateModel) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *TagUpdateModel) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetCode

`func (o *TagUpdateModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TagUpdateModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TagUpdateModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *TagUpdateModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagUpdateModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagUpdateModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TagUpdateModel) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


