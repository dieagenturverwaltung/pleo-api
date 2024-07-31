# CursorPaginatedResponseAggregatedTagModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AggregatedTagModel**](AggregatedTagModel.md) |  | 
**Pagination** | [**CursorPageInfo**](CursorPageInfo.md) |  | 

## Methods

### NewCursorPaginatedResponseAggregatedTagModel

`func NewCursorPaginatedResponseAggregatedTagModel(data []AggregatedTagModel, pagination CursorPageInfo, ) *CursorPaginatedResponseAggregatedTagModel`

NewCursorPaginatedResponseAggregatedTagModel instantiates a new CursorPaginatedResponseAggregatedTagModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPaginatedResponseAggregatedTagModelWithDefaults

`func NewCursorPaginatedResponseAggregatedTagModelWithDefaults() *CursorPaginatedResponseAggregatedTagModel`

NewCursorPaginatedResponseAggregatedTagModelWithDefaults instantiates a new CursorPaginatedResponseAggregatedTagModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CursorPaginatedResponseAggregatedTagModel) GetData() []AggregatedTagModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CursorPaginatedResponseAggregatedTagModel) GetDataOk() (*[]AggregatedTagModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CursorPaginatedResponseAggregatedTagModel) SetData(v []AggregatedTagModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *CursorPaginatedResponseAggregatedTagModel) GetPagination() CursorPageInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CursorPaginatedResponseAggregatedTagModel) GetPaginationOk() (*CursorPageInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CursorPaginatedResponseAggregatedTagModel) SetPagination(v CursorPageInfo)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


