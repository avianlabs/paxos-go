# InitiateSandboxFiatDepositRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | The amount to deposit. | 
**Asset** | **string** | The asset to deposit. Current supported asset: \&quot;USD\&quot;. | 
**MemoId** | **string** | The string that the client must provide in the memo field on their deposit to credit their Paxos platform balance. | 
**FiatNetworkInstructions** | [**FiatNetworkInstructions**](FiatNetworkInstructions.md) |  | 
**FiatAccountOwner** | Pointer to [**FiatAccountOwner**](FiatAccountOwner.md) |  | [optional] 

## Methods

### NewInitiateSandboxFiatDepositRequest

`func NewInitiateSandboxFiatDepositRequest(amount string, asset string, memoId string, fiatNetworkInstructions FiatNetworkInstructions, ) *InitiateSandboxFiatDepositRequest`

NewInitiateSandboxFiatDepositRequest instantiates a new InitiateSandboxFiatDepositRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiateSandboxFiatDepositRequestWithDefaults

`func NewInitiateSandboxFiatDepositRequestWithDefaults() *InitiateSandboxFiatDepositRequest`

NewInitiateSandboxFiatDepositRequestWithDefaults instantiates a new InitiateSandboxFiatDepositRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InitiateSandboxFiatDepositRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InitiateSandboxFiatDepositRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InitiateSandboxFiatDepositRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAsset

`func (o *InitiateSandboxFiatDepositRequest) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *InitiateSandboxFiatDepositRequest) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *InitiateSandboxFiatDepositRequest) SetAsset(v string)`

SetAsset sets Asset field to given value.


### GetMemoId

`func (o *InitiateSandboxFiatDepositRequest) GetMemoId() string`

GetMemoId returns the MemoId field if non-nil, zero value otherwise.

### GetMemoIdOk

`func (o *InitiateSandboxFiatDepositRequest) GetMemoIdOk() (*string, bool)`

GetMemoIdOk returns a tuple with the MemoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoId

`func (o *InitiateSandboxFiatDepositRequest) SetMemoId(v string)`

SetMemoId sets MemoId field to given value.


### GetFiatNetworkInstructions

`func (o *InitiateSandboxFiatDepositRequest) GetFiatNetworkInstructions() FiatNetworkInstructions`

GetFiatNetworkInstructions returns the FiatNetworkInstructions field if non-nil, zero value otherwise.

### GetFiatNetworkInstructionsOk

`func (o *InitiateSandboxFiatDepositRequest) GetFiatNetworkInstructionsOk() (*FiatNetworkInstructions, bool)`

GetFiatNetworkInstructionsOk returns a tuple with the FiatNetworkInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatNetworkInstructions

`func (o *InitiateSandboxFiatDepositRequest) SetFiatNetworkInstructions(v FiatNetworkInstructions)`

SetFiatNetworkInstructions sets FiatNetworkInstructions field to given value.


### GetFiatAccountOwner

`func (o *InitiateSandboxFiatDepositRequest) GetFiatAccountOwner() FiatAccountOwner`

GetFiatAccountOwner returns the FiatAccountOwner field if non-nil, zero value otherwise.

### GetFiatAccountOwnerOk

`func (o *InitiateSandboxFiatDepositRequest) GetFiatAccountOwnerOk() (*FiatAccountOwner, bool)`

GetFiatAccountOwnerOk returns a tuple with the FiatAccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatAccountOwner

`func (o *InitiateSandboxFiatDepositRequest) SetFiatAccountOwner(v FiatAccountOwner)`

SetFiatAccountOwner sets FiatAccountOwner field to given value.

### HasFiatAccountOwner

`func (o *InitiateSandboxFiatDepositRequest) HasFiatAccountOwner() bool`

HasFiatAccountOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


