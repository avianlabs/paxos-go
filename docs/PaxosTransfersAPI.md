# \PaxosTransfersAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaxosTransfer**](PaxosTransfersAPI.md#CreatePaxosTransfer) | **Post** /transfer/paxos | Create Paxos Transfer



## CreatePaxosTransfer

> Transfer CreatePaxosTransfer(ctx).CreatePaxosTransferRequest(createPaxosTransferRequest).Execute()

Create Paxos Transfer



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
	createPaxosTransferRequest := *openapiclient.NewCreatePaxosTransferRequest("FromProfileId_example", "ToProfileId_example", "Amount_example", "Asset_example") // CreatePaxosTransferRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaxosTransfersAPI.CreatePaxosTransfer(context.Background()).CreatePaxosTransferRequest(createPaxosTransferRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaxosTransfersAPI.CreatePaxosTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaxosTransfer`: Transfer
	fmt.Fprintf(os.Stdout, "Response from `PaxosTransfersAPI.CreatePaxosTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaxosTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPaxosTransferRequest** | [**CreatePaxosTransferRequest**](CreatePaxosTransferRequest.md) |  | 

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

