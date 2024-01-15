# HistoricalPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AveragePrice** | Pointer to **string** | Time-weighted average price over the specified time period, beginning at the defined timestamp. For example, if &#x60;increment&#x60; is set to &#x60;ONE_HOUR&#x60;, then &#x60;average_price&#x60; is the time-weighted average for an hour. The entire time period must be completed to display results. | [optional] 
**Timestamp** | Pointer to **time.Time** | Timestamp at the beginning of an increment. | [optional] 

## Methods

### NewHistoricalPrice

`func NewHistoricalPrice() *HistoricalPrice`

NewHistoricalPrice instantiates a new HistoricalPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalPriceWithDefaults

`func NewHistoricalPriceWithDefaults() *HistoricalPrice`

NewHistoricalPriceWithDefaults instantiates a new HistoricalPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAveragePrice

`func (o *HistoricalPrice) GetAveragePrice() string`

GetAveragePrice returns the AveragePrice field if non-nil, zero value otherwise.

### GetAveragePriceOk

`func (o *HistoricalPrice) GetAveragePriceOk() (*string, bool)`

GetAveragePriceOk returns a tuple with the AveragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragePrice

`func (o *HistoricalPrice) SetAveragePrice(v string)`

SetAveragePrice sets AveragePrice field to given value.

### HasAveragePrice

`func (o *HistoricalPrice) HasAveragePrice() bool`

HasAveragePrice returns a boolean if a field has been set.

### GetTimestamp

`func (o *HistoricalPrice) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HistoricalPrice) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HistoricalPrice) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *HistoricalPrice) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


