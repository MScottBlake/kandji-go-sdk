# \DeviceInformationAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceInformationCancelLostMode**](DeviceInformationAPI.md#DeviceInformationCancelLostMode) | **Delete** /api/v1/devices/{device_id}/details/lostmode | Cancel Lost Mode
[**DeviceInformationGetDeviceActivity**](DeviceInformationAPI.md#DeviceInformationGetDeviceActivity) | **Get** /api/v1/devices/{device_id}/activity | Get Device Activity
[**DeviceInformationGetDeviceApps**](DeviceInformationAPI.md#DeviceInformationGetDeviceApps) | **Get** /api/v1/devices/{device_id}/apps | Get Device Apps
[**DeviceInformationGetDeviceDetails**](DeviceInformationAPI.md#DeviceInformationGetDeviceDetails) | **Get** /api/v1/devices/{device_id}/details | Get Device Details
[**DeviceInformationGetDeviceLibraryItems**](DeviceInformationAPI.md#DeviceInformationGetDeviceLibraryItems) | **Get** /api/v1/devices/{device_id}/library-items | Get Device Library Items
[**DeviceInformationGetDeviceLostModeDetails**](DeviceInformationAPI.md#DeviceInformationGetDeviceLostModeDetails) | **Get** /api/v1/devices/{device_id}/details/lostmode | Get Device Lost Mode details
[**DeviceInformationGetDeviceParameters**](DeviceInformationAPI.md#DeviceInformationGetDeviceParameters) | **Get** /api/v1/devices/{device_id}/parameters | Get Device Parameters
[**DeviceInformationGetDeviceStatus**](DeviceInformationAPI.md#DeviceInformationGetDeviceStatus) | **Get** /api/v1/devices/{device_id}/status | Get Device Status
[**DeviceInformationListDevices**](DeviceInformationAPI.md#DeviceInformationListDevices) | **Get** /api/v1/devices | List Devices



## DeviceInformationCancelLostMode

> DeviceInformationCancelLostMode(ctx, deviceId).Execute()

Cancel Lost Mode



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
	r, err := apiClient.DeviceInformationAPI.DeviceInformationCancelLostMode(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.DeviceInformationCancelLostMode``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeviceInformationCancelLostModeRequest struct via the builder pattern


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


## DeviceInformationGetDeviceActivity

> map[string]interface{} DeviceInformationGetDeviceActivity(ctx, deviceId).Limit(limit).Offset(offset).Execute()

Get Device Activity



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
	offset := "0" // string | Specify the starting record to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceInformationAPI.DeviceInformationGetDeviceActivity(context.Background(), deviceId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.DeviceInformationGetDeviceActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceInformationGetDeviceActivity`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceInformationAPI.DeviceInformationGetDeviceActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInformationGetDeviceActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return | 

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


## DeviceInformationGetDeviceApps

> map[string]interface{} DeviceInformationGetDeviceApps(ctx, deviceId).Execute()

Get Device Apps



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
	resp, r, err := apiClient.DeviceInformationAPI.DeviceInformationGetDeviceApps(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.DeviceInformationGetDeviceApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceInformationGetDeviceApps`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceInformationAPI.DeviceInformationGetDeviceApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInformationGetDeviceAppsRequest struct via the builder pattern


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


## DeviceInformationGetDeviceDetails

> map[string]interface{} DeviceInformationGetDeviceDetails(ctx, deviceId).Execute()

Get Device Details



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
	resp, r, err := apiClient.DeviceInformationAPI.DeviceInformationGetDeviceDetails(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.DeviceInformationGetDeviceDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceInformationGetDeviceDetails`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceInformationAPI.DeviceInformationGetDeviceDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInformationGetDeviceDetailsRequest struct via the builder pattern


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


## DeviceInformationGetDeviceLibraryItems

> map[string]interface{} DeviceInformationGetDeviceLibraryItems(ctx, deviceId).Execute()

Get Device Library Items



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
	resp, r, err := apiClient.DeviceInformationAPI.DeviceInformationGetDeviceLibraryItems(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.DeviceInformationGetDeviceLibraryItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceInformationGetDeviceLibraryItems`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceInformationAPI.DeviceInformationGetDeviceLibraryItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInformationGetDeviceLibraryItemsRequest struct via the builder pattern


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


## DeviceInformationGetDeviceLostModeDetails

> map[string]interface{} DeviceInformationGetDeviceLostModeDetails(ctx, deviceId).Execute()

Get Device Lost Mode details



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
	resp, r, err := apiClient.DeviceInformationAPI.DeviceInformationGetDeviceLostModeDetails(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.DeviceInformationGetDeviceLostModeDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceInformationGetDeviceLostModeDetails`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceInformationAPI.DeviceInformationGetDeviceLostModeDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInformationGetDeviceLostModeDetailsRequest struct via the builder pattern


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


## DeviceInformationGetDeviceParameters

> map[string]interface{} DeviceInformationGetDeviceParameters(ctx, deviceId).Execute()

Get Device Parameters



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
	resp, r, err := apiClient.DeviceInformationAPI.DeviceInformationGetDeviceParameters(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.DeviceInformationGetDeviceParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceInformationGetDeviceParameters`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceInformationAPI.DeviceInformationGetDeviceParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInformationGetDeviceParametersRequest struct via the builder pattern


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


## DeviceInformationGetDeviceStatus

> map[string]interface{} DeviceInformationGetDeviceStatus(ctx, deviceId).Execute()

Get Device Status



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
	resp, r, err := apiClient.DeviceInformationAPI.DeviceInformationGetDeviceStatus(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.DeviceInformationGetDeviceStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceInformationGetDeviceStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceInformationAPI.DeviceInformationGetDeviceStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInformationGetDeviceStatusRequest struct via the builder pattern


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


## DeviceInformationListDevices

> map[string]interface{} DeviceInformationListDevices(ctx).Limit(limit).AssetTag(assetTag).BlueprintId(blueprintId).DeviceId(deviceId).DeviceName(deviceName).FilevaultEnabled(filevaultEnabled).MacAddress(macAddress).Model(model).Ordering(ordering).OsVersion(osVersion).Platform(platform).SerialNumber(serialNumber).TagName(tagName).TagNameIn(tagNameIn).TagId(tagId).TagIdIn(tagIdIn).User(user).UserEmail(userEmail).UserId(userId).UserName(userName).Offset(offset).Execute()

List Devices



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
	limit := "300" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results.
	assetTag := "23245" // string |  (optional)
	blueprintId := "91f97957-2353-4f86-a1ab-64d2b044a596" // string | Return results &quot;containing&quot; the specified blueprint id (optional)
	deviceId := "2cfeb3ac-3b5d-423e-bcff-e2676a3a32da" // string |  (optional)
	deviceName := "Johnny's MacBook Pro" // string |  (optional)
	filevaultEnabled := "true" // string | <p>Query for devices that either have FileVault on (true) or off (false). This parameter only applies to macOS. </p> <p>An empty list <code>[]</code> will be returned if no devices are found with the given parameter value.</p> (optional)
	macAddress := "00:0c:29:05:43:b6" // string | Search for a specific device by MAC address (optional)
	model := "MacBook Air (M1, 2020)" // string | Return model results &quot;containing&quot; the specified model string. (optional)
	ordering := "device_id" // string | <p>The <code>ordering</code> parameter can be used to define how the device records are ordered in the response. Prepending a dash (-) to the parameter value will reverse the order of the returned results.</p> <p><code>?ordering=-serial_number</code> will order the response by serial_number in descending order.</p> <p><strong>Possible values</strong></p> <ul> <li><code>asset_tag</code></li> <li><code>blueprint_id</code></li> <li><code>device_id</code></li> <li><code>device_name</code></li> <li><code>last_check_in</code> - agent checkin</li> <li><code>model</code></li> <li><code>platform</code></li> <li><code>os_version</code></li> <li><code>serial_number</code></li> <li><code>user</code></li> </ul> <p>Additionally, multiple values can be combined in a comma separated list to further customize the ordering of the response.</p> <p><code>?ordering=serial_number,platform</code></p> (optional)
	osVersion := "13.2.3" // string | Return all device records with the specified OS version (optional)
	platform := "iPad" // string | Return all records matching a specific platform. Possible values:<code>Mac</code>, <code>iPad</code>, <code>iPhone</code>, <code>AppleTV</code> (optional)
	serialNumber := "VMC5qeJ5pDkp" // string | Search for a specific device by Serial Number. If partial serial number is provided in the query, all device containing the partial string will be returned. (optional)
	tagName := "accuhive_01" // string | Return results for given tag name. Case sensitive. (optional)
	tagNameIn := "accuhive_01, accuhive_02" // string | Return results for given tag names separate by commas. Case sensitive. (optional)
	tagId := "tagId_example" // string | Search for a tag by its ID. Case sensitive. (optional)
	tagIdIn := "tagIdIn_example" // string | Return results for given tag IDs separated by commas. Case sensitive. (optional)
	user := "Art Vandelay" // string | Return results &quot;containing&quot; the user name (optional)
	userEmail := "someUser@Kandji.io" // string | Return results &quot;containing&quot; search on email address (optional)
	userId := "1" // string | &quot;exact&quot; match on kandji user ID number (optional)
	userName := "Vandelay" // string | Return results &quot;containing&quot; the assigned user Display Name (optional)
	offset := "0" // string | Specify the starting record to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceInformationAPI.DeviceInformationListDevices(context.Background()).Limit(limit).AssetTag(assetTag).BlueprintId(blueprintId).DeviceId(deviceId).DeviceName(deviceName).FilevaultEnabled(filevaultEnabled).MacAddress(macAddress).Model(model).Ordering(ordering).OsVersion(osVersion).Platform(platform).SerialNumber(serialNumber).TagName(tagName).TagNameIn(tagNameIn).TagId(tagId).TagIdIn(tagIdIn).User(user).UserEmail(userEmail).UserId(userId).UserName(userName).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.DeviceInformationListDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceInformationListDevices`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceInformationAPI.DeviceInformationListDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInformationListDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **assetTag** | **string** |  | 
 **blueprintId** | **string** | Return results &amp;quot;containing&amp;quot; the specified blueprint id | 
 **deviceId** | **string** |  | 
 **deviceName** | **string** |  | 
 **filevaultEnabled** | **string** | &lt;p&gt;Query for devices that either have FileVault on (true) or off (false). This parameter only applies to macOS. &lt;/p&gt; &lt;p&gt;An empty list &lt;code&gt;[]&lt;/code&gt; will be returned if no devices are found with the given parameter value.&lt;/p&gt; | 
 **macAddress** | **string** | Search for a specific device by MAC address | 
 **model** | **string** | Return model results &amp;quot;containing&amp;quot; the specified model string. | 
 **ordering** | **string** | &lt;p&gt;The &lt;code&gt;ordering&lt;/code&gt; parameter can be used to define how the device records are ordered in the response. Prepending a dash (-) to the parameter value will reverse the order of the returned results.&lt;/p&gt; &lt;p&gt;&lt;code&gt;?ordering&#x3D;-serial_number&lt;/code&gt; will order the response by serial_number in descending order.&lt;/p&gt; &lt;p&gt;&lt;strong&gt;Possible values&lt;/strong&gt;&lt;/p&gt; &lt;ul&gt; &lt;li&gt;&lt;code&gt;asset_tag&lt;/code&gt;&lt;/li&gt; &lt;li&gt;&lt;code&gt;blueprint_id&lt;/code&gt;&lt;/li&gt; &lt;li&gt;&lt;code&gt;device_id&lt;/code&gt;&lt;/li&gt; &lt;li&gt;&lt;code&gt;device_name&lt;/code&gt;&lt;/li&gt; &lt;li&gt;&lt;code&gt;last_check_in&lt;/code&gt; - agent checkin&lt;/li&gt; &lt;li&gt;&lt;code&gt;model&lt;/code&gt;&lt;/li&gt; &lt;li&gt;&lt;code&gt;platform&lt;/code&gt;&lt;/li&gt; &lt;li&gt;&lt;code&gt;os_version&lt;/code&gt;&lt;/li&gt; &lt;li&gt;&lt;code&gt;serial_number&lt;/code&gt;&lt;/li&gt; &lt;li&gt;&lt;code&gt;user&lt;/code&gt;&lt;/li&gt; &lt;/ul&gt; &lt;p&gt;Additionally, multiple values can be combined in a comma separated list to further customize the ordering of the response.&lt;/p&gt; &lt;p&gt;&lt;code&gt;?ordering&#x3D;serial_number,platform&lt;/code&gt;&lt;/p&gt; | 
 **osVersion** | **string** | Return all device records with the specified OS version | 
 **platform** | **string** | Return all records matching a specific platform. Possible values:&lt;code&gt;Mac&lt;/code&gt;, &lt;code&gt;iPad&lt;/code&gt;, &lt;code&gt;iPhone&lt;/code&gt;, &lt;code&gt;AppleTV&lt;/code&gt; | 
 **serialNumber** | **string** | Search for a specific device by Serial Number. If partial serial number is provided in the query, all device containing the partial string will be returned. | 
 **tagName** | **string** | Return results for given tag name. Case sensitive. | 
 **tagNameIn** | **string** | Return results for given tag names separate by commas. Case sensitive. | 
 **tagId** | **string** | Search for a tag by its ID. Case sensitive. | 
 **tagIdIn** | **string** | Return results for given tag IDs separated by commas. Case sensitive. | 
 **user** | **string** | Return results &amp;quot;containing&amp;quot; the user name | 
 **userEmail** | **string** | Return results &amp;quot;containing&amp;quot; search on email address | 
 **userId** | **string** | &amp;quot;exact&amp;quot; match on kandji user ID number | 
 **userName** | **string** | Return results &amp;quot;containing&amp;quot; the assigned user Display Name | 
 **offset** | **string** | Specify the starting record to return | 

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

