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

// PersonDetailsCIPIDType the model 'PersonDetailsCIPIDType'
type PersonDetailsCIPIDType string

// List of PersonDetailsCIPIDType
const (
	PersonDetailsCIPIDType_SSN PersonDetailsCIPIDType = "SSN"
	PersonDetailsCIPIDType_ID_CARD PersonDetailsCIPIDType = "ID_CARD"
	PersonDetailsCIPIDType_ITIN PersonDetailsCIPIDType = "ITIN"
	PersonDetailsCIPIDType_PASSPORT PersonDetailsCIPIDType = "PASSPORT"
	PersonDetailsCIPIDType_DRIVING_LICENSE PersonDetailsCIPIDType = "DRIVING_LICENSE"
	PersonDetailsCIPIDType_VISA PersonDetailsCIPIDType = "VISA"
)

// All allowed values of PersonDetailsCIPIDType enum
var AllowedPersonDetailsCIPIDTypeEnumValues = []PersonDetailsCIPIDType{
	"SSN",
	"ID_CARD",
	"ITIN",
	"PASSPORT",
	"DRIVING_LICENSE",
	"VISA",
}

func (v *PersonDetailsCIPIDType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PersonDetailsCIPIDType(value)
	for _, existing := range AllowedPersonDetailsCIPIDTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PersonDetailsCIPIDType", value)
}

// NewPersonDetailsCIPIDTypeFromValue returns a pointer to a valid PersonDetailsCIPIDType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPersonDetailsCIPIDTypeFromValue(v string) (*PersonDetailsCIPIDType, error) {
	ev := PersonDetailsCIPIDType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PersonDetailsCIPIDType: valid values are %v", v, AllowedPersonDetailsCIPIDTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PersonDetailsCIPIDType) IsValid() bool {
	for _, existing := range AllowedPersonDetailsCIPIDTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PersonDetailsCIPIDType value
func (v PersonDetailsCIPIDType) Ptr() *PersonDetailsCIPIDType {
	return &v
}

type NullablePersonDetailsCIPIDType struct {
	value *PersonDetailsCIPIDType
	isSet bool
}

func (v NullablePersonDetailsCIPIDType) Get() *PersonDetailsCIPIDType {
	return v.value
}

func (v *NullablePersonDetailsCIPIDType) Set(val *PersonDetailsCIPIDType) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonDetailsCIPIDType) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonDetailsCIPIDType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonDetailsCIPIDType(val *PersonDetailsCIPIDType) *NullablePersonDetailsCIPIDType {
	return &NullablePersonDetailsCIPIDType{value: val, isSet: true}
}

func (v NullablePersonDetailsCIPIDType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonDetailsCIPIDType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

