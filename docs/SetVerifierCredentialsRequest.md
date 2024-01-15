# SetVerifierCredentialsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerifierType** | [**IdentityprotoVerifierType**](IdentityprotoVerifierType.md) |  | 
**AuthToken** | Pointer to **string** |  | [optional] 
**AuthSecret** | Pointer to **string** |  | [optional] 

## Methods

### NewSetVerifierCredentialsRequest

`func NewSetVerifierCredentialsRequest(verifierType IdentityprotoVerifierType, ) *SetVerifierCredentialsRequest`

NewSetVerifierCredentialsRequest instantiates a new SetVerifierCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetVerifierCredentialsRequestWithDefaults

`func NewSetVerifierCredentialsRequestWithDefaults() *SetVerifierCredentialsRequest`

NewSetVerifierCredentialsRequestWithDefaults instantiates a new SetVerifierCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerifierType

`func (o *SetVerifierCredentialsRequest) GetVerifierType() IdentityprotoVerifierType`

GetVerifierType returns the VerifierType field if non-nil, zero value otherwise.

### GetVerifierTypeOk

`func (o *SetVerifierCredentialsRequest) GetVerifierTypeOk() (*IdentityprotoVerifierType, bool)`

GetVerifierTypeOk returns a tuple with the VerifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifierType

`func (o *SetVerifierCredentialsRequest) SetVerifierType(v IdentityprotoVerifierType)`

SetVerifierType sets VerifierType field to given value.


### GetAuthToken

`func (o *SetVerifierCredentialsRequest) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *SetVerifierCredentialsRequest) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *SetVerifierCredentialsRequest) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *SetVerifierCredentialsRequest) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetAuthSecret

`func (o *SetVerifierCredentialsRequest) GetAuthSecret() string`

GetAuthSecret returns the AuthSecret field if non-nil, zero value otherwise.

### GetAuthSecretOk

`func (o *SetVerifierCredentialsRequest) GetAuthSecretOk() (*string, bool)`

GetAuthSecretOk returns a tuple with the AuthSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecret

`func (o *SetVerifierCredentialsRequest) SetAuthSecret(v string)`

SetAuthSecret sets AuthSecret field to given value.

### HasAuthSecret

`func (o *SetVerifierCredentialsRequest) HasAuthSecret() bool`

HasAuthSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


