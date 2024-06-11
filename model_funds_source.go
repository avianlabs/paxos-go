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

// FundsSource the model 'FundsSource'
type FundsSource string

// List of FundsSource
const (
	FundsSource_SALARY_DISBURSEMENT FundsSource = "SALARY_DISBURSEMENT"
	FundsSource_INHERITANCE_DISTRIBUTION FundsSource = "INHERITANCE_DISTRIBUTION"
	FundsSource_INVESTMENT_RETURNS FundsSource = "INVESTMENT_RETURNS"
	FundsSource_BUSINESS_DIVIDENDS_PROFITS FundsSource = "BUSINESS_DIVIDENDS_PROFITS"
	FundsSource_PROPERTY_SALE FundsSource = "PROPERTY_SALE"
	FundsSource_LOAN_DISBURSEMENT FundsSource = "LOAN_DISBURSEMENT"
	FundsSource_SAVINGS_ACCOUNT_WITHDRAWAL FundsSource = "SAVINGS_ACCOUNT_WITHDRAWAL"
)

// All allowed values of FundsSource enum
var AllowedFundsSourceEnumValues = []FundsSource{
	"SALARY_DISBURSEMENT",
	"INHERITANCE_DISTRIBUTION",
	"INVESTMENT_RETURNS",
	"BUSINESS_DIVIDENDS_PROFITS",
	"PROPERTY_SALE",
	"LOAN_DISBURSEMENT",
	"SAVINGS_ACCOUNT_WITHDRAWAL",
}

func (v *FundsSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FundsSource(value)
	for _, existing := range AllowedFundsSourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FundsSource", value)
}

// NewFundsSourceFromValue returns a pointer to a valid FundsSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFundsSourceFromValue(v string) (*FundsSource, error) {
	ev := FundsSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FundsSource: valid values are %v", v, AllowedFundsSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FundsSource) IsValid() bool {
	for _, existing := range AllowedFundsSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FundsSource value
func (v FundsSource) Ptr() *FundsSource {
	return &v
}

type NullableFundsSource struct {
	value *FundsSource
	isSet bool
}

func (v NullableFundsSource) Get() *FundsSource {
	return v.value
}

func (v *NullableFundsSource) Set(val *FundsSource) {
	v.value = val
	v.isSet = true
}

func (v NullableFundsSource) IsSet() bool {
	return v.isSet
}

func (v *NullableFundsSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundsSource(val *FundsSource) *NullableFundsSource {
	return &NullableFundsSource{value: val, isSet: true}
}

func (v NullableFundsSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundsSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

