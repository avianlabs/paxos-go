# CreateAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**Account**](Account.md) |  | 
**CreateProfile** | Pointer to **bool** | Create a new profile for this account. Set to &#x60;true&#x60; to track user balances at the Profile level for this account. See also [Profiles](#tag/Profiles). | [optional] 

## Methods

### NewCreateAccountRequest

`func NewCreateAccountRequest(account Account, ) *CreateAccountRequest`

NewCreateAccountRequest instantiates a new CreateAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountRequestWithDefaults

`func NewCreateAccountRequestWithDefaults() *CreateAccountRequest`

NewCreateAccountRequestWithDefaults instantiates a new CreateAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CreateAccountRequest) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateAccountRequest) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateAccountRequest) SetAccount(v Account)`

SetAccount sets Account field to given value.


### GetCreateProfile

`func (o *CreateAccountRequest) GetCreateProfile() bool`

GetCreateProfile returns the CreateProfile field if non-nil, zero value otherwise.

### GetCreateProfileOk

`func (o *CreateAccountRequest) GetCreateProfileOk() (*bool, bool)`

GetCreateProfileOk returns a tuple with the CreateProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateProfile

`func (o *CreateAccountRequest) SetCreateProfile(v bool)`

SetCreateProfile sets CreateProfile field to given value.

### HasCreateProfile

`func (o *CreateAccountRequest) HasCreateProfile() bool`

HasCreateProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


