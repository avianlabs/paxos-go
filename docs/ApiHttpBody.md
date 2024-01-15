# ApiHttpBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | The HTTP Content-Type header value specifying the content type of the body. | [optional] 
**Data** | Pointer to **string** | The HTTP request/response body as raw binary. | [optional] 
**Extensions** | Pointer to [**[]BufAny**](BufAny.md) | Application specific response metadata. Must be set in the first response for streaming APIs. | [optional] 

## Methods

### NewApiHttpBody

`func NewApiHttpBody() *ApiHttpBody`

NewApiHttpBody instantiates a new ApiHttpBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHttpBodyWithDefaults

`func NewApiHttpBodyWithDefaults() *ApiHttpBody`

NewApiHttpBodyWithDefaults instantiates a new ApiHttpBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *ApiHttpBody) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ApiHttpBody) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ApiHttpBody) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ApiHttpBody) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetData

`func (o *ApiHttpBody) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiHttpBody) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiHttpBody) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ApiHttpBody) HasData() bool`

HasData returns a boolean if a field has been set.

### GetExtensions

`func (o *ApiHttpBody) GetExtensions() []BufAny`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ApiHttpBody) GetExtensionsOk() (*[]BufAny, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ApiHttpBody) SetExtensions(v []BufAny)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ApiHttpBody) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


