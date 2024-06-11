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

// WealthSource the model 'WealthSource'
type WealthSource string

// List of WealthSource
const (
	WealthSource_INHERITANCE WealthSource = "INHERITANCE"
	WealthSource_INVESTMENT_GAINS WealthSource = "INVESTMENT_GAINS"
	WealthSource_BUSINESS_OWNERSHIP_DIVIDENDS WealthSource = "BUSINESS_OWNERSHIP_DIVIDENDS"
	WealthSource_EMPLOYMENT_INCOME WealthSource = "EMPLOYMENT_INCOME"
)

// All allowed values of WealthSource enum
var AllowedWealthSourceEnumValues = []WealthSource{
	"INHERITANCE",
	"INVESTMENT_GAINS",
	"BUSINESS_OWNERSHIP_DIVIDENDS",
	"EMPLOYMENT_INCOME",
}

func (v *WealthSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WealthSource(value)
	for _, existing := range AllowedWealthSourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WealthSource", value)
}

// NewWealthSourceFromValue returns a pointer to a valid WealthSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWealthSourceFromValue(v string) (*WealthSource, error) {
	ev := WealthSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WealthSource: valid values are %v", v, AllowedWealthSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WealthSource) IsValid() bool {
	for _, existing := range AllowedWealthSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WealthSource value
func (v WealthSource) Ptr() *WealthSource {
	return &v
}

type NullableWealthSource struct {
	value *WealthSource
	isSet bool
}

func (v NullableWealthSource) Get() *WealthSource {
	return v.value
}

func (v *NullableWealthSource) Set(val *WealthSource) {
	v.value = val
	v.isSet = true
}

func (v NullableWealthSource) IsSet() bool {
	return v.isSet
}

func (v *NullableWealthSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWealthSource(val *WealthSource) *NullableWealthSource {
	return &NullableWealthSource{value: val, isSet: true}
}

func (v NullableWealthSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWealthSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

