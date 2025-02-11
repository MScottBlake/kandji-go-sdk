# GetIpad200Response

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
**User** | Pointer to [**GetIpad200ResponseUser**](GetIpad200ResponseUser.md) |  | [optional] 
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

### NewGetIpad200Response

`func NewGetIpad200Response() *GetIpad200Response`

NewGetIpad200Response instantiates a new GetIpad200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIpad200ResponseWithDefaults

`func NewGetIpad200ResponseWithDefaults() *GetIpad200Response`

NewGetIpad200ResponseWithDefaults instantiates a new GetIpad200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *GetIpad200Response) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GetIpad200Response) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GetIpad200Response) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *GetIpad200Response) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *GetIpad200Response) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *GetIpad200Response) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *GetIpad200Response) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *GetIpad200Response) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetModel

`func (o *GetIpad200Response) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetIpad200Response) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetIpad200Response) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetIpad200Response) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSerialNumber

`func (o *GetIpad200Response) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *GetIpad200Response) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *GetIpad200Response) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *GetIpad200Response) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetPlatform

`func (o *GetIpad200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetIpad200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetIpad200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetIpad200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetOsVersion

`func (o *GetIpad200Response) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *GetIpad200Response) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *GetIpad200Response) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *GetIpad200Response) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetSupplementalBuildVersion

`func (o *GetIpad200Response) GetSupplementalBuildVersion() string`

GetSupplementalBuildVersion returns the SupplementalBuildVersion field if non-nil, zero value otherwise.

### GetSupplementalBuildVersionOk

`func (o *GetIpad200Response) GetSupplementalBuildVersionOk() (*string, bool)`

GetSupplementalBuildVersionOk returns a tuple with the SupplementalBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementalBuildVersion

`func (o *GetIpad200Response) SetSupplementalBuildVersion(v string)`

SetSupplementalBuildVersion sets SupplementalBuildVersion field to given value.

### HasSupplementalBuildVersion

`func (o *GetIpad200Response) HasSupplementalBuildVersion() bool`

HasSupplementalBuildVersion returns a boolean if a field has been set.

### GetSupplementalOsVersionExtra

`func (o *GetIpad200Response) GetSupplementalOsVersionExtra() string`

GetSupplementalOsVersionExtra returns the SupplementalOsVersionExtra field if non-nil, zero value otherwise.

### GetSupplementalOsVersionExtraOk

`func (o *GetIpad200Response) GetSupplementalOsVersionExtraOk() (*string, bool)`

GetSupplementalOsVersionExtraOk returns a tuple with the SupplementalOsVersionExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementalOsVersionExtra

`func (o *GetIpad200Response) SetSupplementalOsVersionExtra(v string)`

SetSupplementalOsVersionExtra sets SupplementalOsVersionExtra field to given value.

### HasSupplementalOsVersionExtra

`func (o *GetIpad200Response) HasSupplementalOsVersionExtra() bool`

HasSupplementalOsVersionExtra returns a boolean if a field has been set.

### GetLastCheckIn

`func (o *GetIpad200Response) GetLastCheckIn() string`

GetLastCheckIn returns the LastCheckIn field if non-nil, zero value otherwise.

### GetLastCheckInOk

`func (o *GetIpad200Response) GetLastCheckInOk() (*string, bool)`

GetLastCheckInOk returns a tuple with the LastCheckIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckIn

`func (o *GetIpad200Response) SetLastCheckIn(v string)`

SetLastCheckIn sets LastCheckIn field to given value.

### HasLastCheckIn

`func (o *GetIpad200Response) HasLastCheckIn() bool`

HasLastCheckIn returns a boolean if a field has been set.

### GetUser

`func (o *GetIpad200Response) GetUser() GetIpad200ResponseUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetIpad200Response) GetUserOk() (*GetIpad200ResponseUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetIpad200Response) SetUser(v GetIpad200ResponseUser)`

SetUser sets User field to given value.

### HasUser

`func (o *GetIpad200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAssetTag

`func (o *GetIpad200Response) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *GetIpad200Response) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *GetIpad200Response) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *GetIpad200Response) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetBlueprintId

`func (o *GetIpad200Response) GetBlueprintId() string`

GetBlueprintId returns the BlueprintId field if non-nil, zero value otherwise.

### GetBlueprintIdOk

`func (o *GetIpad200Response) GetBlueprintIdOk() (*string, bool)`

GetBlueprintIdOk returns a tuple with the BlueprintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintId

`func (o *GetIpad200Response) SetBlueprintId(v string)`

SetBlueprintId sets BlueprintId field to given value.

### HasBlueprintId

`func (o *GetIpad200Response) HasBlueprintId() bool`

HasBlueprintId returns a boolean if a field has been set.

### GetMdmEnabled

`func (o *GetIpad200Response) GetMdmEnabled() int32`

GetMdmEnabled returns the MdmEnabled field if non-nil, zero value otherwise.

### GetMdmEnabledOk

`func (o *GetIpad200Response) GetMdmEnabledOk() (*int32, bool)`

GetMdmEnabledOk returns a tuple with the MdmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmEnabled

`func (o *GetIpad200Response) SetMdmEnabled(v int32)`

SetMdmEnabled sets MdmEnabled field to given value.

### HasMdmEnabled

`func (o *GetIpad200Response) HasMdmEnabled() bool`

HasMdmEnabled returns a boolean if a field has been set.

### GetAgentInstalled

`func (o *GetIpad200Response) GetAgentInstalled() int32`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *GetIpad200Response) GetAgentInstalledOk() (*int32, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *GetIpad200Response) SetAgentInstalled(v int32)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *GetIpad200Response) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetIsMissing

`func (o *GetIpad200Response) GetIsMissing() int32`

GetIsMissing returns the IsMissing field if non-nil, zero value otherwise.

### GetIsMissingOk

`func (o *GetIpad200Response) GetIsMissingOk() (*int32, bool)`

GetIsMissingOk returns a tuple with the IsMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMissing

`func (o *GetIpad200Response) SetIsMissing(v int32)`

SetIsMissing sets IsMissing field to given value.

### HasIsMissing

`func (o *GetIpad200Response) HasIsMissing() bool`

HasIsMissing returns a boolean if a field has been set.

### GetIsRemoved

`func (o *GetIpad200Response) GetIsRemoved() int32`

GetIsRemoved returns the IsRemoved field if non-nil, zero value otherwise.

### GetIsRemovedOk

`func (o *GetIpad200Response) GetIsRemovedOk() (*int32, bool)`

GetIsRemovedOk returns a tuple with the IsRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoved

`func (o *GetIpad200Response) SetIsRemoved(v int32)`

SetIsRemoved sets IsRemoved field to given value.

### HasIsRemoved

`func (o *GetIpad200Response) HasIsRemoved() bool`

HasIsRemoved returns a boolean if a field has been set.

### GetAgentVersion

`func (o *GetIpad200Response) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *GetIpad200Response) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *GetIpad200Response) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *GetIpad200Response) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetFirstEnrollment

`func (o *GetIpad200Response) GetFirstEnrollment() string`

GetFirstEnrollment returns the FirstEnrollment field if non-nil, zero value otherwise.

### GetFirstEnrollmentOk

`func (o *GetIpad200Response) GetFirstEnrollmentOk() (*string, bool)`

GetFirstEnrollmentOk returns a tuple with the FirstEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstEnrollment

`func (o *GetIpad200Response) SetFirstEnrollment(v string)`

SetFirstEnrollment sets FirstEnrollment field to given value.

### HasFirstEnrollment

`func (o *GetIpad200Response) HasFirstEnrollment() bool`

HasFirstEnrollment returns a boolean if a field has been set.

### GetLastEnrollment

`func (o *GetIpad200Response) GetLastEnrollment() string`

GetLastEnrollment returns the LastEnrollment field if non-nil, zero value otherwise.

### GetLastEnrollmentOk

`func (o *GetIpad200Response) GetLastEnrollmentOk() (*string, bool)`

GetLastEnrollmentOk returns a tuple with the LastEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnrollment

`func (o *GetIpad200Response) SetLastEnrollment(v string)`

SetLastEnrollment sets LastEnrollment field to given value.

### HasLastEnrollment

`func (o *GetIpad200Response) HasLastEnrollment() bool`

HasLastEnrollment returns a boolean if a field has been set.

### GetBlueprintName

`func (o *GetIpad200Response) GetBlueprintName() string`

GetBlueprintName returns the BlueprintName field if non-nil, zero value otherwise.

### GetBlueprintNameOk

`func (o *GetIpad200Response) GetBlueprintNameOk() (*string, bool)`

GetBlueprintNameOk returns a tuple with the BlueprintName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintName

`func (o *GetIpad200Response) SetBlueprintName(v string)`

SetBlueprintName sets BlueprintName field to given value.

### HasBlueprintName

`func (o *GetIpad200Response) HasBlueprintName() bool`

HasBlueprintName returns a boolean if a field has been set.

### GetTags

`func (o *GetIpad200Response) GetTags() interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetIpad200Response) GetTagsOk() (*interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetIpad200Response) SetTags(v interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetIpad200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *GetIpad200Response) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *GetIpad200Response) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


