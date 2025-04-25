# \CryptoDepositsAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RejectCryptoDeposit**](CryptoDepositsAPI.md#RejectCryptoDeposit) | **Post** /transfer/crypto-deposits/{id}/reject | Reject Crypto Deposit
[**UpdateCryptoDeposit**](CryptoDepositsAPI.md#UpdateCryptoDeposit) | **Post** /transfer/crypto-deposits/{id}/update | Update Crypto Deposit



## RejectCryptoDeposit

> map[string]interface{} RejectCryptoDeposit(ctx, id).TransferPublicRejectCryptoDepositBody(transferPublicRejectCryptoDepositBody).Execute()

Reject Crypto Deposit



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
	id := "id_example" // string | Paxos ID of the crypto deposit to reject. To retrieve the ID, log in to your Paxos account and go to the [Activity](https://account.paxos.com/wallet/activity) tab. Download the CSV activity file and check the ID column.
	transferPublicRejectCryptoDepositBody := *openapiclient.NewTransferPublicRejectCryptoDepositBody("IdentityId_example") // TransferPublicRejectCryptoDepositBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoDepositsAPI.RejectCryptoDeposit(context.Background(), id).TransferPublicRejectCryptoDepositBody(transferPublicRejectCryptoDepositBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoDepositsAPI.RejectCryptoDeposit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectCryptoDeposit`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CryptoDepositsAPI.RejectCryptoDeposit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Paxos ID of the crypto deposit to reject. To retrieve the ID, log in to your Paxos account and go to the [Activity](https://account.paxos.com/wallet/activity) tab. Download the CSV activity file and check the ID column. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectCryptoDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferPublicRejectCryptoDepositBody** | [**TransferPublicRejectCryptoDepositBody**](TransferPublicRejectCryptoDepositBody.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCryptoDeposit

> map[string]interface{} UpdateCryptoDeposit(ctx, id).TransferPublicUpdateCryptoDepositBody(transferPublicUpdateCryptoDepositBody).Execute()

Update Crypto Deposit



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
	id := "id_example" // string | Paxos ID of the crypto deposit to update. To retrieve the ID, log in to your Paxos account and go to the [Activity](https://account.paxos.com/wallet/activity) tab. Download the CSV activity file and check the ID column.
	transferPublicUpdateCryptoDepositBody := *openapiclient.NewTransferPublicUpdateCryptoDepositBody("IdentityId_example", *openapiclient.NewAddressInfo()) // TransferPublicUpdateCryptoDepositBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoDepositsAPI.UpdateCryptoDeposit(context.Background(), id).TransferPublicUpdateCryptoDepositBody(transferPublicUpdateCryptoDepositBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoDepositsAPI.UpdateCryptoDeposit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCryptoDeposit`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CryptoDepositsAPI.UpdateCryptoDeposit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Paxos ID of the crypto deposit to update. To retrieve the ID, log in to your Paxos account and go to the [Activity](https://account.paxos.com/wallet/activity) tab. Download the CSV activity file and check the ID column. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCryptoDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferPublicUpdateCryptoDepositBody** | [**TransferPublicUpdateCryptoDepositBody**](TransferPublicUpdateCryptoDepositBody.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

