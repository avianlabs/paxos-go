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

// checks if the InstitutionMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstitutionMember{}

// InstitutionMember struct for InstitutionMember
type InstitutionMember struct {
	IdentityId *string `json:"identity_id,omitempty"`
	Roles []InstitutionRoleType `json:"roles,omitempty"`
	// Decimal number representing the percent ownership the identity has in the company--  e.g. 5 represents 5% ownership.
	Ownership *string `json:"ownership,omitempty"`
	Position *string `json:"position,omitempty"`
	Name *string `json:"name,omitempty"`
	SummaryStatus *IdentityStatus `json:"summary_status,omitempty"`
	// Institution member ID. Note: This field is auto-generated. Specifying an ID when creating an institution member is a client error.
	Id *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InstitutionMember InstitutionMember

// NewInstitutionMember instantiates a new InstitutionMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionMember() *InstitutionMember {
	this := InstitutionMember{}
	return &this
}

// NewInstitutionMemberWithDefaults instantiates a new InstitutionMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionMemberWithDefaults() *InstitutionMember {
	this := InstitutionMember{}
	return &this
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *InstitutionMember) GetIdentityId() string {
	if o == nil || IsNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionMember) GetIdentityIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *InstitutionMember) HasIdentityId() bool {
	if o != nil && !IsNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *InstitutionMember) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *InstitutionMember) GetRoles() []InstitutionRoleType {
	if o == nil || IsNil(o.Roles) {
		var ret []InstitutionRoleType
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionMember) GetRolesOk() ([]InstitutionRoleType, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *InstitutionMember) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []InstitutionRoleType and assigns it to the Roles field.
func (o *InstitutionMember) SetRoles(v []InstitutionRoleType) {
	o.Roles = v
}

// GetOwnership returns the Ownership field value if set, zero value otherwise.
func (o *InstitutionMember) GetOwnership() string {
	if o == nil || IsNil(o.Ownership) {
		var ret string
		return ret
	}
	return *o.Ownership
}

// GetOwnershipOk returns a tuple with the Ownership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionMember) GetOwnershipOk() (*string, bool) {
	if o == nil || IsNil(o.Ownership) {
		return nil, false
	}
	return o.Ownership, true
}

// HasOwnership returns a boolean if a field has been set.
func (o *InstitutionMember) HasOwnership() bool {
	if o != nil && !IsNil(o.Ownership) {
		return true
	}

	return false
}

// SetOwnership gets a reference to the given string and assigns it to the Ownership field.
func (o *InstitutionMember) SetOwnership(v string) {
	o.Ownership = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *InstitutionMember) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionMember) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *InstitutionMember) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *InstitutionMember) SetPosition(v string) {
	o.Position = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstitutionMember) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionMember) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstitutionMember) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstitutionMember) SetName(v string) {
	o.Name = &v
}

// GetSummaryStatus returns the SummaryStatus field value if set, zero value otherwise.
func (o *InstitutionMember) GetSummaryStatus() IdentityStatus {
	if o == nil || IsNil(o.SummaryStatus) {
		var ret IdentityStatus
		return ret
	}
	return *o.SummaryStatus
}

// GetSummaryStatusOk returns a tuple with the SummaryStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionMember) GetSummaryStatusOk() (*IdentityStatus, bool) {
	if o == nil || IsNil(o.SummaryStatus) {
		return nil, false
	}
	return o.SummaryStatus, true
}

// HasSummaryStatus returns a boolean if a field has been set.
func (o *InstitutionMember) HasSummaryStatus() bool {
	if o != nil && !IsNil(o.SummaryStatus) {
		return true
	}

	return false
}

// SetSummaryStatus gets a reference to the given IdentityStatus and assigns it to the SummaryStatus field.
func (o *InstitutionMember) SetSummaryStatus(v IdentityStatus) {
	o.SummaryStatus = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InstitutionMember) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionMember) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InstitutionMember) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InstitutionMember) SetId(v string) {
	o.Id = &v
}

func (o InstitutionMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstitutionMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdentityId) {
		toSerialize["identity_id"] = o.IdentityId
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Ownership) {
		toSerialize["ownership"] = o.Ownership
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SummaryStatus) {
		toSerialize["summary_status"] = o.SummaryStatus
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InstitutionMember) UnmarshalJSON(data []byte) (err error) {
	varInstitutionMember := _InstitutionMember{}

	err = json.Unmarshal(data, &varInstitutionMember)

	if err != nil {
		return err
	}

	*o = InstitutionMember(varInstitutionMember)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "identity_id")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "ownership")
		delete(additionalProperties, "position")
		delete(additionalProperties, "name")
		delete(additionalProperties, "summary_status")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstitutionMember struct {
	value *InstitutionMember
	isSet bool
}

func (v NullableInstitutionMember) Get() *InstitutionMember {
	return v.value
}

func (v *NullableInstitutionMember) Set(val *InstitutionMember) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionMember) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionMember(val *InstitutionMember) *NullableInstitutionMember {
	return &NullableInstitutionMember{value: val, isSet: true}
}

func (v NullableInstitutionMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


