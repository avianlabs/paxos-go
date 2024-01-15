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

// TransferDirection Direction of the transfer.
type TransferDirection string

// List of TransferDirection
const (
	TRANSFERDIRECTION_CREDIT TransferDirection = "CREDIT"
	TRANSFERDIRECTION_DEBIT TransferDirection = "DEBIT"
)

// All allowed values of TransferDirection enum
var AllowedTransferDirectionEnumValues = []TransferDirection{
	"CREDIT",
	"DEBIT",
}

func (v *TransferDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransferDirection(value)
	for _, existing := range AllowedTransferDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransferDirection", value)
}

// NewTransferDirectionFromValue returns a pointer to a valid TransferDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferDirectionFromValue(v string) (*TransferDirection, error) {
	ev := TransferDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransferDirection: valid values are %v", v, AllowedTransferDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferDirection) IsValid() bool {
	for _, existing := range AllowedTransferDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferDirection value
func (v TransferDirection) Ptr() *TransferDirection {
	return &v
}

type NullableTransferDirection struct {
	value *TransferDirection
	isSet bool
}

func (v NullableTransferDirection) Get() *TransferDirection {
	return v.value
}

func (v *NullableTransferDirection) Set(val *TransferDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferDirection(val *TransferDirection) *NullableTransferDirection {
	return &NullableTransferDirection{value: val, isSet: true}
}

func (v NullableTransferDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

