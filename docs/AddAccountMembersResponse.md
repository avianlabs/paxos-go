# AddAccountMembersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | ID of account to which members were added. | [optional] 
**Members** | Pointer to [**[]AccountMember**](AccountMember.md) | List of account members that were added to the account. | [optional] 

## Methods

### NewAddAccountMembersResponse

`func NewAddAccountMembersResponse() *AddAccountMembersResponse`

NewAddAccountMembersResponse instantiates a new AddAccountMembersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAccountMembersResponseWithDefaults

`func NewAddAccountMembersResponseWithDefaults() *AddAccountMembersResponse`

NewAddAccountMembersResponseWithDefaults instantiates a new AddAccountMembersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AddAccountMembersResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddAccountMembersResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddAccountMembersResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddAccountMembersResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetMembers

`func (o *AddAccountMembersResponse) GetMembers() []AccountMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *AddAccountMembersResponse) GetMembersOk() (*[]AccountMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *AddAccountMembersResponse) SetMembers(v []AccountMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *AddAccountMembersResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


