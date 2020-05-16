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

//Acquire all software applications.
//
//Returns attributes for all software applications.
//
//	Parameters:
//	- ctx context.Context
//	- szwHostId string
//	- v interface{} <- "KLEVP_EA_PARAM_1" - list of software applications (paramArray|paramParams)
//each element contains attributes from List of attributes of software inventory application.
//
//	Returns:
//(params) contains following attributes:
//	- raw []byte, err error
func (ia *InventoryApi) GetHostInvProducts(ctx context.Context, szwHostId string, v interface{}) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"szwHostId": "%s"
	}`, szwHostId))
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.GetHostInvProducts", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ia.client.Do(ctx, request, &v)
	return raw, err
}

//Acquire software application updates which are installed on specified host.
//
//Acquire software application updates which are installed on specified host.
//
//	Parameters:
//	- ctx	(context.Context) context.
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
func (ia *InventoryApi) GetHostInvPatches(ctx context.Context, szwHostId string, v interface{}) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"szwHostId": "%s"
	}`, szwHostId))
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.GetHostInvPatches", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ia.client.Do(ctx, request, &v)
	return raw, err
}

//Acquire all software application updates.
//
//Returns attributes for all software application updates.
//
//Parameters:
//pParams	reserved. (params)
//Returns:
//(params) contains following attributes:
//"KLEVP_EA_PARAM_1" - list of software application updates (paramArray|paramParams)
//each element contains attributes from List of attributes of software inventory application update.
func (ia *InventoryApi) GetInvPatchesList(ctx context.Context, v interface{}) ([]byte, error) {
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.GetInvPatchesList", nil)
	if err != nil {
		return nil, err
	}

	raw, err := ia.client.Do(ctx, request, &v)
	return raw, err
}

// Acquire all software applications.
// Returns attributes for all software applications.
// Parameters:
// 	- pParams	reserved. (params)
// Returns:
// (params) contains following attributes:
// "KLEVP_EA_PARAM_1" - list of software applications (paramArray|paramParams)
// each element contains attributes from List of attributes of software inventory application.
func (ia *InventoryApi) GetInvProductsList(ctx context.Context, v interface{}) ([]byte, error) {
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.GetInvProductsList", nil)
	if err != nil {
		return nil, err
	}

	raw, err := ia.client.Do(ctx, request, &v)
	return raw, err
}

//Remove from database info about software applications which aren't installed on any host.
//
//Parameters:
//	- pParams	reserved. (params)
func (ia *InventoryApi) DeleteUninstalledApps(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.DeleteUninstalledApps", nil)
	if err != nil {
		return nil, err
	}

	raw, err := ia.client.Do(ctx, request, nil)
	return raw, err
}

//Acquire info about all cleaner ini-files of specified type from SC-server.
//
//Returns info about cleaner ini-files of specified type from SC-server. These files are used to detect and uninstall applications which incompatible with KasperskyLab antivirus applications
//
// Parameters:
//	- wstrType	(string) Type of the ini-file which should be returned. one of the following:
//	- empty string. Then all files will be returned
//	- "uninstall" - ini-file that's may detect and unistall incompatible application
//	- "detect-only" - ini-file that's may only detect incompatible application
//
//pParams	reserved. (params)
//
//Returns:
//	- (array) collection of paramParams objects where each of them has the following structure:
//each element contains attributes from List of attributes of cleaner ini-files.
func (ia *InventoryApi) GetSrvCompetitorIniFileInfoList(ctx context.Context, wstrType string) (*PxgValCIFIL, []byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
		"wstrType": "%s"
	}`, wstrType))
	request, err := http.NewRequest("POST", ia.client.Server+"/api/v1.0/InventoryApi.GetSrvCompetitorIniFileInfoList", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValCIFIL := new(PxgValCIFIL)
	raw, err := ia.client.Do(ctx, request, &pxgValCIFIL)
	return pxgValCIFIL, raw, err
}

/*
	TODO -> func (ia *InventoryApi) GetObservedApps(ctx context.Context) ([]byte, error)
	TODO -> func (ia *InventoryApi) SetObservedApps(ctx context.Context, v interface{}) ([]byte, error)
*/
