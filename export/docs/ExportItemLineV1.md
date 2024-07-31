# ExportItemLineV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**NullableAccountV1**](AccountV1.md) |  | [optional] 
**AccountingEntryLineId** | **string** | This is the Pleo internal identifier of this accounting entry line | 
**LineAmount** | [**ExportItemAmount**](ExportItemAmount.md) |  | 
**NetAmount** | [**ExportItemAmount**](ExportItemAmount.md) |  | 
**Tags** | Pointer to [**[]TagV1**](TagV1.md) | Tags are used to appropriate funds to specific projects, cost centers, departments, etc. This is what is known as \&quot;tags\&quot; in Pleo. | [optional] 
**Tax** | Pointer to [**NullableTaxV1**](TaxV1.md) |  | [optional] 

## Methods

### NewExportItemLineV1

`func NewExportItemLineV1(accountingEntryLineId string, lineAmount ExportItemAmount, netAmount ExportItemAmount, ) *ExportItemLineV1`

NewExportItemLineV1 instantiates a new ExportItemLineV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportItemLineV1WithDefaults

`func NewExportItemLineV1WithDefaults() *ExportItemLineV1`

NewExportItemLineV1WithDefaults instantiates a new ExportItemLineV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ExportItemLineV1) GetAccount() AccountV1`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ExportItemLineV1) GetAccountOk() (*AccountV1, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ExportItemLineV1) SetAccount(v AccountV1)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ExportItemLineV1) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ExportItemLineV1) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ExportItemLineV1) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetAccountingEntryLineId

`func (o *ExportItemLineV1) GetAccountingEntryLineId() string`

GetAccountingEntryLineId returns the AccountingEntryLineId field if non-nil, zero value otherwise.

### GetAccountingEntryLineIdOk

`func (o *ExportItemLineV1) GetAccountingEntryLineIdOk() (*string, bool)`

GetAccountingEntryLineIdOk returns a tuple with the AccountingEntryLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntryLineId

`func (o *ExportItemLineV1) SetAccountingEntryLineId(v string)`

SetAccountingEntryLineId sets AccountingEntryLineId field to given value.


### GetLineAmount

`func (o *ExportItemLineV1) GetLineAmount() ExportItemAmount`

GetLineAmount returns the LineAmount field if non-nil, zero value otherwise.

### GetLineAmountOk

`func (o *ExportItemLineV1) GetLineAmountOk() (*ExportItemAmount, bool)`

GetLineAmountOk returns a tuple with the LineAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAmount

`func (o *ExportItemLineV1) SetLineAmount(v ExportItemAmount)`

SetLineAmount sets LineAmount field to given value.


### GetNetAmount

`func (o *ExportItemLineV1) GetNetAmount() ExportItemAmount`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *ExportItemLineV1) GetNetAmountOk() (*ExportItemAmount, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *ExportItemLineV1) SetNetAmount(v ExportItemAmount)`

SetNetAmount sets NetAmount field to given value.


### GetTags

`func (o *ExportItemLineV1) GetTags() []TagV1`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExportItemLineV1) GetTagsOk() (*[]TagV1, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExportItemLineV1) SetTags(v []TagV1)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExportItemLineV1) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ExportItemLineV1) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ExportItemLineV1) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTax

`func (o *ExportItemLineV1) GetTax() TaxV1`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *ExportItemLineV1) GetTaxOk() (*TaxV1, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *ExportItemLineV1) SetTax(v TaxV1)`

SetTax sets Tax field to given value.

### HasTax

`func (o *ExportItemLineV1) HasTax() bool`

HasTax returns a boolean if a field has been set.

### SetTaxNil

`func (o *ExportItemLineV1) SetTaxNil(b bool)`

 SetTaxNil sets the value for Tax to be an explicit nil

### UnsetTax
`func (o *ExportItemLineV1) UnsetTax()`

UnsetTax ensures that no value is present for Tax, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


