/*
Commerce Layer API

Testing EventsApiService

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

func Test_api_EventsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EventsApiService GETAuthorizationIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var authorizationId interface{}

		httpRes, err := apiClient.EventsApi.GETAuthorizationIdEvents(context.Background(), authorizationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETCaptureIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var captureId interface{}

		httpRes, err := apiClient.EventsApi.GETCaptureIdEvents(context.Background(), captureId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETCleanupIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cleanupId interface{}

		httpRes, err := apiClient.EventsApi.GETCleanupIdEvents(context.Background(), cleanupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETCustomerAddressIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerAddressId interface{}

		httpRes, err := apiClient.EventsApi.GETCustomerAddressIdEvents(context.Background(), customerAddressId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETCustomerIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerId interface{}

		httpRes, err := apiClient.EventsApi.GETCustomerIdEvents(context.Background(), customerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETCustomerPasswordResetIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerPasswordResetId interface{}

		httpRes, err := apiClient.EventsApi.GETCustomerPasswordResetIdEvents(context.Background(), customerPasswordResetId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETCustomerSubscriptionIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerSubscriptionId interface{}

		httpRes, err := apiClient.EventsApi.GETCustomerSubscriptionIdEvents(context.Background(), customerSubscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EventsApi.GETEvents(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETEventsEventId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var eventId interface{}

		resp, httpRes, err := apiClient.EventsApi.GETEventsEventId(context.Background(), eventId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETExportIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var exportId interface{}

		httpRes, err := apiClient.EventsApi.GETExportIdEvents(context.Background(), exportId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETExternalPromotionIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalPromotionId interface{}

		httpRes, err := apiClient.EventsApi.GETExternalPromotionIdEvents(context.Background(), externalPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETFixedAmountPromotionIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fixedAmountPromotionId interface{}

		httpRes, err := apiClient.EventsApi.GETFixedAmountPromotionIdEvents(context.Background(), fixedAmountPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETFixedPricePromotionIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fixedPricePromotionId interface{}

		httpRes, err := apiClient.EventsApi.GETFixedPricePromotionIdEvents(context.Background(), fixedPricePromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETFreeGiftPromotionIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var freeGiftPromotionId interface{}

		httpRes, err := apiClient.EventsApi.GETFreeGiftPromotionIdEvents(context.Background(), freeGiftPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETFreeShippingPromotionIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var freeShippingPromotionId interface{}

		httpRes, err := apiClient.EventsApi.GETFreeShippingPromotionIdEvents(context.Background(), freeShippingPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETGiftCardIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var giftCardId interface{}

		httpRes, err := apiClient.EventsApi.GETGiftCardIdEvents(context.Background(), giftCardId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETImportIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var importId interface{}

		httpRes, err := apiClient.EventsApi.GETImportIdEvents(context.Background(), importId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETInStockSubscriptionIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var inStockSubscriptionId interface{}

		httpRes, err := apiClient.EventsApi.GETInStockSubscriptionIdEvents(context.Background(), inStockSubscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETOrderCopyIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderCopyId interface{}

		httpRes, err := apiClient.EventsApi.GETOrderCopyIdEvents(context.Background(), orderCopyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETOrderFactoryIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderFactoryId interface{}

		httpRes, err := apiClient.EventsApi.GETOrderFactoryIdEvents(context.Background(), orderFactoryId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETOrderIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderId interface{}

		httpRes, err := apiClient.EventsApi.GETOrderIdEvents(context.Background(), orderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETOrderSubscriptionIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderSubscriptionId interface{}

		httpRes, err := apiClient.EventsApi.GETOrderSubscriptionIdEvents(context.Background(), orderSubscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETParcelIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var parcelId interface{}

		httpRes, err := apiClient.EventsApi.GETParcelIdEvents(context.Background(), parcelId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETPercentageDiscountPromotionIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var percentageDiscountPromotionId interface{}

		httpRes, err := apiClient.EventsApi.GETPercentageDiscountPromotionIdEvents(context.Background(), percentageDiscountPromotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETPriceFrequencyTierIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceFrequencyTierId interface{}

		httpRes, err := apiClient.EventsApi.GETPriceFrequencyTierIdEvents(context.Background(), priceFrequencyTierId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETPriceVolumeTierIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var priceVolumeTierId interface{}

		httpRes, err := apiClient.EventsApi.GETPriceVolumeTierIdEvents(context.Background(), priceVolumeTierId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETPromotionIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var promotionId interface{}

		httpRes, err := apiClient.EventsApi.GETPromotionIdEvents(context.Background(), promotionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETRecurringOrderCopyIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var recurringOrderCopyId interface{}

		httpRes, err := apiClient.EventsApi.GETRecurringOrderCopyIdEvents(context.Background(), recurringOrderCopyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETRefundIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var refundId interface{}

		httpRes, err := apiClient.EventsApi.GETRefundIdEvents(context.Background(), refundId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETReturnIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var returnId interface{}

		httpRes, err := apiClient.EventsApi.GETReturnIdEvents(context.Background(), returnId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETShipmentIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var shipmentId interface{}

		httpRes, err := apiClient.EventsApi.GETShipmentIdEvents(context.Background(), shipmentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETStockTransferIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var stockTransferId interface{}

		httpRes, err := apiClient.EventsApi.GETStockTransferIdEvents(context.Background(), stockTransferId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsApiService GETVoidIdEvents", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var voidId interface{}

		httpRes, err := apiClient.EventsApi.GETVoidIdEvents(context.Background(), voidId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
