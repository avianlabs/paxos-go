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

// TimeInForce How long an order will remain active before it expires. - Immediate-or-Cancel (IOC): Cancel if not executed immediately, partial fills allowed. - Good-Til-Canceled (GTC): Order can be canceled at any point until executed. - Good-Til-Time (GTT): Expires if not executed by a specified time. GTT must be greater than 10 seconds after the order is placed, otherwise the order will be rejected. - Fill-or-Kill (FOK): Fill entire order only or cancel entire order, does not allow for partial filling.  **Time in Force validity for Order Types**  | Order type      | Immediate or Cancel (IOC) | Good Til Canceled (GTC) | Good Til Time (GTT) | Fill or Kill (FOK) | | --------------- | ------------------------- | ----------------------- | ------------------- | ------------------ | | Market Order    | Default                   |  -                      |  -                  |  -                 | | Limit Order     | Valid                     | Default                 | Valid               | Valid              | | Post Only Limit | Valid                     | Default                 | Valid               | Valid              | | Stop Market     |  -                        | Default                 | Valid               |  -                 | | Stop Limit      |  -                        | Default                 | Valid               |  -                 |
type TimeInForce string

// List of TimeInForce
const (
	GTC_TimeInForce TimeInForce = "GTC"
	FOK_TimeInForce TimeInForce = "FOK"
	IOC_TimeInForce TimeInForce = "IOC"
	GTT_TimeInForce TimeInForce = "GTT"
)

// All allowed values of TimeInForce enum
var AllowedTimeInForceEnumValues = []TimeInForce{
	"GTC",
	"FOK",
	"IOC",
	"GTT",
}

func (v *TimeInForce) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TimeInForce(value)
	for _, existing := range AllowedTimeInForceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TimeInForce", value)
}

// NewTimeInForceFromValue returns a pointer to a valid TimeInForce
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTimeInForceFromValue(v string) (*TimeInForce, error) {
	ev := TimeInForce(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TimeInForce: valid values are %v", v, AllowedTimeInForceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TimeInForce) IsValid() bool {
	for _, existing := range AllowedTimeInForceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeInForce value
func (v TimeInForce) Ptr() *TimeInForce {
	return &v
}

type NullableTimeInForce struct {
	value *TimeInForce
	isSet bool
}

func (v NullableTimeInForce) Get() *TimeInForce {
	return v.value
}

func (v *NullableTimeInForce) Set(val *TimeInForce) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeInForce) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeInForce) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeInForce(val *TimeInForce) *NullableTimeInForce {
	return &NullableTimeInForce{value: val, isSet: true}
}

func (v NullableTimeInForce) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeInForce) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

