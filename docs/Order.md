# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The UUID of the order. | [optional] 
**ProfileId** | Pointer to **string** | The profile ID the order is associated with. | [optional] 
**RefId** | Pointer to **string** | The idempotence ID for order creation. Can be reused if the order has been closed for more than 24 hours. | [optional] 
**Status** | Pointer to [**OrderStatus**](OrderStatus.md) |  | [optional] 
**Side** | Pointer to [**OrderSide**](OrderSide.md) |  | [optional] 
**Market** | Pointer to [**Market**](Market.md) |  | [optional] 
**Type** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**BaseAmount** | Pointer to **string** | The base amount or purchase amount for a market sell order. | [optional] 
**Price** | Pointer to **string** | The quote price. | [optional] 
**QuoteAmount** | Pointer to **string** | The quote amount of purchase for a market buy order. | [optional] 
**Metadata** | Pointer to **map[string]string** | Client-specified metadata. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time at which the order was created. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | The time at which the order was last modified. | [optional] 
**AmountFilled** | Pointer to **string** | The amount that was filled. | [optional] 
**VolumeWeightedAveragePrice** | Pointer to **string** | The volume-weighted average price. | [optional] 
**TimeInForce** | Pointer to [**TimeInForce**](TimeInForce.md) |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** | The date the order will expire if not completed when specified time in force is GTT. | [optional] 
**IdentityId** | Pointer to **string** | The end user that requests the trade. This field must be used in conjunction with &#x60;identity_account_id&#x60;, otherwise the order is rejected. Depending on your integration type, &#x60;identity_id&#x60; and &#x60;identity_account_id&#x60; may be required. | [optional] 
**IdentityAccountId** | Pointer to **string** | The account under which this order is placed. The provided identity must be allowed to trade on behalf of this account. This field must be used in conjunction with &#x60;identity_id&#x60;, otherwise the order is rejected. Depending on your integration type, &#x60;identity_account_id&#x60; and &#x60;identity_id&#x60; may be required. | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**RecipientProfileId** | Pointer to **string** | The profileId that will receive settled currency (base for buy orders, quote for sell orders). | [optional] 
**IsTriggered** | Pointer to **bool** | Returns &#x60;true&#x60; when a stop order has been triggered. | [optional] 

## Methods

### NewOrder

`func NewOrder() *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Order) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Order) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProfileId

`func (o *Order) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *Order) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *Order) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *Order) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetRefId

`func (o *Order) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *Order) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *Order) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *Order) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetStatus

`func (o *Order) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Order) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSide

`func (o *Order) GetSide() OrderSide`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *Order) GetSideOk() (*OrderSide, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *Order) SetSide(v OrderSide)`

SetSide sets Side field to given value.

### HasSide

`func (o *Order) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetMarket

`func (o *Order) GetMarket() Market`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *Order) GetMarketOk() (*Market, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *Order) SetMarket(v Market)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *Order) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetType

`func (o *Order) GetType() OrderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Order) GetTypeOk() (*OrderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Order) SetType(v OrderType)`

SetType sets Type field to given value.

### HasType

`func (o *Order) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBaseAmount

`func (o *Order) GetBaseAmount() string`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *Order) GetBaseAmountOk() (*string, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *Order) SetBaseAmount(v string)`

SetBaseAmount sets BaseAmount field to given value.

### HasBaseAmount

`func (o *Order) HasBaseAmount() bool`

HasBaseAmount returns a boolean if a field has been set.

### GetPrice

`func (o *Order) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Order) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Order) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Order) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuoteAmount

`func (o *Order) GetQuoteAmount() string`

GetQuoteAmount returns the QuoteAmount field if non-nil, zero value otherwise.

### GetQuoteAmountOk

`func (o *Order) GetQuoteAmountOk() (*string, bool)`

GetQuoteAmountOk returns a tuple with the QuoteAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAmount

`func (o *Order) SetQuoteAmount(v string)`

SetQuoteAmount sets QuoteAmount field to given value.

### HasQuoteAmount

`func (o *Order) HasQuoteAmount() bool`

HasQuoteAmount returns a boolean if a field has been set.

### GetMetadata

`func (o *Order) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Order) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Order) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Order) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Order) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Order) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Order) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Order) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Order) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Order) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Order) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Order) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetAmountFilled

`func (o *Order) GetAmountFilled() string`

GetAmountFilled returns the AmountFilled field if non-nil, zero value otherwise.

### GetAmountFilledOk

`func (o *Order) GetAmountFilledOk() (*string, bool)`

GetAmountFilledOk returns a tuple with the AmountFilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountFilled

`func (o *Order) SetAmountFilled(v string)`

SetAmountFilled sets AmountFilled field to given value.

### HasAmountFilled

`func (o *Order) HasAmountFilled() bool`

HasAmountFilled returns a boolean if a field has been set.

### GetVolumeWeightedAveragePrice

`func (o *Order) GetVolumeWeightedAveragePrice() string`

GetVolumeWeightedAveragePrice returns the VolumeWeightedAveragePrice field if non-nil, zero value otherwise.

### GetVolumeWeightedAveragePriceOk

`func (o *Order) GetVolumeWeightedAveragePriceOk() (*string, bool)`

GetVolumeWeightedAveragePriceOk returns a tuple with the VolumeWeightedAveragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeWeightedAveragePrice

`func (o *Order) SetVolumeWeightedAveragePrice(v string)`

SetVolumeWeightedAveragePrice sets VolumeWeightedAveragePrice field to given value.

### HasVolumeWeightedAveragePrice

`func (o *Order) HasVolumeWeightedAveragePrice() bool`

HasVolumeWeightedAveragePrice returns a boolean if a field has been set.

### GetTimeInForce

`func (o *Order) GetTimeInForce() TimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *Order) GetTimeInForceOk() (*TimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *Order) SetTimeInForce(v TimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *Order) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetExpirationDate

`func (o *Order) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *Order) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *Order) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *Order) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetIdentityId

`func (o *Order) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *Order) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *Order) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *Order) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetIdentityAccountId

`func (o *Order) GetIdentityAccountId() string`

GetIdentityAccountId returns the IdentityAccountId field if non-nil, zero value otherwise.

### GetIdentityAccountIdOk

`func (o *Order) GetIdentityAccountIdOk() (*string, bool)`

GetIdentityAccountIdOk returns a tuple with the IdentityAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAccountId

`func (o *Order) SetIdentityAccountId(v string)`

SetIdentityAccountId sets IdentityAccountId field to given value.

### HasIdentityAccountId

`func (o *Order) HasIdentityAccountId() bool`

HasIdentityAccountId returns a boolean if a field has been set.

### GetStopPrice

`func (o *Order) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *Order) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *Order) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *Order) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetRecipientProfileId

`func (o *Order) GetRecipientProfileId() string`

GetRecipientProfileId returns the RecipientProfileId field if non-nil, zero value otherwise.

### GetRecipientProfileIdOk

`func (o *Order) GetRecipientProfileIdOk() (*string, bool)`

GetRecipientProfileIdOk returns a tuple with the RecipientProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientProfileId

`func (o *Order) SetRecipientProfileId(v string)`

SetRecipientProfileId sets RecipientProfileId field to given value.

### HasRecipientProfileId

`func (o *Order) HasRecipientProfileId() bool`

HasRecipientProfileId returns a boolean if a field has been set.

### GetIsTriggered

`func (o *Order) GetIsTriggered() bool`

GetIsTriggered returns the IsTriggered field if non-nil, zero value otherwise.

### GetIsTriggeredOk

`func (o *Order) GetIsTriggeredOk() (*bool, bool)`

GetIsTriggeredOk returns a tuple with the IsTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTriggered

`func (o *Order) SetIsTriggered(v bool)`

SetIsTriggered sets IsTriggered field to given value.

### HasIsTriggered

`func (o *Order) HasIsTriggered() bool`

HasIsTriggered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


