# TaxCodeSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **bool** | If set to true will only return tax codes that have been archived, if set to false will only return tax codes that have not been archived. By default,both archived and non-archived tax codes will be returned. | [optional] 
**TaxCodeIds** | Pointer to **[]string** | Useful for fetching a list of tax codes given a specific list of IDs. | [optional] 
**Type** | Pointer to **string** | Classification of the tax codes to fetch | [optional] 

## Methods

### NewTaxCodeSearchRequest

`func NewTaxCodeSearchRequest() *TaxCodeSearchRequest`

NewTaxCodeSearchRequest instantiates a new TaxCodeSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxCodeSearchRequestWithDefaults

`func NewTaxCodeSearchRequestWithDefaults() *TaxCodeSearchRequest`

NewTaxCodeSearchRequestWithDefaults instantiates a new TaxCodeSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *TaxCodeSearchRequest) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *TaxCodeSearchRequest) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *TaxCodeSearchRequest) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *TaxCodeSearchRequest) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetTaxCodeIds

`func (o *TaxCodeSearchRequest) GetTaxCodeIds() []string`

GetTaxCodeIds returns the TaxCodeIds field if non-nil, zero value otherwise.

### GetTaxCodeIdsOk

`func (o *TaxCodeSearchRequest) GetTaxCodeIdsOk() (*[]string, bool)`

GetTaxCodeIdsOk returns a tuple with the TaxCodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCodeIds

`func (o *TaxCodeSearchRequest) SetTaxCodeIds(v []string)`

SetTaxCodeIds sets TaxCodeIds field to given value.

### HasTaxCodeIds

`func (o *TaxCodeSearchRequest) HasTaxCodeIds() bool`

HasTaxCodeIds returns a boolean if a field has been set.

### GetType

`func (o *TaxCodeSearchRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxCodeSearchRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxCodeSearchRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaxCodeSearchRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


