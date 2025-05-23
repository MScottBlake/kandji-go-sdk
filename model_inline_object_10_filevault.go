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

// checks if the InlineObject10Filevault type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject10Filevault{}

// InlineObject10Filevault struct for InlineObject10Filevault
type InlineObject10Filevault struct {
	FilevaultEnabled interface{} `json:"filevault_enabled,omitempty"`
	FilevaultRecoverykeyType *string `json:"filevault_recoverykey_type,omitempty"`
	FilevaultPrkEscrowed *int32 `json:"filevault_prk_escrowed,omitempty"`
	FilevaultNextRotation *string `json:"filevault_next_rotation,omitempty"`
	FilevaultRegenRequired *int32 `json:"filevault_regen_required,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineObject10Filevault InlineObject10Filevault

// NewInlineObject10Filevault instantiates a new InlineObject10Filevault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject10Filevault() *InlineObject10Filevault {
	this := InlineObject10Filevault{}
	return &this
}

// NewInlineObject10FilevaultWithDefaults instantiates a new InlineObject10Filevault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject10FilevaultWithDefaults() *InlineObject10Filevault {
	this := InlineObject10Filevault{}
	return &this
}

// GetFilevaultEnabled returns the FilevaultEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject10Filevault) GetFilevaultEnabled() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FilevaultEnabled
}

// GetFilevaultEnabledOk returns a tuple with the FilevaultEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject10Filevault) GetFilevaultEnabledOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FilevaultEnabled) {
		return nil, false
	}
	return &o.FilevaultEnabled, true
}

// HasFilevaultEnabled returns a boolean if a field has been set.
func (o *InlineObject10Filevault) HasFilevaultEnabled() bool {
	if o != nil && !IsNil(o.FilevaultEnabled) {
		return true
	}

	return false
}

// SetFilevaultEnabled gets a reference to the given interface{} and assigns it to the FilevaultEnabled field.
func (o *InlineObject10Filevault) SetFilevaultEnabled(v interface{}) {
	o.FilevaultEnabled = v
}

// GetFilevaultRecoverykeyType returns the FilevaultRecoverykeyType field value if set, zero value otherwise.
func (o *InlineObject10Filevault) GetFilevaultRecoverykeyType() string {
	if o == nil || IsNil(o.FilevaultRecoverykeyType) {
		var ret string
		return ret
	}
	return *o.FilevaultRecoverykeyType
}

// GetFilevaultRecoverykeyTypeOk returns a tuple with the FilevaultRecoverykeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10Filevault) GetFilevaultRecoverykeyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FilevaultRecoverykeyType) {
		return nil, false
	}
	return o.FilevaultRecoverykeyType, true
}

// HasFilevaultRecoverykeyType returns a boolean if a field has been set.
func (o *InlineObject10Filevault) HasFilevaultRecoverykeyType() bool {
	if o != nil && !IsNil(o.FilevaultRecoverykeyType) {
		return true
	}

	return false
}

// SetFilevaultRecoverykeyType gets a reference to the given string and assigns it to the FilevaultRecoverykeyType field.
func (o *InlineObject10Filevault) SetFilevaultRecoverykeyType(v string) {
	o.FilevaultRecoverykeyType = &v
}

// GetFilevaultPrkEscrowed returns the FilevaultPrkEscrowed field value if set, zero value otherwise.
func (o *InlineObject10Filevault) GetFilevaultPrkEscrowed() int32 {
	if o == nil || IsNil(o.FilevaultPrkEscrowed) {
		var ret int32
		return ret
	}
	return *o.FilevaultPrkEscrowed
}

// GetFilevaultPrkEscrowedOk returns a tuple with the FilevaultPrkEscrowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10Filevault) GetFilevaultPrkEscrowedOk() (*int32, bool) {
	if o == nil || IsNil(o.FilevaultPrkEscrowed) {
		return nil, false
	}
	return o.FilevaultPrkEscrowed, true
}

// HasFilevaultPrkEscrowed returns a boolean if a field has been set.
func (o *InlineObject10Filevault) HasFilevaultPrkEscrowed() bool {
	if o != nil && !IsNil(o.FilevaultPrkEscrowed) {
		return true
	}

	return false
}

// SetFilevaultPrkEscrowed gets a reference to the given int32 and assigns it to the FilevaultPrkEscrowed field.
func (o *InlineObject10Filevault) SetFilevaultPrkEscrowed(v int32) {
	o.FilevaultPrkEscrowed = &v
}

// GetFilevaultNextRotation returns the FilevaultNextRotation field value if set, zero value otherwise.
func (o *InlineObject10Filevault) GetFilevaultNextRotation() string {
	if o == nil || IsNil(o.FilevaultNextRotation) {
		var ret string
		return ret
	}
	return *o.FilevaultNextRotation
}

// GetFilevaultNextRotationOk returns a tuple with the FilevaultNextRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10Filevault) GetFilevaultNextRotationOk() (*string, bool) {
	if o == nil || IsNil(o.FilevaultNextRotation) {
		return nil, false
	}
	return o.FilevaultNextRotation, true
}

// HasFilevaultNextRotation returns a boolean if a field has been set.
func (o *InlineObject10Filevault) HasFilevaultNextRotation() bool {
	if o != nil && !IsNil(o.FilevaultNextRotation) {
		return true
	}

	return false
}

// SetFilevaultNextRotation gets a reference to the given string and assigns it to the FilevaultNextRotation field.
func (o *InlineObject10Filevault) SetFilevaultNextRotation(v string) {
	o.FilevaultNextRotation = &v
}

// GetFilevaultRegenRequired returns the FilevaultRegenRequired field value if set, zero value otherwise.
func (o *InlineObject10Filevault) GetFilevaultRegenRequired() int32 {
	if o == nil || IsNil(o.FilevaultRegenRequired) {
		var ret int32
		return ret
	}
	return *o.FilevaultRegenRequired
}

// GetFilevaultRegenRequiredOk returns a tuple with the FilevaultRegenRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10Filevault) GetFilevaultRegenRequiredOk() (*int32, bool) {
	if o == nil || IsNil(o.FilevaultRegenRequired) {
		return nil, false
	}
	return o.FilevaultRegenRequired, true
}

// HasFilevaultRegenRequired returns a boolean if a field has been set.
func (o *InlineObject10Filevault) HasFilevaultRegenRequired() bool {
	if o != nil && !IsNil(o.FilevaultRegenRequired) {
		return true
	}

	return false
}

// SetFilevaultRegenRequired gets a reference to the given int32 and assigns it to the FilevaultRegenRequired field.
func (o *InlineObject10Filevault) SetFilevaultRegenRequired(v int32) {
	o.FilevaultRegenRequired = &v
}

func (o InlineObject10Filevault) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject10Filevault) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FilevaultEnabled != nil {
		toSerialize["filevault_enabled"] = o.FilevaultEnabled
	}
	if !IsNil(o.FilevaultRecoverykeyType) {
		toSerialize["filevault_recoverykey_type"] = o.FilevaultRecoverykeyType
	}
	if !IsNil(o.FilevaultPrkEscrowed) {
		toSerialize["filevault_prk_escrowed"] = o.FilevaultPrkEscrowed
	}
	if !IsNil(o.FilevaultNextRotation) {
		toSerialize["filevault_next_rotation"] = o.FilevaultNextRotation
	}
	if !IsNil(o.FilevaultRegenRequired) {
		toSerialize["filevault_regen_required"] = o.FilevaultRegenRequired
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InlineObject10Filevault) UnmarshalJSON(data []byte) (err error) {
	varInlineObject10Filevault := _InlineObject10Filevault{}

	err = json.Unmarshal(data, &varInlineObject10Filevault)

	if err != nil {
		return err
	}

	*o = InlineObject10Filevault(varInlineObject10Filevault)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filevault_enabled")
		delete(additionalProperties, "filevault_recoverykey_type")
		delete(additionalProperties, "filevault_prk_escrowed")
		delete(additionalProperties, "filevault_next_rotation")
		delete(additionalProperties, "filevault_regen_required")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineObject10Filevault struct {
	value *InlineObject10Filevault
	isSet bool
}

func (v NullableInlineObject10Filevault) Get() *InlineObject10Filevault {
	return v.value
}

func (v *NullableInlineObject10Filevault) Set(val *InlineObject10Filevault) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject10Filevault) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject10Filevault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject10Filevault(val *InlineObject10Filevault) *NullableInlineObject10Filevault {
	return &NullableInlineObject10Filevault{value: val, isSet: true}
}

func (v NullableInlineObject10Filevault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject10Filevault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


