# \FiatTransfersAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFiatAccount**](FiatTransfersAPI.md#CreateFiatAccount) | **Post** /transfer/fiat-accounts | Create Fiat Account
[**CreateFiatDepositInstructions**](FiatTransfersAPI.md#CreateFiatDepositInstructions) | **Post** /transfer/fiat-deposit-instructions | Create Fiat Deposit Instructions
[**CreateFiatWithdrawal**](FiatTransfersAPI.md#CreateFiatWithdrawal) | **Post** /transfer/fiat-withdrawals | Create Fiat Withdrawal
[**DeleteFiatAccount**](FiatTransfersAPI.md#DeleteFiatAccount) | **Delete** /transfer/fiat-accounts/{id} | Delete Fiat Account
[**GetFiatAccount**](FiatTransfersAPI.md#GetFiatAccount) | **Get** /transfer/fiat-accounts/{id} | Get Fiat Account
[**GetFiatDepositInstructions**](FiatTransfersAPI.md#GetFiatDepositInstructions) | **Get** /transfer/fiat-deposit-instructions/{id} | Get Fiat Deposit Instructions
[**ListFiatAccounts**](FiatTransfersAPI.md#ListFiatAccounts) | **Get** /transfer/fiat-accounts | List Fiat Accounts
[**ListFiatDepositInstructions**](FiatTransfersAPI.md#ListFiatDepositInstructions) | **Get** /transfer/fiat-deposit-instructions | List Fiat Deposit Instructions
[**UpdateFiatAccount**](FiatTransfersAPI.md#UpdateFiatAccount) | **Put** /transfer/fiat-accounts/{id} | Update Fiat Account



## CreateFiatAccount

> FiatAccount CreateFiatAccount(ctx).CreateFiatAccountRequest(createFiatAccountRequest).Execute()

Create Fiat Account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/avianlabs/paxos-go"
)

func main() {
	createFiatAccountRequest := *openapiclient.NewCreateFiatAccountRequest(*openapiclient.NewFiatAccountOwner(), *openapiclient.NewFiatNetworkInstructions()) // CreateFiatAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiatTransfersAPI.CreateFiatAccount(context.Background()).CreateFiatAccountRequest(createFiatAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiatTransfersAPI.CreateFiatAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFiatAccount`: FiatAccount
	fmt.Fprintf(os.Stdout, "Response from `FiatTransfersAPI.CreateFiatAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFiatAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFiatAccountRequest** | [**CreateFiatAccountRequest**](CreateFiatAccountRequest.md) |  | 

### Return type

[**FiatAccount**](FiatAccount.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFiatDepositInstructions

> FiatDepositInstructions CreateFiatDepositInstructions(ctx).CreateFiatDepositInstructionsRequest(createFiatDepositInstructionsRequest).Execute()

Create Fiat Deposit Instructions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/avianlabs/paxos-go"
)

func main() {
	createFiatDepositInstructionsRequest := *openapiclient.NewCreateFiatDepositInstructionsRequest("ProfileId_example", openapiclient.FiatNetwork("WIRE")) // CreateFiatDepositInstructionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiatTransfersAPI.CreateFiatDepositInstructions(context.Background()).CreateFiatDepositInstructionsRequest(createFiatDepositInstructionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiatTransfersAPI.CreateFiatDepositInstructions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFiatDepositInstructions`: FiatDepositInstructions
	fmt.Fprintf(os.Stdout, "Response from `FiatTransfersAPI.CreateFiatDepositInstructions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFiatDepositInstructionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFiatDepositInstructionsRequest** | [**CreateFiatDepositInstructionsRequest**](CreateFiatDepositInstructionsRequest.md) |  | 

### Return type

[**FiatDepositInstructions**](FiatDepositInstructions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFiatWithdrawal

> Transfer CreateFiatWithdrawal(ctx).CreateFiatWithdrawalRequest(createFiatWithdrawalRequest).Execute()

Create Fiat Withdrawal



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/avianlabs/paxos-go"
)

func main() {
	createFiatWithdrawalRequest := *openapiclient.NewCreateFiatWithdrawalRequest("Asset_example", "FiatAccountId_example", "ProfileId_example") // CreateFiatWithdrawalRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiatTransfersAPI.CreateFiatWithdrawal(context.Background()).CreateFiatWithdrawalRequest(createFiatWithdrawalRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiatTransfersAPI.CreateFiatWithdrawal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFiatWithdrawal`: Transfer
	fmt.Fprintf(os.Stdout, "Response from `FiatTransfersAPI.CreateFiatWithdrawal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFiatWithdrawalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFiatWithdrawalRequest** | [**CreateFiatWithdrawalRequest**](CreateFiatWithdrawalRequest.md) |  | 

### Return type

[**Transfer**](Transfer.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFiatAccount

> interface{} DeleteFiatAccount(ctx, id).Execute()

Delete Fiat Account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/avianlabs/paxos-go"
)

func main() {
	id := "id_example" // string | The Paxos fiat account ID (UUID).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiatTransfersAPI.DeleteFiatAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiatTransfersAPI.DeleteFiatAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFiatAccount`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FiatTransfersAPI.DeleteFiatAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Paxos fiat account ID (UUID). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFiatAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiatAccount

> FiatAccount GetFiatAccount(ctx, id).Execute()

Get Fiat Account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/avianlabs/paxos-go"
)

func main() {
	id := "id_example" // string | The Paxos fiat account ID (UUID). The Fiat Account ID (`id`) is provided in the response of the [Create Fiat Account](#operation/CreateFiatAccount). Use this ID to retrieve the instructions using [Get Fiat Account](#operation/GetFiatAccount) & [List Fiat Accounts](#operation/ListFiatAccounts).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiatTransfersAPI.GetFiatAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiatTransfersAPI.GetFiatAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiatAccount`: FiatAccount
	fmt.Fprintf(os.Stdout, "Response from `FiatTransfersAPI.GetFiatAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Paxos fiat account ID (UUID). The Fiat Account ID (&#x60;id&#x60;) is provided in the response of the [Create Fiat Account](#operation/CreateFiatAccount). Use this ID to retrieve the instructions using [Get Fiat Account](#operation/GetFiatAccount) &amp; [List Fiat Accounts](#operation/ListFiatAccounts). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiatAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FiatAccount**](FiatAccount.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiatDepositInstructions

> FiatDepositInstructions GetFiatDepositInstructions(ctx, id).Execute()

Get Fiat Deposit Instructions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/avianlabs/paxos-go"
)

func main() {
	id := "id_example" // string | Retrieve the Paxos Fiat Deposit Instructions for the given `id`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiatTransfersAPI.GetFiatDepositInstructions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiatTransfersAPI.GetFiatDepositInstructions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiatDepositInstructions`: FiatDepositInstructions
	fmt.Fprintf(os.Stdout, "Response from `FiatTransfersAPI.GetFiatDepositInstructions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Retrieve the Paxos Fiat Deposit Instructions for the given &#x60;id&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiatDepositInstructionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FiatDepositInstructions**](FiatDepositInstructions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFiatAccounts

> ListFiatAccountsResponse ListFiatAccounts(ctx).Ids(ids).RefIds(refIds).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtEq(createdAtEq).CreatedAtGte(createdAtGte).CreatedAtGt(createdAtGt).Limit(limit).Order(order).OrderBy(orderBy).PageCursor(pageCursor).Execute()

List Fiat Accounts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/avianlabs/paxos-go"
)

func main() {
	ids := []string{"Inner_example"} // []string | Optionally filter by the UUIDs of the accounts. Limit 100. (optional)
	refIds := []string{"Inner_example"} // []string | Optionally filter by the client-specified IDs provided during account creation. Limit 100. (optional)
	createdAtLt := time.Now() // time.Time | Include timestamps strictly less than lt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtLte := time.Now() // time.Time | Include timestamps less than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtEq := time.Now() // time.Time | Include timestamps exactly equal to eq. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtGte := time.Now() // time.Time | Include timestamps greater than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtGt := time.Now() // time.Time | Include timestamps strictly greater than gt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	limit := int32(56) // int32 | Number of results to return. Defaults to 100 if no limit is provided. Maximum 1000. (optional)
	order := "order_example" // string | Return items in ascending (ASC) or descending (DESC) order. Defaults to DESC. (optional)
	orderBy := "orderBy_example" // string | The specific method by which the returned results will be ordered. Defaults to CREATED_AT. (optional)
	pageCursor := "pageCursor_example" // string | Optional: Cursor for getting the next page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiatTransfersAPI.ListFiatAccounts(context.Background()).Ids(ids).RefIds(refIds).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtEq(createdAtEq).CreatedAtGte(createdAtGte).CreatedAtGt(createdAtGt).Limit(limit).Order(order).OrderBy(orderBy).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiatTransfersAPI.ListFiatAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFiatAccounts`: ListFiatAccountsResponse
	fmt.Fprintf(os.Stdout, "Response from `FiatTransfersAPI.ListFiatAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFiatAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | Optionally filter by the UUIDs of the accounts. Limit 100. | 
 **refIds** | **[]string** | Optionally filter by the client-specified IDs provided during account creation. Limit 100. | 
 **createdAtLt** | **time.Time** | Include timestamps strictly less than lt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtLte** | **time.Time** | Include timestamps less than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtEq** | **time.Time** | Include timestamps exactly equal to eq. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtGte** | **time.Time** | Include timestamps greater than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtGt** | **time.Time** | Include timestamps strictly greater than gt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **limit** | **int32** | Number of results to return. Defaults to 100 if no limit is provided. Maximum 1000. | 
 **order** | **string** | Return items in ascending (ASC) or descending (DESC) order. Defaults to DESC. | 
 **orderBy** | **string** | The specific method by which the returned results will be ordered. Defaults to CREATED_AT. | 
 **pageCursor** | **string** | Optional: Cursor for getting the next page of results. | 

### Return type

[**ListFiatAccountsResponse**](ListFiatAccountsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFiatDepositInstructions

> ListFiatDepositInstructionsResponse ListFiatDepositInstructions(ctx).Ids(ids).RefIds(refIds).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtEq(createdAtEq).CreatedAtGte(createdAtGte).CreatedAtGt(createdAtGt).Limit(limit).Order(order).PageCursor(pageCursor).Execute()

List Fiat Deposit Instructions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/avianlabs/paxos-go"
)

func main() {
	ids := []string{"Inner_example"} // []string | Optionally filter by the UUIDs of the instructions. Limit 100. (optional)
	refIds := []string{"Inner_example"} // []string | Optionally filter by the client-specified IDs provided during instructions creation. Limit 100. (optional)
	createdAtLt := time.Now() // time.Time | Include timestamps strictly less than lt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtLte := time.Now() // time.Time | Include timestamps less than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtEq := time.Now() // time.Time | Include timestamps exactly equal to eq. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtGte := time.Now() // time.Time | Include timestamps greater than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtGt := time.Now() // time.Time | Include timestamps strictly greater than gt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	limit := int32(56) // int32 | Number of results to return. Defaults to 100 if no limit is provided. Maximum 1000. (optional)
	order := "order_example" // string | Return items in ascending (ASC) or descending (DESC) order. Defaults to DESC. (optional)
	pageCursor := "pageCursor_example" // string | Optional: Cursor for getting the next page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiatTransfersAPI.ListFiatDepositInstructions(context.Background()).Ids(ids).RefIds(refIds).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtEq(createdAtEq).CreatedAtGte(createdAtGte).CreatedAtGt(createdAtGt).Limit(limit).Order(order).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiatTransfersAPI.ListFiatDepositInstructions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFiatDepositInstructions`: ListFiatDepositInstructionsResponse
	fmt.Fprintf(os.Stdout, "Response from `FiatTransfersAPI.ListFiatDepositInstructions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFiatDepositInstructionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | Optionally filter by the UUIDs of the instructions. Limit 100. | 
 **refIds** | **[]string** | Optionally filter by the client-specified IDs provided during instructions creation. Limit 100. | 
 **createdAtLt** | **time.Time** | Include timestamps strictly less than lt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtLte** | **time.Time** | Include timestamps less than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtEq** | **time.Time** | Include timestamps exactly equal to eq. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtGte** | **time.Time** | Include timestamps greater than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtGt** | **time.Time** | Include timestamps strictly greater than gt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **limit** | **int32** | Number of results to return. Defaults to 100 if no limit is provided. Maximum 1000. | 
 **order** | **string** | Return items in ascending (ASC) or descending (DESC) order. Defaults to DESC. | 
 **pageCursor** | **string** | Optional: Cursor for getting the next page of results. | 

### Return type

[**ListFiatDepositInstructionsResponse**](ListFiatDepositInstructionsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFiatAccount

> FiatAccount UpdateFiatAccount(ctx, id).TransferPublicUpdateFiatAccountBody(transferPublicUpdateFiatAccountBody).Execute()

Update Fiat Account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/avianlabs/paxos-go"
)

func main() {
	id := "id_example" // string | The Paxos fiat account ID (UUID). The Fiat Account ID (`id`) is provided in the response of the [Create Fiat Account](#operation/CreateFiatAccount). Use this ID to retrieve the instructions using [Get Fiat Account](#operation/GetFiatAccount) & [List Fiat Accounts](#operation/ListFiatAccounts).
	transferPublicUpdateFiatAccountBody := *openapiclient.NewTransferPublicUpdateFiatAccountBody() // TransferPublicUpdateFiatAccountBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiatTransfersAPI.UpdateFiatAccount(context.Background(), id).TransferPublicUpdateFiatAccountBody(transferPublicUpdateFiatAccountBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiatTransfersAPI.UpdateFiatAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFiatAccount`: FiatAccount
	fmt.Fprintf(os.Stdout, "Response from `FiatTransfersAPI.UpdateFiatAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Paxos fiat account ID (UUID). The Fiat Account ID (&#x60;id&#x60;) is provided in the response of the [Create Fiat Account](#operation/CreateFiatAccount). Use this ID to retrieve the instructions using [Get Fiat Account](#operation/GetFiatAccount) &amp; [List Fiat Accounts](#operation/ListFiatAccounts). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFiatAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferPublicUpdateFiatAccountBody** | [**TransferPublicUpdateFiatAccountBody**](TransferPublicUpdateFiatAccountBody.md) |  | 

### Return type

[**FiatAccount**](FiatAccount.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

