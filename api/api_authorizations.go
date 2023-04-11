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

// AuthorizationsApiService AuthorizationsApi service
type AuthorizationsApiService service

type AuthorizationsApiGETAuthorizationsRequest struct {
	ctx        context.Context
	ApiService *AuthorizationsApiService
}

func (r AuthorizationsApiGETAuthorizationsRequest) Execute() (*GETAuthorizations200Response, *http.Response, error) {
	return r.ApiService.GETAuthorizationsExecute(r)
}

/*
GETAuthorizations List all authorizations

List all authorizations

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AuthorizationsApiGETAuthorizationsRequest
*/
func (a *AuthorizationsApiService) GETAuthorizations(ctx context.Context) AuthorizationsApiGETAuthorizationsRequest {
	return AuthorizationsApiGETAuthorizationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETAuthorizations200Response
func (a *AuthorizationsApiService) GETAuthorizationsExecute(r AuthorizationsApiGETAuthorizationsRequest) (*GETAuthorizations200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETAuthorizations200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationsApiService.GETAuthorizations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/authorizations"

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

type AuthorizationsApiGETAuthorizationsAuthorizationIdRequest struct {
	ctx             context.Context
	ApiService      *AuthorizationsApiService
	authorizationId interface{}
}

func (r AuthorizationsApiGETAuthorizationsAuthorizationIdRequest) Execute() (*GETAuthorizationsAuthorizationId200Response, *http.Response, error) {
	return r.ApiService.GETAuthorizationsAuthorizationIdExecute(r)
}

/*
GETAuthorizationsAuthorizationId Retrieve an authorization

Retrieve an authorization

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authorizationId The resource's id
	@return AuthorizationsApiGETAuthorizationsAuthorizationIdRequest
*/
func (a *AuthorizationsApiService) GETAuthorizationsAuthorizationId(ctx context.Context, authorizationId interface{}) AuthorizationsApiGETAuthorizationsAuthorizationIdRequest {
	return AuthorizationsApiGETAuthorizationsAuthorizationIdRequest{
		ApiService:      a,
		ctx:             ctx,
		authorizationId: authorizationId,
	}
}

// Execute executes the request
//
//	@return GETAuthorizationsAuthorizationId200Response
func (a *AuthorizationsApiService) GETAuthorizationsAuthorizationIdExecute(r AuthorizationsApiGETAuthorizationsAuthorizationIdRequest) (*GETAuthorizationsAuthorizationId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETAuthorizationsAuthorizationId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationsApiService.GETAuthorizationsAuthorizationId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/authorizations/{authorizationId}"
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

type AuthorizationsApiGETCaptureIdReferenceAuthorizationRequest struct {
	ctx        context.Context
	ApiService *AuthorizationsApiService
	captureId  interface{}
}

func (r AuthorizationsApiGETCaptureIdReferenceAuthorizationRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCaptureIdReferenceAuthorizationExecute(r)
}

/*
GETCaptureIdReferenceAuthorization Retrieve the reference authorization associated to the capture

Retrieve the reference authorization associated to the capture

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param captureId The resource's id
	@return AuthorizationsApiGETCaptureIdReferenceAuthorizationRequest
*/
func (a *AuthorizationsApiService) GETCaptureIdReferenceAuthorization(ctx context.Context, captureId interface{}) AuthorizationsApiGETCaptureIdReferenceAuthorizationRequest {
	return AuthorizationsApiGETCaptureIdReferenceAuthorizationRequest{
		ApiService: a,
		ctx:        ctx,
		captureId:  captureId,
	}
}

// Execute executes the request
func (a *AuthorizationsApiService) GETCaptureIdReferenceAuthorizationExecute(r AuthorizationsApiGETCaptureIdReferenceAuthorizationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationsApiService.GETCaptureIdReferenceAuthorization")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/captures/{captureId}/reference_authorization"
	localVarPath = strings.Replace(localVarPath, "{"+"captureId"+"}", url.PathEscape(parameterValueToString(r.captureId, "captureId")), -1)

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

type AuthorizationsApiGETOrderIdAuthorizationsRequest struct {
	ctx        context.Context
	ApiService *AuthorizationsApiService
	orderId    interface{}
}

func (r AuthorizationsApiGETOrderIdAuthorizationsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderIdAuthorizationsExecute(r)
}

/*
GETOrderIdAuthorizations Retrieve the authorizations associated to the order

Retrieve the authorizations associated to the order

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId The resource's id
	@return AuthorizationsApiGETOrderIdAuthorizationsRequest
*/
func (a *AuthorizationsApiService) GETOrderIdAuthorizations(ctx context.Context, orderId interface{}) AuthorizationsApiGETOrderIdAuthorizationsRequest {
	return AuthorizationsApiGETOrderIdAuthorizationsRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *AuthorizationsApiService) GETOrderIdAuthorizationsExecute(r AuthorizationsApiGETOrderIdAuthorizationsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationsApiService.GETOrderIdAuthorizations")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{orderId}/authorizations"
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

type AuthorizationsApiGETVoidIdReferenceAuthorizationRequest struct {
	ctx        context.Context
	ApiService *AuthorizationsApiService
	voidId     interface{}
}

func (r AuthorizationsApiGETVoidIdReferenceAuthorizationRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETVoidIdReferenceAuthorizationExecute(r)
}

/*
GETVoidIdReferenceAuthorization Retrieve the reference authorization associated to the void

Retrieve the reference authorization associated to the void

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param voidId The resource's id
	@return AuthorizationsApiGETVoidIdReferenceAuthorizationRequest
*/
func (a *AuthorizationsApiService) GETVoidIdReferenceAuthorization(ctx context.Context, voidId interface{}) AuthorizationsApiGETVoidIdReferenceAuthorizationRequest {
	return AuthorizationsApiGETVoidIdReferenceAuthorizationRequest{
		ApiService: a,
		ctx:        ctx,
		voidId:     voidId,
	}
}

// Execute executes the request
func (a *AuthorizationsApiService) GETVoidIdReferenceAuthorizationExecute(r AuthorizationsApiGETVoidIdReferenceAuthorizationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationsApiService.GETVoidIdReferenceAuthorization")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/voids/{voidId}/reference_authorization"
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

type AuthorizationsApiPATCHAuthorizationsAuthorizationIdRequest struct {
	ctx                                       context.Context
	ApiService                                *AuthorizationsApiService
	pATCHAuthorizationsAuthorizationIdRequest *PATCHAuthorizationsAuthorizationIdRequest
	authorizationId                           interface{}
}

func (r AuthorizationsApiPATCHAuthorizationsAuthorizationIdRequest) PATCHAuthorizationsAuthorizationIdRequest(pATCHAuthorizationsAuthorizationIdRequest PATCHAuthorizationsAuthorizationIdRequest) AuthorizationsApiPATCHAuthorizationsAuthorizationIdRequest {
	r.pATCHAuthorizationsAuthorizationIdRequest = &pATCHAuthorizationsAuthorizationIdRequest
	return r
}

func (r AuthorizationsApiPATCHAuthorizationsAuthorizationIdRequest) Execute() (*PATCHAuthorizationsAuthorizationId200Response, *http.Response, error) {
	return r.ApiService.PATCHAuthorizationsAuthorizationIdExecute(r)
}

/*
PATCHAuthorizationsAuthorizationId Update an authorization

Update an authorization

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authorizationId The resource's id
	@return AuthorizationsApiPATCHAuthorizationsAuthorizationIdRequest
*/
func (a *AuthorizationsApiService) PATCHAuthorizationsAuthorizationId(ctx context.Context, authorizationId interface{}) AuthorizationsApiPATCHAuthorizationsAuthorizationIdRequest {
	return AuthorizationsApiPATCHAuthorizationsAuthorizationIdRequest{
		ApiService:      a,
		ctx:             ctx,
		authorizationId: authorizationId,
	}
}

// Execute executes the request
//
//	@return PATCHAuthorizationsAuthorizationId200Response
func (a *AuthorizationsApiService) PATCHAuthorizationsAuthorizationIdExecute(r AuthorizationsApiPATCHAuthorizationsAuthorizationIdRequest) (*PATCHAuthorizationsAuthorizationId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHAuthorizationsAuthorizationId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationsApiService.PATCHAuthorizationsAuthorizationId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/authorizations/{authorizationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"authorizationId"+"}", url.PathEscape(parameterValueToString(r.authorizationId, "authorizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pATCHAuthorizationsAuthorizationIdRequest == nil {
		return localVarReturnValue, nil, reportError("pATCHAuthorizationsAuthorizationIdRequest is required and must be specified")
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
	localVarPostBody = r.pATCHAuthorizationsAuthorizationIdRequest
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
