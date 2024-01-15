# AddInstitutionMembersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstitutionId** | Pointer to **string** | ID of institution identity to which members were added. | [optional] 
**Members** | Pointer to [**[]InstitutionMember**](InstitutionMember.md) | List of institution members that were added to the institution. | [optional] 

## Methods

### NewAddInstitutionMembersResponse

`func NewAddInstitutionMembersResponse() *AddInstitutionMembersResponse`

NewAddInstitutionMembersResponse instantiates a new AddInstitutionMembersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddInstitutionMembersResponseWithDefaults

`func NewAddInstitutionMembersResponseWithDefaults() *AddInstitutionMembersResponse`

NewAddInstitutionMembersResponseWithDefaults instantiates a new AddInstitutionMembersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstitutionId

`func (o *AddInstitutionMembersResponse) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *AddInstitutionMembersResponse) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *AddInstitutionMembersResponse) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.

### HasInstitutionId

`func (o *AddInstitutionMembersResponse) HasInstitutionId() bool`

HasInstitutionId returns a boolean if a field has been set.

### GetMembers

`func (o *AddInstitutionMembersResponse) GetMembers() []InstitutionMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *AddInstitutionMembersResponse) GetMembersOk() (*[]InstitutionMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *AddInstitutionMembersResponse) SetMembers(v []InstitutionMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *AddInstitutionMembersResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


