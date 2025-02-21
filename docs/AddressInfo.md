# AddressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndividualInfo** | Pointer to [**IndividualInfo**](IndividualInfo.md) |  | [optional] 
**InstitutionInfo** | Pointer to [**InstitutionInfo**](InstitutionInfo.md) |  | [optional] 

## Methods

### NewAddressInfo

`func NewAddressInfo() *AddressInfo`

NewAddressInfo instantiates a new AddressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressInfoWithDefaults

`func NewAddressInfoWithDefaults() *AddressInfo`

NewAddressInfoWithDefaults instantiates a new AddressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndividualInfo

`func (o *AddressInfo) GetIndividualInfo() IndividualInfo`

GetIndividualInfo returns the IndividualInfo field if non-nil, zero value otherwise.

### GetIndividualInfoOk

`func (o *AddressInfo) GetIndividualInfoOk() (*IndividualInfo, bool)`

GetIndividualInfoOk returns a tuple with the IndividualInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualInfo

`func (o *AddressInfo) SetIndividualInfo(v IndividualInfo)`

SetIndividualInfo sets IndividualInfo field to given value.

### HasIndividualInfo

`func (o *AddressInfo) HasIndividualInfo() bool`

HasIndividualInfo returns a boolean if a field has been set.

### GetInstitutionInfo

`func (o *AddressInfo) GetInstitutionInfo() InstitutionInfo`

GetInstitutionInfo returns the InstitutionInfo field if non-nil, zero value otherwise.

### GetInstitutionInfoOk

`func (o *AddressInfo) GetInstitutionInfoOk() (*InstitutionInfo, bool)`

GetInstitutionInfoOk returns a tuple with the InstitutionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionInfo

`func (o *AddressInfo) SetInstitutionInfo(v InstitutionInfo)`

SetInstitutionInfo sets InstitutionInfo field to given value.

### HasInstitutionInfo

`func (o *AddressInfo) HasInstitutionInfo() bool`

HasInstitutionInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


