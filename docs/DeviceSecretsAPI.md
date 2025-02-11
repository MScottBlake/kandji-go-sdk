# \DeviceSecretsAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActivationLockBypassCode**](DeviceSecretsAPI.md#GetActivationLockBypassCode) | **Get** /api/v1/devices/{device_id}/secrets/bypasscode | Get Activation Lock Bypass Code
[**GetFilevaultRecoveryKey**](DeviceSecretsAPI.md#GetFilevaultRecoveryKey) | **Get** /api/v1/devices/{device_id}/secrets/filevaultkey | Get FileVault Recovery Key
[**GetRecoveryLockPassword**](DeviceSecretsAPI.md#GetRecoveryLockPassword) | **Get** /api/v1/devices/{device_id}/secrets/recoverypassword | Get Recovery Lock Password
[**GetUnlockPin**](DeviceSecretsAPI.md#GetUnlockPin) | **Get** /api/v1/devices/{device_id}/secrets/unlockpin | Get Unlock Pin



## GetActivationLockBypassCode

> Success200Response5 GetActivationLockBypassCode(ctx, deviceId).Execute()

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
	resp, r, err := apiClient.DeviceSecretsAPI.GetActivationLockBypassCode(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceSecretsAPI.GetActivationLockBypassCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivationLockBypassCode`: Success200Response5
	fmt.Fprintf(os.Stdout, "Response from `DeviceSecretsAPI.GetActivationLockBypassCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivationLockBypassCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Success200Response5**](Success200Response5.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilevaultRecoveryKey

> Success200Response6 GetFilevaultRecoveryKey(ctx, deviceId).Execute()

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
	resp, r, err := apiClient.DeviceSecretsAPI.GetFilevaultRecoveryKey(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceSecretsAPI.GetFilevaultRecoveryKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilevaultRecoveryKey`: Success200Response6
	fmt.Fprintf(os.Stdout, "Response from `DeviceSecretsAPI.GetFilevaultRecoveryKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilevaultRecoveryKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Success200Response6**](Success200Response6.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecoveryLockPassword

> Success200Response8 GetRecoveryLockPassword(ctx, deviceId).Execute()

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
	resp, r, err := apiClient.DeviceSecretsAPI.GetRecoveryLockPassword(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceSecretsAPI.GetRecoveryLockPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecoveryLockPassword`: Success200Response8
	fmt.Fprintf(os.Stdout, "Response from `DeviceSecretsAPI.GetRecoveryLockPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecoveryLockPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Success200Response8**](Success200Response8.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnlockPin

> Success200Response7 GetUnlockPin(ctx, deviceId).Execute()

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
	resp, r, err := apiClient.DeviceSecretsAPI.GetUnlockPin(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceSecretsAPI.GetUnlockPin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUnlockPin`: Success200Response7
	fmt.Fprintf(os.Stdout, "Response from `DeviceSecretsAPI.GetUnlockPin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnlockPinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Success200Response7**](Success200Response7.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

