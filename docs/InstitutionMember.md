# InstitutionMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityId** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to [**[]InstitutionRoleType**](InstitutionRoleType.md) |  | [optional] 
**Ownership** | Pointer to **string** | Decimal number representing the percent ownership the identity has in the company--  e.g. 5 represents 5% ownership. | [optional] 
**Position** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SummaryStatus** | Pointer to [**IdentityStatus**](IdentityStatus.md) |  | [optional] 
**Id** | Pointer to **string** | Institution member ID. Note: This field is auto-generated. Specifying an ID when creating an institution member is a client error. | [optional] 

## Methods

### NewInstitutionMember

`func NewInstitutionMember() *InstitutionMember`

NewInstitutionMember instantiates a new InstitutionMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionMemberWithDefaults

`func NewInstitutionMemberWithDefaults() *InstitutionMember`

NewInstitutionMemberWithDefaults instantiates a new InstitutionMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityId

`func (o *InstitutionMember) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *InstitutionMember) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *InstitutionMember) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *InstitutionMember) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetRoles

`func (o *InstitutionMember) GetRoles() []InstitutionRoleType`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *InstitutionMember) GetRolesOk() (*[]InstitutionRoleType, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *InstitutionMember) SetRoles(v []InstitutionRoleType)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *InstitutionMember) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetOwnership

`func (o *InstitutionMember) GetOwnership() string`

GetOwnership returns the Ownership field if non-nil, zero value otherwise.

### GetOwnershipOk

`func (o *InstitutionMember) GetOwnershipOk() (*string, bool)`

GetOwnershipOk returns a tuple with the Ownership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnership

`func (o *InstitutionMember) SetOwnership(v string)`

SetOwnership sets Ownership field to given value.

### HasOwnership

`func (o *InstitutionMember) HasOwnership() bool`

HasOwnership returns a boolean if a field has been set.

### GetPosition

`func (o *InstitutionMember) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *InstitutionMember) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *InstitutionMember) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *InstitutionMember) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetName

`func (o *InstitutionMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstitutionMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstitutionMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstitutionMember) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSummaryStatus

`func (o *InstitutionMember) GetSummaryStatus() IdentityStatus`

GetSummaryStatus returns the SummaryStatus field if non-nil, zero value otherwise.

### GetSummaryStatusOk

`func (o *InstitutionMember) GetSummaryStatusOk() (*IdentityStatus, bool)`

GetSummaryStatusOk returns a tuple with the SummaryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryStatus

`func (o *InstitutionMember) SetSummaryStatus(v IdentityStatus)`

SetSummaryStatus sets SummaryStatus field to given value.

### HasSummaryStatus

`func (o *InstitutionMember) HasSummaryStatus() bool`

HasSummaryStatus returns a boolean if a field has been set.

### GetId

`func (o *InstitutionMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstitutionMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstitutionMember) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InstitutionMember) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


