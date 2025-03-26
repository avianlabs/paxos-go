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

// checks if the ListAccountsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAccountsResponse{}

// ListAccountsResponse struct for ListAccountsResponse
type ListAccountsResponse struct {
	// Cursor token required for fetching the next page.
	NextPageCursor *string `json:"next_page_cursor,omitempty"`
	// The result list of accounts.
	Items []Account `json:"items,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListAccountsResponse ListAccountsResponse

// NewListAccountsResponse instantiates a new ListAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAccountsResponse() *ListAccountsResponse {
	this := ListAccountsResponse{}
	return &this
}

// NewListAccountsResponseWithDefaults instantiates a new ListAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAccountsResponseWithDefaults() *ListAccountsResponse {
	this := ListAccountsResponse{}
	return &this
}

// GetNextPageCursor returns the NextPageCursor field value if set, zero value otherwise.
func (o *ListAccountsResponse) GetNextPageCursor() string {
	if o == nil || IsNil(o.NextPageCursor) {
		var ret string
		return ret
	}
	return *o.NextPageCursor
}

// GetNextPageCursorOk returns a tuple with the NextPageCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccountsResponse) GetNextPageCursorOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageCursor) {
		return nil, false
	}
	return o.NextPageCursor, true
}

// HasNextPageCursor returns a boolean if a field has been set.
func (o *ListAccountsResponse) HasNextPageCursor() bool {
	if o != nil && !IsNil(o.NextPageCursor) {
		return true
	}

	return false
}

// SetNextPageCursor gets a reference to the given string and assigns it to the NextPageCursor field.
func (o *ListAccountsResponse) SetNextPageCursor(v string) {
	o.NextPageCursor = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListAccountsResponse) GetItems() []Account {
	if o == nil || IsNil(o.Items) {
		var ret []Account
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccountsResponse) GetItemsOk() ([]Account, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListAccountsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Account and assigns it to the Items field.
func (o *ListAccountsResponse) SetItems(v []Account) {
	o.Items = v
}

func (o ListAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAccountsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageCursor) {
		toSerialize["next_page_cursor"] = o.NextPageCursor
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListAccountsResponse) UnmarshalJSON(data []byte) (err error) {
	varListAccountsResponse := _ListAccountsResponse{}

	err = json.Unmarshal(data, &varListAccountsResponse)

	if err != nil {
		return err
	}

	*o = ListAccountsResponse(varListAccountsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next_page_cursor")
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListAccountsResponse struct {
	value *ListAccountsResponse
	isSet bool
}

func (v NullableListAccountsResponse) Get() *ListAccountsResponse {
	return v.value
}

func (v *NullableListAccountsResponse) Set(val *ListAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAccountsResponse(val *ListAccountsResponse) *NullableListAccountsResponse {
	return &NullableListAccountsResponse{value: val, isSet: true}
}

func (v NullableListAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


