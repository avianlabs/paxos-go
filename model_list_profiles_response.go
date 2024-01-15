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

// checks if the ListProfilesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListProfilesResponse{}

// ListProfilesResponse struct for ListProfilesResponse
type ListProfilesResponse struct {
	Items []Profile `json:"items,omitempty"`
	// Cursor token required for fetching the next page.
	NextPageCursor *string `json:"next_page_cursor,omitempty"`
}

// NewListProfilesResponse instantiates a new ListProfilesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListProfilesResponse() *ListProfilesResponse {
	this := ListProfilesResponse{}
	return &this
}

// NewListProfilesResponseWithDefaults instantiates a new ListProfilesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListProfilesResponseWithDefaults() *ListProfilesResponse {
	this := ListProfilesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListProfilesResponse) GetItems() []Profile {
	if o == nil || IsNil(o.Items) {
		var ret []Profile
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListProfilesResponse) GetItemsOk() ([]Profile, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListProfilesResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Profile and assigns it to the Items field.
func (o *ListProfilesResponse) SetItems(v []Profile) {
	o.Items = v
}

// GetNextPageCursor returns the NextPageCursor field value if set, zero value otherwise.
func (o *ListProfilesResponse) GetNextPageCursor() string {
	if o == nil || IsNil(o.NextPageCursor) {
		var ret string
		return ret
	}
	return *o.NextPageCursor
}

// GetNextPageCursorOk returns a tuple with the NextPageCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListProfilesResponse) GetNextPageCursorOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageCursor) {
		return nil, false
	}
	return o.NextPageCursor, true
}

// HasNextPageCursor returns a boolean if a field has been set.
func (o *ListProfilesResponse) HasNextPageCursor() bool {
	if o != nil && !IsNil(o.NextPageCursor) {
		return true
	}

	return false
}

// SetNextPageCursor gets a reference to the given string and assigns it to the NextPageCursor field.
func (o *ListProfilesResponse) SetNextPageCursor(v string) {
	o.NextPageCursor = &v
}

func (o ListProfilesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListProfilesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.NextPageCursor) {
		toSerialize["next_page_cursor"] = o.NextPageCursor
	}
	return toSerialize, nil
}

type NullableListProfilesResponse struct {
	value *ListProfilesResponse
	isSet bool
}

func (v NullableListProfilesResponse) Get() *ListProfilesResponse {
	return v.value
}

func (v *NullableListProfilesResponse) Set(val *ListProfilesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListProfilesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListProfilesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListProfilesResponse(val *ListProfilesResponse) *NullableListProfilesResponse {
	return &NullableListProfilesResponse{value: val, isSet: true}
}

func (v NullableListProfilesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListProfilesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


