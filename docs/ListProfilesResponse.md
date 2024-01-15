# ListProfilesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Profile**](Profile.md) |  | [optional] 
**NextPageCursor** | Pointer to **string** | Cursor token required for fetching the next page. | [optional] 

## Methods

### NewListProfilesResponse

`func NewListProfilesResponse() *ListProfilesResponse`

NewListProfilesResponse instantiates a new ListProfilesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProfilesResponseWithDefaults

`func NewListProfilesResponseWithDefaults() *ListProfilesResponse`

NewListProfilesResponseWithDefaults instantiates a new ListProfilesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListProfilesResponse) GetItems() []Profile`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListProfilesResponse) GetItemsOk() (*[]Profile, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListProfilesResponse) SetItems(v []Profile)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListProfilesResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextPageCursor

`func (o *ListProfilesResponse) GetNextPageCursor() string`

GetNextPageCursor returns the NextPageCursor field if non-nil, zero value otherwise.

### GetNextPageCursorOk

`func (o *ListProfilesResponse) GetNextPageCursorOk() (*string, bool)`

GetNextPageCursorOk returns a tuple with the NextPageCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageCursor

`func (o *ListProfilesResponse) SetNextPageCursor(v string)`

SetNextPageCursor sets NextPageCursor field to given value.

### HasNextPageCursor

`func (o *ListProfilesResponse) HasNextPageCursor() bool`

HasNextPageCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


