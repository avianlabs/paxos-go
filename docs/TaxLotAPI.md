# \TaxLotAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTaxLot**](TaxLotAPI.md#GetTaxLot) | **Get** /tax/tax-lots/{id} | Get Tax Lot
[**ListTaxLots**](TaxLotAPI.md#ListTaxLots) | **Get** /tax/tax-lots | List Tax Lots
[**UpdateTaxLot**](TaxLotAPI.md#UpdateTaxLot) | **Put** /tax/tax-lots/{id} | Update Tax Lot



## GetTaxLot

> TaxLot GetTaxLot(ctx, id).Execute()

Get Tax Lot



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
	id := "id_example" // string | The UUID of the tax lot

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxLotAPI.GetTaxLot(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxLotAPI.GetTaxLot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxLot`: TaxLot
	fmt.Fprintf(os.Stdout, "Response from `TaxLotAPI.GetTaxLot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The UUID of the tax lot | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxLotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaxLot**](TaxLot.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaxLots

> ListTaxLotsResponse ListTaxLots(ctx).TaxLotIds(taxLotIds).ProfileIds(profileIds).AccountIds(accountIds).TransactionTypes(transactionTypes).Status(status).Cryptocurrency(cryptocurrency).CapitalGainType(capitalGainType).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtEq(createdAtEq).CreatedAtGte(createdAtGte).CreatedAtGt(createdAtGt).Limit(limit).Order(order).OrderBy(orderBy).PageCursor(pageCursor).Execute()

List Tax Lots



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
	taxLotIds := []string{"Inner_example"} // []string | A list of tax lot ids. Must be provided if no profile ids nor account ids provided. (optional)
	profileIds := []string{"Inner_example"} // []string | A list of profile ids. Can not be provided together with account ids. (optional)
	accountIds := []string{"Inner_example"} // []string | A list of account ids. Can not be provided together with profile ids. (optional)
	transactionTypes := []string{"TransactionTypes_example"} // []string | A list of transaction types. Requires profile ids or account ids to be provided. (optional)
	status := "status_example" // string | Status of tax lot, OPEN or CLOSED. Requires profile ids or account ids to be provided. (optional)
	cryptocurrency := "cryptocurrency_example" // string | Crypto currency of the tax lot. Requires profile ids or account ids to be provided. (optional)
	capitalGainType := "capitalGainType_example" // string | The capital gain type, values of LONG_TERM, SHORT_TERM, EXEMPT. Requires profile ids or account ids to be provided. (optional)
	createdAtLt := time.Now() // time.Time | Include timestamps strictly less than lt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtLte := time.Now() // time.Time | Include timestamps less than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtEq := time.Now() // time.Time | Include timestamps exactly equal to eq. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtGte := time.Now() // time.Time | Include timestamps greater than or equal to lte. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	createdAtGt := time.Now() // time.Time | Include timestamps strictly greater than gt. RFC3339 format, like `2006-01-02T15:04:05Z`. (optional)
	limit := int32(56) // int32 | Number of results to return. Default value is 100. (optional)
	order := "order_example" // string | Return items in ascending (ASC) or descending (DESC) order. Defaults to DESC. (optional)
	orderBy := "orderBy_example" // string | The specific method by which the returned results will be ordered. (optional)
	pageCursor := "pageCursor_example" // string | Cursor token for fetching the next page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxLotAPI.ListTaxLots(context.Background()).TaxLotIds(taxLotIds).ProfileIds(profileIds).AccountIds(accountIds).TransactionTypes(transactionTypes).Status(status).Cryptocurrency(cryptocurrency).CapitalGainType(capitalGainType).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtEq(createdAtEq).CreatedAtGte(createdAtGte).CreatedAtGt(createdAtGt).Limit(limit).Order(order).OrderBy(orderBy).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxLotAPI.ListTaxLots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTaxLots`: ListTaxLotsResponse
	fmt.Fprintf(os.Stdout, "Response from `TaxLotAPI.ListTaxLots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTaxLotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxLotIds** | **[]string** | A list of tax lot ids. Must be provided if no profile ids nor account ids provided. | 
 **profileIds** | **[]string** | A list of profile ids. Can not be provided together with account ids. | 
 **accountIds** | **[]string** | A list of account ids. Can not be provided together with profile ids. | 
 **transactionTypes** | **[]string** | A list of transaction types. Requires profile ids or account ids to be provided. | 
 **status** | **string** | Status of tax lot, OPEN or CLOSED. Requires profile ids or account ids to be provided. | 
 **cryptocurrency** | **string** | Crypto currency of the tax lot. Requires profile ids or account ids to be provided. | 
 **capitalGainType** | **string** | The capital gain type, values of LONG_TERM, SHORT_TERM, EXEMPT. Requires profile ids or account ids to be provided. | 
 **createdAtLt** | **time.Time** | Include timestamps strictly less than lt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtLte** | **time.Time** | Include timestamps less than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtEq** | **time.Time** | Include timestamps exactly equal to eq. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtGte** | **time.Time** | Include timestamps greater than or equal to lte. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **createdAtGt** | **time.Time** | Include timestamps strictly greater than gt. RFC3339 format, like &#x60;2006-01-02T15:04:05Z&#x60;. | 
 **limit** | **int32** | Number of results to return. Default value is 100. | 
 **order** | **string** | Return items in ascending (ASC) or descending (DESC) order. Defaults to DESC. | 
 **orderBy** | **string** | The specific method by which the returned results will be ordered. | 
 **pageCursor** | **string** | Cursor token for fetching the next page. | 

### Return type

[**ListTaxLotsResponse**](ListTaxLotsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaxLot

> TaxLot UpdateTaxLot(ctx, id).UpdateTaxLotRequest(updateTaxLotRequest).Execute()

Update Tax Lot



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
	id := "id_example" // string | The UUID of the tax lot.
	updateTaxLotRequest := *openapiclient.NewUpdateTaxLotRequest(*openapiclient.NewTaxLot()) // UpdateTaxLotRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxLotAPI.UpdateTaxLot(context.Background(), id).UpdateTaxLotRequest(updateTaxLotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxLotAPI.UpdateTaxLot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTaxLot`: TaxLot
	fmt.Fprintf(os.Stdout, "Response from `TaxLotAPI.UpdateTaxLot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The UUID of the tax lot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaxLotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTaxLotRequest** | [**UpdateTaxLotRequest**](UpdateTaxLotRequest.md) |  | 

### Return type

[**TaxLot**](TaxLot.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

