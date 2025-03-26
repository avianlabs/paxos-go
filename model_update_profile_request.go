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

// checks if the UpdateProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProfileRequest{}

// UpdateProfileRequest struct for UpdateProfileRequest
type UpdateProfileRequest struct {
	Nickname string `json:"nickname"`
	AdditionalProperties map[string]interface{}
}

type _UpdateProfileRequest UpdateProfileRequest

// NewUpdateProfileRequest instantiates a new UpdateProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProfileRequest(nickname string) *UpdateProfileRequest {
	this := UpdateProfileRequest{}
	this.Nickname = nickname
	return &this
}

// NewUpdateProfileRequestWithDefaults instantiates a new UpdateProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProfileRequestWithDefaults() *UpdateProfileRequest {
	this := UpdateProfileRequest{}
	return &this
}

// GetNickname returns the Nickname field value
func (o *UpdateProfileRequest) GetNickname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value
// and a boolean to check if the value has been set.
func (o *UpdateProfileRequest) GetNicknameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nickname, true
}

// SetNickname sets field value
func (o *UpdateProfileRequest) SetNickname(v string) {
	o.Nickname = v
}

func (o UpdateProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nickname"] = o.Nickname

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateProfileRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nickname",
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

	varUpdateProfileRequest := _UpdateProfileRequest{}

	err = json.Unmarshal(data, &varUpdateProfileRequest)

	if err != nil {
		return err
	}

	*o = UpdateProfileRequest(varUpdateProfileRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "nickname")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateProfileRequest struct {
	value *UpdateProfileRequest
	isSet bool
}

func (v NullableUpdateProfileRequest) Get() *UpdateProfileRequest {
	return v.value
}

func (v *NullableUpdateProfileRequest) Set(val *UpdateProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProfileRequest(val *UpdateProfileRequest) *NullableUpdateProfileRequest {
	return &NullableUpdateProfileRequest{value: val, isSet: true}
}

func (v NullableUpdateProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


