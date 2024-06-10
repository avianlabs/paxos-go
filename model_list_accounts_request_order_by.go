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

// ListAccountsRequestOrderBy the model 'ListAccountsRequestOrderBy'
type ListAccountsRequestOrderBy string

// List of ListAccountsRequestOrderBy
const (
	CREATED_AT ListAccountsRequestOrderBy = "CREATED_AT"
)

// All allowed values of ListAccountsRequestOrderBy enum
var AllowedListAccountsRequestOrderByEnumValues = []ListAccountsRequestOrderBy{
	"CREATED_AT",
}

func (v *ListAccountsRequestOrderBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ListAccountsRequestOrderBy(value)
	for _, existing := range AllowedListAccountsRequestOrderByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ListAccountsRequestOrderBy", value)
}

// NewListAccountsRequestOrderByFromValue returns a pointer to a valid ListAccountsRequestOrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListAccountsRequestOrderByFromValue(v string) (*ListAccountsRequestOrderBy, error) {
	ev := ListAccountsRequestOrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListAccountsRequestOrderBy: valid values are %v", v, AllowedListAccountsRequestOrderByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListAccountsRequestOrderBy) IsValid() bool {
	for _, existing := range AllowedListAccountsRequestOrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListAccountsRequestOrderBy value
func (v ListAccountsRequestOrderBy) Ptr() *ListAccountsRequestOrderBy {
	return &v
}

type NullableListAccountsRequestOrderBy struct {
	value *ListAccountsRequestOrderBy
	isSet bool
}

func (v NullableListAccountsRequestOrderBy) Get() *ListAccountsRequestOrderBy {
	return v.value
}

func (v *NullableListAccountsRequestOrderBy) Set(val *ListAccountsRequestOrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullableListAccountsRequestOrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullableListAccountsRequestOrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAccountsRequestOrderBy(val *ListAccountsRequestOrderBy) *NullableListAccountsRequestOrderBy {
	return &NullableListAccountsRequestOrderBy{value: val, isSet: true}
}

func (v NullableListAccountsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAccountsRequestOrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

