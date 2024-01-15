# UpdateAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**Account**](Account.md) |  | 
**SetUserDisabled** | Pointer to **bool** | true if the account is required to be disabled by the API user. | [optional] 

## Methods

### NewUpdateAccountRequest

`func NewUpdateAccountRequest(account Account, ) *UpdateAccountRequest`

NewUpdateAccountRequest instantiates a new UpdateAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountRequestWithDefaults

`func NewUpdateAccountRequestWithDefaults() *UpdateAccountRequest`

NewUpdateAccountRequestWithDefaults instantiates a new UpdateAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *UpdateAccountRequest) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateAccountRequest) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateAccountRequest) SetAccount(v Account)`

SetAccount sets Account field to given value.


### GetSetUserDisabled

`func (o *UpdateAccountRequest) GetSetUserDisabled() bool`

GetSetUserDisabled returns the SetUserDisabled field if non-nil, zero value otherwise.

### GetSetUserDisabledOk

`func (o *UpdateAccountRequest) GetSetUserDisabledOk() (*bool, bool)`

GetSetUserDisabledOk returns a tuple with the SetUserDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetUserDisabled

`func (o *UpdateAccountRequest) SetSetUserDisabled(v bool)`

SetSetUserDisabled sets SetUserDisabled field to given value.

### HasSetUserDisabled

`func (o *UpdateAccountRequest) HasSetUserDisabled() bool`

HasSetUserDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


