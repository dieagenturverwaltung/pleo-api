# \ExportAPIV0Api

All URIs are relative to *https://external.pleo.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExportJob**](ExportAPIV0Api.md#CreateExportJob) | **Post** /v0/export-jobs | Create a new export job
[**CreateExportJobEvent**](ExportAPIV0Api.md#CreateExportJobEvent) | **Post** /v0/export-job-events | Create an Export Job Event
[**GetExportJobById**](ExportAPIV0Api.md#GetExportJobById) | **Get** /v0/export-jobs/{jobId} | Get an Export Job
[**GetExportJobItems**](ExportAPIV0Api.md#GetExportJobItems) | **Get** /v0/export-jobs/{jobId}/items | Get Export Job Items
[**GetExportJobsList**](ExportAPIV0Api.md#GetExportJobsList) | **Get** /v0/export-jobs | Get a list of Export Jobs
[**UpdateExportJobItems**](ExportAPIV0Api.md#UpdateExportJobItems) | **Put** /v0/export-jobs/{jobId}/items | Update Export Job Items



## CreateExportJob

> DataResponseExportJob CreateExportJob(ctx).CreateExportJob(createExportJob).Execute()

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
    resp, r, err := apiClient.ExportAPIV0Api.CreateExportJob(context.Background()).CreateExportJob(createExportJob).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportAPIV0Api.CreateExportJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExportJob`: DataResponseExportJob
    fmt.Fprintf(os.Stdout, "Response from `ExportAPIV0Api.CreateExportJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExportJobRequest struct via the builder pattern


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


## CreateExportJobEvent

> CreateExportJobEvent(ctx).CreateExportJobEvent(createExportJobEvent).Execute()

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
    resp, r, err := apiClient.ExportAPIV0Api.CreateExportJobEvent(context.Background()).CreateExportJobEvent(createExportJobEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportAPIV0Api.CreateExportJobEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExportJobEventRequest struct via the builder pattern


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


## GetExportJobById

> DataResponseExportJob GetExportJobById(ctx, jobId).Execute()

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
    resp, r, err := apiClient.ExportAPIV0Api.GetExportJobById(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportAPIV0Api.GetExportJobById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExportJobById`: DataResponseExportJob
    fmt.Fprintf(os.Stdout, "Response from `ExportAPIV0Api.GetExportJobById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportJobByIdRequest struct via the builder pattern


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


## GetExportJobItems

> CursorPaginatedResponseExportJobItem GetExportJobItems(ctx, jobId).Status(status).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).Execute()

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
    resp, r, err := apiClient.ExportAPIV0Api.GetExportJobItems(context.Background(), jobId).Status(status).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportAPIV0Api.GetExportJobItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExportJobItems`: CursorPaginatedResponseExportJobItem
    fmt.Fprintf(os.Stdout, "Response from `ExportAPIV0Api.GetExportJobItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportJobItemsRequest struct via the builder pattern


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


## GetExportJobsList

> CursorPaginatedResponseExportJob GetExportJobsList(ctx).CompanyId(companyId).Status(status).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).Execute()

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
    resp, r, err := apiClient.ExportAPIV0Api.GetExportJobsList(context.Background()).CompanyId(companyId).Status(status).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportAPIV0Api.GetExportJobsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExportJobsList`: CursorPaginatedResponseExportJob
    fmt.Fprintf(os.Stdout, "Response from `ExportAPIV0Api.GetExportJobsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExportJobsListRequest struct via the builder pattern


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


## UpdateExportJobItems

> ExportJobItemUpdate UpdateExportJobItems(ctx, jobId).UpdateExportJobItem(updateExportJobItem).Execute()

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
    resp, r, err := apiClient.ExportAPIV0Api.UpdateExportJobItems(context.Background(), jobId).UpdateExportJobItem(updateExportJobItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportAPIV0Api.UpdateExportJobItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExportJobItems`: ExportJobItemUpdate
    fmt.Fprintf(os.Stdout, "Response from `ExportAPIV0Api.UpdateExportJobItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExportJobItemsRequest struct via the builder pattern


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

