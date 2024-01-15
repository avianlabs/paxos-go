# FiatDepositInstructions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Fiat Deposit Instructions ID (&#x60;id&#x60;) is provided in the response of the [Create Fiat Deposit Instructions](#operation/CreateFiatDepositInstructions). Use this ID to retrieve the instructions using [Get Fiat Deposit Instructions](#operation/GetFiatDepositInstructions) &amp; [List Fiat Deposit Instructions](#operation/ListFiatDepositInstructions). | 
**ProfileId** | **string** | The Profile (&#x60;profile_id&#x60;) to deposit to. | 
**IdentityId** | **string** | The Identity (&#x60;identity_id&#x60;) of the user making the deposit. | 
**AccountId** | Pointer to **string** | The Account (&#x60;account_id&#x60;) associated with the Identity of the user making the deposit. Required only for customers with [3rd-Party integrations](https://docs.paxos.com/crypto-brokerage/ledger-type#fiat-and-crypto-subledger) making deposits on behalf of their end users. | [optional] 
**RefId** | Pointer to **string** | The optional client-specified ID (for idempotence). | [optional] 
**Status** | [**FiatDepositInstructionsStatus**](FiatDepositInstructionsStatus.md) |  | 
**MemoId** | **string** | The string, unique to the request. To deposit funds into an account, the memo ID must be provided when initiating a &#x60;WIRE&#x60; transfer to Paxos. This is provided after creating [Fiat Deposit Instructions](#/operation/CreateFiatDepositInstructions) The &#x60;memo_id&#x60; can also be found from either [Get Fiat Deposit Instructions](#operation/GetFiatDepositInstructions) or [List Fiat Deposit Instructions](#operation/ListFiatDepositInstructions) for the corresponding ID (Fiat Deposit Instructions ID). | 
**FiatNetworkInstructions** | [**FiatNetworkInstructions**](FiatNetworkInstructions.md) |  | 
**FiatAccountOwner** | Pointer to [**FiatAccountOwner**](FiatAccountOwner.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Optional client-specified metadata. | [optional] 
**CreatedAt** | **time.Time** | The time at which these instructions were created. | 

## Methods

### NewFiatDepositInstructions

`func NewFiatDepositInstructions(id string, profileId string, identityId string, status FiatDepositInstructionsStatus, memoId string, fiatNetworkInstructions FiatNetworkInstructions, createdAt time.Time, ) *FiatDepositInstructions`

NewFiatDepositInstructions instantiates a new FiatDepositInstructions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiatDepositInstructionsWithDefaults

`func NewFiatDepositInstructionsWithDefaults() *FiatDepositInstructions`

NewFiatDepositInstructionsWithDefaults instantiates a new FiatDepositInstructions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FiatDepositInstructions) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FiatDepositInstructions) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FiatDepositInstructions) SetId(v string)`

SetId sets Id field to given value.


### GetProfileId

`func (o *FiatDepositInstructions) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *FiatDepositInstructions) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *FiatDepositInstructions) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetIdentityId

`func (o *FiatDepositInstructions) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *FiatDepositInstructions) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *FiatDepositInstructions) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetAccountId

`func (o *FiatDepositInstructions) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *FiatDepositInstructions) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *FiatDepositInstructions) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *FiatDepositInstructions) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetRefId

`func (o *FiatDepositInstructions) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *FiatDepositInstructions) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *FiatDepositInstructions) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *FiatDepositInstructions) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetStatus

`func (o *FiatDepositInstructions) GetStatus() FiatDepositInstructionsStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FiatDepositInstructions) GetStatusOk() (*FiatDepositInstructionsStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FiatDepositInstructions) SetStatus(v FiatDepositInstructionsStatus)`

SetStatus sets Status field to given value.


### GetMemoId

`func (o *FiatDepositInstructions) GetMemoId() string`

GetMemoId returns the MemoId field if non-nil, zero value otherwise.

### GetMemoIdOk

`func (o *FiatDepositInstructions) GetMemoIdOk() (*string, bool)`

GetMemoIdOk returns a tuple with the MemoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoId

`func (o *FiatDepositInstructions) SetMemoId(v string)`

SetMemoId sets MemoId field to given value.


### GetFiatNetworkInstructions

`func (o *FiatDepositInstructions) GetFiatNetworkInstructions() FiatNetworkInstructions`

GetFiatNetworkInstructions returns the FiatNetworkInstructions field if non-nil, zero value otherwise.

### GetFiatNetworkInstructionsOk

`func (o *FiatDepositInstructions) GetFiatNetworkInstructionsOk() (*FiatNetworkInstructions, bool)`

GetFiatNetworkInstructionsOk returns a tuple with the FiatNetworkInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatNetworkInstructions

`func (o *FiatDepositInstructions) SetFiatNetworkInstructions(v FiatNetworkInstructions)`

SetFiatNetworkInstructions sets FiatNetworkInstructions field to given value.


### GetFiatAccountOwner

`func (o *FiatDepositInstructions) GetFiatAccountOwner() FiatAccountOwner`

GetFiatAccountOwner returns the FiatAccountOwner field if non-nil, zero value otherwise.

### GetFiatAccountOwnerOk

`func (o *FiatDepositInstructions) GetFiatAccountOwnerOk() (*FiatAccountOwner, bool)`

GetFiatAccountOwnerOk returns a tuple with the FiatAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatAccountOwner

`func (o *FiatDepositInstructions) SetFiatAccountOwner(v FiatAccountOwner)`

SetFiatAccountOwner sets FiatAccountOwner field to given value.

### HasFiatAccountOwner

`func (o *FiatDepositInstructions) HasFiatAccountOwner() bool`

HasFiatAccountOwner returns a boolean if a field has been set.

### GetMetadata

`func (o *FiatDepositInstructions) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FiatDepositInstructions) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FiatDepositInstructions) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FiatDepositInstructions) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FiatDepositInstructions) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FiatDepositInstructions) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FiatDepositInstructions) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


