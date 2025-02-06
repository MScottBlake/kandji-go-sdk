# \InHouseAppsAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InhouseAppsCreateInhouseApp**](InHouseAppsAPI.md#InhouseAppsCreateInhouseApp) | **Post** /api/v1/library/ipa-apps | Create In-House App
[**InhouseAppsDeleteInhouseApp**](InHouseAppsAPI.md#InhouseAppsDeleteInhouseApp) | **Delete** /api/v1/library/ipa-apps/{library_item_id} | Delete In-House App
[**InhouseAppsGetInhouseApp**](InHouseAppsAPI.md#InhouseAppsGetInhouseApp) | **Get** /api/v1/library/ipa-apps/{library_item_id} | Get In-House App
[**InhouseAppsListInhouseApps**](InHouseAppsAPI.md#InhouseAppsListInhouseApps) | **Get** /api/v1/library/ipa-apps | List In-House Apps
[**InhouseAppsUpdateInhouseApp**](InHouseAppsAPI.md#InhouseAppsUpdateInhouseApp) | **Patch** /api/v1/library/ipa-apps/{library_item_id} | Update In-House App
[**InhouseAppsUploadInhouseApp**](InHouseAppsAPI.md#InhouseAppsUploadInhouseApp) | **Post** /api/v1/library/ipa-apps/upload | Upload In-House App
[**InhouseAppsUploadInhouseAppStatus**](InHouseAppsAPI.md#InhouseAppsUploadInhouseAppStatus) | **Get** /api/v1/library/ipa-apps/upload/{pending_upload_id}/status | Upload In-House App Status



## InhouseAppsCreateInhouseApp

> InhouseAppsCreateInhouseApp(ctx).ContentType(contentType).Body(body).Execute()

Create In-House App



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
	contentType := "application/json" // string | 
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InHouseAppsAPI.InhouseAppsCreateInhouseApp(context.Background()).ContentType(contentType).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InHouseAppsAPI.InhouseAppsCreateInhouseApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInhouseAppsCreateInhouseAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
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


## InhouseAppsDeleteInhouseApp

> InhouseAppsDeleteInhouseApp(ctx, libraryItemId).Execute()

Delete In-House App



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
	r, err := apiClient.InHouseAppsAPI.InhouseAppsDeleteInhouseApp(context.Background(), libraryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InHouseAppsAPI.InhouseAppsDeleteInhouseApp``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInhouseAppsDeleteInhouseAppRequest struct via the builder pattern


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


## InhouseAppsGetInhouseApp

> map[string]interface{} InhouseAppsGetInhouseApp(ctx, libraryItemId).Execute()

Get In-House App



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
	resp, r, err := apiClient.InHouseAppsAPI.InhouseAppsGetInhouseApp(context.Background(), libraryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InHouseAppsAPI.InhouseAppsGetInhouseApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InhouseAppsGetInhouseApp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `InHouseAppsAPI.InhouseAppsGetInhouseApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInhouseAppsGetInhouseAppRequest struct via the builder pattern


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


## InhouseAppsListInhouseApps

> map[string]interface{} InhouseAppsListInhouseApps(ctx).Page(page).Execute()

List In-House Apps



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
	page := "1" // string | Optional page number. Used when results exceed pagination threshold. A hard upper <code>limit</code> is set at 300 app records returned per request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InHouseAppsAPI.InhouseAppsListInhouseApps(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InHouseAppsAPI.InhouseAppsListInhouseApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InhouseAppsListInhouseApps`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `InHouseAppsAPI.InhouseAppsListInhouseApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInhouseAppsListInhouseAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Optional page number. Used when results exceed pagination threshold. A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 app records returned per request. | 

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


## InhouseAppsUpdateInhouseApp

> InhouseAppsUpdateInhouseApp(ctx, libraryItemId).ContentType(contentType).Body(body).Execute()

Update In-House App



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
	contentType := "application/json" // string | 
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InHouseAppsAPI.InhouseAppsUpdateInhouseApp(context.Background(), libraryItemId).ContentType(contentType).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InHouseAppsAPI.InhouseAppsUpdateInhouseApp``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInhouseAppsUpdateInhouseAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InhouseAppsUploadInhouseApp

> map[string]interface{} InhouseAppsUploadInhouseApp(ctx).ContentType(contentType).Body(body).Execute()

Upload In-House App



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
	contentType := "application/json" // string | 
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InHouseAppsAPI.InhouseAppsUploadInhouseApp(context.Background()).ContentType(contentType).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InHouseAppsAPI.InhouseAppsUploadInhouseApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InhouseAppsUploadInhouseApp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `InHouseAppsAPI.InhouseAppsUploadInhouseApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInhouseAppsUploadInhouseAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
 **body** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InhouseAppsUploadInhouseAppStatus

> map[string]interface{} InhouseAppsUploadInhouseAppStatus(ctx, pendingUploadId).Execute()

Upload In-House App Status



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
	pendingUploadId := "pendingUploadId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InHouseAppsAPI.InhouseAppsUploadInhouseAppStatus(context.Background(), pendingUploadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InHouseAppsAPI.InhouseAppsUploadInhouseAppStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InhouseAppsUploadInhouseAppStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `InHouseAppsAPI.InhouseAppsUploadInhouseAppStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pendingUploadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInhouseAppsUploadInhouseAppStatusRequest struct via the builder pattern


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

