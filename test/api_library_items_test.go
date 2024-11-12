/*
Kandji API

Testing LibraryItemsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package kandji_go_sdk

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/MScottBlake/kandji-go-sdk"
)

func Test_kandji_go_sdk_LibraryItemsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LibraryItemsAPIService GetLibraryItemActivity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var libraryItemId string

		resp, httpRes, err := apiClient.LibraryItemsAPI.GetLibraryItemActivity(context.Background(), libraryItemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryItemsAPIService GetLibraryItemStatuses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var libraryItemId string

		resp, httpRes, err := apiClient.LibraryItemsAPI.GetLibraryItemStatuses(context.Background(), libraryItemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
