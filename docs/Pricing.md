# Pricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**PricePriceMarket**](PricePriceMarket.md) |  | [optional] 
**CurrentPrice** | Pointer to **string** | The current price for the given market. | [optional] 
**YesterdayPrice** | Pointer to **string** | The one-minute average price from 24 hours ago. Updates once per minute. | [optional] 
**SnapshotAt** | Pointer to **time.Time** | The time when the price was generated. RFC3339 format, like &#x60;2023-01-03T18:27:40.294528Z&#x60;. | [optional] 

## Methods

### NewPricing

`func NewPricing() *Pricing`

NewPricing instantiates a new Pricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingWithDefaults

`func NewPricingWithDefaults() *Pricing`

NewPricingWithDefaults instantiates a new Pricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *Pricing) GetMarket() PricePriceMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *Pricing) GetMarketOk() (*PricePriceMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *Pricing) SetMarket(v PricePriceMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *Pricing) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetCurrentPrice

`func (o *Pricing) GetCurrentPrice() string`

GetCurrentPrice returns the CurrentPrice field if non-nil, zero value otherwise.

### GetCurrentPriceOk

`func (o *Pricing) GetCurrentPriceOk() (*string, bool)`

GetCurrentPriceOk returns a tuple with the CurrentPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPrice

`func (o *Pricing) SetCurrentPrice(v string)`

SetCurrentPrice sets CurrentPrice field to given value.

### HasCurrentPrice

`func (o *Pricing) HasCurrentPrice() bool`

HasCurrentPrice returns a boolean if a field has been set.

### GetYesterdayPrice

`func (o *Pricing) GetYesterdayPrice() string`

GetYesterdayPrice returns the YesterdayPrice field if non-nil, zero value otherwise.

### GetYesterdayPriceOk

`func (o *Pricing) GetYesterdayPriceOk() (*string, bool)`

GetYesterdayPriceOk returns a tuple with the YesterdayPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesterdayPrice

`func (o *Pricing) SetYesterdayPrice(v string)`

SetYesterdayPrice sets YesterdayPrice field to given value.

### HasYesterdayPrice

`func (o *Pricing) HasYesterdayPrice() bool`

HasYesterdayPrice returns a boolean if a field has been set.

### GetSnapshotAt

`func (o *Pricing) GetSnapshotAt() time.Time`

GetSnapshotAt returns the SnapshotAt field if non-nil, zero value otherwise.

### GetSnapshotAtOk

`func (o *Pricing) GetSnapshotAtOk() (*time.Time, bool)`

GetSnapshotAtOk returns a tuple with the SnapshotAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotAt

`func (o *Pricing) SetSnapshotAt(v time.Time)`

SetSnapshotAt sets SnapshotAt field to given value.

### HasSnapshotAt

`func (o *Pricing) HasSnapshotAt() bool`

HasSnapshotAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


