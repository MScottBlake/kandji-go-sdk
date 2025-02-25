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

// checks if the InlineObject10Mdm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject10Mdm{}

// InlineObject10Mdm struct for InlineObject10Mdm
type InlineObject10Mdm struct {
	MdmEnabled *string `json:"mdm_enabled,omitempty"`
	Supervised *string `json:"supervised,omitempty"`
	InstallDate *string `json:"install_date,omitempty"`
	LastCheckIn *string `json:"last_check_in,omitempty"`
	MdmEnabledUser interface{} `json:"mdm_enabled_user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineObject10Mdm InlineObject10Mdm

// NewInlineObject10Mdm instantiates a new InlineObject10Mdm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject10Mdm() *InlineObject10Mdm {
	this := InlineObject10Mdm{}
	return &this
}

// NewInlineObject10MdmWithDefaults instantiates a new InlineObject10Mdm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject10MdmWithDefaults() *InlineObject10Mdm {
	this := InlineObject10Mdm{}
	return &this
}

// GetMdmEnabled returns the MdmEnabled field value if set, zero value otherwise.
func (o *InlineObject10Mdm) GetMdmEnabled() string {
	if o == nil || IsNil(o.MdmEnabled) {
		var ret string
		return ret
	}
	return *o.MdmEnabled
}

// GetMdmEnabledOk returns a tuple with the MdmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10Mdm) GetMdmEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.MdmEnabled) {
		return nil, false
	}
	return o.MdmEnabled, true
}

// HasMdmEnabled returns a boolean if a field has been set.
func (o *InlineObject10Mdm) HasMdmEnabled() bool {
	if o != nil && !IsNil(o.MdmEnabled) {
		return true
	}

	return false
}

// SetMdmEnabled gets a reference to the given string and assigns it to the MdmEnabled field.
func (o *InlineObject10Mdm) SetMdmEnabled(v string) {
	o.MdmEnabled = &v
}

// GetSupervised returns the Supervised field value if set, zero value otherwise.
func (o *InlineObject10Mdm) GetSupervised() string {
	if o == nil || IsNil(o.Supervised) {
		var ret string
		return ret
	}
	return *o.Supervised
}

// GetSupervisedOk returns a tuple with the Supervised field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10Mdm) GetSupervisedOk() (*string, bool) {
	if o == nil || IsNil(o.Supervised) {
		return nil, false
	}
	return o.Supervised, true
}

// HasSupervised returns a boolean if a field has been set.
func (o *InlineObject10Mdm) HasSupervised() bool {
	if o != nil && !IsNil(o.Supervised) {
		return true
	}

	return false
}

// SetSupervised gets a reference to the given string and assigns it to the Supervised field.
func (o *InlineObject10Mdm) SetSupervised(v string) {
	o.Supervised = &v
}

// GetInstallDate returns the InstallDate field value if set, zero value otherwise.
func (o *InlineObject10Mdm) GetInstallDate() string {
	if o == nil || IsNil(o.InstallDate) {
		var ret string
		return ret
	}
	return *o.InstallDate
}

// GetInstallDateOk returns a tuple with the InstallDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10Mdm) GetInstallDateOk() (*string, bool) {
	if o == nil || IsNil(o.InstallDate) {
		return nil, false
	}
	return o.InstallDate, true
}

// HasInstallDate returns a boolean if a field has been set.
func (o *InlineObject10Mdm) HasInstallDate() bool {
	if o != nil && !IsNil(o.InstallDate) {
		return true
	}

	return false
}

// SetInstallDate gets a reference to the given string and assigns it to the InstallDate field.
func (o *InlineObject10Mdm) SetInstallDate(v string) {
	o.InstallDate = &v
}

// GetLastCheckIn returns the LastCheckIn field value if set, zero value otherwise.
func (o *InlineObject10Mdm) GetLastCheckIn() string {
	if o == nil || IsNil(o.LastCheckIn) {
		var ret string
		return ret
	}
	return *o.LastCheckIn
}

// GetLastCheckInOk returns a tuple with the LastCheckIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10Mdm) GetLastCheckInOk() (*string, bool) {
	if o == nil || IsNil(o.LastCheckIn) {
		return nil, false
	}
	return o.LastCheckIn, true
}

// HasLastCheckIn returns a boolean if a field has been set.
func (o *InlineObject10Mdm) HasLastCheckIn() bool {
	if o != nil && !IsNil(o.LastCheckIn) {
		return true
	}

	return false
}

// SetLastCheckIn gets a reference to the given string and assigns it to the LastCheckIn field.
func (o *InlineObject10Mdm) SetLastCheckIn(v string) {
	o.LastCheckIn = &v
}

// GetMdmEnabledUser returns the MdmEnabledUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject10Mdm) GetMdmEnabledUser() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MdmEnabledUser
}

// GetMdmEnabledUserOk returns a tuple with the MdmEnabledUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject10Mdm) GetMdmEnabledUserOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MdmEnabledUser) {
		return nil, false
	}
	return &o.MdmEnabledUser, true
}

// HasMdmEnabledUser returns a boolean if a field has been set.
func (o *InlineObject10Mdm) HasMdmEnabledUser() bool {
	if o != nil && !IsNil(o.MdmEnabledUser) {
		return true
	}

	return false
}

// SetMdmEnabledUser gets a reference to the given interface{} and assigns it to the MdmEnabledUser field.
func (o *InlineObject10Mdm) SetMdmEnabledUser(v interface{}) {
	o.MdmEnabledUser = v
}

func (o InlineObject10Mdm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject10Mdm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MdmEnabled) {
		toSerialize["mdm_enabled"] = o.MdmEnabled
	}
	if !IsNil(o.Supervised) {
		toSerialize["supervised"] = o.Supervised
	}
	if !IsNil(o.InstallDate) {
		toSerialize["install_date"] = o.InstallDate
	}
	if !IsNil(o.LastCheckIn) {
		toSerialize["last_check_in"] = o.LastCheckIn
	}
	if o.MdmEnabledUser != nil {
		toSerialize["mdm_enabled_user"] = o.MdmEnabledUser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InlineObject10Mdm) UnmarshalJSON(data []byte) (err error) {
	varInlineObject10Mdm := _InlineObject10Mdm{}

	err = json.Unmarshal(data, &varInlineObject10Mdm)

	if err != nil {
		return err
	}

	*o = InlineObject10Mdm(varInlineObject10Mdm)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mdm_enabled")
		delete(additionalProperties, "supervised")
		delete(additionalProperties, "install_date")
		delete(additionalProperties, "last_check_in")
		delete(additionalProperties, "mdm_enabled_user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineObject10Mdm struct {
	value *InlineObject10Mdm
	isSet bool
}

func (v NullableInlineObject10Mdm) Get() *InlineObject10Mdm {
	return v.value
}

func (v *NullableInlineObject10Mdm) Set(val *InlineObject10Mdm) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject10Mdm) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject10Mdm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject10Mdm(val *InlineObject10Mdm) *NullableInlineObject10Mdm {
	return &NullableInlineObject10Mdm{value: val, isSet: true}
}

func (v NullableInlineObject10Mdm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject10Mdm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


