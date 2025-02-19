# InlineObject10

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**General** | Pointer to [**InlineObject10General**](InlineObject10General.md) |  | [optional] 
**Mdm** | Pointer to [**InlineObject10Mdm**](InlineObject10Mdm.md) |  | [optional] 
**ActivationLock** | Pointer to [**InlineObject10ActivationLock**](InlineObject10ActivationLock.md) |  | [optional] 
**Filevault** | Pointer to [**InlineObject10Filevault**](InlineObject10Filevault.md) |  | [optional] 
**LostMode** | Pointer to [**InlineObject10LostMode**](InlineObject10LostMode.md) |  | [optional] 
**AutomatedDeviceEnrollment** | Pointer to [**InlineObject10AutomatedDeviceEnrollment**](InlineObject10AutomatedDeviceEnrollment.md) |  | [optional] 
**KandjiAgent** | Pointer to [**InlineObject10KandjiAgent**](InlineObject10KandjiAgent.md) |  | [optional] 
**HardwareOverview** | Pointer to [**InlineObject10HardwareOverview**](InlineObject10HardwareOverview.md) |  | [optional] 
**Volumes** | Pointer to **interface{}** |  | [optional] 
**Network** | Pointer to **map[string]interface{}** |  | [optional] 
**RecoveryInformation** | Pointer to [**InlineObject10RecoveryInformation**](InlineObject10RecoveryInformation.md) |  | [optional] 
**Users** | Pointer to [**InlineObject10Users**](InlineObject10Users.md) |  | [optional] 
**InstalledProfiles** | Pointer to **interface{}** |  | [optional] 
**AppleBusinessManager** | Pointer to [**InlineObject10AppleBusinessManager**](InlineObject10AppleBusinessManager.md) |  | [optional] 
**SecurityInformation** | Pointer to [**InlineObject10SecurityInformation**](InlineObject10SecurityInformation.md) |  | [optional] 
**Cellular** | Pointer to [**InlineObject10Cellular**](InlineObject10Cellular.md) |  | [optional] 
**Tags** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewInlineObject10

`func NewInlineObject10() *InlineObject10`

NewInlineObject10 instantiates a new InlineObject10 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject10WithDefaults

`func NewInlineObject10WithDefaults() *InlineObject10`

NewInlineObject10WithDefaults instantiates a new InlineObject10 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneral

`func (o *InlineObject10) GetGeneral() InlineObject10General`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *InlineObject10) GetGeneralOk() (*InlineObject10General, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *InlineObject10) SetGeneral(v InlineObject10General)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *InlineObject10) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetMdm

`func (o *InlineObject10) GetMdm() InlineObject10Mdm`

GetMdm returns the Mdm field if non-nil, zero value otherwise.

### GetMdmOk

`func (o *InlineObject10) GetMdmOk() (*InlineObject10Mdm, bool)`

GetMdmOk returns a tuple with the Mdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdm

`func (o *InlineObject10) SetMdm(v InlineObject10Mdm)`

SetMdm sets Mdm field to given value.

### HasMdm

`func (o *InlineObject10) HasMdm() bool`

HasMdm returns a boolean if a field has been set.

### GetActivationLock

`func (o *InlineObject10) GetActivationLock() InlineObject10ActivationLock`

GetActivationLock returns the ActivationLock field if non-nil, zero value otherwise.

### GetActivationLockOk

`func (o *InlineObject10) GetActivationLockOk() (*InlineObject10ActivationLock, bool)`

GetActivationLockOk returns a tuple with the ActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationLock

`func (o *InlineObject10) SetActivationLock(v InlineObject10ActivationLock)`

SetActivationLock sets ActivationLock field to given value.

### HasActivationLock

`func (o *InlineObject10) HasActivationLock() bool`

HasActivationLock returns a boolean if a field has been set.

### GetFilevault

`func (o *InlineObject10) GetFilevault() InlineObject10Filevault`

GetFilevault returns the Filevault field if non-nil, zero value otherwise.

### GetFilevaultOk

`func (o *InlineObject10) GetFilevaultOk() (*InlineObject10Filevault, bool)`

GetFilevaultOk returns a tuple with the Filevault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilevault

`func (o *InlineObject10) SetFilevault(v InlineObject10Filevault)`

SetFilevault sets Filevault field to given value.

### HasFilevault

`func (o *InlineObject10) HasFilevault() bool`

HasFilevault returns a boolean if a field has been set.

### GetLostMode

`func (o *InlineObject10) GetLostMode() InlineObject10LostMode`

GetLostMode returns the LostMode field if non-nil, zero value otherwise.

### GetLostModeOk

`func (o *InlineObject10) GetLostModeOk() (*InlineObject10LostMode, bool)`

GetLostModeOk returns a tuple with the LostMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostMode

`func (o *InlineObject10) SetLostMode(v InlineObject10LostMode)`

SetLostMode sets LostMode field to given value.

### HasLostMode

`func (o *InlineObject10) HasLostMode() bool`

HasLostMode returns a boolean if a field has been set.

### GetAutomatedDeviceEnrollment

`func (o *InlineObject10) GetAutomatedDeviceEnrollment() InlineObject10AutomatedDeviceEnrollment`

GetAutomatedDeviceEnrollment returns the AutomatedDeviceEnrollment field if non-nil, zero value otherwise.

### GetAutomatedDeviceEnrollmentOk

`func (o *InlineObject10) GetAutomatedDeviceEnrollmentOk() (*InlineObject10AutomatedDeviceEnrollment, bool)`

GetAutomatedDeviceEnrollmentOk returns a tuple with the AutomatedDeviceEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedDeviceEnrollment

`func (o *InlineObject10) SetAutomatedDeviceEnrollment(v InlineObject10AutomatedDeviceEnrollment)`

SetAutomatedDeviceEnrollment sets AutomatedDeviceEnrollment field to given value.

### HasAutomatedDeviceEnrollment

`func (o *InlineObject10) HasAutomatedDeviceEnrollment() bool`

HasAutomatedDeviceEnrollment returns a boolean if a field has been set.

### GetKandjiAgent

`func (o *InlineObject10) GetKandjiAgent() InlineObject10KandjiAgent`

GetKandjiAgent returns the KandjiAgent field if non-nil, zero value otherwise.

### GetKandjiAgentOk

`func (o *InlineObject10) GetKandjiAgentOk() (*InlineObject10KandjiAgent, bool)`

GetKandjiAgentOk returns a tuple with the KandjiAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKandjiAgent

`func (o *InlineObject10) SetKandjiAgent(v InlineObject10KandjiAgent)`

SetKandjiAgent sets KandjiAgent field to given value.

### HasKandjiAgent

`func (o *InlineObject10) HasKandjiAgent() bool`

HasKandjiAgent returns a boolean if a field has been set.

### GetHardwareOverview

`func (o *InlineObject10) GetHardwareOverview() InlineObject10HardwareOverview`

GetHardwareOverview returns the HardwareOverview field if non-nil, zero value otherwise.

### GetHardwareOverviewOk

`func (o *InlineObject10) GetHardwareOverviewOk() (*InlineObject10HardwareOverview, bool)`

GetHardwareOverviewOk returns a tuple with the HardwareOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareOverview

`func (o *InlineObject10) SetHardwareOverview(v InlineObject10HardwareOverview)`

SetHardwareOverview sets HardwareOverview field to given value.

### HasHardwareOverview

`func (o *InlineObject10) HasHardwareOverview() bool`

HasHardwareOverview returns a boolean if a field has been set.

### GetVolumes

`func (o *InlineObject10) GetVolumes() interface{}`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *InlineObject10) GetVolumesOk() (*interface{}, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *InlineObject10) SetVolumes(v interface{})`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *InlineObject10) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *InlineObject10) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *InlineObject10) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil
### GetNetwork

`func (o *InlineObject10) GetNetwork() map[string]interface{}`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineObject10) GetNetworkOk() (*map[string]interface{}, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineObject10) SetNetwork(v map[string]interface{})`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineObject10) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRecoveryInformation

`func (o *InlineObject10) GetRecoveryInformation() InlineObject10RecoveryInformation`

GetRecoveryInformation returns the RecoveryInformation field if non-nil, zero value otherwise.

### GetRecoveryInformationOk

`func (o *InlineObject10) GetRecoveryInformationOk() (*InlineObject10RecoveryInformation, bool)`

GetRecoveryInformationOk returns a tuple with the RecoveryInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryInformation

`func (o *InlineObject10) SetRecoveryInformation(v InlineObject10RecoveryInformation)`

SetRecoveryInformation sets RecoveryInformation field to given value.

### HasRecoveryInformation

`func (o *InlineObject10) HasRecoveryInformation() bool`

HasRecoveryInformation returns a boolean if a field has been set.

### GetUsers

`func (o *InlineObject10) GetUsers() InlineObject10Users`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *InlineObject10) GetUsersOk() (*InlineObject10Users, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *InlineObject10) SetUsers(v InlineObject10Users)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *InlineObject10) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetInstalledProfiles

`func (o *InlineObject10) GetInstalledProfiles() interface{}`

GetInstalledProfiles returns the InstalledProfiles field if non-nil, zero value otherwise.

### GetInstalledProfilesOk

`func (o *InlineObject10) GetInstalledProfilesOk() (*interface{}, bool)`

GetInstalledProfilesOk returns a tuple with the InstalledProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledProfiles

`func (o *InlineObject10) SetInstalledProfiles(v interface{})`

SetInstalledProfiles sets InstalledProfiles field to given value.

### HasInstalledProfiles

`func (o *InlineObject10) HasInstalledProfiles() bool`

HasInstalledProfiles returns a boolean if a field has been set.

### SetInstalledProfilesNil

`func (o *InlineObject10) SetInstalledProfilesNil(b bool)`

 SetInstalledProfilesNil sets the value for InstalledProfiles to be an explicit nil

### UnsetInstalledProfiles
`func (o *InlineObject10) UnsetInstalledProfiles()`

UnsetInstalledProfiles ensures that no value is present for InstalledProfiles, not even an explicit nil
### GetAppleBusinessManager

`func (o *InlineObject10) GetAppleBusinessManager() InlineObject10AppleBusinessManager`

GetAppleBusinessManager returns the AppleBusinessManager field if non-nil, zero value otherwise.

### GetAppleBusinessManagerOk

`func (o *InlineObject10) GetAppleBusinessManagerOk() (*InlineObject10AppleBusinessManager, bool)`

GetAppleBusinessManagerOk returns a tuple with the AppleBusinessManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleBusinessManager

`func (o *InlineObject10) SetAppleBusinessManager(v InlineObject10AppleBusinessManager)`

SetAppleBusinessManager sets AppleBusinessManager field to given value.

### HasAppleBusinessManager

`func (o *InlineObject10) HasAppleBusinessManager() bool`

HasAppleBusinessManager returns a boolean if a field has been set.

### GetSecurityInformation

`func (o *InlineObject10) GetSecurityInformation() InlineObject10SecurityInformation`

GetSecurityInformation returns the SecurityInformation field if non-nil, zero value otherwise.

### GetSecurityInformationOk

`func (o *InlineObject10) GetSecurityInformationOk() (*InlineObject10SecurityInformation, bool)`

GetSecurityInformationOk returns a tuple with the SecurityInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityInformation

`func (o *InlineObject10) SetSecurityInformation(v InlineObject10SecurityInformation)`

SetSecurityInformation sets SecurityInformation field to given value.

### HasSecurityInformation

`func (o *InlineObject10) HasSecurityInformation() bool`

HasSecurityInformation returns a boolean if a field has been set.

### GetCellular

`func (o *InlineObject10) GetCellular() InlineObject10Cellular`

GetCellular returns the Cellular field if non-nil, zero value otherwise.

### GetCellularOk

`func (o *InlineObject10) GetCellularOk() (*InlineObject10Cellular, bool)`

GetCellularOk returns a tuple with the Cellular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellular

`func (o *InlineObject10) SetCellular(v InlineObject10Cellular)`

SetCellular sets Cellular field to given value.

### HasCellular

`func (o *InlineObject10) HasCellular() bool`

HasCellular returns a boolean if a field has been set.

### GetTags

`func (o *InlineObject10) GetTags() interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject10) GetTagsOk() (*interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject10) SetTags(v interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject10) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *InlineObject10) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *InlineObject10) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


