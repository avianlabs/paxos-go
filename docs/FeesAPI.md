# \FeesAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCryptoWithdrawalFee**](FeesAPI.md#CreateCryptoWithdrawalFee) | **Post** /transfer/fees/crypto-withdrawal | Create Crypto Withdrawal Fee



## CreateCryptoWithdrawalFee

> CryptoWithdrawalFee CreateCryptoWithdrawalFee(ctx).CreateCryptoWithdrawalFeeRequest(createCryptoWithdrawalFeeRequest).Execute()

Create Crypto Withdrawal Fee



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
	createCryptoWithdrawalFeeRequest := *openapiclient.NewCreateCryptoWithdrawalFeeRequest("Asset_example", "DestinationAddress_example", openapiclient.CryptoNetwork("BITCOIN")) // CreateCryptoWithdrawalFeeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeesAPI.CreateCryptoWithdrawalFee(context.Background()).CreateCryptoWithdrawalFeeRequest(createCryptoWithdrawalFeeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeesAPI.CreateCryptoWithdrawalFee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCryptoWithdrawalFee`: CryptoWithdrawalFee
	fmt.Fprintf(os.Stdout, "Response from `FeesAPI.CreateCryptoWithdrawalFee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCryptoWithdrawalFeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCryptoWithdrawalFeeRequest** | [**CreateCryptoWithdrawalFeeRequest**](CreateCryptoWithdrawalFeeRequest.md) |  | 

### Return type

[**CryptoWithdrawalFee**](CryptoWithdrawalFee.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

