# ExpenseResponseLines

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountSettled** | [***ExpenseResponseAmountSettled1**](ExpenseResponse_amountSettled_1.md) |  | [default to null]
**AccountId** | **string** | This is the UUID pleo accounting category identifier for this expense line, e.g. entertainment, travel etc. | [optional] [default to null]
**TaxCodeId** | **string** | This is the UUID tax code identifier for this expense line | [optional] [default to null]
**TagIds** | [**[]ExpenseResponseTagIds**](ExpenseResponse_tagIds.md) | These are additional accounting-related information pertaining to the expense, these are usually extracted from the external accounting system. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

