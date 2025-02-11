# GetDeviceLostModeDetails200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**LostMode** | Pointer to [**IphoneOrIpadInLostMode200ResponseLostMode**](IphoneOrIpadInLostMode200ResponseLostMode.md) |  | [optional] 

## Methods

### NewGetDeviceLostModeDetails200Response

`func NewGetDeviceLostModeDetails200Response() *GetDeviceLostModeDetails200Response`

NewGetDeviceLostModeDetails200Response instantiates a new GetDeviceLostModeDetails200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceLostModeDetails200ResponseWithDefaults

`func NewGetDeviceLostModeDetails200ResponseWithDefaults() *GetDeviceLostModeDetails200Response`

NewGetDeviceLostModeDetails200ResponseWithDefaults instantiates a new GetDeviceLostModeDetails200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *GetDeviceLostModeDetails200Response) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GetDeviceLostModeDetails200Response) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GetDeviceLostModeDetails200Response) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *GetDeviceLostModeDetails200Response) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetLostMode

`func (o *GetDeviceLostModeDetails200Response) GetLostMode() IphoneOrIpadInLostMode200ResponseLostMode`

GetLostMode returns the LostMode field if non-nil, zero value otherwise.

### GetLostModeOk

`func (o *GetDeviceLostModeDetails200Response) GetLostModeOk() (*IphoneOrIpadInLostMode200ResponseLostMode, bool)`

GetLostModeOk returns a tuple with the LostMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostMode

`func (o *GetDeviceLostModeDetails200Response) SetLostMode(v IphoneOrIpadInLostMode200ResponseLostMode)`

SetLostMode sets LostMode field to given value.

### HasLostMode

`func (o *GetDeviceLostModeDetails200Response) HasLostMode() bool`

HasLostMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


