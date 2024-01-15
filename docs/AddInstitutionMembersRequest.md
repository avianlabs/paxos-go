# AddInstitutionMembersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstitutionId** | **string** | ID of institution identity to which members will be added. | 
**Members** | [**[]InstitutionMember**](InstitutionMember.md) | A non-empty array of institution members to be added. | 

## Methods

### NewAddInstitutionMembersRequest

`func NewAddInstitutionMembersRequest(institutionId string, members []InstitutionMember, ) *AddInstitutionMembersRequest`

NewAddInstitutionMembersRequest instantiates a new AddInstitutionMembersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddInstitutionMembersRequestWithDefaults

`func NewAddInstitutionMembersRequestWithDefaults() *AddInstitutionMembersRequest`

NewAddInstitutionMembersRequestWithDefaults instantiates a new AddInstitutionMembersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstitutionId

`func (o *AddInstitutionMembersRequest) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *AddInstitutionMembersRequest) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *AddInstitutionMembersRequest) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.


### GetMembers

`func (o *AddInstitutionMembersRequest) GetMembers() []InstitutionMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *AddInstitutionMembersRequest) GetMembersOk() (*[]InstitutionMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *AddInstitutionMembersRequest) SetMembers(v []InstitutionMember)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


