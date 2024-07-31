# \TagGroupsApi

All URIs are relative to *https://external.pleo.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTagGroup**](TagGroupsApi.md#CreateTagGroup) | **Post** /v0/tag-groups | Creates a new tag group resource
[**CreateTagGroupDimension**](TagGroupsApi.md#CreateTagGroupDimension) | **Post** /v0/tag-groups/{groupId}/dimensions | Creates a new tag group dimension
[**DeleteTagGroup**](TagGroupsApi.md#DeleteTagGroup) | **Delete** /v0/tag-groups/{groupId} | Deletes the tag group resource by id
[**DeleteTagGroupDimension**](TagGroupsApi.md#DeleteTagGroupDimension) | **Delete** /v0/tag-groups/{groupId}/dimensions/{dimensionId} | Delete a tag group dimension
[**GetAggregatedTagGroups**](TagGroupsApi.md#GetAggregatedTagGroups) | **Get** /v0/aggregations/tag-groups | Returns all tag groups for given company in an aggregated format
[**GetTagGroup**](TagGroupsApi.md#GetTagGroup) | **Get** /v0/tag-groups/{groupId} | Returns the tag group resource by id
[**GetTagGroupDimension**](TagGroupsApi.md#GetTagGroupDimension) | **Get** /v0/tag-groups/{groupId}/dimensions/{dimensionId} | Get a tag group dimension
[**GetTagGroupDimensions**](TagGroupsApi.md#GetTagGroupDimensions) | **Get** /v0/tag-groups/{groupId}/dimensions | Get a list of tag group dimensions
[**GetTagGroups**](TagGroupsApi.md#GetTagGroups) | **Get** /v0/tag-groups | Returns all tag groups for given company
[**UpdateTagGroup**](TagGroupsApi.md#UpdateTagGroup) | **Put** /v0/tag-groups/{groupId} | Updates the tag group resource by id
[**UpdateTagGroupDimension**](TagGroupsApi.md#UpdateTagGroupDimension) | **Put** /v0/tag-groups/{groupId}/dimensions/{dimensionId} | Update a tag group dimension



## CreateTagGroup

> DataResponseTagGroupModel CreateTagGroup(ctx).CompanyId(companyId).TagGroupCreateModel(tagGroupCreateModel).Execute()

Creates a new tag group resource

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
    tagGroupCreateModel := *openapiclient.NewTagGroupCreateModel(false, "Code_example", "Name_example") // TagGroupCreateModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagGroupsApi.CreateTagGroup(context.Background()).CompanyId(companyId).TagGroupCreateModel(tagGroupCreateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagGroupsApi.CreateTagGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTagGroup`: DataResponseTagGroupModel
    fmt.Fprintf(os.Stdout, "Response from `TagGroupsApi.CreateTagGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 
 **tagGroupCreateModel** | [**TagGroupCreateModel**](TagGroupCreateModel.md) |  | 

### Return type

[**DataResponseTagGroupModel**](DataResponseTagGroupModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTagGroupDimension

> DataResponseTagGroupDimensionModel CreateTagGroupDimension(ctx, groupId).TagGroupDimensionCreateModel(tagGroupDimensionCreateModel).Execute()

Creates a new tag group dimension

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
    tagGroupDimensionCreateModel := *openapiclient.NewTagGroupDimensionCreateModel("5524f270-6c21-11ee-b962-0242ac120002", int32(1), "Project", false) // TagGroupDimensionCreateModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagGroupsApi.CreateTagGroupDimension(context.Background(), groupId).TagGroupDimensionCreateModel(tagGroupDimensionCreateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagGroupsApi.CreateTagGroupDimension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTagGroupDimension`: DataResponseTagGroupDimensionModel
    fmt.Fprintf(os.Stdout, "Response from `TagGroupsApi.CreateTagGroupDimension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagGroupDimensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagGroupDimensionCreateModel** | [**TagGroupDimensionCreateModel**](TagGroupDimensionCreateModel.md) |  | 

### Return type

[**DataResponseTagGroupDimensionModel**](DataResponseTagGroupDimensionModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTagGroup

> bool DeleteTagGroup(ctx, groupId).Execute()

Deletes the tag group resource by id

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagGroupsApi.DeleteTagGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagGroupsApi.DeleteTagGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTagGroup`: bool
    fmt.Fprintf(os.Stdout, "Response from `TagGroupsApi.DeleteTagGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTagGroupDimension

> DeleteTagGroupDimension(ctx, groupId, dimensionId).Execute()

Delete a tag group dimension

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
    dimensionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagGroupsApi.DeleteTagGroupDimension(context.Background(), groupId, dimensionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagGroupsApi.DeleteTagGroupDimension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 
**dimensionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagGroupDimensionRequest struct via the builder pattern


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


## GetAggregatedTagGroups

> DataResponseListAggregatedTagGroupModel GetAggregatedTagGroups(ctx).CompanyId(companyId).OrganizationId(organizationId).TagGroupIds(tagGroupIds).Execute()

Returns all tag groups for given company in an aggregated format



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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    tagGroupIds := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagGroupsApi.GetAggregatedTagGroups(context.Background()).CompanyId(companyId).OrganizationId(organizationId).TagGroupIds(tagGroupIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagGroupsApi.GetAggregatedTagGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAggregatedTagGroups`: DataResponseListAggregatedTagGroupModel
    fmt.Fprintf(os.Stdout, "Response from `TagGroupsApi.GetAggregatedTagGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAggregatedTagGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 
 **organizationId** | **string** |  | 
 **tagGroupIds** | **[]string** |  | 

### Return type

[**DataResponseListAggregatedTagGroupModel**](DataResponseListAggregatedTagGroupModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagGroup

> DataResponseTagGroupModel GetTagGroup(ctx, groupId).Execute()

Returns the tag group resource by id

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagGroupsApi.GetTagGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagGroupsApi.GetTagGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagGroup`: DataResponseTagGroupModel
    fmt.Fprintf(os.Stdout, "Response from `TagGroupsApi.GetTagGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataResponseTagGroupModel**](DataResponseTagGroupModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagGroupDimension

> DataResponseTagGroupDimensionModel GetTagGroupDimension(ctx, groupId, dimensionId).Execute()

Get a tag group dimension

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
    dimensionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagGroupsApi.GetTagGroupDimension(context.Background(), groupId, dimensionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagGroupsApi.GetTagGroupDimension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagGroupDimension`: DataResponseTagGroupDimensionModel
    fmt.Fprintf(os.Stdout, "Response from `TagGroupsApi.GetTagGroupDimension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 
**dimensionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagGroupDimensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataResponseTagGroupDimensionModel**](DataResponseTagGroupDimensionModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagGroupDimensions

> DataResponseListTagGroupDimensionModel GetTagGroupDimensions(ctx, groupId).Execute()

Get a list of tag group dimensions



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagGroupsApi.GetTagGroupDimensions(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagGroupsApi.GetTagGroupDimensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagGroupDimensions`: DataResponseListTagGroupDimensionModel
    fmt.Fprintf(os.Stdout, "Response from `TagGroupsApi.GetTagGroupDimensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagGroupDimensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataResponseListTagGroupDimensionModel**](DataResponseListTagGroupDimensionModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagGroups

> DataResponseListTagGroupModel GetTagGroups(ctx).CompanyId(companyId).OrganizationId(organizationId).Execute()

Returns all tag groups for given company



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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagGroupsApi.GetTagGroups(context.Background()).CompanyId(companyId).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagGroupsApi.GetTagGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagGroups`: DataResponseListTagGroupModel
    fmt.Fprintf(os.Stdout, "Response from `TagGroupsApi.GetTagGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTagGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 
 **organizationId** | **string** |  | 

### Return type

[**DataResponseListTagGroupModel**](DataResponseListTagGroupModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTagGroup

> DataResponseTagGroupModel UpdateTagGroup(ctx, groupId).TagGroupUpdateModel(tagGroupUpdateModel).Execute()

Updates the tag group resource by id

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
    tagGroupUpdateModel := *openapiclient.NewTagGroupUpdateModel(false, "Code_example", map[string]string{"key": "Inner_example"}, "Name_example") // TagGroupUpdateModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagGroupsApi.UpdateTagGroup(context.Background(), groupId).TagGroupUpdateModel(tagGroupUpdateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagGroupsApi.UpdateTagGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTagGroup`: DataResponseTagGroupModel
    fmt.Fprintf(os.Stdout, "Response from `TagGroupsApi.UpdateTagGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagGroupUpdateModel** | [**TagGroupUpdateModel**](TagGroupUpdateModel.md) |  | 

### Return type

[**DataResponseTagGroupModel**](DataResponseTagGroupModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTagGroupDimension

> DataResponseTagGroupDimensionModel UpdateTagGroupDimension(ctx, groupId, dimensionId).TagGroupDimensionUpdateModel(tagGroupDimensionUpdateModel).Execute()

Update a tag group dimension

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
    dimensionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    tagGroupDimensionUpdateModel := *openapiclient.NewTagGroupDimensionUpdateModel("5524f270-6c21-11ee-b962-0242ac120002", int32(1), "Project", false) // TagGroupDimensionUpdateModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagGroupsApi.UpdateTagGroupDimension(context.Background(), groupId, dimensionId).TagGroupDimensionUpdateModel(tagGroupDimensionUpdateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagGroupsApi.UpdateTagGroupDimension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTagGroupDimension`: DataResponseTagGroupDimensionModel
    fmt.Fprintf(os.Stdout, "Response from `TagGroupsApi.UpdateTagGroupDimension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 
**dimensionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagGroupDimensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tagGroupDimensionUpdateModel** | [**TagGroupDimensionUpdateModel**](TagGroupDimensionUpdateModel.md) |  | 

### Return type

[**DataResponseTagGroupDimensionModel**](DataResponseTagGroupDimensionModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

