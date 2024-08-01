# \InstallationsApi

All URIs are relative to *https://external.pleo.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateMyInstallation**](InstallationsApi.md#ActivateMyInstallation) | **Post** /v1/installations/me:activate | Activate the installation for a client
[**CompleteInstallationExternalClientOnlyIdempotently**](InstallationsApi.md#CompleteInstallationExternalClientOnlyIdempotently) | **Put** /v1/installations/completions | Complete the installation of a client
[**CreateClientInstallation**](InstallationsApi.md#CreateClientInstallation) | **Post** /v1/installations/me | Create a client installation
[**DeleteClientInstallation**](InstallationsApi.md#DeleteClientInstallation) | **Delete** /v1/installations/me | Delete client&#39;s installation
[**GetClientInstallation**](InstallationsApi.md#GetClientInstallation) | **Get** /v1/installations/me | Get client installation
[**UpdateClientInstallation**](InstallationsApi.md#UpdateClientInstallation) | **Put** /v1/installations/me | Update a client installation



## ActivateMyInstallation

> InstallationResponse ActivateMyInstallation(ctx).Execute()

Activate the installation for a client



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstallationsApi.ActivateMyInstallation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstallationsApi.ActivateMyInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateMyInstallation`: InstallationResponse
    fmt.Fprintf(os.Stdout, "Response from `InstallationsApi.ActivateMyInstallation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiActivateMyInstallationRequest struct via the builder pattern


### Return type

[**InstallationResponse**](InstallationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteInstallationExternalClientOnlyIdempotently

> InstallationResponse CompleteInstallationExternalClientOnlyIdempotently(ctx).Execute()

Complete the installation of a client



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstallationsApi.CompleteInstallationExternalClientOnlyIdempotently(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstallationsApi.CompleteInstallationExternalClientOnlyIdempotently``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompleteInstallationExternalClientOnlyIdempotently`: InstallationResponse
    fmt.Fprintf(os.Stdout, "Response from `InstallationsApi.CompleteInstallationExternalClientOnlyIdempotently`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteInstallationExternalClientOnlyIdempotentlyRequest struct via the builder pattern


### Return type

[**InstallationResponse**](InstallationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClientInstallation

> InstallationResponse CreateClientInstallation(ctx).CreateClientInstallationRequest(createClientInstallationRequest).Execute()

Create a client installation



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
    createClientInstallationRequest := *openapiclient.NewCreateClientInstallationRequest() // CreateClientInstallationRequest | Installation request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstallationsApi.CreateClientInstallation(context.Background()).CreateClientInstallationRequest(createClientInstallationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstallationsApi.CreateClientInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateClientInstallation`: InstallationResponse
    fmt.Fprintf(os.Stdout, "Response from `InstallationsApi.CreateClientInstallation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientInstallationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createClientInstallationRequest** | [**CreateClientInstallationRequest**](CreateClientInstallationRequest.md) | Installation request | 

### Return type

[**InstallationResponse**](InstallationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClientInstallation

> DeleteClientInstallation(ctx).Execute()

Delete client's installation



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstallationsApi.DeleteClientInstallation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstallationsApi.DeleteClientInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientInstallationRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientInstallation

> InstallationResponse GetClientInstallation(ctx).Execute()

Get client installation



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstallationsApi.GetClientInstallation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstallationsApi.GetClientInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientInstallation`: InstallationResponse
    fmt.Fprintf(os.Stdout, "Response from `InstallationsApi.GetClientInstallation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientInstallationRequest struct via the builder pattern


### Return type

[**InstallationResponse**](InstallationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClientInstallation

> InstallationResponse UpdateClientInstallation(ctx).UpdateInstallationRequest(updateInstallationRequest).Execute()

Update a client installation



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
    updateInstallationRequest := *openapiclient.NewUpdateInstallationRequest("ACTIVATED") // UpdateInstallationRequest | Installation request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstallationsApi.UpdateClientInstallation(context.Background()).UpdateInstallationRequest(updateInstallationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstallationsApi.UpdateClientInstallation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClientInstallation`: InstallationResponse
    fmt.Fprintf(os.Stdout, "Response from `InstallationsApi.UpdateClientInstallation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientInstallationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateInstallationRequest** | [**UpdateInstallationRequest**](UpdateInstallationRequest.md) | Installation request | 

### Return type

[**InstallationResponse**](InstallationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

