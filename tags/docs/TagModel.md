# TagModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | This tag is not used anymore | 
**Code** | **string** | External identifier of the Tag | 
**CreatedAt** | **time.Time** | Creation date and time | 
**GroupId** | **string** | Unique identifier of the Tag Group this tag belongs to | 
**Id** | **string** | Unique identifier of Tag | 
**Name** | Pointer to **string** | User readable name that is used for the possible value within a tag group on an expense | [optional] 
**UpdatedAt** | **time.Time** | Date and time of the last update | 

## Methods

### NewTagModel

`func NewTagModel(archived bool, code string, createdAt time.Time, groupId string, id string, updatedAt time.Time, ) *TagModel`

NewTagModel instantiates a new TagModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagModelWithDefaults

`func NewTagModelWithDefaults() *TagModel`

NewTagModelWithDefaults instantiates a new TagModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *TagModel) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *TagModel) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *TagModel) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetCode

`func (o *TagModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TagModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TagModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetCreatedAt

`func (o *TagModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TagModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TagModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetGroupId

`func (o *TagModel) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *TagModel) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *TagModel) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetId

`func (o *TagModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TagModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TagModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TagModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TagModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TagModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


