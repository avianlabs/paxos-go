# \AccountsAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](AccountsAPI.md#CreateAccount) | **Post** /identity/accounts | Create Account
[**GetAccount**](AccountsAPI.md#GetAccount) | **Get** /identity/accounts/{id} | 
[**ListAccounts**](AccountsAPI.md#ListAccounts) | **Get** /identity/accounts | List Accounts
[**UpdateAccount**](AccountsAPI.md#UpdateAccount) | **Put** /identity/accounts | Update Account



## CreateAccount

> Account CreateAccount(ctx).CreateAccountRequest(createAccountRequest).Execute()

Create Account



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
	createAccountRequest := *openapiclient.NewCreateAccountRequest(*openapiclient.NewAccount()) // CreateAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.CreateAccount(context.Background()).CreateAccountRequest(createAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.CreateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountRequest** | [**CreateAccountRequest**](CreateAccountRequest.md) |  | 

### Return type

[**Account**](Account.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Account GetAccount(ctx, id).Execute()



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
	id := "id_example" // string | uuid id of account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | uuid id of account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> ListAccountsResponse ListAccounts(ctx).PageCursor(pageCursor).Order(order).OrderBy(orderBy).Limit(limit).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtEq(createdAtEq).CreatedAtGte(createdAtGte).CreatedAtGt(createdAtGt).IdentityId(identityId).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtEq(updatedAtEq).UpdatedAtGte(updatedAtGte).UpdatedAtGt(updatedAtGt).Execute()

List Accounts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/avianlabs/paxos-go"
)

func main() {
	pageCursor := "pageCursor_example" // string |  (optional)
	order := "order_example" // string | Return items in ascending (ASC) or descending (DESC) order. Defaults to DESC. (optional)
	orderBy := "orderBy_example" // string | The specific method by which the returned results will be ordered. (optional)
	limit := int32(56) // int32 | Number of results to return. (optional)
	createdAtLt := time.Now() // time.Time | Include timestamps strictly less than lt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtLte := time.Now() // time.Time | Include timestamps less than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtEq := time.Now() // time.Time | Include timestamps exactly equal to eq. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtGte := time.Now() // time.Time | Include timestamps greater than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtGt := time.Now() // time.Time | Include timestamps strictly greater than gt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	identityId := "identityId_example" // string | Optionally filter by primary identity identity. (optional)
	updatedAtLt := time.Now() // time.Time | Include timestamps strictly less than lt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	updatedAtLte := time.Now() // time.Time | Include timestamps less than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	updatedAtEq := time.Now() // time.Time | Include timestamps exactly equal to eq. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	updatedAtGte := time.Now() // time.Time | Include timestamps greater than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	updatedAtGt := time.Now() // time.Time | Include timestamps strictly greater than gt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ListAccounts(context.Background()).PageCursor(pageCursor).Order(order).OrderBy(orderBy).Limit(limit).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtEq(createdAtEq).CreatedAtGte(createdAtGte).CreatedAtGt(createdAtGt).IdentityId(identityId).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtEq(updatedAtEq).UpdatedAtGte(updatedAtGte).UpdatedAtGt(updatedAtGt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccounts`: ListAccountsResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageCursor** | **string** |  | 
 **order** | **string** | Return items in ascending (ASC) or descending (DESC) order. Defaults to DESC. | 
 **orderBy** | **string** | The specific method by which the returned results will be ordered. | 
 **limit** | **int32** | Number of results to return. | 
 **createdAtLt** | **time.Time** | Include timestamps strictly less than lt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtLte** | **time.Time** | Include timestamps less than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtEq** | **time.Time** | Include timestamps exactly equal to eq. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtGte** | **time.Time** | Include timestamps greater than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtGt** | **time.Time** | Include timestamps strictly greater than gt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **identityId** | **string** | Optionally filter by primary identity identity. | 
 **updatedAtLt** | **time.Time** | Include timestamps strictly less than lt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **updatedAtLte** | **time.Time** | Include timestamps less than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **updatedAtEq** | **time.Time** | Include timestamps exactly equal to eq. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **updatedAtGte** | **time.Time** | Include timestamps greater than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **updatedAtGt** | **time.Time** | Include timestamps strictly greater than gt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 

### Return type

[**ListAccountsResponse**](ListAccountsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> Account UpdateAccount(ctx).UpdateAccountRequest(updateAccountRequest).Execute()

Update Account



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
	updateAccountRequest := *openapiclient.NewUpdateAccountRequest(*openapiclient.NewAccount()) // UpdateAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.UpdateAccount(context.Background()).UpdateAccountRequest(updateAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.UpdateAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAccountRequest** | [**UpdateAccountRequest**](UpdateAccountRequest.md) |  | 

### Return type

[**Account**](Account.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

