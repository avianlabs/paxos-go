# \StablecoinConversionAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelStablecoinConversion**](StablecoinConversionAPI.md#CancelStablecoinConversion) | **Delete** /conversion/stablecoins/{id} | Cancel Stablecoin Conversion
[**CreateStablecoinConversion**](StablecoinConversionAPI.md#CreateStablecoinConversion) | **Post** /conversion/stablecoins | Create Stablecoin Conversion
[**GetStablecoinConversion**](StablecoinConversionAPI.md#GetStablecoinConversion) | **Get** /conversion/stablecoins/{id} | Get Stablecoin Conversion
[**ListStablecoinConversions**](StablecoinConversionAPI.md#ListStablecoinConversions) | **Get** /conversion/stablecoins | List Stablecoin Conversions



## CancelStablecoinConversion

> StablecoinConversion CancelStablecoinConversion(ctx, id).Execute()

Cancel Stablecoin Conversion



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
	id := "id_example" // string | System provided UUID for the conversion is provided in the [Create Stablecoin Conversion](#operation/CreateStablecoinConversion) response.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StablecoinConversionAPI.CancelStablecoinConversion(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StablecoinConversionAPI.CancelStablecoinConversion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelStablecoinConversion`: StablecoinConversion
	fmt.Fprintf(os.Stdout, "Response from `StablecoinConversionAPI.CancelStablecoinConversion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | System provided UUID for the conversion is provided in the [Create Stablecoin Conversion](#operation/CreateStablecoinConversion) response. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelStablecoinConversionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StablecoinConversion**](StablecoinConversion.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStablecoinConversion

> StablecoinConversion CreateStablecoinConversion(ctx).CreateStablecoinConversionRequest(createStablecoinConversionRequest).Execute()

Create Stablecoin Conversion



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
	createStablecoinConversionRequest := *openapiclient.NewCreateStablecoinConversionRequest("ProfileId_example", "Amount_example", "SourceAsset_example", "TargetAsset_example") // CreateStablecoinConversionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StablecoinConversionAPI.CreateStablecoinConversion(context.Background()).CreateStablecoinConversionRequest(createStablecoinConversionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StablecoinConversionAPI.CreateStablecoinConversion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStablecoinConversion`: StablecoinConversion
	fmt.Fprintf(os.Stdout, "Response from `StablecoinConversionAPI.CreateStablecoinConversion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStablecoinConversionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStablecoinConversionRequest** | [**CreateStablecoinConversionRequest**](CreateStablecoinConversionRequest.md) |  | 

### Return type

[**StablecoinConversion**](StablecoinConversion.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStablecoinConversion

> StablecoinConversion GetStablecoinConversion(ctx, id).Execute()

Get Stablecoin Conversion



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
	id := "id_example" // string | System provided UUID for the conversion is provided in the [Create Stablecoin Conversion](#operation/CreateStablecoinConversion) response.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StablecoinConversionAPI.GetStablecoinConversion(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StablecoinConversionAPI.GetStablecoinConversion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStablecoinConversion`: StablecoinConversion
	fmt.Fprintf(os.Stdout, "Response from `StablecoinConversionAPI.GetStablecoinConversion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | System provided UUID for the conversion is provided in the [Create Stablecoin Conversion](#operation/CreateStablecoinConversion) response. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStablecoinConversionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StablecoinConversion**](StablecoinConversion.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStablecoinConversions

> ListStablecoinConversionsResponse ListStablecoinConversions(ctx).ProfileId(profileId).RefId(refId).CreatedAtBegin(createdAtBegin).CreatedAtEnd(createdAtEnd).UpdatedAtBegin(updatedAtBegin).UpdatedAtEnd(updatedAtEnd).Order(order).PageCursor(pageCursor).Limit(limit).Execute()

List Stablecoin Conversions



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
	profileId := "profileId_example" // string | The <a href=\"#tag/Profiles\">Profile</a> associated with a conversion. Required in the <a href=\"#operation/CreateStablecoinConversion\">Create Stablecoin Conversion</a> request. (optional)
	refId := "refId_example" // string | Client provided, unique Reference ID included the <a href=\"#operation/CreateStablecoinConversion\">Create Stablecoin Conversion</a> request. (optional)
	createdAtBegin := time.Now() // time.Time | Only return records after this timestamp, inclusive. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtEnd := time.Now() // time.Time | Only return records before this timestamp, inclusive. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	updatedAtBegin := time.Now() // time.Time | Only return records after this timestamp, inclusive. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	updatedAtEnd := time.Now() // time.Time | Only return records before this timestamp, inclusive. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	order := "order_example" // string | Return items in ascending (ASC) or descending (DESC) order by `created_at` timestamp.  Default order is descending (DESC). (optional)
	pageCursor := "pageCursor_example" // string | Cursor token for fetching the next page. (optional)
	limit := int32(56) // int32 | Number of results to return.  Default is 100 items. Maximum is 1000 items. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StablecoinConversionAPI.ListStablecoinConversions(context.Background()).ProfileId(profileId).RefId(refId).CreatedAtBegin(createdAtBegin).CreatedAtEnd(createdAtEnd).UpdatedAtBegin(updatedAtBegin).UpdatedAtEnd(updatedAtEnd).Order(order).PageCursor(pageCursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StablecoinConversionAPI.ListStablecoinConversions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStablecoinConversions`: ListStablecoinConversionsResponse
	fmt.Fprintf(os.Stdout, "Response from `StablecoinConversionAPI.ListStablecoinConversions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStablecoinConversionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | The &lt;a href&#x3D;\&quot;#tag/Profiles\&quot;&gt;Profile&lt;/a&gt; associated with a conversion. Required in the &lt;a href&#x3D;\&quot;#operation/CreateStablecoinConversion\&quot;&gt;Create Stablecoin Conversion&lt;/a&gt; request. | 
 **refId** | **string** | Client provided, unique Reference ID included the &lt;a href&#x3D;\&quot;#operation/CreateStablecoinConversion\&quot;&gt;Create Stablecoin Conversion&lt;/a&gt; request. | 
 **createdAtBegin** | **time.Time** | Only return records after this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtEnd** | **time.Time** | Only return records before this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **updatedAtBegin** | **time.Time** | Only return records after this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **updatedAtEnd** | **time.Time** | Only return records before this timestamp, inclusive. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **order** | **string** | Return items in ascending (ASC) or descending (DESC) order by &#x60;created_at&#x60; timestamp.  Default order is descending (DESC). | 
 **pageCursor** | **string** | Cursor token for fetching the next page. | 
 **limit** | **int32** | Number of results to return.  Default is 100 items. Maximum is 1000 items. | 

### Return type

[**ListStablecoinConversionsResponse**](ListStablecoinConversionsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

