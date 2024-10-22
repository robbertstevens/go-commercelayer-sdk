/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// BraintreePaymentsApiService BraintreePaymentsApi service
type BraintreePaymentsApiService service

type BraintreePaymentsApiDELETEBraintreePaymentsBraintreePaymentIdRequest struct {
	ctx                context.Context
	ApiService         *BraintreePaymentsApiService
	braintreePaymentId interface{}
}

func (r BraintreePaymentsApiDELETEBraintreePaymentsBraintreePaymentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEBraintreePaymentsBraintreePaymentIdExecute(r)
}

/*
DELETEBraintreePaymentsBraintreePaymentId Delete a braintree payment

Delete a braintree payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param braintreePaymentId The resource's id
	@return BraintreePaymentsApiDELETEBraintreePaymentsBraintreePaymentIdRequest
*/
func (a *BraintreePaymentsApiService) DELETEBraintreePaymentsBraintreePaymentId(ctx context.Context, braintreePaymentId interface{}) BraintreePaymentsApiDELETEBraintreePaymentsBraintreePaymentIdRequest {
	return BraintreePaymentsApiDELETEBraintreePaymentsBraintreePaymentIdRequest{
		ApiService:         a,
		ctx:                ctx,
		braintreePaymentId: braintreePaymentId,
	}
}

// Execute executes the request
func (a *BraintreePaymentsApiService) DELETEBraintreePaymentsBraintreePaymentIdExecute(r BraintreePaymentsApiDELETEBraintreePaymentsBraintreePaymentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BraintreePaymentsApiService.DELETEBraintreePaymentsBraintreePaymentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_payments/{braintreePaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"braintreePaymentId"+"}", url.PathEscape(parameterValueToString(r.braintreePaymentId, "braintreePaymentId")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type BraintreePaymentsApiGETBraintreeGatewayIdBraintreePaymentsRequest struct {
	ctx                context.Context
	ApiService         *BraintreePaymentsApiService
	braintreeGatewayId interface{}
}

func (r BraintreePaymentsApiGETBraintreeGatewayIdBraintreePaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETBraintreeGatewayIdBraintreePaymentsExecute(r)
}

/*
GETBraintreeGatewayIdBraintreePayments Retrieve the braintree payments associated to the braintree gateway

Retrieve the braintree payments associated to the braintree gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param braintreeGatewayId The resource's id
	@return BraintreePaymentsApiGETBraintreeGatewayIdBraintreePaymentsRequest
*/
func (a *BraintreePaymentsApiService) GETBraintreeGatewayIdBraintreePayments(ctx context.Context, braintreeGatewayId interface{}) BraintreePaymentsApiGETBraintreeGatewayIdBraintreePaymentsRequest {
	return BraintreePaymentsApiGETBraintreeGatewayIdBraintreePaymentsRequest{
		ApiService:         a,
		ctx:                ctx,
		braintreeGatewayId: braintreeGatewayId,
	}
}

// Execute executes the request
func (a *BraintreePaymentsApiService) GETBraintreeGatewayIdBraintreePaymentsExecute(r BraintreePaymentsApiGETBraintreeGatewayIdBraintreePaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BraintreePaymentsApiService.GETBraintreeGatewayIdBraintreePayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_gateways/{braintreeGatewayId}/braintree_payments"
	localVarPath = strings.Replace(localVarPath, "{"+"braintreeGatewayId"+"}", url.PathEscape(parameterValueToString(r.braintreeGatewayId, "braintreeGatewayId")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type BraintreePaymentsApiGETBraintreePaymentsRequest struct {
	ctx        context.Context
	ApiService *BraintreePaymentsApiService
}

func (r BraintreePaymentsApiGETBraintreePaymentsRequest) Execute() (*GETBraintreePayments200Response, *http.Response, error) {
	return r.ApiService.GETBraintreePaymentsExecute(r)
}

/*
GETBraintreePayments List all braintree payments

List all braintree payments

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BraintreePaymentsApiGETBraintreePaymentsRequest
*/
func (a *BraintreePaymentsApiService) GETBraintreePayments(ctx context.Context) BraintreePaymentsApiGETBraintreePaymentsRequest {
	return BraintreePaymentsApiGETBraintreePaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETBraintreePayments200Response
func (a *BraintreePaymentsApiService) GETBraintreePaymentsExecute(r BraintreePaymentsApiGETBraintreePaymentsRequest) (*GETBraintreePayments200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETBraintreePayments200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BraintreePaymentsApiService.GETBraintreePayments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_payments"

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type BraintreePaymentsApiGETBraintreePaymentsBraintreePaymentIdRequest struct {
	ctx                context.Context
	ApiService         *BraintreePaymentsApiService
	braintreePaymentId interface{}
}

func (r BraintreePaymentsApiGETBraintreePaymentsBraintreePaymentIdRequest) Execute() (*GETBraintreePaymentsBraintreePaymentId200Response, *http.Response, error) {
	return r.ApiService.GETBraintreePaymentsBraintreePaymentIdExecute(r)
}

/*
GETBraintreePaymentsBraintreePaymentId Retrieve a braintree payment

Retrieve a braintree payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param braintreePaymentId The resource's id
	@return BraintreePaymentsApiGETBraintreePaymentsBraintreePaymentIdRequest
*/
func (a *BraintreePaymentsApiService) GETBraintreePaymentsBraintreePaymentId(ctx context.Context, braintreePaymentId interface{}) BraintreePaymentsApiGETBraintreePaymentsBraintreePaymentIdRequest {
	return BraintreePaymentsApiGETBraintreePaymentsBraintreePaymentIdRequest{
		ApiService:         a,
		ctx:                ctx,
		braintreePaymentId: braintreePaymentId,
	}
}

// Execute executes the request
//
//	@return GETBraintreePaymentsBraintreePaymentId200Response
func (a *BraintreePaymentsApiService) GETBraintreePaymentsBraintreePaymentIdExecute(r BraintreePaymentsApiGETBraintreePaymentsBraintreePaymentIdRequest) (*GETBraintreePaymentsBraintreePaymentId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETBraintreePaymentsBraintreePaymentId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BraintreePaymentsApiService.GETBraintreePaymentsBraintreePaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_payments/{braintreePaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"braintreePaymentId"+"}", url.PathEscape(parameterValueToString(r.braintreePaymentId, "braintreePaymentId")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type BraintreePaymentsApiPATCHBraintreePaymentsBraintreePaymentIdRequest struct {
	ctx                    context.Context
	ApiService             *BraintreePaymentsApiService
	braintreePaymentUpdate *BraintreePaymentUpdate
	braintreePaymentId     interface{}
}

func (r BraintreePaymentsApiPATCHBraintreePaymentsBraintreePaymentIdRequest) BraintreePaymentUpdate(braintreePaymentUpdate BraintreePaymentUpdate) BraintreePaymentsApiPATCHBraintreePaymentsBraintreePaymentIdRequest {
	r.braintreePaymentUpdate = &braintreePaymentUpdate
	return r
}

func (r BraintreePaymentsApiPATCHBraintreePaymentsBraintreePaymentIdRequest) Execute() (*PATCHBraintreePaymentsBraintreePaymentId200Response, *http.Response, error) {
	return r.ApiService.PATCHBraintreePaymentsBraintreePaymentIdExecute(r)
}

/*
PATCHBraintreePaymentsBraintreePaymentId Update a braintree payment

Update a braintree payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param braintreePaymentId The resource's id
	@return BraintreePaymentsApiPATCHBraintreePaymentsBraintreePaymentIdRequest
*/
func (a *BraintreePaymentsApiService) PATCHBraintreePaymentsBraintreePaymentId(ctx context.Context, braintreePaymentId interface{}) BraintreePaymentsApiPATCHBraintreePaymentsBraintreePaymentIdRequest {
	return BraintreePaymentsApiPATCHBraintreePaymentsBraintreePaymentIdRequest{
		ApiService:         a,
		ctx:                ctx,
		braintreePaymentId: braintreePaymentId,
	}
}

// Execute executes the request
//
//	@return PATCHBraintreePaymentsBraintreePaymentId200Response
func (a *BraintreePaymentsApiService) PATCHBraintreePaymentsBraintreePaymentIdExecute(r BraintreePaymentsApiPATCHBraintreePaymentsBraintreePaymentIdRequest) (*PATCHBraintreePaymentsBraintreePaymentId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHBraintreePaymentsBraintreePaymentId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BraintreePaymentsApiService.PATCHBraintreePaymentsBraintreePaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_payments/{braintreePaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"braintreePaymentId"+"}", url.PathEscape(parameterValueToString(r.braintreePaymentId, "braintreePaymentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.braintreePaymentUpdate == nil {
		return localVarReturnValue, nil, reportError("braintreePaymentUpdate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json"}

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
	// body params
	localVarPostBody = r.braintreePaymentUpdate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type BraintreePaymentsApiPOSTBraintreePaymentsRequest struct {
	ctx                    context.Context
	ApiService             *BraintreePaymentsApiService
	braintreePaymentCreate *BraintreePaymentCreate
}

func (r BraintreePaymentsApiPOSTBraintreePaymentsRequest) BraintreePaymentCreate(braintreePaymentCreate BraintreePaymentCreate) BraintreePaymentsApiPOSTBraintreePaymentsRequest {
	r.braintreePaymentCreate = &braintreePaymentCreate
	return r
}

func (r BraintreePaymentsApiPOSTBraintreePaymentsRequest) Execute() (*POSTBraintreePayments201Response, *http.Response, error) {
	return r.ApiService.POSTBraintreePaymentsExecute(r)
}

/*
POSTBraintreePayments Create a braintree payment

Create a braintree payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BraintreePaymentsApiPOSTBraintreePaymentsRequest
*/
func (a *BraintreePaymentsApiService) POSTBraintreePayments(ctx context.Context) BraintreePaymentsApiPOSTBraintreePaymentsRequest {
	return BraintreePaymentsApiPOSTBraintreePaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTBraintreePayments201Response
func (a *BraintreePaymentsApiService) POSTBraintreePaymentsExecute(r BraintreePaymentsApiPOSTBraintreePaymentsRequest) (*POSTBraintreePayments201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTBraintreePayments201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BraintreePaymentsApiService.POSTBraintreePayments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_payments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.braintreePaymentCreate == nil {
		return localVarReturnValue, nil, reportError("braintreePaymentCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json"}

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
	// body params
	localVarPostBody = r.braintreePaymentCreate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
