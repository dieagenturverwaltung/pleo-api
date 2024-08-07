/*
Export API

Export OpenAPI definitions

API version: 20.0.0
Contact: apiteam@pleo.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package export

import (
	"encoding/json"
)

// ExportItemAdditionalInformationV1 Additional information applicable to this accounting entry
type ExportItemAdditionalInformationV1 struct {
	// Used to represent attendance involvement or participants at the time of the spend. E.g. if an employee took clients to dinner, this would be used to capture who was present at the dinner.
	Attendees          []*string                    `json:"attendees,omitempty"`
	InvoiceInformation NullableInvoiceInformationV1 `json:"invoiceInformation,omitempty"`
	// A list of reconciliation IDs that were settled by this accounting entry. E.g. if this is a payment, then this would be the list of reconciliation IDs that for invoices were settled by this payment. If this is a reimbursement, it would be a list of the reconciliation IDs of the pocket expenses that the reimbursement settles.
	ReconciledEntries []*string `json:"reconciledEntries,omitempty"`
	// This is an identifier used to reconcile between Pleo and the accounting system. Also known in Pleo as the book ID.
	ReconciliationId string `json:"reconciliationId"`
}

// NewExportItemAdditionalInformationV1 instantiates a new ExportItemAdditionalInformationV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportItemAdditionalInformationV1(reconciliationId string) *ExportItemAdditionalInformationV1 {
	this := ExportItemAdditionalInformationV1{}
	this.ReconciliationId = reconciliationId
	return &this
}

// NewExportItemAdditionalInformationV1WithDefaults instantiates a new ExportItemAdditionalInformationV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportItemAdditionalInformationV1WithDefaults() *ExportItemAdditionalInformationV1 {
	this := ExportItemAdditionalInformationV1{}
	return &this
}

// GetAttendees returns the Attendees field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportItemAdditionalInformationV1) GetAttendees() []*string {
	if o == nil {
		var ret []*string
		return ret
	}
	return o.Attendees
}

// GetAttendeesOk returns a tuple with the Attendees field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportItemAdditionalInformationV1) GetAttendeesOk() ([]*string, bool) {
	if o == nil || o.Attendees == nil {
		return nil, false
	}
	return o.Attendees, true
}

// HasAttendees returns a boolean if a field has been set.
func (o *ExportItemAdditionalInformationV1) HasAttendees() bool {
	if o != nil && o.Attendees != nil {
		return true
	}

	return false
}

// SetAttendees gets a reference to the given []*string and assigns it to the Attendees field.
func (o *ExportItemAdditionalInformationV1) SetAttendees(v []*string) {
	o.Attendees = v
}

// GetInvoiceInformation returns the InvoiceInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportItemAdditionalInformationV1) GetInvoiceInformation() InvoiceInformationV1 {
	if o == nil || o.InvoiceInformation.Get() == nil {
		var ret InvoiceInformationV1
		return ret
	}
	return *o.InvoiceInformation.Get()
}

// GetInvoiceInformationOk returns a tuple with the InvoiceInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportItemAdditionalInformationV1) GetInvoiceInformationOk() (*InvoiceInformationV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceInformation.Get(), o.InvoiceInformation.IsSet()
}

// HasInvoiceInformation returns a boolean if a field has been set.
func (o *ExportItemAdditionalInformationV1) HasInvoiceInformation() bool {
	if o != nil && o.InvoiceInformation.IsSet() {
		return true
	}

	return false
}

// SetInvoiceInformation gets a reference to the given NullableInvoiceInformationV1 and assigns it to the InvoiceInformation field.
func (o *ExportItemAdditionalInformationV1) SetInvoiceInformation(v InvoiceInformationV1) {
	o.InvoiceInformation.Set(&v)
}

// SetInvoiceInformationNil sets the value for InvoiceInformation to be an explicit nil
func (o *ExportItemAdditionalInformationV1) SetInvoiceInformationNil() {
	o.InvoiceInformation.Set(nil)
}

// UnsetInvoiceInformation ensures that no value is present for InvoiceInformation, not even an explicit nil
func (o *ExportItemAdditionalInformationV1) UnsetInvoiceInformation() {
	o.InvoiceInformation.Unset()
}

// GetReconciledEntries returns the ReconciledEntries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportItemAdditionalInformationV1) GetReconciledEntries() []*string {
	if o == nil {
		var ret []*string
		return ret
	}
	return o.ReconciledEntries
}

// GetReconciledEntriesOk returns a tuple with the ReconciledEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportItemAdditionalInformationV1) GetReconciledEntriesOk() ([]*string, bool) {
	if o == nil || o.ReconciledEntries == nil {
		return nil, false
	}
	return o.ReconciledEntries, true
}

// HasReconciledEntries returns a boolean if a field has been set.
func (o *ExportItemAdditionalInformationV1) HasReconciledEntries() bool {
	if o != nil && o.ReconciledEntries != nil {
		return true
	}

	return false
}

// SetReconciledEntries gets a reference to the given []*string and assigns it to the ReconciledEntries field.
func (o *ExportItemAdditionalInformationV1) SetReconciledEntries(v []*string) {
	o.ReconciledEntries = v
}

// GetReconciliationId returns the ReconciliationId field value
func (o *ExportItemAdditionalInformationV1) GetReconciliationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReconciliationId
}

// GetReconciliationIdOk returns a tuple with the ReconciliationId field value
// and a boolean to check if the value has been set.
func (o *ExportItemAdditionalInformationV1) GetReconciliationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReconciliationId, true
}

// SetReconciliationId sets field value
func (o *ExportItemAdditionalInformationV1) SetReconciliationId(v string) {
	o.ReconciliationId = v
}

func (o ExportItemAdditionalInformationV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attendees != nil {
		toSerialize["attendees"] = o.Attendees
	}
	if o.InvoiceInformation.IsSet() {
		toSerialize["invoiceInformation"] = o.InvoiceInformation.Get()
	}
	if o.ReconciledEntries != nil {
		toSerialize["reconciledEntries"] = o.ReconciledEntries
	}
	if true {
		toSerialize["reconciliationId"] = o.ReconciliationId
	}
	return json.Marshal(toSerialize)
}

type NullableExportItemAdditionalInformationV1 struct {
	value *ExportItemAdditionalInformationV1
	isSet bool
}

func (v NullableExportItemAdditionalInformationV1) Get() *ExportItemAdditionalInformationV1 {
	return v.value
}

func (v *NullableExportItemAdditionalInformationV1) Set(val *ExportItemAdditionalInformationV1) {
	v.value = val
	v.isSet = true
}

func (v NullableExportItemAdditionalInformationV1) IsSet() bool {
	return v.isSet
}

func (v *NullableExportItemAdditionalInformationV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportItemAdditionalInformationV1(val *ExportItemAdditionalInformationV1) *NullableExportItemAdditionalInformationV1 {
	return &NullableExportItemAdditionalInformationV1{value: val, isSet: true}
}

func (v NullableExportItemAdditionalInformationV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportItemAdditionalInformationV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
