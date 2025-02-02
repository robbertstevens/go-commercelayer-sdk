/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
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

// WireTransfersApiService WireTransfersApi service
type WireTransfersApiService service

type WireTransfersApiDELETEWireTransfersWireTransferIdRequest struct {
	ctx            context.Context
	ApiService     *WireTransfersApiService
	wireTransferId interface{}
}

func (r WireTransfersApiDELETEWireTransfersWireTransferIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEWireTransfersWireTransferIdExecute(r)
}

/*
DELETEWireTransfersWireTransferId Delete a wire transfer

Delete a wire transfer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param wireTransferId The resource's id
	@return WireTransfersApiDELETEWireTransfersWireTransferIdRequest
*/
func (a *WireTransfersApiService) DELETEWireTransfersWireTransferId(ctx context.Context, wireTransferId interface{}) WireTransfersApiDELETEWireTransfersWireTransferIdRequest {
	return WireTransfersApiDELETEWireTransfersWireTransferIdRequest{
		ApiService:     a,
		ctx:            ctx,
		wireTransferId: wireTransferId,
	}
}

// Execute executes the request
func (a *WireTransfersApiService) DELETEWireTransfersWireTransferIdExecute(r WireTransfersApiDELETEWireTransfersWireTransferIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WireTransfersApiService.DELETEWireTransfersWireTransferId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wire_transfers/{wireTransferId}"
	localVarPath = strings.Replace(localVarPath, "{"+"wireTransferId"+"}", url.PathEscape(parameterValueToString(r.wireTransferId, "wireTransferId")), -1)

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

type WireTransfersApiGETWireTransfersRequest struct {
	ctx        context.Context
	ApiService *WireTransfersApiService
}

func (r WireTransfersApiGETWireTransfersRequest) Execute() (*GETWireTransfers200Response, *http.Response, error) {
	return r.ApiService.GETWireTransfersExecute(r)
}

/*
GETWireTransfers List all wire transfers

List all wire transfers

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return WireTransfersApiGETWireTransfersRequest
*/
func (a *WireTransfersApiService) GETWireTransfers(ctx context.Context) WireTransfersApiGETWireTransfersRequest {
	return WireTransfersApiGETWireTransfersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETWireTransfers200Response
func (a *WireTransfersApiService) GETWireTransfersExecute(r WireTransfersApiGETWireTransfersRequest) (*GETWireTransfers200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETWireTransfers200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WireTransfersApiService.GETWireTransfers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wire_transfers"

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

type WireTransfersApiGETWireTransfersWireTransferIdRequest struct {
	ctx            context.Context
	ApiService     *WireTransfersApiService
	wireTransferId interface{}
}

func (r WireTransfersApiGETWireTransfersWireTransferIdRequest) Execute() (*GETWireTransfersWireTransferId200Response, *http.Response, error) {
	return r.ApiService.GETWireTransfersWireTransferIdExecute(r)
}

/*
GETWireTransfersWireTransferId Retrieve a wire transfer

Retrieve a wire transfer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param wireTransferId The resource's id
	@return WireTransfersApiGETWireTransfersWireTransferIdRequest
*/
func (a *WireTransfersApiService) GETWireTransfersWireTransferId(ctx context.Context, wireTransferId interface{}) WireTransfersApiGETWireTransfersWireTransferIdRequest {
	return WireTransfersApiGETWireTransfersWireTransferIdRequest{
		ApiService:     a,
		ctx:            ctx,
		wireTransferId: wireTransferId,
	}
}

// Execute executes the request
//
//	@return GETWireTransfersWireTransferId200Response
func (a *WireTransfersApiService) GETWireTransfersWireTransferIdExecute(r WireTransfersApiGETWireTransfersWireTransferIdRequest) (*GETWireTransfersWireTransferId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETWireTransfersWireTransferId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WireTransfersApiService.GETWireTransfersWireTransferId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wire_transfers/{wireTransferId}"
	localVarPath = strings.Replace(localVarPath, "{"+"wireTransferId"+"}", url.PathEscape(parameterValueToString(r.wireTransferId, "wireTransferId")), -1)

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

type WireTransfersApiPATCHWireTransfersWireTransferIdRequest struct {
	ctx                context.Context
	ApiService         *WireTransfersApiService
	wireTransferUpdate *WireTransferUpdate
	wireTransferId     interface{}
}

func (r WireTransfersApiPATCHWireTransfersWireTransferIdRequest) WireTransferUpdate(wireTransferUpdate WireTransferUpdate) WireTransfersApiPATCHWireTransfersWireTransferIdRequest {
	r.wireTransferUpdate = &wireTransferUpdate
	return r
}

func (r WireTransfersApiPATCHWireTransfersWireTransferIdRequest) Execute() (*PATCHWireTransfersWireTransferId200Response, *http.Response, error) {
	return r.ApiService.PATCHWireTransfersWireTransferIdExecute(r)
}

/*
PATCHWireTransfersWireTransferId Update a wire transfer

Update a wire transfer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param wireTransferId The resource's id
	@return WireTransfersApiPATCHWireTransfersWireTransferIdRequest
*/
func (a *WireTransfersApiService) PATCHWireTransfersWireTransferId(ctx context.Context, wireTransferId interface{}) WireTransfersApiPATCHWireTransfersWireTransferIdRequest {
	return WireTransfersApiPATCHWireTransfersWireTransferIdRequest{
		ApiService:     a,
		ctx:            ctx,
		wireTransferId: wireTransferId,
	}
}

// Execute executes the request
//
//	@return PATCHWireTransfersWireTransferId200Response
func (a *WireTransfersApiService) PATCHWireTransfersWireTransferIdExecute(r WireTransfersApiPATCHWireTransfersWireTransferIdRequest) (*PATCHWireTransfersWireTransferId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHWireTransfersWireTransferId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WireTransfersApiService.PATCHWireTransfersWireTransferId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wire_transfers/{wireTransferId}"
	localVarPath = strings.Replace(localVarPath, "{"+"wireTransferId"+"}", url.PathEscape(parameterValueToString(r.wireTransferId, "wireTransferId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.wireTransferUpdate == nil {
		return localVarReturnValue, nil, reportError("wireTransferUpdate is required and must be specified")
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
	localVarPostBody = r.wireTransferUpdate
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

type WireTransfersApiPOSTWireTransfersRequest struct {
	ctx                context.Context
	ApiService         *WireTransfersApiService
	wireTransferCreate *WireTransferCreate
}

func (r WireTransfersApiPOSTWireTransfersRequest) WireTransferCreate(wireTransferCreate WireTransferCreate) WireTransfersApiPOSTWireTransfersRequest {
	r.wireTransferCreate = &wireTransferCreate
	return r
}

func (r WireTransfersApiPOSTWireTransfersRequest) Execute() (*POSTWireTransfers201Response, *http.Response, error) {
	return r.ApiService.POSTWireTransfersExecute(r)
}

/*
POSTWireTransfers Create a wire transfer

Create a wire transfer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return WireTransfersApiPOSTWireTransfersRequest
*/
func (a *WireTransfersApiService) POSTWireTransfers(ctx context.Context) WireTransfersApiPOSTWireTransfersRequest {
	return WireTransfersApiPOSTWireTransfersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTWireTransfers201Response
func (a *WireTransfersApiService) POSTWireTransfersExecute(r WireTransfersApiPOSTWireTransfersRequest) (*POSTWireTransfers201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTWireTransfers201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WireTransfersApiService.POSTWireTransfers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wire_transfers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.wireTransferCreate == nil {
		return localVarReturnValue, nil, reportError("wireTransferCreate is required and must be specified")
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
	localVarPostBody = r.wireTransferCreate
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
