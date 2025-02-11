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

// checks if the PrismRequestCategoryExport200ResponseArgs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrismRequestCategoryExport200ResponseArgs{}

// PrismRequestCategoryExport200ResponseArgs struct for PrismRequestCategoryExport200ResponseArgs
type PrismRequestCategoryExport200ResponseArgs struct {
	Filter map[string]interface{} `json:"filter,omitempty"`
	Columns interface{} `json:"columns,omitempty"`
	SortBy *string `json:"sort_by,omitempty"`
	BlueprintIds interface{} `json:"blueprint_ids,omitempty"`
	DeviceFamilies interface{} `json:"device_families,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrismRequestCategoryExport200ResponseArgs PrismRequestCategoryExport200ResponseArgs

// NewPrismRequestCategoryExport200ResponseArgs instantiates a new PrismRequestCategoryExport200ResponseArgs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrismRequestCategoryExport200ResponseArgs() *PrismRequestCategoryExport200ResponseArgs {
	this := PrismRequestCategoryExport200ResponseArgs{}
	return &this
}

// NewPrismRequestCategoryExport200ResponseArgsWithDefaults instantiates a new PrismRequestCategoryExport200ResponseArgs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrismRequestCategoryExport200ResponseArgsWithDefaults() *PrismRequestCategoryExport200ResponseArgs {
	this := PrismRequestCategoryExport200ResponseArgs{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *PrismRequestCategoryExport200ResponseArgs) GetFilter() map[string]interface{} {
	if o == nil || IsNil(o.Filter) {
		var ret map[string]interface{}
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrismRequestCategoryExport200ResponseArgs) GetFilterOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Filter) {
		return map[string]interface{}{}, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *PrismRequestCategoryExport200ResponseArgs) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given map[string]interface{} and assigns it to the Filter field.
func (o *PrismRequestCategoryExport200ResponseArgs) SetFilter(v map[string]interface{}) {
	o.Filter = v
}

// GetColumns returns the Columns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrismRequestCategoryExport200ResponseArgs) GetColumns() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrismRequestCategoryExport200ResponseArgs) GetColumnsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *PrismRequestCategoryExport200ResponseArgs) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given interface{} and assigns it to the Columns field.
func (o *PrismRequestCategoryExport200ResponseArgs) SetColumns(v interface{}) {
	o.Columns = v
}

// GetSortBy returns the SortBy field value if set, zero value otherwise.
func (o *PrismRequestCategoryExport200ResponseArgs) GetSortBy() string {
	if o == nil || IsNil(o.SortBy) {
		var ret string
		return ret
	}
	return *o.SortBy
}

// GetSortByOk returns a tuple with the SortBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrismRequestCategoryExport200ResponseArgs) GetSortByOk() (*string, bool) {
	if o == nil || IsNil(o.SortBy) {
		return nil, false
	}
	return o.SortBy, true
}

// HasSortBy returns a boolean if a field has been set.
func (o *PrismRequestCategoryExport200ResponseArgs) HasSortBy() bool {
	if o != nil && !IsNil(o.SortBy) {
		return true
	}

	return false
}

// SetSortBy gets a reference to the given string and assigns it to the SortBy field.
func (o *PrismRequestCategoryExport200ResponseArgs) SetSortBy(v string) {
	o.SortBy = &v
}

// GetBlueprintIds returns the BlueprintIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrismRequestCategoryExport200ResponseArgs) GetBlueprintIds() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.BlueprintIds
}

// GetBlueprintIdsOk returns a tuple with the BlueprintIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrismRequestCategoryExport200ResponseArgs) GetBlueprintIdsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BlueprintIds) {
		return nil, false
	}
	return &o.BlueprintIds, true
}

// HasBlueprintIds returns a boolean if a field has been set.
func (o *PrismRequestCategoryExport200ResponseArgs) HasBlueprintIds() bool {
	if o != nil && !IsNil(o.BlueprintIds) {
		return true
	}

	return false
}

// SetBlueprintIds gets a reference to the given interface{} and assigns it to the BlueprintIds field.
func (o *PrismRequestCategoryExport200ResponseArgs) SetBlueprintIds(v interface{}) {
	o.BlueprintIds = v
}

// GetDeviceFamilies returns the DeviceFamilies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrismRequestCategoryExport200ResponseArgs) GetDeviceFamilies() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DeviceFamilies
}

// GetDeviceFamiliesOk returns a tuple with the DeviceFamilies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrismRequestCategoryExport200ResponseArgs) GetDeviceFamiliesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DeviceFamilies) {
		return nil, false
	}
	return &o.DeviceFamilies, true
}

// HasDeviceFamilies returns a boolean if a field has been set.
func (o *PrismRequestCategoryExport200ResponseArgs) HasDeviceFamilies() bool {
	if o != nil && !IsNil(o.DeviceFamilies) {
		return true
	}

	return false
}

// SetDeviceFamilies gets a reference to the given interface{} and assigns it to the DeviceFamilies field.
func (o *PrismRequestCategoryExport200ResponseArgs) SetDeviceFamilies(v interface{}) {
	o.DeviceFamilies = v
}

func (o PrismRequestCategoryExport200ResponseArgs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrismRequestCategoryExport200ResponseArgs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if !IsNil(o.SortBy) {
		toSerialize["sort_by"] = o.SortBy
	}
	if o.BlueprintIds != nil {
		toSerialize["blueprint_ids"] = o.BlueprintIds
	}
	if o.DeviceFamilies != nil {
		toSerialize["device_families"] = o.DeviceFamilies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrismRequestCategoryExport200ResponseArgs) UnmarshalJSON(data []byte) (err error) {
	varPrismRequestCategoryExport200ResponseArgs := _PrismRequestCategoryExport200ResponseArgs{}

	err = json.Unmarshal(data, &varPrismRequestCategoryExport200ResponseArgs)

	if err != nil {
		return err
	}

	*o = PrismRequestCategoryExport200ResponseArgs(varPrismRequestCategoryExport200ResponseArgs)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filter")
		delete(additionalProperties, "columns")
		delete(additionalProperties, "sort_by")
		delete(additionalProperties, "blueprint_ids")
		delete(additionalProperties, "device_families")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrismRequestCategoryExport200ResponseArgs struct {
	value *PrismRequestCategoryExport200ResponseArgs
	isSet bool
}

func (v NullablePrismRequestCategoryExport200ResponseArgs) Get() *PrismRequestCategoryExport200ResponseArgs {
	return v.value
}

func (v *NullablePrismRequestCategoryExport200ResponseArgs) Set(val *PrismRequestCategoryExport200ResponseArgs) {
	v.value = val
	v.isSet = true
}

func (v NullablePrismRequestCategoryExport200ResponseArgs) IsSet() bool {
	return v.isSet
}

func (v *NullablePrismRequestCategoryExport200ResponseArgs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrismRequestCategoryExport200ResponseArgs(val *PrismRequestCategoryExport200ResponseArgs) *NullablePrismRequestCategoryExport200ResponseArgs {
	return &NullablePrismRequestCategoryExport200ResponseArgs{value: val, isSet: true}
}

func (v NullablePrismRequestCategoryExport200ResponseArgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrismRequestCategoryExport200ResponseArgs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


