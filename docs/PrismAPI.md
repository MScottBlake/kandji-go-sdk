# \PrismAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrismActivationLock**](PrismAPI.md#PrismActivationLock) | **Get** /api/v1/prism/activation_lock | Activation lock
[**PrismApplicationFirewall**](PrismAPI.md#PrismApplicationFirewall) | **Get** /api/v1/prism/application_firewall | Application firewall
[**PrismApplications**](PrismAPI.md#PrismApplications) | **Get** /api/v1/prism/apps | Applications
[**PrismCertificates**](PrismAPI.md#PrismCertificates) | **Get** /api/v1/prism/certificates | Certificates
[**PrismCount**](PrismAPI.md#PrismCount) | **Get** /api/v1/prism/count | Count
[**PrismDesktopAndScreensaver**](PrismAPI.md#PrismDesktopAndScreensaver) | **Get** /api/v1/prism/desktop_and_screensaver | Desktop and Screensaver
[**PrismDeviceInformation**](PrismAPI.md#PrismDeviceInformation) | **Get** /api/v1/prism/device_information | Device information
[**PrismFilevault**](PrismAPI.md#PrismFilevault) | **Get** /api/v1/prism/filevault | FileVault
[**PrismGatekeeperAndXprotect**](PrismAPI.md#PrismGatekeeperAndXprotect) | **Get** /api/v1/prism/gatekeeper_and_xprotect | Gatekeeper and XProtect
[**PrismGetCategoryExport**](PrismAPI.md#PrismGetCategoryExport) | **Get** /api/v1/prism/export/{export_id} | Get category export
[**PrismInstalledProfiles**](PrismAPI.md#PrismInstalledProfiles) | **Get** /api/v1/prism/installed_profiles | Installed profiles
[**PrismKernelExtensions**](PrismAPI.md#PrismKernelExtensions) | **Get** /api/v1/prism/kernel_extensions | Kernel Extensions
[**PrismLaunchAgentsAndDaemons**](PrismAPI.md#PrismLaunchAgentsAndDaemons) | **Get** /api/v1/prism/launch_agents_and_daemons | Launch Agents and Daemons
[**PrismLocalUsers**](PrismAPI.md#PrismLocalUsers) | **Get** /api/v1/prism/local_users | Local users
[**PrismRequestCategoryExport**](PrismAPI.md#PrismRequestCategoryExport) | **Post** /api/v1/prism/export | Request category export
[**PrismStartupSettings**](PrismAPI.md#PrismStartupSettings) | **Get** /api/v1/prism/startup_settings | Startup settings
[**PrismSystemExtensions**](PrismAPI.md#PrismSystemExtensions) | **Get** /api/v1/prism/system_extensions | System Extensions
[**PrismTransparencyDatabase**](PrismAPI.md#PrismTransparencyDatabase) | **Get** /api/v1/prism/transparency_database | Transparency database



## PrismActivationLock

> PrismActivationLock(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

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
	r, err := apiClient.PrismAPI.PrismActivationLock(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismActivationLock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrismActivationLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

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


## PrismApplicationFirewall

> map[string]interface{} PrismApplicationFirewall(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.PrismAPI.PrismApplicationFirewall(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismApplicationFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrismApplicationFirewall`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.PrismApplicationFirewall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrismApplicationFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

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


## PrismApplications

> map[string]interface{} PrismApplications(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.PrismAPI.PrismApplications(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrismApplications`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.PrismApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrismApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

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


## PrismCertificates

> map[string]interface{} PrismCertificates(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.PrismAPI.PrismCertificates(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrismCertificates`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.PrismCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrismCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

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


## PrismCount

> map[string]interface{} PrismCount(ctx).Category(category).Execute()

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
	resp, r, err := apiClient.PrismAPI.PrismCount(context.Background()).Category(category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrismCount`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.PrismCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrismCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | &lt;p&gt;Return the count of records for the specified category.  If a category contains spaces substitute the spaces for underscores (&amp;quot;_&amp;quot;) when using the API query.&lt;/p&gt; &lt;p&gt;Examples: apps device_information kernel_extensions system_extensions&lt;/p&gt; | 

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


## PrismDesktopAndScreensaver

> map[string]interface{} PrismDesktopAndScreensaver(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.PrismAPI.PrismDesktopAndScreensaver(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismDesktopAndScreensaver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrismDesktopAndScreensaver`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.PrismDesktopAndScreensaver`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrismDesktopAndScreensaverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

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


## PrismDeviceInformation

> map[string]interface{} PrismDeviceInformation(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Body(body).Execute()

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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.PrismDeviceInformation(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismDeviceInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrismDeviceInformation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.PrismDeviceInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrismDeviceInformationRequest struct via the builder pattern


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

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrismFilevault

> map[string]interface{} PrismFilevault(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.PrismAPI.PrismFilevault(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismFilevault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrismFilevault`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.PrismFilevault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrismFilevaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
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


## PrismGatekeeperAndXprotect

> map[string]interface{} PrismGatekeeperAndXprotect(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.PrismAPI.PrismGatekeeperAndXprotect(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismGatekeeperAndXprotect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrismGatekeeperAndXprotect`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.PrismGatekeeperAndXprotect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrismGatekeeperAndXprotectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Results are limited to Mac only as Gatekeeper and XProtect are not applicable for other platfroms. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
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


## PrismGetCategoryExport

> map[string]interface{} PrismGetCategoryExport(ctx, exportId).Execute()

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
	resp, r, err := apiClient.PrismAPI.PrismGetCategoryExport(context.Background(), exportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismGetCategoryExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrismGetCategoryExport`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.PrismGetCategoryExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrismGetCategoryExportRequest struct via the builder pattern


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


## PrismInstalledProfiles

> map[string]interface{} PrismInstalledProfiles(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.PrismAPI.PrismInstalledProfiles(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismInstalledProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrismInstalledProfiles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.PrismInstalledProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrismInstalledProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

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


## PrismKernelExtensions

> map[string]interface{} PrismKernelExtensions(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.PrismAPI.PrismKernelExtensions(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismKernelExtensions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrismKernelExtensions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.PrismKernelExtensions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrismKernelExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | SON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

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


## PrismLaunchAgentsAndDaemons

> map[string]interface{} PrismLaunchAgentsAndDaemons(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.PrismAPI.PrismLaunchAgentsAndDaemons(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismLaunchAgentsAndDaemons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrismLaunchAgentsAndDaemons`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.PrismLaunchAgentsAndDaemons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrismLaunchAgentsAndDaemonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

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


## PrismLocalUsers

> map[string]interface{} PrismLocalUsers(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.PrismAPI.PrismLocalUsers(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismLocalUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrismLocalUsers`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.PrismLocalUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrismLocalUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

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


## PrismRequestCategoryExport

> map[string]interface{} PrismRequestCategoryExport(ctx).Body(body).Execute()

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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrismAPI.PrismRequestCategoryExport(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismRequestCategoryExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrismRequestCategoryExport`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.PrismRequestCategoryExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrismRequestCategoryExportRequest struct via the builder pattern


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


## PrismStartupSettings

> map[string]interface{} PrismStartupSettings(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.PrismAPI.PrismStartupSettings(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismStartupSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrismStartupSettings`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.PrismStartupSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrismStartupSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
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


## PrismSystemExtensions

> map[string]interface{} PrismSystemExtensions(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.PrismAPI.PrismSystemExtensions(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismSystemExtensions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrismSystemExtensions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.PrismSystemExtensions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrismSystemExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

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


## PrismTransparencyDatabase

> map[string]interface{} PrismTransparencyDatabase(ctx).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.PrismAPI.PrismTransparencyDatabase(context.Background()).BlueprintIds(blueprintIds).DeviceFamilies(deviceFamilies).Filter(filter).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrismAPI.PrismTransparencyDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrismTransparencyDatabase`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrismAPI.PrismTransparencyDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrismTransparencyDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintIds** | **string** | Filter results by one or more blueprint IDs separated by commas. | 
 **deviceFamilies** | **string** | Filter results by one or more device families separate by commas. | 
 **filter** | **string** | JSON schema object containing one or more key value pairs. Note: For detailed information on fiters, see the Filters section at the begining of the Visibility API endpoints in this doc. | 
 **sortBy** | **string** | Sort results by the name of a given response body key in either ascending (default behavior) or descending(&lt;code&gt;-&lt;/code&gt;) order. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

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

