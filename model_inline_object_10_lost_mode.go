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

// checks if the InlineObject10LostMode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject10LostMode{}

// InlineObject10LostMode struct for InlineObject10LostMode
type InlineObject10LostMode struct {
	LostModeStatus *string `json:"lost_mode_status,omitempty"`
	EnabledBy *string `json:"enabled_by,omitempty"`
	EnableStatusAt *string `json:"enable_status_at,omitempty"`
	LockScreenMessage *string `json:"lock_screen_message,omitempty"`
	LockScreenPhoneNumber *string `json:"lock_screen_phone_number,omitempty"`
	LockScreenFootnote *string `json:"lock_screen_footnote,omitempty"`
	DisableStatus *string `json:"disable_status,omitempty"`
	DisabledBy *string `json:"disabled_by,omitempty"`
	DisableStatusAt *string `json:"disable_status_at,omitempty"`
	LastLocationStatus *string `json:"last_location_status,omitempty"`
	LastLocationStatusAt *string `json:"last_location_status_at,omitempty"`
	LastLocation *InlineObject10LostModeLastLocation `json:"last_location,omitempty"`
	LastLocationAt *string `json:"last_location_at,omitempty"`
	SoundStatus *string `json:"sound_status,omitempty"`
	SoundStatusAt *string `json:"sound_status_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineObject10LostMode InlineObject10LostMode

// NewInlineObject10LostMode instantiates a new InlineObject10LostMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject10LostMode() *InlineObject10LostMode {
	this := InlineObject10LostMode{}
	return &this
}

// NewInlineObject10LostModeWithDefaults instantiates a new InlineObject10LostMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject10LostModeWithDefaults() *InlineObject10LostMode {
	this := InlineObject10LostMode{}
	return &this
}

// GetLostModeStatus returns the LostModeStatus field value if set, zero value otherwise.
func (o *InlineObject10LostMode) GetLostModeStatus() string {
	if o == nil || IsNil(o.LostModeStatus) {
		var ret string
		return ret
	}
	return *o.LostModeStatus
}

// GetLostModeStatusOk returns a tuple with the LostModeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10LostMode) GetLostModeStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LostModeStatus) {
		return nil, false
	}
	return o.LostModeStatus, true
}

// HasLostModeStatus returns a boolean if a field has been set.
func (o *InlineObject10LostMode) HasLostModeStatus() bool {
	if o != nil && !IsNil(o.LostModeStatus) {
		return true
	}

	return false
}

// SetLostModeStatus gets a reference to the given string and assigns it to the LostModeStatus field.
func (o *InlineObject10LostMode) SetLostModeStatus(v string) {
	o.LostModeStatus = &v
}

// GetEnabledBy returns the EnabledBy field value if set, zero value otherwise.
func (o *InlineObject10LostMode) GetEnabledBy() string {
	if o == nil || IsNil(o.EnabledBy) {
		var ret string
		return ret
	}
	return *o.EnabledBy
}

// GetEnabledByOk returns a tuple with the EnabledBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10LostMode) GetEnabledByOk() (*string, bool) {
	if o == nil || IsNil(o.EnabledBy) {
		return nil, false
	}
	return o.EnabledBy, true
}

// HasEnabledBy returns a boolean if a field has been set.
func (o *InlineObject10LostMode) HasEnabledBy() bool {
	if o != nil && !IsNil(o.EnabledBy) {
		return true
	}

	return false
}

// SetEnabledBy gets a reference to the given string and assigns it to the EnabledBy field.
func (o *InlineObject10LostMode) SetEnabledBy(v string) {
	o.EnabledBy = &v
}

// GetEnableStatusAt returns the EnableStatusAt field value if set, zero value otherwise.
func (o *InlineObject10LostMode) GetEnableStatusAt() string {
	if o == nil || IsNil(o.EnableStatusAt) {
		var ret string
		return ret
	}
	return *o.EnableStatusAt
}

// GetEnableStatusAtOk returns a tuple with the EnableStatusAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10LostMode) GetEnableStatusAtOk() (*string, bool) {
	if o == nil || IsNil(o.EnableStatusAt) {
		return nil, false
	}
	return o.EnableStatusAt, true
}

// HasEnableStatusAt returns a boolean if a field has been set.
func (o *InlineObject10LostMode) HasEnableStatusAt() bool {
	if o != nil && !IsNil(o.EnableStatusAt) {
		return true
	}

	return false
}

// SetEnableStatusAt gets a reference to the given string and assigns it to the EnableStatusAt field.
func (o *InlineObject10LostMode) SetEnableStatusAt(v string) {
	o.EnableStatusAt = &v
}

// GetLockScreenMessage returns the LockScreenMessage field value if set, zero value otherwise.
func (o *InlineObject10LostMode) GetLockScreenMessage() string {
	if o == nil || IsNil(o.LockScreenMessage) {
		var ret string
		return ret
	}
	return *o.LockScreenMessage
}

// GetLockScreenMessageOk returns a tuple with the LockScreenMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10LostMode) GetLockScreenMessageOk() (*string, bool) {
	if o == nil || IsNil(o.LockScreenMessage) {
		return nil, false
	}
	return o.LockScreenMessage, true
}

// HasLockScreenMessage returns a boolean if a field has been set.
func (o *InlineObject10LostMode) HasLockScreenMessage() bool {
	if o != nil && !IsNil(o.LockScreenMessage) {
		return true
	}

	return false
}

// SetLockScreenMessage gets a reference to the given string and assigns it to the LockScreenMessage field.
func (o *InlineObject10LostMode) SetLockScreenMessage(v string) {
	o.LockScreenMessage = &v
}

// GetLockScreenPhoneNumber returns the LockScreenPhoneNumber field value if set, zero value otherwise.
func (o *InlineObject10LostMode) GetLockScreenPhoneNumber() string {
	if o == nil || IsNil(o.LockScreenPhoneNumber) {
		var ret string
		return ret
	}
	return *o.LockScreenPhoneNumber
}

// GetLockScreenPhoneNumberOk returns a tuple with the LockScreenPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10LostMode) GetLockScreenPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.LockScreenPhoneNumber) {
		return nil, false
	}
	return o.LockScreenPhoneNumber, true
}

// HasLockScreenPhoneNumber returns a boolean if a field has been set.
func (o *InlineObject10LostMode) HasLockScreenPhoneNumber() bool {
	if o != nil && !IsNil(o.LockScreenPhoneNumber) {
		return true
	}

	return false
}

// SetLockScreenPhoneNumber gets a reference to the given string and assigns it to the LockScreenPhoneNumber field.
func (o *InlineObject10LostMode) SetLockScreenPhoneNumber(v string) {
	o.LockScreenPhoneNumber = &v
}

// GetLockScreenFootnote returns the LockScreenFootnote field value if set, zero value otherwise.
func (o *InlineObject10LostMode) GetLockScreenFootnote() string {
	if o == nil || IsNil(o.LockScreenFootnote) {
		var ret string
		return ret
	}
	return *o.LockScreenFootnote
}

// GetLockScreenFootnoteOk returns a tuple with the LockScreenFootnote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10LostMode) GetLockScreenFootnoteOk() (*string, bool) {
	if o == nil || IsNil(o.LockScreenFootnote) {
		return nil, false
	}
	return o.LockScreenFootnote, true
}

// HasLockScreenFootnote returns a boolean if a field has been set.
func (o *InlineObject10LostMode) HasLockScreenFootnote() bool {
	if o != nil && !IsNil(o.LockScreenFootnote) {
		return true
	}

	return false
}

// SetLockScreenFootnote gets a reference to the given string and assigns it to the LockScreenFootnote field.
func (o *InlineObject10LostMode) SetLockScreenFootnote(v string) {
	o.LockScreenFootnote = &v
}

// GetDisableStatus returns the DisableStatus field value if set, zero value otherwise.
func (o *InlineObject10LostMode) GetDisableStatus() string {
	if o == nil || IsNil(o.DisableStatus) {
		var ret string
		return ret
	}
	return *o.DisableStatus
}

// GetDisableStatusOk returns a tuple with the DisableStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10LostMode) GetDisableStatusOk() (*string, bool) {
	if o == nil || IsNil(o.DisableStatus) {
		return nil, false
	}
	return o.DisableStatus, true
}

// HasDisableStatus returns a boolean if a field has been set.
func (o *InlineObject10LostMode) HasDisableStatus() bool {
	if o != nil && !IsNil(o.DisableStatus) {
		return true
	}

	return false
}

// SetDisableStatus gets a reference to the given string and assigns it to the DisableStatus field.
func (o *InlineObject10LostMode) SetDisableStatus(v string) {
	o.DisableStatus = &v
}

// GetDisabledBy returns the DisabledBy field value if set, zero value otherwise.
func (o *InlineObject10LostMode) GetDisabledBy() string {
	if o == nil || IsNil(o.DisabledBy) {
		var ret string
		return ret
	}
	return *o.DisabledBy
}

// GetDisabledByOk returns a tuple with the DisabledBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10LostMode) GetDisabledByOk() (*string, bool) {
	if o == nil || IsNil(o.DisabledBy) {
		return nil, false
	}
	return o.DisabledBy, true
}

// HasDisabledBy returns a boolean if a field has been set.
func (o *InlineObject10LostMode) HasDisabledBy() bool {
	if o != nil && !IsNil(o.DisabledBy) {
		return true
	}

	return false
}

// SetDisabledBy gets a reference to the given string and assigns it to the DisabledBy field.
func (o *InlineObject10LostMode) SetDisabledBy(v string) {
	o.DisabledBy = &v
}

// GetDisableStatusAt returns the DisableStatusAt field value if set, zero value otherwise.
func (o *InlineObject10LostMode) GetDisableStatusAt() string {
	if o == nil || IsNil(o.DisableStatusAt) {
		var ret string
		return ret
	}
	return *o.DisableStatusAt
}

// GetDisableStatusAtOk returns a tuple with the DisableStatusAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10LostMode) GetDisableStatusAtOk() (*string, bool) {
	if o == nil || IsNil(o.DisableStatusAt) {
		return nil, false
	}
	return o.DisableStatusAt, true
}

// HasDisableStatusAt returns a boolean if a field has been set.
func (o *InlineObject10LostMode) HasDisableStatusAt() bool {
	if o != nil && !IsNil(o.DisableStatusAt) {
		return true
	}

	return false
}

// SetDisableStatusAt gets a reference to the given string and assigns it to the DisableStatusAt field.
func (o *InlineObject10LostMode) SetDisableStatusAt(v string) {
	o.DisableStatusAt = &v
}

// GetLastLocationStatus returns the LastLocationStatus field value if set, zero value otherwise.
func (o *InlineObject10LostMode) GetLastLocationStatus() string {
	if o == nil || IsNil(o.LastLocationStatus) {
		var ret string
		return ret
	}
	return *o.LastLocationStatus
}

// GetLastLocationStatusOk returns a tuple with the LastLocationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10LostMode) GetLastLocationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LastLocationStatus) {
		return nil, false
	}
	return o.LastLocationStatus, true
}

// HasLastLocationStatus returns a boolean if a field has been set.
func (o *InlineObject10LostMode) HasLastLocationStatus() bool {
	if o != nil && !IsNil(o.LastLocationStatus) {
		return true
	}

	return false
}

// SetLastLocationStatus gets a reference to the given string and assigns it to the LastLocationStatus field.
func (o *InlineObject10LostMode) SetLastLocationStatus(v string) {
	o.LastLocationStatus = &v
}

// GetLastLocationStatusAt returns the LastLocationStatusAt field value if set, zero value otherwise.
func (o *InlineObject10LostMode) GetLastLocationStatusAt() string {
	if o == nil || IsNil(o.LastLocationStatusAt) {
		var ret string
		return ret
	}
	return *o.LastLocationStatusAt
}

// GetLastLocationStatusAtOk returns a tuple with the LastLocationStatusAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10LostMode) GetLastLocationStatusAtOk() (*string, bool) {
	if o == nil || IsNil(o.LastLocationStatusAt) {
		return nil, false
	}
	return o.LastLocationStatusAt, true
}

// HasLastLocationStatusAt returns a boolean if a field has been set.
func (o *InlineObject10LostMode) HasLastLocationStatusAt() bool {
	if o != nil && !IsNil(o.LastLocationStatusAt) {
		return true
	}

	return false
}

// SetLastLocationStatusAt gets a reference to the given string and assigns it to the LastLocationStatusAt field.
func (o *InlineObject10LostMode) SetLastLocationStatusAt(v string) {
	o.LastLocationStatusAt = &v
}

// GetLastLocation returns the LastLocation field value if set, zero value otherwise.
func (o *InlineObject10LostMode) GetLastLocation() InlineObject10LostModeLastLocation {
	if o == nil || IsNil(o.LastLocation) {
		var ret InlineObject10LostModeLastLocation
		return ret
	}
	return *o.LastLocation
}

// GetLastLocationOk returns a tuple with the LastLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10LostMode) GetLastLocationOk() (*InlineObject10LostModeLastLocation, bool) {
	if o == nil || IsNil(o.LastLocation) {
		return nil, false
	}
	return o.LastLocation, true
}

// HasLastLocation returns a boolean if a field has been set.
func (o *InlineObject10LostMode) HasLastLocation() bool {
	if o != nil && !IsNil(o.LastLocation) {
		return true
	}

	return false
}

// SetLastLocation gets a reference to the given InlineObject10LostModeLastLocation and assigns it to the LastLocation field.
func (o *InlineObject10LostMode) SetLastLocation(v InlineObject10LostModeLastLocation) {
	o.LastLocation = &v
}

// GetLastLocationAt returns the LastLocationAt field value if set, zero value otherwise.
func (o *InlineObject10LostMode) GetLastLocationAt() string {
	if o == nil || IsNil(o.LastLocationAt) {
		var ret string
		return ret
	}
	return *o.LastLocationAt
}

// GetLastLocationAtOk returns a tuple with the LastLocationAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10LostMode) GetLastLocationAtOk() (*string, bool) {
	if o == nil || IsNil(o.LastLocationAt) {
		return nil, false
	}
	return o.LastLocationAt, true
}

// HasLastLocationAt returns a boolean if a field has been set.
func (o *InlineObject10LostMode) HasLastLocationAt() bool {
	if o != nil && !IsNil(o.LastLocationAt) {
		return true
	}

	return false
}

// SetLastLocationAt gets a reference to the given string and assigns it to the LastLocationAt field.
func (o *InlineObject10LostMode) SetLastLocationAt(v string) {
	o.LastLocationAt = &v
}

// GetSoundStatus returns the SoundStatus field value if set, zero value otherwise.
func (o *InlineObject10LostMode) GetSoundStatus() string {
	if o == nil || IsNil(o.SoundStatus) {
		var ret string
		return ret
	}
	return *o.SoundStatus
}

// GetSoundStatusOk returns a tuple with the SoundStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10LostMode) GetSoundStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SoundStatus) {
		return nil, false
	}
	return o.SoundStatus, true
}

// HasSoundStatus returns a boolean if a field has been set.
func (o *InlineObject10LostMode) HasSoundStatus() bool {
	if o != nil && !IsNil(o.SoundStatus) {
		return true
	}

	return false
}

// SetSoundStatus gets a reference to the given string and assigns it to the SoundStatus field.
func (o *InlineObject10LostMode) SetSoundStatus(v string) {
	o.SoundStatus = &v
}

// GetSoundStatusAt returns the SoundStatusAt field value if set, zero value otherwise.
func (o *InlineObject10LostMode) GetSoundStatusAt() string {
	if o == nil || IsNil(o.SoundStatusAt) {
		var ret string
		return ret
	}
	return *o.SoundStatusAt
}

// GetSoundStatusAtOk returns a tuple with the SoundStatusAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10LostMode) GetSoundStatusAtOk() (*string, bool) {
	if o == nil || IsNil(o.SoundStatusAt) {
		return nil, false
	}
	return o.SoundStatusAt, true
}

// HasSoundStatusAt returns a boolean if a field has been set.
func (o *InlineObject10LostMode) HasSoundStatusAt() bool {
	if o != nil && !IsNil(o.SoundStatusAt) {
		return true
	}

	return false
}

// SetSoundStatusAt gets a reference to the given string and assigns it to the SoundStatusAt field.
func (o *InlineObject10LostMode) SetSoundStatusAt(v string) {
	o.SoundStatusAt = &v
}

func (o InlineObject10LostMode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject10LostMode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LostModeStatus) {
		toSerialize["lost_mode_status"] = o.LostModeStatus
	}
	if !IsNil(o.EnabledBy) {
		toSerialize["enabled_by"] = o.EnabledBy
	}
	if !IsNil(o.EnableStatusAt) {
		toSerialize["enable_status_at"] = o.EnableStatusAt
	}
	if !IsNil(o.LockScreenMessage) {
		toSerialize["lock_screen_message"] = o.LockScreenMessage
	}
	if !IsNil(o.LockScreenPhoneNumber) {
		toSerialize["lock_screen_phone_number"] = o.LockScreenPhoneNumber
	}
	if !IsNil(o.LockScreenFootnote) {
		toSerialize["lock_screen_footnote"] = o.LockScreenFootnote
	}
	if !IsNil(o.DisableStatus) {
		toSerialize["disable_status"] = o.DisableStatus
	}
	if !IsNil(o.DisabledBy) {
		toSerialize["disabled_by"] = o.DisabledBy
	}
	if !IsNil(o.DisableStatusAt) {
		toSerialize["disable_status_at"] = o.DisableStatusAt
	}
	if !IsNil(o.LastLocationStatus) {
		toSerialize["last_location_status"] = o.LastLocationStatus
	}
	if !IsNil(o.LastLocationStatusAt) {
		toSerialize["last_location_status_at"] = o.LastLocationStatusAt
	}
	if !IsNil(o.LastLocation) {
		toSerialize["last_location"] = o.LastLocation
	}
	if !IsNil(o.LastLocationAt) {
		toSerialize["last_location_at"] = o.LastLocationAt
	}
	if !IsNil(o.SoundStatus) {
		toSerialize["sound_status"] = o.SoundStatus
	}
	if !IsNil(o.SoundStatusAt) {
		toSerialize["sound_status_at"] = o.SoundStatusAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InlineObject10LostMode) UnmarshalJSON(data []byte) (err error) {
	varInlineObject10LostMode := _InlineObject10LostMode{}

	err = json.Unmarshal(data, &varInlineObject10LostMode)

	if err != nil {
		return err
	}

	*o = InlineObject10LostMode(varInlineObject10LostMode)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lost_mode_status")
		delete(additionalProperties, "enabled_by")
		delete(additionalProperties, "enable_status_at")
		delete(additionalProperties, "lock_screen_message")
		delete(additionalProperties, "lock_screen_phone_number")
		delete(additionalProperties, "lock_screen_footnote")
		delete(additionalProperties, "disable_status")
		delete(additionalProperties, "disabled_by")
		delete(additionalProperties, "disable_status_at")
		delete(additionalProperties, "last_location_status")
		delete(additionalProperties, "last_location_status_at")
		delete(additionalProperties, "last_location")
		delete(additionalProperties, "last_location_at")
		delete(additionalProperties, "sound_status")
		delete(additionalProperties, "sound_status_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineObject10LostMode struct {
	value *InlineObject10LostMode
	isSet bool
}

func (v NullableInlineObject10LostMode) Get() *InlineObject10LostMode {
	return v.value
}

func (v *NullableInlineObject10LostMode) Set(val *InlineObject10LostMode) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject10LostMode) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject10LostMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject10LostMode(val *InlineObject10LostMode) *NullableInlineObject10LostMode {
	return &NullableInlineObject10LostMode{value: val, isSet: true}
}

func (v NullableInlineObject10LostMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject10LostMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


