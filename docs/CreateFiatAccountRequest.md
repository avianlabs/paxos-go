# CreateFiatAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefId** | Pointer to **string** | The optional client-specified ID (for idempotence). | [optional] 
**IdentityId** | **string** | The Paxos Identity (&#x60;identity_id&#x60;) of the user&#39;s FiatAccount. | 
**AccountId** | Pointer to **string** | The Paxos Account (&#x60;account_id&#x60;) of the user&#39;s FiatAccount. Required only for customers with [3rd-Party integrations](https://docs.paxos.com/crypto-brokerage/ledger-type#fiat-and-crypto-subledger) making deposits on behalf of their end users. | [optional] 
**FiatAccountOwner** | [**FiatAccountOwner**](FiatAccountOwner.md) |  | 
**FiatNetworkInstructions** | [**FiatNetworkInstructions**](FiatNetworkInstructions.md) |  | 
**Metadata** | Pointer to **map[string]string** | Optional client-specified metadata. Up to 6 key/value pairs may be provided. Each key and value must be less than or equal to 100 characters. | [optional] 

## Methods

### NewCreateFiatAccountRequest

`func NewCreateFiatAccountRequest(identityId string, fiatAccountOwner FiatAccountOwner, fiatNetworkInstructions FiatNetworkInstructions, ) *CreateFiatAccountRequest`

NewCreateFiatAccountRequest instantiates a new CreateFiatAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFiatAccountRequestWithDefaults

`func NewCreateFiatAccountRequestWithDefaults() *CreateFiatAccountRequest`

NewCreateFiatAccountRequestWithDefaults instantiates a new CreateFiatAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefId

`func (o *CreateFiatAccountRequest) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *CreateFiatAccountRequest) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *CreateFiatAccountRequest) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *CreateFiatAccountRequest) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetIdentityId

`func (o *CreateFiatAccountRequest) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *CreateFiatAccountRequest) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *CreateFiatAccountRequest) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetAccountId

`func (o *CreateFiatAccountRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateFiatAccountRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateFiatAccountRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreateFiatAccountRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetFiatAccountOwner

`func (o *CreateFiatAccountRequest) GetFiatAccountOwner() FiatAccountOwner`

GetFiatAccountOwner returns the FiatAccountOwner field if non-nil, zero value otherwise.

### GetFiatAccountOwnerOk

`func (o *CreateFiatAccountRequest) GetFiatAccountOwnerOk() (*FiatAccountOwner, bool)`

GetFiatAccountOwnerOk returns a tuple with the FiatAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatAccountOwner

`func (o *CreateFiatAccountRequest) SetFiatAccountOwner(v FiatAccountOwner)`

SetFiatAccountOwner sets FiatAccountOwner field to given value.


### GetFiatNetworkInstructions

`func (o *CreateFiatAccountRequest) GetFiatNetworkInstructions() FiatNetworkInstructions`

GetFiatNetworkInstructions returns the FiatNetworkInstructions field if non-nil, zero value otherwise.

### GetFiatNetworkInstructionsOk

`func (o *CreateFiatAccountRequest) GetFiatNetworkInstructionsOk() (*FiatNetworkInstructions, bool)`

GetFiatNetworkInstructionsOk returns a tuple with the FiatNetworkInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatNetworkInstructions

`func (o *CreateFiatAccountRequest) SetFiatNetworkInstructions(v FiatNetworkInstructions)`

SetFiatNetworkInstructions sets FiatNetworkInstructions field to given value.


### GetMetadata

`func (o *CreateFiatAccountRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateFiatAccountRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateFiatAccountRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateFiatAccountRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


