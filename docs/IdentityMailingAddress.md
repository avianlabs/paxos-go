# IdentityMailingAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | **string** |  | 
**Address1** | **string** |  | 
**Address2** | Pointer to **string** | To clear address2 (i.e. when updating an identity), set address2 to an empty string (\&quot;\&quot;). | [optional] 
**City** | **string** |  | 
**Province** | **string** |  | 
**ZipCode** | Pointer to **string** |  | [optional] 

## Methods

### NewIdentityMailingAddress

`func NewIdentityMailingAddress(country string, address1 string, city string, province string, ) *IdentityMailingAddress`

NewIdentityMailingAddress instantiates a new IdentityMailingAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityMailingAddressWithDefaults

`func NewIdentityMailingAddressWithDefaults() *IdentityMailingAddress`

NewIdentityMailingAddressWithDefaults instantiates a new IdentityMailingAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *IdentityMailingAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *IdentityMailingAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *IdentityMailingAddress) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetAddress1

`func (o *IdentityMailingAddress) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *IdentityMailingAddress) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *IdentityMailingAddress) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.


### GetAddress2

`func (o *IdentityMailingAddress) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *IdentityMailingAddress) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *IdentityMailingAddress) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *IdentityMailingAddress) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *IdentityMailingAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *IdentityMailingAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *IdentityMailingAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetProvince

`func (o *IdentityMailingAddress) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *IdentityMailingAddress) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *IdentityMailingAddress) SetProvince(v string)`

SetProvince sets Province field to given value.


### GetZipCode

`func (o *IdentityMailingAddress) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *IdentityMailingAddress) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *IdentityMailingAddress) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *IdentityMailingAddress) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


