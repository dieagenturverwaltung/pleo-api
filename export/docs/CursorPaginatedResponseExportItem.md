# CursorPaginatedResponseExportItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ExportItem**](ExportItem.md) |  | 
**Pagination** | [**CursorPageInfo**](CursorPageInfo.md) |  | 

## Methods

### NewCursorPaginatedResponseExportItem

`func NewCursorPaginatedResponseExportItem(data []ExportItem, pagination CursorPageInfo, ) *CursorPaginatedResponseExportItem`

NewCursorPaginatedResponseExportItem instantiates a new CursorPaginatedResponseExportItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPaginatedResponseExportItemWithDefaults

`func NewCursorPaginatedResponseExportItemWithDefaults() *CursorPaginatedResponseExportItem`

NewCursorPaginatedResponseExportItemWithDefaults instantiates a new CursorPaginatedResponseExportItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CursorPaginatedResponseExportItem) GetData() []ExportItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CursorPaginatedResponseExportItem) GetDataOk() (*[]ExportItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CursorPaginatedResponseExportItem) SetData(v []ExportItem)`

SetData sets Data field to given value.


### GetPagination

`func (o *CursorPaginatedResponseExportItem) GetPagination() CursorPageInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CursorPaginatedResponseExportItem) GetPaginationOk() (*CursorPageInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CursorPaginatedResponseExportItem) SetPagination(v CursorPageInfo)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


