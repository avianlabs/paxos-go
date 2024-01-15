# \OrdersAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrder**](OrdersAPI.md#CancelOrder) | **Delete** /profiles/{profile_id}/orders/{id} | Cancel Order
[**CreateOrder**](OrdersAPI.md#CreateOrder) | **Post** /profiles/{profile_id}/orders | Create Order
[**GetOrder**](OrdersAPI.md#GetOrder) | **Get** /profiles/{profile_id}/orders/{id} | Get Order
[**ListExecutions**](OrdersAPI.md#ListExecutions) | **Get** /executions | List Executions
[**ListOrders**](OrdersAPI.md#ListOrders) | **Get** /orders | List Orders



## CancelOrder

> map[string]interface{} CancelOrder(ctx, profileId, id).Execute()

Cancel Order



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
	profileId := "profileId_example" // string | The profile ID associated with the order.
	id := "id_example" // string | The UUID of the Order.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.CancelOrder(context.Background(), profileId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CancelOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelOrder`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CancelOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The profile ID associated with the order. | 
**id** | **string** | The UUID of the Order. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrder

> Order CreateOrder(ctx, profileId).CreateOrderRequest(createOrderRequest).Execute()

Create Order



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
	profileId := "profileId_example" // string | The profileId the order will be associated with.
	createOrderRequest := *openapiclient.NewCreateOrderRequest(openapiclient.OrderSide("BUY"), openapiclient.Market("ETHEUR"), openapiclient.OrderType("LIMIT")) // CreateOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.CreateOrder(context.Background(), profileId).CreateOrderRequest(createOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CreateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrder`: Order
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CreateOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The profileId the order will be associated with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrderRequest** | [**CreateOrderRequest**](CreateOrderRequest.md) |  | 

### Return type

[**Order**](Order.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrder

> Order GetOrder(ctx, profileId, id).Execute()

Get Order



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
	profileId := "profileId_example" // string | The profile ID associated with the order.
	id := "id_example" // string | The UUID of the Order.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.GetOrder(context.Background(), profileId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrder`: Order
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The profile ID associated with the order. | 
**id** | **string** | The UUID of the Order. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Order**](Order.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExecutions

> ListExecutionsResponse ListExecutions(ctx).ProfileId(profileId).AccountId(accountId).OrderId(orderId).SinceExecutionId(sinceExecutionId).RangeBegin(rangeBegin).RangeEnd(rangeEnd).PageCursor(pageCursor).Limit(limit).Execute()

List Executions



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
	profileId := "profileId_example" // string | Filter executions by the Profile ID associated with the orders. (optional)
	accountId := "accountId_example" // string | Filter executions by the Account ID associated with the orders. (optional)
	orderId := "orderId_example" // string | Filter executions for a single order ID. (optional)
	sinceExecutionId := "sinceExecutionId_example" // string | Excludes executions after the given ID. (optional)
	rangeBegin := time.Now() // time.Time | Only return records after this timestamp, inclusive. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	rangeEnd := time.Now() // time.Time | Only return records before this timestamp, inclusive. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	pageCursor := "pageCursor_example" // string | Cursor token for fetching the next page. (optional)
	limit := int32(56) // int32 | Number of results to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.ListExecutions(context.Background()).ProfileId(profileId).AccountId(accountId).OrderId(orderId).SinceExecutionId(sinceExecutionId).RangeBegin(rangeBegin).RangeEnd(rangeEnd).PageCursor(pageCursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.ListExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExecutions`: ListExecutionsResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.ListExecutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | Filter executions by the Profile ID associated with the orders. | 
 **accountId** | **string** | Filter executions by the Account ID associated with the orders. | 
 **orderId** | **string** | Filter executions for a single order ID. | 
 **sinceExecutionId** | **string** | Excludes executions after the given ID. | 
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


## ListOrders

> ListOrdersResponse ListOrders(ctx).ProfileId(profileId).AccountId(accountId).Market(market).Status(status).OrderTimeBegin(orderTimeBegin).OrderTimeEnd(orderTimeEnd).RefIds(refIds).PageCursor(pageCursor).Limit(limit).Execute()

List Orders



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
	profileId := "profileId_example" // string | The profile associated with the orders. (optional)
	accountId := "accountId_example" // string | The account under which these orders are placed. (optional)
	market := "market_example" // string | Filter by the trading pair. (optional)
	status := "status_example" // string | Filter by the status of the order. (optional)
	orderTimeBegin := time.Now() // time.Time | Only return records after this timestamp, inclusive. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	orderTimeEnd := time.Now() // time.Time | Only return records before this timestamp, inclusive. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	refIds := []string{"Inner_example"} // []string | The idempotence IDs provided during order creation. (optional)
	pageCursor := "pageCursor_example" // string | Cursor token for fetching the next page. If using this then do not use paginationLimit and paginationOffset fields. (optional)
	limit := int32(56) // int32 | Number of results to return. If using this then do not use paginationLimit and paginationOffset fields. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.ListOrders(context.Background()).ProfileId(profileId).AccountId(accountId).Market(market).Status(status).OrderTimeBegin(orderTimeBegin).OrderTimeEnd(orderTimeEnd).RefIds(refIds).PageCursor(pageCursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.ListOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrders`: ListOrdersResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.ListOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | The profile associated with the orders. | 
 **accountId** | **string** | The account under which these orders are placed. | 
 **market** | **string** | Filter by the trading pair. | 
 **status** | **string** | Filter by the status of the order. | 
 **orderTimeBegin** | **time.Time** | Only return records after this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **orderTimeEnd** | **time.Time** | Only return records before this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **refIds** | **[]string** | The idempotence IDs provided during order creation. | 
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

