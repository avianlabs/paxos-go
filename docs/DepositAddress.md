# DepositAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The UUID of the deposit address. | 
**ProfileId** | **string** | The ID of the profile that will be credited when a deposit arrives with this address. | 
**CustomerId** | **string** | The Paxos Customer to which the profile belongs. | 
**CryptoNetwork** | [**CryptoNetwork**](CryptoNetwork.md) |  | 
**IdentityId** | Pointer to **string** | The Identity of the end user who will make deposits to this address. | [optional] 
**RefId** | Pointer to **string** | Client-specified ID for replay protection and lookup. | [optional] 
**Metadata** | Pointer to **map[string]string** | Optional client-specified metadata, which is set on deposit address creation. Up to 6 key/value pairs may be returned. Each key and value must be less than or equal to 100 characters. | [optional] 
**Address** | **string** | The crypto address to send to for deposits. | 
**AccountId** | Pointer to **string** | The Account associated to the identity of the user that will be linked to the created address. | [optional] 
**ConversionTargetAsset** | Pointer to [**DepositAddressConversionTargetAsset**](DepositAddressConversionTargetAsset.md) |  | [optional] 

## Methods

### NewDepositAddress

`func NewDepositAddress(id string, profileId string, customerId string, cryptoNetwork CryptoNetwork, address string, ) *DepositAddress`

NewDepositAddress instantiates a new DepositAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositAddressWithDefaults

`func NewDepositAddressWithDefaults() *DepositAddress`

NewDepositAddressWithDefaults instantiates a new DepositAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DepositAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DepositAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DepositAddress) SetId(v string)`

SetId sets Id field to given value.


### GetProfileId

`func (o *DepositAddress) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *DepositAddress) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *DepositAddress) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetCustomerId

`func (o *DepositAddress) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DepositAddress) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DepositAddress) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetCryptoNetwork

`func (o *DepositAddress) GetCryptoNetwork() CryptoNetwork`

GetCryptoNetwork returns the CryptoNetwork field if non-nil, zero value otherwise.

### GetCryptoNetworkOk

`func (o *DepositAddress) GetCryptoNetworkOk() (*CryptoNetwork, bool)`

GetCryptoNetworkOk returns a tuple with the CryptoNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoNetwork

`func (o *DepositAddress) SetCryptoNetwork(v CryptoNetwork)`

SetCryptoNetwork sets CryptoNetwork field to given value.


### GetIdentityId

`func (o *DepositAddress) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *DepositAddress) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *DepositAddress) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *DepositAddress) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetRefId

`func (o *DepositAddress) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *DepositAddress) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *DepositAddress) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *DepositAddress) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetMetadata

`func (o *DepositAddress) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DepositAddress) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DepositAddress) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DepositAddress) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAddress

`func (o *DepositAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DepositAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DepositAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAccountId

`func (o *DepositAddress) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DepositAddress) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DepositAddress) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *DepositAddress) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetConversionTargetAsset

`func (o *DepositAddress) GetConversionTargetAsset() DepositAddressConversionTargetAsset`

GetConversionTargetAsset returns the ConversionTargetAsset field if non-nil, zero value otherwise.

### GetConversionTargetAssetOk

`func (o *DepositAddress) GetConversionTargetAssetOk() (*DepositAddressConversionTargetAsset, bool)`

GetConversionTargetAssetOk returns a tuple with the ConversionTargetAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionTargetAsset

`func (o *DepositAddress) SetConversionTargetAsset(v DepositAddressConversionTargetAsset)`

SetConversionTargetAsset sets ConversionTargetAsset field to given value.

### HasConversionTargetAsset

`func (o *DepositAddress) HasConversionTargetAsset() bool`

HasConversionTargetAsset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


