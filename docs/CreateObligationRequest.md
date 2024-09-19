# CreateObligationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**ObligationDirection**](ObligationDirection.md) |  | [optional] 
**Asset** | Pointer to **string** | Asset which is obliged. | [optional] 
**Amount** | Pointer to **string** | Amount of the asset which is obliged. | [optional] 

## Methods

### NewCreateObligationRequest

`func NewCreateObligationRequest() *CreateObligationRequest`

NewCreateObligationRequest instantiates a new CreateObligationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateObligationRequestWithDefaults

`func NewCreateObligationRequestWithDefaults() *CreateObligationRequest`

NewCreateObligationRequestWithDefaults instantiates a new CreateObligationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *CreateObligationRequest) GetDirection() ObligationDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CreateObligationRequest) GetDirectionOk() (*ObligationDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CreateObligationRequest) SetDirection(v ObligationDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CreateObligationRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetAsset

`func (o *CreateObligationRequest) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *CreateObligationRequest) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *CreateObligationRequest) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *CreateObligationRequest) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAmount

`func (o *CreateObligationRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateObligationRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateObligationRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CreateObligationRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


