/*
Commerce Layer API

Testing MarketsApiService

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

func Test_api_MarketsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MarketsApiService DELETEMarketsMarketId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var marketId interface{}

		httpRes, err := apiClient.MarketsApi.DELETEMarketsMarketId(context.Background(), marketId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETAvalaraAccountIdMarkets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var avalaraAccountId interface{}

		httpRes, err := apiClient.MarketsApi.GETAvalaraAccountIdMarkets(context.Background(), avalaraAccountId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETBillingInfoValidationRuleIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var billingInfoValidationRuleId interface{}

		httpRes, err := apiClient.MarketsApi.GETBillingInfoValidationRuleIdMarket(context.Background(), billingInfoValidationRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETBingGeocoderIdMarkets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bingGeocoderId interface{}

		httpRes, err := apiClient.MarketsApi.GETBingGeocoderIdMarkets(context.Background(), bingGeocoderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETBundleIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var bundleId interface{}

		httpRes, err := apiClient.MarketsApi.GETBundleIdMarket(context.Background(), bundleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETBuyXPayYPromotionIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var buyXPayYPromotionId interface{}

		httpRes, err := apiClient.MarketsApi.GETBuyXPayYPromotionIdMarket(context.Background(), buyXPayYPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETCarrierAccountIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var carrierAccountId interface{}

		httpRes, err := apiClient.MarketsApi.GETCarrierAccountIdMarket(context.Background(), carrierAccountId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETCustomerGroupIdMarkets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerGroupId interface{}

		httpRes, err := apiClient.MarketsApi.GETCustomerGroupIdMarkets(context.Background(), customerGroupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETExternalPromotionIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalPromotionId interface{}

		httpRes, err := apiClient.MarketsApi.GETExternalPromotionIdMarket(context.Background(), externalPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETExternalTaxCalculatorIdMarkets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalTaxCalculatorId interface{}

		httpRes, err := apiClient.MarketsApi.GETExternalTaxCalculatorIdMarkets(context.Background(), externalTaxCalculatorId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETFixedAmountPromotionIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fixedAmountPromotionId interface{}

		httpRes, err := apiClient.MarketsApi.GETFixedAmountPromotionIdMarket(context.Background(), fixedAmountPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETFixedPricePromotionIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fixedPricePromotionId interface{}

		httpRes, err := apiClient.MarketsApi.GETFixedPricePromotionIdMarket(context.Background(), fixedPricePromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETFreeGiftPromotionIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var freeGiftPromotionId interface{}

		httpRes, err := apiClient.MarketsApi.GETFreeGiftPromotionIdMarket(context.Background(), freeGiftPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETFreeShippingPromotionIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var freeShippingPromotionId interface{}

		httpRes, err := apiClient.MarketsApi.GETFreeShippingPromotionIdMarket(context.Background(), freeShippingPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETGeocoderIdMarkets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var geocoderId interface{}

		httpRes, err := apiClient.MarketsApi.GETGeocoderIdMarkets(context.Background(), geocoderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETGiftCardIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var giftCardId interface{}

		httpRes, err := apiClient.MarketsApi.GETGiftCardIdMarket(context.Background(), giftCardId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETGoogleGeocoderIdMarkets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var googleGeocoderId interface{}

		httpRes, err := apiClient.MarketsApi.GETGoogleGeocoderIdMarkets(context.Background(), googleGeocoderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETInStockSubscriptionIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var inStockSubscriptionId interface{}

		httpRes, err := apiClient.MarketsApi.GETInStockSubscriptionIdMarket(context.Background(), inStockSubscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETManualTaxCalculatorIdMarkets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var manualTaxCalculatorId interface{}

		httpRes, err := apiClient.MarketsApi.GETManualTaxCalculatorIdMarkets(context.Background(), manualTaxCalculatorId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETMarkets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MarketsApi.GETMarkets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETMarketsMarketId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var marketId interface{}

		resp, httpRes, err := apiClient.MarketsApi.GETMarketsMarketId(context.Background(), marketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETOrderIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderId interface{}

		httpRes, err := apiClient.MarketsApi.GETOrderIdMarket(context.Background(), orderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETOrderSubscriptionIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderSubscriptionId interface{}

		httpRes, err := apiClient.MarketsApi.GETOrderSubscriptionIdMarket(context.Background(), orderSubscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETPaymentMethodIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var paymentMethodId interface{}

		httpRes, err := apiClient.MarketsApi.GETPaymentMethodIdMarket(context.Background(), paymentMethodId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETPercentageDiscountPromotionIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var percentageDiscountPromotionId interface{}

		httpRes, err := apiClient.MarketsApi.GETPercentageDiscountPromotionIdMarket(context.Background(), percentageDiscountPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETPriceIdJwtMarkets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceId interface{}

		httpRes, err := apiClient.MarketsApi.GETPriceIdJwtMarkets(context.Background(), priceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETPriceListSchedulerIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceListSchedulerId interface{}

		httpRes, err := apiClient.MarketsApi.GETPriceListSchedulerIdMarket(context.Background(), priceListSchedulerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETPromotionIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var promotionId interface{}

		httpRes, err := apiClient.MarketsApi.GETPromotionIdMarket(context.Background(), promotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETShippingMethodIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shippingMethodId interface{}

		httpRes, err := apiClient.MarketsApi.GETShippingMethodIdMarket(context.Background(), shippingMethodId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETSkuIdJwtMarkets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var skuId interface{}

		httpRes, err := apiClient.MarketsApi.GETSkuIdJwtMarkets(context.Background(), skuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETSkuOptionIdMarket", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var skuOptionId interface{}

		httpRes, err := apiClient.MarketsApi.GETSkuOptionIdMarket(context.Background(), skuOptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETSubscriptionModelIdMarkets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var subscriptionModelId interface{}

		httpRes, err := apiClient.MarketsApi.GETSubscriptionModelIdMarkets(context.Background(), subscriptionModelId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETTaxCalculatorIdMarkets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taxCalculatorId interface{}

		httpRes, err := apiClient.MarketsApi.GETTaxCalculatorIdMarkets(context.Background(), taxCalculatorId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService GETTaxjarAccountIdMarkets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var taxjarAccountId interface{}

		httpRes, err := apiClient.MarketsApi.GETTaxjarAccountIdMarkets(context.Background(), taxjarAccountId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService PATCHMarketsMarketId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var marketId interface{}

		resp, httpRes, err := apiClient.MarketsApi.PATCHMarketsMarketId(context.Background(), marketId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketsApiService POSTMarkets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MarketsApi.POSTMarkets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
