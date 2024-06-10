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

// CustomerDueDiligenceNetWorthRange the model 'CustomerDueDiligenceNetWorthRange'
type CustomerDueDiligenceNetWorthRange string

// List of CustomerDueDiligenceNetWorthRange
const (
	_0_TO_100_K CustomerDueDiligenceNetWorthRange = "NET_WORTH_0_TO_100K"
	_100_K_TO_500_K CustomerDueDiligenceNetWorthRange = "NET_WORTH_100K_TO_500K"
	_500_K_TO_1_M CustomerDueDiligenceNetWorthRange = "NET_WORTH_500K_TO_1M"
	_1_M_TO_2_5_M CustomerDueDiligenceNetWorthRange = "NET_WORTH_1M_TO_2_5M"
	_2_5_M_TO_5_M CustomerDueDiligenceNetWorthRange = "NET_WORTH_2_5M_TO_5M"
	_5_M_TO_7_5_M CustomerDueDiligenceNetWorthRange = "NET_WORTH_5M_TO_7_5M"
	_7_5_M_TO_10_M CustomerDueDiligenceNetWorthRange = "NET_WORTH_7_5M_TO_10M"
	_10_M_TO_25_M CustomerDueDiligenceNetWorthRange = "NET_WORTH_10M_TO_25M"
	_25_M_TO_50_M CustomerDueDiligenceNetWorthRange = "NET_WORTH_25M_TO_50M"
	OVER_50_M CustomerDueDiligenceNetWorthRange = "NET_WORTH_OVER_50M"
)

// All allowed values of CustomerDueDiligenceNetWorthRange enum
var AllowedCustomerDueDiligenceNetWorthRangeEnumValues = []CustomerDueDiligenceNetWorthRange{
	"NET_WORTH_0_TO_100K",
	"NET_WORTH_100K_TO_500K",
	"NET_WORTH_500K_TO_1M",
	"NET_WORTH_1M_TO_2_5M",
	"NET_WORTH_2_5M_TO_5M",
	"NET_WORTH_5M_TO_7_5M",
	"NET_WORTH_7_5M_TO_10M",
	"NET_WORTH_10M_TO_25M",
	"NET_WORTH_25M_TO_50M",
	"NET_WORTH_OVER_50M",
}

func (v *CustomerDueDiligenceNetWorthRange) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomerDueDiligenceNetWorthRange(value)
	for _, existing := range AllowedCustomerDueDiligenceNetWorthRangeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomerDueDiligenceNetWorthRange", value)
}

// NewCustomerDueDiligenceNetWorthRangeFromValue returns a pointer to a valid CustomerDueDiligenceNetWorthRange
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomerDueDiligenceNetWorthRangeFromValue(v string) (*CustomerDueDiligenceNetWorthRange, error) {
	ev := CustomerDueDiligenceNetWorthRange(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomerDueDiligenceNetWorthRange: valid values are %v", v, AllowedCustomerDueDiligenceNetWorthRangeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomerDueDiligenceNetWorthRange) IsValid() bool {
	for _, existing := range AllowedCustomerDueDiligenceNetWorthRangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomerDueDiligenceNetWorthRange value
func (v CustomerDueDiligenceNetWorthRange) Ptr() *CustomerDueDiligenceNetWorthRange {
	return &v
}

type NullableCustomerDueDiligenceNetWorthRange struct {
	value *CustomerDueDiligenceNetWorthRange
	isSet bool
}

func (v NullableCustomerDueDiligenceNetWorthRange) Get() *CustomerDueDiligenceNetWorthRange {
	return v.value
}

func (v *NullableCustomerDueDiligenceNetWorthRange) Set(val *CustomerDueDiligenceNetWorthRange) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerDueDiligenceNetWorthRange) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerDueDiligenceNetWorthRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerDueDiligenceNetWorthRange(val *CustomerDueDiligenceNetWorthRange) *NullableCustomerDueDiligenceNetWorthRange {
	return &NullableCustomerDueDiligenceNetWorthRange{value: val, isSet: true}
}

func (v NullableCustomerDueDiligenceNetWorthRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerDueDiligenceNetWorthRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
