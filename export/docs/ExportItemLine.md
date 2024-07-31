# ExportItemLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**NullableAccount**](Account.md) |  | [optional] 
**AccountingEntryLineId** | **string** | This is the Pleo internal identifier of this accounting entry line | 
**LineAmount** | [**ExportItemAmount**](ExportItemAmount.md) |  | 
**NetAmount** | [**ExportItemAmount**](ExportItemAmount.md) |  | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tags are used to appropriate funds to specific projects, cost centers, departments, etc. This is what is known as \&quot;tags\&quot; in Pleo. | [optional] 
**Tax** | Pointer to [**NullableTax**](Tax.md) |  | [optional] 

## Methods

### NewExportItemLine

`func NewExportItemLine(accountingEntryLineId string, lineAmount ExportItemAmount, netAmount ExportItemAmount, ) *ExportItemLine`

NewExportItemLine instantiates a new ExportItemLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportItemLineWithDefaults

`func NewExportItemLineWithDefaults() *ExportItemLine`

NewExportItemLineWithDefaults instantiates a new ExportItemLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ExportItemLine) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ExportItemLine) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ExportItemLine) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ExportItemLine) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ExportItemLine) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ExportItemLine) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetAccountingEntryLineId

`func (o *ExportItemLine) GetAccountingEntryLineId() string`

GetAccountingEntryLineId returns the AccountingEntryLineId field if non-nil, zero value otherwise.

### GetAccountingEntryLineIdOk

`func (o *ExportItemLine) GetAccountingEntryLineIdOk() (*string, bool)`

GetAccountingEntryLineIdOk returns a tuple with the AccountingEntryLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntryLineId

`func (o *ExportItemLine) SetAccountingEntryLineId(v string)`

SetAccountingEntryLineId sets AccountingEntryLineId field to given value.


### GetLineAmount

`func (o *ExportItemLine) GetLineAmount() ExportItemAmount`

GetLineAmount returns the LineAmount field if non-nil, zero value otherwise.

### GetLineAmountOk

`func (o *ExportItemLine) GetLineAmountOk() (*ExportItemAmount, bool)`

GetLineAmountOk returns a tuple with the LineAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineAmount

`func (o *ExportItemLine) SetLineAmount(v ExportItemAmount)`

SetLineAmount sets LineAmount field to given value.


### GetNetAmount

`func (o *ExportItemLine) GetNetAmount() ExportItemAmount`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *ExportItemLine) GetNetAmountOk() (*ExportItemAmount, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *ExportItemLine) SetNetAmount(v ExportItemAmount)`

SetNetAmount sets NetAmount field to given value.


### GetTags

`func (o *ExportItemLine) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExportItemLine) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExportItemLine) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExportItemLine) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ExportItemLine) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ExportItemLine) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTax

`func (o *ExportItemLine) GetTax() Tax`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *ExportItemLine) GetTaxOk() (*Tax, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *ExportItemLine) SetTax(v Tax)`

SetTax sets Tax field to given value.

### HasTax

`func (o *ExportItemLine) HasTax() bool`

HasTax returns a boolean if a field has been set.

### SetTaxNil

`func (o *ExportItemLine) SetTaxNil(b bool)`

 SetTaxNil sets the value for Tax to be an explicit nil

### UnsetTax
`func (o *ExportItemLine) UnsetTax()`

UnsetTax ensures that no value is present for Tax, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


