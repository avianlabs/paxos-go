# TickerRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**PricePriceMarket**](PricePriceMarket.md) |  | [optional] 
**BestBid** | Pointer to [**BookLevel**](BookLevel.md) |  | [optional] 
**BestAsk** | Pointer to [**BookLevel**](BookLevel.md) |  | [optional] 
**LastExecution** | Pointer to [**BookLevel**](BookLevel.md) |  | [optional] 
**LastDay** | Pointer to [**ExchangeStats**](ExchangeStats.md) |  | [optional] 
**Today** | Pointer to [**ExchangeStats**](ExchangeStats.md) |  | [optional] 
**SnapshotAt** | Pointer to **time.Time** | The time at which this data was retrieved. | [optional] 

## Methods

### NewTickerRecord

`func NewTickerRecord() *TickerRecord`

NewTickerRecord instantiates a new TickerRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTickerRecordWithDefaults

`func NewTickerRecordWithDefaults() *TickerRecord`

NewTickerRecordWithDefaults instantiates a new TickerRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *TickerRecord) GetMarket() PricePriceMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *TickerRecord) GetMarketOk() (*PricePriceMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *TickerRecord) SetMarket(v PricePriceMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *TickerRecord) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetBestBid

`func (o *TickerRecord) GetBestBid() BookLevel`

GetBestBid returns the BestBid field if non-nil, zero value otherwise.

### GetBestBidOk

`func (o *TickerRecord) GetBestBidOk() (*BookLevel, bool)`

GetBestBidOk returns a tuple with the BestBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestBid

`func (o *TickerRecord) SetBestBid(v BookLevel)`

SetBestBid sets BestBid field to given value.

### HasBestBid

`func (o *TickerRecord) HasBestBid() bool`

HasBestBid returns a boolean if a field has been set.

### GetBestAsk

`func (o *TickerRecord) GetBestAsk() BookLevel`

GetBestAsk returns the BestAsk field if non-nil, zero value otherwise.

### GetBestAskOk

`func (o *TickerRecord) GetBestAskOk() (*BookLevel, bool)`

GetBestAskOk returns a tuple with the BestAsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestAsk

`func (o *TickerRecord) SetBestAsk(v BookLevel)`

SetBestAsk sets BestAsk field to given value.

### HasBestAsk

`func (o *TickerRecord) HasBestAsk() bool`

HasBestAsk returns a boolean if a field has been set.

### GetLastExecution

`func (o *TickerRecord) GetLastExecution() BookLevel`

GetLastExecution returns the LastExecution field if non-nil, zero value otherwise.

### GetLastExecutionOk

`func (o *TickerRecord) GetLastExecutionOk() (*BookLevel, bool)`

GetLastExecutionOk returns a tuple with the LastExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecution

`func (o *TickerRecord) SetLastExecution(v BookLevel)`

SetLastExecution sets LastExecution field to given value.

### HasLastExecution

`func (o *TickerRecord) HasLastExecution() bool`

HasLastExecution returns a boolean if a field has been set.

### GetLastDay

`func (o *TickerRecord) GetLastDay() ExchangeStats`

GetLastDay returns the LastDay field if non-nil, zero value otherwise.

### GetLastDayOk

`func (o *TickerRecord) GetLastDayOk() (*ExchangeStats, bool)`

GetLastDayOk returns a tuple with the LastDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDay

`func (o *TickerRecord) SetLastDay(v ExchangeStats)`

SetLastDay sets LastDay field to given value.

### HasLastDay

`func (o *TickerRecord) HasLastDay() bool`

HasLastDay returns a boolean if a field has been set.

### GetToday

`func (o *TickerRecord) GetToday() ExchangeStats`

GetToday returns the Today field if non-nil, zero value otherwise.

### GetTodayOk

`func (o *TickerRecord) GetTodayOk() (*ExchangeStats, bool)`

GetTodayOk returns a tuple with the Today field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToday

`func (o *TickerRecord) SetToday(v ExchangeStats)`

SetToday sets Today field to given value.

### HasToday

`func (o *TickerRecord) HasToday() bool`

HasToday returns a boolean if a field has been set.

### GetSnapshotAt

`func (o *TickerRecord) GetSnapshotAt() time.Time`

GetSnapshotAt returns the SnapshotAt field if non-nil, zero value otherwise.

### GetSnapshotAtOk

`func (o *TickerRecord) GetSnapshotAtOk() (*time.Time, bool)`

GetSnapshotAtOk returns a tuple with the SnapshotAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotAt

`func (o *TickerRecord) SetSnapshotAt(v time.Time)`

SetSnapshotAt sets SnapshotAt field to given value.

### HasSnapshotAt

`func (o *TickerRecord) HasSnapshotAt() bool`

HasSnapshotAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


