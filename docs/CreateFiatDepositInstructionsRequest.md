# CreateFiatDepositInstructionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefId** | Pointer to **string** | The optional client-specified ID (for idempotence). | [optional] 
**ProfileId** | **string** | The Profile (&#x60;profile_id&#x60;) to deposit to. | 
**IdentityId** | Pointer to **string** | The Identity (&#x60;identity_id&#x60;) of the user making the deposit. Required only for customers with [3rd-Party integrations](https://docs.paxos.com/crypto-brokerage/ledger-type#fiat-and-crypto-subledger) making deposits on behalf of their end users. | [optional] 
**AccountId** | Pointer to **string** | The Account (&#x60;account_id&#x60;) associated with the Identity of the user making the deposit. Required only for customers with [3rd-Party integrations](https://docs.paxos.com/crypto-brokerage/ledger-type#fiat-and-crypto-subledger) making deposits on behalf of their end users. | [optional] 
**FiatNetwork** | [**FiatNetwork**](FiatNetwork.md) |  | 
**RoutingNumberType** | Pointer to [**FiatWireAccountType**](FiatWireAccountType.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Optional client-specified metadata. Up to 6 key/value pairs may be provided. Each key and value must be less than or equal to 100 characters. | [optional] 

## Methods

### NewCreateFiatDepositInstructionsRequest

`func NewCreateFiatDepositInstructionsRequest(profileId string, fiatNetwork FiatNetwork, ) *CreateFiatDepositInstructionsRequest`

NewCreateFiatDepositInstructionsRequest instantiates a new CreateFiatDepositInstructionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFiatDepositInstructionsRequestWithDefaults

`func NewCreateFiatDepositInstructionsRequestWithDefaults() *CreateFiatDepositInstructionsRequest`

NewCreateFiatDepositInstructionsRequestWithDefaults instantiates a new CreateFiatDepositInstructionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefId

`func (o *CreateFiatDepositInstructionsRequest) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *CreateFiatDepositInstructionsRequest) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *CreateFiatDepositInstructionsRequest) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *CreateFiatDepositInstructionsRequest) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetProfileId

`func (o *CreateFiatDepositInstructionsRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CreateFiatDepositInstructionsRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CreateFiatDepositInstructionsRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetIdentityId

`func (o *CreateFiatDepositInstructionsRequest) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *CreateFiatDepositInstructionsRequest) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *CreateFiatDepositInstructionsRequest) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *CreateFiatDepositInstructionsRequest) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetAccountId

`func (o *CreateFiatDepositInstructionsRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateFiatDepositInstructionsRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateFiatDepositInstructionsRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreateFiatDepositInstructionsRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetFiatNetwork

`func (o *CreateFiatDepositInstructionsRequest) GetFiatNetwork() FiatNetwork`

GetFiatNetwork returns the FiatNetwork field if non-nil, zero value otherwise.

### GetFiatNetworkOk

`func (o *CreateFiatDepositInstructionsRequest) GetFiatNetworkOk() (*FiatNetwork, bool)`

GetFiatNetworkOk returns a tuple with the FiatNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatNetwork

`func (o *CreateFiatDepositInstructionsRequest) SetFiatNetwork(v FiatNetwork)`

SetFiatNetwork sets FiatNetwork field to given value.


### GetRoutingNumberType

`func (o *CreateFiatDepositInstructionsRequest) GetRoutingNumberType() FiatWireAccountType`

GetRoutingNumberType returns the RoutingNumberType field if non-nil, zero value otherwise.

### GetRoutingNumberTypeOk

`func (o *CreateFiatDepositInstructionsRequest) GetRoutingNumberTypeOk() (*FiatWireAccountType, bool)`

GetRoutingNumberTypeOk returns a tuple with the RoutingNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumberType

`func (o *CreateFiatDepositInstructionsRequest) SetRoutingNumberType(v FiatWireAccountType)`

SetRoutingNumberType sets RoutingNumberType field to given value.

### HasRoutingNumberType

`func (o *CreateFiatDepositInstructionsRequest) HasRoutingNumberType() bool`

HasRoutingNumberType returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateFiatDepositInstructionsRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateFiatDepositInstructionsRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateFiatDepositInstructionsRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateFiatDepositInstructionsRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


