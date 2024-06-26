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

// Asset The given asset or assets held in a Paxos balance.
type Asset string

// List of Asset
const (
	Asset_USD Asset = "USD"
	Asset_EUR Asset = "EUR"
	Asset_SGD Asset = "SGD"
	Asset_BTC Asset = "BTC"
	Asset_ETH Asset = "ETH"
	Asset_PAXG Asset = "PAXG"
	Asset_BUSD Asset = "BUSD"
	Asset_BCH Asset = "BCH"
	Asset_LTC Asset = "LTC"
	Asset_USDC Asset = "USDC"
	Asset_USDP Asset = "USDP"
	Asset_AAVE Asset = "AAVE"
	Asset_UNI Asset = "UNI"
	Asset_MATIC Asset = "MATIC"
	Asset_PYUSD Asset = "PYUSD"
	Asset_LINK Asset = "LINK"
	Asset_SOL Asset = "SOL"
)

// All allowed values of Asset enum
var AllowedAssetEnumValues = []Asset{
	"USD",
	"EUR",
	"SGD",
	"BTC",
	"ETH",
	"PAXG",
	"BUSD",
	"BCH",
	"LTC",
	"USDC",
	"USDP",
	"AAVE",
	"UNI",
	"MATIC",
	"PYUSD",
	"LINK",
	"SOL",
}

func (v *Asset) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Asset(value)
	for _, existing := range AllowedAssetEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Asset", value)
}

// NewAssetFromValue returns a pointer to a valid Asset
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssetFromValue(v string) (*Asset, error) {
	ev := Asset(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Asset: valid values are %v", v, AllowedAssetEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Asset) IsValid() bool {
	for _, existing := range AllowedAssetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Asset value
func (v Asset) Ptr() *Asset {
	return &v
}

type NullableAsset struct {
	value *Asset
	isSet bool
}

func (v NullableAsset) Get() *Asset {
	return v.value
}

func (v *NullableAsset) Set(val *Asset) {
	v.value = val
	v.isSet = true
}

func (v NullableAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsset(val *Asset) *NullableAsset {
	return &NullableAsset{value: val, isSet: true}
}

func (v NullableAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

