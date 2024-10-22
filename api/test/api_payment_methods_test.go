/*
Commerce Layer API

Testing PaymentMethodsApiService

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

func Test_api_PaymentMethodsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PaymentMethodsApiService DELETEPaymentMethodsPaymentMethodId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var paymentMethodId interface{}

		httpRes, err := apiClient.PaymentMethodsApi.DELETEPaymentMethodsPaymentMethodId(context.Background(), paymentMethodId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService GETAdyenGatewayIdPaymentMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var adyenGatewayId interface{}

		httpRes, err := apiClient.PaymentMethodsApi.GETAdyenGatewayIdPaymentMethods(context.Background(), adyenGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService GETAxerveGatewayIdPaymentMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var axerveGatewayId interface{}

		httpRes, err := apiClient.PaymentMethodsApi.GETAxerveGatewayIdPaymentMethods(context.Background(), axerveGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService GETBraintreeGatewayIdPaymentMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var braintreeGatewayId interface{}

		httpRes, err := apiClient.PaymentMethodsApi.GETBraintreeGatewayIdPaymentMethods(context.Background(), braintreeGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService GETCheckoutComGatewayIdPaymentMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var checkoutComGatewayId interface{}

		httpRes, err := apiClient.PaymentMethodsApi.GETCheckoutComGatewayIdPaymentMethods(context.Background(), checkoutComGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService GETCustomerPaymentSourceIdPaymentMethod", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customerPaymentSourceId interface{}

		httpRes, err := apiClient.PaymentMethodsApi.GETCustomerPaymentSourceIdPaymentMethod(context.Background(), customerPaymentSourceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService GETExternalGatewayIdPaymentMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalGatewayId interface{}

		httpRes, err := apiClient.PaymentMethodsApi.GETExternalGatewayIdPaymentMethods(context.Background(), externalGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService GETKlarnaGatewayIdPaymentMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var klarnaGatewayId interface{}

		httpRes, err := apiClient.PaymentMethodsApi.GETKlarnaGatewayIdPaymentMethods(context.Background(), klarnaGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService GETManualGatewayIdPaymentMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var manualGatewayId interface{}

		httpRes, err := apiClient.PaymentMethodsApi.GETManualGatewayIdPaymentMethods(context.Background(), manualGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService GETOrderIdAvailablePaymentMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderId interface{}

		httpRes, err := apiClient.PaymentMethodsApi.GETOrderIdAvailablePaymentMethods(context.Background(), orderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService GETOrderIdPaymentMethod", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orderId interface{}

		httpRes, err := apiClient.PaymentMethodsApi.GETOrderIdPaymentMethod(context.Background(), orderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService GETPaymentGatewayIdPaymentMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var paymentGatewayId interface{}

		httpRes, err := apiClient.PaymentMethodsApi.GETPaymentGatewayIdPaymentMethods(context.Background(), paymentGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService GETPaymentMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PaymentMethodsApi.GETPaymentMethods(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService GETPaymentMethodsPaymentMethodId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var paymentMethodId interface{}

		resp, httpRes, err := apiClient.PaymentMethodsApi.GETPaymentMethodsPaymentMethodId(context.Background(), paymentMethodId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService GETPaypalGatewayIdPaymentMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var paypalGatewayId interface{}

		httpRes, err := apiClient.PaymentMethodsApi.GETPaypalGatewayIdPaymentMethods(context.Background(), paypalGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService GETSatispayGatewayIdPaymentMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var satispayGatewayId interface{}

		httpRes, err := apiClient.PaymentMethodsApi.GETSatispayGatewayIdPaymentMethods(context.Background(), satispayGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService GETStripeGatewayIdPaymentMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var stripeGatewayId interface{}

		httpRes, err := apiClient.PaymentMethodsApi.GETStripeGatewayIdPaymentMethods(context.Background(), stripeGatewayId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService PATCHPaymentMethodsPaymentMethodId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var paymentMethodId interface{}

		resp, httpRes, err := apiClient.PaymentMethodsApi.PATCHPaymentMethodsPaymentMethodId(context.Background(), paymentMethodId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentMethodsApiService POSTPaymentMethods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PaymentMethodsApi.POSTPaymentMethods(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
