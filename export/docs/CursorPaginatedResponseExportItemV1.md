# CursorPaginatedResponseExportItemV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ExportItemV1**](ExportItemV1.md) |  | 
**Pagination** | [**CursorPageInfo**](CursorPageInfo.md) |  | 

## Methods

### NewCursorPaginatedResponseExportItemV1

`func NewCursorPaginatedResponseExportItemV1(data []ExportItemV1, pagination CursorPageInfo, ) *CursorPaginatedResponseExportItemV1`

NewCursorPaginatedResponseExportItemV1 instantiates a new CursorPaginatedResponseExportItemV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPaginatedResponseExportItemV1WithDefaults

`func NewCursorPaginatedResponseExportItemV1WithDefaults() *CursorPaginatedResponseExportItemV1`

NewCursorPaginatedResponseExportItemV1WithDefaults instantiates a new CursorPaginatedResponseExportItemV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CursorPaginatedResponseExportItemV1) GetData() []ExportItemV1`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CursorPaginatedResponseExportItemV1) GetDataOk() (*[]ExportItemV1, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CursorPaginatedResponseExportItemV1) SetData(v []ExportItemV1)`

SetData sets Data field to given value.


### GetPagination

`func (o *CursorPaginatedResponseExportItemV1) GetPagination() CursorPageInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CursorPaginatedResponseExportItemV1) GetPaginationOk() (*CursorPageInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CursorPaginatedResponseExportItemV1) SetPagination(v CursorPageInfo)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


