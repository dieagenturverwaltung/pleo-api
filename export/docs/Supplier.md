# Supplier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **NullableString** | Supplier account code in the accounting system | [optional] 
**CategoryCode** | Pointer to **NullableString** | This is the category code that describes the merchant/supplier&#39;s activity | [optional] 
**Country** | Pointer to **NullableString** | The 2-letter country code the merchant or supplier operates in | [optional] 
**Identifier** | Pointer to **NullableString** | Identifier of the supplier in the accounting system | [optional] 
**Name** | Pointer to **NullableString** | Name of the supplier or merchant | [optional] 
**TaxIdentifier** | Pointer to **NullableString** | This is the tax identification of the supplier in their country of operation | [optional] 

## Methods

### NewSupplier

`func NewSupplier() *Supplier`

NewSupplier instantiates a new Supplier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplierWithDefaults

`func NewSupplierWithDefaults() *Supplier`

NewSupplierWithDefaults instantiates a new Supplier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Supplier) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Supplier) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Supplier) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Supplier) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *Supplier) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *Supplier) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetCategoryCode

`func (o *Supplier) GetCategoryCode() string`

GetCategoryCode returns the CategoryCode field if non-nil, zero value otherwise.

### GetCategoryCodeOk

`func (o *Supplier) GetCategoryCodeOk() (*string, bool)`

GetCategoryCodeOk returns a tuple with the CategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryCode

`func (o *Supplier) SetCategoryCode(v string)`

SetCategoryCode sets CategoryCode field to given value.

### HasCategoryCode

`func (o *Supplier) HasCategoryCode() bool`

HasCategoryCode returns a boolean if a field has been set.

### SetCategoryCodeNil

`func (o *Supplier) SetCategoryCodeNil(b bool)`

 SetCategoryCodeNil sets the value for CategoryCode to be an explicit nil

### UnsetCategoryCode
`func (o *Supplier) UnsetCategoryCode()`

UnsetCategoryCode ensures that no value is present for CategoryCode, not even an explicit nil
### GetCountry

`func (o *Supplier) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Supplier) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Supplier) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Supplier) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *Supplier) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Supplier) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetIdentifier

`func (o *Supplier) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Supplier) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Supplier) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Supplier) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *Supplier) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *Supplier) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetName

`func (o *Supplier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Supplier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Supplier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Supplier) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Supplier) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Supplier) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTaxIdentifier

`func (o *Supplier) GetTaxIdentifier() string`

GetTaxIdentifier returns the TaxIdentifier field if non-nil, zero value otherwise.

### GetTaxIdentifierOk

`func (o *Supplier) GetTaxIdentifierOk() (*string, bool)`

GetTaxIdentifierOk returns a tuple with the TaxIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentifier

`func (o *Supplier) SetTaxIdentifier(v string)`

SetTaxIdentifier sets TaxIdentifier field to given value.

### HasTaxIdentifier

`func (o *Supplier) HasTaxIdentifier() bool`

HasTaxIdentifier returns a boolean if a field has been set.

### SetTaxIdentifierNil

`func (o *Supplier) SetTaxIdentifierNil(b bool)`

 SetTaxIdentifierNil sets the value for TaxIdentifier to be an explicit nil

### UnsetTaxIdentifier
`func (o *Supplier) UnsetTaxIdentifier()`

UnsetTaxIdentifier ensures that no value is present for TaxIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


