/*
Kandji API

Testing VulnerabilitiesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package kandji

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/MScottBlake/kandji-go-sdk"
)

func Test_kandji_VulnerabilitiesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VulnerabilitiesAPIService GetVulnerabilityDescription", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cveId string

		resp, httpRes, err := apiClient.VulnerabilitiesAPI.GetVulnerabilityDescription(context.Background(), cveId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VulnerabilitiesAPIService ListAffectedApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cveId string

		resp, httpRes, err := apiClient.VulnerabilitiesAPI.ListAffectedApplications(context.Background(), cveId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VulnerabilitiesAPIService ListAffectedDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var cveId string

		resp, httpRes, err := apiClient.VulnerabilitiesAPI.ListAffectedDevices(context.Background(), cveId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VulnerabilitiesAPIService ListDetections", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VulnerabilitiesAPI.ListDetections(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VulnerabilitiesAPIService ListVulnerabilities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VulnerabilitiesAPI.ListVulnerabilities(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
