# TaxCodesResponseTaxCodes

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique UUID identifier of the tax code. | [default to null]
**Name** | **string** | The name of the tax. | [default to null]
**Code** | **string** | The short code of the tax. | [optional] [default to null]
**Rate** | **float64** | This is the rate of the tax code. | [default to null]
**Type_** | **string** | This is the type of the tax code. Can be \&quot;normal\&quot; or \&quot;reverse\&quot;. | [optional] [default to normal]
**IngoingTaxAccount** | **string** | This is the ingoing account related to the tax codes. | [optional] [default to null]
**OutgoingTaxAccount** | **string** | This is the outgoing account related to the tax codes. | [optional] [default to null]
**Enabled** | **bool** | This value is true if the tax code is not hidden. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Date and time this record was created in the pleo system in the format YYYY-MM-DDTHH:mi:ss.SSSZ | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Date and time this record was last updated in the pleo system in the format YYYY-MM-DDTHH:mi:ss.SSSZ | [optional] [default to null]
**DeletedAt** | [**time.Time**](time.Time.md) | Date and time this record was deleted in the pleo system in the format YYYY-MM-DDTHH:mi:ss.SSSZ | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

