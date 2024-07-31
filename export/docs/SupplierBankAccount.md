# SupplierBankAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **NullableString** | Bank account number | [optional] 
**BankCode** | Pointer to **NullableString** | Bank code | [optional] 
**BankName** | **string** | Name of the bank | 
**Bic** | Pointer to **NullableString** | Branch information | [optional] 
**Country** | **string** | Country the bank is located in. This would be the 2 letter country code | 
**Iban** | Pointer to **NullableString** | Unique international bank account number | [optional] 

## Methods

### NewSupplierBankAccount

`func NewSupplierBankAccount(bankName string, country string, ) *SupplierBankAccount`

NewSupplierBankAccount instantiates a new SupplierBankAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplierBankAccountWithDefaults

`func NewSupplierBankAccountWithDefaults() *SupplierBankAccount`

NewSupplierBankAccountWithDefaults instantiates a new SupplierBankAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *SupplierBankAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *SupplierBankAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *SupplierBankAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *SupplierBankAccount) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### SetAccountNumberNil

`func (o *SupplierBankAccount) SetAccountNumberNil(b bool)`

 SetAccountNumberNil sets the value for AccountNumber to be an explicit nil

### UnsetAccountNumber
`func (o *SupplierBankAccount) UnsetAccountNumber()`

UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
### GetBankCode

`func (o *SupplierBankAccount) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *SupplierBankAccount) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *SupplierBankAccount) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.

### HasBankCode

`func (o *SupplierBankAccount) HasBankCode() bool`

HasBankCode returns a boolean if a field has been set.

### SetBankCodeNil

`func (o *SupplierBankAccount) SetBankCodeNil(b bool)`

 SetBankCodeNil sets the value for BankCode to be an explicit nil

### UnsetBankCode
`func (o *SupplierBankAccount) UnsetBankCode()`

UnsetBankCode ensures that no value is present for BankCode, not even an explicit nil
### GetBankName

`func (o *SupplierBankAccount) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *SupplierBankAccount) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *SupplierBankAccount) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetBic

`func (o *SupplierBankAccount) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *SupplierBankAccount) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *SupplierBankAccount) SetBic(v string)`

SetBic sets Bic field to given value.

### HasBic

`func (o *SupplierBankAccount) HasBic() bool`

HasBic returns a boolean if a field has been set.

### SetBicNil

`func (o *SupplierBankAccount) SetBicNil(b bool)`

 SetBicNil sets the value for Bic to be an explicit nil

### UnsetBic
`func (o *SupplierBankAccount) UnsetBic()`

UnsetBic ensures that no value is present for Bic, not even an explicit nil
### GetCountry

`func (o *SupplierBankAccount) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SupplierBankAccount) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SupplierBankAccount) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetIban

`func (o *SupplierBankAccount) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *SupplierBankAccount) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *SupplierBankAccount) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *SupplierBankAccount) HasIban() bool`

HasIban returns a boolean if a field has been set.

### SetIbanNil

`func (o *SupplierBankAccount) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *SupplierBankAccount) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


