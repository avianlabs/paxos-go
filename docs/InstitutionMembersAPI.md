# \InstitutionMembersAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInstitutionMembers**](InstitutionMembersAPI.md#AddInstitutionMembers) | **Post** /identity/institution-members | Add Institution Members
[**DeleteInstitutionMember**](InstitutionMembersAPI.md#DeleteInstitutionMember) | **Delete** /identity/institution-members/{id} | Remove Institution Member



## AddInstitutionMembers

> AddInstitutionMembersResponse AddInstitutionMembers(ctx).AddInstitutionMembersRequest(addInstitutionMembersRequest).Execute()

Add Institution Members



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
	addInstitutionMembersRequest := *openapiclient.NewAddInstitutionMembersRequest("InstitutionId_example", []openapiclient.InstitutionMember{*openapiclient.NewInstitutionMember()}) // AddInstitutionMembersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionMembersAPI.AddInstitutionMembers(context.Background()).AddInstitutionMembersRequest(addInstitutionMembersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionMembersAPI.AddInstitutionMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddInstitutionMembers`: AddInstitutionMembersResponse
	fmt.Fprintf(os.Stdout, "Response from `InstitutionMembersAPI.AddInstitutionMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddInstitutionMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addInstitutionMembersRequest** | [**AddInstitutionMembersRequest**](AddInstitutionMembersRequest.md) |  | 

### Return type

[**AddInstitutionMembersResponse**](AddInstitutionMembersResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstitutionMember

> map[string]interface{} DeleteInstitutionMember(ctx, id).Execute()

Remove Institution Member



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
	id := "id_example" // string | The institution member ID that should be removed from the institution.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstitutionMembersAPI.DeleteInstitutionMember(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstitutionMembersAPI.DeleteInstitutionMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInstitutionMember`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstitutionMembersAPI.DeleteInstitutionMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The institution member ID that should be removed from the institution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstitutionMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

