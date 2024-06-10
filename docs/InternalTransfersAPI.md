# \InternalTransfersAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInternalTransfer**](InternalTransfersAPI.md#CreateInternalTransfer) | **Post** /transfer/internal | Create Internal Transfer



## CreateInternalTransfer

> Transfer CreateInternalTransfer(ctx).CreateInternalTransferRequest(createInternalTransferRequest).Execute()

Create Internal Transfer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createInternalTransferRequest := *openapiclient.NewCreateInternalTransferRequest("FromProfileId_example", "ToProfileId_example", "Amount_example", "Asset_example") // CreateInternalTransferRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalTransfersAPI.CreateInternalTransfer(context.Background()).CreateInternalTransferRequest(createInternalTransferRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalTransfersAPI.CreateInternalTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInternalTransfer`: Transfer
	fmt.Fprintf(os.Stdout, "Response from `InternalTransfersAPI.CreateInternalTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInternalTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInternalTransferRequest** | [**CreateInternalTransferRequest**](CreateInternalTransferRequest.md) |  | 

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

