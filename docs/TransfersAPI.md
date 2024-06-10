# \TransfersAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransfer**](TransfersAPI.md#GetTransfer) | **Get** /transfer/transfers/{id} | Get Transfer
[**ListTransfers**](TransfersAPI.md#ListTransfers) | **Get** /transfer/transfers | List Transfers



## GetTransfer

> Transfer GetTransfer(ctx, id).Execute()

Get Transfer



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
	id := "id_example" // string | The Paxos transfer UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransfersAPI.GetTransfer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransfersAPI.GetTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransfer`: Transfer
	fmt.Fprintf(os.Stdout, "Response from `TransfersAPI.GetTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Paxos transfer UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transfer**](Transfer.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransfers

> ListTransfersResponse ListTransfers(ctx).ProfileIds(profileIds).IdentityIds(identityIds).RefIds(refIds).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtEq(createdAtEq).CreatedAtGte(createdAtGte).CreatedAtGt(createdAtGt).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtEq(updatedAtEq).UpdatedAtGte(updatedAtGte).UpdatedAtGt(updatedAtGt).Limit(limit).Order(order).OrderBy(orderBy).PageCursor(pageCursor).AccountIds(accountIds).Ids(ids).Type_(type_).GroupIds(groupIds).Execute()

List Transfers



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
	profileIds := []string{"Inner_example"} // []string | Optionally filter by the target profiles of the transfers. Limit 100. Can be combined with created_at.* or updated_at.* filtering options only. (optional)
	identityIds := []string{"Inner_example"} // []string | Optionally filter by the Identities associated with the transfers. Limit 100. Can be combined with created_at.* or updated_at.* filtering options only. (optional)
	refIds := []string{"Inner_example"} // []string | The client-specified IDs provided during transfer creation. Limit 100. Can be combined with created_at.* or updated_at.* filtering options only. (optional)
	createdAtLt := time.Now() // time.Time | Include timestamps strictly less than lt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtLte := time.Now() // time.Time | Include timestamps less than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtEq := time.Now() // time.Time | Include timestamps exactly equal to eq. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtGte := time.Now() // time.Time | Include timestamps greater than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtGt := time.Now() // time.Time | Include timestamps strictly greater than gt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	updatedAtLt := time.Now() // time.Time | Include timestamps strictly less than lt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	updatedAtLte := time.Now() // time.Time | Include timestamps less than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	updatedAtEq := time.Now() // time.Time | Include timestamps exactly equal to eq. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	updatedAtGte := time.Now() // time.Time | Include timestamps greater than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	updatedAtGt := time.Now() // time.Time | Include timestamps strictly greater than gt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	limit := int32(56) // int32 | Number of results to return. Defaults to 100 if no limit is provided. Maximum 1000. (optional)
	order := "order_example" // string | Determines whether the items are returned in ascending (ASC), or descending (DESC) order. Defaults to ASC if `order_by` is UPDATED_AT. Otherwise defaults to DESC. (optional)
	orderBy := "orderBy_example" // string | The specific method by which the returned results will be ordered. When specifying both `order_by` and a timestamp filter, `order_by` must equal the timestamp filter type. For example:  `updated_at.gte = 2022-07-01T03:02:01Z`, `order_by = UPDATED_AT`  Using `updated_at.gte = 2022-07-01T03:02:01Z`, `order_by = CREATED_AT` is an invalid pairing. (optional)
	pageCursor := "pageCursor_example" // string | Optional: Cursor for getting the next page of results. When the number of items returned is fewer than the limit, there is currently no next page. (optional)
	accountIds := []string{"Inner_example"} // []string | Optionally filter by the Accounts associated with the transfers. Limit 100. Can be combined with created_at.* or updated_at.* filtering options only. (optional)
	ids := []string{"Inner_example"} // []string | Optionally filter by the transfer ids. Limit 100. Can be combined with created_at.* or updated_at.* filtering options only. (optional)
	type_ := []string{"Type_example"} // []string | Retrieve all (default) or the specified transfers. Optionally filter by transfer `type`. To retrieve multiple transfer types, use query parameters: e.g., `type=CRYPTO_DEPOSIT&type=CRYPTO_WITHDRAWAL`. Can be combined with `created_at.*` or `updated_at.*` filtering options only. (optional)
	groupIds := []string{"Inner_example"} // []string | Optionally filter by transfer `group_ids`. Limit 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransfersAPI.ListTransfers(context.Background()).ProfileIds(profileIds).IdentityIds(identityIds).RefIds(refIds).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtEq(createdAtEq).CreatedAtGte(createdAtGte).CreatedAtGt(createdAtGt).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtEq(updatedAtEq).UpdatedAtGte(updatedAtGte).UpdatedAtGt(updatedAtGt).Limit(limit).Order(order).OrderBy(orderBy).PageCursor(pageCursor).AccountIds(accountIds).Ids(ids).Type_(type_).GroupIds(groupIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransfersAPI.ListTransfers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransfers`: ListTransfersResponse
	fmt.Fprintf(os.Stdout, "Response from `TransfersAPI.ListTransfers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileIds** | **[]string** | Optionally filter by the target profiles of the transfers. Limit 100. Can be combined with created_at.* or updated_at.* filtering options only. | 
 **identityIds** | **[]string** | Optionally filter by the Identities associated with the transfers. Limit 100. Can be combined with created_at.* or updated_at.* filtering options only. | 
 **refIds** | **[]string** | The client-specified IDs provided during transfer creation. Limit 100. Can be combined with created_at.* or updated_at.* filtering options only. | 
 **createdAtLt** | **time.Time** | Include timestamps strictly less than lt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtLte** | **time.Time** | Include timestamps less than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtEq** | **time.Time** | Include timestamps exactly equal to eq. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtGte** | **time.Time** | Include timestamps greater than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtGt** | **time.Time** | Include timestamps strictly greater than gt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **updatedAtLt** | **time.Time** | Include timestamps strictly less than lt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **updatedAtLte** | **time.Time** | Include timestamps less than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **updatedAtEq** | **time.Time** | Include timestamps exactly equal to eq. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **updatedAtGte** | **time.Time** | Include timestamps greater than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **updatedAtGt** | **time.Time** | Include timestamps strictly greater than gt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **limit** | **int32** | Number of results to return. Defaults to 100 if no limit is provided. Maximum 1000. | 
 **order** | **string** | Determines whether the items are returned in ascending (ASC), or descending (DESC) order. Defaults to ASC if &#x60;order_by&#x60; is UPDATED_AT. Otherwise defaults to DESC. | 
 **orderBy** | **string** | The specific method by which the returned results will be ordered. When specifying both &#x60;order_by&#x60; and a timestamp filter, &#x60;order_by&#x60; must equal the timestamp filter type. For example:  &#x60;updated_at.gte &#x3D; 2022-07-01T03:02:01Z&#x60;, &#x60;order_by &#x3D; UPDATED_AT&#x60;  Using &#x60;updated_at.gte &#x3D; 2022-07-01T03:02:01Z&#x60;, &#x60;order_by &#x3D; CREATED_AT&#x60; is an invalid pairing. | 
 **pageCursor** | **string** | Optional: Cursor for getting the next page of results. When the number of items returned is fewer than the limit, there is currently no next page. | 
 **accountIds** | **[]string** | Optionally filter by the Accounts associated with the transfers. Limit 100. Can be combined with created_at.* or updated_at.* filtering options only. | 
 **ids** | **[]string** | Optionally filter by the transfer ids. Limit 100. Can be combined with created_at.* or updated_at.* filtering options only. | 
 **type_** | **[]string** | Retrieve all (default) or the specified transfers. Optionally filter by transfer &#x60;type&#x60;. To retrieve multiple transfer types, use query parameters: e.g., &#x60;type&#x3D;CRYPTO_DEPOSIT&amp;type&#x3D;CRYPTO_WITHDRAWAL&#x60;. Can be combined with &#x60;created_at.*&#x60; or &#x60;updated_at.*&#x60; filtering options only. | 
 **groupIds** | **[]string** | Optionally filter by transfer &#x60;group_ids&#x60;. Limit 100. | 

### Return type

[**ListTransfersResponse**](ListTransfersResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

