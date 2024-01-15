# FiatAccountOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersonDetails** | Pointer to [**FiatAccountOwnerPersonDetails**](FiatAccountOwnerPersonDetails.md) |  | [optional] 
**InstitutionDetails** | Pointer to [**FiatAccountOwnerInstitutionDetails**](FiatAccountOwnerInstitutionDetails.md) |  | [optional] 

## Methods

### NewFiatAccountOwner

`func NewFiatAccountOwner() *FiatAccountOwner`

NewFiatAccountOwner instantiates a new FiatAccountOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiatAccountOwnerWithDefaults

`func NewFiatAccountOwnerWithDefaults() *FiatAccountOwner`

NewFiatAccountOwnerWithDefaults instantiates a new FiatAccountOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersonDetails

`func (o *FiatAccountOwner) GetPersonDetails() FiatAccountOwnerPersonDetails`

GetPersonDetails returns the PersonDetails field if non-nil, zero value otherwise.

### GetPersonDetailsOk

`func (o *FiatAccountOwner) GetPersonDetailsOk() (*FiatAccountOwnerPersonDetails, bool)`

GetPersonDetailsOk returns a tuple with the PersonDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonDetails

`func (o *FiatAccountOwner) SetPersonDetails(v FiatAccountOwnerPersonDetails)`

SetPersonDetails sets PersonDetails field to given value.

### HasPersonDetails

`func (o *FiatAccountOwner) HasPersonDetails() bool`

HasPersonDetails returns a boolean if a field has been set.

### GetInstitutionDetails

`func (o *FiatAccountOwner) GetInstitutionDetails() FiatAccountOwnerInstitutionDetails`

GetInstitutionDetails returns the InstitutionDetails field if non-nil, zero value otherwise.

### GetInstitutionDetailsOk

`func (o *FiatAccountOwner) GetInstitutionDetailsOk() (*FiatAccountOwnerInstitutionDetails, bool)`

GetInstitutionDetailsOk returns a tuple with the InstitutionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionDetails

`func (o *FiatAccountOwner) SetInstitutionDetails(v FiatAccountOwnerInstitutionDetails)`

SetInstitutionDetails sets InstitutionDetails field to given value.

### HasInstitutionDetails

`func (o *FiatAccountOwner) HasInstitutionDetails() bool`

HasInstitutionDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


