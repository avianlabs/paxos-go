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

// AccountAccountType the model 'AccountAccountType'
type AccountAccountType string

// List of AccountAccountType
const (
	AccountAccountType_BROKERAGE AccountAccountType = "BROKERAGE"
	AccountAccountType_TRADITIONAL_IRA AccountAccountType = "TRADITIONAL_IRA"
	AccountAccountType_ROTH_IRA AccountAccountType = "ROTH_IRA"
	AccountAccountType_SEP_IRA AccountAccountType = "SEP_IRA"
	AccountAccountType_FINANCIAL_ADVISOR AccountAccountType = "FINANCIAL_ADVISOR"
)

// All allowed values of AccountAccountType enum
var AllowedAccountAccountTypeEnumValues = []AccountAccountType{
	"BROKERAGE",
	"TRADITIONAL_IRA",
	"ROTH_IRA",
	"SEP_IRA",
	"FINANCIAL_ADVISOR",
}

func (v *AccountAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountAccountType(value)
	for _, existing := range AllowedAccountAccountTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountAccountType", value)
}

// NewAccountAccountTypeFromValue returns a pointer to a valid AccountAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountAccountTypeFromValue(v string) (*AccountAccountType, error) {
	ev := AccountAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountAccountType: valid values are %v", v, AllowedAccountAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountAccountType) IsValid() bool {
	for _, existing := range AllowedAccountAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountAccountType value
func (v AccountAccountType) Ptr() *AccountAccountType {
	return &v
}

type NullableAccountAccountType struct {
	value *AccountAccountType
	isSet bool
}

func (v NullableAccountAccountType) Get() *AccountAccountType {
	return v.value
}

func (v *NullableAccountAccountType) Set(val *AccountAccountType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAccountType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAccountType(val *AccountAccountType) *NullableAccountAccountType {
	return &NullableAccountAccountType{value: val, isSet: true}
}

func (v NullableAccountAccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

