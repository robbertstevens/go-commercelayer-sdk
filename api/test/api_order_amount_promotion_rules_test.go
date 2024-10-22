/*
Commerce Layer API

Testing OrderAmountPromotionRulesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api

import (
	"context"
	openapiclient "github.com/incentro-dc/go-commercelayer-sdk/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_api_OrderAmountPromotionRulesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrderAmountPromotionRulesApiService DELETEOrderAmountPromotionRulesOrderAmountPromotionRuleId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderAmountPromotionRuleId interface{}

		httpRes, err := apiClient.OrderAmountPromotionRulesApi.DELETEOrderAmountPromotionRulesOrderAmountPromotionRuleId(context.Background(), orderAmountPromotionRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderAmountPromotionRulesApiService GETBuyXPayYPromotionIdOrderAmountPromotionRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var buyXPayYPromotionId interface{}

		httpRes, err := apiClient.OrderAmountPromotionRulesApi.GETBuyXPayYPromotionIdOrderAmountPromotionRule(context.Background(), buyXPayYPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderAmountPromotionRulesApiService GETExternalPromotionIdOrderAmountPromotionRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalPromotionId interface{}

		httpRes, err := apiClient.OrderAmountPromotionRulesApi.GETExternalPromotionIdOrderAmountPromotionRule(context.Background(), externalPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderAmountPromotionRulesApiService GETFixedAmountPromotionIdOrderAmountPromotionRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fixedAmountPromotionId interface{}

		httpRes, err := apiClient.OrderAmountPromotionRulesApi.GETFixedAmountPromotionIdOrderAmountPromotionRule(context.Background(), fixedAmountPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderAmountPromotionRulesApiService GETFixedPricePromotionIdOrderAmountPromotionRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fixedPricePromotionId interface{}

		httpRes, err := apiClient.OrderAmountPromotionRulesApi.GETFixedPricePromotionIdOrderAmountPromotionRule(context.Background(), fixedPricePromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderAmountPromotionRulesApiService GETFreeGiftPromotionIdOrderAmountPromotionRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var freeGiftPromotionId interface{}

		httpRes, err := apiClient.OrderAmountPromotionRulesApi.GETFreeGiftPromotionIdOrderAmountPromotionRule(context.Background(), freeGiftPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderAmountPromotionRulesApiService GETFreeShippingPromotionIdOrderAmountPromotionRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var freeShippingPromotionId interface{}

		httpRes, err := apiClient.OrderAmountPromotionRulesApi.GETFreeShippingPromotionIdOrderAmountPromotionRule(context.Background(), freeShippingPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderAmountPromotionRulesApiService GETOrderAmountPromotionRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OrderAmountPromotionRulesApi.GETOrderAmountPromotionRules(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderAmountPromotionRulesApiService GETOrderAmountPromotionRulesOrderAmountPromotionRuleId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderAmountPromotionRuleId interface{}

		resp, httpRes, err := apiClient.OrderAmountPromotionRulesApi.GETOrderAmountPromotionRulesOrderAmountPromotionRuleId(context.Background(), orderAmountPromotionRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderAmountPromotionRulesApiService GETPercentageDiscountPromotionIdOrderAmountPromotionRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var percentageDiscountPromotionId interface{}

		httpRes, err := apiClient.OrderAmountPromotionRulesApi.GETPercentageDiscountPromotionIdOrderAmountPromotionRule(context.Background(), percentageDiscountPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderAmountPromotionRulesApiService GETPromotionIdOrderAmountPromotionRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var promotionId interface{}

		httpRes, err := apiClient.OrderAmountPromotionRulesApi.GETPromotionIdOrderAmountPromotionRule(context.Background(), promotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderAmountPromotionRulesApiService PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderAmountPromotionRuleId interface{}

		resp, httpRes, err := apiClient.OrderAmountPromotionRulesApi.PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleId(context.Background(), orderAmountPromotionRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrderAmountPromotionRulesApiService POSTOrderAmountPromotionRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OrderAmountPromotionRulesApi.POSTOrderAmountPromotionRules(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
