# TaxCodeCreateUpdateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | Boolean flag used to archive or unarchive a category (when set to true, account is not visible or usable on the platform) | 
**Code** | Pointer to **NullableString** | The accounting system&#39;s internal identifier of the tax code | [optional] 
**CompanyId** | **string** | Company ID the tax code is being created/updated for | 
**IngoingTaxAccount** | Pointer to **NullableString** | Ingoing tax account usually used to account for reverse VAT | [optional] 
**Name** | **string** | Name of the tax code | 
**OutgoingTaxAccount** | Pointer to **NullableString** | Outgoing tax account usually used to account for reverse VAT | [optional] 
**Rate** | **float32** | Percentage rate applied for this tax code (This is represented in decimals and not the percentage. e.g. 20% tax rate would be 0.20) | 
**Type** | **string** | Classification of this tax code | 

## Methods

### NewTaxCodeCreateUpdateModel

`func NewTaxCodeCreateUpdateModel(archived bool, companyId string, name string, rate float32, type_ string, ) *TaxCodeCreateUpdateModel`

NewTaxCodeCreateUpdateModel instantiates a new TaxCodeCreateUpdateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxCodeCreateUpdateModelWithDefaults

`func NewTaxCodeCreateUpdateModelWithDefaults() *TaxCodeCreateUpdateModel`

NewTaxCodeCreateUpdateModelWithDefaults instantiates a new TaxCodeCreateUpdateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *TaxCodeCreateUpdateModel) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *TaxCodeCreateUpdateModel) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *TaxCodeCreateUpdateModel) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetCode

`func (o *TaxCodeCreateUpdateModel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TaxCodeCreateUpdateModel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TaxCodeCreateUpdateModel) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TaxCodeCreateUpdateModel) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *TaxCodeCreateUpdateModel) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *TaxCodeCreateUpdateModel) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCompanyId

`func (o *TaxCodeCreateUpdateModel) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *TaxCodeCreateUpdateModel) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *TaxCodeCreateUpdateModel) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetIngoingTaxAccount

`func (o *TaxCodeCreateUpdateModel) GetIngoingTaxAccount() string`

GetIngoingTaxAccount returns the IngoingTaxAccount field if non-nil, zero value otherwise.

### GetIngoingTaxAccountOk

`func (o *TaxCodeCreateUpdateModel) GetIngoingTaxAccountOk() (*string, bool)`

GetIngoingTaxAccountOk returns a tuple with the IngoingTaxAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngoingTaxAccount

`func (o *TaxCodeCreateUpdateModel) SetIngoingTaxAccount(v string)`

SetIngoingTaxAccount sets IngoingTaxAccount field to given value.

### HasIngoingTaxAccount

`func (o *TaxCodeCreateUpdateModel) HasIngoingTaxAccount() bool`

HasIngoingTaxAccount returns a boolean if a field has been set.

### SetIngoingTaxAccountNil

`func (o *TaxCodeCreateUpdateModel) SetIngoingTaxAccountNil(b bool)`

 SetIngoingTaxAccountNil sets the value for IngoingTaxAccount to be an explicit nil

### UnsetIngoingTaxAccount
`func (o *TaxCodeCreateUpdateModel) UnsetIngoingTaxAccount()`

UnsetIngoingTaxAccount ensures that no value is present for IngoingTaxAccount, not even an explicit nil
### GetName

`func (o *TaxCodeCreateUpdateModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxCodeCreateUpdateModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxCodeCreateUpdateModel) SetName(v string)`

SetName sets Name field to given value.


### GetOutgoingTaxAccount

`func (o *TaxCodeCreateUpdateModel) GetOutgoingTaxAccount() string`

GetOutgoingTaxAccount returns the OutgoingTaxAccount field if non-nil, zero value otherwise.

### GetOutgoingTaxAccountOk

`func (o *TaxCodeCreateUpdateModel) GetOutgoingTaxAccountOk() (*string, bool)`

GetOutgoingTaxAccountOk returns a tuple with the OutgoingTaxAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingTaxAccount

`func (o *TaxCodeCreateUpdateModel) SetOutgoingTaxAccount(v string)`

SetOutgoingTaxAccount sets OutgoingTaxAccount field to given value.

### HasOutgoingTaxAccount

`func (o *TaxCodeCreateUpdateModel) HasOutgoingTaxAccount() bool`

HasOutgoingTaxAccount returns a boolean if a field has been set.

### SetOutgoingTaxAccountNil

`func (o *TaxCodeCreateUpdateModel) SetOutgoingTaxAccountNil(b bool)`

 SetOutgoingTaxAccountNil sets the value for OutgoingTaxAccount to be an explicit nil

### UnsetOutgoingTaxAccount
`func (o *TaxCodeCreateUpdateModel) UnsetOutgoingTaxAccount()`

UnsetOutgoingTaxAccount ensures that no value is present for OutgoingTaxAccount, not even an explicit nil
### GetRate

`func (o *TaxCodeCreateUpdateModel) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *TaxCodeCreateUpdateModel) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *TaxCodeCreateUpdateModel) SetRate(v float32)`

SetRate sets Rate field to given value.


### GetType

`func (o *TaxCodeCreateUpdateModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxCodeCreateUpdateModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxCodeCreateUpdateModel) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


