# FundingSandboxCreateSandboxDepositBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | **string** | The kind of asset to deposit. | 
**Amount** | **string** | The amount to deposit. | 
**CryptoNetwork** | Pointer to [**CryptoNetwork**](CryptoNetwork.md) |  | [optional] 
**CryptoAddress** | Pointer to **string** |  | [optional] 

## Methods

### NewFundingSandboxCreateSandboxDepositBody

`func NewFundingSandboxCreateSandboxDepositBody(asset string, amount string, ) *FundingSandboxCreateSandboxDepositBody`

NewFundingSandboxCreateSandboxDepositBody instantiates a new FundingSandboxCreateSandboxDepositBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundingSandboxCreateSandboxDepositBodyWithDefaults

`func NewFundingSandboxCreateSandboxDepositBodyWithDefaults() *FundingSandboxCreateSandboxDepositBody`

NewFundingSandboxCreateSandboxDepositBodyWithDefaults instantiates a new FundingSandboxCreateSandboxDepositBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *FundingSandboxCreateSandboxDepositBody) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *FundingSandboxCreateSandboxDepositBody) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *FundingSandboxCreateSandboxDepositBody) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetAmount

`func (o *FundingSandboxCreateSandboxDepositBody) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FundingSandboxCreateSandboxDepositBody) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FundingSandboxCreateSandboxDepositBody) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCryptoNetwork

`func (o *FundingSandboxCreateSandboxDepositBody) GetCryptoNetwork() CryptoNetwork`

GetCryptoNetwork returns the CryptoNetwork field if non-nil, zero value otherwise.

### GetCryptoNetworkOk

`func (o *FundingSandboxCreateSandboxDepositBody) GetCryptoNetworkOk() (*CryptoNetwork, bool)`

GetCryptoNetworkOk returns a tuple with the CryptoNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoNetwork

`func (o *FundingSandboxCreateSandboxDepositBody) SetCryptoNetwork(v CryptoNetwork)`

SetCryptoNetwork sets CryptoNetwork field to given value.

### HasCryptoNetwork

`func (o *FundingSandboxCreateSandboxDepositBody) HasCryptoNetwork() bool`

HasCryptoNetwork returns a boolean if a field has been set.

### GetCryptoAddress

`func (o *FundingSandboxCreateSandboxDepositBody) GetCryptoAddress() string`

GetCryptoAddress returns the CryptoAddress field if non-nil, zero value otherwise.

### GetCryptoAddressOk

`func (o *FundingSandboxCreateSandboxDepositBody) GetCryptoAddressOk() (*string, bool)`

GetCryptoAddressOk returns a tuple with the CryptoAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAddress

`func (o *FundingSandboxCreateSandboxDepositBody) SetCryptoAddress(v string)`

SetCryptoAddress sets CryptoAddress field to given value.

### HasCryptoAddress

`func (o *FundingSandboxCreateSandboxDepositBody) HasCryptoAddress() bool`

HasCryptoAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


