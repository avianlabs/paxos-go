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

// checks if the IdentityPublicDocumentUploadBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityPublicDocumentUploadBody{}

// IdentityPublicDocumentUploadBody struct for IdentityPublicDocumentUploadBody
type IdentityPublicDocumentUploadBody struct {
	// The file name.
	Name string `json:"name"`
	// A list of document types contained within the uploaded file.
	DocumentTypes []DocumentType `json:"document_types,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityPublicDocumentUploadBody IdentityPublicDocumentUploadBody

// NewIdentityPublicDocumentUploadBody instantiates a new IdentityPublicDocumentUploadBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityPublicDocumentUploadBody(name string) *IdentityPublicDocumentUploadBody {
	this := IdentityPublicDocumentUploadBody{}
	this.Name = name
	return &this
}

// NewIdentityPublicDocumentUploadBodyWithDefaults instantiates a new IdentityPublicDocumentUploadBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityPublicDocumentUploadBodyWithDefaults() *IdentityPublicDocumentUploadBody {
	this := IdentityPublicDocumentUploadBody{}
	return &this
}

// GetName returns the Name field value
func (o *IdentityPublicDocumentUploadBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IdentityPublicDocumentUploadBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IdentityPublicDocumentUploadBody) SetName(v string) {
	o.Name = v
}

// GetDocumentTypes returns the DocumentTypes field value if set, zero value otherwise.
func (o *IdentityPublicDocumentUploadBody) GetDocumentTypes() []DocumentType {
	if o == nil || IsNil(o.DocumentTypes) {
		var ret []DocumentType
		return ret
	}
	return o.DocumentTypes
}

// GetDocumentTypesOk returns a tuple with the DocumentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityPublicDocumentUploadBody) GetDocumentTypesOk() ([]DocumentType, bool) {
	if o == nil || IsNil(o.DocumentTypes) {
		return nil, false
	}
	return o.DocumentTypes, true
}

// HasDocumentTypes returns a boolean if a field has been set.
func (o *IdentityPublicDocumentUploadBody) HasDocumentTypes() bool {
	if o != nil && !IsNil(o.DocumentTypes) {
		return true
	}

	return false
}

// SetDocumentTypes gets a reference to the given []DocumentType and assigns it to the DocumentTypes field.
func (o *IdentityPublicDocumentUploadBody) SetDocumentTypes(v []DocumentType) {
	o.DocumentTypes = v
}

func (o IdentityPublicDocumentUploadBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityPublicDocumentUploadBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.DocumentTypes) {
		toSerialize["document_types"] = o.DocumentTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityPublicDocumentUploadBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varIdentityPublicDocumentUploadBody := _IdentityPublicDocumentUploadBody{}

	err = json.Unmarshal(data, &varIdentityPublicDocumentUploadBody)

	if err != nil {
		return err
	}

	*o = IdentityPublicDocumentUploadBody(varIdentityPublicDocumentUploadBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "document_types")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityPublicDocumentUploadBody struct {
	value *IdentityPublicDocumentUploadBody
	isSet bool
}

func (v NullableIdentityPublicDocumentUploadBody) Get() *IdentityPublicDocumentUploadBody {
	return v.value
}

func (v *NullableIdentityPublicDocumentUploadBody) Set(val *IdentityPublicDocumentUploadBody) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityPublicDocumentUploadBody) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityPublicDocumentUploadBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityPublicDocumentUploadBody(val *IdentityPublicDocumentUploadBody) *NullableIdentityPublicDocumentUploadBody {
	return &NullableIdentityPublicDocumentUploadBody{value: val, isSet: true}
}

func (v NullableIdentityPublicDocumentUploadBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityPublicDocumentUploadBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


