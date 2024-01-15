# TimestampRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Begin** | Pointer to **time.Time** | Only return records after this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | [optional] 
**End** | Pointer to **time.Time** | Only return records before this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | [optional] 

## Methods

### NewTimestampRange

`func NewTimestampRange() *TimestampRange`

NewTimestampRange instantiates a new TimestampRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimestampRangeWithDefaults

`func NewTimestampRangeWithDefaults() *TimestampRange`

NewTimestampRangeWithDefaults instantiates a new TimestampRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBegin

`func (o *TimestampRange) GetBegin() time.Time`

GetBegin returns the Begin field if non-nil, zero value otherwise.

### GetBeginOk

`func (o *TimestampRange) GetBeginOk() (*time.Time, bool)`

GetBeginOk returns a tuple with the Begin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBegin

`func (o *TimestampRange) SetBegin(v time.Time)`

SetBegin sets Begin field to given value.

### HasBegin

`func (o *TimestampRange) HasBegin() bool`

HasBegin returns a boolean if a field has been set.

### GetEnd

`func (o *TimestampRange) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TimestampRange) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TimestampRange) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *TimestampRange) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


