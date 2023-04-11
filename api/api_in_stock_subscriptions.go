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

// InStockSubscriptionsApiService InStockSubscriptionsApi service
type InStockSubscriptionsApiService service

type InStockSubscriptionsApiDELETEInStockSubscriptionsInStockSubscriptionIdRequest struct {
	ctx                   context.Context
	ApiService            *InStockSubscriptionsApiService
	inStockSubscriptionId interface{}
}

func (r InStockSubscriptionsApiDELETEInStockSubscriptionsInStockSubscriptionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEInStockSubscriptionsInStockSubscriptionIdExecute(r)
}

/*
DELETEInStockSubscriptionsInStockSubscriptionId Delete an in stock subscription

Delete an in stock subscription

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inStockSubscriptionId The resource's id
	@return InStockSubscriptionsApiDELETEInStockSubscriptionsInStockSubscriptionIdRequest
*/
func (a *InStockSubscriptionsApiService) DELETEInStockSubscriptionsInStockSubscriptionId(ctx context.Context, inStockSubscriptionId interface{}) InStockSubscriptionsApiDELETEInStockSubscriptionsInStockSubscriptionIdRequest {
	return InStockSubscriptionsApiDELETEInStockSubscriptionsInStockSubscriptionIdRequest{
		ApiService:            a,
		ctx:                   ctx,
		inStockSubscriptionId: inStockSubscriptionId,
	}
}

// Execute executes the request
func (a *InStockSubscriptionsApiService) DELETEInStockSubscriptionsInStockSubscriptionIdExecute(r InStockSubscriptionsApiDELETEInStockSubscriptionsInStockSubscriptionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InStockSubscriptionsApiService.DELETEInStockSubscriptionsInStockSubscriptionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/in_stock_subscriptions/{inStockSubscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inStockSubscriptionId"+"}", url.PathEscape(parameterValueToString(r.inStockSubscriptionId, "inStockSubscriptionId")), -1)

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

type InStockSubscriptionsApiGETInStockSubscriptionsRequest struct {
	ctx        context.Context
	ApiService *InStockSubscriptionsApiService
}

func (r InStockSubscriptionsApiGETInStockSubscriptionsRequest) Execute() (*GETInStockSubscriptions200Response, *http.Response, error) {
	return r.ApiService.GETInStockSubscriptionsExecute(r)
}

/*
GETInStockSubscriptions List all in stock subscriptions

List all in stock subscriptions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return InStockSubscriptionsApiGETInStockSubscriptionsRequest
*/
func (a *InStockSubscriptionsApiService) GETInStockSubscriptions(ctx context.Context) InStockSubscriptionsApiGETInStockSubscriptionsRequest {
	return InStockSubscriptionsApiGETInStockSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETInStockSubscriptions200Response
func (a *InStockSubscriptionsApiService) GETInStockSubscriptionsExecute(r InStockSubscriptionsApiGETInStockSubscriptionsRequest) (*GETInStockSubscriptions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETInStockSubscriptions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InStockSubscriptionsApiService.GETInStockSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/in_stock_subscriptions"

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

type InStockSubscriptionsApiGETInStockSubscriptionsInStockSubscriptionIdRequest struct {
	ctx                   context.Context
	ApiService            *InStockSubscriptionsApiService
	inStockSubscriptionId interface{}
}

func (r InStockSubscriptionsApiGETInStockSubscriptionsInStockSubscriptionIdRequest) Execute() (*GETInStockSubscriptionsInStockSubscriptionId200Response, *http.Response, error) {
	return r.ApiService.GETInStockSubscriptionsInStockSubscriptionIdExecute(r)
}

/*
GETInStockSubscriptionsInStockSubscriptionId Retrieve an in stock subscription

Retrieve an in stock subscription

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inStockSubscriptionId The resource's id
	@return InStockSubscriptionsApiGETInStockSubscriptionsInStockSubscriptionIdRequest
*/
func (a *InStockSubscriptionsApiService) GETInStockSubscriptionsInStockSubscriptionId(ctx context.Context, inStockSubscriptionId interface{}) InStockSubscriptionsApiGETInStockSubscriptionsInStockSubscriptionIdRequest {
	return InStockSubscriptionsApiGETInStockSubscriptionsInStockSubscriptionIdRequest{
		ApiService:            a,
		ctx:                   ctx,
		inStockSubscriptionId: inStockSubscriptionId,
	}
}

// Execute executes the request
//
//	@return GETInStockSubscriptionsInStockSubscriptionId200Response
func (a *InStockSubscriptionsApiService) GETInStockSubscriptionsInStockSubscriptionIdExecute(r InStockSubscriptionsApiGETInStockSubscriptionsInStockSubscriptionIdRequest) (*GETInStockSubscriptionsInStockSubscriptionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETInStockSubscriptionsInStockSubscriptionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InStockSubscriptionsApiService.GETInStockSubscriptionsInStockSubscriptionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/in_stock_subscriptions/{inStockSubscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inStockSubscriptionId"+"}", url.PathEscape(parameterValueToString(r.inStockSubscriptionId, "inStockSubscriptionId")), -1)

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

type InStockSubscriptionsApiPATCHInStockSubscriptionsInStockSubscriptionIdRequest struct {
	ctx                                                   context.Context
	ApiService                                            *InStockSubscriptionsApiService
	pATCHInStockSubscriptionsInStockSubscriptionIdRequest *PATCHInStockSubscriptionsInStockSubscriptionIdRequest
	inStockSubscriptionId                                 interface{}
}

func (r InStockSubscriptionsApiPATCHInStockSubscriptionsInStockSubscriptionIdRequest) PATCHInStockSubscriptionsInStockSubscriptionIdRequest(pATCHInStockSubscriptionsInStockSubscriptionIdRequest PATCHInStockSubscriptionsInStockSubscriptionIdRequest) InStockSubscriptionsApiPATCHInStockSubscriptionsInStockSubscriptionIdRequest {
	r.pATCHInStockSubscriptionsInStockSubscriptionIdRequest = &pATCHInStockSubscriptionsInStockSubscriptionIdRequest
	return r
}

func (r InStockSubscriptionsApiPATCHInStockSubscriptionsInStockSubscriptionIdRequest) Execute() (*PATCHInStockSubscriptionsInStockSubscriptionId200Response, *http.Response, error) {
	return r.ApiService.PATCHInStockSubscriptionsInStockSubscriptionIdExecute(r)
}

/*
PATCHInStockSubscriptionsInStockSubscriptionId Update an in stock subscription

Update an in stock subscription

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inStockSubscriptionId The resource's id
	@return InStockSubscriptionsApiPATCHInStockSubscriptionsInStockSubscriptionIdRequest
*/
func (a *InStockSubscriptionsApiService) PATCHInStockSubscriptionsInStockSubscriptionId(ctx context.Context, inStockSubscriptionId interface{}) InStockSubscriptionsApiPATCHInStockSubscriptionsInStockSubscriptionIdRequest {
	return InStockSubscriptionsApiPATCHInStockSubscriptionsInStockSubscriptionIdRequest{
		ApiService:            a,
		ctx:                   ctx,
		inStockSubscriptionId: inStockSubscriptionId,
	}
}

// Execute executes the request
//
//	@return PATCHInStockSubscriptionsInStockSubscriptionId200Response
func (a *InStockSubscriptionsApiService) PATCHInStockSubscriptionsInStockSubscriptionIdExecute(r InStockSubscriptionsApiPATCHInStockSubscriptionsInStockSubscriptionIdRequest) (*PATCHInStockSubscriptionsInStockSubscriptionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHInStockSubscriptionsInStockSubscriptionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InStockSubscriptionsApiService.PATCHInStockSubscriptionsInStockSubscriptionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/in_stock_subscriptions/{inStockSubscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inStockSubscriptionId"+"}", url.PathEscape(parameterValueToString(r.inStockSubscriptionId, "inStockSubscriptionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pATCHInStockSubscriptionsInStockSubscriptionIdRequest == nil {
		return localVarReturnValue, nil, reportError("pATCHInStockSubscriptionsInStockSubscriptionIdRequest is required and must be specified")
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
	localVarPostBody = r.pATCHInStockSubscriptionsInStockSubscriptionIdRequest
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

type InStockSubscriptionsApiPOSTInStockSubscriptionsRequest struct {
	ctx                             context.Context
	ApiService                      *InStockSubscriptionsApiService
	pOSTInStockSubscriptionsRequest *POSTInStockSubscriptionsRequest
}

func (r InStockSubscriptionsApiPOSTInStockSubscriptionsRequest) POSTInStockSubscriptionsRequest(pOSTInStockSubscriptionsRequest POSTInStockSubscriptionsRequest) InStockSubscriptionsApiPOSTInStockSubscriptionsRequest {
	r.pOSTInStockSubscriptionsRequest = &pOSTInStockSubscriptionsRequest
	return r
}

func (r InStockSubscriptionsApiPOSTInStockSubscriptionsRequest) Execute() (*POSTInStockSubscriptions201Response, *http.Response, error) {
	return r.ApiService.POSTInStockSubscriptionsExecute(r)
}

/*
POSTInStockSubscriptions Create an in stock subscription

Create an in stock subscription

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return InStockSubscriptionsApiPOSTInStockSubscriptionsRequest
*/
func (a *InStockSubscriptionsApiService) POSTInStockSubscriptions(ctx context.Context) InStockSubscriptionsApiPOSTInStockSubscriptionsRequest {
	return InStockSubscriptionsApiPOSTInStockSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTInStockSubscriptions201Response
func (a *InStockSubscriptionsApiService) POSTInStockSubscriptionsExecute(r InStockSubscriptionsApiPOSTInStockSubscriptionsRequest) (*POSTInStockSubscriptions201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTInStockSubscriptions201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InStockSubscriptionsApiService.POSTInStockSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/in_stock_subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pOSTInStockSubscriptionsRequest == nil {
		return localVarReturnValue, nil, reportError("pOSTInStockSubscriptionsRequest is required and must be specified")
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
	localVarPostBody = r.pOSTInStockSubscriptionsRequest
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
