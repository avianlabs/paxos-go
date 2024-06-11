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

// TaxprotoOrderBy the model 'TaxprotoOrderBy'
type TaxprotoOrderBy string

// List of taxprotoOrderBy
const (
	TaxprotoOrderBy_CREATED_AT TaxprotoOrderBy = "CREATED_AT"
)

// All allowed values of TaxprotoOrderBy enum
var AllowedTaxprotoOrderByEnumValues = []TaxprotoOrderBy{
	"CREATED_AT",
}

func (v *TaxprotoOrderBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TaxprotoOrderBy(value)
	for _, existing := range AllowedTaxprotoOrderByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TaxprotoOrderBy", value)
}

// NewTaxprotoOrderByFromValue returns a pointer to a valid TaxprotoOrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaxprotoOrderByFromValue(v string) (*TaxprotoOrderBy, error) {
	ev := TaxprotoOrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TaxprotoOrderBy: valid values are %v", v, AllowedTaxprotoOrderByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaxprotoOrderBy) IsValid() bool {
	for _, existing := range AllowedTaxprotoOrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to taxprotoOrderBy value
func (v TaxprotoOrderBy) Ptr() *TaxprotoOrderBy {
	return &v
}

type NullableTaxprotoOrderBy struct {
	value *TaxprotoOrderBy
	isSet bool
}

func (v NullableTaxprotoOrderBy) Get() *TaxprotoOrderBy {
	return v.value
}

func (v *NullableTaxprotoOrderBy) Set(val *TaxprotoOrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxprotoOrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxprotoOrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxprotoOrderBy(val *TaxprotoOrderBy) *NullableTaxprotoOrderBy {
	return &NullableTaxprotoOrderBy{value: val, isSet: true}
}

func (v NullableTaxprotoOrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxprotoOrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

