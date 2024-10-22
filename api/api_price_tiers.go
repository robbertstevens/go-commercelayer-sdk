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

// PriceTiersApiService PriceTiersApi service
type PriceTiersApiService service

type PriceTiersApiGETPriceIdPriceTiersRequest struct {
	ctx        context.Context
	ApiService *PriceTiersApiService
	priceId    interface{}
}

func (r PriceTiersApiGETPriceIdPriceTiersRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETPriceIdPriceTiersExecute(r)
}

/*
GETPriceIdPriceTiers Retrieve the price tiers associated to the price

Retrieve the price tiers associated to the price

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param priceId The resource's id
	@return PriceTiersApiGETPriceIdPriceTiersRequest
*/
func (a *PriceTiersApiService) GETPriceIdPriceTiers(ctx context.Context, priceId interface{}) PriceTiersApiGETPriceIdPriceTiersRequest {
	return PriceTiersApiGETPriceIdPriceTiersRequest{
		ApiService: a,
		ctx:        ctx,
		priceId:    priceId,
	}
}

// Execute executes the request
func (a *PriceTiersApiService) GETPriceIdPriceTiersExecute(r PriceTiersApiGETPriceIdPriceTiersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceTiersApiService.GETPriceIdPriceTiers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/prices/{priceId}/price_tiers"
	localVarPath = strings.Replace(localVarPath, "{"+"priceId"+"}", url.PathEscape(parameterValueToString(r.priceId, "priceId")), -1)

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

type PriceTiersApiGETPriceTiersRequest struct {
	ctx        context.Context
	ApiService *PriceTiersApiService
}

func (r PriceTiersApiGETPriceTiersRequest) Execute() (*GETPriceTiers200Response, *http.Response, error) {
	return r.ApiService.GETPriceTiersExecute(r)
}

/*
GETPriceTiers List all price tiers

List all price tiers

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PriceTiersApiGETPriceTiersRequest
*/
func (a *PriceTiersApiService) GETPriceTiers(ctx context.Context) PriceTiersApiGETPriceTiersRequest {
	return PriceTiersApiGETPriceTiersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETPriceTiers200Response
func (a *PriceTiersApiService) GETPriceTiersExecute(r PriceTiersApiGETPriceTiersRequest) (*GETPriceTiers200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETPriceTiers200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceTiersApiService.GETPriceTiers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_tiers"

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

type PriceTiersApiGETPriceTiersPriceTierIdRequest struct {
	ctx         context.Context
	ApiService  *PriceTiersApiService
	priceTierId interface{}
}

func (r PriceTiersApiGETPriceTiersPriceTierIdRequest) Execute() (*GETPriceTiersPriceTierId200Response, *http.Response, error) {
	return r.ApiService.GETPriceTiersPriceTierIdExecute(r)
}

/*
GETPriceTiersPriceTierId Retrieve a price tier

Retrieve a price tier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param priceTierId The resource's id
	@return PriceTiersApiGETPriceTiersPriceTierIdRequest
*/
func (a *PriceTiersApiService) GETPriceTiersPriceTierId(ctx context.Context, priceTierId interface{}) PriceTiersApiGETPriceTiersPriceTierIdRequest {
	return PriceTiersApiGETPriceTiersPriceTierIdRequest{
		ApiService:  a,
		ctx:         ctx,
		priceTierId: priceTierId,
	}
}

// Execute executes the request
//
//	@return GETPriceTiersPriceTierId200Response
func (a *PriceTiersApiService) GETPriceTiersPriceTierIdExecute(r PriceTiersApiGETPriceTiersPriceTierIdRequest) (*GETPriceTiersPriceTierId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETPriceTiersPriceTierId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceTiersApiService.GETPriceTiersPriceTierId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_tiers/{priceTierId}"
	localVarPath = strings.Replace(localVarPath, "{"+"priceTierId"+"}", url.PathEscape(parameterValueToString(r.priceTierId, "priceTierId")), -1)

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
