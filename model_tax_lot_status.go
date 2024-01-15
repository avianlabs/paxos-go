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

// TaxLotStatus the model 'TaxLotStatus'
type TaxLotStatus string

// List of TaxLotStatus
const (
	TAXLOTSTATUS_OPEN TaxLotStatus = "OPEN"
	TAXLOTSTATUS_CLOSED TaxLotStatus = "CLOSED"
)

// All allowed values of TaxLotStatus enum
var AllowedTaxLotStatusEnumValues = []TaxLotStatus{
	"OPEN",
	"CLOSED",
}

func (v *TaxLotStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TaxLotStatus(value)
	for _, existing := range AllowedTaxLotStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TaxLotStatus", value)
}

// NewTaxLotStatusFromValue returns a pointer to a valid TaxLotStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaxLotStatusFromValue(v string) (*TaxLotStatus, error) {
	ev := TaxLotStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TaxLotStatus: valid values are %v", v, AllowedTaxLotStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaxLotStatus) IsValid() bool {
	for _, existing := range AllowedTaxLotStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaxLotStatus value
func (v TaxLotStatus) Ptr() *TaxLotStatus {
	return &v
}

type NullableTaxLotStatus struct {
	value *TaxLotStatus
	isSet bool
}

func (v NullableTaxLotStatus) Get() *TaxLotStatus {
	return v.value
}

func (v *NullableTaxLotStatus) Set(val *TaxLotStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxLotStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxLotStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxLotStatus(val *TaxLotStatus) *NullableTaxLotStatus {
	return &NullableTaxLotStatus{value: val, isSet: true}
}

func (v NullableTaxLotStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxLotStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

