/*
Paxos API

<p>Welcome to Paxos APIs. At Paxos, our mission is to enable the movement of any asset, any time, in a trustworthy way. These APIs serve that mission by making it easier than ever for you to directly integrate our product capabilities into your application, leveraging the speed, stability, and security of the Paxos platform.</p> <p>The documentation that follows gives you access to our Crypto Brokerage, Trading, and Exchange products. It includes APIs for market data, orders, and the held rate quote flow.</p> <p>To test in our sandbox environment, <a href=\"https://account.sandbox.paxos.com\" target=\"_blank\">sign up</a> for an account. For more information about Paxos and our APIs, visit <a href=\"https://www.paxos.com/\" target=\"_blank\">Paxos.com</a>.</p> 

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paxos

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MailingAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MailingAddress{}

// MailingAddress A mailing address.
type MailingAddress struct {
	Country string `json:"country"`
	Address1 *string `json:"address1,omitempty"`
	Address2 *string `json:"address2,omitempty"`
	City *string `json:"city,omitempty"`
	Province *string `json:"province,omitempty"`
	ZipCode *string `json:"zip_code,omitempty"`
}

type _MailingAddress MailingAddress

// NewMailingAddress instantiates a new MailingAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMailingAddress(country string) *MailingAddress {
	this := MailingAddress{}
	this.Country = country
	return &this
}

// NewMailingAddressWithDefaults instantiates a new MailingAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMailingAddressWithDefaults() *MailingAddress {
	this := MailingAddress{}
	return &this
}

// GetCountry returns the Country field value
func (o *MailingAddress) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *MailingAddress) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *MailingAddress) SetCountry(v string) {
	o.Country = v
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *MailingAddress) GetAddress1() string {
	if o == nil || IsNil(o.Address1) {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailingAddress) GetAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.Address1) {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *MailingAddress) HasAddress1() bool {
	if o != nil && !IsNil(o.Address1) {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *MailingAddress) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *MailingAddress) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailingAddress) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *MailingAddress) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *MailingAddress) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *MailingAddress) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailingAddress) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *MailingAddress) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *MailingAddress) SetCity(v string) {
	o.City = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *MailingAddress) GetProvince() string {
	if o == nil || IsNil(o.Province) {
		var ret string
		return ret
	}
	return *o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailingAddress) GetProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.Province) {
		return nil, false
	}
	return o.Province, true
}

// HasProvince returns a boolean if a field has been set.
func (o *MailingAddress) HasProvince() bool {
	if o != nil && !IsNil(o.Province) {
		return true
	}

	return false
}

// SetProvince gets a reference to the given string and assigns it to the Province field.
func (o *MailingAddress) SetProvince(v string) {
	o.Province = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *MailingAddress) GetZipCode() string {
	if o == nil || IsNil(o.ZipCode) {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailingAddress) GetZipCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ZipCode) {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *MailingAddress) HasZipCode() bool {
	if o != nil && !IsNil(o.ZipCode) {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *MailingAddress) SetZipCode(v string) {
	o.ZipCode = &v
}

func (o MailingAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MailingAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["country"] = o.Country
	if !IsNil(o.Address1) {
		toSerialize["address1"] = o.Address1
	}
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Province) {
		toSerialize["province"] = o.Province
	}
	if !IsNil(o.ZipCode) {
		toSerialize["zip_code"] = o.ZipCode
	}
	return toSerialize, nil
}

func (o *MailingAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"country",
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

	varMailingAddress := _MailingAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMailingAddress)

	if err != nil {
		return err
	}

	*o = MailingAddress(varMailingAddress)

	return err
}

type NullableMailingAddress struct {
	value *MailingAddress
	isSet bool
}

func (v NullableMailingAddress) Get() *MailingAddress {
	return v.value
}

func (v *NullableMailingAddress) Set(val *MailingAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableMailingAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableMailingAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMailingAddress(val *MailingAddress) *NullableMailingAddress {
	return &NullableMailingAddress{value: val, isSet: true}
}

func (v NullableMailingAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMailingAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


