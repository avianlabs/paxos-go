# InstitutionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SanctionsVerificationStatus** | Pointer to [**IdentityStatus**](IdentityStatus.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**BusinessAddress** | Pointer to [**IdentityMailingAddress**](IdentityMailingAddress.md) |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**InstitutionType** | Pointer to [**InstitutionType**](InstitutionType.md) |  | [optional] 
**InstitutionSubType** | Pointer to [**InstitutionSubType**](InstitutionSubType.md) |  | [optional] 
**CipId** | Pointer to **string** |  | [optional] 
**CipIdType** | Pointer to [**InstitutionCIPIDType**](InstitutionCIPIDType.md) |  | [optional] 
**CipIdCountry** | Pointer to **string** | Allowed in create and update. Must be an ISO 3166-1 alpha 3 code. | [optional] 
**GovtRegistrationDate** | Pointer to **time.Time** |  | [optional] 
**IncorporationAddress** | Pointer to [**IdentityMailingAddress**](IdentityMailingAddress.md) |  | [optional] 
**RegulationStatus** | Pointer to [**RegulationStatus**](RegulationStatus.md) |  | [optional] 
**TradingType** | Pointer to [**TradingType**](TradingType.md) |  | [optional] 
**ListedExchange** | Pointer to **string** |  | [optional] 
**TickerSymbol** | Pointer to **string** | Ticker symbol of the institution if publicly traded or ticker symbol of the parent institution. | [optional] 
**ParentInstitutionName** | Pointer to **string** |  | [optional] 
**RegulatorName** | Pointer to **string** |  | [optional] 
**RegulatorJurisdiction** | Pointer to **string** |  | [optional] 
**RegulatorRegisterNumber** | Pointer to **string** |  | [optional] 
**DocumentVerificationStatus** | Pointer to [**IdentityStatus**](IdentityStatus.md) |  | [optional] 
**AdditionalScreeningStatus** | Pointer to [**IdentityStatus**](IdentityStatus.md) |  | [optional] 
**DoingBusinessAs** | Pointer to **string** |  | [optional] 

## Methods

### NewInstitutionDetails

`func NewInstitutionDetails() *InstitutionDetails`

NewInstitutionDetails instantiates a new InstitutionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionDetailsWithDefaults

`func NewInstitutionDetailsWithDefaults() *InstitutionDetails`

NewInstitutionDetailsWithDefaults instantiates a new InstitutionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSanctionsVerificationStatus

`func (o *InstitutionDetails) GetSanctionsVerificationStatus() IdentityStatus`

GetSanctionsVerificationStatus returns the SanctionsVerificationStatus field if non-nil, zero value otherwise.

### GetSanctionsVerificationStatusOk

`func (o *InstitutionDetails) GetSanctionsVerificationStatusOk() (*IdentityStatus, bool)`

GetSanctionsVerificationStatusOk returns a tuple with the SanctionsVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctionsVerificationStatus

`func (o *InstitutionDetails) SetSanctionsVerificationStatus(v IdentityStatus)`

SetSanctionsVerificationStatus sets SanctionsVerificationStatus field to given value.

### HasSanctionsVerificationStatus

`func (o *InstitutionDetails) HasSanctionsVerificationStatus() bool`

HasSanctionsVerificationStatus returns a boolean if a field has been set.

### GetName

`func (o *InstitutionDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstitutionDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstitutionDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstitutionDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBusinessAddress

`func (o *InstitutionDetails) GetBusinessAddress() IdentityMailingAddress`

GetBusinessAddress returns the BusinessAddress field if non-nil, zero value otherwise.

### GetBusinessAddressOk

`func (o *InstitutionDetails) GetBusinessAddressOk() (*IdentityMailingAddress, bool)`

GetBusinessAddressOk returns a tuple with the BusinessAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAddress

`func (o *InstitutionDetails) SetBusinessAddress(v IdentityMailingAddress)`

SetBusinessAddress sets BusinessAddress field to given value.

### HasBusinessAddress

`func (o *InstitutionDetails) HasBusinessAddress() bool`

HasBusinessAddress returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *InstitutionDetails) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *InstitutionDetails) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *InstitutionDetails) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *InstitutionDetails) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmail

`func (o *InstitutionDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InstitutionDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InstitutionDetails) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InstitutionDetails) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetInstitutionType

`func (o *InstitutionDetails) GetInstitutionType() InstitutionType`

GetInstitutionType returns the InstitutionType field if non-nil, zero value otherwise.

### GetInstitutionTypeOk

`func (o *InstitutionDetails) GetInstitutionTypeOk() (*InstitutionType, bool)`

GetInstitutionTypeOk returns a tuple with the InstitutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionType

`func (o *InstitutionDetails) SetInstitutionType(v InstitutionType)`

SetInstitutionType sets InstitutionType field to given value.

### HasInstitutionType

`func (o *InstitutionDetails) HasInstitutionType() bool`

HasInstitutionType returns a boolean if a field has been set.

### GetInstitutionSubType

`func (o *InstitutionDetails) GetInstitutionSubType() InstitutionSubType`

GetInstitutionSubType returns the InstitutionSubType field if non-nil, zero value otherwise.

### GetInstitutionSubTypeOk

`func (o *InstitutionDetails) GetInstitutionSubTypeOk() (*InstitutionSubType, bool)`

GetInstitutionSubTypeOk returns a tuple with the InstitutionSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionSubType

`func (o *InstitutionDetails) SetInstitutionSubType(v InstitutionSubType)`

SetInstitutionSubType sets InstitutionSubType field to given value.

### HasInstitutionSubType

`func (o *InstitutionDetails) HasInstitutionSubType() bool`

HasInstitutionSubType returns a boolean if a field has been set.

### GetCipId

`func (o *InstitutionDetails) GetCipId() string`

GetCipId returns the CipId field if non-nil, zero value otherwise.

### GetCipIdOk

`func (o *InstitutionDetails) GetCipIdOk() (*string, bool)`

GetCipIdOk returns a tuple with the CipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipId

`func (o *InstitutionDetails) SetCipId(v string)`

SetCipId sets CipId field to given value.

### HasCipId

`func (o *InstitutionDetails) HasCipId() bool`

HasCipId returns a boolean if a field has been set.

### GetCipIdType

`func (o *InstitutionDetails) GetCipIdType() InstitutionCIPIDType`

GetCipIdType returns the CipIdType field if non-nil, zero value otherwise.

### GetCipIdTypeOk

`func (o *InstitutionDetails) GetCipIdTypeOk() (*InstitutionCIPIDType, bool)`

GetCipIdTypeOk returns a tuple with the CipIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipIdType

`func (o *InstitutionDetails) SetCipIdType(v InstitutionCIPIDType)`

SetCipIdType sets CipIdType field to given value.

### HasCipIdType

`func (o *InstitutionDetails) HasCipIdType() bool`

HasCipIdType returns a boolean if a field has been set.

### GetCipIdCountry

`func (o *InstitutionDetails) GetCipIdCountry() string`

GetCipIdCountry returns the CipIdCountry field if non-nil, zero value otherwise.

### GetCipIdCountryOk

`func (o *InstitutionDetails) GetCipIdCountryOk() (*string, bool)`

GetCipIdCountryOk returns a tuple with the CipIdCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipIdCountry

`func (o *InstitutionDetails) SetCipIdCountry(v string)`

SetCipIdCountry sets CipIdCountry field to given value.

### HasCipIdCountry

`func (o *InstitutionDetails) HasCipIdCountry() bool`

HasCipIdCountry returns a boolean if a field has been set.

### GetGovtRegistrationDate

`func (o *InstitutionDetails) GetGovtRegistrationDate() time.Time`

GetGovtRegistrationDate returns the GovtRegistrationDate field if non-nil, zero value otherwise.

### GetGovtRegistrationDateOk

`func (o *InstitutionDetails) GetGovtRegistrationDateOk() (*time.Time, bool)`

GetGovtRegistrationDateOk returns a tuple with the GovtRegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovtRegistrationDate

`func (o *InstitutionDetails) SetGovtRegistrationDate(v time.Time)`

SetGovtRegistrationDate sets GovtRegistrationDate field to given value.

### HasGovtRegistrationDate

`func (o *InstitutionDetails) HasGovtRegistrationDate() bool`

HasGovtRegistrationDate returns a boolean if a field has been set.

### GetIncorporationAddress

`func (o *InstitutionDetails) GetIncorporationAddress() IdentityMailingAddress`

GetIncorporationAddress returns the IncorporationAddress field if non-nil, zero value otherwise.

### GetIncorporationAddressOk

`func (o *InstitutionDetails) GetIncorporationAddressOk() (*IdentityMailingAddress, bool)`

GetIncorporationAddressOk returns a tuple with the IncorporationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncorporationAddress

`func (o *InstitutionDetails) SetIncorporationAddress(v IdentityMailingAddress)`

SetIncorporationAddress sets IncorporationAddress field to given value.

### HasIncorporationAddress

`func (o *InstitutionDetails) HasIncorporationAddress() bool`

HasIncorporationAddress returns a boolean if a field has been set.

### GetRegulationStatus

`func (o *InstitutionDetails) GetRegulationStatus() RegulationStatus`

GetRegulationStatus returns the RegulationStatus field if non-nil, zero value otherwise.

### GetRegulationStatusOk

`func (o *InstitutionDetails) GetRegulationStatusOk() (*RegulationStatus, bool)`

GetRegulationStatusOk returns a tuple with the RegulationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulationStatus

`func (o *InstitutionDetails) SetRegulationStatus(v RegulationStatus)`

SetRegulationStatus sets RegulationStatus field to given value.

### HasRegulationStatus

`func (o *InstitutionDetails) HasRegulationStatus() bool`

HasRegulationStatus returns a boolean if a field has been set.

### GetTradingType

`func (o *InstitutionDetails) GetTradingType() TradingType`

GetTradingType returns the TradingType field if non-nil, zero value otherwise.

### GetTradingTypeOk

`func (o *InstitutionDetails) GetTradingTypeOk() (*TradingType, bool)`

GetTradingTypeOk returns a tuple with the TradingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingType

`func (o *InstitutionDetails) SetTradingType(v TradingType)`

SetTradingType sets TradingType field to given value.

### HasTradingType

`func (o *InstitutionDetails) HasTradingType() bool`

HasTradingType returns a boolean if a field has been set.

### GetListedExchange

`func (o *InstitutionDetails) GetListedExchange() string`

GetListedExchange returns the ListedExchange field if non-nil, zero value otherwise.

### GetListedExchangeOk

`func (o *InstitutionDetails) GetListedExchangeOk() (*string, bool)`

GetListedExchangeOk returns a tuple with the ListedExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedExchange

`func (o *InstitutionDetails) SetListedExchange(v string)`

SetListedExchange sets ListedExchange field to given value.

### HasListedExchange

`func (o *InstitutionDetails) HasListedExchange() bool`

HasListedExchange returns a boolean if a field has been set.

### GetTickerSymbol

`func (o *InstitutionDetails) GetTickerSymbol() string`

GetTickerSymbol returns the TickerSymbol field if non-nil, zero value otherwise.

### GetTickerSymbolOk

`func (o *InstitutionDetails) GetTickerSymbolOk() (*string, bool)`

GetTickerSymbolOk returns a tuple with the TickerSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerSymbol

`func (o *InstitutionDetails) SetTickerSymbol(v string)`

SetTickerSymbol sets TickerSymbol field to given value.

### HasTickerSymbol

`func (o *InstitutionDetails) HasTickerSymbol() bool`

HasTickerSymbol returns a boolean if a field has been set.

### GetParentInstitutionName

`func (o *InstitutionDetails) GetParentInstitutionName() string`

GetParentInstitutionName returns the ParentInstitutionName field if non-nil, zero value otherwise.

### GetParentInstitutionNameOk

`func (o *InstitutionDetails) GetParentInstitutionNameOk() (*string, bool)`

GetParentInstitutionNameOk returns a tuple with the ParentInstitutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentInstitutionName

`func (o *InstitutionDetails) SetParentInstitutionName(v string)`

SetParentInstitutionName sets ParentInstitutionName field to given value.

### HasParentInstitutionName

`func (o *InstitutionDetails) HasParentInstitutionName() bool`

HasParentInstitutionName returns a boolean if a field has been set.

### GetRegulatorName

`func (o *InstitutionDetails) GetRegulatorName() string`

GetRegulatorName returns the RegulatorName field if non-nil, zero value otherwise.

### GetRegulatorNameOk

`func (o *InstitutionDetails) GetRegulatorNameOk() (*string, bool)`

GetRegulatorNameOk returns a tuple with the RegulatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatorName

`func (o *InstitutionDetails) SetRegulatorName(v string)`

SetRegulatorName sets RegulatorName field to given value.

### HasRegulatorName

`func (o *InstitutionDetails) HasRegulatorName() bool`

HasRegulatorName returns a boolean if a field has been set.

### GetRegulatorJurisdiction

`func (o *InstitutionDetails) GetRegulatorJurisdiction() string`

GetRegulatorJurisdiction returns the RegulatorJurisdiction field if non-nil, zero value otherwise.

### GetRegulatorJurisdictionOk

`func (o *InstitutionDetails) GetRegulatorJurisdictionOk() (*string, bool)`

GetRegulatorJurisdictionOk returns a tuple with the RegulatorJurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatorJurisdiction

`func (o *InstitutionDetails) SetRegulatorJurisdiction(v string)`

SetRegulatorJurisdiction sets RegulatorJurisdiction field to given value.

### HasRegulatorJurisdiction

`func (o *InstitutionDetails) HasRegulatorJurisdiction() bool`

HasRegulatorJurisdiction returns a boolean if a field has been set.

### GetRegulatorRegisterNumber

`func (o *InstitutionDetails) GetRegulatorRegisterNumber() string`

GetRegulatorRegisterNumber returns the RegulatorRegisterNumber field if non-nil, zero value otherwise.

### GetRegulatorRegisterNumberOk

`func (o *InstitutionDetails) GetRegulatorRegisterNumberOk() (*string, bool)`

GetRegulatorRegisterNumberOk returns a tuple with the RegulatorRegisterNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatorRegisterNumber

`func (o *InstitutionDetails) SetRegulatorRegisterNumber(v string)`

SetRegulatorRegisterNumber sets RegulatorRegisterNumber field to given value.

### HasRegulatorRegisterNumber

`func (o *InstitutionDetails) HasRegulatorRegisterNumber() bool`

HasRegulatorRegisterNumber returns a boolean if a field has been set.

### GetDocumentVerificationStatus

`func (o *InstitutionDetails) GetDocumentVerificationStatus() IdentityStatus`

GetDocumentVerificationStatus returns the DocumentVerificationStatus field if non-nil, zero value otherwise.

### GetDocumentVerificationStatusOk

`func (o *InstitutionDetails) GetDocumentVerificationStatusOk() (*IdentityStatus, bool)`

GetDocumentVerificationStatusOk returns a tuple with the DocumentVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentVerificationStatus

`func (o *InstitutionDetails) SetDocumentVerificationStatus(v IdentityStatus)`

SetDocumentVerificationStatus sets DocumentVerificationStatus field to given value.

### HasDocumentVerificationStatus

`func (o *InstitutionDetails) HasDocumentVerificationStatus() bool`

HasDocumentVerificationStatus returns a boolean if a field has been set.

### GetAdditionalScreeningStatus

`func (o *InstitutionDetails) GetAdditionalScreeningStatus() IdentityStatus`

GetAdditionalScreeningStatus returns the AdditionalScreeningStatus field if non-nil, zero value otherwise.

### GetAdditionalScreeningStatusOk

`func (o *InstitutionDetails) GetAdditionalScreeningStatusOk() (*IdentityStatus, bool)`

GetAdditionalScreeningStatusOk returns a tuple with the AdditionalScreeningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalScreeningStatus

`func (o *InstitutionDetails) SetAdditionalScreeningStatus(v IdentityStatus)`

SetAdditionalScreeningStatus sets AdditionalScreeningStatus field to given value.

### HasAdditionalScreeningStatus

`func (o *InstitutionDetails) HasAdditionalScreeningStatus() bool`

HasAdditionalScreeningStatus returns a boolean if a field has been set.

### GetDoingBusinessAs

`func (o *InstitutionDetails) GetDoingBusinessAs() string`

GetDoingBusinessAs returns the DoingBusinessAs field if non-nil, zero value otherwise.

### GetDoingBusinessAsOk

`func (o *InstitutionDetails) GetDoingBusinessAsOk() (*string, bool)`

GetDoingBusinessAsOk returns a tuple with the DoingBusinessAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoingBusinessAs

`func (o *InstitutionDetails) SetDoingBusinessAs(v string)`

SetDoingBusinessAs sets DoingBusinessAs field to given value.

### HasDoingBusinessAs

`func (o *InstitutionDetails) HasDoingBusinessAs() bool`

HasDoingBusinessAs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


