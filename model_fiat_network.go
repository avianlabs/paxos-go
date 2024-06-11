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

// FiatNetwork the model 'FiatNetwork'
type FiatNetwork string

// List of FiatNetwork
const (
	FiatNetwork_WIRE FiatNetwork = "WIRE"
	FiatNetwork_CBIT FiatNetwork = "CBIT"
)

// All allowed values of FiatNetwork enum
var AllowedFiatNetworkEnumValues = []FiatNetwork{
	"WIRE",
	"CBIT",
}

func (v *FiatNetwork) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FiatNetwork(value)
	for _, existing := range AllowedFiatNetworkEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FiatNetwork", value)
}

// NewFiatNetworkFromValue returns a pointer to a valid FiatNetwork
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFiatNetworkFromValue(v string) (*FiatNetwork, error) {
	ev := FiatNetwork(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FiatNetwork: valid values are %v", v, AllowedFiatNetworkEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FiatNetwork) IsValid() bool {
	for _, existing := range AllowedFiatNetworkEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FiatNetwork value
func (v FiatNetwork) Ptr() *FiatNetwork {
	return &v
}

type NullableFiatNetwork struct {
	value *FiatNetwork
	isSet bool
}

func (v NullableFiatNetwork) Get() *FiatNetwork {
	return v.value
}

func (v *NullableFiatNetwork) Set(val *FiatNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableFiatNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableFiatNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiatNetwork(val *FiatNetwork) *NullableFiatNetwork {
	return &NullableFiatNetwork{value: val, isSet: true}
}

func (v NullableFiatNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiatNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

