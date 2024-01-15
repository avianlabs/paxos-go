/*
Paxos API

<p>Welcome to Paxos APIs. At Paxos, our mission is to enable the movement of any asset, any time, in a trustworthy way. These APIs serve that mission by making it easier than ever for you to directly integrate our product capabilities into your application, leveraging the speed, stability, and security of the Paxos platform.</p> <p>The documentation that follows gives you access to our Crypto Brokerage, Trading, and Exchange products. It includes APIs for market data, orders, and the held rate quote flow.</p> <p>To test in our sandbox environment, <a href=\"https://account.sandbox.paxos.com\" target=\"_blank\">sign up</a> for an account. For more information about Paxos and our APIs, visit <a href=\"https://www.paxos.com/\" target=\"_blank\">Paxos.com</a>.</p> 

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paxos

import (
	"encoding/json"
	"time"
)

// checks if the PersonDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonDetails{}

// PersonDetails struct for PersonDetails
type PersonDetails struct {
	IdVerificationStatus *IdentityStatus `json:"id_verification_status,omitempty"`
	SanctionsVerificationStatus *IdentityStatus `json:"sanctions_verification_status,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	DateOfBirth *string `json:"date_of_birth,omitempty"`
	GovtId *string `json:"govt_id,omitempty"`
	Address *IdentityMailingAddress `json:"address,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	Email *string `json:"email,omitempty"`
	// Allowed in create and update. Must be an ISO 3166-1 alpha 3 code.
	Nationality *string `json:"nationality,omitempty"`
	VerifierId *string `json:"verifier_id,omitempty"`
	VerifierType *IdentityprotoVerifierType `json:"verifier_type,omitempty"`
	IdVerificationUrl *string `json:"id_verification_url,omitempty"`
	PassthroughVerifierType *PassthroughVerifierType `json:"passthrough_verifier_type,omitempty"`
	PassthroughVerifiedAt *time.Time `json:"passthrough_verified_at,omitempty"`
	GovtIdType *PersonDetailsCIPIDType `json:"govt_id_type,omitempty"`
	CipId *string `json:"cip_id,omitempty"`
	CipIdType *PersonDetailsCIPIDType `json:"cip_id_type,omitempty"`
	// Allowed in create and update. Must be an ISO 3166-1 alpha 3 code.
	CipIdCountry *string `json:"cip_id_country,omitempty"`
	AdditionalScreeningStatus *IdentityStatus `json:"additional_screening_status,omitempty"`
	// Allowed in create and update.
	Profession *string `json:"profession,omitempty"`
	MiddleName *string `json:"middle_name,omitempty"`
	// Allowed in create and update.
	CountryOfBirth *string `json:"country_of_birth,omitempty"`
}

// NewPersonDetails instantiates a new PersonDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonDetails() *PersonDetails {
	this := PersonDetails{}
	return &this
}

// NewPersonDetailsWithDefaults instantiates a new PersonDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonDetailsWithDefaults() *PersonDetails {
	this := PersonDetails{}
	return &this
}

// GetIdVerificationStatus returns the IdVerificationStatus field value if set, zero value otherwise.
func (o *PersonDetails) GetIdVerificationStatus() IdentityStatus {
	if o == nil || IsNil(o.IdVerificationStatus) {
		var ret IdentityStatus
		return ret
	}
	return *o.IdVerificationStatus
}

// GetIdVerificationStatusOk returns a tuple with the IdVerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetIdVerificationStatusOk() (*IdentityStatus, bool) {
	if o == nil || IsNil(o.IdVerificationStatus) {
		return nil, false
	}
	return o.IdVerificationStatus, true
}

// HasIdVerificationStatus returns a boolean if a field has been set.
func (o *PersonDetails) HasIdVerificationStatus() bool {
	if o != nil && !IsNil(o.IdVerificationStatus) {
		return true
	}

	return false
}

// SetIdVerificationStatus gets a reference to the given IdentityStatus and assigns it to the IdVerificationStatus field.
func (o *PersonDetails) SetIdVerificationStatus(v IdentityStatus) {
	o.IdVerificationStatus = &v
}

// GetSanctionsVerificationStatus returns the SanctionsVerificationStatus field value if set, zero value otherwise.
func (o *PersonDetails) GetSanctionsVerificationStatus() IdentityStatus {
	if o == nil || IsNil(o.SanctionsVerificationStatus) {
		var ret IdentityStatus
		return ret
	}
	return *o.SanctionsVerificationStatus
}

// GetSanctionsVerificationStatusOk returns a tuple with the SanctionsVerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetSanctionsVerificationStatusOk() (*IdentityStatus, bool) {
	if o == nil || IsNil(o.SanctionsVerificationStatus) {
		return nil, false
	}
	return o.SanctionsVerificationStatus, true
}

// HasSanctionsVerificationStatus returns a boolean if a field has been set.
func (o *PersonDetails) HasSanctionsVerificationStatus() bool {
	if o != nil && !IsNil(o.SanctionsVerificationStatus) {
		return true
	}

	return false
}

// SetSanctionsVerificationStatus gets a reference to the given IdentityStatus and assigns it to the SanctionsVerificationStatus field.
func (o *PersonDetails) SetSanctionsVerificationStatus(v IdentityStatus) {
	o.SanctionsVerificationStatus = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *PersonDetails) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PersonDetails) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *PersonDetails) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PersonDetails) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PersonDetails) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PersonDetails) SetLastName(v string) {
	o.LastName = &v
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *PersonDetails) GetDateOfBirth() string {
	if o == nil || IsNil(o.DateOfBirth) {
		var ret string
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetDateOfBirthOk() (*string, bool) {
	if o == nil || IsNil(o.DateOfBirth) {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *PersonDetails) HasDateOfBirth() bool {
	if o != nil && !IsNil(o.DateOfBirth) {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given string and assigns it to the DateOfBirth field.
func (o *PersonDetails) SetDateOfBirth(v string) {
	o.DateOfBirth = &v
}

// GetGovtId returns the GovtId field value if set, zero value otherwise.
func (o *PersonDetails) GetGovtId() string {
	if o == nil || IsNil(o.GovtId) {
		var ret string
		return ret
	}
	return *o.GovtId
}

// GetGovtIdOk returns a tuple with the GovtId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetGovtIdOk() (*string, bool) {
	if o == nil || IsNil(o.GovtId) {
		return nil, false
	}
	return o.GovtId, true
}

// HasGovtId returns a boolean if a field has been set.
func (o *PersonDetails) HasGovtId() bool {
	if o != nil && !IsNil(o.GovtId) {
		return true
	}

	return false
}

// SetGovtId gets a reference to the given string and assigns it to the GovtId field.
func (o *PersonDetails) SetGovtId(v string) {
	o.GovtId = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PersonDetails) GetAddress() IdentityMailingAddress {
	if o == nil || IsNil(o.Address) {
		var ret IdentityMailingAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetAddressOk() (*IdentityMailingAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PersonDetails) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given IdentityMailingAddress and assigns it to the Address field.
func (o *PersonDetails) SetAddress(v IdentityMailingAddress) {
	o.Address = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *PersonDetails) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *PersonDetails) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *PersonDetails) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PersonDetails) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PersonDetails) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PersonDetails) SetEmail(v string) {
	o.Email = &v
}

// GetNationality returns the Nationality field value if set, zero value otherwise.
func (o *PersonDetails) GetNationality() string {
	if o == nil || IsNil(o.Nationality) {
		var ret string
		return ret
	}
	return *o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetNationalityOk() (*string, bool) {
	if o == nil || IsNil(o.Nationality) {
		return nil, false
	}
	return o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *PersonDetails) HasNationality() bool {
	if o != nil && !IsNil(o.Nationality) {
		return true
	}

	return false
}

// SetNationality gets a reference to the given string and assigns it to the Nationality field.
func (o *PersonDetails) SetNationality(v string) {
	o.Nationality = &v
}

// GetVerifierId returns the VerifierId field value if set, zero value otherwise.
func (o *PersonDetails) GetVerifierId() string {
	if o == nil || IsNil(o.VerifierId) {
		var ret string
		return ret
	}
	return *o.VerifierId
}

// GetVerifierIdOk returns a tuple with the VerifierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetVerifierIdOk() (*string, bool) {
	if o == nil || IsNil(o.VerifierId) {
		return nil, false
	}
	return o.VerifierId, true
}

// HasVerifierId returns a boolean if a field has been set.
func (o *PersonDetails) HasVerifierId() bool {
	if o != nil && !IsNil(o.VerifierId) {
		return true
	}

	return false
}

// SetVerifierId gets a reference to the given string and assigns it to the VerifierId field.
func (o *PersonDetails) SetVerifierId(v string) {
	o.VerifierId = &v
}

// GetVerifierType returns the VerifierType field value if set, zero value otherwise.
func (o *PersonDetails) GetVerifierType() IdentityprotoVerifierType {
	if o == nil || IsNil(o.VerifierType) {
		var ret IdentityprotoVerifierType
		return ret
	}
	return *o.VerifierType
}

// GetVerifierTypeOk returns a tuple with the VerifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetVerifierTypeOk() (*IdentityprotoVerifierType, bool) {
	if o == nil || IsNil(o.VerifierType) {
		return nil, false
	}
	return o.VerifierType, true
}

// HasVerifierType returns a boolean if a field has been set.
func (o *PersonDetails) HasVerifierType() bool {
	if o != nil && !IsNil(o.VerifierType) {
		return true
	}

	return false
}

// SetVerifierType gets a reference to the given IdentityprotoVerifierType and assigns it to the VerifierType field.
func (o *PersonDetails) SetVerifierType(v IdentityprotoVerifierType) {
	o.VerifierType = &v
}

// GetIdVerificationUrl returns the IdVerificationUrl field value if set, zero value otherwise.
func (o *PersonDetails) GetIdVerificationUrl() string {
	if o == nil || IsNil(o.IdVerificationUrl) {
		var ret string
		return ret
	}
	return *o.IdVerificationUrl
}

// GetIdVerificationUrlOk returns a tuple with the IdVerificationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetIdVerificationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IdVerificationUrl) {
		return nil, false
	}
	return o.IdVerificationUrl, true
}

// HasIdVerificationUrl returns a boolean if a field has been set.
func (o *PersonDetails) HasIdVerificationUrl() bool {
	if o != nil && !IsNil(o.IdVerificationUrl) {
		return true
	}

	return false
}

// SetIdVerificationUrl gets a reference to the given string and assigns it to the IdVerificationUrl field.
func (o *PersonDetails) SetIdVerificationUrl(v string) {
	o.IdVerificationUrl = &v
}

// GetPassthroughVerifierType returns the PassthroughVerifierType field value if set, zero value otherwise.
func (o *PersonDetails) GetPassthroughVerifierType() PassthroughVerifierType {
	if o == nil || IsNil(o.PassthroughVerifierType) {
		var ret PassthroughVerifierType
		return ret
	}
	return *o.PassthroughVerifierType
}

// GetPassthroughVerifierTypeOk returns a tuple with the PassthroughVerifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetPassthroughVerifierTypeOk() (*PassthroughVerifierType, bool) {
	if o == nil || IsNil(o.PassthroughVerifierType) {
		return nil, false
	}
	return o.PassthroughVerifierType, true
}

// HasPassthroughVerifierType returns a boolean if a field has been set.
func (o *PersonDetails) HasPassthroughVerifierType() bool {
	if o != nil && !IsNil(o.PassthroughVerifierType) {
		return true
	}

	return false
}

// SetPassthroughVerifierType gets a reference to the given PassthroughVerifierType and assigns it to the PassthroughVerifierType field.
func (o *PersonDetails) SetPassthroughVerifierType(v PassthroughVerifierType) {
	o.PassthroughVerifierType = &v
}

// GetPassthroughVerifiedAt returns the PassthroughVerifiedAt field value if set, zero value otherwise.
func (o *PersonDetails) GetPassthroughVerifiedAt() time.Time {
	if o == nil || IsNil(o.PassthroughVerifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.PassthroughVerifiedAt
}

// GetPassthroughVerifiedAtOk returns a tuple with the PassthroughVerifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetPassthroughVerifiedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PassthroughVerifiedAt) {
		return nil, false
	}
	return o.PassthroughVerifiedAt, true
}

// HasPassthroughVerifiedAt returns a boolean if a field has been set.
func (o *PersonDetails) HasPassthroughVerifiedAt() bool {
	if o != nil && !IsNil(o.PassthroughVerifiedAt) {
		return true
	}

	return false
}

// SetPassthroughVerifiedAt gets a reference to the given time.Time and assigns it to the PassthroughVerifiedAt field.
func (o *PersonDetails) SetPassthroughVerifiedAt(v time.Time) {
	o.PassthroughVerifiedAt = &v
}

// GetGovtIdType returns the GovtIdType field value if set, zero value otherwise.
func (o *PersonDetails) GetGovtIdType() PersonDetailsCIPIDType {
	if o == nil || IsNil(o.GovtIdType) {
		var ret PersonDetailsCIPIDType
		return ret
	}
	return *o.GovtIdType
}

// GetGovtIdTypeOk returns a tuple with the GovtIdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetGovtIdTypeOk() (*PersonDetailsCIPIDType, bool) {
	if o == nil || IsNil(o.GovtIdType) {
		return nil, false
	}
	return o.GovtIdType, true
}

// HasGovtIdType returns a boolean if a field has been set.
func (o *PersonDetails) HasGovtIdType() bool {
	if o != nil && !IsNil(o.GovtIdType) {
		return true
	}

	return false
}

// SetGovtIdType gets a reference to the given PersonDetailsCIPIDType and assigns it to the GovtIdType field.
func (o *PersonDetails) SetGovtIdType(v PersonDetailsCIPIDType) {
	o.GovtIdType = &v
}

// GetCipId returns the CipId field value if set, zero value otherwise.
func (o *PersonDetails) GetCipId() string {
	if o == nil || IsNil(o.CipId) {
		var ret string
		return ret
	}
	return *o.CipId
}

// GetCipIdOk returns a tuple with the CipId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetCipIdOk() (*string, bool) {
	if o == nil || IsNil(o.CipId) {
		return nil, false
	}
	return o.CipId, true
}

// HasCipId returns a boolean if a field has been set.
func (o *PersonDetails) HasCipId() bool {
	if o != nil && !IsNil(o.CipId) {
		return true
	}

	return false
}

// SetCipId gets a reference to the given string and assigns it to the CipId field.
func (o *PersonDetails) SetCipId(v string) {
	o.CipId = &v
}

// GetCipIdType returns the CipIdType field value if set, zero value otherwise.
func (o *PersonDetails) GetCipIdType() PersonDetailsCIPIDType {
	if o == nil || IsNil(o.CipIdType) {
		var ret PersonDetailsCIPIDType
		return ret
	}
	return *o.CipIdType
}

// GetCipIdTypeOk returns a tuple with the CipIdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetCipIdTypeOk() (*PersonDetailsCIPIDType, bool) {
	if o == nil || IsNil(o.CipIdType) {
		return nil, false
	}
	return o.CipIdType, true
}

// HasCipIdType returns a boolean if a field has been set.
func (o *PersonDetails) HasCipIdType() bool {
	if o != nil && !IsNil(o.CipIdType) {
		return true
	}

	return false
}

// SetCipIdType gets a reference to the given PersonDetailsCIPIDType and assigns it to the CipIdType field.
func (o *PersonDetails) SetCipIdType(v PersonDetailsCIPIDType) {
	o.CipIdType = &v
}

// GetCipIdCountry returns the CipIdCountry field value if set, zero value otherwise.
func (o *PersonDetails) GetCipIdCountry() string {
	if o == nil || IsNil(o.CipIdCountry) {
		var ret string
		return ret
	}
	return *o.CipIdCountry
}

// GetCipIdCountryOk returns a tuple with the CipIdCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetCipIdCountryOk() (*string, bool) {
	if o == nil || IsNil(o.CipIdCountry) {
		return nil, false
	}
	return o.CipIdCountry, true
}

// HasCipIdCountry returns a boolean if a field has been set.
func (o *PersonDetails) HasCipIdCountry() bool {
	if o != nil && !IsNil(o.CipIdCountry) {
		return true
	}

	return false
}

// SetCipIdCountry gets a reference to the given string and assigns it to the CipIdCountry field.
func (o *PersonDetails) SetCipIdCountry(v string) {
	o.CipIdCountry = &v
}

// GetAdditionalScreeningStatus returns the AdditionalScreeningStatus field value if set, zero value otherwise.
func (o *PersonDetails) GetAdditionalScreeningStatus() IdentityStatus {
	if o == nil || IsNil(o.AdditionalScreeningStatus) {
		var ret IdentityStatus
		return ret
	}
	return *o.AdditionalScreeningStatus
}

// GetAdditionalScreeningStatusOk returns a tuple with the AdditionalScreeningStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetAdditionalScreeningStatusOk() (*IdentityStatus, bool) {
	if o == nil || IsNil(o.AdditionalScreeningStatus) {
		return nil, false
	}
	return o.AdditionalScreeningStatus, true
}

// HasAdditionalScreeningStatus returns a boolean if a field has been set.
func (o *PersonDetails) HasAdditionalScreeningStatus() bool {
	if o != nil && !IsNil(o.AdditionalScreeningStatus) {
		return true
	}

	return false
}

// SetAdditionalScreeningStatus gets a reference to the given IdentityStatus and assigns it to the AdditionalScreeningStatus field.
func (o *PersonDetails) SetAdditionalScreeningStatus(v IdentityStatus) {
	o.AdditionalScreeningStatus = &v
}

// GetProfession returns the Profession field value if set, zero value otherwise.
func (o *PersonDetails) GetProfession() string {
	if o == nil || IsNil(o.Profession) {
		var ret string
		return ret
	}
	return *o.Profession
}

// GetProfessionOk returns a tuple with the Profession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetProfessionOk() (*string, bool) {
	if o == nil || IsNil(o.Profession) {
		return nil, false
	}
	return o.Profession, true
}

// HasProfession returns a boolean if a field has been set.
func (o *PersonDetails) HasProfession() bool {
	if o != nil && !IsNil(o.Profession) {
		return true
	}

	return false
}

// SetProfession gets a reference to the given string and assigns it to the Profession field.
func (o *PersonDetails) SetProfession(v string) {
	o.Profession = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *PersonDetails) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *PersonDetails) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *PersonDetails) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetCountryOfBirth returns the CountryOfBirth field value if set, zero value otherwise.
func (o *PersonDetails) GetCountryOfBirth() string {
	if o == nil || IsNil(o.CountryOfBirth) {
		var ret string
		return ret
	}
	return *o.CountryOfBirth
}

// GetCountryOfBirthOk returns a tuple with the CountryOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetCountryOfBirthOk() (*string, bool) {
	if o == nil || IsNil(o.CountryOfBirth) {
		return nil, false
	}
	return o.CountryOfBirth, true
}

// HasCountryOfBirth returns a boolean if a field has been set.
func (o *PersonDetails) HasCountryOfBirth() bool {
	if o != nil && !IsNil(o.CountryOfBirth) {
		return true
	}

	return false
}

// SetCountryOfBirth gets a reference to the given string and assigns it to the CountryOfBirth field.
func (o *PersonDetails) SetCountryOfBirth(v string) {
	o.CountryOfBirth = &v
}

func (o PersonDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdVerificationStatus) {
		toSerialize["id_verification_status"] = o.IdVerificationStatus
	}
	if !IsNil(o.SanctionsVerificationStatus) {
		toSerialize["sanctions_verification_status"] = o.SanctionsVerificationStatus
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.DateOfBirth) {
		toSerialize["date_of_birth"] = o.DateOfBirth
	}
	if !IsNil(o.GovtId) {
		toSerialize["govt_id"] = o.GovtId
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Nationality) {
		toSerialize["nationality"] = o.Nationality
	}
	if !IsNil(o.VerifierId) {
		toSerialize["verifier_id"] = o.VerifierId
	}
	if !IsNil(o.VerifierType) {
		toSerialize["verifier_type"] = o.VerifierType
	}
	if !IsNil(o.IdVerificationUrl) {
		toSerialize["id_verification_url"] = o.IdVerificationUrl
	}
	if !IsNil(o.PassthroughVerifierType) {
		toSerialize["passthrough_verifier_type"] = o.PassthroughVerifierType
	}
	if !IsNil(o.PassthroughVerifiedAt) {
		toSerialize["passthrough_verified_at"] = o.PassthroughVerifiedAt
	}
	if !IsNil(o.GovtIdType) {
		toSerialize["govt_id_type"] = o.GovtIdType
	}
	if !IsNil(o.CipId) {
		toSerialize["cip_id"] = o.CipId
	}
	if !IsNil(o.CipIdType) {
		toSerialize["cip_id_type"] = o.CipIdType
	}
	if !IsNil(o.CipIdCountry) {
		toSerialize["cip_id_country"] = o.CipIdCountry
	}
	if !IsNil(o.AdditionalScreeningStatus) {
		toSerialize["additional_screening_status"] = o.AdditionalScreeningStatus
	}
	if !IsNil(o.Profession) {
		toSerialize["profession"] = o.Profession
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middle_name"] = o.MiddleName
	}
	if !IsNil(o.CountryOfBirth) {
		toSerialize["country_of_birth"] = o.CountryOfBirth
	}
	return toSerialize, nil
}

type NullablePersonDetails struct {
	value *PersonDetails
	isSet bool
}

func (v NullablePersonDetails) Get() *PersonDetails {
	return v.value
}

func (v *NullablePersonDetails) Set(val *PersonDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonDetails(val *PersonDetails) *NullablePersonDetails {
	return &NullablePersonDetails{value: val, isSet: true}
}

func (v NullablePersonDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


