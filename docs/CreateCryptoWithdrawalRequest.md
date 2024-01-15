# CreateCryptoWithdrawalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefId** | Pointer to **string** | Client-specified ID for replay protection and lookup. | [optional] 
**ProfileId** | **string** | The profile from which to withdraw. | 
**IdentityId** | Pointer to **string** | The Identity of the user making the withdrawal. | [optional] 
**DestinationAddress** | **string** | The destination address. | 
**Asset** | **string** | The asset to withdraw, e.g. \&quot;BTC\&quot; , \&quot;ETH\&quot;. Always specify the &#x60;crypto_network&#x60; for all withdrawals. | 
**BalanceAsset** | Pointer to **string** | The asset&#39;s balance to debit for withdrawals of Paxos-minted USD stablecoin. | [optional] 
**Metadata** | Pointer to **map[string]string** | Optional client-specified metadata. Up to 6 key/value pairs may be provided. Each key and value must be less than or equal to 100 characters. | [optional] 
**AccountId** | Pointer to **string** | The Account associated to the identity of the user making the withdrawal. | [optional] 
**FeeId** | Pointer to **string** | Optional id of the guaranteed fee. | [optional] 
**CryptoNetwork** | [**CryptoNetwork**](CryptoNetwork.md) |  | 
**Amount** | Pointer to **string** | The amount to withdraw. Specify exactly one of &#x60;amount&#x60; or &#x60;total&#x60;, otherwise an error is returned. | [optional] 
**Total** | Pointer to **string** | Total amount to withdraw, including fees. Specify exactly one of &#x60;amount&#x60; or &#x60;total&#x60;, otherwise an error is returned. | [optional] 
**Beneficiary** | Pointer to [**Beneficiary**](Beneficiary.md) |  | [optional] 

## Methods

### NewCreateCryptoWithdrawalRequest

`func NewCreateCryptoWithdrawalRequest(profileId string, destinationAddress string, asset string, cryptoNetwork CryptoNetwork, ) *CreateCryptoWithdrawalRequest`

NewCreateCryptoWithdrawalRequest instantiates a new CreateCryptoWithdrawalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCryptoWithdrawalRequestWithDefaults

`func NewCreateCryptoWithdrawalRequestWithDefaults() *CreateCryptoWithdrawalRequest`

NewCreateCryptoWithdrawalRequestWithDefaults instantiates a new CreateCryptoWithdrawalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefId

`func (o *CreateCryptoWithdrawalRequest) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *CreateCryptoWithdrawalRequest) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *CreateCryptoWithdrawalRequest) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *CreateCryptoWithdrawalRequest) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetProfileId

`func (o *CreateCryptoWithdrawalRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CreateCryptoWithdrawalRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CreateCryptoWithdrawalRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetIdentityId

`func (o *CreateCryptoWithdrawalRequest) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *CreateCryptoWithdrawalRequest) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *CreateCryptoWithdrawalRequest) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *CreateCryptoWithdrawalRequest) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *CreateCryptoWithdrawalRequest) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *CreateCryptoWithdrawalRequest) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *CreateCryptoWithdrawalRequest) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.


### GetAsset

`func (o *CreateCryptoWithdrawalRequest) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *CreateCryptoWithdrawalRequest) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *CreateCryptoWithdrawalRequest) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetBalanceAsset

`func (o *CreateCryptoWithdrawalRequest) GetBalanceAsset() string`

GetBalanceAsset returns the BalanceAsset field if non-nil, zero value otherwise.

### GetBalanceAssetOk

`func (o *CreateCryptoWithdrawalRequest) GetBalanceAssetOk() (*string, bool)`

GetBalanceAssetOk returns a tuple with the BalanceAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAsset

`func (o *CreateCryptoWithdrawalRequest) SetBalanceAsset(v string)`

SetBalanceAsset sets BalanceAsset field to given value.

### HasBalanceAsset

`func (o *CreateCryptoWithdrawalRequest) HasBalanceAsset() bool`

HasBalanceAsset returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateCryptoWithdrawalRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateCryptoWithdrawalRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateCryptoWithdrawalRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateCryptoWithdrawalRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAccountId

`func (o *CreateCryptoWithdrawalRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateCryptoWithdrawalRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateCryptoWithdrawalRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreateCryptoWithdrawalRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetFeeId

`func (o *CreateCryptoWithdrawalRequest) GetFeeId() string`

GetFeeId returns the FeeId field if non-nil, zero value otherwise.

### GetFeeIdOk

`func (o *CreateCryptoWithdrawalRequest) GetFeeIdOk() (*string, bool)`

GetFeeIdOk returns a tuple with the FeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeId

`func (o *CreateCryptoWithdrawalRequest) SetFeeId(v string)`

SetFeeId sets FeeId field to given value.

### HasFeeId

`func (o *CreateCryptoWithdrawalRequest) HasFeeId() bool`

HasFeeId returns a boolean if a field has been set.

### GetCryptoNetwork

`func (o *CreateCryptoWithdrawalRequest) GetCryptoNetwork() CryptoNetwork`

GetCryptoNetwork returns the CryptoNetwork field if non-nil, zero value otherwise.

### GetCryptoNetworkOk

`func (o *CreateCryptoWithdrawalRequest) GetCryptoNetworkOk() (*CryptoNetwork, bool)`

GetCryptoNetworkOk returns a tuple with the CryptoNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoNetwork

`func (o *CreateCryptoWithdrawalRequest) SetCryptoNetwork(v CryptoNetwork)`

SetCryptoNetwork sets CryptoNetwork field to given value.


### GetAmount

`func (o *CreateCryptoWithdrawalRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateCryptoWithdrawalRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateCryptoWithdrawalRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CreateCryptoWithdrawalRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTotal

`func (o *CreateCryptoWithdrawalRequest) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CreateCryptoWithdrawalRequest) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CreateCryptoWithdrawalRequest) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CreateCryptoWithdrawalRequest) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetBeneficiary

`func (o *CreateCryptoWithdrawalRequest) GetBeneficiary() Beneficiary`

GetBeneficiary returns the Beneficiary field if non-nil, zero value otherwise.

### GetBeneficiaryOk

`func (o *CreateCryptoWithdrawalRequest) GetBeneficiaryOk() (*Beneficiary, bool)`

GetBeneficiaryOk returns a tuple with the Beneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiary

`func (o *CreateCryptoWithdrawalRequest) SetBeneficiary(v Beneficiary)`

SetBeneficiary sets Beneficiary field to given value.

### HasBeneficiary

`func (o *CreateCryptoWithdrawalRequest) HasBeneficiary() bool`

HasBeneficiary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


