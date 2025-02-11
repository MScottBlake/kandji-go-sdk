# GetDeviceParameters200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewGetDeviceParameters200Response

`func NewGetDeviceParameters200Response() *GetDeviceParameters200Response`

NewGetDeviceParameters200Response instantiates a new GetDeviceParameters200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceParameters200ResponseWithDefaults

`func NewGetDeviceParameters200ResponseWithDefaults() *GetDeviceParameters200Response`

NewGetDeviceParameters200ResponseWithDefaults instantiates a new GetDeviceParameters200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *GetDeviceParameters200Response) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GetDeviceParameters200Response) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GetDeviceParameters200Response) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *GetDeviceParameters200Response) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetParameters

`func (o *GetDeviceParameters200Response) GetParameters() interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *GetDeviceParameters200Response) GetParametersOk() (*interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *GetDeviceParameters200Response) SetParameters(v interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *GetDeviceParameters200Response) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *GetDeviceParameters200Response) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *GetDeviceParameters200Response) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


