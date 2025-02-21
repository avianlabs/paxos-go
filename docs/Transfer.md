# Transfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Paxos transfer ID. | 
**CustomerId** | **string** | The Paxos customer ID. | 
**ProfileId** | **string** | The target Profile of the transfer. The profile asset balance is debited or credited by the transfer. | 
**IdentityId** | Pointer to **string** | The Paxos ID of the Identity associated with the transfer. | [optional] 
**RefId** | Pointer to **string** | The optional client-specified ID for replay protection and lookup. | [optional] 
**Amount** | **string** | The amount sent in the transfer. | 
**Total** | **string** | The balance change from this transfer: amount - fee for deposits, and amount + fee for withdrawals. Unsigned. | 
**Fee** | **string** | The fee paid for the transfer. | 
**Asset** | **string** | The asset for this transfer. This profile&#39;s balance of this asset will be debited or credited. | 
**BalanceAsset** | Pointer to **string** | The balance_asset represents what asset&#39;s balance was affected at Paxos with this transfer. It only differs from Asset when the transfer includes conversion. | [optional] 
**Direction** | [**TransferDirection**](TransferDirection.md) |  | 
**Type** | [**TransferType**](TransferType.md) |  | 
**Status** | [**TransferStatus**](TransferStatus.md) |  | 
**CreatedAt** | **time.Time** | The time at which this transfer record was created. | 
**UpdatedAt** | **time.Time** | The time at which this transfer record was most recently updated. | 
**Metadata** | Pointer to **map[string]string** | Optional client-specified stored metadata. For deposit event transfers this metadata is copied from the crypto deposit address or fiat deposit memo used for attribution. Up to 6 key/value pairs may be returned. Each key and value must be less than or equal to 100 characters. | [optional] 
**DestinationAddress** | Pointer to **string** | The destination crypto address. | [optional] 
**CryptoNetwork** | Pointer to [**CryptoNetwork**](CryptoNetwork.md) |  | [optional] 
**CryptoTxHash** | Pointer to **string** | For crypto transactions, the on-chain transaction hash. | [optional] 
**CryptoTxIndex** | Pointer to **string** | For crypto transactions, the output index or output address. | [optional] 
**AccountId** | Pointer to **string** | The Paxos ID of the Account associated with the transfer. | [optional] 
**AutoConversion** | Pointer to [**AutoConversion**](AutoConversion.md) |  | [optional] 
**GroupId** | Pointer to **string** | Unique identifier linking the debit and credit sides of an internal or Paxos transfer. | [optional] 
**FiatAccountId** | Pointer to **string** | For fiat withdrawals, the Paxos ID of the owner&#39;s fiat account (UUID). | [optional] 
**SecondaryStatus** | Pointer to [**SecondaryStatus**](SecondaryStatus.md) |  | [optional] 

## Methods

### NewTransfer

`func NewTransfer(id string, customerId string, profileId string, amount string, total string, fee string, asset string, direction TransferDirection, type_ TransferType, status TransferStatus, createdAt time.Time, updatedAt time.Time, ) *Transfer`

NewTransfer instantiates a new Transfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferWithDefaults

`func NewTransferWithDefaults() *Transfer`

NewTransferWithDefaults instantiates a new Transfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transfer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transfer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transfer) SetId(v string)`

SetId sets Id field to given value.


### GetCustomerId

`func (o *Transfer) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Transfer) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Transfer) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetProfileId

`func (o *Transfer) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *Transfer) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *Transfer) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetIdentityId

`func (o *Transfer) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *Transfer) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *Transfer) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *Transfer) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetRefId

`func (o *Transfer) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *Transfer) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *Transfer) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *Transfer) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetAmount

`func (o *Transfer) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Transfer) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Transfer) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetTotal

`func (o *Transfer) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Transfer) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Transfer) SetTotal(v string)`

SetTotal sets Total field to given value.


### GetFee

`func (o *Transfer) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *Transfer) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *Transfer) SetFee(v string)`

SetFee sets Fee field to given value.


### GetAsset

`func (o *Transfer) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *Transfer) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *Transfer) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetBalanceAsset

`func (o *Transfer) GetBalanceAsset() string`

GetBalanceAsset returns the BalanceAsset field if non-nil, zero value otherwise.

### GetBalanceAssetOk

`func (o *Transfer) GetBalanceAssetOk() (*string, bool)`

GetBalanceAssetOk returns a tuple with the BalanceAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAsset

`func (o *Transfer) SetBalanceAsset(v string)`

SetBalanceAsset sets BalanceAsset field to given value.

### HasBalanceAsset

`func (o *Transfer) HasBalanceAsset() bool`

HasBalanceAsset returns a boolean if a field has been set.

### GetDirection

`func (o *Transfer) GetDirection() TransferDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Transfer) GetDirectionOk() (*TransferDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Transfer) SetDirection(v TransferDirection)`

SetDirection sets Direction field to given value.


### GetType

`func (o *Transfer) GetType() TransferType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transfer) GetTypeOk() (*TransferType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transfer) SetType(v TransferType)`

SetType sets Type field to given value.


### GetStatus

`func (o *Transfer) GetStatus() TransferStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transfer) GetStatusOk() (*TransferStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transfer) SetStatus(v TransferStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *Transfer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Transfer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Transfer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Transfer) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Transfer) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Transfer) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetMetadata

`func (o *Transfer) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Transfer) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Transfer) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Transfer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *Transfer) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *Transfer) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *Transfer) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *Transfer) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetCryptoNetwork

`func (o *Transfer) GetCryptoNetwork() CryptoNetwork`

GetCryptoNetwork returns the CryptoNetwork field if non-nil, zero value otherwise.

### GetCryptoNetworkOk

`func (o *Transfer) GetCryptoNetworkOk() (*CryptoNetwork, bool)`

GetCryptoNetworkOk returns a tuple with the CryptoNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoNetwork

`func (o *Transfer) SetCryptoNetwork(v CryptoNetwork)`

SetCryptoNetwork sets CryptoNetwork field to given value.

### HasCryptoNetwork

`func (o *Transfer) HasCryptoNetwork() bool`

HasCryptoNetwork returns a boolean if a field has been set.

### GetCryptoTxHash

`func (o *Transfer) GetCryptoTxHash() string`

GetCryptoTxHash returns the CryptoTxHash field if non-nil, zero value otherwise.

### GetCryptoTxHashOk

`func (o *Transfer) GetCryptoTxHashOk() (*string, bool)`

GetCryptoTxHashOk returns a tuple with the CryptoTxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoTxHash

`func (o *Transfer) SetCryptoTxHash(v string)`

SetCryptoTxHash sets CryptoTxHash field to given value.

### HasCryptoTxHash

`func (o *Transfer) HasCryptoTxHash() bool`

HasCryptoTxHash returns a boolean if a field has been set.

### GetCryptoTxIndex

`func (o *Transfer) GetCryptoTxIndex() string`

GetCryptoTxIndex returns the CryptoTxIndex field if non-nil, zero value otherwise.

### GetCryptoTxIndexOk

`func (o *Transfer) GetCryptoTxIndexOk() (*string, bool)`

GetCryptoTxIndexOk returns a tuple with the CryptoTxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoTxIndex

`func (o *Transfer) SetCryptoTxIndex(v string)`

SetCryptoTxIndex sets CryptoTxIndex field to given value.

### HasCryptoTxIndex

`func (o *Transfer) HasCryptoTxIndex() bool`

HasCryptoTxIndex returns a boolean if a field has been set.

### GetAccountId

`func (o *Transfer) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Transfer) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Transfer) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Transfer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAutoConversion

`func (o *Transfer) GetAutoConversion() AutoConversion`

GetAutoConversion returns the AutoConversion field if non-nil, zero value otherwise.

### GetAutoConversionOk

`func (o *Transfer) GetAutoConversionOk() (*AutoConversion, bool)`

GetAutoConversionOk returns a tuple with the AutoConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConversion

`func (o *Transfer) SetAutoConversion(v AutoConversion)`

SetAutoConversion sets AutoConversion field to given value.

### HasAutoConversion

`func (o *Transfer) HasAutoConversion() bool`

HasAutoConversion returns a boolean if a field has been set.

### GetGroupId

`func (o *Transfer) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Transfer) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Transfer) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Transfer) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetFiatAccountId

`func (o *Transfer) GetFiatAccountId() string`

GetFiatAccountId returns the FiatAccountId field if non-nil, zero value otherwise.

### GetFiatAccountIdOk

`func (o *Transfer) GetFiatAccountIdOk() (*string, bool)`

GetFiatAccountIdOk returns a tuple with the FiatAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatAccountId

`func (o *Transfer) SetFiatAccountId(v string)`

SetFiatAccountId sets FiatAccountId field to given value.

### HasFiatAccountId

`func (o *Transfer) HasFiatAccountId() bool`

HasFiatAccountId returns a boolean if a field has been set.

### GetSecondaryStatus

`func (o *Transfer) GetSecondaryStatus() SecondaryStatus`

GetSecondaryStatus returns the SecondaryStatus field if non-nil, zero value otherwise.

### GetSecondaryStatusOk

`func (o *Transfer) GetSecondaryStatusOk() (*SecondaryStatus, bool)`

GetSecondaryStatusOk returns a tuple with the SecondaryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryStatus

`func (o *Transfer) SetSecondaryStatus(v SecondaryStatus)`

SetSecondaryStatus sets SecondaryStatus field to given value.

### HasSecondaryStatus

`func (o *Transfer) HasSecondaryStatus() bool`

HasSecondaryStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


