/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
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

// ShippingWeightTiersApiService ShippingWeightTiersApi service
type ShippingWeightTiersApiService service

type ShippingWeightTiersApiDELETEShippingWeightTiersShippingWeightTierIdRequest struct {
	ctx                  context.Context
	ApiService           *ShippingWeightTiersApiService
	shippingWeightTierId interface{}
}

func (r ShippingWeightTiersApiDELETEShippingWeightTiersShippingWeightTierIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEShippingWeightTiersShippingWeightTierIdExecute(r)
}

/*
DELETEShippingWeightTiersShippingWeightTierId Delete a shipping weight tier

Delete a shipping weight tier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shippingWeightTierId The resource's id
	@return ShippingWeightTiersApiDELETEShippingWeightTiersShippingWeightTierIdRequest
*/
func (a *ShippingWeightTiersApiService) DELETEShippingWeightTiersShippingWeightTierId(ctx context.Context, shippingWeightTierId interface{}) ShippingWeightTiersApiDELETEShippingWeightTiersShippingWeightTierIdRequest {
	return ShippingWeightTiersApiDELETEShippingWeightTiersShippingWeightTierIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		shippingWeightTierId: shippingWeightTierId,
	}
}

// Execute executes the request
func (a *ShippingWeightTiersApiService) DELETEShippingWeightTiersShippingWeightTierIdExecute(r ShippingWeightTiersApiDELETEShippingWeightTiersShippingWeightTierIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingWeightTiersApiService.DELETEShippingWeightTiersShippingWeightTierId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_weight_tiers/{shippingWeightTierId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingWeightTierId"+"}", url.PathEscape(parameterValueToString(r.shippingWeightTierId, "shippingWeightTierId")), -1)

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

type ShippingWeightTiersApiGETShippingMethodIdShippingWeightTiersRequest struct {
	ctx              context.Context
	ApiService       *ShippingWeightTiersApiService
	shippingMethodId interface{}
}

func (r ShippingWeightTiersApiGETShippingMethodIdShippingWeightTiersRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShippingMethodIdShippingWeightTiersExecute(r)
}

/*
GETShippingMethodIdShippingWeightTiers Retrieve the shipping weight tiers associated to the shipping method

Retrieve the shipping weight tiers associated to the shipping method

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shippingMethodId The resource's id
	@return ShippingWeightTiersApiGETShippingMethodIdShippingWeightTiersRequest
*/
func (a *ShippingWeightTiersApiService) GETShippingMethodIdShippingWeightTiers(ctx context.Context, shippingMethodId interface{}) ShippingWeightTiersApiGETShippingMethodIdShippingWeightTiersRequest {
	return ShippingWeightTiersApiGETShippingMethodIdShippingWeightTiersRequest{
		ApiService:       a,
		ctx:              ctx,
		shippingMethodId: shippingMethodId,
	}
}

// Execute executes the request
func (a *ShippingWeightTiersApiService) GETShippingMethodIdShippingWeightTiersExecute(r ShippingWeightTiersApiGETShippingMethodIdShippingWeightTiersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingWeightTiersApiService.GETShippingMethodIdShippingWeightTiers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_methods/{shippingMethodId}/shipping_weight_tiers"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingMethodId"+"}", url.PathEscape(parameterValueToString(r.shippingMethodId, "shippingMethodId")), -1)

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

type ShippingWeightTiersApiGETShippingWeightTiersRequest struct {
	ctx        context.Context
	ApiService *ShippingWeightTiersApiService
}

func (r ShippingWeightTiersApiGETShippingWeightTiersRequest) Execute() (*GETShippingWeightTiers200Response, *http.Response, error) {
	return r.ApiService.GETShippingWeightTiersExecute(r)
}

/*
GETShippingWeightTiers List all shipping weight tiers

List all shipping weight tiers

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ShippingWeightTiersApiGETShippingWeightTiersRequest
*/
func (a *ShippingWeightTiersApiService) GETShippingWeightTiers(ctx context.Context) ShippingWeightTiersApiGETShippingWeightTiersRequest {
	return ShippingWeightTiersApiGETShippingWeightTiersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETShippingWeightTiers200Response
func (a *ShippingWeightTiersApiService) GETShippingWeightTiersExecute(r ShippingWeightTiersApiGETShippingWeightTiersRequest) (*GETShippingWeightTiers200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETShippingWeightTiers200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingWeightTiersApiService.GETShippingWeightTiers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_weight_tiers"

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

type ShippingWeightTiersApiGETShippingWeightTiersShippingWeightTierIdRequest struct {
	ctx                  context.Context
	ApiService           *ShippingWeightTiersApiService
	shippingWeightTierId interface{}
}

func (r ShippingWeightTiersApiGETShippingWeightTiersShippingWeightTierIdRequest) Execute() (*GETShippingWeightTiersShippingWeightTierId200Response, *http.Response, error) {
	return r.ApiService.GETShippingWeightTiersShippingWeightTierIdExecute(r)
}

/*
GETShippingWeightTiersShippingWeightTierId Retrieve a shipping weight tier

Retrieve a shipping weight tier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shippingWeightTierId The resource's id
	@return ShippingWeightTiersApiGETShippingWeightTiersShippingWeightTierIdRequest
*/
func (a *ShippingWeightTiersApiService) GETShippingWeightTiersShippingWeightTierId(ctx context.Context, shippingWeightTierId interface{}) ShippingWeightTiersApiGETShippingWeightTiersShippingWeightTierIdRequest {
	return ShippingWeightTiersApiGETShippingWeightTiersShippingWeightTierIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		shippingWeightTierId: shippingWeightTierId,
	}
}

// Execute executes the request
//
//	@return GETShippingWeightTiersShippingWeightTierId200Response
func (a *ShippingWeightTiersApiService) GETShippingWeightTiersShippingWeightTierIdExecute(r ShippingWeightTiersApiGETShippingWeightTiersShippingWeightTierIdRequest) (*GETShippingWeightTiersShippingWeightTierId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETShippingWeightTiersShippingWeightTierId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingWeightTiersApiService.GETShippingWeightTiersShippingWeightTierId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_weight_tiers/{shippingWeightTierId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingWeightTierId"+"}", url.PathEscape(parameterValueToString(r.shippingWeightTierId, "shippingWeightTierId")), -1)

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

type ShippingWeightTiersApiPATCHShippingWeightTiersShippingWeightTierIdRequest struct {
	ctx                                                 context.Context
	ApiService                                          *ShippingWeightTiersApiService
	pATCHShippingWeightTiersShippingWeightTierIdRequest *PATCHShippingWeightTiersShippingWeightTierIdRequest
	shippingWeightTierId                                interface{}
}

func (r ShippingWeightTiersApiPATCHShippingWeightTiersShippingWeightTierIdRequest) PATCHShippingWeightTiersShippingWeightTierIdRequest(pATCHShippingWeightTiersShippingWeightTierIdRequest PATCHShippingWeightTiersShippingWeightTierIdRequest) ShippingWeightTiersApiPATCHShippingWeightTiersShippingWeightTierIdRequest {
	r.pATCHShippingWeightTiersShippingWeightTierIdRequest = &pATCHShippingWeightTiersShippingWeightTierIdRequest
	return r
}

func (r ShippingWeightTiersApiPATCHShippingWeightTiersShippingWeightTierIdRequest) Execute() (*PATCHShippingWeightTiersShippingWeightTierId200Response, *http.Response, error) {
	return r.ApiService.PATCHShippingWeightTiersShippingWeightTierIdExecute(r)
}

/*
PATCHShippingWeightTiersShippingWeightTierId Update a shipping weight tier

Update a shipping weight tier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shippingWeightTierId The resource's id
	@return ShippingWeightTiersApiPATCHShippingWeightTiersShippingWeightTierIdRequest
*/
func (a *ShippingWeightTiersApiService) PATCHShippingWeightTiersShippingWeightTierId(ctx context.Context, shippingWeightTierId interface{}) ShippingWeightTiersApiPATCHShippingWeightTiersShippingWeightTierIdRequest {
	return ShippingWeightTiersApiPATCHShippingWeightTiersShippingWeightTierIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		shippingWeightTierId: shippingWeightTierId,
	}
}

// Execute executes the request
//
//	@return PATCHShippingWeightTiersShippingWeightTierId200Response
func (a *ShippingWeightTiersApiService) PATCHShippingWeightTiersShippingWeightTierIdExecute(r ShippingWeightTiersApiPATCHShippingWeightTiersShippingWeightTierIdRequest) (*PATCHShippingWeightTiersShippingWeightTierId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHShippingWeightTiersShippingWeightTierId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingWeightTiersApiService.PATCHShippingWeightTiersShippingWeightTierId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_weight_tiers/{shippingWeightTierId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingWeightTierId"+"}", url.PathEscape(parameterValueToString(r.shippingWeightTierId, "shippingWeightTierId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pATCHShippingWeightTiersShippingWeightTierIdRequest == nil {
		return localVarReturnValue, nil, reportError("pATCHShippingWeightTiersShippingWeightTierIdRequest is required and must be specified")
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
	localVarPostBody = r.pATCHShippingWeightTiersShippingWeightTierIdRequest
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

type ShippingWeightTiersApiPOSTShippingWeightTiersRequest struct {
	ctx                            context.Context
	ApiService                     *ShippingWeightTiersApiService
	pOSTShippingWeightTiersRequest *POSTShippingWeightTiersRequest
}

func (r ShippingWeightTiersApiPOSTShippingWeightTiersRequest) POSTShippingWeightTiersRequest(pOSTShippingWeightTiersRequest POSTShippingWeightTiersRequest) ShippingWeightTiersApiPOSTShippingWeightTiersRequest {
	r.pOSTShippingWeightTiersRequest = &pOSTShippingWeightTiersRequest
	return r
}

func (r ShippingWeightTiersApiPOSTShippingWeightTiersRequest) Execute() (*POSTShippingWeightTiers201Response, *http.Response, error) {
	return r.ApiService.POSTShippingWeightTiersExecute(r)
}

/*
POSTShippingWeightTiers Create a shipping weight tier

Create a shipping weight tier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ShippingWeightTiersApiPOSTShippingWeightTiersRequest
*/
func (a *ShippingWeightTiersApiService) POSTShippingWeightTiers(ctx context.Context) ShippingWeightTiersApiPOSTShippingWeightTiersRequest {
	return ShippingWeightTiersApiPOSTShippingWeightTiersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTShippingWeightTiers201Response
func (a *ShippingWeightTiersApiService) POSTShippingWeightTiersExecute(r ShippingWeightTiersApiPOSTShippingWeightTiersRequest) (*POSTShippingWeightTiers201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTShippingWeightTiers201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingWeightTiersApiService.POSTShippingWeightTiers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_weight_tiers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pOSTShippingWeightTiersRequest == nil {
		return localVarReturnValue, nil, reportError("pOSTShippingWeightTiersRequest is required and must be specified")
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
	localVarPostBody = r.pOSTShippingWeightTiersRequest
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
