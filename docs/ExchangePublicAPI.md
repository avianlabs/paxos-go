# \ExchangePublicAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListProfileExecutions**](ExchangePublicAPI.md#ListProfileExecutions) | **Get** /profiles/{profile_id}/executions | List Executions
[**ListProfileOrders**](ExchangePublicAPI.md#ListProfileOrders) | **Get** /profiles/{profile_id}/orders | List Orders



## ListProfileExecutions

> ListExecutionsResponse ListProfileExecutions(ctx, profileId).OrderId(orderId).SinceExecutionId(sinceExecutionId).RangeBegin(rangeBegin).RangeEnd(rangeEnd).PageCursor(pageCursor).Limit(limit).Execute()

List Executions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	profileId := "profileId_example" // string | The ProfileId associated with the orders.
	orderId := "orderId_example" // string | Filter executions for a single order id. (optional)
	sinceExecutionId := "sinceExecutionId_example" // string | Excludes executions after this id. (optional)
	rangeBegin := time.Now() // time.Time | Only return records after this timestamp, inclusive. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	rangeEnd := time.Now() // time.Time | Only return records before this timestamp, inclusive. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	pageCursor := "pageCursor_example" // string | Cursor token for fetching the next page. (optional)
	limit := int32(56) // int32 | Number of results to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExchangePublicAPI.ListProfileExecutions(context.Background(), profileId).OrderId(orderId).SinceExecutionId(sinceExecutionId).RangeBegin(rangeBegin).RangeEnd(rangeEnd).PageCursor(pageCursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExchangePublicAPI.ListProfileExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProfileExecutions`: ListExecutionsResponse
	fmt.Fprintf(os.Stdout, "Response from `ExchangePublicAPI.ListProfileExecutions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The ProfileId associated with the orders. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProfileExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderId** | **string** | Filter executions for a single order id. | 
 **sinceExecutionId** | **string** | Excludes executions after this id. | 
 **rangeBegin** | **time.Time** | Only return records after this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **rangeEnd** | **time.Time** | Only return records before this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **pageCursor** | **string** | Cursor token for fetching the next page. | 
 **limit** | **int32** | Number of results to return. | 

### Return type

[**ListExecutionsResponse**](ListExecutionsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProfileOrders

> ListOrdersResponse ListProfileOrders(ctx, profileId).Market(market).Status(status).OrderTimeBegin(orderTimeBegin).OrderTimeEnd(orderTimeEnd).RefIds(refIds).PageCursor(pageCursor).Limit(limit).Execute()

List Orders



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	profileId := "profileId_example" // string | The ProfileId associated with the orders.
	market := "market_example" // string | Filter by the trading pair. (optional)
	status := "status_example" // string | Filter by the status of the order. (optional)
	orderTimeBegin := time.Now() // time.Time | Only return records after this timestamp, inclusive. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	orderTimeEnd := time.Now() // time.Time | Only return records before this timestamp, inclusive. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	refIds := []string{"Inner_example"} // []string | The idempotence ids provided during order creation. (optional)
	pageCursor := "pageCursor_example" // string | Cursor token for fetching the next page. If using this then do not use paginationLimit and paginationOffset fields. (optional)
	limit := int32(56) // int32 | Number of results to return. If using this then do not use paginationLimit and paginationOffset fields. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExchangePublicAPI.ListProfileOrders(context.Background(), profileId).Market(market).Status(status).OrderTimeBegin(orderTimeBegin).OrderTimeEnd(orderTimeEnd).RefIds(refIds).PageCursor(pageCursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExchangePublicAPI.ListProfileOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProfileOrders`: ListOrdersResponse
	fmt.Fprintf(os.Stdout, "Response from `ExchangePublicAPI.ListProfileOrders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The ProfileId associated with the orders. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProfileOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** | Filter by the trading pair. | 
 **status** | **string** | Filter by the status of the order. | 
 **orderTimeBegin** | **time.Time** | Only return records after this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **orderTimeEnd** | **time.Time** | Only return records before this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **refIds** | **[]string** | The idempotence ids provided during order creation. | 
 **pageCursor** | **string** | Cursor token for fetching the next page. If using this then do not use paginationLimit and paginationOffset fields. | 
 **limit** | **int32** | Number of results to return. If using this then do not use paginationLimit and paginationOffset fields. | 

### Return type

[**ListOrdersResponse**](ListOrdersResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

