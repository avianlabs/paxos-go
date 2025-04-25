# \ProfilesAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProfile**](ProfilesAPI.md#CreateProfile) | **Post** /profiles | Create Profile
[**DeactivateProfile**](ProfilesAPI.md#DeactivateProfile) | **Put** /profiles/{profile_id}/deactivate | Deactivate Profile
[**GetProfile**](ProfilesAPI.md#GetProfile) | **Get** /profiles/{profile_id} | Get Profile
[**GetProfileBalance**](ProfilesAPI.md#GetProfileBalance) | **Get** /profiles/{profile_id}/balances/{asset} | Get Profile Balance
[**ListProfileBalances**](ProfilesAPI.md#ListProfileBalances) | **Get** /profiles/{profile_id}/balances | List Profile Balances
[**ListProfiles**](ProfilesAPI.md#ListProfiles) | **Get** /profiles | List Profiles
[**UpdateProfile**](ProfilesAPI.md#UpdateProfile) | **Put** /profiles/{profile_id} | Update Profile



## CreateProfile

> Profile CreateProfile(ctx).CreateProfileRequest(createProfileRequest).Execute()

Create Profile



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
	createProfileRequest := *openapiclient.NewCreateProfileRequest("Nickname_example") // CreateProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.CreateProfile(context.Background()).CreateProfileRequest(createProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProfile`: Profile
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProfileRequest** | [**CreateProfileRequest**](CreateProfileRequest.md) |  | 

### Return type

[**Profile**](Profile.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateProfile

> map[string]interface{} DeactivateProfile(ctx, profileId).Execute()

Deactivate Profile



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
	profileId := "profileId_example" // string | The UUID of the profile. The default profile cannot be deactivated. The profile must have a zero balance to be deactivated.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.DeactivateProfile(context.Background(), profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeactivateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateProfile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.DeactivateProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The UUID of the profile. The default profile cannot be deactivated. The profile must have a zero balance to be deactivated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateProfileRequest struct via the builder pattern


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


## GetProfile

> Profile GetProfile(ctx, profileId).IncludeDeactivated(includeDeactivated).Execute()

Get Profile



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
	profileId := "profileId_example" // string | The UUID of the profile, or \"default\" for the default profile.
	includeDeactivated := true // bool | Used to include deactivated profiles in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetProfile(context.Background(), profileId).IncludeDeactivated(includeDeactivated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfile`: Profile
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The UUID of the profile, or \&quot;default\&quot; for the default profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDeactivated** | **bool** | Used to include deactivated profiles in the response. | 

### Return type

[**Profile**](Profile.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileBalance

> ProfileBalance GetProfileBalance(ctx, profileId, asset).Execute()

Get Profile Balance



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
	profileId := "profileId_example" // string | 
	asset := "asset_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.GetProfileBalance(context.Background(), profileId, asset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetProfileBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfileBalance`: ProfileBalance
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetProfileBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 
**asset** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProfileBalance**](ProfileBalance.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProfileBalances

> ListProfileBalancesResponse ListProfileBalances(ctx, profileId).Assets(assets).Execute()

List Profile Balances



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
	profileId := "profileId_example" // string | 
	assets := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ListProfileBalances(context.Background(), profileId).Assets(assets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ListProfileBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProfileBalances`: ListProfileBalancesResponse
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ListProfileBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProfileBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assets** | **[]string** |  | 

### Return type

[**ListProfileBalancesResponse**](ListProfileBalancesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProfiles

> ListProfilesResponse ListProfiles(ctx).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtEq(createdAtEq).CreatedAtGte(createdAtGte).CreatedAtGt(createdAtGt).Limit(limit).Order(order).OrderBy(orderBy).PageCursor(pageCursor).Nickname(nickname).Execute()

List Profiles



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
	createdAtLt := time.Now() // time.Time | Include timestamps strictly less than lt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtLte := time.Now() // time.Time | Include timestamps less than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtEq := time.Now() // time.Time | Include timestamps exactly equal to eq. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtGte := time.Now() // time.Time | Include timestamps greater than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtGt := time.Now() // time.Time | Include timestamps strictly greater than gt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	limit := int32(56) // int32 | Number of results to return. (optional)
	order := "order_example" // string | Return items in ascending (ASC) or descending (DESC) order. Defaults to ASC. (optional)
	orderBy := "orderBy_example" // string | The specific method by which the returned results will be ordered. (optional)
	pageCursor := "pageCursor_example" // string | Cursor token for fetching the next page. (optional)
	nickname := "nickname_example" // string | Optionally filter by Profile display name. Retrieves nickname(s) based on the beginning characters of the given display name (prefix matching). Case insensitive. WIldcards and regular expressions not supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.ListProfiles(context.Background()).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtEq(createdAtEq).CreatedAtGte(createdAtGte).CreatedAtGt(createdAtGt).Limit(limit).Order(order).OrderBy(orderBy).PageCursor(pageCursor).Nickname(nickname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ListProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProfiles`: ListProfilesResponse
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ListProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAtLt** | **time.Time** | Include timestamps strictly less than lt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtLte** | **time.Time** | Include timestamps less than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtEq** | **time.Time** | Include timestamps exactly equal to eq. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtGte** | **time.Time** | Include timestamps greater than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtGt** | **time.Time** | Include timestamps strictly greater than gt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **limit** | **int32** | Number of results to return. | 
 **order** | **string** | Return items in ascending (ASC) or descending (DESC) order. Defaults to ASC. | 
 **orderBy** | **string** | The specific method by which the returned results will be ordered. | 
 **pageCursor** | **string** | Cursor token for fetching the next page. | 
 **nickname** | **string** | Optionally filter by Profile display name. Retrieves nickname(s) based on the beginning characters of the given display name (prefix matching). Case insensitive. WIldcards and regular expressions not supported. | 

### Return type

[**ListProfilesResponse**](ListProfilesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfile

> Profile UpdateProfile(ctx, profileId).ProfilePublicUpdateProfileBody(profilePublicUpdateProfileBody).Execute()

Update Profile



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
	profileId := "profileId_example" // string | 
	profilePublicUpdateProfileBody := *openapiclient.NewProfilePublicUpdateProfileBody("Nickname_example") // ProfilePublicUpdateProfileBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfilesAPI.UpdateProfile(context.Background(), profileId).ProfilePublicUpdateProfileBody(profilePublicUpdateProfileBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.UpdateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProfile`: Profile
	fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.UpdateProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profilePublicUpdateProfileBody** | [**ProfilePublicUpdateProfileBody**](ProfilePublicUpdateProfileBody.md) |  | 

### Return type

[**Profile**](Profile.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

