# DeviceInformationGetDevice200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentInstalled** | Pointer to **int32** |  | [optional] 
**AgentVersion** | Pointer to **string** |  | [optional] 
**ApiLevel** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **string** |  | [optional] 
**BlueprintId** | Pointer to **string** |  | [optional] 
**BlueprintName** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**FirstEnrollment** | Pointer to **string** |  | [optional] 
**FullSoftwareVersion** | Pointer to **string** |  | [optional] 
**IsMissing** | Pointer to **int32** |  | [optional] 
**IsRemoved** | Pointer to **int32** |  | [optional] 
**LastCheckIn** | Pointer to **string** |  | [optional] 
**LastEnrollment** | Pointer to **string** |  | [optional] 
**LostModeStatus** | Pointer to **string** |  | [optional] 
**MdmEnabled** | Pointer to **int32** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**SecurityPatchLevel** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**SupplementalBuildVersion** | Pointer to **string** |  | [optional] 
**SupplementalOsVersionExtra** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **interface{}** |  | [optional] 
**User** | Pointer to [**DeviceInformationGetDevice200ResponseUser**](DeviceInformationGetDevice200ResponseUser.md) |  | [optional] 

## Methods

### NewDeviceInformationGetDevice200Response

`func NewDeviceInformationGetDevice200Response() *DeviceInformationGetDevice200Response`

NewDeviceInformationGetDevice200Response instantiates a new DeviceInformationGetDevice200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInformationGetDevice200ResponseWithDefaults

`func NewDeviceInformationGetDevice200ResponseWithDefaults() *DeviceInformationGetDevice200Response`

NewDeviceInformationGetDevice200ResponseWithDefaults instantiates a new DeviceInformationGetDevice200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentInstalled

`func (o *DeviceInformationGetDevice200Response) GetAgentInstalled() int32`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *DeviceInformationGetDevice200Response) GetAgentInstalledOk() (*int32, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *DeviceInformationGetDevice200Response) SetAgentInstalled(v int32)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *DeviceInformationGetDevice200Response) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetAgentVersion

`func (o *DeviceInformationGetDevice200Response) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *DeviceInformationGetDevice200Response) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *DeviceInformationGetDevice200Response) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *DeviceInformationGetDevice200Response) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetApiLevel

`func (o *DeviceInformationGetDevice200Response) GetApiLevel() string`

GetApiLevel returns the ApiLevel field if non-nil, zero value otherwise.

### GetApiLevelOk

`func (o *DeviceInformationGetDevice200Response) GetApiLevelOk() (*string, bool)`

GetApiLevelOk returns a tuple with the ApiLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiLevel

`func (o *DeviceInformationGetDevice200Response) SetApiLevel(v string)`

SetApiLevel sets ApiLevel field to given value.

### HasApiLevel

`func (o *DeviceInformationGetDevice200Response) HasApiLevel() bool`

HasApiLevel returns a boolean if a field has been set.

### GetAssetTag

`func (o *DeviceInformationGetDevice200Response) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *DeviceInformationGetDevice200Response) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *DeviceInformationGetDevice200Response) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *DeviceInformationGetDevice200Response) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetBlueprintId

`func (o *DeviceInformationGetDevice200Response) GetBlueprintId() string`

GetBlueprintId returns the BlueprintId field if non-nil, zero value otherwise.

### GetBlueprintIdOk

`func (o *DeviceInformationGetDevice200Response) GetBlueprintIdOk() (*string, bool)`

GetBlueprintIdOk returns a tuple with the BlueprintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintId

`func (o *DeviceInformationGetDevice200Response) SetBlueprintId(v string)`

SetBlueprintId sets BlueprintId field to given value.

### HasBlueprintId

`func (o *DeviceInformationGetDevice200Response) HasBlueprintId() bool`

HasBlueprintId returns a boolean if a field has been set.

### GetBlueprintName

`func (o *DeviceInformationGetDevice200Response) GetBlueprintName() string`

GetBlueprintName returns the BlueprintName field if non-nil, zero value otherwise.

### GetBlueprintNameOk

`func (o *DeviceInformationGetDevice200Response) GetBlueprintNameOk() (*string, bool)`

GetBlueprintNameOk returns a tuple with the BlueprintName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintName

`func (o *DeviceInformationGetDevice200Response) SetBlueprintName(v string)`

SetBlueprintName sets BlueprintName field to given value.

### HasBlueprintName

`func (o *DeviceInformationGetDevice200Response) HasBlueprintName() bool`

HasBlueprintName returns a boolean if a field has been set.

### GetDeviceId

`func (o *DeviceInformationGetDevice200Response) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceInformationGetDevice200Response) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceInformationGetDevice200Response) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DeviceInformationGetDevice200Response) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *DeviceInformationGetDevice200Response) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DeviceInformationGetDevice200Response) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DeviceInformationGetDevice200Response) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DeviceInformationGetDevice200Response) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetFirstEnrollment

`func (o *DeviceInformationGetDevice200Response) GetFirstEnrollment() string`

GetFirstEnrollment returns the FirstEnrollment field if non-nil, zero value otherwise.

### GetFirstEnrollmentOk

`func (o *DeviceInformationGetDevice200Response) GetFirstEnrollmentOk() (*string, bool)`

GetFirstEnrollmentOk returns a tuple with the FirstEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstEnrollment

`func (o *DeviceInformationGetDevice200Response) SetFirstEnrollment(v string)`

SetFirstEnrollment sets FirstEnrollment field to given value.

### HasFirstEnrollment

`func (o *DeviceInformationGetDevice200Response) HasFirstEnrollment() bool`

HasFirstEnrollment returns a boolean if a field has been set.

### GetFullSoftwareVersion

`func (o *DeviceInformationGetDevice200Response) GetFullSoftwareVersion() string`

GetFullSoftwareVersion returns the FullSoftwareVersion field if non-nil, zero value otherwise.

### GetFullSoftwareVersionOk

`func (o *DeviceInformationGetDevice200Response) GetFullSoftwareVersionOk() (*string, bool)`

GetFullSoftwareVersionOk returns a tuple with the FullSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullSoftwareVersion

`func (o *DeviceInformationGetDevice200Response) SetFullSoftwareVersion(v string)`

SetFullSoftwareVersion sets FullSoftwareVersion field to given value.

### HasFullSoftwareVersion

`func (o *DeviceInformationGetDevice200Response) HasFullSoftwareVersion() bool`

HasFullSoftwareVersion returns a boolean if a field has been set.

### GetIsMissing

`func (o *DeviceInformationGetDevice200Response) GetIsMissing() int32`

GetIsMissing returns the IsMissing field if non-nil, zero value otherwise.

### GetIsMissingOk

`func (o *DeviceInformationGetDevice200Response) GetIsMissingOk() (*int32, bool)`

GetIsMissingOk returns a tuple with the IsMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMissing

`func (o *DeviceInformationGetDevice200Response) SetIsMissing(v int32)`

SetIsMissing sets IsMissing field to given value.

### HasIsMissing

`func (o *DeviceInformationGetDevice200Response) HasIsMissing() bool`

HasIsMissing returns a boolean if a field has been set.

### GetIsRemoved

`func (o *DeviceInformationGetDevice200Response) GetIsRemoved() int32`

GetIsRemoved returns the IsRemoved field if non-nil, zero value otherwise.

### GetIsRemovedOk

`func (o *DeviceInformationGetDevice200Response) GetIsRemovedOk() (*int32, bool)`

GetIsRemovedOk returns a tuple with the IsRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoved

`func (o *DeviceInformationGetDevice200Response) SetIsRemoved(v int32)`

SetIsRemoved sets IsRemoved field to given value.

### HasIsRemoved

`func (o *DeviceInformationGetDevice200Response) HasIsRemoved() bool`

HasIsRemoved returns a boolean if a field has been set.

### GetLastCheckIn

`func (o *DeviceInformationGetDevice200Response) GetLastCheckIn() string`

GetLastCheckIn returns the LastCheckIn field if non-nil, zero value otherwise.

### GetLastCheckInOk

`func (o *DeviceInformationGetDevice200Response) GetLastCheckInOk() (*string, bool)`

GetLastCheckInOk returns a tuple with the LastCheckIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckIn

`func (o *DeviceInformationGetDevice200Response) SetLastCheckIn(v string)`

SetLastCheckIn sets LastCheckIn field to given value.

### HasLastCheckIn

`func (o *DeviceInformationGetDevice200Response) HasLastCheckIn() bool`

HasLastCheckIn returns a boolean if a field has been set.

### GetLastEnrollment

`func (o *DeviceInformationGetDevice200Response) GetLastEnrollment() string`

GetLastEnrollment returns the LastEnrollment field if non-nil, zero value otherwise.

### GetLastEnrollmentOk

`func (o *DeviceInformationGetDevice200Response) GetLastEnrollmentOk() (*string, bool)`

GetLastEnrollmentOk returns a tuple with the LastEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnrollment

`func (o *DeviceInformationGetDevice200Response) SetLastEnrollment(v string)`

SetLastEnrollment sets LastEnrollment field to given value.

### HasLastEnrollment

`func (o *DeviceInformationGetDevice200Response) HasLastEnrollment() bool`

HasLastEnrollment returns a boolean if a field has been set.

### GetLostModeStatus

`func (o *DeviceInformationGetDevice200Response) GetLostModeStatus() string`

GetLostModeStatus returns the LostModeStatus field if non-nil, zero value otherwise.

### GetLostModeStatusOk

`func (o *DeviceInformationGetDevice200Response) GetLostModeStatusOk() (*string, bool)`

GetLostModeStatusOk returns a tuple with the LostModeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostModeStatus

`func (o *DeviceInformationGetDevice200Response) SetLostModeStatus(v string)`

SetLostModeStatus sets LostModeStatus field to given value.

### HasLostModeStatus

`func (o *DeviceInformationGetDevice200Response) HasLostModeStatus() bool`

HasLostModeStatus returns a boolean if a field has been set.

### GetMdmEnabled

`func (o *DeviceInformationGetDevice200Response) GetMdmEnabled() int32`

GetMdmEnabled returns the MdmEnabled field if non-nil, zero value otherwise.

### GetMdmEnabledOk

`func (o *DeviceInformationGetDevice200Response) GetMdmEnabledOk() (*int32, bool)`

GetMdmEnabledOk returns a tuple with the MdmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmEnabled

`func (o *DeviceInformationGetDevice200Response) SetMdmEnabled(v int32)`

SetMdmEnabled sets MdmEnabled field to given value.

### HasMdmEnabled

`func (o *DeviceInformationGetDevice200Response) HasMdmEnabled() bool`

HasMdmEnabled returns a boolean if a field has been set.

### GetModel

`func (o *DeviceInformationGetDevice200Response) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceInformationGetDevice200Response) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceInformationGetDevice200Response) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DeviceInformationGetDevice200Response) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOsVersion

`func (o *DeviceInformationGetDevice200Response) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DeviceInformationGetDevice200Response) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *DeviceInformationGetDevice200Response) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *DeviceInformationGetDevice200Response) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPlatform

`func (o *DeviceInformationGetDevice200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DeviceInformationGetDevice200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DeviceInformationGetDevice200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DeviceInformationGetDevice200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSecurityPatchLevel

`func (o *DeviceInformationGetDevice200Response) GetSecurityPatchLevel() string`

GetSecurityPatchLevel returns the SecurityPatchLevel field if non-nil, zero value otherwise.

### GetSecurityPatchLevelOk

`func (o *DeviceInformationGetDevice200Response) GetSecurityPatchLevelOk() (*string, bool)`

GetSecurityPatchLevelOk returns a tuple with the SecurityPatchLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPatchLevel

`func (o *DeviceInformationGetDevice200Response) SetSecurityPatchLevel(v string)`

SetSecurityPatchLevel sets SecurityPatchLevel field to given value.

### HasSecurityPatchLevel

`func (o *DeviceInformationGetDevice200Response) HasSecurityPatchLevel() bool`

HasSecurityPatchLevel returns a boolean if a field has been set.

### GetSerialNumber

`func (o *DeviceInformationGetDevice200Response) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DeviceInformationGetDevice200Response) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DeviceInformationGetDevice200Response) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DeviceInformationGetDevice200Response) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSupplementalBuildVersion

`func (o *DeviceInformationGetDevice200Response) GetSupplementalBuildVersion() string`

GetSupplementalBuildVersion returns the SupplementalBuildVersion field if non-nil, zero value otherwise.

### GetSupplementalBuildVersionOk

`func (o *DeviceInformationGetDevice200Response) GetSupplementalBuildVersionOk() (*string, bool)`

GetSupplementalBuildVersionOk returns a tuple with the SupplementalBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementalBuildVersion

`func (o *DeviceInformationGetDevice200Response) SetSupplementalBuildVersion(v string)`

SetSupplementalBuildVersion sets SupplementalBuildVersion field to given value.

### HasSupplementalBuildVersion

`func (o *DeviceInformationGetDevice200Response) HasSupplementalBuildVersion() bool`

HasSupplementalBuildVersion returns a boolean if a field has been set.

### GetSupplementalOsVersionExtra

`func (o *DeviceInformationGetDevice200Response) GetSupplementalOsVersionExtra() string`

GetSupplementalOsVersionExtra returns the SupplementalOsVersionExtra field if non-nil, zero value otherwise.

### GetSupplementalOsVersionExtraOk

`func (o *DeviceInformationGetDevice200Response) GetSupplementalOsVersionExtraOk() (*string, bool)`

GetSupplementalOsVersionExtraOk returns a tuple with the SupplementalOsVersionExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementalOsVersionExtra

`func (o *DeviceInformationGetDevice200Response) SetSupplementalOsVersionExtra(v string)`

SetSupplementalOsVersionExtra sets SupplementalOsVersionExtra field to given value.

### HasSupplementalOsVersionExtra

`func (o *DeviceInformationGetDevice200Response) HasSupplementalOsVersionExtra() bool`

HasSupplementalOsVersionExtra returns a boolean if a field has been set.

### GetTags

`func (o *DeviceInformationGetDevice200Response) GetTags() interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceInformationGetDevice200Response) GetTagsOk() (*interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceInformationGetDevice200Response) SetTags(v interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceInformationGetDevice200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *DeviceInformationGetDevice200Response) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *DeviceInformationGetDevice200Response) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetUser

`func (o *DeviceInformationGetDevice200Response) GetUser() DeviceInformationGetDevice200ResponseUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DeviceInformationGetDevice200Response) GetUserOk() (*DeviceInformationGetDevice200ResponseUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DeviceInformationGetDevice200Response) SetUser(v DeviceInformationGetDevice200ResponseUser)`

SetUser sets User field to given value.

### HasUser

`func (o *DeviceInformationGetDevice200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


