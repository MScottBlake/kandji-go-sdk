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

// checks if the InlineObject10 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject10{}

// InlineObject10 struct for InlineObject10
type InlineObject10 struct {
	General *InlineObject10General `json:"general,omitempty"`
	Mdm *InlineObject10Mdm `json:"mdm,omitempty"`
	ActivationLock *InlineObject10ActivationLock `json:"activation_lock,omitempty"`
	Filevault *InlineObject10Filevault `json:"filevault,omitempty"`
	LostMode *InlineObject10LostMode `json:"lost_mode,omitempty"`
	AutomatedDeviceEnrollment *InlineObject10AutomatedDeviceEnrollment `json:"automated_device_enrollment,omitempty"`
	KandjiAgent *InlineObject10KandjiAgent `json:"kandji_agent,omitempty"`
	HardwareOverview *InlineObject10HardwareOverview `json:"hardware_overview,omitempty"`
	Volumes interface{} `json:"volumes,omitempty"`
	Network map[string]interface{} `json:"network,omitempty"`
	RecoveryInformation *InlineObject10RecoveryInformation `json:"recovery_information,omitempty"`
	Users *InlineObject10Users `json:"users,omitempty"`
	InstalledProfiles interface{} `json:"installed_profiles,omitempty"`
	AppleBusinessManager *InlineObject10AppleBusinessManager `json:"apple_business_manager,omitempty"`
	SecurityInformation *InlineObject10SecurityInformation `json:"security_information,omitempty"`
	Cellular *InlineObject10Cellular `json:"cellular,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineObject10 InlineObject10

// NewInlineObject10 instantiates a new InlineObject10 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject10() *InlineObject10 {
	this := InlineObject10{}
	return &this
}

// NewInlineObject10WithDefaults instantiates a new InlineObject10 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject10WithDefaults() *InlineObject10 {
	this := InlineObject10{}
	return &this
}

// GetGeneral returns the General field value if set, zero value otherwise.
func (o *InlineObject10) GetGeneral() InlineObject10General {
	if o == nil || IsNil(o.General) {
		var ret InlineObject10General
		return ret
	}
	return *o.General
}

// GetGeneralOk returns a tuple with the General field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetGeneralOk() (*InlineObject10General, bool) {
	if o == nil || IsNil(o.General) {
		return nil, false
	}
	return o.General, true
}

// HasGeneral returns a boolean if a field has been set.
func (o *InlineObject10) HasGeneral() bool {
	if o != nil && !IsNil(o.General) {
		return true
	}

	return false
}

// SetGeneral gets a reference to the given InlineObject10General and assigns it to the General field.
func (o *InlineObject10) SetGeneral(v InlineObject10General) {
	o.General = &v
}

// GetMdm returns the Mdm field value if set, zero value otherwise.
func (o *InlineObject10) GetMdm() InlineObject10Mdm {
	if o == nil || IsNil(o.Mdm) {
		var ret InlineObject10Mdm
		return ret
	}
	return *o.Mdm
}

// GetMdmOk returns a tuple with the Mdm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetMdmOk() (*InlineObject10Mdm, bool) {
	if o == nil || IsNil(o.Mdm) {
		return nil, false
	}
	return o.Mdm, true
}

// HasMdm returns a boolean if a field has been set.
func (o *InlineObject10) HasMdm() bool {
	if o != nil && !IsNil(o.Mdm) {
		return true
	}

	return false
}

// SetMdm gets a reference to the given InlineObject10Mdm and assigns it to the Mdm field.
func (o *InlineObject10) SetMdm(v InlineObject10Mdm) {
	o.Mdm = &v
}

// GetActivationLock returns the ActivationLock field value if set, zero value otherwise.
func (o *InlineObject10) GetActivationLock() InlineObject10ActivationLock {
	if o == nil || IsNil(o.ActivationLock) {
		var ret InlineObject10ActivationLock
		return ret
	}
	return *o.ActivationLock
}

// GetActivationLockOk returns a tuple with the ActivationLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetActivationLockOk() (*InlineObject10ActivationLock, bool) {
	if o == nil || IsNil(o.ActivationLock) {
		return nil, false
	}
	return o.ActivationLock, true
}

// HasActivationLock returns a boolean if a field has been set.
func (o *InlineObject10) HasActivationLock() bool {
	if o != nil && !IsNil(o.ActivationLock) {
		return true
	}

	return false
}

// SetActivationLock gets a reference to the given InlineObject10ActivationLock and assigns it to the ActivationLock field.
func (o *InlineObject10) SetActivationLock(v InlineObject10ActivationLock) {
	o.ActivationLock = &v
}

// GetFilevault returns the Filevault field value if set, zero value otherwise.
func (o *InlineObject10) GetFilevault() InlineObject10Filevault {
	if o == nil || IsNil(o.Filevault) {
		var ret InlineObject10Filevault
		return ret
	}
	return *o.Filevault
}

// GetFilevaultOk returns a tuple with the Filevault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetFilevaultOk() (*InlineObject10Filevault, bool) {
	if o == nil || IsNil(o.Filevault) {
		return nil, false
	}
	return o.Filevault, true
}

// HasFilevault returns a boolean if a field has been set.
func (o *InlineObject10) HasFilevault() bool {
	if o != nil && !IsNil(o.Filevault) {
		return true
	}

	return false
}

// SetFilevault gets a reference to the given InlineObject10Filevault and assigns it to the Filevault field.
func (o *InlineObject10) SetFilevault(v InlineObject10Filevault) {
	o.Filevault = &v
}

// GetLostMode returns the LostMode field value if set, zero value otherwise.
func (o *InlineObject10) GetLostMode() InlineObject10LostMode {
	if o == nil || IsNil(o.LostMode) {
		var ret InlineObject10LostMode
		return ret
	}
	return *o.LostMode
}

// GetLostModeOk returns a tuple with the LostMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetLostModeOk() (*InlineObject10LostMode, bool) {
	if o == nil || IsNil(o.LostMode) {
		return nil, false
	}
	return o.LostMode, true
}

// HasLostMode returns a boolean if a field has been set.
func (o *InlineObject10) HasLostMode() bool {
	if o != nil && !IsNil(o.LostMode) {
		return true
	}

	return false
}

// SetLostMode gets a reference to the given InlineObject10LostMode and assigns it to the LostMode field.
func (o *InlineObject10) SetLostMode(v InlineObject10LostMode) {
	o.LostMode = &v
}

// GetAutomatedDeviceEnrollment returns the AutomatedDeviceEnrollment field value if set, zero value otherwise.
func (o *InlineObject10) GetAutomatedDeviceEnrollment() InlineObject10AutomatedDeviceEnrollment {
	if o == nil || IsNil(o.AutomatedDeviceEnrollment) {
		var ret InlineObject10AutomatedDeviceEnrollment
		return ret
	}
	return *o.AutomatedDeviceEnrollment
}

// GetAutomatedDeviceEnrollmentOk returns a tuple with the AutomatedDeviceEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetAutomatedDeviceEnrollmentOk() (*InlineObject10AutomatedDeviceEnrollment, bool) {
	if o == nil || IsNil(o.AutomatedDeviceEnrollment) {
		return nil, false
	}
	return o.AutomatedDeviceEnrollment, true
}

// HasAutomatedDeviceEnrollment returns a boolean if a field has been set.
func (o *InlineObject10) HasAutomatedDeviceEnrollment() bool {
	if o != nil && !IsNil(o.AutomatedDeviceEnrollment) {
		return true
	}

	return false
}

// SetAutomatedDeviceEnrollment gets a reference to the given InlineObject10AutomatedDeviceEnrollment and assigns it to the AutomatedDeviceEnrollment field.
func (o *InlineObject10) SetAutomatedDeviceEnrollment(v InlineObject10AutomatedDeviceEnrollment) {
	o.AutomatedDeviceEnrollment = &v
}

// GetKandjiAgent returns the KandjiAgent field value if set, zero value otherwise.
func (o *InlineObject10) GetKandjiAgent() InlineObject10KandjiAgent {
	if o == nil || IsNil(o.KandjiAgent) {
		var ret InlineObject10KandjiAgent
		return ret
	}
	return *o.KandjiAgent
}

// GetKandjiAgentOk returns a tuple with the KandjiAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetKandjiAgentOk() (*InlineObject10KandjiAgent, bool) {
	if o == nil || IsNil(o.KandjiAgent) {
		return nil, false
	}
	return o.KandjiAgent, true
}

// HasKandjiAgent returns a boolean if a field has been set.
func (o *InlineObject10) HasKandjiAgent() bool {
	if o != nil && !IsNil(o.KandjiAgent) {
		return true
	}

	return false
}

// SetKandjiAgent gets a reference to the given InlineObject10KandjiAgent and assigns it to the KandjiAgent field.
func (o *InlineObject10) SetKandjiAgent(v InlineObject10KandjiAgent) {
	o.KandjiAgent = &v
}

// GetHardwareOverview returns the HardwareOverview field value if set, zero value otherwise.
func (o *InlineObject10) GetHardwareOverview() InlineObject10HardwareOverview {
	if o == nil || IsNil(o.HardwareOverview) {
		var ret InlineObject10HardwareOverview
		return ret
	}
	return *o.HardwareOverview
}

// GetHardwareOverviewOk returns a tuple with the HardwareOverview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetHardwareOverviewOk() (*InlineObject10HardwareOverview, bool) {
	if o == nil || IsNil(o.HardwareOverview) {
		return nil, false
	}
	return o.HardwareOverview, true
}

// HasHardwareOverview returns a boolean if a field has been set.
func (o *InlineObject10) HasHardwareOverview() bool {
	if o != nil && !IsNil(o.HardwareOverview) {
		return true
	}

	return false
}

// SetHardwareOverview gets a reference to the given InlineObject10HardwareOverview and assigns it to the HardwareOverview field.
func (o *InlineObject10) SetHardwareOverview(v InlineObject10HardwareOverview) {
	o.HardwareOverview = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject10) GetVolumes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject10) GetVolumesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Volumes) {
		return nil, false
	}
	return &o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *InlineObject10) HasVolumes() bool {
	if o != nil && !IsNil(o.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given interface{} and assigns it to the Volumes field.
func (o *InlineObject10) SetVolumes(v interface{}) {
	o.Volumes = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineObject10) GetNetwork() map[string]interface{} {
	if o == nil || IsNil(o.Network) {
		var ret map[string]interface{}
		return ret
	}
	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetNetworkOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Network) {
		return map[string]interface{}{}, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineObject10) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given map[string]interface{} and assigns it to the Network field.
func (o *InlineObject10) SetNetwork(v map[string]interface{}) {
	o.Network = v
}

// GetRecoveryInformation returns the RecoveryInformation field value if set, zero value otherwise.
func (o *InlineObject10) GetRecoveryInformation() InlineObject10RecoveryInformation {
	if o == nil || IsNil(o.RecoveryInformation) {
		var ret InlineObject10RecoveryInformation
		return ret
	}
	return *o.RecoveryInformation
}

// GetRecoveryInformationOk returns a tuple with the RecoveryInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetRecoveryInformationOk() (*InlineObject10RecoveryInformation, bool) {
	if o == nil || IsNil(o.RecoveryInformation) {
		return nil, false
	}
	return o.RecoveryInformation, true
}

// HasRecoveryInformation returns a boolean if a field has been set.
func (o *InlineObject10) HasRecoveryInformation() bool {
	if o != nil && !IsNil(o.RecoveryInformation) {
		return true
	}

	return false
}

// SetRecoveryInformation gets a reference to the given InlineObject10RecoveryInformation and assigns it to the RecoveryInformation field.
func (o *InlineObject10) SetRecoveryInformation(v InlineObject10RecoveryInformation) {
	o.RecoveryInformation = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *InlineObject10) GetUsers() InlineObject10Users {
	if o == nil || IsNil(o.Users) {
		var ret InlineObject10Users
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetUsersOk() (*InlineObject10Users, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *InlineObject10) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given InlineObject10Users and assigns it to the Users field.
func (o *InlineObject10) SetUsers(v InlineObject10Users) {
	o.Users = &v
}

// GetInstalledProfiles returns the InstalledProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject10) GetInstalledProfiles() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.InstalledProfiles
}

// GetInstalledProfilesOk returns a tuple with the InstalledProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject10) GetInstalledProfilesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.InstalledProfiles) {
		return nil, false
	}
	return &o.InstalledProfiles, true
}

// HasInstalledProfiles returns a boolean if a field has been set.
func (o *InlineObject10) HasInstalledProfiles() bool {
	if o != nil && !IsNil(o.InstalledProfiles) {
		return true
	}

	return false
}

// SetInstalledProfiles gets a reference to the given interface{} and assigns it to the InstalledProfiles field.
func (o *InlineObject10) SetInstalledProfiles(v interface{}) {
	o.InstalledProfiles = v
}

// GetAppleBusinessManager returns the AppleBusinessManager field value if set, zero value otherwise.
func (o *InlineObject10) GetAppleBusinessManager() InlineObject10AppleBusinessManager {
	if o == nil || IsNil(o.AppleBusinessManager) {
		var ret InlineObject10AppleBusinessManager
		return ret
	}
	return *o.AppleBusinessManager
}

// GetAppleBusinessManagerOk returns a tuple with the AppleBusinessManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetAppleBusinessManagerOk() (*InlineObject10AppleBusinessManager, bool) {
	if o == nil || IsNil(o.AppleBusinessManager) {
		return nil, false
	}
	return o.AppleBusinessManager, true
}

// HasAppleBusinessManager returns a boolean if a field has been set.
func (o *InlineObject10) HasAppleBusinessManager() bool {
	if o != nil && !IsNil(o.AppleBusinessManager) {
		return true
	}

	return false
}

// SetAppleBusinessManager gets a reference to the given InlineObject10AppleBusinessManager and assigns it to the AppleBusinessManager field.
func (o *InlineObject10) SetAppleBusinessManager(v InlineObject10AppleBusinessManager) {
	o.AppleBusinessManager = &v
}

// GetSecurityInformation returns the SecurityInformation field value if set, zero value otherwise.
func (o *InlineObject10) GetSecurityInformation() InlineObject10SecurityInformation {
	if o == nil || IsNil(o.SecurityInformation) {
		var ret InlineObject10SecurityInformation
		return ret
	}
	return *o.SecurityInformation
}

// GetSecurityInformationOk returns a tuple with the SecurityInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetSecurityInformationOk() (*InlineObject10SecurityInformation, bool) {
	if o == nil || IsNil(o.SecurityInformation) {
		return nil, false
	}
	return o.SecurityInformation, true
}

// HasSecurityInformation returns a boolean if a field has been set.
func (o *InlineObject10) HasSecurityInformation() bool {
	if o != nil && !IsNil(o.SecurityInformation) {
		return true
	}

	return false
}

// SetSecurityInformation gets a reference to the given InlineObject10SecurityInformation and assigns it to the SecurityInformation field.
func (o *InlineObject10) SetSecurityInformation(v InlineObject10SecurityInformation) {
	o.SecurityInformation = &v
}

// GetCellular returns the Cellular field value if set, zero value otherwise.
func (o *InlineObject10) GetCellular() InlineObject10Cellular {
	if o == nil || IsNil(o.Cellular) {
		var ret InlineObject10Cellular
		return ret
	}
	return *o.Cellular
}

// GetCellularOk returns a tuple with the Cellular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetCellularOk() (*InlineObject10Cellular, bool) {
	if o == nil || IsNil(o.Cellular) {
		return nil, false
	}
	return o.Cellular, true
}

// HasCellular returns a boolean if a field has been set.
func (o *InlineObject10) HasCellular() bool {
	if o != nil && !IsNil(o.Cellular) {
		return true
	}

	return false
}

// SetCellular gets a reference to the given InlineObject10Cellular and assigns it to the Cellular field.
func (o *InlineObject10) SetCellular(v InlineObject10Cellular) {
	o.Cellular = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject10) GetTags() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject10) GetTagsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject10) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given interface{} and assigns it to the Tags field.
func (o *InlineObject10) SetTags(v interface{}) {
	o.Tags = v
}

func (o InlineObject10) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject10) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.General) {
		toSerialize["general"] = o.General
	}
	if !IsNil(o.Mdm) {
		toSerialize["mdm"] = o.Mdm
	}
	if !IsNil(o.ActivationLock) {
		toSerialize["activation_lock"] = o.ActivationLock
	}
	if !IsNil(o.Filevault) {
		toSerialize["filevault"] = o.Filevault
	}
	if !IsNil(o.LostMode) {
		toSerialize["lost_mode"] = o.LostMode
	}
	if !IsNil(o.AutomatedDeviceEnrollment) {
		toSerialize["automated_device_enrollment"] = o.AutomatedDeviceEnrollment
	}
	if !IsNil(o.KandjiAgent) {
		toSerialize["kandji_agent"] = o.KandjiAgent
	}
	if !IsNil(o.HardwareOverview) {
		toSerialize["hardware_overview"] = o.HardwareOverview
	}
	if o.Volumes != nil {
		toSerialize["volumes"] = o.Volumes
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.RecoveryInformation) {
		toSerialize["recovery_information"] = o.RecoveryInformation
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if o.InstalledProfiles != nil {
		toSerialize["installed_profiles"] = o.InstalledProfiles
	}
	if !IsNil(o.AppleBusinessManager) {
		toSerialize["apple_business_manager"] = o.AppleBusinessManager
	}
	if !IsNil(o.SecurityInformation) {
		toSerialize["security_information"] = o.SecurityInformation
	}
	if !IsNil(o.Cellular) {
		toSerialize["cellular"] = o.Cellular
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InlineObject10) UnmarshalJSON(data []byte) (err error) {
	varInlineObject10 := _InlineObject10{}

	err = json.Unmarshal(data, &varInlineObject10)

	if err != nil {
		return err
	}

	*o = InlineObject10(varInlineObject10)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "general")
		delete(additionalProperties, "mdm")
		delete(additionalProperties, "activation_lock")
		delete(additionalProperties, "filevault")
		delete(additionalProperties, "lost_mode")
		delete(additionalProperties, "automated_device_enrollment")
		delete(additionalProperties, "kandji_agent")
		delete(additionalProperties, "hardware_overview")
		delete(additionalProperties, "volumes")
		delete(additionalProperties, "network")
		delete(additionalProperties, "recovery_information")
		delete(additionalProperties, "users")
		delete(additionalProperties, "installed_profiles")
		delete(additionalProperties, "apple_business_manager")
		delete(additionalProperties, "security_information")
		delete(additionalProperties, "cellular")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineObject10 struct {
	value *InlineObject10
	isSet bool
}

func (v NullableInlineObject10) Get() *InlineObject10 {
	return v.value
}

func (v *NullableInlineObject10) Set(val *InlineObject10) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject10) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject10) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject10(val *InlineObject10) *NullableInlineObject10 {
	return &NullableInlineObject10{value: val, isSet: true}
}

func (v NullableInlineObject10) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject10) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


