# CreateCryptoWithdrawalFeeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | **string** | The currency to withdraw. | 
**DestinationAddress** | **string** | The destination address. | 
**CryptoNetwork** | [**CryptoNetwork**](CryptoNetwork.md) |  | 
**Amount** | Pointer to **string** | The amount to withdraw. Must be greater than &#x60;0&#x60;. Specify exactly one of &#x60;amount&#x60; or &#x60;total&#x60;, otherwise an error is returned. | [optional] 
**Total** | Pointer to **string** | Total amount to withdraw, including fees. Must be greater than &#x60;0&#x60;. Specify exactly one of &#x60;total &#x60; or &#x60;amount&#x60;, otherwise an error is returned. | [optional] 

## Methods

### NewCreateCryptoWithdrawalFeeRequest

`func NewCreateCryptoWithdrawalFeeRequest(asset string, destinationAddress string, cryptoNetwork CryptoNetwork, ) *CreateCryptoWithdrawalFeeRequest`

NewCreateCryptoWithdrawalFeeRequest instantiates a new CreateCryptoWithdrawalFeeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCryptoWithdrawalFeeRequestWithDefaults

`func NewCreateCryptoWithdrawalFeeRequestWithDefaults() *CreateCryptoWithdrawalFeeRequest`

NewCreateCryptoWithdrawalFeeRequestWithDefaults instantiates a new CreateCryptoWithdrawalFeeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *CreateCryptoWithdrawalFeeRequest) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *CreateCryptoWithdrawalFeeRequest) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *CreateCryptoWithdrawalFeeRequest) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetDestinationAddress

`func (o *CreateCryptoWithdrawalFeeRequest) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *CreateCryptoWithdrawalFeeRequest) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *CreateCryptoWithdrawalFeeRequest) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.


### GetCryptoNetwork

`func (o *CreateCryptoWithdrawalFeeRequest) GetCryptoNetwork() CryptoNetwork`

GetCryptoNetwork returns the CryptoNetwork field if non-nil, zero value otherwise.

### GetCryptoNetworkOk

`func (o *CreateCryptoWithdrawalFeeRequest) GetCryptoNetworkOk() (*CryptoNetwork, bool)`

GetCryptoNetworkOk returns a tuple with the CryptoNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoNetwork

`func (o *CreateCryptoWithdrawalFeeRequest) SetCryptoNetwork(v CryptoNetwork)`

SetCryptoNetwork sets CryptoNetwork field to given value.


### GetAmount

`func (o *CreateCryptoWithdrawalFeeRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateCryptoWithdrawalFeeRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateCryptoWithdrawalFeeRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CreateCryptoWithdrawalFeeRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTotal

`func (o *CreateCryptoWithdrawalFeeRequest) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CreateCryptoWithdrawalFeeRequest) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CreateCryptoWithdrawalFeeRequest) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CreateCryptoWithdrawalFeeRequest) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


