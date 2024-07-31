# CursorPageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentRequestPagination** | [**CursorPageCurrentRequestInfo**](CursorPageCurrentRequestInfo.md) |  | 
**EndCursor** | Pointer to **string** |  | [optional] 
**HasNextPage** | **bool** |  | 
**HasPreviousPage** | **bool** |  | 
**StartCursor** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewCursorPageInfo

`func NewCursorPageInfo(currentRequestPagination CursorPageCurrentRequestInfo, hasNextPage bool, hasPreviousPage bool, ) *CursorPageInfo`

NewCursorPageInfo instantiates a new CursorPageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPageInfoWithDefaults

`func NewCursorPageInfoWithDefaults() *CursorPageInfo`

NewCursorPageInfoWithDefaults instantiates a new CursorPageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentRequestPagination

`func (o *CursorPageInfo) GetCurrentRequestPagination() CursorPageCurrentRequestInfo`

GetCurrentRequestPagination returns the CurrentRequestPagination field if non-nil, zero value otherwise.

### GetCurrentRequestPaginationOk

`func (o *CursorPageInfo) GetCurrentRequestPaginationOk() (*CursorPageCurrentRequestInfo, bool)`

GetCurrentRequestPaginationOk returns a tuple with the CurrentRequestPagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRequestPagination

`func (o *CursorPageInfo) SetCurrentRequestPagination(v CursorPageCurrentRequestInfo)`

SetCurrentRequestPagination sets CurrentRequestPagination field to given value.


### GetEndCursor

`func (o *CursorPageInfo) GetEndCursor() string`

GetEndCursor returns the EndCursor field if non-nil, zero value otherwise.

### GetEndCursorOk

`func (o *CursorPageInfo) GetEndCursorOk() (*string, bool)`

GetEndCursorOk returns a tuple with the EndCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCursor

`func (o *CursorPageInfo) SetEndCursor(v string)`

SetEndCursor sets EndCursor field to given value.

### HasEndCursor

`func (o *CursorPageInfo) HasEndCursor() bool`

HasEndCursor returns a boolean if a field has been set.

### GetHasNextPage

`func (o *CursorPageInfo) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *CursorPageInfo) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *CursorPageInfo) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetHasPreviousPage

`func (o *CursorPageInfo) GetHasPreviousPage() bool`

GetHasPreviousPage returns the HasPreviousPage field if non-nil, zero value otherwise.

### GetHasPreviousPageOk

`func (o *CursorPageInfo) GetHasPreviousPageOk() (*bool, bool)`

GetHasPreviousPageOk returns a tuple with the HasPreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreviousPage

`func (o *CursorPageInfo) SetHasPreviousPage(v bool)`

SetHasPreviousPage sets HasPreviousPage field to given value.


### GetStartCursor

`func (o *CursorPageInfo) GetStartCursor() string`

GetStartCursor returns the StartCursor field if non-nil, zero value otherwise.

### GetStartCursorOk

`func (o *CursorPageInfo) GetStartCursorOk() (*string, bool)`

GetStartCursorOk returns a tuple with the StartCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCursor

`func (o *CursorPageInfo) SetStartCursor(v string)`

SetStartCursor sets StartCursor field to given value.

### HasStartCursor

`func (o *CursorPageInfo) HasStartCursor() bool`

HasStartCursor returns a boolean if a field has been set.

### GetTotal

`func (o *CursorPageInfo) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CursorPageInfo) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CursorPageInfo) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CursorPageInfo) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


