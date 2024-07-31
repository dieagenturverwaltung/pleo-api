# \ExportItemsV0Api

All URIs are relative to *https://external.pleo.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExportItems**](ExportItemsV0Api.md#GetExportItems) | **Get** /v0/export-items | Get Export Items



## GetExportItems

> CursorPaginatedResponseExportItem GetExportItems(ctx).JobId(jobId).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).Execute()

Get Export Items



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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job ID which the accounting entries are linked to (optional)
    before := "before_example" // string | Lower bound of the page of data to return (cannot be used together with [after] or [offset]) (optional)
    after := "after_example" // string | Upper bound of the page of data to return (cannot be used together with [before] or [offset]) (optional)
    offset := int64(789) // int64 | Offset of the page of data to return (cannot be used together with [before] or [after]) (optional)
    limit := int32(56) // int32 | The maximum amount of items to return (optional)
    sortingKeys := []string{"Inner_example"} // []string | The keys to sort the results by (optional)
    sortingOrder := []openapiclient.PageOrder{openapiclient.PageOrder("ASC")} // []PageOrder | The order to sort the results by. Must be the same length as [sortingKeys]; one order per key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportItemsV0Api.GetExportItems(context.Background()).JobId(jobId).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportItemsV0Api.GetExportItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExportItems`: CursorPaginatedResponseExportItem
    fmt.Fprintf(os.Stdout, "Response from `ExportItemsV0Api.GetExportItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExportItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **string** | Job ID which the accounting entries are linked to | 
 **before** | **string** | Lower bound of the page of data to return (cannot be used together with [after] or [offset]) | 
 **after** | **string** | Upper bound of the page of data to return (cannot be used together with [before] or [offset]) | 
 **offset** | **int64** | Offset of the page of data to return (cannot be used together with [before] or [after]) | 
 **limit** | **int32** | The maximum amount of items to return | 
 **sortingKeys** | **[]string** | The keys to sort the results by | 
 **sortingOrder** | [**[]PageOrder**](PageOrder.md) | The order to sort the results by. Must be the same length as [sortingKeys]; one order per key | 

### Return type

[**CursorPaginatedResponseExportItem**](CursorPaginatedResponseExportItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

