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

// PriceListSchedulersApiService PriceListSchedulersApi service
type PriceListSchedulersApiService service

type PriceListSchedulersApiDELETEPriceListSchedulersPriceListSchedulerIdRequest struct {
	ctx                  context.Context
	ApiService           *PriceListSchedulersApiService
	priceListSchedulerId interface{}
}

func (r PriceListSchedulersApiDELETEPriceListSchedulersPriceListSchedulerIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEPriceListSchedulersPriceListSchedulerIdExecute(r)
}

/*
DELETEPriceListSchedulersPriceListSchedulerId Delete a price list scheduler

Delete a price list scheduler

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param priceListSchedulerId The resource's id
	@return PriceListSchedulersApiDELETEPriceListSchedulersPriceListSchedulerIdRequest
*/
func (a *PriceListSchedulersApiService) DELETEPriceListSchedulersPriceListSchedulerId(ctx context.Context, priceListSchedulerId interface{}) PriceListSchedulersApiDELETEPriceListSchedulersPriceListSchedulerIdRequest {
	return PriceListSchedulersApiDELETEPriceListSchedulersPriceListSchedulerIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		priceListSchedulerId: priceListSchedulerId,
	}
}

// Execute executes the request
func (a *PriceListSchedulersApiService) DELETEPriceListSchedulersPriceListSchedulerIdExecute(r PriceListSchedulersApiDELETEPriceListSchedulersPriceListSchedulerIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceListSchedulersApiService.DELETEPriceListSchedulersPriceListSchedulerId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_list_schedulers/{priceListSchedulerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"priceListSchedulerId"+"}", url.PathEscape(parameterValueToString(r.priceListSchedulerId, "priceListSchedulerId")), -1)

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

type PriceListSchedulersApiGETMarketIdPriceListSchedulersRequest struct {
	ctx        context.Context
	ApiService *PriceListSchedulersApiService
	marketId   interface{}
}

func (r PriceListSchedulersApiGETMarketIdPriceListSchedulersRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETMarketIdPriceListSchedulersExecute(r)
}

/*
GETMarketIdPriceListSchedulers Retrieve the price list schedulers associated to the market

Retrieve the price list schedulers associated to the market

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param marketId The resource's id
	@return PriceListSchedulersApiGETMarketIdPriceListSchedulersRequest
*/
func (a *PriceListSchedulersApiService) GETMarketIdPriceListSchedulers(ctx context.Context, marketId interface{}) PriceListSchedulersApiGETMarketIdPriceListSchedulersRequest {
	return PriceListSchedulersApiGETMarketIdPriceListSchedulersRequest{
		ApiService: a,
		ctx:        ctx,
		marketId:   marketId,
	}
}

// Execute executes the request
func (a *PriceListSchedulersApiService) GETMarketIdPriceListSchedulersExecute(r PriceListSchedulersApiGETMarketIdPriceListSchedulersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceListSchedulersApiService.GETMarketIdPriceListSchedulers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/markets/{marketId}/price_list_schedulers"
	localVarPath = strings.Replace(localVarPath, "{"+"marketId"+"}", url.PathEscape(parameterValueToString(r.marketId, "marketId")), -1)

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

type PriceListSchedulersApiGETPriceListIdPriceListSchedulersRequest struct {
	ctx         context.Context
	ApiService  *PriceListSchedulersApiService
	priceListId interface{}
}

func (r PriceListSchedulersApiGETPriceListIdPriceListSchedulersRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETPriceListIdPriceListSchedulersExecute(r)
}

/*
GETPriceListIdPriceListSchedulers Retrieve the price list schedulers associated to the price list

Retrieve the price list schedulers associated to the price list

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param priceListId The resource's id
	@return PriceListSchedulersApiGETPriceListIdPriceListSchedulersRequest
*/
func (a *PriceListSchedulersApiService) GETPriceListIdPriceListSchedulers(ctx context.Context, priceListId interface{}) PriceListSchedulersApiGETPriceListIdPriceListSchedulersRequest {
	return PriceListSchedulersApiGETPriceListIdPriceListSchedulersRequest{
		ApiService:  a,
		ctx:         ctx,
		priceListId: priceListId,
	}
}

// Execute executes the request
func (a *PriceListSchedulersApiService) GETPriceListIdPriceListSchedulersExecute(r PriceListSchedulersApiGETPriceListIdPriceListSchedulersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceListSchedulersApiService.GETPriceListIdPriceListSchedulers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_lists/{priceListId}/price_list_schedulers"
	localVarPath = strings.Replace(localVarPath, "{"+"priceListId"+"}", url.PathEscape(parameterValueToString(r.priceListId, "priceListId")), -1)

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

type PriceListSchedulersApiGETPriceListSchedulersRequest struct {
	ctx        context.Context
	ApiService *PriceListSchedulersApiService
}

func (r PriceListSchedulersApiGETPriceListSchedulersRequest) Execute() (*GETPriceListSchedulers200Response, *http.Response, error) {
	return r.ApiService.GETPriceListSchedulersExecute(r)
}

/*
GETPriceListSchedulers List all price list schedulers

List all price list schedulers

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PriceListSchedulersApiGETPriceListSchedulersRequest
*/
func (a *PriceListSchedulersApiService) GETPriceListSchedulers(ctx context.Context) PriceListSchedulersApiGETPriceListSchedulersRequest {
	return PriceListSchedulersApiGETPriceListSchedulersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETPriceListSchedulers200Response
func (a *PriceListSchedulersApiService) GETPriceListSchedulersExecute(r PriceListSchedulersApiGETPriceListSchedulersRequest) (*GETPriceListSchedulers200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETPriceListSchedulers200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceListSchedulersApiService.GETPriceListSchedulers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_list_schedulers"

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

type PriceListSchedulersApiGETPriceListSchedulersPriceListSchedulerIdRequest struct {
	ctx                  context.Context
	ApiService           *PriceListSchedulersApiService
	priceListSchedulerId interface{}
}

func (r PriceListSchedulersApiGETPriceListSchedulersPriceListSchedulerIdRequest) Execute() (*GETPriceListSchedulersPriceListSchedulerId200Response, *http.Response, error) {
	return r.ApiService.GETPriceListSchedulersPriceListSchedulerIdExecute(r)
}

/*
GETPriceListSchedulersPriceListSchedulerId Retrieve a price list scheduler

Retrieve a price list scheduler

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param priceListSchedulerId The resource's id
	@return PriceListSchedulersApiGETPriceListSchedulersPriceListSchedulerIdRequest
*/
func (a *PriceListSchedulersApiService) GETPriceListSchedulersPriceListSchedulerId(ctx context.Context, priceListSchedulerId interface{}) PriceListSchedulersApiGETPriceListSchedulersPriceListSchedulerIdRequest {
	return PriceListSchedulersApiGETPriceListSchedulersPriceListSchedulerIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		priceListSchedulerId: priceListSchedulerId,
	}
}

// Execute executes the request
//
//	@return GETPriceListSchedulersPriceListSchedulerId200Response
func (a *PriceListSchedulersApiService) GETPriceListSchedulersPriceListSchedulerIdExecute(r PriceListSchedulersApiGETPriceListSchedulersPriceListSchedulerIdRequest) (*GETPriceListSchedulersPriceListSchedulerId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETPriceListSchedulersPriceListSchedulerId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceListSchedulersApiService.GETPriceListSchedulersPriceListSchedulerId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_list_schedulers/{priceListSchedulerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"priceListSchedulerId"+"}", url.PathEscape(parameterValueToString(r.priceListSchedulerId, "priceListSchedulerId")), -1)

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

type PriceListSchedulersApiPATCHPriceListSchedulersPriceListSchedulerIdRequest struct {
	ctx                      context.Context
	ApiService               *PriceListSchedulersApiService
	priceListSchedulerUpdate *PriceListSchedulerUpdate
	priceListSchedulerId     interface{}
}

func (r PriceListSchedulersApiPATCHPriceListSchedulersPriceListSchedulerIdRequest) PriceListSchedulerUpdate(priceListSchedulerUpdate PriceListSchedulerUpdate) PriceListSchedulersApiPATCHPriceListSchedulersPriceListSchedulerIdRequest {
	r.priceListSchedulerUpdate = &priceListSchedulerUpdate
	return r
}

func (r PriceListSchedulersApiPATCHPriceListSchedulersPriceListSchedulerIdRequest) Execute() (*PATCHPriceListSchedulersPriceListSchedulerId200Response, *http.Response, error) {
	return r.ApiService.PATCHPriceListSchedulersPriceListSchedulerIdExecute(r)
}

/*
PATCHPriceListSchedulersPriceListSchedulerId Update a price list scheduler

Update a price list scheduler

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param priceListSchedulerId The resource's id
	@return PriceListSchedulersApiPATCHPriceListSchedulersPriceListSchedulerIdRequest
*/
func (a *PriceListSchedulersApiService) PATCHPriceListSchedulersPriceListSchedulerId(ctx context.Context, priceListSchedulerId interface{}) PriceListSchedulersApiPATCHPriceListSchedulersPriceListSchedulerIdRequest {
	return PriceListSchedulersApiPATCHPriceListSchedulersPriceListSchedulerIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		priceListSchedulerId: priceListSchedulerId,
	}
}

// Execute executes the request
//
//	@return PATCHPriceListSchedulersPriceListSchedulerId200Response
func (a *PriceListSchedulersApiService) PATCHPriceListSchedulersPriceListSchedulerIdExecute(r PriceListSchedulersApiPATCHPriceListSchedulersPriceListSchedulerIdRequest) (*PATCHPriceListSchedulersPriceListSchedulerId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHPriceListSchedulersPriceListSchedulerId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceListSchedulersApiService.PATCHPriceListSchedulersPriceListSchedulerId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_list_schedulers/{priceListSchedulerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"priceListSchedulerId"+"}", url.PathEscape(parameterValueToString(r.priceListSchedulerId, "priceListSchedulerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.priceListSchedulerUpdate == nil {
		return localVarReturnValue, nil, reportError("priceListSchedulerUpdate is required and must be specified")
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
	localVarPostBody = r.priceListSchedulerUpdate
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

type PriceListSchedulersApiPOSTPriceListSchedulersRequest struct {
	ctx                      context.Context
	ApiService               *PriceListSchedulersApiService
	priceListSchedulerCreate *PriceListSchedulerCreate
}

func (r PriceListSchedulersApiPOSTPriceListSchedulersRequest) PriceListSchedulerCreate(priceListSchedulerCreate PriceListSchedulerCreate) PriceListSchedulersApiPOSTPriceListSchedulersRequest {
	r.priceListSchedulerCreate = &priceListSchedulerCreate
	return r
}

func (r PriceListSchedulersApiPOSTPriceListSchedulersRequest) Execute() (*POSTPriceListSchedulers201Response, *http.Response, error) {
	return r.ApiService.POSTPriceListSchedulersExecute(r)
}

/*
POSTPriceListSchedulers Create a price list scheduler

Create a price list scheduler

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PriceListSchedulersApiPOSTPriceListSchedulersRequest
*/
func (a *PriceListSchedulersApiService) POSTPriceListSchedulers(ctx context.Context) PriceListSchedulersApiPOSTPriceListSchedulersRequest {
	return PriceListSchedulersApiPOSTPriceListSchedulersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTPriceListSchedulers201Response
func (a *PriceListSchedulersApiService) POSTPriceListSchedulersExecute(r PriceListSchedulersApiPOSTPriceListSchedulersRequest) (*POSTPriceListSchedulers201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTPriceListSchedulers201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceListSchedulersApiService.POSTPriceListSchedulers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_list_schedulers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.priceListSchedulerCreate == nil {
		return localVarReturnValue, nil, reportError("priceListSchedulerCreate is required and must be specified")
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
	localVarPostBody = r.priceListSchedulerCreate
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
