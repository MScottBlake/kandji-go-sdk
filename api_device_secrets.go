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


// DeviceSecretsAPIService DeviceSecretsAPI service
type DeviceSecretsAPIService service

type ApiGetActivationLockBypassCodeRequest struct {
	ctx context.Context
	ApiService *DeviceSecretsAPIService
	deviceId string
}

func (r ApiGetActivationLockBypassCodeRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetActivationLockBypassCodeExecute(r)
}

/*
GetActivationLockBypassCode Get Activation Lock Bypass Code

<p>This request allows you to retrieve the Activation Lock Bypass code.</p>
<p>user_based_albc is the user-based Activation Lock bypass code for when Activation Lock is enabled using an personal Apple ID and Find My.</p>
<p>device_based_albc is the device-based Activation Lock bypass code for when Activation Lock is enabled by the MDM server.</p>
<h3 id=&quot;request-parameters&quot;>Request Parameters</h3>
<p><code>device_id</code> (path parameter): The unique identifier of the device.</p>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceId
 @return ApiGetActivationLockBypassCodeRequest
*/
func (a *DeviceSecretsAPIService) GetActivationLockBypassCode(ctx context.Context, deviceId string) ApiGetActivationLockBypassCodeRequest {
	return ApiGetActivationLockBypassCodeRequest{
		ApiService: a,
		ctx: ctx,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DeviceSecretsAPIService) GetActivationLockBypassCodeExecute(r ApiGetActivationLockBypassCodeRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceSecretsAPIService.GetActivationLockBypassCode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/devices/{device_id}/secrets/bypasscode"
	localVarPath = strings.Replace(localVarPath, "{"+"device_id"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

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

type ApiGetFilevaultRecoveryKeyRequest struct {
	ctx context.Context
	ApiService *DeviceSecretsAPIService
	deviceId string
}

func (r ApiGetFilevaultRecoveryKeyRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetFilevaultRecoveryKeyExecute(r)
}

/*
GetFilevaultRecoveryKey Get FileVault Recovery Key

<p>This request allows you to retrieve the FileVault Recovery key for a macOS device.</p>
<h3 id=&quot;request-parameters&quot;>Request Parameters</h3>
<p><code>device_id</code> (path parameter): The unique identifier of the device.</p>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceId
 @return ApiGetFilevaultRecoveryKeyRequest
*/
func (a *DeviceSecretsAPIService) GetFilevaultRecoveryKey(ctx context.Context, deviceId string) ApiGetFilevaultRecoveryKeyRequest {
	return ApiGetFilevaultRecoveryKeyRequest{
		ApiService: a,
		ctx: ctx,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DeviceSecretsAPIService) GetFilevaultRecoveryKeyExecute(r ApiGetFilevaultRecoveryKeyRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceSecretsAPIService.GetFilevaultRecoveryKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/devices/{device_id}/secrets/filevaultkey"
	localVarPath = strings.Replace(localVarPath, "{"+"device_id"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

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

type ApiGetRecoveryLockPasswordRequest struct {
	ctx context.Context
	ApiService *DeviceSecretsAPIService
	deviceId string
}

func (r ApiGetRecoveryLockPasswordRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetRecoveryLockPasswordExecute(r)
}

/*
GetRecoveryLockPassword Get Recovery Lock Password

<p>This request returns the Recovery Lock password for a Mac with an Apple Silicon processor and the legacy EFI firmware password for a Mac with an Intel processor.</p>
<p>For more details on setting and managing Recovery passwords, see this <a href=&quot;https://support.kandji.io/support/solutions/articles/72000560472-configure-the-recovery-password-library-item&quot;>Kandji support article</a>.</p>
<h3 id=&quot;request-parameters&quot;>Request Parameters</h3>
<p><code>device_id</code> (path parameter): The unique identifier of the device.</p>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceId
 @return ApiGetRecoveryLockPasswordRequest
*/
func (a *DeviceSecretsAPIService) GetRecoveryLockPassword(ctx context.Context, deviceId string) ApiGetRecoveryLockPasswordRequest {
	return ApiGetRecoveryLockPasswordRequest{
		ApiService: a,
		ctx: ctx,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DeviceSecretsAPIService) GetRecoveryLockPasswordExecute(r ApiGetRecoveryLockPasswordRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceSecretsAPIService.GetRecoveryLockPassword")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/devices/{device_id}/secrets/recoverypassword"
	localVarPath = strings.Replace(localVarPath, "{"+"device_id"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

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

type ApiGetUnlockPinRequest struct {
	ctx context.Context
	ApiService *DeviceSecretsAPIService
	deviceId string
}

func (r ApiGetUnlockPinRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetUnlockPinExecute(r)
}

/*
GetUnlockPin Get Unlock Pin

<p>This request allows you to retrieve the device unlock pin for a macOS device.</p>
<h3 id=&quot;request-parameters&quot;>Request Parameters</h3>
<p><code>device_id</code> (path parameter): The unique identifier of the device.</p>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param deviceId
 @return ApiGetUnlockPinRequest
*/
func (a *DeviceSecretsAPIService) GetUnlockPin(ctx context.Context, deviceId string) ApiGetUnlockPinRequest {
	return ApiGetUnlockPinRequest{
		ApiService: a,
		ctx: ctx,
		deviceId: deviceId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DeviceSecretsAPIService) GetUnlockPinExecute(r ApiGetUnlockPinRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceSecretsAPIService.GetUnlockPin")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/devices/{device_id}/secrets/unlockpin"
	localVarPath = strings.Replace(localVarPath, "{"+"device_id"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

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
