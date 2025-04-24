# ExchangePublicCreateOrderBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefId** | Pointer to **string** | The idempotence ID for order creation. Can be reused if the order has been closed for more than 24 hours. | [optional] 
**Side** | [**OrderSide**](OrderSide.md) |  | 
**Market** | [**Market**](Market.md) |  | 
**Type** | [**OrderType**](OrderType.md) |  | 
**BaseAmount** | Pointer to **string** | The base currency amount for any limit order or the exact amount to sell for a market sell order. | [optional] 
**Price** | Pointer to **string** | The quote price. | [optional] 
**QuoteAmount** | Pointer to **string** | The quote currency amount of purchase for a market buy order. | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata to store on the quote and created order. Up to 6 key/value pairs may be stored, with each key and value at most 100 characters. | [optional] 
**AwaitFillMillis** | Pointer to **int64** | The amount of time to wait for the order to fill, in milliseconds. When &#x60;await_fill_millis&#x60; is set to a non-zero value, the Create Order call does not return immediately on order creation. Instead, the call blocks until either:   1. The order has filled completely   2. The time &#x60;await_fill_millis&#x60; has elapsed The maximum wait timeout is 10 seconds (10000 milliseconds). | [optional] 
**TimeInForce** | Pointer to [**TimeInForce**](TimeInForce.md) |  | [optional] 
**ExpirationDate** | Pointer to **string** | The date the order will expire if not completed when specified time in force is GTT. Format is a unix timestamp in milliseconds (13-digits) UTC (total milliseconds that have elapsed since January 1st, 1970 UTC). | [optional] 
**IdentityId** | Pointer to **string** | The end user that requests the trade. This field must be used in conjunction with &#x60;identity_account_id&#x60;, otherwise the order is rejected. Depending on your integration type, &#x60;identity_id&#x60; and &#x60;identity_account_id&#x60; may be required. | [optional] 
**IdentityAccountId** | Pointer to **string** | The account under which this order is placed. The provided identity must be allowed to trade on behalf of this account. This field must be used in conjunction with &#x60;identity_id&#x60;, otherwise the order is rejected. Depending on your integration type, &#x60;identity_account_id&#x60; and &#x60;identity_id&#x60; may be required. | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**RecipientProfileId** | Pointer to **string** | The profileId that will receive settled currency (base for buy orders, quote for sell orders). | [optional] 
**SelfMatchPreventionId** | Pointer to **string** | The string field used to prevent matching against an opposite side order submitted by the same Crypto Brokerage customer. If this field is not submitted, an order that matches against another order submitted by the same customer will cancel the original resting order. Up to 36 characters are supported. This field requires additional permissions only available to certain accounts. Reach out to your Paxos Representative for more information. | [optional] 

## Methods

### NewExchangePublicCreateOrderBody

`func NewExchangePublicCreateOrderBody(side OrderSide, market Market, type_ OrderType, ) *ExchangePublicCreateOrderBody`

NewExchangePublicCreateOrderBody instantiates a new ExchangePublicCreateOrderBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangePublicCreateOrderBodyWithDefaults

`func NewExchangePublicCreateOrderBodyWithDefaults() *ExchangePublicCreateOrderBody`

NewExchangePublicCreateOrderBodyWithDefaults instantiates a new ExchangePublicCreateOrderBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefId

`func (o *ExchangePublicCreateOrderBody) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ExchangePublicCreateOrderBody) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ExchangePublicCreateOrderBody) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ExchangePublicCreateOrderBody) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetSide

`func (o *ExchangePublicCreateOrderBody) GetSide() OrderSide`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *ExchangePublicCreateOrderBody) GetSideOk() (*OrderSide, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *ExchangePublicCreateOrderBody) SetSide(v OrderSide)`

SetSide sets Side field to given value.


### GetMarket

`func (o *ExchangePublicCreateOrderBody) GetMarket() Market`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *ExchangePublicCreateOrderBody) GetMarketOk() (*Market, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *ExchangePublicCreateOrderBody) SetMarket(v Market)`

SetMarket sets Market field to given value.


### GetType

`func (o *ExchangePublicCreateOrderBody) GetType() OrderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExchangePublicCreateOrderBody) GetTypeOk() (*OrderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExchangePublicCreateOrderBody) SetType(v OrderType)`

SetType sets Type field to given value.


### GetBaseAmount

`func (o *ExchangePublicCreateOrderBody) GetBaseAmount() string`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *ExchangePublicCreateOrderBody) GetBaseAmountOk() (*string, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *ExchangePublicCreateOrderBody) SetBaseAmount(v string)`

SetBaseAmount sets BaseAmount field to given value.

### HasBaseAmount

`func (o *ExchangePublicCreateOrderBody) HasBaseAmount() bool`

HasBaseAmount returns a boolean if a field has been set.

### GetPrice

`func (o *ExchangePublicCreateOrderBody) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ExchangePublicCreateOrderBody) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ExchangePublicCreateOrderBody) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ExchangePublicCreateOrderBody) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuoteAmount

`func (o *ExchangePublicCreateOrderBody) GetQuoteAmount() string`

GetQuoteAmount returns the QuoteAmount field if non-nil, zero value otherwise.

### GetQuoteAmountOk

`func (o *ExchangePublicCreateOrderBody) GetQuoteAmountOk() (*string, bool)`

GetQuoteAmountOk returns a tuple with the QuoteAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAmount

`func (o *ExchangePublicCreateOrderBody) SetQuoteAmount(v string)`

SetQuoteAmount sets QuoteAmount field to given value.

### HasQuoteAmount

`func (o *ExchangePublicCreateOrderBody) HasQuoteAmount() bool`

HasQuoteAmount returns a boolean if a field has been set.

### GetMetadata

`func (o *ExchangePublicCreateOrderBody) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ExchangePublicCreateOrderBody) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ExchangePublicCreateOrderBody) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ExchangePublicCreateOrderBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAwaitFillMillis

`func (o *ExchangePublicCreateOrderBody) GetAwaitFillMillis() int64`

GetAwaitFillMillis returns the AwaitFillMillis field if non-nil, zero value otherwise.

### GetAwaitFillMillisOk

`func (o *ExchangePublicCreateOrderBody) GetAwaitFillMillisOk() (*int64, bool)`

GetAwaitFillMillisOk returns a tuple with the AwaitFillMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaitFillMillis

`func (o *ExchangePublicCreateOrderBody) SetAwaitFillMillis(v int64)`

SetAwaitFillMillis sets AwaitFillMillis field to given value.

### HasAwaitFillMillis

`func (o *ExchangePublicCreateOrderBody) HasAwaitFillMillis() bool`

HasAwaitFillMillis returns a boolean if a field has been set.

### GetTimeInForce

`func (o *ExchangePublicCreateOrderBody) GetTimeInForce() TimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *ExchangePublicCreateOrderBody) GetTimeInForceOk() (*TimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *ExchangePublicCreateOrderBody) SetTimeInForce(v TimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *ExchangePublicCreateOrderBody) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ExchangePublicCreateOrderBody) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ExchangePublicCreateOrderBody) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ExchangePublicCreateOrderBody) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ExchangePublicCreateOrderBody) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetIdentityId

`func (o *ExchangePublicCreateOrderBody) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *ExchangePublicCreateOrderBody) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *ExchangePublicCreateOrderBody) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *ExchangePublicCreateOrderBody) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetIdentityAccountId

`func (o *ExchangePublicCreateOrderBody) GetIdentityAccountId() string`

GetIdentityAccountId returns the IdentityAccountId field if non-nil, zero value otherwise.

### GetIdentityAccountIdOk

`func (o *ExchangePublicCreateOrderBody) GetIdentityAccountIdOk() (*string, bool)`

GetIdentityAccountIdOk returns a tuple with the IdentityAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAccountId

`func (o *ExchangePublicCreateOrderBody) SetIdentityAccountId(v string)`

SetIdentityAccountId sets IdentityAccountId field to given value.

### HasIdentityAccountId

`func (o *ExchangePublicCreateOrderBody) HasIdentityAccountId() bool`

HasIdentityAccountId returns a boolean if a field has been set.

### GetStopPrice

`func (o *ExchangePublicCreateOrderBody) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *ExchangePublicCreateOrderBody) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *ExchangePublicCreateOrderBody) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *ExchangePublicCreateOrderBody) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetRecipientProfileId

`func (o *ExchangePublicCreateOrderBody) GetRecipientProfileId() string`

GetRecipientProfileId returns the RecipientProfileId field if non-nil, zero value otherwise.

### GetRecipientProfileIdOk

`func (o *ExchangePublicCreateOrderBody) GetRecipientProfileIdOk() (*string, bool)`

GetRecipientProfileIdOk returns a tuple with the RecipientProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientProfileId

`func (o *ExchangePublicCreateOrderBody) SetRecipientProfileId(v string)`

SetRecipientProfileId sets RecipientProfileId field to given value.

### HasRecipientProfileId

`func (o *ExchangePublicCreateOrderBody) HasRecipientProfileId() bool`

HasRecipientProfileId returns a boolean if a field has been set.

### GetSelfMatchPreventionId

`func (o *ExchangePublicCreateOrderBody) GetSelfMatchPreventionId() string`

GetSelfMatchPreventionId returns the SelfMatchPreventionId field if non-nil, zero value otherwise.

### GetSelfMatchPreventionIdOk

`func (o *ExchangePublicCreateOrderBody) GetSelfMatchPreventionIdOk() (*string, bool)`

GetSelfMatchPreventionIdOk returns a tuple with the SelfMatchPreventionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfMatchPreventionId

`func (o *ExchangePublicCreateOrderBody) SetSelfMatchPreventionId(v string)`

SetSelfMatchPreventionId sets SelfMatchPreventionId field to given value.

### HasSelfMatchPreventionId

`func (o *ExchangePublicCreateOrderBody) HasSelfMatchPreventionId() bool`

HasSelfMatchPreventionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


