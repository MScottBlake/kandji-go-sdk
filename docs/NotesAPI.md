# \NotesAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotesCreateDeviceNote**](NotesAPI.md#NotesCreateDeviceNote) | **Post** /api/v1/devices/{device_id}/notes | Create Device Note
[**NotesDeleteDeviceNote**](NotesAPI.md#NotesDeleteDeviceNote) | **Delete** /api/v1/devices/{device_id}/notes/{note_id} | Delete Device Note
[**NotesGetDeviceNotes**](NotesAPI.md#NotesGetDeviceNotes) | **Get** /api/v1/devices/{device_id}/notes | Get Device Notes
[**NotesRetrieveDeviceNote**](NotesAPI.md#NotesRetrieveDeviceNote) | **Get** /api/v1/devices/{device_id}/notes/{note_id} | Retrieve Device Note
[**NotesUpdateDeviceNote**](NotesAPI.md#NotesUpdateDeviceNote) | **Patch** /api/v1/devices/{device_id}/notes/{note_id} | Update Device Note



## NotesCreateDeviceNote

> map[string]interface{} NotesCreateDeviceNote(ctx, deviceId).Body(body).Execute()

Create Device Note



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
	resp, r, err := apiClient.NotesAPI.NotesCreateDeviceNote(context.Background(), deviceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.NotesCreateDeviceNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotesCreateDeviceNote`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.NotesCreateDeviceNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotesCreateDeviceNoteRequest struct via the builder pattern


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


## NotesDeleteDeviceNote

> map[string]interface{} NotesDeleteDeviceNote(ctx, deviceId, noteId).Execute()

Delete Device Note



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
	noteId := "noteId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.NotesDeleteDeviceNote(context.Background(), deviceId, noteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.NotesDeleteDeviceNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotesDeleteDeviceNote`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.NotesDeleteDeviceNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 
**noteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotesDeleteDeviceNoteRequest struct via the builder pattern


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


## NotesGetDeviceNotes

> map[string]interface{} NotesGetDeviceNotes(ctx, deviceId).Execute()

Get Device Notes



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
	resp, r, err := apiClient.NotesAPI.NotesGetDeviceNotes(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.NotesGetDeviceNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotesGetDeviceNotes`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.NotesGetDeviceNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotesGetDeviceNotesRequest struct via the builder pattern


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


## NotesRetrieveDeviceNote

> map[string]interface{} NotesRetrieveDeviceNote(ctx, deviceId, noteId).Execute()

Retrieve Device Note



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
	noteId := "noteId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.NotesRetrieveDeviceNote(context.Background(), deviceId, noteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.NotesRetrieveDeviceNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotesRetrieveDeviceNote`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.NotesRetrieveDeviceNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 
**noteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotesRetrieveDeviceNoteRequest struct via the builder pattern


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


## NotesUpdateDeviceNote

> map[string]interface{} NotesUpdateDeviceNote(ctx, deviceId, noteId).Body(body).Execute()

Update Device Note



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
	noteId := "noteId_example" // string | 
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.NotesUpdateDeviceNote(context.Background(), deviceId, noteId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.NotesUpdateDeviceNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotesUpdateDeviceNote`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.NotesUpdateDeviceNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 
**noteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotesUpdateDeviceNoteRequest struct via the builder pattern


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

