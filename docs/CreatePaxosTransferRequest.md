# CreatePaxosTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefId** | Pointer to **string** | Client-specified ID for replay protection and lookup. | [optional] 
**FromProfileId** | **string** | The profile from which to send funds. | 
**ToProfileId** | **string** | The destination profile. | 
**Amount** | **string** | The amount to transfer. | 
**Asset** | **string** | The asset to transfer, e.g. \&quot;USD\&quot;, \&quot;BTC\&quot;, \&quot;ETH\&quot;. | 
**Metadata** | Pointer to **map[string]string** | Optional client-specified metadata on the sender side of the transaction. Up to 6 key/value pairs may be provided. Each key and value must be less than or equal to 100 characters. | [optional] 
**RecipientMetadata** | Pointer to **map[string]string** | Optional client-specified metadata for the recipient side of the transaction. Up to 6 key/value pairs may be provided. Each key and value must be less than or equal to 100 characters. | [optional] 

## Methods

### NewCreatePaxosTransferRequest

`func NewCreatePaxosTransferRequest(fromProfileId string, toProfileId string, amount string, asset string, ) *CreatePaxosTransferRequest`

NewCreatePaxosTransferRequest instantiates a new CreatePaxosTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaxosTransferRequestWithDefaults

`func NewCreatePaxosTransferRequestWithDefaults() *CreatePaxosTransferRequest`

NewCreatePaxosTransferRequestWithDefaults instantiates a new CreatePaxosTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefId

`func (o *CreatePaxosTransferRequest) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *CreatePaxosTransferRequest) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *CreatePaxosTransferRequest) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *CreatePaxosTransferRequest) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetFromProfileId

`func (o *CreatePaxosTransferRequest) GetFromProfileId() string`

GetFromProfileId returns the FromProfileId field if non-nil, zero value otherwise.

### GetFromProfileIdOk

`func (o *CreatePaxosTransferRequest) GetFromProfileIdOk() (*string, bool)`

GetFromProfileIdOk returns a tuple with the FromProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromProfileId

`func (o *CreatePaxosTransferRequest) SetFromProfileId(v string)`

SetFromProfileId sets FromProfileId field to given value.


### GetToProfileId

`func (o *CreatePaxosTransferRequest) GetToProfileId() string`

GetToProfileId returns the ToProfileId field if non-nil, zero value otherwise.

### GetToProfileIdOk

`func (o *CreatePaxosTransferRequest) GetToProfileIdOk() (*string, bool)`

GetToProfileIdOk returns a tuple with the ToProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToProfileId

`func (o *CreatePaxosTransferRequest) SetToProfileId(v string)`

SetToProfileId sets ToProfileId field to given value.


### GetAmount

`func (o *CreatePaxosTransferRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreatePaxosTransferRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreatePaxosTransferRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAsset

`func (o *CreatePaxosTransferRequest) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *CreatePaxosTransferRequest) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *CreatePaxosTransferRequest) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetMetadata

`func (o *CreatePaxosTransferRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreatePaxosTransferRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreatePaxosTransferRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreatePaxosTransferRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRecipientMetadata

`func (o *CreatePaxosTransferRequest) GetRecipientMetadata() map[string]string`

GetRecipientMetadata returns the RecipientMetadata field if non-nil, zero value otherwise.

### GetRecipientMetadataOk

`func (o *CreatePaxosTransferRequest) GetRecipientMetadataOk() (*map[string]string, bool)`

GetRecipientMetadataOk returns a tuple with the RecipientMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientMetadata

`func (o *CreatePaxosTransferRequest) SetRecipientMetadata(v map[string]string)`

SetRecipientMetadata sets RecipientMetadata field to given value.

### HasRecipientMetadata

`func (o *CreatePaxosTransferRequest) HasRecipientMetadata() bool`

HasRecipientMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


