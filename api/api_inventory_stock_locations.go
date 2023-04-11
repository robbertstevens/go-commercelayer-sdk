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

// InventoryStockLocationsApiService InventoryStockLocationsApi service
type InventoryStockLocationsApiService service

type InventoryStockLocationsApiDELETEInventoryStockLocationsInventoryStockLocationIdRequest struct {
	ctx                      context.Context
	ApiService               *InventoryStockLocationsApiService
	inventoryStockLocationId interface{}
}

func (r InventoryStockLocationsApiDELETEInventoryStockLocationsInventoryStockLocationIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEInventoryStockLocationsInventoryStockLocationIdExecute(r)
}

/*
DELETEInventoryStockLocationsInventoryStockLocationId Delete an inventory stock location

Delete an inventory stock location

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inventoryStockLocationId The resource's id
	@return InventoryStockLocationsApiDELETEInventoryStockLocationsInventoryStockLocationIdRequest
*/
func (a *InventoryStockLocationsApiService) DELETEInventoryStockLocationsInventoryStockLocationId(ctx context.Context, inventoryStockLocationId interface{}) InventoryStockLocationsApiDELETEInventoryStockLocationsInventoryStockLocationIdRequest {
	return InventoryStockLocationsApiDELETEInventoryStockLocationsInventoryStockLocationIdRequest{
		ApiService:               a,
		ctx:                      ctx,
		inventoryStockLocationId: inventoryStockLocationId,
	}
}

// Execute executes the request
func (a *InventoryStockLocationsApiService) DELETEInventoryStockLocationsInventoryStockLocationIdExecute(r InventoryStockLocationsApiDELETEInventoryStockLocationsInventoryStockLocationIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryStockLocationsApiService.DELETEInventoryStockLocationsInventoryStockLocationId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_stock_locations/{inventoryStockLocationId}"
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

type InventoryStockLocationsApiGETInventoryModelIdInventoryStockLocationsRequest struct {
	ctx              context.Context
	ApiService       *InventoryStockLocationsApiService
	inventoryModelId interface{}
}

func (r InventoryStockLocationsApiGETInventoryModelIdInventoryStockLocationsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETInventoryModelIdInventoryStockLocationsExecute(r)
}

/*
GETInventoryModelIdInventoryStockLocations Retrieve the inventory stock locations associated to the inventory model

Retrieve the inventory stock locations associated to the inventory model

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inventoryModelId The resource's id
	@return InventoryStockLocationsApiGETInventoryModelIdInventoryStockLocationsRequest
*/
func (a *InventoryStockLocationsApiService) GETInventoryModelIdInventoryStockLocations(ctx context.Context, inventoryModelId interface{}) InventoryStockLocationsApiGETInventoryModelIdInventoryStockLocationsRequest {
	return InventoryStockLocationsApiGETInventoryModelIdInventoryStockLocationsRequest{
		ApiService:       a,
		ctx:              ctx,
		inventoryModelId: inventoryModelId,
	}
}

// Execute executes the request
func (a *InventoryStockLocationsApiService) GETInventoryModelIdInventoryStockLocationsExecute(r InventoryStockLocationsApiGETInventoryModelIdInventoryStockLocationsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryStockLocationsApiService.GETInventoryModelIdInventoryStockLocations")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_models/{inventoryModelId}/inventory_stock_locations"
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

type InventoryStockLocationsApiGETInventoryStockLocationsRequest struct {
	ctx        context.Context
	ApiService *InventoryStockLocationsApiService
}

func (r InventoryStockLocationsApiGETInventoryStockLocationsRequest) Execute() (*GETInventoryStockLocations200Response, *http.Response, error) {
	return r.ApiService.GETInventoryStockLocationsExecute(r)
}

/*
GETInventoryStockLocations List all inventory stock locations

List all inventory stock locations

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return InventoryStockLocationsApiGETInventoryStockLocationsRequest
*/
func (a *InventoryStockLocationsApiService) GETInventoryStockLocations(ctx context.Context) InventoryStockLocationsApiGETInventoryStockLocationsRequest {
	return InventoryStockLocationsApiGETInventoryStockLocationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETInventoryStockLocations200Response
func (a *InventoryStockLocationsApiService) GETInventoryStockLocationsExecute(r InventoryStockLocationsApiGETInventoryStockLocationsRequest) (*GETInventoryStockLocations200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETInventoryStockLocations200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryStockLocationsApiService.GETInventoryStockLocations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_stock_locations"

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

type InventoryStockLocationsApiGETInventoryStockLocationsInventoryStockLocationIdRequest struct {
	ctx                      context.Context
	ApiService               *InventoryStockLocationsApiService
	inventoryStockLocationId interface{}
}

func (r InventoryStockLocationsApiGETInventoryStockLocationsInventoryStockLocationIdRequest) Execute() (*GETInventoryStockLocationsInventoryStockLocationId200Response, *http.Response, error) {
	return r.ApiService.GETInventoryStockLocationsInventoryStockLocationIdExecute(r)
}

/*
GETInventoryStockLocationsInventoryStockLocationId Retrieve an inventory stock location

Retrieve an inventory stock location

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inventoryStockLocationId The resource's id
	@return InventoryStockLocationsApiGETInventoryStockLocationsInventoryStockLocationIdRequest
*/
func (a *InventoryStockLocationsApiService) GETInventoryStockLocationsInventoryStockLocationId(ctx context.Context, inventoryStockLocationId interface{}) InventoryStockLocationsApiGETInventoryStockLocationsInventoryStockLocationIdRequest {
	return InventoryStockLocationsApiGETInventoryStockLocationsInventoryStockLocationIdRequest{
		ApiService:               a,
		ctx:                      ctx,
		inventoryStockLocationId: inventoryStockLocationId,
	}
}

// Execute executes the request
//
//	@return GETInventoryStockLocationsInventoryStockLocationId200Response
func (a *InventoryStockLocationsApiService) GETInventoryStockLocationsInventoryStockLocationIdExecute(r InventoryStockLocationsApiGETInventoryStockLocationsInventoryStockLocationIdRequest) (*GETInventoryStockLocationsInventoryStockLocationId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETInventoryStockLocationsInventoryStockLocationId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryStockLocationsApiService.GETInventoryStockLocationsInventoryStockLocationId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_stock_locations/{inventoryStockLocationId}"
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

type InventoryStockLocationsApiGETStockLocationIdInventoryStockLocationsRequest struct {
	ctx             context.Context
	ApiService      *InventoryStockLocationsApiService
	stockLocationId interface{}
}

func (r InventoryStockLocationsApiGETStockLocationIdInventoryStockLocationsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStockLocationIdInventoryStockLocationsExecute(r)
}

/*
GETStockLocationIdInventoryStockLocations Retrieve the inventory stock locations associated to the stock location

Retrieve the inventory stock locations associated to the stock location

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stockLocationId The resource's id
	@return InventoryStockLocationsApiGETStockLocationIdInventoryStockLocationsRequest
*/
func (a *InventoryStockLocationsApiService) GETStockLocationIdInventoryStockLocations(ctx context.Context, stockLocationId interface{}) InventoryStockLocationsApiGETStockLocationIdInventoryStockLocationsRequest {
	return InventoryStockLocationsApiGETStockLocationIdInventoryStockLocationsRequest{
		ApiService:      a,
		ctx:             ctx,
		stockLocationId: stockLocationId,
	}
}

// Execute executes the request
func (a *InventoryStockLocationsApiService) GETStockLocationIdInventoryStockLocationsExecute(r InventoryStockLocationsApiGETStockLocationIdInventoryStockLocationsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryStockLocationsApiService.GETStockLocationIdInventoryStockLocations")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_locations/{stockLocationId}/inventory_stock_locations"
	localVarPath = strings.Replace(localVarPath, "{"+"stockLocationId"+"}", url.PathEscape(parameterValueToString(r.stockLocationId, "stockLocationId")), -1)

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

type InventoryStockLocationsApiPATCHInventoryStockLocationsInventoryStockLocationIdRequest struct {
	ctx                                                         context.Context
	ApiService                                                  *InventoryStockLocationsApiService
	pATCHInventoryStockLocationsInventoryStockLocationIdRequest *PATCHInventoryStockLocationsInventoryStockLocationIdRequest
	inventoryStockLocationId                                    interface{}
}

func (r InventoryStockLocationsApiPATCHInventoryStockLocationsInventoryStockLocationIdRequest) PATCHInventoryStockLocationsInventoryStockLocationIdRequest(pATCHInventoryStockLocationsInventoryStockLocationIdRequest PATCHInventoryStockLocationsInventoryStockLocationIdRequest) InventoryStockLocationsApiPATCHInventoryStockLocationsInventoryStockLocationIdRequest {
	r.pATCHInventoryStockLocationsInventoryStockLocationIdRequest = &pATCHInventoryStockLocationsInventoryStockLocationIdRequest
	return r
}

func (r InventoryStockLocationsApiPATCHInventoryStockLocationsInventoryStockLocationIdRequest) Execute() (*PATCHInventoryStockLocationsInventoryStockLocationId200Response, *http.Response, error) {
	return r.ApiService.PATCHInventoryStockLocationsInventoryStockLocationIdExecute(r)
}

/*
PATCHInventoryStockLocationsInventoryStockLocationId Update an inventory stock location

Update an inventory stock location

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inventoryStockLocationId The resource's id
	@return InventoryStockLocationsApiPATCHInventoryStockLocationsInventoryStockLocationIdRequest
*/
func (a *InventoryStockLocationsApiService) PATCHInventoryStockLocationsInventoryStockLocationId(ctx context.Context, inventoryStockLocationId interface{}) InventoryStockLocationsApiPATCHInventoryStockLocationsInventoryStockLocationIdRequest {
	return InventoryStockLocationsApiPATCHInventoryStockLocationsInventoryStockLocationIdRequest{
		ApiService:               a,
		ctx:                      ctx,
		inventoryStockLocationId: inventoryStockLocationId,
	}
}

// Execute executes the request
//
//	@return PATCHInventoryStockLocationsInventoryStockLocationId200Response
func (a *InventoryStockLocationsApiService) PATCHInventoryStockLocationsInventoryStockLocationIdExecute(r InventoryStockLocationsApiPATCHInventoryStockLocationsInventoryStockLocationIdRequest) (*PATCHInventoryStockLocationsInventoryStockLocationId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHInventoryStockLocationsInventoryStockLocationId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryStockLocationsApiService.PATCHInventoryStockLocationsInventoryStockLocationId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_stock_locations/{inventoryStockLocationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryStockLocationId"+"}", url.PathEscape(parameterValueToString(r.inventoryStockLocationId, "inventoryStockLocationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pATCHInventoryStockLocationsInventoryStockLocationIdRequest == nil {
		return localVarReturnValue, nil, reportError("pATCHInventoryStockLocationsInventoryStockLocationIdRequest is required and must be specified")
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
	localVarPostBody = r.pATCHInventoryStockLocationsInventoryStockLocationIdRequest
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

type InventoryStockLocationsApiPOSTInventoryStockLocationsRequest struct {
	ctx                                context.Context
	ApiService                         *InventoryStockLocationsApiService
	pOSTInventoryStockLocationsRequest *POSTInventoryStockLocationsRequest
}

func (r InventoryStockLocationsApiPOSTInventoryStockLocationsRequest) POSTInventoryStockLocationsRequest(pOSTInventoryStockLocationsRequest POSTInventoryStockLocationsRequest) InventoryStockLocationsApiPOSTInventoryStockLocationsRequest {
	r.pOSTInventoryStockLocationsRequest = &pOSTInventoryStockLocationsRequest
	return r
}

func (r InventoryStockLocationsApiPOSTInventoryStockLocationsRequest) Execute() (*POSTInventoryStockLocations201Response, *http.Response, error) {
	return r.ApiService.POSTInventoryStockLocationsExecute(r)
}

/*
POSTInventoryStockLocations Create an inventory stock location

Create an inventory stock location

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return InventoryStockLocationsApiPOSTInventoryStockLocationsRequest
*/
func (a *InventoryStockLocationsApiService) POSTInventoryStockLocations(ctx context.Context) InventoryStockLocationsApiPOSTInventoryStockLocationsRequest {
	return InventoryStockLocationsApiPOSTInventoryStockLocationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTInventoryStockLocations201Response
func (a *InventoryStockLocationsApiService) POSTInventoryStockLocationsExecute(r InventoryStockLocationsApiPOSTInventoryStockLocationsRequest) (*POSTInventoryStockLocations201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTInventoryStockLocations201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryStockLocationsApiService.POSTInventoryStockLocations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_stock_locations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pOSTInventoryStockLocationsRequest == nil {
		return localVarReturnValue, nil, reportError("pOSTInventoryStockLocationsRequest is required and must be specified")
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
	localVarPostBody = r.pOSTInventoryStockLocationsRequest
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
