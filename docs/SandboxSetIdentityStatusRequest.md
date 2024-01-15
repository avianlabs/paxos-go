# SandboxSetIdentityStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdVerificationStatus** | Pointer to [**IdentityStatus**](IdentityStatus.md) |  | [optional] 
**SanctionsVerificationStatus** | Pointer to [**IdentityStatus**](IdentityStatus.md) |  | [optional] 
**UserDisabled** | Pointer to [**SetDisable**](SetDisable.md) |  | [optional] 
**AdminDisabled** | Pointer to [**SetDisable**](SetDisable.md) |  | [optional] 
**DocumentVerificationStatus** | Pointer to [**IdentityStatus**](IdentityStatus.md) |  | [optional] 
**AdditionalScreeningStatus** | Pointer to [**IdentityStatus**](IdentityStatus.md) |  | [optional] 

## Methods

### NewSandboxSetIdentityStatusRequest

`func NewSandboxSetIdentityStatusRequest() *SandboxSetIdentityStatusRequest`

NewSandboxSetIdentityStatusRequest instantiates a new SandboxSetIdentityStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxSetIdentityStatusRequestWithDefaults

`func NewSandboxSetIdentityStatusRequestWithDefaults() *SandboxSetIdentityStatusRequest`

NewSandboxSetIdentityStatusRequestWithDefaults instantiates a new SandboxSetIdentityStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdVerificationStatus

`func (o *SandboxSetIdentityStatusRequest) GetIdVerificationStatus() IdentityStatus`

GetIdVerificationStatus returns the IdVerificationStatus field if non-nil, zero value otherwise.

### GetIdVerificationStatusOk

`func (o *SandboxSetIdentityStatusRequest) GetIdVerificationStatusOk() (*IdentityStatus, bool)`

GetIdVerificationStatusOk returns a tuple with the IdVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdVerificationStatus

`func (o *SandboxSetIdentityStatusRequest) SetIdVerificationStatus(v IdentityStatus)`

SetIdVerificationStatus sets IdVerificationStatus field to given value.

### HasIdVerificationStatus

`func (o *SandboxSetIdentityStatusRequest) HasIdVerificationStatus() bool`

HasIdVerificationStatus returns a boolean if a field has been set.

### GetSanctionsVerificationStatus

`func (o *SandboxSetIdentityStatusRequest) GetSanctionsVerificationStatus() IdentityStatus`

GetSanctionsVerificationStatus returns the SanctionsVerificationStatus field if non-nil, zero value otherwise.

### GetSanctionsVerificationStatusOk

`func (o *SandboxSetIdentityStatusRequest) GetSanctionsVerificationStatusOk() (*IdentityStatus, bool)`

GetSanctionsVerificationStatusOk returns a tuple with the SanctionsVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctionsVerificationStatus

`func (o *SandboxSetIdentityStatusRequest) SetSanctionsVerificationStatus(v IdentityStatus)`

SetSanctionsVerificationStatus sets SanctionsVerificationStatus field to given value.

### HasSanctionsVerificationStatus

`func (o *SandboxSetIdentityStatusRequest) HasSanctionsVerificationStatus() bool`

HasSanctionsVerificationStatus returns a boolean if a field has been set.

### GetUserDisabled

`func (o *SandboxSetIdentityStatusRequest) GetUserDisabled() SetDisable`

GetUserDisabled returns the UserDisabled field if non-nil, zero value otherwise.

### GetUserDisabledOk

`func (o *SandboxSetIdentityStatusRequest) GetUserDisabledOk() (*SetDisable, bool)`

GetUserDisabledOk returns a tuple with the UserDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisabled

`func (o *SandboxSetIdentityStatusRequest) SetUserDisabled(v SetDisable)`

SetUserDisabled sets UserDisabled field to given value.

### HasUserDisabled

`func (o *SandboxSetIdentityStatusRequest) HasUserDisabled() bool`

HasUserDisabled returns a boolean if a field has been set.

### GetAdminDisabled

`func (o *SandboxSetIdentityStatusRequest) GetAdminDisabled() SetDisable`

GetAdminDisabled returns the AdminDisabled field if non-nil, zero value otherwise.

### GetAdminDisabledOk

`func (o *SandboxSetIdentityStatusRequest) GetAdminDisabledOk() (*SetDisable, bool)`

GetAdminDisabledOk returns a tuple with the AdminDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminDisabled

`func (o *SandboxSetIdentityStatusRequest) SetAdminDisabled(v SetDisable)`

SetAdminDisabled sets AdminDisabled field to given value.

### HasAdminDisabled

`func (o *SandboxSetIdentityStatusRequest) HasAdminDisabled() bool`

HasAdminDisabled returns a boolean if a field has been set.

### GetDocumentVerificationStatus

`func (o *SandboxSetIdentityStatusRequest) GetDocumentVerificationStatus() IdentityStatus`

GetDocumentVerificationStatus returns the DocumentVerificationStatus field if non-nil, zero value otherwise.

### GetDocumentVerificationStatusOk

`func (o *SandboxSetIdentityStatusRequest) GetDocumentVerificationStatusOk() (*IdentityStatus, bool)`

GetDocumentVerificationStatusOk returns a tuple with the DocumentVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentVerificationStatus

`func (o *SandboxSetIdentityStatusRequest) SetDocumentVerificationStatus(v IdentityStatus)`

SetDocumentVerificationStatus sets DocumentVerificationStatus field to given value.

### HasDocumentVerificationStatus

`func (o *SandboxSetIdentityStatusRequest) HasDocumentVerificationStatus() bool`

HasDocumentVerificationStatus returns a boolean if a field has been set.

### GetAdditionalScreeningStatus

`func (o *SandboxSetIdentityStatusRequest) GetAdditionalScreeningStatus() IdentityStatus`

GetAdditionalScreeningStatus returns the AdditionalScreeningStatus field if non-nil, zero value otherwise.

### GetAdditionalScreeningStatusOk

`func (o *SandboxSetIdentityStatusRequest) GetAdditionalScreeningStatusOk() (*IdentityStatus, bool)`

GetAdditionalScreeningStatusOk returns a tuple with the AdditionalScreeningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalScreeningStatus

`func (o *SandboxSetIdentityStatusRequest) SetAdditionalScreeningStatus(v IdentityStatus)`

SetAdditionalScreeningStatus sets AdditionalScreeningStatus field to given value.

### HasAdditionalScreeningStatus

`func (o *SandboxSetIdentityStatusRequest) HasAdditionalScreeningStatus() bool`

HasAdditionalScreeningStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


