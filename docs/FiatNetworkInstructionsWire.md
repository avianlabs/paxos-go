# FiatNetworkInstructionsWire

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | The FiatAccount owner&#39;s bank account number. | 
**FiatAccountOwnerAddress** | [**MailingAddress**](MailingAddress.md) |  | 
**RoutingDetails** | [**WireRoutingDetails**](WireRoutingDetails.md) |  | 
**IntermediaryRoutingDetails** | Pointer to [**WireRoutingDetails**](WireRoutingDetails.md) |  | [optional] 

## Methods

### NewFiatNetworkInstructionsWire

`func NewFiatNetworkInstructionsWire(accountNumber string, fiatAccountOwnerAddress MailingAddress, routingDetails WireRoutingDetails, ) *FiatNetworkInstructionsWire`

NewFiatNetworkInstructionsWire instantiates a new FiatNetworkInstructionsWire object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiatNetworkInstructionsWireWithDefaults

`func NewFiatNetworkInstructionsWireWithDefaults() *FiatNetworkInstructionsWire`

NewFiatNetworkInstructionsWireWithDefaults instantiates a new FiatNetworkInstructionsWire object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *FiatNetworkInstructionsWire) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *FiatNetworkInstructionsWire) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *FiatNetworkInstructionsWire) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetFiatAccountOwnerAddress

`func (o *FiatNetworkInstructionsWire) GetFiatAccountOwnerAddress() MailingAddress`

GetFiatAccountOwnerAddress returns the FiatAccountOwnerAddress field if non-nil, zero value otherwise.

### GetFiatAccountOwnerAddressOk

`func (o *FiatNetworkInstructionsWire) GetFiatAccountOwnerAddressOk() (*MailingAddress, bool)`

GetFiatAccountOwnerAddressOk returns a tuple with the FiatAccountOwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatAccountOwnerAddress

`func (o *FiatNetworkInstructionsWire) SetFiatAccountOwnerAddress(v MailingAddress)`

SetFiatAccountOwnerAddress sets FiatAccountOwnerAddress field to given value.


### GetRoutingDetails

`func (o *FiatNetworkInstructionsWire) GetRoutingDetails() WireRoutingDetails`

GetRoutingDetails returns the RoutingDetails field if non-nil, zero value otherwise.

### GetRoutingDetailsOk

`func (o *FiatNetworkInstructionsWire) GetRoutingDetailsOk() (*WireRoutingDetails, bool)`

GetRoutingDetailsOk returns a tuple with the RoutingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingDetails

`func (o *FiatNetworkInstructionsWire) SetRoutingDetails(v WireRoutingDetails)`

SetRoutingDetails sets RoutingDetails field to given value.


### GetIntermediaryRoutingDetails

`func (o *FiatNetworkInstructionsWire) GetIntermediaryRoutingDetails() WireRoutingDetails`

GetIntermediaryRoutingDetails returns the IntermediaryRoutingDetails field if non-nil, zero value otherwise.

### GetIntermediaryRoutingDetailsOk

`func (o *FiatNetworkInstructionsWire) GetIntermediaryRoutingDetailsOk() (*WireRoutingDetails, bool)`

GetIntermediaryRoutingDetailsOk returns a tuple with the IntermediaryRoutingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediaryRoutingDetails

`func (o *FiatNetworkInstructionsWire) SetIntermediaryRoutingDetails(v WireRoutingDetails)`

SetIntermediaryRoutingDetails sets IntermediaryRoutingDetails field to given value.

### HasIntermediaryRoutingDetails

`func (o *FiatNetworkInstructionsWire) HasIntermediaryRoutingDetails() bool`

HasIntermediaryRoutingDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


