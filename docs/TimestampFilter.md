# TimestampFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lt** | Pointer to **time.Time** | Include timestamps strictly less than lt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | [optional] 
**Lte** | Pointer to **time.Time** | Include timestamps less than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | [optional] 
**Eq** | Pointer to **time.Time** | Include timestamps exactly equal to eq. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | [optional] 
**Gte** | Pointer to **time.Time** | Include timestamps greater than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | [optional] 
**Gt** | Pointer to **time.Time** | Include timestamps strictly greater than gt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | [optional] 

## Methods

### NewTimestampFilter

`func NewTimestampFilter() *TimestampFilter`

NewTimestampFilter instantiates a new TimestampFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimestampFilterWithDefaults

`func NewTimestampFilterWithDefaults() *TimestampFilter`

NewTimestampFilterWithDefaults instantiates a new TimestampFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLt

`func (o *TimestampFilter) GetLt() time.Time`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *TimestampFilter) GetLtOk() (*time.Time, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *TimestampFilter) SetLt(v time.Time)`

SetLt sets Lt field to given value.

### HasLt

`func (o *TimestampFilter) HasLt() bool`

HasLt returns a boolean if a field has been set.

### GetLte

`func (o *TimestampFilter) GetLte() time.Time`

GetLte returns the Lte field if non-nil, zero value otherwise.

### GetLteOk

`func (o *TimestampFilter) GetLteOk() (*time.Time, bool)`

GetLteOk returns a tuple with the Lte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLte

`func (o *TimestampFilter) SetLte(v time.Time)`

SetLte sets Lte field to given value.

### HasLte

`func (o *TimestampFilter) HasLte() bool`

HasLte returns a boolean if a field has been set.

### GetEq

`func (o *TimestampFilter) GetEq() time.Time`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *TimestampFilter) GetEqOk() (*time.Time, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *TimestampFilter) SetEq(v time.Time)`

SetEq sets Eq field to given value.

### HasEq

`func (o *TimestampFilter) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetGte

`func (o *TimestampFilter) GetGte() time.Time`

GetGte returns the Gte field if non-nil, zero value otherwise.

### GetGteOk

`func (o *TimestampFilter) GetGteOk() (*time.Time, bool)`

GetGteOk returns a tuple with the Gte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGte

`func (o *TimestampFilter) SetGte(v time.Time)`

SetGte sets Gte field to given value.

### HasGte

`func (o *TimestampFilter) HasGte() bool`

HasGte returns a boolean if a field has been set.

### GetGt

`func (o *TimestampFilter) GetGt() time.Time`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *TimestampFilter) GetGtOk() (*time.Time, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *TimestampFilter) SetGt(v time.Time)`

SetGt sets Gt field to given value.

### HasGt

`func (o *TimestampFilter) HasGt() bool`

HasGt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


