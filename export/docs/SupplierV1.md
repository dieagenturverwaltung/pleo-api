# SupplierV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **NullableString** | Supplier account code in the accounting system | [optional] 
**CategoryCode** | Pointer to **NullableString** | This is the category code that describes the merchant/supplier&#39;s activity | [optional] 
**Code** | Pointer to **NullableString** | Identifier of the supplier in the accounting system | [optional] 
**Country** | Pointer to **NullableString** | The 2-letter country code the merchant or supplier operates in | [optional] 
**Name** | Pointer to **NullableString** | Name of the supplier or merchant | [optional] 
**TaxIdentifier** | Pointer to **NullableString** | This is the tax identification of the supplier in their country of operation | [optional] 

## Methods

### NewSupplierV1

`func NewSupplierV1() *SupplierV1`

NewSupplierV1 instantiates a new SupplierV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplierV1WithDefaults

`func NewSupplierV1WithDefaults() *SupplierV1`

NewSupplierV1WithDefaults instantiates a new SupplierV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *SupplierV1) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SupplierV1) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SupplierV1) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SupplierV1) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *SupplierV1) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *SupplierV1) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetCategoryCode

`func (o *SupplierV1) GetCategoryCode() string`

GetCategoryCode returns the CategoryCode field if non-nil, zero value otherwise.

### GetCategoryCodeOk

`func (o *SupplierV1) GetCategoryCodeOk() (*string, bool)`

GetCategoryCodeOk returns a tuple with the CategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryCode

`func (o *SupplierV1) SetCategoryCode(v string)`

SetCategoryCode sets CategoryCode field to given value.

### HasCategoryCode

`func (o *SupplierV1) HasCategoryCode() bool`

HasCategoryCode returns a boolean if a field has been set.

### SetCategoryCodeNil

`func (o *SupplierV1) SetCategoryCodeNil(b bool)`

 SetCategoryCodeNil sets the value for CategoryCode to be an explicit nil

### UnsetCategoryCode
`func (o *SupplierV1) UnsetCategoryCode()`

UnsetCategoryCode ensures that no value is present for CategoryCode, not even an explicit nil
### GetCode

`func (o *SupplierV1) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SupplierV1) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SupplierV1) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SupplierV1) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *SupplierV1) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *SupplierV1) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCountry

`func (o *SupplierV1) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SupplierV1) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SupplierV1) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SupplierV1) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *SupplierV1) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *SupplierV1) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetName

`func (o *SupplierV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SupplierV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SupplierV1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SupplierV1) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SupplierV1) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SupplierV1) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTaxIdentifier

`func (o *SupplierV1) GetTaxIdentifier() string`

GetTaxIdentifier returns the TaxIdentifier field if non-nil, zero value otherwise.

### GetTaxIdentifierOk

`func (o *SupplierV1) GetTaxIdentifierOk() (*string, bool)`

GetTaxIdentifierOk returns a tuple with the TaxIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentifier

`func (o *SupplierV1) SetTaxIdentifier(v string)`

SetTaxIdentifier sets TaxIdentifier field to given value.

### HasTaxIdentifier

`func (o *SupplierV1) HasTaxIdentifier() bool`

HasTaxIdentifier returns a boolean if a field has been set.

### SetTaxIdentifierNil

`func (o *SupplierV1) SetTaxIdentifierNil(b bool)`

 SetTaxIdentifierNil sets the value for TaxIdentifier to be an explicit nil

### UnsetTaxIdentifier
`func (o *SupplierV1) UnsetTaxIdentifier()`

UnsetTaxIdentifier ensures that no value is present for TaxIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


