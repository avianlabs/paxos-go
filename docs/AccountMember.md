# AccountMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**AccountMemberAccountRoleType**](AccountMemberAccountRoleType.md) |  | [optional] 
**Roles** | Pointer to [**[]AccountMemberAccountRoleType**](AccountMemberAccountRoleType.md) |  | [optional] 
**Id** | Pointer to **string** | Account member ID. Note: This field is auto-generated. Specifying an ID when creating an account member is a client error. | [optional] 

## Methods

### NewAccountMember

`func NewAccountMember() *AccountMember`

NewAccountMember instantiates a new AccountMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountMemberWithDefaults

`func NewAccountMemberWithDefaults() *AccountMember`

NewAccountMemberWithDefaults instantiates a new AccountMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityId

`func (o *AccountMember) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *AccountMember) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *AccountMember) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *AccountMember) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetType

`func (o *AccountMember) GetType() AccountMemberAccountRoleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountMember) GetTypeOk() (*AccountMemberAccountRoleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountMember) SetType(v AccountMemberAccountRoleType)`

SetType sets Type field to given value.

### HasType

`func (o *AccountMember) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRoles

`func (o *AccountMember) GetRoles() []AccountMemberAccountRoleType`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AccountMember) GetRolesOk() (*[]AccountMemberAccountRoleType, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AccountMember) SetRoles(v []AccountMemberAccountRoleType)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AccountMember) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetId

`func (o *AccountMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountMember) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountMember) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


