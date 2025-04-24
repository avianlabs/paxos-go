# \SandboxDepositsAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSandboxDeposit**](SandboxDepositsAPI.md#CreateSandboxDeposit) | **Post** /sandbox/profiles/{profile_id}/deposit | Create Sandbox Deposit



## CreateSandboxDeposit

> CreateSandboxDepositResponse CreateSandboxDeposit(ctx, profileId).FundingSandboxCreateSandboxDepositBody(fundingSandboxCreateSandboxDepositBody).Execute()

Create Sandbox Deposit



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
	profileId := "profileId_example" // string | The ID of the profile that will credited with the assets.
	fundingSandboxCreateSandboxDepositBody := *openapiclient.NewFundingSandboxCreateSandboxDepositBody("Asset_example", "Amount_example") // FundingSandboxCreateSandboxDepositBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxDepositsAPI.CreateSandboxDeposit(context.Background(), profileId).FundingSandboxCreateSandboxDepositBody(fundingSandboxCreateSandboxDepositBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxDepositsAPI.CreateSandboxDeposit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSandboxDeposit`: CreateSandboxDepositResponse
	fmt.Fprintf(os.Stdout, "Response from `SandboxDepositsAPI.CreateSandboxDeposit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The ID of the profile that will credited with the assets. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSandboxDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fundingSandboxCreateSandboxDepositBody** | [**FundingSandboxCreateSandboxDepositBody**](FundingSandboxCreateSandboxDepositBody.md) |  | 

### Return type

[**CreateSandboxDepositResponse**](CreateSandboxDepositResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

