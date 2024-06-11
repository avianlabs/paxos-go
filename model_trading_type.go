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

// TradingType the model 'TradingType'
type TradingType string

// List of TradingType
const (
	TradingType_PRIVATE TradingType = "PRIVATE"
	TradingType_PUBLIC TradingType = "PUBLIC"
	TradingType_PUBLICLY_TRADED_SUBSIDIARY TradingType = "PUBLICLY_TRADED_SUBSIDIARY"
)

// All allowed values of TradingType enum
var AllowedTradingTypeEnumValues = []TradingType{
	"PRIVATE",
	"PUBLIC",
	"PUBLICLY_TRADED_SUBSIDIARY",
}

func (v *TradingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TradingType(value)
	for _, existing := range AllowedTradingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TradingType", value)
}

// NewTradingTypeFromValue returns a pointer to a valid TradingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTradingTypeFromValue(v string) (*TradingType, error) {
	ev := TradingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TradingType: valid values are %v", v, AllowedTradingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TradingType) IsValid() bool {
	for _, existing := range AllowedTradingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TradingType value
func (v TradingType) Ptr() *TradingType {
	return &v
}

type NullableTradingType struct {
	value *TradingType
	isSet bool
}

func (v NullableTradingType) Get() *TradingType {
	return v.value
}

func (v *NullableTradingType) Set(val *TradingType) {
	v.value = val
	v.isSet = true
}

func (v NullableTradingType) IsSet() bool {
	return v.isSet
}

func (v *NullableTradingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradingType(val *TradingType) *NullableTradingType {
	return &NullableTradingType{value: val, isSet: true}
}

func (v NullableTradingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

