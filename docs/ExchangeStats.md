# ExchangeStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**High** | Pointer to **string** | Highest price in range. | [optional] 
**Low** | Pointer to **string** | Lowest price in range. | [optional] 
**Open** | Pointer to **string** | First price in range. | [optional] 
**Volume** | Pointer to **string** |  | [optional] 
**VolumeWeightedAveragePrice** | Pointer to **string** | Volume-Weighted Average Price over Time Period. | [optional] 
**Range** | Pointer to [**TimestampRange**](TimestampRange.md) |  | [optional] 

## Methods

### NewExchangeStats

`func NewExchangeStats() *ExchangeStats`

NewExchangeStats instantiates a new ExchangeStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeStatsWithDefaults

`func NewExchangeStatsWithDefaults() *ExchangeStats`

NewExchangeStatsWithDefaults instantiates a new ExchangeStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHigh

`func (o *ExchangeStats) GetHigh() string`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *ExchangeStats) GetHighOk() (*string, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *ExchangeStats) SetHigh(v string)`

SetHigh sets High field to given value.

### HasHigh

`func (o *ExchangeStats) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetLow

`func (o *ExchangeStats) GetLow() string`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *ExchangeStats) GetLowOk() (*string, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *ExchangeStats) SetLow(v string)`

SetLow sets Low field to given value.

### HasLow

`func (o *ExchangeStats) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetOpen

`func (o *ExchangeStats) GetOpen() string`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *ExchangeStats) GetOpenOk() (*string, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *ExchangeStats) SetOpen(v string)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *ExchangeStats) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetVolume

`func (o *ExchangeStats) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *ExchangeStats) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *ExchangeStats) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *ExchangeStats) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetVolumeWeightedAveragePrice

`func (o *ExchangeStats) GetVolumeWeightedAveragePrice() string`

GetVolumeWeightedAveragePrice returns the VolumeWeightedAveragePrice field if non-nil, zero value otherwise.

### GetVolumeWeightedAveragePriceOk

`func (o *ExchangeStats) GetVolumeWeightedAveragePriceOk() (*string, bool)`

GetVolumeWeightedAveragePriceOk returns a tuple with the VolumeWeightedAveragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeWeightedAveragePrice

`func (o *ExchangeStats) SetVolumeWeightedAveragePrice(v string)`

SetVolumeWeightedAveragePrice sets VolumeWeightedAveragePrice field to given value.

### HasVolumeWeightedAveragePrice

`func (o *ExchangeStats) HasVolumeWeightedAveragePrice() bool`

HasVolumeWeightedAveragePrice returns a boolean if a field has been set.

### GetRange

`func (o *ExchangeStats) GetRange() TimestampRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *ExchangeStats) GetRangeOk() (*TimestampRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *ExchangeStats) SetRange(v TimestampRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *ExchangeStats) HasRange() bool`

HasRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


