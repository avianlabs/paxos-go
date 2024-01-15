# \PricingAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListHistoricalPrices**](PricingAPI.md#ListHistoricalPrices) | **Get** /markets/{market}/historical-prices | List Historical Prices
[**ListPrices**](PricingAPI.md#ListPrices) | **Get** /all-markets/prices | List Prices
[**ListTickers**](PricingAPI.md#ListTickers) | **Get** /all-markets/ticker | List Tickers



## ListHistoricalPrices

> ListHistoricalPricesResponse ListHistoricalPrices(ctx, market).MaxDataPoints(maxDataPoints).RangeBegin(rangeBegin).RangeEnd(rangeEnd).PaginationLimit(paginationLimit).PaginationOffset(paginationOffset).Increment(increment).Execute()

List Historical Prices



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
	market := "market_example" // string | Market of Order Book.
	maxDataPoints := int64(789) // int64 | Maximum number of data points to return.  The time frame of the increments will be inferred by finding the most granular increment without breaching the `max_data_points` value based on a set of recording data points. (optional)
	rangeBegin := time.Now() // time.Time | Only return records after this timestamp, inclusive. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	rangeEnd := time.Now() // time.Time | Only return records before this timestamp, inclusive. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	paginationLimit := int32(56) // int32 | Number of results to return (optional)
	paginationOffset := int32(56) // int32 | Number of results to skip (optional)
	increment := "increment_example" // string | Time increment between prices. Returns data exclusive from `range.begin` and `range.end`. Do not use with `max_data_points`, as this results in an error. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingAPI.ListHistoricalPrices(context.Background(), market).MaxDataPoints(maxDataPoints).RangeBegin(rangeBegin).RangeEnd(rangeEnd).PaginationLimit(paginationLimit).PaginationOffset(paginationOffset).Increment(increment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.ListHistoricalPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHistoricalPrices`: ListHistoricalPricesResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.ListHistoricalPrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**market** | **string** | Market of Order Book. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListHistoricalPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxDataPoints** | **int64** | Maximum number of data points to return.  The time frame of the increments will be inferred by finding the most granular increment without breaching the &#x60;max_data_points&#x60; value based on a set of recording data points. | 
 **rangeBegin** | **time.Time** | Only return records after this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **rangeEnd** | **time.Time** | Only return records before this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **paginationLimit** | **int32** | Number of results to return | 
 **paginationOffset** | **int32** | Number of results to skip | 
 **increment** | **string** | Time increment between prices. Returns data exclusive from &#x60;range.begin&#x60; and &#x60;range.end&#x60;. Do not use with &#x60;max_data_points&#x60;, as this results in an error. | 

### Return type

[**ListHistoricalPricesResponse**](ListHistoricalPricesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrices

> ListPricesResponse ListPrices(ctx).Markets(markets).Execute()

List Prices



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
	markets := []string{"Markets_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingAPI.ListPrices(context.Background()).Markets(markets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.ListPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPrices`: ListPricesResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.ListPrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markets** | **[]string** |  | 

### Return type

[**ListPricesResponse**](ListPricesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTickers

> ListTickersResponse ListTickers(ctx).Execute()

List Tickers



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingAPI.ListTickers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.ListTickers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTickers`: ListTickersResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.ListTickers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTickersRequest struct via the builder pattern


### Return type

[**ListTickersResponse**](ListTickersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

