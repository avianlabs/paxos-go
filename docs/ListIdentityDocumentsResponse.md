# ListIdentityDocumentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documents** | Pointer to [**[]DocumentDescription**](DocumentDescription.md) |  | [optional] 
**PendingDocuments** | Pointer to [**[]DocumentType**](DocumentType.md) |  | [optional] 

## Methods

### NewListIdentityDocumentsResponse

`func NewListIdentityDocumentsResponse() *ListIdentityDocumentsResponse`

NewListIdentityDocumentsResponse instantiates a new ListIdentityDocumentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIdentityDocumentsResponseWithDefaults

`func NewListIdentityDocumentsResponseWithDefaults() *ListIdentityDocumentsResponse`

NewListIdentityDocumentsResponseWithDefaults instantiates a new ListIdentityDocumentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocuments

`func (o *ListIdentityDocumentsResponse) GetDocuments() []DocumentDescription`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *ListIdentityDocumentsResponse) GetDocumentsOk() (*[]DocumentDescription, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *ListIdentityDocumentsResponse) SetDocuments(v []DocumentDescription)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *ListIdentityDocumentsResponse) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetPendingDocuments

`func (o *ListIdentityDocumentsResponse) GetPendingDocuments() []DocumentType`

GetPendingDocuments returns the PendingDocuments field if non-nil, zero value otherwise.

### GetPendingDocumentsOk

`func (o *ListIdentityDocumentsResponse) GetPendingDocumentsOk() (*[]DocumentType, bool)`

GetPendingDocumentsOk returns a tuple with the PendingDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDocuments

`func (o *ListIdentityDocumentsResponse) SetPendingDocuments(v []DocumentType)`

SetPendingDocuments sets PendingDocuments field to given value.

### HasPendingDocuments

`func (o *ListIdentityDocumentsResponse) HasPendingDocuments() bool`

HasPendingDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


