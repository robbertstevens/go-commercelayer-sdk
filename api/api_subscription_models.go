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

// SubscriptionModelsApiService SubscriptionModelsApi service
type SubscriptionModelsApiService service

type SubscriptionModelsApiDELETESubscriptionModelsSubscriptionModelIdRequest struct {
	ctx                 context.Context
	ApiService          *SubscriptionModelsApiService
	subscriptionModelId interface{}
}

func (r SubscriptionModelsApiDELETESubscriptionModelsSubscriptionModelIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETESubscriptionModelsSubscriptionModelIdExecute(r)
}

/*
DELETESubscriptionModelsSubscriptionModelId Delete a subscription model

Delete a subscription model

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param subscriptionModelId The resource's id
	@return SubscriptionModelsApiDELETESubscriptionModelsSubscriptionModelIdRequest
*/
func (a *SubscriptionModelsApiService) DELETESubscriptionModelsSubscriptionModelId(ctx context.Context, subscriptionModelId interface{}) SubscriptionModelsApiDELETESubscriptionModelsSubscriptionModelIdRequest {
	return SubscriptionModelsApiDELETESubscriptionModelsSubscriptionModelIdRequest{
		ApiService:          a,
		ctx:                 ctx,
		subscriptionModelId: subscriptionModelId,
	}
}

// Execute executes the request
func (a *SubscriptionModelsApiService) DELETESubscriptionModelsSubscriptionModelIdExecute(r SubscriptionModelsApiDELETESubscriptionModelsSubscriptionModelIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionModelsApiService.DELETESubscriptionModelsSubscriptionModelId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription_models/{subscriptionModelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionModelId"+"}", url.PathEscape(parameterValueToString(r.subscriptionModelId, "subscriptionModelId")), -1)

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

type SubscriptionModelsApiGETMarketIdSubscriptionModelRequest struct {
	ctx        context.Context
	ApiService *SubscriptionModelsApiService
	marketId   interface{}
}

func (r SubscriptionModelsApiGETMarketIdSubscriptionModelRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETMarketIdSubscriptionModelExecute(r)
}

/*
GETMarketIdSubscriptionModel Retrieve the subscription model associated to the market

Retrieve the subscription model associated to the market

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param marketId The resource's id
	@return SubscriptionModelsApiGETMarketIdSubscriptionModelRequest
*/
func (a *SubscriptionModelsApiService) GETMarketIdSubscriptionModel(ctx context.Context, marketId interface{}) SubscriptionModelsApiGETMarketIdSubscriptionModelRequest {
	return SubscriptionModelsApiGETMarketIdSubscriptionModelRequest{
		ApiService: a,
		ctx:        ctx,
		marketId:   marketId,
	}
}

// Execute executes the request
func (a *SubscriptionModelsApiService) GETMarketIdSubscriptionModelExecute(r SubscriptionModelsApiGETMarketIdSubscriptionModelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionModelsApiService.GETMarketIdSubscriptionModel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/markets/{marketId}/subscription_model"
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

type SubscriptionModelsApiGETOrderSubscriptionIdSubscriptionModelRequest struct {
	ctx                 context.Context
	ApiService          *SubscriptionModelsApiService
	orderSubscriptionId interface{}
}

func (r SubscriptionModelsApiGETOrderSubscriptionIdSubscriptionModelRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderSubscriptionIdSubscriptionModelExecute(r)
}

/*
GETOrderSubscriptionIdSubscriptionModel Retrieve the subscription model associated to the order subscription

Retrieve the subscription model associated to the order subscription

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderSubscriptionId The resource's id
	@return SubscriptionModelsApiGETOrderSubscriptionIdSubscriptionModelRequest
*/
func (a *SubscriptionModelsApiService) GETOrderSubscriptionIdSubscriptionModel(ctx context.Context, orderSubscriptionId interface{}) SubscriptionModelsApiGETOrderSubscriptionIdSubscriptionModelRequest {
	return SubscriptionModelsApiGETOrderSubscriptionIdSubscriptionModelRequest{
		ApiService:          a,
		ctx:                 ctx,
		orderSubscriptionId: orderSubscriptionId,
	}
}

// Execute executes the request
func (a *SubscriptionModelsApiService) GETOrderSubscriptionIdSubscriptionModelExecute(r SubscriptionModelsApiGETOrderSubscriptionIdSubscriptionModelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionModelsApiService.GETOrderSubscriptionIdSubscriptionModel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order_subscriptions/{orderSubscriptionId}/subscription_model"
	localVarPath = strings.Replace(localVarPath, "{"+"orderSubscriptionId"+"}", url.PathEscape(parameterValueToString(r.orderSubscriptionId, "orderSubscriptionId")), -1)

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

type SubscriptionModelsApiGETSubscriptionModelsRequest struct {
	ctx        context.Context
	ApiService *SubscriptionModelsApiService
}

func (r SubscriptionModelsApiGETSubscriptionModelsRequest) Execute() (*GETSubscriptionModels200Response, *http.Response, error) {
	return r.ApiService.GETSubscriptionModelsExecute(r)
}

/*
GETSubscriptionModels List all subscription models

List all subscription models

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SubscriptionModelsApiGETSubscriptionModelsRequest
*/
func (a *SubscriptionModelsApiService) GETSubscriptionModels(ctx context.Context) SubscriptionModelsApiGETSubscriptionModelsRequest {
	return SubscriptionModelsApiGETSubscriptionModelsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETSubscriptionModels200Response
func (a *SubscriptionModelsApiService) GETSubscriptionModelsExecute(r SubscriptionModelsApiGETSubscriptionModelsRequest) (*GETSubscriptionModels200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETSubscriptionModels200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionModelsApiService.GETSubscriptionModels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription_models"

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

type SubscriptionModelsApiGETSubscriptionModelsSubscriptionModelIdRequest struct {
	ctx                 context.Context
	ApiService          *SubscriptionModelsApiService
	subscriptionModelId interface{}
}

func (r SubscriptionModelsApiGETSubscriptionModelsSubscriptionModelIdRequest) Execute() (*GETSubscriptionModelsSubscriptionModelId200Response, *http.Response, error) {
	return r.ApiService.GETSubscriptionModelsSubscriptionModelIdExecute(r)
}

/*
GETSubscriptionModelsSubscriptionModelId Retrieve a subscription model

Retrieve a subscription model

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param subscriptionModelId The resource's id
	@return SubscriptionModelsApiGETSubscriptionModelsSubscriptionModelIdRequest
*/
func (a *SubscriptionModelsApiService) GETSubscriptionModelsSubscriptionModelId(ctx context.Context, subscriptionModelId interface{}) SubscriptionModelsApiGETSubscriptionModelsSubscriptionModelIdRequest {
	return SubscriptionModelsApiGETSubscriptionModelsSubscriptionModelIdRequest{
		ApiService:          a,
		ctx:                 ctx,
		subscriptionModelId: subscriptionModelId,
	}
}

// Execute executes the request
//
//	@return GETSubscriptionModelsSubscriptionModelId200Response
func (a *SubscriptionModelsApiService) GETSubscriptionModelsSubscriptionModelIdExecute(r SubscriptionModelsApiGETSubscriptionModelsSubscriptionModelIdRequest) (*GETSubscriptionModelsSubscriptionModelId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETSubscriptionModelsSubscriptionModelId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionModelsApiService.GETSubscriptionModelsSubscriptionModelId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription_models/{subscriptionModelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionModelId"+"}", url.PathEscape(parameterValueToString(r.subscriptionModelId, "subscriptionModelId")), -1)

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

type SubscriptionModelsApiPATCHSubscriptionModelsSubscriptionModelIdRequest struct {
	ctx                                               context.Context
	ApiService                                        *SubscriptionModelsApiService
	pATCHSubscriptionModelsSubscriptionModelIdRequest *PATCHSubscriptionModelsSubscriptionModelIdRequest
	subscriptionModelId                               interface{}
}

func (r SubscriptionModelsApiPATCHSubscriptionModelsSubscriptionModelIdRequest) PATCHSubscriptionModelsSubscriptionModelIdRequest(pATCHSubscriptionModelsSubscriptionModelIdRequest PATCHSubscriptionModelsSubscriptionModelIdRequest) SubscriptionModelsApiPATCHSubscriptionModelsSubscriptionModelIdRequest {
	r.pATCHSubscriptionModelsSubscriptionModelIdRequest = &pATCHSubscriptionModelsSubscriptionModelIdRequest
	return r
}

func (r SubscriptionModelsApiPATCHSubscriptionModelsSubscriptionModelIdRequest) Execute() (*PATCHSubscriptionModelsSubscriptionModelId200Response, *http.Response, error) {
	return r.ApiService.PATCHSubscriptionModelsSubscriptionModelIdExecute(r)
}

/*
PATCHSubscriptionModelsSubscriptionModelId Update a subscription model

Update a subscription model

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param subscriptionModelId The resource's id
	@return SubscriptionModelsApiPATCHSubscriptionModelsSubscriptionModelIdRequest
*/
func (a *SubscriptionModelsApiService) PATCHSubscriptionModelsSubscriptionModelId(ctx context.Context, subscriptionModelId interface{}) SubscriptionModelsApiPATCHSubscriptionModelsSubscriptionModelIdRequest {
	return SubscriptionModelsApiPATCHSubscriptionModelsSubscriptionModelIdRequest{
		ApiService:          a,
		ctx:                 ctx,
		subscriptionModelId: subscriptionModelId,
	}
}

// Execute executes the request
//
//	@return PATCHSubscriptionModelsSubscriptionModelId200Response
func (a *SubscriptionModelsApiService) PATCHSubscriptionModelsSubscriptionModelIdExecute(r SubscriptionModelsApiPATCHSubscriptionModelsSubscriptionModelIdRequest) (*PATCHSubscriptionModelsSubscriptionModelId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHSubscriptionModelsSubscriptionModelId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionModelsApiService.PATCHSubscriptionModelsSubscriptionModelId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription_models/{subscriptionModelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionModelId"+"}", url.PathEscape(parameterValueToString(r.subscriptionModelId, "subscriptionModelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pATCHSubscriptionModelsSubscriptionModelIdRequest == nil {
		return localVarReturnValue, nil, reportError("pATCHSubscriptionModelsSubscriptionModelIdRequest is required and must be specified")
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
	localVarPostBody = r.pATCHSubscriptionModelsSubscriptionModelIdRequest
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

type SubscriptionModelsApiPOSTSubscriptionModelsRequest struct {
	ctx                           context.Context
	ApiService                    *SubscriptionModelsApiService
	pOSTSubscriptionModelsRequest *POSTSubscriptionModelsRequest
}

func (r SubscriptionModelsApiPOSTSubscriptionModelsRequest) POSTSubscriptionModelsRequest(pOSTSubscriptionModelsRequest POSTSubscriptionModelsRequest) SubscriptionModelsApiPOSTSubscriptionModelsRequest {
	r.pOSTSubscriptionModelsRequest = &pOSTSubscriptionModelsRequest
	return r
}

func (r SubscriptionModelsApiPOSTSubscriptionModelsRequest) Execute() (*POSTSubscriptionModels201Response, *http.Response, error) {
	return r.ApiService.POSTSubscriptionModelsExecute(r)
}

/*
POSTSubscriptionModels Create a subscription model

Create a subscription model

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SubscriptionModelsApiPOSTSubscriptionModelsRequest
*/
func (a *SubscriptionModelsApiService) POSTSubscriptionModels(ctx context.Context) SubscriptionModelsApiPOSTSubscriptionModelsRequest {
	return SubscriptionModelsApiPOSTSubscriptionModelsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTSubscriptionModels201Response
func (a *SubscriptionModelsApiService) POSTSubscriptionModelsExecute(r SubscriptionModelsApiPOSTSubscriptionModelsRequest) (*POSTSubscriptionModels201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTSubscriptionModels201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionModelsApiService.POSTSubscriptionModels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription_models"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pOSTSubscriptionModelsRequest == nil {
		return localVarReturnValue, nil, reportError("pOSTSubscriptionModelsRequest is required and must be specified")
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
	localVarPostBody = r.pOSTSubscriptionModelsRequest
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
