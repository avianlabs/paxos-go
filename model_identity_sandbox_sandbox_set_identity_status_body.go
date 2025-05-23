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

// checks if the IdentitySandboxSandboxSetIdentityStatusBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentitySandboxSandboxSetIdentityStatusBody{}

// IdentitySandboxSandboxSetIdentityStatusBody struct for IdentitySandboxSandboxSetIdentityStatusBody
type IdentitySandboxSandboxSetIdentityStatusBody struct {
	IdVerificationStatus *IdentityStatus `json:"id_verification_status,omitempty"`
	SanctionsVerificationStatus *IdentityStatus `json:"sanctions_verification_status,omitempty"`
	UserDisabled *SetDisable `json:"user_disabled,omitempty"`
	AdminDisabled *SetDisable `json:"admin_disabled,omitempty"`
	DocumentVerificationStatus *IdentityStatus `json:"document_verification_status,omitempty"`
	AdditionalScreeningStatus *IdentityStatus `json:"additional_screening_status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentitySandboxSandboxSetIdentityStatusBody IdentitySandboxSandboxSetIdentityStatusBody

// NewIdentitySandboxSandboxSetIdentityStatusBody instantiates a new IdentitySandboxSandboxSetIdentityStatusBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySandboxSandboxSetIdentityStatusBody() *IdentitySandboxSandboxSetIdentityStatusBody {
	this := IdentitySandboxSandboxSetIdentityStatusBody{}
	return &this
}

// NewIdentitySandboxSandboxSetIdentityStatusBodyWithDefaults instantiates a new IdentitySandboxSandboxSetIdentityStatusBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySandboxSandboxSetIdentityStatusBodyWithDefaults() *IdentitySandboxSandboxSetIdentityStatusBody {
	this := IdentitySandboxSandboxSetIdentityStatusBody{}
	return &this
}

// GetIdVerificationStatus returns the IdVerificationStatus field value if set, zero value otherwise.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) GetIdVerificationStatus() IdentityStatus {
	if o == nil || IsNil(o.IdVerificationStatus) {
		var ret IdentityStatus
		return ret
	}
	return *o.IdVerificationStatus
}

// GetIdVerificationStatusOk returns a tuple with the IdVerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) GetIdVerificationStatusOk() (*IdentityStatus, bool) {
	if o == nil || IsNil(o.IdVerificationStatus) {
		return nil, false
	}
	return o.IdVerificationStatus, true
}

// HasIdVerificationStatus returns a boolean if a field has been set.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) HasIdVerificationStatus() bool {
	if o != nil && !IsNil(o.IdVerificationStatus) {
		return true
	}

	return false
}

// SetIdVerificationStatus gets a reference to the given IdentityStatus and assigns it to the IdVerificationStatus field.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) SetIdVerificationStatus(v IdentityStatus) {
	o.IdVerificationStatus = &v
}

// GetSanctionsVerificationStatus returns the SanctionsVerificationStatus field value if set, zero value otherwise.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) GetSanctionsVerificationStatus() IdentityStatus {
	if o == nil || IsNil(o.SanctionsVerificationStatus) {
		var ret IdentityStatus
		return ret
	}
	return *o.SanctionsVerificationStatus
}

// GetSanctionsVerificationStatusOk returns a tuple with the SanctionsVerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) GetSanctionsVerificationStatusOk() (*IdentityStatus, bool) {
	if o == nil || IsNil(o.SanctionsVerificationStatus) {
		return nil, false
	}
	return o.SanctionsVerificationStatus, true
}

// HasSanctionsVerificationStatus returns a boolean if a field has been set.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) HasSanctionsVerificationStatus() bool {
	if o != nil && !IsNil(o.SanctionsVerificationStatus) {
		return true
	}

	return false
}

// SetSanctionsVerificationStatus gets a reference to the given IdentityStatus and assigns it to the SanctionsVerificationStatus field.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) SetSanctionsVerificationStatus(v IdentityStatus) {
	o.SanctionsVerificationStatus = &v
}

// GetUserDisabled returns the UserDisabled field value if set, zero value otherwise.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) GetUserDisabled() SetDisable {
	if o == nil || IsNil(o.UserDisabled) {
		var ret SetDisable
		return ret
	}
	return *o.UserDisabled
}

// GetUserDisabledOk returns a tuple with the UserDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) GetUserDisabledOk() (*SetDisable, bool) {
	if o == nil || IsNil(o.UserDisabled) {
		return nil, false
	}
	return o.UserDisabled, true
}

// HasUserDisabled returns a boolean if a field has been set.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) HasUserDisabled() bool {
	if o != nil && !IsNil(o.UserDisabled) {
		return true
	}

	return false
}

// SetUserDisabled gets a reference to the given SetDisable and assigns it to the UserDisabled field.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) SetUserDisabled(v SetDisable) {
	o.UserDisabled = &v
}

// GetAdminDisabled returns the AdminDisabled field value if set, zero value otherwise.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) GetAdminDisabled() SetDisable {
	if o == nil || IsNil(o.AdminDisabled) {
		var ret SetDisable
		return ret
	}
	return *o.AdminDisabled
}

// GetAdminDisabledOk returns a tuple with the AdminDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) GetAdminDisabledOk() (*SetDisable, bool) {
	if o == nil || IsNil(o.AdminDisabled) {
		return nil, false
	}
	return o.AdminDisabled, true
}

// HasAdminDisabled returns a boolean if a field has been set.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) HasAdminDisabled() bool {
	if o != nil && !IsNil(o.AdminDisabled) {
		return true
	}

	return false
}

// SetAdminDisabled gets a reference to the given SetDisable and assigns it to the AdminDisabled field.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) SetAdminDisabled(v SetDisable) {
	o.AdminDisabled = &v
}

// GetDocumentVerificationStatus returns the DocumentVerificationStatus field value if set, zero value otherwise.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) GetDocumentVerificationStatus() IdentityStatus {
	if o == nil || IsNil(o.DocumentVerificationStatus) {
		var ret IdentityStatus
		return ret
	}
	return *o.DocumentVerificationStatus
}

// GetDocumentVerificationStatusOk returns a tuple with the DocumentVerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) GetDocumentVerificationStatusOk() (*IdentityStatus, bool) {
	if o == nil || IsNil(o.DocumentVerificationStatus) {
		return nil, false
	}
	return o.DocumentVerificationStatus, true
}

// HasDocumentVerificationStatus returns a boolean if a field has been set.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) HasDocumentVerificationStatus() bool {
	if o != nil && !IsNil(o.DocumentVerificationStatus) {
		return true
	}

	return false
}

// SetDocumentVerificationStatus gets a reference to the given IdentityStatus and assigns it to the DocumentVerificationStatus field.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) SetDocumentVerificationStatus(v IdentityStatus) {
	o.DocumentVerificationStatus = &v
}

// GetAdditionalScreeningStatus returns the AdditionalScreeningStatus field value if set, zero value otherwise.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) GetAdditionalScreeningStatus() IdentityStatus {
	if o == nil || IsNil(o.AdditionalScreeningStatus) {
		var ret IdentityStatus
		return ret
	}
	return *o.AdditionalScreeningStatus
}

// GetAdditionalScreeningStatusOk returns a tuple with the AdditionalScreeningStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) GetAdditionalScreeningStatusOk() (*IdentityStatus, bool) {
	if o == nil || IsNil(o.AdditionalScreeningStatus) {
		return nil, false
	}
	return o.AdditionalScreeningStatus, true
}

// HasAdditionalScreeningStatus returns a boolean if a field has been set.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) HasAdditionalScreeningStatus() bool {
	if o != nil && !IsNil(o.AdditionalScreeningStatus) {
		return true
	}

	return false
}

// SetAdditionalScreeningStatus gets a reference to the given IdentityStatus and assigns it to the AdditionalScreeningStatus field.
func (o *IdentitySandboxSandboxSetIdentityStatusBody) SetAdditionalScreeningStatus(v IdentityStatus) {
	o.AdditionalScreeningStatus = &v
}

func (o IdentitySandboxSandboxSetIdentityStatusBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentitySandboxSandboxSetIdentityStatusBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdVerificationStatus) {
		toSerialize["id_verification_status"] = o.IdVerificationStatus
	}
	if !IsNil(o.SanctionsVerificationStatus) {
		toSerialize["sanctions_verification_status"] = o.SanctionsVerificationStatus
	}
	if !IsNil(o.UserDisabled) {
		toSerialize["user_disabled"] = o.UserDisabled
	}
	if !IsNil(o.AdminDisabled) {
		toSerialize["admin_disabled"] = o.AdminDisabled
	}
	if !IsNil(o.DocumentVerificationStatus) {
		toSerialize["document_verification_status"] = o.DocumentVerificationStatus
	}
	if !IsNil(o.AdditionalScreeningStatus) {
		toSerialize["additional_screening_status"] = o.AdditionalScreeningStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentitySandboxSandboxSetIdentityStatusBody) UnmarshalJSON(data []byte) (err error) {
	varIdentitySandboxSandboxSetIdentityStatusBody := _IdentitySandboxSandboxSetIdentityStatusBody{}

	err = json.Unmarshal(data, &varIdentitySandboxSandboxSetIdentityStatusBody)

	if err != nil {
		return err
	}

	*o = IdentitySandboxSandboxSetIdentityStatusBody(varIdentitySandboxSandboxSetIdentityStatusBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id_verification_status")
		delete(additionalProperties, "sanctions_verification_status")
		delete(additionalProperties, "user_disabled")
		delete(additionalProperties, "admin_disabled")
		delete(additionalProperties, "document_verification_status")
		delete(additionalProperties, "additional_screening_status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentitySandboxSandboxSetIdentityStatusBody struct {
	value *IdentitySandboxSandboxSetIdentityStatusBody
	isSet bool
}

func (v NullableIdentitySandboxSandboxSetIdentityStatusBody) Get() *IdentitySandboxSandboxSetIdentityStatusBody {
	return v.value
}

func (v *NullableIdentitySandboxSandboxSetIdentityStatusBody) Set(val *IdentitySandboxSandboxSetIdentityStatusBody) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySandboxSandboxSetIdentityStatusBody) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySandboxSandboxSetIdentityStatusBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySandboxSandboxSetIdentityStatusBody(val *IdentitySandboxSandboxSetIdentityStatusBody) *NullableIdentitySandboxSandboxSetIdentityStatusBody {
	return &NullableIdentitySandboxSandboxSetIdentityStatusBody{value: val, isSet: true}
}

func (v NullableIdentitySandboxSandboxSetIdentityStatusBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySandboxSandboxSetIdentityStatusBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


