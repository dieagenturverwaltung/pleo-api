# {{classname}}

All URIs are relative to *https://openapi.pleo.io/v1/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTag**](TagGroupsApi.md#CreateTag) | **Post** /tag-groups/{tagGroupId}/tags | Create a new tag
[**CreateTagGroup**](TagGroupsApi.md#CreateTagGroup) | **Post** /tag-groups | Create a new tag group.
[**DeleteTag**](TagGroupsApi.md#DeleteTag) | **Delete** /tag-groups/{tagGroupId}/tags/{tagId} | Delete a tag
[**DeleteTagGroup**](TagGroupsApi.md#DeleteTagGroup) | **Delete** /tag-groups/{tagGroupId} | Delete or archive a tag group.
[**GetTag**](TagGroupsApi.md#GetTag) | **Get** /tag-groups/{tagGroupId}/tags/{tagId} | Get a tag from a tag group
[**GetTagGroup**](TagGroupsApi.md#GetTagGroup) | **Get** /tag-groups/{tagGroupId} | Get a tag group
[**GetTagGroups**](TagGroupsApi.md#GetTagGroups) | **Get** /tag-groups | Get all tag groups belonging to the specified company.
[**GetTags**](TagGroupsApi.md#GetTags) | **Get** /tag-groups/{tagGroupId}/tags | Get all tags for a specified tag group
[**UpdateTag**](TagGroupsApi.md#UpdateTag) | **Patch** /tag-groups/{tagGroupId}/tags/{tagId} | Update a tag
[**UpdateTagGroup**](TagGroupsApi.md#UpdateTagGroup) | **Patch** /tag-groups/{tagGroupId} | Update a tag group

# **CreateTag**
> TagResponse CreateTag(ctx, body, tagGroupId)
Create a new tag

This endpoint allows for the creation of a new Tag and its attribute values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagCreationRequest**](TagCreationRequest.md)|  | 
  **tagGroupId** | **string**| Unique UUID of the tag group. | 

### Return type

[**TagResponse**](TagResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTagGroup**
> TagGroupResponse CreateTagGroup(ctx, body)
Create a new tag group.

This endpoint allows for the creation of a new Tag Group with its attribute values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagGroupCreationRequest**](TagGroupCreationRequest.md)|  | 

### Return type

[**TagGroupResponse**](TagGroupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTag**
> interface{} DeleteTag(ctx, tagGroupId, tagId)
Delete a tag

. If this tag is not already in use, the tag will be deleted, otherwise an error will be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagGroupId** | **string**| Unique UUID of the tag group. | 
  **tagId** | **string**| Unique UUID of the tag. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTagGroup**
> interface{} DeleteTagGroup(ctx, tagGroupId, optional)
Delete or archive a tag group.

This endpoint will delete or archive a Tag Group and all associated attributes and Tags.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagGroupId** | **string**| Unique UUID of the tag group. | 
 **optional** | ***TagGroupsApiDeleteTagGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagGroupsApiDeleteTagGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archive** | **optional.Bool**| This flag allows you to archive the tag group as opposed to a forced delete of the tag group and it&#x27;s associations. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTag**
> TagResponse GetTag(ctx, tagGroupId, tagId)
Get a tag from a tag group

This endpoint returns a specific tag with it's attribute values for a specified tag group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagGroupId** | **string**| Unique UUID of the tag group. | 
  **tagId** | **string**| Unique UUID of the tag. | 

### Return type

[**TagResponse**](TagResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagGroup**
> TagGroupResponse GetTagGroup(ctx, tagGroupId)
Get a tag group

This endpoint will return a single Tag Group, as well as all of it's attributes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagGroupId** | **string**| Unique UUID of the tag group. | 

### Return type

[**TagGroupResponse**](TagGroupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagGroups**
> []TagGroupsResponseInner GetTagGroups(ctx, optional)
Get all tag groups belonging to the specified company.

This endpoint returns all tag groups belonging to the company as well as their attributes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TagGroupsApiGetTagGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagGroupsApiGetTagGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archived** | **optional.Bool**| This flag allows you fetch only tag groups that have been archived if set to true or those that are still actively in use if set to false. By default it will only fetch active tag groups. | 

### Return type

[**[]TagGroupsResponseInner**](array.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTags**
> []TagsResponseInner GetTags(ctx, tagGroupId, optional)
Get all tags for a specified tag group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagGroupId** | **string**| Unique UUID of the tag group. | 
 **optional** | ***TagGroupsApiGetTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagGroupsApiGetTagsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hidden** | **optional.Bool**| This flag allows you fetch only tags that have been hidden/archived if set to true or tags that are still actively in use if set to false. By default it will only fetch active tags. | 

### Return type

[**[]TagsResponseInner**](array.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTag**
> TagResponse UpdateTag(ctx, body, tagGroupId, tagId)
Update a tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagUpdateRequest**](TagUpdateRequest.md)|  | 
  **tagGroupId** | **string**| Unique UUID of the tag group. | 
  **tagId** | **string**| Unique UUID of the tag. | 

### Return type

[**TagResponse**](TagResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTagGroup**
> TagGroupResponse UpdateTagGroup(ctx, body, tagGroupId)
Update a tag group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateTagGroupRequest**](UpdateTagGroupRequest.md)|  | 
  **tagGroupId** | **string**| Unique UUID of the tag group. | 

### Return type

[**TagGroupResponse**](TagGroupResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

