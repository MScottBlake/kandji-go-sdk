/*
Kandji API

<html><head></head><body><h1 id=&quot;welcome-to-the-kandji-api-documentation&quot;>Welcome to the Kandji API Documentation</h1> <p>You can find your API URL in Settings &gt; Access. The API URL will follow the below formats.</p> <ul> <li><p>US - <code>https://SubDomain.api.kandji.io</code></p> </li> <li><p>EU - <code>https://SubDomain.api.eu.kandji.io</code></p> </li> </ul> <p>For information on how to obtain an API token, please refer to the following support article.</p> <p><a href=&quot;https://support.kandji.io/api&quot;>https://support.kandji.io/api</a></p> <h4 id=&quot;rate-limit&quot;>Rate Limit</h4> <p>The Kandji API currently has an API rate limit of 10,000 requests per hour per customer.</p> <h4 id=&quot;request-methods&quot;>Request Methods</h4> <p>HTTP request methods supported by the Kandji API.</p> <div class=&quot;click-to-expand-wrapper is-table-wrapper&quot;><table> <thead> <tr> <th>Method</th> <th>Definition</th> </tr> </thead> <tbody> <tr> <td>GET</td> <td>The <code>GET</code> method requests a representation of the specified resource.</td> </tr> <tr> <td>POST</td> <td>The <code>POST</code> method submits an entity to the specified resource.</td> </tr> <tr> <td>PATCH</td> <td>The <code>PATCH</code> method applies partial modifications to a resource.</td> </tr> <tr> <td>DELETE</td> <td>The <code>DELETE</code> method deletes the specified resource.</td> </tr> </tbody> </table> </div><h4 id=&quot;response-codes&quot;>Response codes</h4> <p>Not all response codes apply to every endpoint.</p> <div class=&quot;click-to-expand-wrapper is-table-wrapper&quot;><table> <thead> <tr> <th>Code</th> <th>Response</th> </tr> </thead> <tbody> <tr> <td>200</td> <td>OK</td> </tr> <tr> <td>201</td> <td>Created</td> </tr> <tr> <td>204</td> <td>No content</td> </tr> <tr> <td></td> <td>Typical response when sending the DELETE method.</td> </tr> <tr> <td>400</td> <td>Bad Request</td> </tr> <tr> <td></td> <td>&quot;Command already running&quot; - The command may already be running in a <em>Pending</em> state waiting on the device.</td> </tr> <tr> <td></td> <td>&quot;Command is not allowed for current device&quot; - The command may not be compatible with the target device.</td> </tr> <tr> <td></td> <td>&quot;JSON parse error - Expecting ',' delimiter: line 3 column 2 (char 65)&quot;</td> </tr> <tr> <td>401</td> <td>Unauthorized</td> </tr> <tr> <td></td> <td>This error can occur if the token is incorrect, was revoked, or the token has expired.</td> </tr> <tr> <td>403</td> <td>Forbidden</td> </tr> <tr> <td></td> <td>The request was understood but cannot be authorized.</td> </tr> <tr> <td>404</td> <td>Not found</td> </tr> <tr> <td></td> <td>Unable to locate the resource in the Kandji tenant.</td> </tr> <tr> <td>415</td> <td>Unsupported Media Type</td> </tr> <tr> <td></td> <td>The request contains a media type which the server or resource does not support.</td> </tr> <tr> <td>500</td> <td>Internal server error</td> </tr> <tr> <td>503</td> <td>Service unavailable</td> </tr> <tr> <td></td> <td>This error can occur if a file upload is still being processed via the custom apps API.</td> </tr> </tbody> </table> </div><h4 id=&quot;data-structure&quot;>Data structure</h4> <p>The API returns all structured responses in JSON schema format.</p> <h4 id=&quot;examples&quot;>Examples</h4> <p>Code examples using the API can be found in the Kandji support <a href=&quot;https://github.com/kandji-inc/support/tree/main/api-tools&quot;>GitHub</a>.</p> </body></html>

API version: 1.0.0
Contact: mitchelsblake@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kandji

import (
	"encoding/json"
)

// checks if the ThreatsGetBehavioralDetections200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreatsGetBehavioralDetections200Response{}

// ThreatsGetBehavioralDetections200Response struct for ThreatsGetBehavioralDetections200Response
type ThreatsGetBehavioralDetections200Response struct {
	MaliciousCount *int32 `json:"malicious_count,omitempty"`
	Next interface{} `json:"next,omitempty"`
	Previous interface{} `json:"previous,omitempty"`
	Results interface{} `json:"results,omitempty"`
	SuspiciousCount *int32 `json:"suspicious_count,omitempty"`
	TotalCount *int32 `json:"total_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ThreatsGetBehavioralDetections200Response ThreatsGetBehavioralDetections200Response

// NewThreatsGetBehavioralDetections200Response instantiates a new ThreatsGetBehavioralDetections200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreatsGetBehavioralDetections200Response() *ThreatsGetBehavioralDetections200Response {
	this := ThreatsGetBehavioralDetections200Response{}
	return &this
}

// NewThreatsGetBehavioralDetections200ResponseWithDefaults instantiates a new ThreatsGetBehavioralDetections200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreatsGetBehavioralDetections200ResponseWithDefaults() *ThreatsGetBehavioralDetections200Response {
	this := ThreatsGetBehavioralDetections200Response{}
	return &this
}

// GetMaliciousCount returns the MaliciousCount field value if set, zero value otherwise.
func (o *ThreatsGetBehavioralDetections200Response) GetMaliciousCount() int32 {
	if o == nil || IsNil(o.MaliciousCount) {
		var ret int32
		return ret
	}
	return *o.MaliciousCount
}

// GetMaliciousCountOk returns a tuple with the MaliciousCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatsGetBehavioralDetections200Response) GetMaliciousCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MaliciousCount) {
		return nil, false
	}
	return o.MaliciousCount, true
}

// HasMaliciousCount returns a boolean if a field has been set.
func (o *ThreatsGetBehavioralDetections200Response) HasMaliciousCount() bool {
	if o != nil && !IsNil(o.MaliciousCount) {
		return true
	}

	return false
}

// SetMaliciousCount gets a reference to the given int32 and assigns it to the MaliciousCount field.
func (o *ThreatsGetBehavioralDetections200Response) SetMaliciousCount(v int32) {
	o.MaliciousCount = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreatsGetBehavioralDetections200Response) GetNext() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreatsGetBehavioralDetections200Response) GetNextOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return &o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ThreatsGetBehavioralDetections200Response) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given interface{} and assigns it to the Next field.
func (o *ThreatsGetBehavioralDetections200Response) SetNext(v interface{}) {
	o.Next = v
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreatsGetBehavioralDetections200Response) GetPrevious() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreatsGetBehavioralDetections200Response) GetPreviousOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return &o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *ThreatsGetBehavioralDetections200Response) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given interface{} and assigns it to the Previous field.
func (o *ThreatsGetBehavioralDetections200Response) SetPrevious(v interface{}) {
	o.Previous = v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreatsGetBehavioralDetections200Response) GetResults() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreatsGetBehavioralDetections200Response) GetResultsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return &o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ThreatsGetBehavioralDetections200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given interface{} and assigns it to the Results field.
func (o *ThreatsGetBehavioralDetections200Response) SetResults(v interface{}) {
	o.Results = v
}

// GetSuspiciousCount returns the SuspiciousCount field value if set, zero value otherwise.
func (o *ThreatsGetBehavioralDetections200Response) GetSuspiciousCount() int32 {
	if o == nil || IsNil(o.SuspiciousCount) {
		var ret int32
		return ret
	}
	return *o.SuspiciousCount
}

// GetSuspiciousCountOk returns a tuple with the SuspiciousCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatsGetBehavioralDetections200Response) GetSuspiciousCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SuspiciousCount) {
		return nil, false
	}
	return o.SuspiciousCount, true
}

// HasSuspiciousCount returns a boolean if a field has been set.
func (o *ThreatsGetBehavioralDetections200Response) HasSuspiciousCount() bool {
	if o != nil && !IsNil(o.SuspiciousCount) {
		return true
	}

	return false
}

// SetSuspiciousCount gets a reference to the given int32 and assigns it to the SuspiciousCount field.
func (o *ThreatsGetBehavioralDetections200Response) SetSuspiciousCount(v int32) {
	o.SuspiciousCount = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ThreatsGetBehavioralDetections200Response) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatsGetBehavioralDetections200Response) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ThreatsGetBehavioralDetections200Response) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ThreatsGetBehavioralDetections200Response) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o ThreatsGetBehavioralDetections200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreatsGetBehavioralDetections200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaliciousCount) {
		toSerialize["malicious_count"] = o.MaliciousCount
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.SuspiciousCount) {
		toSerialize["suspicious_count"] = o.SuspiciousCount
	}
	if !IsNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ThreatsGetBehavioralDetections200Response) UnmarshalJSON(data []byte) (err error) {
	varThreatsGetBehavioralDetections200Response := _ThreatsGetBehavioralDetections200Response{}

	err = json.Unmarshal(data, &varThreatsGetBehavioralDetections200Response)

	if err != nil {
		return err
	}

	*o = ThreatsGetBehavioralDetections200Response(varThreatsGetBehavioralDetections200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "malicious_count")
		delete(additionalProperties, "next")
		delete(additionalProperties, "previous")
		delete(additionalProperties, "results")
		delete(additionalProperties, "suspicious_count")
		delete(additionalProperties, "total_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableThreatsGetBehavioralDetections200Response struct {
	value *ThreatsGetBehavioralDetections200Response
	isSet bool
}

func (v NullableThreatsGetBehavioralDetections200Response) Get() *ThreatsGetBehavioralDetections200Response {
	return v.value
}

func (v *NullableThreatsGetBehavioralDetections200Response) Set(val *ThreatsGetBehavioralDetections200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableThreatsGetBehavioralDetections200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableThreatsGetBehavioralDetections200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreatsGetBehavioralDetections200Response(val *ThreatsGetBehavioralDetections200Response) *NullableThreatsGetBehavioralDetections200Response {
	return &NullableThreatsGetBehavioralDetections200Response{value: val, isSet: true}
}

func (v NullableThreatsGetBehavioralDetections200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreatsGetBehavioralDetections200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


