# WireRoutingDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoutingNumberType** | [**FiatWireAccountType**](FiatWireAccountType.md) |  | 
**RoutingNumber** | **string** | The routing number. | 
**BankName** | **string** | The name of the bank. | 
**BankAddress** | [**MailingAddress**](MailingAddress.md) |  | 

## Methods

### NewWireRoutingDetails

`func NewWireRoutingDetails(routingNumberType FiatWireAccountType, routingNumber string, bankName string, bankAddress MailingAddress, ) *WireRoutingDetails`

NewWireRoutingDetails instantiates a new WireRoutingDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireRoutingDetailsWithDefaults

`func NewWireRoutingDetailsWithDefaults() *WireRoutingDetails`

NewWireRoutingDetailsWithDefaults instantiates a new WireRoutingDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutingNumberType

`func (o *WireRoutingDetails) GetRoutingNumberType() FiatWireAccountType`

GetRoutingNumberType returns the RoutingNumberType field if non-nil, zero value otherwise.

### GetRoutingNumberTypeOk

`func (o *WireRoutingDetails) GetRoutingNumberTypeOk() (*FiatWireAccountType, bool)`

GetRoutingNumberTypeOk returns a tuple with the RoutingNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumberType

`func (o *WireRoutingDetails) SetRoutingNumberType(v FiatWireAccountType)`

SetRoutingNumberType sets RoutingNumberType field to given value.


### GetRoutingNumber

`func (o *WireRoutingDetails) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *WireRoutingDetails) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *WireRoutingDetails) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.


### GetBankName

`func (o *WireRoutingDetails) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *WireRoutingDetails) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *WireRoutingDetails) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetBankAddress

`func (o *WireRoutingDetails) GetBankAddress() MailingAddress`

GetBankAddress returns the BankAddress field if non-nil, zero value otherwise.

### GetBankAddressOk

`func (o *WireRoutingDetails) GetBankAddressOk() (*MailingAddress, bool)`

GetBankAddressOk returns a tuple with the BankAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAddress

`func (o *WireRoutingDetails) SetBankAddress(v MailingAddress)`

SetBankAddress sets BankAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


