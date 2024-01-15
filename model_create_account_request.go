/*
Paxos API

<p>Welcome to Paxos APIs. At Paxos, our mission is to enable the movement of any asset, any time, in a trustworthy way. These APIs serve that mission by making it easier than ever for you to directly integrate our product capabilities into your application, leveraging the speed, stability, and security of the Paxos platform.</p> <p>The documentation that follows gives you access to our Crypto Brokerage, Trading, and Exchange products. It includes APIs for market data, orders, and the held rate quote flow.</p> <p>To test in our sandbox environment, <a href=\"https://account.sandbox.paxos.com\" target=\"_blank\">sign up</a> for an account. For more information about Paxos and our APIs, visit <a href=\"https://www.paxos.com/\" target=\"_blank\">Paxos.com</a>.</p> 

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paxos

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateAccountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccountRequest{}

// CreateAccountRequest struct for CreateAccountRequest
type CreateAccountRequest struct {
	Account Account `json:"account"`
	// Create a new profile for this account. Set to `true` to track user balances at the Profile level for this account. See also [Profiles](#tag/Profiles).
	CreateProfile *bool `json:"create_profile,omitempty"`
}

type _CreateAccountRequest CreateAccountRequest

// NewCreateAccountRequest instantiates a new CreateAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountRequest(account Account) *CreateAccountRequest {
	this := CreateAccountRequest{}
	this.Account = account
	return &this
}

// NewCreateAccountRequestWithDefaults instantiates a new CreateAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountRequestWithDefaults() *CreateAccountRequest {
	this := CreateAccountRequest{}
	return &this
}

// GetAccount returns the Account field value
func (o *CreateAccountRequest) GetAccount() Account {
	if o == nil {
		var ret Account
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetAccountOk() (*Account, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *CreateAccountRequest) SetAccount(v Account) {
	o.Account = v
}

// GetCreateProfile returns the CreateProfile field value if set, zero value otherwise.
func (o *CreateAccountRequest) GetCreateProfile() bool {
	if o == nil || IsNil(o.CreateProfile) {
		var ret bool
		return ret
	}
	return *o.CreateProfile
}

// GetCreateProfileOk returns a tuple with the CreateProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountRequest) GetCreateProfileOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateProfile) {
		return nil, false
	}
	return o.CreateProfile, true
}

// HasCreateProfile returns a boolean if a field has been set.
func (o *CreateAccountRequest) HasCreateProfile() bool {
	if o != nil && !IsNil(o.CreateProfile) {
		return true
	}

	return false
}

// SetCreateProfile gets a reference to the given bool and assigns it to the CreateProfile field.
func (o *CreateAccountRequest) SetCreateProfile(v bool) {
	o.CreateProfile = &v
}

func (o CreateAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account"] = o.Account
	if !IsNil(o.CreateProfile) {
		toSerialize["create_profile"] = o.CreateProfile
	}
	return toSerialize, nil
}

func (o *CreateAccountRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account",
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

	varCreateAccountRequest := _CreateAccountRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateAccountRequest)

	if err != nil {
		return err
	}

	*o = CreateAccountRequest(varCreateAccountRequest)

	return err
}

type NullableCreateAccountRequest struct {
	value *CreateAccountRequest
	isSet bool
}

func (v NullableCreateAccountRequest) Get() *CreateAccountRequest {
	return v.value
}

func (v *NullableCreateAccountRequest) Set(val *CreateAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountRequest(val *CreateAccountRequest) *NullableCreateAccountRequest {
	return &NullableCreateAccountRequest{value: val, isSet: true}
}

func (v NullableCreateAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


