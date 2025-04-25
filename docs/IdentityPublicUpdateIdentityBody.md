# IdentityPublicUpdateIdentityBody

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
**CustomerDueDiligence** | Pointer to [**CustomerDueDiligence**](CustomerDueDiligence.md) |  | [optional] 
**IsMerchant** | Pointer to **bool** | Set to true to indicate that this identity is a merchant. | [optional] 
**LastKycRefreshDate** | Pointer to **time.Time** | Set to the timestamp the identity has last undergone a periodic kyc refresh. If unset, the update is not for periodic kyc refresh. RFC3339 format, like &#x60;YYYY-MM-DDTHH:MM:SS.sssZ&#x60;. ex: &#x60;2006-01-02T15:04:05Z&#x60;. | [optional] 

## Methods

### NewIdentityPublicUpdateIdentityBody

`func NewIdentityPublicUpdateIdentityBody() *IdentityPublicUpdateIdentityBody`

NewIdentityPublicUpdateIdentityBody instantiates a new IdentityPublicUpdateIdentityBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityPublicUpdateIdentityBodyWithDefaults

`func NewIdentityPublicUpdateIdentityBodyWithDefaults() *IdentityPublicUpdateIdentityBody`

NewIdentityPublicUpdateIdentityBodyWithDefaults instantiates a new IdentityPublicUpdateIdentityBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersonDetails

`func (o *IdentityPublicUpdateIdentityBody) GetPersonDetails() PersonDetails`

GetPersonDetails returns the PersonDetails field if non-nil, zero value otherwise.

### GetPersonDetailsOk

`func (o *IdentityPublicUpdateIdentityBody) GetPersonDetailsOk() (*PersonDetails, bool)`

GetPersonDetailsOk returns a tuple with the PersonDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonDetails

`func (o *IdentityPublicUpdateIdentityBody) SetPersonDetails(v PersonDetails)`

SetPersonDetails sets PersonDetails field to given value.

### HasPersonDetails

`func (o *IdentityPublicUpdateIdentityBody) HasPersonDetails() bool`

HasPersonDetails returns a boolean if a field has been set.

### GetMetadata

`func (o *IdentityPublicUpdateIdentityBody) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IdentityPublicUpdateIdentityBody) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IdentityPublicUpdateIdentityBody) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IdentityPublicUpdateIdentityBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSetUserDisabled

`func (o *IdentityPublicUpdateIdentityBody) GetSetUserDisabled() bool`

GetSetUserDisabled returns the SetUserDisabled field if non-nil, zero value otherwise.

### GetSetUserDisabledOk

`func (o *IdentityPublicUpdateIdentityBody) GetSetUserDisabledOk() (*bool, bool)`

GetSetUserDisabledOk returns a tuple with the SetUserDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetUserDisabled

`func (o *IdentityPublicUpdateIdentityBody) SetSetUserDisabled(v bool)`

SetSetUserDisabled sets SetUserDisabled field to given value.

### HasSetUserDisabled

`func (o *IdentityPublicUpdateIdentityBody) HasSetUserDisabled() bool`

HasSetUserDisabled returns a boolean if a field has been set.

### GetInstitutionDetails

`func (o *IdentityPublicUpdateIdentityBody) GetInstitutionDetails() InstitutionDetails`

GetInstitutionDetails returns the InstitutionDetails field if non-nil, zero value otherwise.

### GetInstitutionDetailsOk

`func (o *IdentityPublicUpdateIdentityBody) GetInstitutionDetailsOk() (*InstitutionDetails, bool)`

GetInstitutionDetailsOk returns a tuple with the InstitutionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionDetails

`func (o *IdentityPublicUpdateIdentityBody) SetInstitutionDetails(v InstitutionDetails)`

SetInstitutionDetails sets InstitutionDetails field to given value.

### HasInstitutionDetails

`func (o *IdentityPublicUpdateIdentityBody) HasInstitutionDetails() bool`

HasInstitutionDetails returns a boolean if a field has been set.

### GetRefId

`func (o *IdentityPublicUpdateIdentityBody) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *IdentityPublicUpdateIdentityBody) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *IdentityPublicUpdateIdentityBody) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *IdentityPublicUpdateIdentityBody) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetTaxDetails

`func (o *IdentityPublicUpdateIdentityBody) GetTaxDetails() []TaxDetail`

GetTaxDetails returns the TaxDetails field if non-nil, zero value otherwise.

### GetTaxDetailsOk

`func (o *IdentityPublicUpdateIdentityBody) GetTaxDetailsOk() (*[]TaxDetail, bool)`

GetTaxDetailsOk returns a tuple with the TaxDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDetails

`func (o *IdentityPublicUpdateIdentityBody) SetTaxDetails(v []TaxDetail)`

SetTaxDetails sets TaxDetails field to given value.

### HasTaxDetails

`func (o *IdentityPublicUpdateIdentityBody) HasTaxDetails() bool`

HasTaxDetails returns a boolean if a field has been set.

### GetSetTaxDetailsNotRequired

`func (o *IdentityPublicUpdateIdentityBody) GetSetTaxDetailsNotRequired() bool`

GetSetTaxDetailsNotRequired returns the SetTaxDetailsNotRequired field if non-nil, zero value otherwise.

### GetSetTaxDetailsNotRequiredOk

`func (o *IdentityPublicUpdateIdentityBody) GetSetTaxDetailsNotRequiredOk() (*bool, bool)`

GetSetTaxDetailsNotRequiredOk returns a tuple with the SetTaxDetailsNotRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetTaxDetailsNotRequired

`func (o *IdentityPublicUpdateIdentityBody) SetSetTaxDetailsNotRequired(v bool)`

SetSetTaxDetailsNotRequired sets SetTaxDetailsNotRequired field to given value.

### HasSetTaxDetailsNotRequired

`func (o *IdentityPublicUpdateIdentityBody) HasSetTaxDetailsNotRequired() bool`

HasSetTaxDetailsNotRequired returns a boolean if a field has been set.

### GetCustomerDueDiligence

`func (o *IdentityPublicUpdateIdentityBody) GetCustomerDueDiligence() CustomerDueDiligence`

GetCustomerDueDiligence returns the CustomerDueDiligence field if non-nil, zero value otherwise.

### GetCustomerDueDiligenceOk

`func (o *IdentityPublicUpdateIdentityBody) GetCustomerDueDiligenceOk() (*CustomerDueDiligence, bool)`

GetCustomerDueDiligenceOk returns a tuple with the CustomerDueDiligence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDueDiligence

`func (o *IdentityPublicUpdateIdentityBody) SetCustomerDueDiligence(v CustomerDueDiligence)`

SetCustomerDueDiligence sets CustomerDueDiligence field to given value.

### HasCustomerDueDiligence

`func (o *IdentityPublicUpdateIdentityBody) HasCustomerDueDiligence() bool`

HasCustomerDueDiligence returns a boolean if a field has been set.

### GetIsMerchant

`func (o *IdentityPublicUpdateIdentityBody) GetIsMerchant() bool`

GetIsMerchant returns the IsMerchant field if non-nil, zero value otherwise.

### GetIsMerchantOk

`func (o *IdentityPublicUpdateIdentityBody) GetIsMerchantOk() (*bool, bool)`

GetIsMerchantOk returns a tuple with the IsMerchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMerchant

`func (o *IdentityPublicUpdateIdentityBody) SetIsMerchant(v bool)`

SetIsMerchant sets IsMerchant field to given value.

### HasIsMerchant

`func (o *IdentityPublicUpdateIdentityBody) HasIsMerchant() bool`

HasIsMerchant returns a boolean if a field has been set.

### GetLastKycRefreshDate

`func (o *IdentityPublicUpdateIdentityBody) GetLastKycRefreshDate() time.Time`

GetLastKycRefreshDate returns the LastKycRefreshDate field if non-nil, zero value otherwise.

### GetLastKycRefreshDateOk

`func (o *IdentityPublicUpdateIdentityBody) GetLastKycRefreshDateOk() (*time.Time, bool)`

GetLastKycRefreshDateOk returns a tuple with the LastKycRefreshDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKycRefreshDate

`func (o *IdentityPublicUpdateIdentityBody) SetLastKycRefreshDate(v time.Time)`

SetLastKycRefreshDate sets LastKycRefreshDate field to given value.

### HasLastKycRefreshDate

`func (o *IdentityPublicUpdateIdentityBody) HasLastKycRefreshDate() bool`

HasLastKycRefreshDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


