# TransferLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the limit. Single-transaction limits will not have an id, but only a rule id. | [optional] 
**RuleId** | Pointer to **string** | The limit rule this limit applies to. | [optional] 
**TransferTypes** | Pointer to [**[]TransferType**](TransferType.md) | One or more transfer types that the limit applies to, ordered alphabetically. | [optional] 
**PeriodSeconds** | Pointer to **string** | The number of seconds that the limit is measured over. Single-transaction limits have period 0. | [optional] 
**Limit** | Pointer to **string** | The limit. | [optional] 
**Remaining** | Pointer to **string** | The limit, minus any transactions that have been done in the period. Single-transaction limits always have remaining set to the limit. | [optional] 
**LimitAsset** | Pointer to **string** | The asset that limit and remaining are given in, e.g. \&quot;USD\&quot;, \&quot;BTC\&quot;, \&quot;ETH\&quot;. | [optional] 

## Methods

### NewTransferLimit

`func NewTransferLimit() *TransferLimit`

NewTransferLimit instantiates a new TransferLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferLimitWithDefaults

`func NewTransferLimitWithDefaults() *TransferLimit`

NewTransferLimitWithDefaults instantiates a new TransferLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransferLimit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferLimit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferLimit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransferLimit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRuleId

`func (o *TransferLimit) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *TransferLimit) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *TransferLimit) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *TransferLimit) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetTransferTypes

`func (o *TransferLimit) GetTransferTypes() []TransferType`

GetTransferTypes returns the TransferTypes field if non-nil, zero value otherwise.

### GetTransferTypesOk

`func (o *TransferLimit) GetTransferTypesOk() (*[]TransferType, bool)`

GetTransferTypesOk returns a tuple with the TransferTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferTypes

`func (o *TransferLimit) SetTransferTypes(v []TransferType)`

SetTransferTypes sets TransferTypes field to given value.

### HasTransferTypes

`func (o *TransferLimit) HasTransferTypes() bool`

HasTransferTypes returns a boolean if a field has been set.

### GetPeriodSeconds

`func (o *TransferLimit) GetPeriodSeconds() string`

GetPeriodSeconds returns the PeriodSeconds field if non-nil, zero value otherwise.

### GetPeriodSecondsOk

`func (o *TransferLimit) GetPeriodSecondsOk() (*string, bool)`

GetPeriodSecondsOk returns a tuple with the PeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodSeconds

`func (o *TransferLimit) SetPeriodSeconds(v string)`

SetPeriodSeconds sets PeriodSeconds field to given value.

### HasPeriodSeconds

`func (o *TransferLimit) HasPeriodSeconds() bool`

HasPeriodSeconds returns a boolean if a field has been set.

### GetLimit

`func (o *TransferLimit) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TransferLimit) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TransferLimit) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TransferLimit) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetRemaining

`func (o *TransferLimit) GetRemaining() string`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *TransferLimit) GetRemainingOk() (*string, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *TransferLimit) SetRemaining(v string)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *TransferLimit) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.

### GetLimitAsset

`func (o *TransferLimit) GetLimitAsset() string`

GetLimitAsset returns the LimitAsset field if non-nil, zero value otherwise.

### GetLimitAssetOk

`func (o *TransferLimit) GetLimitAssetOk() (*string, bool)`

GetLimitAssetOk returns a tuple with the LimitAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitAsset

`func (o *TransferLimit) SetLimitAsset(v string)`

SetLimitAsset sets LimitAsset field to given value.

### HasLimitAsset

`func (o *TransferLimit) HasLimitAsset() bool`

HasLimitAsset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


