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


// StockLineItemsApiService StockLineItemsApi service
type StockLineItemsApiService service

type StockLineItemsApiGETLineItemIdStockLineItemsRequest struct {
	ctx context.Context
	ApiService *StockLineItemsApiService
	lineItemId string
}

func (r StockLineItemsApiGETLineItemIdStockLineItemsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETLineItemIdStockLineItemsExecute(r)
}

/*
GETLineItemIdStockLineItems Retrieve the stock line items associated to the line item

Retrieve the stock line items associated to the line item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param lineItemId The resource's id
 @return StockLineItemsApiGETLineItemIdStockLineItemsRequest
*/
func (a *StockLineItemsApiService) GETLineItemIdStockLineItems(ctx context.Context, lineItemId string) StockLineItemsApiGETLineItemIdStockLineItemsRequest {
	return StockLineItemsApiGETLineItemIdStockLineItemsRequest{
		ApiService: a,
		ctx: ctx,
		lineItemId: lineItemId,
	}
}

// Execute executes the request
func (a *StockLineItemsApiService) GETLineItemIdStockLineItemsExecute(r StockLineItemsApiGETLineItemIdStockLineItemsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.GETLineItemIdStockLineItems")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/line_items/{lineItemId}/stock_line_items"
	localVarPath = strings.Replace(localVarPath, "{"+"lineItemId"+"}", url.PathEscape(parameterToString(r.lineItemId, "")), -1)

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

type StockLineItemsApiGETParcelLineItemIdStockLineItemRequest struct {
	ctx context.Context
	ApiService *StockLineItemsApiService
	parcelLineItemId string
}

func (r StockLineItemsApiGETParcelLineItemIdStockLineItemRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETParcelLineItemIdStockLineItemExecute(r)
}

/*
GETParcelLineItemIdStockLineItem Retrieve the stock line item associated to the parcel line item

Retrieve the stock line item associated to the parcel line item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parcelLineItemId The resource's id
 @return StockLineItemsApiGETParcelLineItemIdStockLineItemRequest
*/
func (a *StockLineItemsApiService) GETParcelLineItemIdStockLineItem(ctx context.Context, parcelLineItemId string) StockLineItemsApiGETParcelLineItemIdStockLineItemRequest {
	return StockLineItemsApiGETParcelLineItemIdStockLineItemRequest{
		ApiService: a,
		ctx: ctx,
		parcelLineItemId: parcelLineItemId,
	}
}

// Execute executes the request
func (a *StockLineItemsApiService) GETParcelLineItemIdStockLineItemExecute(r StockLineItemsApiGETParcelLineItemIdStockLineItemRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.GETParcelLineItemIdStockLineItem")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parcel_line_items/{parcelLineItemId}/stock_line_item"
	localVarPath = strings.Replace(localVarPath, "{"+"parcelLineItemId"+"}", url.PathEscape(parameterToString(r.parcelLineItemId, "")), -1)

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

type StockLineItemsApiGETShipmentIdStockLineItemsRequest struct {
	ctx context.Context
	ApiService *StockLineItemsApiService
	shipmentId string
}

func (r StockLineItemsApiGETShipmentIdStockLineItemsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShipmentIdStockLineItemsExecute(r)
}

/*
GETShipmentIdStockLineItems Retrieve the stock line items associated to the shipment

Retrieve the stock line items associated to the shipment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shipmentId The resource's id
 @return StockLineItemsApiGETShipmentIdStockLineItemsRequest
*/
func (a *StockLineItemsApiService) GETShipmentIdStockLineItems(ctx context.Context, shipmentId string) StockLineItemsApiGETShipmentIdStockLineItemsRequest {
	return StockLineItemsApiGETShipmentIdStockLineItemsRequest{
		ApiService: a,
		ctx: ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
func (a *StockLineItemsApiService) GETShipmentIdStockLineItemsExecute(r StockLineItemsApiGETShipmentIdStockLineItemsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.GETShipmentIdStockLineItems")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipments/{shipmentId}/stock_line_items"
	localVarPath = strings.Replace(localVarPath, "{"+"shipmentId"+"}", url.PathEscape(parameterToString(r.shipmentId, "")), -1)

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

type StockLineItemsApiGETStockLineItemsRequest struct {
	ctx context.Context
	ApiService *StockLineItemsApiService
}

func (r StockLineItemsApiGETStockLineItemsRequest) Execute() (*GETStockLineItems200Response, *http.Response, error) {
	return r.ApiService.GETStockLineItemsExecute(r)
}

/*
GETStockLineItems List all stock line items

List all stock line items

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return StockLineItemsApiGETStockLineItemsRequest
*/
func (a *StockLineItemsApiService) GETStockLineItems(ctx context.Context) StockLineItemsApiGETStockLineItemsRequest {
	return StockLineItemsApiGETStockLineItemsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GETStockLineItems200Response
func (a *StockLineItemsApiService) GETStockLineItemsExecute(r StockLineItemsApiGETStockLineItemsRequest) (*GETStockLineItems200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GETStockLineItems200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.GETStockLineItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_line_items"

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

type StockLineItemsApiGETStockLineItemsStockLineItemIdRequest struct {
	ctx context.Context
	ApiService *StockLineItemsApiService
	stockLineItemId string
}

func (r StockLineItemsApiGETStockLineItemsStockLineItemIdRequest) Execute() (*GETStockLineItemsStockLineItemId200Response, *http.Response, error) {
	return r.ApiService.GETStockLineItemsStockLineItemIdExecute(r)
}

/*
GETStockLineItemsStockLineItemId Retrieve a stock line item

Retrieve a stock line item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stockLineItemId The resource's id
 @return StockLineItemsApiGETStockLineItemsStockLineItemIdRequest
*/
func (a *StockLineItemsApiService) GETStockLineItemsStockLineItemId(ctx context.Context, stockLineItemId string) StockLineItemsApiGETStockLineItemsStockLineItemIdRequest {
	return StockLineItemsApiGETStockLineItemsStockLineItemIdRequest{
		ApiService: a,
		ctx: ctx,
		stockLineItemId: stockLineItemId,
	}
}

// Execute executes the request
//  @return GETStockLineItemsStockLineItemId200Response
func (a *StockLineItemsApiService) GETStockLineItemsStockLineItemIdExecute(r StockLineItemsApiGETStockLineItemsStockLineItemIdRequest) (*GETStockLineItemsStockLineItemId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GETStockLineItemsStockLineItemId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.GETStockLineItemsStockLineItemId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_line_items/{stockLineItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stockLineItemId"+"}", url.PathEscape(parameterToString(r.stockLineItemId, "")), -1)

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
