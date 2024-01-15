# CreateStablecoinConversionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** | The Profile associated with a conversion. | 
**Amount** | **string** | Asset amount to convert. &lt;a href&#x3D;\&quot;/stablecoin/conversion/asset\&quot;&gt;Details&lt;/a&gt;. | 
**SourceAsset** | **string** | The asset to convert from. &lt;a href&#x3D;\&quot;/stablecoin/conversion/asset\&quot;&gt;Details&lt;/a&gt;. | 
**TargetAsset** | **string** | The asset to convert to. &lt;a href&#x3D;\&quot;/stablecoin/conversion/asset\&quot;&gt;Details&lt;/a&gt;. | 
**RefId** | Pointer to **string** | Client provided, unique Reference ID for lookup and replay protection. | [optional] 
**IdentityId** | Pointer to **string** | The Identity ID associated with the user requesting the conversion. Required only for customers with [3rd-Party integrations](/crypto-brokerage/ledger-type#fiat-and-crypto-subledger). | [optional] 
**AccountId** | Pointer to **string** | The Account ID associated with the user requesting the conversion. Required only for customers with [3rd-Party integrations](/crypto-brokerage/ledger-type#fiat-and-crypto-subledger). | [optional] 
**Metadata** | Pointer to **map[string]string** | Optional client-specified stored metadata. | [optional] 
**RecipientProfileId** | Pointer to **string** | For directed settlement, the receiving side &#x60;profile_id&#x60;. | [optional] 

## Methods

### NewCreateStablecoinConversionRequest

`func NewCreateStablecoinConversionRequest(profileId string, amount string, sourceAsset string, targetAsset string, ) *CreateStablecoinConversionRequest`

NewCreateStablecoinConversionRequest instantiates a new CreateStablecoinConversionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStablecoinConversionRequestWithDefaults

`func NewCreateStablecoinConversionRequestWithDefaults() *CreateStablecoinConversionRequest`

NewCreateStablecoinConversionRequestWithDefaults instantiates a new CreateStablecoinConversionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *CreateStablecoinConversionRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CreateStablecoinConversionRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CreateStablecoinConversionRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetAmount

`func (o *CreateStablecoinConversionRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateStablecoinConversionRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateStablecoinConversionRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetSourceAsset

`func (o *CreateStablecoinConversionRequest) GetSourceAsset() string`

GetSourceAsset returns the SourceAsset field if non-nil, zero value otherwise.

### GetSourceAssetOk

`func (o *CreateStablecoinConversionRequest) GetSourceAssetOk() (*string, bool)`

GetSourceAssetOk returns a tuple with the SourceAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAsset

`func (o *CreateStablecoinConversionRequest) SetSourceAsset(v string)`

SetSourceAsset sets SourceAsset field to given value.


### GetTargetAsset

`func (o *CreateStablecoinConversionRequest) GetTargetAsset() string`

GetTargetAsset returns the TargetAsset field if non-nil, zero value otherwise.

### GetTargetAssetOk

`func (o *CreateStablecoinConversionRequest) GetTargetAssetOk() (*string, bool)`

GetTargetAssetOk returns a tuple with the TargetAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAsset

`func (o *CreateStablecoinConversionRequest) SetTargetAsset(v string)`

SetTargetAsset sets TargetAsset field to given value.


### GetRefId

`func (o *CreateStablecoinConversionRequest) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *CreateStablecoinConversionRequest) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *CreateStablecoinConversionRequest) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *CreateStablecoinConversionRequest) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetIdentityId

`func (o *CreateStablecoinConversionRequest) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *CreateStablecoinConversionRequest) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *CreateStablecoinConversionRequest) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *CreateStablecoinConversionRequest) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetAccountId

`func (o *CreateStablecoinConversionRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateStablecoinConversionRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateStablecoinConversionRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreateStablecoinConversionRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateStablecoinConversionRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateStablecoinConversionRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateStablecoinConversionRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateStablecoinConversionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRecipientProfileId

`func (o *CreateStablecoinConversionRequest) GetRecipientProfileId() string`

GetRecipientProfileId returns the RecipientProfileId field if non-nil, zero value otherwise.

### GetRecipientProfileIdOk

`func (o *CreateStablecoinConversionRequest) GetRecipientProfileIdOk() (*string, bool)`

GetRecipientProfileIdOk returns a tuple with the RecipientProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientProfileId

`func (o *CreateStablecoinConversionRequest) SetRecipientProfileId(v string)`

SetRecipientProfileId sets RecipientProfileId field to given value.

### HasRecipientProfileId

`func (o *CreateStablecoinConversionRequest) HasRecipientProfileId() bool`

HasRecipientProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


