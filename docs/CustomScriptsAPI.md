# \CustomScriptsAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomScriptsCreateCustomScript**](CustomScriptsAPI.md#CustomScriptsCreateCustomScript) | **Post** /api/v1/library/custom-scripts | Create Custom Script
[**CustomScriptsDeleteCustomScript**](CustomScriptsAPI.md#CustomScriptsDeleteCustomScript) | **Delete** /api/v1/library/custom-scripts/{library_item_id} | Delete Custom Script
[**CustomScriptsGetCustomScript**](CustomScriptsAPI.md#CustomScriptsGetCustomScript) | **Get** /api/v1/library/custom-scripts/{library_item_id} | Get Custom Script
[**CustomScriptsListCustomScripts**](CustomScriptsAPI.md#CustomScriptsListCustomScripts) | **Get** /api/v1/library/custom-scripts | List Custom Scripts
[**CustomScriptsUpdateCustomScript**](CustomScriptsAPI.md#CustomScriptsUpdateCustomScript) | **Patch** /api/v1/library/custom-scripts/{library_item_id} | Update Custom Script



## CustomScriptsCreateCustomScript

> map[string]interface{} CustomScriptsCreateCustomScript(ctx).Body(body).Execute()

Create Custom Script



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
	resp, r, err := apiClient.CustomScriptsAPI.CustomScriptsCreateCustomScript(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomScriptsAPI.CustomScriptsCreateCustomScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomScriptsCreateCustomScript`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomScriptsAPI.CustomScriptsCreateCustomScript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomScriptsCreateCustomScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## CustomScriptsDeleteCustomScript

> CustomScriptsDeleteCustomScript(ctx, libraryItemId).Execute()

Delete Custom Script



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
	r, err := apiClient.CustomScriptsAPI.CustomScriptsDeleteCustomScript(context.Background(), libraryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomScriptsAPI.CustomScriptsDeleteCustomScript``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCustomScriptsDeleteCustomScriptRequest struct via the builder pattern


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


## CustomScriptsGetCustomScript

> map[string]interface{} CustomScriptsGetCustomScript(ctx, libraryItemId).Execute()

Get Custom Script



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
	resp, r, err := apiClient.CustomScriptsAPI.CustomScriptsGetCustomScript(context.Background(), libraryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomScriptsAPI.CustomScriptsGetCustomScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomScriptsGetCustomScript`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomScriptsAPI.CustomScriptsGetCustomScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomScriptsGetCustomScriptRequest struct via the builder pattern


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


## CustomScriptsListCustomScripts

> map[string]interface{} CustomScriptsListCustomScripts(ctx).Page(page).Execute()

List Custom Scripts



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
	page := "1" // string | Optional page number. Used when results exceed pagination threshold. A hard upper limit is set at 300 device records returned per request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomScriptsAPI.CustomScriptsListCustomScripts(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomScriptsAPI.CustomScriptsListCustomScripts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomScriptsListCustomScripts`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomScriptsAPI.CustomScriptsListCustomScripts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomScriptsListCustomScriptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Optional page number. Used when results exceed pagination threshold. A hard upper limit is set at 300 device records returned per request. | 

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


## CustomScriptsUpdateCustomScript

> map[string]interface{} CustomScriptsUpdateCustomScript(ctx, libraryItemId).Body(body).Execute()

Update Custom Script



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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomScriptsAPI.CustomScriptsUpdateCustomScript(context.Background(), libraryItemId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomScriptsAPI.CustomScriptsUpdateCustomScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomScriptsUpdateCustomScript`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomScriptsAPI.CustomScriptsUpdateCustomScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomScriptsUpdateCustomScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

