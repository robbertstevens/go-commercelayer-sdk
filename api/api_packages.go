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

// PackagesApiService PackagesApi service
type PackagesApiService service

type PackagesApiDELETEPackagesPackageIdRequest struct {
	ctx        context.Context
	ApiService *PackagesApiService
	packageId  string
}

func (r PackagesApiDELETEPackagesPackageIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEPackagesPackageIdExecute(r)
}

/*
DELETEPackagesPackageId Delete a package

Delete a package

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param packageId The resource's id
 @return PackagesApiDELETEPackagesPackageIdRequest
*/
func (a *PackagesApiService) DELETEPackagesPackageId(ctx context.Context, packageId string) PackagesApiDELETEPackagesPackageIdRequest {
	return PackagesApiDELETEPackagesPackageIdRequest{
		ApiService: a,
		ctx:        ctx,
		packageId:  packageId,
	}
}

// Execute executes the request
func (a *PackagesApiService) DELETEPackagesPackageIdExecute(r PackagesApiDELETEPackagesPackageIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackagesApiService.DELETEPackagesPackageId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/packages/{packageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"packageId"+"}", url.PathEscape(parameterToString(r.packageId, "")), -1)

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

type PackagesApiGETPackagesRequest struct {
	ctx        context.Context
	ApiService *PackagesApiService
}

func (r PackagesApiGETPackagesRequest) Execute() (*GETPackages200Response, *http.Response, error) {
	return r.ApiService.GETPackagesExecute(r)
}

/*
GETPackages List all packages

List all packages

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PackagesApiGETPackagesRequest
*/
func (a *PackagesApiService) GETPackages(ctx context.Context) PackagesApiGETPackagesRequest {
	return PackagesApiGETPackagesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GETPackages200Response
func (a *PackagesApiService) GETPackagesExecute(r PackagesApiGETPackagesRequest) (*GETPackages200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETPackages200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackagesApiService.GETPackages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/packages"

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

type PackagesApiGETPackagesPackageIdRequest struct {
	ctx        context.Context
	ApiService *PackagesApiService
	packageId  string
}

func (r PackagesApiGETPackagesPackageIdRequest) Execute() (*GETPackagesPackageId200Response, *http.Response, error) {
	return r.ApiService.GETPackagesPackageIdExecute(r)
}

/*
GETPackagesPackageId Retrieve a package

Retrieve a package

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param packageId The resource's id
 @return PackagesApiGETPackagesPackageIdRequest
*/
func (a *PackagesApiService) GETPackagesPackageId(ctx context.Context, packageId string) PackagesApiGETPackagesPackageIdRequest {
	return PackagesApiGETPackagesPackageIdRequest{
		ApiService: a,
		ctx:        ctx,
		packageId:  packageId,
	}
}

// Execute executes the request
//  @return GETPackagesPackageId200Response
func (a *PackagesApiService) GETPackagesPackageIdExecute(r PackagesApiGETPackagesPackageIdRequest) (*GETPackagesPackageId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETPackagesPackageId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackagesApiService.GETPackagesPackageId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/packages/{packageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"packageId"+"}", url.PathEscape(parameterToString(r.packageId, "")), -1)

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

type PackagesApiGETParcelIdPackageRequest struct {
	ctx        context.Context
	ApiService *PackagesApiService
	parcelId   string
}

func (r PackagesApiGETParcelIdPackageRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETParcelIdPackageExecute(r)
}

/*
GETParcelIdPackage Retrieve the package associated to the parcel

Retrieve the package associated to the parcel

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parcelId The resource's id
 @return PackagesApiGETParcelIdPackageRequest
*/
func (a *PackagesApiService) GETParcelIdPackage(ctx context.Context, parcelId string) PackagesApiGETParcelIdPackageRequest {
	return PackagesApiGETParcelIdPackageRequest{
		ApiService: a,
		ctx:        ctx,
		parcelId:   parcelId,
	}
}

// Execute executes the request
func (a *PackagesApiService) GETParcelIdPackageExecute(r PackagesApiGETParcelIdPackageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackagesApiService.GETParcelIdPackage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parcels/{parcelId}/package"
	localVarPath = strings.Replace(localVarPath, "{"+"parcelId"+"}", url.PathEscape(parameterToString(r.parcelId, "")), -1)

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

type PackagesApiPATCHPackagesPackageIdRequest struct {
	ctx           context.Context
	ApiService    *PackagesApiService
	packageUpdate *PackageUpdate
	packageId     string
}

func (r PackagesApiPATCHPackagesPackageIdRequest) PackageUpdate(packageUpdate PackageUpdate) PackagesApiPATCHPackagesPackageIdRequest {
	r.packageUpdate = &packageUpdate
	return r
}

func (r PackagesApiPATCHPackagesPackageIdRequest) Execute() (*PATCHPackagesPackageId200Response, *http.Response, error) {
	return r.ApiService.PATCHPackagesPackageIdExecute(r)
}

/*
PATCHPackagesPackageId Update a package

Update a package

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param packageId The resource's id
 @return PackagesApiPATCHPackagesPackageIdRequest
*/
func (a *PackagesApiService) PATCHPackagesPackageId(ctx context.Context, packageId string) PackagesApiPATCHPackagesPackageIdRequest {
	return PackagesApiPATCHPackagesPackageIdRequest{
		ApiService: a,
		ctx:        ctx,
		packageId:  packageId,
	}
}

// Execute executes the request
//  @return PATCHPackagesPackageId200Response
func (a *PackagesApiService) PATCHPackagesPackageIdExecute(r PackagesApiPATCHPackagesPackageIdRequest) (*PATCHPackagesPackageId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHPackagesPackageId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackagesApiService.PATCHPackagesPackageId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/packages/{packageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"packageId"+"}", url.PathEscape(parameterToString(r.packageId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.packageUpdate == nil {
		return localVarReturnValue, nil, reportError("packageUpdate is required and must be specified")
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
	localVarPostBody = r.packageUpdate
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

type PackagesApiPOSTPackagesRequest struct {
	ctx           context.Context
	ApiService    *PackagesApiService
	packageCreate *PackageCreate
}

func (r PackagesApiPOSTPackagesRequest) PackageCreate(packageCreate PackageCreate) PackagesApiPOSTPackagesRequest {
	r.packageCreate = &packageCreate
	return r
}

func (r PackagesApiPOSTPackagesRequest) Execute() (*POSTPackages201Response, *http.Response, error) {
	return r.ApiService.POSTPackagesExecute(r)
}

/*
POSTPackages Create a package

Create a package

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PackagesApiPOSTPackagesRequest
*/
func (a *PackagesApiService) POSTPackages(ctx context.Context) PackagesApiPOSTPackagesRequest {
	return PackagesApiPOSTPackagesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return POSTPackages201Response
func (a *PackagesApiService) POSTPackagesExecute(r PackagesApiPOSTPackagesRequest) (*POSTPackages201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTPackages201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackagesApiService.POSTPackages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/packages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.packageCreate == nil {
		return localVarReturnValue, nil, reportError("packageCreate is required and must be specified")
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
	localVarPostBody = r.packageCreate
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
