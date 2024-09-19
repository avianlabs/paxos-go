# SecondaryStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**SecondaryStatusName**](SecondaryStatusName.md) |  | [optional] 
**Detail** | Pointer to **string** | Additional information about the current status of the transfer (e.g. if information is missing). | [optional] 

## Methods

### NewSecondaryStatus

`func NewSecondaryStatus() *SecondaryStatus`

NewSecondaryStatus instantiates a new SecondaryStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecondaryStatusWithDefaults

`func NewSecondaryStatusWithDefaults() *SecondaryStatus`

NewSecondaryStatusWithDefaults instantiates a new SecondaryStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecondaryStatus) GetName() SecondaryStatusName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecondaryStatus) GetNameOk() (*SecondaryStatusName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecondaryStatus) SetName(v SecondaryStatusName)`

SetName sets Name field to given value.

### HasName

`func (o *SecondaryStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDetail

`func (o *SecondaryStatus) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SecondaryStatus) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SecondaryStatus) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SecondaryStatus) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


