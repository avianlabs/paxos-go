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

// checks if the UpdateFiatAccountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFiatAccountRequest{}

// UpdateFiatAccountRequest struct for UpdateFiatAccountRequest
type UpdateFiatAccountRequest struct {
	FiatAccountOwner *FiatAccountOwner `json:"fiat_account_owner,omitempty"`
	FiatNetworkInstructions *UpdateFiatAccountRequestUpdateFiatNetworkInstructions `json:"fiat_network_instructions,omitempty"`
	// Optional client-specified metadata. Up to 6 key/value pairs may be provided. Each key and value must be less than or equal to 100 characters.
	Metadata *map[string]string `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateFiatAccountRequest UpdateFiatAccountRequest

// NewUpdateFiatAccountRequest instantiates a new UpdateFiatAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFiatAccountRequest() *UpdateFiatAccountRequest {
	this := UpdateFiatAccountRequest{}
	return &this
}

// NewUpdateFiatAccountRequestWithDefaults instantiates a new UpdateFiatAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFiatAccountRequestWithDefaults() *UpdateFiatAccountRequest {
	this := UpdateFiatAccountRequest{}
	return &this
}

// GetFiatAccountOwner returns the FiatAccountOwner field value if set, zero value otherwise.
func (o *UpdateFiatAccountRequest) GetFiatAccountOwner() FiatAccountOwner {
	if o == nil || IsNil(o.FiatAccountOwner) {
		var ret FiatAccountOwner
		return ret
	}
	return *o.FiatAccountOwner
}

// GetFiatAccountOwnerOk returns a tuple with the FiatAccountOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFiatAccountRequest) GetFiatAccountOwnerOk() (*FiatAccountOwner, bool) {
	if o == nil || IsNil(o.FiatAccountOwner) {
		return nil, false
	}
	return o.FiatAccountOwner, true
}

// HasFiatAccountOwner returns a boolean if a field has been set.
func (o *UpdateFiatAccountRequest) HasFiatAccountOwner() bool {
	if o != nil && !IsNil(o.FiatAccountOwner) {
		return true
	}

	return false
}

// SetFiatAccountOwner gets a reference to the given FiatAccountOwner and assigns it to the FiatAccountOwner field.
func (o *UpdateFiatAccountRequest) SetFiatAccountOwner(v FiatAccountOwner) {
	o.FiatAccountOwner = &v
}

// GetFiatNetworkInstructions returns the FiatNetworkInstructions field value if set, zero value otherwise.
func (o *UpdateFiatAccountRequest) GetFiatNetworkInstructions() UpdateFiatAccountRequestUpdateFiatNetworkInstructions {
	if o == nil || IsNil(o.FiatNetworkInstructions) {
		var ret UpdateFiatAccountRequestUpdateFiatNetworkInstructions
		return ret
	}
	return *o.FiatNetworkInstructions
}

// GetFiatNetworkInstructionsOk returns a tuple with the FiatNetworkInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFiatAccountRequest) GetFiatNetworkInstructionsOk() (*UpdateFiatAccountRequestUpdateFiatNetworkInstructions, bool) {
	if o == nil || IsNil(o.FiatNetworkInstructions) {
		return nil, false
	}
	return o.FiatNetworkInstructions, true
}

// HasFiatNetworkInstructions returns a boolean if a field has been set.
func (o *UpdateFiatAccountRequest) HasFiatNetworkInstructions() bool {
	if o != nil && !IsNil(o.FiatNetworkInstructions) {
		return true
	}

	return false
}

// SetFiatNetworkInstructions gets a reference to the given UpdateFiatAccountRequestUpdateFiatNetworkInstructions and assigns it to the FiatNetworkInstructions field.
func (o *UpdateFiatAccountRequest) SetFiatNetworkInstructions(v UpdateFiatAccountRequestUpdateFiatNetworkInstructions) {
	o.FiatNetworkInstructions = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdateFiatAccountRequest) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFiatAccountRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdateFiatAccountRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *UpdateFiatAccountRequest) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o UpdateFiatAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFiatAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FiatAccountOwner) {
		toSerialize["fiat_account_owner"] = o.FiatAccountOwner
	}
	if !IsNil(o.FiatNetworkInstructions) {
		toSerialize["fiat_network_instructions"] = o.FiatNetworkInstructions
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateFiatAccountRequest) UnmarshalJSON(data []byte) (err error) {
	varUpdateFiatAccountRequest := _UpdateFiatAccountRequest{}

	err = json.Unmarshal(data, &varUpdateFiatAccountRequest)

	if err != nil {
		return err
	}

	*o = UpdateFiatAccountRequest(varUpdateFiatAccountRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fiat_account_owner")
		delete(additionalProperties, "fiat_network_instructions")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateFiatAccountRequest struct {
	value *UpdateFiatAccountRequest
	isSet bool
}

func (v NullableUpdateFiatAccountRequest) Get() *UpdateFiatAccountRequest {
	return v.value
}

func (v *NullableUpdateFiatAccountRequest) Set(val *UpdateFiatAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFiatAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFiatAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFiatAccountRequest(val *UpdateFiatAccountRequest) *NullableUpdateFiatAccountRequest {
	return &NullableUpdateFiatAccountRequest{value: val, isSet: true}
}

func (v NullableUpdateFiatAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFiatAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


