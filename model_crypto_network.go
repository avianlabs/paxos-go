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

// CryptoNetwork A CryptoNetwork is a blockchain transmitting cryptocurrencies.
type CryptoNetwork string

// List of CryptoNetwork
const (
	CryptoNetwork_BITCOIN CryptoNetwork = "BITCOIN"
	CryptoNetwork_ETHEREUM CryptoNetwork = "ETHEREUM"
	CryptoNetwork_BITCOIN_CASH CryptoNetwork = "BITCOIN_CASH"
	CryptoNetwork_LITECOIN CryptoNetwork = "LITECOIN"
	CryptoNetwork_SOLANA CryptoNetwork = "SOLANA"
	CryptoNetwork_POLYGON_POS CryptoNetwork = "POLYGON_POS"
	CryptoNetwork_ARBITRUM_ONE CryptoNetwork = "ARBITRUM_ONE"
)

// All allowed values of CryptoNetwork enum
var AllowedCryptoNetworkEnumValues = []CryptoNetwork{
	"BITCOIN",
	"ETHEREUM",
	"BITCOIN_CASH",
	"LITECOIN",
	"SOLANA",
	"POLYGON_POS",
	"ARBITRUM_ONE",
}

func (v *CryptoNetwork) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CryptoNetwork(value)
	for _, existing := range AllowedCryptoNetworkEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CryptoNetwork", value)
}

// NewCryptoNetworkFromValue returns a pointer to a valid CryptoNetwork
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCryptoNetworkFromValue(v string) (*CryptoNetwork, error) {
	ev := CryptoNetwork(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CryptoNetwork: valid values are %v", v, AllowedCryptoNetworkEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CryptoNetwork) IsValid() bool {
	for _, existing := range AllowedCryptoNetworkEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CryptoNetwork value
func (v CryptoNetwork) Ptr() *CryptoNetwork {
	return &v
}

type NullableCryptoNetwork struct {
	value *CryptoNetwork
	isSet bool
}

func (v NullableCryptoNetwork) Get() *CryptoNetwork {
	return v.value
}

func (v *NullableCryptoNetwork) Set(val *CryptoNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptoNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptoNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptoNetwork(val *CryptoNetwork) *NullableCryptoNetwork {
	return &NullableCryptoNetwork{value: val, isSet: true}
}

func (v NullableCryptoNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptoNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

