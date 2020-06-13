/*
 * MIT License
 *
 * Copyright (c) [2020] [Semchenko Aleksandr]
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package kaspersky

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"net/http"
)

//	UserDevicesApi Class Reference
//
//	Interface to unified mobile device management
//
//	List of all members:
type UserDevicesApi service

//UserDevicesApi.DeleteCommand
//Delete a command previously posted to the specified device.
//
//	Parameters:
//	- c_wstrCommandGuid	(string) globally unique command instance id
//	- bForced	(bool) delete command without waiting of real removing from products
//
//	Returns:
//	- (bool) true if the command has been successfully deleted
func (uda *UserDevicesApi) DeleteCommand(ctx context.Context, c_wstrCommandGuid string, bForced bool) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"c_wstrCommandGuid" : "%s", "bForced" : %v }`, c_wstrCommandGuid, bForced))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.DeleteCommand",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.DeleteDevice
//Delete device.
//
//	Parameters:
//	- lDeviceId	(int64) device id
func (uda *UserDevicesApi) DeleteDevice(ctx context.Context, lDeviceId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lDeviceId": %d }`, lDeviceId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.DeleteDevice",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.DeleteEnrollmentPackage
//Delete enrollment package.
//
//	Parameters:
//	- lEnrPkgId	(int64) enrollment package id
func (uda *UserDevicesApi) DeleteEnrollmentPackage(ctx context.Context, lEnrPkgId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lEnrPkgId": %d }`, lEnrPkgId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.DeleteEnrollmentPackage",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GenerateQRCode
//Generates QR code.
//
//Generates QR code from any string
//
//	Parameters:
//	- strInputData	(string) input data
//	- lQRCodeSize	(int64) image size in pixels, for example: value 200 means image size 200x200
//	- lImageFormat	(int64) image format, possible value: 3 = Png
//
//	Returns:
//	- pResult (binary) contains QR-Code image binary data
func (uda *UserDevicesApi) GenerateQRCode(ctx context.Context, strInputData string, lQRCodeSize,
	lImageFormat int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strInputData": "%s", "lQRCodeSize": %d, "lImageFormat": %d }`, strInputData,
		lQRCodeSize, lImageFormat))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GenerateQRCode",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetCommands
//Acquire states of all commands posted to the specified device.
//
//	Parameters:
//	- lDeviceId	(int64) device id
//
//	Returns:
//	- (array) array, each elemen describes the command and its state (see Device commands)
//
//	+---------------------------+--------+---------------------------------------------------------------------------------------------+
//	|            Id             |  Type  |                                         Description                                         |
//	+---------------------------+--------+---------------------------------------------------------------------------------------------+
//	| "KLMDM_CMD_GUID"          | string | Command unique identifier                                                                   |
//	| "KLMDM_CMD_TYPE"          | string | Command type, like "Wipe", "Block" and etc.                                                 |
//	| "KLMDM_CMD_ARG"           | params | Command arguments, it will be passed directly to device. It can be NULL for some commands.  |
//	| "KLMDM_CMD_MDM_PROTOCOLS" | int64  | Bit mask means which protocols will be used to process command.!Not used yet!               |
//	| "KLMDM_CMD_PROCESSFLAGS"  | int64  | Command process flags mean algorithm of command processing. Bit mask values:                |
//	|                           |        | 0x0001: Sequentially - Process commands sequentially, otherwise process command in parallel |
//	|                           |        | 0x0002: ProcessAll - Process commands by all protocol, otherwise by only one                |
//	|                           |        | 0x0004: RecallOnFailure - Try to recall command on failure                                  |
//	| "KLMDM_CMD_STATUS"        | int64  | Command processing state. Enum values:                                                      |
//	|                           |        | 1: Processing                                                                               |
//	|                           |        | 2: Completed                                                                                |
//	|                           |        | 3: Failed                                                                                   |
//	|                           |        | 4: Removing                                                                                 |
//	|                           |        | 5: Removed                                                                                  |
//	|                           |        | 6: RemoveFailed                                                                             |
//	| "KLMDM_CMD_RESULT_INFO"   | string | Additional info about completed or failed command                                           |
//	| "KLMDM_CMD_RESULT_DATA"   | params | Result data of completed command                                                            |
// +---------------------------+--------+---------------------------------------------------------------------------------------------+
func (uda *UserDevicesApi) GetCommands(ctx context.Context, lDeviceId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lDeviceId": %d }`, lDeviceId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetCommands",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetCommandsLibrary
//Acquire library of all available commands.
//
//Acquires list contains commands info reqired to display and launch commands
//
//	Returns:
//	- (array) array, each element is a container describes the command (see Device commands library)
//	+------------------------------+--------+--------------------------------------------------------+
//	|              Id              |  Type  |                      Description                       |
//	+------------------------------+--------+--------------------------------------------------------+
//	| "KLMDM_CMD_FLAG"             | int64  | Bit flag of command                                    |
//	| "KLMDM_CMD_TYPE"             | string | Command type, like "Wipe", "Block" and etc.            |
//	| "KLMDM_CMD_DEF_DISPLAY_NAME" | string | Default display name, localized by the Security Center |
//	+------------------------------+--------+--------------------------------------------------------+
func (uda *UserDevicesApi) GetCommandsLibrary(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetCommandsLibrary",
		nil)
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetDecipheredCommandList
//Calculate commands array according to bit mask of supported commands.
//
//Makes commands array according to bit mask of commands supported by device
//
//	Parameters:
//	- llCommandFlags	(int64) bit mask of commands supported by device
//	- pCommandsLibrary	(array) array of commands from the commands library,
//	each element is a container (paramParams) describing the command (see Device commands library).
//
//	Returns:
//	- (array) array calculated according llCommandFlags,
//	each element is a container (paramParams) describing the command (see Device commands library).
func (uda *UserDevicesApi) GetDecipheredCommandList(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetDecipheredCommandList", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetDevice
//Acquire properties of the specified user device.
//
//	Parameters:
//	- lDeviceId	(int64) device id
//
//	Returns:
//	- (params) device info, container Params contains attributes from List of device attributes
func (uda *UserDevicesApi) GetDevice(ctx context.Context, lDeviceId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lDeviceId": %d }`, lDeviceId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetDevice",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetDevices
//Acquire properties of all registered devices owned by specified user.
//
//	Parameters:
//	- pUserId	(binary) user identifier, binary data as array of bytes; empty means current user
//
//	Returns:
//	- (array) array, each element is a container params contains attributes from List of device attributes
func (uda *UserDevicesApi) GetDevices(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetDevices", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetDevicesExtraData
//Gets additional data from devices such as installed applications, profiles, certificates and etc. !
//
//	Parameters:
//	- pDeviceIds	(array) array of devices integer identity
//	- pCategories	(array) array of categories such as installed applications, profiles,
//	certificates (see Extra Data Categories)
//
//	Returns:
//	- (array) contains additional data from devices like this: Extra Data Example
func (uda *UserDevicesApi) GetDevicesExtraData(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetDevicesExtraData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetEnrollmentPackage
//Get the info of enrollment package created for the device.
//
//	Parameters:
//	- llEnrPkgId	(int64) enrollment package id
//	Returns:
//	- (params) container with enrollment package info (see Device enrollment packages info)
func (uda *UserDevicesApi) GetEnrollmentPackage(ctx context.Context, llEnrPkgId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"llEnrPkgId": %d }`, llEnrPkgId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetEnrollmentPackage",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetEnrollmentPackageFileData
//Get enrollment package file data.
//
//	Parameters:
//	- c_wstrPackageId		(string) enrollment package id extruded from HTTP request
//	- c_wstrPackageFileType	(string) enrollment package file type, example: "iOS4", "iOS5", "Andr4", "WPhone2"
//	- lBuffOffset			(int64) start position
//	- lBuffSize				(int64) number of bytes to read
//
//	Returns:
//	- (binary) contains requested data chunk of enrollment package file
func (uda *UserDevicesApi) GetEnrollmentPackageFileData(ctx context.Context, c_wstrPackageId,
	c_wstrPackageFileType string, lBuffOffset, lBuffSize int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
		"c_wstrPackageId": "%s",
		"c_wstrPackageFileType": "%s",
		"lQRCodeSize": %d, 
		"lImageFormat":%d
	}`, c_wstrPackageId, c_wstrPackageFileType, lBuffOffset, lBuffSize))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetEnrollmentPackageFileData",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetEnrollmentPackageFileInfo
//Get enrollment package file data.
//
//	Parameters:
//	- c_wstrPackageId	(string) enrollment package id extruded from HTTP request
//	- c_wstrPackageFileType	(string) enrollment package file type, example: "iOS4", "iOS5", "Andr4", "WPhone2"
//	- lBuffOffset	(int64) start position
//	- lBuffSize	(int64) number of bytes to read
//
//	Returns:
//	- (binary) contains requested data chunk of enrollment package file
func (uda *UserDevicesApi) GetEnrollmentPackageFileInfo(ctx context.Context, c_wstrPackageId,
	c_wstrUserAgent, c_wstrPackageFileType string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
		"c_wstrPackageId": "%s",
		"c_wstrUserAgent": "%s",
		"c_wstrPackageFileType": "%s"
	}`, c_wstrPackageId, c_wstrUserAgent, c_wstrPackageFileType))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetEnrollmentPackageFileInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetEnrollmentPackages
//Get the list of enrollment packages created for a device.
//
//	Parameters:
//	- pUserId	(binary) user identifier, binary data as array of bytes; empty means current user
//
//	Returns:
//	- (array) array of containers Params with enrollment packages info (see Device enrollment packages info)
func (uda *UserDevicesApi) GetEnrollmentPackages(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetEnrollmentPackages", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetJournalCommandResult
//Returns command result for specific journal record.
//
//	Parameters:
//	- llJrnlId	(int64) journal record id
//
//	Returns:
//	- (params) command result
func (uda *UserDevicesApi) GetJournalCommandResult(ctx context.Context, llJrnlId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"llJrnlId": %d }`, llJrnlId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetJournalCommandResult",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetJournalRecords
//Acquire records from journal.
//
//Acquire records from journal about completed or failed commands posted to device
//
//	Parameters:
//	- lDeviceId	(int64) device id
//
//	Returns:
//	- (array) array, each element contains record about completed or failed commands (see Journal records)
func (uda *UserDevicesApi) GetJournalRecords(ctx context.Context, lDeviceId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lDeviceId": %d }`, lDeviceId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetJournalRecords",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}
	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetJournalRecords2
//Acquire records from journal without command result.
//
//Acquires records from journal about completed or failed commands posted to device without command result.
//To get command result for specific record you can use UserDevicesApi.GetJournalCommandResult
//
//	Parameters:
//	- lDeviceId	(int64) device id
//
//	Returns:
//	- (array) array, each element contains record about completed or failed commands (see Journal records)
func (uda *UserDevicesApi) GetJournalRecords2(ctx context.Context, lDeviceId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lDeviceId": %d }`, lDeviceId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetJournalRecords2",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetLatestDeviceActivityDate
//Acquire latest device activity date.
//
//	Parameters:
//	- lDeviceId	(int64) device id
//
//	Returns:
//	- (datetime) latest device activity date
func (uda *UserDevicesApi) GetLatestDeviceActivityDate(ctx context.Context, lDeviceId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lDeviceId": %d }`, lDeviceId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetLatestDeviceActivityDate",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetMobileAgentSettingStorageData
//Get mobile agent setting storage data.
//
//	Parameters:
//	- lDeviceId	(int64) device id
//	- c_wstrSectionName	(string) required section name
//
//	Returns:
//	- (params) mobile agent setting storage data
func (uda *UserDevicesApi) GetMobileAgentSettingStorageData(ctx context.Context, lDeviceId int64,
	c_wstrSectionName string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lDeviceId": %d, "c_wstrSectionName" : "%s" }`, lDeviceId, c_wstrSectionName))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetMobileAgentSettingStorageData",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}
	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetMultitenancyServerSettings
//Retrieves multitenancy server settings.
//
//	Parameters:
//	- c_wstrMtncServerId	(string) string id of multitenancy server which can be retrieved from method
//	GetMultitenancyServersInfo(value c_szwMtncSrvInfoStrId),
//	if c_wstrMtncServerId is empty then settings will be retrieved from the first available multitenancy server
//
//	Returns:
//	- (params) container which contains current settings of the multitenancy server,
//	now settings available only for KLUMDM.MDMProtocol_IOSMDM (see MDM4IOS multitenancy server settings)
func (uda *UserDevicesApi) GetMultitenancyServerSettings(ctx context.Context, c_wstrMtncServerId string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"c_wstrMtncServerId": "%s" }`, c_wstrMtncServerId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetMultitenancyServerSettings",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetMultitenancyServersInfo
//Retrieves multitenancy servers list.
//
//	Parameters:
//	- nProtocolIds	(int) bit mask means which multitenancy server protocols required
//
//	Returns:
//	- (array) array, each element contains information about multitenancy servers (see Multitenancy servers info)
func (uda *UserDevicesApi) GetMultitenancyServersInfo(ctx context.Context, nProtocolIds int64) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"nProtocolIds" : %d }`, nProtocolIds))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetMultitenancyServersInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GetSafeBrowserAutoinstallFlag
//Returns flag which means install or don't install SafeBrowser automatically when device connects first time.
//
//	Returns:
//	- (bool) flag which means install or don't install SafeBrowser automatically
func (uda *UserDevicesApi) GetSafeBrowserAutoinstallFlag(ctx context.Context) (*PxgValBool, []byte,
	error) {
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetSafeBrowserAutoinstallFlag",
		nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := uda.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

//UserDevicesApi.GetSyncInfo
//Retrieves group synchronization info for UMDM policy.
//
//	Parameters:
//	- nGroupId	(int64) Group identifier.
//	- nGSyncId	[int64] Synchronization identifier.
//	- pFields	(array) array of requested attributes
//
//	Returns:
//	- (params) Group synchronization info for requested attributes
//
//	Exceptions:
//	- STDE_NOTFOUND	if group with requested attributes was not found
func (uda *UserDevicesApi) GetSyncInfo(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetSyncInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.GlueDevices
//Glue information on a device got from different sources
//
//	Parameters:
//	- lDevice1Id	(int64) first device id
//	- lDevice2Id	(int64) second device id
//
//	Deprecated:
func (uda *UserDevicesApi) GlueDevices(ctx context.Context, lDevice1Id, lDevice2Id int64) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"lDevice1Id": %d, "lDevice2Id" : %d }`, lDevice1Id, lDevice2Id))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GlueDevices",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.PostCommand
//Post a command to the specified device.
//
//Note, you can't sent more then one command with same type to the specified device
//while command with same type is not completed, instead you may call method RecallCommand
//
//	Parameters:
//	- lDeviceId			(int64) device id
//	- c_wstrCommandGuid	(string) globally unique command instance id.
//	- c_wstrCommandType	(string) command type, you can retrieve it by calling GetCommandsLibrary
//	- pArguments		(params) command arguments
//	- lMdmProtocols		(int64) bit mask means which protocols will be used to process command (DEPRECATED, just set 0)
//	- lProcessFlags		(int64) command process flags mean algorithm to process command (DEPRECATED, just set 0)
func (uda *UserDevicesApi) PostCommand(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.PostCommand", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.RecallCommand
//Recall command previously posted to the specified device.
//
//	Parameters:
//	- c_wstrCommandGuid	(string) globally unique command instance id
//
//	Returns:
//	- (bool) true if the command has not been delivered to the device yet
func (uda *UserDevicesApi) RecallCommand(ctx context.Context, c_wstrCommandGuid string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"c_wstrCommandGuid": "%s"}`, c_wstrCommandGuid))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.RecallCommand",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.SetMultitenancyServerSettings
//Set multitenancy server settings.
//
//	Parameters:
//	- c_wstrMtncServerId	(string) string id of multitenancy server which can be retrieved from method
//	GetMultitenancyServersInfo(value c_szwMtncSrvInfoStrId),
//	if c_wstrMtncServerId is empty then settings will be retrieved from the first available multitenancy server
//	- pSettings	(params) contains settings to multitenancy server,
//	now settings available only for KLUMDM::MDMProtocol_IOSMDM (see MDM4IOS multitenancy server settings)
func (uda *UserDevicesApi) SetMultitenancyServerSettings(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.SetMultitenancyServerSettings", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.SetSafeBrowserAutoinstallFlag
//Set flag which means install or don't install SafeBrowser automatically when device connects first time.
//
//	Parameters:
//	- bInstall	(bool) flag means install or don't install SafeBrowser automatically
func (uda *UserDevicesApi) SetSafeBrowserAutoinstallFlag(ctx context.Context, bInstall bool) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"bInstall": %v}`, bInstall))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.SetSafeBrowserAutoinstallFlag",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.SspLoginAllowed
//Check user permission to login to SSP.
//
//	Exceptions:
//	- STDE_NOACCESS	if login to SSP is not allowed
func (uda *UserDevicesApi) SspLoginAllowed(ctx context.Context) ([]byte,
	error) {
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.SspLoginAllowed",
		nil)
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//UserDevicesApi.UpdateDevice
//Modify properties of the specified user device.
//
//	Parameters:
//	- lDeviceId	(int64) device id
//	- pDevice	(params) device info, container Params contains attributes: List of device attributes
func (uda *UserDevicesApi) UpdateDevice(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.UpdateDevice", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}
