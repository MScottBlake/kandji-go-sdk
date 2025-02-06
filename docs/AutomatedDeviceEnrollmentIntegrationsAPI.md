# \AutomatedDeviceEnrollmentIntegrationsAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration**](AutomatedDeviceEnrollmentIntegrationsAPI.md#AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration) | **Post** /api/v1/integrations/apple/ade/ | Create ADE integration
[**AutomatedDeviceEnrollmentIntegrationsDeleteAdeIntegration**](AutomatedDeviceEnrollmentIntegrationsAPI.md#AutomatedDeviceEnrollmentIntegrationsDeleteAdeIntegration) | **Delete** /api/v1/integrations/apple/ade/{ade_token_id} | Delete ADE integration
[**AutomatedDeviceEnrollmentIntegrationsDownloadAdePublicKey**](AutomatedDeviceEnrollmentIntegrationsAPI.md#AutomatedDeviceEnrollmentIntegrationsDownloadAdePublicKey) | **Get** /api/v1/integrations/apple/ade/public_key/ | Download ADE public key
[**AutomatedDeviceEnrollmentIntegrationsGetAdeDevice**](AutomatedDeviceEnrollmentIntegrationsAPI.md#AutomatedDeviceEnrollmentIntegrationsGetAdeDevice) | **Get** /api/v1/integrations/apple/ade/devices/{device_id} | Get ADE device
[**AutomatedDeviceEnrollmentIntegrationsGetAdeIntegration**](AutomatedDeviceEnrollmentIntegrationsAPI.md#AutomatedDeviceEnrollmentIntegrationsGetAdeIntegration) | **Get** /api/v1/integrations/apple/ade/{ade_token_id} | Get ADE integration
[**AutomatedDeviceEnrollmentIntegrationsListAdeDevices**](AutomatedDeviceEnrollmentIntegrationsAPI.md#AutomatedDeviceEnrollmentIntegrationsListAdeDevices) | **Get** /api/v1/integrations/apple/ade/devices | List ADE devices
[**AutomatedDeviceEnrollmentIntegrationsListAdeIntegrations**](AutomatedDeviceEnrollmentIntegrationsAPI.md#AutomatedDeviceEnrollmentIntegrationsListAdeIntegrations) | **Get** /api/v1/integrations/apple/ade | List ADE integrations
[**AutomatedDeviceEnrollmentIntegrationsListDevicesAssociatedToAdeToken**](AutomatedDeviceEnrollmentIntegrationsAPI.md#AutomatedDeviceEnrollmentIntegrationsListDevicesAssociatedToAdeToken) | **Get** /api/v1/integrations/apple/ade/{ade_token_id}/devices | List devices associated to ADE token
[**AutomatedDeviceEnrollmentIntegrationsRenewAdeIntegration**](AutomatedDeviceEnrollmentIntegrationsAPI.md#AutomatedDeviceEnrollmentIntegrationsRenewAdeIntegration) | **Post** /api/v1/integrations/apple/ade/{ade_token_id}/renew | Renew ADE integration
[**AutomatedDeviceEnrollmentIntegrationsUpdateAdeDevice**](AutomatedDeviceEnrollmentIntegrationsAPI.md#AutomatedDeviceEnrollmentIntegrationsUpdateAdeDevice) | **Patch** /api/v1/integrations/apple/ade/devices/{device_id} | Update ADE device
[**AutomatedDeviceEnrollmentIntegrationsUpdateAdeIntegration**](AutomatedDeviceEnrollmentIntegrationsAPI.md#AutomatedDeviceEnrollmentIntegrationsUpdateAdeIntegration) | **Patch** /api/v1/integrations/apple/ade/{ade_token_id} | Update ADE integration



## AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration

> map[string]interface{} AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration(ctx).BlueprintId(blueprintId).Phone(phone).Email(email).File(file).Execute()

Create ADE integration



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
	blueprintId := "blueprintId_example" // string | 
	phone := "phone_example" // string | 
	email := "email_example" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | This is the MDM server token file(.p7m) download from ABM. Once downloaded from ABM, the file can be uploaded via API.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration(context.Background()).BlueprintId(blueprintId).Phone(phone).Email(email).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsCreateAdeIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutomatedDeviceEnrollmentIntegrationsCreateAdeIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintId** | **string** |  | 
 **phone** | **string** |  | 
 **email** | **string** |  | 
 **file** | ***os.File** | This is the MDM server token file(.p7m) download from ABM. Once downloaded from ABM, the file can be uploaded via API. | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomatedDeviceEnrollmentIntegrationsDeleteAdeIntegration

> AutomatedDeviceEnrollmentIntegrationsDeleteAdeIntegration(ctx, adeTokenId).Execute()

Delete ADE integration



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
	adeTokenId := "adeTokenId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsDeleteAdeIntegration(context.Background(), adeTokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsDeleteAdeIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adeTokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomatedDeviceEnrollmentIntegrationsDeleteAdeIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomatedDeviceEnrollmentIntegrationsDownloadAdePublicKey

> string AutomatedDeviceEnrollmentIntegrationsDownloadAdePublicKey(ctx).Execute()

Download ADE public key



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsDownloadAdePublicKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsDownloadAdePublicKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomatedDeviceEnrollmentIntegrationsDownloadAdePublicKey`: string
	fmt.Fprintf(os.Stdout, "Response from `AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsDownloadAdePublicKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAutomatedDeviceEnrollmentIntegrationsDownloadAdePublicKeyRequest struct via the builder pattern


### Return type

**string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-x509-ca-cert

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomatedDeviceEnrollmentIntegrationsGetAdeDevice

> map[string]interface{} AutomatedDeviceEnrollmentIntegrationsGetAdeDevice(ctx, deviceId).Execute()

Get ADE device



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
	deviceId := "deviceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsGetAdeDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsGetAdeDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomatedDeviceEnrollmentIntegrationsGetAdeDevice`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsGetAdeDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomatedDeviceEnrollmentIntegrationsGetAdeDeviceRequest struct via the builder pattern


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


## AutomatedDeviceEnrollmentIntegrationsGetAdeIntegration

> AutomatedDeviceEnrollmentIntegrationsGetAdeIntegration(ctx, adeTokenId).Execute()

Get ADE integration



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
	adeTokenId := "adeTokenId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsGetAdeIntegration(context.Background(), adeTokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsGetAdeIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adeTokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomatedDeviceEnrollmentIntegrationsGetAdeIntegrationRequest struct via the builder pattern


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


## AutomatedDeviceEnrollmentIntegrationsListAdeDevices

> map[string]interface{} AutomatedDeviceEnrollmentIntegrationsListAdeDevices(ctx).BlueprintId(blueprintId).UserId(userId).DepAccount(depAccount).DeviceFamily(deviceFamily).Model(model).Os(os).ProfileStatus(profileStatus).SerialNumber(serialNumber).Page(page).Execute()

List ADE devices



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
	blueprintId := "fce0cc58-caa5-40d2-a0d7-a0b257127ec5" // string | Return results &quot;containing&quot; the specified blueprint id (optional)
	userId := "8136" // string | &quot;exact&quot; match on kandji user ID number (optional)
	depAccount := "depAccount_example" // string | The ADE token UUID (optional)
	deviceFamily := "deviceFamily_example" // string | Mac, iPhone, iPad, AppleTV, iPod (optional)
	model := "MacBook Air" // string | Return model results &quot;containing&quot; the specified model string. - &quot;iPad (8th Generation)&quot;, &quot;MacBook Air&quot; (optional)
	os := "os_example" // string | OSX, iOS, tvOS (optional)
	profileStatus := "profileStatus_example" // string | The automated device enrollment profile assignment status - assigned, empty, pushed, removed (optional)
	serialNumber := "serialNumber_example" // string | Search for a specific device by Serial Number. If partial serial number is provided in the query, all device containing the partial string will be returned. (optional)
	page := "1" // string | Use the <code>page</code> parameter to page through results or to request a specific page. By default, if a page is not specified, page 1 is returned. Note: 300 device records are returned per page of results. Alternatively, the <code>next</code> and <code>previous</code> key attributes in the response can be used to request the next page of results or return to the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsListAdeDevices(context.Background()).BlueprintId(blueprintId).UserId(userId).DepAccount(depAccount).DeviceFamily(deviceFamily).Model(model).Os(os).ProfileStatus(profileStatus).SerialNumber(serialNumber).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsListAdeDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomatedDeviceEnrollmentIntegrationsListAdeDevices`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsListAdeDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutomatedDeviceEnrollmentIntegrationsListAdeDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintId** | **string** | Return results &amp;quot;containing&amp;quot; the specified blueprint id | 
 **userId** | **string** | &amp;quot;exact&amp;quot; match on kandji user ID number | 
 **depAccount** | **string** | The ADE token UUID | 
 **deviceFamily** | **string** | Mac, iPhone, iPad, AppleTV, iPod | 
 **model** | **string** | Return model results &amp;quot;containing&amp;quot; the specified model string. - &amp;quot;iPad (8th Generation)&amp;quot;, &amp;quot;MacBook Air&amp;quot; | 
 **os** | **string** | OSX, iOS, tvOS | 
 **profileStatus** | **string** | The automated device enrollment profile assignment status - assigned, empty, pushed, removed | 
 **serialNumber** | **string** | Search for a specific device by Serial Number. If partial serial number is provided in the query, all device containing the partial string will be returned. | 
 **page** | **string** | Use the &lt;code&gt;page&lt;/code&gt; parameter to page through results or to request a specific page. By default, if a page is not specified, page 1 is returned. Note: 300 device records are returned per page of results. Alternatively, the &lt;code&gt;next&lt;/code&gt; and &lt;code&gt;previous&lt;/code&gt; key attributes in the response can be used to request the next page of results or return to the previous page. | 

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


## AutomatedDeviceEnrollmentIntegrationsListAdeIntegrations

> AutomatedDeviceEnrollmentIntegrationsListAdeIntegrations(ctx).Execute()

List ADE integrations



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsListAdeIntegrations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsListAdeIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAutomatedDeviceEnrollmentIntegrationsListAdeIntegrationsRequest struct via the builder pattern


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


## AutomatedDeviceEnrollmentIntegrationsListDevicesAssociatedToAdeToken

> map[string]interface{} AutomatedDeviceEnrollmentIntegrationsListDevicesAssociatedToAdeToken(ctx, adeTokenId).Page(page).Execute()

List devices associated to ADE token



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
	adeTokenId := "adeTokenId_example" // string | 
	page := "1" // string | Use the <code>page</code> parameter to page through results or to request a specific page. By default, if a page is not specified, page 1 is returned. Note: 300 device records are returned per page of results. Alternatively, the <code>next</code> and <code>previous</code> key attributes in the response can be used to request the next page of results or return to the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsListDevicesAssociatedToAdeToken(context.Background(), adeTokenId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsListDevicesAssociatedToAdeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomatedDeviceEnrollmentIntegrationsListDevicesAssociatedToAdeToken`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsListDevicesAssociatedToAdeToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adeTokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomatedDeviceEnrollmentIntegrationsListDevicesAssociatedToAdeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Use the &lt;code&gt;page&lt;/code&gt; parameter to page through results or to request a specific page. By default, if a page is not specified, page 1 is returned. Note: 300 device records are returned per page of results. Alternatively, the &lt;code&gt;next&lt;/code&gt; and &lt;code&gt;previous&lt;/code&gt; key attributes in the response can be used to request the next page of results or return to the previous page. | 

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


## AutomatedDeviceEnrollmentIntegrationsRenewAdeIntegration

> AutomatedDeviceEnrollmentIntegrationsRenewAdeIntegration(ctx, adeTokenId).BlueprintId(blueprintId).Phone(phone).Email(email).File(file).Execute()

Renew ADE integration



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
	adeTokenId := "adeTokenId_example" // string | 
	blueprintId := "blueprintId_example" // string | 
	phone := "phone_example" // string | 
	email := "email_example" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | This is the MDM server token file(.p7m) download from ABM. Once downloaded from ABM, the file can be uploaded via API.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsRenewAdeIntegration(context.Background(), adeTokenId).BlueprintId(blueprintId).Phone(phone).Email(email).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsRenewAdeIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adeTokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomatedDeviceEnrollmentIntegrationsRenewAdeIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blueprintId** | **string** |  | 
 **phone** | **string** |  | 
 **email** | **string** |  | 
 **file** | ***os.File** | This is the MDM server token file(.p7m) download from ABM. Once downloaded from ABM, the file can be uploaded via API. | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomatedDeviceEnrollmentIntegrationsUpdateAdeDevice

> map[string]interface{} AutomatedDeviceEnrollmentIntegrationsUpdateAdeDevice(ctx, deviceId).Body(body).Execute()

Update ADE device



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
	deviceId := "deviceId_example" // string | 
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsUpdateAdeDevice(context.Background(), deviceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsUpdateAdeDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomatedDeviceEnrollmentIntegrationsUpdateAdeDevice`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsUpdateAdeDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomatedDeviceEnrollmentIntegrationsUpdateAdeDeviceRequest struct via the builder pattern


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


## AutomatedDeviceEnrollmentIntegrationsUpdateAdeIntegration

> AutomatedDeviceEnrollmentIntegrationsUpdateAdeIntegration(ctx, adeTokenId).Body(body).Execute()

Update ADE integration



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
	adeTokenId := "adeTokenId_example" // string | 
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsUpdateAdeIntegration(context.Background(), adeTokenId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.AutomatedDeviceEnrollmentIntegrationsUpdateAdeIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adeTokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomatedDeviceEnrollmentIntegrationsUpdateAdeIntegrationRequest struct via the builder pattern


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

