# {{classname}}

All URIs are relative to *https://openapi.pleo.io/v1/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmployee**](EmployeesApi.md#CreateEmployee) | **Post** /employees | Create an employee for a company
[**DeleteEmployee**](EmployeesApi.md#DeleteEmployee) | **Delete** /employees/{employeeId} | Delete an employee for a company
[**GetEmployee**](EmployeesApi.md#GetEmployee) | **Get** /employees/{employeeId} | Get an employee for a company
[**GetEmployees**](EmployeesApi.md#GetEmployees) | **Get** /employees | Get company employees
[**UpdateEmployee**](EmployeesApi.md#UpdateEmployee) | **Put** /employees/{employeeId} | Update an employee for a company

# **CreateEmployee**
> EmployeeCreationResponse CreateEmployee(ctx, body)
Create an employee for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmployeeCreationRequest**](EmployeeCreationRequest.md)|  | 

### Return type

[**EmployeeCreationResponse**](EmployeeCreationResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEmployee**
> interface{} DeleteEmployee(ctx, employeeId)
Delete an employee for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **employeeId** | **string**| Unique UUID of the employee. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmployee**
> EmployeeResponse GetEmployee(ctx, employeeId)
Get an employee for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **employeeId** | **string**| Unique UUID of the employee. | 

### Return type

[**EmployeeResponse**](EmployeeResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmployees**
> EmployeesResponse GetEmployees(ctx, optional)
Get company employees

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmployeesApiGetEmployeesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EmployeesApiGetEmployeesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageOffset** | **optional.Float64**| This is the pagination offset value. The record number you would like to start from. This offset value starts at 0. | 
 **pageSize** | **optional.Float64**| The number of employees to return for each page. | 

### Return type

[**EmployeesResponse**](EmployeesResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEmployee**
> EmployeeResponse UpdateEmployee(ctx, body, employeeId)
Update an employee for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EmployeesEmployeeIdBody**](EmployeesEmployeeIdBody.md)|  | 
  **employeeId** | **string**| Unique UUID of the employee. | 

### Return type

[**EmployeeResponse**](EmployeeResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

