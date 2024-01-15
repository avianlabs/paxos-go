# CryptoWithdrawalFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the guaranteed fee. | 
**Fee** | **string** | The guaranteed fee value, in the same currency. | 
**Asset** | **string** | The currency to withdraw. | 
**ExpiresAt** | **time.Time** | The expiration timestamp of the created fee. | 
**DestinationAddress** | **string** | The destination address. | 
**CryptoNetwork** | [**CryptoNetwork**](CryptoNetwork.md) |  | 
**Amount** | Pointer to **string** | The quoted amount to withdraw for which the fee is valid. Specify exactly one of &#x60;amount&#x60; or &#x60;total&#x60;, otherwise an error is returned. | [optional] 
**Total** | Pointer to **string** | Total amount to withdraw, including fees. Specify exactly one of &#x60;amount&#x60; or &#x60;total&#x60;, otherwise an error is returned. | [optional] 

## Methods

### NewCryptoWithdrawalFee

`func NewCryptoWithdrawalFee(id string, fee string, asset string, expiresAt time.Time, destinationAddress string, cryptoNetwork CryptoNetwork, ) *CryptoWithdrawalFee`

NewCryptoWithdrawalFee instantiates a new CryptoWithdrawalFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoWithdrawalFeeWithDefaults

`func NewCryptoWithdrawalFeeWithDefaults() *CryptoWithdrawalFee`

NewCryptoWithdrawalFeeWithDefaults instantiates a new CryptoWithdrawalFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CryptoWithdrawalFee) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CryptoWithdrawalFee) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CryptoWithdrawalFee) SetId(v string)`

SetId sets Id field to given value.


### GetFee

`func (o *CryptoWithdrawalFee) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CryptoWithdrawalFee) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CryptoWithdrawalFee) SetFee(v string)`

SetFee sets Fee field to given value.


### GetAsset

`func (o *CryptoWithdrawalFee) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *CryptoWithdrawalFee) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *CryptoWithdrawalFee) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetExpiresAt

`func (o *CryptoWithdrawalFee) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CryptoWithdrawalFee) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CryptoWithdrawalFee) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetDestinationAddress

`func (o *CryptoWithdrawalFee) GetDestinationAddress() string`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *CryptoWithdrawalFee) GetDestinationAddressOk() (*string, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *CryptoWithdrawalFee) SetDestinationAddress(v string)`

SetDestinationAddress sets DestinationAddress field to given value.


### GetCryptoNetwork

`func (o *CryptoWithdrawalFee) GetCryptoNetwork() CryptoNetwork`

GetCryptoNetwork returns the CryptoNetwork field if non-nil, zero value otherwise.

### GetCryptoNetworkOk

`func (o *CryptoWithdrawalFee) GetCryptoNetworkOk() (*CryptoNetwork, bool)`

GetCryptoNetworkOk returns a tuple with the CryptoNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoNetwork

`func (o *CryptoWithdrawalFee) SetCryptoNetwork(v CryptoNetwork)`

SetCryptoNetwork sets CryptoNetwork field to given value.


### GetAmount

`func (o *CryptoWithdrawalFee) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CryptoWithdrawalFee) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CryptoWithdrawalFee) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CryptoWithdrawalFee) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTotal

`func (o *CryptoWithdrawalFee) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CryptoWithdrawalFee) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CryptoWithdrawalFee) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CryptoWithdrawalFee) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


