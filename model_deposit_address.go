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

// checks if the DepositAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DepositAddress{}

// DepositAddress struct for DepositAddress
type DepositAddress struct {
	// The UUID of the deposit address.
	Id string `json:"id"`
	// The ID of the profile that will be credited when a deposit arrives with this address.
	ProfileId string `json:"profile_id"`
	// The Paxos Customer to which the profile belongs.
	CustomerId string `json:"customer_id"`
	CryptoNetwork CryptoNetwork `json:"crypto_network"`
	// The Identity of the end user who will make deposits to this address.
	IdentityId *string `json:"identity_id,omitempty"`
	// Client-specified ID for replay protection and lookup.
	RefId *string `json:"ref_id,omitempty"`
	// Optional client-specified metadata, which is set on deposit address creation. Up to 6 key/value pairs may be returned. Each key and value must be less than or equal to 100 characters.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// The crypto address to send to for deposits.
	Address string `json:"address"`
	// The Account associated to the identity of the user that will be linked to the created address.
	AccountId *string `json:"account_id,omitempty"`
	ConversionTargetAsset DepositAddressConversionTargetAsset `json:"conversion_target_asset"`
	// List of networks compatible with the created address. Any of these networks can be used to deposit to the address.
	CompatibleCryptoNetworks []CryptoNetwork `json:"compatible_crypto_networks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DepositAddress DepositAddress

// NewDepositAddress instantiates a new DepositAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositAddress(id string, profileId string, customerId string, cryptoNetwork CryptoNetwork, address string, conversionTargetAsset DepositAddressConversionTargetAsset) *DepositAddress {
	this := DepositAddress{}
	this.Id = id
	this.ProfileId = profileId
	this.CustomerId = customerId
	this.CryptoNetwork = cryptoNetwork
	this.Address = address
	this.ConversionTargetAsset = conversionTargetAsset
	return &this
}

// NewDepositAddressWithDefaults instantiates a new DepositAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositAddressWithDefaults() *DepositAddress {
	this := DepositAddress{}
	return &this
}

// GetId returns the Id field value
func (o *DepositAddress) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DepositAddress) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DepositAddress) SetId(v string) {
	o.Id = v
}

// GetProfileId returns the ProfileId field value
func (o *DepositAddress) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value
// and a boolean to check if the value has been set.
func (o *DepositAddress) GetProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileId, true
}

// SetProfileId sets field value
func (o *DepositAddress) SetProfileId(v string) {
	o.ProfileId = v
}

// GetCustomerId returns the CustomerId field value
func (o *DepositAddress) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *DepositAddress) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *DepositAddress) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetCryptoNetwork returns the CryptoNetwork field value
func (o *DepositAddress) GetCryptoNetwork() CryptoNetwork {
	if o == nil {
		var ret CryptoNetwork
		return ret
	}

	return o.CryptoNetwork
}

// GetCryptoNetworkOk returns a tuple with the CryptoNetwork field value
// and a boolean to check if the value has been set.
func (o *DepositAddress) GetCryptoNetworkOk() (*CryptoNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CryptoNetwork, true
}

// SetCryptoNetwork sets field value
func (o *DepositAddress) SetCryptoNetwork(v CryptoNetwork) {
	o.CryptoNetwork = v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *DepositAddress) GetIdentityId() string {
	if o == nil || IsNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositAddress) GetIdentityIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *DepositAddress) HasIdentityId() bool {
	if o != nil && !IsNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *DepositAddress) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *DepositAddress) GetRefId() string {
	if o == nil || IsNil(o.RefId) {
		var ret string
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositAddress) GetRefIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefId) {
		return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *DepositAddress) HasRefId() bool {
	if o != nil && !IsNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given string and assigns it to the RefId field.
func (o *DepositAddress) SetRefId(v string) {
	o.RefId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DepositAddress) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositAddress) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DepositAddress) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *DepositAddress) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetAddress returns the Address field value
func (o *DepositAddress) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *DepositAddress) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *DepositAddress) SetAddress(v string) {
	o.Address = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *DepositAddress) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositAddress) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *DepositAddress) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *DepositAddress) SetAccountId(v string) {
	o.AccountId = &v
}

// GetConversionTargetAsset returns the ConversionTargetAsset field value
func (o *DepositAddress) GetConversionTargetAsset() DepositAddressConversionTargetAsset {
	if o == nil {
		var ret DepositAddressConversionTargetAsset
		return ret
	}

	return o.ConversionTargetAsset
}

// GetConversionTargetAssetOk returns a tuple with the ConversionTargetAsset field value
// and a boolean to check if the value has been set.
func (o *DepositAddress) GetConversionTargetAssetOk() (*DepositAddressConversionTargetAsset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversionTargetAsset, true
}

// SetConversionTargetAsset sets field value
func (o *DepositAddress) SetConversionTargetAsset(v DepositAddressConversionTargetAsset) {
	o.ConversionTargetAsset = v
}

// GetCompatibleCryptoNetworks returns the CompatibleCryptoNetworks field value if set, zero value otherwise.
func (o *DepositAddress) GetCompatibleCryptoNetworks() []CryptoNetwork {
	if o == nil || IsNil(o.CompatibleCryptoNetworks) {
		var ret []CryptoNetwork
		return ret
	}
	return o.CompatibleCryptoNetworks
}

// GetCompatibleCryptoNetworksOk returns a tuple with the CompatibleCryptoNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositAddress) GetCompatibleCryptoNetworksOk() ([]CryptoNetwork, bool) {
	if o == nil || IsNil(o.CompatibleCryptoNetworks) {
		return nil, false
	}
	return o.CompatibleCryptoNetworks, true
}

// HasCompatibleCryptoNetworks returns a boolean if a field has been set.
func (o *DepositAddress) HasCompatibleCryptoNetworks() bool {
	if o != nil && !IsNil(o.CompatibleCryptoNetworks) {
		return true
	}

	return false
}

// SetCompatibleCryptoNetworks gets a reference to the given []CryptoNetwork and assigns it to the CompatibleCryptoNetworks field.
func (o *DepositAddress) SetCompatibleCryptoNetworks(v []CryptoNetwork) {
	o.CompatibleCryptoNetworks = v
}

func (o DepositAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["profile_id"] = o.ProfileId
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["crypto_network"] = o.CryptoNetwork
	if !IsNil(o.IdentityId) {
		toSerialize["identity_id"] = o.IdentityId
	}
	if !IsNil(o.RefId) {
		toSerialize["ref_id"] = o.RefId
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["address"] = o.Address
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	toSerialize["conversion_target_asset"] = o.ConversionTargetAsset
	if !IsNil(o.CompatibleCryptoNetworks) {
		toSerialize["compatible_crypto_networks"] = o.CompatibleCryptoNetworks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DepositAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"profile_id",
		"customer_id",
		"crypto_network",
		"address",
		"conversion_target_asset",
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

	varDepositAddress := _DepositAddress{}

	err = json.Unmarshal(data, &varDepositAddress)

	if err != nil {
		return err
	}

	*o = DepositAddress(varDepositAddress)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "profile_id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "crypto_network")
		delete(additionalProperties, "identity_id")
		delete(additionalProperties, "ref_id")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "address")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "conversion_target_asset")
		delete(additionalProperties, "compatible_crypto_networks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepositAddress struct {
	value *DepositAddress
	isSet bool
}

func (v NullableDepositAddress) Get() *DepositAddress {
	return v.value
}

func (v *NullableDepositAddress) Set(val *DepositAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositAddress(val *DepositAddress) *NullableDepositAddress {
	return &NullableDepositAddress{value: val, isSet: true}
}

func (v NullableDepositAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


