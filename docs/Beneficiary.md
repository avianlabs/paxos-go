# Beneficiary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersonDetails** | Pointer to [**BeneficiaryPerson**](BeneficiaryPerson.md) |  | [optional] 
**InstitutionDetails** | Pointer to [**BeneficiaryInstitution**](BeneficiaryInstitution.md) |  | [optional] 

## Methods

### NewBeneficiary

`func NewBeneficiary() *Beneficiary`

NewBeneficiary instantiates a new Beneficiary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeneficiaryWithDefaults

`func NewBeneficiaryWithDefaults() *Beneficiary`

NewBeneficiaryWithDefaults instantiates a new Beneficiary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersonDetails

`func (o *Beneficiary) GetPersonDetails() BeneficiaryPerson`

GetPersonDetails returns the PersonDetails field if non-nil, zero value otherwise.

### GetPersonDetailsOk

`func (o *Beneficiary) GetPersonDetailsOk() (*BeneficiaryPerson, bool)`

GetPersonDetailsOk returns a tuple with the PersonDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonDetails

`func (o *Beneficiary) SetPersonDetails(v BeneficiaryPerson)`

SetPersonDetails sets PersonDetails field to given value.

### HasPersonDetails

`func (o *Beneficiary) HasPersonDetails() bool`

HasPersonDetails returns a boolean if a field has been set.

### GetInstitutionDetails

`func (o *Beneficiary) GetInstitutionDetails() BeneficiaryInstitution`

GetInstitutionDetails returns the InstitutionDetails field if non-nil, zero value otherwise.

### GetInstitutionDetailsOk

`func (o *Beneficiary) GetInstitutionDetailsOk() (*BeneficiaryInstitution, bool)`

GetInstitutionDetailsOk returns a tuple with the InstitutionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionDetails

`func (o *Beneficiary) SetInstitutionDetails(v BeneficiaryInstitution)`

SetInstitutionDetails sets InstitutionDetails field to given value.

### HasInstitutionDetails

`func (o *Beneficiary) HasInstitutionDetails() bool`

HasInstitutionDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


