# GetOrderBookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asks** | Pointer to [**[]BookLevel**](BookLevel.md) | All Asks. | [optional] 
**Bids** | Pointer to [**[]BookLevel**](BookLevel.md) | All Bids. | [optional] 
**Market** | Pointer to [**Market**](Market.md) |  | [optional] 

## Methods

### NewGetOrderBookResponse

`func NewGetOrderBookResponse() *GetOrderBookResponse`

NewGetOrderBookResponse instantiates a new GetOrderBookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderBookResponseWithDefaults

`func NewGetOrderBookResponseWithDefaults() *GetOrderBookResponse`

NewGetOrderBookResponseWithDefaults instantiates a new GetOrderBookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsks

`func (o *GetOrderBookResponse) GetAsks() []BookLevel`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *GetOrderBookResponse) GetAsksOk() (*[]BookLevel, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *GetOrderBookResponse) SetAsks(v []BookLevel)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *GetOrderBookResponse) HasAsks() bool`

HasAsks returns a boolean if a field has been set.

### GetBids

`func (o *GetOrderBookResponse) GetBids() []BookLevel`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *GetOrderBookResponse) GetBidsOk() (*[]BookLevel, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *GetOrderBookResponse) SetBids(v []BookLevel)`

SetBids sets Bids field to given value.

### HasBids

`func (o *GetOrderBookResponse) HasBids() bool`

HasBids returns a boolean if a field has been set.

### GetMarket

`func (o *GetOrderBookResponse) GetMarket() Market`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *GetOrderBookResponse) GetMarketOk() (*Market, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *GetOrderBookResponse) SetMarket(v Market)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *GetOrderBookResponse) HasMarket() bool`

HasMarket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


