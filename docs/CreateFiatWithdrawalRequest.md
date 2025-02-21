# CreateFiatWithdrawalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefId** | Pointer to **string** | The optional client-specified ID (for idempotence). | [optional] 
**Amount** | Pointer to **string** | Amount to withdraw, excluding fees. Specify exactly one of &#x60;amount&#x60; or &#x60;total&#x60;. When &#x60;amount&#x60; is specified, Paxos initiates the withdrawal for &#x60;amount&#x60; and then charges fees. | [optional] 
**Asset** | **string** | The asset to withdraw. Current supported asset: \&quot;USD\&quot;. | 
**FiatAccountId** | **string** | The fiat account (&#x60;fiat_account_id&#x60;) destination. | 
**ProfileId** | **string** | The Profile (&#x60;profile_id&#x60;) to withdraw from. | 
**IdentityId** | Pointer to **string** | The Identity (&#x60;identity_id&#x60;) of the user making the withdrawal. Required only for customers with [3rd-Party integrations](https://docs.paxos.com/crypto-brokerage/ledger-type#fiat-and-crypto-subledger) initiating transfers on behalf of their end users. | [optional] 
**AccountId** | Pointer to **string** | The Account (&#x60;account_id&#x60;) associated with the Identity of the user making the withdrawal. Required only for customers with [3rd-Party integrations](https://docs.paxos.com/crypto-brokerage/ledger-type#fiat-and-crypto-subledger) initiating transfers on behalf of their end users. | [optional] 
**Metadata** | Pointer to **map[string]string** | Optional client-specified metadata. Up to 6 key/value pairs may be provided. Each key and value must be less than or equal to 100 characters. | [optional] 
**Memo** | Pointer to **string** | Optional additional memo to be set on the outgoing wire. Only used for wire withdrawals. | [optional] 
**Total** | Pointer to **string** | Total to withdraw, including fees. Specify exactly one of &#x60;amount&#x60; or &#x60;total&#x60;. When &#x60;total&#x60; is specified, Paxos initiates the withdrawal for &#x60;total&#x60; minus the fee. | [optional] 

## Methods

### NewCreateFiatWithdrawalRequest

`func NewCreateFiatWithdrawalRequest(asset string, fiatAccountId string, profileId string, ) *CreateFiatWithdrawalRequest`

NewCreateFiatWithdrawalRequest instantiates a new CreateFiatWithdrawalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFiatWithdrawalRequestWithDefaults

`func NewCreateFiatWithdrawalRequestWithDefaults() *CreateFiatWithdrawalRequest`

NewCreateFiatWithdrawalRequestWithDefaults instantiates a new CreateFiatWithdrawalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefId

`func (o *CreateFiatWithdrawalRequest) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *CreateFiatWithdrawalRequest) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *CreateFiatWithdrawalRequest) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *CreateFiatWithdrawalRequest) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetAmount

`func (o *CreateFiatWithdrawalRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateFiatWithdrawalRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateFiatWithdrawalRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CreateFiatWithdrawalRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAsset

`func (o *CreateFiatWithdrawalRequest) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *CreateFiatWithdrawalRequest) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *CreateFiatWithdrawalRequest) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetFiatAccountId

`func (o *CreateFiatWithdrawalRequest) GetFiatAccountId() string`

GetFiatAccountId returns the FiatAccountId field if non-nil, zero value otherwise.

### GetFiatAccountIdOk

`func (o *CreateFiatWithdrawalRequest) GetFiatAccountIdOk() (*string, bool)`

GetFiatAccountIdOk returns a tuple with the FiatAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatAccountId

`func (o *CreateFiatWithdrawalRequest) SetFiatAccountId(v string)`

SetFiatAccountId sets FiatAccountId field to given value.


### GetProfileId

`func (o *CreateFiatWithdrawalRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CreateFiatWithdrawalRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CreateFiatWithdrawalRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetIdentityId

`func (o *CreateFiatWithdrawalRequest) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *CreateFiatWithdrawalRequest) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *CreateFiatWithdrawalRequest) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *CreateFiatWithdrawalRequest) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetAccountId

`func (o *CreateFiatWithdrawalRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateFiatWithdrawalRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateFiatWithdrawalRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreateFiatWithdrawalRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateFiatWithdrawalRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateFiatWithdrawalRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateFiatWithdrawalRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateFiatWithdrawalRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMemo

`func (o *CreateFiatWithdrawalRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *CreateFiatWithdrawalRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *CreateFiatWithdrawalRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *CreateFiatWithdrawalRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetTotal

`func (o *CreateFiatWithdrawalRequest) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CreateFiatWithdrawalRequest) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CreateFiatWithdrawalRequest) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CreateFiatWithdrawalRequest) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


