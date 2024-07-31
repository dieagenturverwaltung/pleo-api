# CursorPaginatedResponseTagModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TagModel**](TagModel.md) |  | 
**Pagination** | [**CursorPageInfo**](CursorPageInfo.md) |  | 

## Methods

### NewCursorPaginatedResponseTagModel

`func NewCursorPaginatedResponseTagModel(data []TagModel, pagination CursorPageInfo, ) *CursorPaginatedResponseTagModel`

NewCursorPaginatedResponseTagModel instantiates a new CursorPaginatedResponseTagModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPaginatedResponseTagModelWithDefaults

`func NewCursorPaginatedResponseTagModelWithDefaults() *CursorPaginatedResponseTagModel`

NewCursorPaginatedResponseTagModelWithDefaults instantiates a new CursorPaginatedResponseTagModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CursorPaginatedResponseTagModel) GetData() []TagModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CursorPaginatedResponseTagModel) GetDataOk() (*[]TagModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CursorPaginatedResponseTagModel) SetData(v []TagModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *CursorPaginatedResponseTagModel) GetPagination() CursorPageInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CursorPaginatedResponseTagModel) GetPaginationOk() (*CursorPageInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CursorPaginatedResponseTagModel) SetPagination(v CursorPageInfo)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


