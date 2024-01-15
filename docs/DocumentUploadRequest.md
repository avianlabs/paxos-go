# DocumentUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The file name. | 
**DocumentTypes** | Pointer to [**[]DocumentType**](DocumentType.md) | A list of document types contained within the uploaded file. | [optional] 

## Methods

### NewDocumentUploadRequest

`func NewDocumentUploadRequest(name string, ) *DocumentUploadRequest`

NewDocumentUploadRequest instantiates a new DocumentUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentUploadRequestWithDefaults

`func NewDocumentUploadRequestWithDefaults() *DocumentUploadRequest`

NewDocumentUploadRequestWithDefaults instantiates a new DocumentUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DocumentUploadRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentUploadRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentUploadRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDocumentTypes

`func (o *DocumentUploadRequest) GetDocumentTypes() []DocumentType`

GetDocumentTypes returns the DocumentTypes field if non-nil, zero value otherwise.

### GetDocumentTypesOk

`func (o *DocumentUploadRequest) GetDocumentTypesOk() (*[]DocumentType, bool)`

GetDocumentTypesOk returns a tuple with the DocumentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypes

`func (o *DocumentUploadRequest) SetDocumentTypes(v []DocumentType)`

SetDocumentTypes sets DocumentTypes field to given value.

### HasDocumentTypes

`func (o *DocumentUploadRequest) HasDocumentTypes() bool`

HasDocumentTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


