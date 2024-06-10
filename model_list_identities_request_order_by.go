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

// ListIdentitiesRequestOrderBy the model 'ListIdentitiesRequestOrderBy'
type ListIdentitiesRequestOrderBy string

// List of ListIdentitiesRequestOrderBy
const (
	CREATED_AT_ListIdentitiesRequestOrderBy ListIdentitiesRequestOrderBy = "CREATED_AT"
)

// All allowed values of ListIdentitiesRequestOrderBy enum
var AllowedListIdentitiesRequestOrderByEnumValues = []ListIdentitiesRequestOrderBy{
	"CREATED_AT",
}

func (v *ListIdentitiesRequestOrderBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ListIdentitiesRequestOrderBy(value)
	for _, existing := range AllowedListIdentitiesRequestOrderByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ListIdentitiesRequestOrderBy", value)
}

// NewListIdentitiesRequestOrderByFromValue returns a pointer to a valid ListIdentitiesRequestOrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListIdentitiesRequestOrderByFromValue(v string) (*ListIdentitiesRequestOrderBy, error) {
	ev := ListIdentitiesRequestOrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListIdentitiesRequestOrderBy: valid values are %v", v, AllowedListIdentitiesRequestOrderByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListIdentitiesRequestOrderBy) IsValid() bool {
	for _, existing := range AllowedListIdentitiesRequestOrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListIdentitiesRequestOrderBy value
func (v ListIdentitiesRequestOrderBy) Ptr() *ListIdentitiesRequestOrderBy {
	return &v
}

type NullableListIdentitiesRequestOrderBy struct {
	value *ListIdentitiesRequestOrderBy
	isSet bool
}

func (v NullableListIdentitiesRequestOrderBy) Get() *ListIdentitiesRequestOrderBy {
	return v.value
}

func (v *NullableListIdentitiesRequestOrderBy) Set(val *ListIdentitiesRequestOrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullableListIdentitiesRequestOrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullableListIdentitiesRequestOrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListIdentitiesRequestOrderBy(val *ListIdentitiesRequestOrderBy) *NullableListIdentitiesRequestOrderBy {
	return &NullableListIdentitiesRequestOrderBy{value: val, isSet: true}
}

func (v NullableListIdentitiesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListIdentitiesRequestOrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

