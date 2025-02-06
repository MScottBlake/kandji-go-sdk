# \SelfServiceAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SelfServiceListSelfServiceCategories**](SelfServiceAPI.md#SelfServiceListSelfServiceCategories) | **Get** /api/v1/self-service/categories | List Self Service Categories



## SelfServiceListSelfServiceCategories

> map[string]interface{} SelfServiceListSelfServiceCategories(ctx).Execute()

List Self Service Categories



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
	resp, r, err := apiClient.SelfServiceAPI.SelfServiceListSelfServiceCategories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceAPI.SelfServiceListSelfServiceCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelfServiceListSelfServiceCategories`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SelfServiceAPI.SelfServiceListSelfServiceCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSelfServiceListSelfServiceCategoriesRequest struct via the builder pattern


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

