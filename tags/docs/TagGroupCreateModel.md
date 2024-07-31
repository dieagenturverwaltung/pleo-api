# TagGroupCreateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** |  | 
**Code** | **string** |  | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewTagGroupCreateModel

`func NewTagGroupCreateModel(archived bool, code string, name string, ) *TagGroupCreateModel`

NewTagGroupCreateModel instantiates a new TagGroupCreateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagGroupCreateModelWithDefaults

`func NewTagGroupCreateModelWithDefaults() *TagGroupCreateModel`

NewTagGroupCreateModelWithDefaults instantiates a new TagGroupCreateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *TagGroupCreateModel) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *TagGroupCreateModel) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *TagGroupCreateModel) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetCode

`func (o *TagGroupCreateModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TagGroupCreateModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TagGroupCreateModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetMetadata

`func (o *TagGroupCreateModel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TagGroupCreateModel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TagGroupCreateModel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TagGroupCreateModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *TagGroupCreateModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagGroupCreateModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagGroupCreateModel) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


