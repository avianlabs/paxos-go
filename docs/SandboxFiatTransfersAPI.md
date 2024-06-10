# \SandboxFiatTransfersAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InitiateSandboxFiatDeposit**](SandboxFiatTransfersAPI.md#InitiateSandboxFiatDeposit) | **Post** /sandbox/fiat-deposits | Initiate Sandbox Fiat Deposit



## InitiateSandboxFiatDeposit

> map[string]interface{} InitiateSandboxFiatDeposit(ctx).InitiateSandboxFiatDepositRequest(initiateSandboxFiatDepositRequest).Execute()

Initiate Sandbox Fiat Deposit



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
	initiateSandboxFiatDepositRequest := *openapiclient.NewInitiateSandboxFiatDepositRequest("Amount_example", "Asset_example", "MemoId_example", *openapiclient.NewFiatNetworkInstructions()) // InitiateSandboxFiatDepositRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxFiatTransfersAPI.InitiateSandboxFiatDeposit(context.Background()).InitiateSandboxFiatDepositRequest(initiateSandboxFiatDepositRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxFiatTransfersAPI.InitiateSandboxFiatDeposit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateSandboxFiatDeposit`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SandboxFiatTransfersAPI.InitiateSandboxFiatDeposit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateSandboxFiatDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **initiateSandboxFiatDepositRequest** | [**InitiateSandboxFiatDepositRequest**](InitiateSandboxFiatDepositRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

