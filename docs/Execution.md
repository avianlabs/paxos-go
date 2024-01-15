# Execution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** | The UUID of the order associated with the execution. | [optional] 
**ExecutedAt** | Pointer to **time.Time** | Timestamp of the execution. | [optional] 
**Market** | Pointer to [**Market**](Market.md) |  | [optional] 
**Side** | Pointer to [**OrderSide**](OrderSide.md) |  | [optional] 
**Amount** | Pointer to **string** | Execution amount. | [optional] 
**Price** | Pointer to **string** | Execution price. | [optional] 
**Commission** | Pointer to **string** | Amount of commission paid. | [optional] 
**CommissionAsset** | Pointer to [**Asset**](Asset.md) |  | [optional] 
**Rebate** | Pointer to **string** | Amount of rebate applied. | [optional] 
**RebateAsset** | Pointer to [**Asset**](Asset.md) |  | [optional] 
**AccountId** | Pointer to **string** | Account ID associated with the execution. | [optional] 
**GrossTradeAmount** | Pointer to **string** | The total asset traded (asset amount multiplied by price paid). | [optional] 

## Methods

### NewExecution

`func NewExecution() *Execution`

NewExecution instantiates a new Execution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionWithDefaults

`func NewExecutionWithDefaults() *Execution`

NewExecutionWithDefaults instantiates a new Execution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *Execution) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *Execution) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *Execution) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *Execution) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetOrderId

`func (o *Execution) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Execution) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Execution) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *Execution) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetExecutedAt

`func (o *Execution) GetExecutedAt() time.Time`

GetExecutedAt returns the ExecutedAt field if non-nil, zero value otherwise.

### GetExecutedAtOk

`func (o *Execution) GetExecutedAtOk() (*time.Time, bool)`

GetExecutedAtOk returns a tuple with the ExecutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAt

`func (o *Execution) SetExecutedAt(v time.Time)`

SetExecutedAt sets ExecutedAt field to given value.

### HasExecutedAt

`func (o *Execution) HasExecutedAt() bool`

HasExecutedAt returns a boolean if a field has been set.

### GetMarket

`func (o *Execution) GetMarket() Market`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *Execution) GetMarketOk() (*Market, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *Execution) SetMarket(v Market)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *Execution) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetSide

`func (o *Execution) GetSide() OrderSide`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *Execution) GetSideOk() (*OrderSide, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *Execution) SetSide(v OrderSide)`

SetSide sets Side field to given value.

### HasSide

`func (o *Execution) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetAmount

`func (o *Execution) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Execution) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Execution) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Execution) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPrice

`func (o *Execution) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Execution) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Execution) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Execution) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCommission

`func (o *Execution) GetCommission() string`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *Execution) GetCommissionOk() (*string, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *Execution) SetCommission(v string)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *Execution) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetCommissionAsset

`func (o *Execution) GetCommissionAsset() Asset`

GetCommissionAsset returns the CommissionAsset field if non-nil, zero value otherwise.

### GetCommissionAssetOk

`func (o *Execution) GetCommissionAssetOk() (*Asset, bool)`

GetCommissionAssetOk returns a tuple with the CommissionAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAsset

`func (o *Execution) SetCommissionAsset(v Asset)`

SetCommissionAsset sets CommissionAsset field to given value.

### HasCommissionAsset

`func (o *Execution) HasCommissionAsset() bool`

HasCommissionAsset returns a boolean if a field has been set.

### GetRebate

`func (o *Execution) GetRebate() string`

GetRebate returns the Rebate field if non-nil, zero value otherwise.

### GetRebateOk

`func (o *Execution) GetRebateOk() (*string, bool)`

GetRebateOk returns a tuple with the Rebate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebate

`func (o *Execution) SetRebate(v string)`

SetRebate sets Rebate field to given value.

### HasRebate

`func (o *Execution) HasRebate() bool`

HasRebate returns a boolean if a field has been set.

### GetRebateAsset

`func (o *Execution) GetRebateAsset() Asset`

GetRebateAsset returns the RebateAsset field if non-nil, zero value otherwise.

### GetRebateAssetOk

`func (o *Execution) GetRebateAssetOk() (*Asset, bool)`

GetRebateAssetOk returns a tuple with the RebateAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebateAsset

`func (o *Execution) SetRebateAsset(v Asset)`

SetRebateAsset sets RebateAsset field to given value.

### HasRebateAsset

`func (o *Execution) HasRebateAsset() bool`

HasRebateAsset returns a boolean if a field has been set.

### GetAccountId

`func (o *Execution) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Execution) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Execution) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Execution) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetGrossTradeAmount

`func (o *Execution) GetGrossTradeAmount() string`

GetGrossTradeAmount returns the GrossTradeAmount field if non-nil, zero value otherwise.

### GetGrossTradeAmountOk

`func (o *Execution) GetGrossTradeAmountOk() (*string, bool)`

GetGrossTradeAmountOk returns a tuple with the GrossTradeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossTradeAmount

`func (o *Execution) SetGrossTradeAmount(v string)`

SetGrossTradeAmount sets GrossTradeAmount field to given value.

### HasGrossTradeAmount

`func (o *Execution) HasGrossTradeAmount() bool`

HasGrossTradeAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


