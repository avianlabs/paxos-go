/*
Paxos API

<p>Welcome to Paxos APIs. At Paxos, our mission is to enable the movement of any asset, any time, in a trustworthy way. These APIs serve that mission by making it easier than ever for you to directly integrate our product capabilities into your application, leveraging the speed, stability, and security of the Paxos platform.</p> <p>The documentation that follows gives you access to our Crypto Brokerage, Trading, and Exchange products. It includes APIs for market data, orders, and the held rate quote flow.</p> <p>To test in our sandbox environment, <a href=\"https://account.sandbox.paxos.com\" target=\"_blank\">sign up</a> for an account. For more information about Paxos and our APIs, visit <a href=\"https://www.paxos.com/\" target=\"_blank\">Paxos.com</a>.</p> 

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paxos

import (
	"encoding/json"
	"time"
)

// checks if the Pricing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pricing{}

// Pricing struct for Pricing
type Pricing struct {
	Market *PricePriceMarket `json:"market,omitempty"`
	// The current price for the given market.
	CurrentPrice *string `json:"current_price,omitempty"`
	// The one-minute average price from 24 hours ago. Updates once per minute.
	YesterdayPrice *string `json:"yesterday_price,omitempty"`
	// The time when the price was generated. RFC3339 format, like `2023-01-03T18:27:40.294528Z`.
	SnapshotAt *time.Time `json:"snapshot_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Pricing Pricing

// NewPricing instantiates a new Pricing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricing() *Pricing {
	this := Pricing{}
	return &this
}

// NewPricingWithDefaults instantiates a new Pricing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricingWithDefaults() *Pricing {
	this := Pricing{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *Pricing) GetMarket() PricePriceMarket {
	if o == nil || IsNil(o.Market) {
		var ret PricePriceMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pricing) GetMarketOk() (*PricePriceMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *Pricing) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given PricePriceMarket and assigns it to the Market field.
func (o *Pricing) SetMarket(v PricePriceMarket) {
	o.Market = &v
}

// GetCurrentPrice returns the CurrentPrice field value if set, zero value otherwise.
func (o *Pricing) GetCurrentPrice() string {
	if o == nil || IsNil(o.CurrentPrice) {
		var ret string
		return ret
	}
	return *o.CurrentPrice
}

// GetCurrentPriceOk returns a tuple with the CurrentPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pricing) GetCurrentPriceOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentPrice) {
		return nil, false
	}
	return o.CurrentPrice, true
}

// HasCurrentPrice returns a boolean if a field has been set.
func (o *Pricing) HasCurrentPrice() bool {
	if o != nil && !IsNil(o.CurrentPrice) {
		return true
	}

	return false
}

// SetCurrentPrice gets a reference to the given string and assigns it to the CurrentPrice field.
func (o *Pricing) SetCurrentPrice(v string) {
	o.CurrentPrice = &v
}

// GetYesterdayPrice returns the YesterdayPrice field value if set, zero value otherwise.
func (o *Pricing) GetYesterdayPrice() string {
	if o == nil || IsNil(o.YesterdayPrice) {
		var ret string
		return ret
	}
	return *o.YesterdayPrice
}

// GetYesterdayPriceOk returns a tuple with the YesterdayPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pricing) GetYesterdayPriceOk() (*string, bool) {
	if o == nil || IsNil(o.YesterdayPrice) {
		return nil, false
	}
	return o.YesterdayPrice, true
}

// HasYesterdayPrice returns a boolean if a field has been set.
func (o *Pricing) HasYesterdayPrice() bool {
	if o != nil && !IsNil(o.YesterdayPrice) {
		return true
	}

	return false
}

// SetYesterdayPrice gets a reference to the given string and assigns it to the YesterdayPrice field.
func (o *Pricing) SetYesterdayPrice(v string) {
	o.YesterdayPrice = &v
}

// GetSnapshotAt returns the SnapshotAt field value if set, zero value otherwise.
func (o *Pricing) GetSnapshotAt() time.Time {
	if o == nil || IsNil(o.SnapshotAt) {
		var ret time.Time
		return ret
	}
	return *o.SnapshotAt
}

// GetSnapshotAtOk returns a tuple with the SnapshotAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pricing) GetSnapshotAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SnapshotAt) {
		return nil, false
	}
	return o.SnapshotAt, true
}

// HasSnapshotAt returns a boolean if a field has been set.
func (o *Pricing) HasSnapshotAt() bool {
	if o != nil && !IsNil(o.SnapshotAt) {
		return true
	}

	return false
}

// SetSnapshotAt gets a reference to the given time.Time and assigns it to the SnapshotAt field.
func (o *Pricing) SetSnapshotAt(v time.Time) {
	o.SnapshotAt = &v
}

func (o Pricing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pricing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.CurrentPrice) {
		toSerialize["current_price"] = o.CurrentPrice
	}
	if !IsNil(o.YesterdayPrice) {
		toSerialize["yesterday_price"] = o.YesterdayPrice
	}
	if !IsNil(o.SnapshotAt) {
		toSerialize["snapshot_at"] = o.SnapshotAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Pricing) UnmarshalJSON(data []byte) (err error) {
	varPricing := _Pricing{}

	err = json.Unmarshal(data, &varPricing)

	if err != nil {
		return err
	}

	*o = Pricing(varPricing)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "market")
		delete(additionalProperties, "current_price")
		delete(additionalProperties, "yesterday_price")
		delete(additionalProperties, "snapshot_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePricing struct {
	value *Pricing
	isSet bool
}

func (v NullablePricing) Get() *Pricing {
	return v.value
}

func (v *NullablePricing) Set(val *Pricing) {
	v.value = val
	v.isSet = true
}

func (v NullablePricing) IsSet() bool {
	return v.isSet
}

func (v *NullablePricing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricing(val *Pricing) *NullablePricing {
	return &NullablePricing{value: val, isSet: true}
}

func (v NullablePricing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


