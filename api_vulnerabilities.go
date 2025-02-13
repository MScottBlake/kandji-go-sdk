/*
Kandji API

<html><head></head><body><h1 id=&quot;welcome-to-the-kandji-api-documentation&quot;>Welcome to the Kandji API Documentation</h1> <p>You can find your API URL in Settings &gt; Access. The API URL will follow the below formats.</p> <ul> <li><p>US - <code>https://SubDomain.api.kandji.io</code></p> </li> <li><p>EU - <code>https://SubDomain.api.eu.kandji.io</code></p> </li> </ul> <p>For information on how to obtain an API token, please refer to the following support article.</p> <p><a href=&quot;https://support.kandji.io/api&quot;>https://support.kandji.io/api</a></p> <h4 id=&quot;rate-limit&quot;>Rate Limit</h4> <p>The Kandji API currently has an API rate limit of 10,000 requests per hour per customer.</p> <h4 id=&quot;request-methods&quot;>Request Methods</h4> <p>HTTP request methods supported by the Kandji API.</p> <div class=&quot;click-to-expand-wrapper is-table-wrapper&quot;><table> <thead> <tr> <th>Method</th> <th>Definition</th> </tr> </thead> <tbody> <tr> <td>GET</td> <td>The <code>GET</code> method requests a representation of the specified resource.</td> </tr> <tr> <td>POST</td> <td>The <code>POST</code> method submits an entity to the specified resource.</td> </tr> <tr> <td>PATCH</td> <td>The <code>PATCH</code> method applies partial modifications to a resource.</td> </tr> <tr> <td>DELETE</td> <td>The <code>DELETE</code> method deletes the specified resource.</td> </tr> </tbody> </table> </div><h4 id=&quot;response-codes&quot;>Response codes</h4> <p>Not all response codes apply to every endpoint.</p> <div class=&quot;click-to-expand-wrapper is-table-wrapper&quot;><table> <thead> <tr> <th>Code</th> <th>Response</th> </tr> </thead> <tbody> <tr> <td>200</td> <td>OK</td> </tr> <tr> <td>201</td> <td>Created</td> </tr> <tr> <td>204</td> <td>No content</td> </tr> <tr> <td></td> <td>Typical response when sending the DELETE method.</td> </tr> <tr> <td>400</td> <td>Bad Request</td> </tr> <tr> <td></td> <td>&quot;Command already running&quot; - The command may already be running in a <em>Pending</em> state waiting on the device.</td> </tr> <tr> <td></td> <td>&quot;Command is not allowed for current device&quot; - The command may not be compatible with the target device.</td> </tr> <tr> <td></td> <td>&quot;JSON parse error - Expecting ',' delimiter: line 3 column 2 (char 65)&quot;</td> </tr> <tr> <td>401</td> <td>Unauthorized</td> </tr> <tr> <td></td> <td>This error can occur if the token is incorrect, was revoked, or the token has expired.</td> </tr> <tr> <td>403</td> <td>Forbidden</td> </tr> <tr> <td></td> <td>The request was understood but cannot be authorized.</td> </tr> <tr> <td>404</td> <td>Not found</td> </tr> <tr> <td></td> <td>Unable to locate the resource in the Kandji tenant.</td> </tr> <tr> <td>415</td> <td>Unsupported Media Type</td> </tr> <tr> <td></td> <td>The request contains a media type which the server or resource does not support.</td> </tr> <tr> <td>500</td> <td>Internal server error</td> </tr> <tr> <td>503</td> <td>Service unavailable</td> </tr> <tr> <td></td> <td>This error can occur if a file upload is still being processed via the custom apps API.</td> </tr> </tbody> </table> </div><h4 id=&quot;data-structure&quot;>Data structure</h4> <p>The API returns all structured responses in JSON schema format.</p> <h4 id=&quot;examples&quot;>Examples</h4> <p>Code examples using the API can be found in the Kandji support <a href=&quot;https://github.com/kandji-inc/support/tree/main/api-tools&quot;>GitHub</a>.</p> </body></html>

API version: 1.0.0
Contact: mitchelsblake@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kandji

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type VulnerabilitiesAPI interface {

	/*
	GetVulnerabilityDescription Get Vulnerability Description

	This endpoint makes a request to retrieve information about a cve and summary information about detections for a tenants fleet.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cveId
	@return ApiGetVulnerabilityDescriptionRequest
	*/
	GetVulnerabilityDescription(ctx context.Context, cveId string) ApiGetVulnerabilityDescriptionRequest

	// GetVulnerabilityDescriptionExecute executes the request
	//  @return VulnerabilitiesGetVulnerabilityDescription200Response
	GetVulnerabilityDescriptionExecute(r ApiGetVulnerabilityDescriptionRequest) (*VulnerabilitiesGetVulnerabilityDescription200Response, *http.Response, error)

	/*
	ListAffectedApplications List Affected Applications

	This endpoint makes a request to retrieve a list of applications impacted by a specified <code>cve_id</code> vulnerability for a tenants fleet.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cveId
	@return ApiListAffectedApplicationsRequest
	*/
	ListAffectedApplications(ctx context.Context, cveId string) ApiListAffectedApplicationsRequest

	// ListAffectedApplicationsExecute executes the request
	//  @return map[string]interface{}
	ListAffectedApplicationsExecute(r ApiListAffectedApplicationsRequest) (map[string]interface{}, *http.Response, error)

	/*
	ListAffectedDevices List Affected Devices

	This endpoint makes a request to retrieve a list of devices impacted by a specified <code>cve_id</code> vulnerability for a tenants fleet.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cveId
	@return ApiListAffectedDevicesRequest
	*/
	ListAffectedDevices(ctx context.Context, cveId string) ApiListAffectedDevicesRequest

	// ListAffectedDevicesExecute executes the request
	//  @return map[string]interface{}
	ListAffectedDevicesExecute(r ApiListAffectedDevicesRequest) (map[string]interface{}, *http.Response, error)

	/*
	ListDetections List Detections

	This endpoint makes a request to retrieve a list of all vulnerability detections across the device fleet.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListDetectionsRequest
	*/
	ListDetections(ctx context.Context) ApiListDetectionsRequest

	// ListDetectionsExecute executes the request
	//  @return VulnerabilitiesListDetections200Response
	ListDetectionsExecute(r ApiListDetectionsRequest) (*VulnerabilitiesListDetections200Response, *http.Response, error)

	/*
	ListVulnerabilities List Vulnerabilities

	This endpoint makes a request to retrieve a list of all vulnerabilities grouped by cve.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListVulnerabilitiesRequest
	*/
	ListVulnerabilities(ctx context.Context) ApiListVulnerabilitiesRequest

	// ListVulnerabilitiesExecute executes the request
	//  @return VulnerabilitiesListVulnerabilities200Response
	ListVulnerabilitiesExecute(r ApiListVulnerabilitiesRequest) (*VulnerabilitiesListVulnerabilities200Response, *http.Response, error)
}

// VulnerabilitiesAPIService VulnerabilitiesAPI service
type VulnerabilitiesAPIService service

type ApiGetVulnerabilityDescriptionRequest struct {
	ctx context.Context
	ApiService VulnerabilitiesAPI
	cveId string
}

func (r ApiGetVulnerabilityDescriptionRequest) Execute() (*VulnerabilitiesGetVulnerabilityDescription200Response, *http.Response, error) {
	return r.ApiService.GetVulnerabilityDescriptionExecute(r)
}

/*
GetVulnerabilityDescription Get Vulnerability Description

This endpoint makes a request to retrieve information about a cve and summary information about detections for a tenants fleet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cveId
 @return ApiGetVulnerabilityDescriptionRequest
*/
func (a *VulnerabilitiesAPIService) GetVulnerabilityDescription(ctx context.Context, cveId string) ApiGetVulnerabilityDescriptionRequest {
	return ApiGetVulnerabilityDescriptionRequest{
		ApiService: a,
		ctx: ctx,
		cveId: cveId,
	}
}

// Execute executes the request
//  @return VulnerabilitiesGetVulnerabilityDescription200Response
func (a *VulnerabilitiesAPIService) GetVulnerabilityDescriptionExecute(r ApiGetVulnerabilityDescriptionRequest) (*VulnerabilitiesGetVulnerabilityDescription200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VulnerabilitiesGetVulnerabilityDescription200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VulnerabilitiesAPIService.GetVulnerabilityDescription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/vulnerability-management/vulnerabilities/{cve_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"cve_id"+"}", url.PathEscape(parameterValueToString(r.cveId, "cveId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ApiListAffectedApplicationsRequest struct {
	ctx context.Context
	ApiService VulnerabilitiesAPI
	cveId string
	page *string
	size *string
	sortBy *string
	filter *string
}

// The page number of the response.
func (r ApiListAffectedApplicationsRequest) Page(page string) ApiListAffectedApplicationsRequest {
	r.page = &page
	return r
}

// A hard upper limit is set at 50  records returned per request. If more records are expected, pagination should be used using the URL value returned in the next attribute. Additionally, filters can be added to a request to limit the results.
func (r ApiListAffectedApplicationsRequest) Size(size string) ApiListAffectedApplicationsRequest {
	r.size = &size
	return r
}

// Field to sort by. Example: sort_by&#x3D;app_name.
func (r ApiListAffectedApplicationsRequest) SortBy(sortBy string) ApiListAffectedApplicationsRequest {
	r.sortBy = &sortBy
	return r
}

// Filterable columns: blueprint_id updated_at
func (r ApiListAffectedApplicationsRequest) Filter(filter string) ApiListAffectedApplicationsRequest {
	r.filter = &filter
	return r
}

func (r ApiListAffectedApplicationsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.ListAffectedApplicationsExecute(r)
}

/*
ListAffectedApplications List Affected Applications

This endpoint makes a request to retrieve a list of applications impacted by a specified <code>cve_id</code> vulnerability for a tenants fleet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cveId
 @return ApiListAffectedApplicationsRequest
*/
func (a *VulnerabilitiesAPIService) ListAffectedApplications(ctx context.Context, cveId string) ApiListAffectedApplicationsRequest {
	return ApiListAffectedApplicationsRequest{
		ApiService: a,
		ctx: ctx,
		cveId: cveId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *VulnerabilitiesAPIService) ListAffectedApplicationsExecute(r ApiListAffectedApplicationsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VulnerabilitiesAPIService.ListAffectedApplications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/vulnerability-management/vulnerabilities/{cve_id}/applications"
	localVarPath = strings.Replace(localVarPath, "{"+"cve_id"+"}", url.PathEscape(parameterValueToString(r.cveId, "cveId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.size != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_by", r.sortBy, "form", "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ApiListAffectedDevicesRequest struct {
	ctx context.Context
	ApiService VulnerabilitiesAPI
	cveId string
	page *string
	size *string
	sortBy *string
	filter *string
}

// The page number of the response.
func (r ApiListAffectedDevicesRequest) Page(page string) ApiListAffectedDevicesRequest {
	r.page = &page
	return r
}

// A hard upper limit is set at 50  records returned per request. If more records are expected, pagination should be used using the URL value returned in the next attribute. Additionally, filters can be added to a request to limit the results.
func (r ApiListAffectedDevicesRequest) Size(size string) ApiListAffectedDevicesRequest {
	r.size = &size
	return r
}

// Field to sort by. Example: sort_by&#x3D;app_name.
func (r ApiListAffectedDevicesRequest) SortBy(sortBy string) ApiListAffectedDevicesRequest {
	r.sortBy = &sortBy
	return r
}

// Filterable columns: blueprint_id updated_at
func (r ApiListAffectedDevicesRequest) Filter(filter string) ApiListAffectedDevicesRequest {
	r.filter = &filter
	return r
}

func (r ApiListAffectedDevicesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.ListAffectedDevicesExecute(r)
}

/*
ListAffectedDevices List Affected Devices

This endpoint makes a request to retrieve a list of devices impacted by a specified <code>cve_id</code> vulnerability for a tenants fleet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param cveId
 @return ApiListAffectedDevicesRequest
*/
func (a *VulnerabilitiesAPIService) ListAffectedDevices(ctx context.Context, cveId string) ApiListAffectedDevicesRequest {
	return ApiListAffectedDevicesRequest{
		ApiService: a,
		ctx: ctx,
		cveId: cveId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *VulnerabilitiesAPIService) ListAffectedDevicesExecute(r ApiListAffectedDevicesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VulnerabilitiesAPIService.ListAffectedDevices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/vulnerability-management/vulnerabilities/{cve_id}/devices"
	localVarPath = strings.Replace(localVarPath, "{"+"cve_id"+"}", url.PathEscape(parameterValueToString(r.cveId, "cveId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.size != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_by", r.sortBy, "form", "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ApiListDetectionsRequest struct {
	ctx context.Context
	ApiService VulnerabilitiesAPI
	after *string
	limit *string
	filter *string
}

// Cursor token.
func (r ApiListDetectionsRequest) After(after string) ApiListDetectionsRequest {
	r.after = &after
	return r
}

// A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300  records returned per request. If more records are expected, pagination should be used using the URL value returned in the next attribute. Additionally, filters can be added to a request to limit the results.
func (r ApiListDetectionsRequest) Limit(limit string) ApiListDetectionsRequest {
	r.limit = &limit
	return r
}

// Can filter on any key attribute within the response.
func (r ApiListDetectionsRequest) Filter(filter string) ApiListDetectionsRequest {
	r.filter = &filter
	return r
}

func (r ApiListDetectionsRequest) Execute() (*VulnerabilitiesListDetections200Response, *http.Response, error) {
	return r.ApiService.ListDetectionsExecute(r)
}

/*
ListDetections List Detections

This endpoint makes a request to retrieve a list of all vulnerability detections across the device fleet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDetectionsRequest
*/
func (a *VulnerabilitiesAPIService) ListDetections(ctx context.Context) ApiListDetectionsRequest {
	return ApiListDetectionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return VulnerabilitiesListDetections200Response
func (a *VulnerabilitiesAPIService) ListDetectionsExecute(r ApiListDetectionsRequest) (*VulnerabilitiesListDetections200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VulnerabilitiesListDetections200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VulnerabilitiesAPIService.ListDetections")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/vulnerability-management/detections"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.after != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ApiListVulnerabilitiesRequest struct {
	ctx context.Context
	ApiService VulnerabilitiesAPI
	page *string
	size *string
	sortBy *string
	filter *string
}

// The page number of the response.
func (r ApiListVulnerabilitiesRequest) Page(page string) ApiListVulnerabilitiesRequest {
	r.page = &page
	return r
}

// A hard upper limit is set at 50  records returned per request. If more records are expected, pagination should be used using the URL value returned in the next attribute. Additionally, filters can be added to a request to limit the results.
func (r ApiListVulnerabilitiesRequest) Size(size string) ApiListVulnerabilitiesRequest {
	r.size = &size
	return r
}

// Field to sort by. Example: sort_by&#x3D;cve_id.
func (r ApiListVulnerabilitiesRequest) SortBy(sortBy string) ApiListVulnerabilitiesRequest {
	r.sortBy = &sortBy
	return r
}

// &lt;p&gt;Filterable columns:&lt;/p&gt; &lt;p&gt;cve_id app_name severity first_detection_date latest_detection_date&lt;/p&gt;
func (r ApiListVulnerabilitiesRequest) Filter(filter string) ApiListVulnerabilitiesRequest {
	r.filter = &filter
	return r
}

func (r ApiListVulnerabilitiesRequest) Execute() (*VulnerabilitiesListVulnerabilities200Response, *http.Response, error) {
	return r.ApiService.ListVulnerabilitiesExecute(r)
}

/*
ListVulnerabilities List Vulnerabilities

This endpoint makes a request to retrieve a list of all vulnerabilities grouped by cve.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListVulnerabilitiesRequest
*/
func (a *VulnerabilitiesAPIService) ListVulnerabilities(ctx context.Context) ApiListVulnerabilitiesRequest {
	return ApiListVulnerabilitiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return VulnerabilitiesListVulnerabilities200Response
func (a *VulnerabilitiesAPIService) ListVulnerabilitiesExecute(r ApiListVulnerabilitiesRequest) (*VulnerabilitiesListVulnerabilities200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VulnerabilitiesListVulnerabilities200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VulnerabilitiesAPIService.ListVulnerabilities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/vulnerability-management/vulnerabilities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.size != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_by", r.sortBy, "form", "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
