# \PrismAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivationLock**](PrismAPI.md#ActivationLock) | **Get** /api/v1/prism/activation_lock | Activation lock
[**ApplicationFirewall**](PrismAPI.md#ApplicationFirewall) | **Get** /api/v1/prism/application_firewall | Application firewall
[**Applications**](PrismAPI.md#Applications) | **Get** /api/v1/prism/apps | Applications
[**Certificates**](PrismAPI.md#Certificates) | **Get** /api/v1/prism/certificates | Certificates
[**Count**](PrismAPI.md#Count) | **Get** /api/v1/prism/count | Count
[**DesktopAndScreensaver**](PrismAPI.md#DesktopAndScreensaver) | **Get** /api/v1/prism/desktop_and_screensaver | Desktop and Screensaver
[**DeviceInformation**](PrismAPI.md#DeviceInformation) | **Get** /api/v1/prism/device_information | Device information
[**Filevault**](PrismAPI.md#Filevault) | **Get** /api/v1/prism/filevault | FileVault
[**GatekeeperAndXprotect**](PrismAPI.md#GatekeeperAndXprotect) | **Get** /api/v1/prism/gatekeeper_and_xprotect | Gatekeeper and XProtect
[**GetCategoryExport**](PrismAPI.md#GetCategoryExport) | **Get** /api/v1/prism/export/{export_id} | Get category export
[**InstalledProfiles**](PrismAPI.md#InstalledProfiles) | **Get** /api/v1/prism/installed_profiles | Installed profiles
[**KernelExtensions**](PrismAPI.md#KernelExtensions) | **Get** /api/v1/prism/kernel_extensions | Kernel Extensions
[**LaunchAgentsAndDaemons**](PrismAPI.md#LaunchAgentsAndDaemons) | **Get** /api/v1/prism/launch_agents_and_daemons | Launch Agents and Daemons
[**LocalUsers**](PrismAPI.md#LocalUsers) | **Get** /api/v1/prism/local_users | Local users
[**RequestCategoryExport**](PrismAPI.md#RequestCategoryExport) | **Post** /api/v1/prism/export | Request category export
[**StartupSettings**](PrismAPI.md#StartupSettings) | **Get** /api/v1/prism/startup_settings | Startup settings
[**SystemExtensions**](PrismAPI.md#SystemExtensions) | **Get** /api/v1/prism/system_extensions | System Extensions
[**TransparencyDatabase**](PrismAPI.md#TransparencyDatabase) | **Get** /api/v1/prism/transparency_database | Transparency database



## ActivationLock

> PrismActivationLock200Response ActivationLock(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

Activation lock



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
	blueprintIds := "blueprint_id, blueprint_id,blueprint_id" // string | Filter results by one or more blueprint IDs separated by commas. (optional)
	deviceFamilies := "Mac,iPhone,iPad" // string | Filter results by one or more device families separate by commas. (optional)
	filter := "filter_example" // string | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. (optional)
	sortBy := "sortBy_example" // string | Sort results by the name of a given response body key in either ascending (default behavior) or descending(<code>-</code>) order. (optional)
	limit := "limit_example" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "offset_example" // string | Specify the starting record to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.ActivationLock(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.ActivationLock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivationLock`: PrismActivationLock200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.ActivationLock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivationLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

### Return type

[**PrismActivationLock200Response**](PrismActivationLock200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationFirewall

> PrismActivationLock200Response ApplicationFirewall(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

Application firewall



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
	blueprintIds := "blueprint_id, blueprint_id,blueprint_id" // string | Filter results by one or more blueprint IDs separated by commas. (optional)
	deviceFamilies := "Mac" // string | Filter results by one or more device families separate by commas. (optional)
	filter := "filter_example" // string | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. (optional)
	sortBy := "sortBy_example" // string | Sort results by the name of a given response body key in either ascending (default behavior) or descending(<code>-</code>) order. (optional)
	limit := "limit_example" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "offset_example" // string | Specify the starting record to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.ApplicationFirewall(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.ApplicationFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplicationFirewall`: PrismActivationLock200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.ApplicationFirewall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

### Return type

[**PrismActivationLock200Response**](PrismActivationLock200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Applications

> PrismApplications200Response Applications(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

Applications



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
	blueprintIds := "blueprint_id, blueprint_id,blueprint_id" // string | Filter results by one or more blueprint IDs separated by commas. (optional)
	deviceFamilies := "Mac,iPhone,iPad,tvOS" // string | Filter results by one or more device families separate by commas. (optional)
	filter := "{"name":{"not_in":["Okta Verify"]},"device__name":{"not_in":["testuser’s MacBook Air"]}}" // string | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. (optional)
	sortBy := "sortBy_example" // string | Sort results by the name of a given response body key in either ascending (default behavior) or descending(<code>-</code>) order. (optional)
	limit := "limit_example" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "offset_example" // string | Specify the starting record to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.Applications(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.Applications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Applications`: PrismApplications200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.Applications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

### Return type

[**PrismApplications200Response**](PrismApplications200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Certificates

> PrismActivationLock200Response Certificates(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

Certificates



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
	blueprintIds := "blueprint_id, blueprint_id,blueprint_id" // string | Filter results by one or more blueprint IDs separated by commas. (optional)
	deviceFamilies := "Mac,iPhone,iPad,tvOS" // string | Filter results by one or more device families separate by commas. (optional)
	filter := "filter_example" // string | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. (optional)
	sortBy := "sortBy_example" // string | Sort results by the name of a given response body key in either ascending (default behavior) or descending(<code>-</code>) order. (optional)
	limit := "limit_example" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "offset_example" // string | Specify the starting record to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.Certificates(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.Certificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Certificates`: PrismActivationLock200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.Certificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

### Return type

[**PrismActivationLock200Response**](PrismActivationLock200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Count

> PrismCount200Response Count(ctx).Category(category).Execute()

Count



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
	category := "apps" // string | <p>Return the count of records for the specified category.  If a category contains spaces substitute the spaces for underscores (&quot;_&quot;) when using the API query.</p> <p>Examples: apps device_information kernel_extensions system_extensions</p>

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.Count(context.Background()).Category(category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.Count``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Count`: PrismCount200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.Count`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | &lt;p&gt;Return the count of records for the specified category.  If a category contains spaces substitute the spaces for underscores (&amp;quot;_&amp;quot;) when using the API query.&lt;/p&gt; &lt;p&gt;Examples: apps device_information kernel_extensions system_extensions&lt;/p&gt; | 

### Return type

[**PrismCount200Response**](PrismCount200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DesktopAndScreensaver

> PrismActivationLock200Response DesktopAndScreensaver(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

Desktop and Screensaver



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
	blueprintIds := "blueprint_id, blueprint_id,blueprint_id" // string | Filter results by one or more blueprint IDs separated by commas. (optional)
	deviceFamilies := "Mac,iPhone,iPad,tvOS" // string | Filter results by one or more device families separate by commas. (optional)
	filter := "filter_example" // string | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. (optional)
	sortBy := "sortBy_example" // string | Sort results by the name of a given response body key in either ascending (default behavior) or descending(<code>-</code>) order. (optional)
	limit := "limit_example" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "offset_example" // string | Specify the starting record to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.DesktopAndScreensaver(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.DesktopAndScreensaver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DesktopAndScreensaver`: PrismActivationLock200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.DesktopAndScreensaver`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDesktopAndScreensaverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

### Return type

[**PrismActivationLock200Response**](PrismActivationLock200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceInformation

> PrismDeviceInformation200Response DeviceInformation(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Body(body).Execute()

Device information



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
	blueprintIds := "14afabf2-7599-47af-a942-bf7f0b8fedf8" // string | Filter results by one or more blueprint IDs separated by commas. (optional)
	deviceFamilies := "Mac,iPhone,iPad,tvOS" // string | Filter results by one or more device families separate by commas. (optional)
	filter := "{"device__name":{"in":["testusers's MacBook Air"]},"updated_at":{"gte":"2023-09-03T04:00:00.000Z","lte":"2023-09-04T04:00:00.000Z"}}" // string | <p>JSON schema object containing one or more key value pairs.</p> <p>Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc.</p> (optional)
	sortBy := "serial_number" // string | Sort results by the name of a given response body key in either ascending (default behavior) or descending(<code>-</code>) order. (optional)
	limit := "limit_example" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "offset_example" // string | Specify the starting record to return (optional)
	body := "{}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.DeviceInformation(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.DeviceInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceInformation`: PrismDeviceInformation200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.DeviceInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | &lt;p&gt;JSON schema object containing one or more key value pairs.&lt;/p&gt; &lt;p&gt;Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc.&lt;/p&gt; | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return | 
 **body** | **string** |  | 

### Return type

[**PrismDeviceInformation200Response**](PrismDeviceInformation200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Filevault

> PrismDeviceInformation200Response Filevault(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

FileVault



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
	blueprintIds := "blueprint_id, blueprint_id,blueprint_id" // string | Filter results by one or more blueprint IDs separated by commas. (optional)
	deviceFamilies := "Mac,iPhone,iPad,tvOS" // string | Filter results by one or more device families separate by commas. (optional)
	filter := "filter_example" // string | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. (optional)
	sortBy := "sortBy_example" // string | Sort results by the name of a given response body key in either ascending (default behavior) or descending(<code>-</code>) order. (optional)
	limit := "limit_example" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "offset_example" // string | Specify the starting record to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.Filevault(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.Filevault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Filevault`: PrismDeviceInformation200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.Filevault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilevaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return | 

### Return type

[**PrismDeviceInformation200Response**](PrismDeviceInformation200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatekeeperAndXprotect

> PrismActivationLock200Response GatekeeperAndXprotect(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

Gatekeeper and XProtect



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
	blueprintIds := "blueprint_id, blueprint_id,blueprint_id" // string | Filter results by one or more blueprint IDs separated by commas. (optional)
	deviceFamilies := "Mac" // string | Results are limited to Mac only as Gatekeeper and XProtect are not applicable for other platfroms. (optional)
	filter := "filter_example" // string | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. (optional)
	sortBy := "sortBy_example" // string | Sort results by the name of a given response body key in either ascending (default behavior) or descending(<code>-</code>) order. (optional)
	limit := "limit_example" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "offset_example" // string | Specify the starting record to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.GatekeeperAndXprotect(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.GatekeeperAndXprotect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatekeeperAndXprotect`: PrismActivationLock200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.GatekeeperAndXprotect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatekeeperAndXprotectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Results are limited to Mac only as Gatekeeper and XProtect are not applicable for other platfroms. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return | 

### Return type

[**PrismActivationLock200Response**](PrismActivationLock200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoryExport

> PrismGetCategoryExport200Response GetCategoryExport(ctx, exportId).Execute()

Get category export



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
	exportId := "exportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.GetCategoryExport(context.Background(), exportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.GetCategoryExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategoryExport`: PrismGetCategoryExport200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.GetCategoryExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrismGetCategoryExport200Response**](PrismGetCategoryExport200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstalledProfiles

> PrismActivationLock200Response InstalledProfiles(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

Installed profiles



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
	blueprintIds := "blueprint_id, blueprint_id,blueprint_id" // string | Filter results by one or more blueprint IDs separated by commas. (optional)
	deviceFamilies := "Mac,iPhone,iPad,tvOS" // string | Filter results by one or more device families separate by commas. (optional)
	filter := "filter_example" // string | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. (optional)
	sortBy := "sortBy_example" // string | Sort results by the name of a given response body key in either ascending (default behavior) or descending(<code>-</code>) order. (optional)
	limit := "limit_example" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "offset_example" // string | Specify the starting record to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.InstalledProfiles(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.InstalledProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstalledProfiles`: PrismActivationLock200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.InstalledProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstalledProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

### Return type

[**PrismActivationLock200Response**](PrismActivationLock200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KernelExtensions

> PrismActivationLock200Response KernelExtensions(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

Kernel Extensions



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
	blueprintIds := "blueprint_id, blueprint_id,blueprint_id" // string | Filter results by one or more blueprint IDs separated by commas. (optional)
	deviceFamilies := "Mac,iPhone,iPad,tvOS" // string | Filter results by one or more device families separate by commas. (optional)
	filter := "filter_example" // string | SON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. (optional)
	sortBy := "sortBy_example" // string | Sort results by the name of a given response body key in either ascending (default behavior) or descending(<code>-</code>) order. (optional)
	limit := "limit_example" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "offset_example" // string | Specify the starting record to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.KernelExtensions(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.KernelExtensions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KernelExtensions`: PrismActivationLock200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.KernelExtensions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKernelExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | SON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

### Return type

[**PrismActivationLock200Response**](PrismActivationLock200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchAgentsAndDaemons

> PrismActivationLock200Response LaunchAgentsAndDaemons(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

Launch Agents and Daemons



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
	blueprintIds := "blueprint_id, blueprint_id,blueprint_id" // string | Filter results by one or more blueprint IDs separated by commas. (optional)
	deviceFamilies := "Mac,iPhone,iPad,tvOS" // string | Filter results by one or more device families separate by commas. (optional)
	filter := "filter_example" // string | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. (optional)
	sortBy := "sortBy_example" // string | Sort results by the name of a given response body key in either ascending (default behavior) or descending(<code>-</code>) order. (optional)
	limit := "limit_example" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "offset_example" // string | Specify the starting record to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.LaunchAgentsAndDaemons(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.LaunchAgentsAndDaemons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchAgentsAndDaemons`: PrismActivationLock200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.LaunchAgentsAndDaemons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLaunchAgentsAndDaemonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

### Return type

[**PrismActivationLock200Response**](PrismActivationLock200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocalUsers

> PrismLocalUsers200Response LocalUsers(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

Local users



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
	blueprintIds := "blueprint_id, blueprint_id,blueprint_id" // string | Filter results by one or more blueprint IDs separated by commas. (optional)
	deviceFamilies := "Mac,iPhone,iPad,tvOS" // string | Filter results by one or more device families separate by commas. (optional)
	filter := "filter_example" // string | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. (optional)
	sortBy := "sortBy_example" // string | Sort results by the name of a given response body key in either ascending (default behavior) or descending(<code>-</code>) order. (optional)
	limit := "limit_example" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "offset_example" // string | Specify the starting record to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.LocalUsers(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.LocalUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocalUsers`: PrismLocalUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.LocalUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocalUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

### Return type

[**PrismLocalUsers200Response**](PrismLocalUsers200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestCategoryExport

> PrismRequestCategoryExport200Response RequestCategoryExport(ctx).Body(body).Execute()

Request category export



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
	body := "{"blueprint_ids":["string","string","string"],"category":"device_information","device_families":["Mac"],"filter":{},"sort_by":"device__nam"}" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.RequestCategoryExport(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.RequestCategoryExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestCategoryExport`: PrismRequestCategoryExport200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.RequestCategoryExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestCategoryExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

[**PrismRequestCategoryExport200Response**](PrismRequestCategoryExport200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartupSettings

> PrismActivationLock200Response StartupSettings(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

Startup settings



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
	blueprintIds := "blueprint_id, blueprint_id,blueprint_id" // string | Filter results by one or more blueprint IDs separated by commas. (optional)
	deviceFamilies := "Mac,iPhone,iPad,tvOS" // string | Filter results by one or more device families separate by commas. (optional)
	filter := "filter_example" // string | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. (optional)
	sortBy := "sortBy_example" // string | Sort results by the name of a given response body key in either ascending (default behavior) or descending(<code>-</code>) order. (optional)
	limit := "limit_example" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "offset_example" // string | Specify the starting record to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.StartupSettings(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.StartupSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartupSettings`: PrismActivationLock200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.StartupSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartupSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return | 

### Return type

[**PrismActivationLock200Response**](PrismActivationLock200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemExtensions

> PrismActivationLock200Response SystemExtensions(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

System Extensions



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
	blueprintIds := "blueprint_id, blueprint_id,blueprint_id" // string | Filter results by one or more blueprint IDs separated by commas. (optional)
	deviceFamilies := "Mac,iPhone,iPad,tvOS" // string | Filter results by one or more device families separate by commas. (optional)
	filter := "filter_example" // string | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. (optional)
	sortBy := "sortBy_example" // string | Sort results by the name of a given response body key in either ascending (default behavior) or descending(<code>-</code>) order. (optional)
	limit := "limit_example" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "offset_example" // string | Specify the starting record to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.SystemExtensions(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.SystemExtensions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemExtensions`: PrismActivationLock200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.SystemExtensions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

### Return type

[**PrismActivationLock200Response**](PrismActivationLock200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransparencyDatabase

> PrismActivationLock200Response TransparencyDatabase(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

Transparency database



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
	blueprintIds := "blueprint_id, blueprint_id,blueprint_id" // string | Filter results by one or more blueprint IDs separated by commas. (optional)
	deviceFamilies := "Mac,iPhone,iPad,tvOS" // string | Filter results by one or more device families separate by commas. (optional)
	filter := "filter_example" // string | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. (optional)
	sortBy := "sortBy_example" // string | Sort results by the name of a given response body key in either ascending (default behavior) or descending(<code>-</code>) order. (optional)
	limit := "limit_example" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "offset_example" // string | Specify the starting record to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.TransparencyDatabase(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.TransparencyDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransparencyDatabase`: PrismActivationLock200Response
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.TransparencyDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransparencyDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

### Return type

[**PrismActivationLock200Response**](PrismActivationLock200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

