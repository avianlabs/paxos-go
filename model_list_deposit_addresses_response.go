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

// checks if the ListDepositAddressesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDepositAddressesResponse{}

// ListDepositAddressesResponse struct for ListDepositAddressesResponse
type ListDepositAddressesResponse struct {
	Items []DepositAddress `json:"items,omitempty"`
	NextPageCursor *string `json:"next_page_cursor,omitempty"`
}

// NewListDepositAddressesResponse instantiates a new ListDepositAddressesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDepositAddressesResponse() *ListDepositAddressesResponse {
	this := ListDepositAddressesResponse{}
	return &this
}

// NewListDepositAddressesResponseWithDefaults instantiates a new ListDepositAddressesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDepositAddressesResponseWithDefaults() *ListDepositAddressesResponse {
	this := ListDepositAddressesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListDepositAddressesResponse) GetItems() []DepositAddress {
	if o == nil || IsNil(o.Items) {
		var ret []DepositAddress
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDepositAddressesResponse) GetItemsOk() ([]DepositAddress, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListDepositAddressesResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []DepositAddress and assigns it to the Items field.
func (o *ListDepositAddressesResponse) SetItems(v []DepositAddress) {
	o.Items = v
}

// GetNextPageCursor returns the NextPageCursor field value if set, zero value otherwise.
func (o *ListDepositAddressesResponse) GetNextPageCursor() string {
	if o == nil || IsNil(o.NextPageCursor) {
		var ret string
		return ret
	}
	return *o.NextPageCursor
}

// GetNextPageCursorOk returns a tuple with the NextPageCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDepositAddressesResponse) GetNextPageCursorOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageCursor) {
		return nil, false
	}
	return o.NextPageCursor, true
}

// HasNextPageCursor returns a boolean if a field has been set.
func (o *ListDepositAddressesResponse) HasNextPageCursor() bool {
	if o != nil && !IsNil(o.NextPageCursor) {
		return true
	}

	return false
}

// SetNextPageCursor gets a reference to the given string and assigns it to the NextPageCursor field.
func (o *ListDepositAddressesResponse) SetNextPageCursor(v string) {
	o.NextPageCursor = &v
}

func (o ListDepositAddressesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDepositAddressesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.NextPageCursor) {
		toSerialize["next_page_cursor"] = o.NextPageCursor
	}
	return toSerialize, nil
}

type NullableListDepositAddressesResponse struct {
	value *ListDepositAddressesResponse
	isSet bool
}

func (v NullableListDepositAddressesResponse) Get() *ListDepositAddressesResponse {
	return v.value
}

func (v *NullableListDepositAddressesResponse) Set(val *ListDepositAddressesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDepositAddressesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDepositAddressesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDepositAddressesResponse(val *ListDepositAddressesResponse) *NullableListDepositAddressesResponse {
	return &NullableListDepositAddressesResponse{value: val, isSet: true}
}

func (v NullableListDepositAddressesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDepositAddressesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


