# \DeviceActionsAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceActionsClearPasscode**](DeviceActionsAPI.md#DeviceActionsClearPasscode) | **Post** /api/v1/devices/{device_id}/action/clearpasscode | Clear Passcode
[**DeviceActionsDeleteDevice**](DeviceActionsAPI.md#DeviceActionsDeleteDevice) | **Delete** /api/v1/devices/{device_id} | Delete Device
[**DeviceActionsDeleteUser**](DeviceActionsAPI.md#DeviceActionsDeleteUser) | **Post** /api/v1/devices/{device_id}/action/deleteuser | Delete User
[**DeviceActionsEraseDevice**](DeviceActionsAPI.md#DeviceActionsEraseDevice) | **Post** /api/v1/devices/{device_id}/action/erase | Erase Device
[**DeviceActionsGetDeviceCommands**](DeviceActionsAPI.md#DeviceActionsGetDeviceCommands) | **Get** /api/v1/devices/{device_id}/commands | Get Device Commands
[**DeviceActionsLockDevice**](DeviceActionsAPI.md#DeviceActionsLockDevice) | **Post** /api/v1/devices/{device_id}/action/lock | Lock Device
[**DeviceActionsReinstallAgent**](DeviceActionsAPI.md#DeviceActionsReinstallAgent) | **Post** /api/v1/devices/{device_id}/action/reinstallagent | Reinstall Agent
[**DeviceActionsRemoteDesktop**](DeviceActionsAPI.md#DeviceActionsRemoteDesktop) | **Post** /api/v1/devices/{device_id}/action/remotedesktop | Remote Desktop
[**DeviceActionsRenewMdmProfile**](DeviceActionsAPI.md#DeviceActionsRenewMdmProfile) | **Post** /api/v1/devices/{device_id}/action/renewmdmprofile | Renew MDM Profile
[**DeviceActionsRestartDevice**](DeviceActionsAPI.md#DeviceActionsRestartDevice) | **Post** /api/v1/devices/{device_id}/action/restart | Restart Device
[**DeviceActionsSendBlankpush**](DeviceActionsAPI.md#DeviceActionsSendBlankpush) | **Post** /api/v1/devices/{device_id}/action/blankpush | Send Blankpush
[**DeviceActionsSetName**](DeviceActionsAPI.md#DeviceActionsSetName) | **Post** /api/v1/devices/{device_id}/action/setname | Set Name
[**DeviceActionsShutdown**](DeviceActionsAPI.md#DeviceActionsShutdown) | **Post** /api/v1/devices/{device_id}/action/shutdown | Shutdown
[**DeviceActionsUnlockAccount**](DeviceActionsAPI.md#DeviceActionsUnlockAccount) | **Post** /api/v1/devices/{device_id}/action/unlockaccount | Unlock Account
[**DeviceActionsUpdateInventory**](DeviceActionsAPI.md#DeviceActionsUpdateInventory) | **Post** /api/v1/devices/{device_id}/action/updateinventory | Update Inventory



## DeviceActionsClearPasscode

> DeviceActionsClearPasscode(ctx, deviceId).Execute()

Clear Passcode



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
	r, err := apiClient.DeviceActionsAPI.DeviceActionsClearPasscode(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceActionsAPI.DeviceActionsClearPasscode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceActionsClearPasscodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceActionsDeleteDevice

> DeviceActionsDeleteDevice(ctx, deviceId).Execute()

Delete Device



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
	r, err := apiClient.DeviceActionsAPI.DeviceActionsDeleteDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceActionsAPI.DeviceActionsDeleteDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceActionsDeleteDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceActionsDeleteUser

> DeviceActionsDeleteUser(ctx, deviceId).Body(body).Execute()

Delete User



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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceActionsAPI.DeviceActionsDeleteUser(context.Background(), deviceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceActionsAPI.DeviceActionsDeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceActionsDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceActionsEraseDevice

> DeviceActionsEraseDevice(ctx, deviceId).Body(body).Execute()

Erase Device



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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceActionsAPI.DeviceActionsEraseDevice(context.Background(), deviceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceActionsAPI.DeviceActionsEraseDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceActionsEraseDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceActionsGetDeviceCommands

> DeviceActionsGetDeviceCommands(ctx, deviceId).Limit(limit).Offset(offset).Execute()

Get Device Commands



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
	limit := "300" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results.
	offset := "offset_example" // string | Specify the starting record to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceActionsAPI.DeviceActionsGetDeviceCommands(context.Background(), deviceId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceActionsAPI.DeviceActionsGetDeviceCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceActionsGetDeviceCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceActionsLockDevice

> map[string]interface{} DeviceActionsLockDevice(ctx, deviceId).Body(body).Execute()

Lock Device



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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceActionsAPI.DeviceActionsLockDevice(context.Background(), deviceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceActionsAPI.DeviceActionsLockDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceActionsLockDevice`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceActionsAPI.DeviceActionsLockDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceActionsLockDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceActionsReinstallAgent

> DeviceActionsReinstallAgent(ctx, deviceId).Execute()

Reinstall Agent



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
	r, err := apiClient.DeviceActionsAPI.DeviceActionsReinstallAgent(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceActionsAPI.DeviceActionsReinstallAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceActionsReinstallAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceActionsRemoteDesktop

> DeviceActionsRemoteDesktop(ctx, deviceId).Body(body).Execute()

Remote Desktop



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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceActionsAPI.DeviceActionsRemoteDesktop(context.Background(), deviceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceActionsAPI.DeviceActionsRemoteDesktop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceActionsRemoteDesktopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceActionsRenewMdmProfile

> DeviceActionsRenewMdmProfile(ctx, deviceId).Execute()

Renew MDM Profile



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
	r, err := apiClient.DeviceActionsAPI.DeviceActionsRenewMdmProfile(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceActionsAPI.DeviceActionsRenewMdmProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceActionsRenewMdmProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceActionsRestartDevice

> DeviceActionsRestartDevice(ctx, deviceId).Body(body).Execute()

Restart Device



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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceActionsAPI.DeviceActionsRestartDevice(context.Background(), deviceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceActionsAPI.DeviceActionsRestartDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceActionsRestartDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceActionsSendBlankpush

> DeviceActionsSendBlankpush(ctx, deviceId).Execute()

Send Blankpush



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
	r, err := apiClient.DeviceActionsAPI.DeviceActionsSendBlankpush(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceActionsAPI.DeviceActionsSendBlankpush``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceActionsSendBlankpushRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceActionsSetName

> DeviceActionsSetName(ctx, deviceId).Body(body).Execute()

Set Name



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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceActionsAPI.DeviceActionsSetName(context.Background(), deviceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceActionsAPI.DeviceActionsSetName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceActionsSetNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceActionsShutdown

> DeviceActionsShutdown(ctx, deviceId).Execute()

Shutdown



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
	r, err := apiClient.DeviceActionsAPI.DeviceActionsShutdown(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceActionsAPI.DeviceActionsShutdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceActionsShutdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceActionsUnlockAccount

> DeviceActionsUnlockAccount(ctx, deviceId).Body(body).Execute()

Unlock Account



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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceActionsAPI.DeviceActionsUnlockAccount(context.Background(), deviceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceActionsAPI.DeviceActionsUnlockAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceActionsUnlockAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceActionsUpdateInventory

> DeviceActionsUpdateInventory(ctx, deviceId).Execute()

Update Inventory



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
	r, err := apiClient.DeviceActionsAPI.DeviceActionsUpdateInventory(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceActionsAPI.DeviceActionsUpdateInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceActionsUpdateInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

