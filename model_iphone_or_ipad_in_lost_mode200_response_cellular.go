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

// checks if the IphoneOrIpadInLostMode200ResponseCellular type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IphoneOrIpadInLostMode200ResponseCellular{}

// IphoneOrIpadInLostMode200ResponseCellular struct for IphoneOrIpadInLostMode200ResponseCellular
type IphoneOrIpadInLostMode200ResponseCellular struct {
	VoiceRoaming *int32 `json:"voice_roaming,omitempty"`
	DataRoaming *int32 `json:"data_roaming,omitempty"`
	CellularTechnology *int32 `json:"cellular_technology,omitempty"`
	Subscriptions interface{} `json:"subscriptions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IphoneOrIpadInLostMode200ResponseCellular IphoneOrIpadInLostMode200ResponseCellular

// NewIphoneOrIpadInLostMode200ResponseCellular instantiates a new IphoneOrIpadInLostMode200ResponseCellular object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIphoneOrIpadInLostMode200ResponseCellular() *IphoneOrIpadInLostMode200ResponseCellular {
	this := IphoneOrIpadInLostMode200ResponseCellular{}
	return &this
}

// NewIphoneOrIpadInLostMode200ResponseCellularWithDefaults instantiates a new IphoneOrIpadInLostMode200ResponseCellular object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIphoneOrIpadInLostMode200ResponseCellularWithDefaults() *IphoneOrIpadInLostMode200ResponseCellular {
	this := IphoneOrIpadInLostMode200ResponseCellular{}
	return &this
}

// GetVoiceRoaming returns the VoiceRoaming field value if set, zero value otherwise.
func (o *IphoneOrIpadInLostMode200ResponseCellular) GetVoiceRoaming() int32 {
	if o == nil || IsNil(o.VoiceRoaming) {
		var ret int32
		return ret
	}
	return *o.VoiceRoaming
}

// GetVoiceRoamingOk returns a tuple with the VoiceRoaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IphoneOrIpadInLostMode200ResponseCellular) GetVoiceRoamingOk() (*int32, bool) {
	if o == nil || IsNil(o.VoiceRoaming) {
		return nil, false
	}
	return o.VoiceRoaming, true
}

// HasVoiceRoaming returns a boolean if a field has been set.
func (o *IphoneOrIpadInLostMode200ResponseCellular) HasVoiceRoaming() bool {
	if o != nil && !IsNil(o.VoiceRoaming) {
		return true
	}

	return false
}

// SetVoiceRoaming gets a reference to the given int32 and assigns it to the VoiceRoaming field.
func (o *IphoneOrIpadInLostMode200ResponseCellular) SetVoiceRoaming(v int32) {
	o.VoiceRoaming = &v
}

// GetDataRoaming returns the DataRoaming field value if set, zero value otherwise.
func (o *IphoneOrIpadInLostMode200ResponseCellular) GetDataRoaming() int32 {
	if o == nil || IsNil(o.DataRoaming) {
		var ret int32
		return ret
	}
	return *o.DataRoaming
}

// GetDataRoamingOk returns a tuple with the DataRoaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IphoneOrIpadInLostMode200ResponseCellular) GetDataRoamingOk() (*int32, bool) {
	if o == nil || IsNil(o.DataRoaming) {
		return nil, false
	}
	return o.DataRoaming, true
}

// HasDataRoaming returns a boolean if a field has been set.
func (o *IphoneOrIpadInLostMode200ResponseCellular) HasDataRoaming() bool {
	if o != nil && !IsNil(o.DataRoaming) {
		return true
	}

	return false
}

// SetDataRoaming gets a reference to the given int32 and assigns it to the DataRoaming field.
func (o *IphoneOrIpadInLostMode200ResponseCellular) SetDataRoaming(v int32) {
	o.DataRoaming = &v
}

// GetCellularTechnology returns the CellularTechnology field value if set, zero value otherwise.
func (o *IphoneOrIpadInLostMode200ResponseCellular) GetCellularTechnology() int32 {
	if o == nil || IsNil(o.CellularTechnology) {
		var ret int32
		return ret
	}
	return *o.CellularTechnology
}

// GetCellularTechnologyOk returns a tuple with the CellularTechnology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IphoneOrIpadInLostMode200ResponseCellular) GetCellularTechnologyOk() (*int32, bool) {
	if o == nil || IsNil(o.CellularTechnology) {
		return nil, false
	}
	return o.CellularTechnology, true
}

// HasCellularTechnology returns a boolean if a field has been set.
func (o *IphoneOrIpadInLostMode200ResponseCellular) HasCellularTechnology() bool {
	if o != nil && !IsNil(o.CellularTechnology) {
		return true
	}

	return false
}

// SetCellularTechnology gets a reference to the given int32 and assigns it to the CellularTechnology field.
func (o *IphoneOrIpadInLostMode200ResponseCellular) SetCellularTechnology(v int32) {
	o.CellularTechnology = &v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IphoneOrIpadInLostMode200ResponseCellular) GetSubscriptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IphoneOrIpadInLostMode200ResponseCellular) GetSubscriptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Subscriptions) {
		return nil, false
	}
	return &o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *IphoneOrIpadInLostMode200ResponseCellular) HasSubscriptions() bool {
	if o != nil && !IsNil(o.Subscriptions) {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given interface{} and assigns it to the Subscriptions field.
func (o *IphoneOrIpadInLostMode200ResponseCellular) SetSubscriptions(v interface{}) {
	o.Subscriptions = v
}

func (o IphoneOrIpadInLostMode200ResponseCellular) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IphoneOrIpadInLostMode200ResponseCellular) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VoiceRoaming) {
		toSerialize["voice_roaming"] = o.VoiceRoaming
	}
	if !IsNil(o.DataRoaming) {
		toSerialize["data_roaming"] = o.DataRoaming
	}
	if !IsNil(o.CellularTechnology) {
		toSerialize["cellular_technology"] = o.CellularTechnology
	}
	if o.Subscriptions != nil {
		toSerialize["subscriptions"] = o.Subscriptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IphoneOrIpadInLostMode200ResponseCellular) UnmarshalJSON(data []byte) (err error) {
	varIphoneOrIpadInLostMode200ResponseCellular := _IphoneOrIpadInLostMode200ResponseCellular{}

	err = json.Unmarshal(data, &varIphoneOrIpadInLostMode200ResponseCellular)

	if err != nil {
		return err
	}

	*o = IphoneOrIpadInLostMode200ResponseCellular(varIphoneOrIpadInLostMode200ResponseCellular)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "voice_roaming")
		delete(additionalProperties, "data_roaming")
		delete(additionalProperties, "cellular_technology")
		delete(additionalProperties, "subscriptions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIphoneOrIpadInLostMode200ResponseCellular struct {
	value *IphoneOrIpadInLostMode200ResponseCellular
	isSet bool
}

func (v NullableIphoneOrIpadInLostMode200ResponseCellular) Get() *IphoneOrIpadInLostMode200ResponseCellular {
	return v.value
}

func (v *NullableIphoneOrIpadInLostMode200ResponseCellular) Set(val *IphoneOrIpadInLostMode200ResponseCellular) {
	v.value = val
	v.isSet = true
}

func (v NullableIphoneOrIpadInLostMode200ResponseCellular) IsSet() bool {
	return v.isSet
}

func (v *NullableIphoneOrIpadInLostMode200ResponseCellular) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIphoneOrIpadInLostMode200ResponseCellular(val *IphoneOrIpadInLostMode200ResponseCellular) *NullableIphoneOrIpadInLostMode200ResponseCellular {
	return &NullableIphoneOrIpadInLostMode200ResponseCellular{value: val, isSet: true}
}

func (v NullableIphoneOrIpadInLostMode200ResponseCellular) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIphoneOrIpadInLostMode200ResponseCellular) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


