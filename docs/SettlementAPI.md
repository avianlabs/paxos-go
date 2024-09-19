# \SettlementAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AffirmTransaction**](SettlementAPI.md#AffirmTransaction) | **Put** /settlement/transactions/{transaction_id}/affirm | Affirm Transaction
[**CancelTransaction**](SettlementAPI.md#CancelTransaction) | **Delete** /settlement/transactions/{transaction_id} | Cancel Transaction
[**CreateTransaction**](SettlementAPI.md#CreateTransaction) | **Post** /settlement/transactions | Create Transaction
[**GetTransaction**](SettlementAPI.md#GetTransaction) | **Get** /settlement/transactions/{transaction_id} | Get Transaction
[**ListTransactions**](SettlementAPI.md#ListTransactions) | **Get** /settlement/transactions | List Transactions



## AffirmTransaction

> Transaction AffirmTransaction(ctx, transactionId).Execute()

Affirm Transaction



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
	transactionId := "transactionId_example" // string | The unique `id` in the `Transaction` object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettlementAPI.AffirmTransaction(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettlementAPI.AffirmTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AffirmTransaction`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `SettlementAPI.AffirmTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The unique &#x60;id&#x60; in the &#x60;Transaction&#x60; object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAffirmTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelTransaction

> Transaction CancelTransaction(ctx, transactionId).Execute()

Cancel Transaction



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
	transactionId := "transactionId_example" // string | The unique `id` in the `Transaction` object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettlementAPI.CancelTransaction(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettlementAPI.CancelTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelTransaction`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `SettlementAPI.CancelTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The unique &#x60;id&#x60; in the &#x60;Transaction&#x60; object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTransaction

> Transaction CreateTransaction(ctx).CreateTransactionRequest(createTransactionRequest).Execute()

Create Transaction



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
	createTransactionRequest := *openapiclient.NewCreateTransactionRequest() // CreateTransactionRequest | Request to create a bilateral settlement transaction with one or more legs.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettlementAPI.CreateTransaction(context.Background()).CreateTransactionRequest(createTransactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettlementAPI.CreateTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTransaction`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `SettlementAPI.CreateTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTransactionRequest** | [**CreateTransactionRequest**](CreateTransactionRequest.md) | Request to create a bilateral settlement transaction with one or more legs. | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransaction

> Transaction GetTransaction(ctx, transactionId).Execute()

Get Transaction



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
	transactionId := "transactionId_example" // string | The unique `id` in the `Transaction` object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettlementAPI.GetTransaction(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettlementAPI.GetTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransaction`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `SettlementAPI.GetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The unique &#x60;id&#x60; in the &#x60;Transaction&#x60; object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> ListTransactionsResponse ListTransactions(ctx).Statuses(statuses).SourceProfileId(sourceProfileId).TargetProfileId(targetProfileId).Limit(limit).PageCursor(pageCursor).Execute()

List Transactions



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
	statuses := []string{"Statuses_example"} // []string | Transaction statuses to filter on   - PENDING: Initial state of a settlement transaction upon creation.  - SETTLED: Indicates all obligations belong to the settlement transaction have been enacted.  - EXPIRED: Indicates the settlement transaction is no longer eligible for settlement.  - CANCELLED: Indicates the settlement transaction was cancelled by the source profile.  - AFFIRMED: Indicates the settlement transaction will be eligible for settlement once within the window. (optional)
	sourceProfileId := "sourceProfileId_example" // string | The `profile_id` of the entity on the submitting side of the transaction. (optional)
	targetProfileId := "targetProfileId_example" // string | The `profile_id` of the entity on the receiving side of the transaction. (optional)
	limit := int64(789) // int64 | Number of results to return. Defaults to 100 if no limit is provided. (optional)
	pageCursor := "pageCursor_example" // string | Cursor for getting the next page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettlementAPI.ListTransactions(context.Background()).Statuses(statuses).SourceProfileId(sourceProfileId).TargetProfileId(targetProfileId).Limit(limit).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettlementAPI.ListTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransactions`: ListTransactionsResponse
	fmt.Fprintf(os.Stdout, "Response from `SettlementAPI.ListTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statuses** | **[]string** | Transaction statuses to filter on   - PENDING: Initial state of a settlement transaction upon creation.  - SETTLED: Indicates all obligations belong to the settlement transaction have been enacted.  - EXPIRED: Indicates the settlement transaction is no longer eligible for settlement.  - CANCELLED: Indicates the settlement transaction was cancelled by the source profile.  - AFFIRMED: Indicates the settlement transaction will be eligible for settlement once within the window. | 
 **sourceProfileId** | **string** | The &#x60;profile_id&#x60; of the entity on the submitting side of the transaction. | 
 **targetProfileId** | **string** | The &#x60;profile_id&#x60; of the entity on the receiving side of the transaction. | 
 **limit** | **int64** | Number of results to return. Defaults to 100 if no limit is provided. | 
 **pageCursor** | **string** | Cursor for getting the next page of results. | 

### Return type

[**ListTransactionsResponse**](ListTransactionsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

