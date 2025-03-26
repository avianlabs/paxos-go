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

// checks if the FiatAccountOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FiatAccountOwner{}

// FiatAccountOwner struct for FiatAccountOwner
type FiatAccountOwner struct {
	PersonDetails *FiatAccountOwnerPersonDetails `json:"person_details,omitempty"`
	InstitutionDetails *FiatAccountOwnerInstitutionDetails `json:"institution_details,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FiatAccountOwner FiatAccountOwner

// NewFiatAccountOwner instantiates a new FiatAccountOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiatAccountOwner() *FiatAccountOwner {
	this := FiatAccountOwner{}
	return &this
}

// NewFiatAccountOwnerWithDefaults instantiates a new FiatAccountOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiatAccountOwnerWithDefaults() *FiatAccountOwner {
	this := FiatAccountOwner{}
	return &this
}

// GetPersonDetails returns the PersonDetails field value if set, zero value otherwise.
func (o *FiatAccountOwner) GetPersonDetails() FiatAccountOwnerPersonDetails {
	if o == nil || IsNil(o.PersonDetails) {
		var ret FiatAccountOwnerPersonDetails
		return ret
	}
	return *o.PersonDetails
}

// GetPersonDetailsOk returns a tuple with the PersonDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatAccountOwner) GetPersonDetailsOk() (*FiatAccountOwnerPersonDetails, bool) {
	if o == nil || IsNil(o.PersonDetails) {
		return nil, false
	}
	return o.PersonDetails, true
}

// HasPersonDetails returns a boolean if a field has been set.
func (o *FiatAccountOwner) HasPersonDetails() bool {
	if o != nil && !IsNil(o.PersonDetails) {
		return true
	}

	return false
}

// SetPersonDetails gets a reference to the given FiatAccountOwnerPersonDetails and assigns it to the PersonDetails field.
func (o *FiatAccountOwner) SetPersonDetails(v FiatAccountOwnerPersonDetails) {
	o.PersonDetails = &v
}

// GetInstitutionDetails returns the InstitutionDetails field value if set, zero value otherwise.
func (o *FiatAccountOwner) GetInstitutionDetails() FiatAccountOwnerInstitutionDetails {
	if o == nil || IsNil(o.InstitutionDetails) {
		var ret FiatAccountOwnerInstitutionDetails
		return ret
	}
	return *o.InstitutionDetails
}

// GetInstitutionDetailsOk returns a tuple with the InstitutionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatAccountOwner) GetInstitutionDetailsOk() (*FiatAccountOwnerInstitutionDetails, bool) {
	if o == nil || IsNil(o.InstitutionDetails) {
		return nil, false
	}
	return o.InstitutionDetails, true
}

// HasInstitutionDetails returns a boolean if a field has been set.
func (o *FiatAccountOwner) HasInstitutionDetails() bool {
	if o != nil && !IsNil(o.InstitutionDetails) {
		return true
	}

	return false
}

// SetInstitutionDetails gets a reference to the given FiatAccountOwnerInstitutionDetails and assigns it to the InstitutionDetails field.
func (o *FiatAccountOwner) SetInstitutionDetails(v FiatAccountOwnerInstitutionDetails) {
	o.InstitutionDetails = &v
}

func (o FiatAccountOwner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FiatAccountOwner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PersonDetails) {
		toSerialize["person_details"] = o.PersonDetails
	}
	if !IsNil(o.InstitutionDetails) {
		toSerialize["institution_details"] = o.InstitutionDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FiatAccountOwner) UnmarshalJSON(data []byte) (err error) {
	varFiatAccountOwner := _FiatAccountOwner{}

	err = json.Unmarshal(data, &varFiatAccountOwner)

	if err != nil {
		return err
	}

	*o = FiatAccountOwner(varFiatAccountOwner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "person_details")
		delete(additionalProperties, "institution_details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFiatAccountOwner struct {
	value *FiatAccountOwner
	isSet bool
}

func (v NullableFiatAccountOwner) Get() *FiatAccountOwner {
	return v.value
}

func (v *NullableFiatAccountOwner) Set(val *FiatAccountOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableFiatAccountOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableFiatAccountOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiatAccountOwner(val *FiatAccountOwner) *NullableFiatAccountOwner {
	return &NullableFiatAccountOwner{value: val, isSet: true}
}

func (v NullableFiatAccountOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiatAccountOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


