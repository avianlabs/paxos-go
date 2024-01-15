# \IdentityDocumentsAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocumentUpload**](IdentityDocumentsAPI.md#DocumentUpload) | **Put** /identity/identities/{identity_id}/documents | Document Upload
[**ListIdentityDocuments**](IdentityDocumentsAPI.md#ListIdentityDocuments) | **Get** /identity/identities/{identity_id}/documents | List Identity Documents



## DocumentUpload

> DocumentUploadResponse DocumentUpload(ctx, identityId).DocumentUploadRequest(documentUploadRequest).Execute()

Document Upload



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
	identityId := "identityId_example" // string | The id of the identity the document is associated with.
	documentUploadRequest := *openapiclient.NewDocumentUploadRequest("Name_example") // DocumentUploadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityDocumentsAPI.DocumentUpload(context.Background(), identityId).DocumentUploadRequest(documentUploadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityDocumentsAPI.DocumentUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocumentUpload`: DocumentUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `IdentityDocumentsAPI.DocumentUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | The id of the identity the document is associated with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocumentUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **documentUploadRequest** | [**DocumentUploadRequest**](DocumentUploadRequest.md) |  | 

### Return type

[**DocumentUploadResponse**](DocumentUploadResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentityDocuments

> ListIdentityDocumentsResponse ListIdentityDocuments(ctx, identityId).IncludePendingDocs(includePendingDocs).Execute()

List Identity Documents



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
	identityId := "identityId_example" // string | id associated with the identity
	includePendingDocs := true // bool | Add a list of pending document types for the identity (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityDocumentsAPI.ListIdentityDocuments(context.Background(), identityId).IncludePendingDocs(includePendingDocs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityDocumentsAPI.ListIdentityDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdentityDocuments`: ListIdentityDocumentsResponse
	fmt.Fprintf(os.Stdout, "Response from `IdentityDocumentsAPI.ListIdentityDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | id associated with the identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIdentityDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includePendingDocs** | **bool** | Add a list of pending document types for the identity | 

### Return type

[**ListIdentityDocumentsResponse**](ListIdentityDocumentsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

