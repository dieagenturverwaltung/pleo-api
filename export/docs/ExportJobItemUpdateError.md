# ExportJobItemUpdateError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountingEntryId** | **string** | Accounting entry ID for which this error occurred | 
**Message** | **string** | Descriptive error message | 
**Type** | **string** | Descriptive error type | 

## Methods

### NewExportJobItemUpdateError

`func NewExportJobItemUpdateError(accountingEntryId string, message string, type_ string, ) *ExportJobItemUpdateError`

NewExportJobItemUpdateError instantiates a new ExportJobItemUpdateError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportJobItemUpdateErrorWithDefaults

`func NewExportJobItemUpdateErrorWithDefaults() *ExportJobItemUpdateError`

NewExportJobItemUpdateErrorWithDefaults instantiates a new ExportJobItemUpdateError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountingEntryId

`func (o *ExportJobItemUpdateError) GetAccountingEntryId() string`

GetAccountingEntryId returns the AccountingEntryId field if non-nil, zero value otherwise.

### GetAccountingEntryIdOk

`func (o *ExportJobItemUpdateError) GetAccountingEntryIdOk() (*string, bool)`

GetAccountingEntryIdOk returns a tuple with the AccountingEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingEntryId

`func (o *ExportJobItemUpdateError) SetAccountingEntryId(v string)`

SetAccountingEntryId sets AccountingEntryId field to given value.


### GetMessage

`func (o *ExportJobItemUpdateError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ExportJobItemUpdateError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ExportJobItemUpdateError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetType

`func (o *ExportJobItemUpdateError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportJobItemUpdateError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportJobItemUpdateError) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


