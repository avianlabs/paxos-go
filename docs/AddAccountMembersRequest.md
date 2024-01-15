# AddAccountMembersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | ID of account to which members will be added. | 
**Members** | [**[]AccountMember**](AccountMember.md) | A non-empty array of account members to be added. | 

## Methods

### NewAddAccountMembersRequest

`func NewAddAccountMembersRequest(accountId string, members []AccountMember, ) *AddAccountMembersRequest`

NewAddAccountMembersRequest instantiates a new AddAccountMembersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAccountMembersRequestWithDefaults

`func NewAddAccountMembersRequestWithDefaults() *AddAccountMembersRequest`

NewAddAccountMembersRequestWithDefaults instantiates a new AddAccountMembersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AddAccountMembersRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddAccountMembersRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddAccountMembersRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetMembers

`func (o *AddAccountMembersRequest) GetMembers() []AccountMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *AddAccountMembersRequest) GetMembersOk() (*[]AccountMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *AddAccountMembersRequest) SetMembers(v []AccountMember)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


