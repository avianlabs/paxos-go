# Identity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**SummaryStatus** | Pointer to [**IdentityStatus**](IdentityStatus.md) |  | [optional] 
**UserDisabled** | Pointer to **bool** |  | [optional] 
**AdminDisabled** | Pointer to **bool** |  | [optional] 
**PersonDetails** | Pointer to [**PersonDetails**](PersonDetails.md) |  | [optional] 
**Type** | Pointer to [**IdentityType**](IdentityType.md) |  | [optional] 
**RefId** | Pointer to **string** | A user-facing ID to prevent duplicate account creation. Unique for all accounts created by the same API user. | [optional] 
**InstitutionDetails** | Pointer to [**InstitutionDetails**](InstitutionDetails.md) |  | [optional] 
**InstitutionMembers** | Pointer to [**[]InstitutionMember**](InstitutionMember.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | the time at which the identity is created at. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | the time at which the identity is updated at. | [optional] 
**TaxDetails** | Pointer to [**[]TaxDetail**](TaxDetail.md) |  | [optional] 
**TaxDetailsNotRequired** | Pointer to **bool** |  | [optional] 
**SummaryTinVerificationStatus** | Pointer to [**TINVerificationStatus**](TINVerificationStatus.md) |  | [optional] 
**CustomerDueDiligence** | Pointer to [**CustomerDueDiligence**](CustomerDueDiligence.md) |  | [optional] 
**IsMerchant** | Pointer to **bool** | True if the identity is a merchant. | [optional] 

## Methods

### NewIdentity

`func NewIdentity(id string, ) *Identity`

NewIdentity instantiates a new Identity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityWithDefaults

`func NewIdentityWithDefaults() *Identity`

NewIdentityWithDefaults instantiates a new Identity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Identity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Identity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Identity) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *Identity) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Identity) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Identity) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Identity) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSummaryStatus

`func (o *Identity) GetSummaryStatus() IdentityStatus`

GetSummaryStatus returns the SummaryStatus field if non-nil, zero value otherwise.

### GetSummaryStatusOk

`func (o *Identity) GetSummaryStatusOk() (*IdentityStatus, bool)`

GetSummaryStatusOk returns a tuple with the SummaryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryStatus

`func (o *Identity) SetSummaryStatus(v IdentityStatus)`

SetSummaryStatus sets SummaryStatus field to given value.

### HasSummaryStatus

`func (o *Identity) HasSummaryStatus() bool`

HasSummaryStatus returns a boolean if a field has been set.

### GetUserDisabled

`func (o *Identity) GetUserDisabled() bool`

GetUserDisabled returns the UserDisabled field if non-nil, zero value otherwise.

### GetUserDisabledOk

`func (o *Identity) GetUserDisabledOk() (*bool, bool)`

GetUserDisabledOk returns a tuple with the UserDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisabled

`func (o *Identity) SetUserDisabled(v bool)`

SetUserDisabled sets UserDisabled field to given value.

### HasUserDisabled

`func (o *Identity) HasUserDisabled() bool`

HasUserDisabled returns a boolean if a field has been set.

### GetAdminDisabled

`func (o *Identity) GetAdminDisabled() bool`

GetAdminDisabled returns the AdminDisabled field if non-nil, zero value otherwise.

### GetAdminDisabledOk

`func (o *Identity) GetAdminDisabledOk() (*bool, bool)`

GetAdminDisabledOk returns a tuple with the AdminDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminDisabled

`func (o *Identity) SetAdminDisabled(v bool)`

SetAdminDisabled sets AdminDisabled field to given value.

### HasAdminDisabled

`func (o *Identity) HasAdminDisabled() bool`

HasAdminDisabled returns a boolean if a field has been set.

### GetPersonDetails

`func (o *Identity) GetPersonDetails() PersonDetails`

GetPersonDetails returns the PersonDetails field if non-nil, zero value otherwise.

### GetPersonDetailsOk

`func (o *Identity) GetPersonDetailsOk() (*PersonDetails, bool)`

GetPersonDetailsOk returns a tuple with the PersonDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonDetails

`func (o *Identity) SetPersonDetails(v PersonDetails)`

SetPersonDetails sets PersonDetails field to given value.

### HasPersonDetails

`func (o *Identity) HasPersonDetails() bool`

HasPersonDetails returns a boolean if a field has been set.

### GetType

`func (o *Identity) GetType() IdentityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Identity) GetTypeOk() (*IdentityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Identity) SetType(v IdentityType)`

SetType sets Type field to given value.

### HasType

`func (o *Identity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRefId

`func (o *Identity) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *Identity) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *Identity) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *Identity) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetInstitutionDetails

`func (o *Identity) GetInstitutionDetails() InstitutionDetails`

GetInstitutionDetails returns the InstitutionDetails field if non-nil, zero value otherwise.

### GetInstitutionDetailsOk

`func (o *Identity) GetInstitutionDetailsOk() (*InstitutionDetails, bool)`

GetInstitutionDetailsOk returns a tuple with the InstitutionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionDetails

`func (o *Identity) SetInstitutionDetails(v InstitutionDetails)`

SetInstitutionDetails sets InstitutionDetails field to given value.

### HasInstitutionDetails

`func (o *Identity) HasInstitutionDetails() bool`

HasInstitutionDetails returns a boolean if a field has been set.

### GetInstitutionMembers

`func (o *Identity) GetInstitutionMembers() []InstitutionMember`

GetInstitutionMembers returns the InstitutionMembers field if non-nil, zero value otherwise.

### GetInstitutionMembersOk

`func (o *Identity) GetInstitutionMembersOk() (*[]InstitutionMember, bool)`

GetInstitutionMembersOk returns a tuple with the InstitutionMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionMembers

`func (o *Identity) SetInstitutionMembers(v []InstitutionMember)`

SetInstitutionMembers sets InstitutionMembers field to given value.

### HasInstitutionMembers

`func (o *Identity) HasInstitutionMembers() bool`

HasInstitutionMembers returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Identity) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Identity) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Identity) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Identity) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Identity) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Identity) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Identity) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Identity) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTaxDetails

`func (o *Identity) GetTaxDetails() []TaxDetail`

GetTaxDetails returns the TaxDetails field if non-nil, zero value otherwise.

### GetTaxDetailsOk

`func (o *Identity) GetTaxDetailsOk() (*[]TaxDetail, bool)`

GetTaxDetailsOk returns a tuple with the TaxDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDetails

`func (o *Identity) SetTaxDetails(v []TaxDetail)`

SetTaxDetails sets TaxDetails field to given value.

### HasTaxDetails

`func (o *Identity) HasTaxDetails() bool`

HasTaxDetails returns a boolean if a field has been set.

### GetTaxDetailsNotRequired

`func (o *Identity) GetTaxDetailsNotRequired() bool`

GetTaxDetailsNotRequired returns the TaxDetailsNotRequired field if non-nil, zero value otherwise.

### GetTaxDetailsNotRequiredOk

`func (o *Identity) GetTaxDetailsNotRequiredOk() (*bool, bool)`

GetTaxDetailsNotRequiredOk returns a tuple with the TaxDetailsNotRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDetailsNotRequired

`func (o *Identity) SetTaxDetailsNotRequired(v bool)`

SetTaxDetailsNotRequired sets TaxDetailsNotRequired field to given value.

### HasTaxDetailsNotRequired

`func (o *Identity) HasTaxDetailsNotRequired() bool`

HasTaxDetailsNotRequired returns a boolean if a field has been set.

### GetSummaryTinVerificationStatus

`func (o *Identity) GetSummaryTinVerificationStatus() TINVerificationStatus`

GetSummaryTinVerificationStatus returns the SummaryTinVerificationStatus field if non-nil, zero value otherwise.

### GetSummaryTinVerificationStatusOk

`func (o *Identity) GetSummaryTinVerificationStatusOk() (*TINVerificationStatus, bool)`

GetSummaryTinVerificationStatusOk returns a tuple with the SummaryTinVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryTinVerificationStatus

`func (o *Identity) SetSummaryTinVerificationStatus(v TINVerificationStatus)`

SetSummaryTinVerificationStatus sets SummaryTinVerificationStatus field to given value.

### HasSummaryTinVerificationStatus

`func (o *Identity) HasSummaryTinVerificationStatus() bool`

HasSummaryTinVerificationStatus returns a boolean if a field has been set.

### GetCustomerDueDiligence

`func (o *Identity) GetCustomerDueDiligence() CustomerDueDiligence`

GetCustomerDueDiligence returns the CustomerDueDiligence field if non-nil, zero value otherwise.

### GetCustomerDueDiligenceOk

`func (o *Identity) GetCustomerDueDiligenceOk() (*CustomerDueDiligence, bool)`

GetCustomerDueDiligenceOk returns a tuple with the CustomerDueDiligence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDueDiligence

`func (o *Identity) SetCustomerDueDiligence(v CustomerDueDiligence)`

SetCustomerDueDiligence sets CustomerDueDiligence field to given value.

### HasCustomerDueDiligence

`func (o *Identity) HasCustomerDueDiligence() bool`

HasCustomerDueDiligence returns a boolean if a field has been set.

### GetIsMerchant

`func (o *Identity) GetIsMerchant() bool`

GetIsMerchant returns the IsMerchant field if non-nil, zero value otherwise.

### GetIsMerchantOk

`func (o *Identity) GetIsMerchantOk() (*bool, bool)`

GetIsMerchantOk returns a tuple with the IsMerchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMerchant

`func (o *Identity) SetIsMerchant(v bool)`

SetIsMerchant sets IsMerchant field to given value.

### HasIsMerchant

`func (o *Identity) HasIsMerchant() bool`

HasIsMerchant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


