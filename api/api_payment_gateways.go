/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// PaymentGatewaysApiService PaymentGatewaysApi service
type PaymentGatewaysApiService service

type PaymentGatewaysApiGETAdyenPaymentIdPaymentGatewayRequest struct {
	ctx context.Context
	ApiService *PaymentGatewaysApiService
	adyenPaymentId string
}

func (r PaymentGatewaysApiGETAdyenPaymentIdPaymentGatewayRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETAdyenPaymentIdPaymentGatewayExecute(r)
}

/*
GETAdyenPaymentIdPaymentGateway Retrieve the payment gateway associated to the adyen payment

Retrieve the payment gateway associated to the adyen payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param adyenPaymentId The resource's id
 @return PaymentGatewaysApiGETAdyenPaymentIdPaymentGatewayRequest
*/
func (a *PaymentGatewaysApiService) GETAdyenPaymentIdPaymentGateway(ctx context.Context, adyenPaymentId string) PaymentGatewaysApiGETAdyenPaymentIdPaymentGatewayRequest {
	return PaymentGatewaysApiGETAdyenPaymentIdPaymentGatewayRequest{
		ApiService: a,
		ctx: ctx,
		adyenPaymentId: adyenPaymentId,
	}
}

// Execute executes the request
func (a *PaymentGatewaysApiService) GETAdyenPaymentIdPaymentGatewayExecute(r PaymentGatewaysApiGETAdyenPaymentIdPaymentGatewayRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentGatewaysApiService.GETAdyenPaymentIdPaymentGateway")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_payments/{adyenPaymentId}/payment_gateway"
	localVarPath = strings.Replace(localVarPath, "{"+"adyenPaymentId"+"}", url.PathEscape(parameterToString(r.adyenPaymentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type PaymentGatewaysApiGETBraintreePaymentIdPaymentGatewayRequest struct {
	ctx context.Context
	ApiService *PaymentGatewaysApiService
	braintreePaymentId string
}

func (r PaymentGatewaysApiGETBraintreePaymentIdPaymentGatewayRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETBraintreePaymentIdPaymentGatewayExecute(r)
}

/*
GETBraintreePaymentIdPaymentGateway Retrieve the payment gateway associated to the braintree payment

Retrieve the payment gateway associated to the braintree payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param braintreePaymentId The resource's id
 @return PaymentGatewaysApiGETBraintreePaymentIdPaymentGatewayRequest
*/
func (a *PaymentGatewaysApiService) GETBraintreePaymentIdPaymentGateway(ctx context.Context, braintreePaymentId string) PaymentGatewaysApiGETBraintreePaymentIdPaymentGatewayRequest {
	return PaymentGatewaysApiGETBraintreePaymentIdPaymentGatewayRequest{
		ApiService: a,
		ctx: ctx,
		braintreePaymentId: braintreePaymentId,
	}
}

// Execute executes the request
func (a *PaymentGatewaysApiService) GETBraintreePaymentIdPaymentGatewayExecute(r PaymentGatewaysApiGETBraintreePaymentIdPaymentGatewayRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentGatewaysApiService.GETBraintreePaymentIdPaymentGateway")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_payments/{braintreePaymentId}/payment_gateway"
	localVarPath = strings.Replace(localVarPath, "{"+"braintreePaymentId"+"}", url.PathEscape(parameterToString(r.braintreePaymentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type PaymentGatewaysApiGETCheckoutComPaymentIdPaymentGatewayRequest struct {
	ctx context.Context
	ApiService *PaymentGatewaysApiService
	checkoutComPaymentId string
}

func (r PaymentGatewaysApiGETCheckoutComPaymentIdPaymentGatewayRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCheckoutComPaymentIdPaymentGatewayExecute(r)
}

/*
GETCheckoutComPaymentIdPaymentGateway Retrieve the payment gateway associated to the checkout.com payment

Retrieve the payment gateway associated to the checkout.com payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param checkoutComPaymentId The resource's id
 @return PaymentGatewaysApiGETCheckoutComPaymentIdPaymentGatewayRequest
*/
func (a *PaymentGatewaysApiService) GETCheckoutComPaymentIdPaymentGateway(ctx context.Context, checkoutComPaymentId string) PaymentGatewaysApiGETCheckoutComPaymentIdPaymentGatewayRequest {
	return PaymentGatewaysApiGETCheckoutComPaymentIdPaymentGatewayRequest{
		ApiService: a,
		ctx: ctx,
		checkoutComPaymentId: checkoutComPaymentId,
	}
}

// Execute executes the request
func (a *PaymentGatewaysApiService) GETCheckoutComPaymentIdPaymentGatewayExecute(r PaymentGatewaysApiGETCheckoutComPaymentIdPaymentGatewayRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentGatewaysApiService.GETCheckoutComPaymentIdPaymentGateway")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_payments/{checkoutComPaymentId}/payment_gateway"
	localVarPath = strings.Replace(localVarPath, "{"+"checkoutComPaymentId"+"}", url.PathEscape(parameterToString(r.checkoutComPaymentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type PaymentGatewaysApiGETExternalPaymentIdPaymentGatewayRequest struct {
	ctx context.Context
	ApiService *PaymentGatewaysApiService
	externalPaymentId string
}

func (r PaymentGatewaysApiGETExternalPaymentIdPaymentGatewayRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETExternalPaymentIdPaymentGatewayExecute(r)
}

/*
GETExternalPaymentIdPaymentGateway Retrieve the payment gateway associated to the external payment

Retrieve the payment gateway associated to the external payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalPaymentId The resource's id
 @return PaymentGatewaysApiGETExternalPaymentIdPaymentGatewayRequest
*/
func (a *PaymentGatewaysApiService) GETExternalPaymentIdPaymentGateway(ctx context.Context, externalPaymentId string) PaymentGatewaysApiGETExternalPaymentIdPaymentGatewayRequest {
	return PaymentGatewaysApiGETExternalPaymentIdPaymentGatewayRequest{
		ApiService: a,
		ctx: ctx,
		externalPaymentId: externalPaymentId,
	}
}

// Execute executes the request
func (a *PaymentGatewaysApiService) GETExternalPaymentIdPaymentGatewayExecute(r PaymentGatewaysApiGETExternalPaymentIdPaymentGatewayRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentGatewaysApiService.GETExternalPaymentIdPaymentGateway")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_payments/{externalPaymentId}/payment_gateway"
	localVarPath = strings.Replace(localVarPath, "{"+"externalPaymentId"+"}", url.PathEscape(parameterToString(r.externalPaymentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type PaymentGatewaysApiGETKlarnaPaymentIdPaymentGatewayRequest struct {
	ctx context.Context
	ApiService *PaymentGatewaysApiService
	klarnaPaymentId string
}

func (r PaymentGatewaysApiGETKlarnaPaymentIdPaymentGatewayRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETKlarnaPaymentIdPaymentGatewayExecute(r)
}

/*
GETKlarnaPaymentIdPaymentGateway Retrieve the payment gateway associated to the klarna payment

Retrieve the payment gateway associated to the klarna payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param klarnaPaymentId The resource's id
 @return PaymentGatewaysApiGETKlarnaPaymentIdPaymentGatewayRequest
*/
func (a *PaymentGatewaysApiService) GETKlarnaPaymentIdPaymentGateway(ctx context.Context, klarnaPaymentId string) PaymentGatewaysApiGETKlarnaPaymentIdPaymentGatewayRequest {
	return PaymentGatewaysApiGETKlarnaPaymentIdPaymentGatewayRequest{
		ApiService: a,
		ctx: ctx,
		klarnaPaymentId: klarnaPaymentId,
	}
}

// Execute executes the request
func (a *PaymentGatewaysApiService) GETKlarnaPaymentIdPaymentGatewayExecute(r PaymentGatewaysApiGETKlarnaPaymentIdPaymentGatewayRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentGatewaysApiService.GETKlarnaPaymentIdPaymentGateway")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/klarna_payments/{klarnaPaymentId}/payment_gateway"
	localVarPath = strings.Replace(localVarPath, "{"+"klarnaPaymentId"+"}", url.PathEscape(parameterToString(r.klarnaPaymentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type PaymentGatewaysApiGETPaymentGatewaysRequest struct {
	ctx context.Context
	ApiService *PaymentGatewaysApiService
}

func (r PaymentGatewaysApiGETPaymentGatewaysRequest) Execute() (*GETPaymentGateways200Response, *http.Response, error) {
	return r.ApiService.GETPaymentGatewaysExecute(r)
}

/*
GETPaymentGateways List all payment gateways

List all payment gateways

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PaymentGatewaysApiGETPaymentGatewaysRequest
*/
func (a *PaymentGatewaysApiService) GETPaymentGateways(ctx context.Context) PaymentGatewaysApiGETPaymentGatewaysRequest {
	return PaymentGatewaysApiGETPaymentGatewaysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GETPaymentGateways200Response
func (a *PaymentGatewaysApiService) GETPaymentGatewaysExecute(r PaymentGatewaysApiGETPaymentGatewaysRequest) (*GETPaymentGateways200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GETPaymentGateways200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentGatewaysApiService.GETPaymentGateways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment_gateways"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PaymentGatewaysApiGETPaymentGatewaysPaymentGatewayIdRequest struct {
	ctx context.Context
	ApiService *PaymentGatewaysApiService
	paymentGatewayId string
}

func (r PaymentGatewaysApiGETPaymentGatewaysPaymentGatewayIdRequest) Execute() (*GETPaymentGatewaysPaymentGatewayId200Response, *http.Response, error) {
	return r.ApiService.GETPaymentGatewaysPaymentGatewayIdExecute(r)
}

/*
GETPaymentGatewaysPaymentGatewayId Retrieve a payment gateway

Retrieve a payment gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentGatewayId The resource's id
 @return PaymentGatewaysApiGETPaymentGatewaysPaymentGatewayIdRequest
*/
func (a *PaymentGatewaysApiService) GETPaymentGatewaysPaymentGatewayId(ctx context.Context, paymentGatewayId string) PaymentGatewaysApiGETPaymentGatewaysPaymentGatewayIdRequest {
	return PaymentGatewaysApiGETPaymentGatewaysPaymentGatewayIdRequest{
		ApiService: a,
		ctx: ctx,
		paymentGatewayId: paymentGatewayId,
	}
}

// Execute executes the request
//  @return GETPaymentGatewaysPaymentGatewayId200Response
func (a *PaymentGatewaysApiService) GETPaymentGatewaysPaymentGatewayIdExecute(r PaymentGatewaysApiGETPaymentGatewaysPaymentGatewayIdRequest) (*GETPaymentGatewaysPaymentGatewayId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GETPaymentGatewaysPaymentGatewayId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentGatewaysApiService.GETPaymentGatewaysPaymentGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment_gateways/{paymentGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"paymentGatewayId"+"}", url.PathEscape(parameterToString(r.paymentGatewayId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PaymentGatewaysApiGETPaymentMethodIdPaymentGatewayRequest struct {
	ctx context.Context
	ApiService *PaymentGatewaysApiService
	paymentMethodId string
}

func (r PaymentGatewaysApiGETPaymentMethodIdPaymentGatewayRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETPaymentMethodIdPaymentGatewayExecute(r)
}

/*
GETPaymentMethodIdPaymentGateway Retrieve the payment gateway associated to the payment method

Retrieve the payment gateway associated to the payment method

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentMethodId The resource's id
 @return PaymentGatewaysApiGETPaymentMethodIdPaymentGatewayRequest
*/
func (a *PaymentGatewaysApiService) GETPaymentMethodIdPaymentGateway(ctx context.Context, paymentMethodId string) PaymentGatewaysApiGETPaymentMethodIdPaymentGatewayRequest {
	return PaymentGatewaysApiGETPaymentMethodIdPaymentGatewayRequest{
		ApiService: a,
		ctx: ctx,
		paymentMethodId: paymentMethodId,
	}
}

// Execute executes the request
func (a *PaymentGatewaysApiService) GETPaymentMethodIdPaymentGatewayExecute(r PaymentGatewaysApiGETPaymentMethodIdPaymentGatewayRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentGatewaysApiService.GETPaymentMethodIdPaymentGateway")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment_methods/{paymentMethodId}/payment_gateway"
	localVarPath = strings.Replace(localVarPath, "{"+"paymentMethodId"+"}", url.PathEscape(parameterToString(r.paymentMethodId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type PaymentGatewaysApiGETPaypalPaymentIdPaymentGatewayRequest struct {
	ctx context.Context
	ApiService *PaymentGatewaysApiService
	paypalPaymentId string
}

func (r PaymentGatewaysApiGETPaypalPaymentIdPaymentGatewayRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETPaypalPaymentIdPaymentGatewayExecute(r)
}

/*
GETPaypalPaymentIdPaymentGateway Retrieve the payment gateway associated to the paypal payment

Retrieve the payment gateway associated to the paypal payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paypalPaymentId The resource's id
 @return PaymentGatewaysApiGETPaypalPaymentIdPaymentGatewayRequest
*/
func (a *PaymentGatewaysApiService) GETPaypalPaymentIdPaymentGateway(ctx context.Context, paypalPaymentId string) PaymentGatewaysApiGETPaypalPaymentIdPaymentGatewayRequest {
	return PaymentGatewaysApiGETPaypalPaymentIdPaymentGatewayRequest{
		ApiService: a,
		ctx: ctx,
		paypalPaymentId: paypalPaymentId,
	}
}

// Execute executes the request
func (a *PaymentGatewaysApiService) GETPaypalPaymentIdPaymentGatewayExecute(r PaymentGatewaysApiGETPaypalPaymentIdPaymentGatewayRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentGatewaysApiService.GETPaypalPaymentIdPaymentGateway")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/paypal_payments/{paypalPaymentId}/payment_gateway"
	localVarPath = strings.Replace(localVarPath, "{"+"paypalPaymentId"+"}", url.PathEscape(parameterToString(r.paypalPaymentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type PaymentGatewaysApiGETStripePaymentIdPaymentGatewayRequest struct {
	ctx context.Context
	ApiService *PaymentGatewaysApiService
	stripePaymentId string
}

func (r PaymentGatewaysApiGETStripePaymentIdPaymentGatewayRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStripePaymentIdPaymentGatewayExecute(r)
}

/*
GETStripePaymentIdPaymentGateway Retrieve the payment gateway associated to the stripe payment

Retrieve the payment gateway associated to the stripe payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stripePaymentId The resource's id
 @return PaymentGatewaysApiGETStripePaymentIdPaymentGatewayRequest
*/
func (a *PaymentGatewaysApiService) GETStripePaymentIdPaymentGateway(ctx context.Context, stripePaymentId string) PaymentGatewaysApiGETStripePaymentIdPaymentGatewayRequest {
	return PaymentGatewaysApiGETStripePaymentIdPaymentGatewayRequest{
		ApiService: a,
		ctx: ctx,
		stripePaymentId: stripePaymentId,
	}
}

// Execute executes the request
func (a *PaymentGatewaysApiService) GETStripePaymentIdPaymentGatewayExecute(r PaymentGatewaysApiGETStripePaymentIdPaymentGatewayRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentGatewaysApiService.GETStripePaymentIdPaymentGateway")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stripe_payments/{stripePaymentId}/payment_gateway"
	localVarPath = strings.Replace(localVarPath, "{"+"stripePaymentId"+"}", url.PathEscape(parameterToString(r.stripePaymentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
