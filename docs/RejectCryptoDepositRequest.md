# RejectCryptoDepositRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityId** | **string** | The Identity (&#x60;identity_id&#x60;) of the end user that requests to reject the deposit. | 

## Methods

### NewRejectCryptoDepositRequest

`func NewRejectCryptoDepositRequest(identityId string, ) *RejectCryptoDepositRequest`

NewRejectCryptoDepositRequest instantiates a new RejectCryptoDepositRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRejectCryptoDepositRequestWithDefaults

`func NewRejectCryptoDepositRequestWithDefaults() *RejectCryptoDepositRequest`

NewRejectCryptoDepositRequestWithDefaults instantiates a new RejectCryptoDepositRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityId

`func (o *RejectCryptoDepositRequest) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *RejectCryptoDepositRequest) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *RejectCryptoDepositRequest) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


