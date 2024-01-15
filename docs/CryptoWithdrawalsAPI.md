# \CryptoWithdrawalsAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCryptoWithdrawal**](CryptoWithdrawalsAPI.md#CreateCryptoWithdrawal) | **Post** /transfer/crypto-withdrawals | Create Crypto Withdrawal



## CreateCryptoWithdrawal

> Transfer CreateCryptoWithdrawal(ctx).CreateCryptoWithdrawalRequest(createCryptoWithdrawalRequest).Execute()

Create Crypto Withdrawal



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
	createCryptoWithdrawalRequest := *openapiclient.NewCreateCryptoWithdrawalRequest("ProfileId_example", "DestinationAddress_example", "Asset_example", openapiclient.CryptoNetwork("BITCOIN")) // CreateCryptoWithdrawalRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoWithdrawalsAPI.CreateCryptoWithdrawal(context.Background()).CreateCryptoWithdrawalRequest(createCryptoWithdrawalRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoWithdrawalsAPI.CreateCryptoWithdrawal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCryptoWithdrawal`: Transfer
	fmt.Fprintf(os.Stdout, "Response from `CryptoWithdrawalsAPI.CreateCryptoWithdrawal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCryptoWithdrawalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCryptoWithdrawalRequest** | [**CreateCryptoWithdrawalRequest**](CreateCryptoWithdrawalRequest.md) |  | 

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

