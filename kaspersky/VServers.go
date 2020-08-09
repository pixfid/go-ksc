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

// VServers service allows to create and destroy virtual servers, acquire and modify their attributes.
type VServers service

// VServersInfos struct using in VServers
type VServersInfos struct {
	VServersInfo *[]VServersInfo `json:"PxgRetVal,omitempty"`
}

type VServersInfo struct {
	Type    *string  `json:"type,omitempty"`
	VServer *VServer `json:"value,omitempty"`
}

// GetVServers Acquire virtual servers for the specified group.
func (vs *VServers) GetVServers(ctx context.Context, lParentGroup int64) (*VServersInfos, error) {
	postData := []byte(fmt.Sprintf(`{"lParentGroup": %d}`, lParentGroup))
	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.GetVServers", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	vServersInfos := new(VServersInfos)
	_, err = vs.client.Do(ctx, request, &vServersInfos)
	return vServersInfos, err
}

// VServerInfo struct using in AddVServerInfo
type VServerInfo struct {
	VServer *VServer `json:"PxgRetVal,omitempty"`
}

type VServer struct {
	// KlvsrvCreated Creation time.
	KlvsrvCreated *DateTime `json:"KLVSRV_CREATED,omitempty"`

	// KlvsrvDN Virtual server display name.
	KlvsrvDN *string `json:"KLVSRV_DN,omitempty"`

	// KlvsrvEnabled If the specified virtual server is enabled.
	KlvsrvEnabled *bool `json:"KLVSRV_ENABLED,omitempty"`

	// KlvsrvGroups Id of virtual server's group "Managed computers".
	KlvsrvGroups *int64 `json:"KLVSRV_GROUPS,omitempty"`

	// KlvsrvGrp Id of parent group
	KlvsrvGrp *int64 `json:"KLVSRV_GRP,omitempty"`

	// KlvsrvHstUid Host name of the virtual server
	KlvsrvHstUid *string `json:"KLVSRV_HST_UID,omitempty"`

	// KlvsrvID Virtual server ID
	KlvsrvID *int64 `json:"KLVSRV_ID,omitempty"`

	// KlvsrvLicEnabled Allow automatic deployment of license keys from the virtual Server to its devices.
	KlvsrvLicEnabled *bool `json:"KLVSRV_LIC_ENABLED,omitempty"`

	// KlvsrvNewHostsProhibited New managed hosts cannot be added to the VS due to the limitation violation
	KlvsrvNewHostsProhibited *bool `json:"KLVSRV_NEW_HOSTS_PROHIBITED,omitempty"`

	// KlvsrvSuper Id of virtual server's group "Master server".
	KlvsrvSuper *int64 `json:"KLVSRV_SUPER,omitempty"`

	// KlvsrvTooMuchHosts Number of managed hosts connected to a VS has already reached 10% of the limit
	//(see LP_VS_MaxCountOfHosts) , read-only attribute.
	KlvsrvTooMuchHosts *bool `json:"KLVSRV_TOO_MUCH_HOSTS,omitempty"`

	// KlvsrvUid Virtual server name - a globally unique string
	KlvsrvUid *string `json:"KLVSRV_UID,omitempty"`

	// KlvsrvUnassigned Id of virtual server's group "Unassigned computers".
	KlvsrvUnassigned *int64 `json:"KLVSRV_UNASSIGNED,omitempty"`
}

// AddVServerInfo Register new virtual server.
func (vs *VServers) AddVServerInfo(ctx context.Context, strDisplayName string, lParentGroup int64) (*VServerInfo, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"lParentGroup": %d, "strDisplayName" : "%s"}`, lParentGroup, strDisplayName))
	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.AddVServerInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	vServerInfo := new(VServerInfo)
	raw, err := vs.client.Do(ctx, request, &vServerInfo)
	return vServerInfo, raw, err
}

// DelVServer Unregister specified Virtual Server.
func (vs *VServers) DelVServer(ctx context.Context, lVServer int64) (*WActionGUID, error) {
	postData := []byte(fmt.Sprintf(`{"lVServer": %d}`, lVServer))
	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.DelVServer", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	wActionGUID := new(WActionGUID)
	_, err = vs.client.Do(ctx, request, &wActionGUID)
	return wActionGUID, err
}

// VServerPermissions struct
type VServerPermissions struct {
	Permissions *Permissions `json:"PxgRetVal,omitempty"`
}

// GetPermissions Return ACL for the specified virtual server.
func (vs *VServers) GetPermissions(ctx context.Context, lVServer int64) (*VServerPermissions, error) {
	postData := []byte(fmt.Sprintf(`{"lVServer": %d}`, lVServer))
	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.GetPermissions", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	vServerPermissions := new(VServerPermissions)
	_, err = vs.client.Do(ctx, request, &vServerPermissions)
	return vServerPermissions, err
}

// VServerInfoParams struct
type VServerInfoParams struct {
	// LVServer virtual server id
	LVServer int64 `json:"lVServer,omitempty"`

	// PFields2Return attributes to acquire
	PFields2Return []string `json:"pFields2Return"`
}

// GetVServerInfo Returns info about the specified virtual server.
func (vs *VServers) GetVServerInfo(ctx context.Context, params VServerInfoParams) (*VServerInfo, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.GetVServerInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	vServerInfo := new(VServerInfo)
	_, err = vs.client.Do(ctx, request, &vServerInfo)
	return vServerInfo, err
}

// MoveVServer Moves specified virtual server to new parent group
func (vs *VServers) MoveVServer(ctx context.Context, lVServer int64, lNewParentGroup int64) (*WActionGUID, error) {
	postData := []byte(fmt.Sprintf(`{"lVServer": %d, "lNewParentGroup" : %d}`, lVServer, lNewParentGroup))
	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.MoveVServer", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	wActionGUID := new(WActionGUID)
	_, err = vs.client.Do(ctx, request, &wActionGUID)
	return wActionGUID, err
}

// RecallCertAndCloseConnections Function recalls Network Agent certificate
// from the specified virtual server and closes active connections from such Network Agents.
func (vs *VServers) RecallCertAndCloseConnections(ctx context.Context, lVServer int64) error {
	postData := []byte(fmt.Sprintf(`{"lVServer": %d}`, lVServer))
	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.RecallCertAndCloseConnections", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = vs.client.Do(ctx, request, nil)
	return err
}

// UpdateVServerInfoParams struct
type UpdateVServerInfoParams struct {
	// LVServer virtual server id
	LVServer int64 `json:"lVServer,omitempty"`

	//VSInfo a container containing no-read-only attributes "KLVSRV_*"
	VSInfo VSInfo `json:"pInfo,omitempty"`
}

// VSInfo struct
type VSInfo struct {
	// KlvsrvDN Virtual server display name.
	KlvsrvDN string `json:"KLVSRV_DN,omitempty"`

	// KlvsrvEnabled If the specified virtual server is enabled.
	KlvsrvEnabled bool `json:"KLVSRV_ENABLED,omitempty"`

	// KlvsrvLicEnabled Allow automatic deployment of license keys from the virtual Server to its devices.
	KlvsrvLicEnabled bool `json:"KLVSRV_LIC_ENABLED,omitempty"`

	// KlvsrvCreated Creation time.
	KlvsrvCreated *DateTime `json:"KLVSRV_CREATED,omitempty"`

	// KlvsrvCustomInfo Custom info XML.
	KlvsrvCustomInfo string `json:"KLVSRV_CUSTOM_INFO,omitempty"`

	// KlvsrvNewHostsProhibited New managed hosts cannot be added to the VS due to the limitation violation
	//(see LP_VS_MaxCountOfHosts), read-only attribute.
	KlvsrvNewHostsProhibited bool `json:"KLVSRV_NEW_HOSTS_PROHIBITED,omitempty"`

	// KlvsrvTooMuchHosts Number of managed hosts connected to a VS has already reached 10% of the limit
	//(see LP_VS_MaxCountOfHosts) , read-only attribute.
	KlvsrvTooMuchHosts bool `json:"KLVSRV_TOO_MUCH_HOSTS,omitempty"`
}

// UpdateVServerInfo Modifies attributes of the specified virtual server
func (vs *VServers) UpdateVServerInfo(ctx context.Context, params UpdateVServerInfoParams) error {
	postData, err := json.Marshal(params)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers.UpdateVServerInfo", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = vs.client.Do(ctx, request, nil)
	return err
}

// ACLParams structure describing ACL permission
type ACLParams struct {
	// LVServer virtual server id
	LVServer int64 `json:"lVServer,omitempty"`

	// Permissions Permissions structure
	Permissions Permissions `json:"pPermissions,omitempty"`

	// Protection if true checks if the user does not denies access to the server to itself
	Protection bool `json:"bProtection,omitempty"`
}

type Permissions struct {
	// KlsplInherited If the permissions list is inherited.
	KlsplInherited bool `json:"KLSPL_INHERITED,omitempty"`

	//KlsplPermissions Permission list
	KlsplPermissions []KlsplPermission `json:"KLSPL_PERMISSIONS"`

	// KlsplPossibleRights Bitmask of access rights available for editing
	KlsplPossibleRights int64 `json:"KLSPL_POSSIBLE_RIGHTS,omitempty"`
}

//KlsplPermission container of permissions entry attributes
type KlsplPermission struct {
	Type string `json:"type,omitempty"`
	// KlsplPermissionValue permissions entry
	KlsplPermissionValue KlsplPermissionValue `json:"value,omitempty"`
}

type KlsplPermissionValue struct {
	// KlspluSid Security ID (if OS account).
	KlspluSid string `json:"KLSPLU_SID,omitempty"`

	// KlspluName User or group name (if OS account).
	KlspluName string `json:"KLSPLU_NAME,omitempty"`

	// KlspluIsGroup Is group (if OS account).
	KlspluIsGroup bool `json:"KLSPLU_IS_GROUP,omitempty"`

	// KlspluAkuserID Id of AK user.
	KlspluAkuserID int64 `json:"KLSPLU_AKUSER_ID,omitempty"`

	// KlspluMayWrite The entry may be modified.
	KlspluMayWrite bool `json:"KLSPLU_MAY_WRITE,omitempty"`

	// KlspluMayRemove The entry may be removed.
	KlspluMayRemove bool `json:"KLSPLU_MAY_REMOVE,omitempty"`

	// KlspluAllowmask Allow access mask
	KlspluAllowmask int64 `json:"KLSPLU_ALLOWMASK,omitempty"`

	// KlspluDenymask Deny access mask
	KlspluDenymask int64 `json:"KLSPLU_DENYMASK,omitempty"`
}

// SetPermissions Set ACL for the specified virtual server.
func (vs *VServers) SetPermissions(ctx context.Context, params ACLParams) ([]byte, error) {
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
