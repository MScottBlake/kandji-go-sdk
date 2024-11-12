/*
Kandji API

Testing AutomatedDeviceEnrollmentIntegrationsAPIService

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

func Test_kandji_go_sdk_AutomatedDeviceEnrollmentIntegrationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AutomatedDeviceEnrollmentIntegrationsAPIService CreateAdeIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.CreateAdeIntegration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutomatedDeviceEnrollmentIntegrationsAPIService DeleteAdeIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var adeTokenId string

		httpRes, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.DeleteAdeIntegration(context.Background(), adeTokenId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutomatedDeviceEnrollmentIntegrationsAPIService DownloadAdePublicKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.DownloadAdePublicKey(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutomatedDeviceEnrollmentIntegrationsAPIService GetAdeDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		resp, httpRes, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.GetAdeDevice(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutomatedDeviceEnrollmentIntegrationsAPIService GetAdeIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var adeTokenId string

		httpRes, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.GetAdeIntegration(context.Background(), adeTokenId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutomatedDeviceEnrollmentIntegrationsAPIService ListAdeDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.ListAdeDevices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutomatedDeviceEnrollmentIntegrationsAPIService ListAdeIntegrations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.ListAdeIntegrations(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutomatedDeviceEnrollmentIntegrationsAPIService ListDevicesAssociatedToAdeToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var adeTokenId string

		resp, httpRes, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.ListDevicesAssociatedToAdeToken(context.Background(), adeTokenId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutomatedDeviceEnrollmentIntegrationsAPIService RenewAdeIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var adeTokenId string

		httpRes, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.RenewAdeIntegration(context.Background(), adeTokenId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutomatedDeviceEnrollmentIntegrationsAPIService UpdateAdeDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		resp, httpRes, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.UpdateAdeDevice(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AutomatedDeviceEnrollmentIntegrationsAPIService UpdateAdeIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var adeTokenId string

		httpRes, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.UpdateAdeIntegration(context.Background(), adeTokenId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
