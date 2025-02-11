# Success200Response1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlueprintId** | Pointer to **string** |  | [optional] 
**MdmDevice** | Pointer to [**Success200Response1MdmDevice**](Success200Response1MdmDevice.md) |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**DepAccount** | Pointer to [**Success200Response1DepAccount**](Success200Response1DepAccount.md) |  | [optional] 
**AssetTag** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DeviceAssignedBy** | Pointer to **string** |  | [optional] 
**DeviceAssignedDate** | Pointer to **string** |  | [optional] 
**DeviceFamily** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**ProfileAssignTime** | Pointer to **string** |  | [optional] 
**ProfilePushTime** | Pointer to **string** |  | [optional] 
**ProfileStatus** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastAssignmentStatus** | Pointer to **string** |  | [optional] 
**FailedAssignmentAttempts** | Pointer to **int32** |  | [optional] 
**AssignmentStatusReceivedAt** | Pointer to **string** |  | [optional] 
**Blueprint** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewSuccess200Response1

`func NewSuccess200Response1() *Success200Response1`

NewSuccess200Response1 instantiates a new Success200Response1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccess200Response1WithDefaults

`func NewSuccess200Response1WithDefaults() *Success200Response1`

NewSuccess200Response1WithDefaults instantiates a new Success200Response1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlueprintId

`func (o *Success200Response1) GetBlueprintId() string`

GetBlueprintId returns the BlueprintId field if non-nil, zero value otherwise.

### GetBlueprintIdOk

`func (o *Success200Response1) GetBlueprintIdOk() (*string, bool)`

GetBlueprintIdOk returns a tuple with the BlueprintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintId

`func (o *Success200Response1) SetBlueprintId(v string)`

SetBlueprintId sets BlueprintId field to given value.

### HasBlueprintId

`func (o *Success200Response1) HasBlueprintId() bool`

HasBlueprintId returns a boolean if a field has been set.

### GetMdmDevice

`func (o *Success200Response1) GetMdmDevice() Success200Response1MdmDevice`

GetMdmDevice returns the MdmDevice field if non-nil, zero value otherwise.

### GetMdmDeviceOk

`func (o *Success200Response1) GetMdmDeviceOk() (*Success200Response1MdmDevice, bool)`

GetMdmDeviceOk returns a tuple with the MdmDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmDevice

`func (o *Success200Response1) SetMdmDevice(v Success200Response1MdmDevice)`

SetMdmDevice sets MdmDevice field to given value.

### HasMdmDevice

`func (o *Success200Response1) HasMdmDevice() bool`

HasMdmDevice returns a boolean if a field has been set.

### GetUserId

`func (o *Success200Response1) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Success200Response1) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Success200Response1) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Success200Response1) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetDepAccount

`func (o *Success200Response1) GetDepAccount() Success200Response1DepAccount`

GetDepAccount returns the DepAccount field if non-nil, zero value otherwise.

### GetDepAccountOk

`func (o *Success200Response1) GetDepAccountOk() (*Success200Response1DepAccount, bool)`

GetDepAccountOk returns a tuple with the DepAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepAccount

`func (o *Success200Response1) SetDepAccount(v Success200Response1DepAccount)`

SetDepAccount sets DepAccount field to given value.

### HasDepAccount

`func (o *Success200Response1) HasDepAccount() bool`

HasDepAccount returns a boolean if a field has been set.

### GetAssetTag

`func (o *Success200Response1) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *Success200Response1) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *Success200Response1) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *Success200Response1) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetColor

`func (o *Success200Response1) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Success200Response1) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Success200Response1) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Success200Response1) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *Success200Response1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Success200Response1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Success200Response1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Success200Response1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceAssignedBy

`func (o *Success200Response1) GetDeviceAssignedBy() string`

GetDeviceAssignedBy returns the DeviceAssignedBy field if non-nil, zero value otherwise.

### GetDeviceAssignedByOk

`func (o *Success200Response1) GetDeviceAssignedByOk() (*string, bool)`

GetDeviceAssignedByOk returns a tuple with the DeviceAssignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAssignedBy

`func (o *Success200Response1) SetDeviceAssignedBy(v string)`

SetDeviceAssignedBy sets DeviceAssignedBy field to given value.

### HasDeviceAssignedBy

`func (o *Success200Response1) HasDeviceAssignedBy() bool`

HasDeviceAssignedBy returns a boolean if a field has been set.

### GetDeviceAssignedDate

`func (o *Success200Response1) GetDeviceAssignedDate() string`

GetDeviceAssignedDate returns the DeviceAssignedDate field if non-nil, zero value otherwise.

### GetDeviceAssignedDateOk

`func (o *Success200Response1) GetDeviceAssignedDateOk() (*string, bool)`

GetDeviceAssignedDateOk returns a tuple with the DeviceAssignedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAssignedDate

`func (o *Success200Response1) SetDeviceAssignedDate(v string)`

SetDeviceAssignedDate sets DeviceAssignedDate field to given value.

### HasDeviceAssignedDate

`func (o *Success200Response1) HasDeviceAssignedDate() bool`

HasDeviceAssignedDate returns a boolean if a field has been set.

### GetDeviceFamily

`func (o *Success200Response1) GetDeviceFamily() string`

GetDeviceFamily returns the DeviceFamily field if non-nil, zero value otherwise.

### GetDeviceFamilyOk

`func (o *Success200Response1) GetDeviceFamilyOk() (*string, bool)`

GetDeviceFamilyOk returns a tuple with the DeviceFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFamily

`func (o *Success200Response1) SetDeviceFamily(v string)`

SetDeviceFamily sets DeviceFamily field to given value.

### HasDeviceFamily

`func (o *Success200Response1) HasDeviceFamily() bool`

HasDeviceFamily returns a boolean if a field has been set.

### GetModel

`func (o *Success200Response1) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Success200Response1) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Success200Response1) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Success200Response1) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOs

`func (o *Success200Response1) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Success200Response1) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Success200Response1) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *Success200Response1) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetProfileAssignTime

`func (o *Success200Response1) GetProfileAssignTime() string`

GetProfileAssignTime returns the ProfileAssignTime field if non-nil, zero value otherwise.

### GetProfileAssignTimeOk

`func (o *Success200Response1) GetProfileAssignTimeOk() (*string, bool)`

GetProfileAssignTimeOk returns a tuple with the ProfileAssignTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileAssignTime

`func (o *Success200Response1) SetProfileAssignTime(v string)`

SetProfileAssignTime sets ProfileAssignTime field to given value.

### HasProfileAssignTime

`func (o *Success200Response1) HasProfileAssignTime() bool`

HasProfileAssignTime returns a boolean if a field has been set.

### GetProfilePushTime

`func (o *Success200Response1) GetProfilePushTime() string`

GetProfilePushTime returns the ProfilePushTime field if non-nil, zero value otherwise.

### GetProfilePushTimeOk

`func (o *Success200Response1) GetProfilePushTimeOk() (*string, bool)`

GetProfilePushTimeOk returns a tuple with the ProfilePushTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePushTime

`func (o *Success200Response1) SetProfilePushTime(v string)`

SetProfilePushTime sets ProfilePushTime field to given value.

### HasProfilePushTime

`func (o *Success200Response1) HasProfilePushTime() bool`

HasProfilePushTime returns a boolean if a field has been set.

### GetProfileStatus

`func (o *Success200Response1) GetProfileStatus() string`

GetProfileStatus returns the ProfileStatus field if non-nil, zero value otherwise.

### GetProfileStatusOk

`func (o *Success200Response1) GetProfileStatusOk() (*string, bool)`

GetProfileStatusOk returns a tuple with the ProfileStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileStatus

`func (o *Success200Response1) SetProfileStatus(v string)`

SetProfileStatus sets ProfileStatus field to given value.

### HasProfileStatus

`func (o *Success200Response1) HasProfileStatus() bool`

HasProfileStatus returns a boolean if a field has been set.

### GetSerialNumber

`func (o *Success200Response1) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Success200Response1) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Success200Response1) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Success200Response1) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetId

`func (o *Success200Response1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Success200Response1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Success200Response1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Success200Response1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastAssignmentStatus

`func (o *Success200Response1) GetLastAssignmentStatus() string`

GetLastAssignmentStatus returns the LastAssignmentStatus field if non-nil, zero value otherwise.

### GetLastAssignmentStatusOk

`func (o *Success200Response1) GetLastAssignmentStatusOk() (*string, bool)`

GetLastAssignmentStatusOk returns a tuple with the LastAssignmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAssignmentStatus

`func (o *Success200Response1) SetLastAssignmentStatus(v string)`

SetLastAssignmentStatus sets LastAssignmentStatus field to given value.

### HasLastAssignmentStatus

`func (o *Success200Response1) HasLastAssignmentStatus() bool`

HasLastAssignmentStatus returns a boolean if a field has been set.

### GetFailedAssignmentAttempts

`func (o *Success200Response1) GetFailedAssignmentAttempts() int32`

GetFailedAssignmentAttempts returns the FailedAssignmentAttempts field if non-nil, zero value otherwise.

### GetFailedAssignmentAttemptsOk

`func (o *Success200Response1) GetFailedAssignmentAttemptsOk() (*int32, bool)`

GetFailedAssignmentAttemptsOk returns a tuple with the FailedAssignmentAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAssignmentAttempts

`func (o *Success200Response1) SetFailedAssignmentAttempts(v int32)`

SetFailedAssignmentAttempts sets FailedAssignmentAttempts field to given value.

### HasFailedAssignmentAttempts

`func (o *Success200Response1) HasFailedAssignmentAttempts() bool`

HasFailedAssignmentAttempts returns a boolean if a field has been set.

### GetAssignmentStatusReceivedAt

`func (o *Success200Response1) GetAssignmentStatusReceivedAt() string`

GetAssignmentStatusReceivedAt returns the AssignmentStatusReceivedAt field if non-nil, zero value otherwise.

### GetAssignmentStatusReceivedAtOk

`func (o *Success200Response1) GetAssignmentStatusReceivedAtOk() (*string, bool)`

GetAssignmentStatusReceivedAtOk returns a tuple with the AssignmentStatusReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentStatusReceivedAt

`func (o *Success200Response1) SetAssignmentStatusReceivedAt(v string)`

SetAssignmentStatusReceivedAt sets AssignmentStatusReceivedAt field to given value.

### HasAssignmentStatusReceivedAt

`func (o *Success200Response1) HasAssignmentStatusReceivedAt() bool`

HasAssignmentStatusReceivedAt returns a boolean if a field has been set.

### GetBlueprint

`func (o *Success200Response1) GetBlueprint() string`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *Success200Response1) GetBlueprintOk() (*string, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *Success200Response1) SetBlueprint(v string)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *Success200Response1) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetUser

`func (o *Success200Response1) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Success200Response1) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Success200Response1) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Success200Response1) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


