# CreateTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefId** | Pointer to **string** | Idempotency key. | [optional] 
**SettlementWindowStart** | Pointer to **time.Time** | The start of the window which the transaction is eligible for settlement. If omitted, transactions are immediately eligible for settlement upon success. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | [optional] 
**SettlementWindowEnd** | Pointer to **time.Time** | The end of the window which the transaction is eligible for settlement. Transactions which are not cancelled or settled by this time will expire. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | [optional] 
**SourceProfileId** | Pointer to **string** | The &#x60;profile_id&#x60; of the entity submitting the transaction. | [optional] 
**TargetProfileId** | Pointer to **string** | The &#x60;profile_id&#x60; of the entity receiving the transaction. | [optional] 
**Legs** | Pointer to [**[]CreateObligationRequest**](CreateObligationRequest.md) | The obligations (representing one-way asset movements) to be settled atomically. | [optional] 

## Methods

### NewCreateTransactionRequest

`func NewCreateTransactionRequest() *CreateTransactionRequest`

NewCreateTransactionRequest instantiates a new CreateTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransactionRequestWithDefaults

`func NewCreateTransactionRequestWithDefaults() *CreateTransactionRequest`

NewCreateTransactionRequestWithDefaults instantiates a new CreateTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefId

`func (o *CreateTransactionRequest) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *CreateTransactionRequest) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *CreateTransactionRequest) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *CreateTransactionRequest) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetSettlementWindowStart

`func (o *CreateTransactionRequest) GetSettlementWindowStart() time.Time`

GetSettlementWindowStart returns the SettlementWindowStart field if non-nil, zero value otherwise.

### GetSettlementWindowStartOk

`func (o *CreateTransactionRequest) GetSettlementWindowStartOk() (*time.Time, bool)`

GetSettlementWindowStartOk returns a tuple with the SettlementWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementWindowStart

`func (o *CreateTransactionRequest) SetSettlementWindowStart(v time.Time)`

SetSettlementWindowStart sets SettlementWindowStart field to given value.

### HasSettlementWindowStart

`func (o *CreateTransactionRequest) HasSettlementWindowStart() bool`

HasSettlementWindowStart returns a boolean if a field has been set.

### GetSettlementWindowEnd

`func (o *CreateTransactionRequest) GetSettlementWindowEnd() time.Time`

GetSettlementWindowEnd returns the SettlementWindowEnd field if non-nil, zero value otherwise.

### GetSettlementWindowEndOk

`func (o *CreateTransactionRequest) GetSettlementWindowEndOk() (*time.Time, bool)`

GetSettlementWindowEndOk returns a tuple with the SettlementWindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementWindowEnd

`func (o *CreateTransactionRequest) SetSettlementWindowEnd(v time.Time)`

SetSettlementWindowEnd sets SettlementWindowEnd field to given value.

### HasSettlementWindowEnd

`func (o *CreateTransactionRequest) HasSettlementWindowEnd() bool`

HasSettlementWindowEnd returns a boolean if a field has been set.

### GetSourceProfileId

`func (o *CreateTransactionRequest) GetSourceProfileId() string`

GetSourceProfileId returns the SourceProfileId field if non-nil, zero value otherwise.

### GetSourceProfileIdOk

`func (o *CreateTransactionRequest) GetSourceProfileIdOk() (*string, bool)`

GetSourceProfileIdOk returns a tuple with the SourceProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProfileId

`func (o *CreateTransactionRequest) SetSourceProfileId(v string)`

SetSourceProfileId sets SourceProfileId field to given value.

### HasSourceProfileId

`func (o *CreateTransactionRequest) HasSourceProfileId() bool`

HasSourceProfileId returns a boolean if a field has been set.

### GetTargetProfileId

`func (o *CreateTransactionRequest) GetTargetProfileId() string`

GetTargetProfileId returns the TargetProfileId field if non-nil, zero value otherwise.

### GetTargetProfileIdOk

`func (o *CreateTransactionRequest) GetTargetProfileIdOk() (*string, bool)`

GetTargetProfileIdOk returns a tuple with the TargetProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProfileId

`func (o *CreateTransactionRequest) SetTargetProfileId(v string)`

SetTargetProfileId sets TargetProfileId field to given value.

### HasTargetProfileId

`func (o *CreateTransactionRequest) HasTargetProfileId() bool`

HasTargetProfileId returns a boolean if a field has been set.

### GetLegs

`func (o *CreateTransactionRequest) GetLegs() []CreateObligationRequest`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *CreateTransactionRequest) GetLegsOk() (*[]CreateObligationRequest, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *CreateTransactionRequest) SetLegs(v []CreateObligationRequest)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *CreateTransactionRequest) HasLegs() bool`

HasLegs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


