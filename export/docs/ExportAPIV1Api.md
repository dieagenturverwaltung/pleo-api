# \ExportAPIV1Api

All URIs are relative to *https://external.pleo.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExportJobEventV1**](ExportAPIV1Api.md#CreateExportJobEventV1) | **Post** /v1/export-job-events | Create an Export Job Event
[**CreateExportJobV1**](ExportAPIV1Api.md#CreateExportJobV1) | **Post** /v1/export-jobs | Create a new export job
[**GetExportJobByIdV1**](ExportAPIV1Api.md#GetExportJobByIdV1) | **Get** /v1/export-jobs/{jobId} | Get an Export Job
[**GetExportJobItemsV1**](ExportAPIV1Api.md#GetExportJobItemsV1) | **Get** /v1/export-jobs/{jobId}/items | Get Export Job Items
[**GetExportJobsListV1**](ExportAPIV1Api.md#GetExportJobsListV1) | **Get** /v1/export-jobs | Get a list of Export Jobs
[**UpdateExportJobItemsV1**](ExportAPIV1Api.md#UpdateExportJobItemsV1) | **Put** /v1/export-jobs/{jobId}/items | Update Export Job Items



## CreateExportJobEventV1

> CreateExportJobEventV1(ctx).CreateExportJobEvent(createExportJobEvent).Execute()

Create an Export Job Event



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
    createExportJobEvent := *openapiclient.NewCreateExportJobEvent("Event_example", "JobId_example") // CreateExportJobEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportAPIV1Api.CreateExportJobEventV1(context.Background()).CreateExportJobEvent(createExportJobEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportAPIV1Api.CreateExportJobEventV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExportJobEventV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createExportJobEvent** | [**CreateExportJobEvent**](CreateExportJobEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExportJobV1

> DataResponseExportJob CreateExportJobV1(ctx).CreateExportJob(createExportJob).Execute()

Create a new export job

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
    createExportJob := *openapiclient.NewCreateExportJob([]string{"AccountingEntryIds_example"}, "CompanyId_example") // CreateExportJob |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportAPIV1Api.CreateExportJobV1(context.Background()).CreateExportJob(createExportJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportAPIV1Api.CreateExportJobV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExportJobV1`: DataResponseExportJob
    fmt.Fprintf(os.Stdout, "Response from `ExportAPIV1Api.CreateExportJobV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExportJobV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createExportJob** | [**CreateExportJob**](CreateExportJob.md) |  | 

### Return type

[**DataResponseExportJob**](DataResponseExportJob.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExportJobByIdV1

> DataResponseExportJob GetExportJobByIdV1(ctx, jobId).Execute()

Get an Export Job



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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportAPIV1Api.GetExportJobByIdV1(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportAPIV1Api.GetExportJobByIdV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExportJobByIdV1`: DataResponseExportJob
    fmt.Fprintf(os.Stdout, "Response from `ExportAPIV1Api.GetExportJobByIdV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportJobByIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataResponseExportJob**](DataResponseExportJob.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExportJobItemsV1

> CursorPaginatedResponseExportJobItem GetExportJobItemsV1(ctx, jobId).Status(status).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).Execute()

Get Export Job Items



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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    status := "status_example" // string | Filter by the export job item status (optional)
    before := "before_example" // string | Lower bound of the page of data to return (cannot be used together with [after] or [offset]) (optional)
    after := "after_example" // string | Upper bound of the page of data to return (cannot be used together with [before] or [offset]) (optional)
    offset := int64(789) // int64 | Offset of the page of data to return (cannot be used together with [before] or [after]) (optional)
    limit := int32(56) // int32 | The maximum amount of items to return (optional)
    sortingKeys := []string{"Inner_example"} // []string | The keys to sort the results by (optional)
    sortingOrder := []openapiclient.PageOrder{openapiclient.PageOrder("ASC")} // []PageOrder | The order to sort the results by. Must be the same length as [sortingKeys]; one order per key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportAPIV1Api.GetExportJobItemsV1(context.Background(), jobId).Status(status).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportAPIV1Api.GetExportJobItemsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExportJobItemsV1`: CursorPaginatedResponseExportJobItem
    fmt.Fprintf(os.Stdout, "Response from `ExportAPIV1Api.GetExportJobItemsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportJobItemsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter by the export job item status | 
 **before** | **string** | Lower bound of the page of data to return (cannot be used together with [after] or [offset]) | 
 **after** | **string** | Upper bound of the page of data to return (cannot be used together with [before] or [offset]) | 
 **offset** | **int64** | Offset of the page of data to return (cannot be used together with [before] or [after]) | 
 **limit** | **int32** | The maximum amount of items to return | 
 **sortingKeys** | **[]string** | The keys to sort the results by | 
 **sortingOrder** | [**[]PageOrder**](PageOrder.md) | The order to sort the results by. Must be the same length as [sortingKeys]; one order per key | 

### Return type

[**CursorPaginatedResponseExportJobItem**](CursorPaginatedResponseExportJobItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExportJobsListV1

> CursorPaginatedResponseExportJob GetExportJobsListV1(ctx).CompanyId(companyId).Status(status).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).Execute()

Get a list of Export Jobs



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
    companyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Fetch a list of jobs for a given company (optional)
    status := "status_example" // string | Fetch a list of jobs for a specific status (optional)
    before := "before_example" // string | Lower bound of the page of data to return (cannot be used together with [after] or [offset]) (optional)
    after := "after_example" // string | Upper bound of the page of data to return (cannot be used together with [before] or [offset]) (optional)
    offset := int64(789) // int64 | Offset of the page of data to return (cannot be used together with [before] or [after]) (optional)
    limit := int32(56) // int32 | The maximum amount of items to return (optional)
    sortingKeys := []string{"Inner_example"} // []string | The keys to sort the results by (optional)
    sortingOrder := []openapiclient.PageOrder{openapiclient.PageOrder("ASC")} // []PageOrder | The order to sort the results by. Must be the same length as [sortingKeys]; one order per key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportAPIV1Api.GetExportJobsListV1(context.Background()).CompanyId(companyId).Status(status).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportAPIV1Api.GetExportJobsListV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExportJobsListV1`: CursorPaginatedResponseExportJob
    fmt.Fprintf(os.Stdout, "Response from `ExportAPIV1Api.GetExportJobsListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExportJobsListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Fetch a list of jobs for a given company | 
 **status** | **string** | Fetch a list of jobs for a specific status | 
 **before** | **string** | Lower bound of the page of data to return (cannot be used together with [after] or [offset]) | 
 **after** | **string** | Upper bound of the page of data to return (cannot be used together with [before] or [offset]) | 
 **offset** | **int64** | Offset of the page of data to return (cannot be used together with [before] or [after]) | 
 **limit** | **int32** | The maximum amount of items to return | 
 **sortingKeys** | **[]string** | The keys to sort the results by | 
 **sortingOrder** | [**[]PageOrder**](PageOrder.md) | The order to sort the results by. Must be the same length as [sortingKeys]; one order per key | 

### Return type

[**CursorPaginatedResponseExportJob**](CursorPaginatedResponseExportJob.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExportJobItemsV1

> ExportJobItemUpdate UpdateExportJobItemsV1(ctx, jobId).UpdateExportJobItem(updateExportJobItem).Execute()

Update Export Job Items



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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    updateExportJobItem := []openapiclient.UpdateExportJobItem{*openapiclient.NewUpdateExportJobItem("AccountingEntryId_example", "Status_example")} // []UpdateExportJobItem |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportAPIV1Api.UpdateExportJobItemsV1(context.Background(), jobId).UpdateExportJobItem(updateExportJobItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportAPIV1Api.UpdateExportJobItemsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExportJobItemsV1`: ExportJobItemUpdate
    fmt.Fprintf(os.Stdout, "Response from `ExportAPIV1Api.UpdateExportJobItemsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExportJobItemsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateExportJobItem** | [**[]UpdateExportJobItem**](UpdateExportJobItem.md) |  | 

### Return type

[**ExportJobItemUpdate**](ExportJobItemUpdate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

