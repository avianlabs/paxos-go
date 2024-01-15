# TaxLot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The tax lot ID. | [optional] 
**ProfileId** | Pointer to **string** | Profile ID. | [optional] 
**AccountId** | Pointer to **string** | Account ID. | [optional] 
**AcquiredAt** | Pointer to **time.Time** | The timestamp at which the asset was acquired. | [optional] 
**Cryptocurrency** | Pointer to **string** | The currency for numeric fields in Crypto. | [optional] 
**FiatCurrency** | Pointer to **string** | The currency for numeric fields in Fiat. | [optional] 
**CostBasis** | Pointer to **string** | The cost basis. | [optional] 
**CostBasisPerCoin** | Pointer to **string** | The cost basis per crypto coin. | [optional] 
**Quantity** | Pointer to **string** | The quantity of total crypto coins. | [optional] 
**TransactionType** | Pointer to [**TransactionType**](TransactionType.md) |  | [optional] 
**TransactionId** | Pointer to **string** | Transaction id. | [optional] 
**CapitalGain** | Pointer to **string** | The capital gain for the tax lot, estimated with current crypto price for open tax lots. | [optional] 
**CapitalGainType** | Pointer to [**CapitalGainType**](CapitalGainType.md) |  | [optional] 
**PercentageCapitalGain** | Pointer to **string** | The percentage of capital gain over cost basis. | [optional] 
**CurrentMarketValue** | Pointer to **string** | The current market value calculated by current Crypto price. | [optional] 
**DaysToLongTerm** | Pointer to **string** | The number of days left until capital gain type is LONG_TERM. Return 999 if acquisition time is unknown. | [optional] 
**Status** | Pointer to [**TaxLotStatus**](TaxLotStatus.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp at which tax lot is created. | [optional] 
**ClosedAt** | Pointer to **time.Time** | The timestamp at which tax lot is closed. | [optional] 

## Methods

### NewTaxLot

`func NewTaxLot() *TaxLot`

NewTaxLot instantiates a new TaxLot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxLotWithDefaults

`func NewTaxLotWithDefaults() *TaxLot`

NewTaxLotWithDefaults instantiates a new TaxLot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxLot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxLot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxLot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxLot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProfileId

`func (o *TaxLot) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *TaxLot) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *TaxLot) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *TaxLot) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetAccountId

`func (o *TaxLot) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TaxLot) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TaxLot) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TaxLot) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAcquiredAt

`func (o *TaxLot) GetAcquiredAt() time.Time`

GetAcquiredAt returns the AcquiredAt field if non-nil, zero value otherwise.

### GetAcquiredAtOk

`func (o *TaxLot) GetAcquiredAtOk() (*time.Time, bool)`

GetAcquiredAtOk returns a tuple with the AcquiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquiredAt

`func (o *TaxLot) SetAcquiredAt(v time.Time)`

SetAcquiredAt sets AcquiredAt field to given value.

### HasAcquiredAt

`func (o *TaxLot) HasAcquiredAt() bool`

HasAcquiredAt returns a boolean if a field has been set.

### GetCryptocurrency

`func (o *TaxLot) GetCryptocurrency() string`

GetCryptocurrency returns the Cryptocurrency field if non-nil, zero value otherwise.

### GetCryptocurrencyOk

`func (o *TaxLot) GetCryptocurrencyOk() (*string, bool)`

GetCryptocurrencyOk returns a tuple with the Cryptocurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptocurrency

`func (o *TaxLot) SetCryptocurrency(v string)`

SetCryptocurrency sets Cryptocurrency field to given value.

### HasCryptocurrency

`func (o *TaxLot) HasCryptocurrency() bool`

HasCryptocurrency returns a boolean if a field has been set.

### GetFiatCurrency

`func (o *TaxLot) GetFiatCurrency() string`

GetFiatCurrency returns the FiatCurrency field if non-nil, zero value otherwise.

### GetFiatCurrencyOk

`func (o *TaxLot) GetFiatCurrencyOk() (*string, bool)`

GetFiatCurrencyOk returns a tuple with the FiatCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatCurrency

`func (o *TaxLot) SetFiatCurrency(v string)`

SetFiatCurrency sets FiatCurrency field to given value.

### HasFiatCurrency

`func (o *TaxLot) HasFiatCurrency() bool`

HasFiatCurrency returns a boolean if a field has been set.

### GetCostBasis

`func (o *TaxLot) GetCostBasis() string`

GetCostBasis returns the CostBasis field if non-nil, zero value otherwise.

### GetCostBasisOk

`func (o *TaxLot) GetCostBasisOk() (*string, bool)`

GetCostBasisOk returns a tuple with the CostBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostBasis

`func (o *TaxLot) SetCostBasis(v string)`

SetCostBasis sets CostBasis field to given value.

### HasCostBasis

`func (o *TaxLot) HasCostBasis() bool`

HasCostBasis returns a boolean if a field has been set.

### GetCostBasisPerCoin

`func (o *TaxLot) GetCostBasisPerCoin() string`

GetCostBasisPerCoin returns the CostBasisPerCoin field if non-nil, zero value otherwise.

### GetCostBasisPerCoinOk

`func (o *TaxLot) GetCostBasisPerCoinOk() (*string, bool)`

GetCostBasisPerCoinOk returns a tuple with the CostBasisPerCoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostBasisPerCoin

`func (o *TaxLot) SetCostBasisPerCoin(v string)`

SetCostBasisPerCoin sets CostBasisPerCoin field to given value.

### HasCostBasisPerCoin

`func (o *TaxLot) HasCostBasisPerCoin() bool`

HasCostBasisPerCoin returns a boolean if a field has been set.

### GetQuantity

`func (o *TaxLot) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *TaxLot) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *TaxLot) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *TaxLot) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTransactionType

`func (o *TaxLot) GetTransactionType() TransactionType`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *TaxLot) GetTransactionTypeOk() (*TransactionType, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *TaxLot) SetTransactionType(v TransactionType)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *TaxLot) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetTransactionId

`func (o *TaxLot) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TaxLot) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TaxLot) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *TaxLot) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetCapitalGain

`func (o *TaxLot) GetCapitalGain() string`

GetCapitalGain returns the CapitalGain field if non-nil, zero value otherwise.

### GetCapitalGainOk

`func (o *TaxLot) GetCapitalGainOk() (*string, bool)`

GetCapitalGainOk returns a tuple with the CapitalGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalGain

`func (o *TaxLot) SetCapitalGain(v string)`

SetCapitalGain sets CapitalGain field to given value.

### HasCapitalGain

`func (o *TaxLot) HasCapitalGain() bool`

HasCapitalGain returns a boolean if a field has been set.

### GetCapitalGainType

`func (o *TaxLot) GetCapitalGainType() CapitalGainType`

GetCapitalGainType returns the CapitalGainType field if non-nil, zero value otherwise.

### GetCapitalGainTypeOk

`func (o *TaxLot) GetCapitalGainTypeOk() (*CapitalGainType, bool)`

GetCapitalGainTypeOk returns a tuple with the CapitalGainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalGainType

`func (o *TaxLot) SetCapitalGainType(v CapitalGainType)`

SetCapitalGainType sets CapitalGainType field to given value.

### HasCapitalGainType

`func (o *TaxLot) HasCapitalGainType() bool`

HasCapitalGainType returns a boolean if a field has been set.

### GetPercentageCapitalGain

`func (o *TaxLot) GetPercentageCapitalGain() string`

GetPercentageCapitalGain returns the PercentageCapitalGain field if non-nil, zero value otherwise.

### GetPercentageCapitalGainOk

`func (o *TaxLot) GetPercentageCapitalGainOk() (*string, bool)`

GetPercentageCapitalGainOk returns a tuple with the PercentageCapitalGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageCapitalGain

`func (o *TaxLot) SetPercentageCapitalGain(v string)`

SetPercentageCapitalGain sets PercentageCapitalGain field to given value.

### HasPercentageCapitalGain

`func (o *TaxLot) HasPercentageCapitalGain() bool`

HasPercentageCapitalGain returns a boolean if a field has been set.

### GetCurrentMarketValue

`func (o *TaxLot) GetCurrentMarketValue() string`

GetCurrentMarketValue returns the CurrentMarketValue field if non-nil, zero value otherwise.

### GetCurrentMarketValueOk

`func (o *TaxLot) GetCurrentMarketValueOk() (*string, bool)`

GetCurrentMarketValueOk returns a tuple with the CurrentMarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMarketValue

`func (o *TaxLot) SetCurrentMarketValue(v string)`

SetCurrentMarketValue sets CurrentMarketValue field to given value.

### HasCurrentMarketValue

`func (o *TaxLot) HasCurrentMarketValue() bool`

HasCurrentMarketValue returns a boolean if a field has been set.

### GetDaysToLongTerm

`func (o *TaxLot) GetDaysToLongTerm() string`

GetDaysToLongTerm returns the DaysToLongTerm field if non-nil, zero value otherwise.

### GetDaysToLongTermOk

`func (o *TaxLot) GetDaysToLongTermOk() (*string, bool)`

GetDaysToLongTermOk returns a tuple with the DaysToLongTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToLongTerm

`func (o *TaxLot) SetDaysToLongTerm(v string)`

SetDaysToLongTerm sets DaysToLongTerm field to given value.

### HasDaysToLongTerm

`func (o *TaxLot) HasDaysToLongTerm() bool`

HasDaysToLongTerm returns a boolean if a field has been set.

### GetStatus

`func (o *TaxLot) GetStatus() TaxLotStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaxLot) GetStatusOk() (*TaxLotStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaxLot) SetStatus(v TaxLotStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaxLot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaxLot) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaxLot) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaxLot) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaxLot) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetClosedAt

`func (o *TaxLot) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *TaxLot) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *TaxLot) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *TaxLot) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


