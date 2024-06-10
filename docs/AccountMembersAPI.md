# \AccountMembersAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccountMembers**](AccountMembersAPI.md#AddAccountMembers) | **Post** /identity/account-members | Add Account Members
[**DeleteAccountMember**](AccountMembersAPI.md#DeleteAccountMember) | **Delete** /identity/account-members/{id} | Remove Account Member



## AddAccountMembers

> AddAccountMembersResponse AddAccountMembers(ctx).AddAccountMembersRequest(addAccountMembersRequest).Execute()

Add Account Members



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
	addAccountMembersRequest := *openapiclient.NewAddAccountMembersRequest("AccountId_example", []openapiclient.AccountMember{*openapiclient.NewAccountMember()}) // AddAccountMembersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountMembersAPI.AddAccountMembers(context.Background()).AddAccountMembersRequest(addAccountMembersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountMembersAPI.AddAccountMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAccountMembers`: AddAccountMembersResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountMembersAPI.AddAccountMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAccountMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addAccountMembersRequest** | [**AddAccountMembersRequest**](AddAccountMembersRequest.md) |  | 

### Return type

[**AddAccountMembersResponse**](AddAccountMembersResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountMember

> map[string]interface{} DeleteAccountMember(ctx, id).Execute()

Remove Account Member



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
	id := "id_example" // string | The account member ID that should be removed from the account.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountMembersAPI.DeleteAccountMember(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountMembersAPI.DeleteAccountMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccountMember`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountMembersAPI.DeleteAccountMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The account member ID that should be removed from the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountMemberRequest struct via the builder pattern


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

