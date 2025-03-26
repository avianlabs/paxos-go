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

// checks if the BeneficiaryPerson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BeneficiaryPerson{}

// BeneficiaryPerson struct for BeneficiaryPerson
type BeneficiaryPerson struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BeneficiaryPerson BeneficiaryPerson

// NewBeneficiaryPerson instantiates a new BeneficiaryPerson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeneficiaryPerson() *BeneficiaryPerson {
	this := BeneficiaryPerson{}
	return &this
}

// NewBeneficiaryPersonWithDefaults instantiates a new BeneficiaryPerson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeneficiaryPersonWithDefaults() *BeneficiaryPerson {
	this := BeneficiaryPerson{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *BeneficiaryPerson) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeneficiaryPerson) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *BeneficiaryPerson) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *BeneficiaryPerson) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *BeneficiaryPerson) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeneficiaryPerson) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *BeneficiaryPerson) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *BeneficiaryPerson) SetLastName(v string) {
	o.LastName = &v
}

func (o BeneficiaryPerson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BeneficiaryPerson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BeneficiaryPerson) UnmarshalJSON(data []byte) (err error) {
	varBeneficiaryPerson := _BeneficiaryPerson{}

	err = json.Unmarshal(data, &varBeneficiaryPerson)

	if err != nil {
		return err
	}

	*o = BeneficiaryPerson(varBeneficiaryPerson)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "first_name")
		delete(additionalProperties, "last_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBeneficiaryPerson struct {
	value *BeneficiaryPerson
	isSet bool
}

func (v NullableBeneficiaryPerson) Get() *BeneficiaryPerson {
	return v.value
}

func (v *NullableBeneficiaryPerson) Set(val *BeneficiaryPerson) {
	v.value = val
	v.isSet = true
}

func (v NullableBeneficiaryPerson) IsSet() bool {
	return v.isSet
}

func (v *NullableBeneficiaryPerson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeneficiaryPerson(val *BeneficiaryPerson) *NullableBeneficiaryPerson {
	return &NullableBeneficiaryPerson{value: val, isSet: true}
}

func (v NullableBeneficiaryPerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeneficiaryPerson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


