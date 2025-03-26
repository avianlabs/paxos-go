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

// checks if the AutoConversion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoConversion{}

// AutoConversion struct for AutoConversion
type AutoConversion struct {
	FromTransferId *string `json:"from_transfer_id,omitempty"`
	ToTransferId *string `json:"to_transfer_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AutoConversion AutoConversion

// NewAutoConversion instantiates a new AutoConversion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoConversion() *AutoConversion {
	this := AutoConversion{}
	return &this
}

// NewAutoConversionWithDefaults instantiates a new AutoConversion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoConversionWithDefaults() *AutoConversion {
	this := AutoConversion{}
	return &this
}

// GetFromTransferId returns the FromTransferId field value if set, zero value otherwise.
func (o *AutoConversion) GetFromTransferId() string {
	if o == nil || IsNil(o.FromTransferId) {
		var ret string
		return ret
	}
	return *o.FromTransferId
}

// GetFromTransferIdOk returns a tuple with the FromTransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoConversion) GetFromTransferIdOk() (*string, bool) {
	if o == nil || IsNil(o.FromTransferId) {
		return nil, false
	}
	return o.FromTransferId, true
}

// HasFromTransferId returns a boolean if a field has been set.
func (o *AutoConversion) HasFromTransferId() bool {
	if o != nil && !IsNil(o.FromTransferId) {
		return true
	}

	return false
}

// SetFromTransferId gets a reference to the given string and assigns it to the FromTransferId field.
func (o *AutoConversion) SetFromTransferId(v string) {
	o.FromTransferId = &v
}

// GetToTransferId returns the ToTransferId field value if set, zero value otherwise.
func (o *AutoConversion) GetToTransferId() string {
	if o == nil || IsNil(o.ToTransferId) {
		var ret string
		return ret
	}
	return *o.ToTransferId
}

// GetToTransferIdOk returns a tuple with the ToTransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoConversion) GetToTransferIdOk() (*string, bool) {
	if o == nil || IsNil(o.ToTransferId) {
		return nil, false
	}
	return o.ToTransferId, true
}

// HasToTransferId returns a boolean if a field has been set.
func (o *AutoConversion) HasToTransferId() bool {
	if o != nil && !IsNil(o.ToTransferId) {
		return true
	}

	return false
}

// SetToTransferId gets a reference to the given string and assigns it to the ToTransferId field.
func (o *AutoConversion) SetToTransferId(v string) {
	o.ToTransferId = &v
}

func (o AutoConversion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoConversion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FromTransferId) {
		toSerialize["from_transfer_id"] = o.FromTransferId
	}
	if !IsNil(o.ToTransferId) {
		toSerialize["to_transfer_id"] = o.ToTransferId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AutoConversion) UnmarshalJSON(data []byte) (err error) {
	varAutoConversion := _AutoConversion{}

	err = json.Unmarshal(data, &varAutoConversion)

	if err != nil {
		return err
	}

	*o = AutoConversion(varAutoConversion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "from_transfer_id")
		delete(additionalProperties, "to_transfer_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAutoConversion struct {
	value *AutoConversion
	isSet bool
}

func (v NullableAutoConversion) Get() *AutoConversion {
	return v.value
}

func (v *NullableAutoConversion) Set(val *AutoConversion) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoConversion) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoConversion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoConversion(val *AutoConversion) *NullableAutoConversion {
	return &NullableAutoConversion{value: val, isSet: true}
}

func (v NullableAutoConversion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoConversion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


