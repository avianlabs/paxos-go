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

// ListDepositAddressesRequestOrderBy Query filter order.
type ListDepositAddressesRequestOrderBy string

// List of ListDepositAddressesRequestOrderBy
const (
	LISTDEPOSITADDRESSESREQUESTORDERBY_CREATED_AT ListDepositAddressesRequestOrderBy = "CREATED_AT"
)

// All allowed values of ListDepositAddressesRequestOrderBy enum
var AllowedListDepositAddressesRequestOrderByEnumValues = []ListDepositAddressesRequestOrderBy{
	"CREATED_AT",
}

func (v *ListDepositAddressesRequestOrderBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ListDepositAddressesRequestOrderBy(value)
	for _, existing := range AllowedListDepositAddressesRequestOrderByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ListDepositAddressesRequestOrderBy", value)
}

// NewListDepositAddressesRequestOrderByFromValue returns a pointer to a valid ListDepositAddressesRequestOrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListDepositAddressesRequestOrderByFromValue(v string) (*ListDepositAddressesRequestOrderBy, error) {
	ev := ListDepositAddressesRequestOrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListDepositAddressesRequestOrderBy: valid values are %v", v, AllowedListDepositAddressesRequestOrderByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListDepositAddressesRequestOrderBy) IsValid() bool {
	for _, existing := range AllowedListDepositAddressesRequestOrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListDepositAddressesRequestOrderBy value
func (v ListDepositAddressesRequestOrderBy) Ptr() *ListDepositAddressesRequestOrderBy {
	return &v
}

type NullableListDepositAddressesRequestOrderBy struct {
	value *ListDepositAddressesRequestOrderBy
	isSet bool
}

func (v NullableListDepositAddressesRequestOrderBy) Get() *ListDepositAddressesRequestOrderBy {
	return v.value
}

func (v *NullableListDepositAddressesRequestOrderBy) Set(val *ListDepositAddressesRequestOrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullableListDepositAddressesRequestOrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullableListDepositAddressesRequestOrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDepositAddressesRequestOrderBy(val *ListDepositAddressesRequestOrderBy) *NullableListDepositAddressesRequestOrderBy {
	return &NullableListDepositAddressesRequestOrderBy{value: val, isSet: true}
}

func (v NullableListDepositAddressesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDepositAddressesRequestOrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

