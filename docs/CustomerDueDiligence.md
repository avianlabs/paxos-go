# CustomerDueDiligence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | A list of alternate names or aliases associated with the Identity. | [optional] 
**EstimatedNetWorth** | Pointer to [**CustomerDueDiligenceNetWorthRange**](CustomerDueDiligenceNetWorthRange.md) |  | [optional] 
**EstimatedYearlyIncome** | Pointer to [**CustomerDueDiligenceYearlyIncomeRange**](CustomerDueDiligenceYearlyIncomeRange.md) |  | [optional] 
**ExpectedTransferValue** | Pointer to [**CustomerDueDiligenceTransferValueRange**](CustomerDueDiligenceTransferValueRange.md) |  | [optional] 
**SourceOfWealth** | Pointer to [**WealthSource**](WealthSource.md) |  | [optional] 
**SourceOfFunds** | Pointer to [**FundsSource**](FundsSource.md) |  | [optional] 
**PurposeOfAccount** | Pointer to [**AccountPurpose**](AccountPurpose.md) |  | [optional] 
**EmploymentStatus** | Pointer to [**EmploymentStatus**](EmploymentStatus.md) |  | [optional] 
**EmploymentIndustrySector** | Pointer to [**InstitutionSubType**](InstitutionSubType.md) |  | [optional] 
**IndustrySector** | Pointer to [**InstitutionSubType**](InstitutionSubType.md) |  | [optional] 
**HasUnderlyingTrustStructure** | Pointer to **bool** | Whether or not the institution tied to the Identity has an underlying trust structure. | [optional] 
**HasNomineeShareholders** | Pointer to **bool** | Whether or not the institution tied to the Identity has nominee shareholders. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCustomerDueDiligence

`func NewCustomerDueDiligence() *CustomerDueDiligence`

NewCustomerDueDiligence instantiates a new CustomerDueDiligence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerDueDiligenceWithDefaults

`func NewCustomerDueDiligenceWithDefaults() *CustomerDueDiligence`

NewCustomerDueDiligenceWithDefaults instantiates a new CustomerDueDiligence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *CustomerDueDiligence) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *CustomerDueDiligence) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *CustomerDueDiligence) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *CustomerDueDiligence) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetEstimatedNetWorth

`func (o *CustomerDueDiligence) GetEstimatedNetWorth() CustomerDueDiligenceNetWorthRange`

GetEstimatedNetWorth returns the EstimatedNetWorth field if non-nil, zero value otherwise.

### GetEstimatedNetWorthOk

`func (o *CustomerDueDiligence) GetEstimatedNetWorthOk() (*CustomerDueDiligenceNetWorthRange, bool)`

GetEstimatedNetWorthOk returns a tuple with the EstimatedNetWorth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedNetWorth

`func (o *CustomerDueDiligence) SetEstimatedNetWorth(v CustomerDueDiligenceNetWorthRange)`

SetEstimatedNetWorth sets EstimatedNetWorth field to given value.

### HasEstimatedNetWorth

`func (o *CustomerDueDiligence) HasEstimatedNetWorth() bool`

HasEstimatedNetWorth returns a boolean if a field has been set.

### GetEstimatedYearlyIncome

`func (o *CustomerDueDiligence) GetEstimatedYearlyIncome() CustomerDueDiligenceYearlyIncomeRange`

GetEstimatedYearlyIncome returns the EstimatedYearlyIncome field if non-nil, zero value otherwise.

### GetEstimatedYearlyIncomeOk

`func (o *CustomerDueDiligence) GetEstimatedYearlyIncomeOk() (*CustomerDueDiligenceYearlyIncomeRange, bool)`

GetEstimatedYearlyIncomeOk returns a tuple with the EstimatedYearlyIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedYearlyIncome

`func (o *CustomerDueDiligence) SetEstimatedYearlyIncome(v CustomerDueDiligenceYearlyIncomeRange)`

SetEstimatedYearlyIncome sets EstimatedYearlyIncome field to given value.

### HasEstimatedYearlyIncome

`func (o *CustomerDueDiligence) HasEstimatedYearlyIncome() bool`

HasEstimatedYearlyIncome returns a boolean if a field has been set.

### GetExpectedTransferValue

`func (o *CustomerDueDiligence) GetExpectedTransferValue() CustomerDueDiligenceTransferValueRange`

GetExpectedTransferValue returns the ExpectedTransferValue field if non-nil, zero value otherwise.

### GetExpectedTransferValueOk

`func (o *CustomerDueDiligence) GetExpectedTransferValueOk() (*CustomerDueDiligenceTransferValueRange, bool)`

GetExpectedTransferValueOk returns a tuple with the ExpectedTransferValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedTransferValue

`func (o *CustomerDueDiligence) SetExpectedTransferValue(v CustomerDueDiligenceTransferValueRange)`

SetExpectedTransferValue sets ExpectedTransferValue field to given value.

### HasExpectedTransferValue

`func (o *CustomerDueDiligence) HasExpectedTransferValue() bool`

HasExpectedTransferValue returns a boolean if a field has been set.

### GetSourceOfWealth

`func (o *CustomerDueDiligence) GetSourceOfWealth() WealthSource`

GetSourceOfWealth returns the SourceOfWealth field if non-nil, zero value otherwise.

### GetSourceOfWealthOk

`func (o *CustomerDueDiligence) GetSourceOfWealthOk() (*WealthSource, bool)`

GetSourceOfWealthOk returns a tuple with the SourceOfWealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfWealth

`func (o *CustomerDueDiligence) SetSourceOfWealth(v WealthSource)`

SetSourceOfWealth sets SourceOfWealth field to given value.

### HasSourceOfWealth

`func (o *CustomerDueDiligence) HasSourceOfWealth() bool`

HasSourceOfWealth returns a boolean if a field has been set.

### GetSourceOfFunds

`func (o *CustomerDueDiligence) GetSourceOfFunds() FundsSource`

GetSourceOfFunds returns the SourceOfFunds field if non-nil, zero value otherwise.

### GetSourceOfFundsOk

`func (o *CustomerDueDiligence) GetSourceOfFundsOk() (*FundsSource, bool)`

GetSourceOfFundsOk returns a tuple with the SourceOfFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFunds

`func (o *CustomerDueDiligence) SetSourceOfFunds(v FundsSource)`

SetSourceOfFunds sets SourceOfFunds field to given value.

### HasSourceOfFunds

`func (o *CustomerDueDiligence) HasSourceOfFunds() bool`

HasSourceOfFunds returns a boolean if a field has been set.

### GetPurposeOfAccount

`func (o *CustomerDueDiligence) GetPurposeOfAccount() AccountPurpose`

GetPurposeOfAccount returns the PurposeOfAccount field if non-nil, zero value otherwise.

### GetPurposeOfAccountOk

`func (o *CustomerDueDiligence) GetPurposeOfAccountOk() (*AccountPurpose, bool)`

GetPurposeOfAccountOk returns a tuple with the PurposeOfAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeOfAccount

`func (o *CustomerDueDiligence) SetPurposeOfAccount(v AccountPurpose)`

SetPurposeOfAccount sets PurposeOfAccount field to given value.

### HasPurposeOfAccount

`func (o *CustomerDueDiligence) HasPurposeOfAccount() bool`

HasPurposeOfAccount returns a boolean if a field has been set.

### GetEmploymentStatus

`func (o *CustomerDueDiligence) GetEmploymentStatus() EmploymentStatus`

GetEmploymentStatus returns the EmploymentStatus field if non-nil, zero value otherwise.

### GetEmploymentStatusOk

`func (o *CustomerDueDiligence) GetEmploymentStatusOk() (*EmploymentStatus, bool)`

GetEmploymentStatusOk returns a tuple with the EmploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentStatus

`func (o *CustomerDueDiligence) SetEmploymentStatus(v EmploymentStatus)`

SetEmploymentStatus sets EmploymentStatus field to given value.

### HasEmploymentStatus

`func (o *CustomerDueDiligence) HasEmploymentStatus() bool`

HasEmploymentStatus returns a boolean if a field has been set.

### GetEmploymentIndustrySector

`func (o *CustomerDueDiligence) GetEmploymentIndustrySector() InstitutionSubType`

GetEmploymentIndustrySector returns the EmploymentIndustrySector field if non-nil, zero value otherwise.

### GetEmploymentIndustrySectorOk

`func (o *CustomerDueDiligence) GetEmploymentIndustrySectorOk() (*InstitutionSubType, bool)`

GetEmploymentIndustrySectorOk returns a tuple with the EmploymentIndustrySector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentIndustrySector

`func (o *CustomerDueDiligence) SetEmploymentIndustrySector(v InstitutionSubType)`

SetEmploymentIndustrySector sets EmploymentIndustrySector field to given value.

### HasEmploymentIndustrySector

`func (o *CustomerDueDiligence) HasEmploymentIndustrySector() bool`

HasEmploymentIndustrySector returns a boolean if a field has been set.

### GetIndustrySector

`func (o *CustomerDueDiligence) GetIndustrySector() InstitutionSubType`

GetIndustrySector returns the IndustrySector field if non-nil, zero value otherwise.

### GetIndustrySectorOk

`func (o *CustomerDueDiligence) GetIndustrySectorOk() (*InstitutionSubType, bool)`

GetIndustrySectorOk returns a tuple with the IndustrySector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustrySector

`func (o *CustomerDueDiligence) SetIndustrySector(v InstitutionSubType)`

SetIndustrySector sets IndustrySector field to given value.

### HasIndustrySector

`func (o *CustomerDueDiligence) HasIndustrySector() bool`

HasIndustrySector returns a boolean if a field has been set.

### GetHasUnderlyingTrustStructure

`func (o *CustomerDueDiligence) GetHasUnderlyingTrustStructure() bool`

GetHasUnderlyingTrustStructure returns the HasUnderlyingTrustStructure field if non-nil, zero value otherwise.

### GetHasUnderlyingTrustStructureOk

`func (o *CustomerDueDiligence) GetHasUnderlyingTrustStructureOk() (*bool, bool)`

GetHasUnderlyingTrustStructureOk returns a tuple with the HasUnderlyingTrustStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUnderlyingTrustStructure

`func (o *CustomerDueDiligence) SetHasUnderlyingTrustStructure(v bool)`

SetHasUnderlyingTrustStructure sets HasUnderlyingTrustStructure field to given value.

### HasHasUnderlyingTrustStructure

`func (o *CustomerDueDiligence) HasHasUnderlyingTrustStructure() bool`

HasHasUnderlyingTrustStructure returns a boolean if a field has been set.

### GetHasNomineeShareholders

`func (o *CustomerDueDiligence) GetHasNomineeShareholders() bool`

GetHasNomineeShareholders returns the HasNomineeShareholders field if non-nil, zero value otherwise.

### GetHasNomineeShareholdersOk

`func (o *CustomerDueDiligence) GetHasNomineeShareholdersOk() (*bool, bool)`

GetHasNomineeShareholdersOk returns a tuple with the HasNomineeShareholders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNomineeShareholders

`func (o *CustomerDueDiligence) SetHasNomineeShareholders(v bool)`

SetHasNomineeShareholders sets HasNomineeShareholders field to given value.

### HasHasNomineeShareholders

`func (o *CustomerDueDiligence) HasHasNomineeShareholders() bool`

HasHasNomineeShareholders returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomerDueDiligence) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerDueDiligence) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerDueDiligence) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomerDueDiligence) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


