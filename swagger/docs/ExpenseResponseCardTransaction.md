# ExpenseResponseCardTransaction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | This is the transaction state of this card transaction. | [default to null]
**AuthorizedAt** | [**time.Time**](time.Time.md) | Date and time the transaction was authorized in the format YYYY-MM-DDTHH:mi:ss.SSSZ. | [optional] [default to null]
**SettledAt** | [**time.Time**](time.Time.md) | Date and time the transaction was settled in the pleo wallet currency in the format YYYY-MM-DDTHH:mi:ss.SSSZ. | [optional] [default to null]
**ReversedAt** | [**time.Time**](time.Time.md) | Date and time the transaction was reversed if the transaction is in a reversal state in the format YYYY-MM-DDTHH:mi:ss.SSSZ. | [optional] [default to null]
**Merchant** | [***ExpenseResponseCardTransactionMerchant**](ExpenseResponse_cardTransaction_merchant.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

