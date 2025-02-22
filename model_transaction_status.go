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

// TransactionStatus - PENDING: Initial state of a settlement transaction upon creation.  - SETTLED: Indicates all obligations belong to the settlement transaction have been enacted.  - EXPIRED: Indicates the settlement transaction is no longer eligible for settlement.  - CANCELLED: Indicates the settlement transaction was cancelled by the source profile.  - AFFIRMED: Indicates the settlement transaction will be eligible for settlement once within the window.
type TransactionStatus string

// List of TransactionStatus
const (
	TransactionStatus_PENDING TransactionStatus = "PENDING"
	TransactionStatus_SETTLED TransactionStatus = "SETTLED"
	TransactionStatus_EXPIRED TransactionStatus = "EXPIRED"
	TransactionStatus_CANCELLED TransactionStatus = "CANCELLED"
	TransactionStatus_AFFIRMED TransactionStatus = "AFFIRMED"
)

// All allowed values of TransactionStatus enum
var AllowedTransactionStatusEnumValues = []TransactionStatus{
	"PENDING",
	"SETTLED",
	"EXPIRED",
	"CANCELLED",
	"AFFIRMED",
}

func (v *TransactionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionStatus(value)
	for _, existing := range AllowedTransactionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransactionStatus", value)
}

// NewTransactionStatusFromValue returns a pointer to a valid TransactionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionStatusFromValue(v string) (*TransactionStatus, error) {
	ev := TransactionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransactionStatus: valid values are %v", v, AllowedTransactionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionStatus) IsValid() bool {
	for _, existing := range AllowedTransactionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionStatus value
func (v TransactionStatus) Ptr() *TransactionStatus {
	return &v
}

type NullableTransactionStatus struct {
	value *TransactionStatus
	isSet bool
}

func (v NullableTransactionStatus) Get() *TransactionStatus {
	return v.value
}

func (v *NullableTransactionStatus) Set(val *TransactionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionStatus(val *TransactionStatus) *NullableTransactionStatus {
	return &NullableTransactionStatus{value: val, isSet: true}
}

func (v NullableTransactionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

