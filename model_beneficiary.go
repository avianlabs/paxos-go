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

// checks if the Beneficiary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Beneficiary{}

// Beneficiary struct for Beneficiary
type Beneficiary struct {
	PersonDetails *BeneficiaryPerson `json:"person_details,omitempty"`
	InstitutionDetails *BeneficiaryInstitution `json:"institution_details,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Beneficiary Beneficiary

// NewBeneficiary instantiates a new Beneficiary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeneficiary() *Beneficiary {
	this := Beneficiary{}
	return &this
}

// NewBeneficiaryWithDefaults instantiates a new Beneficiary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeneficiaryWithDefaults() *Beneficiary {
	this := Beneficiary{}
	return &this
}

// GetPersonDetails returns the PersonDetails field value if set, zero value otherwise.
func (o *Beneficiary) GetPersonDetails() BeneficiaryPerson {
	if o == nil || IsNil(o.PersonDetails) {
		var ret BeneficiaryPerson
		return ret
	}
	return *o.PersonDetails
}

// GetPersonDetailsOk returns a tuple with the PersonDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Beneficiary) GetPersonDetailsOk() (*BeneficiaryPerson, bool) {
	if o == nil || IsNil(o.PersonDetails) {
		return nil, false
	}
	return o.PersonDetails, true
}

// HasPersonDetails returns a boolean if a field has been set.
func (o *Beneficiary) HasPersonDetails() bool {
	if o != nil && !IsNil(o.PersonDetails) {
		return true
	}

	return false
}

// SetPersonDetails gets a reference to the given BeneficiaryPerson and assigns it to the PersonDetails field.
func (o *Beneficiary) SetPersonDetails(v BeneficiaryPerson) {
	o.PersonDetails = &v
}

// GetInstitutionDetails returns the InstitutionDetails field value if set, zero value otherwise.
func (o *Beneficiary) GetInstitutionDetails() BeneficiaryInstitution {
	if o == nil || IsNil(o.InstitutionDetails) {
		var ret BeneficiaryInstitution
		return ret
	}
	return *o.InstitutionDetails
}

// GetInstitutionDetailsOk returns a tuple with the InstitutionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Beneficiary) GetInstitutionDetailsOk() (*BeneficiaryInstitution, bool) {
	if o == nil || IsNil(o.InstitutionDetails) {
		return nil, false
	}
	return o.InstitutionDetails, true
}

// HasInstitutionDetails returns a boolean if a field has been set.
func (o *Beneficiary) HasInstitutionDetails() bool {
	if o != nil && !IsNil(o.InstitutionDetails) {
		return true
	}

	return false
}

// SetInstitutionDetails gets a reference to the given BeneficiaryInstitution and assigns it to the InstitutionDetails field.
func (o *Beneficiary) SetInstitutionDetails(v BeneficiaryInstitution) {
	o.InstitutionDetails = &v
}

func (o Beneficiary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Beneficiary) ToMap() (map[string]interface{}, error) {
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

func (o *Beneficiary) UnmarshalJSON(data []byte) (err error) {
	varBeneficiary := _Beneficiary{}

	err = json.Unmarshal(data, &varBeneficiary)

	if err != nil {
		return err
	}

	*o = Beneficiary(varBeneficiary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "person_details")
		delete(additionalProperties, "institution_details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBeneficiary struct {
	value *Beneficiary
	isSet bool
}

func (v NullableBeneficiary) Get() *Beneficiary {
	return v.value
}

func (v *NullableBeneficiary) Set(val *Beneficiary) {
	v.value = val
	v.isSet = true
}

func (v NullableBeneficiary) IsSet() bool {
	return v.isSet
}

func (v *NullableBeneficiary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeneficiary(val *Beneficiary) *NullableBeneficiary {
	return &NullableBeneficiary{value: val, isSet: true}
}

func (v NullableBeneficiary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeneficiary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


