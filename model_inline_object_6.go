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

// checks if the InlineObject6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject6{}

// InlineObject6 struct for InlineObject6
type InlineObject6 struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Icon *string `json:"icon,omitempty"`
	Color *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Params map[string]interface{} `json:"params,omitempty"`
	ComputersCount *int32 `json:"computers_count,omitempty"`
	EnrollmentCode *InlineObject5EnrollmentCode `json:"enrollment_code,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineObject6 InlineObject6

// NewInlineObject6 instantiates a new InlineObject6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject6() *InlineObject6 {
	this := InlineObject6{}
	return &this
}

// NewInlineObject6WithDefaults instantiates a new InlineObject6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject6WithDefaults() *InlineObject6 {
	this := InlineObject6{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineObject6) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineObject6) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineObject6) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject6) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject6) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject6) SetName(v string) {
	o.Name = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *InlineObject6) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *InlineObject6) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *InlineObject6) SetIcon(v string) {
	o.Icon = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *InlineObject6) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *InlineObject6) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *InlineObject6) SetColor(v string) {
	o.Color = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject6) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject6) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject6) SetDescription(v string) {
	o.Description = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *InlineObject6) GetParams() map[string]interface{} {
	if o == nil || IsNil(o.Params) {
		var ret map[string]interface{}
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Params) {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *InlineObject6) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]interface{} and assigns it to the Params field.
func (o *InlineObject6) SetParams(v map[string]interface{}) {
	o.Params = v
}

// GetComputersCount returns the ComputersCount field value if set, zero value otherwise.
func (o *InlineObject6) GetComputersCount() int32 {
	if o == nil || IsNil(o.ComputersCount) {
		var ret int32
		return ret
	}
	return *o.ComputersCount
}

// GetComputersCountOk returns a tuple with the ComputersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetComputersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ComputersCount) {
		return nil, false
	}
	return o.ComputersCount, true
}

// HasComputersCount returns a boolean if a field has been set.
func (o *InlineObject6) HasComputersCount() bool {
	if o != nil && !IsNil(o.ComputersCount) {
		return true
	}

	return false
}

// SetComputersCount gets a reference to the given int32 and assigns it to the ComputersCount field.
func (o *InlineObject6) SetComputersCount(v int32) {
	o.ComputersCount = &v
}

// GetEnrollmentCode returns the EnrollmentCode field value if set, zero value otherwise.
func (o *InlineObject6) GetEnrollmentCode() InlineObject5EnrollmentCode {
	if o == nil || IsNil(o.EnrollmentCode) {
		var ret InlineObject5EnrollmentCode
		return ret
	}
	return *o.EnrollmentCode
}

// GetEnrollmentCodeOk returns a tuple with the EnrollmentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetEnrollmentCodeOk() (*InlineObject5EnrollmentCode, bool) {
	if o == nil || IsNil(o.EnrollmentCode) {
		return nil, false
	}
	return o.EnrollmentCode, true
}

// HasEnrollmentCode returns a boolean if a field has been set.
func (o *InlineObject6) HasEnrollmentCode() bool {
	if o != nil && !IsNil(o.EnrollmentCode) {
		return true
	}

	return false
}

// SetEnrollmentCode gets a reference to the given InlineObject5EnrollmentCode and assigns it to the EnrollmentCode field.
func (o *InlineObject6) SetEnrollmentCode(v InlineObject5EnrollmentCode) {
	o.EnrollmentCode = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineObject6) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineObject6) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineObject6) SetType(v string) {
	o.Type = &v
}

func (o InlineObject6) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	if !IsNil(o.ComputersCount) {
		toSerialize["computers_count"] = o.ComputersCount
	}
	if !IsNil(o.EnrollmentCode) {
		toSerialize["enrollment_code"] = o.EnrollmentCode
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InlineObject6) UnmarshalJSON(data []byte) (err error) {
	varInlineObject6 := _InlineObject6{}

	err = json.Unmarshal(data, &varInlineObject6)

	if err != nil {
		return err
	}

	*o = InlineObject6(varInlineObject6)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "icon")
		delete(additionalProperties, "color")
		delete(additionalProperties, "description")
		delete(additionalProperties, "params")
		delete(additionalProperties, "computers_count")
		delete(additionalProperties, "enrollment_code")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineObject6 struct {
	value *InlineObject6
	isSet bool
}

func (v NullableInlineObject6) Get() *InlineObject6 {
	return v.value
}

func (v *NullableInlineObject6) Set(val *InlineObject6) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject6) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject6(val *InlineObject6) *NullableInlineObject6 {
	return &NullableInlineObject6{value: val, isSet: true}
}

func (v NullableInlineObject6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


