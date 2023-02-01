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

// CustomerPaymentSourcesApiService CustomerPaymentSourcesApi service
type CustomerPaymentSourcesApiService service

type CustomerPaymentSourcesApiDELETECustomerPaymentSourcesCustomerPaymentSourceIdRequest struct {
	ctx                     context.Context
	ApiService              *CustomerPaymentSourcesApiService
	customerPaymentSourceId string
}

func (r CustomerPaymentSourcesApiDELETECustomerPaymentSourcesCustomerPaymentSourceIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETECustomerPaymentSourcesCustomerPaymentSourceIdExecute(r)
}

/*
DELETECustomerPaymentSourcesCustomerPaymentSourceId Delete a customer payment source

Delete a customer payment source

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerPaymentSourceId The resource's id
	@return CustomerPaymentSourcesApiDELETECustomerPaymentSourcesCustomerPaymentSourceIdRequest
*/
func (a *CustomerPaymentSourcesApiService) DELETECustomerPaymentSourcesCustomerPaymentSourceId(ctx context.Context, customerPaymentSourceId string) CustomerPaymentSourcesApiDELETECustomerPaymentSourcesCustomerPaymentSourceIdRequest {
	return CustomerPaymentSourcesApiDELETECustomerPaymentSourcesCustomerPaymentSourceIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		customerPaymentSourceId: customerPaymentSourceId,
	}
}

// Execute executes the request
func (a *CustomerPaymentSourcesApiService) DELETECustomerPaymentSourcesCustomerPaymentSourceIdExecute(r CustomerPaymentSourcesApiDELETECustomerPaymentSourcesCustomerPaymentSourceIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPaymentSourcesApiService.DELETECustomerPaymentSourcesCustomerPaymentSourceId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_payment_sources/{customerPaymentSourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerPaymentSourceId"+"}", url.PathEscape(parameterToString(r.customerPaymentSourceId, "")), -1)

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

type CustomerPaymentSourcesApiGETCustomerIdCustomerPaymentSourcesRequest struct {
	ctx        context.Context
	ApiService *CustomerPaymentSourcesApiService
	customerId string
}

func (r CustomerPaymentSourcesApiGETCustomerIdCustomerPaymentSourcesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCustomerIdCustomerPaymentSourcesExecute(r)
}

/*
GETCustomerIdCustomerPaymentSources Retrieve the customer payment sources associated to the customer

Retrieve the customer payment sources associated to the customer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerId The resource's id
	@return CustomerPaymentSourcesApiGETCustomerIdCustomerPaymentSourcesRequest
*/
func (a *CustomerPaymentSourcesApiService) GETCustomerIdCustomerPaymentSources(ctx context.Context, customerId string) CustomerPaymentSourcesApiGETCustomerIdCustomerPaymentSourcesRequest {
	return CustomerPaymentSourcesApiGETCustomerIdCustomerPaymentSourcesRequest{
		ApiService: a,
		ctx:        ctx,
		customerId: customerId,
	}
}

// Execute executes the request
func (a *CustomerPaymentSourcesApiService) GETCustomerIdCustomerPaymentSourcesExecute(r CustomerPaymentSourcesApiGETCustomerIdCustomerPaymentSourcesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPaymentSourcesApiService.GETCustomerIdCustomerPaymentSources")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customerId}/customer_payment_sources"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", url.PathEscape(parameterToString(r.customerId, "")), -1)

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

type CustomerPaymentSourcesApiGETCustomerPaymentSourcesRequest struct {
	ctx        context.Context
	ApiService *CustomerPaymentSourcesApiService
}

func (r CustomerPaymentSourcesApiGETCustomerPaymentSourcesRequest) Execute() (*GETCustomerPaymentSources200Response, *http.Response, error) {
	return r.ApiService.GETCustomerPaymentSourcesExecute(r)
}

/*
GETCustomerPaymentSources List all customer payment sources

List all customer payment sources

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CustomerPaymentSourcesApiGETCustomerPaymentSourcesRequest
*/
func (a *CustomerPaymentSourcesApiService) GETCustomerPaymentSources(ctx context.Context) CustomerPaymentSourcesApiGETCustomerPaymentSourcesRequest {
	return CustomerPaymentSourcesApiGETCustomerPaymentSourcesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETCustomerPaymentSources200Response
func (a *CustomerPaymentSourcesApiService) GETCustomerPaymentSourcesExecute(r CustomerPaymentSourcesApiGETCustomerPaymentSourcesRequest) (*GETCustomerPaymentSources200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCustomerPaymentSources200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPaymentSourcesApiService.GETCustomerPaymentSources")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_payment_sources"

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

type CustomerPaymentSourcesApiGETCustomerPaymentSourcesCustomerPaymentSourceIdRequest struct {
	ctx                     context.Context
	ApiService              *CustomerPaymentSourcesApiService
	customerPaymentSourceId string
}

func (r CustomerPaymentSourcesApiGETCustomerPaymentSourcesCustomerPaymentSourceIdRequest) Execute() (*GETCustomerPaymentSourcesCustomerPaymentSourceId200Response, *http.Response, error) {
	return r.ApiService.GETCustomerPaymentSourcesCustomerPaymentSourceIdExecute(r)
}

/*
GETCustomerPaymentSourcesCustomerPaymentSourceId Retrieve a customer payment source

Retrieve a customer payment source

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerPaymentSourceId The resource's id
	@return CustomerPaymentSourcesApiGETCustomerPaymentSourcesCustomerPaymentSourceIdRequest
*/
func (a *CustomerPaymentSourcesApiService) GETCustomerPaymentSourcesCustomerPaymentSourceId(ctx context.Context, customerPaymentSourceId string) CustomerPaymentSourcesApiGETCustomerPaymentSourcesCustomerPaymentSourceIdRequest {
	return CustomerPaymentSourcesApiGETCustomerPaymentSourcesCustomerPaymentSourceIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		customerPaymentSourceId: customerPaymentSourceId,
	}
}

// Execute executes the request
//
//	@return GETCustomerPaymentSourcesCustomerPaymentSourceId200Response
func (a *CustomerPaymentSourcesApiService) GETCustomerPaymentSourcesCustomerPaymentSourceIdExecute(r CustomerPaymentSourcesApiGETCustomerPaymentSourcesCustomerPaymentSourceIdRequest) (*GETCustomerPaymentSourcesCustomerPaymentSourceId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCustomerPaymentSourcesCustomerPaymentSourceId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPaymentSourcesApiService.GETCustomerPaymentSourcesCustomerPaymentSourceId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_payment_sources/{customerPaymentSourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerPaymentSourceId"+"}", url.PathEscape(parameterToString(r.customerPaymentSourceId, "")), -1)

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

type CustomerPaymentSourcesApiGETExternalPaymentIdWalletRequest struct {
	ctx               context.Context
	ApiService        *CustomerPaymentSourcesApiService
	externalPaymentId string
}

func (r CustomerPaymentSourcesApiGETExternalPaymentIdWalletRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETExternalPaymentIdWalletExecute(r)
}

/*
GETExternalPaymentIdWallet Retrieve the wallet associated to the external payment

Retrieve the wallet associated to the external payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param externalPaymentId The resource's id
	@return CustomerPaymentSourcesApiGETExternalPaymentIdWalletRequest
*/
func (a *CustomerPaymentSourcesApiService) GETExternalPaymentIdWallet(ctx context.Context, externalPaymentId string) CustomerPaymentSourcesApiGETExternalPaymentIdWalletRequest {
	return CustomerPaymentSourcesApiGETExternalPaymentIdWalletRequest{
		ApiService:        a,
		ctx:               ctx,
		externalPaymentId: externalPaymentId,
	}
}

// Execute executes the request
func (a *CustomerPaymentSourcesApiService) GETExternalPaymentIdWalletExecute(r CustomerPaymentSourcesApiGETExternalPaymentIdWalletRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPaymentSourcesApiService.GETExternalPaymentIdWallet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_payments/{externalPaymentId}/wallet"
	localVarPath = strings.Replace(localVarPath, "{"+"externalPaymentId"+"}", url.PathEscape(parameterToString(r.externalPaymentId, "")), -1)

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

type CustomerPaymentSourcesApiGETOrderIdAvailableCustomerPaymentSourcesRequest struct {
	ctx        context.Context
	ApiService *CustomerPaymentSourcesApiService
	orderId    string
}

func (r CustomerPaymentSourcesApiGETOrderIdAvailableCustomerPaymentSourcesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderIdAvailableCustomerPaymentSourcesExecute(r)
}

/*
GETOrderIdAvailableCustomerPaymentSources Retrieve the available customer payment sources associated to the order

Retrieve the available customer payment sources associated to the order

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId The resource's id
	@return CustomerPaymentSourcesApiGETOrderIdAvailableCustomerPaymentSourcesRequest
*/
func (a *CustomerPaymentSourcesApiService) GETOrderIdAvailableCustomerPaymentSources(ctx context.Context, orderId string) CustomerPaymentSourcesApiGETOrderIdAvailableCustomerPaymentSourcesRequest {
	return CustomerPaymentSourcesApiGETOrderIdAvailableCustomerPaymentSourcesRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *CustomerPaymentSourcesApiService) GETOrderIdAvailableCustomerPaymentSourcesExecute(r CustomerPaymentSourcesApiGETOrderIdAvailableCustomerPaymentSourcesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPaymentSourcesApiService.GETOrderIdAvailableCustomerPaymentSources")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{orderId}/available_customer_payment_sources"
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

type CustomerPaymentSourcesApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest struct {
	ctx                         context.Context
	ApiService                  *CustomerPaymentSourcesApiService
	customerPaymentSourceUpdate *CustomerPaymentSourceUpdate
	customerPaymentSourceId     string
}

func (r CustomerPaymentSourcesApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) CustomerPaymentSourceUpdate(customerPaymentSourceUpdate CustomerPaymentSourceUpdate) CustomerPaymentSourcesApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest {
	r.customerPaymentSourceUpdate = &customerPaymentSourceUpdate
	return r
}

func (r CustomerPaymentSourcesApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) Execute() (*POSTCustomerPaymentSources201Response, *http.Response, error) {
	return r.ApiService.PATCHCustomerPaymentSourcesCustomerPaymentSourceIdExecute(r)
}

/*
PATCHCustomerPaymentSourcesCustomerPaymentSourceId Update a customer payment source

Update a customer payment source

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerPaymentSourceId The resource's id
	@return CustomerPaymentSourcesApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest
*/
func (a *CustomerPaymentSourcesApiService) PATCHCustomerPaymentSourcesCustomerPaymentSourceId(ctx context.Context, customerPaymentSourceId string) CustomerPaymentSourcesApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest {
	return CustomerPaymentSourcesApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		customerPaymentSourceId: customerPaymentSourceId,
	}
}

// Execute executes the request
//
//	@return POSTCustomerPaymentSources201Response
func (a *CustomerPaymentSourcesApiService) PATCHCustomerPaymentSourcesCustomerPaymentSourceIdExecute(r CustomerPaymentSourcesApiPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) (*POSTCustomerPaymentSources201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTCustomerPaymentSources201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPaymentSourcesApiService.PATCHCustomerPaymentSourcesCustomerPaymentSourceId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_payment_sources/{customerPaymentSourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerPaymentSourceId"+"}", url.PathEscape(parameterToString(r.customerPaymentSourceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerPaymentSourceUpdate == nil {
		return localVarReturnValue, nil, reportError("customerPaymentSourceUpdate is required and must be specified")
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
	localVarPostBody = r.customerPaymentSourceUpdate
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

type CustomerPaymentSourcesApiPOSTCustomerPaymentSourcesRequest struct {
	ctx                         context.Context
	ApiService                  *CustomerPaymentSourcesApiService
	customerPaymentSourceCreate *CustomerPaymentSourceCreate
}

func (r CustomerPaymentSourcesApiPOSTCustomerPaymentSourcesRequest) CustomerPaymentSourceCreate(customerPaymentSourceCreate CustomerPaymentSourceCreate) CustomerPaymentSourcesApiPOSTCustomerPaymentSourcesRequest {
	r.customerPaymentSourceCreate = &customerPaymentSourceCreate
	return r
}

func (r CustomerPaymentSourcesApiPOSTCustomerPaymentSourcesRequest) Execute() (*POSTCustomerPaymentSources201Response, *http.Response, error) {
	return r.ApiService.POSTCustomerPaymentSourcesExecute(r)
}

/*
POSTCustomerPaymentSources Create a customer payment source

Create a customer payment source

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CustomerPaymentSourcesApiPOSTCustomerPaymentSourcesRequest
*/
func (a *CustomerPaymentSourcesApiService) POSTCustomerPaymentSources(ctx context.Context) CustomerPaymentSourcesApiPOSTCustomerPaymentSourcesRequest {
	return CustomerPaymentSourcesApiPOSTCustomerPaymentSourcesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTCustomerPaymentSources201Response
func (a *CustomerPaymentSourcesApiService) POSTCustomerPaymentSourcesExecute(r CustomerPaymentSourcesApiPOSTCustomerPaymentSourcesRequest) (*POSTCustomerPaymentSources201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTCustomerPaymentSources201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerPaymentSourcesApiService.POSTCustomerPaymentSources")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_payment_sources"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerPaymentSourceCreate == nil {
		return localVarReturnValue, nil, reportError("customerPaymentSourceCreate is required and must be specified")
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
	localVarPostBody = r.customerPaymentSourceCreate
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
