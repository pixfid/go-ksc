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

//	InventoryApi Class Reference
//
//	Interface for working with Software Inventory subsystem..
//
//	Interface allows to get information about software applications
//	that's are installed on client hosts and modify some settings for Software Inventory subsystem.
//	To get additional information you also can use srvview (see Software inventory list)
//
//	List of all members.
type InventoryApi service

//	Acquire all software applications.
//
//	Returns attributes for all software applications.
//
//	Parameters:
//	- szwHostId string
//	- out interface{} <- "KLEVP_EA_PARAM_1" - list of software applications (paramArray|paramParams)
//each element contains attributes from List of attributes of software inventory application.
//
//	Returns:
//	- (params) contains following attributes:
func (ia *InventoryApi) GetHostInvProducts(ctx context.Context, szwHostId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwHostId": "%s"}`, szwHostId))
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.GetHostInvProducts", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ia.client.Do(ctx, request, nil)
	return raw, err
}

//	Acquire software application updates which are installed on specified host.
//
//	Acquire software application updates which are installed on specified host.
//
//	Parameters:
//	- szwHostId	(string) host name, a unique server-generated string (see KLHST_WKS_HOSTNAME attribute).
//	(It is NOT the same as computer network name (DNS-, FQDN-, NetBIOS-name))
//
//	- interface{} "KLEVP_EA_PARAM_1" - list of software application updates (paramArray|paramParams)
//
//each element contains attributes from List of attributes of software inventory application update.
//	pParams	reserved. (params)
//
//	Returns:
//	- raw []byte, err error
func (ia *InventoryApi) GetHostInvPatches(ctx context.Context, szwHostId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwHostId": "%s"}`, szwHostId))
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.GetHostInvPatches", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ia.client.Do(ctx, request, nil)
	return raw, err
}

//	Acquire all software application updates.
//
//	Returns attributes for all software application updates.
//
//	Parameters:
//	- pParams	reserved. (params)
//
//	Returns:
//	- (params) contains following attributes:
//	"KLEVP_EA_PARAM_1" - list of software application updates (paramArray|paramParams)
//	each element contains attributes from List of attributes of software inventory application update.
func (ia *InventoryApi) GetInvPatchesList(ctx context.Context, params Null) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.GetInvPatchesList", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ia.client.Do(ctx, request, nil)
	return raw, err
}

// Acquire all software applications.
// Returns attributes for all software applications.
//
//	Parameters:
// 	- pParams	reserved. (params)
//
//	Returns:
//	- (params) contains following attributes:
//	|- "KLEVP_EA_PARAM_1" - list of software applications (paramArray|paramParams)
// each element contains attributes from List of attributes of software inventory application.
func (ia *InventoryApi) GetInvProductsList(ctx context.Context, params Null) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.GetInvProductsList", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ia.client.Do(ctx, request, nil)
	return raw, err
}

//	Remove from database info about software applications which aren't installed on any host.
//
//	Parameters:
//	- pParams	reserved. (params)
func (ia *InventoryApi) DeleteUninstalledApps(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.DeleteUninstalledApps", nil)
	if err != nil {
		return nil, err
	}

	raw, err := ia.client.Do(ctx, request, nil)
	return raw, err
}

//	Acquire info about all cleaner ini-files of specified type from SC-server.
//
//	Returns info about cleaner ini-files of specified type from SC-server. These files are used to detect and uninstall applications which incompatible with KasperskyLab antivirus applications
//
// Parameters:
//	- wstrType	(string) Type of the ini-file which should be returned. one of the following:
//	- empty string. Then all files will be returned
//	- "uninstall" - ini-file that's may detect and unistall incompatible application
//	- "detect-only" - ini-file that's may only detect incompatible application
//
//	- pParams	reserved. (params)
//
//	Returns:
//	- (array) collection of paramParams objects where each of them has the following structure:
//each element contains attributes from List of attributes of cleaner ini-files.
func (ia *InventoryApi) GetSrvCompetitorIniFileInfoList(ctx context.Context, wstrType string) (*PxgValCIFIL, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrType": "%s"}`, wstrType))
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.GetSrvCompetitorIniFileInfoList", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValCIFIL := new(PxgValCIFIL)
	raw, err := ia.client.Do(ctx, request, &pxgValCIFIL)
	return pxgValCIFIL, raw, err
}

//	Acquire list of observed applications.
//
//	Acquire list of observed applications.
//	When observed application is installed on any host it is published
//	"KLNAG_EV_INV_OBS_APP_INSTALLED" event.
//	When observed application is uninstalled on any host
//	it is published "KLNAG_EV_INV_OBS_APP_UNINSTALLED" event.
//
//	Parameters:
//	- pParams	reserved. (params)
//
//	Returns:
//	- (array) collection of paramString application string Id. (
//	see "ProductID" from List of attributes of software inventory application )
func (ia *InventoryApi) GetObservedApps(ctx context.Context, params Null) (*PxgValArrayOfString, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.GetObservedApps", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValArrayOfString := new(PxgValArrayOfString)
	raw, err := ia.client.Do(ctx, request, &pxgValArrayOfString)
	return pxgValArrayOfString, raw, err
}

//ObservedAppsParams struct using in InventoryApi.SetObservedApps
type ObservedAppsParams struct {
	//collection of (string) application string Id.
	PAppIDS []string `json:"pAppIds"`

	//reserved. (params)
	PParams Null `json:"pParams"`
}

//	Set list of observed applications.
//
//	When observed application is installed on any host it is published "KLNAG_EV_INV_OBS_APP_INSTALLED" event.
//	When observed application is uninstalled on any host it is published "KLNAG_EV_INV_OBS_APP_UNINSTALLED" event.
//
//	Parameters:
//	- pAppIds	(array) collection of (string) application string Id. (
//	see "ProductID" from List of attributes of software inventory application )
//	- pParams	reserved. (params)
func (ia *InventoryApi) SetObservedApps(ctx context.Context, params Null) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.SetObservedApps", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ia.client.Do(ctx, request, nil)
	return raw, err
}
