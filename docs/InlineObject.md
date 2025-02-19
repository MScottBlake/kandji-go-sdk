# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Blueprint** | Pointer to [**InlineObjectBlueprint**](InlineObjectBlueprint.md) |  | [optional] 
**AccessTokenExpiry** | Pointer to **string** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**ServerUuid** | Pointer to **string** |  | [optional] 
**AdminId** | Pointer to **string** |  | [optional] 
**OrgName** | Pointer to **string** |  | [optional] 
**OrgEmail** | Pointer to **string** |  | [optional] 
**OrgPhone** | Pointer to **string** |  | [optional] 
**OrgAddress** | Pointer to **string** |  | [optional] 
**OrgType** | Pointer to **string** |  | [optional] 
**StokenFileName** | Pointer to **string** |  | [optional] 
**LastDeviceSync** | Pointer to **interface{}** |  | [optional] 
**Defaults** | Pointer to [**InlineObjectDefaults**](InlineObjectDefaults.md) |  | [optional] 
**DaysLeft** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusReason** | Pointer to **interface{}** |  | [optional] 
**StatusReceivedAt** | Pointer to **interface{}** |  | [optional] 
**DeviceCounts** | Pointer to [**InlineObjectDeviceCounts**](InlineObjectDeviceCounts.md) |  | [optional] 

## Methods

### NewInlineObject

`func NewInlineObject() *InlineObject`

NewInlineObject instantiates a new InlineObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObjectWithDefaults

`func NewInlineObjectWithDefaults() *InlineObject`

NewInlineObjectWithDefaults instantiates a new InlineObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBlueprint

`func (o *InlineObject) GetBlueprint() InlineObjectBlueprint`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *InlineObject) GetBlueprintOk() (*InlineObjectBlueprint, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *InlineObject) SetBlueprint(v InlineObjectBlueprint)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *InlineObject) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetAccessTokenExpiry

`func (o *InlineObject) GetAccessTokenExpiry() string`

GetAccessTokenExpiry returns the AccessTokenExpiry field if non-nil, zero value otherwise.

### GetAccessTokenExpiryOk

`func (o *InlineObject) GetAccessTokenExpiryOk() (*string, bool)`

GetAccessTokenExpiryOk returns a tuple with the AccessTokenExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpiry

`func (o *InlineObject) SetAccessTokenExpiry(v string)`

SetAccessTokenExpiry sets AccessTokenExpiry field to given value.

### HasAccessTokenExpiry

`func (o *InlineObject) HasAccessTokenExpiry() bool`

HasAccessTokenExpiry returns a boolean if a field has been set.

### GetServerName

`func (o *InlineObject) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *InlineObject) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *InlineObject) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *InlineObject) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerUuid

`func (o *InlineObject) GetServerUuid() string`

GetServerUuid returns the ServerUuid field if non-nil, zero value otherwise.

### GetServerUuidOk

`func (o *InlineObject) GetServerUuidOk() (*string, bool)`

GetServerUuidOk returns a tuple with the ServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUuid

`func (o *InlineObject) SetServerUuid(v string)`

SetServerUuid sets ServerUuid field to given value.

### HasServerUuid

`func (o *InlineObject) HasServerUuid() bool`

HasServerUuid returns a boolean if a field has been set.

### GetAdminId

`func (o *InlineObject) GetAdminId() string`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *InlineObject) GetAdminIdOk() (*string, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *InlineObject) SetAdminId(v string)`

SetAdminId sets AdminId field to given value.

### HasAdminId

`func (o *InlineObject) HasAdminId() bool`

HasAdminId returns a boolean if a field has been set.

### GetOrgName

`func (o *InlineObject) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *InlineObject) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *InlineObject) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *InlineObject) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetOrgEmail

`func (o *InlineObject) GetOrgEmail() string`

GetOrgEmail returns the OrgEmail field if non-nil, zero value otherwise.

### GetOrgEmailOk

`func (o *InlineObject) GetOrgEmailOk() (*string, bool)`

GetOrgEmailOk returns a tuple with the OrgEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgEmail

`func (o *InlineObject) SetOrgEmail(v string)`

SetOrgEmail sets OrgEmail field to given value.

### HasOrgEmail

`func (o *InlineObject) HasOrgEmail() bool`

HasOrgEmail returns a boolean if a field has been set.

### GetOrgPhone

`func (o *InlineObject) GetOrgPhone() string`

GetOrgPhone returns the OrgPhone field if non-nil, zero value otherwise.

### GetOrgPhoneOk

`func (o *InlineObject) GetOrgPhoneOk() (*string, bool)`

GetOrgPhoneOk returns a tuple with the OrgPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgPhone

`func (o *InlineObject) SetOrgPhone(v string)`

SetOrgPhone sets OrgPhone field to given value.

### HasOrgPhone

`func (o *InlineObject) HasOrgPhone() bool`

HasOrgPhone returns a boolean if a field has been set.

### GetOrgAddress

`func (o *InlineObject) GetOrgAddress() string`

GetOrgAddress returns the OrgAddress field if non-nil, zero value otherwise.

### GetOrgAddressOk

`func (o *InlineObject) GetOrgAddressOk() (*string, bool)`

GetOrgAddressOk returns a tuple with the OrgAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAddress

`func (o *InlineObject) SetOrgAddress(v string)`

SetOrgAddress sets OrgAddress field to given value.

### HasOrgAddress

`func (o *InlineObject) HasOrgAddress() bool`

HasOrgAddress returns a boolean if a field has been set.

### GetOrgType

`func (o *InlineObject) GetOrgType() string`

GetOrgType returns the OrgType field if non-nil, zero value otherwise.

### GetOrgTypeOk

`func (o *InlineObject) GetOrgTypeOk() (*string, bool)`

GetOrgTypeOk returns a tuple with the OrgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgType

`func (o *InlineObject) SetOrgType(v string)`

SetOrgType sets OrgType field to given value.

### HasOrgType

`func (o *InlineObject) HasOrgType() bool`

HasOrgType returns a boolean if a field has been set.

### GetStokenFileName

`func (o *InlineObject) GetStokenFileName() string`

GetStokenFileName returns the StokenFileName field if non-nil, zero value otherwise.

### GetStokenFileNameOk

`func (o *InlineObject) GetStokenFileNameOk() (*string, bool)`

GetStokenFileNameOk returns a tuple with the StokenFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStokenFileName

`func (o *InlineObject) SetStokenFileName(v string)`

SetStokenFileName sets StokenFileName field to given value.

### HasStokenFileName

`func (o *InlineObject) HasStokenFileName() bool`

HasStokenFileName returns a boolean if a field has been set.

### GetLastDeviceSync

`func (o *InlineObject) GetLastDeviceSync() interface{}`

GetLastDeviceSync returns the LastDeviceSync field if non-nil, zero value otherwise.

### GetLastDeviceSyncOk

`func (o *InlineObject) GetLastDeviceSyncOk() (*interface{}, bool)`

GetLastDeviceSyncOk returns a tuple with the LastDeviceSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeviceSync

`func (o *InlineObject) SetLastDeviceSync(v interface{})`

SetLastDeviceSync sets LastDeviceSync field to given value.

### HasLastDeviceSync

`func (o *InlineObject) HasLastDeviceSync() bool`

HasLastDeviceSync returns a boolean if a field has been set.

### SetLastDeviceSyncNil

`func (o *InlineObject) SetLastDeviceSyncNil(b bool)`

 SetLastDeviceSyncNil sets the value for LastDeviceSync to be an explicit nil

### UnsetLastDeviceSync
`func (o *InlineObject) UnsetLastDeviceSync()`

UnsetLastDeviceSync ensures that no value is present for LastDeviceSync, not even an explicit nil
### GetDefaults

`func (o *InlineObject) GetDefaults() InlineObjectDefaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *InlineObject) GetDefaultsOk() (*InlineObjectDefaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *InlineObject) SetDefaults(v InlineObjectDefaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *InlineObject) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetDaysLeft

`func (o *InlineObject) GetDaysLeft() int32`

GetDaysLeft returns the DaysLeft field if non-nil, zero value otherwise.

### GetDaysLeftOk

`func (o *InlineObject) GetDaysLeftOk() (*int32, bool)`

GetDaysLeftOk returns a tuple with the DaysLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysLeft

`func (o *InlineObject) SetDaysLeft(v int32)`

SetDaysLeft sets DaysLeft field to given value.

### HasDaysLeft

`func (o *InlineObject) HasDaysLeft() bool`

HasDaysLeft returns a boolean if a field has been set.

### GetStatus

`func (o *InlineObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineObject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *InlineObject) GetStatusReason() interface{}`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *InlineObject) GetStatusReasonOk() (*interface{}, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *InlineObject) SetStatusReason(v interface{})`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *InlineObject) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### SetStatusReasonNil

`func (o *InlineObject) SetStatusReasonNil(b bool)`

 SetStatusReasonNil sets the value for StatusReason to be an explicit nil

### UnsetStatusReason
`func (o *InlineObject) UnsetStatusReason()`

UnsetStatusReason ensures that no value is present for StatusReason, not even an explicit nil
### GetStatusReceivedAt

`func (o *InlineObject) GetStatusReceivedAt() interface{}`

GetStatusReceivedAt returns the StatusReceivedAt field if non-nil, zero value otherwise.

### GetStatusReceivedAtOk

`func (o *InlineObject) GetStatusReceivedAtOk() (*interface{}, bool)`

GetStatusReceivedAtOk returns a tuple with the StatusReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReceivedAt

`func (o *InlineObject) SetStatusReceivedAt(v interface{})`

SetStatusReceivedAt sets StatusReceivedAt field to given value.

### HasStatusReceivedAt

`func (o *InlineObject) HasStatusReceivedAt() bool`

HasStatusReceivedAt returns a boolean if a field has been set.

### SetStatusReceivedAtNil

`func (o *InlineObject) SetStatusReceivedAtNil(b bool)`

 SetStatusReceivedAtNil sets the value for StatusReceivedAt to be an explicit nil

### UnsetStatusReceivedAt
`func (o *InlineObject) UnsetStatusReceivedAt()`

UnsetStatusReceivedAt ensures that no value is present for StatusReceivedAt, not even an explicit nil
### GetDeviceCounts

`func (o *InlineObject) GetDeviceCounts() InlineObjectDeviceCounts`

GetDeviceCounts returns the DeviceCounts field if non-nil, zero value otherwise.

### GetDeviceCountsOk

`func (o *InlineObject) GetDeviceCountsOk() (*InlineObjectDeviceCounts, bool)`

GetDeviceCountsOk returns a tuple with the DeviceCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCounts

`func (o *InlineObject) SetDeviceCounts(v InlineObjectDeviceCounts)`

SetDeviceCounts sets DeviceCounts field to given value.

### HasDeviceCounts

`func (o *InlineObject) HasDeviceCounts() bool`

HasDeviceCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


