/*
Kandji API

Testing TagsAPIService

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

func Test_kandji_go_sdk_TagsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TagsAPIService CreateTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TagsAPI.CreateTag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService DeleteTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tagId string

		httpRes, err := apiClient.TagsAPI.DeleteTag(context.Background(), tagId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService GetTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TagsAPI.GetTags(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService UpdateTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tagId string

		resp, httpRes, err := apiClient.TagsAPI.UpdateTag(context.Background(), tagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}