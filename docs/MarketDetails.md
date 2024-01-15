# MarketDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**Market**](Market.md) |  | [optional] 
**BaseAsset** | Pointer to [**Asset**](Asset.md) |  | [optional] 
**QuoteAsset** | Pointer to [**Asset**](Asset.md) |  | [optional] 
**TickRate** | Pointer to **string** |  | [optional] 

## Methods

### NewMarketDetails

`func NewMarketDetails() *MarketDetails`

NewMarketDetails instantiates a new MarketDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketDetailsWithDefaults

`func NewMarketDetailsWithDefaults() *MarketDetails`

NewMarketDetailsWithDefaults instantiates a new MarketDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *MarketDetails) GetMarket() Market`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *MarketDetails) GetMarketOk() (*Market, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *MarketDetails) SetMarket(v Market)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *MarketDetails) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetBaseAsset

`func (o *MarketDetails) GetBaseAsset() Asset`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *MarketDetails) GetBaseAssetOk() (*Asset, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *MarketDetails) SetBaseAsset(v Asset)`

SetBaseAsset sets BaseAsset field to given value.

### HasBaseAsset

`func (o *MarketDetails) HasBaseAsset() bool`

HasBaseAsset returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *MarketDetails) GetQuoteAsset() Asset`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *MarketDetails) GetQuoteAssetOk() (*Asset, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *MarketDetails) SetQuoteAsset(v Asset)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *MarketDetails) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetTickRate

`func (o *MarketDetails) GetTickRate() string`

GetTickRate returns the TickRate field if non-nil, zero value otherwise.

### GetTickRateOk

`func (o *MarketDetails) GetTickRateOk() (*string, bool)`

GetTickRateOk returns a tuple with the TickRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickRate

`func (o *MarketDetails) SetTickRate(v string)`

SetTickRate sets TickRate field to given value.

### HasTickRate

`func (o *MarketDetails) HasTickRate() bool`

HasTickRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


