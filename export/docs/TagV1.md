# TagV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | This is the identifier to the specific cost allocation item. e.g. the specific project, department, etc. | 
**GroupCode** | **string** | Allocation grouping, e.g. project, cost center, department or if the if dimensions are being used, this would be the dimension group identifier. | 
**Id** | **string** | This is the Pleo internal identifier of the tag | 

## Methods

### NewTagV1

`func NewTagV1(code string, groupCode string, id string, ) *TagV1`

NewTagV1 instantiates a new TagV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagV1WithDefaults

`func NewTagV1WithDefaults() *TagV1`

NewTagV1WithDefaults instantiates a new TagV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TagV1) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TagV1) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TagV1) SetCode(v string)`

SetCode sets Code field to given value.


### GetGroupCode

`func (o *TagV1) GetGroupCode() string`

GetGroupCode returns the GroupCode field if non-nil, zero value otherwise.

### GetGroupCodeOk

`func (o *TagV1) GetGroupCodeOk() (*string, bool)`

GetGroupCodeOk returns a tuple with the GroupCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCode

`func (o *TagV1) SetGroupCode(v string)`

SetGroupCode sets GroupCode field to given value.


### GetId

`func (o *TagV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagV1) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


