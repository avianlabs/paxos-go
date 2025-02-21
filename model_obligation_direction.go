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

// ObligationDirection Direction of asset delivery for an obligation.   - DELIVER: Assets are delivered from the source profile to the target profile.  - RECEIVE: Assets are delivered from the target profile to the source profile.
type ObligationDirection string

// List of ObligationDirection
const (
	ObligationDirection_DELIVER ObligationDirection = "DELIVER"
	ObligationDirection_RECEIVE ObligationDirection = "RECEIVE"
)

// All allowed values of ObligationDirection enum
var AllowedObligationDirectionEnumValues = []ObligationDirection{
	"DELIVER",
	"RECEIVE",
}

func (v *ObligationDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ObligationDirection(value)
	for _, existing := range AllowedObligationDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ObligationDirection", value)
}

// NewObligationDirectionFromValue returns a pointer to a valid ObligationDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewObligationDirectionFromValue(v string) (*ObligationDirection, error) {
	ev := ObligationDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ObligationDirection: valid values are %v", v, AllowedObligationDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ObligationDirection) IsValid() bool {
	for _, existing := range AllowedObligationDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObligationDirection value
func (v ObligationDirection) Ptr() *ObligationDirection {
	return &v
}

type NullableObligationDirection struct {
	value *ObligationDirection
	isSet bool
}

func (v NullableObligationDirection) Get() *ObligationDirection {
	return v.value
}

func (v *NullableObligationDirection) Set(val *ObligationDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableObligationDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableObligationDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObligationDirection(val *ObligationDirection) *NullableObligationDirection {
	return &NullableObligationDirection{value: val, isSet: true}
}

func (v NullableObligationDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObligationDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

