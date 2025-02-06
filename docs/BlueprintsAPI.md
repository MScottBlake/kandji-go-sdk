# \BlueprintsAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlueprintsAssignLibraryItem**](BlueprintsAPI.md#BlueprintsAssignLibraryItem) | **Post** /api/v1/blueprints/{blueprint_id}/assign-library-item | Assign Library Item
[**BlueprintsCreateBlueprint**](BlueprintsAPI.md#BlueprintsCreateBlueprint) | **Post** /api/v1/blueprints | Create Blueprint
[**BlueprintsDeleteBlueprint**](BlueprintsAPI.md#BlueprintsDeleteBlueprint) | **Delete** /api/v1/blueprints/{blueprint_id} | Delete Blueprint
[**BlueprintsGetBlueprint**](BlueprintsAPI.md#BlueprintsGetBlueprint) | **Get** /api/v1/blueprints/{blueprint_id} | Get Blueprint
[**BlueprintsGetBlueprintTemplates**](BlueprintsAPI.md#BlueprintsGetBlueprintTemplates) | **Get** /api/v1/blueprints/templates/ | Get Blueprint Templates
[**BlueprintsGetManualEnrollmentProfile**](BlueprintsAPI.md#BlueprintsGetManualEnrollmentProfile) | **Get** /api/v1/blueprints/{blueprint_id}/ota-enrollment-profile | Get Manual Enrollment Profile
[**BlueprintsListBlueprints**](BlueprintsAPI.md#BlueprintsListBlueprints) | **Get** /api/v1/blueprints | List Blueprints
[**BlueprintsListLibraryItems**](BlueprintsAPI.md#BlueprintsListLibraryItems) | **Get** /api/v1/blueprints/{blueprint_id}/list-library-items | List Library Items
[**BlueprintsRemoveLibraryItem**](BlueprintsAPI.md#BlueprintsRemoveLibraryItem) | **Post** /api/v1/blueprints/{blueprint_id}/remove-library-item | Remove Library Item
[**BlueprintsUpdateBlueprint**](BlueprintsAPI.md#BlueprintsUpdateBlueprint) | **Patch** /api/v1/blueprints/{blueprint_id} | Update Blueprint



## BlueprintsAssignLibraryItem

> map[string]interface{} BlueprintsAssignLibraryItem(ctx, blueprintId).Body(body).Execute()

Assign Library Item



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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintsAPI.BlueprintsAssignLibraryItem(context.Background(), blueprintId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.BlueprintsAssignLibraryItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlueprintsAssignLibraryItem`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BlueprintsAPI.BlueprintsAssignLibraryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blueprintId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlueprintsAssignLibraryItemRequest struct via the builder pattern


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


## BlueprintsCreateBlueprint

> map[string]interface{} BlueprintsCreateBlueprint(ctx).Name(name).EnrollmentCodeIsActive(enrollmentCodeIsActive).EnrollmentCodeCode(enrollmentCodeCode).SourceType(sourceType).SourceId(sourceId).Type_(type_).Execute()

Create Blueprint



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
	name := "name_example" // string | (required) Set the name of the Blueprint. The name provided must be unique.
	enrollmentCodeIsActive := "enrollmentCodeIsActive_example" // string | (required) Enable or Disable the Blueprint for manual device enrollment from the enrollment portal
	enrollmentCodeCode := "enrollmentCodeCode_example" // string | Optionally, set the enrollment code of the Blueprint. This key is not required. If an enrollment code is not supplied in the payload body, it will be randomly generated. The enrollment code will be returned in the response and visible in the Web app.
	sourceType := "sourceType_example" // string | Set the source to create the blueprint from. Possible options: <code>template</code> and <code>blueprint</code>.
	sourceId := "sourceId_example" // string | Set either the source template ID, or the source Blueprint ID to clone an existing template or blueprint.
	type_ := "type__example" // string | Choose the type of blueprint to create. Options: <code>classic</code> or <code>map</code>

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintsAPI.BlueprintsCreateBlueprint(context.Background()).Name(name).EnrollmentCodeIsActive(enrollmentCodeIsActive).EnrollmentCodeCode(enrollmentCodeCode).SourceType(sourceType).SourceId(sourceId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.BlueprintsCreateBlueprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlueprintsCreateBlueprint`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BlueprintsAPI.BlueprintsCreateBlueprint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlueprintsCreateBlueprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | (required) Set the name of the Blueprint. The name provided must be unique. | 
 **enrollmentCodeIsActive** | **string** | (required) Enable or Disable the Blueprint for manual device enrollment from the enrollment portal | 
 **enrollmentCodeCode** | **string** | Optionally, set the enrollment code of the Blueprint. This key is not required. If an enrollment code is not supplied in the payload body, it will be randomly generated. The enrollment code will be returned in the response and visible in the Web app. | 
 **sourceType** | **string** | Set the source to create the blueprint from. Possible options: &lt;code&gt;template&lt;/code&gt; and &lt;code&gt;blueprint&lt;/code&gt;. | 
 **sourceId** | **string** | Set either the source template ID, or the source Blueprint ID to clone an existing template or blueprint. | 
 **type_** | **string** | Choose the type of blueprint to create. Options: &lt;code&gt;classic&lt;/code&gt; or &lt;code&gt;map&lt;/code&gt; | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlueprintsDeleteBlueprint

> BlueprintsDeleteBlueprint(ctx, blueprintId).Execute()

Delete Blueprint



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlueprintsAPI.BlueprintsDeleteBlueprint(context.Background(), blueprintId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.BlueprintsDeleteBlueprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blueprintId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlueprintsDeleteBlueprintRequest struct via the builder pattern


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


## BlueprintsGetBlueprint

> BlueprintsGetBlueprint(ctx, blueprintId).Execute()

Get Blueprint



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlueprintsAPI.BlueprintsGetBlueprint(context.Background(), blueprintId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.BlueprintsGetBlueprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blueprintId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlueprintsGetBlueprintRequest struct via the builder pattern


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


## BlueprintsGetBlueprintTemplates

> BlueprintsGetBlueprintTemplates(ctx).Limit(limit).Offset(offset).Execute()

Get Blueprint Templates

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
	limit := "100" // string | Number of results to return per page. (optional)
	offset := "400" // string | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlueprintsAPI.BlueprintsGetBlueprintTemplates(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.BlueprintsGetBlueprintTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlueprintsGetBlueprintTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | Number of results to return per page. | 
 **offset** | **string** | The initial index from which to return the results. | 

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


## BlueprintsGetManualEnrollmentProfile

> string BlueprintsGetManualEnrollmentProfile(ctx, blueprintId).Sso(sso).Execute()

Get Manual Enrollment Profile



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
	sso := "true" // string | Use the <code>sso</code> query parameter, set to <code>true</code>, to return a URL instead of the manual enrollment profile. This parameter should only be used for blueprints in which &quot;Require Authentication&quot; is configured for Manual Enrollment. The returned URL must be used to authenticate via SSO to receive an enrollment profile. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintsAPI.BlueprintsGetManualEnrollmentProfile(context.Background(), blueprintId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.BlueprintsGetManualEnrollmentProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlueprintsGetManualEnrollmentProfile`: string
	fmt.Fprintf(os.Stdout, "Response from `BlueprintsAPI.BlueprintsGetManualEnrollmentProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blueprintId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlueprintsGetManualEnrollmentProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sso** | **string** | Use the &lt;code&gt;sso&lt;/code&gt; query parameter, set to &lt;code&gt;true&lt;/code&gt;, to return a URL instead of the manual enrollment profile. This parameter should only be used for blueprints in which &amp;quot;Require Authentication&amp;quot; is configured for Manual Enrollment. The returned URL must be used to authenticate via SSO to receive an enrollment profile. | 

### Return type

**string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-apple-aspen-config

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlueprintsListBlueprints

> map[string]interface{} BlueprintsListBlueprints(ctx).Id(id).IdIn(idIn).Name(name).Limit(limit).Offset(offset).Execute()

List Blueprints



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
	id := "97e4e175-1631-43f6-a02b-33fd1c748ab8" // string | Look up a specific Blueprint by its ID (optional)
	idIn := "11f4eb9a-10ed-4c3d-a7c1-fb87f95743fb,6391086e-85a1-4820-813c-f9c75025fff4" // string | Specify a list of Blueprint IDs to limit the results to.  Multiple values may be separated by commas. There is a double underscore (<code>__</code>) between id and in (optional)
	name := "testing_blueprint" // string | Return Blueprint names &quot;containing&quot; the specified search string. (optional)
	limit := "100" // string | Number of results to return per page. (optional)
	offset := "400" // string | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintsAPI.BlueprintsListBlueprints(context.Background()).Id(id).IdIn(idIn).Name(name).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.BlueprintsListBlueprints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlueprintsListBlueprints`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BlueprintsAPI.BlueprintsListBlueprints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlueprintsListBlueprintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Look up a specific Blueprint by its ID | 
 **idIn** | **string** | Specify a list of Blueprint IDs to limit the results to.  Multiple values may be separated by commas. There is a double underscore (&lt;code&gt;__&lt;/code&gt;) between id and in | 
 **name** | **string** | Return Blueprint names &amp;quot;containing&amp;quot; the specified search string. | 
 **limit** | **string** | Number of results to return per page. | 
 **offset** | **string** | The initial index from which to return the results. | 

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


## BlueprintsListLibraryItems

> map[string]interface{} BlueprintsListLibraryItems(ctx, blueprintId).Execute()

List Library Items



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintsAPI.BlueprintsListLibraryItems(context.Background(), blueprintId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.BlueprintsListLibraryItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlueprintsListLibraryItems`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BlueprintsAPI.BlueprintsListLibraryItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blueprintId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlueprintsListLibraryItemsRequest struct via the builder pattern


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


## BlueprintsRemoveLibraryItem

> map[string]interface{} BlueprintsRemoveLibraryItem(ctx, blueprintId).Body(body).Execute()

Remove Library Item



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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintsAPI.BlueprintsRemoveLibraryItem(context.Background(), blueprintId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.BlueprintsRemoveLibraryItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlueprintsRemoveLibraryItem`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BlueprintsAPI.BlueprintsRemoveLibraryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blueprintId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlueprintsRemoveLibraryItemRequest struct via the builder pattern


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


## BlueprintsUpdateBlueprint

> map[string]interface{} BlueprintsUpdateBlueprint(ctx, blueprintId).Name(name).Description(description).EnrollmentCodeCode(enrollmentCodeCode).EnrollmentCodeIsActive(enrollmentCodeIsActive).Execute()

Update Blueprint



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
	name := "name_example" // string | Update the name of the Blueprint
	description := "description_example" // string | Update the description of the Blueprint 
	enrollmentCodeCode := "enrollmentCodeCode_example" // string | Update the enrollment code of the Blueprint 
	enrollmentCodeIsActive := "enrollmentCodeIsActive_example" // string | Disable the Blueprint for manual device enrollment from the enrollment portal.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintsAPI.BlueprintsUpdateBlueprint(context.Background(), blueprintId).Name(name).Description(description).EnrollmentCodeCode(enrollmentCodeCode).EnrollmentCodeIsActive(enrollmentCodeIsActive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.BlueprintsUpdateBlueprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlueprintsUpdateBlueprint`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BlueprintsAPI.BlueprintsUpdateBlueprint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blueprintId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlueprintsUpdateBlueprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Update the name of the Blueprint | 
 **description** | **string** | Update the description of the Blueprint  | 
 **enrollmentCodeCode** | **string** | Update the enrollment code of the Blueprint  | 
 **enrollmentCodeIsActive** | **string** | Disable the Blueprint for manual device enrollment from the enrollment portal. | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

