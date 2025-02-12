# DeviceInformationGetDeviceDetails200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**General** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseGeneral**](DeviceInformationGetDeviceDetails200ResponseGeneral.md) |  | [optional] 
**Mdm** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseMdm**](DeviceInformationGetDeviceDetails200ResponseMdm.md) |  | [optional] 
**ActivationLock** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseActivationLock**](DeviceInformationGetDeviceDetails200ResponseActivationLock.md) |  | [optional] 
**Filevault** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseFilevault**](DeviceInformationGetDeviceDetails200ResponseFilevault.md) |  | [optional] 
**LostMode** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseLostMode**](DeviceInformationGetDeviceDetails200ResponseLostMode.md) |  | [optional] 
**AutomatedDeviceEnrollment** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseAutomatedDeviceEnrollment**](DeviceInformationGetDeviceDetails200ResponseAutomatedDeviceEnrollment.md) |  | [optional] 
**KandjiAgent** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseKandjiAgent**](DeviceInformationGetDeviceDetails200ResponseKandjiAgent.md) |  | [optional] 
**HardwareOverview** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseHardwareOverview**](DeviceInformationGetDeviceDetails200ResponseHardwareOverview.md) |  | [optional] 
**Volumes** | Pointer to **interface{}** |  | [optional] 
**Network** | Pointer to **map[string]interface{}** |  | [optional] 
**RecoveryInformation** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseRecoveryInformation**](DeviceInformationGetDeviceDetails200ResponseRecoveryInformation.md) |  | [optional] 
**Users** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseUsers**](DeviceInformationGetDeviceDetails200ResponseUsers.md) |  | [optional] 
**InstalledProfiles** | Pointer to **interface{}** |  | [optional] 
**AppleBusinessManager** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseAppleBusinessManager**](DeviceInformationGetDeviceDetails200ResponseAppleBusinessManager.md) |  | [optional] 
**SecurityInformation** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseSecurityInformation**](DeviceInformationGetDeviceDetails200ResponseSecurityInformation.md) |  | [optional] 
**Cellular** | Pointer to [**DeviceInformationGetDeviceDetails200ResponseCellular**](DeviceInformationGetDeviceDetails200ResponseCellular.md) |  | [optional] 
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

### GetMdm

`func (o *DeviceInformationGetDeviceDetails200Response) GetMdm() DeviceInformationGetDeviceDetails200ResponseMdm`

GetMdm returns the Mdm field if non-nil, zero value otherwise.

### GetMdmOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetMdmOk() (*DeviceInformationGetDeviceDetails200ResponseMdm, bool)`

GetMdmOk returns a tuple with the Mdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdm

`func (o *DeviceInformationGetDeviceDetails200Response) SetMdm(v DeviceInformationGetDeviceDetails200ResponseMdm)`

SetMdm sets Mdm field to given value.

### HasMdm

`func (o *DeviceInformationGetDeviceDetails200Response) HasMdm() bool`

HasMdm returns a boolean if a field has been set.

### GetActivationLock

`func (o *DeviceInformationGetDeviceDetails200Response) GetActivationLock() DeviceInformationGetDeviceDetails200ResponseActivationLock`

GetActivationLock returns the ActivationLock field if non-nil, zero value otherwise.

### GetActivationLockOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetActivationLockOk() (*DeviceInformationGetDeviceDetails200ResponseActivationLock, bool)`

GetActivationLockOk returns a tuple with the ActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationLock

`func (o *DeviceInformationGetDeviceDetails200Response) SetActivationLock(v DeviceInformationGetDeviceDetails200ResponseActivationLock)`

SetActivationLock sets ActivationLock field to given value.

### HasActivationLock

`func (o *DeviceInformationGetDeviceDetails200Response) HasActivationLock() bool`

HasActivationLock returns a boolean if a field has been set.

### GetFilevault

`func (o *DeviceInformationGetDeviceDetails200Response) GetFilevault() DeviceInformationGetDeviceDetails200ResponseFilevault`

GetFilevault returns the Filevault field if non-nil, zero value otherwise.

### GetFilevaultOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetFilevaultOk() (*DeviceInformationGetDeviceDetails200ResponseFilevault, bool)`

GetFilevaultOk returns a tuple with the Filevault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilevault

`func (o *DeviceInformationGetDeviceDetails200Response) SetFilevault(v DeviceInformationGetDeviceDetails200ResponseFilevault)`

SetFilevault sets Filevault field to given value.

### HasFilevault

`func (o *DeviceInformationGetDeviceDetails200Response) HasFilevault() bool`

HasFilevault returns a boolean if a field has been set.

### GetLostMode

`func (o *DeviceInformationGetDeviceDetails200Response) GetLostMode() DeviceInformationGetDeviceDetails200ResponseLostMode`

GetLostMode returns the LostMode field if non-nil, zero value otherwise.

### GetLostModeOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetLostModeOk() (*DeviceInformationGetDeviceDetails200ResponseLostMode, bool)`

GetLostModeOk returns a tuple with the LostMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostMode

`func (o *DeviceInformationGetDeviceDetails200Response) SetLostMode(v DeviceInformationGetDeviceDetails200ResponseLostMode)`

SetLostMode sets LostMode field to given value.

### HasLostMode

`func (o *DeviceInformationGetDeviceDetails200Response) HasLostMode() bool`

HasLostMode returns a boolean if a field has been set.

### GetAutomatedDeviceEnrollment

`func (o *DeviceInformationGetDeviceDetails200Response) GetAutomatedDeviceEnrollment() DeviceInformationGetDeviceDetails200ResponseAutomatedDeviceEnrollment`

GetAutomatedDeviceEnrollment returns the AutomatedDeviceEnrollment field if non-nil, zero value otherwise.

### GetAutomatedDeviceEnrollmentOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetAutomatedDeviceEnrollmentOk() (*DeviceInformationGetDeviceDetails200ResponseAutomatedDeviceEnrollment, bool)`

GetAutomatedDeviceEnrollmentOk returns a tuple with the AutomatedDeviceEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedDeviceEnrollment

`func (o *DeviceInformationGetDeviceDetails200Response) SetAutomatedDeviceEnrollment(v DeviceInformationGetDeviceDetails200ResponseAutomatedDeviceEnrollment)`

SetAutomatedDeviceEnrollment sets AutomatedDeviceEnrollment field to given value.

### HasAutomatedDeviceEnrollment

`func (o *DeviceInformationGetDeviceDetails200Response) HasAutomatedDeviceEnrollment() bool`

HasAutomatedDeviceEnrollment returns a boolean if a field has been set.

### GetKandjiAgent

`func (o *DeviceInformationGetDeviceDetails200Response) GetKandjiAgent() DeviceInformationGetDeviceDetails200ResponseKandjiAgent`

GetKandjiAgent returns the KandjiAgent field if non-nil, zero value otherwise.

### GetKandjiAgentOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetKandjiAgentOk() (*DeviceInformationGetDeviceDetails200ResponseKandjiAgent, bool)`

GetKandjiAgentOk returns a tuple with the KandjiAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKandjiAgent

`func (o *DeviceInformationGetDeviceDetails200Response) SetKandjiAgent(v DeviceInformationGetDeviceDetails200ResponseKandjiAgent)`

SetKandjiAgent sets KandjiAgent field to given value.

### HasKandjiAgent

`func (o *DeviceInformationGetDeviceDetails200Response) HasKandjiAgent() bool`

HasKandjiAgent returns a boolean if a field has been set.

### GetHardwareOverview

`func (o *DeviceInformationGetDeviceDetails200Response) GetHardwareOverview() DeviceInformationGetDeviceDetails200ResponseHardwareOverview`

GetHardwareOverview returns the HardwareOverview field if non-nil, zero value otherwise.

### GetHardwareOverviewOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetHardwareOverviewOk() (*DeviceInformationGetDeviceDetails200ResponseHardwareOverview, bool)`

GetHardwareOverviewOk returns a tuple with the HardwareOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareOverview

`func (o *DeviceInformationGetDeviceDetails200Response) SetHardwareOverview(v DeviceInformationGetDeviceDetails200ResponseHardwareOverview)`

SetHardwareOverview sets HardwareOverview field to given value.

### HasHardwareOverview

`func (o *DeviceInformationGetDeviceDetails200Response) HasHardwareOverview() bool`

HasHardwareOverview returns a boolean if a field has been set.

### GetVolumes

`func (o *DeviceInformationGetDeviceDetails200Response) GetVolumes() interface{}`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetVolumesOk() (*interface{}, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *DeviceInformationGetDeviceDetails200Response) SetVolumes(v interface{})`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *DeviceInformationGetDeviceDetails200Response) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *DeviceInformationGetDeviceDetails200Response) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *DeviceInformationGetDeviceDetails200Response) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil
### GetNetwork

`func (o *DeviceInformationGetDeviceDetails200Response) GetNetwork() map[string]interface{}`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetNetworkOk() (*map[string]interface{}, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *DeviceInformationGetDeviceDetails200Response) SetNetwork(v map[string]interface{})`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *DeviceInformationGetDeviceDetails200Response) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRecoveryInformation

`func (o *DeviceInformationGetDeviceDetails200Response) GetRecoveryInformation() DeviceInformationGetDeviceDetails200ResponseRecoveryInformation`

GetRecoveryInformation returns the RecoveryInformation field if non-nil, zero value otherwise.

### GetRecoveryInformationOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetRecoveryInformationOk() (*DeviceInformationGetDeviceDetails200ResponseRecoveryInformation, bool)`

GetRecoveryInformationOk returns a tuple with the RecoveryInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryInformation

`func (o *DeviceInformationGetDeviceDetails200Response) SetRecoveryInformation(v DeviceInformationGetDeviceDetails200ResponseRecoveryInformation)`

SetRecoveryInformation sets RecoveryInformation field to given value.

### HasRecoveryInformation

`func (o *DeviceInformationGetDeviceDetails200Response) HasRecoveryInformation() bool`

HasRecoveryInformation returns a boolean if a field has been set.

### GetUsers

`func (o *DeviceInformationGetDeviceDetails200Response) GetUsers() DeviceInformationGetDeviceDetails200ResponseUsers`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetUsersOk() (*DeviceInformationGetDeviceDetails200ResponseUsers, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *DeviceInformationGetDeviceDetails200Response) SetUsers(v DeviceInformationGetDeviceDetails200ResponseUsers)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *DeviceInformationGetDeviceDetails200Response) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetInstalledProfiles

`func (o *DeviceInformationGetDeviceDetails200Response) GetInstalledProfiles() interface{}`

GetInstalledProfiles returns the InstalledProfiles field if non-nil, zero value otherwise.

### GetInstalledProfilesOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetInstalledProfilesOk() (*interface{}, bool)`

GetInstalledProfilesOk returns a tuple with the InstalledProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledProfiles

`func (o *DeviceInformationGetDeviceDetails200Response) SetInstalledProfiles(v interface{})`

SetInstalledProfiles sets InstalledProfiles field to given value.

### HasInstalledProfiles

`func (o *DeviceInformationGetDeviceDetails200Response) HasInstalledProfiles() bool`

HasInstalledProfiles returns a boolean if a field has been set.

### SetInstalledProfilesNil

`func (o *DeviceInformationGetDeviceDetails200Response) SetInstalledProfilesNil(b bool)`

 SetInstalledProfilesNil sets the value for InstalledProfiles to be an explicit nil

### UnsetInstalledProfiles
`func (o *DeviceInformationGetDeviceDetails200Response) UnsetInstalledProfiles()`

UnsetInstalledProfiles ensures that no value is present for InstalledProfiles, not even an explicit nil
### GetAppleBusinessManager

`func (o *DeviceInformationGetDeviceDetails200Response) GetAppleBusinessManager() DeviceInformationGetDeviceDetails200ResponseAppleBusinessManager`

GetAppleBusinessManager returns the AppleBusinessManager field if non-nil, zero value otherwise.

### GetAppleBusinessManagerOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetAppleBusinessManagerOk() (*DeviceInformationGetDeviceDetails200ResponseAppleBusinessManager, bool)`

GetAppleBusinessManagerOk returns a tuple with the AppleBusinessManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleBusinessManager

`func (o *DeviceInformationGetDeviceDetails200Response) SetAppleBusinessManager(v DeviceInformationGetDeviceDetails200ResponseAppleBusinessManager)`

SetAppleBusinessManager sets AppleBusinessManager field to given value.

### HasAppleBusinessManager

`func (o *DeviceInformationGetDeviceDetails200Response) HasAppleBusinessManager() bool`

HasAppleBusinessManager returns a boolean if a field has been set.

### GetSecurityInformation

`func (o *DeviceInformationGetDeviceDetails200Response) GetSecurityInformation() DeviceInformationGetDeviceDetails200ResponseSecurityInformation`

GetSecurityInformation returns the SecurityInformation field if non-nil, zero value otherwise.

### GetSecurityInformationOk

`func (o *DeviceInformationGetDeviceDetails200Response) GetSecurityInformationOk() (*DeviceInformationGetDeviceDetails200ResponseSecurityInformation, bool)`

GetSecurityInformationOk returns a tuple with the SecurityInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityInformation

`func (o *DeviceInformationGetDeviceDetails200Response) SetSecurityInformation(v DeviceInformationGetDeviceDetails200ResponseSecurityInformation)`

SetSecurityInformation sets SecurityInformation field to given value.

### HasSecurityInformation

`func (o *DeviceInformationGetDeviceDetails200Response) HasSecurityInformation() bool`

HasSecurityInformation returns a boolean if a field has been set.

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


