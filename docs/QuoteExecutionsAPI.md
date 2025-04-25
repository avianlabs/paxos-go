# \QuoteExecutionsAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuoteExecution**](QuoteExecutionsAPI.md#CreateQuoteExecution) | **Post** /profiles/{profile_id}/quote-executions | Create Quote Execution
[**GetQuoteExecution**](QuoteExecutionsAPI.md#GetQuoteExecution) | **Get** /profiles/{profile_id}/quote-executions/{id} | Get Quote Execution
[**ListQuoteExecutions**](QuoteExecutionsAPI.md#ListQuoteExecutions) | **Get** /profiles/{profile_id}/quote-executions | List Quote Executions



## CreateQuoteExecution

> QuoteExecution CreateQuoteExecution(ctx, profileId).ExchangePublicCreateQuoteExecutionBody(exchangePublicCreateQuoteExecutionBody).Execute()

Create Quote Execution



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
	profileId := "profileId_example" // string | The ID of the profile under which to execute this order.
	exchangePublicCreateQuoteExecutionBody := *openapiclient.NewExchangePublicCreateQuoteExecutionBody("QuoteId_example") // ExchangePublicCreateQuoteExecutionBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuoteExecutionsAPI.CreateQuoteExecution(context.Background(), profileId).ExchangePublicCreateQuoteExecutionBody(exchangePublicCreateQuoteExecutionBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteExecutionsAPI.CreateQuoteExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateQuoteExecution`: QuoteExecution
	fmt.Fprintf(os.Stdout, "Response from `QuoteExecutionsAPI.CreateQuoteExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The ID of the profile under which to execute this order. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuoteExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exchangePublicCreateQuoteExecutionBody** | [**ExchangePublicCreateQuoteExecutionBody**](ExchangePublicCreateQuoteExecutionBody.md) |  | 

### Return type

[**QuoteExecution**](QuoteExecution.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuoteExecution

> QuoteExecution GetQuoteExecution(ctx, profileId, id).Execute()

Get Quote Execution



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
	profileId := "profileId_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuoteExecutionsAPI.GetQuoteExecution(context.Background(), profileId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteExecutionsAPI.GetQuoteExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuoteExecution`: QuoteExecution
	fmt.Fprintf(os.Stdout, "Response from `QuoteExecutionsAPI.GetQuoteExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**QuoteExecution**](QuoteExecution.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQuoteExecutions

> ListQuoteExecutionsResponse ListQuoteExecutions(ctx, profileId).Side(side).Market(market).Status(status).CreatedAtBegin(createdAtBegin).CreatedAtEnd(createdAtEnd).Limit(limit).Order(order).OrderBy(orderBy).PageCursor(pageCursor).Execute()

List Quote Executions



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
	profileId := "profileId_example" // string | The profile ID associated with the orders.
	side := "side_example" // string | Filter by buy or sell side. (optional)
	market := "market_example" // string | Filter by the trading pair. (optional)
	status := "status_example" // string | Filter by the status of the order. (optional)
	createdAtBegin := time.Now() // time.Time | Only return records after this timestamp, inclusive. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtEnd := time.Now() // time.Time | Only return records before this timestamp, inclusive. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	limit := int32(56) // int32 | Number of results to return. (optional)
	order := "order_example" // string | Return items in ascending (ASC) or descending (DESC) order. Default order is descending (DESC). (optional)
	orderBy := "orderBy_example" // string | The specific method by which the returned results will be ordered. (optional)
	pageCursor := "pageCursor_example" // string | Cursor token for fetching the next page. If using this then do not use `pagination.limit` and `pagination.offset` fields. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuoteExecutionsAPI.ListQuoteExecutions(context.Background(), profileId).Side(side).Market(market).Status(status).CreatedAtBegin(createdAtBegin).CreatedAtEnd(createdAtEnd).Limit(limit).Order(order).OrderBy(orderBy).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteExecutionsAPI.ListQuoteExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQuoteExecutions`: ListQuoteExecutionsResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteExecutionsAPI.ListQuoteExecutions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The profile ID associated with the orders. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListQuoteExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **side** | **string** | Filter by buy or sell side. | 
 **market** | **string** | Filter by the trading pair. | 
 **status** | **string** | Filter by the status of the order. | 
 **createdAtBegin** | **time.Time** | Only return records after this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtEnd** | **time.Time** | Only return records before this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **limit** | **int32** | Number of results to return. | 
 **order** | **string** | Return items in ascending (ASC) or descending (DESC) order. Default order is descending (DESC). | 
 **orderBy** | **string** | The specific method by which the returned results will be ordered. | 
 **pageCursor** | **string** | Cursor token for fetching the next page. If using this then do not use &#x60;pagination.limit&#x60; and &#x60;pagination.offset&#x60; fields. | 

### Return type

[**ListQuoteExecutionsResponse**](ListQuoteExecutionsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

