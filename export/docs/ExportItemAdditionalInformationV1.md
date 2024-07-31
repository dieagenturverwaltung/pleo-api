# ExportItemAdditionalInformationV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attendees** | Pointer to **[]string** | Used to represent attendance involvement or participants at the time of the spend. E.g. if an employee took clients to dinner, this would be used to capture who was present at the dinner. | [optional] 
**InvoiceInformation** | Pointer to [**NullableInvoiceInformationV1**](InvoiceInformationV1.md) |  | [optional] 
**ReconciledEntries** | Pointer to **[]string** | A list of reconciliation IDs that were settled by this accounting entry. E.g. if this is a payment, then this would be the list of reconciliation IDs that for invoices were settled by this payment. If this is a reimbursement, it would be a list of the reconciliation IDs of the pocket expenses that the reimbursement settles. | [optional] 
**ReconciliationId** | **string** | This is an identifier used to reconcile between Pleo and the accounting system. Also known in Pleo as the book ID. | 

## Methods

### NewExportItemAdditionalInformationV1

`func NewExportItemAdditionalInformationV1(reconciliationId string, ) *ExportItemAdditionalInformationV1`

NewExportItemAdditionalInformationV1 instantiates a new ExportItemAdditionalInformationV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportItemAdditionalInformationV1WithDefaults

`func NewExportItemAdditionalInformationV1WithDefaults() *ExportItemAdditionalInformationV1`

NewExportItemAdditionalInformationV1WithDefaults instantiates a new ExportItemAdditionalInformationV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttendees

`func (o *ExportItemAdditionalInformationV1) GetAttendees() []*string`

GetAttendees returns the Attendees field if non-nil, zero value otherwise.

### GetAttendeesOk

`func (o *ExportItemAdditionalInformationV1) GetAttendeesOk() (*[]*string, bool)`

GetAttendeesOk returns a tuple with the Attendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendees

`func (o *ExportItemAdditionalInformationV1) SetAttendees(v []*string)`

SetAttendees sets Attendees field to given value.

### HasAttendees

`func (o *ExportItemAdditionalInformationV1) HasAttendees() bool`

HasAttendees returns a boolean if a field has been set.

### SetAttendeesNil

`func (o *ExportItemAdditionalInformationV1) SetAttendeesNil(b bool)`

 SetAttendeesNil sets the value for Attendees to be an explicit nil

### UnsetAttendees
`func (o *ExportItemAdditionalInformationV1) UnsetAttendees()`

UnsetAttendees ensures that no value is present for Attendees, not even an explicit nil
### GetInvoiceInformation

`func (o *ExportItemAdditionalInformationV1) GetInvoiceInformation() InvoiceInformationV1`

GetInvoiceInformation returns the InvoiceInformation field if non-nil, zero value otherwise.

### GetInvoiceInformationOk

`func (o *ExportItemAdditionalInformationV1) GetInvoiceInformationOk() (*InvoiceInformationV1, bool)`

GetInvoiceInformationOk returns a tuple with the InvoiceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceInformation

`func (o *ExportItemAdditionalInformationV1) SetInvoiceInformation(v InvoiceInformationV1)`

SetInvoiceInformation sets InvoiceInformation field to given value.

### HasInvoiceInformation

`func (o *ExportItemAdditionalInformationV1) HasInvoiceInformation() bool`

HasInvoiceInformation returns a boolean if a field has been set.

### SetInvoiceInformationNil

`func (o *ExportItemAdditionalInformationV1) SetInvoiceInformationNil(b bool)`

 SetInvoiceInformationNil sets the value for InvoiceInformation to be an explicit nil

### UnsetInvoiceInformation
`func (o *ExportItemAdditionalInformationV1) UnsetInvoiceInformation()`

UnsetInvoiceInformation ensures that no value is present for InvoiceInformation, not even an explicit nil
### GetReconciledEntries

`func (o *ExportItemAdditionalInformationV1) GetReconciledEntries() []*string`

GetReconciledEntries returns the ReconciledEntries field if non-nil, zero value otherwise.

### GetReconciledEntriesOk

`func (o *ExportItemAdditionalInformationV1) GetReconciledEntriesOk() (*[]*string, bool)`

GetReconciledEntriesOk returns a tuple with the ReconciledEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciledEntries

`func (o *ExportItemAdditionalInformationV1) SetReconciledEntries(v []*string)`

SetReconciledEntries sets ReconciledEntries field to given value.

### HasReconciledEntries

`func (o *ExportItemAdditionalInformationV1) HasReconciledEntries() bool`

HasReconciledEntries returns a boolean if a field has been set.

### SetReconciledEntriesNil

`func (o *ExportItemAdditionalInformationV1) SetReconciledEntriesNil(b bool)`

 SetReconciledEntriesNil sets the value for ReconciledEntries to be an explicit nil

### UnsetReconciledEntries
`func (o *ExportItemAdditionalInformationV1) UnsetReconciledEntries()`

UnsetReconciledEntries ensures that no value is present for ReconciledEntries, not even an explicit nil
### GetReconciliationId

`func (o *ExportItemAdditionalInformationV1) GetReconciliationId() string`

GetReconciliationId returns the ReconciliationId field if non-nil, zero value otherwise.

### GetReconciliationIdOk

`func (o *ExportItemAdditionalInformationV1) GetReconciliationIdOk() (*string, bool)`

GetReconciliationIdOk returns a tuple with the ReconciliationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationId

`func (o *ExportItemAdditionalInformationV1) SetReconciliationId(v string)`

SetReconciliationId sets ReconciliationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


