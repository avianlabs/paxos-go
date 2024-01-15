# DocumentUploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | Pointer to **string** | The id assigned to the file. | [optional] 
**Name** | Pointer to **string** | The file name. | [optional] 
**UploadUrl** | Pointer to **string** | The URI where we expect the file to be uploaded into. | [optional] 

## Methods

### NewDocumentUploadResponse

`func NewDocumentUploadResponse() *DocumentUploadResponse`

NewDocumentUploadResponse instantiates a new DocumentUploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentUploadResponseWithDefaults

`func NewDocumentUploadResponseWithDefaults() *DocumentUploadResponse`

NewDocumentUploadResponseWithDefaults instantiates a new DocumentUploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *DocumentUploadResponse) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *DocumentUploadResponse) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *DocumentUploadResponse) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *DocumentUploadResponse) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetName

`func (o *DocumentUploadResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentUploadResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentUploadResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocumentUploadResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUploadUrl

`func (o *DocumentUploadResponse) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *DocumentUploadResponse) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *DocumentUploadResponse) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *DocumentUploadResponse) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


