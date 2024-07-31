# AggregatedTagModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | This tag is not used anymore | 
**Code** | **string** | External identifier of the Tag | 
**CompanyId** | **string** | Unique identifier of the company the Tag Group belongs to | 
**CreatedAt** | **time.Time** | Creation date and time | 
**Dimensions** | [**[]AggregatedTagTagDimensionModel**](AggregatedTagTagDimensionModel.md) | List of all the dimensions and dimension values associated with the tag | 
**Group** | [**AggregatedTagTagGroupModel**](AggregatedTagTagGroupModel.md) |  | 
**Id** | **string** | Unique identifier of Tag | 
**Name** | Pointer to **string** | User readable name that is used for the possible value within a tag group on an expense | [optional] 
**UpdatedAt** | **time.Time** | Date and time of the last update | 

## Methods

### NewAggregatedTagModel

`func NewAggregatedTagModel(archived bool, code string, companyId string, createdAt time.Time, dimensions []AggregatedTagTagDimensionModel, group AggregatedTagTagGroupModel, id string, updatedAt time.Time, ) *AggregatedTagModel`

NewAggregatedTagModel instantiates a new AggregatedTagModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregatedTagModelWithDefaults

`func NewAggregatedTagModelWithDefaults() *AggregatedTagModel`

NewAggregatedTagModelWithDefaults instantiates a new AggregatedTagModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *AggregatedTagModel) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *AggregatedTagModel) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *AggregatedTagModel) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetCode

`func (o *AggregatedTagModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AggregatedTagModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AggregatedTagModel) SetCode(v string)`

SetCode sets Code field to given value.


### GetCompanyId

`func (o *AggregatedTagModel) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *AggregatedTagModel) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *AggregatedTagModel) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetCreatedAt

`func (o *AggregatedTagModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AggregatedTagModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AggregatedTagModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDimensions

`func (o *AggregatedTagModel) GetDimensions() []AggregatedTagTagDimensionModel`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *AggregatedTagModel) GetDimensionsOk() (*[]AggregatedTagTagDimensionModel, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *AggregatedTagModel) SetDimensions(v []AggregatedTagTagDimensionModel)`

SetDimensions sets Dimensions field to given value.


### GetGroup

`func (o *AggregatedTagModel) GetGroup() AggregatedTagTagGroupModel`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AggregatedTagModel) GetGroupOk() (*AggregatedTagTagGroupModel, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AggregatedTagModel) SetGroup(v AggregatedTagTagGroupModel)`

SetGroup sets Group field to given value.


### GetId

`func (o *AggregatedTagModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AggregatedTagModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AggregatedTagModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AggregatedTagModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AggregatedTagModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AggregatedTagModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AggregatedTagModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AggregatedTagModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AggregatedTagModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AggregatedTagModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


