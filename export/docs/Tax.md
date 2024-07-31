# Tax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**ExportItemAmount**](ExportItemAmount.md) |  | 
**Code** | Pointer to **NullableString** | Tax code | [optional] 
**Rate** | **float32** | Tax rate. This is represented in decimals and not the percentage. e.g. 20% tax rate would be 0.20 | 
**Type** | **string** |  | 

## Methods

### NewTax

`func NewTax(amount ExportItemAmount, rate float32, type_ string, ) *Tax`

NewTax instantiates a new Tax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxWithDefaults

`func NewTaxWithDefaults() *Tax`

NewTaxWithDefaults instantiates a new Tax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Tax) GetAmount() ExportItemAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Tax) GetAmountOk() (*ExportItemAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Tax) SetAmount(v ExportItemAmount)`

SetAmount sets Amount field to given value.


### GetCode

`func (o *Tax) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Tax) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Tax) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Tax) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Tax) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Tax) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetRate

`func (o *Tax) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *Tax) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *Tax) SetRate(v float32)`

SetRate sets Rate field to given value.


### GetType

`func (o *Tax) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Tax) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Tax) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


