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

// ListTaxFormsRequestOrderBy the model 'ListTaxFormsRequestOrderBy'
type ListTaxFormsRequestOrderBy string

// List of ListTaxFormsRequestOrderBy
const (
	ListTaxFormsRequestOrderBy_ID ListTaxFormsRequestOrderBy = "ID"
)

// All allowed values of ListTaxFormsRequestOrderBy enum
var AllowedListTaxFormsRequestOrderByEnumValues = []ListTaxFormsRequestOrderBy{
	"ID",
}

func (v *ListTaxFormsRequestOrderBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ListTaxFormsRequestOrderBy(value)
	for _, existing := range AllowedListTaxFormsRequestOrderByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ListTaxFormsRequestOrderBy", value)
}

// NewListTaxFormsRequestOrderByFromValue returns a pointer to a valid ListTaxFormsRequestOrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListTaxFormsRequestOrderByFromValue(v string) (*ListTaxFormsRequestOrderBy, error) {
	ev := ListTaxFormsRequestOrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListTaxFormsRequestOrderBy: valid values are %v", v, AllowedListTaxFormsRequestOrderByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListTaxFormsRequestOrderBy) IsValid() bool {
	for _, existing := range AllowedListTaxFormsRequestOrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListTaxFormsRequestOrderBy value
func (v ListTaxFormsRequestOrderBy) Ptr() *ListTaxFormsRequestOrderBy {
	return &v
}

type NullableListTaxFormsRequestOrderBy struct {
	value *ListTaxFormsRequestOrderBy
	isSet bool
}

func (v NullableListTaxFormsRequestOrderBy) Get() *ListTaxFormsRequestOrderBy {
	return v.value
}

func (v *NullableListTaxFormsRequestOrderBy) Set(val *ListTaxFormsRequestOrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullableListTaxFormsRequestOrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullableListTaxFormsRequestOrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTaxFormsRequestOrderBy(val *ListTaxFormsRequestOrderBy) *NullableListTaxFormsRequestOrderBy {
	return &NullableListTaxFormsRequestOrderBy{value: val, isSet: true}
}

func (v NullableListTaxFormsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTaxFormsRequestOrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

