# \QuotesAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListQuotes**](QuotesAPI.md#ListQuotes) | **Get** /quotes | List Quotes



## ListQuotes

> ListQuotesResponse ListQuotes(ctx).Markets(markets).Execute()

List Quotes



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
	resp, r, err := apiClient.QuotesAPI.ListQuotes(context.Background()).Markets(markets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ListQuotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQuotes`: ListQuotesResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ListQuotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markets** | **[]string** |  | 

### Return type

[**ListQuotesResponse**](ListQuotesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

