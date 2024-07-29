# {{classname}}

All URIs are relative to *https://openapi.pleo.io/v1/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExpenseReceipt**](ReceiptsApi.md#GetExpenseReceipt) | **Get** /expenses/{expenseId}/receipts/{receiptId} | Get an expense receipt
[**GetExpenseReceipts**](ReceiptsApi.md#GetExpenseReceipts) | **Get** /expenses/{expenseId}/receipts | Get expense receipts

# **GetExpenseReceipt**
> Receipt GetExpenseReceipt(ctx, expenseId, receiptId)
Get an expense receipt

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **expenseId** | **string**| Unique UUID of the expense. | 
  **receiptId** | **string**| Unique UUID of the receipt. | 

### Return type

[**Receipt**](Receipt.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExpenseReceipts**
> []ReceiptsInner GetExpenseReceipts(ctx, expenseId)
Get expense receipts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **expenseId** | **string**| Unique UUID of the expense. | 

### Return type

[**[]ReceiptsInner**](array.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

