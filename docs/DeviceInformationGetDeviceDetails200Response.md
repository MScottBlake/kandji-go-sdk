# DeviceInformationGetDeviceDetails200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cellular** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseCellular**](DeviceInformationGetDeviceDetails200ResponseCellular.md) |  | [optional] 
**General** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseGeneral**](DeviceInformationGetDeviceDetails200ResponseGeneral.md) |  | [optional] 
**Hardware** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseHardware**](DeviceInformationGetDeviceDetails200ResponseHardware.md) |  | [optional] 
**Management** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseManagement**](DeviceInformationGetDeviceDetails200ResponseManagement.md) |  | [optional] 
**Security** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseSecurity**](DeviceInformationGetDeviceDetails200ResponseSecurity.md) |  | [optional] 
**Tags** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewDeviceInformationGetDeviceDetails200Response

`func NewDeviceInformationGetDeviceDetails200Response() *DeviceInformationGetDeviceDetails200Response`

NewDeviceInformationGetDeviceDetails200Response instantiates a new DeviceInformationGetDeviceDetails200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInformationGetDeviceDetails200ResponseWithDefaults

`func NewDeviceInformationGetDeviceDetails200ResponseWithDefaults() *DeviceInformationGetDeviceDetails200Response`

NewDeviceInformationGetDeviceDetails200ResponseWithDefaults instantiates a new DeviceInformationGetDeviceDetails200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCellular

`func (o *DeviceInformationGetDeviceDetails200Response) GetCellular() DeviceInformationGetDeviceDetails200ResponseCellular`

GetCellular returns the Cellular field if non-nil, zero value otherwise.

### GetCellularOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetCellularOk() (*DeviceInformationGetDeviceDetails200ResponseCellular, bool)`

GetCellularOk returns a tuple with the Cellular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellular

`func (o *DeviceInformationGetDeviceDetails200Response) SetCellular(v DeviceInformationGetDeviceDetails200ResponseCellular)`

SetCellular sets Cellular field to given value.

### HasCellular

`func (o *DeviceInformationGetDeviceDetails200Response) HasCellular() bool`

HasCellular returns a boolean if a field has been set.

### GetGeneral

`func (o *DeviceInformationGetDeviceDetails200Response) GetGeneral() DeviceInformationGetDeviceDetails200ResponseGeneral`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetGeneralOk() (*DeviceInformationGetDeviceDetails200ResponseGeneral, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *DeviceInformationGetDeviceDetails200Response) SetGeneral(v DeviceInformationGetDeviceDetails200ResponseGeneral)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *DeviceInformationGetDeviceDetails200Response) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetHardware

`func (o *DeviceInformationGetDeviceDetails200Response) GetHardware() DeviceInformationGetDeviceDetails200ResponseHardware`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetHardwareOk() (*DeviceInformationGetDeviceDetails200ResponseHardware, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *DeviceInformationGetDeviceDetails200Response) SetHardware(v DeviceInformationGetDeviceDetails200ResponseHardware)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *DeviceInformationGetDeviceDetails200Response) HasHardware() bool`

HasHardware returns a boolean if a field has been set.

### GetManagement

`func (o *DeviceInformationGetDeviceDetails200Response) GetManagement() DeviceInformationGetDeviceDetails200ResponseManagement`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetManagementOk() (*DeviceInformationGetDeviceDetails200ResponseManagement, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *DeviceInformationGetDeviceDetails200Response) SetManagement(v DeviceInformationGetDeviceDetails200ResponseManagement)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *DeviceInformationGetDeviceDetails200Response) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetSecurity

`func (o *DeviceInformationGetDeviceDetails200Response) GetSecurity() DeviceInformationGetDeviceDetails200ResponseSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetSecurityOk() (*DeviceInformationGetDeviceDetails200ResponseSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *DeviceInformationGetDeviceDetails200Response) SetSecurity(v DeviceInformationGetDeviceDetails200ResponseSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *DeviceInformationGetDeviceDetails200Response) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetTags

`func (o *DeviceInformationGetDeviceDetails200Response) GetTags() interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetTagsOk() (*interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceInformationGetDeviceDetails200Response) SetTags(v interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceInformationGetDeviceDetails200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *DeviceInformationGetDeviceDetails200Response) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *DeviceInformationGetDeviceDetails200Response) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


