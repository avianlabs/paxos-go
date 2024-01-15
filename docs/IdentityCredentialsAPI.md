# \IdentityCredentialsAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetryIdVerification**](IdentityCredentialsAPI.md#RetryIdVerification) | **Post** /identity/identities/{id}/retry-id-verification | Retry Id Verification
[**SetVerifierCredentials**](IdentityCredentialsAPI.md#SetVerifierCredentials) | **Post** /identity/verifier-credentials | Set Verifier Credentials



## RetryIdVerification

> RetryIdVerificationResponse RetryIdVerification(ctx, id).Body(body).Execute()

Retry Id Verification



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
	id := "id_example" // string | id associated with the identity
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityCredentialsAPI.RetryIdVerification(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityCredentialsAPI.RetryIdVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryIdVerification`: RetryIdVerificationResponse
	fmt.Fprintf(os.Stdout, "Response from `IdentityCredentialsAPI.RetryIdVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id associated with the identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryIdVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**RetryIdVerificationResponse**](RetryIdVerificationResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVerifierCredentials

> map[string]interface{} SetVerifierCredentials(ctx).SetVerifierCredentialsRequest(setVerifierCredentialsRequest).Execute()

Set Verifier Credentials



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
	setVerifierCredentialsRequest := *openapiclient.NewSetVerifierCredentialsRequest(openapiclient.identityprotoVerifierType("JUMIO")) // SetVerifierCredentialsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityCredentialsAPI.SetVerifierCredentials(context.Background()).SetVerifierCredentialsRequest(setVerifierCredentialsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityCredentialsAPI.SetVerifierCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVerifierCredentials`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IdentityCredentialsAPI.SetVerifierCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetVerifierCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setVerifierCredentialsRequest** | [**SetVerifierCredentialsRequest**](SetVerifierCredentialsRequest.md) |  | 

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

