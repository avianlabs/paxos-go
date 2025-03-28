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

// checks if the ListTaxFormsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTaxFormsResponse{}

// ListTaxFormsResponse struct for ListTaxFormsResponse
type ListTaxFormsResponse struct {
	// List of tax form URLs. The size shall not exceed `users_limit` times `form_types`.
	TaxFormUrls []TaxFormURL `json:"tax_form_urls,omitempty"`
	// Cursor token required for fetching the next page.
	NextPageCursor *string `json:"next_page_cursor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListTaxFormsResponse ListTaxFormsResponse

// NewListTaxFormsResponse instantiates a new ListTaxFormsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTaxFormsResponse() *ListTaxFormsResponse {
	this := ListTaxFormsResponse{}
	return &this
}

// NewListTaxFormsResponseWithDefaults instantiates a new ListTaxFormsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTaxFormsResponseWithDefaults() *ListTaxFormsResponse {
	this := ListTaxFormsResponse{}
	return &this
}

// GetTaxFormUrls returns the TaxFormUrls field value if set, zero value otherwise.
func (o *ListTaxFormsResponse) GetTaxFormUrls() []TaxFormURL {
	if o == nil || IsNil(o.TaxFormUrls) {
		var ret []TaxFormURL
		return ret
	}
	return o.TaxFormUrls
}

// GetTaxFormUrlsOk returns a tuple with the TaxFormUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaxFormsResponse) GetTaxFormUrlsOk() ([]TaxFormURL, bool) {
	if o == nil || IsNil(o.TaxFormUrls) {
		return nil, false
	}
	return o.TaxFormUrls, true
}

// HasTaxFormUrls returns a boolean if a field has been set.
func (o *ListTaxFormsResponse) HasTaxFormUrls() bool {
	if o != nil && !IsNil(o.TaxFormUrls) {
		return true
	}

	return false
}

// SetTaxFormUrls gets a reference to the given []TaxFormURL and assigns it to the TaxFormUrls field.
func (o *ListTaxFormsResponse) SetTaxFormUrls(v []TaxFormURL) {
	o.TaxFormUrls = v
}

// GetNextPageCursor returns the NextPageCursor field value if set, zero value otherwise.
func (o *ListTaxFormsResponse) GetNextPageCursor() string {
	if o == nil || IsNil(o.NextPageCursor) {
		var ret string
		return ret
	}
	return *o.NextPageCursor
}

// GetNextPageCursorOk returns a tuple with the NextPageCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTaxFormsResponse) GetNextPageCursorOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageCursor) {
		return nil, false
	}
	return o.NextPageCursor, true
}

// HasNextPageCursor returns a boolean if a field has been set.
func (o *ListTaxFormsResponse) HasNextPageCursor() bool {
	if o != nil && !IsNil(o.NextPageCursor) {
		return true
	}

	return false
}

// SetNextPageCursor gets a reference to the given string and assigns it to the NextPageCursor field.
func (o *ListTaxFormsResponse) SetNextPageCursor(v string) {
	o.NextPageCursor = &v
}

func (o ListTaxFormsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTaxFormsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaxFormUrls) {
		toSerialize["tax_form_urls"] = o.TaxFormUrls
	}
	if !IsNil(o.NextPageCursor) {
		toSerialize["next_page_cursor"] = o.NextPageCursor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListTaxFormsResponse) UnmarshalJSON(data []byte) (err error) {
	varListTaxFormsResponse := _ListTaxFormsResponse{}

	err = json.Unmarshal(data, &varListTaxFormsResponse)

	if err != nil {
		return err
	}

	*o = ListTaxFormsResponse(varListTaxFormsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tax_form_urls")
		delete(additionalProperties, "next_page_cursor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListTaxFormsResponse struct {
	value *ListTaxFormsResponse
	isSet bool
}

func (v NullableListTaxFormsResponse) Get() *ListTaxFormsResponse {
	return v.value
}

func (v *NullableListTaxFormsResponse) Set(val *ListTaxFormsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListTaxFormsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListTaxFormsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTaxFormsResponse(val *ListTaxFormsResponse) *NullableListTaxFormsResponse {
	return &NullableListTaxFormsResponse{value: val, isSet: true}
}

func (v NullableListTaxFormsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTaxFormsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


