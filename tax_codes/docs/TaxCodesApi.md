# \TaxCodesApi

All URIs are relative to *https://external.pleo.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaxCode**](TaxCodesApi.md#CreateTaxCode) | **Post** /v0/tax-codes | Create a new tax code
[**DeleteTaxCode**](TaxCodesApi.md#DeleteTaxCode) | **Delete** /v0/tax-codes/{taxCodeId} | Delete an tax code
[**GetTaxCode**](TaxCodesApi.md#GetTaxCode) | **Get** /v0/tax-codes/{taxCodeId} | Returns a given tax code
[**GetTaxCodes**](TaxCodesApi.md#GetTaxCodes) | **Post** /v0/tax-codes:search | Returns a list of tax codes
[**UpdateTaxCode**](TaxCodesApi.md#UpdateTaxCode) | **Put** /v0/tax-codes/{taxCodeId} | Update an tax code



## CreateTaxCode

> DataResponseTaxCodeModel CreateTaxCode(ctx).TaxCodeCreateUpdateModel(taxCodeCreateUpdateModel).Execute()

Create a new tax code

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    taxCodeCreateUpdateModel := *openapiclient.NewTaxCodeCreateUpdateModel(false, "CompanyId_example", "Standardmoms (25 %)", float32(123), "Type_example") // TaxCodeCreateUpdateModel | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCodesApi.CreateTaxCode(context.Background()).TaxCodeCreateUpdateModel(taxCodeCreateUpdateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCodesApi.CreateTaxCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTaxCode`: DataResponseTaxCodeModel
    fmt.Fprintf(os.Stdout, "Response from `TaxCodesApi.CreateTaxCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaxCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxCodeCreateUpdateModel** | [**TaxCodeCreateUpdateModel**](TaxCodeCreateUpdateModel.md) |  | 

### Return type

[**DataResponseTaxCodeModel**](DataResponseTaxCodeModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTaxCode

> DeleteTaxCode(ctx, taxCodeId).Execute()

Delete an tax code

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    taxCodeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCodesApi.DeleteTaxCode(context.Background(), taxCodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCodesApi.DeleteTaxCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxCodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaxCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxCode

> DataResponseTaxCodeModel GetTaxCode(ctx, taxCodeId).Execute()

Returns a given tax code

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    taxCodeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCodesApi.GetTaxCode(context.Background(), taxCodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCodesApi.GetTaxCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaxCode`: DataResponseTaxCodeModel
    fmt.Fprintf(os.Stdout, "Response from `TaxCodesApi.GetTaxCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxCodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataResponseTaxCodeModel**](DataResponseTaxCodeModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxCodes

> CursorPaginatedResponseTaxCodeModel GetTaxCodes(ctx).CompanyId(companyId).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).TaxCodeSearchRequest(taxCodeSearchRequest).Execute()

Returns a list of tax codes

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    before := "before_example" // string | Lower bound of the page of data to return (cannot be used together with [after] or [offset]) (optional)
    after := "after_example" // string | Upper bound of the page of data to return (cannot be used together with [before] or [offset]) (optional)
    offset := int64(789) // int64 | Offset of the page of data to return (cannot be used together with [before] or [after]) (optional)
    limit := int32(56) // int32 | The maximum amount of items to return (optional)
    sortingKeys := []string{"Inner_example"} // []string | The keys to sort the results by (optional)
    sortingOrder := []openapiclient.PageOrder{openapiclient.PageOrder("ASC")} // []PageOrder | The order to sort the results by. Must be the same length as [sortingKeys]; one order per key (optional)
    taxCodeSearchRequest := *openapiclient.NewTaxCodeSearchRequest() // TaxCodeSearchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCodesApi.GetTaxCodes(context.Background()).CompanyId(companyId).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).TaxCodeSearchRequest(taxCodeSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCodesApi.GetTaxCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaxCodes`: CursorPaginatedResponseTaxCodeModel
    fmt.Fprintf(os.Stdout, "Response from `TaxCodesApi.GetTaxCodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 
 **before** | **string** | Lower bound of the page of data to return (cannot be used together with [after] or [offset]) | 
 **after** | **string** | Upper bound of the page of data to return (cannot be used together with [before] or [offset]) | 
 **offset** | **int64** | Offset of the page of data to return (cannot be used together with [before] or [after]) | 
 **limit** | **int32** | The maximum amount of items to return | 
 **sortingKeys** | **[]string** | The keys to sort the results by | 
 **sortingOrder** | [**[]PageOrder**](PageOrder.md) | The order to sort the results by. Must be the same length as [sortingKeys]; one order per key | 
 **taxCodeSearchRequest** | [**TaxCodeSearchRequest**](TaxCodeSearchRequest.md) |  | 

### Return type

[**CursorPaginatedResponseTaxCodeModel**](CursorPaginatedResponseTaxCodeModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaxCode

> DataResponseTaxCodeModel UpdateTaxCode(ctx, taxCodeId).TaxCodeCreateUpdateModel(taxCodeCreateUpdateModel).Execute()

Update an tax code

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    taxCodeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    taxCodeCreateUpdateModel := *openapiclient.NewTaxCodeCreateUpdateModel(false, "CompanyId_example", "Standardmoms (25 %)", float32(123), "Type_example") // TaxCodeCreateUpdateModel | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCodesApi.UpdateTaxCode(context.Background(), taxCodeId).TaxCodeCreateUpdateModel(taxCodeCreateUpdateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCodesApi.UpdateTaxCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTaxCode`: DataResponseTaxCodeModel
    fmt.Fprintf(os.Stdout, "Response from `TaxCodesApi.UpdateTaxCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxCodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaxCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taxCodeCreateUpdateModel** | [**TaxCodeCreateUpdateModel**](TaxCodeCreateUpdateModel.md) |  | 

### Return type

[**DataResponseTaxCodeModel**](DataResponseTaxCodeModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

