# PersonDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdVerificationStatus** | Pointer to [**IdentityStatus**](IdentityStatus.md) |  | [optional] 
**SanctionsVerificationStatus** | Pointer to [**IdentityStatus**](IdentityStatus.md) |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**DateOfBirth** | Pointer to **string** |  | [optional] 
**GovtId** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**IdentityMailingAddress**](IdentityMailingAddress.md) |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Nationality** | Pointer to **string** | Allowed in create and update. Must be an ISO 3166-1 alpha 3 code. | [optional] 
**VerifierId** | Pointer to **string** |  | [optional] 
**VerifierType** | Pointer to [**IdentityprotoVerifierType**](IdentityprotoVerifierType.md) |  | [optional] 
**IdVerificationUrl** | Pointer to **string** |  | [optional] 
**PassthroughVerifierType** | Pointer to [**PassthroughVerifierType**](PassthroughVerifierType.md) |  | [optional] 
**PassthroughVerifiedAt** | Pointer to **time.Time** |  | [optional] 
**GovtIdType** | Pointer to [**PersonDetailsCIPIDType**](PersonDetailsCIPIDType.md) |  | [optional] 
**CipId** | Pointer to **string** |  | [optional] 
**CipIdType** | Pointer to [**PersonDetailsCIPIDType**](PersonDetailsCIPIDType.md) |  | [optional] 
**CipIdCountry** | Pointer to **string** | Allowed in create and update. Must be an ISO 3166-1 alpha 3 code. | [optional] 
**AdditionalScreeningStatus** | Pointer to [**IdentityStatus**](IdentityStatus.md) |  | [optional] 
**Profession** | Pointer to **string** | Allowed in create and update. | [optional] 
**MiddleName** | Pointer to **string** |  | [optional] 
**CountryOfBirth** | Pointer to **string** | Allowed in create and update. | [optional] 

## Methods

### NewPersonDetails

`func NewPersonDetails() *PersonDetails`

NewPersonDetails instantiates a new PersonDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonDetailsWithDefaults

`func NewPersonDetailsWithDefaults() *PersonDetails`

NewPersonDetailsWithDefaults instantiates a new PersonDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdVerificationStatus

`func (o *PersonDetails) GetIdVerificationStatus() IdentityStatus`

GetIdVerificationStatus returns the IdVerificationStatus field if non-nil, zero value otherwise.

### GetIdVerificationStatusOk

`func (o *PersonDetails) GetIdVerificationStatusOk() (*IdentityStatus, bool)`

GetIdVerificationStatusOk returns a tuple with the IdVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdVerificationStatus

`func (o *PersonDetails) SetIdVerificationStatus(v IdentityStatus)`

SetIdVerificationStatus sets IdVerificationStatus field to given value.

### HasIdVerificationStatus

`func (o *PersonDetails) HasIdVerificationStatus() bool`

HasIdVerificationStatus returns a boolean if a field has been set.

### GetSanctionsVerificationStatus

`func (o *PersonDetails) GetSanctionsVerificationStatus() IdentityStatus`

GetSanctionsVerificationStatus returns the SanctionsVerificationStatus field if non-nil, zero value otherwise.

### GetSanctionsVerificationStatusOk

`func (o *PersonDetails) GetSanctionsVerificationStatusOk() (*IdentityStatus, bool)`

GetSanctionsVerificationStatusOk returns a tuple with the SanctionsVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctionsVerificationStatus

`func (o *PersonDetails) SetSanctionsVerificationStatus(v IdentityStatus)`

SetSanctionsVerificationStatus sets SanctionsVerificationStatus field to given value.

### HasSanctionsVerificationStatus

`func (o *PersonDetails) HasSanctionsVerificationStatus() bool`

HasSanctionsVerificationStatus returns a boolean if a field has been set.

### GetFirstName

`func (o *PersonDetails) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PersonDetails) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PersonDetails) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PersonDetails) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PersonDetails) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PersonDetails) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PersonDetails) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PersonDetails) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *PersonDetails) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *PersonDetails) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *PersonDetails) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *PersonDetails) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetGovtId

`func (o *PersonDetails) GetGovtId() string`

GetGovtId returns the GovtId field if non-nil, zero value otherwise.

### GetGovtIdOk

`func (o *PersonDetails) GetGovtIdOk() (*string, bool)`

GetGovtIdOk returns a tuple with the GovtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovtId

`func (o *PersonDetails) SetGovtId(v string)`

SetGovtId sets GovtId field to given value.

### HasGovtId

`func (o *PersonDetails) HasGovtId() bool`

HasGovtId returns a boolean if a field has been set.

### GetAddress

`func (o *PersonDetails) GetAddress() IdentityMailingAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PersonDetails) GetAddressOk() (*IdentityMailingAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PersonDetails) SetAddress(v IdentityMailingAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PersonDetails) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *PersonDetails) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PersonDetails) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PersonDetails) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PersonDetails) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmail

`func (o *PersonDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PersonDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PersonDetails) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PersonDetails) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetNationality

`func (o *PersonDetails) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *PersonDetails) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *PersonDetails) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *PersonDetails) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetVerifierId

`func (o *PersonDetails) GetVerifierId() string`

GetVerifierId returns the VerifierId field if non-nil, zero value otherwise.

### GetVerifierIdOk

`func (o *PersonDetails) GetVerifierIdOk() (*string, bool)`

GetVerifierIdOk returns a tuple with the VerifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifierId

`func (o *PersonDetails) SetVerifierId(v string)`

SetVerifierId sets VerifierId field to given value.

### HasVerifierId

`func (o *PersonDetails) HasVerifierId() bool`

HasVerifierId returns a boolean if a field has been set.

### GetVerifierType

`func (o *PersonDetails) GetVerifierType() IdentityprotoVerifierType`

GetVerifierType returns the VerifierType field if non-nil, zero value otherwise.

### GetVerifierTypeOk

`func (o *PersonDetails) GetVerifierTypeOk() (*IdentityprotoVerifierType, bool)`

GetVerifierTypeOk returns a tuple with the VerifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifierType

`func (o *PersonDetails) SetVerifierType(v IdentityprotoVerifierType)`

SetVerifierType sets VerifierType field to given value.

### HasVerifierType

`func (o *PersonDetails) HasVerifierType() bool`

HasVerifierType returns a boolean if a field has been set.

### GetIdVerificationUrl

`func (o *PersonDetails) GetIdVerificationUrl() string`

GetIdVerificationUrl returns the IdVerificationUrl field if non-nil, zero value otherwise.

### GetIdVerificationUrlOk

`func (o *PersonDetails) GetIdVerificationUrlOk() (*string, bool)`

GetIdVerificationUrlOk returns a tuple with the IdVerificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdVerificationUrl

`func (o *PersonDetails) SetIdVerificationUrl(v string)`

SetIdVerificationUrl sets IdVerificationUrl field to given value.

### HasIdVerificationUrl

`func (o *PersonDetails) HasIdVerificationUrl() bool`

HasIdVerificationUrl returns a boolean if a field has been set.

### GetPassthroughVerifierType

`func (o *PersonDetails) GetPassthroughVerifierType() PassthroughVerifierType`

GetPassthroughVerifierType returns the PassthroughVerifierType field if non-nil, zero value otherwise.

### GetPassthroughVerifierTypeOk

`func (o *PersonDetails) GetPassthroughVerifierTypeOk() (*PassthroughVerifierType, bool)`

GetPassthroughVerifierTypeOk returns a tuple with the PassthroughVerifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthroughVerifierType

`func (o *PersonDetails) SetPassthroughVerifierType(v PassthroughVerifierType)`

SetPassthroughVerifierType sets PassthroughVerifierType field to given value.

### HasPassthroughVerifierType

`func (o *PersonDetails) HasPassthroughVerifierType() bool`

HasPassthroughVerifierType returns a boolean if a field has been set.

### GetPassthroughVerifiedAt

`func (o *PersonDetails) GetPassthroughVerifiedAt() time.Time`

GetPassthroughVerifiedAt returns the PassthroughVerifiedAt field if non-nil, zero value otherwise.

### GetPassthroughVerifiedAtOk

`func (o *PersonDetails) GetPassthroughVerifiedAtOk() (*time.Time, bool)`

GetPassthroughVerifiedAtOk returns a tuple with the PassthroughVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassthroughVerifiedAt

`func (o *PersonDetails) SetPassthroughVerifiedAt(v time.Time)`

SetPassthroughVerifiedAt sets PassthroughVerifiedAt field to given value.

### HasPassthroughVerifiedAt

`func (o *PersonDetails) HasPassthroughVerifiedAt() bool`

HasPassthroughVerifiedAt returns a boolean if a field has been set.

### GetGovtIdType

`func (o *PersonDetails) GetGovtIdType() PersonDetailsCIPIDType`

GetGovtIdType returns the GovtIdType field if non-nil, zero value otherwise.

### GetGovtIdTypeOk

`func (o *PersonDetails) GetGovtIdTypeOk() (*PersonDetailsCIPIDType, bool)`

GetGovtIdTypeOk returns a tuple with the GovtIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovtIdType

`func (o *PersonDetails) SetGovtIdType(v PersonDetailsCIPIDType)`

SetGovtIdType sets GovtIdType field to given value.

### HasGovtIdType

`func (o *PersonDetails) HasGovtIdType() bool`

HasGovtIdType returns a boolean if a field has been set.

### GetCipId

`func (o *PersonDetails) GetCipId() string`

GetCipId returns the CipId field if non-nil, zero value otherwise.

### GetCipIdOk

`func (o *PersonDetails) GetCipIdOk() (*string, bool)`

GetCipIdOk returns a tuple with the CipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipId

`func (o *PersonDetails) SetCipId(v string)`

SetCipId sets CipId field to given value.

### HasCipId

`func (o *PersonDetails) HasCipId() bool`

HasCipId returns a boolean if a field has been set.

### GetCipIdType

`func (o *PersonDetails) GetCipIdType() PersonDetailsCIPIDType`

GetCipIdType returns the CipIdType field if non-nil, zero value otherwise.

### GetCipIdTypeOk

`func (o *PersonDetails) GetCipIdTypeOk() (*PersonDetailsCIPIDType, bool)`

GetCipIdTypeOk returns a tuple with the CipIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipIdType

`func (o *PersonDetails) SetCipIdType(v PersonDetailsCIPIDType)`

SetCipIdType sets CipIdType field to given value.

### HasCipIdType

`func (o *PersonDetails) HasCipIdType() bool`

HasCipIdType returns a boolean if a field has been set.

### GetCipIdCountry

`func (o *PersonDetails) GetCipIdCountry() string`

GetCipIdCountry returns the CipIdCountry field if non-nil, zero value otherwise.

### GetCipIdCountryOk

`func (o *PersonDetails) GetCipIdCountryOk() (*string, bool)`

GetCipIdCountryOk returns a tuple with the CipIdCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipIdCountry

`func (o *PersonDetails) SetCipIdCountry(v string)`

SetCipIdCountry sets CipIdCountry field to given value.

### HasCipIdCountry

`func (o *PersonDetails) HasCipIdCountry() bool`

HasCipIdCountry returns a boolean if a field has been set.

### GetAdditionalScreeningStatus

`func (o *PersonDetails) GetAdditionalScreeningStatus() IdentityStatus`

GetAdditionalScreeningStatus returns the AdditionalScreeningStatus field if non-nil, zero value otherwise.

### GetAdditionalScreeningStatusOk

`func (o *PersonDetails) GetAdditionalScreeningStatusOk() (*IdentityStatus, bool)`

GetAdditionalScreeningStatusOk returns a tuple with the AdditionalScreeningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalScreeningStatus

`func (o *PersonDetails) SetAdditionalScreeningStatus(v IdentityStatus)`

SetAdditionalScreeningStatus sets AdditionalScreeningStatus field to given value.

### HasAdditionalScreeningStatus

`func (o *PersonDetails) HasAdditionalScreeningStatus() bool`

HasAdditionalScreeningStatus returns a boolean if a field has been set.

### GetProfession

`func (o *PersonDetails) GetProfession() string`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *PersonDetails) GetProfessionOk() (*string, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfession

`func (o *PersonDetails) SetProfession(v string)`

SetProfession sets Profession field to given value.

### HasProfession

`func (o *PersonDetails) HasProfession() bool`

HasProfession returns a boolean if a field has been set.

### GetMiddleName

`func (o *PersonDetails) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *PersonDetails) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *PersonDetails) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *PersonDetails) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetCountryOfBirth

`func (o *PersonDetails) GetCountryOfBirth() string`

GetCountryOfBirth returns the CountryOfBirth field if non-nil, zero value otherwise.

### GetCountryOfBirthOk

`func (o *PersonDetails) GetCountryOfBirthOk() (*string, bool)`

GetCountryOfBirthOk returns a tuple with the CountryOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfBirth

`func (o *PersonDetails) SetCountryOfBirth(v string)`

SetCountryOfBirth sets CountryOfBirth field to given value.

### HasCountryOfBirth

`func (o *PersonDetails) HasCountryOfBirth() bool`

HasCountryOfBirth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


