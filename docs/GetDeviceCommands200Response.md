# GetDeviceCommands200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**Commands** | Pointer to [**GetDeviceActivity200ResponseActivity**](GetDeviceActivity200ResponseActivity.md) |  | [optional] 

## Methods

### NewGetDeviceCommands200Response

`func NewGetDeviceCommands200Response() *GetDeviceCommands200Response`

NewGetDeviceCommands200Response instantiates a new GetDeviceCommands200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceCommands200ResponseWithDefaults

`func NewGetDeviceCommands200ResponseWithDefaults() *GetDeviceCommands200Response`

NewGetDeviceCommands200ResponseWithDefaults instantiates a new GetDeviceCommands200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *GetDeviceCommands200Response) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GetDeviceCommands200Response) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GetDeviceCommands200Response) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *GetDeviceCommands200Response) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetCommands

`func (o *GetDeviceCommands200Response) GetCommands() GetDeviceActivity200ResponseActivity`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *GetDeviceCommands200Response) GetCommandsOk() (*GetDeviceActivity200ResponseActivity, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *GetDeviceCommands200Response) SetCommands(v GetDeviceActivity200ResponseActivity)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *GetDeviceCommands200Response) HasCommands() bool`

HasCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


