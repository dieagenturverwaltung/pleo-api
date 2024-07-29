# {{classname}}

All URIs are relative to *https://openapi.pleo.io/v1/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountGroupV11**](DeprecatedApi.md#CreateAccountGroupV11) | **Post** /account-categories | Create an account group for a company
[**DeleteAccountGroupV11**](DeprecatedApi.md#DeleteAccountGroupV11) | **Delete** /account-categories/{accountCategoryId} | Delete an account group for a company
[**GetAccountGroupV11**](DeprecatedApi.md#GetAccountGroupV11) | **Get** /account-categories/{accountCategoryId} | Get an account group for a company
[**GetAccountGroupsV11**](DeprecatedApi.md#GetAccountGroupsV11) | **Get** /account-categories | Get all company account groups
[**UpdateAccountGroupV11**](DeprecatedApi.md#UpdateAccountGroupV11) | **Put** /account-categories/{accountCategoryId} | Update an account group for a company

# **CreateAccountGroupV11**
> AccountGroupResponse CreateAccountGroupV11(ctx, body)
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

# **DeleteAccountGroupV11**
> interface{} DeleteAccountGroupV11(ctx, accountCategoryId)
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

# **GetAccountGroupV11**
> GetAccountGroupResponse GetAccountGroupV11(ctx, accountCategoryId)
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

# **GetAccountGroupsV11**
> AccountGroupsResponse GetAccountGroupsV11(ctx, optional)
Get all company account groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeprecatedApiGetAccountGroupsV11Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeprecatedApiGetAccountGroupsV11Opts struct
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

# **UpdateAccountGroupV11**
> AccountGroupResponse UpdateAccountGroupV11(ctx, body, accountCategoryId)
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

