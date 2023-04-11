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

// InventoryModelsApiService InventoryModelsApi service
type InventoryModelsApiService service

type InventoryModelsApiDELETEInventoryModelsInventoryModelIdRequest struct {
	ctx              context.Context
	ApiService       *InventoryModelsApiService
	inventoryModelId interface{}
}

func (r InventoryModelsApiDELETEInventoryModelsInventoryModelIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEInventoryModelsInventoryModelIdExecute(r)
}

/*
DELETEInventoryModelsInventoryModelId Delete an inventory model

Delete an inventory model

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inventoryModelId The resource's id
	@return InventoryModelsApiDELETEInventoryModelsInventoryModelIdRequest
*/
func (a *InventoryModelsApiService) DELETEInventoryModelsInventoryModelId(ctx context.Context, inventoryModelId interface{}) InventoryModelsApiDELETEInventoryModelsInventoryModelIdRequest {
	return InventoryModelsApiDELETEInventoryModelsInventoryModelIdRequest{
		ApiService:       a,
		ctx:              ctx,
		inventoryModelId: inventoryModelId,
	}
}

// Execute executes the request
func (a *InventoryModelsApiService) DELETEInventoryModelsInventoryModelIdExecute(r InventoryModelsApiDELETEInventoryModelsInventoryModelIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryModelsApiService.DELETEInventoryModelsInventoryModelId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_models/{inventoryModelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryModelId"+"}", url.PathEscape(parameterValueToString(r.inventoryModelId, "inventoryModelId")), -1)

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

type InventoryModelsApiGETInventoryModelsRequest struct {
	ctx        context.Context
	ApiService *InventoryModelsApiService
}

func (r InventoryModelsApiGETInventoryModelsRequest) Execute() (*GETInventoryModels200Response, *http.Response, error) {
	return r.ApiService.GETInventoryModelsExecute(r)
}

/*
GETInventoryModels List all inventory models

List all inventory models

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return InventoryModelsApiGETInventoryModelsRequest
*/
func (a *InventoryModelsApiService) GETInventoryModels(ctx context.Context) InventoryModelsApiGETInventoryModelsRequest {
	return InventoryModelsApiGETInventoryModelsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETInventoryModels200Response
func (a *InventoryModelsApiService) GETInventoryModelsExecute(r InventoryModelsApiGETInventoryModelsRequest) (*GETInventoryModels200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETInventoryModels200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryModelsApiService.GETInventoryModels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_models"

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

type InventoryModelsApiGETInventoryModelsInventoryModelIdRequest struct {
	ctx              context.Context
	ApiService       *InventoryModelsApiService
	inventoryModelId interface{}
}

func (r InventoryModelsApiGETInventoryModelsInventoryModelIdRequest) Execute() (*GETInventoryModelsInventoryModelId200Response, *http.Response, error) {
	return r.ApiService.GETInventoryModelsInventoryModelIdExecute(r)
}

/*
GETInventoryModelsInventoryModelId Retrieve an inventory model

Retrieve an inventory model

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inventoryModelId The resource's id
	@return InventoryModelsApiGETInventoryModelsInventoryModelIdRequest
*/
func (a *InventoryModelsApiService) GETInventoryModelsInventoryModelId(ctx context.Context, inventoryModelId interface{}) InventoryModelsApiGETInventoryModelsInventoryModelIdRequest {
	return InventoryModelsApiGETInventoryModelsInventoryModelIdRequest{
		ApiService:       a,
		ctx:              ctx,
		inventoryModelId: inventoryModelId,
	}
}

// Execute executes the request
//
//	@return GETInventoryModelsInventoryModelId200Response
func (a *InventoryModelsApiService) GETInventoryModelsInventoryModelIdExecute(r InventoryModelsApiGETInventoryModelsInventoryModelIdRequest) (*GETInventoryModelsInventoryModelId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETInventoryModelsInventoryModelId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryModelsApiService.GETInventoryModelsInventoryModelId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_models/{inventoryModelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryModelId"+"}", url.PathEscape(parameterValueToString(r.inventoryModelId, "inventoryModelId")), -1)

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

type InventoryModelsApiGETInventoryReturnLocationIdInventoryModelRequest struct {
	ctx                       context.Context
	ApiService                *InventoryModelsApiService
	inventoryReturnLocationId interface{}
}

func (r InventoryModelsApiGETInventoryReturnLocationIdInventoryModelRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETInventoryReturnLocationIdInventoryModelExecute(r)
}

/*
GETInventoryReturnLocationIdInventoryModel Retrieve the inventory model associated to the inventory return location

Retrieve the inventory model associated to the inventory return location

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inventoryReturnLocationId The resource's id
	@return InventoryModelsApiGETInventoryReturnLocationIdInventoryModelRequest
*/
func (a *InventoryModelsApiService) GETInventoryReturnLocationIdInventoryModel(ctx context.Context, inventoryReturnLocationId interface{}) InventoryModelsApiGETInventoryReturnLocationIdInventoryModelRequest {
	return InventoryModelsApiGETInventoryReturnLocationIdInventoryModelRequest{
		ApiService:                a,
		ctx:                       ctx,
		inventoryReturnLocationId: inventoryReturnLocationId,
	}
}

// Execute executes the request
func (a *InventoryModelsApiService) GETInventoryReturnLocationIdInventoryModelExecute(r InventoryModelsApiGETInventoryReturnLocationIdInventoryModelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryModelsApiService.GETInventoryReturnLocationIdInventoryModel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_return_locations/{inventoryReturnLocationId}/inventory_model"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryReturnLocationId"+"}", url.PathEscape(parameterValueToString(r.inventoryReturnLocationId, "inventoryReturnLocationId")), -1)

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

type InventoryModelsApiGETInventoryStockLocationIdInventoryModelRequest struct {
	ctx                      context.Context
	ApiService               *InventoryModelsApiService
	inventoryStockLocationId interface{}
}

func (r InventoryModelsApiGETInventoryStockLocationIdInventoryModelRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETInventoryStockLocationIdInventoryModelExecute(r)
}

/*
GETInventoryStockLocationIdInventoryModel Retrieve the inventory model associated to the inventory stock location

Retrieve the inventory model associated to the inventory stock location

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inventoryStockLocationId The resource's id
	@return InventoryModelsApiGETInventoryStockLocationIdInventoryModelRequest
*/
func (a *InventoryModelsApiService) GETInventoryStockLocationIdInventoryModel(ctx context.Context, inventoryStockLocationId interface{}) InventoryModelsApiGETInventoryStockLocationIdInventoryModelRequest {
	return InventoryModelsApiGETInventoryStockLocationIdInventoryModelRequest{
		ApiService:               a,
		ctx:                      ctx,
		inventoryStockLocationId: inventoryStockLocationId,
	}
}

// Execute executes the request
func (a *InventoryModelsApiService) GETInventoryStockLocationIdInventoryModelExecute(r InventoryModelsApiGETInventoryStockLocationIdInventoryModelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryModelsApiService.GETInventoryStockLocationIdInventoryModel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_stock_locations/{inventoryStockLocationId}/inventory_model"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryStockLocationId"+"}", url.PathEscape(parameterValueToString(r.inventoryStockLocationId, "inventoryStockLocationId")), -1)

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

type InventoryModelsApiGETMarketIdInventoryModelRequest struct {
	ctx        context.Context
	ApiService *InventoryModelsApiService
	marketId   interface{}
}

func (r InventoryModelsApiGETMarketIdInventoryModelRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETMarketIdInventoryModelExecute(r)
}

/*
GETMarketIdInventoryModel Retrieve the inventory model associated to the market

Retrieve the inventory model associated to the market

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param marketId The resource's id
	@return InventoryModelsApiGETMarketIdInventoryModelRequest
*/
func (a *InventoryModelsApiService) GETMarketIdInventoryModel(ctx context.Context, marketId interface{}) InventoryModelsApiGETMarketIdInventoryModelRequest {
	return InventoryModelsApiGETMarketIdInventoryModelRequest{
		ApiService: a,
		ctx:        ctx,
		marketId:   marketId,
	}
}

// Execute executes the request
func (a *InventoryModelsApiService) GETMarketIdInventoryModelExecute(r InventoryModelsApiGETMarketIdInventoryModelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryModelsApiService.GETMarketIdInventoryModel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/markets/{marketId}/inventory_model"
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

type InventoryModelsApiPATCHInventoryModelsInventoryModelIdRequest struct {
	ctx                                         context.Context
	ApiService                                  *InventoryModelsApiService
	pATCHInventoryModelsInventoryModelIdRequest *PATCHInventoryModelsInventoryModelIdRequest
	inventoryModelId                            interface{}
}

func (r InventoryModelsApiPATCHInventoryModelsInventoryModelIdRequest) PATCHInventoryModelsInventoryModelIdRequest(pATCHInventoryModelsInventoryModelIdRequest PATCHInventoryModelsInventoryModelIdRequest) InventoryModelsApiPATCHInventoryModelsInventoryModelIdRequest {
	r.pATCHInventoryModelsInventoryModelIdRequest = &pATCHInventoryModelsInventoryModelIdRequest
	return r
}

func (r InventoryModelsApiPATCHInventoryModelsInventoryModelIdRequest) Execute() (*PATCHInventoryModelsInventoryModelId200Response, *http.Response, error) {
	return r.ApiService.PATCHInventoryModelsInventoryModelIdExecute(r)
}

/*
PATCHInventoryModelsInventoryModelId Update an inventory model

Update an inventory model

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inventoryModelId The resource's id
	@return InventoryModelsApiPATCHInventoryModelsInventoryModelIdRequest
*/
func (a *InventoryModelsApiService) PATCHInventoryModelsInventoryModelId(ctx context.Context, inventoryModelId interface{}) InventoryModelsApiPATCHInventoryModelsInventoryModelIdRequest {
	return InventoryModelsApiPATCHInventoryModelsInventoryModelIdRequest{
		ApiService:       a,
		ctx:              ctx,
		inventoryModelId: inventoryModelId,
	}
}

// Execute executes the request
//
//	@return PATCHInventoryModelsInventoryModelId200Response
func (a *InventoryModelsApiService) PATCHInventoryModelsInventoryModelIdExecute(r InventoryModelsApiPATCHInventoryModelsInventoryModelIdRequest) (*PATCHInventoryModelsInventoryModelId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHInventoryModelsInventoryModelId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryModelsApiService.PATCHInventoryModelsInventoryModelId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_models/{inventoryModelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryModelId"+"}", url.PathEscape(parameterValueToString(r.inventoryModelId, "inventoryModelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pATCHInventoryModelsInventoryModelIdRequest == nil {
		return localVarReturnValue, nil, reportError("pATCHInventoryModelsInventoryModelIdRequest is required and must be specified")
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
	localVarPostBody = r.pATCHInventoryModelsInventoryModelIdRequest
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

type InventoryModelsApiPOSTInventoryModelsRequest struct {
	ctx                        context.Context
	ApiService                 *InventoryModelsApiService
	pOSTInventoryModelsRequest *POSTInventoryModelsRequest
}

func (r InventoryModelsApiPOSTInventoryModelsRequest) POSTInventoryModelsRequest(pOSTInventoryModelsRequest POSTInventoryModelsRequest) InventoryModelsApiPOSTInventoryModelsRequest {
	r.pOSTInventoryModelsRequest = &pOSTInventoryModelsRequest
	return r
}

func (r InventoryModelsApiPOSTInventoryModelsRequest) Execute() (*POSTInventoryModels201Response, *http.Response, error) {
	return r.ApiService.POSTInventoryModelsExecute(r)
}

/*
POSTInventoryModels Create an inventory model

Create an inventory model

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return InventoryModelsApiPOSTInventoryModelsRequest
*/
func (a *InventoryModelsApiService) POSTInventoryModels(ctx context.Context) InventoryModelsApiPOSTInventoryModelsRequest {
	return InventoryModelsApiPOSTInventoryModelsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTInventoryModels201Response
func (a *InventoryModelsApiService) POSTInventoryModelsExecute(r InventoryModelsApiPOSTInventoryModelsRequest) (*POSTInventoryModels201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTInventoryModels201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryModelsApiService.POSTInventoryModels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_models"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pOSTInventoryModelsRequest == nil {
		return localVarReturnValue, nil, reportError("pOSTInventoryModelsRequest is required and must be specified")
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
	localVarPostBody = r.pOSTInventoryModelsRequest
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
