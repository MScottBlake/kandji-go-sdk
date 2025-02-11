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

// checks if the PrismLocalUsers200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrismLocalUsers200Response{}

// PrismLocalUsers200Response struct for PrismLocalUsers200Response
type PrismLocalUsers200Response struct {
	Offset *int32 `json:"offset,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Total *int32 `json:"total,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Cursor *string `json:"cursor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrismLocalUsers200Response PrismLocalUsers200Response

// NewPrismLocalUsers200Response instantiates a new PrismLocalUsers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrismLocalUsers200Response() *PrismLocalUsers200Response {
	this := PrismLocalUsers200Response{}
	return &this
}

// NewPrismLocalUsers200ResponseWithDefaults instantiates a new PrismLocalUsers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrismLocalUsers200ResponseWithDefaults() *PrismLocalUsers200Response {
	this := PrismLocalUsers200Response{}
	return &this
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *PrismLocalUsers200Response) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrismLocalUsers200Response) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *PrismLocalUsers200Response) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *PrismLocalUsers200Response) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PrismLocalUsers200Response) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrismLocalUsers200Response) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PrismLocalUsers200Response) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PrismLocalUsers200Response) SetLimit(v int32) {
	o.Limit = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *PrismLocalUsers200Response) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrismLocalUsers200Response) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *PrismLocalUsers200Response) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *PrismLocalUsers200Response) SetTotal(v int32) {
	o.Total = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrismLocalUsers200Response) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrismLocalUsers200Response) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PrismLocalUsers200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *PrismLocalUsers200Response) SetData(v interface{}) {
	o.Data = v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *PrismLocalUsers200Response) GetCursor() string {
	if o == nil || IsNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrismLocalUsers200Response) GetCursorOk() (*string, bool) {
	if o == nil || IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *PrismLocalUsers200Response) HasCursor() bool {
	if o != nil && !IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *PrismLocalUsers200Response) SetCursor(v string) {
	o.Cursor = &v
}

func (o PrismLocalUsers200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrismLocalUsers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrismLocalUsers200Response) UnmarshalJSON(data []byte) (err error) {
	varPrismLocalUsers200Response := _PrismLocalUsers200Response{}

	err = json.Unmarshal(data, &varPrismLocalUsers200Response)

	if err != nil {
		return err
	}

	*o = PrismLocalUsers200Response(varPrismLocalUsers200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "offset")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "total")
		delete(additionalProperties, "data")
		delete(additionalProperties, "cursor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrismLocalUsers200Response struct {
	value *PrismLocalUsers200Response
	isSet bool
}

func (v NullablePrismLocalUsers200Response) Get() *PrismLocalUsers200Response {
	return v.value
}

func (v *NullablePrismLocalUsers200Response) Set(val *PrismLocalUsers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePrismLocalUsers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePrismLocalUsers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrismLocalUsers200Response(val *PrismLocalUsers200Response) *NullablePrismLocalUsers200Response {
	return &NullablePrismLocalUsers200Response{value: val, isSet: true}
}

func (v NullablePrismLocalUsers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrismLocalUsers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


