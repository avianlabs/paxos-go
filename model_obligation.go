/*
Paxos API

<p>Welcome to Paxos APIs. At Paxos, our mission is to enable the movement of any asset, any time, in a trustworthy way. These APIs serve that mission by making it easier than ever for you to directly integrate our product capabilities into your application, leveraging the speed, stability, and security of the Paxos platform.</p> <p>The documentation that follows gives you access to our Crypto Brokerage, Trading, and Exchange products. It includes APIs for market data, orders, and the held rate quote flow.</p> <p>To test in our sandbox environment, <a href=\"https://account.sandbox.paxos.com\" target=\"_blank\">sign up</a> for an account. For more information about Paxos and our APIs, visit <a href=\"https://www.paxos.com/\" target=\"_blank\">Paxos.com</a>.</p> 

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paxos

import (
	"encoding/json"
	"fmt"
)

// checks if the Obligation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Obligation{}

// Obligation struct for Obligation
type Obligation struct {
	Id string `json:"id"`
	Direction ObligationDirection `json:"direction"`
	// Asset which is obliged.
	Asset string `json:"asset"`
	// Amount of the asset which is obliged.
	Amount string `json:"amount" validate:"regexp=^[0-9]*\\\\.?[0-9]+$"`
	AdditionalProperties map[string]interface{}
}

type _Obligation Obligation

// NewObligation instantiates a new Obligation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObligation(id string, direction ObligationDirection, asset string, amount string) *Obligation {
	this := Obligation{}
	this.Id = id
	this.Direction = direction
	this.Asset = asset
	this.Amount = amount
	return &this
}

// NewObligationWithDefaults instantiates a new Obligation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObligationWithDefaults() *Obligation {
	this := Obligation{}
	return &this
}

// GetId returns the Id field value
func (o *Obligation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Obligation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Obligation) SetId(v string) {
	o.Id = v
}

// GetDirection returns the Direction field value
func (o *Obligation) GetDirection() ObligationDirection {
	if o == nil {
		var ret ObligationDirection
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *Obligation) GetDirectionOk() (*ObligationDirection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *Obligation) SetDirection(v ObligationDirection) {
	o.Direction = v
}

// GetAsset returns the Asset field value
func (o *Obligation) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *Obligation) GetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *Obligation) SetAsset(v string) {
	o.Asset = v
}

// GetAmount returns the Amount field value
func (o *Obligation) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Obligation) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Obligation) SetAmount(v string) {
	o.Amount = v
}

func (o Obligation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Obligation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["direction"] = o.Direction
	toSerialize["asset"] = o.Asset
	toSerialize["amount"] = o.Amount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Obligation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"direction",
		"asset",
		"amount",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varObligation := _Obligation{}

	err = json.Unmarshal(data, &varObligation)

	if err != nil {
		return err
	}

	*o = Obligation(varObligation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "direction")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableObligation struct {
	value *Obligation
	isSet bool
}

func (v NullableObligation) Get() *Obligation {
	return v.value
}

func (v *NullableObligation) Set(val *Obligation) {
	v.value = val
	v.isSet = true
}

func (v NullableObligation) IsSet() bool {
	return v.isSet
}

func (v *NullableObligation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObligation(val *Obligation) *NullableObligation {
	return &NullableObligation{value: val, isSet: true}
}

func (v NullableObligation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObligation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


