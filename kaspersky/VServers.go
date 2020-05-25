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

//	VServers Class Reference
//
//	Virtual servers processing.
//
//	Detailed Description
//
//	Allows to create and destroy virtual servers, acquire and modify their attributes.
type VServers service

//	Acquire virtual servers for the specified group.
//
//	Returns array of virtual servers for the specified group
//
//	Parameters:
//	- lParentGroup	(int64) id of parent group, -1 means 'from all groups'
//
//	Returns:
//	- (array) array, each element is a container KLPAR.ParamsPtr containing attributes "KLVSRV_*"
//	(see List of virtual server attributes).
func (vs *VServers) GetVServers(ctx context.Context, lParentGroup int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lParentGroup": %d}`, lParentGroup))
	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.GetVServers", bytes.NewBuffer(postData))

	raw, err := vs.client.Do(ctx, request, nil)
	return raw, err
}

//VServerInfo struct
type VServerInfo struct {
	VServer *VServer `json:"PxgRetVal,omitempty"`
}

//	VServer struct
type VServer struct {
	//Creation time.
	KlvsrvCreated *KlvsrvCreated `json:"KLVSRV_CREATED,omitempty"`

	//Virtual server display name.
	KlvsrvDN *string `json:"KLVSRV_DN,omitempty"`

	//If the specified virtual server is enabled.
	KlvsrvEnabled *bool `json:"KLVSRV_ENABLED,omitempty"`

	//Id of virtual server's group "Managed computers".
	KlvsrvGroups *int64 `json:"KLVSRV_GROUPS,omitempty"`

	//Id of parent group
	KlvsrvGrp *int64 `json:"KLVSRV_GRP,omitempty"`

	//Host name of the virtual server (See KLHST_WKS_HOSTNAME attribute )
	KlvsrvHstUid *string `json:"KLVSRV_HST_UID,omitempty"`

	//Virtual server ID
	KlvsrvID *int64 `json:"KLVSRV_ID,omitempty"`

	//Allow automatic deployment of license keys from the virtual Server to its devices.
	KlvsrvLicEnabled *bool `json:"KLVSRV_LIC_ENABLED,omitempty"`

	//New managed hosts cannot be added to the VS due to the limitation violation
	//(see LP_VS_MaxCountOfHosts), read-only attribute.
	KlvsrvNewHostsProhibited *bool `json:"KLVSRV_NEW_HOSTS_PROHIBITED,omitempty"`

	//Id of virtual server's group "Master server".
	KlvsrvSuper *int64 `json:"KLVSRV_SUPER,omitempty"`

	//Number of managed hosts connected to a VS has already reached 10% of the limit
	//(see LP_VS_MaxCountOfHosts) , read-only attribute.
	KlvsrvTooMuchHosts *bool `json:"KLVSRV_TOO_MUCH_HOSTS,omitempty"`

	//Virtual server name - a globally unique string
	KlvsrvUid *string `json:"KLVSRV_UID,omitempty"`

	//Id of virtual server's group "Unassigned computers".
	KlvsrvUnassigned *int64 `json:"KLVSRV_UNASSIGNED,omitempty"`
}

type KlvsrvCreated struct {
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

//	Register new virtual server.
//
//	Parameters:
//	- strDisplayName	(string) virtual server display name, if display name is non-unique,
//	it will be modified to become unique
//	- lParentGroup	(int64) virtual server parent group
//
//	Returns:
//	- (params) a container KLPAR.ParamsPtr containing attributes "KLVSRV_ID" and "KLVSRV_DN" (
//	see List of virtual server attributes).
//
func (vs *VServers) AddVServerInfo(ctx context.Context, strDisplayName string, lParentGroup int64) (*VServer, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"lParentGroup": %d, "strDisplayName" : "%s"}`, lParentGroup, strDisplayName))
	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.AddVServerInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	vServer := new(VServer)
	raw, err := vs.client.Do(ctx, request, &vServer)
	return vServer, raw, err
}

//	Unregister specified Virtual Server.
//
//	Unregisters specified Virtual Server
//
//	Parameters:
//	- lVServer	(int64) Virtual Server id
//	- [out]	strActionGuid	(string) id of asynchronous operation,
//	to get status use AsyncActionStateChecker.CheckActionState
func (vs *VServers) DelVServer(ctx context.Context, lVServer int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lVServer": %d}`, lVServer))
	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.DelVServer", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vs.client.Do(ctx, request, nil)
	return raw, err
}

//	Return ACL for the specified virtual server.
//
//	Returns ACL for the specified virtual server
//
//	Parameters:
//	- lVServer	(int64) virtual server id
func (vs *VServers) GetPermissions(ctx context.Context, lVServer int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lVServer": %d}`, lVServer))
	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.GetPermissions", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vs.client.Do(ctx, request, nil)
	return raw, err
}

//HostInfoParams struct
type VServerInfoParams struct {
	LVServer       int64    `json:"lVServer,omitempty"`
	PFields2Return []string `json:"pFields2Return"`
}

//	Acquire info on virtual server.
//
//	Returns info about the specified virtual server
//
//	Parameters:
//	- params VServerInfoParams
//
//	Returns:
//	- (params) a container containing attributes "KLVSRV_*" (see List of virtual server attributes)
func (vs *VServers) GetVServerInfo(ctx context.Context, params VServerInfoParams) (*VServerInfo, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.GetVServerInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	vServerInfo := new(VServerInfo)
	raw, err := vs.client.Do(ctx, request, &vServerInfo)
	return vServerInfo, raw, err
}

//	Moves specified virtual server.
//
//	Moves specified virtual server
//
//	Parameters:
//	- lVServer	(int64) in Virtual Server id
//	- lNewParentGroup	(int) in New group
//
//	Return:
//	- strActionGuid	(string) id of asynchronous operation,
//	to get status use AsyncActionStateChecker.CheckActionState
func (vs *VServers) MoveVServer(ctx context.Context, lVServer int64, lNewParentGroup int64) (*WActionGUID, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"lVServer": %d, "lNewParentGroup" : %d}`, lVServer, lNewParentGroup))
	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.MoveVServer", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	wActionGUID := new(WActionGUID)
	raw, err := vs.client.Do(ctx, request, &wActionGUID)
	return wActionGUID, raw, err
}

//	Function recalls Network Agent certificate from the specified virtual server
//	and closes active connections from such Network Agents.
//
//	Parameters:
//	- lVServer	(int64) virtual server id
func (vs *VServers) RecallCertAndCloseConnections(ctx context.Context, lVServer int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lVServer": %d}`, lVServer))
	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.RecallCertAndCloseConnections", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vs.client.Do(ctx, request, nil)
	return raw, err
}

//UpdateVServerInfoParams struct
type UpdateVServerInfoParams struct {
	LVServer *int64  `json:"lVServer,omitempty"`
	VSInfo   *VSInfo `json:"pInfo,omitempty"`
}

//VSInfo struct
type VSInfo struct {
	//Virtual server display name.
	KlvsrvDN *string `json:"KLVSRV_DN,omitempty"`

	//If the specified virtual server is enabled.
	KlvsrvEnabled *bool `json:"KLVSRV_ENABLED,omitempty"`

	//Allow automatic deployment of license keys from the virtual Server to its devices.
	KlvsrvLicEnabled *bool `json:"KLVSRV_LIC_ENABLED,omitempty"`

	//Creation time.
	KlvsrvCreated *KlvsrvCreated `json:"KLVSRV_CREATED,omitempty"`

	//Custom info XML.
	KlvsrvCustomInfo *string `json:"KLVSRV_CUSTOM_INFO,omitempty"`

	//New managed hosts cannot be added to the VS due to the limitation violation
	//(see LP_VS_MaxCountOfHosts), read-only attribute.
	KlvsrvNewHostsProhibited *bool `json:"KLVSRV_NEW_HOSTS_PROHIBITED,omitempty"`

	//Number of managed hosts connected to a VS has already reached 10% of the limit
	//(see LP_VS_MaxCountOfHosts) , read-only attribute.
	KlvsrvTooMuchHosts *bool `json:"KLVSRV_TOO_MUCH_HOSTS,omitempty"`
}

//	Modify virtual server attributes.
//
//	Modifies attributes of the specified virtual server
//
//	Parameters:
//	- params UpdateVServerInfoParams
func (vs *VServers) UpdateVServerInfo(ctx context.Context, params UpdateVServerInfoParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.UpdateVServerInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vs.client.Do(ctx, request, nil)
	return raw, err
}

//	Set ACL for the specified virtual server.
//
//	Sets ACL for the specified virtual server
//
//	Parameters:
//	- params interface{}
func (vs *VServers) SetPermissions(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.SetPermissions", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vs.client.Do(ctx, request, nil)
	return raw, err
}
