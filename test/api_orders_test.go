/*
Paxos API

Testing OrdersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package paxos

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_paxos_OrdersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrdersAPIService CancelOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var profileId string
		var id string

		resp, httpRes, err := apiClient.OrdersAPI.CancelOrder(context.Background(), profileId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersAPIService CreateOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var profileId string

		resp, httpRes, err := apiClient.OrdersAPI.CreateOrder(context.Background(), profileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersAPIService GetOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var profileId string
		var id string

		resp, httpRes, err := apiClient.OrdersAPI.GetOrder(context.Background(), profileId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersAPIService ListExecutions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrdersAPI.ListExecutions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersAPIService ListOrders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrdersAPI.ListOrders(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
