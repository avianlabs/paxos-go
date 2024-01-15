# Profile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Nickname** | **string** |  | 
**Type** | [**ProfileType**](ProfileType.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 

## Methods

### NewProfile

`func NewProfile(id string, nickname string, type_ ProfileType, ) *Profile`

NewProfile instantiates a new Profile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileWithDefaults

`func NewProfileWithDefaults() *Profile`

NewProfileWithDefaults instantiates a new Profile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Profile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Profile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Profile) SetId(v string)`

SetId sets Id field to given value.


### GetNickname

`func (o *Profile) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *Profile) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *Profile) SetNickname(v string)`

SetNickname sets Nickname field to given value.


### GetType

`func (o *Profile) GetType() ProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Profile) GetTypeOk() (*ProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Profile) SetType(v ProfileType)`

SetType sets Type field to given value.


### GetDescription

`func (o *Profile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Profile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Profile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Profile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccountId

`func (o *Profile) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Profile) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Profile) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Profile) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


