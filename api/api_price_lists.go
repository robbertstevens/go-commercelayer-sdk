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

// PriceListsApiService PriceListsApi service
type PriceListsApiService service

type PriceListsApiDELETEPriceListsPriceListIdRequest struct {
	ctx         context.Context
	ApiService  *PriceListsApiService
	priceListId interface{}
}

func (r PriceListsApiDELETEPriceListsPriceListIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEPriceListsPriceListIdExecute(r)
}

/*
DELETEPriceListsPriceListId Delete a price list

Delete a price list

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param priceListId The resource's id
	@return PriceListsApiDELETEPriceListsPriceListIdRequest
*/
func (a *PriceListsApiService) DELETEPriceListsPriceListId(ctx context.Context, priceListId interface{}) PriceListsApiDELETEPriceListsPriceListIdRequest {
	return PriceListsApiDELETEPriceListsPriceListIdRequest{
		ApiService:  a,
		ctx:         ctx,
		priceListId: priceListId,
	}
}

// Execute executes the request
func (a *PriceListsApiService) DELETEPriceListsPriceListIdExecute(r PriceListsApiDELETEPriceListsPriceListIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceListsApiService.DELETEPriceListsPriceListId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_lists/{priceListId}"
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

type PriceListsApiGETMarketIdBasePriceListRequest struct {
	ctx        context.Context
	ApiService *PriceListsApiService
	marketId   interface{}
}

func (r PriceListsApiGETMarketIdBasePriceListRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETMarketIdBasePriceListExecute(r)
}

/*
GETMarketIdBasePriceList Retrieve the base price list associated to the market

Retrieve the base price list associated to the market

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param marketId The resource's id
	@return PriceListsApiGETMarketIdBasePriceListRequest
*/
func (a *PriceListsApiService) GETMarketIdBasePriceList(ctx context.Context, marketId interface{}) PriceListsApiGETMarketIdBasePriceListRequest {
	return PriceListsApiGETMarketIdBasePriceListRequest{
		ApiService: a,
		ctx:        ctx,
		marketId:   marketId,
	}
}

// Execute executes the request
func (a *PriceListsApiService) GETMarketIdBasePriceListExecute(r PriceListsApiGETMarketIdBasePriceListRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceListsApiService.GETMarketIdBasePriceList")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/markets/{marketId}/base_price_list"
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

type PriceListsApiGETMarketIdPriceListRequest struct {
	ctx        context.Context
	ApiService *PriceListsApiService
	marketId   interface{}
}

func (r PriceListsApiGETMarketIdPriceListRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETMarketIdPriceListExecute(r)
}

/*
GETMarketIdPriceList Retrieve the price list associated to the market

Retrieve the price list associated to the market

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param marketId The resource's id
	@return PriceListsApiGETMarketIdPriceListRequest
*/
func (a *PriceListsApiService) GETMarketIdPriceList(ctx context.Context, marketId interface{}) PriceListsApiGETMarketIdPriceListRequest {
	return PriceListsApiGETMarketIdPriceListRequest{
		ApiService: a,
		ctx:        ctx,
		marketId:   marketId,
	}
}

// Execute executes the request
func (a *PriceListsApiService) GETMarketIdPriceListExecute(r PriceListsApiGETMarketIdPriceListRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceListsApiService.GETMarketIdPriceList")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/markets/{marketId}/price_list"
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

type PriceListsApiGETPriceIdPriceListRequest struct {
	ctx        context.Context
	ApiService *PriceListsApiService
	priceId    interface{}
}

func (r PriceListsApiGETPriceIdPriceListRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETPriceIdPriceListExecute(r)
}

/*
GETPriceIdPriceList Retrieve the price list associated to the price

Retrieve the price list associated to the price

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param priceId The resource's id
	@return PriceListsApiGETPriceIdPriceListRequest
*/
func (a *PriceListsApiService) GETPriceIdPriceList(ctx context.Context, priceId interface{}) PriceListsApiGETPriceIdPriceListRequest {
	return PriceListsApiGETPriceIdPriceListRequest{
		ApiService: a,
		ctx:        ctx,
		priceId:    priceId,
	}
}

// Execute executes the request
func (a *PriceListsApiService) GETPriceIdPriceListExecute(r PriceListsApiGETPriceIdPriceListRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceListsApiService.GETPriceIdPriceList")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/prices/{priceId}/price_list"
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

type PriceListsApiGETPriceListSchedulerIdPriceListRequest struct {
	ctx                  context.Context
	ApiService           *PriceListsApiService
	priceListSchedulerId interface{}
}

func (r PriceListsApiGETPriceListSchedulerIdPriceListRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETPriceListSchedulerIdPriceListExecute(r)
}

/*
GETPriceListSchedulerIdPriceList Retrieve the price list associated to the price list scheduler

Retrieve the price list associated to the price list scheduler

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param priceListSchedulerId The resource's id
	@return PriceListsApiGETPriceListSchedulerIdPriceListRequest
*/
func (a *PriceListsApiService) GETPriceListSchedulerIdPriceList(ctx context.Context, priceListSchedulerId interface{}) PriceListsApiGETPriceListSchedulerIdPriceListRequest {
	return PriceListsApiGETPriceListSchedulerIdPriceListRequest{
		ApiService:           a,
		ctx:                  ctx,
		priceListSchedulerId: priceListSchedulerId,
	}
}

// Execute executes the request
func (a *PriceListsApiService) GETPriceListSchedulerIdPriceListExecute(r PriceListsApiGETPriceListSchedulerIdPriceListRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceListsApiService.GETPriceListSchedulerIdPriceList")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_list_schedulers/{priceListSchedulerId}/price_list"
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

type PriceListsApiGETPriceListsRequest struct {
	ctx        context.Context
	ApiService *PriceListsApiService
}

func (r PriceListsApiGETPriceListsRequest) Execute() (*GETPriceLists200Response, *http.Response, error) {
	return r.ApiService.GETPriceListsExecute(r)
}

/*
GETPriceLists List all price lists

List all price lists

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PriceListsApiGETPriceListsRequest
*/
func (a *PriceListsApiService) GETPriceLists(ctx context.Context) PriceListsApiGETPriceListsRequest {
	return PriceListsApiGETPriceListsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETPriceLists200Response
func (a *PriceListsApiService) GETPriceListsExecute(r PriceListsApiGETPriceListsRequest) (*GETPriceLists200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETPriceLists200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceListsApiService.GETPriceLists")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_lists"

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

type PriceListsApiGETPriceListsPriceListIdRequest struct {
	ctx         context.Context
	ApiService  *PriceListsApiService
	priceListId interface{}
}

func (r PriceListsApiGETPriceListsPriceListIdRequest) Execute() (*GETPriceListsPriceListId200Response, *http.Response, error) {
	return r.ApiService.GETPriceListsPriceListIdExecute(r)
}

/*
GETPriceListsPriceListId Retrieve a price list

Retrieve a price list

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param priceListId The resource's id
	@return PriceListsApiGETPriceListsPriceListIdRequest
*/
func (a *PriceListsApiService) GETPriceListsPriceListId(ctx context.Context, priceListId interface{}) PriceListsApiGETPriceListsPriceListIdRequest {
	return PriceListsApiGETPriceListsPriceListIdRequest{
		ApiService:  a,
		ctx:         ctx,
		priceListId: priceListId,
	}
}

// Execute executes the request
//
//	@return GETPriceListsPriceListId200Response
func (a *PriceListsApiService) GETPriceListsPriceListIdExecute(r PriceListsApiGETPriceListsPriceListIdRequest) (*GETPriceListsPriceListId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETPriceListsPriceListId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceListsApiService.GETPriceListsPriceListId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_lists/{priceListId}"
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

type PriceListsApiPATCHPriceListsPriceListIdRequest struct {
	ctx             context.Context
	ApiService      *PriceListsApiService
	priceListUpdate *PriceListUpdate
	priceListId     interface{}
}

func (r PriceListsApiPATCHPriceListsPriceListIdRequest) PriceListUpdate(priceListUpdate PriceListUpdate) PriceListsApiPATCHPriceListsPriceListIdRequest {
	r.priceListUpdate = &priceListUpdate
	return r
}

func (r PriceListsApiPATCHPriceListsPriceListIdRequest) Execute() (*PATCHPriceListsPriceListId200Response, *http.Response, error) {
	return r.ApiService.PATCHPriceListsPriceListIdExecute(r)
}

/*
PATCHPriceListsPriceListId Update a price list

Update a price list

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param priceListId The resource's id
	@return PriceListsApiPATCHPriceListsPriceListIdRequest
*/
func (a *PriceListsApiService) PATCHPriceListsPriceListId(ctx context.Context, priceListId interface{}) PriceListsApiPATCHPriceListsPriceListIdRequest {
	return PriceListsApiPATCHPriceListsPriceListIdRequest{
		ApiService:  a,
		ctx:         ctx,
		priceListId: priceListId,
	}
}

// Execute executes the request
//
//	@return PATCHPriceListsPriceListId200Response
func (a *PriceListsApiService) PATCHPriceListsPriceListIdExecute(r PriceListsApiPATCHPriceListsPriceListIdRequest) (*PATCHPriceListsPriceListId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHPriceListsPriceListId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceListsApiService.PATCHPriceListsPriceListId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_lists/{priceListId}"
	localVarPath = strings.Replace(localVarPath, "{"+"priceListId"+"}", url.PathEscape(parameterValueToString(r.priceListId, "priceListId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.priceListUpdate == nil {
		return localVarReturnValue, nil, reportError("priceListUpdate is required and must be specified")
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
	localVarPostBody = r.priceListUpdate
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

type PriceListsApiPOSTPriceListsRequest struct {
	ctx             context.Context
	ApiService      *PriceListsApiService
	priceListCreate *PriceListCreate
}

func (r PriceListsApiPOSTPriceListsRequest) PriceListCreate(priceListCreate PriceListCreate) PriceListsApiPOSTPriceListsRequest {
	r.priceListCreate = &priceListCreate
	return r
}

func (r PriceListsApiPOSTPriceListsRequest) Execute() (*POSTPriceLists201Response, *http.Response, error) {
	return r.ApiService.POSTPriceListsExecute(r)
}

/*
POSTPriceLists Create a price list

Create a price list

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PriceListsApiPOSTPriceListsRequest
*/
func (a *PriceListsApiService) POSTPriceLists(ctx context.Context) PriceListsApiPOSTPriceListsRequest {
	return PriceListsApiPOSTPriceListsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTPriceLists201Response
func (a *PriceListsApiService) POSTPriceListsExecute(r PriceListsApiPOSTPriceListsRequest) (*POSTPriceLists201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTPriceLists201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceListsApiService.POSTPriceLists")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_lists"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.priceListCreate == nil {
		return localVarReturnValue, nil, reportError("priceListCreate is required and must be specified")
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
	localVarPostBody = r.priceListCreate
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
