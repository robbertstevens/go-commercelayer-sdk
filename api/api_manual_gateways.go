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

// ManualGatewaysApiService ManualGatewaysApi service
type ManualGatewaysApiService service

type ManualGatewaysApiDELETEManualGatewaysManualGatewayIdRequest struct {
	ctx             context.Context
	ApiService      *ManualGatewaysApiService
	manualGatewayId interface{}
}

func (r ManualGatewaysApiDELETEManualGatewaysManualGatewayIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEManualGatewaysManualGatewayIdExecute(r)
}

/*
DELETEManualGatewaysManualGatewayId Delete a manual gateway

Delete a manual gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param manualGatewayId The resource's id
	@return ManualGatewaysApiDELETEManualGatewaysManualGatewayIdRequest
*/
func (a *ManualGatewaysApiService) DELETEManualGatewaysManualGatewayId(ctx context.Context, manualGatewayId interface{}) ManualGatewaysApiDELETEManualGatewaysManualGatewayIdRequest {
	return ManualGatewaysApiDELETEManualGatewaysManualGatewayIdRequest{
		ApiService:      a,
		ctx:             ctx,
		manualGatewayId: manualGatewayId,
	}
}

// Execute executes the request
func (a *ManualGatewaysApiService) DELETEManualGatewaysManualGatewayIdExecute(r ManualGatewaysApiDELETEManualGatewaysManualGatewayIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManualGatewaysApiService.DELETEManualGatewaysManualGatewayId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/manual_gateways/{manualGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"manualGatewayId"+"}", url.PathEscape(parameterValueToString(r.manualGatewayId, "manualGatewayId")), -1)

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

type ManualGatewaysApiGETManualGatewaysRequest struct {
	ctx        context.Context
	ApiService *ManualGatewaysApiService
}

func (r ManualGatewaysApiGETManualGatewaysRequest) Execute() (*GETManualGateways200Response, *http.Response, error) {
	return r.ApiService.GETManualGatewaysExecute(r)
}

/*
GETManualGateways List all manual gateways

List all manual gateways

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ManualGatewaysApiGETManualGatewaysRequest
*/
func (a *ManualGatewaysApiService) GETManualGateways(ctx context.Context) ManualGatewaysApiGETManualGatewaysRequest {
	return ManualGatewaysApiGETManualGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETManualGateways200Response
func (a *ManualGatewaysApiService) GETManualGatewaysExecute(r ManualGatewaysApiGETManualGatewaysRequest) (*GETManualGateways200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETManualGateways200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManualGatewaysApiService.GETManualGateways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/manual_gateways"

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

type ManualGatewaysApiGETManualGatewaysManualGatewayIdRequest struct {
	ctx             context.Context
	ApiService      *ManualGatewaysApiService
	manualGatewayId interface{}
}

func (r ManualGatewaysApiGETManualGatewaysManualGatewayIdRequest) Execute() (*GETManualGatewaysManualGatewayId200Response, *http.Response, error) {
	return r.ApiService.GETManualGatewaysManualGatewayIdExecute(r)
}

/*
GETManualGatewaysManualGatewayId Retrieve a manual gateway

Retrieve a manual gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param manualGatewayId The resource's id
	@return ManualGatewaysApiGETManualGatewaysManualGatewayIdRequest
*/
func (a *ManualGatewaysApiService) GETManualGatewaysManualGatewayId(ctx context.Context, manualGatewayId interface{}) ManualGatewaysApiGETManualGatewaysManualGatewayIdRequest {
	return ManualGatewaysApiGETManualGatewaysManualGatewayIdRequest{
		ApiService:      a,
		ctx:             ctx,
		manualGatewayId: manualGatewayId,
	}
}

// Execute executes the request
//
//	@return GETManualGatewaysManualGatewayId200Response
func (a *ManualGatewaysApiService) GETManualGatewaysManualGatewayIdExecute(r ManualGatewaysApiGETManualGatewaysManualGatewayIdRequest) (*GETManualGatewaysManualGatewayId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETManualGatewaysManualGatewayId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManualGatewaysApiService.GETManualGatewaysManualGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/manual_gateways/{manualGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"manualGatewayId"+"}", url.PathEscape(parameterValueToString(r.manualGatewayId, "manualGatewayId")), -1)

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

type ManualGatewaysApiPATCHManualGatewaysManualGatewayIdRequest struct {
	ctx                 context.Context
	ApiService          *ManualGatewaysApiService
	manualGatewayUpdate *ManualGatewayUpdate
	manualGatewayId     interface{}
}

func (r ManualGatewaysApiPATCHManualGatewaysManualGatewayIdRequest) ManualGatewayUpdate(manualGatewayUpdate ManualGatewayUpdate) ManualGatewaysApiPATCHManualGatewaysManualGatewayIdRequest {
	r.manualGatewayUpdate = &manualGatewayUpdate
	return r
}

func (r ManualGatewaysApiPATCHManualGatewaysManualGatewayIdRequest) Execute() (*PATCHManualGatewaysManualGatewayId200Response, *http.Response, error) {
	return r.ApiService.PATCHManualGatewaysManualGatewayIdExecute(r)
}

/*
PATCHManualGatewaysManualGatewayId Update a manual gateway

Update a manual gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param manualGatewayId The resource's id
	@return ManualGatewaysApiPATCHManualGatewaysManualGatewayIdRequest
*/
func (a *ManualGatewaysApiService) PATCHManualGatewaysManualGatewayId(ctx context.Context, manualGatewayId interface{}) ManualGatewaysApiPATCHManualGatewaysManualGatewayIdRequest {
	return ManualGatewaysApiPATCHManualGatewaysManualGatewayIdRequest{
		ApiService:      a,
		ctx:             ctx,
		manualGatewayId: manualGatewayId,
	}
}

// Execute executes the request
//
//	@return PATCHManualGatewaysManualGatewayId200Response
func (a *ManualGatewaysApiService) PATCHManualGatewaysManualGatewayIdExecute(r ManualGatewaysApiPATCHManualGatewaysManualGatewayIdRequest) (*PATCHManualGatewaysManualGatewayId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHManualGatewaysManualGatewayId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManualGatewaysApiService.PATCHManualGatewaysManualGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/manual_gateways/{manualGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"manualGatewayId"+"}", url.PathEscape(parameterValueToString(r.manualGatewayId, "manualGatewayId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.manualGatewayUpdate == nil {
		return localVarReturnValue, nil, reportError("manualGatewayUpdate is required and must be specified")
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
	localVarPostBody = r.manualGatewayUpdate
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

type ManualGatewaysApiPOSTManualGatewaysRequest struct {
	ctx                 context.Context
	ApiService          *ManualGatewaysApiService
	manualGatewayCreate *ManualGatewayCreate
}

func (r ManualGatewaysApiPOSTManualGatewaysRequest) ManualGatewayCreate(manualGatewayCreate ManualGatewayCreate) ManualGatewaysApiPOSTManualGatewaysRequest {
	r.manualGatewayCreate = &manualGatewayCreate
	return r
}

func (r ManualGatewaysApiPOSTManualGatewaysRequest) Execute() (*POSTManualGateways201Response, *http.Response, error) {
	return r.ApiService.POSTManualGatewaysExecute(r)
}

/*
POSTManualGateways Create a manual gateway

Create a manual gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ManualGatewaysApiPOSTManualGatewaysRequest
*/
func (a *ManualGatewaysApiService) POSTManualGateways(ctx context.Context) ManualGatewaysApiPOSTManualGatewaysRequest {
	return ManualGatewaysApiPOSTManualGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTManualGateways201Response
func (a *ManualGatewaysApiService) POSTManualGatewaysExecute(r ManualGatewaysApiPOSTManualGatewaysRequest) (*POSTManualGateways201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTManualGateways201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManualGatewaysApiService.POSTManualGateways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/manual_gateways"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.manualGatewayCreate == nil {
		return localVarReturnValue, nil, reportError("manualGatewayCreate is required and must be specified")
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
	localVarPostBody = r.manualGatewayCreate
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
