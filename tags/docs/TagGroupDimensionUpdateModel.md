# TagGroupDimensionUpdateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | External identifier of the Tag group Dimension used for mapping to accounting system | 
**DisplayOrder** | **int32** | Determines the order in which the dimensions are displayed in the UI. Starts from 1 | 
**Name** | **string** | User readable name of Tag Group Dimension | 
**Visible** | **bool** | Determines if this dimension is displayed in the UI | 

## Methods

### NewTagGroupDimensionUpdateModel

`func NewTagGroupDimensionUpdateModel(code string, displayOrder int32, name string, visible bool, ) *TagGroupDimensionUpdateModel`

NewTagGroupDimensionUpdateModel instantiates a new TagGroupDimensionUpdateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagGroupDimensionUpdateModelWithDefaults

`func NewTagGroupDimensionUpdateModelWithDefaults() *TagGroupDimensionUpdateModel`

NewTagGroupDimensionUpdateModelWithDefaults instantiates a new TagGroupDimensionUpdateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TagGroupDimensionUpdateModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TagGroupDimensionUpdateModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TagGroupDimensionUpdateModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetDisplayOrder

`func (o *TagGroupDimensionUpdateModel) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *TagGroupDimensionUpdateModel) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *TagGroupDimensionUpdateModel) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.


### GetName

`func (o *TagGroupDimensionUpdateModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagGroupDimensionUpdateModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagGroupDimensionUpdateModel) SetName(v string)`

SetName sets Name field to given value.


### GetVisible

`func (o *TagGroupDimensionUpdateModel) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *TagGroupDimensionUpdateModel) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *TagGroupDimensionUpdateModel) SetVisible(v bool)`

SetVisible sets Visible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


