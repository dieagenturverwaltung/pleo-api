# ExportItemMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attendees** | Pointer to **[]string** | Used to represent attendance involvement or participants at the time of the spend. E.g. if an employee took clients to dinner, this would be used to capture who was present at the dinner. | [optional] 
**InvoiceInformation** | Pointer to [**NullableInvoiceInformation**](InvoiceInformation.md) |  | [optional] 
**ReconciliationId** | **string** | This is an identifier used to reconcile between Pleo and the accounting system. Also known in Pleo as the book ID. | 

## Methods

### NewExportItemMetadata

`func NewExportItemMetadata(reconciliationId string, ) *ExportItemMetadata`

NewExportItemMetadata instantiates a new ExportItemMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportItemMetadataWithDefaults

`func NewExportItemMetadataWithDefaults() *ExportItemMetadata`

NewExportItemMetadataWithDefaults instantiates a new ExportItemMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttendees

`func (o *ExportItemMetadata) GetAttendees() []*string`

GetAttendees returns the Attendees field if non-nil, zero value otherwise.

### GetAttendeesOk

`func (o *ExportItemMetadata) GetAttendeesOk() (*[]*string, bool)`

GetAttendeesOk returns a tuple with the Attendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendees

`func (o *ExportItemMetadata) SetAttendees(v []*string)`

SetAttendees sets Attendees field to given value.

### HasAttendees

`func (o *ExportItemMetadata) HasAttendees() bool`

HasAttendees returns a boolean if a field has been set.

### SetAttendeesNil

`func (o *ExportItemMetadata) SetAttendeesNil(b bool)`

 SetAttendeesNil sets the value for Attendees to be an explicit nil

### UnsetAttendees
`func (o *ExportItemMetadata) UnsetAttendees()`

UnsetAttendees ensures that no value is present for Attendees, not even an explicit nil
### GetInvoiceInformation

`func (o *ExportItemMetadata) GetInvoiceInformation() InvoiceInformation`

GetInvoiceInformation returns the InvoiceInformation field if non-nil, zero value otherwise.

### GetInvoiceInformationOk

`func (o *ExportItemMetadata) GetInvoiceInformationOk() (*InvoiceInformation, bool)`

GetInvoiceInformationOk returns a tuple with the InvoiceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceInformation

`func (o *ExportItemMetadata) SetInvoiceInformation(v InvoiceInformation)`

SetInvoiceInformation sets InvoiceInformation field to given value.

### HasInvoiceInformation

`func (o *ExportItemMetadata) HasInvoiceInformation() bool`

HasInvoiceInformation returns a boolean if a field has been set.

### SetInvoiceInformationNil

`func (o *ExportItemMetadata) SetInvoiceInformationNil(b bool)`

 SetInvoiceInformationNil sets the value for InvoiceInformation to be an explicit nil

### UnsetInvoiceInformation
`func (o *ExportItemMetadata) UnsetInvoiceInformation()`

UnsetInvoiceInformation ensures that no value is present for InvoiceInformation, not even an explicit nil
### GetReconciliationId

`func (o *ExportItemMetadata) GetReconciliationId() string`

GetReconciliationId returns the ReconciliationId field if non-nil, zero value otherwise.

### GetReconciliationIdOk

`func (o *ExportItemMetadata) GetReconciliationIdOk() (*string, bool)`

GetReconciliationIdOk returns a tuple with the ReconciliationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationId

`func (o *ExportItemMetadata) SetReconciliationId(v string)`

SetReconciliationId sets ReconciliationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


