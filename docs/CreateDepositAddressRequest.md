# CreateDepositAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** | The target profile for deposit crediting. | 
**CryptoNetwork** | [**CryptoNetwork**](CryptoNetwork.md) |  | 
**IdentityId** | Pointer to **string** | The Identity of the end user who will make deposits to the created address. | [optional] 
**RefId** | Pointer to **string** | Client-specified ID for replay protection and lookup. | [optional] 
**Metadata** | Pointer to **map[string]string** | Optional client-specified metadata, which will be added to the deposit address but not to deposit transfers to the created address. Up to 6 key/value pairs may be provided. Each key and value must be less than or equal to 100 characters. | [optional] 
**AccountId** | Pointer to **string** | The Account associated to the identity of the user that will be linked to the created address. | [optional] 
**ConversionTargetAsset** | Pointer to [**DepositAddressConversionTargetAsset**](DepositAddressConversionTargetAsset.md) |  | [optional] 

## Methods

### NewCreateDepositAddressRequest

`func NewCreateDepositAddressRequest(profileId string, cryptoNetwork CryptoNetwork, ) *CreateDepositAddressRequest`

NewCreateDepositAddressRequest instantiates a new CreateDepositAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDepositAddressRequestWithDefaults

`func NewCreateDepositAddressRequestWithDefaults() *CreateDepositAddressRequest`

NewCreateDepositAddressRequestWithDefaults instantiates a new CreateDepositAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *CreateDepositAddressRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CreateDepositAddressRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CreateDepositAddressRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetCryptoNetwork

`func (o *CreateDepositAddressRequest) GetCryptoNetwork() CryptoNetwork`

GetCryptoNetwork returns the CryptoNetwork field if non-nil, zero value otherwise.

### GetCryptoNetworkOk

`func (o *CreateDepositAddressRequest) GetCryptoNetworkOk() (*CryptoNetwork, bool)`

GetCryptoNetworkOk returns a tuple with the CryptoNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoNetwork

`func (o *CreateDepositAddressRequest) SetCryptoNetwork(v CryptoNetwork)`

SetCryptoNetwork sets CryptoNetwork field to given value.


### GetIdentityId

`func (o *CreateDepositAddressRequest) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *CreateDepositAddressRequest) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *CreateDepositAddressRequest) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *CreateDepositAddressRequest) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetRefId

`func (o *CreateDepositAddressRequest) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *CreateDepositAddressRequest) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *CreateDepositAddressRequest) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *CreateDepositAddressRequest) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateDepositAddressRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateDepositAddressRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateDepositAddressRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateDepositAddressRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAccountId

`func (o *CreateDepositAddressRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateDepositAddressRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateDepositAddressRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreateDepositAddressRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetConversionTargetAsset

`func (o *CreateDepositAddressRequest) GetConversionTargetAsset() DepositAddressConversionTargetAsset`

GetConversionTargetAsset returns the ConversionTargetAsset field if non-nil, zero value otherwise.

### GetConversionTargetAssetOk

`func (o *CreateDepositAddressRequest) GetConversionTargetAssetOk() (*DepositAddressConversionTargetAsset, bool)`

GetConversionTargetAssetOk returns a tuple with the ConversionTargetAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionTargetAsset

`func (o *CreateDepositAddressRequest) SetConversionTargetAsset(v DepositAddressConversionTargetAsset)`

SetConversionTargetAsset sets ConversionTargetAsset field to given value.

### HasConversionTargetAsset

`func (o *CreateDepositAddressRequest) HasConversionTargetAsset() bool`

HasConversionTargetAsset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


