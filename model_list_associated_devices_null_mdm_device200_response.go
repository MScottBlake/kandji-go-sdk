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

// checks if the ListAssociatedDevicesNullMdmDevice200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAssociatedDevicesNullMdmDevice200Response{}

// ListAssociatedDevicesNullMdmDevice200Response struct for ListAssociatedDevicesNullMdmDevice200Response
type ListAssociatedDevicesNullMdmDevice200Response struct {
	Count *int32 `json:"count,omitempty"`
	Next *string `json:"next,omitempty"`
	Previous interface{} `json:"previous,omitempty"`
	Results interface{} `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListAssociatedDevicesNullMdmDevice200Response ListAssociatedDevicesNullMdmDevice200Response

// NewListAssociatedDevicesNullMdmDevice200Response instantiates a new ListAssociatedDevicesNullMdmDevice200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAssociatedDevicesNullMdmDevice200Response() *ListAssociatedDevicesNullMdmDevice200Response {
	this := ListAssociatedDevicesNullMdmDevice200Response{}
	return &this
}

// NewListAssociatedDevicesNullMdmDevice200ResponseWithDefaults instantiates a new ListAssociatedDevicesNullMdmDevice200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAssociatedDevicesNullMdmDevice200ResponseWithDefaults() *ListAssociatedDevicesNullMdmDevice200Response {
	this := ListAssociatedDevicesNullMdmDevice200Response{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListAssociatedDevicesNullMdmDevice200Response) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAssociatedDevicesNullMdmDevice200Response) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListAssociatedDevicesNullMdmDevice200Response) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ListAssociatedDevicesNullMdmDevice200Response) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListAssociatedDevicesNullMdmDevice200Response) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAssociatedDevicesNullMdmDevice200Response) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListAssociatedDevicesNullMdmDevice200Response) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *ListAssociatedDevicesNullMdmDevice200Response) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListAssociatedDevicesNullMdmDevice200Response) GetPrevious() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListAssociatedDevicesNullMdmDevice200Response) GetPreviousOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return &o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *ListAssociatedDevicesNullMdmDevice200Response) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given interface{} and assigns it to the Previous field.
func (o *ListAssociatedDevicesNullMdmDevice200Response) SetPrevious(v interface{}) {
	o.Previous = v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListAssociatedDevicesNullMdmDevice200Response) GetResults() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListAssociatedDevicesNullMdmDevice200Response) GetResultsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return &o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListAssociatedDevicesNullMdmDevice200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given interface{} and assigns it to the Results field.
func (o *ListAssociatedDevicesNullMdmDevice200Response) SetResults(v interface{}) {
	o.Results = v
}

func (o ListAssociatedDevicesNullMdmDevice200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAssociatedDevicesNullMdmDevice200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListAssociatedDevicesNullMdmDevice200Response) UnmarshalJSON(data []byte) (err error) {
	varListAssociatedDevicesNullMdmDevice200Response := _ListAssociatedDevicesNullMdmDevice200Response{}

	err = json.Unmarshal(data, &varListAssociatedDevicesNullMdmDevice200Response)

	if err != nil {
		return err
	}

	*o = ListAssociatedDevicesNullMdmDevice200Response(varListAssociatedDevicesNullMdmDevice200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "next")
		delete(additionalProperties, "previous")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListAssociatedDevicesNullMdmDevice200Response struct {
	value *ListAssociatedDevicesNullMdmDevice200Response
	isSet bool
}

func (v NullableListAssociatedDevicesNullMdmDevice200Response) Get() *ListAssociatedDevicesNullMdmDevice200Response {
	return v.value
}

func (v *NullableListAssociatedDevicesNullMdmDevice200Response) Set(val *ListAssociatedDevicesNullMdmDevice200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListAssociatedDevicesNullMdmDevice200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListAssociatedDevicesNullMdmDevice200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAssociatedDevicesNullMdmDevice200Response(val *ListAssociatedDevicesNullMdmDevice200Response) *NullableListAssociatedDevicesNullMdmDevice200Response {
	return &NullableListAssociatedDevicesNullMdmDevice200Response{value: val, isSet: true}
}

func (v NullableListAssociatedDevicesNullMdmDevice200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAssociatedDevicesNullMdmDevice200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


