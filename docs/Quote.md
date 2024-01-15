# Quote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The UUID of the quote. | 
**Market** | [**Market**](Market.md) |  | 
**Side** | [**OrderSide**](OrderSide.md) |  | 
**Price** | **string** | The guaranteed price, held until expires_at. | 
**BaseAsset** | [**Asset**](Asset.md) |  | 
**QuoteAsset** | [**Asset**](Asset.md) |  | 
**CreatedAt** | **time.Time** | The time at which the quote was first offered. | 
**ExpiresAt** | **time.Time** | The time at which the quote expires. | 

## Methods

### NewQuote

`func NewQuote(id string, market Market, side OrderSide, price string, baseAsset Asset, quoteAsset Asset, createdAt time.Time, expiresAt time.Time, ) *Quote`

NewQuote instantiates a new Quote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteWithDefaults

`func NewQuoteWithDefaults() *Quote`

NewQuoteWithDefaults instantiates a new Quote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Quote) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Quote) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Quote) SetId(v string)`

SetId sets Id field to given value.


### GetMarket

`func (o *Quote) GetMarket() Market`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *Quote) GetMarketOk() (*Market, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *Quote) SetMarket(v Market)`

SetMarket sets Market field to given value.


### GetSide

`func (o *Quote) GetSide() OrderSide`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *Quote) GetSideOk() (*OrderSide, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *Quote) SetSide(v OrderSide)`

SetSide sets Side field to given value.


### GetPrice

`func (o *Quote) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Quote) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Quote) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetBaseAsset

`func (o *Quote) GetBaseAsset() Asset`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *Quote) GetBaseAssetOk() (*Asset, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *Quote) SetBaseAsset(v Asset)`

SetBaseAsset sets BaseAsset field to given value.


### GetQuoteAsset

`func (o *Quote) GetQuoteAsset() Asset`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *Quote) GetQuoteAssetOk() (*Asset, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *Quote) SetQuoteAsset(v Asset)`

SetQuoteAsset sets QuoteAsset field to given value.


### GetCreatedAt

`func (o *Quote) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Quote) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Quote) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *Quote) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Quote) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Quote) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


