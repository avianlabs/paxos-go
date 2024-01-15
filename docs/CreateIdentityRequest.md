# CreateIdentityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersonDetails** | Pointer to [**PersonDetails**](PersonDetails.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**RefId** | Pointer to **string** | A user-facing ID to prevent duplicate identity creation. Unique for all identities created by the same API user. | [optional] 
**InstitutionDetails** | Pointer to [**InstitutionDetails**](InstitutionDetails.md) |  | [optional] 
**InstitutionMembers** | Pointer to [**[]InstitutionMember**](InstitutionMember.md) |  | [optional] 
**TaxDetails** | Pointer to [**[]TaxDetail**](TaxDetail.md) | List of tax details associated with the identity. Must be set if tax_details_not_required is false or not set. | [optional] 
**TaxDetailsNotRequired** | Pointer to **bool** | Set to true if tax details are not legally required. | [optional] 

## Methods

### NewCreateIdentityRequest

`func NewCreateIdentityRequest() *CreateIdentityRequest`

NewCreateIdentityRequest instantiates a new CreateIdentityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIdentityRequestWithDefaults

`func NewCreateIdentityRequestWithDefaults() *CreateIdentityRequest`

NewCreateIdentityRequestWithDefaults instantiates a new CreateIdentityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersonDetails

`func (o *CreateIdentityRequest) GetPersonDetails() PersonDetails`

GetPersonDetails returns the PersonDetails field if non-nil, zero value otherwise.

### GetPersonDetailsOk

`func (o *CreateIdentityRequest) GetPersonDetailsOk() (*PersonDetails, bool)`

GetPersonDetailsOk returns a tuple with the PersonDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonDetails

`func (o *CreateIdentityRequest) SetPersonDetails(v PersonDetails)`

SetPersonDetails sets PersonDetails field to given value.

### HasPersonDetails

`func (o *CreateIdentityRequest) HasPersonDetails() bool`

HasPersonDetails returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateIdentityRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateIdentityRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateIdentityRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateIdentityRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRefId

`func (o *CreateIdentityRequest) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *CreateIdentityRequest) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *CreateIdentityRequest) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *CreateIdentityRequest) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetInstitutionDetails

`func (o *CreateIdentityRequest) GetInstitutionDetails() InstitutionDetails`

GetInstitutionDetails returns the InstitutionDetails field if non-nil, zero value otherwise.

### GetInstitutionDetailsOk

`func (o *CreateIdentityRequest) GetInstitutionDetailsOk() (*InstitutionDetails, bool)`

GetInstitutionDetailsOk returns a tuple with the InstitutionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionDetails

`func (o *CreateIdentityRequest) SetInstitutionDetails(v InstitutionDetails)`

SetInstitutionDetails sets InstitutionDetails field to given value.

### HasInstitutionDetails

`func (o *CreateIdentityRequest) HasInstitutionDetails() bool`

HasInstitutionDetails returns a boolean if a field has been set.

### GetInstitutionMembers

`func (o *CreateIdentityRequest) GetInstitutionMembers() []InstitutionMember`

GetInstitutionMembers returns the InstitutionMembers field if non-nil, zero value otherwise.

### GetInstitutionMembersOk

`func (o *CreateIdentityRequest) GetInstitutionMembersOk() (*[]InstitutionMember, bool)`

GetInstitutionMembersOk returns a tuple with the InstitutionMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionMembers

`func (o *CreateIdentityRequest) SetInstitutionMembers(v []InstitutionMember)`

SetInstitutionMembers sets InstitutionMembers field to given value.

### HasInstitutionMembers

`func (o *CreateIdentityRequest) HasInstitutionMembers() bool`

HasInstitutionMembers returns a boolean if a field has been set.

### GetTaxDetails

`func (o *CreateIdentityRequest) GetTaxDetails() []TaxDetail`

GetTaxDetails returns the TaxDetails field if non-nil, zero value otherwise.

### GetTaxDetailsOk

`func (o *CreateIdentityRequest) GetTaxDetailsOk() (*[]TaxDetail, bool)`

GetTaxDetailsOk returns a tuple with the TaxDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDetails

`func (o *CreateIdentityRequest) SetTaxDetails(v []TaxDetail)`

SetTaxDetails sets TaxDetails field to given value.

### HasTaxDetails

`func (o *CreateIdentityRequest) HasTaxDetails() bool`

HasTaxDetails returns a boolean if a field has been set.

### GetTaxDetailsNotRequired

`func (o *CreateIdentityRequest) GetTaxDetailsNotRequired() bool`

GetTaxDetailsNotRequired returns the TaxDetailsNotRequired field if non-nil, zero value otherwise.

### GetTaxDetailsNotRequiredOk

`func (o *CreateIdentityRequest) GetTaxDetailsNotRequiredOk() (*bool, bool)`

GetTaxDetailsNotRequiredOk returns a tuple with the TaxDetailsNotRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDetailsNotRequired

`func (o *CreateIdentityRequest) SetTaxDetailsNotRequired(v bool)`

SetTaxDetailsNotRequired sets TaxDetailsNotRequired field to given value.

### HasTaxDetailsNotRequired

`func (o *CreateIdentityRequest) HasTaxDetailsNotRequired() bool`

HasTaxDetailsNotRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


