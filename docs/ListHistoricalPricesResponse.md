# ListHistoricalPricesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]HistoricalPrice**](HistoricalPrice.md) | List of historical prices. | [optional] 

## Methods

### NewListHistoricalPricesResponse

`func NewListHistoricalPricesResponse() *ListHistoricalPricesResponse`

NewListHistoricalPricesResponse instantiates a new ListHistoricalPricesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHistoricalPricesResponseWithDefaults

`func NewListHistoricalPricesResponseWithDefaults() *ListHistoricalPricesResponse`

NewListHistoricalPricesResponseWithDefaults instantiates a new ListHistoricalPricesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ListHistoricalPricesResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListHistoricalPricesResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListHistoricalPricesResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListHistoricalPricesResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItems

`func (o *ListHistoricalPricesResponse) GetItems() []HistoricalPrice`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListHistoricalPricesResponse) GetItemsOk() (*[]HistoricalPrice, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListHistoricalPricesResponse) SetItems(v []HistoricalPrice)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListHistoricalPricesResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


