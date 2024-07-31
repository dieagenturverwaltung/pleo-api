# TagGroupDimensionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | External identifier of the Tag group Dimension used for mapping to accounting system | 
**CreatedAt** | **time.Time** | Creation date and time | 
**DisplayOrder** | **int32** | Determines the order in which the dimensions are displayed in the UI. Starts from 1 | 
**GroupId** | **string** | Unique identifier of the Tag Group this dimension belongs to | 
**Id** | **string** | Unique identifier of Tag Group Dimension (generated on creation) | 
**Name** | **string** | User readable name of Tag Group Dimension | 
**UpdatedAt** | **time.Time** | Date and time of the last update | 
**Visible** | **bool** | Determines if this dimension is displayed in the UI | 

## Methods

### NewTagGroupDimensionModel

`func NewTagGroupDimensionModel(code string, createdAt time.Time, displayOrder int32, groupId string, id string, name string, updatedAt time.Time, visible bool, ) *TagGroupDimensionModel`

NewTagGroupDimensionModel instantiates a new TagGroupDimensionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagGroupDimensionModelWithDefaults

`func NewTagGroupDimensionModelWithDefaults() *TagGroupDimensionModel`

NewTagGroupDimensionModelWithDefaults instantiates a new TagGroupDimensionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TagGroupDimensionModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TagGroupDimensionModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TagGroupDimensionModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetCreatedAt

`func (o *TagGroupDimensionModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TagGroupDimensionModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TagGroupDimensionModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDisplayOrder

`func (o *TagGroupDimensionModel) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *TagGroupDimensionModel) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *TagGroupDimensionModel) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.


### GetGroupId

`func (o *TagGroupDimensionModel) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *TagGroupDimensionModel) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *TagGroupDimensionModel) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetId

`func (o *TagGroupDimensionModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagGroupDimensionModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagGroupDimensionModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TagGroupDimensionModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagGroupDimensionModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagGroupDimensionModel) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *TagGroupDimensionModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TagGroupDimensionModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TagGroupDimensionModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVisible

`func (o *TagGroupDimensionModel) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *TagGroupDimensionModel) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *TagGroupDimensionModel) SetVisible(v bool)`

SetVisible sets Visible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


