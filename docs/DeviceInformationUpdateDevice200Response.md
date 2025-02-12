# DeviceInformationUpdateDevice200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**LastCheckIn** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **string** |  | [optional] 
**BlueprintId** | Pointer to **string** |  | [optional] 
**MdmEnabled** | Pointer to **int32** |  | [optional] 
**AgentInstalled** | Pointer to **int32** |  | [optional] 
**IsMissing** | Pointer to **int32** |  | [optional] 
**IsRemoved** | Pointer to **int32** |  | [optional] 
**AgentVersion** | Pointer to **string** |  | [optional] 
**FirstEnrollment** | Pointer to **string** |  | [optional] 
**LastEnrollment** | Pointer to **string** |  | [optional] 
**BlueprintName** | Pointer to **string** |  | [optional] 
**LostModeStatus** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewDeviceInformationUpdateDevice200Response

`func NewDeviceInformationUpdateDevice200Response() *DeviceInformationUpdateDevice200Response`

NewDeviceInformationUpdateDevice200Response instantiates a new DeviceInformationUpdateDevice200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInformationUpdateDevice200ResponseWithDefaults

`func NewDeviceInformationUpdateDevice200ResponseWithDefaults() *DeviceInformationUpdateDevice200Response`

NewDeviceInformationUpdateDevice200ResponseWithDefaults instantiates a new DeviceInformationUpdateDevice200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *DeviceInformationUpdateDevice200Response) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceInformationUpdateDevice200Response) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceInformationUpdateDevice200Response) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DeviceInformationUpdateDevice200Response) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *DeviceInformationUpdateDevice200Response) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DeviceInformationUpdateDevice200Response) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DeviceInformationUpdateDevice200Response) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DeviceInformationUpdateDevice200Response) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetModel

`func (o *DeviceInformationUpdateDevice200Response) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceInformationUpdateDevice200Response) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceInformationUpdateDevice200Response) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DeviceInformationUpdateDevice200Response) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSerialNumber

`func (o *DeviceInformationUpdateDevice200Response) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DeviceInformationUpdateDevice200Response) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DeviceInformationUpdateDevice200Response) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DeviceInformationUpdateDevice200Response) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetPlatform

`func (o *DeviceInformationUpdateDevice200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DeviceInformationUpdateDevice200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DeviceInformationUpdateDevice200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DeviceInformationUpdateDevice200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetOsVersion

`func (o *DeviceInformationUpdateDevice200Response) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DeviceInformationUpdateDevice200Response) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *DeviceInformationUpdateDevice200Response) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *DeviceInformationUpdateDevice200Response) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetLastCheckIn

`func (o *DeviceInformationUpdateDevice200Response) GetLastCheckIn() string`

GetLastCheckIn returns the LastCheckIn field if non-nil, zero value otherwise.

### GetLastCheckInOk

`func (o *DeviceInformationUpdateDevice200Response) GetLastCheckInOk() (*string, bool)`

GetLastCheckInOk returns a tuple with the LastCheckIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckIn

`func (o *DeviceInformationUpdateDevice200Response) SetLastCheckIn(v string)`

SetLastCheckIn sets LastCheckIn field to given value.

### HasLastCheckIn

`func (o *DeviceInformationUpdateDevice200Response) HasLastCheckIn() bool`

HasLastCheckIn returns a boolean if a field has been set.

### GetUser

`func (o *DeviceInformationUpdateDevice200Response) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DeviceInformationUpdateDevice200Response) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DeviceInformationUpdateDevice200Response) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *DeviceInformationUpdateDevice200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAssetTag

`func (o *DeviceInformationUpdateDevice200Response) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *DeviceInformationUpdateDevice200Response) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *DeviceInformationUpdateDevice200Response) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *DeviceInformationUpdateDevice200Response) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetBlueprintId

`func (o *DeviceInformationUpdateDevice200Response) GetBlueprintId() string`

GetBlueprintId returns the BlueprintId field if non-nil, zero value otherwise.

### GetBlueprintIdOk

`func (o *DeviceInformationUpdateDevice200Response) GetBlueprintIdOk() (*string, bool)`

GetBlueprintIdOk returns a tuple with the BlueprintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintId

`func (o *DeviceInformationUpdateDevice200Response) SetBlueprintId(v string)`

SetBlueprintId sets BlueprintId field to given value.

### HasBlueprintId

`func (o *DeviceInformationUpdateDevice200Response) HasBlueprintId() bool`

HasBlueprintId returns a boolean if a field has been set.

### GetMdmEnabled

`func (o *DeviceInformationUpdateDevice200Response) GetMdmEnabled() int32`

GetMdmEnabled returns the MdmEnabled field if non-nil, zero value otherwise.

### GetMdmEnabledOk

`func (o *DeviceInformationUpdateDevice200Response) GetMdmEnabledOk() (*int32, bool)`

GetMdmEnabledOk returns a tuple with the MdmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmEnabled

`func (o *DeviceInformationUpdateDevice200Response) SetMdmEnabled(v int32)`

SetMdmEnabled sets MdmEnabled field to given value.

### HasMdmEnabled

`func (o *DeviceInformationUpdateDevice200Response) HasMdmEnabled() bool`

HasMdmEnabled returns a boolean if a field has been set.

### GetAgentInstalled

`func (o *DeviceInformationUpdateDevice200Response) GetAgentInstalled() int32`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *DeviceInformationUpdateDevice200Response) GetAgentInstalledOk() (*int32, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *DeviceInformationUpdateDevice200Response) SetAgentInstalled(v int32)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *DeviceInformationUpdateDevice200Response) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetIsMissing

`func (o *DeviceInformationUpdateDevice200Response) GetIsMissing() int32`

GetIsMissing returns the IsMissing field if non-nil, zero value otherwise.

### GetIsMissingOk

`func (o *DeviceInformationUpdateDevice200Response) GetIsMissingOk() (*int32, bool)`

GetIsMissingOk returns a tuple with the IsMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMissing

`func (o *DeviceInformationUpdateDevice200Response) SetIsMissing(v int32)`

SetIsMissing sets IsMissing field to given value.

### HasIsMissing

`func (o *DeviceInformationUpdateDevice200Response) HasIsMissing() bool`

HasIsMissing returns a boolean if a field has been set.

### GetIsRemoved

`func (o *DeviceInformationUpdateDevice200Response) GetIsRemoved() int32`

GetIsRemoved returns the IsRemoved field if non-nil, zero value otherwise.

### GetIsRemovedOk

`func (o *DeviceInformationUpdateDevice200Response) GetIsRemovedOk() (*int32, bool)`

GetIsRemovedOk returns a tuple with the IsRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoved

`func (o *DeviceInformationUpdateDevice200Response) SetIsRemoved(v int32)`

SetIsRemoved sets IsRemoved field to given value.

### HasIsRemoved

`func (o *DeviceInformationUpdateDevice200Response) HasIsRemoved() bool`

HasIsRemoved returns a boolean if a field has been set.

### GetAgentVersion

`func (o *DeviceInformationUpdateDevice200Response) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *DeviceInformationUpdateDevice200Response) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *DeviceInformationUpdateDevice200Response) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *DeviceInformationUpdateDevice200Response) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetFirstEnrollment

`func (o *DeviceInformationUpdateDevice200Response) GetFirstEnrollment() string`

GetFirstEnrollment returns the FirstEnrollment field if non-nil, zero value otherwise.

### GetFirstEnrollmentOk

`func (o *DeviceInformationUpdateDevice200Response) GetFirstEnrollmentOk() (*string, bool)`

GetFirstEnrollmentOk returns a tuple with the FirstEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstEnrollment

`func (o *DeviceInformationUpdateDevice200Response) SetFirstEnrollment(v string)`

SetFirstEnrollment sets FirstEnrollment field to given value.

### HasFirstEnrollment

`func (o *DeviceInformationUpdateDevice200Response) HasFirstEnrollment() bool`

HasFirstEnrollment returns a boolean if a field has been set.

### GetLastEnrollment

`func (o *DeviceInformationUpdateDevice200Response) GetLastEnrollment() string`

GetLastEnrollment returns the LastEnrollment field if non-nil, zero value otherwise.

### GetLastEnrollmentOk

`func (o *DeviceInformationUpdateDevice200Response) GetLastEnrollmentOk() (*string, bool)`

GetLastEnrollmentOk returns a tuple with the LastEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnrollment

`func (o *DeviceInformationUpdateDevice200Response) SetLastEnrollment(v string)`

SetLastEnrollment sets LastEnrollment field to given value.

### HasLastEnrollment

`func (o *DeviceInformationUpdateDevice200Response) HasLastEnrollment() bool`

HasLastEnrollment returns a boolean if a field has been set.

### GetBlueprintName

`func (o *DeviceInformationUpdateDevice200Response) GetBlueprintName() string`

GetBlueprintName returns the BlueprintName field if non-nil, zero value otherwise.

### GetBlueprintNameOk

`func (o *DeviceInformationUpdateDevice200Response) GetBlueprintNameOk() (*string, bool)`

GetBlueprintNameOk returns a tuple with the BlueprintName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintName

`func (o *DeviceInformationUpdateDevice200Response) SetBlueprintName(v string)`

SetBlueprintName sets BlueprintName field to given value.

### HasBlueprintName

`func (o *DeviceInformationUpdateDevice200Response) HasBlueprintName() bool`

HasBlueprintName returns a boolean if a field has been set.

### GetLostModeStatus

`func (o *DeviceInformationUpdateDevice200Response) GetLostModeStatus() string`

GetLostModeStatus returns the LostModeStatus field if non-nil, zero value otherwise.

### GetLostModeStatusOk

`func (o *DeviceInformationUpdateDevice200Response) GetLostModeStatusOk() (*string, bool)`

GetLostModeStatusOk returns a tuple with the LostModeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostModeStatus

`func (o *DeviceInformationUpdateDevice200Response) SetLostModeStatus(v string)`

SetLostModeStatus sets LostModeStatus field to given value.

### HasLostModeStatus

`func (o *DeviceInformationUpdateDevice200Response) HasLostModeStatus() bool`

HasLostModeStatus returns a boolean if a field has been set.

### GetTags

`func (o *DeviceInformationUpdateDevice200Response) GetTags() interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceInformationUpdateDevice200Response) GetTagsOk() (*interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceInformationUpdateDevice200Response) SetTags(v interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceInformationUpdateDevice200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *DeviceInformationUpdateDevice200Response) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *DeviceInformationUpdateDevice200Response) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


