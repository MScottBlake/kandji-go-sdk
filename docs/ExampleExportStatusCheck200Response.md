# ExampleExportStatusCheck200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Args** | Pointer to [**DeviceInfoForAllIpads200ResponseArgs**](DeviceInfoForAllIpads200ResponseArgs.md) |  | [optional] 
**ErrorMsg** | Pointer to **interface{}** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**SignedUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewExampleExportStatusCheck200Response

`func NewExampleExportStatusCheck200Response() *ExampleExportStatusCheck200Response`

NewExampleExportStatusCheck200Response instantiates a new ExampleExportStatusCheck200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExampleExportStatusCheck200ResponseWithDefaults

`func NewExampleExportStatusCheck200ResponseWithDefaults() *ExampleExportStatusCheck200Response`

NewExampleExportStatusCheck200ResponseWithDefaults instantiates a new ExampleExportStatusCheck200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExampleExportStatusCheck200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExampleExportStatusCheck200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExampleExportStatusCheck200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExampleExportStatusCheck200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ExampleExportStatusCheck200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExampleExportStatusCheck200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExampleExportStatusCheck200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExampleExportStatusCheck200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCategory

`func (o *ExampleExportStatusCheck200Response) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ExampleExportStatusCheck200Response) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ExampleExportStatusCheck200Response) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ExampleExportStatusCheck200Response) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetArgs

`func (o *ExampleExportStatusCheck200Response) GetArgs() DeviceInfoForAllIpads200ResponseArgs`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ExampleExportStatusCheck200Response) GetArgsOk() (*DeviceInfoForAllIpads200ResponseArgs, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ExampleExportStatusCheck200Response) SetArgs(v DeviceInfoForAllIpads200ResponseArgs)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *ExampleExportStatusCheck200Response) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetErrorMsg

`func (o *ExampleExportStatusCheck200Response) GetErrorMsg() interface{}`

GetErrorMsg returns the ErrorMsg field if non-nil, zero value otherwise.

### GetErrorMsgOk

`func (o *ExampleExportStatusCheck200Response) GetErrorMsgOk() (*interface{}, bool)`

GetErrorMsgOk returns a tuple with the ErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsg

`func (o *ExampleExportStatusCheck200Response) SetErrorMsg(v interface{})`

SetErrorMsg sets ErrorMsg field to given value.

### HasErrorMsg

`func (o *ExampleExportStatusCheck200Response) HasErrorMsg() bool`

HasErrorMsg returns a boolean if a field has been set.

### SetErrorMsgNil

`func (o *ExampleExportStatusCheck200Response) SetErrorMsgNil(b bool)`

 SetErrorMsgNil sets the value for ErrorMsg to be an explicit nil

### UnsetErrorMsg
`func (o *ExampleExportStatusCheck200Response) UnsetErrorMsg()`

UnsetErrorMsg ensures that no value is present for ErrorMsg, not even an explicit nil
### GetPath

`func (o *ExampleExportStatusCheck200Response) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ExampleExportStatusCheck200Response) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ExampleExportStatusCheck200Response) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ExampleExportStatusCheck200Response) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSignedUrl

`func (o *ExampleExportStatusCheck200Response) GetSignedUrl() string`

GetSignedUrl returns the SignedUrl field if non-nil, zero value otherwise.

### GetSignedUrlOk

`func (o *ExampleExportStatusCheck200Response) GetSignedUrlOk() (*string, bool)`

GetSignedUrlOk returns a tuple with the SignedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedUrl

`func (o *ExampleExportStatusCheck200Response) SetSignedUrl(v string)`

SetSignedUrl sets SignedUrl field to given value.

### HasSignedUrl

`func (o *ExampleExportStatusCheck200Response) HasSignedUrl() bool`

HasSignedUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ExampleExportStatusCheck200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExampleExportStatusCheck200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExampleExportStatusCheck200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ExampleExportStatusCheck200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ExampleExportStatusCheck200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExampleExportStatusCheck200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExampleExportStatusCheck200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ExampleExportStatusCheck200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


