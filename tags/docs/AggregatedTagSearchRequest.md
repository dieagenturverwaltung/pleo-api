# AggregatedTagSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagIds** | Pointer to **[]string** | Useful for fetching a list of tags given a specific list of IDs. | [optional] 
**TextSearch** | Pointer to **string** | Free text search for tags by name | [optional] 

## Methods

### NewAggregatedTagSearchRequest

`func NewAggregatedTagSearchRequest() *AggregatedTagSearchRequest`

NewAggregatedTagSearchRequest instantiates a new AggregatedTagSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregatedTagSearchRequestWithDefaults

`func NewAggregatedTagSearchRequestWithDefaults() *AggregatedTagSearchRequest`

NewAggregatedTagSearchRequestWithDefaults instantiates a new AggregatedTagSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagIds

`func (o *AggregatedTagSearchRequest) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *AggregatedTagSearchRequest) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *AggregatedTagSearchRequest) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *AggregatedTagSearchRequest) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTextSearch

`func (o *AggregatedTagSearchRequest) GetTextSearch() string`

GetTextSearch returns the TextSearch field if non-nil, zero value otherwise.

### GetTextSearchOk

`func (o *AggregatedTagSearchRequest) GetTextSearchOk() (*string, bool)`

GetTextSearchOk returns a tuple with the TextSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextSearch

`func (o *AggregatedTagSearchRequest) SetTextSearch(v string)`

SetTextSearch sets TextSearch field to given value.

### HasTextSearch

`func (o *AggregatedTagSearchRequest) HasTextSearch() bool`

HasTextSearch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


