# Success200Response16

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int32** |  | [optional] 
**Archived** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Department** | Pointer to **interface{}** |  | [optional] 
**DeprecatedUserId** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Integration** | Pointer to [**Success200Response16Integration**](Success200Response16Integration.md) |  | [optional] 
**JobTitle** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DeviceCount** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewSuccess200Response16

`func NewSuccess200Response16() *Success200Response16`

NewSuccess200Response16 instantiates a new Success200Response16 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccess200Response16WithDefaults

`func NewSuccess200Response16WithDefaults() *Success200Response16`

NewSuccess200Response16WithDefaults instantiates a new Success200Response16 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *Success200Response16) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Success200Response16) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Success200Response16) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *Success200Response16) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetArchived

`func (o *Success200Response16) GetArchived() int32`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Success200Response16) GetArchivedOk() (*int32, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Success200Response16) SetArchived(v int32)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Success200Response16) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Success200Response16) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Success200Response16) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Success200Response16) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Success200Response16) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDepartment

`func (o *Success200Response16) GetDepartment() interface{}`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Success200Response16) GetDepartmentOk() (*interface{}, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Success200Response16) SetDepartment(v interface{})`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Success200Response16) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *Success200Response16) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *Success200Response16) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDeprecatedUserId

`func (o *Success200Response16) GetDeprecatedUserId() string`

GetDeprecatedUserId returns the DeprecatedUserId field if non-nil, zero value otherwise.

### GetDeprecatedUserIdOk

`func (o *Success200Response16) GetDeprecatedUserIdOk() (*string, bool)`

GetDeprecatedUserIdOk returns a tuple with the DeprecatedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedUserId

`func (o *Success200Response16) SetDeprecatedUserId(v string)`

SetDeprecatedUserId sets DeprecatedUserId field to given value.

### HasDeprecatedUserId

`func (o *Success200Response16) HasDeprecatedUserId() bool`

HasDeprecatedUserId returns a boolean if a field has been set.

### GetEmail

`func (o *Success200Response16) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Success200Response16) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Success200Response16) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Success200Response16) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *Success200Response16) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Success200Response16) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Success200Response16) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Success200Response16) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntegration

`func (o *Success200Response16) GetIntegration() Success200Response16Integration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *Success200Response16) GetIntegrationOk() (*Success200Response16Integration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *Success200Response16) SetIntegration(v Success200Response16Integration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *Success200Response16) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetJobTitle

`func (o *Success200Response16) GetJobTitle() interface{}`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *Success200Response16) GetJobTitleOk() (*interface{}, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *Success200Response16) SetJobTitle(v interface{})`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *Success200Response16) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *Success200Response16) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *Success200Response16) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetName

`func (o *Success200Response16) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Success200Response16) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Success200Response16) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Success200Response16) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeviceCount

`func (o *Success200Response16) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *Success200Response16) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *Success200Response16) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *Success200Response16) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Success200Response16) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Success200Response16) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Success200Response16) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Success200Response16) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


