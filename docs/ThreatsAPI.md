# \ThreatsAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBehavioralDetections**](ThreatsAPI.md#GetBehavioralDetections) | **Get** /api/v1/behavioral-detections | Get Behavioral Detections
[**GetThreatDetails**](ThreatsAPI.md#GetThreatDetails) | **Get** /api/v1/threat-details | Get Threat Details



## GetBehavioralDetections

> ThreatsGetBehavioralDetections200Response GetBehavioralDetections(ctx).ThreatId(threatId).Classification(classification).Status(status).DateRange(dateRange).DetectionDate(detectionDate).DeviceId(deviceId).MalwareFamily(malwareFamily).ParentProcessName(parentProcessName).TargetProcessName(targetProcessName).InformationalTags(informationalTags).Term(term).SortBy(sortBy).Limit(limit).Offset(offset).Execute()

Get Behavioral Detections



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
	threatId := "threatId_example" // string | Filter by a specific threat ID (threat_id=Kandji_BD_0096). (optional)
	classification := "classification_example" // string | Filter by threat classification (classification=malicious). (optional)
	status := "status_example" // string | Filter by threat status (threat_status=blocked) (optional)
	dateRange := "dateRange_example" // string | Return all records within a specified number of days (Int) (optional)
	detectionDate := "detectionDate_example" // string | two query params detection_date_from and detection_date_to (optional)
	deviceId := "deviceId_example" // string | Search for a specific device by the device id (uuid). (optional)
	malwareFamily := "malwareFamily_example" // string | Filter by malware family (malware_family=TrickBot). (optional)
	parentProcessName := "parentProcessName_example" // string | Filter by parent process (parent_process_name=bash). (optional)
	targetProcessName := "targetProcessName_example" // string | Filter by target process (target_process_name=python). (optional)
	informationalTags := "informationalTags_example" // string | Filter by tags (informational_tags=exploit,privilege_escalation). (optional)
	term := "term_example" // string | Search term to filter threat results. Device name, file hash, image path (optional)
	sortBy := "status" // string | <p>Detections can be sorted by any of the following keys. Prepending a dash (-) to the parameter value will reverse the order of the returned results. ?sort_by=-device_name will order the response by device_name in descending order.</p> <ul> <li>threat_id</li> <li>classification</li> <li>device_name</li> <li>parent_process_name</li> <li>target_process_name</li> <li>detection_date</li> <li>status</li> </ul> (optional)
	limit := "1000" // string | A hard upper <code>limit</code> is set at 1000 records returned per request. If more records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters.  Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "offset_example" // string | Specify the starting record to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreatsAPI.GetBehavioralDetections(context.Background()).ThreatId(threatId).Classification(classification).Status(status).DateRange(dateRange).DetectionDate(detectionDate).DeviceId(deviceId).MalwareFamily(malwareFamily).ParentProcessName(parentProcessName).TargetProcessName(targetProcessName).InformationalTags(informationalTags).Term(term).SortBy(sortBy).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatsAPI.GetBehavioralDetections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBehavioralDetections`: ThreatsGetBehavioralDetections200Response
	fmt.Fprintf(os.Stdout, "Response from `ThreatsAPI.GetBehavioralDetections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBehavioralDetectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threatId** | **string** | Filter by a specific threat ID (threat_id&#x3D;Kandji_BD_0096). | 
 **classification** | **string** | Filter by threat classification (classification&#x3D;malicious). | 
 **status** | **string** | Filter by threat status (threat_status&#x3D;blocked) | 
 **dateRange** | **string** | Return all records within a specified number of days (Int) | 
 **detectionDate** | **string** | two query params detection_date_from and detection_date_to | 
 **deviceId** | **string** | Search for a specific device by the device id (uuid). | 
 **malwareFamily** | **string** | Filter by malware family (malware_family&#x3D;TrickBot). | 
 **parentProcessName** | **string** | Filter by parent process (parent_process_name&#x3D;bash). | 
 **targetProcessName** | **string** | Filter by target process (target_process_name&#x3D;python). | 
 **informationalTags** | **string** | Filter by tags (informational_tags&#x3D;exploit,privilege_escalation). | 
 **term** | **string** | Search term to filter threat results. Device name, file hash, image path | 
 **sortBy** | **string** | &lt;p&gt;Detections can be sorted by any of the following keys. Prepending a dash (-) to the parameter value will reverse the order of the returned results. ?sort_by&#x3D;-device_name will order the response by device_name in descending order.&lt;/p&gt; &lt;ul&gt; &lt;li&gt;threat_id&lt;/li&gt; &lt;li&gt;classification&lt;/li&gt; &lt;li&gt;device_name&lt;/li&gt; &lt;li&gt;parent_process_name&lt;/li&gt; &lt;li&gt;target_process_name&lt;/li&gt; &lt;li&gt;detection_date&lt;/li&gt; &lt;li&gt;status&lt;/li&gt; &lt;/ul&gt; | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 1000 records returned per request. If more records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters.  Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return. | 

### Return type

[**ThreatsGetBehavioralDetections200Response**](ThreatsGetBehavioralDetections200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThreatDetails

> ThreatsGetThreatDetails200Response GetThreatDetails(ctx).Classification(classification).DateRange(dateRange).DeviceId(deviceId).Status(status).SortBy(sortBy).Term(term).Limit(limit).Offset(offset).Execute()

Get Threat Details



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
	classification := "malware" // string | Return all records matching a specified classification. The following classification options are available: <code>malware</code> and <code>pup</code>. Leave this parameter empty to return all classification types. (optional)
	dateRange := "7" // string | Return all records within a specified number of days. Any positive number of days may be specified. Examples: <code>7</code>, <code>30</code>, <code>60</code>, <code>90</code>, <code>180</code>, or <code>365</code>. (optional)
	deviceId := "15fcec08-xxxx-xxxx-xxxx-7c2f950910eb" // string |  (optional)
	status := "quarantined" // string | Return all records matching a specified status. The following status options are available: <code>quarantined</code>, <code>not_quarantined</code>, or <code>released</code>. Leave this parameter empty to return all status types. (optional)
	sortBy := "status" // string | <p>Results can be sorted with the following options: </p> <ul> <li>threat_name</li> <li>classification</li> <li>device_name</li> <li>process_name</li> <li>process_owner</li> <li>detection_date</li> <li>status</li> </ul> <p>Prepending a dash (-) to the parameter value will reverse the order of the returned results.</p> <p><code>?sort_by=-device_name</code> will order the response by device_name in descending order.</p> (optional)
	term := "Chrome" // string | <p>Search term to filter threat results.</p> <p>The response will include anything matching the following fields: <code>device_name</code>, <code>file_hash</code>, and <code>file_path</code>.</p> <p>So if you search for <code>bad file</code>, the results will include anywhere <code>bad file</code> exists in the three fields above.</p> (optional)
	limit := "1000" // string | <p>A hard upper <code>limit</code> is set at 1000 records returned per request. If more records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. </p> <p>Additionally, parameter queries can be added to a request to limit the results.</p> (optional)
	offset := "1" // string | Specify the starting record to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreatsAPI.GetThreatDetails(context.Background()).Classification(classification).DateRange(dateRange).DeviceId(deviceId).Status(status).SortBy(sortBy).Term(term).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatsAPI.GetThreatDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThreatDetails`: ThreatsGetThreatDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `ThreatsAPI.GetThreatDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetThreatDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **classification** | **string** | Return all records matching a specified classification. The following classification options are available: &lt;code&gt;malware&lt;/code&gt; and &lt;code&gt;pup&lt;/code&gt;. Leave this parameter empty to return all classification types. | 
 **dateRange** | **string** | Return all records within a specified number of days. Any positive number of days may be specified. Examples: &lt;code&gt;7&lt;/code&gt;, &lt;code&gt;30&lt;/code&gt;, &lt;code&gt;60&lt;/code&gt;, &lt;code&gt;90&lt;/code&gt;, &lt;code&gt;180&lt;/code&gt;, or &lt;code&gt;365&lt;/code&gt;. | 
 **deviceId** | **string** |  | 
 **status** | **string** | Return all records matching a specified status. The following status options are available: &lt;code&gt;quarantined&lt;/code&gt;, &lt;code&gt;not_quarantined&lt;/code&gt;, or &lt;code&gt;released&lt;/code&gt;. Leave this parameter empty to return all status types. | 
 **sortBy** | **string** | &lt;p&gt;Results can be sorted with the following options: &lt;/p&gt; &lt;ul&gt; &lt;li&gt;threat_name&lt;/li&gt; &lt;li&gt;classification&lt;/li&gt; &lt;li&gt;device_name&lt;/li&gt; &lt;li&gt;process_name&lt;/li&gt; &lt;li&gt;process_owner&lt;/li&gt; &lt;li&gt;detection_date&lt;/li&gt; &lt;li&gt;status&lt;/li&gt; &lt;/ul&gt; &lt;p&gt;Prepending a dash (-) to the parameter value will reverse the order of the returned results.&lt;/p&gt; &lt;p&gt;&lt;code&gt;?sort_by&#x3D;-device_name&lt;/code&gt; will order the response by device_name in descending order.&lt;/p&gt; | 
 **term** | **string** | &lt;p&gt;Search term to filter threat results.&lt;/p&gt; &lt;p&gt;The response will include anything matching the following fields: &lt;code&gt;device_name&lt;/code&gt;, &lt;code&gt;file_hash&lt;/code&gt;, and &lt;code&gt;file_path&lt;/code&gt;.&lt;/p&gt; &lt;p&gt;So if you search for &lt;code&gt;bad file&lt;/code&gt;, the results will include anywhere &lt;code&gt;bad file&lt;/code&gt; exists in the three fields above.&lt;/p&gt; | 
 **limit** | **string** | &lt;p&gt;A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 1000 records returned per request. If more records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. &lt;/p&gt; &lt;p&gt;Additionally, parameter queries can be added to a request to limit the results.&lt;/p&gt; | 
 **offset** | **string** | Specify the starting record to return | 

### Return type

[**ThreatsGetThreatDetails200Response**](ThreatsGetThreatDetails200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

