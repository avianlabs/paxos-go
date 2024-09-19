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

// ListTransfersRequestOrderBy Query filter order.
type ListTransfersRequestOrderBy string

// List of ListTransfersRequestOrderBy
const (
	CREATED_AT ListTransfersRequestOrderBy = "CREATED_AT"
	UPDATED_AT ListTransfersRequestOrderBy = "UPDATED_AT"
)

// All allowed values of ListTransfersRequestOrderBy enum
var AllowedListTransfersRequestOrderByEnumValues = []ListTransfersRequestOrderBy{
	"CREATED_AT",
	"UPDATED_AT",
}

func (v *ListTransfersRequestOrderBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ListTransfersRequestOrderBy(value)
	for _, existing := range AllowedListTransfersRequestOrderByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ListTransfersRequestOrderBy", value)
}

// NewListTransfersRequestOrderByFromValue returns a pointer to a valid ListTransfersRequestOrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListTransfersRequestOrderByFromValue(v string) (*ListTransfersRequestOrderBy, error) {
	ev := ListTransfersRequestOrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListTransfersRequestOrderBy: valid values are %v", v, AllowedListTransfersRequestOrderByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListTransfersRequestOrderBy) IsValid() bool {
	for _, existing := range AllowedListTransfersRequestOrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListTransfersRequestOrderBy value
func (v ListTransfersRequestOrderBy) Ptr() *ListTransfersRequestOrderBy {
	return &v
}

type NullableListTransfersRequestOrderBy struct {
	value *ListTransfersRequestOrderBy
	isSet bool
}

func (v NullableListTransfersRequestOrderBy) Get() *ListTransfersRequestOrderBy {
	return v.value
}

func (v *NullableListTransfersRequestOrderBy) Set(val *ListTransfersRequestOrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransfersRequestOrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransfersRequestOrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransfersRequestOrderBy(val *ListTransfersRequestOrderBy) *NullableListTransfersRequestOrderBy {
	return &NullableListTransfersRequestOrderBy{value: val, isSet: true}
}

func (v NullableListTransfersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransfersRequestOrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

