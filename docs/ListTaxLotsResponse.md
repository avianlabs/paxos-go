# ListTaxLotsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageCursor** | Pointer to **string** | Cursor token required for fetching the next page. | [optional] 
**Items** | Pointer to [**[]TaxLot**](TaxLot.md) | The result list of tax lots. | [optional] 

## Methods

### NewListTaxLotsResponse

`func NewListTaxLotsResponse() *ListTaxLotsResponse`

NewListTaxLotsResponse instantiates a new ListTaxLotsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTaxLotsResponseWithDefaults

`func NewListTaxLotsResponseWithDefaults() *ListTaxLotsResponse`

NewListTaxLotsResponseWithDefaults instantiates a new ListTaxLotsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageCursor

`func (o *ListTaxLotsResponse) GetNextPageCursor() string`

GetNextPageCursor returns the NextPageCursor field if non-nil, zero value otherwise.

### GetNextPageCursorOk

`func (o *ListTaxLotsResponse) GetNextPageCursorOk() (*string, bool)`

GetNextPageCursorOk returns a tuple with the NextPageCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageCursor

`func (o *ListTaxLotsResponse) SetNextPageCursor(v string)`

SetNextPageCursor sets NextPageCursor field to given value.

### HasNextPageCursor

`func (o *ListTaxLotsResponse) HasNextPageCursor() bool`

HasNextPageCursor returns a boolean if a field has been set.

### GetItems

`func (o *ListTaxLotsResponse) GetItems() []TaxLot`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListTaxLotsResponse) GetItemsOk() (*[]TaxLot, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListTaxLotsResponse) SetItems(v []TaxLot)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListTaxLotsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


