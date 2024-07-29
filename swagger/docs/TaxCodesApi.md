# {{classname}}

All URIs are relative to *https://openapi.pleo.io/v1/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTaxCode**](TaxCodesApi.md#GetTaxCode) | **Get** /tax-codes/{taxCodeId} | Get a tax code for a company
[**GetTaxCodes**](TaxCodesApi.md#GetTaxCodes) | **Get** /tax-codes | Get tax codes for a company

# **GetTaxCode**
> TaxCodeResponse GetTaxCode(ctx, taxCodeId)
Get a tax code for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taxCodeId** | **string**| Unique UUID of the tax code. | 

### Return type

[**TaxCodeResponse**](TaxCodeResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaxCodes**
> TaxCodesResponse GetTaxCodes(ctx, optional)
Get tax codes for a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TaxCodesApiGetTaxCodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaxCodesApiGetTaxCodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **force** | **optional.Bool**| Force tax code list to reload. | 
 **hidden** | **optional.Bool**| Include hidden tax codes. | 

### Return type

[**TaxCodesResponse**](TaxCodesResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

