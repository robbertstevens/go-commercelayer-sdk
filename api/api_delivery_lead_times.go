/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
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

// DeliveryLeadTimesApiService DeliveryLeadTimesApi service
type DeliveryLeadTimesApiService service

type DeliveryLeadTimesApiDELETEDeliveryLeadTimesDeliveryLeadTimeIdRequest struct {
	ctx                context.Context
	ApiService         *DeliveryLeadTimesApiService
	deliveryLeadTimeId string
}

func (r DeliveryLeadTimesApiDELETEDeliveryLeadTimesDeliveryLeadTimeIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEDeliveryLeadTimesDeliveryLeadTimeIdExecute(r)
}

/*
DELETEDeliveryLeadTimesDeliveryLeadTimeId Delete a delivery lead time

Delete a delivery lead time

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deliveryLeadTimeId The resource's id
 @return DeliveryLeadTimesApiDELETEDeliveryLeadTimesDeliveryLeadTimeIdRequest
*/
func (a *DeliveryLeadTimesApiService) DELETEDeliveryLeadTimesDeliveryLeadTimeId(ctx context.Context, deliveryLeadTimeId string) DeliveryLeadTimesApiDELETEDeliveryLeadTimesDeliveryLeadTimeIdRequest {
	return DeliveryLeadTimesApiDELETEDeliveryLeadTimesDeliveryLeadTimeIdRequest{
		ApiService:         a,
		ctx:                ctx,
		deliveryLeadTimeId: deliveryLeadTimeId,
	}
}

// Execute executes the request
func (a *DeliveryLeadTimesApiService) DELETEDeliveryLeadTimesDeliveryLeadTimeIdExecute(r DeliveryLeadTimesApiDELETEDeliveryLeadTimesDeliveryLeadTimeIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeliveryLeadTimesApiService.DELETEDeliveryLeadTimesDeliveryLeadTimeId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/delivery_lead_times/{deliveryLeadTimeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"deliveryLeadTimeId"+"}", url.PathEscape(parameterToString(r.deliveryLeadTimeId, "")), -1)

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

type DeliveryLeadTimesApiGETDeliveryLeadTimesRequest struct {
	ctx        context.Context
	ApiService *DeliveryLeadTimesApiService
}

func (r DeliveryLeadTimesApiGETDeliveryLeadTimesRequest) Execute() (*GETDeliveryLeadTimes200Response, *http.Response, error) {
	return r.ApiService.GETDeliveryLeadTimesExecute(r)
}

/*
GETDeliveryLeadTimes List all delivery lead times

List all delivery lead times

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DeliveryLeadTimesApiGETDeliveryLeadTimesRequest
*/
func (a *DeliveryLeadTimesApiService) GETDeliveryLeadTimes(ctx context.Context) DeliveryLeadTimesApiGETDeliveryLeadTimesRequest {
	return DeliveryLeadTimesApiGETDeliveryLeadTimesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GETDeliveryLeadTimes200Response
func (a *DeliveryLeadTimesApiService) GETDeliveryLeadTimesExecute(r DeliveryLeadTimesApiGETDeliveryLeadTimesRequest) (*GETDeliveryLeadTimes200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETDeliveryLeadTimes200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeliveryLeadTimesApiService.GETDeliveryLeadTimes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/delivery_lead_times"

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

type DeliveryLeadTimesApiGETDeliveryLeadTimesDeliveryLeadTimeIdRequest struct {
	ctx                context.Context
	ApiService         *DeliveryLeadTimesApiService
	deliveryLeadTimeId string
}

func (r DeliveryLeadTimesApiGETDeliveryLeadTimesDeliveryLeadTimeIdRequest) Execute() (*GETDeliveryLeadTimesDeliveryLeadTimeId200Response, *http.Response, error) {
	return r.ApiService.GETDeliveryLeadTimesDeliveryLeadTimeIdExecute(r)
}

/*
GETDeliveryLeadTimesDeliveryLeadTimeId Retrieve a delivery lead time

Retrieve a delivery lead time

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deliveryLeadTimeId The resource's id
 @return DeliveryLeadTimesApiGETDeliveryLeadTimesDeliveryLeadTimeIdRequest
*/
func (a *DeliveryLeadTimesApiService) GETDeliveryLeadTimesDeliveryLeadTimeId(ctx context.Context, deliveryLeadTimeId string) DeliveryLeadTimesApiGETDeliveryLeadTimesDeliveryLeadTimeIdRequest {
	return DeliveryLeadTimesApiGETDeliveryLeadTimesDeliveryLeadTimeIdRequest{
		ApiService:         a,
		ctx:                ctx,
		deliveryLeadTimeId: deliveryLeadTimeId,
	}
}

// Execute executes the request
//  @return GETDeliveryLeadTimesDeliveryLeadTimeId200Response
func (a *DeliveryLeadTimesApiService) GETDeliveryLeadTimesDeliveryLeadTimeIdExecute(r DeliveryLeadTimesApiGETDeliveryLeadTimesDeliveryLeadTimeIdRequest) (*GETDeliveryLeadTimesDeliveryLeadTimeId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETDeliveryLeadTimesDeliveryLeadTimeId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeliveryLeadTimesApiService.GETDeliveryLeadTimesDeliveryLeadTimeId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/delivery_lead_times/{deliveryLeadTimeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"deliveryLeadTimeId"+"}", url.PathEscape(parameterToString(r.deliveryLeadTimeId, "")), -1)

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

type DeliveryLeadTimesApiGETShipmentIdDeliveryLeadTimeRequest struct {
	ctx        context.Context
	ApiService *DeliveryLeadTimesApiService
	shipmentId string
}

func (r DeliveryLeadTimesApiGETShipmentIdDeliveryLeadTimeRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShipmentIdDeliveryLeadTimeExecute(r)
}

/*
GETShipmentIdDeliveryLeadTime Retrieve the delivery lead time associated to the shipment

Retrieve the delivery lead time associated to the shipment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shipmentId The resource's id
 @return DeliveryLeadTimesApiGETShipmentIdDeliveryLeadTimeRequest
*/
func (a *DeliveryLeadTimesApiService) GETShipmentIdDeliveryLeadTime(ctx context.Context, shipmentId string) DeliveryLeadTimesApiGETShipmentIdDeliveryLeadTimeRequest {
	return DeliveryLeadTimesApiGETShipmentIdDeliveryLeadTimeRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
func (a *DeliveryLeadTimesApiService) GETShipmentIdDeliveryLeadTimeExecute(r DeliveryLeadTimesApiGETShipmentIdDeliveryLeadTimeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeliveryLeadTimesApiService.GETShipmentIdDeliveryLeadTime")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipments/{shipmentId}/delivery_lead_time"
	localVarPath = strings.Replace(localVarPath, "{"+"shipmentId"+"}", url.PathEscape(parameterToString(r.shipmentId, "")), -1)

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

type DeliveryLeadTimesApiGETShippingMethodIdDeliveryLeadTimeForShipmentRequest struct {
	ctx              context.Context
	ApiService       *DeliveryLeadTimesApiService
	shippingMethodId string
}

func (r DeliveryLeadTimesApiGETShippingMethodIdDeliveryLeadTimeForShipmentRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShippingMethodIdDeliveryLeadTimeForShipmentExecute(r)
}

/*
GETShippingMethodIdDeliveryLeadTimeForShipment Retrieve the delivery lead time for shipment associated to the shipping method

Retrieve the delivery lead time for shipment associated to the shipping method

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingMethodId The resource's id
 @return DeliveryLeadTimesApiGETShippingMethodIdDeliveryLeadTimeForShipmentRequest
*/
func (a *DeliveryLeadTimesApiService) GETShippingMethodIdDeliveryLeadTimeForShipment(ctx context.Context, shippingMethodId string) DeliveryLeadTimesApiGETShippingMethodIdDeliveryLeadTimeForShipmentRequest {
	return DeliveryLeadTimesApiGETShippingMethodIdDeliveryLeadTimeForShipmentRequest{
		ApiService:       a,
		ctx:              ctx,
		shippingMethodId: shippingMethodId,
	}
}

// Execute executes the request
func (a *DeliveryLeadTimesApiService) GETShippingMethodIdDeliveryLeadTimeForShipmentExecute(r DeliveryLeadTimesApiGETShippingMethodIdDeliveryLeadTimeForShipmentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeliveryLeadTimesApiService.GETShippingMethodIdDeliveryLeadTimeForShipment")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_methods/{shippingMethodId}/delivery_lead_time_for_shipment"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingMethodId"+"}", url.PathEscape(parameterToString(r.shippingMethodId, "")), -1)

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

type DeliveryLeadTimesApiGETSkuIdDeliveryLeadTimesRequest struct {
	ctx        context.Context
	ApiService *DeliveryLeadTimesApiService
	skuId      string
}

func (r DeliveryLeadTimesApiGETSkuIdDeliveryLeadTimesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETSkuIdDeliveryLeadTimesExecute(r)
}

/*
GETSkuIdDeliveryLeadTimes Retrieve the delivery lead times associated to the SKU

Retrieve the delivery lead times associated to the SKU

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skuId The resource's id
 @return DeliveryLeadTimesApiGETSkuIdDeliveryLeadTimesRequest
*/
func (a *DeliveryLeadTimesApiService) GETSkuIdDeliveryLeadTimes(ctx context.Context, skuId string) DeliveryLeadTimesApiGETSkuIdDeliveryLeadTimesRequest {
	return DeliveryLeadTimesApiGETSkuIdDeliveryLeadTimesRequest{
		ApiService: a,
		ctx:        ctx,
		skuId:      skuId,
	}
}

// Execute executes the request
func (a *DeliveryLeadTimesApiService) GETSkuIdDeliveryLeadTimesExecute(r DeliveryLeadTimesApiGETSkuIdDeliveryLeadTimesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeliveryLeadTimesApiService.GETSkuIdDeliveryLeadTimes")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/skus/{skuId}/delivery_lead_times"
	localVarPath = strings.Replace(localVarPath, "{"+"skuId"+"}", url.PathEscape(parameterToString(r.skuId, "")), -1)

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

type DeliveryLeadTimesApiPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest struct {
	ctx                    context.Context
	ApiService             *DeliveryLeadTimesApiService
	deliveryLeadTimeUpdate *DeliveryLeadTimeUpdate
	deliveryLeadTimeId     string
}

func (r DeliveryLeadTimesApiPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest) DeliveryLeadTimeUpdate(deliveryLeadTimeUpdate DeliveryLeadTimeUpdate) DeliveryLeadTimesApiPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest {
	r.deliveryLeadTimeUpdate = &deliveryLeadTimeUpdate
	return r
}

func (r DeliveryLeadTimesApiPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest) Execute() (*PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response, *http.Response, error) {
	return r.ApiService.PATCHDeliveryLeadTimesDeliveryLeadTimeIdExecute(r)
}

/*
PATCHDeliveryLeadTimesDeliveryLeadTimeId Update a delivery lead time

Update a delivery lead time

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deliveryLeadTimeId The resource's id
 @return DeliveryLeadTimesApiPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest
*/
func (a *DeliveryLeadTimesApiService) PATCHDeliveryLeadTimesDeliveryLeadTimeId(ctx context.Context, deliveryLeadTimeId string) DeliveryLeadTimesApiPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest {
	return DeliveryLeadTimesApiPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest{
		ApiService:         a,
		ctx:                ctx,
		deliveryLeadTimeId: deliveryLeadTimeId,
	}
}

// Execute executes the request
//  @return PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response
func (a *DeliveryLeadTimesApiService) PATCHDeliveryLeadTimesDeliveryLeadTimeIdExecute(r DeliveryLeadTimesApiPATCHDeliveryLeadTimesDeliveryLeadTimeIdRequest) (*PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeliveryLeadTimesApiService.PATCHDeliveryLeadTimesDeliveryLeadTimeId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/delivery_lead_times/{deliveryLeadTimeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"deliveryLeadTimeId"+"}", url.PathEscape(parameterToString(r.deliveryLeadTimeId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deliveryLeadTimeUpdate == nil {
		return localVarReturnValue, nil, reportError("deliveryLeadTimeUpdate is required and must be specified")
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
	localVarPostBody = r.deliveryLeadTimeUpdate
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

type DeliveryLeadTimesApiPOSTDeliveryLeadTimesRequest struct {
	ctx                    context.Context
	ApiService             *DeliveryLeadTimesApiService
	deliveryLeadTimeCreate *DeliveryLeadTimeCreate
}

func (r DeliveryLeadTimesApiPOSTDeliveryLeadTimesRequest) DeliveryLeadTimeCreate(deliveryLeadTimeCreate DeliveryLeadTimeCreate) DeliveryLeadTimesApiPOSTDeliveryLeadTimesRequest {
	r.deliveryLeadTimeCreate = &deliveryLeadTimeCreate
	return r
}

func (r DeliveryLeadTimesApiPOSTDeliveryLeadTimesRequest) Execute() (*POSTDeliveryLeadTimes201Response, *http.Response, error) {
	return r.ApiService.POSTDeliveryLeadTimesExecute(r)
}

/*
POSTDeliveryLeadTimes Create a delivery lead time

Create a delivery lead time

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DeliveryLeadTimesApiPOSTDeliveryLeadTimesRequest
*/
func (a *DeliveryLeadTimesApiService) POSTDeliveryLeadTimes(ctx context.Context) DeliveryLeadTimesApiPOSTDeliveryLeadTimesRequest {
	return DeliveryLeadTimesApiPOSTDeliveryLeadTimesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return POSTDeliveryLeadTimes201Response
func (a *DeliveryLeadTimesApiService) POSTDeliveryLeadTimesExecute(r DeliveryLeadTimesApiPOSTDeliveryLeadTimesRequest) (*POSTDeliveryLeadTimes201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTDeliveryLeadTimes201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeliveryLeadTimesApiService.POSTDeliveryLeadTimes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/delivery_lead_times"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deliveryLeadTimeCreate == nil {
		return localVarReturnValue, nil, reportError("deliveryLeadTimeCreate is required and must be specified")
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
	localVarPostBody = r.deliveryLeadTimeCreate
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
