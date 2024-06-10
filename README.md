# Go API client for paxos

<p>Welcome to Paxos APIs. At Paxos, our mission is to enable the movement of any asset, any time, in a trustworthy way. These APIs serve that mission by making it easier than ever for you to directly integrate our product capabilities into your application, leveraging the speed, stability, and security of the Paxos platform.</p> <p>The documentation that follows gives you access to our Crypto Brokerage, Trading, and Exchange products. It includes APIs for market data, orders, and the held rate quote flow.</p> <p>To test in our sandbox environment, <a href=\"https://account.sandbox.paxos.com\" target=\"_blank\">sign up</a> for an account. For more information about Paxos and our APIs, visit <a href=\"https://www.paxos.com/\" target=\"_blank\">Paxos.com</a>.</p>

### Command
Generation is performed with the following command:
```sh
openapi-generator generate -i paxos-v2.openapi.json -g go --package-name paxos
```

OpenAPI spec is found [here](https://docs.paxos.com/api/v2).


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0
- Package version: 1.0.0
- Generator version: 7.6.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import paxos "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `paxos.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), paxos.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `paxos.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), paxos.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `paxos.ContextOperationServerIndices` and `paxos.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), paxos.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), paxos.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.paxos.com/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountMembersAPI* | [**AddAccountMembers**](docs/AccountMembersAPI.md#addaccountmembers) | **Post** /identity/account-members | Add Account Members
*AccountMembersAPI* | [**DeleteAccountMember**](docs/AccountMembersAPI.md#deleteaccountmember) | **Delete** /identity/account-members/{id} | Remove Account Member
*AccountsAPI* | [**CreateAccount**](docs/AccountsAPI.md#createaccount) | **Post** /identity/accounts | Create Account
*AccountsAPI* | [**GetAccount**](docs/AccountsAPI.md#getaccount) | **Get** /identity/accounts/{id} | 
*AccountsAPI* | [**ListAccounts**](docs/AccountsAPI.md#listaccounts) | **Get** /identity/accounts | List Accounts
*AccountsAPI* | [**UpdateAccount**](docs/AccountsAPI.md#updateaccount) | **Put** /identity/accounts | Update Account
*CryptoWithdrawalsAPI* | [**CreateCryptoWithdrawal**](docs/CryptoWithdrawalsAPI.md#createcryptowithdrawal) | **Post** /transfer/crypto-withdrawals | Create Crypto Withdrawal
*DepositAddressesAPI* | [**CreateDepositAddress**](docs/DepositAddressesAPI.md#createdepositaddress) | **Post** /transfer/deposit-addresses | Create Deposit Address
*DepositAddressesAPI* | [**ListDepositAddresses**](docs/DepositAddressesAPI.md#listdepositaddresses) | **Get** /transfer/deposit-addresses | List Deposit Addresses
*ExchangePublicAPI* | [**ListProfileExecutions**](docs/ExchangePublicAPI.md#listprofileexecutions) | **Get** /profiles/{profile_id}/executions | List Executions
*ExchangePublicAPI* | [**ListProfileOrders**](docs/ExchangePublicAPI.md#listprofileorders) | **Get** /profiles/{profile_id}/orders | List Orders
*FeesAPI* | [**CreateCryptoWithdrawalFee**](docs/FeesAPI.md#createcryptowithdrawalfee) | **Post** /transfer/fees/crypto-withdrawal | Create Crypto Withdrawal Fee
*FiatTransfersAPI* | [**CreateFiatAccount**](docs/FiatTransfersAPI.md#createfiataccount) | **Post** /transfer/fiat-accounts | Create Fiat Account
*FiatTransfersAPI* | [**CreateFiatDepositInstructions**](docs/FiatTransfersAPI.md#createfiatdepositinstructions) | **Post** /transfer/fiat-deposit-instructions | Create Fiat Deposit Instructions
*FiatTransfersAPI* | [**CreateFiatWithdrawal**](docs/FiatTransfersAPI.md#createfiatwithdrawal) | **Post** /transfer/fiat-withdrawals | Create Fiat Withdrawal
*FiatTransfersAPI* | [**DeleteFiatAccount**](docs/FiatTransfersAPI.md#deletefiataccount) | **Delete** /transfer/fiat-accounts/{id} | Delete Fiat Account
*FiatTransfersAPI* | [**GetFiatAccount**](docs/FiatTransfersAPI.md#getfiataccount) | **Get** /transfer/fiat-accounts/{id} | Get Fiat Account
*FiatTransfersAPI* | [**GetFiatDepositInstructions**](docs/FiatTransfersAPI.md#getfiatdepositinstructions) | **Get** /transfer/fiat-deposit-instructions/{id} | Get Fiat Deposit Instructions
*FiatTransfersAPI* | [**ListFiatAccounts**](docs/FiatTransfersAPI.md#listfiataccounts) | **Get** /transfer/fiat-accounts | List Fiat Accounts
*FiatTransfersAPI* | [**ListFiatDepositInstructions**](docs/FiatTransfersAPI.md#listfiatdepositinstructions) | **Get** /transfer/fiat-deposit-instructions | List Fiat Deposit Instructions
*FiatTransfersAPI* | [**UpdateFiatAccount**](docs/FiatTransfersAPI.md#updatefiataccount) | **Put** /transfer/fiat-accounts/{id} | Update Fiat Account
*IdentityAPI* | [**CreateIdentity**](docs/IdentityAPI.md#createidentity) | **Post** /identity/identities | Create Identity
*IdentityAPI* | [**GetIdentity**](docs/IdentityAPI.md#getidentity) | **Get** /identity/identities/{id} | Get Identity
*IdentityAPI* | [**ListIdentities**](docs/IdentityAPI.md#listidentities) | **Get** /identity/identities | List Identities
*IdentityAPI* | [**UpdateIdentity**](docs/IdentityAPI.md#updateidentity) | **Put** /identity/identities/{id} | Update Identity
*IdentityCredentialsAPI* | [**RetryIdVerification**](docs/IdentityCredentialsAPI.md#retryidverification) | **Post** /identity/identities/{id}/retry-id-verification | Retry Id Verification
*IdentityCredentialsAPI* | [**SetVerifierCredentials**](docs/IdentityCredentialsAPI.md#setverifiercredentials) | **Post** /identity/verifier-credentials | Set Verifier Credentials
*IdentityDocumentsAPI* | [**DocumentUpload**](docs/IdentityDocumentsAPI.md#documentupload) | **Put** /identity/identities/{identity_id}/documents | Document Upload
*IdentityDocumentsAPI* | [**ListIdentityDocuments**](docs/IdentityDocumentsAPI.md#listidentitydocuments) | **Get** /identity/identities/{identity_id}/documents | List Identity Documents
*InstitutionMembersAPI* | [**AddInstitutionMembers**](docs/InstitutionMembersAPI.md#addinstitutionmembers) | **Post** /identity/institution-members | Add Institution Members
*InstitutionMembersAPI* | [**DeleteInstitutionMember**](docs/InstitutionMembersAPI.md#deleteinstitutionmember) | **Delete** /identity/institution-members/{id} | Remove Institution Member
*InternalTransfersAPI* | [**CreateInternalTransfer**](docs/InternalTransfersAPI.md#createinternaltransfer) | **Post** /transfer/internal | Create Internal Transfer
*LimitsAPI* | [**ListTransferLimits**](docs/LimitsAPI.md#listtransferlimits) | **Get** /transfer/limits/utilizations | List Transfer Limits
*MarketDataAPI* | [**GetOrderBook**](docs/MarketDataAPI.md#getorderbook) | **Get** /markets/{market}/order-book | Get Order Book
*MarketDataAPI* | [**GetTicker**](docs/MarketDataAPI.md#getticker) | **Get** /markets/{market}/ticker | Get Ticker
*MarketDataAPI* | [**ListMarkets**](docs/MarketDataAPI.md#listmarkets) | **Get** /markets | List Markets
*MarketDataAPI* | [**ListRecentExecutions**](docs/MarketDataAPI.md#listrecentexecutions) | **Get** /markets/{market}/recent-executions | List Recent Executions
*OrdersAPI* | [**CancelOrder**](docs/OrdersAPI.md#cancelorder) | **Delete** /profiles/{profile_id}/orders/{id} | Cancel Order
*OrdersAPI* | [**CreateOrder**](docs/OrdersAPI.md#createorder) | **Post** /profiles/{profile_id}/orders | Create Order
*OrdersAPI* | [**GetOrder**](docs/OrdersAPI.md#getorder) | **Get** /profiles/{profile_id}/orders/{id} | Get Order
*OrdersAPI* | [**ListExecutions**](docs/OrdersAPI.md#listexecutions) | **Get** /executions | List Executions
*OrdersAPI* | [**ListOrders**](docs/OrdersAPI.md#listorders) | **Get** /orders | List Orders
*PricingAPI* | [**ListHistoricalPrices**](docs/PricingAPI.md#listhistoricalprices) | **Get** /markets/{market}/historical-prices | List Historical Prices
*PricingAPI* | [**ListPrices**](docs/PricingAPI.md#listprices) | **Get** /all-markets/prices | List Prices
*PricingAPI* | [**ListTickers**](docs/PricingAPI.md#listtickers) | **Get** /all-markets/ticker | List Tickers
*ProfilesAPI* | [**CreateProfile**](docs/ProfilesAPI.md#createprofile) | **Post** /profiles | Create Profile
*ProfilesAPI* | [**GetProfile**](docs/ProfilesAPI.md#getprofile) | **Get** /profiles/{profile_id} | Get Profile
*ProfilesAPI* | [**GetProfileBalance**](docs/ProfilesAPI.md#getprofilebalance) | **Get** /profiles/{profile_id}/balances/{asset} | Get Profile Balance
*ProfilesAPI* | [**ListProfileBalances**](docs/ProfilesAPI.md#listprofilebalances) | **Get** /profiles/{profile_id}/balances | List Profile Balances
*ProfilesAPI* | [**ListProfiles**](docs/ProfilesAPI.md#listprofiles) | **Get** /profiles | List Profiles
*QuoteExecutionsAPI* | [**CreateQuoteExecution**](docs/QuoteExecutionsAPI.md#createquoteexecution) | **Post** /profiles/{profile_id}/quote-executions | Create Quote Execution
*QuoteExecutionsAPI* | [**GetQuoteExecution**](docs/QuoteExecutionsAPI.md#getquoteexecution) | **Get** /profiles/{profile_id}/quote-executions/{id} | Get Quote Execution
*QuoteExecutionsAPI* | [**ListQuoteExecutions**](docs/QuoteExecutionsAPI.md#listquoteexecutions) | **Get** /profiles/{profile_id}/quote-executions | List Quote Executions
*QuotesAPI* | [**ListQuotes**](docs/QuotesAPI.md#listquotes) | **Get** /quotes | List Quotes
*SandboxDepositsAPI* | [**CreateSandboxDeposit**](docs/SandboxDepositsAPI.md#createsandboxdeposit) | **Post** /sandbox/profiles/{profile_id}/deposit | Create Sandbox Deposit
*SandboxFiatTransfersAPI* | [**InitiateSandboxFiatDeposit**](docs/SandboxFiatTransfersAPI.md#initiatesandboxfiatdeposit) | **Post** /sandbox/fiat-deposits | Initiate Sandbox Fiat Deposit
*SandboxIdentityAPI* | [**SandboxSetIdentityStatus**](docs/SandboxIdentityAPI.md#sandboxsetidentitystatus) | **Put** /identity/identities/{id}/sandbox-status | Sandbox Set Identity Status
*StablecoinConversionAPI* | [**CancelStablecoinConversion**](docs/StablecoinConversionAPI.md#cancelstablecoinconversion) | **Delete** /conversion/stablecoins/{id} | Cancel Stablecoin Conversion
*StablecoinConversionAPI* | [**CreateStablecoinConversion**](docs/StablecoinConversionAPI.md#createstablecoinconversion) | **Post** /conversion/stablecoins | Create Stablecoin Conversion
*StablecoinConversionAPI* | [**GetStablecoinConversion**](docs/StablecoinConversionAPI.md#getstablecoinconversion) | **Get** /conversion/stablecoins/{id} | Get Stablecoin Conversion
*StablecoinConversionAPI* | [**ListStablecoinConversions**](docs/StablecoinConversionAPI.md#liststablecoinconversions) | **Get** /conversion/stablecoins | List Stablecoin Conversions
*TaxFormsAPI* | [**ListTaxFormRevisions**](docs/TaxFormsAPI.md#listtaxformrevisions) | **Get** /tax/tax-form-revisions | List Tax Form Revisions
*TaxFormsAPI* | [**ListTaxForms**](docs/TaxFormsAPI.md#listtaxforms) | **Get** /tax/tax-forms | List Tax Forms
*TaxLotAPI* | [**GetTaxLot**](docs/TaxLotAPI.md#gettaxlot) | **Get** /tax/tax-lots/{id} | Get Tax Lot
*TaxLotAPI* | [**ListTaxLots**](docs/TaxLotAPI.md#listtaxlots) | **Get** /tax/tax-lots | List Tax Lots
*TaxLotAPI* | [**UpdateTaxLot**](docs/TaxLotAPI.md#updatetaxlot) | **Put** /tax/tax-lots/{id} | Update Tax Lot
*TransfersAPI* | [**GetTransfer**](docs/TransfersAPI.md#gettransfer) | **Get** /transfer/transfers/{id} | Get Transfer
*TransfersAPI* | [**ListTransfers**](docs/TransfersAPI.md#listtransfers) | **Get** /transfer/transfers | List Transfers


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountAccountType](docs/AccountAccountType.md)
 - [AccountMember](docs/AccountMember.md)
 - [AccountMemberAccountRoleType](docs/AccountMemberAccountRoleType.md)
 - [AccountPurpose](docs/AccountPurpose.md)
 - [AddAccountMembersRequest](docs/AddAccountMembersRequest.md)
 - [AddAccountMembersResponse](docs/AddAccountMembersResponse.md)
 - [AddInstitutionMembersRequest](docs/AddInstitutionMembersRequest.md)
 - [AddInstitutionMembersResponse](docs/AddInstitutionMembersResponse.md)
 - [ApiHttpBody](docs/ApiHttpBody.md)
 - [Asset](docs/Asset.md)
 - [AutoConversion](docs/AutoConversion.md)
 - [Beneficiary](docs/Beneficiary.md)
 - [BeneficiaryInstitution](docs/BeneficiaryInstitution.md)
 - [BeneficiaryPerson](docs/BeneficiaryPerson.md)
 - [BookLevel](docs/BookLevel.md)
 - [BufAny](docs/BufAny.md)
 - [CapitalGainType](docs/CapitalGainType.md)
 - [CreateAccountRequest](docs/CreateAccountRequest.md)
 - [CreateCryptoWithdrawalFeeRequest](docs/CreateCryptoWithdrawalFeeRequest.md)
 - [CreateCryptoWithdrawalRequest](docs/CreateCryptoWithdrawalRequest.md)
 - [CreateDepositAddressRequest](docs/CreateDepositAddressRequest.md)
 - [CreateFiatAccountRequest](docs/CreateFiatAccountRequest.md)
 - [CreateFiatDepositInstructionsRequest](docs/CreateFiatDepositInstructionsRequest.md)
 - [CreateFiatWithdrawalRequest](docs/CreateFiatWithdrawalRequest.md)
 - [CreateIdentityRequest](docs/CreateIdentityRequest.md)
 - [CreateInternalTransferRequest](docs/CreateInternalTransferRequest.md)
 - [CreateOrderRequest](docs/CreateOrderRequest.md)
 - [CreateProfileRequest](docs/CreateProfileRequest.md)
 - [CreateQuoteExecutionRequest](docs/CreateQuoteExecutionRequest.md)
 - [CreateSandboxDepositRequest](docs/CreateSandboxDepositRequest.md)
 - [CreateSandboxDepositResponse](docs/CreateSandboxDepositResponse.md)
 - [CreateStablecoinConversionRequest](docs/CreateStablecoinConversionRequest.md)
 - [CryptoNetwork](docs/CryptoNetwork.md)
 - [CryptoWithdrawalFee](docs/CryptoWithdrawalFee.md)
 - [CustomerDueDiligence](docs/CustomerDueDiligence.md)
 - [CustomerDueDiligenceNetWorthRange](docs/CustomerDueDiligenceNetWorthRange.md)
 - [CustomerDueDiligenceTransferValueRange](docs/CustomerDueDiligenceTransferValueRange.md)
 - [CustomerDueDiligenceYearlyIncomeRange](docs/CustomerDueDiligenceYearlyIncomeRange.md)
 - [DepositAddress](docs/DepositAddress.md)
 - [DepositAddressConversionTargetAsset](docs/DepositAddressConversionTargetAsset.md)
 - [DocumentDescription](docs/DocumentDescription.md)
 - [DocumentType](docs/DocumentType.md)
 - [DocumentUploadRequest](docs/DocumentUploadRequest.md)
 - [DocumentUploadResponse](docs/DocumentUploadResponse.md)
 - [EmploymentStatus](docs/EmploymentStatus.md)
 - [ExchangeStats](docs/ExchangeStats.md)
 - [Execution](docs/Execution.md)
 - [FiatAccount](docs/FiatAccount.md)
 - [FiatAccountOwner](docs/FiatAccountOwner.md)
 - [FiatAccountOwnerInstitutionDetails](docs/FiatAccountOwnerInstitutionDetails.md)
 - [FiatAccountOwnerPersonDetails](docs/FiatAccountOwnerPersonDetails.md)
 - [FiatAccountStatus](docs/FiatAccountStatus.md)
 - [FiatDepositInstructions](docs/FiatDepositInstructions.md)
 - [FiatDepositInstructionsStatus](docs/FiatDepositInstructionsStatus.md)
 - [FiatNetwork](docs/FiatNetwork.md)
 - [FiatNetworkInstructions](docs/FiatNetworkInstructions.md)
 - [FiatNetworkInstructionsCbit](docs/FiatNetworkInstructionsCbit.md)
 - [FiatNetworkInstructionsWire](docs/FiatNetworkInstructionsWire.md)
 - [FiatWireAccountType](docs/FiatWireAccountType.md)
 - [FundsSource](docs/FundsSource.md)
 - [GetOrderBookResponse](docs/GetOrderBookResponse.md)
 - [HistoricalPrice](docs/HistoricalPrice.md)
 - [Identity](docs/Identity.md)
 - [IdentityMailingAddress](docs/IdentityMailingAddress.md)
 - [IdentityStatus](docs/IdentityStatus.md)
 - [IdentityType](docs/IdentityType.md)
 - [IdentityprotoVerifierType](docs/IdentityprotoVerifierType.md)
 - [Increment](docs/Increment.md)
 - [InitiateSandboxFiatDepositRequest](docs/InitiateSandboxFiatDepositRequest.md)
 - [InstitutionCIPIDType](docs/InstitutionCIPIDType.md)
 - [InstitutionDetails](docs/InstitutionDetails.md)
 - [InstitutionMember](docs/InstitutionMember.md)
 - [InstitutionRoleType](docs/InstitutionRoleType.md)
 - [InstitutionSubType](docs/InstitutionSubType.md)
 - [InstitutionType](docs/InstitutionType.md)
 - [ListAccountsRequestOrderBy](docs/ListAccountsRequestOrderBy.md)
 - [ListAccountsResponse](docs/ListAccountsResponse.md)
 - [ListDepositAddressesRequestOrderBy](docs/ListDepositAddressesRequestOrderBy.md)
 - [ListDepositAddressesResponse](docs/ListDepositAddressesResponse.md)
 - [ListExecutionsResponse](docs/ListExecutionsResponse.md)
 - [ListFiatAccountsRequestOrderBy](docs/ListFiatAccountsRequestOrderBy.md)
 - [ListFiatAccountsResponse](docs/ListFiatAccountsResponse.md)
 - [ListFiatDepositInstructionsResponse](docs/ListFiatDepositInstructionsResponse.md)
 - [ListHistoricalPricesResponse](docs/ListHistoricalPricesResponse.md)
 - [ListIdentitiesRequestOrderBy](docs/ListIdentitiesRequestOrderBy.md)
 - [ListIdentitiesResponse](docs/ListIdentitiesResponse.md)
 - [ListIdentityDocumentsResponse](docs/ListIdentityDocumentsResponse.md)
 - [ListMarketsResponse](docs/ListMarketsResponse.md)
 - [ListOrdersResponse](docs/ListOrdersResponse.md)
 - [ListPricesResponse](docs/ListPricesResponse.md)
 - [ListProfileBalancesResponse](docs/ListProfileBalancesResponse.md)
 - [ListProfilesRequestOrderBy](docs/ListProfilesRequestOrderBy.md)
 - [ListProfilesResponse](docs/ListProfilesResponse.md)
 - [ListQuoteExecutionsResponse](docs/ListQuoteExecutionsResponse.md)
 - [ListQuotesResponse](docs/ListQuotesResponse.md)
 - [ListRecentExecutionsResponse](docs/ListRecentExecutionsResponse.md)
 - [ListStablecoinConversionsResponse](docs/ListStablecoinConversionsResponse.md)
 - [ListTaxFormRevisionsResponse](docs/ListTaxFormRevisionsResponse.md)
 - [ListTaxFormsRequestOrderBy](docs/ListTaxFormsRequestOrderBy.md)
 - [ListTaxFormsResponse](docs/ListTaxFormsResponse.md)
 - [ListTaxLotsResponse](docs/ListTaxLotsResponse.md)
 - [ListTickersResponse](docs/ListTickersResponse.md)
 - [ListTransferLimitsResponse](docs/ListTransferLimitsResponse.md)
 - [ListTransfersRequestOrderBy](docs/ListTransfersRequestOrderBy.md)
 - [ListTransfersResponse](docs/ListTransfersResponse.md)
 - [MailingAddress](docs/MailingAddress.md)
 - [Market](docs/Market.md)
 - [MarketDetails](docs/MarketDetails.md)
 - [Order](docs/Order.md)
 - [OrderBy](docs/OrderBy.md)
 - [OrderSide](docs/OrderSide.md)
 - [OrderStatus](docs/OrderStatus.md)
 - [OrderType](docs/OrderType.md)
 - [Pagination](docs/Pagination.md)
 - [PassthroughVerificationField](docs/PassthroughVerificationField.md)
 - [PassthroughVerifierType](docs/PassthroughVerifierType.md)
 - [PersonDetails](docs/PersonDetails.md)
 - [PersonDetailsCIPIDType](docs/PersonDetailsCIPIDType.md)
 - [PricePriceMarket](docs/PricePriceMarket.md)
 - [Pricing](docs/Pricing.md)
 - [Problem](docs/Problem.md)
 - [Profile](docs/Profile.md)
 - [ProfileBalance](docs/ProfileBalance.md)
 - [ProfileType](docs/ProfileType.md)
 - [Quote](docs/Quote.md)
 - [QuoteExecution](docs/QuoteExecution.md)
 - [QuoteExecutionStatus](docs/QuoteExecutionStatus.md)
 - [RecentExecution](docs/RecentExecution.md)
 - [RegulationStatus](docs/RegulationStatus.md)
 - [RetryIdVerificationResponse](docs/RetryIdVerificationResponse.md)
 - [SandboxSetIdentityStatusRequest](docs/SandboxSetIdentityStatusRequest.md)
 - [SetDisable](docs/SetDisable.md)
 - [SetVerifierCredentialsRequest](docs/SetVerifierCredentialsRequest.md)
 - [SortOrder](docs/SortOrder.md)
 - [StablecoinConversion](docs/StablecoinConversion.md)
 - [TINVerificationStatus](docs/TINVerificationStatus.md)
 - [TaxDetail](docs/TaxDetail.md)
 - [TaxFormType](docs/TaxFormType.md)
 - [TaxFormURL](docs/TaxFormURL.md)
 - [TaxLot](docs/TaxLot.md)
 - [TaxLotStatus](docs/TaxLotStatus.md)
 - [TaxprotoOrderBy](docs/TaxprotoOrderBy.md)
 - [TickerRecord](docs/TickerRecord.md)
 - [TimeInForce](docs/TimeInForce.md)
 - [TimestampFilter](docs/TimestampFilter.md)
 - [TimestampRange](docs/TimestampRange.md)
 - [TradingType](docs/TradingType.md)
 - [TransactionType](docs/TransactionType.md)
 - [Transfer](docs/Transfer.md)
 - [TransferDirection](docs/TransferDirection.md)
 - [TransferLimit](docs/TransferLimit.md)
 - [TransferStatus](docs/TransferStatus.md)
 - [TransferType](docs/TransferType.md)
 - [UpdateAccountRequest](docs/UpdateAccountRequest.md)
 - [UpdateFiatAccountRequest](docs/UpdateFiatAccountRequest.md)
 - [UpdateFiatAccountRequestUpdateFiatNetworkInstructions](docs/UpdateFiatAccountRequestUpdateFiatNetworkInstructions.md)
 - [UpdateFiatNetworkInstructionsUpdateWire](docs/UpdateFiatNetworkInstructionsUpdateWire.md)
 - [UpdateIdentityRequest](docs/UpdateIdentityRequest.md)
 - [UpdateTaxLotRequest](docs/UpdateTaxLotRequest.md)
 - [UpdateWireUpdateRoutingDetails](docs/UpdateWireUpdateRoutingDetails.md)
 - [WealthSource](docs/WealthSource.md)
 - [WireRoutingDetails](docs/WireRoutingDetails.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### OAuth2


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **funding:read_profile**: Read profile and balance data
 - **exchange:read_quote**: Read quotes and quote history
 - **exchange:write_quote_execution**: Execute on quoted prices
 - **exchange:read_quote_execution**: Read quote executions
 - **exchange:write_order**: Create and Cancel Orders
 - **exchange:read_order**: Read Orders and Executions
 - **exchange:historical_prices**: Read Historical Pricing Data
 - **transfer:read_transfer**: Read deposit and withdrawal transfers
 - **transfer:read_deposit_address**: Read deposit addresses
 - **transfer:write_deposit_address**: Create and manage deposit addresses
 - **fee:write_crypto_withdrawal_fee**: Create crypto withdrawal fees
 - **transfer:write_crypto_withdrawal**: Create crypto withdrawal transfers
 - **conversion:read_conversion_stablecoin**: Retrieve requested or completed conversions
 - **conversion:write_conversion_stablecoin**: Create or cancel a conversion request

Example

```go
auth := context.WithValue(context.Background(), paxos.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, paxos.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



