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

// VoidsApiService VoidsApi service
type VoidsApiService service

type VoidsApiGETAuthorizationIdVoidsRequest struct {
	ctx             context.Context
	ApiService      *VoidsApiService
	authorizationId interface{}
}

func (r VoidsApiGETAuthorizationIdVoidsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETAuthorizationIdVoidsExecute(r)
}

/*
GETAuthorizationIdVoids Retrieve the voids associated to the authorization

Retrieve the voids associated to the authorization

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authorizationId The resource's id
	@return VoidsApiGETAuthorizationIdVoidsRequest
*/
func (a *VoidsApiService) GETAuthorizationIdVoids(ctx context.Context, authorizationId interface{}) VoidsApiGETAuthorizationIdVoidsRequest {
	return VoidsApiGETAuthorizationIdVoidsRequest{
		ApiService:      a,
		ctx:             ctx,
		authorizationId: authorizationId,
	}
}

// Execute executes the request
func (a *VoidsApiService) GETAuthorizationIdVoidsExecute(r VoidsApiGETAuthorizationIdVoidsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VoidsApiService.GETAuthorizationIdVoids")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/authorizations/{authorizationId}/voids"
	localVarPath = strings.Replace(localVarPath, "{"+"authorizationId"+"}", url.PathEscape(parameterValueToString(r.authorizationId, "authorizationId")), -1)

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

type VoidsApiGETOrderIdVoidsRequest struct {
	ctx        context.Context
	ApiService *VoidsApiService
	orderId    interface{}
}

func (r VoidsApiGETOrderIdVoidsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderIdVoidsExecute(r)
}

/*
GETOrderIdVoids Retrieve the voids associated to the order

Retrieve the voids associated to the order

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId The resource's id
	@return VoidsApiGETOrderIdVoidsRequest
*/
func (a *VoidsApiService) GETOrderIdVoids(ctx context.Context, orderId interface{}) VoidsApiGETOrderIdVoidsRequest {
	return VoidsApiGETOrderIdVoidsRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *VoidsApiService) GETOrderIdVoidsExecute(r VoidsApiGETOrderIdVoidsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VoidsApiService.GETOrderIdVoids")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{orderId}/voids"
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

type VoidsApiGETVoidsRequest struct {
	ctx        context.Context
	ApiService *VoidsApiService
}

func (r VoidsApiGETVoidsRequest) Execute() (*GETVoids200Response, *http.Response, error) {
	return r.ApiService.GETVoidsExecute(r)
}

/*
GETVoids List all voids

List all voids

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return VoidsApiGETVoidsRequest
*/
func (a *VoidsApiService) GETVoids(ctx context.Context) VoidsApiGETVoidsRequest {
	return VoidsApiGETVoidsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETVoids200Response
func (a *VoidsApiService) GETVoidsExecute(r VoidsApiGETVoidsRequest) (*GETVoids200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETVoids200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VoidsApiService.GETVoids")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/voids"

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

type VoidsApiGETVoidsVoidIdRequest struct {
	ctx        context.Context
	ApiService *VoidsApiService
	voidId     interface{}
}

func (r VoidsApiGETVoidsVoidIdRequest) Execute() (*GETVoidsVoidId200Response, *http.Response, error) {
	return r.ApiService.GETVoidsVoidIdExecute(r)
}

/*
GETVoidsVoidId Retrieve a void

Retrieve a void

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param voidId The resource's id
	@return VoidsApiGETVoidsVoidIdRequest
*/
func (a *VoidsApiService) GETVoidsVoidId(ctx context.Context, voidId interface{}) VoidsApiGETVoidsVoidIdRequest {
	return VoidsApiGETVoidsVoidIdRequest{
		ApiService: a,
		ctx:        ctx,
		voidId:     voidId,
	}
}

// Execute executes the request
//
//	@return GETVoidsVoidId200Response
func (a *VoidsApiService) GETVoidsVoidIdExecute(r VoidsApiGETVoidsVoidIdRequest) (*GETVoidsVoidId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETVoidsVoidId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VoidsApiService.GETVoidsVoidId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/voids/{voidId}"
	localVarPath = strings.Replace(localVarPath, "{"+"voidId"+"}", url.PathEscape(parameterValueToString(r.voidId, "voidId")), -1)

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

type VoidsApiPATCHVoidsVoidIdRequest struct {
	ctx        context.Context
	ApiService *VoidsApiService
	voidUpdate *VoidUpdate
	voidId     interface{}
}

func (r VoidsApiPATCHVoidsVoidIdRequest) VoidUpdate(voidUpdate VoidUpdate) VoidsApiPATCHVoidsVoidIdRequest {
	r.voidUpdate = &voidUpdate
	return r
}

func (r VoidsApiPATCHVoidsVoidIdRequest) Execute() (*PATCHVoidsVoidId200Response, *http.Response, error) {
	return r.ApiService.PATCHVoidsVoidIdExecute(r)
}

/*
PATCHVoidsVoidId Update a void

Update a void

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param voidId The resource's id
	@return VoidsApiPATCHVoidsVoidIdRequest
*/
func (a *VoidsApiService) PATCHVoidsVoidId(ctx context.Context, voidId interface{}) VoidsApiPATCHVoidsVoidIdRequest {
	return VoidsApiPATCHVoidsVoidIdRequest{
		ApiService: a,
		ctx:        ctx,
		voidId:     voidId,
	}
}

// Execute executes the request
//
//	@return PATCHVoidsVoidId200Response
func (a *VoidsApiService) PATCHVoidsVoidIdExecute(r VoidsApiPATCHVoidsVoidIdRequest) (*PATCHVoidsVoidId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHVoidsVoidId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VoidsApiService.PATCHVoidsVoidId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/voids/{voidId}"
	localVarPath = strings.Replace(localVarPath, "{"+"voidId"+"}", url.PathEscape(parameterValueToString(r.voidId, "voidId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.voidUpdate == nil {
		return localVarReturnValue, nil, reportError("voidUpdate is required and must be specified")
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
	localVarPostBody = r.voidUpdate
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
