# UpdateUserAssignment200Response

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

### NewUpdateUserAssignment200Response

`func NewUpdateUserAssignment200Response() *UpdateUserAssignment200Response`

NewUpdateUserAssignment200Response instantiates a new UpdateUserAssignment200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserAssignment200ResponseWithDefaults

`func NewUpdateUserAssignment200ResponseWithDefaults() *UpdateUserAssignment200Response`

NewUpdateUserAssignment200ResponseWithDefaults instantiates a new UpdateUserAssignment200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlueprintId

`func (o *UpdateUserAssignment200Response) GetBlueprintId() string`

GetBlueprintId returns the BlueprintId field if non-nil, zero value otherwise.

### GetBlueprintIdOk

`func (o *UpdateUserAssignment200Response) GetBlueprintIdOk() (*string, bool)`

GetBlueprintIdOk returns a tuple with the BlueprintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintId

`func (o *UpdateUserAssignment200Response) SetBlueprintId(v string)`

SetBlueprintId sets BlueprintId field to given value.

### HasBlueprintId

`func (o *UpdateUserAssignment200Response) HasBlueprintId() bool`

HasBlueprintId returns a boolean if a field has been set.

### GetMdmDevice

`func (o *UpdateUserAssignment200Response) GetMdmDevice() Success200Response1MdmDevice`

GetMdmDevice returns the MdmDevice field if non-nil, zero value otherwise.

### GetMdmDeviceOk

`func (o *UpdateUserAssignment200Response) GetMdmDeviceOk() (*Success200Response1MdmDevice, bool)`

GetMdmDeviceOk returns a tuple with the MdmDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmDevice

`func (o *UpdateUserAssignment200Response) SetMdmDevice(v Success200Response1MdmDevice)`

SetMdmDevice sets MdmDevice field to given value.

### HasMdmDevice

`func (o *UpdateUserAssignment200Response) HasMdmDevice() bool`

HasMdmDevice returns a boolean if a field has been set.

### GetUserId

`func (o *UpdateUserAssignment200Response) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdateUserAssignment200Response) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdateUserAssignment200Response) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UpdateUserAssignment200Response) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetDepAccount

`func (o *UpdateUserAssignment200Response) GetDepAccount() Success200Response1DepAccount`

GetDepAccount returns the DepAccount field if non-nil, zero value otherwise.

### GetDepAccountOk

`func (o *UpdateUserAssignment200Response) GetDepAccountOk() (*Success200Response1DepAccount, bool)`

GetDepAccountOk returns a tuple with the DepAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepAccount

`func (o *UpdateUserAssignment200Response) SetDepAccount(v Success200Response1DepAccount)`

SetDepAccount sets DepAccount field to given value.

### HasDepAccount

`func (o *UpdateUserAssignment200Response) HasDepAccount() bool`

HasDepAccount returns a boolean if a field has been set.

### GetAssetTag

`func (o *UpdateUserAssignment200Response) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *UpdateUserAssignment200Response) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *UpdateUserAssignment200Response) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *UpdateUserAssignment200Response) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### GetColor

`func (o *UpdateUserAssignment200Response) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *UpdateUserAssignment200Response) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *UpdateUserAssignment200Response) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *UpdateUserAssignment200Response) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateUserAssignment200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateUserAssignment200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateUserAssignment200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateUserAssignment200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceAssignedBy

`func (o *UpdateUserAssignment200Response) GetDeviceAssignedBy() string`

GetDeviceAssignedBy returns the DeviceAssignedBy field if non-nil, zero value otherwise.

### GetDeviceAssignedByOk

`func (o *UpdateUserAssignment200Response) GetDeviceAssignedByOk() (*string, bool)`

GetDeviceAssignedByOk returns a tuple with the DeviceAssignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAssignedBy

`func (o *UpdateUserAssignment200Response) SetDeviceAssignedBy(v string)`

SetDeviceAssignedBy sets DeviceAssignedBy field to given value.

### HasDeviceAssignedBy

`func (o *UpdateUserAssignment200Response) HasDeviceAssignedBy() bool`

HasDeviceAssignedBy returns a boolean if a field has been set.

### GetDeviceAssignedDate

`func (o *UpdateUserAssignment200Response) GetDeviceAssignedDate() string`

GetDeviceAssignedDate returns the DeviceAssignedDate field if non-nil, zero value otherwise.

### GetDeviceAssignedDateOk

`func (o *UpdateUserAssignment200Response) GetDeviceAssignedDateOk() (*string, bool)`

GetDeviceAssignedDateOk returns a tuple with the DeviceAssignedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAssignedDate

`func (o *UpdateUserAssignment200Response) SetDeviceAssignedDate(v string)`

SetDeviceAssignedDate sets DeviceAssignedDate field to given value.

### HasDeviceAssignedDate

`func (o *UpdateUserAssignment200Response) HasDeviceAssignedDate() bool`

HasDeviceAssignedDate returns a boolean if a field has been set.

### GetDeviceFamily

`func (o *UpdateUserAssignment200Response) GetDeviceFamily() string`

GetDeviceFamily returns the DeviceFamily field if non-nil, zero value otherwise.

### GetDeviceFamilyOk

`func (o *UpdateUserAssignment200Response) GetDeviceFamilyOk() (*string, bool)`

GetDeviceFamilyOk returns a tuple with the DeviceFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFamily

`func (o *UpdateUserAssignment200Response) SetDeviceFamily(v string)`

SetDeviceFamily sets DeviceFamily field to given value.

### HasDeviceFamily

`func (o *UpdateUserAssignment200Response) HasDeviceFamily() bool`

HasDeviceFamily returns a boolean if a field has been set.

### GetModel

`func (o *UpdateUserAssignment200Response) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *UpdateUserAssignment200Response) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *UpdateUserAssignment200Response) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *UpdateUserAssignment200Response) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOs

`func (o *UpdateUserAssignment200Response) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *UpdateUserAssignment200Response) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *UpdateUserAssignment200Response) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *UpdateUserAssignment200Response) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetProfileAssignTime

`func (o *UpdateUserAssignment200Response) GetProfileAssignTime() string`

GetProfileAssignTime returns the ProfileAssignTime field if non-nil, zero value otherwise.

### GetProfileAssignTimeOk

`func (o *UpdateUserAssignment200Response) GetProfileAssignTimeOk() (*string, bool)`

GetProfileAssignTimeOk returns a tuple with the ProfileAssignTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileAssignTime

`func (o *UpdateUserAssignment200Response) SetProfileAssignTime(v string)`

SetProfileAssignTime sets ProfileAssignTime field to given value.

### HasProfileAssignTime

`func (o *UpdateUserAssignment200Response) HasProfileAssignTime() bool`

HasProfileAssignTime returns a boolean if a field has been set.

### GetProfilePushTime

`func (o *UpdateUserAssignment200Response) GetProfilePushTime() string`

GetProfilePushTime returns the ProfilePushTime field if non-nil, zero value otherwise.

### GetProfilePushTimeOk

`func (o *UpdateUserAssignment200Response) GetProfilePushTimeOk() (*string, bool)`

GetProfilePushTimeOk returns a tuple with the ProfilePushTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePushTime

`func (o *UpdateUserAssignment200Response) SetProfilePushTime(v string)`

SetProfilePushTime sets ProfilePushTime field to given value.

### HasProfilePushTime

`func (o *UpdateUserAssignment200Response) HasProfilePushTime() bool`

HasProfilePushTime returns a boolean if a field has been set.

### GetProfileStatus

`func (o *UpdateUserAssignment200Response) GetProfileStatus() string`

GetProfileStatus returns the ProfileStatus field if non-nil, zero value otherwise.

### GetProfileStatusOk

`func (o *UpdateUserAssignment200Response) GetProfileStatusOk() (*string, bool)`

GetProfileStatusOk returns a tuple with the ProfileStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileStatus

`func (o *UpdateUserAssignment200Response) SetProfileStatus(v string)`

SetProfileStatus sets ProfileStatus field to given value.

### HasProfileStatus

`func (o *UpdateUserAssignment200Response) HasProfileStatus() bool`

HasProfileStatus returns a boolean if a field has been set.

### GetSerialNumber

`func (o *UpdateUserAssignment200Response) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *UpdateUserAssignment200Response) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *UpdateUserAssignment200Response) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *UpdateUserAssignment200Response) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetId

`func (o *UpdateUserAssignment200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateUserAssignment200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateUserAssignment200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateUserAssignment200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastAssignmentStatus

`func (o *UpdateUserAssignment200Response) GetLastAssignmentStatus() string`

GetLastAssignmentStatus returns the LastAssignmentStatus field if non-nil, zero value otherwise.

### GetLastAssignmentStatusOk

`func (o *UpdateUserAssignment200Response) GetLastAssignmentStatusOk() (*string, bool)`

GetLastAssignmentStatusOk returns a tuple with the LastAssignmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAssignmentStatus

`func (o *UpdateUserAssignment200Response) SetLastAssignmentStatus(v string)`

SetLastAssignmentStatus sets LastAssignmentStatus field to given value.

### HasLastAssignmentStatus

`func (o *UpdateUserAssignment200Response) HasLastAssignmentStatus() bool`

HasLastAssignmentStatus returns a boolean if a field has been set.

### GetFailedAssignmentAttempts

`func (o *UpdateUserAssignment200Response) GetFailedAssignmentAttempts() int32`

GetFailedAssignmentAttempts returns the FailedAssignmentAttempts field if non-nil, zero value otherwise.

### GetFailedAssignmentAttemptsOk

`func (o *UpdateUserAssignment200Response) GetFailedAssignmentAttemptsOk() (*int32, bool)`

GetFailedAssignmentAttemptsOk returns a tuple with the FailedAssignmentAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAssignmentAttempts

`func (o *UpdateUserAssignment200Response) SetFailedAssignmentAttempts(v int32)`

SetFailedAssignmentAttempts sets FailedAssignmentAttempts field to given value.

### HasFailedAssignmentAttempts

`func (o *UpdateUserAssignment200Response) HasFailedAssignmentAttempts() bool`

HasFailedAssignmentAttempts returns a boolean if a field has been set.

### GetAssignmentStatusReceivedAt

`func (o *UpdateUserAssignment200Response) GetAssignmentStatusReceivedAt() string`

GetAssignmentStatusReceivedAt returns the AssignmentStatusReceivedAt field if non-nil, zero value otherwise.

### GetAssignmentStatusReceivedAtOk

`func (o *UpdateUserAssignment200Response) GetAssignmentStatusReceivedAtOk() (*string, bool)`

GetAssignmentStatusReceivedAtOk returns a tuple with the AssignmentStatusReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentStatusReceivedAt

`func (o *UpdateUserAssignment200Response) SetAssignmentStatusReceivedAt(v string)`

SetAssignmentStatusReceivedAt sets AssignmentStatusReceivedAt field to given value.

### HasAssignmentStatusReceivedAt

`func (o *UpdateUserAssignment200Response) HasAssignmentStatusReceivedAt() bool`

HasAssignmentStatusReceivedAt returns a boolean if a field has been set.

### GetBlueprint

`func (o *UpdateUserAssignment200Response) GetBlueprint() string`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *UpdateUserAssignment200Response) GetBlueprintOk() (*string, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *UpdateUserAssignment200Response) SetBlueprint(v string)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *UpdateUserAssignment200Response) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetUser

`func (o *UpdateUserAssignment200Response) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UpdateUserAssignment200Response) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UpdateUserAssignment200Response) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *UpdateUserAssignment200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


