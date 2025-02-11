# \UsersAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUser**](UsersAPI.md#DeleteUser) | **Delete** /api/v1/users/{user_id} | Delete User
[**GetUser**](UsersAPI.md#GetUser) | **Get** /api/v1/users/{user_id} | Get User
[**ListUsers**](UsersAPI.md#ListUsers) | **Get** /api/v1/users | List Users



## DeleteUser

> DeleteUser(ctx, userId).Execute()

Delete User



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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.DeleteUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> Success200Response16 GetUser(ctx, userId).Execute()

Get User



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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: Success200Response16
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Success200Response16**](Success200Response16.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> Success200Response15 ListUsers(ctx).Email(email).Id(id).IntegrationId(integrationId).Archived(archived).Execute()

List Users



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
	email := "test_user_01@accuhive.io" // string | Returns users with email addresses containing the provided string. (optional)
	id := "69c009ca-1f78-4bdf-bb93-08d6d39041db" // string | Search for a user matching the provided UUID value. (optional)
	integrationId := "f7461096-4ef9-43aa-88e9-ca1967ba0b38" // string | Search for a integration matching the provided UUID value. (optional)
	archived := "false" // string | Return only users that are either archived (true) or not archived (false). Archived users are users that appear in the Kandji Users module under the Archived tab. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.ListUsers(context.Background()).Email(email).Id(id).IntegrationId(integrationId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsers`: Success200Response15
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Returns users with email addresses containing the provided string. | 
 **id** | **string** | Search for a user matching the provided UUID value. | 
 **integrationId** | **string** | Search for a integration matching the provided UUID value. | 
 **archived** | **string** | Return only users that are either archived (true) or not archived (false). Archived users are users that appear in the Kandji Users module under the Archived tab. | 

### Return type

[**Success200Response15**](Success200Response15.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

