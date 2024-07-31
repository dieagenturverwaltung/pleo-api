# CursorPageCurrentRequestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **string** |  | [optional] 
**Before** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**Parameters** | **map[string][]string** |  | 
**SortingKeys** | Pointer to **[]string** |  | [optional] 
**SortingOrder** | Pointer to [**[]PageOrder**](PageOrder.md) |  | [optional] 

## Methods

### NewCursorPageCurrentRequestInfo

`func NewCursorPageCurrentRequestInfo(parameters map[string][]string, ) *CursorPageCurrentRequestInfo`

NewCursorPageCurrentRequestInfo instantiates a new CursorPageCurrentRequestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPageCurrentRequestInfoWithDefaults

`func NewCursorPageCurrentRequestInfoWithDefaults() *CursorPageCurrentRequestInfo`

NewCursorPageCurrentRequestInfoWithDefaults instantiates a new CursorPageCurrentRequestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *CursorPageCurrentRequestInfo) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *CursorPageCurrentRequestInfo) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *CursorPageCurrentRequestInfo) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *CursorPageCurrentRequestInfo) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetBefore

`func (o *CursorPageCurrentRequestInfo) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *CursorPageCurrentRequestInfo) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *CursorPageCurrentRequestInfo) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *CursorPageCurrentRequestInfo) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetLimit

`func (o *CursorPageCurrentRequestInfo) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CursorPageCurrentRequestInfo) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CursorPageCurrentRequestInfo) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CursorPageCurrentRequestInfo) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *CursorPageCurrentRequestInfo) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CursorPageCurrentRequestInfo) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CursorPageCurrentRequestInfo) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CursorPageCurrentRequestInfo) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetParameters

`func (o *CursorPageCurrentRequestInfo) GetParameters() map[string][]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CursorPageCurrentRequestInfo) GetParametersOk() (*map[string][]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CursorPageCurrentRequestInfo) SetParameters(v map[string][]string)`

SetParameters sets Parameters field to given value.


### GetSortingKeys

`func (o *CursorPageCurrentRequestInfo) GetSortingKeys() []string`

GetSortingKeys returns the SortingKeys field if non-nil, zero value otherwise.

### GetSortingKeysOk

`func (o *CursorPageCurrentRequestInfo) GetSortingKeysOk() (*[]string, bool)`

GetSortingKeysOk returns a tuple with the SortingKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingKeys

`func (o *CursorPageCurrentRequestInfo) SetSortingKeys(v []string)`

SetSortingKeys sets SortingKeys field to given value.

### HasSortingKeys

`func (o *CursorPageCurrentRequestInfo) HasSortingKeys() bool`

HasSortingKeys returns a boolean if a field has been set.

### GetSortingOrder

`func (o *CursorPageCurrentRequestInfo) GetSortingOrder() []PageOrder`

GetSortingOrder returns the SortingOrder field if non-nil, zero value otherwise.

### GetSortingOrderOk

`func (o *CursorPageCurrentRequestInfo) GetSortingOrderOk() (*[]PageOrder, bool)`

GetSortingOrderOk returns a tuple with the SortingOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingOrder

`func (o *CursorPageCurrentRequestInfo) SetSortingOrder(v []PageOrder)`

SetSortingOrder sets SortingOrder field to given value.

### HasSortingOrder

`func (o *CursorPageCurrentRequestInfo) HasSortingOrder() bool`

HasSortingOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


