# ExportJobItemUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ExportJobItem**](ExportJobItem.md) |  | 
**Errors** | [**[]ExportJobItemUpdateError**](ExportJobItemUpdateError.md) | If any errors occurred in updating the export job items, they will be listed here | 

## Methods

### NewExportJobItemUpdate

`func NewExportJobItemUpdate(data []ExportJobItem, errors []ExportJobItemUpdateError, ) *ExportJobItemUpdate`

NewExportJobItemUpdate instantiates a new ExportJobItemUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportJobItemUpdateWithDefaults

`func NewExportJobItemUpdateWithDefaults() *ExportJobItemUpdate`

NewExportJobItemUpdateWithDefaults instantiates a new ExportJobItemUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ExportJobItemUpdate) GetData() []ExportJobItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExportJobItemUpdate) GetDataOk() (*[]ExportJobItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExportJobItemUpdate) SetData(v []ExportJobItem)`

SetData sets Data field to given value.


### GetErrors

`func (o *ExportJobItemUpdate) GetErrors() []ExportJobItemUpdateError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ExportJobItemUpdate) GetErrorsOk() (*[]ExportJobItemUpdateError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ExportJobItemUpdate) SetErrors(v []ExportJobItemUpdateError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


