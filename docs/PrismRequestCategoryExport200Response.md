# PrismRequestCategoryExport200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Args** | Pointer to [**PrismRequestCategoryExport200ResponseArgs**](PrismRequestCategoryExport200ResponseArgs.md) |  | [optional] 
**ErrorMsg** | Pointer to **interface{}** |  | [optional] 
**Path** | Pointer to **interface{}** |  | [optional] 
**SignedUrl** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewPrismRequestCategoryExport200Response

`func NewPrismRequestCategoryExport200Response() *PrismRequestCategoryExport200Response`

NewPrismRequestCategoryExport200Response instantiates a new PrismRequestCategoryExport200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismRequestCategoryExport200ResponseWithDefaults

`func NewPrismRequestCategoryExport200ResponseWithDefaults() *PrismRequestCategoryExport200Response`

NewPrismRequestCategoryExport200ResponseWithDefaults instantiates a new PrismRequestCategoryExport200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrismRequestCategoryExport200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrismRequestCategoryExport200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrismRequestCategoryExport200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrismRequestCategoryExport200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *PrismRequestCategoryExport200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrismRequestCategoryExport200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrismRequestCategoryExport200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrismRequestCategoryExport200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCategory

`func (o *PrismRequestCategoryExport200Response) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PrismRequestCategoryExport200Response) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PrismRequestCategoryExport200Response) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PrismRequestCategoryExport200Response) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetArgs

`func (o *PrismRequestCategoryExport200Response) GetArgs() PrismRequestCategoryExport200ResponseArgs`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *PrismRequestCategoryExport200Response) GetArgsOk() (*PrismRequestCategoryExport200ResponseArgs, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *PrismRequestCategoryExport200Response) SetArgs(v PrismRequestCategoryExport200ResponseArgs)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *PrismRequestCategoryExport200Response) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetErrorMsg

`func (o *PrismRequestCategoryExport200Response) GetErrorMsg() interface{}`

GetErrorMsg returns the ErrorMsg field if non-nil, zero value otherwise.

### GetErrorMsgOk

`func (o *PrismRequestCategoryExport200Response) GetErrorMsgOk() (*interface{}, bool)`

GetErrorMsgOk returns a tuple with the ErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsg

`func (o *PrismRequestCategoryExport200Response) SetErrorMsg(v interface{})`

SetErrorMsg sets ErrorMsg field to given value.

### HasErrorMsg

`func (o *PrismRequestCategoryExport200Response) HasErrorMsg() bool`

HasErrorMsg returns a boolean if a field has been set.

### SetErrorMsgNil

`func (o *PrismRequestCategoryExport200Response) SetErrorMsgNil(b bool)`

 SetErrorMsgNil sets the value for ErrorMsg to be an explicit nil

### UnsetErrorMsg
`func (o *PrismRequestCategoryExport200Response) UnsetErrorMsg()`

UnsetErrorMsg ensures that no value is present for ErrorMsg, not even an explicit nil
### GetPath

`func (o *PrismRequestCategoryExport200Response) GetPath() interface{}`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PrismRequestCategoryExport200Response) GetPathOk() (*interface{}, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PrismRequestCategoryExport200Response) SetPath(v interface{})`

SetPath sets Path field to given value.

### HasPath

`func (o *PrismRequestCategoryExport200Response) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *PrismRequestCategoryExport200Response) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *PrismRequestCategoryExport200Response) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetSignedUrl

`func (o *PrismRequestCategoryExport200Response) GetSignedUrl() interface{}`

GetSignedUrl returns the SignedUrl field if non-nil, zero value otherwise.

### GetSignedUrlOk

`func (o *PrismRequestCategoryExport200Response) GetSignedUrlOk() (*interface{}, bool)`

GetSignedUrlOk returns a tuple with the SignedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedUrl

`func (o *PrismRequestCategoryExport200Response) SetSignedUrl(v interface{})`

SetSignedUrl sets SignedUrl field to given value.

### HasSignedUrl

`func (o *PrismRequestCategoryExport200Response) HasSignedUrl() bool`

HasSignedUrl returns a boolean if a field has been set.

### SetSignedUrlNil

`func (o *PrismRequestCategoryExport200Response) SetSignedUrlNil(b bool)`

 SetSignedUrlNil sets the value for SignedUrl to be an explicit nil

### UnsetSignedUrl
`func (o *PrismRequestCategoryExport200Response) UnsetSignedUrl()`

UnsetSignedUrl ensures that no value is present for SignedUrl, not even an explicit nil
### GetCreatedAt

`func (o *PrismRequestCategoryExport200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PrismRequestCategoryExport200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PrismRequestCategoryExport200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PrismRequestCategoryExport200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PrismRequestCategoryExport200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PrismRequestCategoryExport200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PrismRequestCategoryExport200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PrismRequestCategoryExport200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


