# ListTransfersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Transfer**](Transfer.md) |  | [optional] 
**NextPageCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListTransfersResponse

`func NewListTransfersResponse() *ListTransfersResponse`

NewListTransfersResponse instantiates a new ListTransfersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransfersResponseWithDefaults

`func NewListTransfersResponseWithDefaults() *ListTransfersResponse`

NewListTransfersResponseWithDefaults instantiates a new ListTransfersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListTransfersResponse) GetItems() []Transfer`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListTransfersResponse) GetItemsOk() (*[]Transfer, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListTransfersResponse) SetItems(v []Transfer)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListTransfersResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextPageCursor

`func (o *ListTransfersResponse) GetNextPageCursor() string`

GetNextPageCursor returns the NextPageCursor field if non-nil, zero value otherwise.

### GetNextPageCursorOk

`func (o *ListTransfersResponse) GetNextPageCursorOk() (*string, bool)`

GetNextPageCursorOk returns a tuple with the NextPageCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageCursor

`func (o *ListTransfersResponse) SetNextPageCursor(v string)`

SetNextPageCursor sets NextPageCursor field to given value.

### HasNextPageCursor

`func (o *ListTransfersResponse) HasNextPageCursor() bool`

HasNextPageCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


