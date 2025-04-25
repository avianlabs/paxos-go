# \SandboxIdentityAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SandboxSetIdentityStatus**](SandboxIdentityAPI.md#SandboxSetIdentityStatus) | **Put** /identity/identities/{id}/sandbox-status | Sandbox Set Identity Status



## SandboxSetIdentityStatus

> map[string]interface{} SandboxSetIdentityStatus(ctx, id).IdentitySandboxSandboxSetIdentityStatusBody(identitySandboxSandboxSetIdentityStatusBody).Execute()

Sandbox Set Identity Status



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
	id := "id_example" // string | 
	identitySandboxSandboxSetIdentityStatusBody := *openapiclient.NewIdentitySandboxSandboxSetIdentityStatusBody() // IdentitySandboxSandboxSetIdentityStatusBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxIdentityAPI.SandboxSetIdentityStatus(context.Background(), id).IdentitySandboxSandboxSetIdentityStatusBody(identitySandboxSandboxSetIdentityStatusBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxIdentityAPI.SandboxSetIdentityStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SandboxSetIdentityStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SandboxIdentityAPI.SandboxSetIdentityStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSandboxSetIdentityStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identitySandboxSandboxSetIdentityStatusBody** | [**IdentitySandboxSandboxSetIdentityStatusBody**](IdentitySandboxSandboxSetIdentityStatusBody.md) |  | 

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

