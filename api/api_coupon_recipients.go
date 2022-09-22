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

// CouponRecipientsApiService CouponRecipientsApi service
type CouponRecipientsApiService service

type CouponRecipientsApiDELETECouponRecipientsCouponRecipientIdRequest struct {
	ctx               context.Context
	ApiService        *CouponRecipientsApiService
	couponRecipientId string
}

func (r CouponRecipientsApiDELETECouponRecipientsCouponRecipientIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETECouponRecipientsCouponRecipientIdExecute(r)
}

/*
DELETECouponRecipientsCouponRecipientId Delete a coupon recipient

Delete a coupon recipient

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param couponRecipientId The resource's id
 @return CouponRecipientsApiDELETECouponRecipientsCouponRecipientIdRequest
*/
func (a *CouponRecipientsApiService) DELETECouponRecipientsCouponRecipientId(ctx context.Context, couponRecipientId string) CouponRecipientsApiDELETECouponRecipientsCouponRecipientIdRequest {
	return CouponRecipientsApiDELETECouponRecipientsCouponRecipientIdRequest{
		ApiService:        a,
		ctx:               ctx,
		couponRecipientId: couponRecipientId,
	}
}

// Execute executes the request
func (a *CouponRecipientsApiService) DELETECouponRecipientsCouponRecipientIdExecute(r CouponRecipientsApiDELETECouponRecipientsCouponRecipientIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CouponRecipientsApiService.DELETECouponRecipientsCouponRecipientId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/coupon_recipients/{couponRecipientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"couponRecipientId"+"}", url.PathEscape(parameterToString(r.couponRecipientId, "")), -1)

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

type CouponRecipientsApiGETCouponRecipientsRequest struct {
	ctx        context.Context
	ApiService *CouponRecipientsApiService
}

func (r CouponRecipientsApiGETCouponRecipientsRequest) Execute() (*GETCouponRecipients200Response, *http.Response, error) {
	return r.ApiService.GETCouponRecipientsExecute(r)
}

/*
GETCouponRecipients List all coupon recipients

List all coupon recipients

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CouponRecipientsApiGETCouponRecipientsRequest
*/
func (a *CouponRecipientsApiService) GETCouponRecipients(ctx context.Context) CouponRecipientsApiGETCouponRecipientsRequest {
	return CouponRecipientsApiGETCouponRecipientsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GETCouponRecipients200Response
func (a *CouponRecipientsApiService) GETCouponRecipientsExecute(r CouponRecipientsApiGETCouponRecipientsRequest) (*GETCouponRecipients200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCouponRecipients200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CouponRecipientsApiService.GETCouponRecipients")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/coupon_recipients"

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

type CouponRecipientsApiGETCouponRecipientsCouponRecipientIdRequest struct {
	ctx               context.Context
	ApiService        *CouponRecipientsApiService
	couponRecipientId string
}

func (r CouponRecipientsApiGETCouponRecipientsCouponRecipientIdRequest) Execute() (*GETCouponRecipientsCouponRecipientId200Response, *http.Response, error) {
	return r.ApiService.GETCouponRecipientsCouponRecipientIdExecute(r)
}

/*
GETCouponRecipientsCouponRecipientId Retrieve a coupon recipient

Retrieve a coupon recipient

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param couponRecipientId The resource's id
 @return CouponRecipientsApiGETCouponRecipientsCouponRecipientIdRequest
*/
func (a *CouponRecipientsApiService) GETCouponRecipientsCouponRecipientId(ctx context.Context, couponRecipientId string) CouponRecipientsApiGETCouponRecipientsCouponRecipientIdRequest {
	return CouponRecipientsApiGETCouponRecipientsCouponRecipientIdRequest{
		ApiService:        a,
		ctx:               ctx,
		couponRecipientId: couponRecipientId,
	}
}

// Execute executes the request
//  @return GETCouponRecipientsCouponRecipientId200Response
func (a *CouponRecipientsApiService) GETCouponRecipientsCouponRecipientIdExecute(r CouponRecipientsApiGETCouponRecipientsCouponRecipientIdRequest) (*GETCouponRecipientsCouponRecipientId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCouponRecipientsCouponRecipientId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CouponRecipientsApiService.GETCouponRecipientsCouponRecipientId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/coupon_recipients/{couponRecipientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"couponRecipientId"+"}", url.PathEscape(parameterToString(r.couponRecipientId, "")), -1)

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

type CouponRecipientsApiPATCHCouponRecipientsCouponRecipientIdRequest struct {
	ctx                   context.Context
	ApiService            *CouponRecipientsApiService
	couponRecipientUpdate *CouponRecipientUpdate
	couponRecipientId     string
}

func (r CouponRecipientsApiPATCHCouponRecipientsCouponRecipientIdRequest) CouponRecipientUpdate(couponRecipientUpdate CouponRecipientUpdate) CouponRecipientsApiPATCHCouponRecipientsCouponRecipientIdRequest {
	r.couponRecipientUpdate = &couponRecipientUpdate
	return r
}

func (r CouponRecipientsApiPATCHCouponRecipientsCouponRecipientIdRequest) Execute() (*PATCHCouponRecipientsCouponRecipientId200Response, *http.Response, error) {
	return r.ApiService.PATCHCouponRecipientsCouponRecipientIdExecute(r)
}

/*
PATCHCouponRecipientsCouponRecipientId Update a coupon recipient

Update a coupon recipient

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param couponRecipientId The resource's id
 @return CouponRecipientsApiPATCHCouponRecipientsCouponRecipientIdRequest
*/
func (a *CouponRecipientsApiService) PATCHCouponRecipientsCouponRecipientId(ctx context.Context, couponRecipientId string) CouponRecipientsApiPATCHCouponRecipientsCouponRecipientIdRequest {
	return CouponRecipientsApiPATCHCouponRecipientsCouponRecipientIdRequest{
		ApiService:        a,
		ctx:               ctx,
		couponRecipientId: couponRecipientId,
	}
}

// Execute executes the request
//  @return PATCHCouponRecipientsCouponRecipientId200Response
func (a *CouponRecipientsApiService) PATCHCouponRecipientsCouponRecipientIdExecute(r CouponRecipientsApiPATCHCouponRecipientsCouponRecipientIdRequest) (*PATCHCouponRecipientsCouponRecipientId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHCouponRecipientsCouponRecipientId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CouponRecipientsApiService.PATCHCouponRecipientsCouponRecipientId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/coupon_recipients/{couponRecipientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"couponRecipientId"+"}", url.PathEscape(parameterToString(r.couponRecipientId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.couponRecipientUpdate == nil {
		return localVarReturnValue, nil, reportError("couponRecipientUpdate is required and must be specified")
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
	localVarPostBody = r.couponRecipientUpdate
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

type CouponRecipientsApiPOSTCouponRecipientsRequest struct {
	ctx                   context.Context
	ApiService            *CouponRecipientsApiService
	couponRecipientCreate *CouponRecipientCreate
}

func (r CouponRecipientsApiPOSTCouponRecipientsRequest) CouponRecipientCreate(couponRecipientCreate CouponRecipientCreate) CouponRecipientsApiPOSTCouponRecipientsRequest {
	r.couponRecipientCreate = &couponRecipientCreate
	return r
}

func (r CouponRecipientsApiPOSTCouponRecipientsRequest) Execute() (*POSTCouponRecipients201Response, *http.Response, error) {
	return r.ApiService.POSTCouponRecipientsExecute(r)
}

/*
POSTCouponRecipients Create a coupon recipient

Create a coupon recipient

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CouponRecipientsApiPOSTCouponRecipientsRequest
*/
func (a *CouponRecipientsApiService) POSTCouponRecipients(ctx context.Context) CouponRecipientsApiPOSTCouponRecipientsRequest {
	return CouponRecipientsApiPOSTCouponRecipientsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return POSTCouponRecipients201Response
func (a *CouponRecipientsApiService) POSTCouponRecipientsExecute(r CouponRecipientsApiPOSTCouponRecipientsRequest) (*POSTCouponRecipients201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTCouponRecipients201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CouponRecipientsApiService.POSTCouponRecipients")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/coupon_recipients"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.couponRecipientCreate == nil {
		return localVarReturnValue, nil, reportError("couponRecipientCreate is required and must be specified")
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
	localVarPostBody = r.couponRecipientCreate
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
