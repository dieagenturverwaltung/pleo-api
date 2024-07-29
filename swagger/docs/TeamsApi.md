# {{classname}}

All URIs are relative to *https://openapi.pleo.io/v1/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTeamEmployee**](TeamsApi.md#AddTeamEmployee) | **Put** /teams/{teamId}/employees/{employeeId} | Add an employee to a team
[**CreateTeam**](TeamsApi.md#CreateTeam) | **Post** /teams | Create a team for a company
[**DeleteTeam**](TeamsApi.md#DeleteTeam) | **Delete** /teams/{teamId} | Delete a team for a company
[**DeleteTeamEmployee**](TeamsApi.md#DeleteTeamEmployee) | **Delete** /teams/{teamId}/employees/{employeeId} | Remove an employee from a team
[**GetTeam**](TeamsApi.md#GetTeam) | **Get** /teams/{teamId} | Get a team for a company
[**GetTeams**](TeamsApi.md#GetTeams) | **Get** /teams | Get all teams for a company
[**UpdateTeam**](TeamsApi.md#UpdateTeam) | **Put** /teams/{teamId} | Update a team for a company

# **AddTeamEmployee**
> interface{} AddTeamEmployee(ctx, teamId, employeeId)
Add an employee to a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**| Unique UUID of the team. | 
  **employeeId** | **string**| Unique UUID of the employee. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTeam**
> TeamResponse CreateTeam(ctx, body)
Create a team for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TeamCreationRequest**](TeamCreationRequest.md)|  | 

### Return type

[**TeamResponse**](TeamResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTeam**
> interface{} DeleteTeam(ctx, teamId)
Delete a team for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**| Unique UUID of the team. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTeamEmployee**
> interface{} DeleteTeamEmployee(ctx, teamId, employeeId)
Remove an employee from a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**| Unique UUID of the team. | 
  **employeeId** | **string**| Unique UUID of the employee. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeam**
> TeamResponse GetTeam(ctx, teamId)
Get a team for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**| Unique UUID of the team. | 

### Return type

[**TeamResponse**](TeamResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeams**
> TeamsResponse GetTeams(ctx, optional)
Get all teams for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamsApiGetTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiGetTeamsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageOffset** | **optional.Float64**| This is the pagination offset value. The record number you would like to start from. This offset value starts at 0. | 
 **pageSize** | **optional.Float64**| The number of teams to return for each page. | [default to 25]

### Return type

[**TeamsResponse**](TeamsResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTeam**
> TeamResponse UpdateTeam(ctx, body, teamId)
Update a team for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TeamUpdateRequest**](TeamUpdateRequest.md)|  | 
  **teamId** | **string**| Unique UUID of the team. | 

### Return type

[**TeamResponse**](TeamResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

