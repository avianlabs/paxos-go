/*
Paxos API

<p>Welcome to Paxos APIs. At Paxos, our mission is to enable the movement of any asset, any time, in a trustworthy way. These APIs serve that mission by making it easier than ever for you to directly integrate our product capabilities into your application, leveraging the speed, stability, and security of the Paxos platform.</p> <p>The documentation that follows gives you access to our Crypto Brokerage, Trading, and Exchange products. It includes APIs for market data, orders, and the held rate quote flow.</p> <p>To test in our sandbox environment, <a href=\"https://account.sandbox.paxos.com\" target=\"_blank\">sign up</a> for an account. For more information about Paxos and our APIs, visit <a href=\"https://www.paxos.com/\" target=\"_blank\">Paxos.com</a>.</p> 

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paxos

import (
	"encoding/json"
)

// checks if the AddressInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressInfo{}

// AddressInfo struct for AddressInfo
type AddressInfo struct {
	IndividualInfo *IndividualInfo `json:"individual_info,omitempty"`
	InstitutionInfo *InstitutionInfo `json:"institution_info,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AddressInfo AddressInfo

// NewAddressInfo instantiates a new AddressInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressInfo() *AddressInfo {
	this := AddressInfo{}
	return &this
}

// NewAddressInfoWithDefaults instantiates a new AddressInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressInfoWithDefaults() *AddressInfo {
	this := AddressInfo{}
	return &this
}

// GetIndividualInfo returns the IndividualInfo field value if set, zero value otherwise.
func (o *AddressInfo) GetIndividualInfo() IndividualInfo {
	if o == nil || IsNil(o.IndividualInfo) {
		var ret IndividualInfo
		return ret
	}
	return *o.IndividualInfo
}

// GetIndividualInfoOk returns a tuple with the IndividualInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetIndividualInfoOk() (*IndividualInfo, bool) {
	if o == nil || IsNil(o.IndividualInfo) {
		return nil, false
	}
	return o.IndividualInfo, true
}

// HasIndividualInfo returns a boolean if a field has been set.
func (o *AddressInfo) HasIndividualInfo() bool {
	if o != nil && !IsNil(o.IndividualInfo) {
		return true
	}

	return false
}

// SetIndividualInfo gets a reference to the given IndividualInfo and assigns it to the IndividualInfo field.
func (o *AddressInfo) SetIndividualInfo(v IndividualInfo) {
	o.IndividualInfo = &v
}

// GetInstitutionInfo returns the InstitutionInfo field value if set, zero value otherwise.
func (o *AddressInfo) GetInstitutionInfo() InstitutionInfo {
	if o == nil || IsNil(o.InstitutionInfo) {
		var ret InstitutionInfo
		return ret
	}
	return *o.InstitutionInfo
}

// GetInstitutionInfoOk returns a tuple with the InstitutionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInfo) GetInstitutionInfoOk() (*InstitutionInfo, bool) {
	if o == nil || IsNil(o.InstitutionInfo) {
		return nil, false
	}
	return o.InstitutionInfo, true
}

// HasInstitutionInfo returns a boolean if a field has been set.
func (o *AddressInfo) HasInstitutionInfo() bool {
	if o != nil && !IsNil(o.InstitutionInfo) {
		return true
	}

	return false
}

// SetInstitutionInfo gets a reference to the given InstitutionInfo and assigns it to the InstitutionInfo field.
func (o *AddressInfo) SetInstitutionInfo(v InstitutionInfo) {
	o.InstitutionInfo = &v
}

func (o AddressInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IndividualInfo) {
		toSerialize["individual_info"] = o.IndividualInfo
	}
	if !IsNil(o.InstitutionInfo) {
		toSerialize["institution_info"] = o.InstitutionInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AddressInfo) UnmarshalJSON(data []byte) (err error) {
	varAddressInfo := _AddressInfo{}

	err = json.Unmarshal(data, &varAddressInfo)

	if err != nil {
		return err
	}

	*o = AddressInfo(varAddressInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "individual_info")
		delete(additionalProperties, "institution_info")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAddressInfo struct {
	value *AddressInfo
	isSet bool
}

func (v NullableAddressInfo) Get() *AddressInfo {
	return v.value
}

func (v *NullableAddressInfo) Set(val *AddressInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressInfo(val *AddressInfo) *NullableAddressInfo {
	return &NullableAddressInfo{value: val, isSet: true}
}

func (v NullableAddressInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


