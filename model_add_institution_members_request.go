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

// checks if the AddInstitutionMembersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddInstitutionMembersRequest{}

// AddInstitutionMembersRequest struct for AddInstitutionMembersRequest
type AddInstitutionMembersRequest struct {
	// ID of institution identity to which members will be added.
	InstitutionId string `json:"institution_id"`
	// A non-empty array of institution members to be added.
	Members []InstitutionMember `json:"members"`
	AdditionalProperties map[string]interface{}
}

type _AddInstitutionMembersRequest AddInstitutionMembersRequest

// NewAddInstitutionMembersRequest instantiates a new AddInstitutionMembersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddInstitutionMembersRequest(institutionId string, members []InstitutionMember) *AddInstitutionMembersRequest {
	this := AddInstitutionMembersRequest{}
	this.InstitutionId = institutionId
	this.Members = members
	return &this
}

// NewAddInstitutionMembersRequestWithDefaults instantiates a new AddInstitutionMembersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddInstitutionMembersRequestWithDefaults() *AddInstitutionMembersRequest {
	this := AddInstitutionMembersRequest{}
	return &this
}

// GetInstitutionId returns the InstitutionId field value
func (o *AddInstitutionMembersRequest) GetInstitutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value
// and a boolean to check if the value has been set.
func (o *AddInstitutionMembersRequest) GetInstitutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstitutionId, true
}

// SetInstitutionId sets field value
func (o *AddInstitutionMembersRequest) SetInstitutionId(v string) {
	o.InstitutionId = v
}

// GetMembers returns the Members field value
func (o *AddInstitutionMembersRequest) GetMembers() []InstitutionMember {
	if o == nil {
		var ret []InstitutionMember
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *AddInstitutionMembersRequest) GetMembersOk() ([]InstitutionMember, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *AddInstitutionMembersRequest) SetMembers(v []InstitutionMember) {
	o.Members = v
}

func (o AddInstitutionMembersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddInstitutionMembersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["institution_id"] = o.InstitutionId
	toSerialize["members"] = o.Members

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AddInstitutionMembersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"institution_id",
		"members",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAddInstitutionMembersRequest := _AddInstitutionMembersRequest{}

	err = json.Unmarshal(data, &varAddInstitutionMembersRequest)

	if err != nil {
		return err
	}

	*o = AddInstitutionMembersRequest(varAddInstitutionMembersRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "institution_id")
		delete(additionalProperties, "members")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAddInstitutionMembersRequest struct {
	value *AddInstitutionMembersRequest
	isSet bool
}

func (v NullableAddInstitutionMembersRequest) Get() *AddInstitutionMembersRequest {
	return v.value
}

func (v *NullableAddInstitutionMembersRequest) Set(val *AddInstitutionMembersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddInstitutionMembersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddInstitutionMembersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddInstitutionMembersRequest(val *AddInstitutionMembersRequest) *NullableAddInstitutionMembersRequest {
	return &NullableAddInstitutionMembersRequest{value: val, isSet: true}
}

func (v NullableAddInstitutionMembersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddInstitutionMembersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


