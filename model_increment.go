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

// Increment the model 'Increment'
type Increment string

// List of Increment
const (
	ONE_MINUTE Increment = "ONE_MINUTE"
	FIVE_MINUTES Increment = "FIVE_MINUTES"
	FIFTEEN_MINUTES Increment = "FIFTEEN_MINUTES"
	THIRTY_MINUTES Increment = "THIRTY_MINUTES"
	ONE_HOUR Increment = "ONE_HOUR"
	TWO_HOURS Increment = "TWO_HOURS"
	TWELVE_HOURS Increment = "TWELVE_HOURS"
	ONE_DAY Increment = "ONE_DAY"
	ONE_WEEK Increment = "ONE_WEEK"
	TWO_WEEKS Increment = "TWO_WEEKS"
	FOUR_WEEKS Increment = "FOUR_WEEKS"
)

// All allowed values of Increment enum
var AllowedIncrementEnumValues = []Increment{
	"ONE_MINUTE",
	"FIVE_MINUTES",
	"FIFTEEN_MINUTES",
	"THIRTY_MINUTES",
	"ONE_HOUR",
	"TWO_HOURS",
	"TWELVE_HOURS",
	"ONE_DAY",
	"ONE_WEEK",
	"TWO_WEEKS",
	"FOUR_WEEKS",
}

func (v *Increment) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Increment(value)
	for _, existing := range AllowedIncrementEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Increment", value)
}

// NewIncrementFromValue returns a pointer to a valid Increment
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIncrementFromValue(v string) (*Increment, error) {
	ev := Increment(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Increment: valid values are %v", v, AllowedIncrementEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Increment) IsValid() bool {
	for _, existing := range AllowedIncrementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Increment value
func (v Increment) Ptr() *Increment {
	return &v
}

type NullableIncrement struct {
	value *Increment
	isSet bool
}

func (v NullableIncrement) Get() *Increment {
	return v.value
}

func (v *NullableIncrement) Set(val *Increment) {
	v.value = val
	v.isSet = true
}

func (v NullableIncrement) IsSet() bool {
	return v.isSet
}

func (v *NullableIncrement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncrement(val *Increment) *NullableIncrement {
	return &NullableIncrement{value: val, isSet: true}
}

func (v NullableIncrement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncrement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

