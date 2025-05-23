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

// checks if the SecondaryStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecondaryStatus{}

// SecondaryStatus Secondary status for the transfer, used for more granular explanation of the transfer status.
type SecondaryStatus struct {
	Name *SecondaryStatusName `json:"name,omitempty"`
	// Additional information about the current status of the transfer (e.g. if information is missing).
	Detail *string `json:"detail,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecondaryStatus SecondaryStatus

// NewSecondaryStatus instantiates a new SecondaryStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecondaryStatus() *SecondaryStatus {
	this := SecondaryStatus{}
	return &this
}

// NewSecondaryStatusWithDefaults instantiates a new SecondaryStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecondaryStatusWithDefaults() *SecondaryStatus {
	this := SecondaryStatus{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecondaryStatus) GetName() SecondaryStatusName {
	if o == nil || IsNil(o.Name) {
		var ret SecondaryStatusName
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecondaryStatus) GetNameOk() (*SecondaryStatusName, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecondaryStatus) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given SecondaryStatusName and assigns it to the Name field.
func (o *SecondaryStatus) SetName(v SecondaryStatusName) {
	o.Name = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *SecondaryStatus) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecondaryStatus) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *SecondaryStatus) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *SecondaryStatus) SetDetail(v string) {
	o.Detail = &v
}

func (o SecondaryStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecondaryStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecondaryStatus) UnmarshalJSON(data []byte) (err error) {
	varSecondaryStatus := _SecondaryStatus{}

	err = json.Unmarshal(data, &varSecondaryStatus)

	if err != nil {
		return err
	}

	*o = SecondaryStatus(varSecondaryStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "detail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecondaryStatus struct {
	value *SecondaryStatus
	isSet bool
}

func (v NullableSecondaryStatus) Get() *SecondaryStatus {
	return v.value
}

func (v *NullableSecondaryStatus) Set(val *SecondaryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSecondaryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSecondaryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecondaryStatus(val *SecondaryStatus) *NullableSecondaryStatus {
	return &NullableSecondaryStatus{value: val, isSet: true}
}

func (v NullableSecondaryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecondaryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


