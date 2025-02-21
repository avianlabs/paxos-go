# QuoteExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The UUID of the quote execution. | 
**ProfileId** | **string** | The ID of the profile under which this order executed. | 
**QuoteId** | **string** | The ID of the quote used to create this quote execution. | 
**RefId** | Pointer to **string** | A unique identifier for the quote execution creation (for idempotence). | [optional] 
**Status** | [**QuoteExecutionStatus**](QuoteExecutionStatus.md) |  | 
**Market** | [**Market**](Market.md) |  | 
**Side** | [**OrderSide**](OrderSide.md) |  | 
**Price** | **string** | The guaranteed price of the quote. | 
**BaseAmount** | **string** | The amount of assets (crypto) in the transaction. | 
**BaseAsset** | **string** | The \&quot;base\&quot; side of the trading pair (crypto - like BTC, ETH, PAXG). | 
**QuoteAmount** | **string** | The amount of quote currency (cash) in the transaction. | 
**QuoteAsset** | **string** | The \&quot;quote\&quot; side of the trading pair (fiat - like USD, EUR, SGD). | 
**CreatedAt** | **time.Time** | The time at which the quote execution was created. | 
**SettledAt** | Pointer to **time.Time** | The time at which the quote execution was settled, completing its lifecycle. | [optional] 
**Metadata** | Pointer to **map[string]string** | Client-specified metadata. | [optional] 
**IdentityId** | Pointer to **string** | The identity under which this quote execution is placed. | [optional] 
**AccountId** | Pointer to **string** | The account under which this quote execution is placed. | [optional] 
**RecipientProfileId** | Pointer to **string** | The ID of the profile under which to deposit the funds. | [optional] 

## Methods

### NewQuoteExecution

`func NewQuoteExecution(id string, profileId string, quoteId string, status QuoteExecutionStatus, market Market, side OrderSide, price string, baseAmount string, baseAsset string, quoteAmount string, quoteAsset string, createdAt time.Time, ) *QuoteExecution`

NewQuoteExecution instantiates a new QuoteExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteExecutionWithDefaults

`func NewQuoteExecutionWithDefaults() *QuoteExecution`

NewQuoteExecutionWithDefaults instantiates a new QuoteExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuoteExecution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuoteExecution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuoteExecution) SetId(v string)`

SetId sets Id field to given value.


### GetProfileId

`func (o *QuoteExecution) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *QuoteExecution) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *QuoteExecution) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetQuoteId

`func (o *QuoteExecution) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *QuoteExecution) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *QuoteExecution) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetRefId

`func (o *QuoteExecution) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *QuoteExecution) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *QuoteExecution) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *QuoteExecution) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetStatus

`func (o *QuoteExecution) GetStatus() QuoteExecutionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QuoteExecution) GetStatusOk() (*QuoteExecutionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QuoteExecution) SetStatus(v QuoteExecutionStatus)`

SetStatus sets Status field to given value.


### GetMarket

`func (o *QuoteExecution) GetMarket() Market`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *QuoteExecution) GetMarketOk() (*Market, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *QuoteExecution) SetMarket(v Market)`

SetMarket sets Market field to given value.


### GetSide

`func (o *QuoteExecution) GetSide() OrderSide`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *QuoteExecution) GetSideOk() (*OrderSide, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *QuoteExecution) SetSide(v OrderSide)`

SetSide sets Side field to given value.


### GetPrice

`func (o *QuoteExecution) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QuoteExecution) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QuoteExecution) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetBaseAmount

`func (o *QuoteExecution) GetBaseAmount() string`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *QuoteExecution) GetBaseAmountOk() (*string, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *QuoteExecution) SetBaseAmount(v string)`

SetBaseAmount sets BaseAmount field to given value.


### GetBaseAsset

`func (o *QuoteExecution) GetBaseAsset() string`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *QuoteExecution) GetBaseAssetOk() (*string, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *QuoteExecution) SetBaseAsset(v string)`

SetBaseAsset sets BaseAsset field to given value.


### GetQuoteAmount

`func (o *QuoteExecution) GetQuoteAmount() string`

GetQuoteAmount returns the QuoteAmount field if non-nil, zero value otherwise.

### GetQuoteAmountOk

`func (o *QuoteExecution) GetQuoteAmountOk() (*string, bool)`

GetQuoteAmountOk returns a tuple with the QuoteAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAmount

`func (o *QuoteExecution) SetQuoteAmount(v string)`

SetQuoteAmount sets QuoteAmount field to given value.


### GetQuoteAsset

`func (o *QuoteExecution) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *QuoteExecution) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *QuoteExecution) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.


### GetCreatedAt

`func (o *QuoteExecution) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QuoteExecution) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QuoteExecution) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSettledAt

`func (o *QuoteExecution) GetSettledAt() time.Time`

GetSettledAt returns the SettledAt field if non-nil, zero value otherwise.

### GetSettledAtOk

`func (o *QuoteExecution) GetSettledAtOk() (*time.Time, bool)`

GetSettledAtOk returns a tuple with the SettledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAt

`func (o *QuoteExecution) SetSettledAt(v time.Time)`

SetSettledAt sets SettledAt field to given value.

### HasSettledAt

`func (o *QuoteExecution) HasSettledAt() bool`

HasSettledAt returns a boolean if a field has been set.

### GetMetadata

`func (o *QuoteExecution) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *QuoteExecution) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *QuoteExecution) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *QuoteExecution) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetIdentityId

`func (o *QuoteExecution) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *QuoteExecution) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *QuoteExecution) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *QuoteExecution) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetAccountId

`func (o *QuoteExecution) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *QuoteExecution) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *QuoteExecution) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *QuoteExecution) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetRecipientProfileId

`func (o *QuoteExecution) GetRecipientProfileId() string`

GetRecipientProfileId returns the RecipientProfileId field if non-nil, zero value otherwise.

### GetRecipientProfileIdOk

`func (o *QuoteExecution) GetRecipientProfileIdOk() (*string, bool)`

GetRecipientProfileIdOk returns a tuple with the RecipientProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientProfileId

`func (o *QuoteExecution) SetRecipientProfileId(v string)`

SetRecipientProfileId sets RecipientProfileId field to given value.

### HasRecipientProfileId

`func (o *QuoteExecution) HasRecipientProfileId() bool`

HasRecipientProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


