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

// checks if the DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation{}

// DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation struct for DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation
type DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation struct {
	Latitude *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation

// NewDeviceInformationGetDeviceDetails200ResponseLostModeLastLocation instantiates a new DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceInformationGetDeviceDetails200ResponseLostModeLastLocation() *DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation {
	this := DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation{}
	return &this
}

// NewDeviceInformationGetDeviceDetails200ResponseLostModeLastLocationWithDefaults instantiates a new DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceInformationGetDeviceDetails200ResponseLostModeLastLocationWithDefaults() *DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation {
	this := DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation{}
	return &this
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) GetLatitude() string {
	if o == nil || IsNil(o.Latitude) {
		var ret string
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) GetLatitudeOk() (*string, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given string and assigns it to the Latitude field.
func (o *DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) SetLatitude(v string) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) GetLongitude() string {
	if o == nil || IsNil(o.Longitude) {
		var ret string
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) GetLongitudeOk() (*string, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given string and assigns it to the Longitude field.
func (o *DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) SetLongitude(v string) {
	o.Longitude = &v
}

func (o DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) UnmarshalJSON(data []byte) (err error) {
	varDeviceInformationGetDeviceDetails200ResponseLostModeLastLocation := _DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation{}

	err = json.Unmarshal(data, &varDeviceInformationGetDeviceDetails200ResponseLostModeLastLocation)

	if err != nil {
		return err
	}

	*o = DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation(varDeviceInformationGetDeviceDetails200ResponseLostModeLastLocation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "latitude")
		delete(additionalProperties, "longitude")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceInformationGetDeviceDetails200ResponseLostModeLastLocation struct {
	value *DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation
	isSet bool
}

func (v NullableDeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) Get() *DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation {
	return v.value
}

func (v *NullableDeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) Set(val *DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceInformationGetDeviceDetails200ResponseLostModeLastLocation(val *DeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) *NullableDeviceInformationGetDeviceDetails200ResponseLostModeLastLocation {
	return &NullableDeviceInformationGetDeviceDetails200ResponseLostModeLastLocation{value: val, isSet: true}
}

func (v NullableDeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceInformationGetDeviceDetails200ResponseLostModeLastLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


