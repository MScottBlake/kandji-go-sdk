# \CustomProfilesAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomProfilesCreateCustomProfile**](CustomProfilesAPI.md#CustomProfilesCreateCustomProfile) | **Post** /api/v1/library/custom-profiles | Create Custom Profile
[**CustomProfilesDeleteCustomProfile**](CustomProfilesAPI.md#CustomProfilesDeleteCustomProfile) | **Delete** /api/v1/library/custom-profiles/{library_item_id} | Delete Custom Profile
[**CustomProfilesGetCustomProfile**](CustomProfilesAPI.md#CustomProfilesGetCustomProfile) | **Get** /api/v1/library/custom-profiles/{library_item_id} | Get Custom Profile
[**CustomProfilesListCustomProfiles**](CustomProfilesAPI.md#CustomProfilesListCustomProfiles) | **Get** /api/v1/library/custom-profiles | List Custom Profiles
[**CustomProfilesUpdateCustomProfile**](CustomProfilesAPI.md#CustomProfilesUpdateCustomProfile) | **Patch** /api/v1/library/custom-profiles/{library_item_id} | Update Custom Profile



## CustomProfilesCreateCustomProfile

> map[string]interface{} CustomProfilesCreateCustomProfile(ctx).Name(name).File(file).Active(active).Execute()

Create Custom Profile



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
	name := "name_example" // string | (Required) The profile name
	file := os.NewFile(1234, "some_file") // *os.File | (Required) The path to the profile's .mobileconfig file
	active := "active_example" // string | (Optional, default=true) Whether this library item is active

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomProfilesAPI.CustomProfilesCreateCustomProfile(context.Background()).Name(name).File(file).Active(active).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomProfilesAPI.CustomProfilesCreateCustomProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomProfilesCreateCustomProfile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomProfilesAPI.CustomProfilesCreateCustomProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomProfilesCreateCustomProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | (Required) The profile name | 
 **file** | ***os.File** | (Required) The path to the profile&#39;s .mobileconfig file | 
 **active** | **string** | (Optional, default&#x3D;true) Whether this library item is active | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomProfilesDeleteCustomProfile

> CustomProfilesDeleteCustomProfile(ctx, libraryItemId).Execute()

Delete Custom Profile



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
	r, err := apiClient.CustomProfilesAPI.CustomProfilesDeleteCustomProfile(context.Background(), libraryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomProfilesAPI.CustomProfilesDeleteCustomProfile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCustomProfilesDeleteCustomProfileRequest struct via the builder pattern


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


## CustomProfilesGetCustomProfile

> map[string]interface{} CustomProfilesGetCustomProfile(ctx, libraryItemId).Execute()

Get Custom Profile



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
	resp, r, err := apiClient.CustomProfilesAPI.CustomProfilesGetCustomProfile(context.Background(), libraryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomProfilesAPI.CustomProfilesGetCustomProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomProfilesGetCustomProfile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomProfilesAPI.CustomProfilesGetCustomProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomProfilesGetCustomProfileRequest struct via the builder pattern


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


## CustomProfilesListCustomProfiles

> map[string]interface{} CustomProfilesListCustomProfiles(ctx).Page(page).Execute()

List Custom Profiles



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
	resp, r, err := apiClient.CustomProfilesAPI.CustomProfilesListCustomProfiles(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomProfilesAPI.CustomProfilesListCustomProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomProfilesListCustomProfiles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomProfilesAPI.CustomProfilesListCustomProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomProfilesListCustomProfilesRequest struct via the builder pattern


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


## CustomProfilesUpdateCustomProfile

> map[string]interface{} CustomProfilesUpdateCustomProfile(ctx, libraryItemId).Execute()

Update Custom Profile



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
	resp, r, err := apiClient.CustomProfilesAPI.CustomProfilesUpdateCustomProfile(context.Background(), libraryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomProfilesAPI.CustomProfilesUpdateCustomProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomProfilesUpdateCustomProfile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomProfilesAPI.CustomProfilesUpdateCustomProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomProfilesUpdateCustomProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

