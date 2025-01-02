/*
Kandji API

Testing InHouseAppsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package kandji_sdk

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/MScottBlake/kandji-go-sdk"
)

func Test_kandji_sdk_InHouseAppsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InHouseAppsAPIService CreateInhouseApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.InHouseAppsAPI.CreateInhouseApp(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InHouseAppsAPIService DeleteInhouseApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var libraryItemId string

		httpRes, err := apiClient.InHouseAppsAPI.DeleteInhouseApp(context.Background(), libraryItemId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InHouseAppsAPIService GetInhouseApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var libraryItemId string

		resp, httpRes, err := apiClient.InHouseAppsAPI.GetInhouseApp(context.Background(), libraryItemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InHouseAppsAPIService ListInhouseApps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InHouseAppsAPI.ListInhouseApps(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InHouseAppsAPIService UpdateInhouseApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var libraryItemId string

		httpRes, err := apiClient.InHouseAppsAPI.UpdateInhouseApp(context.Background(), libraryItemId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InHouseAppsAPIService UploadInhouseApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InHouseAppsAPI.UploadInhouseApp(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InHouseAppsAPIService UploadInhouseAppStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pendingUploadId string

		resp, httpRes, err := apiClient.InHouseAppsAPI.UploadInhouseAppStatus(context.Background(), pendingUploadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}