# AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Blueprint** | Pointer to [**AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseBlueprint**](AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseBlueprint.md) |  | [optional] 
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
**Defaults** | Pointer to [**AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseDefaults**](AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseDefaults.md) |  | [optional] 
**DaysLeft** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusReason** | Pointer to **interface{}** |  | [optional] 
**StatusReceivedAt** | Pointer to **interface{}** |  | [optional] 
**DeviceCounts** | Pointer to [**AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseDeviceCounts**](AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseDeviceCounts.md) |  | [optional] 

## Methods

### NewAutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response

`func NewAutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response() *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response`

NewAutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response instantiates a new AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseWithDefaults

`func NewAutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseWithDefaults() *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response`

NewAutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseWithDefaults instantiates a new AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBlueprint

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetBlueprint() AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseBlueprint`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetBlueprintOk() (*AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseBlueprint, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetBlueprint(v AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseBlueprint)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetAccessTokenExpiry

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetAccessTokenExpiry() string`

GetAccessTokenExpiry returns the AccessTokenExpiry field if non-nil, zero value otherwise.

### GetAccessTokenExpiryOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetAccessTokenExpiryOk() (*string, bool)`

GetAccessTokenExpiryOk returns a tuple with the AccessTokenExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpiry

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetAccessTokenExpiry(v string)`

SetAccessTokenExpiry sets AccessTokenExpiry field to given value.

### HasAccessTokenExpiry

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasAccessTokenExpiry() bool`

HasAccessTokenExpiry returns a boolean if a field has been set.

### GetServerName

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerUuid

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetServerUuid() string`

GetServerUuid returns the ServerUuid field if non-nil, zero value otherwise.

### GetServerUuidOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetServerUuidOk() (*string, bool)`

GetServerUuidOk returns a tuple with the ServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUuid

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetServerUuid(v string)`

SetServerUuid sets ServerUuid field to given value.

### HasServerUuid

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasServerUuid() bool`

HasServerUuid returns a boolean if a field has been set.

### GetAdminId

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetAdminId() string`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetAdminIdOk() (*string, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetAdminId(v string)`

SetAdminId sets AdminId field to given value.

### HasAdminId

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasAdminId() bool`

HasAdminId returns a boolean if a field has been set.

### GetOrgName

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetOrgEmail

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetOrgEmail() string`

GetOrgEmail returns the OrgEmail field if non-nil, zero value otherwise.

### GetOrgEmailOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetOrgEmailOk() (*string, bool)`

GetOrgEmailOk returns a tuple with the OrgEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgEmail

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetOrgEmail(v string)`

SetOrgEmail sets OrgEmail field to given value.

### HasOrgEmail

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasOrgEmail() bool`

HasOrgEmail returns a boolean if a field has been set.

### GetOrgPhone

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetOrgPhone() string`

GetOrgPhone returns the OrgPhone field if non-nil, zero value otherwise.

### GetOrgPhoneOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetOrgPhoneOk() (*string, bool)`

GetOrgPhoneOk returns a tuple with the OrgPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgPhone

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetOrgPhone(v string)`

SetOrgPhone sets OrgPhone field to given value.

### HasOrgPhone

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasOrgPhone() bool`

HasOrgPhone returns a boolean if a field has been set.

### GetOrgAddress

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetOrgAddress() string`

GetOrgAddress returns the OrgAddress field if non-nil, zero value otherwise.

### GetOrgAddressOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetOrgAddressOk() (*string, bool)`

GetOrgAddressOk returns a tuple with the OrgAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAddress

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetOrgAddress(v string)`

SetOrgAddress sets OrgAddress field to given value.

### HasOrgAddress

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasOrgAddress() bool`

HasOrgAddress returns a boolean if a field has been set.

### GetOrgType

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetOrgType() string`

GetOrgType returns the OrgType field if non-nil, zero value otherwise.

### GetOrgTypeOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetOrgTypeOk() (*string, bool)`

GetOrgTypeOk returns a tuple with the OrgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgType

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetOrgType(v string)`

SetOrgType sets OrgType field to given value.

### HasOrgType

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasOrgType() bool`

HasOrgType returns a boolean if a field has been set.

### GetStokenFileName

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetStokenFileName() string`

GetStokenFileName returns the StokenFileName field if non-nil, zero value otherwise.

### GetStokenFileNameOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetStokenFileNameOk() (*string, bool)`

GetStokenFileNameOk returns a tuple with the StokenFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStokenFileName

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetStokenFileName(v string)`

SetStokenFileName sets StokenFileName field to given value.

### HasStokenFileName

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasStokenFileName() bool`

HasStokenFileName returns a boolean if a field has been set.

### GetLastDeviceSync

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetLastDeviceSync() interface{}`

GetLastDeviceSync returns the LastDeviceSync field if non-nil, zero value otherwise.

### GetLastDeviceSyncOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetLastDeviceSyncOk() (*interface{}, bool)`

GetLastDeviceSyncOk returns a tuple with the LastDeviceSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeviceSync

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetLastDeviceSync(v interface{})`

SetLastDeviceSync sets LastDeviceSync field to given value.

### HasLastDeviceSync

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasLastDeviceSync() bool`

HasLastDeviceSync returns a boolean if a field has been set.

### SetLastDeviceSyncNil

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetLastDeviceSyncNil(b bool)`

 SetLastDeviceSyncNil sets the value for LastDeviceSync to be an explicit nil

### UnsetLastDeviceSync
`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) UnsetLastDeviceSync()`

UnsetLastDeviceSync ensures that no value is present for LastDeviceSync, not even an explicit nil
### GetDefaults

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetDefaults() AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseDefaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetDefaultsOk() (*AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseDefaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetDefaults(v AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseDefaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetDaysLeft

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetDaysLeft() int32`

GetDaysLeft returns the DaysLeft field if non-nil, zero value otherwise.

### GetDaysLeftOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetDaysLeftOk() (*int32, bool)`

GetDaysLeftOk returns a tuple with the DaysLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysLeft

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetDaysLeft(v int32)`

SetDaysLeft sets DaysLeft field to given value.

### HasDaysLeft

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasDaysLeft() bool`

HasDaysLeft returns a boolean if a field has been set.

### GetStatus

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetStatusReason() interface{}`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetStatusReasonOk() (*interface{}, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetStatusReason(v interface{})`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### SetStatusReasonNil

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetStatusReasonNil(b bool)`

 SetStatusReasonNil sets the value for StatusReason to be an explicit nil

### UnsetStatusReason
`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) UnsetStatusReason()`

UnsetStatusReason ensures that no value is present for StatusReason, not even an explicit nil
### GetStatusReceivedAt

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetStatusReceivedAt() interface{}`

GetStatusReceivedAt returns the StatusReceivedAt field if non-nil, zero value otherwise.

### GetStatusReceivedAtOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetStatusReceivedAtOk() (*interface{}, bool)`

GetStatusReceivedAtOk returns a tuple with the StatusReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReceivedAt

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetStatusReceivedAt(v interface{})`

SetStatusReceivedAt sets StatusReceivedAt field to given value.

### HasStatusReceivedAt

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasStatusReceivedAt() bool`

HasStatusReceivedAt returns a boolean if a field has been set.

### SetStatusReceivedAtNil

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetStatusReceivedAtNil(b bool)`

 SetStatusReceivedAtNil sets the value for StatusReceivedAt to be an explicit nil

### UnsetStatusReceivedAt
`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) UnsetStatusReceivedAt()`

UnsetStatusReceivedAt ensures that no value is present for StatusReceivedAt, not even an explicit nil
### GetDeviceCounts

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetDeviceCounts() AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseDeviceCounts`

GetDeviceCounts returns the DeviceCounts field if non-nil, zero value otherwise.

### GetDeviceCountsOk

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) GetDeviceCountsOk() (*AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseDeviceCounts, bool)`

GetDeviceCountsOk returns a tuple with the DeviceCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCounts

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) SetDeviceCounts(v AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200ResponseDeviceCounts)`

SetDeviceCounts sets DeviceCounts field to given value.

### HasDeviceCounts

`func (o *AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration200Response) HasDeviceCounts() bool`

HasDeviceCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


