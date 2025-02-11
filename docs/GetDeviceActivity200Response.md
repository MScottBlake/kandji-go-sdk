# GetDeviceActivity200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**Activity** | Pointer to [**GetDeviceActivity200ResponseActivity**](GetDeviceActivity200ResponseActivity.md) |  | [optional] 

## Methods

### NewGetDeviceActivity200Response

`func NewGetDeviceActivity200Response() *GetDeviceActivity200Response`

NewGetDeviceActivity200Response instantiates a new GetDeviceActivity200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceActivity200ResponseWithDefaults

`func NewGetDeviceActivity200ResponseWithDefaults() *GetDeviceActivity200Response`

NewGetDeviceActivity200ResponseWithDefaults instantiates a new GetDeviceActivity200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *GetDeviceActivity200Response) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GetDeviceActivity200Response) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GetDeviceActivity200Response) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *GetDeviceActivity200Response) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetActivity

`func (o *GetDeviceActivity200Response) GetActivity() GetDeviceActivity200ResponseActivity`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *GetDeviceActivity200Response) GetActivityOk() (*GetDeviceActivity200ResponseActivity, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *GetDeviceActivity200Response) SetActivity(v GetDeviceActivity200ResponseActivity)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *GetDeviceActivity200Response) HasActivity() bool`

HasActivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


