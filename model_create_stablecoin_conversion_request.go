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

// checks if the CreateStablecoinConversionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateStablecoinConversionRequest{}

// CreateStablecoinConversionRequest struct for CreateStablecoinConversionRequest
type CreateStablecoinConversionRequest struct {
	// The Profile associated with a conversion.
	ProfileId string `json:"profile_id"`
	// Asset amount to convert. <a href=\"/stablecoin/conversion/asset\">Details</a>.
	Amount string `json:"amount" validate:"regexp=^[0-9]*\\\\.?[0-9]+$"`
	// The asset to convert from. <a href=\"/stablecoin/conversion/asset\">Details</a>.
	SourceAsset string `json:"source_asset"`
	// The asset to convert to. <a href=\"/stablecoin/conversion/asset\">Details</a>.
	TargetAsset string `json:"target_asset"`
	// Client provided, unique Reference ID for lookup and replay protection.
	RefId *string `json:"ref_id,omitempty"`
	// The Identity ID associated with the user requesting the conversion. Required only for customers with [3rd-Party integrations](/crypto-brokerage/ledger-type#fiat-and-crypto-subledger).
	IdentityId *string `json:"identity_id,omitempty"`
	// The Account ID associated with the user requesting the conversion. Required only for customers with [3rd-Party integrations](/crypto-brokerage/ledger-type#fiat-and-crypto-subledger).
	AccountId *string `json:"account_id,omitempty"`
	// Optional client-specified stored metadata.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// For directed settlement, the receiving side `profile_id`.
	RecipientProfileId *string `json:"recipient_profile_id,omitempty"`
}

type _CreateStablecoinConversionRequest CreateStablecoinConversionRequest

// NewCreateStablecoinConversionRequest instantiates a new CreateStablecoinConversionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStablecoinConversionRequest(profileId string, amount string, sourceAsset string, targetAsset string) *CreateStablecoinConversionRequest {
	this := CreateStablecoinConversionRequest{}
	this.ProfileId = profileId
	this.Amount = amount
	this.SourceAsset = sourceAsset
	this.TargetAsset = targetAsset
	return &this
}

// NewCreateStablecoinConversionRequestWithDefaults instantiates a new CreateStablecoinConversionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStablecoinConversionRequestWithDefaults() *CreateStablecoinConversionRequest {
	this := CreateStablecoinConversionRequest{}
	return &this
}

// GetProfileId returns the ProfileId field value
func (o *CreateStablecoinConversionRequest) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value
// and a boolean to check if the value has been set.
func (o *CreateStablecoinConversionRequest) GetProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileId, true
}

// SetProfileId sets field value
func (o *CreateStablecoinConversionRequest) SetProfileId(v string) {
	o.ProfileId = v
}

// GetAmount returns the Amount field value
func (o *CreateStablecoinConversionRequest) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreateStablecoinConversionRequest) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreateStablecoinConversionRequest) SetAmount(v string) {
	o.Amount = v
}

// GetSourceAsset returns the SourceAsset field value
func (o *CreateStablecoinConversionRequest) GetSourceAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceAsset
}

// GetSourceAssetOk returns a tuple with the SourceAsset field value
// and a boolean to check if the value has been set.
func (o *CreateStablecoinConversionRequest) GetSourceAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceAsset, true
}

// SetSourceAsset sets field value
func (o *CreateStablecoinConversionRequest) SetSourceAsset(v string) {
	o.SourceAsset = v
}

// GetTargetAsset returns the TargetAsset field value
func (o *CreateStablecoinConversionRequest) GetTargetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetAsset
}

// GetTargetAssetOk returns a tuple with the TargetAsset field value
// and a boolean to check if the value has been set.
func (o *CreateStablecoinConversionRequest) GetTargetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetAsset, true
}

// SetTargetAsset sets field value
func (o *CreateStablecoinConversionRequest) SetTargetAsset(v string) {
	o.TargetAsset = v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *CreateStablecoinConversionRequest) GetRefId() string {
	if o == nil || IsNil(o.RefId) {
		var ret string
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStablecoinConversionRequest) GetRefIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefId) {
		return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *CreateStablecoinConversionRequest) HasRefId() bool {
	if o != nil && !IsNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given string and assigns it to the RefId field.
func (o *CreateStablecoinConversionRequest) SetRefId(v string) {
	o.RefId = &v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *CreateStablecoinConversionRequest) GetIdentityId() string {
	if o == nil || IsNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStablecoinConversionRequest) GetIdentityIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *CreateStablecoinConversionRequest) HasIdentityId() bool {
	if o != nil && !IsNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *CreateStablecoinConversionRequest) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *CreateStablecoinConversionRequest) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStablecoinConversionRequest) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *CreateStablecoinConversionRequest) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *CreateStablecoinConversionRequest) SetAccountId(v string) {
	o.AccountId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateStablecoinConversionRequest) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStablecoinConversionRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateStablecoinConversionRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CreateStablecoinConversionRequest) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetRecipientProfileId returns the RecipientProfileId field value if set, zero value otherwise.
func (o *CreateStablecoinConversionRequest) GetRecipientProfileId() string {
	if o == nil || IsNil(o.RecipientProfileId) {
		var ret string
		return ret
	}
	return *o.RecipientProfileId
}

// GetRecipientProfileIdOk returns a tuple with the RecipientProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStablecoinConversionRequest) GetRecipientProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientProfileId) {
		return nil, false
	}
	return o.RecipientProfileId, true
}

// HasRecipientProfileId returns a boolean if a field has been set.
func (o *CreateStablecoinConversionRequest) HasRecipientProfileId() bool {
	if o != nil && !IsNil(o.RecipientProfileId) {
		return true
	}

	return false
}

// SetRecipientProfileId gets a reference to the given string and assigns it to the RecipientProfileId field.
func (o *CreateStablecoinConversionRequest) SetRecipientProfileId(v string) {
	o.RecipientProfileId = &v
}

func (o CreateStablecoinConversionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateStablecoinConversionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profile_id"] = o.ProfileId
	toSerialize["amount"] = o.Amount
	toSerialize["source_asset"] = o.SourceAsset
	toSerialize["target_asset"] = o.TargetAsset
	if !IsNil(o.RefId) {
		toSerialize["ref_id"] = o.RefId
	}
	if !IsNil(o.IdentityId) {
		toSerialize["identity_id"] = o.IdentityId
	}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.RecipientProfileId) {
		toSerialize["recipient_profile_id"] = o.RecipientProfileId
	}
	return toSerialize, nil
}

func (o *CreateStablecoinConversionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"profile_id",
		"amount",
		"source_asset",
		"target_asset",
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

	varCreateStablecoinConversionRequest := _CreateStablecoinConversionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateStablecoinConversionRequest)

	if err != nil {
		return err
	}

	*o = CreateStablecoinConversionRequest(varCreateStablecoinConversionRequest)

	return err
}

type NullableCreateStablecoinConversionRequest struct {
	value *CreateStablecoinConversionRequest
	isSet bool
}

func (v NullableCreateStablecoinConversionRequest) Get() *CreateStablecoinConversionRequest {
	return v.value
}

func (v *NullableCreateStablecoinConversionRequest) Set(val *CreateStablecoinConversionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStablecoinConversionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStablecoinConversionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStablecoinConversionRequest(val *CreateStablecoinConversionRequest) *NullableCreateStablecoinConversionRequest {
	return &NullableCreateStablecoinConversionRequest{value: val, isSet: true}
}

func (v NullableCreateStablecoinConversionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateStablecoinConversionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


