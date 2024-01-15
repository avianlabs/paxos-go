# ListTaxFormsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxFormUrls** | Pointer to [**[]TaxFormURL**](TaxFormURL.md) | List of tax form URLs. The size shall not exceed &#x60;users_limit&#x60; times &#x60;form_types&#x60;. | [optional] 
**NextPageCursor** | Pointer to **string** | Cursor token required for fetching the next page. | [optional] 

## Methods

### NewListTaxFormsResponse

`func NewListTaxFormsResponse() *ListTaxFormsResponse`

NewListTaxFormsResponse instantiates a new ListTaxFormsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTaxFormsResponseWithDefaults

`func NewListTaxFormsResponseWithDefaults() *ListTaxFormsResponse`

NewListTaxFormsResponseWithDefaults instantiates a new ListTaxFormsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxFormUrls

`func (o *ListTaxFormsResponse) GetTaxFormUrls() []TaxFormURL`

GetTaxFormUrls returns the TaxFormUrls field if non-nil, zero value otherwise.

### GetTaxFormUrlsOk

`func (o *ListTaxFormsResponse) GetTaxFormUrlsOk() (*[]TaxFormURL, bool)`

GetTaxFormUrlsOk returns a tuple with the TaxFormUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxFormUrls

`func (o *ListTaxFormsResponse) SetTaxFormUrls(v []TaxFormURL)`

SetTaxFormUrls sets TaxFormUrls field to given value.

### HasTaxFormUrls

`func (o *ListTaxFormsResponse) HasTaxFormUrls() bool`

HasTaxFormUrls returns a boolean if a field has been set.

### GetNextPageCursor

`func (o *ListTaxFormsResponse) GetNextPageCursor() string`

GetNextPageCursor returns the NextPageCursor field if non-nil, zero value otherwise.

### GetNextPageCursorOk

`func (o *ListTaxFormsResponse) GetNextPageCursorOk() (*string, bool)`

GetNextPageCursorOk returns a tuple with the NextPageCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageCursor

`func (o *ListTaxFormsResponse) SetNextPageCursor(v string)`

SetNextPageCursor sets NextPageCursor field to given value.

### HasNextPageCursor

`func (o *ListTaxFormsResponse) HasNextPageCursor() bool`

HasNextPageCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


