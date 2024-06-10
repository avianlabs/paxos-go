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

// AccountPurpose the model 'AccountPurpose'
type AccountPurpose string

// List of AccountPurpose
const (
	INVESTMENT_TRADING AccountPurpose = "INVESTMENT_TRADING"
	SAVINGS AccountPurpose = "SAVINGS"
	STABLECOIN_PURCHASE_REDEMPTION AccountPurpose = "STABLECOIN_PURCHASE_REDEMPTION"
)

// All allowed values of AccountPurpose enum
var AllowedAccountPurposeEnumValues = []AccountPurpose{
	"INVESTMENT_TRADING",
	"SAVINGS",
	"STABLECOIN_PURCHASE_REDEMPTION",
}

func (v *AccountPurpose) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountPurpose(value)
	for _, existing := range AllowedAccountPurposeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountPurpose", value)
}

// NewAccountPurposeFromValue returns a pointer to a valid AccountPurpose
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountPurposeFromValue(v string) (*AccountPurpose, error) {
	ev := AccountPurpose(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountPurpose: valid values are %v", v, AllowedAccountPurposeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountPurpose) IsValid() bool {
	for _, existing := range AllowedAccountPurposeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountPurpose value
func (v AccountPurpose) Ptr() *AccountPurpose {
	return &v
}

type NullableAccountPurpose struct {
	value *AccountPurpose
	isSet bool
}

func (v NullableAccountPurpose) Get() *AccountPurpose {
	return v.value
}

func (v *NullableAccountPurpose) Set(val *AccountPurpose) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountPurpose) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountPurpose) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountPurpose(val *AccountPurpose) *NullableAccountPurpose {
	return &NullableAccountPurpose{value: val, isSet: true}
}

func (v NullableAccountPurpose) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountPurpose) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
