# CursorPaginatedResponseTaxCodeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TaxCodeModel**](TaxCodeModel.md) |  | 
**Pagination** | [**CursorPageInfo**](CursorPageInfo.md) |  | 

## Methods

### NewCursorPaginatedResponseTaxCodeModel

`func NewCursorPaginatedResponseTaxCodeModel(data []TaxCodeModel, pagination CursorPageInfo, ) *CursorPaginatedResponseTaxCodeModel`

NewCursorPaginatedResponseTaxCodeModel instantiates a new CursorPaginatedResponseTaxCodeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPaginatedResponseTaxCodeModelWithDefaults

`func NewCursorPaginatedResponseTaxCodeModelWithDefaults() *CursorPaginatedResponseTaxCodeModel`

NewCursorPaginatedResponseTaxCodeModelWithDefaults instantiates a new CursorPaginatedResponseTaxCodeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CursorPaginatedResponseTaxCodeModel) GetData() []TaxCodeModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CursorPaginatedResponseTaxCodeModel) GetDataOk() (*[]TaxCodeModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CursorPaginatedResponseTaxCodeModel) SetData(v []TaxCodeModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *CursorPaginatedResponseTaxCodeModel) GetPagination() CursorPageInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CursorPaginatedResponseTaxCodeModel) GetPaginationOk() (*CursorPageInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CursorPaginatedResponseTaxCodeModel) SetPagination(v CursorPageInfo)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


