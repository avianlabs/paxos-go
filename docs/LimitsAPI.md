# \LimitsAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListTransferLimits**](LimitsAPI.md#ListTransferLimits) | **Get** /transfer/limits/utilizations | List Transfer Limits



## ListTransferLimits

> ListTransferLimitsResponse ListTransferLimits(ctx).TransferType(transferType).Execute()

List Transfer Limits



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
	transferType := "transferType_example" // string | Type of transfer to fetch limits for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LimitsAPI.ListTransferLimits(context.Background()).TransferType(transferType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LimitsAPI.ListTransferLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransferLimits`: ListTransferLimitsResponse
	fmt.Fprintf(os.Stdout, "Response from `LimitsAPI.ListTransferLimits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransferLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferType** | **string** | Type of transfer to fetch limits for. | 

### Return type

[**ListTransferLimitsResponse**](ListTransferLimitsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

