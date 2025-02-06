# \CustomAppsAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomAppsCreateCustomApp**](CustomAppsAPI.md#CustomAppsCreateCustomApp) | **Post** /api/v1/library/custom-apps | Create Custom App
[**CustomAppsDeleteCustomApp**](CustomAppsAPI.md#CustomAppsDeleteCustomApp) | **Delete** /api/v1/library/custom-apps/{library_item_id} | Delete Custom App
[**CustomAppsGetCustomApp**](CustomAppsAPI.md#CustomAppsGetCustomApp) | **Get** /api/v1/library/custom-apps/{library_item_id} | Get Custom App
[**CustomAppsListCustomApps**](CustomAppsAPI.md#CustomAppsListCustomApps) | **Get** /api/v1/library/custom-apps | List Custom Apps
[**CustomAppsUpdateCustomApp**](CustomAppsAPI.md#CustomAppsUpdateCustomApp) | **Patch** /api/v1/library/custom-apps/{library_item_id} | Update Custom App
[**CustomAppsUploadCustomApp**](CustomAppsAPI.md#CustomAppsUploadCustomApp) | **Post** /api/v1/library/custom-apps/upload | Upload Custom App



## CustomAppsCreateCustomApp

> map[string]interface{} CustomAppsCreateCustomApp(ctx).Name(name).FileKey(fileKey).InstallType(installType).InstallEnforcement(installEnforcement).ShowInSelfService(showInSelfService).SelfServiceCategoryId(selfServiceCategoryId).SelfServiceRecommended(selfServiceRecommended).Execute()

Create Custom App



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
	name := "name_example" // string | (Required) The name for this Custom App
	fileKey := "fileKey_example" // string | (Required) The S3 key from the <code>Upload Custom App</code> endpont used to upload the custom app file.
	installType := "installType_example" // string | (Required) Options are package, zip, image
	installEnforcement := "installEnforcement_example" // string | (Required) Options are install_once, continuously_enforce, no_enforcement
	showInSelfService := "showInSelfService_example" // string | (Optional, default=false) Displays this app in Self Service
	selfServiceCategoryId := "selfServiceCategoryId_example" // string | (Required for show_in_self_service=true) Self Service Category (by ID) to display app in
	selfServiceRecommended := "selfServiceRecommended_example" // string | (Optional, default=false) Adds recommended flag to app in Self Service

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomAppsAPI.CustomAppsCreateCustomApp(context.Background()).Name(name).FileKey(fileKey).InstallType(installType).InstallEnforcement(installEnforcement).ShowInSelfService(showInSelfService).SelfServiceCategoryId(selfServiceCategoryId).SelfServiceRecommended(selfServiceRecommended).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAppsAPI.CustomAppsCreateCustomApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomAppsCreateCustomApp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomAppsAPI.CustomAppsCreateCustomApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomAppsCreateCustomAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | (Required) The name for this Custom App | 
 **fileKey** | **string** | (Required) The S3 key from the &lt;code&gt;Upload Custom App&lt;/code&gt; endpont used to upload the custom app file. | 
 **installType** | **string** | (Required) Options are package, zip, image | 
 **installEnforcement** | **string** | (Required) Options are install_once, continuously_enforce, no_enforcement | 
 **showInSelfService** | **string** | (Optional, default&#x3D;false) Displays this app in Self Service | 
 **selfServiceCategoryId** | **string** | (Required for show_in_self_service&#x3D;true) Self Service Category (by ID) to display app in | 
 **selfServiceRecommended** | **string** | (Optional, default&#x3D;false) Adds recommended flag to app in Self Service | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomAppsDeleteCustomApp

> CustomAppsDeleteCustomApp(ctx, libraryItemId).Execute()

Delete Custom App



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
	libraryItemId := "libraryItemId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomAppsAPI.CustomAppsDeleteCustomApp(context.Background(), libraryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAppsAPI.CustomAppsDeleteCustomApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomAppsDeleteCustomAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomAppsGetCustomApp

> map[string]interface{} CustomAppsGetCustomApp(ctx, libraryItemId).Execute()

Get Custom App



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
	libraryItemId := "libraryItemId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomAppsAPI.CustomAppsGetCustomApp(context.Background(), libraryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAppsAPI.CustomAppsGetCustomApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomAppsGetCustomApp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomAppsAPI.CustomAppsGetCustomApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomAppsGetCustomAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomAppsListCustomApps

> map[string]interface{} CustomAppsListCustomApps(ctx).Page(page).Execute()

List Custom Apps



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
	page := "1" // string | Optional page number. Used when results exceed pagination threshold. A hard upper <code>limit</code> is set at 300 device records returned per request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomAppsAPI.CustomAppsListCustomApps(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAppsAPI.CustomAppsListCustomApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomAppsListCustomApps`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomAppsAPI.CustomAppsListCustomApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomAppsListCustomAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Optional page number. Used when results exceed pagination threshold. A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomAppsUpdateCustomApp

> map[string]interface{} CustomAppsUpdateCustomApp(ctx, libraryItemId).Name(name).Active(active).Execute()

Update Custom App



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
	libraryItemId := "libraryItemId_example" // string | 
	name := "name_example" // string | Renaming a Custom App
	active := "active_example" // string | (Optional, default=true) Whether this Custom App is active and installable

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomAppsAPI.CustomAppsUpdateCustomApp(context.Background(), libraryItemId).Name(name).Active(active).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAppsAPI.CustomAppsUpdateCustomApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomAppsUpdateCustomApp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomAppsAPI.CustomAppsUpdateCustomApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomAppsUpdateCustomAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Renaming a Custom App | 
 **active** | **string** | (Optional, default&#x3D;true) Whether this Custom App is active and installable | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomAppsUploadCustomApp

> CustomAppsUploadCustomApp(ctx).Body(body).Execute()

Upload Custom App



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
	r, err := apiClient.CustomAppsAPI.CustomAppsUploadCustomApp(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAppsAPI.CustomAppsUploadCustomApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomAppsUploadCustomAppRequest struct via the builder pattern


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

