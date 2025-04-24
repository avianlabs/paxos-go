# IdentityPublicDocumentUploadBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The file name. | 
**DocumentTypes** | Pointer to [**[]DocumentType**](DocumentType.md) | A list of document types contained within the uploaded file. | [optional] 

## Methods

### NewIdentityPublicDocumentUploadBody

`func NewIdentityPublicDocumentUploadBody(name string, ) *IdentityPublicDocumentUploadBody`

NewIdentityPublicDocumentUploadBody instantiates a new IdentityPublicDocumentUploadBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityPublicDocumentUploadBodyWithDefaults

`func NewIdentityPublicDocumentUploadBodyWithDefaults() *IdentityPublicDocumentUploadBody`

NewIdentityPublicDocumentUploadBodyWithDefaults instantiates a new IdentityPublicDocumentUploadBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdentityPublicDocumentUploadBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityPublicDocumentUploadBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityPublicDocumentUploadBody) SetName(v string)`

SetName sets Name field to given value.


### GetDocumentTypes

`func (o *IdentityPublicDocumentUploadBody) GetDocumentTypes() []DocumentType`

GetDocumentTypes returns the DocumentTypes field if non-nil, zero value otherwise.

### GetDocumentTypesOk

`func (o *IdentityPublicDocumentUploadBody) GetDocumentTypesOk() (*[]DocumentType, bool)`

GetDocumentTypesOk returns a tuple with the DocumentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypes

`func (o *IdentityPublicDocumentUploadBody) SetDocumentTypes(v []DocumentType)`

SetDocumentTypes sets DocumentTypes field to given value.

### HasDocumentTypes

`func (o *IdentityPublicDocumentUploadBody) HasDocumentTypes() bool`

HasDocumentTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


