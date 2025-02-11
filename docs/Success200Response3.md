# Success200Response3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **map[string]interface{}** |  | [optional] 
**ComputersCount** | Pointer to **int32** |  | [optional] 
**MissingComputersCount** | Pointer to **int32** |  | [optional] 
**EnrollmentCode** | Pointer to [**Success201ResponseEnrollmentCode**](Success201ResponseEnrollmentCode.md) |  | [optional] 

## Methods

### NewSuccess200Response3

`func NewSuccess200Response3() *Success200Response3`

NewSuccess200Response3 instantiates a new Success200Response3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccess200Response3WithDefaults

`func NewSuccess200Response3WithDefaults() *Success200Response3`

NewSuccess200Response3WithDefaults instantiates a new Success200Response3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Success200Response3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Success200Response3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Success200Response3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Success200Response3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Success200Response3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Success200Response3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Success200Response3) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Success200Response3) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIcon

`func (o *Success200Response3) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Success200Response3) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Success200Response3) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Success200Response3) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetColor

`func (o *Success200Response3) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Success200Response3) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Success200Response3) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Success200Response3) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *Success200Response3) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Success200Response3) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Success200Response3) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Success200Response3) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParams

`func (o *Success200Response3) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Success200Response3) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Success200Response3) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *Success200Response3) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetComputersCount

`func (o *Success200Response3) GetComputersCount() int32`

GetComputersCount returns the ComputersCount field if non-nil, zero value otherwise.

### GetComputersCountOk

`func (o *Success200Response3) GetComputersCountOk() (*int32, bool)`

GetComputersCountOk returns a tuple with the ComputersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputersCount

`func (o *Success200Response3) SetComputersCount(v int32)`

SetComputersCount sets ComputersCount field to given value.

### HasComputersCount

`func (o *Success200Response3) HasComputersCount() bool`

HasComputersCount returns a boolean if a field has been set.

### GetMissingComputersCount

`func (o *Success200Response3) GetMissingComputersCount() int32`

GetMissingComputersCount returns the MissingComputersCount field if non-nil, zero value otherwise.

### GetMissingComputersCountOk

`func (o *Success200Response3) GetMissingComputersCountOk() (*int32, bool)`

GetMissingComputersCountOk returns a tuple with the MissingComputersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingComputersCount

`func (o *Success200Response3) SetMissingComputersCount(v int32)`

SetMissingComputersCount sets MissingComputersCount field to given value.

### HasMissingComputersCount

`func (o *Success200Response3) HasMissingComputersCount() bool`

HasMissingComputersCount returns a boolean if a field has been set.

### GetEnrollmentCode

`func (o *Success200Response3) GetEnrollmentCode() Success201ResponseEnrollmentCode`

GetEnrollmentCode returns the EnrollmentCode field if non-nil, zero value otherwise.

### GetEnrollmentCodeOk

`func (o *Success200Response3) GetEnrollmentCodeOk() (*Success201ResponseEnrollmentCode, bool)`

GetEnrollmentCodeOk returns a tuple with the EnrollmentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCode

`func (o *Success200Response3) SetEnrollmentCode(v Success201ResponseEnrollmentCode)`

SetEnrollmentCode sets EnrollmentCode field to given value.

### HasEnrollmentCode

`func (o *Success200Response3) HasEnrollmentCode() bool`

HasEnrollmentCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


