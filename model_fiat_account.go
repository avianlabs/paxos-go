/*
Paxos API

<p>Welcome to Paxos APIs. At Paxos, our mission is to enable the movement of any asset, any time, in a trustworthy way. These APIs serve that mission by making it easier than ever for you to directly integrate our product capabilities into your application, leveraging the speed, stability, and security of the Paxos platform.</p> <p>The documentation that follows gives you access to our Crypto Brokerage, Trading, and Exchange products. It includes APIs for market data, orders, and the held rate quote flow.</p> <p>To test in our sandbox environment, <a href=\"https://account.sandbox.paxos.com\" target=\"_blank\">sign up</a> for an account. For more information about Paxos and our APIs, visit <a href=\"https://www.paxos.com/\" target=\"_blank\">Paxos.com</a>.</p> 

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paxos

import (
	"encoding/json"
	"time"
)

// checks if the FiatAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FiatAccount{}

// FiatAccount struct for FiatAccount
type FiatAccount struct {
	// The Paxos FiatAccount ID (UUID).
	Id *string `json:"id,omitempty"`
	// The optional client-specified ID (for idempotence).
	RefId *string `json:"ref_id,omitempty"`
	// The Paxos Identity (`identity_id`) of the user's FiatAccount.
	IdentityId *string `json:"identity_id,omitempty"`
	// The Paxos Account (`account_id`) of the user's FiatAccount. Required only for customers with [3rd-Party integrations](https://docs.paxos.com/crypto-brokerage/ledger-type#fiat-and-crypto-subledger) making deposits on behalf of their end users.
	AccountId *string `json:"account_id,omitempty"`
	FiatAccountOwner *FiatAccountOwner `json:"fiat_account_owner,omitempty"`
	FiatNetworkInstructions *FiatNetworkInstructions `json:"fiat_network_instructions,omitempty"`
	Status *FiatAccountStatus `json:"status,omitempty"`
	// Optional client-specified metadata. Up to 6 key/value pairs may be returned. Each key and value must be less than or equal to 100 characters.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// The time at which this FiatAccount was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The time at which this FiatAccount record was most recently updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewFiatAccount instantiates a new FiatAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiatAccount() *FiatAccount {
	this := FiatAccount{}
	return &this
}

// NewFiatAccountWithDefaults instantiates a new FiatAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiatAccountWithDefaults() *FiatAccount {
	this := FiatAccount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FiatAccount) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatAccount) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FiatAccount) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FiatAccount) SetId(v string) {
	o.Id = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *FiatAccount) GetRefId() string {
	if o == nil || IsNil(o.RefId) {
		var ret string
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatAccount) GetRefIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefId) {
		return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *FiatAccount) HasRefId() bool {
	if o != nil && !IsNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given string and assigns it to the RefId field.
func (o *FiatAccount) SetRefId(v string) {
	o.RefId = &v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *FiatAccount) GetIdentityId() string {
	if o == nil || IsNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatAccount) GetIdentityIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *FiatAccount) HasIdentityId() bool {
	if o != nil && !IsNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *FiatAccount) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *FiatAccount) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatAccount) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *FiatAccount) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *FiatAccount) SetAccountId(v string) {
	o.AccountId = &v
}

// GetFiatAccountOwner returns the FiatAccountOwner field value if set, zero value otherwise.
func (o *FiatAccount) GetFiatAccountOwner() FiatAccountOwner {
	if o == nil || IsNil(o.FiatAccountOwner) {
		var ret FiatAccountOwner
		return ret
	}
	return *o.FiatAccountOwner
}

// GetFiatAccountOwnerOk returns a tuple with the FiatAccountOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatAccount) GetFiatAccountOwnerOk() (*FiatAccountOwner, bool) {
	if o == nil || IsNil(o.FiatAccountOwner) {
		return nil, false
	}
	return o.FiatAccountOwner, true
}

// HasFiatAccountOwner returns a boolean if a field has been set.
func (o *FiatAccount) HasFiatAccountOwner() bool {
	if o != nil && !IsNil(o.FiatAccountOwner) {
		return true
	}

	return false
}

// SetFiatAccountOwner gets a reference to the given FiatAccountOwner and assigns it to the FiatAccountOwner field.
func (o *FiatAccount) SetFiatAccountOwner(v FiatAccountOwner) {
	o.FiatAccountOwner = &v
}

// GetFiatNetworkInstructions returns the FiatNetworkInstructions field value if set, zero value otherwise.
func (o *FiatAccount) GetFiatNetworkInstructions() FiatNetworkInstructions {
	if o == nil || IsNil(o.FiatNetworkInstructions) {
		var ret FiatNetworkInstructions
		return ret
	}
	return *o.FiatNetworkInstructions
}

// GetFiatNetworkInstructionsOk returns a tuple with the FiatNetworkInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatAccount) GetFiatNetworkInstructionsOk() (*FiatNetworkInstructions, bool) {
	if o == nil || IsNil(o.FiatNetworkInstructions) {
		return nil, false
	}
	return o.FiatNetworkInstructions, true
}

// HasFiatNetworkInstructions returns a boolean if a field has been set.
func (o *FiatAccount) HasFiatNetworkInstructions() bool {
	if o != nil && !IsNil(o.FiatNetworkInstructions) {
		return true
	}

	return false
}

// SetFiatNetworkInstructions gets a reference to the given FiatNetworkInstructions and assigns it to the FiatNetworkInstructions field.
func (o *FiatAccount) SetFiatNetworkInstructions(v FiatNetworkInstructions) {
	o.FiatNetworkInstructions = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FiatAccount) GetStatus() FiatAccountStatus {
	if o == nil || IsNil(o.Status) {
		var ret FiatAccountStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatAccount) GetStatusOk() (*FiatAccountStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FiatAccount) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given FiatAccountStatus and assigns it to the Status field.
func (o *FiatAccount) SetStatus(v FiatAccountStatus) {
	o.Status = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *FiatAccount) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatAccount) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *FiatAccount) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *FiatAccount) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FiatAccount) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatAccount) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FiatAccount) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FiatAccount) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FiatAccount) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiatAccount) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FiatAccount) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FiatAccount) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o FiatAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FiatAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RefId) {
		toSerialize["ref_id"] = o.RefId
	}
	if !IsNil(o.IdentityId) {
		toSerialize["identity_id"] = o.IdentityId
	}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.FiatAccountOwner) {
		toSerialize["fiat_account_owner"] = o.FiatAccountOwner
	}
	if !IsNil(o.FiatNetworkInstructions) {
		toSerialize["fiat_network_instructions"] = o.FiatNetworkInstructions
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableFiatAccount struct {
	value *FiatAccount
	isSet bool
}

func (v NullableFiatAccount) Get() *FiatAccount {
	return v.value
}

func (v *NullableFiatAccount) Set(val *FiatAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableFiatAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableFiatAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiatAccount(val *FiatAccount) *NullableFiatAccount {
	return &NullableFiatAccount{value: val, isSet: true}
}

func (v NullableFiatAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiatAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


