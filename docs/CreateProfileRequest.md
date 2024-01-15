# CreateProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nickname** | **string** | The display name of the profile. Must be unique. | 
**Description** | Pointer to **string** | The description of the created profile. | [optional] 

## Methods

### NewCreateProfileRequest

`func NewCreateProfileRequest(nickname string, ) *CreateProfileRequest`

NewCreateProfileRequest instantiates a new CreateProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProfileRequestWithDefaults

`func NewCreateProfileRequestWithDefaults() *CreateProfileRequest`

NewCreateProfileRequestWithDefaults instantiates a new CreateProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNickname

`func (o *CreateProfileRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *CreateProfileRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *CreateProfileRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.


### GetDescription

`func (o *CreateProfileRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProfileRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProfileRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateProfileRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


