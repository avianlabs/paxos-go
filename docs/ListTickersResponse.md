# ListTickersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tickers** | Pointer to [**[]TickerRecord**](TickerRecord.md) |  | [optional] 

## Methods

### NewListTickersResponse

`func NewListTickersResponse() *ListTickersResponse`

NewListTickersResponse instantiates a new ListTickersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTickersResponseWithDefaults

`func NewListTickersResponseWithDefaults() *ListTickersResponse`

NewListTickersResponseWithDefaults instantiates a new ListTickersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTickers

`func (o *ListTickersResponse) GetTickers() []TickerRecord`

GetTickers returns the Tickers field if non-nil, zero value otherwise.

### GetTickersOk

`func (o *ListTickersResponse) GetTickersOk() (*[]TickerRecord, bool)`

GetTickersOk returns a tuple with the Tickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickers

`func (o *ListTickersResponse) SetTickers(v []TickerRecord)`

SetTickers sets Tickers field to given value.

### HasTickers

`func (o *ListTickersResponse) HasTickers() bool`

HasTickers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


