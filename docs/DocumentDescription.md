# DocumentDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DocumentTypes** | Pointer to [**[]DocumentType**](DocumentType.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ArchivedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDocumentDescription

`func NewDocumentDescription() *DocumentDescription`

NewDocumentDescription instantiates a new DocumentDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentDescriptionWithDefaults

`func NewDocumentDescriptionWithDefaults() *DocumentDescription`

NewDocumentDescriptionWithDefaults instantiates a new DocumentDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *DocumentDescription) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *DocumentDescription) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *DocumentDescription) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *DocumentDescription) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetName

`func (o *DocumentDescription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentDescription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentDescription) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocumentDescription) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDocumentTypes

`func (o *DocumentDescription) GetDocumentTypes() []DocumentType`

GetDocumentTypes returns the DocumentTypes field if non-nil, zero value otherwise.

### GetDocumentTypesOk

`func (o *DocumentDescription) GetDocumentTypesOk() (*[]DocumentType, bool)`

GetDocumentTypesOk returns a tuple with the DocumentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypes

`func (o *DocumentDescription) SetDocumentTypes(v []DocumentType)`

SetDocumentTypes sets DocumentTypes field to given value.

### HasDocumentTypes

`func (o *DocumentDescription) HasDocumentTypes() bool`

HasDocumentTypes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DocumentDescription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentDescription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentDescription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocumentDescription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *DocumentDescription) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *DocumentDescription) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *DocumentDescription) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *DocumentDescription) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


