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

// checks if the ExchangePublicCreateOrderBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangePublicCreateOrderBody{}

// ExchangePublicCreateOrderBody struct for ExchangePublicCreateOrderBody
type ExchangePublicCreateOrderBody struct {
	// The idempotence ID for order creation. Can be reused if the order has been closed for more than 24 hours.
	RefId *string `json:"ref_id,omitempty"`
	Side OrderSide `json:"side"`
	Market Market `json:"market"`
	Type OrderType `json:"type"`
	// The base currency amount for any limit order or the exact amount to sell for a market sell order.
	BaseAmount *string `json:"base_amount,omitempty" validate:"regexp=^[0-9]*\\\\.?[0-9]+$"`
	// The quote price.
	Price *string `json:"price,omitempty"`
	// The quote currency amount of purchase for a market buy order.
	QuoteAmount *string `json:"quote_amount,omitempty" validate:"regexp=^[0-9]*\\\\.?[0-9]+$"`
	// Metadata to store on the quote and created order. Up to 6 key/value pairs may be stored, with each key and value at most 100 characters.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// The amount of time to wait for the order to fill, in milliseconds. When `await_fill_millis` is set to a non-zero value, the Create Order call does not return immediately on order creation. Instead, the call blocks until either:   1. The order has filled completely   2. The time `await_fill_millis` has elapsed The maximum wait timeout is 10 seconds (10000 milliseconds).
	AwaitFillMillis *int64 `json:"await_fill_millis,omitempty"`
	TimeInForce *TimeInForce `json:"time_in_force,omitempty"`
	// The date the order will expire if not completed when specified time in force is GTT. Format is a unix timestamp in milliseconds (13-digits) UTC (total milliseconds that have elapsed since January 1st, 1970 UTC).
	ExpirationDate *string `json:"expiration_date,omitempty"`
	// The end user that requests the trade. This field must be used in conjunction with `identity_account_id`, otherwise the order is rejected. Depending on your integration type, `identity_id` and `identity_account_id` may be required.
	IdentityId *string `json:"identity_id,omitempty"`
	// The account under which this order is placed. The provided identity must be allowed to trade on behalf of this account. This field must be used in conjunction with `identity_id`, otherwise the order is rejected. Depending on your integration type, `identity_account_id` and `identity_id` may be required.
	IdentityAccountId *string `json:"identity_account_id,omitempty"`
	StopPrice *string `json:"stop_price,omitempty"`
	// The profileId that will receive settled currency (base for buy orders, quote for sell orders).
	RecipientProfileId *string `json:"recipient_profile_id,omitempty"`
	// The string field used to prevent matching against an opposite side order submitted by the same Crypto Brokerage customer. If this field is not submitted, an order that matches against another order submitted by the same customer will cancel the original resting order. Up to 36 characters are supported. This field requires additional permissions only available to certain accounts. Reach out to your Paxos Representative for more information.
	SelfMatchPreventionId *string `json:"self_match_prevention_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExchangePublicCreateOrderBody ExchangePublicCreateOrderBody

// NewExchangePublicCreateOrderBody instantiates a new ExchangePublicCreateOrderBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangePublicCreateOrderBody(side OrderSide, market Market, type_ OrderType) *ExchangePublicCreateOrderBody {
	this := ExchangePublicCreateOrderBody{}
	this.Side = side
	this.Market = market
	this.Type = type_
	return &this
}

// NewExchangePublicCreateOrderBodyWithDefaults instantiates a new ExchangePublicCreateOrderBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangePublicCreateOrderBodyWithDefaults() *ExchangePublicCreateOrderBody {
	this := ExchangePublicCreateOrderBody{}
	return &this
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *ExchangePublicCreateOrderBody) GetRefId() string {
	if o == nil || IsNil(o.RefId) {
		var ret string
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangePublicCreateOrderBody) GetRefIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefId) {
		return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *ExchangePublicCreateOrderBody) HasRefId() bool {
	if o != nil && !IsNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given string and assigns it to the RefId field.
func (o *ExchangePublicCreateOrderBody) SetRefId(v string) {
	o.RefId = &v
}

// GetSide returns the Side field value
func (o *ExchangePublicCreateOrderBody) GetSide() OrderSide {
	if o == nil {
		var ret OrderSide
		return ret
	}

	return o.Side
}

// GetSideOk returns a tuple with the Side field value
// and a boolean to check if the value has been set.
func (o *ExchangePublicCreateOrderBody) GetSideOk() (*OrderSide, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Side, true
}

// SetSide sets field value
func (o *ExchangePublicCreateOrderBody) SetSide(v OrderSide) {
	o.Side = v
}

// GetMarket returns the Market field value
func (o *ExchangePublicCreateOrderBody) GetMarket() Market {
	if o == nil {
		var ret Market
		return ret
	}

	return o.Market
}

// GetMarketOk returns a tuple with the Market field value
// and a boolean to check if the value has been set.
func (o *ExchangePublicCreateOrderBody) GetMarketOk() (*Market, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Market, true
}

// SetMarket sets field value
func (o *ExchangePublicCreateOrderBody) SetMarket(v Market) {
	o.Market = v
}

// GetType returns the Type field value
func (o *ExchangePublicCreateOrderBody) GetType() OrderType {
	if o == nil {
		var ret OrderType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExchangePublicCreateOrderBody) GetTypeOk() (*OrderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExchangePublicCreateOrderBody) SetType(v OrderType) {
	o.Type = v
}

// GetBaseAmount returns the BaseAmount field value if set, zero value otherwise.
func (o *ExchangePublicCreateOrderBody) GetBaseAmount() string {
	if o == nil || IsNil(o.BaseAmount) {
		var ret string
		return ret
	}
	return *o.BaseAmount
}

// GetBaseAmountOk returns a tuple with the BaseAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangePublicCreateOrderBody) GetBaseAmountOk() (*string, bool) {
	if o == nil || IsNil(o.BaseAmount) {
		return nil, false
	}
	return o.BaseAmount, true
}

// HasBaseAmount returns a boolean if a field has been set.
func (o *ExchangePublicCreateOrderBody) HasBaseAmount() bool {
	if o != nil && !IsNil(o.BaseAmount) {
		return true
	}

	return false
}

// SetBaseAmount gets a reference to the given string and assigns it to the BaseAmount field.
func (o *ExchangePublicCreateOrderBody) SetBaseAmount(v string) {
	o.BaseAmount = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ExchangePublicCreateOrderBody) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangePublicCreateOrderBody) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ExchangePublicCreateOrderBody) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *ExchangePublicCreateOrderBody) SetPrice(v string) {
	o.Price = &v
}

// GetQuoteAmount returns the QuoteAmount field value if set, zero value otherwise.
func (o *ExchangePublicCreateOrderBody) GetQuoteAmount() string {
	if o == nil || IsNil(o.QuoteAmount) {
		var ret string
		return ret
	}
	return *o.QuoteAmount
}

// GetQuoteAmountOk returns a tuple with the QuoteAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangePublicCreateOrderBody) GetQuoteAmountOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteAmount) {
		return nil, false
	}
	return o.QuoteAmount, true
}

// HasQuoteAmount returns a boolean if a field has been set.
func (o *ExchangePublicCreateOrderBody) HasQuoteAmount() bool {
	if o != nil && !IsNil(o.QuoteAmount) {
		return true
	}

	return false
}

// SetQuoteAmount gets a reference to the given string and assigns it to the QuoteAmount field.
func (o *ExchangePublicCreateOrderBody) SetQuoteAmount(v string) {
	o.QuoteAmount = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ExchangePublicCreateOrderBody) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangePublicCreateOrderBody) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ExchangePublicCreateOrderBody) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *ExchangePublicCreateOrderBody) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetAwaitFillMillis returns the AwaitFillMillis field value if set, zero value otherwise.
func (o *ExchangePublicCreateOrderBody) GetAwaitFillMillis() int64 {
	if o == nil || IsNil(o.AwaitFillMillis) {
		var ret int64
		return ret
	}
	return *o.AwaitFillMillis
}

// GetAwaitFillMillisOk returns a tuple with the AwaitFillMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangePublicCreateOrderBody) GetAwaitFillMillisOk() (*int64, bool) {
	if o == nil || IsNil(o.AwaitFillMillis) {
		return nil, false
	}
	return o.AwaitFillMillis, true
}

// HasAwaitFillMillis returns a boolean if a field has been set.
func (o *ExchangePublicCreateOrderBody) HasAwaitFillMillis() bool {
	if o != nil && !IsNil(o.AwaitFillMillis) {
		return true
	}

	return false
}

// SetAwaitFillMillis gets a reference to the given int64 and assigns it to the AwaitFillMillis field.
func (o *ExchangePublicCreateOrderBody) SetAwaitFillMillis(v int64) {
	o.AwaitFillMillis = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *ExchangePublicCreateOrderBody) GetTimeInForce() TimeInForce {
	if o == nil || IsNil(o.TimeInForce) {
		var ret TimeInForce
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangePublicCreateOrderBody) GetTimeInForceOk() (*TimeInForce, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *ExchangePublicCreateOrderBody) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given TimeInForce and assigns it to the TimeInForce field.
func (o *ExchangePublicCreateOrderBody) SetTimeInForce(v TimeInForce) {
	o.TimeInForce = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *ExchangePublicCreateOrderBody) GetExpirationDate() string {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangePublicCreateOrderBody) GetExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *ExchangePublicCreateOrderBody) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *ExchangePublicCreateOrderBody) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *ExchangePublicCreateOrderBody) GetIdentityId() string {
	if o == nil || IsNil(o.IdentityId) {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangePublicCreateOrderBody) GetIdentityIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityId) {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *ExchangePublicCreateOrderBody) HasIdentityId() bool {
	if o != nil && !IsNil(o.IdentityId) {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *ExchangePublicCreateOrderBody) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetIdentityAccountId returns the IdentityAccountId field value if set, zero value otherwise.
func (o *ExchangePublicCreateOrderBody) GetIdentityAccountId() string {
	if o == nil || IsNil(o.IdentityAccountId) {
		var ret string
		return ret
	}
	return *o.IdentityAccountId
}

// GetIdentityAccountIdOk returns a tuple with the IdentityAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangePublicCreateOrderBody) GetIdentityAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityAccountId) {
		return nil, false
	}
	return o.IdentityAccountId, true
}

// HasIdentityAccountId returns a boolean if a field has been set.
func (o *ExchangePublicCreateOrderBody) HasIdentityAccountId() bool {
	if o != nil && !IsNil(o.IdentityAccountId) {
		return true
	}

	return false
}

// SetIdentityAccountId gets a reference to the given string and assigns it to the IdentityAccountId field.
func (o *ExchangePublicCreateOrderBody) SetIdentityAccountId(v string) {
	o.IdentityAccountId = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *ExchangePublicCreateOrderBody) GetStopPrice() string {
	if o == nil || IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangePublicCreateOrderBody) GetStopPriceOk() (*string, bool) {
	if o == nil || IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *ExchangePublicCreateOrderBody) HasStopPrice() bool {
	if o != nil && !IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *ExchangePublicCreateOrderBody) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetRecipientProfileId returns the RecipientProfileId field value if set, zero value otherwise.
func (o *ExchangePublicCreateOrderBody) GetRecipientProfileId() string {
	if o == nil || IsNil(o.RecipientProfileId) {
		var ret string
		return ret
	}
	return *o.RecipientProfileId
}

// GetRecipientProfileIdOk returns a tuple with the RecipientProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangePublicCreateOrderBody) GetRecipientProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientProfileId) {
		return nil, false
	}
	return o.RecipientProfileId, true
}

// HasRecipientProfileId returns a boolean if a field has been set.
func (o *ExchangePublicCreateOrderBody) HasRecipientProfileId() bool {
	if o != nil && !IsNil(o.RecipientProfileId) {
		return true
	}

	return false
}

// SetRecipientProfileId gets a reference to the given string and assigns it to the RecipientProfileId field.
func (o *ExchangePublicCreateOrderBody) SetRecipientProfileId(v string) {
	o.RecipientProfileId = &v
}

// GetSelfMatchPreventionId returns the SelfMatchPreventionId field value if set, zero value otherwise.
func (o *ExchangePublicCreateOrderBody) GetSelfMatchPreventionId() string {
	if o == nil || IsNil(o.SelfMatchPreventionId) {
		var ret string
		return ret
	}
	return *o.SelfMatchPreventionId
}

// GetSelfMatchPreventionIdOk returns a tuple with the SelfMatchPreventionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangePublicCreateOrderBody) GetSelfMatchPreventionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SelfMatchPreventionId) {
		return nil, false
	}
	return o.SelfMatchPreventionId, true
}

// HasSelfMatchPreventionId returns a boolean if a field has been set.
func (o *ExchangePublicCreateOrderBody) HasSelfMatchPreventionId() bool {
	if o != nil && !IsNil(o.SelfMatchPreventionId) {
		return true
	}

	return false
}

// SetSelfMatchPreventionId gets a reference to the given string and assigns it to the SelfMatchPreventionId field.
func (o *ExchangePublicCreateOrderBody) SetSelfMatchPreventionId(v string) {
	o.SelfMatchPreventionId = &v
}

func (o ExchangePublicCreateOrderBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangePublicCreateOrderBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RefId) {
		toSerialize["ref_id"] = o.RefId
	}
	toSerialize["side"] = o.Side
	toSerialize["market"] = o.Market
	toSerialize["type"] = o.Type
	if !IsNil(o.BaseAmount) {
		toSerialize["base_amount"] = o.BaseAmount
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.QuoteAmount) {
		toSerialize["quote_amount"] = o.QuoteAmount
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.AwaitFillMillis) {
		toSerialize["await_fill_millis"] = o.AwaitFillMillis
	}
	if !IsNil(o.TimeInForce) {
		toSerialize["time_in_force"] = o.TimeInForce
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	if !IsNil(o.IdentityId) {
		toSerialize["identity_id"] = o.IdentityId
	}
	if !IsNil(o.IdentityAccountId) {
		toSerialize["identity_account_id"] = o.IdentityAccountId
	}
	if !IsNil(o.StopPrice) {
		toSerialize["stop_price"] = o.StopPrice
	}
	if !IsNil(o.RecipientProfileId) {
		toSerialize["recipient_profile_id"] = o.RecipientProfileId
	}
	if !IsNil(o.SelfMatchPreventionId) {
		toSerialize["self_match_prevention_id"] = o.SelfMatchPreventionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExchangePublicCreateOrderBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"side",
		"market",
		"type",
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

	varExchangePublicCreateOrderBody := _ExchangePublicCreateOrderBody{}

	err = json.Unmarshal(data, &varExchangePublicCreateOrderBody)

	if err != nil {
		return err
	}

	*o = ExchangePublicCreateOrderBody(varExchangePublicCreateOrderBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ref_id")
		delete(additionalProperties, "side")
		delete(additionalProperties, "market")
		delete(additionalProperties, "type")
		delete(additionalProperties, "base_amount")
		delete(additionalProperties, "price")
		delete(additionalProperties, "quote_amount")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "await_fill_millis")
		delete(additionalProperties, "time_in_force")
		delete(additionalProperties, "expiration_date")
		delete(additionalProperties, "identity_id")
		delete(additionalProperties, "identity_account_id")
		delete(additionalProperties, "stop_price")
		delete(additionalProperties, "recipient_profile_id")
		delete(additionalProperties, "self_match_prevention_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangePublicCreateOrderBody struct {
	value *ExchangePublicCreateOrderBody
	isSet bool
}

func (v NullableExchangePublicCreateOrderBody) Get() *ExchangePublicCreateOrderBody {
	return v.value
}

func (v *NullableExchangePublicCreateOrderBody) Set(val *ExchangePublicCreateOrderBody) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangePublicCreateOrderBody) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangePublicCreateOrderBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangePublicCreateOrderBody(val *ExchangePublicCreateOrderBody) *NullableExchangePublicCreateOrderBody {
	return &NullableExchangePublicCreateOrderBody{value: val, isSet: true}
}

func (v NullableExchangePublicCreateOrderBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangePublicCreateOrderBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


