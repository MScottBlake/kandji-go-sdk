# \VulnerabilitiesAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVulnerabilityDescription**](VulnerabilitiesAPI.md#GetVulnerabilityDescription) | **Get** /api/v1/vulnerability-management/vulnerabilities/{cve_id} | Get Vulnerability Description
[**ListAffectedApps**](VulnerabilitiesAPI.md#ListAffectedApps) | **Get** /api/v1/vulnerability-management/vulnerabilities/{cve_id}/applications | List Affected Apps
[**ListAffectedDevices**](VulnerabilitiesAPI.md#ListAffectedDevices) | **Get** /api/v1/vulnerability-management/vulnerabilities/{cve_id}/devices | List Affected Devices
[**ListDetections**](VulnerabilitiesAPI.md#ListDetections) | **Get** /api/v1/vulnerability-management/detections | List Detections
[**ListVulnerabilities**](VulnerabilitiesAPI.md#ListVulnerabilities) | **Get** /api/v1/vulnerability-management/vulnerabilities | List Vulnerabilities



## GetVulnerabilityDescription

> VulnerabilitiesGetVulnerabilityDescription200Response GetVulnerabilityDescription(ctx, cveId).Execute()

Get Vulnerability Description



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
	cveId := "cveId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnerabilitiesAPI.GetVulnerabilityDescription(context.Background(), cveId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesAPI.GetVulnerabilityDescription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVulnerabilityDescription`: VulnerabilitiesGetVulnerabilityDescription200Response
	fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesAPI.GetVulnerabilityDescription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cveId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVulnerabilityDescriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VulnerabilitiesGetVulnerabilityDescription200Response**](VulnerabilitiesGetVulnerabilityDescription200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAffectedApps

> VulnerabilitiesListAffectedApps200Response ListAffectedApps(ctx, cveId).Page(page).Size(size).SortBy(sortBy).Filter(filter).Execute()

List Affected Apps



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
	cveId := "cveId_example" // string | 
	page := "page_example" // string | The page number of the response. (optional)
	size := "50" // string | A hard upper limit is set at 50  records returned per request. If more records are expected, pagination should be used using the URL value returned in the next attribute. Additionally, filters can be added to a request to limit the results. (optional)
	sortBy := "cve_id" // string | Field to sort by. Example: sort_by=app_name. (optional)
	filter := "{"created_at":{"gte":"2025-05-23T17:11:31.816587Z"}}" // string | <p>Filterable columns:</p> <ul> <li>blueprint_id</li> <li>created_at</li> </ul> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnerabilitiesAPI.ListAffectedApps(context.Background(), cveId).Page(page).Size(size).SortBy(sortBy).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesAPI.ListAffectedApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAffectedApps`: VulnerabilitiesListAffectedApps200Response
	fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesAPI.ListAffectedApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cveId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAffectedAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | The page number of the response. | 
 **size** | **string** | A hard upper limit is set at 50  records returned per request. If more records are expected, pagination should be used using the URL value returned in the next attribute. Additionally, filters can be added to a request to limit the results. | 
 **sortBy** | **string** | Field to sort by. Example: sort_by&#x3D;app_name. | 
 **filter** | **string** | &lt;p&gt;Filterable columns:&lt;/p&gt; &lt;ul&gt; &lt;li&gt;blueprint_id&lt;/li&gt; &lt;li&gt;created_at&lt;/li&gt; &lt;/ul&gt; | 

### Return type

[**VulnerabilitiesListAffectedApps200Response**](VulnerabilitiesListAffectedApps200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAffectedDevices

> VulnerabilitiesListAffectedApps200Response ListAffectedDevices(ctx, cveId).Page(page).Size(size).SortBy(sortBy).Filter(filter).Execute()

List Affected Devices



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
	cveId := "cveId_example" // string | 
	page := "page_example" // string | The page number of the response. (optional)
	size := "50" // string | A hard upper limit is set at 50  records returned per request. If more records are expected, pagination should be used using the URL value returned in the next attribute. Additionally, filters can be added to a request to limit the results. (optional)
	sortBy := "cve_id" // string | Field to sort by. Example: sort_by=app_name. (optional)
	filter := "{"detection_datetime":{"gte":"2025-05-23T17:11:31.816587Z"}}" // string | <p>Filterable columns:</p> <ul> <li>blueprint_id</li> <li>detection_datetime</li> </ul> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnerabilitiesAPI.ListAffectedDevices(context.Background(), cveId).Page(page).Size(size).SortBy(sortBy).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesAPI.ListAffectedDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAffectedDevices`: VulnerabilitiesListAffectedApps200Response
	fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesAPI.ListAffectedDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cveId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAffectedDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | The page number of the response. | 
 **size** | **string** | A hard upper limit is set at 50  records returned per request. If more records are expected, pagination should be used using the URL value returned in the next attribute. Additionally, filters can be added to a request to limit the results. | 
 **sortBy** | **string** | Field to sort by. Example: sort_by&#x3D;app_name. | 
 **filter** | **string** | &lt;p&gt;Filterable columns:&lt;/p&gt; &lt;ul&gt; &lt;li&gt;blueprint_id&lt;/li&gt; &lt;li&gt;detection_datetime&lt;/li&gt; &lt;/ul&gt; | 

### Return type

[**VulnerabilitiesListAffectedApps200Response**](VulnerabilitiesListAffectedApps200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDetections

> VulnerabilitiesListDetections200Response ListDetections(ctx).After(after).Size(size).Filter(filter).Execute()

List Detections



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
	after := "after_example" // string | Cursor token. (optional)
	size := "300" // string | A hard upper <code>limit</code> is set at 300  records returned per request. If more records are expected, pagination should be used using the URL value returned in the next attribute. Additionally, filters can be added to a request to limit the results. (optional)
	filter := "{"cve_id":{"in":["CVE-2024-24795"]}}" // string | <p>Filter on any key attribute within the response.</p> <ul> <li>device_id</li> <li>device_name</li> <li>device_serial_number</li> <li>device_model</li> <li>device_os_version</li> <li>blueprint_id</li> <li>blueprint_name</li> <li>name</li> <li>path</li> <li>version</li> <li>bundle_id</li> <li>cve_id</li> <li>cve_description</li> <li>cve_link</li> <li>cvss_score</li> <li>cvss_severity</li> <li>detection_datetime</li> <li>cve_published_at</li> <li>cve_modified_at</li> </ul> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnerabilitiesAPI.ListDetections(context.Background()).After(after).Size(size).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesAPI.ListDetections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDetections`: VulnerabilitiesListDetections200Response
	fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesAPI.ListDetections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDetectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | Cursor token. | 
 **size** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300  records returned per request. If more records are expected, pagination should be used using the URL value returned in the next attribute. Additionally, filters can be added to a request to limit the results. | 
 **filter** | **string** | &lt;p&gt;Filter on any key attribute within the response.&lt;/p&gt; &lt;ul&gt; &lt;li&gt;device_id&lt;/li&gt; &lt;li&gt;device_name&lt;/li&gt; &lt;li&gt;device_serial_number&lt;/li&gt; &lt;li&gt;device_model&lt;/li&gt; &lt;li&gt;device_os_version&lt;/li&gt; &lt;li&gt;blueprint_id&lt;/li&gt; &lt;li&gt;blueprint_name&lt;/li&gt; &lt;li&gt;name&lt;/li&gt; &lt;li&gt;path&lt;/li&gt; &lt;li&gt;version&lt;/li&gt; &lt;li&gt;bundle_id&lt;/li&gt; &lt;li&gt;cve_id&lt;/li&gt; &lt;li&gt;cve_description&lt;/li&gt; &lt;li&gt;cve_link&lt;/li&gt; &lt;li&gt;cvss_score&lt;/li&gt; &lt;li&gt;cvss_severity&lt;/li&gt; &lt;li&gt;detection_datetime&lt;/li&gt; &lt;li&gt;cve_published_at&lt;/li&gt; &lt;li&gt;cve_modified_at&lt;/li&gt; &lt;/ul&gt; | 

### Return type

[**VulnerabilitiesListDetections200Response**](VulnerabilitiesListDetections200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVulnerabilities

> map[string]interface{} ListVulnerabilities(ctx).Page(page).Size(size).SortBy(sortBy).Filter(filter).Execute()

List Vulnerabilities



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
	page := "page_example" // string | The page number of the response. (optional)
	size := "50" // string | A hard upper limit is set at 50  records returned per request. If more records are expected, pagination should be used using the URL value returned in the next attribute. Additionally, filters can be added to a request to limit the results. (optional)
	sortBy := "cve_id" // string | <p>Field to sort by.</p> <ul> <li>cve_id</li> <li>software (the name of the software)</li> <li>cvss_severity</li> <li>first_detection_date</li> <li>latest_detection_date</li> </ul> (optional)
	filter := "{"cve_id":{"in":["CVE-2024-24795"]}}" // string | <p>Filterable columns</p> <ul> <li>cve_id</li> <li>app_name</li> <li>severity</li> <li>first_detection_date</li> <li>latest_detection_date</li> </ul> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnerabilitiesAPI.ListVulnerabilities(context.Background()).Page(page).Size(size).SortBy(sortBy).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesAPI.ListVulnerabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVulnerabilities`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesAPI.ListVulnerabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVulnerabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | The page number of the response. | 
 **size** | **string** | A hard upper limit is set at 50  records returned per request. If more records are expected, pagination should be used using the URL value returned in the next attribute. Additionally, filters can be added to a request to limit the results. | 
 **sortBy** | **string** | &lt;p&gt;Field to sort by.&lt;/p&gt; &lt;ul&gt; &lt;li&gt;cve_id&lt;/li&gt; &lt;li&gt;software (the name of the software)&lt;/li&gt; &lt;li&gt;cvss_severity&lt;/li&gt; &lt;li&gt;first_detection_date&lt;/li&gt; &lt;li&gt;latest_detection_date&lt;/li&gt; &lt;/ul&gt; | 
 **filter** | **string** | &lt;p&gt;Filterable columns&lt;/p&gt; &lt;ul&gt; &lt;li&gt;cve_id&lt;/li&gt; &lt;li&gt;app_name&lt;/li&gt; &lt;li&gt;severity&lt;/li&gt; &lt;li&gt;first_detection_date&lt;/li&gt; &lt;li&gt;latest_detection_date&lt;/li&gt; &lt;/ul&gt; | 

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

