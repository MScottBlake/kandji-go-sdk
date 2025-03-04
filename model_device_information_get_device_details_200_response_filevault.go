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

// checks if the DeviceInformationGetDeviceDetails200ResponseFilevault type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceInformationGetDeviceDetails200ResponseFilevault{}

// DeviceInformationGetDeviceDetails200ResponseFilevault struct for DeviceInformationGetDeviceDetails200ResponseFilevault
type DeviceInformationGetDeviceDetails200ResponseFilevault struct {
	FilevaultEnabled interface{} `json:"filevault_enabled,omitempty"`
	FilevaultNextRotation *string `json:"filevault_next_rotation,omitempty"`
	FilevaultPrkEscrowed *int32 `json:"filevault_prk_escrowed,omitempty"`
	FilevaultRecoverykeyType *string `json:"filevault_recoverykey_type,omitempty"`
	FilevaultRegenRequired *int32 `json:"filevault_regen_required,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceInformationGetDeviceDetails200ResponseFilevault DeviceInformationGetDeviceDetails200ResponseFilevault

// NewDeviceInformationGetDeviceDetails200ResponseFilevault instantiates a new DeviceInformationGetDeviceDetails200ResponseFilevault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceInformationGetDeviceDetails200ResponseFilevault() *DeviceInformationGetDeviceDetails200ResponseFilevault {
	this := DeviceInformationGetDeviceDetails200ResponseFilevault{}
	return &this
}

// NewDeviceInformationGetDeviceDetails200ResponseFilevaultWithDefaults instantiates a new DeviceInformationGetDeviceDetails200ResponseFilevault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceInformationGetDeviceDetails200ResponseFilevaultWithDefaults() *DeviceInformationGetDeviceDetails200ResponseFilevault {
	this := DeviceInformationGetDeviceDetails200ResponseFilevault{}
	return &this
}

// GetFilevaultEnabled returns the FilevaultEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) GetFilevaultEnabled() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FilevaultEnabled
}

// GetFilevaultEnabledOk returns a tuple with the FilevaultEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) GetFilevaultEnabledOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FilevaultEnabled) {
		return nil, false
	}
	return &o.FilevaultEnabled, true
}

// HasFilevaultEnabled returns a boolean if a field has been set.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) HasFilevaultEnabled() bool {
	if o != nil && !IsNil(o.FilevaultEnabled) {
		return true
	}

	return false
}

// SetFilevaultEnabled gets a reference to the given interface{} and assigns it to the FilevaultEnabled field.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) SetFilevaultEnabled(v interface{}) {
	o.FilevaultEnabled = v
}

// GetFilevaultNextRotation returns the FilevaultNextRotation field value if set, zero value otherwise.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) GetFilevaultNextRotation() string {
	if o == nil || IsNil(o.FilevaultNextRotation) {
		var ret string
		return ret
	}
	return *o.FilevaultNextRotation
}

// GetFilevaultNextRotationOk returns a tuple with the FilevaultNextRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) GetFilevaultNextRotationOk() (*string, bool) {
	if o == nil || IsNil(o.FilevaultNextRotation) {
		return nil, false
	}
	return o.FilevaultNextRotation, true
}

// HasFilevaultNextRotation returns a boolean if a field has been set.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) HasFilevaultNextRotation() bool {
	if o != nil && !IsNil(o.FilevaultNextRotation) {
		return true
	}

	return false
}

// SetFilevaultNextRotation gets a reference to the given string and assigns it to the FilevaultNextRotation field.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) SetFilevaultNextRotation(v string) {
	o.FilevaultNextRotation = &v
}

// GetFilevaultPrkEscrowed returns the FilevaultPrkEscrowed field value if set, zero value otherwise.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) GetFilevaultPrkEscrowed() int32 {
	if o == nil || IsNil(o.FilevaultPrkEscrowed) {
		var ret int32
		return ret
	}
	return *o.FilevaultPrkEscrowed
}

// GetFilevaultPrkEscrowedOk returns a tuple with the FilevaultPrkEscrowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) GetFilevaultPrkEscrowedOk() (*int32, bool) {
	if o == nil || IsNil(o.FilevaultPrkEscrowed) {
		return nil, false
	}
	return o.FilevaultPrkEscrowed, true
}

// HasFilevaultPrkEscrowed returns a boolean if a field has been set.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) HasFilevaultPrkEscrowed() bool {
	if o != nil && !IsNil(o.FilevaultPrkEscrowed) {
		return true
	}

	return false
}

// SetFilevaultPrkEscrowed gets a reference to the given int32 and assigns it to the FilevaultPrkEscrowed field.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) SetFilevaultPrkEscrowed(v int32) {
	o.FilevaultPrkEscrowed = &v
}

// GetFilevaultRecoverykeyType returns the FilevaultRecoverykeyType field value if set, zero value otherwise.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) GetFilevaultRecoverykeyType() string {
	if o == nil || IsNil(o.FilevaultRecoverykeyType) {
		var ret string
		return ret
	}
	return *o.FilevaultRecoverykeyType
}

// GetFilevaultRecoverykeyTypeOk returns a tuple with the FilevaultRecoverykeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) GetFilevaultRecoverykeyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FilevaultRecoverykeyType) {
		return nil, false
	}
	return o.FilevaultRecoverykeyType, true
}

// HasFilevaultRecoverykeyType returns a boolean if a field has been set.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) HasFilevaultRecoverykeyType() bool {
	if o != nil && !IsNil(o.FilevaultRecoverykeyType) {
		return true
	}

	return false
}

// SetFilevaultRecoverykeyType gets a reference to the given string and assigns it to the FilevaultRecoverykeyType field.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) SetFilevaultRecoverykeyType(v string) {
	o.FilevaultRecoverykeyType = &v
}

// GetFilevaultRegenRequired returns the FilevaultRegenRequired field value if set, zero value otherwise.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) GetFilevaultRegenRequired() int32 {
	if o == nil || IsNil(o.FilevaultRegenRequired) {
		var ret int32
		return ret
	}
	return *o.FilevaultRegenRequired
}

// GetFilevaultRegenRequiredOk returns a tuple with the FilevaultRegenRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) GetFilevaultRegenRequiredOk() (*int32, bool) {
	if o == nil || IsNil(o.FilevaultRegenRequired) {
		return nil, false
	}
	return o.FilevaultRegenRequired, true
}

// HasFilevaultRegenRequired returns a boolean if a field has been set.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) HasFilevaultRegenRequired() bool {
	if o != nil && !IsNil(o.FilevaultRegenRequired) {
		return true
	}

	return false
}

// SetFilevaultRegenRequired gets a reference to the given int32 and assigns it to the FilevaultRegenRequired field.
func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) SetFilevaultRegenRequired(v int32) {
	o.FilevaultRegenRequired = &v
}

func (o DeviceInformationGetDeviceDetails200ResponseFilevault) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceInformationGetDeviceDetails200ResponseFilevault) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FilevaultEnabled != nil {
		toSerialize["filevault_enabled"] = o.FilevaultEnabled
	}
	if !IsNil(o.FilevaultNextRotation) {
		toSerialize["filevault_next_rotation"] = o.FilevaultNextRotation
	}
	if !IsNil(o.FilevaultPrkEscrowed) {
		toSerialize["filevault_prk_escrowed"] = o.FilevaultPrkEscrowed
	}
	if !IsNil(o.FilevaultRecoverykeyType) {
		toSerialize["filevault_recoverykey_type"] = o.FilevaultRecoverykeyType
	}
	if !IsNil(o.FilevaultRegenRequired) {
		toSerialize["filevault_regen_required"] = o.FilevaultRegenRequired
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceInformationGetDeviceDetails200ResponseFilevault) UnmarshalJSON(data []byte) (err error) {
	varDeviceInformationGetDeviceDetails200ResponseFilevault := _DeviceInformationGetDeviceDetails200ResponseFilevault{}

	err = json.Unmarshal(data, &varDeviceInformationGetDeviceDetails200ResponseFilevault)

	if err != nil {
		return err
	}

	*o = DeviceInformationGetDeviceDetails200ResponseFilevault(varDeviceInformationGetDeviceDetails200ResponseFilevault)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filevault_enabled")
		delete(additionalProperties, "filevault_next_rotation")
		delete(additionalProperties, "filevault_prk_escrowed")
		delete(additionalProperties, "filevault_recoverykey_type")
		delete(additionalProperties, "filevault_regen_required")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceInformationGetDeviceDetails200ResponseFilevault struct {
	value *DeviceInformationGetDeviceDetails200ResponseFilevault
	isSet bool
}

func (v NullableDeviceInformationGetDeviceDetails200ResponseFilevault) Get() *DeviceInformationGetDeviceDetails200ResponseFilevault {
	return v.value
}

func (v *NullableDeviceInformationGetDeviceDetails200ResponseFilevault) Set(val *DeviceInformationGetDeviceDetails200ResponseFilevault) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceInformationGetDeviceDetails200ResponseFilevault) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceInformationGetDeviceDetails200ResponseFilevault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceInformationGetDeviceDetails200ResponseFilevault(val *DeviceInformationGetDeviceDetails200ResponseFilevault) *NullableDeviceInformationGetDeviceDetails200ResponseFilevault {
	return &NullableDeviceInformationGetDeviceDetails200ResponseFilevault{value: val, isSet: true}
}

func (v NullableDeviceInformationGetDeviceDetails200ResponseFilevault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceInformationGetDeviceDetails200ResponseFilevault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


