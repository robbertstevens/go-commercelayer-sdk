/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
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

// LineItemsApiService LineItemsApi service
type LineItemsApiService service

type LineItemsApiDELETELineItemsLineItemIdRequest struct {
	ctx        context.Context
	ApiService *LineItemsApiService
	lineItemId string
}

func (r LineItemsApiDELETELineItemsLineItemIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETELineItemsLineItemIdExecute(r)
}

/*
DELETELineItemsLineItemId Delete a line item

Delete a line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param lineItemId The resource's id
	@return LineItemsApiDELETELineItemsLineItemIdRequest
*/
func (a *LineItemsApiService) DELETELineItemsLineItemId(ctx context.Context, lineItemId string) LineItemsApiDELETELineItemsLineItemIdRequest {
	return LineItemsApiDELETELineItemsLineItemIdRequest{
		ApiService: a,
		ctx:        ctx,
		lineItemId: lineItemId,
	}
}

// Execute executes the request
func (a *LineItemsApiService) DELETELineItemsLineItemIdExecute(r LineItemsApiDELETELineItemsLineItemIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineItemsApiService.DELETELineItemsLineItemId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/line_items/{lineItemId}"
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

type LineItemsApiGETLineItemOptionIdLineItemRequest struct {
	ctx              context.Context
	ApiService       *LineItemsApiService
	lineItemOptionId string
}

func (r LineItemsApiGETLineItemOptionIdLineItemRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETLineItemOptionIdLineItemExecute(r)
}

/*
GETLineItemOptionIdLineItem Retrieve the line item associated to the line item option

Retrieve the line item associated to the line item option

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param lineItemOptionId The resource's id
	@return LineItemsApiGETLineItemOptionIdLineItemRequest
*/
func (a *LineItemsApiService) GETLineItemOptionIdLineItem(ctx context.Context, lineItemOptionId string) LineItemsApiGETLineItemOptionIdLineItemRequest {
	return LineItemsApiGETLineItemOptionIdLineItemRequest{
		ApiService:       a,
		ctx:              ctx,
		lineItemOptionId: lineItemOptionId,
	}
}

// Execute executes the request
func (a *LineItemsApiService) GETLineItemOptionIdLineItemExecute(r LineItemsApiGETLineItemOptionIdLineItemRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineItemsApiService.GETLineItemOptionIdLineItem")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/line_item_options/{lineItemOptionId}/line_item"
	localVarPath = strings.Replace(localVarPath, "{"+"lineItemOptionId"+"}", url.PathEscape(parameterToString(r.lineItemOptionId, "")), -1)

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

type LineItemsApiGETLineItemsRequest struct {
	ctx        context.Context
	ApiService *LineItemsApiService
}

func (r LineItemsApiGETLineItemsRequest) Execute() (*GETLineItems200Response, *http.Response, error) {
	return r.ApiService.GETLineItemsExecute(r)
}

/*
GETLineItems List all line items

List all line items

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return LineItemsApiGETLineItemsRequest
*/
func (a *LineItemsApiService) GETLineItems(ctx context.Context) LineItemsApiGETLineItemsRequest {
	return LineItemsApiGETLineItemsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETLineItems200Response
func (a *LineItemsApiService) GETLineItemsExecute(r LineItemsApiGETLineItemsRequest) (*GETLineItems200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETLineItems200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineItemsApiService.GETLineItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/line_items"

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

type LineItemsApiGETLineItemsLineItemIdRequest struct {
	ctx        context.Context
	ApiService *LineItemsApiService
	lineItemId string
}

func (r LineItemsApiGETLineItemsLineItemIdRequest) Execute() (*GETLineItemsLineItemId200Response, *http.Response, error) {
	return r.ApiService.GETLineItemsLineItemIdExecute(r)
}

/*
GETLineItemsLineItemId Retrieve a line item

Retrieve a line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param lineItemId The resource's id
	@return LineItemsApiGETLineItemsLineItemIdRequest
*/
func (a *LineItemsApiService) GETLineItemsLineItemId(ctx context.Context, lineItemId string) LineItemsApiGETLineItemsLineItemIdRequest {
	return LineItemsApiGETLineItemsLineItemIdRequest{
		ApiService: a,
		ctx:        ctx,
		lineItemId: lineItemId,
	}
}

// Execute executes the request
//
//	@return GETLineItemsLineItemId200Response
func (a *LineItemsApiService) GETLineItemsLineItemIdExecute(r LineItemsApiGETLineItemsLineItemIdRequest) (*GETLineItemsLineItemId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETLineItemsLineItemId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineItemsApiService.GETLineItemsLineItemId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/line_items/{lineItemId}"
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

type LineItemsApiGETOrderIdLineItemsRequest struct {
	ctx        context.Context
	ApiService *LineItemsApiService
	orderId    string
}

func (r LineItemsApiGETOrderIdLineItemsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderIdLineItemsExecute(r)
}

/*
GETOrderIdLineItems Retrieve the line items associated to the order

Retrieve the line items associated to the order

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId The resource's id
	@return LineItemsApiGETOrderIdLineItemsRequest
*/
func (a *LineItemsApiService) GETOrderIdLineItems(ctx context.Context, orderId string) LineItemsApiGETOrderIdLineItemsRequest {
	return LineItemsApiGETOrderIdLineItemsRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *LineItemsApiService) GETOrderIdLineItemsExecute(r LineItemsApiGETOrderIdLineItemsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineItemsApiService.GETOrderIdLineItems")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{orderId}/line_items"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterToString(r.orderId, "")), -1)

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

type LineItemsApiGETReturnLineItemIdLineItemRequest struct {
	ctx              context.Context
	ApiService       *LineItemsApiService
	returnLineItemId string
}

func (r LineItemsApiGETReturnLineItemIdLineItemRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETReturnLineItemIdLineItemExecute(r)
}

/*
GETReturnLineItemIdLineItem Retrieve the line item associated to the return line item

Retrieve the line item associated to the return line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param returnLineItemId The resource's id
	@return LineItemsApiGETReturnLineItemIdLineItemRequest
*/
func (a *LineItemsApiService) GETReturnLineItemIdLineItem(ctx context.Context, returnLineItemId string) LineItemsApiGETReturnLineItemIdLineItemRequest {
	return LineItemsApiGETReturnLineItemIdLineItemRequest{
		ApiService:       a,
		ctx:              ctx,
		returnLineItemId: returnLineItemId,
	}
}

// Execute executes the request
func (a *LineItemsApiService) GETReturnLineItemIdLineItemExecute(r LineItemsApiGETReturnLineItemIdLineItemRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineItemsApiService.GETReturnLineItemIdLineItem")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return_line_items/{returnLineItemId}/line_item"
	localVarPath = strings.Replace(localVarPath, "{"+"returnLineItemId"+"}", url.PathEscape(parameterToString(r.returnLineItemId, "")), -1)

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

type LineItemsApiGETStockLineItemIdLineItemRequest struct {
	ctx             context.Context
	ApiService      *LineItemsApiService
	stockLineItemId string
}

func (r LineItemsApiGETStockLineItemIdLineItemRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStockLineItemIdLineItemExecute(r)
}

/*
GETStockLineItemIdLineItem Retrieve the line item associated to the stock line item

Retrieve the line item associated to the stock line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stockLineItemId The resource's id
	@return LineItemsApiGETStockLineItemIdLineItemRequest
*/
func (a *LineItemsApiService) GETStockLineItemIdLineItem(ctx context.Context, stockLineItemId string) LineItemsApiGETStockLineItemIdLineItemRequest {
	return LineItemsApiGETStockLineItemIdLineItemRequest{
		ApiService:      a,
		ctx:             ctx,
		stockLineItemId: stockLineItemId,
	}
}

// Execute executes the request
func (a *LineItemsApiService) GETStockLineItemIdLineItemExecute(r LineItemsApiGETStockLineItemIdLineItemRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineItemsApiService.GETStockLineItemIdLineItem")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_line_items/{stockLineItemId}/line_item"
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

type LineItemsApiGETStockTransferIdLineItemRequest struct {
	ctx             context.Context
	ApiService      *LineItemsApiService
	stockTransferId string
}

func (r LineItemsApiGETStockTransferIdLineItemRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStockTransferIdLineItemExecute(r)
}

/*
GETStockTransferIdLineItem Retrieve the line item associated to the stock transfer

Retrieve the line item associated to the stock transfer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stockTransferId The resource's id
	@return LineItemsApiGETStockTransferIdLineItemRequest
*/
func (a *LineItemsApiService) GETStockTransferIdLineItem(ctx context.Context, stockTransferId string) LineItemsApiGETStockTransferIdLineItemRequest {
	return LineItemsApiGETStockTransferIdLineItemRequest{
		ApiService:      a,
		ctx:             ctx,
		stockTransferId: stockTransferId,
	}
}

// Execute executes the request
func (a *LineItemsApiService) GETStockTransferIdLineItemExecute(r LineItemsApiGETStockTransferIdLineItemRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineItemsApiService.GETStockTransferIdLineItem")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_transfers/{stockTransferId}/line_item"
	localVarPath = strings.Replace(localVarPath, "{"+"stockTransferId"+"}", url.PathEscape(parameterToString(r.stockTransferId, "")), -1)

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

type LineItemsApiPATCHLineItemsLineItemIdRequest struct {
	ctx            context.Context
	ApiService     *LineItemsApiService
	lineItemUpdate *LineItemUpdate
	lineItemId     string
}

func (r LineItemsApiPATCHLineItemsLineItemIdRequest) LineItemUpdate(lineItemUpdate LineItemUpdate) LineItemsApiPATCHLineItemsLineItemIdRequest {
	r.lineItemUpdate = &lineItemUpdate
	return r
}

func (r LineItemsApiPATCHLineItemsLineItemIdRequest) Execute() (*PATCHLineItemsLineItemId200Response, *http.Response, error) {
	return r.ApiService.PATCHLineItemsLineItemIdExecute(r)
}

/*
PATCHLineItemsLineItemId Update a line item

Update a line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param lineItemId The resource's id
	@return LineItemsApiPATCHLineItemsLineItemIdRequest
*/
func (a *LineItemsApiService) PATCHLineItemsLineItemId(ctx context.Context, lineItemId string) LineItemsApiPATCHLineItemsLineItemIdRequest {
	return LineItemsApiPATCHLineItemsLineItemIdRequest{
		ApiService: a,
		ctx:        ctx,
		lineItemId: lineItemId,
	}
}

// Execute executes the request
//
//	@return PATCHLineItemsLineItemId200Response
func (a *LineItemsApiService) PATCHLineItemsLineItemIdExecute(r LineItemsApiPATCHLineItemsLineItemIdRequest) (*PATCHLineItemsLineItemId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHLineItemsLineItemId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineItemsApiService.PATCHLineItemsLineItemId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/line_items/{lineItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"lineItemId"+"}", url.PathEscape(parameterToString(r.lineItemId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lineItemUpdate == nil {
		return localVarReturnValue, nil, reportError("lineItemUpdate is required and must be specified")
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
	localVarPostBody = r.lineItemUpdate
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

type LineItemsApiPOSTLineItemsRequest struct {
	ctx            context.Context
	ApiService     *LineItemsApiService
	lineItemCreate *LineItemCreate
}

func (r LineItemsApiPOSTLineItemsRequest) LineItemCreate(lineItemCreate LineItemCreate) LineItemsApiPOSTLineItemsRequest {
	r.lineItemCreate = &lineItemCreate
	return r
}

func (r LineItemsApiPOSTLineItemsRequest) Execute() (*POSTLineItems201Response, *http.Response, error) {
	return r.ApiService.POSTLineItemsExecute(r)
}

/*
POSTLineItems Create a line item

Create a line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return LineItemsApiPOSTLineItemsRequest
*/
func (a *LineItemsApiService) POSTLineItems(ctx context.Context) LineItemsApiPOSTLineItemsRequest {
	return LineItemsApiPOSTLineItemsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTLineItems201Response
func (a *LineItemsApiService) POSTLineItemsExecute(r LineItemsApiPOSTLineItemsRequest) (*POSTLineItems201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTLineItems201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineItemsApiService.POSTLineItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/line_items"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lineItemCreate == nil {
		return localVarReturnValue, nil, reportError("lineItemCreate is required and must be specified")
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
	localVarPostBody = r.lineItemCreate
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
