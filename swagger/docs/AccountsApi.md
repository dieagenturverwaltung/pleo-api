# {{classname}}

All URIs are relative to *https://openapi.pleo.io/v1/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](AccountsApi.md#CreateAccount) | **Post** /accounts | Create an account for a company
[**CreateAccountGroup**](AccountsApi.md#CreateAccountGroup) | **Post** /account-groups | Create an account group for a company
[**DeleteAccount**](AccountsApi.md#DeleteAccount) | **Delete** /accounts/{accountId} | Delete an account for a company
[**DeleteAccountGroup**](AccountsApi.md#DeleteAccountGroup) | **Delete** /account-groups/{accountCategoryId} | Delete an account group for a company
[**GetAccount**](AccountsApi.md#GetAccount) | **Get** /accounts/{accountId} | Get an account for a company
[**GetAccountGroup**](AccountsApi.md#GetAccountGroup) | **Get** /account-groups/{accountCategoryId} | Get an account group for a company
[**GetAccountGroups**](AccountsApi.md#GetAccountGroups) | **Get** /account-groups | Get all company account groups
[**GetAccounts**](AccountsApi.md#GetAccounts) | **Get** /accounts | Get company accounts
[**UpdateAccount**](AccountsApi.md#UpdateAccount) | **Put** /accounts/{accountId} | Update an account for a company
[**UpdateAccountGroup**](AccountsApi.md#UpdateAccountGroup) | **Put** /account-groups/{accountCategoryId} | Update an account group for a company

# **CreateAccount**
> AccountResponse CreateAccount(ctx, body)
Create an account for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountCreationRequest**](AccountCreationRequest.md)|  | 

### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccountGroup**
> AccountGroupResponse CreateAccountGroup(ctx, body)
Create an account group for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountGroupCreationRequest**](AccountGroupCreationRequest.md)|  | 

### Return type

[**AccountGroupResponse**](AccountGroupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccount**
> interface{} DeleteAccount(ctx, accountId)
Delete an account for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Unique UUID of the account. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountGroup**
> interface{} DeleteAccountGroup(ctx, accountCategoryId)
Delete an account group for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountCategoryId** | **string**| Unique UUID of the account. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccount**
> AccountResponse GetAccount(ctx, accountId)
Get an account for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| Unique UUID of the account. | 

### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountGroup**
> GetAccountGroupResponse GetAccountGroup(ctx, accountCategoryId)
Get an account group for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountCategoryId** | **string**| Unique UUID of the account. | 

### Return type

[**GetAccountGroupResponse**](GetAccountGroupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountGroups**
> AccountGroupsResponse GetAccountGroups(ctx, optional)
Get all company account groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsApiGetAccountGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsApiGetAccountGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageOffset** | **optional.Float64**| This is the pagination offset value. The record number you would like to start from. This offset value starts at 0. | 
 **pageSize** | **optional.Float64**| The number of account groups to return for each page. | 

### Return type

[**AccountGroupsResponse**](AccountGroupsResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccounts**
> AccountsResponse GetAccounts(ctx, optional)
Get company accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountsApiGetAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsApiGetAccountsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageOffset** | **optional.Float64**| This is the pagination offset value. The record number you would like to start from. This offset value starts at 0. | 
 **pageSize** | **optional.Float64**| The number of accounts to return for each page. | 

### Return type

[**AccountsResponse**](AccountsResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccount**
> AccountResponse UpdateAccount(ctx, body, accountId)
Update an account for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountCreationRequest**](AccountCreationRequest.md)|  | 
  **accountId** | **string**| Unique UUID of the account. | 

### Return type

[**AccountResponse**](AccountResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountGroup**
> AccountGroupResponse UpdateAccountGroup(ctx, body, accountCategoryId)
Update an account group for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountGroupCreationRequest**](AccountGroupCreationRequest.md)|  | 
  **accountCategoryId** | **string**| Unique UUID of the account. | 

### Return type

[**AccountGroupResponse**](AccountGroupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

