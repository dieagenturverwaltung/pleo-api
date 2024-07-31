# AggregatedTagGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | This tag group is not used anymore | 
**Code** | **string** | External identifier of the Tag group / Dimension used for mapping to accounting system | 
**CompanyId** | **string** | Unique identifier of the company the Tag Group belongs to | 
**CreatedAt** | **time.Time** | Creation date and time | 
**Dimensions** | [**[]AggregatedTagGroupDimensionModel**](AggregatedTagGroupDimensionModel.md) | List of all the dimensions associated with this tag group | 
**Id** | **string** | Unique identifier of Tag Group (generated on creation) | 
**Name** | **string** | User readable name of Tag Group | 
**UpdatedAt** | **time.Time** | Date and time of the last update | 

## Methods

### NewAggregatedTagGroupModel

`func NewAggregatedTagGroupModel(archived bool, code string, companyId string, createdAt time.Time, dimensions []AggregatedTagGroupDimensionModel, id string, name string, updatedAt time.Time, ) *AggregatedTagGroupModel`

NewAggregatedTagGroupModel instantiates a new AggregatedTagGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregatedTagGroupModelWithDefaults

`func NewAggregatedTagGroupModelWithDefaults() *AggregatedTagGroupModel`

NewAggregatedTagGroupModelWithDefaults instantiates a new AggregatedTagGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *AggregatedTagGroupModel) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *AggregatedTagGroupModel) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *AggregatedTagGroupModel) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetCode

`func (o *AggregatedTagGroupModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AggregatedTagGroupModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AggregatedTagGroupModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetCompanyId

`func (o *AggregatedTagGroupModel) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *AggregatedTagGroupModel) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *AggregatedTagGroupModel) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetCreatedAt

`func (o *AggregatedTagGroupModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AggregatedTagGroupModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AggregatedTagGroupModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDimensions

`func (o *AggregatedTagGroupModel) GetDimensions() []AggregatedTagGroupDimensionModel`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *AggregatedTagGroupModel) GetDimensionsOk() (*[]AggregatedTagGroupDimensionModel, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *AggregatedTagGroupModel) SetDimensions(v []AggregatedTagGroupDimensionModel)`

SetDimensions sets Dimensions field to given value.


### GetId

`func (o *AggregatedTagGroupModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AggregatedTagGroupModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AggregatedTagGroupModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AggregatedTagGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AggregatedTagGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AggregatedTagGroupModel) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *AggregatedTagGroupModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AggregatedTagGroupModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AggregatedTagGroupModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


