# InvoiceInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DueDate** | Pointer to **NullableTime** | Date the payment is due | [optional] 
**InvoiceDate** | **time.Time** | Date the invoice was issued | 
**InvoiceNumber** | **string** | The invoice number | 
**PaymentDate** | Pointer to **NullableTime** | Date the invoice was paid if the status is paid | [optional] 
**Status** | **string** | The invoice status | 
**SupplierBankAccount** | Pointer to [**NullableSupplierBankAccount**](SupplierBankAccount.md) |  | [optional] 

## Methods

### NewInvoiceInformation

`func NewInvoiceInformation(invoiceDate time.Time, invoiceNumber string, status string, ) *InvoiceInformation`

NewInvoiceInformation instantiates a new InvoiceInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceInformationWithDefaults

`func NewInvoiceInformationWithDefaults() *InvoiceInformation`

NewInvoiceInformationWithDefaults instantiates a new InvoiceInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDueDate

`func (o *InvoiceInformation) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceInformation) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceInformation) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvoiceInformation) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *InvoiceInformation) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *InvoiceInformation) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetInvoiceDate

`func (o *InvoiceInformation) GetInvoiceDate() time.Time`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *InvoiceInformation) GetInvoiceDateOk() (*time.Time, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *InvoiceInformation) SetInvoiceDate(v time.Time)`

SetInvoiceDate sets InvoiceDate field to given value.


### GetInvoiceNumber

`func (o *InvoiceInformation) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *InvoiceInformation) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *InvoiceInformation) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetPaymentDate

`func (o *InvoiceInformation) GetPaymentDate() time.Time`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *InvoiceInformation) GetPaymentDateOk() (*time.Time, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *InvoiceInformation) SetPaymentDate(v time.Time)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *InvoiceInformation) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### SetPaymentDateNil

`func (o *InvoiceInformation) SetPaymentDateNil(b bool)`

 SetPaymentDateNil sets the value for PaymentDate to be an explicit nil

### UnsetPaymentDate
`func (o *InvoiceInformation) UnsetPaymentDate()`

UnsetPaymentDate ensures that no value is present for PaymentDate, not even an explicit nil
### GetStatus

`func (o *InvoiceInformation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceInformation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceInformation) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSupplierBankAccount

`func (o *InvoiceInformation) GetSupplierBankAccount() SupplierBankAccount`

GetSupplierBankAccount returns the SupplierBankAccount field if non-nil, zero value otherwise.

### GetSupplierBankAccountOk

`func (o *InvoiceInformation) GetSupplierBankAccountOk() (*SupplierBankAccount, bool)`

GetSupplierBankAccountOk returns a tuple with the SupplierBankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierBankAccount

`func (o *InvoiceInformation) SetSupplierBankAccount(v SupplierBankAccount)`

SetSupplierBankAccount sets SupplierBankAccount field to given value.

### HasSupplierBankAccount

`func (o *InvoiceInformation) HasSupplierBankAccount() bool`

HasSupplierBankAccount returns a boolean if a field has been set.

### SetSupplierBankAccountNil

`func (o *InvoiceInformation) SetSupplierBankAccountNil(b bool)`

 SetSupplierBankAccountNil sets the value for SupplierBankAccount to be an explicit nil

### UnsetSupplierBankAccount
`func (o *InvoiceInformation) UnsetSupplierBankAccount()`

UnsetSupplierBankAccount ensures that no value is present for SupplierBankAccount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


