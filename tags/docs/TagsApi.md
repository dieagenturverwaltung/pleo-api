# \TagsApi

All URIs are relative to *https://external.pleo.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDimensionValue**](TagsApi.md#CreateDimensionValue) | **Post** /v0/tags/{tagId}/dimensions/{dimensionId} | Creates a dimension value
[**CreateTag**](TagsApi.md#CreateTag) | **Post** /v0/tag-groups/{groupId}/tags | Creates a new tag sub-resource under the given tag group
[**DeleteDimensionValue**](TagsApi.md#DeleteDimensionValue) | **Delete** /v0/tags/{tagId}/dimensions/{dimensionId} | Deletes a dimension value
[**DeleteTag**](TagsApi.md#DeleteTag) | **Delete** /v0/tags/{tagId} | Deletes tag by id
[**GetDimensionValue**](TagsApi.md#GetDimensionValue) | **Get** /v0/tags/{tagId}/dimensions/{dimensionId} | Gets a dimension value
[**GetTag**](TagsApi.md#GetTag) | **Get** /v0/tags/{tagId} | Returns tag by id
[**GetTagDimensionValues**](TagsApi.md#GetTagDimensionValues) | **Get** /v0/tags/{tagId}/dimensions | Fetches dimension values
[**GetTags**](TagsApi.md#GetTags) | **Get** /v0/tag-groups/{groupId}/tags | Returns tags for given tag group
[**SearchAggregatedTags**](TagsApi.md#SearchAggregatedTags) | **Post** /v0/aggregations/tags | Search and return aggregated tags
[**SearchTags**](TagsApi.md#SearchTags) | **Get** /v0/tags | Search tags
[**UpdateDimensionValue**](TagsApi.md#UpdateDimensionValue) | **Put** /v0/tags/{tagId}/dimensions/{dimensionId} | Updates a dimension value
[**UpdateTag**](TagsApi.md#UpdateTag) | **Put** /v0/tags/{tagId} | Updates tag by id



## CreateDimensionValue

> DataResponseTagDimensionValueModel CreateDimensionValue(ctx, tagId, dimensionId).TagDimensionValueCreateModel(tagDimensionValueCreateModel).Execute()

Creates a dimension value

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
    tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    dimensionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    tagDimensionValueCreateModel := *openapiclient.NewTagDimensionValueCreateModel("Value_example") // TagDimensionValueCreateModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.CreateDimensionValue(context.Background(), tagId, dimensionId).TagDimensionValueCreateModel(tagDimensionValueCreateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.CreateDimensionValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDimensionValue`: DataResponseTagDimensionValueModel
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.CreateDimensionValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 
**dimensionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDimensionValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tagDimensionValueCreateModel** | [**TagDimensionValueCreateModel**](TagDimensionValueCreateModel.md) |  | 

### Return type

[**DataResponseTagDimensionValueModel**](DataResponseTagDimensionValueModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTag

> DataResponseTagModel CreateTag(ctx, groupId).TagCreateModel(tagCreateModel).Execute()

Creates a new tag sub-resource under the given tag group

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
    groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    tagCreateModel := *openapiclient.NewTagCreateModel(false, "12345") // TagCreateModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.CreateTag(context.Background(), groupId).TagCreateModel(tagCreateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.CreateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTag`: DataResponseTagModel
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.CreateTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagCreateModel** | [**TagCreateModel**](TagCreateModel.md) |  | 

### Return type

[**DataResponseTagModel**](DataResponseTagModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDimensionValue

> DataResponseTagDimensionValueModel DeleteDimensionValue(ctx, tagId, dimensionId).Execute()

Deletes a dimension value

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
    tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    dimensionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.DeleteDimensionValue(context.Background(), tagId, dimensionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.DeleteDimensionValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDimensionValue`: DataResponseTagDimensionValueModel
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.DeleteDimensionValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 
**dimensionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDimensionValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataResponseTagDimensionValueModel**](DataResponseTagDimensionValueModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTag

> DeleteTag(ctx, tagId).Execute()

Deletes tag by id

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
    tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.DeleteTag(context.Background(), tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.DeleteTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


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


## GetDimensionValue

> DataResponseTagDimensionValueModel GetDimensionValue(ctx, tagId, dimensionId).Execute()

Gets a dimension value

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
    tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    dimensionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.GetDimensionValue(context.Background(), tagId, dimensionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GetDimensionValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDimensionValue`: DataResponseTagDimensionValueModel
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GetDimensionValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 
**dimensionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDimensionValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataResponseTagDimensionValueModel**](DataResponseTagDimensionValueModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTag

> DataResponseTagModel GetTag(ctx, tagId).Execute()

Returns tag by id

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
    tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.GetTag(context.Background(), tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GetTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTag`: DataResponseTagModel
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GetTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataResponseTagModel**](DataResponseTagModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagDimensionValues

> DataResponseListTagDimensionValueModel GetTagDimensionValues(ctx, tagId).Execute()

Fetches dimension values

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
    tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.GetTagDimensionValues(context.Background(), tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GetTagDimensionValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagDimensionValues`: DataResponseListTagDimensionValueModel
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GetTagDimensionValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagDimensionValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataResponseListTagDimensionValueModel**](DataResponseListTagDimensionValueModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> CursorPaginatedResponseTagModel GetTags(ctx, groupId).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).IncludeArchived(includeArchived).Execute()

Returns tags for given tag group

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
    groupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    before := "before_example" // string | Lower bound of the page of data to return (cannot be used together with [after] or [offset]) (optional)
    after := "after_example" // string | Upper bound of the page of data to return (cannot be used together with [before] or [offset]) (optional)
    offset := int64(789) // int64 | Offset of the page of data to return (cannot be used together with [before] or [after]) (optional)
    limit := int32(56) // int32 | The maximum amount of items to return (optional)
    sortingKeys := []string{"Inner_example"} // []string | The keys to sort the results by (optional)
    sortingOrder := []openapiclient.PageOrder{openapiclient.PageOrder("ASC")} // []PageOrder | The order to sort the results by. Must be the same length as [sortingKeys]; one order per key (optional)
    includeArchived := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.GetTags(context.Background(), groupId).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).IncludeArchived(includeArchived).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTags`: CursorPaginatedResponseTagModel
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GetTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **before** | **string** | Lower bound of the page of data to return (cannot be used together with [after] or [offset]) | 
 **after** | **string** | Upper bound of the page of data to return (cannot be used together with [before] or [offset]) | 
 **offset** | **int64** | Offset of the page of data to return (cannot be used together with [before] or [after]) | 
 **limit** | **int32** | The maximum amount of items to return | 
 **sortingKeys** | **[]string** | The keys to sort the results by | 
 **sortingOrder** | [**[]PageOrder**](PageOrder.md) | The order to sort the results by. Must be the same length as [sortingKeys]; one order per key | 
 **includeArchived** | **bool** |  | 

### Return type

[**CursorPaginatedResponseTagModel**](CursorPaginatedResponseTagModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAggregatedTags

> CursorPaginatedResponseAggregatedTagModel SearchAggregatedTags(ctx).CompanyId(companyId).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).AggregatedTagSearchRequest(aggregatedTagSearchRequest).Execute()

Search and return aggregated tags

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
    aggregatedTagSearchRequest := *openapiclient.NewAggregatedTagSearchRequest() // AggregatedTagSearchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.SearchAggregatedTags(context.Background()).CompanyId(companyId).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).AggregatedTagSearchRequest(aggregatedTagSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.SearchAggregatedTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAggregatedTags`: CursorPaginatedResponseAggregatedTagModel
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.SearchAggregatedTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAggregatedTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 
 **before** | **string** | Lower bound of the page of data to return (cannot be used together with [after] or [offset]) | 
 **after** | **string** | Upper bound of the page of data to return (cannot be used together with [before] or [offset]) | 
 **offset** | **int64** | Offset of the page of data to return (cannot be used together with [before] or [after]) | 
 **limit** | **int32** | The maximum amount of items to return | 
 **sortingKeys** | **[]string** | The keys to sort the results by | 
 **sortingOrder** | [**[]PageOrder**](PageOrder.md) | The order to sort the results by. Must be the same length as [sortingKeys]; one order per key | 
 **aggregatedTagSearchRequest** | [**AggregatedTagSearchRequest**](AggregatedTagSearchRequest.md) |  | 

### Return type

[**CursorPaginatedResponseAggregatedTagModel**](CursorPaginatedResponseAggregatedTagModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTags

> CursorPaginatedResponseTagModel SearchTags(ctx).CompanyId(companyId).TextSearch(textSearch).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).Execute()

Search tags

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
    textSearch := "textSearch_example" // string |  (optional)
    before := "before_example" // string | Lower bound of the page of data to return (cannot be used together with [after] or [offset]) (optional)
    after := "after_example" // string | Upper bound of the page of data to return (cannot be used together with [before] or [offset]) (optional)
    offset := int64(789) // int64 | Offset of the page of data to return (cannot be used together with [before] or [after]) (optional)
    limit := int32(56) // int32 | The maximum amount of items to return (optional)
    sortingKeys := []string{"Inner_example"} // []string | The keys to sort the results by (optional)
    sortingOrder := []openapiclient.PageOrder{openapiclient.PageOrder("ASC")} // []PageOrder | The order to sort the results by. Must be the same length as [sortingKeys]; one order per key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.SearchTags(context.Background()).CompanyId(companyId).TextSearch(textSearch).Before(before).After(after).Offset(offset).Limit(limit).SortingKeys(sortingKeys).SortingOrder(sortingOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.SearchTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchTags`: CursorPaginatedResponseTagModel
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.SearchTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 
 **textSearch** | **string** |  | 
 **before** | **string** | Lower bound of the page of data to return (cannot be used together with [after] or [offset]) | 
 **after** | **string** | Upper bound of the page of data to return (cannot be used together with [before] or [offset]) | 
 **offset** | **int64** | Offset of the page of data to return (cannot be used together with [before] or [after]) | 
 **limit** | **int32** | The maximum amount of items to return | 
 **sortingKeys** | **[]string** | The keys to sort the results by | 
 **sortingOrder** | [**[]PageOrder**](PageOrder.md) | The order to sort the results by. Must be the same length as [sortingKeys]; one order per key | 

### Return type

[**CursorPaginatedResponseTagModel**](CursorPaginatedResponseTagModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDimensionValue

> DataResponseTagDimensionValueModel UpdateDimensionValue(ctx, tagId, dimensionId).TagDimensionValueUpdateModel(tagDimensionValueUpdateModel).Execute()

Updates a dimension value

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
    tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    dimensionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    tagDimensionValueUpdateModel := *openapiclient.NewTagDimensionValueUpdateModel("Value_example") // TagDimensionValueUpdateModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.UpdateDimensionValue(context.Background(), tagId, dimensionId).TagDimensionValueUpdateModel(tagDimensionValueUpdateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.UpdateDimensionValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDimensionValue`: DataResponseTagDimensionValueModel
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.UpdateDimensionValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 
**dimensionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDimensionValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tagDimensionValueUpdateModel** | [**TagDimensionValueUpdateModel**](TagDimensionValueUpdateModel.md) |  | 

### Return type

[**DataResponseTagDimensionValueModel**](DataResponseTagDimensionValueModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTag

> DataResponseTagModel UpdateTag(ctx, tagId).TagUpdateModel(tagUpdateModel).Execute()

Updates tag by id

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
    tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    tagUpdateModel := *openapiclient.NewTagUpdateModel(false, "12345") // TagUpdateModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.UpdateTag(context.Background(), tagId).TagUpdateModel(tagUpdateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.UpdateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTag`: DataResponseTagModel
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.UpdateTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagUpdateModel** | [**TagUpdateModel**](TagUpdateModel.md) |  | 

### Return type

[**DataResponseTagModel**](DataResponseTagModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

