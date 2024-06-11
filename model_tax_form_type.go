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

// TaxFormType the model 'TaxFormType'
type TaxFormType string

// List of TaxFormType
const (
	TaxFormType_B TaxFormType = "FORM_1099_B"
	TaxFormType_MISC TaxFormType = "FORM_1099_MISC"
)

// All allowed values of TaxFormType enum
var AllowedTaxFormTypeEnumValues = []TaxFormType{
	"FORM_1099_B",
	"FORM_1099_MISC",
}

func (v *TaxFormType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TaxFormType(value)
	for _, existing := range AllowedTaxFormTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TaxFormType", value)
}

// NewTaxFormTypeFromValue returns a pointer to a valid TaxFormType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaxFormTypeFromValue(v string) (*TaxFormType, error) {
	ev := TaxFormType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TaxFormType: valid values are %v", v, AllowedTaxFormTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaxFormType) IsValid() bool {
	for _, existing := range AllowedTaxFormTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaxFormType value
func (v TaxFormType) Ptr() *TaxFormType {
	return &v
}

type NullableTaxFormType struct {
	value *TaxFormType
	isSet bool
}

func (v NullableTaxFormType) Get() *TaxFormType {
	return v.value
}

func (v *NullableTaxFormType) Set(val *TaxFormType) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxFormType) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxFormType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxFormType(val *TaxFormType) *NullableTaxFormType {
	return &NullableTaxFormType{value: val, isSet: true}
}

func (v NullableTaxFormType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxFormType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

