/*
 *
 * 	Copyright (C) 2020  <Semchenko Aleksandr>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.If not, see <http://www.gnu.org/licenses/>.
 * /
 */

package kaspersky

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//	LicenseKeys Class Reference
//
//	Operating with keys
//
//	Public Member Functions:
type LicenseKeys struct {
	client *Client
}

//	Install a key on the server.
//
//	Parameters:
//	- pKeyInfo	(params) input data container, mandatory:
//
//Installing a key by specifying contents of a key file:
//
//	 ● "KLLICSRV_KEYDATA" - key data container, mandatory (paramParams).
//	-	◯ "KLLIC_IFKEYFILE" - set to true in this case, mandatory (paramBool)
//	-	◯ "KLLIC_LICFILE" - keyfile body, mandatory (paramBinary)
//
//	Installing a key by specifying just an activation 2.0 code:
//
//	 ● "KLLICSRV_KEYDATA" - key data container, mandatory (paramParams).
//	-	◯ "KLLIC_IFKEYFILE" - set to false in this case (paramBool)
//	-	◯ "KLLIC_LICFILE" - ASCII-encoded string with activation code in format of XXXXX-XXXXX-XXXXX-XXXXX (
//	exactly 23 characters long buffer) (paramBinary)
//
//	Returns:
//
//	(string) serial number of the installed license.
//
//	Exceptions:
//
//	Throws	exception in case of error.
func (lk *LicenseKeys) InstallKey(ctx context.Context, pKeyInfo interface{}) bool {

	postData, _ := json.Marshal(pKeyInfo)

	request, err := http.NewRequest("POST", lk.client.Server+"/api/v1.0/LicenseKeys.InstallKey",
		bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := lk.client.Do(ctx, request, nil)

	if raw != nil {
		return true
	}
	return false
}

//TODO Download license key files from activation key servers V1.
//
//Parameters:
//strActivationCode	(string) activation code in format XXXXX-XXXXX-XXXXX-XXXXX, mandatory.
//Returns:
//(array) of paramBinary, array of license key files related to the specified code.
//Exceptions:
//Throws	exception in case of error.
func (lk *LicenseKeys) DownloadKeyFiles(ctx context.Context, wstrActivationCode string) bool {
	postData := []byte(fmt.Sprintf(`
	{
	"wstrActivationCode": "%s"
	}`, wstrActivationCode))

	request, err := http.NewRequest("POST", lk.client.Server+"/api/v1.0/LicenseKeys.DownloadKeyFiles",
		bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := lk.client.Do(ctx, request, nil)

	if raw != nil {
		return true
	}
	return false
}

//AcquireKeyHostsParams struct
type AcquireKeyHostsParams struct {
	PInData        PInData  `json:"pInData"`
	PFields        []string `json:"pFields"`
	PFieldsToOrder []string `json:"pFieldsToOrder"`
	LTimeoutSEC    int64    `json:"lTimeoutSec"`
}

type PInData struct {
	KllicSerial string `json:"KLLIC_SERIAL"`
}

type HostsKeyIterator struct {
	LKeyCount    int64  `json:"lKeyCount"`
	WstrIterator string `json:"wstrIterator"`
}

//	Get an array of hosts that are currently using the specified key.
//
//	Parameters:
//	- pInData	(params) container of input attributes:
//	- "KLLIC_SERIAL" - serial number of the key which the hosts array will be returned for, mandatory (string).
//	- pFields	(array) of string, array of host attribute names to return.
//	See List of host-specific license attributes for attribute names.
//	- pFieldsToOrder	(array) of string, array of host attributes to be used for ordering
//	- pOptions	(params) currently ignored.
//	- lTimeoutSec	(int64) iterator timeout in seconds. Output iterator will be available for this time long.
//		Default is zero which means 15 minutes long.
//
//	Return:
//	- lKeyCount	(int64) count of keys returned via enumerator.
//	- wstrIterator	(string) forward-iterator name for accesing key attributes through GroupSyncIterator.
//
//	Note:
//	Call GroupSyncIterator::GetNextItems to iterate. Call GroupSyncIterator::ReleaseIterator when you are done.
func (lk *LicenseKeys) AcquireKeyHosts(ctx context.Context, params AcquireKeyHostsParams) (*HostsKeyIterator,
	[]byte, error) {
	postData, _ := json.Marshal(params)

	request, err := http.NewRequest("POST", lk.client.Server+"/api/v1.0/LicenseKeys.AcquireKeyHosts",
		bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	hostsKeyIterator := new(HostsKeyIterator)
	raw, err := lk.client.Do(ctx, request, &hostsKeyIterator)

	return hostsKeyIterator, raw, nil
}

//	EnumKeysParams struct
type EnumKeysParams struct {
	PFields     []string `json:"pFields"`
	LTimeoutSEC int64    `json:"lTimeoutSec,omitempty"`
}

// Enumerate keys.
//
// Parameters:
//	- pFields	(array) of string, array of attribute names to return,
//	mandatory . See List of license key attributes for attribute names.
//	- pFieldsToOrder	(array) of string, array of attributes to be used for ordering, mandatory
//	- pOptions	(params) enumeration options, optional :
//	"KLLICSRV_ENOPT_INSTALLED_ONLY" - true if you are enumerating only installed keys.
//	Default is false - method will return all keys including ones not-installed on adm. server (product reported keys).
//	- lTimeoutSec	(int64) iterator timeout in seconds, optional.
//
//	Return: Output iterator will be available for this time long. Default is zero which means 15 minutes long.
//	- lKeyCount	(int64) count of keys returned via enumerator.
//	- wstrIterator	(string) forward-iterator name for accesing key attributes through GroupSyncIterator.
//
//	Call GroupSyncIterator::GetNextItems to iterate. Call GroupSyncIterator::ReleaseIterator when you are done.
//
//Exceptions:
//	Throws	exception in case of error.
func (lk *LicenseKeys) EnumKeys(ctx context.Context, params EnumKeysParams, v interface{}) ([]byte, error) {

	postData, _ := json.Marshal(params)

	request, err := http.NewRequest("POST", lk.client.Server+"/api/v1.0/LicenseKeys.EnumKeys",
		bytes.NewBuffer(postData))

	raw, err := lk.client.Do(ctx, request, &v)
	return raw, err
}

//KeyDataParams struct
type KeyDataParams struct {
	PKeyInfo PKeyInfo `json:"pKeyInfo,omitempty"`
}

type PKeyInfo struct {
	KllicSerial     string `json:"KLLIC_SERIAL,omitempty"`
	KllicsrvKeydata bool   `json:"KLLICSRV_KEYDATA,omitempty"`
}

//	Get data of a key.
//
//	Parameters:
//	pKeyInfo	(params) container which must contain "KLLIC_SERIAL" attribute to specify the interested license.
//
//	For any attribute to query you must put such attribute with any value into the container pKeyInfo.
//	In particular, if you need the key body then put into
//	pKeyInfo container the attribute with name "KLLICSRV_KEYDATA" of type (bool).
//	If the license key has been uploaded to the Administration Server
//	so that and Administration Server has license key body then it will be returned in "KLLICSRV_KEYDATA" attribute.
//	Note that the additional "ExportLicense" access right must be set up to the user under which this call is made.
//
//	Returns:
//	(params) container with the requested key attribute values. See List of license key attributes for attribute names.
func (lk *LicenseKeys) GetKeyData(ctx context.Context, params KeyDataParams, v interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)

	request, err := http.NewRequest("POST", lk.client.Server+"/api/v1.0/LicenseKeys.GetKeyData",
		bytes.NewBuffer(postData))

	raw, err := lk.client.Do(ctx, request, &v)
	return raw, err
}

//TODO func (lk *LicenseKeys) AdjustKey(ctx context.Context, params Params, v interface{})
//TODO func (lk *LicenseKeys) CheckIfSaasLicenseIsValid(ctx context.Context, params Params, v interface{})
//TODO func (lk *LicenseKeys) SaasTryToInstall(ctx context.Context, params Params, v interface{})
//TODO func (lk *LicenseKeys) SaasTryToUninstall(ctx context.Context, params Params, v interface{})
//TODO func (lk *LicenseKeys) UninstallKey(ctx context.Context, params Params, v interface{})
