# TaxV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**ExportItemAmount**](ExportItemAmount.md) |  | 
**Code** | Pointer to **NullableString** | Tax code | [optional] 
**Id** | **string** | This is the Pleo internal identifier of the tax code | 
**Rate** | **float32** | Tax rate. This is represented in decimals and not the percentage. e.g. 20% tax rate would be 0.20 | 
**Type** | **string** | Tax type | 

## Methods

### NewTaxV1

`func NewTaxV1(amount ExportItemAmount, id string, rate float32, type_ string, ) *TaxV1`

NewTaxV1 instantiates a new TaxV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxV1WithDefaults

`func NewTaxV1WithDefaults() *TaxV1`

NewTaxV1WithDefaults instantiates a new TaxV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TaxV1) GetAmount() ExportItemAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TaxV1) GetAmountOk() (*ExportItemAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TaxV1) SetAmount(v ExportItemAmount)`

SetAmount sets Amount field to given value.


### GetCode

`func (o *TaxV1) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TaxV1) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TaxV1) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TaxV1) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *TaxV1) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *TaxV1) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetId

`func (o *TaxV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxV1) SetId(v string)`

SetId sets Id field to given value.


### GetRate

`func (o *TaxV1) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *TaxV1) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *TaxV1) SetRate(v float32)`

SetRate sets Rate field to given value.


### GetType

`func (o *TaxV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxV1) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


