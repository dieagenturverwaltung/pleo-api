# TagGroupDimensionCreateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | External identifier of the Tag group Dimension used for mapping to accounting system | 
**DisplayOrder** | **int32** | Determines the order in which the dimensions are displayed in the UI. Starts from 1 | 
**Name** | **string** | User readable name of Tag Group Dimension | 
**Visible** | **bool** | Determines if this dimension is displayed in the UI | 

## Methods

### NewTagGroupDimensionCreateModel

`func NewTagGroupDimensionCreateModel(code string, displayOrder int32, name string, visible bool, ) *TagGroupDimensionCreateModel`

NewTagGroupDimensionCreateModel instantiates a new TagGroupDimensionCreateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagGroupDimensionCreateModelWithDefaults

`func NewTagGroupDimensionCreateModelWithDefaults() *TagGroupDimensionCreateModel`

NewTagGroupDimensionCreateModelWithDefaults instantiates a new TagGroupDimensionCreateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TagGroupDimensionCreateModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TagGroupDimensionCreateModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TagGroupDimensionCreateModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetDisplayOrder

`func (o *TagGroupDimensionCreateModel) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *TagGroupDimensionCreateModel) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *TagGroupDimensionCreateModel) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.


### GetName

`func (o *TagGroupDimensionCreateModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagGroupDimensionCreateModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagGroupDimensionCreateModel) SetName(v string)`

SetName sets Name field to given value.


### GetVisible

`func (o *TagGroupDimensionCreateModel) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *TagGroupDimensionCreateModel) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *TagGroupDimensionCreateModel) SetVisible(v bool)`

SetVisible sets Visible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


