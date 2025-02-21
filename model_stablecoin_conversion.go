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

// checks if the StablecoinConversion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StablecoinConversion{}

// StablecoinConversion struct for StablecoinConversion
type StablecoinConversion struct {
	// System provided UUID for the conversion is provided in the [Create Stablecoin Conversion](#operation/CreateStablecoinConversion) response.  Required parameter for the <a href=\"#operation/GetStablecoinConversion\">Get Stablecoin Conversion</a> request.
	Id *string `json:"id,omitempty"`
	// The Profile associated with a conversion. Required in the <a href=\"#operation/CreateStablecoinConversion\">Create Stablecoin Conversion</a> request.
	ProfileId *string `json:"profile_id,omitempty"`
	// Asset amount to convert. <a href=\"https://docs.paxos.com/developer/convert#assets\">Details</a>.
	Amount *string `json:"amount,omitempty"`
	// The asset to convert from. <a href=\"https://docs.paxos.com/developer/convert#assets\">Details</a>.
	SourceAsset *string `json:"source_asset,omitempty"`
	// The asset to convert to. <a href=\"https://docs.paxos.com/developer/convert#assets\">Details</a>.
	TargetAsset *string `json:"target_asset,omitempty"`
	// The current status of the conversion. <a href=\"https://docs.paxos.com/developer/convert#stablecoin-conversion-statuses\">Details</a>.
	Status *string `json:"status,omitempty"`
	// Client provided, unique Reference ID included the <a href=\"#operation/CreateStablecoinConversion\">Create Stablecoin Conversion</a> request.
	RefId *string `json:"ref_id,omitempty"`
	// The Identity ID associated with the user requesting the conversion. Required only for customers with [3rd-Party integrations](/crypto-brokerage/ledger-type#fiat-and-crypto-subledger).
	IdentityId *string `json:"identity_id,omitempty"`
	// The Account ID associated with the user requesting the conversion. Required only for customers with [3rd-Party integrations](/crypto-brokerage/ledger-type#fiat-and-crypto-subledger).
	AccountId *string `json:"account_id,omitempty"`
	// The time at which the conversion was requested. See RFC3339 format, like `2006-01-02T15:04:05Z`.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The time at which the conversion was last updated. RFC3339 format, like `2006-01-02T15:04:05Z`.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The time at which the conversion was settled. <a href=\"https://docs.paxos.com/developer/convert#stablecoin-conversion-statuses\">Details</a>. RFC3339 format, like `2006-01-02T15:04:05Z`.
	SettledAt *time.Time `json:"settled_at,omitempty"`
	// The time at which the conversion has been cancelled. <a href=\"https://docs.paxos.com/developer/convert#stablecoin-conversion-statuses\">Details</a>. RFC3339 format, like `2006-01-02T15:04:05Z`.
	CancelledAt *time.Time `json:"cancelled_at,omitempty"`
	// Optional client-specified stored metadata.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// For directed settlement, the receiving side `profile_id`.
	RecipientProfileId *string `json:"recipient_profile_id,omitempty"`
}

// NewStablecoinConversion instantiates a new StablecoinConversion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStablecoinConversion() *StablecoinConversion {
	this := StablecoinConversion{}
	return &this
}

// NewStablecoinConversionWithDefaults instantiates a new StablecoinConversion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStablecoinConversionWithDefaults() *StablecoinConversion {
	this := StablecoinConversion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StablecoinConversion) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StablecoinConversion) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StablecoinConversion) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StablecoinConversion) SetId(v string) {
	o.Id = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *StablecoinConversion) GetProfileId() string {
	if o == nil || IsNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StablecoinConversion) GetProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *StablecoinConversion) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *StablecoinConversion) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *StablecoinConversion) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StablecoinConversion) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *StablecoinConversion) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *StablecoinConversion) SetAmount(v string) {
	o.Amount = &v
}

// GetSourceAsset returns the SourceAsset field value if set, zero value otherwise.
func (o *StablecoinConversion) GetSourceAsset() string {
	if o == nil || IsNil(o.SourceAsset) {
		var ret string
		return ret
	}
	return *o.SourceAsset
}

// GetSourceAssetOk returns a tuple with the SourceAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StablecoinConversion) GetSourceAssetOk() (*string, bool) {
	if o == nil || IsNil(o.SourceAsset) {
		return nil, false
	}
	return o.SourceAsset, true
}

// HasSourceAsset returns a boolean if a field has been set.
func (o *StablecoinConversion) HasSourceAsset() bool {
	if o != nil && !IsNil(o.SourceAsset) {
		return true
	}

	return false
}

// SetSourceAsset gets a reference to the given string and assigns it to the SourceAsset field.
func (o *StablecoinConversion) SetSourceAsset(v string) {
	o.SourceAsset = &v
}

// GetTargetAsset returns the TargetAsset field value if set, zero value otherwise.
func (o *StablecoinConversion) GetTargetAsset() string {
	if o == nil || IsNil(o.TargetAsset) {
		var ret string
		return ret
	}
	return *o.TargetAsset
}

// GetTargetAssetOk returns a tuple with the TargetAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StablecoinConversion) GetTargetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.TargetAsset) {
		return nil, false
	}
	return o.TargetAsset, true
}

// HasTargetAsset returns a boolean if a field has been set.
func (o *StablecoinConversion) HasTargetAsset() bool {
	if o != nil && !IsNil(o.TargetAsset) {
		return true
	}

	return false
}

// SetTargetAsset gets a reference to the given string and assigns it to the TargetAsset field.
func (o *StablecoinConversion) SetTargetAsset(v string) {
	o.TargetAsset = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StablecoinConversion) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StablecoinConversion) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StablecoinConversion) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StablecoinConversion) SetStatus(v string) {
	o.Status = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *StablecoinConversion) GetRefId() string {
	if o == nil || IsNil(o.RefId) {
		var ret string
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StablecoinConversion) GetRefIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefId) {
		return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *StablecoinConversion) HasRefId() bool {
	if o != nil && !IsNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given string and assigns it to the RefId field.
func (o *StablecoinConversion) SetRefId(v string) {
	o.RefId = &v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *StablecoinConversion) GetIdentityId() string {
	if o == nil || IsNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StablecoinConversion) GetIdentityIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *StablecoinConversion) HasIdentityId() bool {
	if o != nil && !IsNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *StablecoinConversion) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *StablecoinConversion) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StablecoinConversion) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *StablecoinConversion) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *StablecoinConversion) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *StablecoinConversion) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StablecoinConversion) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *StablecoinConversion) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *StablecoinConversion) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *StablecoinConversion) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StablecoinConversion) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *StablecoinConversion) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *StablecoinConversion) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetSettledAt returns the SettledAt field value if set, zero value otherwise.
func (o *StablecoinConversion) GetSettledAt() time.Time {
	if o == nil || IsNil(o.SettledAt) {
		var ret time.Time
		return ret
	}
	return *o.SettledAt
}

// GetSettledAtOk returns a tuple with the SettledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StablecoinConversion) GetSettledAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SettledAt) {
		return nil, false
	}
	return o.SettledAt, true
}

// HasSettledAt returns a boolean if a field has been set.
func (o *StablecoinConversion) HasSettledAt() bool {
	if o != nil && !IsNil(o.SettledAt) {
		return true
	}

	return false
}

// SetSettledAt gets a reference to the given time.Time and assigns it to the SettledAt field.
func (o *StablecoinConversion) SetSettledAt(v time.Time) {
	o.SettledAt = &v
}

// GetCancelledAt returns the CancelledAt field value if set, zero value otherwise.
func (o *StablecoinConversion) GetCancelledAt() time.Time {
	if o == nil || IsNil(o.CancelledAt) {
		var ret time.Time
		return ret
	}
	return *o.CancelledAt
}

// GetCancelledAtOk returns a tuple with the CancelledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StablecoinConversion) GetCancelledAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CancelledAt) {
		return nil, false
	}
	return o.CancelledAt, true
}

// HasCancelledAt returns a boolean if a field has been set.
func (o *StablecoinConversion) HasCancelledAt() bool {
	if o != nil && !IsNil(o.CancelledAt) {
		return true
	}

	return false
}

// SetCancelledAt gets a reference to the given time.Time and assigns it to the CancelledAt field.
func (o *StablecoinConversion) SetCancelledAt(v time.Time) {
	o.CancelledAt = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *StablecoinConversion) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StablecoinConversion) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *StablecoinConversion) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *StablecoinConversion) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetRecipientProfileId returns the RecipientProfileId field value if set, zero value otherwise.
func (o *StablecoinConversion) GetRecipientProfileId() string {
	if o == nil || IsNil(o.RecipientProfileId) {
		var ret string
		return ret
	}
	return *o.RecipientProfileId
}

// GetRecipientProfileIdOk returns a tuple with the RecipientProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StablecoinConversion) GetRecipientProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientProfileId) {
		return nil, false
	}
	return o.RecipientProfileId, true
}

// HasRecipientProfileId returns a boolean if a field has been set.
func (o *StablecoinConversion) HasRecipientProfileId() bool {
	if o != nil && !IsNil(o.RecipientProfileId) {
		return true
	}

	return false
}

// SetRecipientProfileId gets a reference to the given string and assigns it to the RecipientProfileId field.
func (o *StablecoinConversion) SetRecipientProfileId(v string) {
	o.RecipientProfileId = &v
}

func (o StablecoinConversion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StablecoinConversion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProfileId) {
		toSerialize["profile_id"] = o.ProfileId
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.SourceAsset) {
		toSerialize["source_asset"] = o.SourceAsset
	}
	if !IsNil(o.TargetAsset) {
		toSerialize["target_asset"] = o.TargetAsset
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
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
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.SettledAt) {
		toSerialize["settled_at"] = o.SettledAt
	}
	if !IsNil(o.CancelledAt) {
		toSerialize["cancelled_at"] = o.CancelledAt
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.RecipientProfileId) {
		toSerialize["recipient_profile_id"] = o.RecipientProfileId
	}
	return toSerialize, nil
}

type NullableStablecoinConversion struct {
	value *StablecoinConversion
	isSet bool
}

func (v NullableStablecoinConversion) Get() *StablecoinConversion {
	return v.value
}

func (v *NullableStablecoinConversion) Set(val *StablecoinConversion) {
	v.value = val
	v.isSet = true
}

func (v NullableStablecoinConversion) IsSet() bool {
	return v.isSet
}

func (v *NullableStablecoinConversion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStablecoinConversion(val *StablecoinConversion) *NullableStablecoinConversion {
	return &NullableStablecoinConversion{value: val, isSet: true}
}

func (v NullableStablecoinConversion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStablecoinConversion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


