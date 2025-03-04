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

// checks if the DeviceActionsGetDeviceCommands200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceActionsGetDeviceCommands200Response{}

// DeviceActionsGetDeviceCommands200Response struct for DeviceActionsGetDeviceCommands200Response
type DeviceActionsGetDeviceCommands200Response struct {
	Commands *BlueprintsListBlueprints200Response `json:"commands,omitempty"`
	DeviceId *string `json:"device_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceActionsGetDeviceCommands200Response DeviceActionsGetDeviceCommands200Response

// NewDeviceActionsGetDeviceCommands200Response instantiates a new DeviceActionsGetDeviceCommands200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceActionsGetDeviceCommands200Response() *DeviceActionsGetDeviceCommands200Response {
	this := DeviceActionsGetDeviceCommands200Response{}
	return &this
}

// NewDeviceActionsGetDeviceCommands200ResponseWithDefaults instantiates a new DeviceActionsGetDeviceCommands200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceActionsGetDeviceCommands200ResponseWithDefaults() *DeviceActionsGetDeviceCommands200Response {
	this := DeviceActionsGetDeviceCommands200Response{}
	return &this
}

// GetCommands returns the Commands field value if set, zero value otherwise.
func (o *DeviceActionsGetDeviceCommands200Response) GetCommands() BlueprintsListBlueprints200Response {
	if o == nil || IsNil(o.Commands) {
		var ret BlueprintsListBlueprints200Response
		return ret
	}
	return *o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceActionsGetDeviceCommands200Response) GetCommandsOk() (*BlueprintsListBlueprints200Response, bool) {
	if o == nil || IsNil(o.Commands) {
		return nil, false
	}
	return o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *DeviceActionsGetDeviceCommands200Response) HasCommands() bool {
	if o != nil && !IsNil(o.Commands) {
		return true
	}

	return false
}

// SetCommands gets a reference to the given BlueprintsListBlueprints200Response and assigns it to the Commands field.
func (o *DeviceActionsGetDeviceCommands200Response) SetCommands(v BlueprintsListBlueprints200Response) {
	o.Commands = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *DeviceActionsGetDeviceCommands200Response) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceActionsGetDeviceCommands200Response) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *DeviceActionsGetDeviceCommands200Response) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *DeviceActionsGetDeviceCommands200Response) SetDeviceId(v string) {
	o.DeviceId = &v
}

func (o DeviceActionsGetDeviceCommands200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceActionsGetDeviceCommands200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Commands) {
		toSerialize["commands"] = o.Commands
	}
	if !IsNil(o.DeviceId) {
		toSerialize["device_id"] = o.DeviceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceActionsGetDeviceCommands200Response) UnmarshalJSON(data []byte) (err error) {
	varDeviceActionsGetDeviceCommands200Response := _DeviceActionsGetDeviceCommands200Response{}

	err = json.Unmarshal(data, &varDeviceActionsGetDeviceCommands200Response)

	if err != nil {
		return err
	}

	*o = DeviceActionsGetDeviceCommands200Response(varDeviceActionsGetDeviceCommands200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "commands")
		delete(additionalProperties, "device_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceActionsGetDeviceCommands200Response struct {
	value *DeviceActionsGetDeviceCommands200Response
	isSet bool
}

func (v NullableDeviceActionsGetDeviceCommands200Response) Get() *DeviceActionsGetDeviceCommands200Response {
	return v.value
}

func (v *NullableDeviceActionsGetDeviceCommands200Response) Set(val *DeviceActionsGetDeviceCommands200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceActionsGetDeviceCommands200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceActionsGetDeviceCommands200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceActionsGetDeviceCommands200Response(val *DeviceActionsGetDeviceCommands200Response) *NullableDeviceActionsGetDeviceCommands200Response {
	return &NullableDeviceActionsGetDeviceCommands200Response{value: val, isSet: true}
}

func (v NullableDeviceActionsGetDeviceCommands200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceActionsGetDeviceCommands200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


