# Obligation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Direction** | [**ObligationDirection**](ObligationDirection.md) |  | 
**Asset** | **string** | Asset which is obliged. | 
**Amount** | **string** | Amount of the asset which is obliged. | 

## Methods

### NewObligation

`func NewObligation(id string, direction ObligationDirection, asset string, amount string, ) *Obligation`

NewObligation instantiates a new Obligation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObligationWithDefaults

`func NewObligationWithDefaults() *Obligation`

NewObligationWithDefaults instantiates a new Obligation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Obligation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Obligation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Obligation) SetId(v string)`

SetId sets Id field to given value.


### GetDirection

`func (o *Obligation) GetDirection() ObligationDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Obligation) GetDirectionOk() (*ObligationDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Obligation) SetDirection(v ObligationDirection)`

SetDirection sets Direction field to given value.


### GetAsset

`func (o *Obligation) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *Obligation) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *Obligation) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetAmount

`func (o *Obligation) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Obligation) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Obligation) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


