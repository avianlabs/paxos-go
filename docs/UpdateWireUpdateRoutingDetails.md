# UpdateWireUpdateRoutingDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankName** | Pointer to **string** | The name of the bank. | [optional] 
**BankAddress** | Pointer to [**MailingAddress**](MailingAddress.md) |  | [optional] 

## Methods

### NewUpdateWireUpdateRoutingDetails

`func NewUpdateWireUpdateRoutingDetails() *UpdateWireUpdateRoutingDetails`

NewUpdateWireUpdateRoutingDetails instantiates a new UpdateWireUpdateRoutingDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWireUpdateRoutingDetailsWithDefaults

`func NewUpdateWireUpdateRoutingDetailsWithDefaults() *UpdateWireUpdateRoutingDetails`

NewUpdateWireUpdateRoutingDetailsWithDefaults instantiates a new UpdateWireUpdateRoutingDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankName

`func (o *UpdateWireUpdateRoutingDetails) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *UpdateWireUpdateRoutingDetails) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *UpdateWireUpdateRoutingDetails) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *UpdateWireUpdateRoutingDetails) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetBankAddress

`func (o *UpdateWireUpdateRoutingDetails) GetBankAddress() MailingAddress`

GetBankAddress returns the BankAddress field if non-nil, zero value otherwise.

### GetBankAddressOk

`func (o *UpdateWireUpdateRoutingDetails) GetBankAddressOk() (*MailingAddress, bool)`

GetBankAddressOk returns a tuple with the BankAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAddress

`func (o *UpdateWireUpdateRoutingDetails) SetBankAddress(v MailingAddress)`

SetBankAddress sets BankAddress field to given value.

### HasBankAddress

`func (o *UpdateWireUpdateRoutingDetails) HasBankAddress() bool`

HasBankAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


