# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The transaction identifier used to query or refer to a specific transaction. | 
**RefId** | **string** | Idempotency key. | 
**SettlementWindowStart** | **time.Time** | The start of the window which the transaction is eligible for settlement. If omitted, transactions are immediately eligible for settlement upon success. RFC3339 format, like &#x60;YYYY-MM-DDTHH:MM:SS.sssZ&#x60;. ex: &#x60;2006-01-02T15:04:05Z&#x60;. | 
**SettlementWindowEnd** | **time.Time** | The end of the window which the transaction is eligible for settlement. Transactions which are not cancelled or settled by this time will expire. RFC3339 format, like &#x60;YYYY-MM-DDTHH:MM:SS.sssZ&#x60;. ex: &#x60;2006-01-02T15:04:05Z&#x60;. | 
**SourceProfileId** | **string** | The Profile ID (profile_id) of the entity submitting the transaction. | 
**TargetProfileId** | **string** | The Profile ID (profile_id) of the entity receiving the transaction. | 
**Legs** | [**[]Obligation**](Obligation.md) | The obligations (representing one-way asset movements) to be settled atomically. | 
**Status** | [**TransactionStatus**](TransactionStatus.md) |  | 
**CreatedAt** | **time.Time** | The timestamp when the transaction was first created, RFC3339 format, like &#x60;YYYY-MM-DDTHH:MM:SS.sssZ&#x60;. | 
**UpdatedAt** | **time.Time** | The timestamp when the transaction was last updated, RFC3339 format, like &#x60;YYYY-MM-DDTHH:MM:SS.sssZ&#x60;. | 

## Methods

### NewTransaction

`func NewTransaction(id string, refId string, settlementWindowStart time.Time, settlementWindowEnd time.Time, sourceProfileId string, targetProfileId string, legs []Obligation, status TransactionStatus, createdAt time.Time, updatedAt time.Time, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v string)`

SetId sets Id field to given value.


### GetRefId

`func (o *Transaction) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *Transaction) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *Transaction) SetRefId(v string)`

SetRefId sets RefId field to given value.


### GetSettlementWindowStart

`func (o *Transaction) GetSettlementWindowStart() time.Time`

GetSettlementWindowStart returns the SettlementWindowStart field if non-nil, zero value otherwise.

### GetSettlementWindowStartOk

`func (o *Transaction) GetSettlementWindowStartOk() (*time.Time, bool)`

GetSettlementWindowStartOk returns a tuple with the SettlementWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementWindowStart

`func (o *Transaction) SetSettlementWindowStart(v time.Time)`

SetSettlementWindowStart sets SettlementWindowStart field to given value.


### GetSettlementWindowEnd

`func (o *Transaction) GetSettlementWindowEnd() time.Time`

GetSettlementWindowEnd returns the SettlementWindowEnd field if non-nil, zero value otherwise.

### GetSettlementWindowEndOk

`func (o *Transaction) GetSettlementWindowEndOk() (*time.Time, bool)`

GetSettlementWindowEndOk returns a tuple with the SettlementWindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementWindowEnd

`func (o *Transaction) SetSettlementWindowEnd(v time.Time)`

SetSettlementWindowEnd sets SettlementWindowEnd field to given value.


### GetSourceProfileId

`func (o *Transaction) GetSourceProfileId() string`

GetSourceProfileId returns the SourceProfileId field if non-nil, zero value otherwise.

### GetSourceProfileIdOk

`func (o *Transaction) GetSourceProfileIdOk() (*string, bool)`

GetSourceProfileIdOk returns a tuple with the SourceProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProfileId

`func (o *Transaction) SetSourceProfileId(v string)`

SetSourceProfileId sets SourceProfileId field to given value.


### GetTargetProfileId

`func (o *Transaction) GetTargetProfileId() string`

GetTargetProfileId returns the TargetProfileId field if non-nil, zero value otherwise.

### GetTargetProfileIdOk

`func (o *Transaction) GetTargetProfileIdOk() (*string, bool)`

GetTargetProfileIdOk returns a tuple with the TargetProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProfileId

`func (o *Transaction) SetTargetProfileId(v string)`

SetTargetProfileId sets TargetProfileId field to given value.


### GetLegs

`func (o *Transaction) GetLegs() []Obligation`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *Transaction) GetLegsOk() (*[]Obligation, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *Transaction) SetLegs(v []Obligation)`

SetLegs sets Legs field to given value.


### GetStatus

`func (o *Transaction) GetStatus() TransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*TransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v TransactionStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *Transaction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Transaction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Transaction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Transaction) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Transaction) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Transaction) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


