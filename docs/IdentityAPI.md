# \IdentityAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentity**](IdentityAPI.md#CreateIdentity) | **Post** /identity/identities | Create Identity
[**GetIdentity**](IdentityAPI.md#GetIdentity) | **Get** /identity/identities/{id} | Get Identity
[**ListIdentities**](IdentityAPI.md#ListIdentities) | **Get** /identity/identities | List Identities
[**UpdateIdentity**](IdentityAPI.md#UpdateIdentity) | **Put** /identity/identities/{id} | Update Identity



## CreateIdentity

> Identity CreateIdentity(ctx).CreateIdentityRequest(createIdentityRequest).Execute()

Create Identity



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
	createIdentityRequest := *openapiclient.NewCreateIdentityRequest() // CreateIdentityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAPI.CreateIdentity(context.Background()).CreateIdentityRequest(createIdentityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPI.CreateIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIdentity`: Identity
	fmt.Fprintf(os.Stdout, "Response from `IdentityAPI.CreateIdentity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIdentityRequest** | [**CreateIdentityRequest**](CreateIdentityRequest.md) |  | 

### Return type

[**Identity**](Identity.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentity

> Identity GetIdentity(ctx, id).IncludeDetails(includeDetails).IncludeInstitutionMembers(includeInstitutionMembers).Execute()

Get Identity



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
	includeDetails := true // bool | query param; details are encrypted, so we do not want to include them by default (optional)
	includeInstitutionMembers := true // bool | query param; to include institution members for institution identity (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAPI.GetIdentity(context.Background(), id).IncludeDetails(includeDetails).IncludeInstitutionMembers(includeInstitutionMembers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPI.GetIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentity`: Identity
	fmt.Fprintf(os.Stdout, "Response from `IdentityAPI.GetIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id associated with the identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDetails** | **bool** | query param; details are encrypted, so we do not want to include them by default | 
 **includeInstitutionMembers** | **bool** | query param; to include institution members for institution identity | 

### Return type

[**Identity**](Identity.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentities

> ListIdentitiesResponse ListIdentities(ctx).SummaryStatus(summaryStatus).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtEq(createdAtEq).CreatedAtGte(createdAtGte).CreatedAtGt(createdAtGt).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtEq(updatedAtEq).UpdatedAtGte(updatedAtGte).UpdatedAtGt(updatedAtGt).Limit(limit).Order(order).OrderBy(orderBy).PageCursor(pageCursor).IdentityType(identityType).Execute()

List Identities



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
	summaryStatus := "summaryStatus_example" // string | Summary Status of the Identity. (optional)
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
	limit := int32(56) // int32 | Number of results to return. (optional)
	order := "order_example" // string | Return items in ascending (ASC) or descending (DESC) order. Defaults to DESC. (optional)
	orderBy := "orderBy_example" // string | The specific method by which the returned results will be ordered. (optional)
	pageCursor := "pageCursor_example" // string | Cursor token for fetching the next page. (optional)
	identityType := "identityType_example" // string | Optionally filter by Identity type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAPI.ListIdentities(context.Background()).SummaryStatus(summaryStatus).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtEq(createdAtEq).CreatedAtGte(createdAtGte).CreatedAtGt(createdAtGt).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtEq(updatedAtEq).UpdatedAtGte(updatedAtGte).UpdatedAtGt(updatedAtGt).Limit(limit).Order(order).OrderBy(orderBy).PageCursor(pageCursor).IdentityType(identityType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPI.ListIdentities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdentities`: ListIdentitiesResponse
	fmt.Fprintf(os.Stdout, "Response from `IdentityAPI.ListIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **summaryStatus** | **string** | Summary Status of the Identity. | 
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
 **limit** | **int32** | Number of results to return. | 
 **order** | **string** | Return items in ascending (ASC) or descending (DESC) order. Defaults to DESC. | 
 **orderBy** | **string** | The specific method by which the returned results will be ordered. | 
 **pageCursor** | **string** | Cursor token for fetching the next page. | 
 **identityType** | **string** | Optionally filter by Identity type | 

### Return type

[**ListIdentitiesResponse**](ListIdentitiesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentity

> Identity UpdateIdentity(ctx, id).UpdateIdentityRequest(updateIdentityRequest).Execute()

Update Identity



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
	updateIdentityRequest := *openapiclient.NewUpdateIdentityRequest() // UpdateIdentityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityAPI.UpdateIdentity(context.Background(), id).UpdateIdentityRequest(updateIdentityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityAPI.UpdateIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIdentity`: Identity
	fmt.Fprintf(os.Stdout, "Response from `IdentityAPI.UpdateIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIdentityRequest** | [**UpdateIdentityRequest**](UpdateIdentityRequest.md) |  | 

### Return type

[**Identity**](Identity.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

