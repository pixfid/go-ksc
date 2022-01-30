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

// UserDevicesApi Interface to unified mobile device management.
type UserDevicesApi service

// DeleteCommand
// Delete a command previously posted to the specified device.
func (uda *UserDevicesApi) DeleteCommand(ctx context.Context, c_wstrCommandGuid string, bForced bool) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"c_wstrCommandGuid" : "%s", "bForced" : %v }`, c_wstrCommandGuid, bForced))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.DeleteCommand",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// DeleteDevice
// Delete device.
func (uda *UserDevicesApi) DeleteDevice(ctx context.Context, lDeviceId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lDeviceId": %d }`, lDeviceId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.DeleteDevice",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// DeleteEnrollmentPackage
// Delete enrollment package.
func (uda *UserDevicesApi) DeleteEnrollmentPackage(ctx context.Context, lEnrPkgId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lEnrPkgId": %d }`, lEnrPkgId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.DeleteEnrollmentPackage",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// GenerateQRCode
// Generates QR code from any string
func (uda *UserDevicesApi) GenerateQRCode(ctx context.Context, strInputData string, lQRCodeSize, lImageFormat int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strInputData": "%s", "lQRCodeSize": %d, "lImageFormat": %d }`, strInputData, lQRCodeSize, lImageFormat))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GenerateQRCode",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// GetCommands
// Acquire states of all commands posted to the specified device.
func (uda *UserDevicesApi) GetCommands(ctx context.Context, lDeviceId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lDeviceId": %d }`, lDeviceId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetCommands",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// CommandsLibrary struct
type CommandsLibrary struct {
	PxgRetVal []PxgRetVal `json:"PxgRetVal"`
}

type PxgRetVal struct {
	Type  *string      `json:"type,omitempty"`
	Value *CmdLibValue `json:"value,omitempty"`
}

type CmdLibValue struct {
	KlmdmCmdDefDisplayName *string       `json:"KLMDM_CMD_DEF_DISPLAY_NAME,omitempty"`
	KlmdmCmdFlag           *KlmdmCmdFlag `json:"KLMDM_CMD_FLAG,omitempty"`
	KlmdmCmdType           *string       `json:"KLMDM_CMD_TYPE,omitempty"`
}

type KlmdmCmdFlag struct {
	Type  *string `json:"type,omitempty"`
	Value *int64  `json:"value,omitempty"`
}

// GetCommandsLibrary
// Acquires list contains commands info reqired to display and launch commands
func (uda *UserDevicesApi) GetCommandsLibrary(ctx context.Context) (*CommandsLibrary, error) {
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetCommandsLibrary", nil)
	if err != nil {
		return nil, err
	}

	commandsLibrary := new(CommandsLibrary)
	_, err = uda.client.Request(ctx, request, &commandsLibrary)
	return commandsLibrary, err
}

// GetDecipheredCommandList
// Calculate commands array according to bit mask of supported commands.
// Makes commands array according to bit mask of commands supported by device
func (uda *UserDevicesApi) GetDecipheredCommandList(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetDecipheredCommandList", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// GetDevice Acquire properties of the specified user device.
//	Deprecated: Previously used for the SSP. Now, use the SrvView UmdmDevices instead.
func (uda *UserDevicesApi) GetDevice(ctx context.Context, lDeviceId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lDeviceId": %d }`, lDeviceId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetDevice",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// UserID struct
type UserID struct {
	PUserID string `json:"pUserId,omitempty"`
}

// GetDevices Acquire properties of all registered devices owned by specified user.
//	Deprecated: Previously used for the SSP. Now, use the SrvView UmdmDevices instead.
func (uda *UserDevicesApi) GetDevices(ctx context.Context, params UserID) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetDevices", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// GetDevicesExtraData
// Gets additional data from devices such as installed applications, profiles, certificates and etc. !
func (uda *UserDevicesApi) GetDevicesExtraData(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetDevicesExtraData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// GetEnrollmentPackage Get the info of enrollment package created for the device.
func (uda *UserDevicesApi) GetEnrollmentPackage(ctx context.Context, llEnrPkgId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"llEnrPkgId": %d }`, llEnrPkgId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetEnrollmentPackage",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// GetEnrollmentPackageFileData
// Get enrollment package file data.
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

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// GetEnrollmentPackageFileInfo
// Get enrollment package file data.
func (uda *UserDevicesApi) GetEnrollmentPackageFileInfo(ctx context.Context, c_wstrPackageId, c_wstrUserAgent, c_wstrPackageFileType string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"c_wstrPackageId": "%s", "c_wstrUserAgent": "%s", "c_wstrPackageFileType": "%s"	}`, c_wstrPackageId, c_wstrUserAgent, c_wstrPackageFileType))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetEnrollmentPackageFileInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// GetEnrollmentPackages
// Get the list of enrollment packages created for a device.
func (uda *UserDevicesApi) GetEnrollmentPackages(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetEnrollmentPackages", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// GetJournalCommandResult
// Returns command result for specific journal record.
func (uda *UserDevicesApi) GetJournalCommandResult(ctx context.Context, llJrnlId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"llJrnlId": %d }`, llJrnlId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetJournalCommandResult",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// GetJournalRecords
// Acquire records from journal about completed or failed commands posted to device
func (uda *UserDevicesApi) GetJournalRecords(ctx context.Context, lDeviceId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lDeviceId": %d }`, lDeviceId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetJournalRecords",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}
	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// GetJournalRecords2
// Acquires records from journal about completed or failed commands posted to device without command result.
// To get command result for specific record you can use UserDevicesApi.GetJournalCommandResult
func (uda *UserDevicesApi) GetJournalRecords2(ctx context.Context, lDeviceId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lDeviceId": %d }`, lDeviceId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetJournalRecords2",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// GetLatestDeviceActivityDate Acquire latest device activity date.
//	Deprecated: Previously used for the SSP. Now, use the SrvView UmdmDevices instead.
func (uda *UserDevicesApi) GetLatestDeviceActivityDate(ctx context.Context, lDeviceId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lDeviceId": %d }`, lDeviceId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetLatestDeviceActivityDate",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// GetMobileAgentSettingStorageData Get mobile agent setting storage data.
func (uda *UserDevicesApi) GetMobileAgentSettingStorageData(ctx context.Context, lDeviceId int64,
	c_wstrSectionName string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lDeviceId": %d, "c_wstrSectionName" : "%s" }`, lDeviceId, c_wstrSectionName))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetMobileAgentSettingStorageData",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}
	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// GetMultitenancyServerSettings
// Retrieves multitenancy server settings.
func (uda *UserDevicesApi) GetMultitenancyServerSettings(ctx context.Context, c_wstrMtncServerId string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"c_wstrMtncServerId": "%s" }`, c_wstrMtncServerId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetMultitenancyServerSettings",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// GetMultitenancyServersInfo
// Retrieves multitenancy servers list.
func (uda *UserDevicesApi) GetMultitenancyServersInfo(ctx context.Context, nProtocolIds int64) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"nProtocolIds" : %d }`, nProtocolIds))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetMultitenancyServersInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// GetSafeBrowserAutoinstallFlag
// Returns flag which means install or don't install SafeBrowser automatically when device connects first time.
func (uda *UserDevicesApi) GetSafeBrowserAutoinstallFlag(ctx context.Context) (*PxgValBool, []byte,
	error) {
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetSafeBrowserAutoinstallFlag",
		nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := uda.client.Request(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

// GetSyncInfo
// Retrieves group synchronization info for UMDM policy.
func (uda *UserDevicesApi) GetSyncInfo(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GetSyncInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// GlueDevices Glue information on a device got from different sources
//	Deprecated: Not used.
func (uda *UserDevicesApi) GlueDevices(ctx context.Context, lDevice1Id, lDevice2Id int64) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"lDevice1Id": %d, "lDevice2Id" : %d }`, lDevice1Id, lDevice2Id))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.GlueDevices",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// PostCommand Post a command to the specified device.
// Note, you can't sent more then one command with same type to the specified device
// while command with same type is not completed, instead you may call method RecallCommand
func (uda *UserDevicesApi) PostCommand(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.PostCommand", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// RecallCommand previously posted to the specified device.
func (uda *UserDevicesApi) RecallCommand(ctx context.Context, c_wstrCommandGuid string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"c_wstrCommandGuid": "%s"}`, c_wstrCommandGuid))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.RecallCommand",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// SetMultitenancyServerSettings
// Set multitenancy server settings.
func (uda *UserDevicesApi) SetMultitenancyServerSettings(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.SetMultitenancyServerSettings", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// SetSafeBrowserAutoinstallFlag Set flag which means install or don't install SafeBrowser automatically when device connects first time.
//	Deprecated: The SafeBrowser is not supported now
func (uda *UserDevicesApi) SetSafeBrowserAutoinstallFlag(ctx context.Context, bInstall bool) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"bInstall": %v}`, bInstall))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.SetSafeBrowserAutoinstallFlag",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// SspLoginAllowed Check user permission to login to SSP.
//	Deprecated: The SSP is not supported now
func (uda *UserDevicesApi) SspLoginAllowed(ctx context.Context) ([]byte,
	error) {
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.SspLoginAllowed",
		nil)
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}

// UpdateDevice Modify properties of the specified user device.
//	Deprecated: Previously used for the SSP
func (uda *UserDevicesApi) UpdateDevice(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/UserDevicesApi.UpdateDevice", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Request(ctx, request, nil)
	return raw, err
}
