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


// ManualTaxCalculatorsApiService ManualTaxCalculatorsApi service
type ManualTaxCalculatorsApiService service

type ManualTaxCalculatorsApiDELETEManualTaxCalculatorsManualTaxCalculatorIdRequest struct {
	ctx context.Context
	ApiService *ManualTaxCalculatorsApiService
	manualTaxCalculatorId string
}

func (r ManualTaxCalculatorsApiDELETEManualTaxCalculatorsManualTaxCalculatorIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEManualTaxCalculatorsManualTaxCalculatorIdExecute(r)
}

/*
DELETEManualTaxCalculatorsManualTaxCalculatorId Delete a manual tax calculator

Delete a manual tax calculator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param manualTaxCalculatorId The resource's id
 @return ManualTaxCalculatorsApiDELETEManualTaxCalculatorsManualTaxCalculatorIdRequest
*/
func (a *ManualTaxCalculatorsApiService) DELETEManualTaxCalculatorsManualTaxCalculatorId(ctx context.Context, manualTaxCalculatorId string) ManualTaxCalculatorsApiDELETEManualTaxCalculatorsManualTaxCalculatorIdRequest {
	return ManualTaxCalculatorsApiDELETEManualTaxCalculatorsManualTaxCalculatorIdRequest{
		ApiService: a,
		ctx: ctx,
		manualTaxCalculatorId: manualTaxCalculatorId,
	}
}

// Execute executes the request
func (a *ManualTaxCalculatorsApiService) DELETEManualTaxCalculatorsManualTaxCalculatorIdExecute(r ManualTaxCalculatorsApiDELETEManualTaxCalculatorsManualTaxCalculatorIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManualTaxCalculatorsApiService.DELETEManualTaxCalculatorsManualTaxCalculatorId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/manual_tax_calculators/{manualTaxCalculatorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"manualTaxCalculatorId"+"}", url.PathEscape(parameterToString(r.manualTaxCalculatorId, "")), -1)

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

type ManualTaxCalculatorsApiGETManualTaxCalculatorsRequest struct {
	ctx context.Context
	ApiService *ManualTaxCalculatorsApiService
}

func (r ManualTaxCalculatorsApiGETManualTaxCalculatorsRequest) Execute() (*GETManualTaxCalculators200Response, *http.Response, error) {
	return r.ApiService.GETManualTaxCalculatorsExecute(r)
}

/*
GETManualTaxCalculators List all manual tax calculators

List all manual tax calculators

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ManualTaxCalculatorsApiGETManualTaxCalculatorsRequest
*/
func (a *ManualTaxCalculatorsApiService) GETManualTaxCalculators(ctx context.Context) ManualTaxCalculatorsApiGETManualTaxCalculatorsRequest {
	return ManualTaxCalculatorsApiGETManualTaxCalculatorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GETManualTaxCalculators200Response
func (a *ManualTaxCalculatorsApiService) GETManualTaxCalculatorsExecute(r ManualTaxCalculatorsApiGETManualTaxCalculatorsRequest) (*GETManualTaxCalculators200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GETManualTaxCalculators200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManualTaxCalculatorsApiService.GETManualTaxCalculators")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/manual_tax_calculators"

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

type ManualTaxCalculatorsApiGETManualTaxCalculatorsManualTaxCalculatorIdRequest struct {
	ctx context.Context
	ApiService *ManualTaxCalculatorsApiService
	manualTaxCalculatorId string
}

func (r ManualTaxCalculatorsApiGETManualTaxCalculatorsManualTaxCalculatorIdRequest) Execute() (*GETManualTaxCalculatorsManualTaxCalculatorId200Response, *http.Response, error) {
	return r.ApiService.GETManualTaxCalculatorsManualTaxCalculatorIdExecute(r)
}

/*
GETManualTaxCalculatorsManualTaxCalculatorId Retrieve a manual tax calculator

Retrieve a manual tax calculator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param manualTaxCalculatorId The resource's id
 @return ManualTaxCalculatorsApiGETManualTaxCalculatorsManualTaxCalculatorIdRequest
*/
func (a *ManualTaxCalculatorsApiService) GETManualTaxCalculatorsManualTaxCalculatorId(ctx context.Context, manualTaxCalculatorId string) ManualTaxCalculatorsApiGETManualTaxCalculatorsManualTaxCalculatorIdRequest {
	return ManualTaxCalculatorsApiGETManualTaxCalculatorsManualTaxCalculatorIdRequest{
		ApiService: a,
		ctx: ctx,
		manualTaxCalculatorId: manualTaxCalculatorId,
	}
}

// Execute executes the request
//  @return GETManualTaxCalculatorsManualTaxCalculatorId200Response
func (a *ManualTaxCalculatorsApiService) GETManualTaxCalculatorsManualTaxCalculatorIdExecute(r ManualTaxCalculatorsApiGETManualTaxCalculatorsManualTaxCalculatorIdRequest) (*GETManualTaxCalculatorsManualTaxCalculatorId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GETManualTaxCalculatorsManualTaxCalculatorId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManualTaxCalculatorsApiService.GETManualTaxCalculatorsManualTaxCalculatorId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/manual_tax_calculators/{manualTaxCalculatorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"manualTaxCalculatorId"+"}", url.PathEscape(parameterToString(r.manualTaxCalculatorId, "")), -1)

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

type ManualTaxCalculatorsApiGETTaxRuleIdManualTaxCalculatorRequest struct {
	ctx context.Context
	ApiService *ManualTaxCalculatorsApiService
	taxRuleId string
}

func (r ManualTaxCalculatorsApiGETTaxRuleIdManualTaxCalculatorRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETTaxRuleIdManualTaxCalculatorExecute(r)
}

/*
GETTaxRuleIdManualTaxCalculator Retrieve the manual tax calculator associated to the tax rule

Retrieve the manual tax calculator associated to the tax rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taxRuleId The resource's id
 @return ManualTaxCalculatorsApiGETTaxRuleIdManualTaxCalculatorRequest
*/
func (a *ManualTaxCalculatorsApiService) GETTaxRuleIdManualTaxCalculator(ctx context.Context, taxRuleId string) ManualTaxCalculatorsApiGETTaxRuleIdManualTaxCalculatorRequest {
	return ManualTaxCalculatorsApiGETTaxRuleIdManualTaxCalculatorRequest{
		ApiService: a,
		ctx: ctx,
		taxRuleId: taxRuleId,
	}
}

// Execute executes the request
func (a *ManualTaxCalculatorsApiService) GETTaxRuleIdManualTaxCalculatorExecute(r ManualTaxCalculatorsApiGETTaxRuleIdManualTaxCalculatorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManualTaxCalculatorsApiService.GETTaxRuleIdManualTaxCalculator")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tax_rules/{taxRuleId}/manual_tax_calculator"
	localVarPath = strings.Replace(localVarPath, "{"+"taxRuleId"+"}", url.PathEscape(parameterToString(r.taxRuleId, "")), -1)

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

type ManualTaxCalculatorsApiPATCHManualTaxCalculatorsManualTaxCalculatorIdRequest struct {
	ctx context.Context
	ApiService *ManualTaxCalculatorsApiService
	manualTaxCalculatorUpdate *ManualTaxCalculatorUpdate
	manualTaxCalculatorId string
}

func (r ManualTaxCalculatorsApiPATCHManualTaxCalculatorsManualTaxCalculatorIdRequest) ManualTaxCalculatorUpdate(manualTaxCalculatorUpdate ManualTaxCalculatorUpdate) ManualTaxCalculatorsApiPATCHManualTaxCalculatorsManualTaxCalculatorIdRequest {
	r.manualTaxCalculatorUpdate = &manualTaxCalculatorUpdate
	return r
}

func (r ManualTaxCalculatorsApiPATCHManualTaxCalculatorsManualTaxCalculatorIdRequest) Execute() (*PATCHManualTaxCalculatorsManualTaxCalculatorId200Response, *http.Response, error) {
	return r.ApiService.PATCHManualTaxCalculatorsManualTaxCalculatorIdExecute(r)
}

/*
PATCHManualTaxCalculatorsManualTaxCalculatorId Update a manual tax calculator

Update a manual tax calculator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param manualTaxCalculatorId The resource's id
 @return ManualTaxCalculatorsApiPATCHManualTaxCalculatorsManualTaxCalculatorIdRequest
*/
func (a *ManualTaxCalculatorsApiService) PATCHManualTaxCalculatorsManualTaxCalculatorId(ctx context.Context, manualTaxCalculatorId string) ManualTaxCalculatorsApiPATCHManualTaxCalculatorsManualTaxCalculatorIdRequest {
	return ManualTaxCalculatorsApiPATCHManualTaxCalculatorsManualTaxCalculatorIdRequest{
		ApiService: a,
		ctx: ctx,
		manualTaxCalculatorId: manualTaxCalculatorId,
	}
}

// Execute executes the request
//  @return PATCHManualTaxCalculatorsManualTaxCalculatorId200Response
func (a *ManualTaxCalculatorsApiService) PATCHManualTaxCalculatorsManualTaxCalculatorIdExecute(r ManualTaxCalculatorsApiPATCHManualTaxCalculatorsManualTaxCalculatorIdRequest) (*PATCHManualTaxCalculatorsManualTaxCalculatorId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PATCHManualTaxCalculatorsManualTaxCalculatorId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManualTaxCalculatorsApiService.PATCHManualTaxCalculatorsManualTaxCalculatorId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/manual_tax_calculators/{manualTaxCalculatorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"manualTaxCalculatorId"+"}", url.PathEscape(parameterToString(r.manualTaxCalculatorId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.manualTaxCalculatorUpdate == nil {
		return localVarReturnValue, nil, reportError("manualTaxCalculatorUpdate is required and must be specified")
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
	localVarPostBody = r.manualTaxCalculatorUpdate
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

type ManualTaxCalculatorsApiPOSTManualTaxCalculatorsRequest struct {
	ctx context.Context
	ApiService *ManualTaxCalculatorsApiService
	manualTaxCalculatorCreate *ManualTaxCalculatorCreate
}

func (r ManualTaxCalculatorsApiPOSTManualTaxCalculatorsRequest) ManualTaxCalculatorCreate(manualTaxCalculatorCreate ManualTaxCalculatorCreate) ManualTaxCalculatorsApiPOSTManualTaxCalculatorsRequest {
	r.manualTaxCalculatorCreate = &manualTaxCalculatorCreate
	return r
}

func (r ManualTaxCalculatorsApiPOSTManualTaxCalculatorsRequest) Execute() (*POSTManualTaxCalculators201Response, *http.Response, error) {
	return r.ApiService.POSTManualTaxCalculatorsExecute(r)
}

/*
POSTManualTaxCalculators Create a manual tax calculator

Create a manual tax calculator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ManualTaxCalculatorsApiPOSTManualTaxCalculatorsRequest
*/
func (a *ManualTaxCalculatorsApiService) POSTManualTaxCalculators(ctx context.Context) ManualTaxCalculatorsApiPOSTManualTaxCalculatorsRequest {
	return ManualTaxCalculatorsApiPOSTManualTaxCalculatorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return POSTManualTaxCalculators201Response
func (a *ManualTaxCalculatorsApiService) POSTManualTaxCalculatorsExecute(r ManualTaxCalculatorsApiPOSTManualTaxCalculatorsRequest) (*POSTManualTaxCalculators201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *POSTManualTaxCalculators201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManualTaxCalculatorsApiService.POSTManualTaxCalculators")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/manual_tax_calculators"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.manualTaxCalculatorCreate == nil {
		return localVarReturnValue, nil, reportError("manualTaxCalculatorCreate is required and must be specified")
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
	localVarPostBody = r.manualTaxCalculatorCreate
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
