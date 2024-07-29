# ExpenseResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique UUID identifier of the expense | [default to null]
**EmployeeId** | **string** | Unique UUID identifier of the employee that performed the expense | [optional] [default to null]
**EmployeeCode** | **string** | External identifier of the employee that performed the expense | [optional] [default to null]
**DepartmentId** | **string** | Unique UUID identifier of the department for which the employee executing the expense belongs to | [optional] [default to null]
**PerformedAt** | [**time.Time**](time.Time.md) | The date the expense was performed in the format YYYY-MM-DDTHH:mi:ss.SSSZ | [default to null]
**AmountOriginal** | [***ExpenseResponseAmountOriginal**](ExpenseResponse_amountOriginal.md) |  | [default to null]
**AmountSettled** | [***ExpenseResponseAmountSettled**](ExpenseResponse_amountSettled.md) |  | [optional] [default to null]
**Note** | **string** | Additional comments on the expense added to enhance the purpose of the expense | [optional] [default to null]
**Type_** | **string** | The type of this expense. | [default to null]
**CardTransaction** | [***ExpenseResponseCardTransaction**](ExpenseResponse_cardTransaction.md) |  | [optional] [default to null]
**AccountId** | **string** | This is the UUID pleo accounting category identifier for this expense, e.g. entertainment, travel etc. | [optional] [default to null]
**TaxCodeId** | **string** | This is the UUID tax code identifier for this expense | [optional] [default to null]
**TagIds** | [**[]ExpenseResponseTagIds**](ExpenseResponse_tagIds.md) | These are additional accounting-related information pertaining to the expense, these are usually extracted from the external accounting system. | [optional] [default to null]
**Lines** | [**[]ExpenseResponseLines**](ExpenseResponse_lines.md) | This is a breakdown of the expense lines for this expense, could also be potentially empty for non-card related expenses | [default to null]
**ReceiptIds** | **[]string** | Unique identifiers (UUIDs) for the receipts attached to this expense | [optional] [default to null]
**SettledExpenseIds** | **[]string** | Unique identifiers (UUIDs) of the related settled Expenses in case the expense is a reimbursement of *TRANSFER family | [optional] [default to null]
**Status** | **string** | This is related to export status information | [default to null]
**Supplier** | [***ExpenseResponseSupplier**](ExpenseResponse_supplier.md) |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Date and time this record was created in the pleo system in the format YYYY-MM-DDTHH:mi:ss.SSSZ | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Date and time this record was last updated in the pleo system in the format YYYY-MM-DDTHH:mi:ss.SSSZ | [default to null]
**DeletedAt** | [**time.Time**](time.Time.md) | Date and time this record was deleted in the pleo system in the format YYYY-MM-DDTHH:mi:ss.SSSZ | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

