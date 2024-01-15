# StablecoinConversion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | System provided UUID for the conversion is provided in the [Create Stablecoin Conversion](#operation/CreateStablecoinConversion) response.  Required parameter for the &lt;a href&#x3D;\&quot;#operation/GetStablecoinConversion\&quot;&gt;Get Stablecoin Conversion&lt;/a&gt; request. | [optional] 
**ProfileId** | Pointer to **string** | The Profile associated with a conversion. Required in the &lt;a href&#x3D;\&quot;#operation/CreateStablecoinConversion\&quot;&gt;Create Stablecoin Conversion&lt;/a&gt; request. | [optional] 
**Amount** | Pointer to **string** | Asset amount to convert. &lt;a href&#x3D;\&quot;/stablecoin/conversion/asset\&quot;&gt;Details&lt;/a&gt;. | [optional] 
**SourceAsset** | Pointer to **string** | The asset to convert from. &lt;a href&#x3D;\&quot;/stablecoin/conversion/asset\&quot;&gt;Details&lt;/a&gt;. | [optional] 
**TargetAsset** | Pointer to **string** | The asset to convert to. &lt;a href&#x3D;\&quot;/stablecoin/conversion/asset\&quot;&gt;Details&lt;/a&gt;. | [optional] 
**Status** | Pointer to **string** | The current status of the conversion. &lt;a href&#x3D;\&quot;/stablecoin/conversion/status\&quot;&gt;Details&lt;/a&gt;. | [optional] 
**RefId** | Pointer to **string** | Client provided, unique Reference ID included the &lt;a href&#x3D;\&quot;#operation/CreateStablecoinConversion\&quot;&gt;Create Stablecoin Conversion&lt;/a&gt; request. | [optional] 
**IdentityId** | Pointer to **string** | The Identity ID associated with the user requesting the conversion. Required only for customers with [3rd-Party integrations](/crypto-brokerage/ledger-type#fiat-and-crypto-subledger). | [optional] 
**AccountId** | Pointer to **string** | The Account ID associated with the user requesting the conversion. Required only for customers with [3rd-Party integrations](/crypto-brokerage/ledger-type#fiat-and-crypto-subledger). | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time at which the conversion was requested. See RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time at which the conversion was last updated. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | [optional] 
**SettledAt** | Pointer to **time.Time** | The time at which the conversion was settled. &lt;a href&#x3D;\&quot;/stablecoin/conversion/status\&quot;&gt;Details&lt;/a&gt;. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | [optional] 
**CancelledAt** | Pointer to **time.Time** | The time at which the conversion has been cancelled. &lt;a href&#x3D;\&quot;/stablecoin/conversion/status\&quot;&gt;Details&lt;/a&gt;. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | [optional] 
**Metadata** | Pointer to **map[string]string** | Optional client-specified stored metadata. | [optional] 

## Methods

### NewStablecoinConversion

`func NewStablecoinConversion() *StablecoinConversion`

NewStablecoinConversion instantiates a new StablecoinConversion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStablecoinConversionWithDefaults

`func NewStablecoinConversionWithDefaults() *StablecoinConversion`

NewStablecoinConversionWithDefaults instantiates a new StablecoinConversion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StablecoinConversion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StablecoinConversion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StablecoinConversion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StablecoinConversion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProfileId

`func (o *StablecoinConversion) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *StablecoinConversion) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *StablecoinConversion) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *StablecoinConversion) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetAmount

`func (o *StablecoinConversion) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StablecoinConversion) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StablecoinConversion) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *StablecoinConversion) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetSourceAsset

`func (o *StablecoinConversion) GetSourceAsset() string`

GetSourceAsset returns the SourceAsset field if non-nil, zero value otherwise.

### GetSourceAssetOk

`func (o *StablecoinConversion) GetSourceAssetOk() (*string, bool)`

GetSourceAssetOk returns a tuple with the SourceAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAsset

`func (o *StablecoinConversion) SetSourceAsset(v string)`

SetSourceAsset sets SourceAsset field to given value.

### HasSourceAsset

`func (o *StablecoinConversion) HasSourceAsset() bool`

HasSourceAsset returns a boolean if a field has been set.

### GetTargetAsset

`func (o *StablecoinConversion) GetTargetAsset() string`

GetTargetAsset returns the TargetAsset field if non-nil, zero value otherwise.

### GetTargetAssetOk

`func (o *StablecoinConversion) GetTargetAssetOk() (*string, bool)`

GetTargetAssetOk returns a tuple with the TargetAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAsset

`func (o *StablecoinConversion) SetTargetAsset(v string)`

SetTargetAsset sets TargetAsset field to given value.

### HasTargetAsset

`func (o *StablecoinConversion) HasTargetAsset() bool`

HasTargetAsset returns a boolean if a field has been set.

### GetStatus

`func (o *StablecoinConversion) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StablecoinConversion) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StablecoinConversion) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StablecoinConversion) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRefId

`func (o *StablecoinConversion) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *StablecoinConversion) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *StablecoinConversion) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *StablecoinConversion) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetIdentityId

`func (o *StablecoinConversion) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *StablecoinConversion) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *StablecoinConversion) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *StablecoinConversion) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetAccountId

`func (o *StablecoinConversion) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *StablecoinConversion) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *StablecoinConversion) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *StablecoinConversion) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StablecoinConversion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StablecoinConversion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StablecoinConversion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StablecoinConversion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *StablecoinConversion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StablecoinConversion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StablecoinConversion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StablecoinConversion) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetSettledAt

`func (o *StablecoinConversion) GetSettledAt() time.Time`

GetSettledAt returns the SettledAt field if non-nil, zero value otherwise.

### GetSettledAtOk

`func (o *StablecoinConversion) GetSettledAtOk() (*time.Time, bool)`

GetSettledAtOk returns a tuple with the SettledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAt

`func (o *StablecoinConversion) SetSettledAt(v time.Time)`

SetSettledAt sets SettledAt field to given value.

### HasSettledAt

`func (o *StablecoinConversion) HasSettledAt() bool`

HasSettledAt returns a boolean if a field has been set.

### GetCancelledAt

`func (o *StablecoinConversion) GetCancelledAt() time.Time`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *StablecoinConversion) GetCancelledAtOk() (*time.Time, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *StablecoinConversion) SetCancelledAt(v time.Time)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *StablecoinConversion) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### GetMetadata

`func (o *StablecoinConversion) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StablecoinConversion) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StablecoinConversion) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StablecoinConversion) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


