# FiatAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Paxos FiatAccount ID (UUID). | [optional] 
**RefId** | Pointer to **string** | The optional client-specified ID (for idempotence). | [optional] 
**IdentityId** | Pointer to **string** | The Paxos Identity (&#x60;identity_id&#x60;) of the user&#39;s FiatAccount. | [optional] 
**AccountId** | Pointer to **string** | The Paxos Account (&#x60;account_id&#x60;) of the user&#39;s FiatAccount. Required only for customers with [3rd-Party integrations](https://docs.paxos.com/crypto-brokerage/ledger-type#fiat-and-crypto-subledger) initiating transfers on behalf of their end users. | [optional] 
**FiatAccountOwner** | Pointer to [**FiatAccountOwner**](FiatAccountOwner.md) |  | [optional] 
**FiatNetworkInstructions** | Pointer to [**FiatNetworkInstructions**](FiatNetworkInstructions.md) |  | [optional] 
**Status** | Pointer to [**FiatAccountStatus**](FiatAccountStatus.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Optional client-specified metadata. Up to 6 key/value pairs may be returned. Each key and value must be less than or equal to 100 characters. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time at which this FiatAccount was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time at which this FiatAccount record was most recently updated. | [optional] 

## Methods

### NewFiatAccount

`func NewFiatAccount() *FiatAccount`

NewFiatAccount instantiates a new FiatAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiatAccountWithDefaults

`func NewFiatAccountWithDefaults() *FiatAccount`

NewFiatAccountWithDefaults instantiates a new FiatAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FiatAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FiatAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FiatAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FiatAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRefId

`func (o *FiatAccount) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *FiatAccount) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *FiatAccount) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *FiatAccount) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetIdentityId

`func (o *FiatAccount) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *FiatAccount) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *FiatAccount) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *FiatAccount) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetAccountId

`func (o *FiatAccount) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *FiatAccount) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *FiatAccount) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *FiatAccount) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetFiatAccountOwner

`func (o *FiatAccount) GetFiatAccountOwner() FiatAccountOwner`

GetFiatAccountOwner returns the FiatAccountOwner field if non-nil, zero value otherwise.

### GetFiatAccountOwnerOk

`func (o *FiatAccount) GetFiatAccountOwnerOk() (*FiatAccountOwner, bool)`

GetFiatAccountOwnerOk returns a tuple with the FiatAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatAccountOwner

`func (o *FiatAccount) SetFiatAccountOwner(v FiatAccountOwner)`

SetFiatAccountOwner sets FiatAccountOwner field to given value.

### HasFiatAccountOwner

`func (o *FiatAccount) HasFiatAccountOwner() bool`

HasFiatAccountOwner returns a boolean if a field has been set.

### GetFiatNetworkInstructions

`func (o *FiatAccount) GetFiatNetworkInstructions() FiatNetworkInstructions`

GetFiatNetworkInstructions returns the FiatNetworkInstructions field if non-nil, zero value otherwise.

### GetFiatNetworkInstructionsOk

`func (o *FiatAccount) GetFiatNetworkInstructionsOk() (*FiatNetworkInstructions, bool)`

GetFiatNetworkInstructionsOk returns a tuple with the FiatNetworkInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatNetworkInstructions

`func (o *FiatAccount) SetFiatNetworkInstructions(v FiatNetworkInstructions)`

SetFiatNetworkInstructions sets FiatNetworkInstructions field to given value.

### HasFiatNetworkInstructions

`func (o *FiatAccount) HasFiatNetworkInstructions() bool`

HasFiatNetworkInstructions returns a boolean if a field has been set.

### GetStatus

`func (o *FiatAccount) GetStatus() FiatAccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FiatAccount) GetStatusOk() (*FiatAccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FiatAccount) SetStatus(v FiatAccountStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FiatAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMetadata

`func (o *FiatAccount) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FiatAccount) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FiatAccount) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FiatAccount) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FiatAccount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FiatAccount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FiatAccount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FiatAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FiatAccount) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FiatAccount) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FiatAccount) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FiatAccount) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


