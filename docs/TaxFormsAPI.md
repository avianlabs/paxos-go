# \TaxFormsAPI

All URIs are relative to *https://api.paxos.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListTaxFormRevisions**](TaxFormsAPI.md#ListTaxFormRevisions) | **Get** /tax/tax-form-revisions | List Tax Form Revisions
[**ListTaxForms**](TaxFormsAPI.md#ListTaxForms) | **Get** /tax/tax-forms | List Tax Forms



## ListTaxFormRevisions

> ListTaxFormRevisionsResponse ListTaxFormRevisions(ctx).AccountId(accountId).TaxYear(taxYear).FormTypes(formTypes).Revision(revision).Execute()

List Tax Form Revisions



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
	accountId := "accountId_example" // string | Required. Account ID
	taxYear := "taxYear_example" // string | Required. Tax Year
	formTypes := []string{"FormTypes_example"} // []string | Tax forms 1009 B or Misc (optional)
	revision := "revision_example" // string | Revision (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxFormsAPI.ListTaxFormRevisions(context.Background()).AccountId(accountId).TaxYear(taxYear).FormTypes(formTypes).Revision(revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxFormsAPI.ListTaxFormRevisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTaxFormRevisions`: ListTaxFormRevisionsResponse
	fmt.Fprintf(os.Stdout, "Response from `TaxFormsAPI.ListTaxFormRevisions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTaxFormRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Required. Account ID | 
 **taxYear** | **string** | Required. Tax Year | 
 **formTypes** | **[]string** | Tax forms 1009 B or Misc | 
 **revision** | **string** | Revision | 

### Return type

[**ListTaxFormRevisionsResponse**](ListTaxFormRevisionsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaxForms

> ListTaxFormsResponse ListTaxForms(ctx).TaxYear(taxYear).AccountIds(accountIds).FormTypes(formTypes).UsersLimit(usersLimit).OrderBy(orderBy).Order(order).PageCursor(pageCursor).Execute()

List Tax Forms



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
	taxYear := "taxYear_example" // string | Required. Tax year
	accountIds := []string{"Inner_example"} // []string | Optional. A list of Account IDs. Maximum 50. (optional)
	formTypes := []string{"FormTypes_example"} // []string | Form types (optional)
	usersLimit := int32(56) // int32 | Number of results to return. Defaults to 50. (optional)
	orderBy := "orderBy_example" // string | The specific method by which the returned results will be ordered. Defaults to ID. (optional)
	order := "order_example" // string | Return items in ascending (ASC) or descending (DESC) order. Defaults to ASC. (optional)
	pageCursor := "pageCursor_example" // string | Cursor token for fetching the next page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxFormsAPI.ListTaxForms(context.Background()).TaxYear(taxYear).AccountIds(accountIds).FormTypes(formTypes).UsersLimit(usersLimit).OrderBy(orderBy).Order(order).PageCursor(pageCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxFormsAPI.ListTaxForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTaxForms`: ListTaxFormsResponse
	fmt.Fprintf(os.Stdout, "Response from `TaxFormsAPI.ListTaxForms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTaxFormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxYear** | **string** | Required. Tax year | 
 **accountIds** | **[]string** | Optional. A list of Account IDs. Maximum 50. | 
 **formTypes** | **[]string** | Form types | 
 **usersLimit** | **int32** | Number of results to return. Defaults to 50. | 
 **orderBy** | **string** | The specific method by which the returned results will be ordered. Defaults to ID. | 
 **order** | **string** | Return items in ascending (ASC) or descending (DESC) order. Defaults to ASC. | 
 **pageCursor** | **string** | Cursor token for fetching the next page. | 

### Return type

[**ListTaxFormsResponse**](ListTaxFormsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

