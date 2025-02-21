# UpdateCryptoDepositRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityId** | **string** | The Identity (&#x60;identity_id&#x60;) of the end user updating the deposit. | 
**OriginatorAddressInfo** | [**AddressInfo**](AddressInfo.md) |  | 

## Methods

### NewUpdateCryptoDepositRequest

`func NewUpdateCryptoDepositRequest(identityId string, originatorAddressInfo AddressInfo, ) *UpdateCryptoDepositRequest`

NewUpdateCryptoDepositRequest instantiates a new UpdateCryptoDepositRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCryptoDepositRequestWithDefaults

`func NewUpdateCryptoDepositRequestWithDefaults() *UpdateCryptoDepositRequest`

NewUpdateCryptoDepositRequestWithDefaults instantiates a new UpdateCryptoDepositRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityId

`func (o *UpdateCryptoDepositRequest) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *UpdateCryptoDepositRequest) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *UpdateCryptoDepositRequest) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetOriginatorAddressInfo

`func (o *UpdateCryptoDepositRequest) GetOriginatorAddressInfo() AddressInfo`

GetOriginatorAddressInfo returns the OriginatorAddressInfo field if non-nil, zero value otherwise.

### GetOriginatorAddressInfoOk

`func (o *UpdateCryptoDepositRequest) GetOriginatorAddressInfoOk() (*AddressInfo, bool)`

GetOriginatorAddressInfoOk returns a tuple with the OriginatorAddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorAddressInfo

`func (o *UpdateCryptoDepositRequest) SetOriginatorAddressInfo(v AddressInfo)`

SetOriginatorAddressInfo sets OriginatorAddressInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


