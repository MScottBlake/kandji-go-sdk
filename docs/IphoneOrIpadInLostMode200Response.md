# IphoneOrIpadInLostMode200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**General** | Pointer to [**IphoneOrIpadInLostMode200ResponseGeneral**](IphoneOrIpadInLostMode200ResponseGeneral.md) |  | [optional] 
**Mdm** | Pointer to [**IphoneOrIpadInLostMode200ResponseMdm**](IphoneOrIpadInLostMode200ResponseMdm.md) |  | [optional] 
**ActivationLock** | Pointer to [**IphoneOrIpadInLostMode200ResponseActivationLock**](IphoneOrIpadInLostMode200ResponseActivationLock.md) |  | [optional] 
**Filevault** | Pointer to [**IphoneOrIpadInLostMode200ResponseFilevault**](IphoneOrIpadInLostMode200ResponseFilevault.md) |  | [optional] 
**LostMode** | Pointer to [**IphoneOrIpadInLostMode200ResponseLostMode**](IphoneOrIpadInLostMode200ResponseLostMode.md) |  | [optional] 
**AutomatedDeviceEnrollment** | Pointer to [**IphoneOrIpadInLostMode200ResponseAutomatedDeviceEnrollment**](IphoneOrIpadInLostMode200ResponseAutomatedDeviceEnrollment.md) |  | [optional] 
**KandjiAgent** | Pointer to [**IphoneOrIpadInLostMode200ResponseKandjiAgent**](IphoneOrIpadInLostMode200ResponseKandjiAgent.md) |  | [optional] 
**HardwareOverview** | Pointer to [**IphoneOrIpadInLostMode200ResponseHardwareOverview**](IphoneOrIpadInLostMode200ResponseHardwareOverview.md) |  | [optional] 
**Volumes** | Pointer to **interface{}** |  | [optional] 
**Network** | Pointer to **map[string]interface{}** |  | [optional] 
**RecoveryInformation** | Pointer to [**IphoneOrIpadInLostMode200ResponseRecoveryInformation**](IphoneOrIpadInLostMode200ResponseRecoveryInformation.md) |  | [optional] 
**Users** | Pointer to [**IphoneOrIpadInLostMode200ResponseUsers**](IphoneOrIpadInLostMode200ResponseUsers.md) |  | [optional] 
**InstalledProfiles** | Pointer to **interface{}** |  | [optional] 
**AppleBusinessManager** | Pointer to [**IphoneOrIpadInLostMode200ResponseAppleBusinessManager**](IphoneOrIpadInLostMode200ResponseAppleBusinessManager.md) |  | [optional] 
**SecurityInformation** | Pointer to [**IphoneOrIpadInLostMode200ResponseSecurityInformation**](IphoneOrIpadInLostMode200ResponseSecurityInformation.md) |  | [optional] 
**Cellular** | Pointer to [**IphoneOrIpadInLostMode200ResponseCellular**](IphoneOrIpadInLostMode200ResponseCellular.md) |  | [optional] 
**Tags** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewIphoneOrIpadInLostMode200Response

`func NewIphoneOrIpadInLostMode200Response() *IphoneOrIpadInLostMode200Response`

NewIphoneOrIpadInLostMode200Response instantiates a new IphoneOrIpadInLostMode200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIphoneOrIpadInLostMode200ResponseWithDefaults

`func NewIphoneOrIpadInLostMode200ResponseWithDefaults() *IphoneOrIpadInLostMode200Response`

NewIphoneOrIpadInLostMode200ResponseWithDefaults instantiates a new IphoneOrIpadInLostMode200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneral

`func (o *IphoneOrIpadInLostMode200Response) GetGeneral() IphoneOrIpadInLostMode200ResponseGeneral`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *IphoneOrIpadInLostMode200Response) GetGeneralOk() (*IphoneOrIpadInLostMode200ResponseGeneral, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *IphoneOrIpadInLostMode200Response) SetGeneral(v IphoneOrIpadInLostMode200ResponseGeneral)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *IphoneOrIpadInLostMode200Response) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetMdm

`func (o *IphoneOrIpadInLostMode200Response) GetMdm() IphoneOrIpadInLostMode200ResponseMdm`

GetMdm returns the Mdm field if non-nil, zero value otherwise.

### GetMdmOk

`func (o *IphoneOrIpadInLostMode200Response) GetMdmOk() (*IphoneOrIpadInLostMode200ResponseMdm, bool)`

GetMdmOk returns a tuple with the Mdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdm

`func (o *IphoneOrIpadInLostMode200Response) SetMdm(v IphoneOrIpadInLostMode200ResponseMdm)`

SetMdm sets Mdm field to given value.

### HasMdm

`func (o *IphoneOrIpadInLostMode200Response) HasMdm() bool`

HasMdm returns a boolean if a field has been set.

### GetActivationLock

`func (o *IphoneOrIpadInLostMode200Response) GetActivationLock() IphoneOrIpadInLostMode200ResponseActivationLock`

GetActivationLock returns the ActivationLock field if non-nil, zero value otherwise.

### GetActivationLockOk

`func (o *IphoneOrIpadInLostMode200Response) GetActivationLockOk() (*IphoneOrIpadInLostMode200ResponseActivationLock, bool)`

GetActivationLockOk returns a tuple with the ActivationLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationLock

`func (o *IphoneOrIpadInLostMode200Response) SetActivationLock(v IphoneOrIpadInLostMode200ResponseActivationLock)`

SetActivationLock sets ActivationLock field to given value.

### HasActivationLock

`func (o *IphoneOrIpadInLostMode200Response) HasActivationLock() bool`

HasActivationLock returns a boolean if a field has been set.

### GetFilevault

`func (o *IphoneOrIpadInLostMode200Response) GetFilevault() IphoneOrIpadInLostMode200ResponseFilevault`

GetFilevault returns the Filevault field if non-nil, zero value otherwise.

### GetFilevaultOk

`func (o *IphoneOrIpadInLostMode200Response) GetFilevaultOk() (*IphoneOrIpadInLostMode200ResponseFilevault, bool)`

GetFilevaultOk returns a tuple with the Filevault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilevault

`func (o *IphoneOrIpadInLostMode200Response) SetFilevault(v IphoneOrIpadInLostMode200ResponseFilevault)`

SetFilevault sets Filevault field to given value.

### HasFilevault

`func (o *IphoneOrIpadInLostMode200Response) HasFilevault() bool`

HasFilevault returns a boolean if a field has been set.

### GetLostMode

`func (o *IphoneOrIpadInLostMode200Response) GetLostMode() IphoneOrIpadInLostMode200ResponseLostMode`

GetLostMode returns the LostMode field if non-nil, zero value otherwise.

### GetLostModeOk

`func (o *IphoneOrIpadInLostMode200Response) GetLostModeOk() (*IphoneOrIpadInLostMode200ResponseLostMode, bool)`

GetLostModeOk returns a tuple with the LostMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostMode

`func (o *IphoneOrIpadInLostMode200Response) SetLostMode(v IphoneOrIpadInLostMode200ResponseLostMode)`

SetLostMode sets LostMode field to given value.

### HasLostMode

`func (o *IphoneOrIpadInLostMode200Response) HasLostMode() bool`

HasLostMode returns a boolean if a field has been set.

### GetAutomatedDeviceEnrollment

`func (o *IphoneOrIpadInLostMode200Response) GetAutomatedDeviceEnrollment() IphoneOrIpadInLostMode200ResponseAutomatedDeviceEnrollment`

GetAutomatedDeviceEnrollment returns the AutomatedDeviceEnrollment field if non-nil, zero value otherwise.

### GetAutomatedDeviceEnrollmentOk

`func (o *IphoneOrIpadInLostMode200Response) GetAutomatedDeviceEnrollmentOk() (*IphoneOrIpadInLostMode200ResponseAutomatedDeviceEnrollment, bool)`

GetAutomatedDeviceEnrollmentOk returns a tuple with the AutomatedDeviceEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedDeviceEnrollment

`func (o *IphoneOrIpadInLostMode200Response) SetAutomatedDeviceEnrollment(v IphoneOrIpadInLostMode200ResponseAutomatedDeviceEnrollment)`

SetAutomatedDeviceEnrollment sets AutomatedDeviceEnrollment field to given value.

### HasAutomatedDeviceEnrollment

`func (o *IphoneOrIpadInLostMode200Response) HasAutomatedDeviceEnrollment() bool`

HasAutomatedDeviceEnrollment returns a boolean if a field has been set.

### GetKandjiAgent

`func (o *IphoneOrIpadInLostMode200Response) GetKandjiAgent() IphoneOrIpadInLostMode200ResponseKandjiAgent`

GetKandjiAgent returns the KandjiAgent field if non-nil, zero value otherwise.

### GetKandjiAgentOk

`func (o *IphoneOrIpadInLostMode200Response) GetKandjiAgentOk() (*IphoneOrIpadInLostMode200ResponseKandjiAgent, bool)`

GetKandjiAgentOk returns a tuple with the KandjiAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKandjiAgent

`func (o *IphoneOrIpadInLostMode200Response) SetKandjiAgent(v IphoneOrIpadInLostMode200ResponseKandjiAgent)`

SetKandjiAgent sets KandjiAgent field to given value.

### HasKandjiAgent

`func (o *IphoneOrIpadInLostMode200Response) HasKandjiAgent() bool`

HasKandjiAgent returns a boolean if a field has been set.

### GetHardwareOverview

`func (o *IphoneOrIpadInLostMode200Response) GetHardwareOverview() IphoneOrIpadInLostMode200ResponseHardwareOverview`

GetHardwareOverview returns the HardwareOverview field if non-nil, zero value otherwise.

### GetHardwareOverviewOk

`func (o *IphoneOrIpadInLostMode200Response) GetHardwareOverviewOk() (*IphoneOrIpadInLostMode200ResponseHardwareOverview, bool)`

GetHardwareOverviewOk returns a tuple with the HardwareOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareOverview

`func (o *IphoneOrIpadInLostMode200Response) SetHardwareOverview(v IphoneOrIpadInLostMode200ResponseHardwareOverview)`

SetHardwareOverview sets HardwareOverview field to given value.

### HasHardwareOverview

`func (o *IphoneOrIpadInLostMode200Response) HasHardwareOverview() bool`

HasHardwareOverview returns a boolean if a field has been set.

### GetVolumes

`func (o *IphoneOrIpadInLostMode200Response) GetVolumes() interface{}`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *IphoneOrIpadInLostMode200Response) GetVolumesOk() (*interface{}, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *IphoneOrIpadInLostMode200Response) SetVolumes(v interface{})`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *IphoneOrIpadInLostMode200Response) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *IphoneOrIpadInLostMode200Response) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *IphoneOrIpadInLostMode200Response) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil
### GetNetwork

`func (o *IphoneOrIpadInLostMode200Response) GetNetwork() map[string]interface{}`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *IphoneOrIpadInLostMode200Response) GetNetworkOk() (*map[string]interface{}, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *IphoneOrIpadInLostMode200Response) SetNetwork(v map[string]interface{})`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *IphoneOrIpadInLostMode200Response) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRecoveryInformation

`func (o *IphoneOrIpadInLostMode200Response) GetRecoveryInformation() IphoneOrIpadInLostMode200ResponseRecoveryInformation`

GetRecoveryInformation returns the RecoveryInformation field if non-nil, zero value otherwise.

### GetRecoveryInformationOk

`func (o *IphoneOrIpadInLostMode200Response) GetRecoveryInformationOk() (*IphoneOrIpadInLostMode200ResponseRecoveryInformation, bool)`

GetRecoveryInformationOk returns a tuple with the RecoveryInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryInformation

`func (o *IphoneOrIpadInLostMode200Response) SetRecoveryInformation(v IphoneOrIpadInLostMode200ResponseRecoveryInformation)`

SetRecoveryInformation sets RecoveryInformation field to given value.

### HasRecoveryInformation

`func (o *IphoneOrIpadInLostMode200Response) HasRecoveryInformation() bool`

HasRecoveryInformation returns a boolean if a field has been set.

### GetUsers

`func (o *IphoneOrIpadInLostMode200Response) GetUsers() IphoneOrIpadInLostMode200ResponseUsers`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IphoneOrIpadInLostMode200Response) GetUsersOk() (*IphoneOrIpadInLostMode200ResponseUsers, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IphoneOrIpadInLostMode200Response) SetUsers(v IphoneOrIpadInLostMode200ResponseUsers)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IphoneOrIpadInLostMode200Response) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetInstalledProfiles

`func (o *IphoneOrIpadInLostMode200Response) GetInstalledProfiles() interface{}`

GetInstalledProfiles returns the InstalledProfiles field if non-nil, zero value otherwise.

### GetInstalledProfilesOk

`func (o *IphoneOrIpadInLostMode200Response) GetInstalledProfilesOk() (*interface{}, bool)`

GetInstalledProfilesOk returns a tuple with the InstalledProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledProfiles

`func (o *IphoneOrIpadInLostMode200Response) SetInstalledProfiles(v interface{})`

SetInstalledProfiles sets InstalledProfiles field to given value.

### HasInstalledProfiles

`func (o *IphoneOrIpadInLostMode200Response) HasInstalledProfiles() bool`

HasInstalledProfiles returns a boolean if a field has been set.

### SetInstalledProfilesNil

`func (o *IphoneOrIpadInLostMode200Response) SetInstalledProfilesNil(b bool)`

 SetInstalledProfilesNil sets the value for InstalledProfiles to be an explicit nil

### UnsetInstalledProfiles
`func (o *IphoneOrIpadInLostMode200Response) UnsetInstalledProfiles()`

UnsetInstalledProfiles ensures that no value is present for InstalledProfiles, not even an explicit nil
### GetAppleBusinessManager

`func (o *IphoneOrIpadInLostMode200Response) GetAppleBusinessManager() IphoneOrIpadInLostMode200ResponseAppleBusinessManager`

GetAppleBusinessManager returns the AppleBusinessManager field if non-nil, zero value otherwise.

### GetAppleBusinessManagerOk

`func (o *IphoneOrIpadInLostMode200Response) GetAppleBusinessManagerOk() (*IphoneOrIpadInLostMode200ResponseAppleBusinessManager, bool)`

GetAppleBusinessManagerOk returns a tuple with the AppleBusinessManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleBusinessManager

`func (o *IphoneOrIpadInLostMode200Response) SetAppleBusinessManager(v IphoneOrIpadInLostMode200ResponseAppleBusinessManager)`

SetAppleBusinessManager sets AppleBusinessManager field to given value.

### HasAppleBusinessManager

`func (o *IphoneOrIpadInLostMode200Response) HasAppleBusinessManager() bool`

HasAppleBusinessManager returns a boolean if a field has been set.

### GetSecurityInformation

`func (o *IphoneOrIpadInLostMode200Response) GetSecurityInformation() IphoneOrIpadInLostMode200ResponseSecurityInformation`

GetSecurityInformation returns the SecurityInformation field if non-nil, zero value otherwise.

### GetSecurityInformationOk

`func (o *IphoneOrIpadInLostMode200Response) GetSecurityInformationOk() (*IphoneOrIpadInLostMode200ResponseSecurityInformation, bool)`

GetSecurityInformationOk returns a tuple with the SecurityInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityInformation

`func (o *IphoneOrIpadInLostMode200Response) SetSecurityInformation(v IphoneOrIpadInLostMode200ResponseSecurityInformation)`

SetSecurityInformation sets SecurityInformation field to given value.

### HasSecurityInformation

`func (o *IphoneOrIpadInLostMode200Response) HasSecurityInformation() bool`

HasSecurityInformation returns a boolean if a field has been set.

### GetCellular

`func (o *IphoneOrIpadInLostMode200Response) GetCellular() IphoneOrIpadInLostMode200ResponseCellular`

GetCellular returns the Cellular field if non-nil, zero value otherwise.

### GetCellularOk

`func (o *IphoneOrIpadInLostMode200Response) GetCellularOk() (*IphoneOrIpadInLostMode200ResponseCellular, bool)`

GetCellularOk returns a tuple with the Cellular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellular

`func (o *IphoneOrIpadInLostMode200Response) SetCellular(v IphoneOrIpadInLostMode200ResponseCellular)`

SetCellular sets Cellular field to given value.

### HasCellular

`func (o *IphoneOrIpadInLostMode200Response) HasCellular() bool`

HasCellular returns a boolean if a field has been set.

### GetTags

`func (o *IphoneOrIpadInLostMode200Response) GetTags() interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IphoneOrIpadInLostMode200Response) GetTagsOk() (*interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IphoneOrIpadInLostMode200Response) SetTags(v interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *IphoneOrIpadInLostMode200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *IphoneOrIpadInLostMode200Response) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *IphoneOrIpadInLostMode200Response) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


