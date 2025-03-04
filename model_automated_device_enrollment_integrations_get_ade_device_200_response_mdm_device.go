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

// checks if the AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice{}

// AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice struct for AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice
type AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice struct {
	DeferredInstall *int32 `json:"deferred_install,omitempty"`
	EnrolledAt *string `json:"enrolled_at,omitempty"`
	EnrollmentStatus *int32 `json:"enrollment_status,omitempty"`
	Id *string `json:"id,omitempty"`
	IsMissing *int32 `json:"is_missing,omitempty"`
	IsRemoved *int32 `json:"is_removed,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice

// NewAutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice instantiates a new AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice() *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice {
	this := AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice{}
	return &this
}

// NewAutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDeviceWithDefaults instantiates a new AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDeviceWithDefaults() *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice {
	this := AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice{}
	return &this
}

// GetDeferredInstall returns the DeferredInstall field value if set, zero value otherwise.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) GetDeferredInstall() int32 {
	if o == nil || IsNil(o.DeferredInstall) {
		var ret int32
		return ret
	}
	return *o.DeferredInstall
}

// GetDeferredInstallOk returns a tuple with the DeferredInstall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) GetDeferredInstallOk() (*int32, bool) {
	if o == nil || IsNil(o.DeferredInstall) {
		return nil, false
	}
	return o.DeferredInstall, true
}

// HasDeferredInstall returns a boolean if a field has been set.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) HasDeferredInstall() bool {
	if o != nil && !IsNil(o.DeferredInstall) {
		return true
	}

	return false
}

// SetDeferredInstall gets a reference to the given int32 and assigns it to the DeferredInstall field.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) SetDeferredInstall(v int32) {
	o.DeferredInstall = &v
}

// GetEnrolledAt returns the EnrolledAt field value if set, zero value otherwise.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) GetEnrolledAt() string {
	if o == nil || IsNil(o.EnrolledAt) {
		var ret string
		return ret
	}
	return *o.EnrolledAt
}

// GetEnrolledAtOk returns a tuple with the EnrolledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) GetEnrolledAtOk() (*string, bool) {
	if o == nil || IsNil(o.EnrolledAt) {
		return nil, false
	}
	return o.EnrolledAt, true
}

// HasEnrolledAt returns a boolean if a field has been set.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) HasEnrolledAt() bool {
	if o != nil && !IsNil(o.EnrolledAt) {
		return true
	}

	return false
}

// SetEnrolledAt gets a reference to the given string and assigns it to the EnrolledAt field.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) SetEnrolledAt(v string) {
	o.EnrolledAt = &v
}

// GetEnrollmentStatus returns the EnrollmentStatus field value if set, zero value otherwise.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) GetEnrollmentStatus() int32 {
	if o == nil || IsNil(o.EnrollmentStatus) {
		var ret int32
		return ret
	}
	return *o.EnrollmentStatus
}

// GetEnrollmentStatusOk returns a tuple with the EnrollmentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) GetEnrollmentStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.EnrollmentStatus) {
		return nil, false
	}
	return o.EnrollmentStatus, true
}

// HasEnrollmentStatus returns a boolean if a field has been set.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) HasEnrollmentStatus() bool {
	if o != nil && !IsNil(o.EnrollmentStatus) {
		return true
	}

	return false
}

// SetEnrollmentStatus gets a reference to the given int32 and assigns it to the EnrollmentStatus field.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) SetEnrollmentStatus(v int32) {
	o.EnrollmentStatus = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) SetId(v string) {
	o.Id = &v
}

// GetIsMissing returns the IsMissing field value if set, zero value otherwise.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) GetIsMissing() int32 {
	if o == nil || IsNil(o.IsMissing) {
		var ret int32
		return ret
	}
	return *o.IsMissing
}

// GetIsMissingOk returns a tuple with the IsMissing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) GetIsMissingOk() (*int32, bool) {
	if o == nil || IsNil(o.IsMissing) {
		return nil, false
	}
	return o.IsMissing, true
}

// HasIsMissing returns a boolean if a field has been set.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) HasIsMissing() bool {
	if o != nil && !IsNil(o.IsMissing) {
		return true
	}

	return false
}

// SetIsMissing gets a reference to the given int32 and assigns it to the IsMissing field.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) SetIsMissing(v int32) {
	o.IsMissing = &v
}

// GetIsRemoved returns the IsRemoved field value if set, zero value otherwise.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) GetIsRemoved() int32 {
	if o == nil || IsNil(o.IsRemoved) {
		var ret int32
		return ret
	}
	return *o.IsRemoved
}

// GetIsRemovedOk returns a tuple with the IsRemoved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) GetIsRemovedOk() (*int32, bool) {
	if o == nil || IsNil(o.IsRemoved) {
		return nil, false
	}
	return o.IsRemoved, true
}

// HasIsRemoved returns a boolean if a field has been set.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) HasIsRemoved() bool {
	if o != nil && !IsNil(o.IsRemoved) {
		return true
	}

	return false
}

// SetIsRemoved gets a reference to the given int32 and assigns it to the IsRemoved field.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) SetIsRemoved(v int32) {
	o.IsRemoved = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) SetName(v string) {
	o.Name = &v
}

func (o AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeferredInstall) {
		toSerialize["deferred_install"] = o.DeferredInstall
	}
	if !IsNil(o.EnrolledAt) {
		toSerialize["enrolled_at"] = o.EnrolledAt
	}
	if !IsNil(o.EnrollmentStatus) {
		toSerialize["enrollment_status"] = o.EnrollmentStatus
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsMissing) {
		toSerialize["is_missing"] = o.IsMissing
	}
	if !IsNil(o.IsRemoved) {
		toSerialize["is_removed"] = o.IsRemoved
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) UnmarshalJSON(data []byte) (err error) {
	varAutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice := _AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice{}

	err = json.Unmarshal(data, &varAutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice)

	if err != nil {
		return err
	}

	*o = AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice(varAutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "deferred_install")
		delete(additionalProperties, "enrolled_at")
		delete(additionalProperties, "enrollment_status")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_missing")
		delete(additionalProperties, "is_removed")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice struct {
	value *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice
	isSet bool
}

func (v NullableAutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) Get() *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice {
	return v.value
}

func (v *NullableAutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) Set(val *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice(val *AutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) *NullableAutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice {
	return &NullableAutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice{value: val, isSet: true}
}

func (v NullableAutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomatedDeviceEnrollmentIntegrationsGetAdeDevice200ResponseMdmDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


