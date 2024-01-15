# TaxDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxPayerId** | Pointer to **string** |  | [optional] 
**TaxPayerCountry** | Pointer to **string** |  | [optional] 
**TinVerificationStatus** | Pointer to [**TINVerificationStatus**](TINVerificationStatus.md) |  | [optional] 

## Methods

### NewTaxDetail

`func NewTaxDetail() *TaxDetail`

NewTaxDetail instantiates a new TaxDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxDetailWithDefaults

`func NewTaxDetailWithDefaults() *TaxDetail`

NewTaxDetailWithDefaults instantiates a new TaxDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxPayerId

`func (o *TaxDetail) GetTaxPayerId() string`

GetTaxPayerId returns the TaxPayerId field if non-nil, zero value otherwise.

### GetTaxPayerIdOk

`func (o *TaxDetail) GetTaxPayerIdOk() (*string, bool)`

GetTaxPayerIdOk returns a tuple with the TaxPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPayerId

`func (o *TaxDetail) SetTaxPayerId(v string)`

SetTaxPayerId sets TaxPayerId field to given value.

### HasTaxPayerId

`func (o *TaxDetail) HasTaxPayerId() bool`

HasTaxPayerId returns a boolean if a field has been set.

### GetTaxPayerCountry

`func (o *TaxDetail) GetTaxPayerCountry() string`

GetTaxPayerCountry returns the TaxPayerCountry field if non-nil, zero value otherwise.

### GetTaxPayerCountryOk

`func (o *TaxDetail) GetTaxPayerCountryOk() (*string, bool)`

GetTaxPayerCountryOk returns a tuple with the TaxPayerCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPayerCountry

`func (o *TaxDetail) SetTaxPayerCountry(v string)`

SetTaxPayerCountry sets TaxPayerCountry field to given value.

### HasTaxPayerCountry

`func (o *TaxDetail) HasTaxPayerCountry() bool`

HasTaxPayerCountry returns a boolean if a field has been set.

### GetTinVerificationStatus

`func (o *TaxDetail) GetTinVerificationStatus() TINVerificationStatus`

GetTinVerificationStatus returns the TinVerificationStatus field if non-nil, zero value otherwise.

### GetTinVerificationStatusOk

`func (o *TaxDetail) GetTinVerificationStatusOk() (*TINVerificationStatus, bool)`

GetTinVerificationStatusOk returns a tuple with the TinVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTinVerificationStatus

`func (o *TaxDetail) SetTinVerificationStatus(v TINVerificationStatus)`

SetTinVerificationStatus sets TinVerificationStatus field to given value.

### HasTinVerificationStatus

`func (o *TaxDetail) HasTinVerificationStatus() bool`

HasTinVerificationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


