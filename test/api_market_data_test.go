/*
Paxos API

Testing MarketDataAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package paxos

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/avianlabs/paxos-go"
)

func Test_paxos_MarketDataAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MarketDataAPIService GetOrderBook", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var market string

		resp, httpRes, err := apiClient.MarketDataAPI.GetOrderBook(context.Background(), market).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketDataAPIService GetTicker", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var market string

		resp, httpRes, err := apiClient.MarketDataAPI.GetTicker(context.Background(), market).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketDataAPIService ListMarkets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MarketDataAPI.ListMarkets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketDataAPIService ListRecentExecutions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var market string

		resp, httpRes, err := apiClient.MarketDataAPI.ListRecentExecutions(context.Background(), market).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
