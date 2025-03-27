# \AuditLogAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAuditEvents**](AuditLogAPI.md#ListAuditEvents) | **Get** /api/v1/audit/events | List audit events



## ListAuditEvents

> AuditLogListAuditEvents200Response ListAuditEvents(ctx).Limit(limit).SortBy(sortBy).StartDate(startDate).EndDate(endDate).Cursor(cursor).Execute()

List audit events



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
	limit := "500" // string | A max upper <code>limit</code> is set at 500 records returned per request. Pagination should be used using the cursor in the <code>next</code> key to request more results. Additionally, parameter queries can be added to a request to filter the results.
	sortBy := "-occurred_at" // string | Sort results by <code>occurred_at</code>, <code>id</code> either ascending (default behavior) or descending(-) order.
	startDate := "2024-11-26T22:58:26.239570Z" // string | Filter by start date in datetime or year-month-day (2024-11-26) formats (optional)
	endDate := "2024-12-06T17:48:41.784439Z" // string | Filter by start date in datetime or year-month-day (2024-12-06) formats (optional)
	cursor := "cursor_example" // string | You can pass the next cursor as a parameter or use the URL in the next key to get the next page of results or to start from where you left off last. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogAPI.ListAuditEvents(context.Background()).Limit(limit).SortBy(sortBy).StartDate(startDate).EndDate(endDate).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogAPI.ListAuditEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAuditEvents`: AuditLogListAuditEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditLogAPI.ListAuditEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAuditEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | A max upper &lt;code&gt;limit&lt;/code&gt; is set at 500 records returned per request. Pagination should be used using the cursor in the &lt;code&gt;next&lt;/code&gt; key to request more results. Additionally, parameter queries can be added to a request to filter the results. | 
 **sortBy** | **string** | Sort results by &lt;code&gt;occurred_at&lt;/code&gt;, &lt;code&gt;id&lt;/code&gt; either ascending (default behavior) or descending(-) order. | 
 **startDate** | **string** | Filter by start date in datetime or year-month-day (2024-11-26) formats | 
 **endDate** | **string** | Filter by start date in datetime or year-month-day (2024-12-06) formats | 
 **cursor** | **string** | You can pass the next cursor as a parameter or use the URL in the next key to get the next page of results or to start from where you left off last. | 

### Return type

[**AuditLogListAuditEvents200Response**](AuditLogListAuditEvents200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

