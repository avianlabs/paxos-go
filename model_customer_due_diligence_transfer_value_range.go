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

// CustomerDueDiligenceTransferValueRange the model 'CustomerDueDiligenceTransferValueRange'
type CustomerDueDiligenceTransferValueRange string

// List of CustomerDueDiligenceTransferValueRange
const (
	CustomerDueDiligenceTransferValueRange__0_TO_25_K CustomerDueDiligenceTransferValueRange = "TRANSFER_VALUE_0_TO_25K"
	CustomerDueDiligenceTransferValueRange__25_K_TO_50_K CustomerDueDiligenceTransferValueRange = "TRANSFER_VALUE_25K_TO_50K"
	CustomerDueDiligenceTransferValueRange__50_K_TO_100_K CustomerDueDiligenceTransferValueRange = "TRANSFER_VALUE_50K_TO_100K"
	CustomerDueDiligenceTransferValueRange__100_K_TO_250_K CustomerDueDiligenceTransferValueRange = "TRANSFER_VALUE_100K_TO_250K"
	CustomerDueDiligenceTransferValueRange__250_K_TO_500_K CustomerDueDiligenceTransferValueRange = "TRANSFER_VALUE_250K_TO_500K"
	CustomerDueDiligenceTransferValueRange__500_K_TO_750_K CustomerDueDiligenceTransferValueRange = "TRANSFER_VALUE_500K_TO_750K"
	CustomerDueDiligenceTransferValueRange__750_K_TO_1_M CustomerDueDiligenceTransferValueRange = "TRANSFER_VALUE_750K_TO_1M"
	CustomerDueDiligenceTransferValueRange__1_M_TO_2_5_M CustomerDueDiligenceTransferValueRange = "TRANSFER_VALUE_1M_TO_2_5M"
	CustomerDueDiligenceTransferValueRange__2_5_M_TO_5_M CustomerDueDiligenceTransferValueRange = "TRANSFER_VALUE_2_5M_TO_5M"
	CustomerDueDiligenceTransferValueRange_ABOVE_5_M CustomerDueDiligenceTransferValueRange = "TRANSFER_VALUE_ABOVE_5M"
)

// All allowed values of CustomerDueDiligenceTransferValueRange enum
var AllowedCustomerDueDiligenceTransferValueRangeEnumValues = []CustomerDueDiligenceTransferValueRange{
	"TRANSFER_VALUE_0_TO_25K",
	"TRANSFER_VALUE_25K_TO_50K",
	"TRANSFER_VALUE_50K_TO_100K",
	"TRANSFER_VALUE_100K_TO_250K",
	"TRANSFER_VALUE_250K_TO_500K",
	"TRANSFER_VALUE_500K_TO_750K",
	"TRANSFER_VALUE_750K_TO_1M",
	"TRANSFER_VALUE_1M_TO_2_5M",
	"TRANSFER_VALUE_2_5M_TO_5M",
	"TRANSFER_VALUE_ABOVE_5M",
}

func (v *CustomerDueDiligenceTransferValueRange) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomerDueDiligenceTransferValueRange(value)
	for _, existing := range AllowedCustomerDueDiligenceTransferValueRangeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomerDueDiligenceTransferValueRange", value)
}

// NewCustomerDueDiligenceTransferValueRangeFromValue returns a pointer to a valid CustomerDueDiligenceTransferValueRange
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomerDueDiligenceTransferValueRangeFromValue(v string) (*CustomerDueDiligenceTransferValueRange, error) {
	ev := CustomerDueDiligenceTransferValueRange(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomerDueDiligenceTransferValueRange: valid values are %v", v, AllowedCustomerDueDiligenceTransferValueRangeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomerDueDiligenceTransferValueRange) IsValid() bool {
	for _, existing := range AllowedCustomerDueDiligenceTransferValueRangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomerDueDiligenceTransferValueRange value
func (v CustomerDueDiligenceTransferValueRange) Ptr() *CustomerDueDiligenceTransferValueRange {
	return &v
}

type NullableCustomerDueDiligenceTransferValueRange struct {
	value *CustomerDueDiligenceTransferValueRange
	isSet bool
}

func (v NullableCustomerDueDiligenceTransferValueRange) Get() *CustomerDueDiligenceTransferValueRange {
	return v.value
}

func (v *NullableCustomerDueDiligenceTransferValueRange) Set(val *CustomerDueDiligenceTransferValueRange) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerDueDiligenceTransferValueRange) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerDueDiligenceTransferValueRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerDueDiligenceTransferValueRange(val *CustomerDueDiligenceTransferValueRange) *NullableCustomerDueDiligenceTransferValueRange {
	return &NullableCustomerDueDiligenceTransferValueRange{value: val, isSet: true}
}

func (v NullableCustomerDueDiligenceTransferValueRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerDueDiligenceTransferValueRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

