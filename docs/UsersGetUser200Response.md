# UsersGetUser200Response

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
**Integration** | Pointer to [**UsersGetUser200ResponseIntegration**](UsersGetUser200ResponseIntegration.md) |  | [optional] 
**JobTitle** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DeviceCount** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewUsersGetUser200Response

`func NewUsersGetUser200Response() *UsersGetUser200Response`

NewUsersGetUser200Response instantiates a new UsersGetUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersGetUser200ResponseWithDefaults

`func NewUsersGetUser200ResponseWithDefaults() *UsersGetUser200Response`

NewUsersGetUser200ResponseWithDefaults instantiates a new UsersGetUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UsersGetUser200Response) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UsersGetUser200Response) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UsersGetUser200Response) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *UsersGetUser200Response) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetArchived

`func (o *UsersGetUser200Response) GetArchived() int32`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UsersGetUser200Response) GetArchivedOk() (*int32, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UsersGetUser200Response) SetArchived(v int32)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UsersGetUser200Response) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UsersGetUser200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UsersGetUser200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UsersGetUser200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UsersGetUser200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDepartment

`func (o *UsersGetUser200Response) GetDepartment() interface{}`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *UsersGetUser200Response) GetDepartmentOk() (*interface{}, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *UsersGetUser200Response) SetDepartment(v interface{})`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *UsersGetUser200Response) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *UsersGetUser200Response) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *UsersGetUser200Response) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDeprecatedUserId

`func (o *UsersGetUser200Response) GetDeprecatedUserId() string`

GetDeprecatedUserId returns the DeprecatedUserId field if non-nil, zero value otherwise.

### GetDeprecatedUserIdOk

`func (o *UsersGetUser200Response) GetDeprecatedUserIdOk() (*string, bool)`

GetDeprecatedUserIdOk returns a tuple with the DeprecatedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedUserId

`func (o *UsersGetUser200Response) SetDeprecatedUserId(v string)`

SetDeprecatedUserId sets DeprecatedUserId field to given value.

### HasDeprecatedUserId

`func (o *UsersGetUser200Response) HasDeprecatedUserId() bool`

HasDeprecatedUserId returns a boolean if a field has been set.

### GetEmail

`func (o *UsersGetUser200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UsersGetUser200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UsersGetUser200Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UsersGetUser200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *UsersGetUser200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsersGetUser200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsersGetUser200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UsersGetUser200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntegration

`func (o *UsersGetUser200Response) GetIntegration() UsersGetUser200ResponseIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *UsersGetUser200Response) GetIntegrationOk() (*UsersGetUser200ResponseIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *UsersGetUser200Response) SetIntegration(v UsersGetUser200ResponseIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *UsersGetUser200Response) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetJobTitle

`func (o *UsersGetUser200Response) GetJobTitle() interface{}`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *UsersGetUser200Response) GetJobTitleOk() (*interface{}, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *UsersGetUser200Response) SetJobTitle(v interface{})`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *UsersGetUser200Response) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *UsersGetUser200Response) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *UsersGetUser200Response) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetName

`func (o *UsersGetUser200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsersGetUser200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsersGetUser200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UsersGetUser200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeviceCount

`func (o *UsersGetUser200Response) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *UsersGetUser200Response) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *UsersGetUser200Response) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *UsersGetUser200Response) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UsersGetUser200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UsersGetUser200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UsersGetUser200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UsersGetUser200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


