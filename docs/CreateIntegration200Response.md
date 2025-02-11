# CreateIntegration200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Blueprint** | Pointer to [**CreateIntegration200ResponseBlueprint**](CreateIntegration200ResponseBlueprint.md) |  | [optional] 
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
**Defaults** | Pointer to [**CreateIntegration200ResponseDefaults**](CreateIntegration200ResponseDefaults.md) |  | [optional] 
**DaysLeft** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusReason** | Pointer to **interface{}** |  | [optional] 
**StatusReceivedAt** | Pointer to **interface{}** |  | [optional] 
**DeviceCounts** | Pointer to [**CreateIntegration200ResponseDeviceCounts**](CreateIntegration200ResponseDeviceCounts.md) |  | [optional] 

## Methods

### NewCreateIntegration200Response

`func NewCreateIntegration200Response() *CreateIntegration200Response`

NewCreateIntegration200Response instantiates a new CreateIntegration200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIntegration200ResponseWithDefaults

`func NewCreateIntegration200ResponseWithDefaults() *CreateIntegration200Response`

NewCreateIntegration200ResponseWithDefaults instantiates a new CreateIntegration200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateIntegration200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateIntegration200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateIntegration200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateIntegration200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBlueprint

`func (o *CreateIntegration200Response) GetBlueprint() CreateIntegration200ResponseBlueprint`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *CreateIntegration200Response) GetBlueprintOk() (*CreateIntegration200ResponseBlueprint, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *CreateIntegration200Response) SetBlueprint(v CreateIntegration200ResponseBlueprint)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *CreateIntegration200Response) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetAccessTokenExpiry

`func (o *CreateIntegration200Response) GetAccessTokenExpiry() string`

GetAccessTokenExpiry returns the AccessTokenExpiry field if non-nil, zero value otherwise.

### GetAccessTokenExpiryOk

`func (o *CreateIntegration200Response) GetAccessTokenExpiryOk() (*string, bool)`

GetAccessTokenExpiryOk returns a tuple with the AccessTokenExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpiry

`func (o *CreateIntegration200Response) SetAccessTokenExpiry(v string)`

SetAccessTokenExpiry sets AccessTokenExpiry field to given value.

### HasAccessTokenExpiry

`func (o *CreateIntegration200Response) HasAccessTokenExpiry() bool`

HasAccessTokenExpiry returns a boolean if a field has been set.

### GetServerName

`func (o *CreateIntegration200Response) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *CreateIntegration200Response) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *CreateIntegration200Response) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *CreateIntegration200Response) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerUuid

`func (o *CreateIntegration200Response) GetServerUuid() string`

GetServerUuid returns the ServerUuid field if non-nil, zero value otherwise.

### GetServerUuidOk

`func (o *CreateIntegration200Response) GetServerUuidOk() (*string, bool)`

GetServerUuidOk returns a tuple with the ServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUuid

`func (o *CreateIntegration200Response) SetServerUuid(v string)`

SetServerUuid sets ServerUuid field to given value.

### HasServerUuid

`func (o *CreateIntegration200Response) HasServerUuid() bool`

HasServerUuid returns a boolean if a field has been set.

### GetAdminId

`func (o *CreateIntegration200Response) GetAdminId() string`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *CreateIntegration200Response) GetAdminIdOk() (*string, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *CreateIntegration200Response) SetAdminId(v string)`

SetAdminId sets AdminId field to given value.

### HasAdminId

`func (o *CreateIntegration200Response) HasAdminId() bool`

HasAdminId returns a boolean if a field has been set.

### GetOrgName

`func (o *CreateIntegration200Response) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *CreateIntegration200Response) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *CreateIntegration200Response) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *CreateIntegration200Response) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetOrgEmail

`func (o *CreateIntegration200Response) GetOrgEmail() string`

GetOrgEmail returns the OrgEmail field if non-nil, zero value otherwise.

### GetOrgEmailOk

`func (o *CreateIntegration200Response) GetOrgEmailOk() (*string, bool)`

GetOrgEmailOk returns a tuple with the OrgEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgEmail

`func (o *CreateIntegration200Response) SetOrgEmail(v string)`

SetOrgEmail sets OrgEmail field to given value.

### HasOrgEmail

`func (o *CreateIntegration200Response) HasOrgEmail() bool`

HasOrgEmail returns a boolean if a field has been set.

### GetOrgPhone

`func (o *CreateIntegration200Response) GetOrgPhone() string`

GetOrgPhone returns the OrgPhone field if non-nil, zero value otherwise.

### GetOrgPhoneOk

`func (o *CreateIntegration200Response) GetOrgPhoneOk() (*string, bool)`

GetOrgPhoneOk returns a tuple with the OrgPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgPhone

`func (o *CreateIntegration200Response) SetOrgPhone(v string)`

SetOrgPhone sets OrgPhone field to given value.

### HasOrgPhone

`func (o *CreateIntegration200Response) HasOrgPhone() bool`

HasOrgPhone returns a boolean if a field has been set.

### GetOrgAddress

`func (o *CreateIntegration200Response) GetOrgAddress() string`

GetOrgAddress returns the OrgAddress field if non-nil, zero value otherwise.

### GetOrgAddressOk

`func (o *CreateIntegration200Response) GetOrgAddressOk() (*string, bool)`

GetOrgAddressOk returns a tuple with the OrgAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAddress

`func (o *CreateIntegration200Response) SetOrgAddress(v string)`

SetOrgAddress sets OrgAddress field to given value.

### HasOrgAddress

`func (o *CreateIntegration200Response) HasOrgAddress() bool`

HasOrgAddress returns a boolean if a field has been set.

### GetOrgType

`func (o *CreateIntegration200Response) GetOrgType() string`

GetOrgType returns the OrgType field if non-nil, zero value otherwise.

### GetOrgTypeOk

`func (o *CreateIntegration200Response) GetOrgTypeOk() (*string, bool)`

GetOrgTypeOk returns a tuple with the OrgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgType

`func (o *CreateIntegration200Response) SetOrgType(v string)`

SetOrgType sets OrgType field to given value.

### HasOrgType

`func (o *CreateIntegration200Response) HasOrgType() bool`

HasOrgType returns a boolean if a field has been set.

### GetStokenFileName

`func (o *CreateIntegration200Response) GetStokenFileName() string`

GetStokenFileName returns the StokenFileName field if non-nil, zero value otherwise.

### GetStokenFileNameOk

`func (o *CreateIntegration200Response) GetStokenFileNameOk() (*string, bool)`

GetStokenFileNameOk returns a tuple with the StokenFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStokenFileName

`func (o *CreateIntegration200Response) SetStokenFileName(v string)`

SetStokenFileName sets StokenFileName field to given value.

### HasStokenFileName

`func (o *CreateIntegration200Response) HasStokenFileName() bool`

HasStokenFileName returns a boolean if a field has been set.

### GetLastDeviceSync

`func (o *CreateIntegration200Response) GetLastDeviceSync() interface{}`

GetLastDeviceSync returns the LastDeviceSync field if non-nil, zero value otherwise.

### GetLastDeviceSyncOk

`func (o *CreateIntegration200Response) GetLastDeviceSyncOk() (*interface{}, bool)`

GetLastDeviceSyncOk returns a tuple with the LastDeviceSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeviceSync

`func (o *CreateIntegration200Response) SetLastDeviceSync(v interface{})`

SetLastDeviceSync sets LastDeviceSync field to given value.

### HasLastDeviceSync

`func (o *CreateIntegration200Response) HasLastDeviceSync() bool`

HasLastDeviceSync returns a boolean if a field has been set.

### SetLastDeviceSyncNil

`func (o *CreateIntegration200Response) SetLastDeviceSyncNil(b bool)`

 SetLastDeviceSyncNil sets the value for LastDeviceSync to be an explicit nil

### UnsetLastDeviceSync
`func (o *CreateIntegration200Response) UnsetLastDeviceSync()`

UnsetLastDeviceSync ensures that no value is present for LastDeviceSync, not even an explicit nil
### GetDefaults

`func (o *CreateIntegration200Response) GetDefaults() CreateIntegration200ResponseDefaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *CreateIntegration200Response) GetDefaultsOk() (*CreateIntegration200ResponseDefaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *CreateIntegration200Response) SetDefaults(v CreateIntegration200ResponseDefaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *CreateIntegration200Response) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetDaysLeft

`func (o *CreateIntegration200Response) GetDaysLeft() int32`

GetDaysLeft returns the DaysLeft field if non-nil, zero value otherwise.

### GetDaysLeftOk

`func (o *CreateIntegration200Response) GetDaysLeftOk() (*int32, bool)`

GetDaysLeftOk returns a tuple with the DaysLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysLeft

`func (o *CreateIntegration200Response) SetDaysLeft(v int32)`

SetDaysLeft sets DaysLeft field to given value.

### HasDaysLeft

`func (o *CreateIntegration200Response) HasDaysLeft() bool`

HasDaysLeft returns a boolean if a field has been set.

### GetStatus

`func (o *CreateIntegration200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateIntegration200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateIntegration200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateIntegration200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *CreateIntegration200Response) GetStatusReason() interface{}`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *CreateIntegration200Response) GetStatusReasonOk() (*interface{}, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *CreateIntegration200Response) SetStatusReason(v interface{})`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *CreateIntegration200Response) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### SetStatusReasonNil

`func (o *CreateIntegration200Response) SetStatusReasonNil(b bool)`

 SetStatusReasonNil sets the value for StatusReason to be an explicit nil

### UnsetStatusReason
`func (o *CreateIntegration200Response) UnsetStatusReason()`

UnsetStatusReason ensures that no value is present for StatusReason, not even an explicit nil
### GetStatusReceivedAt

`func (o *CreateIntegration200Response) GetStatusReceivedAt() interface{}`

GetStatusReceivedAt returns the StatusReceivedAt field if non-nil, zero value otherwise.

### GetStatusReceivedAtOk

`func (o *CreateIntegration200Response) GetStatusReceivedAtOk() (*interface{}, bool)`

GetStatusReceivedAtOk returns a tuple with the StatusReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReceivedAt

`func (o *CreateIntegration200Response) SetStatusReceivedAt(v interface{})`

SetStatusReceivedAt sets StatusReceivedAt field to given value.

### HasStatusReceivedAt

`func (o *CreateIntegration200Response) HasStatusReceivedAt() bool`

HasStatusReceivedAt returns a boolean if a field has been set.

### SetStatusReceivedAtNil

`func (o *CreateIntegration200Response) SetStatusReceivedAtNil(b bool)`

 SetStatusReceivedAtNil sets the value for StatusReceivedAt to be an explicit nil

### UnsetStatusReceivedAt
`func (o *CreateIntegration200Response) UnsetStatusReceivedAt()`

UnsetStatusReceivedAt ensures that no value is present for StatusReceivedAt, not even an explicit nil
### GetDeviceCounts

`func (o *CreateIntegration200Response) GetDeviceCounts() CreateIntegration200ResponseDeviceCounts`

GetDeviceCounts returns the DeviceCounts field if non-nil, zero value otherwise.

### GetDeviceCountsOk

`func (o *CreateIntegration200Response) GetDeviceCountsOk() (*CreateIntegration200ResponseDeviceCounts, bool)`

GetDeviceCountsOk returns a tuple with the DeviceCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCounts

`func (o *CreateIntegration200Response) SetDeviceCounts(v CreateIntegration200ResponseDeviceCounts)`

SetDeviceCounts sets DeviceCounts field to given value.

### HasDeviceCounts

`func (o *CreateIntegration200Response) HasDeviceCounts() bool`

HasDeviceCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


