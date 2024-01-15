# ListIdentitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageCursor** | Pointer to **string** | Cursor token required for fetching the next page. | [optional] 
**Items** | Pointer to [**[]Identity**](Identity.md) | The result list of identities. | [optional] 

## Methods

### NewListIdentitiesResponse

`func NewListIdentitiesResponse() *ListIdentitiesResponse`

NewListIdentitiesResponse instantiates a new ListIdentitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIdentitiesResponseWithDefaults

`func NewListIdentitiesResponseWithDefaults() *ListIdentitiesResponse`

NewListIdentitiesResponseWithDefaults instantiates a new ListIdentitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageCursor

`func (o *ListIdentitiesResponse) GetNextPageCursor() string`

GetNextPageCursor returns the NextPageCursor field if non-nil, zero value otherwise.

### GetNextPageCursorOk

`func (o *ListIdentitiesResponse) GetNextPageCursorOk() (*string, bool)`

GetNextPageCursorOk returns a tuple with the NextPageCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageCursor

`func (o *ListIdentitiesResponse) SetNextPageCursor(v string)`

SetNextPageCursor sets NextPageCursor field to given value.

### HasNextPageCursor

`func (o *ListIdentitiesResponse) HasNextPageCursor() bool`

HasNextPageCursor returns a boolean if a field has been set.

### GetItems

`func (o *ListIdentitiesResponse) GetItems() []Identity`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListIdentitiesResponse) GetItemsOk() (*[]Identity, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListIdentitiesResponse) SetItems(v []Identity)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListIdentitiesResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


