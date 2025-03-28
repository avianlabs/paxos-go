/*
Paxos API

<p>Welcome to Paxos APIs. At Paxos, our mission is to enable the movement of any asset, any time, in a trustworthy way. These APIs serve that mission by making it easier than ever for you to directly integrate our product capabilities into your application, leveraging the speed, stability, and security of the Paxos platform.</p> <p>The documentation that follows gives you access to our Crypto Brokerage, Trading, and Exchange products. It includes APIs for market data, orders, and the held rate quote flow.</p> <p>To test in our sandbox environment, <a href=\"https://account.sandbox.paxos.com\" target=\"_blank\">sign up</a> for an account. For more information about Paxos and our APIs, visit <a href=\"https://www.paxos.com/\" target=\"_blank\">Paxos.com</a>.</p> 

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paxos

import (
	"encoding/json"
)

// checks if the ExchangeStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeStats{}

// ExchangeStats struct for ExchangeStats
type ExchangeStats struct {
	// Highest price in range.
	High *string `json:"high,omitempty"`
	// Lowest price in range.
	Low *string `json:"low,omitempty"`
	// First price in range.
	Open *string `json:"open,omitempty"`
	Volume *string `json:"volume,omitempty"`
	// Volume-Weighted Average Price over Time Period.
	VolumeWeightedAveragePrice *string `json:"volume_weighted_average_price,omitempty"`
	Range *TimestampRange `json:"range,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExchangeStats ExchangeStats

// NewExchangeStats instantiates a new ExchangeStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeStats() *ExchangeStats {
	this := ExchangeStats{}
	return &this
}

// NewExchangeStatsWithDefaults instantiates a new ExchangeStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeStatsWithDefaults() *ExchangeStats {
	this := ExchangeStats{}
	return &this
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *ExchangeStats) GetHigh() string {
	if o == nil || IsNil(o.High) {
		var ret string
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeStats) GetHighOk() (*string, bool) {
	if o == nil || IsNil(o.High) {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *ExchangeStats) HasHigh() bool {
	if o != nil && !IsNil(o.High) {
		return true
	}

	return false
}

// SetHigh gets a reference to the given string and assigns it to the High field.
func (o *ExchangeStats) SetHigh(v string) {
	o.High = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *ExchangeStats) GetLow() string {
	if o == nil || IsNil(o.Low) {
		var ret string
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeStats) GetLowOk() (*string, bool) {
	if o == nil || IsNil(o.Low) {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *ExchangeStats) HasLow() bool {
	if o != nil && !IsNil(o.Low) {
		return true
	}

	return false
}

// SetLow gets a reference to the given string and assigns it to the Low field.
func (o *ExchangeStats) SetLow(v string) {
	o.Low = &v
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *ExchangeStats) GetOpen() string {
	if o == nil || IsNil(o.Open) {
		var ret string
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeStats) GetOpenOk() (*string, bool) {
	if o == nil || IsNil(o.Open) {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *ExchangeStats) HasOpen() bool {
	if o != nil && !IsNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given string and assigns it to the Open field.
func (o *ExchangeStats) SetOpen(v string) {
	o.Open = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *ExchangeStats) GetVolume() string {
	if o == nil || IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeStats) GetVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *ExchangeStats) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *ExchangeStats) SetVolume(v string) {
	o.Volume = &v
}

// GetVolumeWeightedAveragePrice returns the VolumeWeightedAveragePrice field value if set, zero value otherwise.
func (o *ExchangeStats) GetVolumeWeightedAveragePrice() string {
	if o == nil || IsNil(o.VolumeWeightedAveragePrice) {
		var ret string
		return ret
	}
	return *o.VolumeWeightedAveragePrice
}

// GetVolumeWeightedAveragePriceOk returns a tuple with the VolumeWeightedAveragePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeStats) GetVolumeWeightedAveragePriceOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeWeightedAveragePrice) {
		return nil, false
	}
	return o.VolumeWeightedAveragePrice, true
}

// HasVolumeWeightedAveragePrice returns a boolean if a field has been set.
func (o *ExchangeStats) HasVolumeWeightedAveragePrice() bool {
	if o != nil && !IsNil(o.VolumeWeightedAveragePrice) {
		return true
	}

	return false
}

// SetVolumeWeightedAveragePrice gets a reference to the given string and assigns it to the VolumeWeightedAveragePrice field.
func (o *ExchangeStats) SetVolumeWeightedAveragePrice(v string) {
	o.VolumeWeightedAveragePrice = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *ExchangeStats) GetRange() TimestampRange {
	if o == nil || IsNil(o.Range) {
		var ret TimestampRange
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeStats) GetRangeOk() (*TimestampRange, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *ExchangeStats) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given TimestampRange and assigns it to the Range field.
func (o *ExchangeStats) SetRange(v TimestampRange) {
	o.Range = &v
}

func (o ExchangeStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.High) {
		toSerialize["high"] = o.High
	}
	if !IsNil(o.Low) {
		toSerialize["low"] = o.Low
	}
	if !IsNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	if !IsNil(o.VolumeWeightedAveragePrice) {
		toSerialize["volume_weighted_average_price"] = o.VolumeWeightedAveragePrice
	}
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExchangeStats) UnmarshalJSON(data []byte) (err error) {
	varExchangeStats := _ExchangeStats{}

	err = json.Unmarshal(data, &varExchangeStats)

	if err != nil {
		return err
	}

	*o = ExchangeStats(varExchangeStats)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "high")
		delete(additionalProperties, "low")
		delete(additionalProperties, "open")
		delete(additionalProperties, "volume")
		delete(additionalProperties, "volume_weighted_average_price")
		delete(additionalProperties, "range")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeStats struct {
	value *ExchangeStats
	isSet bool
}

func (v NullableExchangeStats) Get() *ExchangeStats {
	return v.value
}

func (v *NullableExchangeStats) Set(val *ExchangeStats) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeStats) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeStats(val *ExchangeStats) *NullableExchangeStats {
	return &NullableExchangeStats{value: val, isSet: true}
}

func (v NullableExchangeStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


