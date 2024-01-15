# ListQuoteExecutionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]QuoteExecution**](QuoteExecution.md) |  | [optional] 
**NextPageCursor** | Pointer to **string** | Cursor token required for fetching the next page. | [optional] 

## Methods

### NewListQuoteExecutionsResponse

`func NewListQuoteExecutionsResponse() *ListQuoteExecutionsResponse`

NewListQuoteExecutionsResponse instantiates a new ListQuoteExecutionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListQuoteExecutionsResponseWithDefaults

`func NewListQuoteExecutionsResponseWithDefaults() *ListQuoteExecutionsResponse`

NewListQuoteExecutionsResponseWithDefaults instantiates a new ListQuoteExecutionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ListQuoteExecutionsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListQuoteExecutionsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListQuoteExecutionsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListQuoteExecutionsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItems

`func (o *ListQuoteExecutionsResponse) GetItems() []QuoteExecution`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListQuoteExecutionsResponse) GetItemsOk() (*[]QuoteExecution, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListQuoteExecutionsResponse) SetItems(v []QuoteExecution)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListQuoteExecutionsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextPageCursor

`func (o *ListQuoteExecutionsResponse) GetNextPageCursor() string`

GetNextPageCursor returns the NextPageCursor field if non-nil, zero value otherwise.

### GetNextPageCursorOk

`func (o *ListQuoteExecutionsResponse) GetNextPageCursorOk() (*string, bool)`

GetNextPageCursorOk returns a tuple with the NextPageCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageCursor

`func (o *ListQuoteExecutionsResponse) SetNextPageCursor(v string)`

SetNextPageCursor sets NextPageCursor field to given value.

### HasNextPageCursor

`func (o *ListQuoteExecutionsResponse) HasNextPageCursor() bool`

HasNextPageCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


