# \DepositAddressesAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDepositAddress**](DepositAddressesAPI.md#CreateDepositAddress) | **Post** /transfer/deposit-addresses | Create Deposit Address
[**ListDepositAddresses**](DepositAddressesAPI.md#ListDepositAddresses) | **Get** /transfer/deposit-addresses | List Deposit Addresses



## CreateDepositAddress

> DepositAddress CreateDepositAddress(ctx).CreateDepositAddressRequest(createDepositAddressRequest).Execute()

Create Deposit Address



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
	createDepositAddressRequest := *openapiclient.NewCreateDepositAddressRequest("ProfileId_example", openapiclient.CryptoNetwork("BITCOIN")) // CreateDepositAddressRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepositAddressesAPI.CreateDepositAddress(context.Background()).CreateDepositAddressRequest(createDepositAddressRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepositAddressesAPI.CreateDepositAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDepositAddress`: DepositAddress
	fmt.Fprintf(os.Stdout, "Response from `DepositAddressesAPI.CreateDepositAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDepositAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDepositAddressRequest** | [**CreateDepositAddressRequest**](CreateDepositAddressRequest.md) |  | 

### Return type

[**DepositAddress**](DepositAddress.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDepositAddresses

> ListDepositAddressesResponse ListDepositAddresses(ctx).ProfileIds(profileIds).IdentityIds(identityIds).Ids(ids).RefIds(refIds).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtEq(createdAtEq).CreatedAtGte(createdAtGte).CreatedAtGt(createdAtGt).Limit(limit).Order(order).OrderBy(orderBy).PageCursor(pageCursor).AccountIds(accountIds).Execute()

List Deposit Addresses



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	profileIds := []string{"Inner_example"} // []string | Optionally filter by the profile IDs associated with the deposit addresses. Limit 100. (optional)
	identityIds := []string{"Inner_example"} // []string | Optionally filter by the Identities associated with the deposit addresses. Limit 100. (optional)
	ids := []string{"Inner_example"} // []string | Optionally filter by the UUIDs of the deposit addresses. Limit 100. (optional)
	refIds := []string{"Inner_example"} // []string | The client-specified IDs provided during transfer creation. Limit 100. (optional)
	createdAtLt := time.Now() // time.Time | Include timestamps strictly less than lt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtLte := time.Now() // time.Time | Include timestamps less than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtEq := time.Now() // time.Time | Include timestamps exactly equal to eq. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtGte := time.Now() // time.Time | Include timestamps greater than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtGt := time.Now() // time.Time | Include timestamps strictly greater than gt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	limit := int32(56) // int32 | Number of results to return. Defaults to 100 if no limit is provided. Maximum 1000. (optional)
	order := "order_example" // string | Return items in ascending (ASC) or descending (DESC) order. Defaults to DESC. (optional)
	orderBy := "orderBy_example" // string | The specific method by which the returned results will be ordered. (optional)
	pageCursor := "pageCursor_example" // string | Optional: Cursor for getting the next page of results. (optional)
	accountIds := []string{"Inner_example"} // []string | Optionally filter by the Accounts associated with the deposit address. Limit 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepositAddressesAPI.ListDepositAddresses(context.Background()).ProfileIds(profileIds).IdentityIds(identityIds).Ids(ids).RefIds(refIds).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtEq(createdAtEq).CreatedAtGte(createdAtGte).CreatedAtGt(createdAtGt).Limit(limit).Order(order).OrderBy(orderBy).PageCursor(pageCursor).AccountIds(accountIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepositAddressesAPI.ListDepositAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDepositAddresses`: ListDepositAddressesResponse
	fmt.Fprintf(os.Stdout, "Response from `DepositAddressesAPI.ListDepositAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDepositAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileIds** | **[]string** | Optionally filter by the profile IDs associated with the deposit addresses. Limit 100. | 
 **identityIds** | **[]string** | Optionally filter by the Identities associated with the deposit addresses. Limit 100. | 
 **ids** | **[]string** | Optionally filter by the UUIDs of the deposit addresses. Limit 100. | 
 **refIds** | **[]string** | The client-specified IDs provided during transfer creation. Limit 100. | 
 **createdAtLt** | **time.Time** | Include timestamps strictly less than lt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtLte** | **time.Time** | Include timestamps less than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtEq** | **time.Time** | Include timestamps exactly equal to eq. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtGte** | **time.Time** | Include timestamps greater than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtGt** | **time.Time** | Include timestamps strictly greater than gt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **limit** | **int32** | Number of results to return. Defaults to 100 if no limit is provided. Maximum 1000. | 
 **order** | **string** | Return items in ascending (ASC) or descending (DESC) order. Defaults to DESC. | 
 **orderBy** | **string** | The specific method by which the returned results will be ordered. | 
 **pageCursor** | **string** | Optional: Cursor for getting the next page of results. | 
 **accountIds** | **[]string** | Optionally filter by the Accounts associated with the deposit address. Limit 100. | 

### Return type

[**ListDepositAddressesResponse**](ListDepositAddressesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

