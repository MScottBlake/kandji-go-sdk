# InlineObject7

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
**EnrollmentCode** | Pointer to [**InlineObject5EnrollmentCode**](InlineObject5EnrollmentCode.md) |  | [optional] 

## Methods

### NewInlineObject7

`func NewInlineObject7() *InlineObject7`

NewInlineObject7 instantiates a new InlineObject7 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject7WithDefaults

`func NewInlineObject7WithDefaults() *InlineObject7`

NewInlineObject7WithDefaults instantiates a new InlineObject7 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineObject7) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineObject7) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineObject7) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineObject7) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineObject7) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject7) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject7) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject7) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIcon

`func (o *InlineObject7) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *InlineObject7) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *InlineObject7) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *InlineObject7) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetColor

`func (o *InlineObject7) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *InlineObject7) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *InlineObject7) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *InlineObject7) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *InlineObject7) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineObject7) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineObject7) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineObject7) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParams

`func (o *InlineObject7) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *InlineObject7) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *InlineObject7) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *InlineObject7) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetComputersCount

`func (o *InlineObject7) GetComputersCount() int32`

GetComputersCount returns the ComputersCount field if non-nil, zero value otherwise.

### GetComputersCountOk

`func (o *InlineObject7) GetComputersCountOk() (*int32, bool)`

GetComputersCountOk returns a tuple with the ComputersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputersCount

`func (o *InlineObject7) SetComputersCount(v int32)`

SetComputersCount sets ComputersCount field to given value.

### HasComputersCount

`func (o *InlineObject7) HasComputersCount() bool`

HasComputersCount returns a boolean if a field has been set.

### GetMissingComputersCount

`func (o *InlineObject7) GetMissingComputersCount() int32`

GetMissingComputersCount returns the MissingComputersCount field if non-nil, zero value otherwise.

### GetMissingComputersCountOk

`func (o *InlineObject7) GetMissingComputersCountOk() (*int32, bool)`

GetMissingComputersCountOk returns a tuple with the MissingComputersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingComputersCount

`func (o *InlineObject7) SetMissingComputersCount(v int32)`

SetMissingComputersCount sets MissingComputersCount field to given value.

### HasMissingComputersCount

`func (o *InlineObject7) HasMissingComputersCount() bool`

HasMissingComputersCount returns a boolean if a field has been set.

### GetEnrollmentCode

`func (o *InlineObject7) GetEnrollmentCode() InlineObject5EnrollmentCode`

GetEnrollmentCode returns the EnrollmentCode field if non-nil, zero value otherwise.

### GetEnrollmentCodeOk

`func (o *InlineObject7) GetEnrollmentCodeOk() (*InlineObject5EnrollmentCode, bool)`

GetEnrollmentCodeOk returns a tuple with the EnrollmentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentCode

`func (o *InlineObject7) SetEnrollmentCode(v InlineObject5EnrollmentCode)`

SetEnrollmentCode sets EnrollmentCode field to given value.

### HasEnrollmentCode

`func (o *InlineObject7) HasEnrollmentCode() bool`

HasEnrollmentCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


