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

// TaxjarAccountsApiService TaxjarAccountsApi service
type TaxjarAccountsApiService service

type TaxjarAccountsApiDELETETaxjarAccountsTaxjarAccountIdRequest struct {
	ctx             context.Context
	ApiService      *TaxjarAccountsApiService
	taxjarAccountId interface{}
}

func (r TaxjarAccountsApiDELETETaxjarAccountsTaxjarAccountIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETETaxjarAccountsTaxjarAccountIdExecute(r)
}

/*
DELETETaxjarAccountsTaxjarAccountId Delete a taxjar account

Delete a taxjar account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param taxjarAccountId The resource's id
	@return TaxjarAccountsApiDELETETaxjarAccountsTaxjarAccountIdRequest
*/
func (a *TaxjarAccountsApiService) DELETETaxjarAccountsTaxjarAccountId(ctx context.Context, taxjarAccountId interface{}) TaxjarAccountsApiDELETETaxjarAccountsTaxjarAccountIdRequest {
	return TaxjarAccountsApiDELETETaxjarAccountsTaxjarAccountIdRequest{
		ApiService:      a,
		ctx:             ctx,
		taxjarAccountId: taxjarAccountId,
	}
}

// Execute executes the request
func (a *TaxjarAccountsApiService) DELETETaxjarAccountsTaxjarAccountIdExecute(r TaxjarAccountsApiDELETETaxjarAccountsTaxjarAccountIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxjarAccountsApiService.DELETETaxjarAccountsTaxjarAccountId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/taxjar_accounts/{taxjarAccountId}"
	localVarPath = strings.Replace(localVarPath, "{"+"taxjarAccountId"+"}", url.PathEscape(parameterValueToString(r.taxjarAccountId, "taxjarAccountId")), -1)

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

type TaxjarAccountsApiGETTaxjarAccountsRequest struct {
	ctx        context.Context
	ApiService *TaxjarAccountsApiService
}

func (r TaxjarAccountsApiGETTaxjarAccountsRequest) Execute() (*GETTaxjarAccounts200Response, *http.Response, error) {
	return r.ApiService.GETTaxjarAccountsExecute(r)
}

/*
GETTaxjarAccounts List all taxjar accounts

List all taxjar accounts

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TaxjarAccountsApiGETTaxjarAccountsRequest
*/
func (a *TaxjarAccountsApiService) GETTaxjarAccounts(ctx context.Context) TaxjarAccountsApiGETTaxjarAccountsRequest {
	return TaxjarAccountsApiGETTaxjarAccountsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETTaxjarAccounts200Response
func (a *TaxjarAccountsApiService) GETTaxjarAccountsExecute(r TaxjarAccountsApiGETTaxjarAccountsRequest) (*GETTaxjarAccounts200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETTaxjarAccounts200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxjarAccountsApiService.GETTaxjarAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/taxjar_accounts"

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

type TaxjarAccountsApiGETTaxjarAccountsTaxjarAccountIdRequest struct {
	ctx             context.Context
	ApiService      *TaxjarAccountsApiService
	taxjarAccountId interface{}
}

func (r TaxjarAccountsApiGETTaxjarAccountsTaxjarAccountIdRequest) Execute() (*GETTaxjarAccountsTaxjarAccountId200Response, *http.Response, error) {
	return r.ApiService.GETTaxjarAccountsTaxjarAccountIdExecute(r)
}

/*
GETTaxjarAccountsTaxjarAccountId Retrieve a taxjar account

Retrieve a taxjar account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param taxjarAccountId The resource's id
	@return TaxjarAccountsApiGETTaxjarAccountsTaxjarAccountIdRequest
*/
func (a *TaxjarAccountsApiService) GETTaxjarAccountsTaxjarAccountId(ctx context.Context, taxjarAccountId interface{}) TaxjarAccountsApiGETTaxjarAccountsTaxjarAccountIdRequest {
	return TaxjarAccountsApiGETTaxjarAccountsTaxjarAccountIdRequest{
		ApiService:      a,
		ctx:             ctx,
		taxjarAccountId: taxjarAccountId,
	}
}

// Execute executes the request
//
//	@return GETTaxjarAccountsTaxjarAccountId200Response
func (a *TaxjarAccountsApiService) GETTaxjarAccountsTaxjarAccountIdExecute(r TaxjarAccountsApiGETTaxjarAccountsTaxjarAccountIdRequest) (*GETTaxjarAccountsTaxjarAccountId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETTaxjarAccountsTaxjarAccountId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxjarAccountsApiService.GETTaxjarAccountsTaxjarAccountId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/taxjar_accounts/{taxjarAccountId}"
	localVarPath = strings.Replace(localVarPath, "{"+"taxjarAccountId"+"}", url.PathEscape(parameterValueToString(r.taxjarAccountId, "taxjarAccountId")), -1)

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

type TaxjarAccountsApiPATCHTaxjarAccountsTaxjarAccountIdRequest struct {
	ctx                 context.Context
	ApiService          *TaxjarAccountsApiService
	taxjarAccountUpdate *TaxjarAccountUpdate
	taxjarAccountId     interface{}
}

func (r TaxjarAccountsApiPATCHTaxjarAccountsTaxjarAccountIdRequest) TaxjarAccountUpdate(taxjarAccountUpdate TaxjarAccountUpdate) TaxjarAccountsApiPATCHTaxjarAccountsTaxjarAccountIdRequest {
	r.taxjarAccountUpdate = &taxjarAccountUpdate
	return r
}

func (r TaxjarAccountsApiPATCHTaxjarAccountsTaxjarAccountIdRequest) Execute() (*PATCHTaxjarAccountsTaxjarAccountId200Response, *http.Response, error) {
	return r.ApiService.PATCHTaxjarAccountsTaxjarAccountIdExecute(r)
}

/*
PATCHTaxjarAccountsTaxjarAccountId Update a taxjar account

Update a taxjar account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param taxjarAccountId The resource's id
	@return TaxjarAccountsApiPATCHTaxjarAccountsTaxjarAccountIdRequest
*/
func (a *TaxjarAccountsApiService) PATCHTaxjarAccountsTaxjarAccountId(ctx context.Context, taxjarAccountId interface{}) TaxjarAccountsApiPATCHTaxjarAccountsTaxjarAccountIdRequest {
	return TaxjarAccountsApiPATCHTaxjarAccountsTaxjarAccountIdRequest{
		ApiService:      a,
		ctx:             ctx,
		taxjarAccountId: taxjarAccountId,
	}
}

// Execute executes the request
//
//	@return PATCHTaxjarAccountsTaxjarAccountId200Response
func (a *TaxjarAccountsApiService) PATCHTaxjarAccountsTaxjarAccountIdExecute(r TaxjarAccountsApiPATCHTaxjarAccountsTaxjarAccountIdRequest) (*PATCHTaxjarAccountsTaxjarAccountId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHTaxjarAccountsTaxjarAccountId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxjarAccountsApiService.PATCHTaxjarAccountsTaxjarAccountId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/taxjar_accounts/{taxjarAccountId}"
	localVarPath = strings.Replace(localVarPath, "{"+"taxjarAccountId"+"}", url.PathEscape(parameterValueToString(r.taxjarAccountId, "taxjarAccountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.taxjarAccountUpdate == nil {
		return localVarReturnValue, nil, reportError("taxjarAccountUpdate is required and must be specified")
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
	localVarPostBody = r.taxjarAccountUpdate
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

type TaxjarAccountsApiPOSTTaxjarAccountsRequest struct {
	ctx                 context.Context
	ApiService          *TaxjarAccountsApiService
	taxjarAccountCreate *TaxjarAccountCreate
}

func (r TaxjarAccountsApiPOSTTaxjarAccountsRequest) TaxjarAccountCreate(taxjarAccountCreate TaxjarAccountCreate) TaxjarAccountsApiPOSTTaxjarAccountsRequest {
	r.taxjarAccountCreate = &taxjarAccountCreate
	return r
}

func (r TaxjarAccountsApiPOSTTaxjarAccountsRequest) Execute() (*POSTTaxjarAccounts201Response, *http.Response, error) {
	return r.ApiService.POSTTaxjarAccountsExecute(r)
}

/*
POSTTaxjarAccounts Create a taxjar account

Create a taxjar account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TaxjarAccountsApiPOSTTaxjarAccountsRequest
*/
func (a *TaxjarAccountsApiService) POSTTaxjarAccounts(ctx context.Context) TaxjarAccountsApiPOSTTaxjarAccountsRequest {
	return TaxjarAccountsApiPOSTTaxjarAccountsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTTaxjarAccounts201Response
func (a *TaxjarAccountsApiService) POSTTaxjarAccountsExecute(r TaxjarAccountsApiPOSTTaxjarAccountsRequest) (*POSTTaxjarAccounts201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTTaxjarAccounts201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxjarAccountsApiService.POSTTaxjarAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/taxjar_accounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.taxjarAccountCreate == nil {
		return localVarReturnValue, nil, reportError("taxjarAccountCreate is required and must be specified")
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
	localVarPostBody = r.taxjarAccountCreate
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
