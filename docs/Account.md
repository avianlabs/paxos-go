# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id used for all other interactions with this account. | [optional] 
**IdentityId** | Pointer to **string** | The primary identity associated with the account. | [optional] 
**Description** | Pointer to **string** | A free-text description of the account. | [optional] 
**AdminDisabled** | Pointer to **bool** | true if the account has been disabled by a Paxos admin. | [optional] 
**UserDisabled** | Pointer to **bool** | true if the account has been disabled by the API user. | [optional] 
**Metadata** | Pointer to **map[string]string** | API User-facing metadata. | [optional] 
**RefId** | Pointer to **string** | A user-facing ID to prevent duplicate account creation. Unique for all accounts created by the same API user. | [optional] 
**Members** | Pointer to [**[]AccountMember**](AccountMember.md) | Additional Identities with varying types of access to this account. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time this account was created. | [optional] 
**SummaryStatus** | Pointer to [**IdentityStatus**](IdentityStatus.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time this account was updated. | [optional] 
**ProfileId** | Pointer to **string** | The ID of the profile created for the account when &#x60;create_profile&#x3D;true&#x60;. The [Profile](#tag/Profiles) type is &#x60;NORMAL&#x60;. The field is omitted when the account has no associated [Profile](#tag/Profiles). | [optional] 
**Type** | Pointer to [**AccountAccountType**](AccountAccountType.md) |  | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Account) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityId

`func (o *Account) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *Account) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *Account) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *Account) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetDescription

`func (o *Account) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Account) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Account) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Account) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdminDisabled

`func (o *Account) GetAdminDisabled() bool`

GetAdminDisabled returns the AdminDisabled field if non-nil, zero value otherwise.

### GetAdminDisabledOk

`func (o *Account) GetAdminDisabledOk() (*bool, bool)`

GetAdminDisabledOk returns a tuple with the AdminDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminDisabled

`func (o *Account) SetAdminDisabled(v bool)`

SetAdminDisabled sets AdminDisabled field to given value.

### HasAdminDisabled

`func (o *Account) HasAdminDisabled() bool`

HasAdminDisabled returns a boolean if a field has been set.

### GetUserDisabled

`func (o *Account) GetUserDisabled() bool`

GetUserDisabled returns the UserDisabled field if non-nil, zero value otherwise.

### GetUserDisabledOk

`func (o *Account) GetUserDisabledOk() (*bool, bool)`

GetUserDisabledOk returns a tuple with the UserDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisabled

`func (o *Account) SetUserDisabled(v bool)`

SetUserDisabled sets UserDisabled field to given value.

### HasUserDisabled

`func (o *Account) HasUserDisabled() bool`

HasUserDisabled returns a boolean if a field has been set.

### GetMetadata

`func (o *Account) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Account) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Account) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Account) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRefId

`func (o *Account) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *Account) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *Account) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *Account) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetMembers

`func (o *Account) GetMembers() []AccountMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Account) GetMembersOk() (*[]AccountMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Account) SetMembers(v []AccountMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Account) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Account) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Account) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Account) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Account) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSummaryStatus

`func (o *Account) GetSummaryStatus() IdentityStatus`

GetSummaryStatus returns the SummaryStatus field if non-nil, zero value otherwise.

### GetSummaryStatusOk

`func (o *Account) GetSummaryStatusOk() (*IdentityStatus, bool)`

GetSummaryStatusOk returns a tuple with the SummaryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryStatus

`func (o *Account) SetSummaryStatus(v IdentityStatus)`

SetSummaryStatus sets SummaryStatus field to given value.

### HasSummaryStatus

`func (o *Account) HasSummaryStatus() bool`

HasSummaryStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Account) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Account) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Account) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Account) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetProfileId

`func (o *Account) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *Account) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *Account) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *Account) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetType

`func (o *Account) GetType() AccountAccountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Account) GetTypeOk() (*AccountAccountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Account) SetType(v AccountAccountType)`

SetType sets Type field to given value.

### HasType

`func (o *Account) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


