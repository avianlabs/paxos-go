# CreateQuoteExecutionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteId** | **string** | The ID of the held Quote for buying or selling some asset. | 
**RefId** | Pointer to **string** | A unique identifier for the quote execution (for idempotence). | [optional] 
**BaseAmount** | Pointer to **string** | The amount of the base asset (crypto) to buy or sell using the specified quote. The maximum precision is 8 decimals. | [optional] 
**QuoteAmount** | Pointer to **string** | The amount of the quote asset (fiat) to spend or acquire using the specified quote. The maximum precision is 2 decimals. | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata to store on the quote execution. Up to 6 key/value pairs may be stored, with each key and value at most 100 characters. | [optional] 
**IdentityId** | Pointer to **string** | The end user that requests the quote execution. | [optional] 
**AccountId** | Pointer to **string** | The account under which this quote execution is placed. The provided identity must be allowed to trade on behalf of this account. | [optional] 
**RecipientProfileId** | Pointer to **string** | The ID of the profile under which to deposit the funds. | [optional] 

## Methods

### NewCreateQuoteExecutionRequest

`func NewCreateQuoteExecutionRequest(quoteId string, ) *CreateQuoteExecutionRequest`

NewCreateQuoteExecutionRequest instantiates a new CreateQuoteExecutionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQuoteExecutionRequestWithDefaults

`func NewCreateQuoteExecutionRequestWithDefaults() *CreateQuoteExecutionRequest`

NewCreateQuoteExecutionRequestWithDefaults instantiates a new CreateQuoteExecutionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteId

`func (o *CreateQuoteExecutionRequest) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *CreateQuoteExecutionRequest) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *CreateQuoteExecutionRequest) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetRefId

`func (o *CreateQuoteExecutionRequest) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *CreateQuoteExecutionRequest) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *CreateQuoteExecutionRequest) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *CreateQuoteExecutionRequest) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetBaseAmount

`func (o *CreateQuoteExecutionRequest) GetBaseAmount() string`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *CreateQuoteExecutionRequest) GetBaseAmountOk() (*string, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *CreateQuoteExecutionRequest) SetBaseAmount(v string)`

SetBaseAmount sets BaseAmount field to given value.

### HasBaseAmount

`func (o *CreateQuoteExecutionRequest) HasBaseAmount() bool`

HasBaseAmount returns a boolean if a field has been set.

### GetQuoteAmount

`func (o *CreateQuoteExecutionRequest) GetQuoteAmount() string`

GetQuoteAmount returns the QuoteAmount field if non-nil, zero value otherwise.

### GetQuoteAmountOk

`func (o *CreateQuoteExecutionRequest) GetQuoteAmountOk() (*string, bool)`

GetQuoteAmountOk returns a tuple with the QuoteAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAmount

`func (o *CreateQuoteExecutionRequest) SetQuoteAmount(v string)`

SetQuoteAmount sets QuoteAmount field to given value.

### HasQuoteAmount

`func (o *CreateQuoteExecutionRequest) HasQuoteAmount() bool`

HasQuoteAmount returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateQuoteExecutionRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateQuoteExecutionRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateQuoteExecutionRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateQuoteExecutionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetIdentityId

`func (o *CreateQuoteExecutionRequest) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *CreateQuoteExecutionRequest) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *CreateQuoteExecutionRequest) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *CreateQuoteExecutionRequest) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetAccountId

`func (o *CreateQuoteExecutionRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateQuoteExecutionRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateQuoteExecutionRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreateQuoteExecutionRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetRecipientProfileId

`func (o *CreateQuoteExecutionRequest) GetRecipientProfileId() string`

GetRecipientProfileId returns the RecipientProfileId field if non-nil, zero value otherwise.

### GetRecipientProfileIdOk

`func (o *CreateQuoteExecutionRequest) GetRecipientProfileIdOk() (*string, bool)`

GetRecipientProfileIdOk returns a tuple with the RecipientProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientProfileId

`func (o *CreateQuoteExecutionRequest) SetRecipientProfileId(v string)`

SetRecipientProfileId sets RecipientProfileId field to given value.

### HasRecipientProfileId

`func (o *CreateQuoteExecutionRequest) HasRecipientProfileId() bool`

HasRecipientProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


