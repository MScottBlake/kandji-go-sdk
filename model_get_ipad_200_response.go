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

// checks if the GetIpad200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIpad200Response{}

// GetIpad200Response struct for GetIpad200Response
type GetIpad200Response struct {
	DeviceId *string `json:"device_id,omitempty"`
	DeviceName *string `json:"device_name,omitempty"`
	Model *string `json:"model,omitempty"`
	SerialNumber *string `json:"serial_number,omitempty"`
	Platform *string `json:"platform,omitempty"`
	OsVersion *string `json:"os_version,omitempty"`
	SupplementalBuildVersion *string `json:"supplemental_build_version,omitempty"`
	SupplementalOsVersionExtra *string `json:"supplemental_os_version_extra,omitempty"`
	LastCheckIn *string `json:"last_check_in,omitempty"`
	User *GetIpad200ResponseUser `json:"user,omitempty"`
	AssetTag *string `json:"asset_tag,omitempty"`
	BlueprintId *string `json:"blueprint_id,omitempty"`
	MdmEnabled *int32 `json:"mdm_enabled,omitempty"`
	AgentInstalled *int32 `json:"agent_installed,omitempty"`
	IsMissing *int32 `json:"is_missing,omitempty"`
	IsRemoved *int32 `json:"is_removed,omitempty"`
	AgentVersion *string `json:"agent_version,omitempty"`
	FirstEnrollment *string `json:"first_enrollment,omitempty"`
	LastEnrollment *string `json:"last_enrollment,omitempty"`
	BlueprintName *string `json:"blueprint_name,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIpad200Response GetIpad200Response

// NewGetIpad200Response instantiates a new GetIpad200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIpad200Response() *GetIpad200Response {
	this := GetIpad200Response{}
	return &this
}

// NewGetIpad200ResponseWithDefaults instantiates a new GetIpad200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIpad200ResponseWithDefaults() *GetIpad200Response {
	this := GetIpad200Response{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *GetIpad200Response) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *GetIpad200Response) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *GetIpad200Response) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *GetIpad200Response) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetDeviceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *GetIpad200Response) HasDeviceName() bool {
	if o != nil && !IsNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *GetIpad200Response) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GetIpad200Response) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GetIpad200Response) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GetIpad200Response) SetModel(v string) {
	o.Model = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *GetIpad200Response) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *GetIpad200Response) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *GetIpad200Response) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *GetIpad200Response) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *GetIpad200Response) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *GetIpad200Response) SetPlatform(v string) {
	o.Platform = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *GetIpad200Response) GetOsVersion() string {
	if o == nil || IsNil(o.OsVersion) {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetOsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *GetIpad200Response) HasOsVersion() bool {
	if o != nil && !IsNil(o.OsVersion) {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *GetIpad200Response) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetSupplementalBuildVersion returns the SupplementalBuildVersion field value if set, zero value otherwise.
func (o *GetIpad200Response) GetSupplementalBuildVersion() string {
	if o == nil || IsNil(o.SupplementalBuildVersion) {
		var ret string
		return ret
	}
	return *o.SupplementalBuildVersion
}

// GetSupplementalBuildVersionOk returns a tuple with the SupplementalBuildVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetSupplementalBuildVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SupplementalBuildVersion) {
		return nil, false
	}
	return o.SupplementalBuildVersion, true
}

// HasSupplementalBuildVersion returns a boolean if a field has been set.
func (o *GetIpad200Response) HasSupplementalBuildVersion() bool {
	if o != nil && !IsNil(o.SupplementalBuildVersion) {
		return true
	}

	return false
}

// SetSupplementalBuildVersion gets a reference to the given string and assigns it to the SupplementalBuildVersion field.
func (o *GetIpad200Response) SetSupplementalBuildVersion(v string) {
	o.SupplementalBuildVersion = &v
}

// GetSupplementalOsVersionExtra returns the SupplementalOsVersionExtra field value if set, zero value otherwise.
func (o *GetIpad200Response) GetSupplementalOsVersionExtra() string {
	if o == nil || IsNil(o.SupplementalOsVersionExtra) {
		var ret string
		return ret
	}
	return *o.SupplementalOsVersionExtra
}

// GetSupplementalOsVersionExtraOk returns a tuple with the SupplementalOsVersionExtra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetSupplementalOsVersionExtraOk() (*string, bool) {
	if o == nil || IsNil(o.SupplementalOsVersionExtra) {
		return nil, false
	}
	return o.SupplementalOsVersionExtra, true
}

// HasSupplementalOsVersionExtra returns a boolean if a field has been set.
func (o *GetIpad200Response) HasSupplementalOsVersionExtra() bool {
	if o != nil && !IsNil(o.SupplementalOsVersionExtra) {
		return true
	}

	return false
}

// SetSupplementalOsVersionExtra gets a reference to the given string and assigns it to the SupplementalOsVersionExtra field.
func (o *GetIpad200Response) SetSupplementalOsVersionExtra(v string) {
	o.SupplementalOsVersionExtra = &v
}

// GetLastCheckIn returns the LastCheckIn field value if set, zero value otherwise.
func (o *GetIpad200Response) GetLastCheckIn() string {
	if o == nil || IsNil(o.LastCheckIn) {
		var ret string
		return ret
	}
	return *o.LastCheckIn
}

// GetLastCheckInOk returns a tuple with the LastCheckIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetLastCheckInOk() (*string, bool) {
	if o == nil || IsNil(o.LastCheckIn) {
		return nil, false
	}
	return o.LastCheckIn, true
}

// HasLastCheckIn returns a boolean if a field has been set.
func (o *GetIpad200Response) HasLastCheckIn() bool {
	if o != nil && !IsNil(o.LastCheckIn) {
		return true
	}

	return false
}

// SetLastCheckIn gets a reference to the given string and assigns it to the LastCheckIn field.
func (o *GetIpad200Response) SetLastCheckIn(v string) {
	o.LastCheckIn = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *GetIpad200Response) GetUser() GetIpad200ResponseUser {
	if o == nil || IsNil(o.User) {
		var ret GetIpad200ResponseUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetUserOk() (*GetIpad200ResponseUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *GetIpad200Response) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given GetIpad200ResponseUser and assigns it to the User field.
func (o *GetIpad200Response) SetUser(v GetIpad200ResponseUser) {
	o.User = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise.
func (o *GetIpad200Response) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag) {
		var ret string
		return ret
	}
	return *o.AssetTag
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetAssetTagOk() (*string, bool) {
	if o == nil || IsNil(o.AssetTag) {
		return nil, false
	}
	return o.AssetTag, true
}

// HasAssetTag returns a boolean if a field has been set.
func (o *GetIpad200Response) HasAssetTag() bool {
	if o != nil && !IsNil(o.AssetTag) {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given string and assigns it to the AssetTag field.
func (o *GetIpad200Response) SetAssetTag(v string) {
	o.AssetTag = &v
}

// GetBlueprintId returns the BlueprintId field value if set, zero value otherwise.
func (o *GetIpad200Response) GetBlueprintId() string {
	if o == nil || IsNil(o.BlueprintId) {
		var ret string
		return ret
	}
	return *o.BlueprintId
}

// GetBlueprintIdOk returns a tuple with the BlueprintId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetBlueprintIdOk() (*string, bool) {
	if o == nil || IsNil(o.BlueprintId) {
		return nil, false
	}
	return o.BlueprintId, true
}

// HasBlueprintId returns a boolean if a field has been set.
func (o *GetIpad200Response) HasBlueprintId() bool {
	if o != nil && !IsNil(o.BlueprintId) {
		return true
	}

	return false
}

// SetBlueprintId gets a reference to the given string and assigns it to the BlueprintId field.
func (o *GetIpad200Response) SetBlueprintId(v string) {
	o.BlueprintId = &v
}

// GetMdmEnabled returns the MdmEnabled field value if set, zero value otherwise.
func (o *GetIpad200Response) GetMdmEnabled() int32 {
	if o == nil || IsNil(o.MdmEnabled) {
		var ret int32
		return ret
	}
	return *o.MdmEnabled
}

// GetMdmEnabledOk returns a tuple with the MdmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetMdmEnabledOk() (*int32, bool) {
	if o == nil || IsNil(o.MdmEnabled) {
		return nil, false
	}
	return o.MdmEnabled, true
}

// HasMdmEnabled returns a boolean if a field has been set.
func (o *GetIpad200Response) HasMdmEnabled() bool {
	if o != nil && !IsNil(o.MdmEnabled) {
		return true
	}

	return false
}

// SetMdmEnabled gets a reference to the given int32 and assigns it to the MdmEnabled field.
func (o *GetIpad200Response) SetMdmEnabled(v int32) {
	o.MdmEnabled = &v
}

// GetAgentInstalled returns the AgentInstalled field value if set, zero value otherwise.
func (o *GetIpad200Response) GetAgentInstalled() int32 {
	if o == nil || IsNil(o.AgentInstalled) {
		var ret int32
		return ret
	}
	return *o.AgentInstalled
}

// GetAgentInstalledOk returns a tuple with the AgentInstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetAgentInstalledOk() (*int32, bool) {
	if o == nil || IsNil(o.AgentInstalled) {
		return nil, false
	}
	return o.AgentInstalled, true
}

// HasAgentInstalled returns a boolean if a field has been set.
func (o *GetIpad200Response) HasAgentInstalled() bool {
	if o != nil && !IsNil(o.AgentInstalled) {
		return true
	}

	return false
}

// SetAgentInstalled gets a reference to the given int32 and assigns it to the AgentInstalled field.
func (o *GetIpad200Response) SetAgentInstalled(v int32) {
	o.AgentInstalled = &v
}

// GetIsMissing returns the IsMissing field value if set, zero value otherwise.
func (o *GetIpad200Response) GetIsMissing() int32 {
	if o == nil || IsNil(o.IsMissing) {
		var ret int32
		return ret
	}
	return *o.IsMissing
}

// GetIsMissingOk returns a tuple with the IsMissing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetIsMissingOk() (*int32, bool) {
	if o == nil || IsNil(o.IsMissing) {
		return nil, false
	}
	return o.IsMissing, true
}

// HasIsMissing returns a boolean if a field has been set.
func (o *GetIpad200Response) HasIsMissing() bool {
	if o != nil && !IsNil(o.IsMissing) {
		return true
	}

	return false
}

// SetIsMissing gets a reference to the given int32 and assigns it to the IsMissing field.
func (o *GetIpad200Response) SetIsMissing(v int32) {
	o.IsMissing = &v
}

// GetIsRemoved returns the IsRemoved field value if set, zero value otherwise.
func (o *GetIpad200Response) GetIsRemoved() int32 {
	if o == nil || IsNil(o.IsRemoved) {
		var ret int32
		return ret
	}
	return *o.IsRemoved
}

// GetIsRemovedOk returns a tuple with the IsRemoved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetIsRemovedOk() (*int32, bool) {
	if o == nil || IsNil(o.IsRemoved) {
		return nil, false
	}
	return o.IsRemoved, true
}

// HasIsRemoved returns a boolean if a field has been set.
func (o *GetIpad200Response) HasIsRemoved() bool {
	if o != nil && !IsNil(o.IsRemoved) {
		return true
	}

	return false
}

// SetIsRemoved gets a reference to the given int32 and assigns it to the IsRemoved field.
func (o *GetIpad200Response) SetIsRemoved(v int32) {
	o.IsRemoved = &v
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise.
func (o *GetIpad200Response) GetAgentVersion() string {
	if o == nil || IsNil(o.AgentVersion) {
		var ret string
		return ret
	}
	return *o.AgentVersion
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetAgentVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AgentVersion) {
		return nil, false
	}
	return o.AgentVersion, true
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *GetIpad200Response) HasAgentVersion() bool {
	if o != nil && !IsNil(o.AgentVersion) {
		return true
	}

	return false
}

// SetAgentVersion gets a reference to the given string and assigns it to the AgentVersion field.
func (o *GetIpad200Response) SetAgentVersion(v string) {
	o.AgentVersion = &v
}

// GetFirstEnrollment returns the FirstEnrollment field value if set, zero value otherwise.
func (o *GetIpad200Response) GetFirstEnrollment() string {
	if o == nil || IsNil(o.FirstEnrollment) {
		var ret string
		return ret
	}
	return *o.FirstEnrollment
}

// GetFirstEnrollmentOk returns a tuple with the FirstEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetFirstEnrollmentOk() (*string, bool) {
	if o == nil || IsNil(o.FirstEnrollment) {
		return nil, false
	}
	return o.FirstEnrollment, true
}

// HasFirstEnrollment returns a boolean if a field has been set.
func (o *GetIpad200Response) HasFirstEnrollment() bool {
	if o != nil && !IsNil(o.FirstEnrollment) {
		return true
	}

	return false
}

// SetFirstEnrollment gets a reference to the given string and assigns it to the FirstEnrollment field.
func (o *GetIpad200Response) SetFirstEnrollment(v string) {
	o.FirstEnrollment = &v
}

// GetLastEnrollment returns the LastEnrollment field value if set, zero value otherwise.
func (o *GetIpad200Response) GetLastEnrollment() string {
	if o == nil || IsNil(o.LastEnrollment) {
		var ret string
		return ret
	}
	return *o.LastEnrollment
}

// GetLastEnrollmentOk returns a tuple with the LastEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetLastEnrollmentOk() (*string, bool) {
	if o == nil || IsNil(o.LastEnrollment) {
		return nil, false
	}
	return o.LastEnrollment, true
}

// HasLastEnrollment returns a boolean if a field has been set.
func (o *GetIpad200Response) HasLastEnrollment() bool {
	if o != nil && !IsNil(o.LastEnrollment) {
		return true
	}

	return false
}

// SetLastEnrollment gets a reference to the given string and assigns it to the LastEnrollment field.
func (o *GetIpad200Response) SetLastEnrollment(v string) {
	o.LastEnrollment = &v
}

// GetBlueprintName returns the BlueprintName field value if set, zero value otherwise.
func (o *GetIpad200Response) GetBlueprintName() string {
	if o == nil || IsNil(o.BlueprintName) {
		var ret string
		return ret
	}
	return *o.BlueprintName
}

// GetBlueprintNameOk returns a tuple with the BlueprintName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpad200Response) GetBlueprintNameOk() (*string, bool) {
	if o == nil || IsNil(o.BlueprintName) {
		return nil, false
	}
	return o.BlueprintName, true
}

// HasBlueprintName returns a boolean if a field has been set.
func (o *GetIpad200Response) HasBlueprintName() bool {
	if o != nil && !IsNil(o.BlueprintName) {
		return true
	}

	return false
}

// SetBlueprintName gets a reference to the given string and assigns it to the BlueprintName field.
func (o *GetIpad200Response) SetBlueprintName(v string) {
	o.BlueprintName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetIpad200Response) GetTags() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetIpad200Response) GetTagsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetIpad200Response) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given interface{} and assigns it to the Tags field.
func (o *GetIpad200Response) SetTags(v interface{}) {
	o.Tags = v
}

func (o GetIpad200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIpad200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceId) {
		toSerialize["device_id"] = o.DeviceId
	}
	if !IsNil(o.DeviceName) {
		toSerialize["device_name"] = o.DeviceName
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.OsVersion) {
		toSerialize["os_version"] = o.OsVersion
	}
	if !IsNil(o.SupplementalBuildVersion) {
		toSerialize["supplemental_build_version"] = o.SupplementalBuildVersion
	}
	if !IsNil(o.SupplementalOsVersionExtra) {
		toSerialize["supplemental_os_version_extra"] = o.SupplementalOsVersionExtra
	}
	if !IsNil(o.LastCheckIn) {
		toSerialize["last_check_in"] = o.LastCheckIn
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.AssetTag) {
		toSerialize["asset_tag"] = o.AssetTag
	}
	if !IsNil(o.BlueprintId) {
		toSerialize["blueprint_id"] = o.BlueprintId
	}
	if !IsNil(o.MdmEnabled) {
		toSerialize["mdm_enabled"] = o.MdmEnabled
	}
	if !IsNil(o.AgentInstalled) {
		toSerialize["agent_installed"] = o.AgentInstalled
	}
	if !IsNil(o.IsMissing) {
		toSerialize["is_missing"] = o.IsMissing
	}
	if !IsNil(o.IsRemoved) {
		toSerialize["is_removed"] = o.IsRemoved
	}
	if !IsNil(o.AgentVersion) {
		toSerialize["agent_version"] = o.AgentVersion
	}
	if !IsNil(o.FirstEnrollment) {
		toSerialize["first_enrollment"] = o.FirstEnrollment
	}
	if !IsNil(o.LastEnrollment) {
		toSerialize["last_enrollment"] = o.LastEnrollment
	}
	if !IsNil(o.BlueprintName) {
		toSerialize["blueprint_name"] = o.BlueprintName
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetIpad200Response) UnmarshalJSON(data []byte) (err error) {
	varGetIpad200Response := _GetIpad200Response{}

	err = json.Unmarshal(data, &varGetIpad200Response)

	if err != nil {
		return err
	}

	*o = GetIpad200Response(varGetIpad200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_id")
		delete(additionalProperties, "device_name")
		delete(additionalProperties, "model")
		delete(additionalProperties, "serial_number")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "os_version")
		delete(additionalProperties, "supplemental_build_version")
		delete(additionalProperties, "supplemental_os_version_extra")
		delete(additionalProperties, "last_check_in")
		delete(additionalProperties, "user")
		delete(additionalProperties, "asset_tag")
		delete(additionalProperties, "blueprint_id")
		delete(additionalProperties, "mdm_enabled")
		delete(additionalProperties, "agent_installed")
		delete(additionalProperties, "is_missing")
		delete(additionalProperties, "is_removed")
		delete(additionalProperties, "agent_version")
		delete(additionalProperties, "first_enrollment")
		delete(additionalProperties, "last_enrollment")
		delete(additionalProperties, "blueprint_name")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIpad200Response struct {
	value *GetIpad200Response
	isSet bool
}

func (v NullableGetIpad200Response) Get() *GetIpad200Response {
	return v.value
}

func (v *NullableGetIpad200Response) Set(val *GetIpad200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIpad200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIpad200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIpad200Response(val *GetIpad200Response) *NullableGetIpad200Response {
	return &NullableGetIpad200Response{value: val, isSet: true}
}

func (v NullableGetIpad200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIpad200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


