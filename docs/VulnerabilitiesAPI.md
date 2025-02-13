# \VulnerabilitiesAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVulnerabilityDescription**](VulnerabilitiesAPI.md#GetVulnerabilityDescription) | **Get** /api/v1/vulnerability-management/vulnerabilities/{cve_id} | Get Vulnerability Description
[**ListAffectedApplications**](VulnerabilitiesAPI.md#ListAffectedApplications) | **Get** /api/v1/vulnerability-management/vulnerabilities/{cve_id}/applications | List Affected Applications
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


## ListAffectedApplications

> map[string]interface{} ListAffectedApplications(ctx, cveId).Page(page).Size(size).SortBy(sortBy).Filter(filter).Execute()

List Affected Applications



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
	filter := "device_serial_number" // string | Filterable columns: blueprint_id updated_at (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnerabilitiesAPI.ListAffectedApplications(context.Background(), cveId).Page(page).Size(size).SortBy(sortBy).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesAPI.ListAffectedApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAffectedApplications`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesAPI.ListAffectedApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cveId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAffectedApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | The page number of the response. | 
 **size** | **string** | A hard upper limit is set at 50  records returned per request. If more records are expected, pagination should be used using the URL value returned in the next attribute. Additionally, filters can be added to a request to limit the results. | 
 **sortBy** | **string** | Field to sort by. Example: sort_by&#x3D;app_name. | 
 **filter** | **string** | Filterable columns: blueprint_id updated_at | 

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


## ListAffectedDevices

> map[string]interface{} ListAffectedDevices(ctx, cveId).Page(page).Size(size).SortBy(sortBy).Filter(filter).Execute()

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
	filter := "device_serial_number" // string | Filterable columns: blueprint_id updated_at (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnerabilitiesAPI.ListAffectedDevices(context.Background(), cveId).Page(page).Size(size).SortBy(sortBy).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesAPI.ListAffectedDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAffectedDevices`: map[string]interface{}
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
 **filter** | **string** | Filterable columns: blueprint_id updated_at | 

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


## ListDetections

> VulnerabilitiesListDetections200Response ListDetections(ctx).After(after).Limit(limit).Filter(filter).Execute()

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
	limit := "300" // string | A hard upper <code>limit</code> is set at 300  records returned per request. If more records are expected, pagination should be used using the URL value returned in the next attribute. Additionally, filters can be added to a request to limit the results. (optional)
	filter := "device_serial_number" // string | Can filter on any key attribute within the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnerabilitiesAPI.ListDetections(context.Background()).After(after).Limit(limit).Filter(filter).Execute()
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
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300  records returned per request. If more records are expected, pagination should be used using the URL value returned in the next attribute. Additionally, filters can be added to a request to limit the results. | 
 **filter** | **string** | Can filter on any key attribute within the response. | 

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

> VulnerabilitiesListVulnerabilities200Response ListVulnerabilities(ctx).Page(page).Size(size).SortBy(sortBy).Filter(filter).Execute()

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
	sortBy := "cve_id" // string | Field to sort by. Example: sort_by=cve_id. (optional)
	filter := "device_serial_number" // string | <p>Filterable columns:</p> <p>cve_id app_name severity first_detection_date latest_detection_date</p> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnerabilitiesAPI.ListVulnerabilities(context.Background()).Page(page).Size(size).SortBy(sortBy).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesAPI.ListVulnerabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVulnerabilities`: VulnerabilitiesListVulnerabilities200Response
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
 **sortBy** | **string** | Field to sort by. Example: sort_by&#x3D;cve_id. | 
 **filter** | **string** | &lt;p&gt;Filterable columns:&lt;/p&gt; &lt;p&gt;cve_id app_name severity first_detection_date latest_detection_date&lt;/p&gt; | 

### Return type

[**VulnerabilitiesListVulnerabilities200Response**](VulnerabilitiesListVulnerabilities200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

