# UpdateFiatAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FiatAccountOwner** | Pointer to [**FiatAccountOwner**](FiatAccountOwner.md) |  | [optional] 
**FiatNetworkInstructions** | Pointer to [**UpdateFiatAccountRequestUpdateFiatNetworkInstructions**](UpdateFiatAccountRequestUpdateFiatNetworkInstructions.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Optional client-specified metadata. Up to 6 key/value pairs may be provided. Each key and value must be less than or equal to 100 characters. | [optional] 

## Methods

### NewUpdateFiatAccountRequest

`func NewUpdateFiatAccountRequest() *UpdateFiatAccountRequest`

NewUpdateFiatAccountRequest instantiates a new UpdateFiatAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFiatAccountRequestWithDefaults

`func NewUpdateFiatAccountRequestWithDefaults() *UpdateFiatAccountRequest`

NewUpdateFiatAccountRequestWithDefaults instantiates a new UpdateFiatAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiatAccountOwner

`func (o *UpdateFiatAccountRequest) GetFiatAccountOwner() FiatAccountOwner`

GetFiatAccountOwner returns the FiatAccountOwner field if non-nil, zero value otherwise.

### GetFiatAccountOwnerOk

`func (o *UpdateFiatAccountRequest) GetFiatAccountOwnerOk() (*FiatAccountOwner, bool)`

GetFiatAccountOwnerOk returns a tuple with the FiatAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatAccountOwner

`func (o *UpdateFiatAccountRequest) SetFiatAccountOwner(v FiatAccountOwner)`

SetFiatAccountOwner sets FiatAccountOwner field to given value.

### HasFiatAccountOwner

`func (o *UpdateFiatAccountRequest) HasFiatAccountOwner() bool`

HasFiatAccountOwner returns a boolean if a field has been set.

### GetFiatNetworkInstructions

`func (o *UpdateFiatAccountRequest) GetFiatNetworkInstructions() UpdateFiatAccountRequestUpdateFiatNetworkInstructions`

GetFiatNetworkInstructions returns the FiatNetworkInstructions field if non-nil, zero value otherwise.

### GetFiatNetworkInstructionsOk

`func (o *UpdateFiatAccountRequest) GetFiatNetworkInstructionsOk() (*UpdateFiatAccountRequestUpdateFiatNetworkInstructions, bool)`

GetFiatNetworkInstructionsOk returns a tuple with the FiatNetworkInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatNetworkInstructions

`func (o *UpdateFiatAccountRequest) SetFiatNetworkInstructions(v UpdateFiatAccountRequestUpdateFiatNetworkInstructions)`

SetFiatNetworkInstructions sets FiatNetworkInstructions field to given value.

### HasFiatNetworkInstructions

`func (o *UpdateFiatAccountRequest) HasFiatNetworkInstructions() bool`

HasFiatNetworkInstructions returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateFiatAccountRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateFiatAccountRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateFiatAccountRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateFiatAccountRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


