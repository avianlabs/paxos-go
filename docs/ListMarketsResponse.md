# ListMarketsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Markets** | Pointer to [**[]MarketDetails**](MarketDetails.md) |  | [optional] 

## Methods

### NewListMarketsResponse

`func NewListMarketsResponse() *ListMarketsResponse`

NewListMarketsResponse instantiates a new ListMarketsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMarketsResponseWithDefaults

`func NewListMarketsResponseWithDefaults() *ListMarketsResponse`

NewListMarketsResponseWithDefaults instantiates a new ListMarketsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarkets

`func (o *ListMarketsResponse) GetMarkets() []MarketDetails`

GetMarkets returns the Markets field if non-nil, zero value otherwise.

### GetMarketsOk

`func (o *ListMarketsResponse) GetMarketsOk() (*[]MarketDetails, bool)`

GetMarketsOk returns a tuple with the Markets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkets

`func (o *ListMarketsResponse) SetMarkets(v []MarketDetails)`

SetMarkets sets Markets field to given value.

### HasMarkets

`func (o *ListMarketsResponse) HasMarkets() bool`

HasMarkets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


