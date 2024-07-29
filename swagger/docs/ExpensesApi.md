# {{classname}}

All URIs are relative to *https://openapi.pleo.io/v1/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExpense**](ExpensesApi.md#GetExpense) | **Get** /expenses/{expenseId} | Get an expense for a company
[**GetExpenses**](ExpensesApi.md#GetExpenses) | **Get** /expenses | Get expenses for a company
[**UpdateExpenses**](ExpensesApi.md#UpdateExpenses) | **Put** /expenses | Update a given list of expenses. 

# **GetExpense**
> ExpenseResponse GetExpense(ctx, expenseId)
Get an expense for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **expenseId** | **string**| Unique UUID of the expense. | 

### Return type

[**ExpenseResponse**](ExpenseResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExpenses**
> ExpensesResponse GetExpenses(ctx, optional)
Get expenses for a company

By default does not return expenses with Transaction State AUTHORIZATION or AUTHORIZATION_REVERSAL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExpensesApiGetExpensesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExpensesApiGetExpensesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageOffset** | **optional.Float64**| This is the pagination offset value. The record number you would like to start from. This offset value starts at 0. | 
 **pageSize** | **optional.Float64**| The number of expenses to return for each page. | 
 **dateFrom** | **optional.Time**| Date and time to start the expense search from. Format is: YYYY-MM-DDTHH:mi:ss.SSSZ | 
 **dateTo** | **optional.Time**| Date and time to start the expense search from. Format is: YYYY-MM-DDTHH:mi:ss.SSSZ | 
 **types** | [**optional.Interface of []string**](string.md)| An array of the expense types to filter on. | 
 **accountId** | **optional.String**| Unique UUID of the expense category account to filter the expenses by. | 
 **tagId** | **optional.String**| Unique UUID of the tag the expense belongs to, to filter the expenses by. | 
 **transactionStates** | [**optional.Interface of []string**](string.md)| An array of transaction states to filter the expenses by. | 
 **status** | [**optional.Interface of []string**](string.md)| An array of the export status to filter on. | 

### Return type

[**ExpensesResponse**](ExpensesResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateExpenses**
> interface{} UpdateExpenses(ctx, body)
Update a given list of expenses. 

At the moment, it is only possible to update expenses' status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExpensesUpdateRequest**](ExpensesUpdateRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

