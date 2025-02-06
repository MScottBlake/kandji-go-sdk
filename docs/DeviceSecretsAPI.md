# \DeviceSecretsAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceSecretsGetActivationLockBypassCode**](DeviceSecretsAPI.md#DeviceSecretsGetActivationLockBypassCode) | **Get** /api/v1/devices/{device_id}/secrets/bypasscode | Get Activation Lock Bypass Code
[**DeviceSecretsGetFilevaultRecoveryKey**](DeviceSecretsAPI.md#DeviceSecretsGetFilevaultRecoveryKey) | **Get** /api/v1/devices/{device_id}/secrets/filevaultkey | Get FileVault Recovery Key
[**DeviceSecretsGetRecoveryLockPassword**](DeviceSecretsAPI.md#DeviceSecretsGetRecoveryLockPassword) | **Get** /api/v1/devices/{device_id}/secrets/recoverypassword | Get Recovery Lock Password
[**DeviceSecretsGetUnlockPin**](DeviceSecretsAPI.md#DeviceSecretsGetUnlockPin) | **Get** /api/v1/devices/{device_id}/secrets/unlockpin | Get Unlock Pin



## DeviceSecretsGetActivationLockBypassCode

> map[string]interface{} DeviceSecretsGetActivationLockBypassCode(ctx, deviceId).Execute()

Get Activation Lock Bypass Code



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/MScottBlake/kandji-go-sdk"
)

func main() {
	deviceId := "deviceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceSecretsAPI.DeviceSecretsGetActivationLockBypassCode(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceSecretsAPI.DeviceSecretsGetActivationLockBypassCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceSecretsGetActivationLockBypassCode`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceSecretsAPI.DeviceSecretsGetActivationLockBypassCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceSecretsGetActivationLockBypassCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceSecretsGetFilevaultRecoveryKey

> map[string]interface{} DeviceSecretsGetFilevaultRecoveryKey(ctx, deviceId).Execute()

Get FileVault Recovery Key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/MScottBlake/kandji-go-sdk"
)

func main() {
	deviceId := "deviceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceSecretsAPI.DeviceSecretsGetFilevaultRecoveryKey(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceSecretsAPI.DeviceSecretsGetFilevaultRecoveryKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceSecretsGetFilevaultRecoveryKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceSecretsAPI.DeviceSecretsGetFilevaultRecoveryKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceSecretsGetFilevaultRecoveryKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceSecretsGetRecoveryLockPassword

> map[string]interface{} DeviceSecretsGetRecoveryLockPassword(ctx, deviceId).Execute()

Get Recovery Lock Password



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/MScottBlake/kandji-go-sdk"
)

func main() {
	deviceId := "deviceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceSecretsAPI.DeviceSecretsGetRecoveryLockPassword(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceSecretsAPI.DeviceSecretsGetRecoveryLockPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceSecretsGetRecoveryLockPassword`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceSecretsAPI.DeviceSecretsGetRecoveryLockPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceSecretsGetRecoveryLockPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceSecretsGetUnlockPin

> map[string]interface{} DeviceSecretsGetUnlockPin(ctx, deviceId).Execute()

Get Unlock Pin



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/MScottBlake/kandji-go-sdk"
)

func main() {
	deviceId := "deviceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceSecretsAPI.DeviceSecretsGetUnlockPin(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceSecretsAPI.DeviceSecretsGetUnlockPin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceSecretsGetUnlockPin`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceSecretsAPI.DeviceSecretsGetUnlockPin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceSecretsGetUnlockPinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

