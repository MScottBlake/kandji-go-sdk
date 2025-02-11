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

// checks if the PrismRequestCategoryExport400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrismRequestCategoryExport400Response{}

// PrismRequestCategoryExport400Response struct for PrismRequestCategoryExport400Response
type PrismRequestCategoryExport400Response struct {
	Detail *string `json:"detail,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrismRequestCategoryExport400Response PrismRequestCategoryExport400Response

// NewPrismRequestCategoryExport400Response instantiates a new PrismRequestCategoryExport400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrismRequestCategoryExport400Response() *PrismRequestCategoryExport400Response {
	this := PrismRequestCategoryExport400Response{}
	return &this
}

// NewPrismRequestCategoryExport400ResponseWithDefaults instantiates a new PrismRequestCategoryExport400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrismRequestCategoryExport400ResponseWithDefaults() *PrismRequestCategoryExport400Response {
	this := PrismRequestCategoryExport400Response{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *PrismRequestCategoryExport400Response) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrismRequestCategoryExport400Response) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *PrismRequestCategoryExport400Response) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *PrismRequestCategoryExport400Response) SetDetail(v string) {
	o.Detail = &v
}

func (o PrismRequestCategoryExport400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrismRequestCategoryExport400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrismRequestCategoryExport400Response) UnmarshalJSON(data []byte) (err error) {
	varPrismRequestCategoryExport400Response := _PrismRequestCategoryExport400Response{}

	err = json.Unmarshal(data, &varPrismRequestCategoryExport400Response)

	if err != nil {
		return err
	}

	*o = PrismRequestCategoryExport400Response(varPrismRequestCategoryExport400Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "detail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrismRequestCategoryExport400Response struct {
	value *PrismRequestCategoryExport400Response
	isSet bool
}

func (v NullablePrismRequestCategoryExport400Response) Get() *PrismRequestCategoryExport400Response {
	return v.value
}

func (v *NullablePrismRequestCategoryExport400Response) Set(val *PrismRequestCategoryExport400Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePrismRequestCategoryExport400Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePrismRequestCategoryExport400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrismRequestCategoryExport400Response(val *PrismRequestCategoryExport400Response) *NullablePrismRequestCategoryExport400Response {
	return &NullablePrismRequestCategoryExport400Response{value: val, isSet: true}
}

func (v NullablePrismRequestCategoryExport400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrismRequestCategoryExport400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


