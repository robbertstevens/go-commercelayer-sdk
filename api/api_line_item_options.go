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

// LineItemOptionsApiService LineItemOptionsApi service
type LineItemOptionsApiService service

type LineItemOptionsApiDELETELineItemOptionsLineItemOptionIdRequest struct {
	ctx              context.Context
	ApiService       *LineItemOptionsApiService
	lineItemOptionId interface{}
}

func (r LineItemOptionsApiDELETELineItemOptionsLineItemOptionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETELineItemOptionsLineItemOptionIdExecute(r)
}

/*
DELETELineItemOptionsLineItemOptionId Delete a line item option

Delete a line item option

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param lineItemOptionId The resource's id
	@return LineItemOptionsApiDELETELineItemOptionsLineItemOptionIdRequest
*/
func (a *LineItemOptionsApiService) DELETELineItemOptionsLineItemOptionId(ctx context.Context, lineItemOptionId interface{}) LineItemOptionsApiDELETELineItemOptionsLineItemOptionIdRequest {
	return LineItemOptionsApiDELETELineItemOptionsLineItemOptionIdRequest{
		ApiService:       a,
		ctx:              ctx,
		lineItemOptionId: lineItemOptionId,
	}
}

// Execute executes the request
func (a *LineItemOptionsApiService) DELETELineItemOptionsLineItemOptionIdExecute(r LineItemOptionsApiDELETELineItemOptionsLineItemOptionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineItemOptionsApiService.DELETELineItemOptionsLineItemOptionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/line_item_options/{lineItemOptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"lineItemOptionId"+"}", url.PathEscape(parameterValueToString(r.lineItemOptionId, "lineItemOptionId")), -1)

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

type LineItemOptionsApiGETLineItemIdLineItemOptionsRequest struct {
	ctx        context.Context
	ApiService *LineItemOptionsApiService
	lineItemId interface{}
}

func (r LineItemOptionsApiGETLineItemIdLineItemOptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETLineItemIdLineItemOptionsExecute(r)
}

/*
GETLineItemIdLineItemOptions Retrieve the line item options associated to the line item

Retrieve the line item options associated to the line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param lineItemId The resource's id
	@return LineItemOptionsApiGETLineItemIdLineItemOptionsRequest
*/
func (a *LineItemOptionsApiService) GETLineItemIdLineItemOptions(ctx context.Context, lineItemId interface{}) LineItemOptionsApiGETLineItemIdLineItemOptionsRequest {
	return LineItemOptionsApiGETLineItemIdLineItemOptionsRequest{
		ApiService: a,
		ctx:        ctx,
		lineItemId: lineItemId,
	}
}

// Execute executes the request
func (a *LineItemOptionsApiService) GETLineItemIdLineItemOptionsExecute(r LineItemOptionsApiGETLineItemIdLineItemOptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineItemOptionsApiService.GETLineItemIdLineItemOptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/line_items/{lineItemId}/line_item_options"
	localVarPath = strings.Replace(localVarPath, "{"+"lineItemId"+"}", url.PathEscape(parameterValueToString(r.lineItemId, "lineItemId")), -1)

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

type LineItemOptionsApiGETLineItemOptionsRequest struct {
	ctx        context.Context
	ApiService *LineItemOptionsApiService
}

func (r LineItemOptionsApiGETLineItemOptionsRequest) Execute() (*GETLineItemOptions200Response, *http.Response, error) {
	return r.ApiService.GETLineItemOptionsExecute(r)
}

/*
GETLineItemOptions List all line item options

List all line item options

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return LineItemOptionsApiGETLineItemOptionsRequest
*/
func (a *LineItemOptionsApiService) GETLineItemOptions(ctx context.Context) LineItemOptionsApiGETLineItemOptionsRequest {
	return LineItemOptionsApiGETLineItemOptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETLineItemOptions200Response
func (a *LineItemOptionsApiService) GETLineItemOptionsExecute(r LineItemOptionsApiGETLineItemOptionsRequest) (*GETLineItemOptions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETLineItemOptions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineItemOptionsApiService.GETLineItemOptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/line_item_options"

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

type LineItemOptionsApiGETLineItemOptionsLineItemOptionIdRequest struct {
	ctx              context.Context
	ApiService       *LineItemOptionsApiService
	lineItemOptionId interface{}
}

func (r LineItemOptionsApiGETLineItemOptionsLineItemOptionIdRequest) Execute() (*GETLineItemOptionsLineItemOptionId200Response, *http.Response, error) {
	return r.ApiService.GETLineItemOptionsLineItemOptionIdExecute(r)
}

/*
GETLineItemOptionsLineItemOptionId Retrieve a line item option

Retrieve a line item option

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param lineItemOptionId The resource's id
	@return LineItemOptionsApiGETLineItemOptionsLineItemOptionIdRequest
*/
func (a *LineItemOptionsApiService) GETLineItemOptionsLineItemOptionId(ctx context.Context, lineItemOptionId interface{}) LineItemOptionsApiGETLineItemOptionsLineItemOptionIdRequest {
	return LineItemOptionsApiGETLineItemOptionsLineItemOptionIdRequest{
		ApiService:       a,
		ctx:              ctx,
		lineItemOptionId: lineItemOptionId,
	}
}

// Execute executes the request
//
//	@return GETLineItemOptionsLineItemOptionId200Response
func (a *LineItemOptionsApiService) GETLineItemOptionsLineItemOptionIdExecute(r LineItemOptionsApiGETLineItemOptionsLineItemOptionIdRequest) (*GETLineItemOptionsLineItemOptionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETLineItemOptionsLineItemOptionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineItemOptionsApiService.GETLineItemOptionsLineItemOptionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/line_item_options/{lineItemOptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"lineItemOptionId"+"}", url.PathEscape(parameterValueToString(r.lineItemOptionId, "lineItemOptionId")), -1)

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

type LineItemOptionsApiGETOrderIdLineItemOptionsRequest struct {
	ctx        context.Context
	ApiService *LineItemOptionsApiService
	orderId    interface{}
}

func (r LineItemOptionsApiGETOrderIdLineItemOptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderIdLineItemOptionsExecute(r)
}

/*
GETOrderIdLineItemOptions Retrieve the line item options associated to the order

Retrieve the line item options associated to the order

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId The resource's id
	@return LineItemOptionsApiGETOrderIdLineItemOptionsRequest
*/
func (a *LineItemOptionsApiService) GETOrderIdLineItemOptions(ctx context.Context, orderId interface{}) LineItemOptionsApiGETOrderIdLineItemOptionsRequest {
	return LineItemOptionsApiGETOrderIdLineItemOptionsRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *LineItemOptionsApiService) GETOrderIdLineItemOptionsExecute(r LineItemOptionsApiGETOrderIdLineItemOptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineItemOptionsApiService.GETOrderIdLineItemOptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{orderId}/line_item_options"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterValueToString(r.orderId, "orderId")), -1)

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

type LineItemOptionsApiPATCHLineItemOptionsLineItemOptionIdRequest struct {
	ctx                  context.Context
	ApiService           *LineItemOptionsApiService
	lineItemOptionUpdate *LineItemOptionUpdate
	lineItemOptionId     interface{}
}

func (r LineItemOptionsApiPATCHLineItemOptionsLineItemOptionIdRequest) LineItemOptionUpdate(lineItemOptionUpdate LineItemOptionUpdate) LineItemOptionsApiPATCHLineItemOptionsLineItemOptionIdRequest {
	r.lineItemOptionUpdate = &lineItemOptionUpdate
	return r
}

func (r LineItemOptionsApiPATCHLineItemOptionsLineItemOptionIdRequest) Execute() (*PATCHLineItemOptionsLineItemOptionId200Response, *http.Response, error) {
	return r.ApiService.PATCHLineItemOptionsLineItemOptionIdExecute(r)
}

/*
PATCHLineItemOptionsLineItemOptionId Update a line item option

Update a line item option

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param lineItemOptionId The resource's id
	@return LineItemOptionsApiPATCHLineItemOptionsLineItemOptionIdRequest
*/
func (a *LineItemOptionsApiService) PATCHLineItemOptionsLineItemOptionId(ctx context.Context, lineItemOptionId interface{}) LineItemOptionsApiPATCHLineItemOptionsLineItemOptionIdRequest {
	return LineItemOptionsApiPATCHLineItemOptionsLineItemOptionIdRequest{
		ApiService:       a,
		ctx:              ctx,
		lineItemOptionId: lineItemOptionId,
	}
}

// Execute executes the request
//
//	@return PATCHLineItemOptionsLineItemOptionId200Response
func (a *LineItemOptionsApiService) PATCHLineItemOptionsLineItemOptionIdExecute(r LineItemOptionsApiPATCHLineItemOptionsLineItemOptionIdRequest) (*PATCHLineItemOptionsLineItemOptionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHLineItemOptionsLineItemOptionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineItemOptionsApiService.PATCHLineItemOptionsLineItemOptionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/line_item_options/{lineItemOptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"lineItemOptionId"+"}", url.PathEscape(parameterValueToString(r.lineItemOptionId, "lineItemOptionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lineItemOptionUpdate == nil {
		return localVarReturnValue, nil, reportError("lineItemOptionUpdate is required and must be specified")
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
	localVarPostBody = r.lineItemOptionUpdate
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

type LineItemOptionsApiPOSTLineItemOptionsRequest struct {
	ctx                  context.Context
	ApiService           *LineItemOptionsApiService
	lineItemOptionCreate *LineItemOptionCreate
}

func (r LineItemOptionsApiPOSTLineItemOptionsRequest) LineItemOptionCreate(lineItemOptionCreate LineItemOptionCreate) LineItemOptionsApiPOSTLineItemOptionsRequest {
	r.lineItemOptionCreate = &lineItemOptionCreate
	return r
}

func (r LineItemOptionsApiPOSTLineItemOptionsRequest) Execute() (*POSTLineItemOptions201Response, *http.Response, error) {
	return r.ApiService.POSTLineItemOptionsExecute(r)
}

/*
POSTLineItemOptions Create a line item option

Create a line item option

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return LineItemOptionsApiPOSTLineItemOptionsRequest
*/
func (a *LineItemOptionsApiService) POSTLineItemOptions(ctx context.Context) LineItemOptionsApiPOSTLineItemOptionsRequest {
	return LineItemOptionsApiPOSTLineItemOptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTLineItemOptions201Response
func (a *LineItemOptionsApiService) POSTLineItemOptionsExecute(r LineItemOptionsApiPOSTLineItemOptionsRequest) (*POSTLineItemOptions201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTLineItemOptions201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineItemOptionsApiService.POSTLineItemOptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/line_item_options"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lineItemOptionCreate == nil {
		return localVarReturnValue, nil, reportError("lineItemOptionCreate is required and must be specified")
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
	localVarPostBody = r.lineItemOptionCreate
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
