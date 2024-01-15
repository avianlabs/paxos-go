# CreateSandboxDepositRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | [**Asset**](Asset.md) |  | 
**Amount** | **string** | The amount to deposit. | 
**CryptoNetwork** | Pointer to [**CryptoNetwork**](CryptoNetwork.md) |  | [optional] 

## Methods

### NewCreateSandboxDepositRequest

`func NewCreateSandboxDepositRequest(asset Asset, amount string, ) *CreateSandboxDepositRequest`

NewCreateSandboxDepositRequest instantiates a new CreateSandboxDepositRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSandboxDepositRequestWithDefaults

`func NewCreateSandboxDepositRequestWithDefaults() *CreateSandboxDepositRequest`

NewCreateSandboxDepositRequestWithDefaults instantiates a new CreateSandboxDepositRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *CreateSandboxDepositRequest) GetAsset() Asset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *CreateSandboxDepositRequest) GetAssetOk() (*Asset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *CreateSandboxDepositRequest) SetAsset(v Asset)`

SetAsset sets Asset field to given value.


### GetAmount

`func (o *CreateSandboxDepositRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateSandboxDepositRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateSandboxDepositRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCryptoNetwork

`func (o *CreateSandboxDepositRequest) GetCryptoNetwork() CryptoNetwork`

GetCryptoNetwork returns the CryptoNetwork field if non-nil, zero value otherwise.

### GetCryptoNetworkOk

`func (o *CreateSandboxDepositRequest) GetCryptoNetworkOk() (*CryptoNetwork, bool)`

GetCryptoNetworkOk returns a tuple with the CryptoNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoNetwork

`func (o *CreateSandboxDepositRequest) SetCryptoNetwork(v CryptoNetwork)`

SetCryptoNetwork sets CryptoNetwork field to given value.

### HasCryptoNetwork

`func (o *CreateSandboxDepositRequest) HasCryptoNetwork() bool`

HasCryptoNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


