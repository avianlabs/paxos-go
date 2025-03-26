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
	"fmt"
)

// checks if the Transaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transaction{}

// Transaction Transaction returned on create (Create Transaction), list (List Transactions) and get (Get Transaction) requests.
type Transaction struct {
	// The transaction identifier used to query or refer to a specific transaction.
	Id string `json:"id"`
	// Idempotency key.
	RefId string `json:"ref_id"`
	// The start of the window which the transaction is eligible for settlement. If omitted, transactions are immediately eligible for settlement upon success. RFC3339 format, like `YYYY-MM-DDTHH:MM:SS.sssZ`. ex: `2006-01-02T15:04:05Z`.
	SettlementWindowStart time.Time `json:"settlement_window_start"`
	// The end of the window which the transaction is eligible for settlement. Transactions which are not cancelled or settled by this time will expire. RFC3339 format, like `YYYY-MM-DDTHH:MM:SS.sssZ`. ex: `2006-01-02T15:04:05Z`.
	SettlementWindowEnd time.Time `json:"settlement_window_end"`
	// The Profile ID (profile_id) of the entity submitting the transaction.
	SourceProfileId string `json:"source_profile_id"`
	// The Profile ID (profile_id) of the entity receiving the transaction.
	TargetProfileId string `json:"target_profile_id"`
	// The obligations (representing one-way asset movements) to be settled atomically.
	Legs []Obligation `json:"legs"`
	Status TransactionStatus `json:"status"`
	// The timestamp when the transaction was first created, RFC3339 format, like `YYYY-MM-DDTHH:MM:SS.sssZ`.
	CreatedAt time.Time `json:"created_at"`
	// The timestamp when the transaction was last updated, RFC3339 format, like `YYYY-MM-DDTHH:MM:SS.sssZ`.
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _Transaction Transaction

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction(id string, refId string, settlementWindowStart time.Time, settlementWindowEnd time.Time, sourceProfileId string, targetProfileId string, legs []Obligation, status TransactionStatus, createdAt time.Time, updatedAt time.Time) *Transaction {
	this := Transaction{}
	this.Id = id
	this.RefId = refId
	this.SettlementWindowStart = settlementWindowStart
	this.SettlementWindowEnd = settlementWindowEnd
	this.SourceProfileId = sourceProfileId
	this.TargetProfileId = targetProfileId
	this.Legs = legs
	this.Status = status
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	return &this
}

// GetId returns the Id field value
func (o *Transaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Transaction) SetId(v string) {
	o.Id = v
}

// GetRefId returns the RefId field value
func (o *Transaction) GetRefId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetRefIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefId, true
}

// SetRefId sets field value
func (o *Transaction) SetRefId(v string) {
	o.RefId = v
}

// GetSettlementWindowStart returns the SettlementWindowStart field value
func (o *Transaction) GetSettlementWindowStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SettlementWindowStart
}

// GetSettlementWindowStartOk returns a tuple with the SettlementWindowStart field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetSettlementWindowStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SettlementWindowStart, true
}

// SetSettlementWindowStart sets field value
func (o *Transaction) SetSettlementWindowStart(v time.Time) {
	o.SettlementWindowStart = v
}

// GetSettlementWindowEnd returns the SettlementWindowEnd field value
func (o *Transaction) GetSettlementWindowEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SettlementWindowEnd
}

// GetSettlementWindowEndOk returns a tuple with the SettlementWindowEnd field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetSettlementWindowEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SettlementWindowEnd, true
}

// SetSettlementWindowEnd sets field value
func (o *Transaction) SetSettlementWindowEnd(v time.Time) {
	o.SettlementWindowEnd = v
}

// GetSourceProfileId returns the SourceProfileId field value
func (o *Transaction) GetSourceProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceProfileId
}

// GetSourceProfileIdOk returns a tuple with the SourceProfileId field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetSourceProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceProfileId, true
}

// SetSourceProfileId sets field value
func (o *Transaction) SetSourceProfileId(v string) {
	o.SourceProfileId = v
}

// GetTargetProfileId returns the TargetProfileId field value
func (o *Transaction) GetTargetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetProfileId
}

// GetTargetProfileIdOk returns a tuple with the TargetProfileId field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetTargetProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetProfileId, true
}

// SetTargetProfileId sets field value
func (o *Transaction) SetTargetProfileId(v string) {
	o.TargetProfileId = v
}

// GetLegs returns the Legs field value
func (o *Transaction) GetLegs() []Obligation {
	if o == nil {
		var ret []Obligation
		return ret
	}

	return o.Legs
}

// GetLegsOk returns a tuple with the Legs field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetLegsOk() ([]Obligation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Legs, true
}

// SetLegs sets field value
func (o *Transaction) SetLegs(v []Obligation) {
	o.Legs = v
}

// GetStatus returns the Status field value
func (o *Transaction) GetStatus() TransactionStatus {
	if o == nil {
		var ret TransactionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetStatusOk() (*TransactionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Transaction) SetStatus(v TransactionStatus) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Transaction) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Transaction) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Transaction) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Transaction) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["ref_id"] = o.RefId
	toSerialize["settlement_window_start"] = o.SettlementWindowStart
	toSerialize["settlement_window_end"] = o.SettlementWindowEnd
	toSerialize["source_profile_id"] = o.SourceProfileId
	toSerialize["target_profile_id"] = o.TargetProfileId
	toSerialize["legs"] = o.Legs
	toSerialize["status"] = o.Status
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Transaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"ref_id",
		"settlement_window_start",
		"settlement_window_end",
		"source_profile_id",
		"target_profile_id",
		"legs",
		"status",
		"created_at",
		"updated_at",
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

	varTransaction := _Transaction{}

	err = json.Unmarshal(data, &varTransaction)

	if err != nil {
		return err
	}

	*o = Transaction(varTransaction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "ref_id")
		delete(additionalProperties, "settlement_window_start")
		delete(additionalProperties, "settlement_window_end")
		delete(additionalProperties, "source_profile_id")
		delete(additionalProperties, "target_profile_id")
		delete(additionalProperties, "legs")
		delete(additionalProperties, "status")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


