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

// OrderType Trade type.
type OrderType string

// List of OrderType
const (
	LIMIT_OrderType OrderType = "LIMIT"
	MARKET_OrderType OrderType = "MARKET"
	POST_ONLY_LIMIT_OrderType OrderType = "POST_ONLY_LIMIT"
	STOP_MARKET_OrderType OrderType = "STOP_MARKET"
	STOP_LIMIT_OrderType OrderType = "STOP_LIMIT"
)

// All allowed values of OrderType enum
var AllowedOrderTypeEnumValues = []OrderType{
	"LIMIT",
	"MARKET",
	"POST_ONLY_LIMIT",
	"STOP_MARKET",
	"STOP_LIMIT",
}

func (v *OrderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderType(value)
	for _, existing := range AllowedOrderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderType", value)
}

// NewOrderTypeFromValue returns a pointer to a valid OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderTypeFromValue(v string) (*OrderType, error) {
	ev := OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderType: valid values are %v", v, AllowedOrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderType) IsValid() bool {
	for _, existing := range AllowedOrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderType value
func (v OrderType) Ptr() *OrderType {
	return &v
}

type NullableOrderType struct {
	value *OrderType
	isSet bool
}

func (v NullableOrderType) Get() *OrderType {
	return v.value
}

func (v *NullableOrderType) Set(val *OrderType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderType(val *OrderType) *NullableOrderType {
	return &NullableOrderType{value: val, isSet: true}
}

func (v NullableOrderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

