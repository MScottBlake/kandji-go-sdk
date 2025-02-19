# InlineObject8

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**SupplementalBuildVersion** | Pointer to **string** |  | [optional] 
**SupplementalOsVersionExtra** | Pointer to **string** |  | [optional] 
**LastCheckIn** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**InlineObject8User**](InlineObject8User.md) |  | [optional] 
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
**Tags** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewInlineObject8

`func NewInlineObject8() *InlineObject8`

NewInlineObject8 instantiates a new InlineObject8 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject8WithDefaults

`func NewInlineObject8WithDefaults() *InlineObject8`

NewInlineObject8WithDefaults instantiates a new InlineObject8 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *InlineObject8) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *InlineObject8) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *InlineObject8) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *InlineObject8) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *InlineObject8) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *InlineObject8) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *InlineObject8) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *InlineObject8) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetModel

`func (o *InlineObject8) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineObject8) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineObject8) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineObject8) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSerialNumber

`func (o *InlineObject8) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *InlineObject8) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *InlineObject8) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *InlineObject8) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetPlatform

`func (o *InlineObject8) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *InlineObject8) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *InlineObject8) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *InlineObject8) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetOsVersion

`func (o *InlineObject8) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *InlineObject8) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *InlineObject8) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *InlineObject8) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetSupplementalBuildVersion

`func (o *InlineObject8) GetSupplementalBuildVersion() string`

GetSupplementalBuildVersion returns the SupplementalBuildVersion field if non-nil, zero value otherwise.

### GetSupplementalBuildVersionOk

`func (o *InlineObject8) GetSupplementalBuildVersionOk() (*string, bool)`

GetSupplementalBuildVersionOk returns a tuple with the SupplementalBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementalBuildVersion

`func (o *InlineObject8) SetSupplementalBuildVersion(v string)`

SetSupplementalBuildVersion sets SupplementalBuildVersion field to given value.

### HasSupplementalBuildVersion

`func (o *InlineObject8) HasSupplementalBuildVersion() bool`

HasSupplementalBuildVersion returns a boolean if a field has been set.

### GetSupplementalOsVersionExtra

`func (o *InlineObject8) GetSupplementalOsVersionExtra() string`

GetSupplementalOsVersionExtra returns the SupplementalOsVersionExtra field if non-nil, zero value otherwise.

### GetSupplementalOsVersionExtraOk

`func (o *InlineObject8) GetSupplementalOsVersionExtraOk() (*string, bool)`

GetSupplementalOsVersionExtraOk returns a tuple with the SupplementalOsVersionExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementalOsVersionExtra

`func (o *InlineObject8) SetSupplementalOsVersionExtra(v string)`

SetSupplementalOsVersionExtra sets SupplementalOsVersionExtra field to given value.

### HasSupplementalOsVersionExtra

`func (o *InlineObject8) HasSupplementalOsVersionExtra() bool`

HasSupplementalOsVersionExtra returns a boolean if a field has been set.

### GetLastCheckIn

`func (o *InlineObject8) GetLastCheckIn() string`

GetLastCheckIn returns the LastCheckIn field if non-nil, zero value otherwise.

### GetLastCheckInOk

`func (o *InlineObject8) GetLastCheckInOk() (*string, bool)`

GetLastCheckInOk returns a tuple with the LastCheckIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckIn

`func (o *InlineObject8) SetLastCheckIn(v string)`

SetLastCheckIn sets LastCheckIn field to given value.

### HasLastCheckIn

`func (o *InlineObject8) HasLastCheckIn() bool`

HasLastCheckIn returns a boolean if a field has been set.

### GetUser

`func (o *InlineObject8) GetUser() InlineObject8User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *InlineObject8) GetUserOk() (*InlineObject8User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *InlineObject8) SetUser(v InlineObject8User)`

SetUser sets User field to given value.

### HasUser

`func (o *InlineObject8) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAssetTag

`func (o *InlineObject8) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *InlineObject8) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *InlineObject8) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *InlineObject8) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetBlueprintId

`func (o *InlineObject8) GetBlueprintId() string`

GetBlueprintId returns the BlueprintId field if non-nil, zero value otherwise.

### GetBlueprintIdOk

`func (o *InlineObject8) GetBlueprintIdOk() (*string, bool)`

GetBlueprintIdOk returns a tuple with the BlueprintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintId

`func (o *InlineObject8) SetBlueprintId(v string)`

SetBlueprintId sets BlueprintId field to given value.

### HasBlueprintId

`func (o *InlineObject8) HasBlueprintId() bool`

HasBlueprintId returns a boolean if a field has been set.

### GetMdmEnabled

`func (o *InlineObject8) GetMdmEnabled() int32`

GetMdmEnabled returns the MdmEnabled field if non-nil, zero value otherwise.

### GetMdmEnabledOk

`func (o *InlineObject8) GetMdmEnabledOk() (*int32, bool)`

GetMdmEnabledOk returns a tuple with the MdmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmEnabled

`func (o *InlineObject8) SetMdmEnabled(v int32)`

SetMdmEnabled sets MdmEnabled field to given value.

### HasMdmEnabled

`func (o *InlineObject8) HasMdmEnabled() bool`

HasMdmEnabled returns a boolean if a field has been set.

### GetAgentInstalled

`func (o *InlineObject8) GetAgentInstalled() int32`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *InlineObject8) GetAgentInstalledOk() (*int32, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *InlineObject8) SetAgentInstalled(v int32)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *InlineObject8) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetIsMissing

`func (o *InlineObject8) GetIsMissing() int32`

GetIsMissing returns the IsMissing field if non-nil, zero value otherwise.

### GetIsMissingOk

`func (o *InlineObject8) GetIsMissingOk() (*int32, bool)`

GetIsMissingOk returns a tuple with the IsMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMissing

`func (o *InlineObject8) SetIsMissing(v int32)`

SetIsMissing sets IsMissing field to given value.

### HasIsMissing

`func (o *InlineObject8) HasIsMissing() bool`

HasIsMissing returns a boolean if a field has been set.

### GetIsRemoved

`func (o *InlineObject8) GetIsRemoved() int32`

GetIsRemoved returns the IsRemoved field if non-nil, zero value otherwise.

### GetIsRemovedOk

`func (o *InlineObject8) GetIsRemovedOk() (*int32, bool)`

GetIsRemovedOk returns a tuple with the IsRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoved

`func (o *InlineObject8) SetIsRemoved(v int32)`

SetIsRemoved sets IsRemoved field to given value.

### HasIsRemoved

`func (o *InlineObject8) HasIsRemoved() bool`

HasIsRemoved returns a boolean if a field has been set.

### GetAgentVersion

`func (o *InlineObject8) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *InlineObject8) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *InlineObject8) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *InlineObject8) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetFirstEnrollment

`func (o *InlineObject8) GetFirstEnrollment() string`

GetFirstEnrollment returns the FirstEnrollment field if non-nil, zero value otherwise.

### GetFirstEnrollmentOk

`func (o *InlineObject8) GetFirstEnrollmentOk() (*string, bool)`

GetFirstEnrollmentOk returns a tuple with the FirstEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstEnrollment

`func (o *InlineObject8) SetFirstEnrollment(v string)`

SetFirstEnrollment sets FirstEnrollment field to given value.

### HasFirstEnrollment

`func (o *InlineObject8) HasFirstEnrollment() bool`

HasFirstEnrollment returns a boolean if a field has been set.

### GetLastEnrollment

`func (o *InlineObject8) GetLastEnrollment() string`

GetLastEnrollment returns the LastEnrollment field if non-nil, zero value otherwise.

### GetLastEnrollmentOk

`func (o *InlineObject8) GetLastEnrollmentOk() (*string, bool)`

GetLastEnrollmentOk returns a tuple with the LastEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnrollment

`func (o *InlineObject8) SetLastEnrollment(v string)`

SetLastEnrollment sets LastEnrollment field to given value.

### HasLastEnrollment

`func (o *InlineObject8) HasLastEnrollment() bool`

HasLastEnrollment returns a boolean if a field has been set.

### GetBlueprintName

`func (o *InlineObject8) GetBlueprintName() string`

GetBlueprintName returns the BlueprintName field if non-nil, zero value otherwise.

### GetBlueprintNameOk

`func (o *InlineObject8) GetBlueprintNameOk() (*string, bool)`

GetBlueprintNameOk returns a tuple with the BlueprintName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintName

`func (o *InlineObject8) SetBlueprintName(v string)`

SetBlueprintName sets BlueprintName field to given value.

### HasBlueprintName

`func (o *InlineObject8) HasBlueprintName() bool`

HasBlueprintName returns a boolean if a field has been set.

### GetTags

`func (o *InlineObject8) GetTags() interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject8) GetTagsOk() (*interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject8) SetTags(v interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject8) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *InlineObject8) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *InlineObject8) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


