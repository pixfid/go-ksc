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

//	LicenseKeys Class Reference
//
//	Operating with keys
//
//	Public Member Functions:
type LicenseKeys service

//	Install a key on the server.
//
//	Parameters:
//	- pKeyInfo	(params) input data container, mandatory:
//
//Installing a key by specifying contents of a key file:
//
//	 - "KLLICSRV_KEYDATA" - key data container, mandatory (paramParams).
//		|- "KLLIC_IFKEYFILE" - set to true in this case, mandatory (paramBool)
//		|- "KLLIC_LICFILE" - keyfile body, mandatory (paramBinary)
//
//	Installing a key by specifying just an activation 2.0 code:
//
//	- "KLLICSRV_KEYDATA" - key data container, mandatory (paramParams).
//		|-	"KLLIC_IFKEYFILE" - set to false in this case (paramBool)
//		|-"KLLIC_LICFILE" - ASCII-encoded string with activation code in format of XXXXX-XXXXX-XXXXX-XXXXX (
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
		return false
	}

	raw, err := lk.client.Do(ctx, request, nil)
	if raw != nil {
		return true
	}
	return false
}

//
//	Parameters:
//	- strActivationCode	(string) activation code in format XXXXX-XXXXX-XXXXX-XXXXX, mandatory.
//	Returns:
//	- (array) of paramBinary, array of license key files related to the specified code.
//
//	Exceptions:
//	- Throws	exception in case of error.
func (lk *LicenseKeys) DownloadKeyFiles(ctx context.Context, wstrActivationCode string) bool {
	postData := []byte(fmt.Sprintf(`
	{
	"wstrActivationCode": "%s"
	}`, wstrActivationCode))
	request, err := http.NewRequest("POST", lk.client.Server+"/api/v1.0/LicenseKeys.DownloadKeyFiles",
		bytes.NewBuffer(postData))
	if err != nil {
		return false
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
//	Call GroupSyncIterator.GetNextItems to iterate. Call GroupSyncIterator.ReleaseIterator when you are done.
func (lk *LicenseKeys) AcquireKeyHosts(ctx context.Context, params AcquireKeyHostsParams) (*HostsKeyIterator,
	[]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", lk.client.Server+"/api/v1.0/LicenseKeys.AcquireKeyHosts",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
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
//	Call GroupSyncIterator.GetNextItems to iterate. Call GroupSyncIterator.ReleaseIterator when you are done.
//
//Exceptions:
//	Throws	exception in case of error.
func (lk *LicenseKeys) EnumKeys(ctx context.Context, params EnumKeysParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", lk.client.Server+"/api/v1.0/LicenseKeys.EnumKeys",
		bytes.NewBuffer(postData))

	raw, err := lk.client.Do(ctx, request, nil)
	return raw, err
}

//KeyDataParams struct
type KeyDataParams struct {
	//	PKeyInfo container which must contain "KLLIC_SERIAL" attribute to specify the interested license.
	//	For any attribute to query you must put such attribute with any value into the container pKeyInfo.
	//	In particular, if you need the key body then put into pKeyInfo container the attribute with name "KLLICSRV_KEYDATA" of type
	PKeyInfo PKeyInfo `json:"pKeyInfo,omitempty"`
}

type PKeyInfo struct {
	//	KllicSerial
	KllicSerial string `json:"KLLIC_SERIAL,omitempty"`

	//	KllicsrvKeydata
	KllicsrvKeydata bool `json:"KLLICSRV_KEYDATA,omitempty"`
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
func (lk *LicenseKeys) GetKeyData(ctx context.Context, params KeyDataParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", lk.client.Server+"/api/v1.0/LicenseKeys.GetKeyData",
		bytes.NewBuffer(postData))

	raw, err := lk.client.Do(ctx, request, nil)
	return raw, err
}

//	Uninstall an adm. server's license.
//
//	Parameters:
//	- bCurrent	(bool) true if the current license should be uninstalled, false to uninstall the reserved one.
//
//	Exceptions:
//	- Throws	exception in case of error.
//
//	See also:
//	SaasTryToInstall
func (lk *LicenseKeys) SaasTryToUninstall(ctx context.Context, bCurrent bool) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"bCurrent": %v}`, bCurrent))
	request, err := http.NewRequest("POST", lk.client.Server+"/api/v1.0/LicenseKeys.SaasTryToUninstall",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lk.client.Do(ctx, request, nil)
	return raw, err
}

//	AdjustKeyParams struct
type AdjustKeyParams struct {
	//	PKData container of input attributes, mandatory.
	//	See List of license key attributes for attribute names
	PKData *PKData `json:"pData,omitempty"`
}

//	PKData struct
type PKData struct {
	// KllicsrvAutokey License serial number (mandatory)
	KllicsrvAutokey bool `json:"KLLICSRV_AUTOKEY,omitempty"`

	// KllicSerial true if license can be deployed automatically,
	// false otherwise (string, mandatory)
	KllicSerial string `json:"KLLIC_SERIAL,omitempty"`
}

//	Adjust adm. server's license attributes.
//
//	Parameters:
//	- pData	(params) container of input attributes, mandatory.
//	See List of license key attributes for attribute names.
//	Supported attributes:
//	- "KLLIC_SERIAL" - (string) License serial number (mandatory)
//	- "KLLICSRV_AUTOKEY" - (bool) true if license can be deployed automatically, false otherwise (paramString,
//	mandatory)
//
//	Exceptions:
//	- Throws	exception in case of error.
func (lk *LicenseKeys) AdjustKey(ctx context.Context, params AdjustKeyParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", lk.client.Server+"/api/v1.0/LicenseKeys.AdjustKey",
		bytes.NewBuffer(postData))

	raw, err := lk.client.Do(ctx, request, nil)
	return raw, err
}

//SaasKeyParam struct using in LicenseKeys.SaasTryToInstall
type SaasKeyParam struct {
	//input data container, mandatory.
	SaasKeyPInData *SaasKeyPInData `json:"pInData,omitempty"`

	//true if license should be installed in place of active one,
	//false to install it as a reserved one.
	BAsCurrent bool `json:"bAsCurrent,omitempty"`
}

type SaasKeyPInData struct {
	//serial number of the license being checked, mandatory (paramString).
	//The license must be placed in the license store before installing (see InstallKey).
	KllicSerial string `json:"KLLIC_SERIAL,omitempty"`
}

//	Install an adm. server's license.
//
//	Parameters:
//	- params SaasKeyParam
//		|- pInData	(params) input data container, mandatory. Attributes are allowed here:
//			|- "KLLIC_SERIAL" - serial number of the license being checked,
//	mandatory (paramString). The license must be placed in the license store before installing (see InstallKey).
//		|- bAsCurrent	(boolean) true if license should be installed in place of active one,
//	false to install it as a reserved one.
//
//	Exceptions:
//	- Throws exception in case of error.
//
//	See also:
//	- LicenseKeys.SaasTryToUninstall
func (lk *LicenseKeys) SaasTryToInstall(ctx context.Context, params SaasKeyParam) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", lk.client.Server+"/api/v1.0/LicenseKeys.SaasTryToInstall",
		bytes.NewBuffer(postData))

	raw, err := lk.client.Do(ctx, request, nil)
	return raw, err
}

//	Check if license can be installed to the adm. server.
//	License is treated as a valid one if it is suitable for being used by the adm.
//	server itself (its AppId is 1017, it is not expired etc.)
//
//	Parameters:
//	- params SaasKeyParam
//	|- pInData	(params) container of input attributes, mandatory. Attributes are allowed here:
//		|- "KLLIC_SERIAL" - serial number of the license being checked,
//		mandatory (paramString). The license must be placed in the license store before checking (see InstallKey).
//	|- bAsCurrent	(boolean) true if license should be checked in place of active one, false otherwise.
//
//	Exceptions:
//	- KLSTD::STDE_NOTPERM	license is not valid, see Some error definitions.
//	- KLERR::Error*	if an error occured during the checking process.
func (lk *LicenseKeys) CheckIfSaasLicenseIsValid(ctx context.Context, params SaasKeyParam) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", lk.client.Server+"/api/v1.0/LicenseKeys.CheckIfSaasLicenseIsValid",
		bytes.NewBuffer(postData))

	raw, err := lk.client.Do(ctx, request, nil)
	return raw, err
}

//	Uninstall an adm. server's license.
//
//	Parameters:
//	- bCurrent	(boolean) true if the current license should be uninstalled, false to uninstall the reserved one.
//
//	Exceptions:
//	- Throws	exception in case of error.
//
//	See also:
//	- LicenseKeys.SaasTryToInstall
func (lk *LicenseKeys) UninstallKey(ctx context.Context, bCurrent bool) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"bCurrent": %v}`, bCurrent))
	request, err := http.NewRequest("POST", lk.client.Server+"/api/v1.0/LicenseKeys.UninstallKey",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lk.client.Do(ctx, request, nil)
	return raw, err
}
