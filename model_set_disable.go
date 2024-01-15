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

// checks if the SetDisable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetDisable{}

// SetDisable struct for SetDisable
type SetDisable struct {
	Disabled *bool `json:"disabled,omitempty"`
}

// NewSetDisable instantiates a new SetDisable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetDisable() *SetDisable {
	this := SetDisable{}
	return &this
}

// NewSetDisableWithDefaults instantiates a new SetDisable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetDisableWithDefaults() *SetDisable {
	this := SetDisable{}
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *SetDisable) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetDisable) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *SetDisable) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *SetDisable) SetDisabled(v bool) {
	o.Disabled = &v
}

func (o SetDisable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetDisable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	return toSerialize, nil
}

type NullableSetDisable struct {
	value *SetDisable
	isSet bool
}

func (v NullableSetDisable) Get() *SetDisable {
	return v.value
}

func (v *NullableSetDisable) Set(val *SetDisable) {
	v.value = val
	v.isSet = true
}

func (v NullableSetDisable) IsSet() bool {
	return v.isSet
}

func (v *NullableSetDisable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetDisable(val *SetDisable) *NullableSetDisable {
	return &NullableSetDisable{value: val, isSet: true}
}

func (v NullableSetDisable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetDisable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


