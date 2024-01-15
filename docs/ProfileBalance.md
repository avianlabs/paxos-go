# ProfileBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | [**Asset**](Asset.md) |  | 
**Available** | **string** | The available asset balance for new debit requests. | 
**Trading** | **string** | The asset amount committed to pending orders. | 

## Methods

### NewProfileBalance

`func NewProfileBalance(asset Asset, available string, trading string, ) *ProfileBalance`

NewProfileBalance instantiates a new ProfileBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileBalanceWithDefaults

`func NewProfileBalanceWithDefaults() *ProfileBalance`

NewProfileBalanceWithDefaults instantiates a new ProfileBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *ProfileBalance) GetAsset() Asset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *ProfileBalance) GetAssetOk() (*Asset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *ProfileBalance) SetAsset(v Asset)`

SetAsset sets Asset field to given value.


### GetAvailable

`func (o *ProfileBalance) GetAvailable() string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ProfileBalance) GetAvailableOk() (*string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ProfileBalance) SetAvailable(v string)`

SetAvailable sets Available field to given value.


### GetTrading

`func (o *ProfileBalance) GetTrading() string`

GetTrading returns the Trading field if non-nil, zero value otherwise.

### GetTradingOk

`func (o *ProfileBalance) GetTradingOk() (*string, bool)`

GetTradingOk returns a tuple with the Trading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrading

`func (o *ProfileBalance) SetTrading(v string)`

SetTrading sets Trading field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


