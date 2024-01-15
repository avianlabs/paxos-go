# BookLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **string** | Price at level. | [optional] 
**Amount** | Pointer to **string** | Amount of orders at pricing level. | [optional] 

## Methods

### NewBookLevel

`func NewBookLevel() *BookLevel`

NewBookLevel instantiates a new BookLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookLevelWithDefaults

`func NewBookLevelWithDefaults() *BookLevel`

NewBookLevelWithDefaults instantiates a new BookLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *BookLevel) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BookLevel) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BookLevel) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BookLevel) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetAmount

`func (o *BookLevel) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BookLevel) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BookLevel) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BookLevel) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


