# TagDimensionValueModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Creation date and time | 
**DimensionId** | **string** | Unique identifier of the Tag Group Dimension this dimension value is associated with | 
**TagId** | **string** | Unique identifier of the Tag this dimension value belongs to | 
**UpdatedAt** | **time.Time** | Date and time of the last update | 
**Value** | **string** | Value of the given dimension for the given tag | 

## Methods

### NewTagDimensionValueModel

`func NewTagDimensionValueModel(createdAt time.Time, dimensionId string, tagId string, updatedAt time.Time, value string, ) *TagDimensionValueModel`

NewTagDimensionValueModel instantiates a new TagDimensionValueModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagDimensionValueModelWithDefaults

`func NewTagDimensionValueModelWithDefaults() *TagDimensionValueModel`

NewTagDimensionValueModelWithDefaults instantiates a new TagDimensionValueModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TagDimensionValueModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TagDimensionValueModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TagDimensionValueModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDimensionId

`func (o *TagDimensionValueModel) GetDimensionId() string`

GetDimensionId returns the DimensionId field if non-nil, zero value otherwise.

### GetDimensionIdOk

`func (o *TagDimensionValueModel) GetDimensionIdOk() (*string, bool)`

GetDimensionIdOk returns a tuple with the DimensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionId

`func (o *TagDimensionValueModel) SetDimensionId(v string)`

SetDimensionId sets DimensionId field to given value.


### GetTagId

`func (o *TagDimensionValueModel) GetTagId() string`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *TagDimensionValueModel) GetTagIdOk() (*string, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *TagDimensionValueModel) SetTagId(v string)`

SetTagId sets TagId field to given value.


### GetUpdatedAt

`func (o *TagDimensionValueModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TagDimensionValueModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TagDimensionValueModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetValue

`func (o *TagDimensionValueModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TagDimensionValueModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TagDimensionValueModel) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


