# Go API client for kandji

<html><head></head><body><h1 id=&quot;welcome-to-the-kandji-api-documentation&quot;>Welcome to the Kandji API Documentation</h1>
<p>You can find your API URL in Settings &gt; Access. The API URL will follow the below formats.</p>
<ul>
<li><p>US - <code>https://SubDomain.api.kandji.io</code></p>
</li>
<li><p>EU - <code>https://SubDomain.api.eu.kandji.io</code></p>
</li>
</ul>
<p>For information on how to obtain an API token, please refer to the following support article.</p>
<p><a href=&quot;https://support.kandji.io/api&quot;>https://support.kandji.io/api</a></p>
<h4 id=&quot;rate-limit&quot;>Rate Limit</h4>
<p>The Kandji API currently has an API rate limit of 10,000 requests per hour per customer.</p>
<h4 id=&quot;request-methods&quot;>Request Methods</h4>
<p>HTTP request methods supported by the Kandji API.</p>
<div class=&quot;click-to-expand-wrapper is-table-wrapper&quot;><table>
<thead>
<tr>
<th>Method</th>
<th>Definition</th>
</tr>
</thead>
<tbody>
<tr>
<td>GET</td>
<td>The <code>GET</code> method requests a representation of the specified resource.</td>
</tr>
<tr>
<td>POST</td>
<td>The <code>POST</code> method submits an entity to the specified resource.</td>
</tr>
<tr>
<td>PATCH</td>
<td>The <code>PATCH</code> method applies partial modifications to a resource.</td>
</tr>
<tr>
<td>DELETE</td>
<td>The <code>DELETE</code> method deletes the specified resource.</td>
</tr>
</tbody>
</table>
</div><h4 id=&quot;response-codes&quot;>Response codes</h4>
<p>Not all response codes apply to every endpoint.</p>
<div class=&quot;click-to-expand-wrapper is-table-wrapper&quot;><table>
<thead>
<tr>
<th>Code</th>
<th>Response</th>
</tr>
</thead>
<tbody>
<tr>
<td>200</td>
<td>OK</td>
</tr>
<tr>
<td>201</td>
<td>Created</td>
</tr>
<tr>
<td>204</td>
<td>No content</td>
</tr>
<tr>
<td></td>
<td>Typical response when sending the DELETE method.</td>
</tr>
<tr>
<td>400</td>
<td>Bad Request</td>
</tr>
<tr>
<td></td>
<td>&quot;Command already running&quot; - The command may already be running in a <em>Pending</em> state waiting on the device.</td>
</tr>
<tr>
<td></td>
<td>&quot;Command is not allowed for current device&quot; - The command may not be compatible with the target device.</td>
</tr>
<tr>
<td></td>
<td>&quot;JSON parse error - Expecting ',' delimiter: line 3 column 2 (char 65)&quot;</td>
</tr>
<tr>
<td>401</td>
<td>Unauthorized</td>
</tr>
<tr>
<td></td>
<td>This error can occur if the token is incorrect, was revoked, or the token has expired.</td>
</tr>
<tr>
<td>403</td>
<td>Forbidden</td>
</tr>
<tr>
<td></td>
<td>The request was understood but cannot be authorized.</td>
</tr>
<tr>
<td>404</td>
<td>Not found</td>
</tr>
<tr>
<td></td>
<td>Unable to locate the resource in the Kandji tenant.</td>
</tr>
<tr>
<td>415</td>
<td>Unsupported Media Type</td>
</tr>
<tr>
<td></td>
<td>The request contains a media type which the server or resource does not support.</td>
</tr>
<tr>
<td>500</td>
<td>Internal server error</td>
</tr>
<tr>
<td>503</td>
<td>Service unavailable</td>
</tr>
<tr>
<td></td>
<td>This error can occur if a file upload is still being processed via the custom apps API.</td>
</tr>
</tbody>
</table>
</div><h4 id=&quot;data-structure&quot;>Data structure</h4>
<p>The API returns all structured responses in JSON schema format.</p>
<h4 id=&quot;examples&quot;>Examples</h4>
<p>Code examples using the API can be found in the Kandji support <a href=&quot;https://github.com/kandji-inc/support/tree/main/api-tools&quot;>GitHub</a>.</p>
</body></html>

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Generator version: 7.11.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://github.com/MScottBlake/kandji-openapi](https://github.com/MScottBlake/kandji-openapi)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import kandji "github.com/MScottBlake/kandji-go-sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `kandji.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), kandji.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `kandji.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), kandji.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `kandji.ContextOperationServerIndices` and `kandji.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), kandji.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), kandji.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AutomatedDeviceEnrollmentIntegrationsAPI* | [**CreateAdeIntegration**](docs/AutomatedDeviceEnrollmentIntegrationsAPI.md#createadeintegration) | **Post** /api/v1/integrations/apple/ade/ | Create ADE integration
*AutomatedDeviceEnrollmentIntegrationsAPI* | [**DeleteAdeIntegration**](docs/AutomatedDeviceEnrollmentIntegrationsAPI.md#deleteadeintegration) | **Delete** /api/v1/integrations/apple/ade/{ade_token_id} | Delete ADE integration
*AutomatedDeviceEnrollmentIntegrationsAPI* | [**DownloadAdePublicKey**](docs/AutomatedDeviceEnrollmentIntegrationsAPI.md#downloadadepublickey) | **Get** /api/v1/integrations/apple/ade/public_key/ | Download ADE public key
*AutomatedDeviceEnrollmentIntegrationsAPI* | [**GetAdeDevice**](docs/AutomatedDeviceEnrollmentIntegrationsAPI.md#getadedevice) | **Get** /api/v1/integrations/apple/ade/devices/{device_id} | Get ADE device
*AutomatedDeviceEnrollmentIntegrationsAPI* | [**GetAdeIntegration**](docs/AutomatedDeviceEnrollmentIntegrationsAPI.md#getadeintegration) | **Get** /api/v1/integrations/apple/ade/{ade_token_id} | Get ADE integration
*AutomatedDeviceEnrollmentIntegrationsAPI* | [**ListAdeDevices**](docs/AutomatedDeviceEnrollmentIntegrationsAPI.md#listadedevices) | **Get** /api/v1/integrations/apple/ade/devices | List ADE devices
*AutomatedDeviceEnrollmentIntegrationsAPI* | [**ListAdeIntegrations**](docs/AutomatedDeviceEnrollmentIntegrationsAPI.md#listadeintegrations) | **Get** /api/v1/integrations/apple/ade | List ADE integrations
*AutomatedDeviceEnrollmentIntegrationsAPI* | [**ListDevicesAssociatedToAdeToken**](docs/AutomatedDeviceEnrollmentIntegrationsAPI.md#listdevicesassociatedtoadetoken) | **Get** /api/v1/integrations/apple/ade/{ade_token_id}/devices | List devices associated to ADE token
*AutomatedDeviceEnrollmentIntegrationsAPI* | [**RenewAdeIntegration**](docs/AutomatedDeviceEnrollmentIntegrationsAPI.md#renewadeintegration) | **Post** /api/v1/integrations/apple/ade/{ade_token_id}/renew | Renew ADE integration
*AutomatedDeviceEnrollmentIntegrationsAPI* | [**UpdateAdeDevice**](docs/AutomatedDeviceEnrollmentIntegrationsAPI.md#updateadedevice) | **Patch** /api/v1/integrations/apple/ade/devices/{device_id} | Update ADE device
*AutomatedDeviceEnrollmentIntegrationsAPI* | [**UpdateAdeIntegration**](docs/AutomatedDeviceEnrollmentIntegrationsAPI.md#updateadeintegration) | **Patch** /api/v1/integrations/apple/ade/{ade_token_id} | Update ADE integration
*BlueprintsAPI* | [**AssignLibraryItem**](docs/BlueprintsAPI.md#assignlibraryitem) | **Post** /api/v1/blueprints/{blueprint_id}/assign-library-item | Assign Library Item
*BlueprintsAPI* | [**CreateBlueprint**](docs/BlueprintsAPI.md#createblueprint) | **Post** /api/v1/blueprints | Create Blueprint
*BlueprintsAPI* | [**DeleteBlueprint**](docs/BlueprintsAPI.md#deleteblueprint) | **Delete** /api/v1/blueprints/{blueprint_id} | Delete Blueprint
*BlueprintsAPI* | [**GetBlueprint**](docs/BlueprintsAPI.md#getblueprint) | **Get** /api/v1/blueprints/{blueprint_id} | Get Blueprint
*BlueprintsAPI* | [**GetBlueprintTemplates**](docs/BlueprintsAPI.md#getblueprinttemplates) | **Get** /api/v1/blueprints/templates/ | Get Blueprint Templates
*BlueprintsAPI* | [**GetManualEnrollmentProfile**](docs/BlueprintsAPI.md#getmanualenrollmentprofile) | **Get** /api/v1/blueprints/{blueprint_id}/ota-enrollment-profile | Get Manual Enrollment Profile
*BlueprintsAPI* | [**ListBlueprints**](docs/BlueprintsAPI.md#listblueprints) | **Get** /api/v1/blueprints | List Blueprints
*BlueprintsAPI* | [**ListLibraryItems**](docs/BlueprintsAPI.md#listlibraryitems) | **Get** /api/v1/blueprints/{blueprint_id}/list-library-items | List Library Items
*BlueprintsAPI* | [**RemoveLibraryItem**](docs/BlueprintsAPI.md#removelibraryitem) | **Post** /api/v1/blueprints/{blueprint_id}/remove-library-item | Remove Library Item
*BlueprintsAPI* | [**UpdateBlueprint**](docs/BlueprintsAPI.md#updateblueprint) | **Patch** /api/v1/blueprints/{blueprint_id} | Update Blueprint
*DeviceActionsAPI* | [**ClearPasscode**](docs/DeviceActionsAPI.md#clearpasscode) | **Post** /api/v1/devices/{device_id}/action/clearpasscode | Clear Passcode
*DeviceActionsAPI* | [**DeleteDevice**](docs/DeviceActionsAPI.md#deletedevice) | **Delete** /api/v1/devices/{device_id} | Delete Device
*DeviceActionsAPI* | [**DeleteUser**](docs/DeviceActionsAPI.md#deleteuser) | **Post** /api/v1/devices/{device_id}/action/deleteuser | Delete User
*DeviceActionsAPI* | [**EraseDevice**](docs/DeviceActionsAPI.md#erasedevice) | **Post** /api/v1/devices/{device_id}/action/erase | Erase Device
*DeviceActionsAPI* | [**GetDeviceCommands**](docs/DeviceActionsAPI.md#getdevicecommands) | **Get** /api/v1/devices/{device_id}/commands | Get Device Commands
*DeviceActionsAPI* | [**LockDevice**](docs/DeviceActionsAPI.md#lockdevice) | **Post** /api/v1/devices/{device_id}/action/lock | Lock Device
*DeviceActionsAPI* | [**ReinstallAgent**](docs/DeviceActionsAPI.md#reinstallagent) | **Post** /api/v1/devices/{device_id}/action/reinstallagent | Reinstall Agent
*DeviceActionsAPI* | [**RemoteDesktop**](docs/DeviceActionsAPI.md#remotedesktop) | **Post** /api/v1/devices/{device_id}/action/remotedesktop | Remote Desktop
*DeviceActionsAPI* | [**RenewMdmProfile**](docs/DeviceActionsAPI.md#renewmdmprofile) | **Post** /api/v1/devices/{device_id}/action/renewmdmprofile | Renew MDM Profile
*DeviceActionsAPI* | [**RestartDevice**](docs/DeviceActionsAPI.md#restartdevice) | **Post** /api/v1/devices/{device_id}/action/restart | Restart Device
*DeviceActionsAPI* | [**SendBlankpush**](docs/DeviceActionsAPI.md#sendblankpush) | **Post** /api/v1/devices/{device_id}/action/blankpush | Send Blankpush
*DeviceActionsAPI* | [**SetName**](docs/DeviceActionsAPI.md#setname) | **Post** /api/v1/devices/{device_id}/action/setname | Set Name
*DeviceActionsAPI* | [**Shutdown**](docs/DeviceActionsAPI.md#shutdown) | **Post** /api/v1/devices/{device_id}/action/shutdown | Shutdown
*DeviceActionsAPI* | [**UnlockAccount**](docs/DeviceActionsAPI.md#unlockaccount) | **Post** /api/v1/devices/{device_id}/action/unlockaccount | Unlock Account
*DeviceActionsAPI* | [**UpdateInventory**](docs/DeviceActionsAPI.md#updateinventory) | **Post** /api/v1/devices/{device_id}/action/updateinventory | Update Inventory
*DeviceInformationAPI* | [**CancelLostMode**](docs/DeviceInformationAPI.md#cancellostmode) | **Delete** /api/v1/devices/{device_id}/details/lostmode | Cancel Lost Mode
*DeviceInformationAPI* | [**GetDevice**](docs/DeviceInformationAPI.md#getdevice) | **Get** /api/v1/devices/{device_id} | Get Device
*DeviceInformationAPI* | [**GetDeviceActivity**](docs/DeviceInformationAPI.md#getdeviceactivity) | **Get** /api/v1/devices/{device_id}/activity | Get Device Activity
*DeviceInformationAPI* | [**GetDeviceApps**](docs/DeviceInformationAPI.md#getdeviceapps) | **Get** /api/v1/devices/{device_id}/apps | Get Device Apps
*DeviceInformationAPI* | [**GetDeviceDetails**](docs/DeviceInformationAPI.md#getdevicedetails) | **Get** /api/v1/devices/{device_id}/details | Get Device Details
*DeviceInformationAPI* | [**GetDeviceLibraryItems**](docs/DeviceInformationAPI.md#getdevicelibraryitems) | **Get** /api/v1/devices/{device_id}/library-items | Get Device Library Items
*DeviceInformationAPI* | [**GetDeviceLostModeDetails**](docs/DeviceInformationAPI.md#getdevicelostmodedetails) | **Get** /api/v1/devices/{device_id}/details/lostmode | Get Device Lost Mode details
*DeviceInformationAPI* | [**GetDeviceParameters**](docs/DeviceInformationAPI.md#getdeviceparameters) | **Get** /api/v1/devices/{device_id}/parameters | Get Device Parameters
*DeviceInformationAPI* | [**GetDeviceStatus**](docs/DeviceInformationAPI.md#getdevicestatus) | **Get** /api/v1/devices/{device_id}/status | Get Device Status
*DeviceInformationAPI* | [**ListDevices**](docs/DeviceInformationAPI.md#listdevices) | **Get** /api/v1/devices | List Devices
*DeviceInformationAPI* | [**UpdateDevice**](docs/DeviceInformationAPI.md#updatedevice) | **Patch** /api/v1/devices/{device_id} | Update Device
*DeviceSecretsAPI* | [**GetActivationLockBypassCode**](docs/DeviceSecretsAPI.md#getactivationlockbypasscode) | **Get** /api/v1/devices/{device_id}/secrets/bypasscode | Get Activation Lock Bypass Code
*DeviceSecretsAPI* | [**GetFilevaultRecoveryKey**](docs/DeviceSecretsAPI.md#getfilevaultrecoverykey) | **Get** /api/v1/devices/{device_id}/secrets/filevaultkey | Get FileVault Recovery Key
*DeviceSecretsAPI* | [**GetRecoveryLockPassword**](docs/DeviceSecretsAPI.md#getrecoverylockpassword) | **Get** /api/v1/devices/{device_id}/secrets/recoverypassword | Get Recovery Lock Password
*DeviceSecretsAPI* | [**GetUnlockPin**](docs/DeviceSecretsAPI.md#getunlockpin) | **Get** /api/v1/devices/{device_id}/secrets/unlockpin | Get Unlock Pin
*LibraryItemsAPI* | [**GetLibraryItemActivity**](docs/LibraryItemsAPI.md#getlibraryitemactivity) | **Get** /api/v1/library/library-items/{library_item_id}/activity | Get Library Item Activity
*LibraryItemsAPI* | [**GetLibraryItemStatuses**](docs/LibraryItemsAPI.md#getlibraryitemstatuses) | **Get** /api/v1/library/library-items/{library_item_id}/status | Get Library Item Statuses
*PrismAPI* | [**ActivationLock**](docs/PrismAPI.md#activationlock) | **Get** /api/v1/prism/activation_lock | Activation lock
*PrismAPI* | [**ApplicationFirewall**](docs/PrismAPI.md#applicationfirewall) | **Get** /api/v1/prism/application_firewall | Application firewall
*PrismAPI* | [**Applications**](docs/PrismAPI.md#applications) | **Get** /api/v1/prism/apps | Applications
*PrismAPI* | [**Certificates**](docs/PrismAPI.md#certificates) | **Get** /api/v1/prism/certificates | Certificates
*PrismAPI* | [**Count**](docs/PrismAPI.md#count) | **Get** /api/v1/prism/count | Count
*PrismAPI* | [**DesktopAndScreensaver**](docs/PrismAPI.md#desktopandscreensaver) | **Get** /api/v1/prism/desktop_and_screensaver | Desktop and Screensaver
*PrismAPI* | [**DeviceInformation**](docs/PrismAPI.md#deviceinformation) | **Get** /api/v1/prism/device_information | Device information
*PrismAPI* | [**Filevault**](docs/PrismAPI.md#filevault) | **Get** /api/v1/prism/filevault | FileVault
*PrismAPI* | [**GatekeeperAndXprotect**](docs/PrismAPI.md#gatekeeperandxprotect) | **Get** /api/v1/prism/gatekeeper_and_xprotect | Gatekeeper and XProtect
*PrismAPI* | [**GetCategoryExport**](docs/PrismAPI.md#getcategoryexport) | **Get** /api/v1/prism/export/{export_id} | Get category export
*PrismAPI* | [**InstalledProfiles**](docs/PrismAPI.md#installedprofiles) | **Get** /api/v1/prism/installed_profiles | Installed profiles
*PrismAPI* | [**KernelExtensions**](docs/PrismAPI.md#kernelextensions) | **Get** /api/v1/prism/kernel_extensions | Kernel Extensions
*PrismAPI* | [**LaunchAgentsAndDaemons**](docs/PrismAPI.md#launchagentsanddaemons) | **Get** /api/v1/prism/launch_agents_and_daemons | Launch Agents and Daemons
*PrismAPI* | [**LocalUsers**](docs/PrismAPI.md#localusers) | **Get** /api/v1/prism/local_users | Local users
*PrismAPI* | [**RequestCategoryExport**](docs/PrismAPI.md#requestcategoryexport) | **Post** /api/v1/prism/export | Request category export
*PrismAPI* | [**StartupSettings**](docs/PrismAPI.md#startupsettings) | **Get** /api/v1/prism/startup_settings | Startup settings
*PrismAPI* | [**SystemExtensions**](docs/PrismAPI.md#systemextensions) | **Get** /api/v1/prism/system_extensions | System Extensions
*PrismAPI* | [**TransparencyDatabase**](docs/PrismAPI.md#transparencydatabase) | **Get** /api/v1/prism/transparency_database | Transparency database
*SettingsAPI* | [**Licensing**](docs/SettingsAPI.md#licensing) | **Get** /api/v1/settings/licensing | Licensing
*TagsAPI* | [**CreateTag**](docs/TagsAPI.md#createtag) | **Post** /api/v1/tags | Create Tag
*TagsAPI* | [**DeleteTag**](docs/TagsAPI.md#deletetag) | **Delete** /api/v1/tags/{tag_id} | Delete Tag
*TagsAPI* | [**GetTags**](docs/TagsAPI.md#gettags) | **Get** /api/v1/tags | Get Tags
*TagsAPI* | [**UpdateTag**](docs/TagsAPI.md#updatetag) | **Patch** /api/v1/tags/{tag_id} | Update Tag
*ThreatsAPI* | [**GetThreatDetails**](docs/ThreatsAPI.md#getthreatdetails) | **Get** /api/v1/threat-details | Get Threat Details
*UsersAPI* | [**DeleteUser**](docs/UsersAPI.md#deleteuser) | **Delete** /api/v1/users/{user_id} | Delete User
*UsersAPI* | [**GetUser**](docs/UsersAPI.md#getuser) | **Get** /api/v1/users/{user_id} | Get User
*UsersAPI* | [**ListUsers**](docs/UsersAPI.md#listusers) | **Get** /api/v1/users | List Users


## Documentation For Models

 - [BadCategoryRequest400Response](docs/BadCategoryRequest400Response.md)
 - [ClearAllTags200Response](docs/ClearAllTags200Response.md)
 - [CreateIntegration200Response](docs/CreateIntegration200Response.md)
 - [CreateIntegration200ResponseBlueprint](docs/CreateIntegration200ResponseBlueprint.md)
 - [CreateIntegration200ResponseDefaults](docs/CreateIntegration200ResponseDefaults.md)
 - [CreateIntegration200ResponseDeviceCounts](docs/CreateIntegration200ResponseDeviceCounts.md)
 - [DeviceInfoForAllIpads200Response](docs/DeviceInfoForAllIpads200Response.md)
 - [DeviceInfoForAllIpads200ResponseArgs](docs/DeviceInfoForAllIpads200ResponseArgs.md)
 - [EdrStatus200Response](docs/EdrStatus200Response.md)
 - [EdrStatus200Response1](docs/EdrStatus200Response1.md)
 - [ExampleExportStatusCheck200Response](docs/ExampleExportStatusCheck200Response.md)
 - [GetDeviceActivity200Response](docs/GetDeviceActivity200Response.md)
 - [GetDeviceActivity200ResponseActivity](docs/GetDeviceActivity200ResponseActivity.md)
 - [GetDeviceCommands200Response](docs/GetDeviceCommands200Response.md)
 - [GetDeviceLostModeDetails200Response](docs/GetDeviceLostModeDetails200Response.md)
 - [GetDeviceParameters200Response](docs/GetDeviceParameters200Response.md)
 - [GetDeviceStatus200Response](docs/GetDeviceStatus200Response.md)
 - [GetDevicesInABlueprintSortedBySerialNumber200Response](docs/GetDevicesInABlueprintSortedBySerialNumber200Response.md)
 - [GetIpad200Response](docs/GetIpad200Response.md)
 - [GetIpad200ResponseUser](docs/GetIpad200ResponseUser.md)
 - [InvalidUuid400Response](docs/InvalidUuid400Response.md)
 - [IpadApps200Response](docs/IpadApps200Response.md)
 - [IphoneOrIpadInLostMode200Response](docs/IphoneOrIpadInLostMode200Response.md)
 - [IphoneOrIpadInLostMode200ResponseActivationLock](docs/IphoneOrIpadInLostMode200ResponseActivationLock.md)
 - [IphoneOrIpadInLostMode200ResponseAppleBusinessManager](docs/IphoneOrIpadInLostMode200ResponseAppleBusinessManager.md)
 - [IphoneOrIpadInLostMode200ResponseAutomatedDeviceEnrollment](docs/IphoneOrIpadInLostMode200ResponseAutomatedDeviceEnrollment.md)
 - [IphoneOrIpadInLostMode200ResponseCellular](docs/IphoneOrIpadInLostMode200ResponseCellular.md)
 - [IphoneOrIpadInLostMode200ResponseFilevault](docs/IphoneOrIpadInLostMode200ResponseFilevault.md)
 - [IphoneOrIpadInLostMode200ResponseGeneral](docs/IphoneOrIpadInLostMode200ResponseGeneral.md)
 - [IphoneOrIpadInLostMode200ResponseHardwareOverview](docs/IphoneOrIpadInLostMode200ResponseHardwareOverview.md)
 - [IphoneOrIpadInLostMode200ResponseKandjiAgent](docs/IphoneOrIpadInLostMode200ResponseKandjiAgent.md)
 - [IphoneOrIpadInLostMode200ResponseLostMode](docs/IphoneOrIpadInLostMode200ResponseLostMode.md)
 - [IphoneOrIpadInLostMode200ResponseLostModeLastLocation](docs/IphoneOrIpadInLostMode200ResponseLostModeLastLocation.md)
 - [IphoneOrIpadInLostMode200ResponseMdm](docs/IphoneOrIpadInLostMode200ResponseMdm.md)
 - [IphoneOrIpadInLostMode200ResponseRecoveryInformation](docs/IphoneOrIpadInLostMode200ResponseRecoveryInformation.md)
 - [IphoneOrIpadInLostMode200ResponseSecurityInformation](docs/IphoneOrIpadInLostMode200ResponseSecurityInformation.md)
 - [IphoneOrIpadInLostMode200ResponseUsers](docs/IphoneOrIpadInLostMode200ResponseUsers.md)
 - [ListAssociatedDevicesNullMdmDevice200Response](docs/ListAssociatedDevicesNullMdmDevice200Response.md)
 - [StartupSettings200Response](docs/StartupSettings200Response.md)
 - [Success200Response](docs/Success200Response.md)
 - [Success200Response1](docs/Success200Response1.md)
 - [Success200Response10](docs/Success200Response10.md)
 - [Success200Response11](docs/Success200Response11.md)
 - [Success200Response12](docs/Success200Response12.md)
 - [Success200Response13](docs/Success200Response13.md)
 - [Success200Response14](docs/Success200Response14.md)
 - [Success200Response15](docs/Success200Response15.md)
 - [Success200Response16](docs/Success200Response16.md)
 - [Success200Response16Integration](docs/Success200Response16Integration.md)
 - [Success200Response17](docs/Success200Response17.md)
 - [Success200Response17Counts](docs/Success200Response17Counts.md)
 - [Success200Response17Limits](docs/Success200Response17Limits.md)
 - [Success200Response17LimitsMaxDevicesPerPlatform](docs/Success200Response17LimitsMaxDevicesPerPlatform.md)
 - [Success200Response1DepAccount](docs/Success200Response1DepAccount.md)
 - [Success200Response1MdmDevice](docs/Success200Response1MdmDevice.md)
 - [Success200Response2](docs/Success200Response2.md)
 - [Success200Response3](docs/Success200Response3.md)
 - [Success200Response4](docs/Success200Response4.md)
 - [Success200Response5](docs/Success200Response5.md)
 - [Success200Response6](docs/Success200Response6.md)
 - [Success200Response7](docs/Success200Response7.md)
 - [Success200Response8](docs/Success200Response8.md)
 - [Success200Response9](docs/Success200Response9.md)
 - [Success201Response](docs/Success201Response.md)
 - [Success201Response1](docs/Success201Response1.md)
 - [Success201ResponseEnrollmentCode](docs/Success201ResponseEnrollmentCode.md)
 - [SuccessNoKernelExtensions200Response](docs/SuccessNoKernelExtensions200Response.md)
 - [SystemExtensions200Response](docs/SystemExtensions200Response.md)
 - [TransparencyDatabase200Response](docs/TransparencyDatabase200Response.md)
 - [UpdateUserAssignment200Response](docs/UpdateUserAssignment200Response.md)
 - [UserAssignedToDevice400Response](docs/UserAssignedToDevice400Response.md)
 - [UsingTermParam200Response](docs/UsingTermParam200Response.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### bearer

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), kandji.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

mitchelsblake@gmail.com

