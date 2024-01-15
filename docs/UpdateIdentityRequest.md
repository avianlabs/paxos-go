# UpdateIdentityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersonDetails** | Pointer to [**PersonDetails**](PersonDetails.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**SetUserDisabled** | Pointer to **bool** | &#x60;true&#x60; disables the identity. &#x60;false&#x60; re-enables it, unless it has been disabled by a Paxos admin. | [optional] 
**InstitutionDetails** | Pointer to [**InstitutionDetails**](InstitutionDetails.md) |  | [optional] 
**RefId** | Pointer to **string** | A user-facing ID to prevent duplicate identity creation. Unique for all identities created by the same API user. | [optional] 
**TaxDetails** | Pointer to [**[]TaxDetail**](TaxDetail.md) |  | [optional] 
**SetTaxDetailsNotRequired** | Pointer to **bool** | Set to true if tax details are not legally required. | [optional] 

## Methods

### NewUpdateIdentityRequest

`func NewUpdateIdentityRequest() *UpdateIdentityRequest`

NewUpdateIdentityRequest instantiates a new UpdateIdentityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIdentityRequestWithDefaults

`func NewUpdateIdentityRequestWithDefaults() *UpdateIdentityRequest`

NewUpdateIdentityRequestWithDefaults instantiates a new UpdateIdentityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersonDetails

`func (o *UpdateIdentityRequest) GetPersonDetails() PersonDetails`

GetPersonDetails returns the PersonDetails field if non-nil, zero value otherwise.

### GetPersonDetailsOk

`func (o *UpdateIdentityRequest) GetPersonDetailsOk() (*PersonDetails, bool)`

GetPersonDetailsOk returns a tuple with the PersonDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonDetails

`func (o *UpdateIdentityRequest) SetPersonDetails(v PersonDetails)`

SetPersonDetails sets PersonDetails field to given value.

### HasPersonDetails

`func (o *UpdateIdentityRequest) HasPersonDetails() bool`

HasPersonDetails returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateIdentityRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateIdentityRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateIdentityRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateIdentityRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSetUserDisabled

`func (o *UpdateIdentityRequest) GetSetUserDisabled() bool`

GetSetUserDisabled returns the SetUserDisabled field if non-nil, zero value otherwise.

### GetSetUserDisabledOk

`func (o *UpdateIdentityRequest) GetSetUserDisabledOk() (*bool, bool)`

GetSetUserDisabledOk returns a tuple with the SetUserDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetUserDisabled

`func (o *UpdateIdentityRequest) SetSetUserDisabled(v bool)`

SetSetUserDisabled sets SetUserDisabled field to given value.

### HasSetUserDisabled

`func (o *UpdateIdentityRequest) HasSetUserDisabled() bool`

HasSetUserDisabled returns a boolean if a field has been set.

### GetInstitutionDetails

`func (o *UpdateIdentityRequest) GetInstitutionDetails() InstitutionDetails`

GetInstitutionDetails returns the InstitutionDetails field if non-nil, zero value otherwise.

### GetInstitutionDetailsOk

`func (o *UpdateIdentityRequest) GetInstitutionDetailsOk() (*InstitutionDetails, bool)`

GetInstitutionDetailsOk returns a tuple with the InstitutionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionDetails

`func (o *UpdateIdentityRequest) SetInstitutionDetails(v InstitutionDetails)`

SetInstitutionDetails sets InstitutionDetails field to given value.

### HasInstitutionDetails

`func (o *UpdateIdentityRequest) HasInstitutionDetails() bool`

HasInstitutionDetails returns a boolean if a field has been set.

### GetRefId

`func (o *UpdateIdentityRequest) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UpdateIdentityRequest) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UpdateIdentityRequest) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UpdateIdentityRequest) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetTaxDetails

`func (o *UpdateIdentityRequest) GetTaxDetails() []TaxDetail`

GetTaxDetails returns the TaxDetails field if non-nil, zero value otherwise.

### GetTaxDetailsOk

`func (o *UpdateIdentityRequest) GetTaxDetailsOk() (*[]TaxDetail, bool)`

GetTaxDetailsOk returns a tuple with the TaxDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDetails

`func (o *UpdateIdentityRequest) SetTaxDetails(v []TaxDetail)`

SetTaxDetails sets TaxDetails field to given value.

### HasTaxDetails

`func (o *UpdateIdentityRequest) HasTaxDetails() bool`

HasTaxDetails returns a boolean if a field has been set.

### GetSetTaxDetailsNotRequired

`func (o *UpdateIdentityRequest) GetSetTaxDetailsNotRequired() bool`

GetSetTaxDetailsNotRequired returns the SetTaxDetailsNotRequired field if non-nil, zero value otherwise.

### GetSetTaxDetailsNotRequiredOk

`func (o *UpdateIdentityRequest) GetSetTaxDetailsNotRequiredOk() (*bool, bool)`

GetSetTaxDetailsNotRequiredOk returns a tuple with the SetTaxDetailsNotRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetTaxDetailsNotRequired

`func (o *UpdateIdentityRequest) SetSetTaxDetailsNotRequired(v bool)`

SetSetTaxDetailsNotRequired sets SetTaxDetailsNotRequired field to given value.

### HasSetTaxDetailsNotRequired

`func (o *UpdateIdentityRequest) HasSetTaxDetailsNotRequired() bool`

HasSetTaxDetailsNotRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


