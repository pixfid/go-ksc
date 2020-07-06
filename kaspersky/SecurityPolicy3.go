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

// SecurityPolicy3 Allows to manage security groups of internal users.
//
// Use srvview SplUserGroupSrvViewName to get information about relationship between users and groups.
// To use this class, a caller must has 'Change security rights' permissions (access mask 0x400) in
// User Permissions functional area on Administration server.
type SecurityPolicy3 service

// SecurityGroupParams struct
type SecurityGroupParams struct {
	// LVsID id of a virtual server, a value = 0 means main server.
	// It is ignored in case of connection to virtual server.
	LVsID      int64      `json:"lVsId,omitempty"`
	PGrpParams PGrpParams `json:"pGrpParams,omitempty"`
}

type PGrpParams struct {
	Type  string    `json:"type,omitempty"`
	Value PGrpValue `json:"value,omitempty"`
}

type PGrpValue struct {
	// KlsplSECGrpName name of a group,
	KlsplSECGrpName string `json:"KLSPL_SEC_GRP_NAME,omitempty"`

	// KlsplSECGrpDesc description of a group.
	KlsplSECGrpDesc string `json:"KLSPL_SEC_GRP_DESC,omitempty"`
}

// AddSecurityGroup Creates a security group on a server.
// lVsId	(int64) id of a virtual server, a value = 0 means main server.
//
// It is ignored in case of connection to virtual server.
//
func (sp *SecurityPolicy3) AddSecurityGroup(ctx context.Context, params SecurityGroupParams) (*PxgValInt, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy3.AddSecurityGroup", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgValInt := new(PxgValInt)
	_, err = sp.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, err
}

// AddUserIntoSecurityGroup Adds user into a security group.
//
// If a group or user does not exist a error occurs.
//
// Connection to a virtual server has access only to groups and users which located on this virtual server.
//
// A user located on a virtual server can be added only into a group located on this virtual server, otherwise a error occurs.
//
// A user located on a main server can be added only into a group located on this main server, otherwise a error occurs.
func (sp *SecurityPolicy3) AddUserIntoSecurityGroup(ctx context.Context, lUserId, lGrpId int64) error {
	postData := []byte(fmt.Sprintf(`{"lUserId": %d, "lGrpId": %d}`, lUserId, lGrpId))
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy3.AddUserIntoSecurityGroup", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = sp.client.Do(ctx, request, nil)
	return err
}

// CloseUserConnections Closes user connections.
func (sp *SecurityPolicy3) CloseUserConnections(ctx context.Context, lUserId int64) error {
	postData := []byte(fmt.Sprintf(`{"lUserId": %d }`, lUserId))
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy3.CloseUserConnections", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = sp.client.Do(ctx, request, nil)
	return err
}

// DeleteSecurityGroup Delete a security group.
//
// Connection to a virtual server has access only to groups which located on this virtual server.
func (sp *SecurityPolicy3) DeleteSecurityGroup(ctx context.Context, lGrpId int64) error {
	postData := []byte(fmt.Sprintf(`{"lGrpId": %d}`, lGrpId))
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy3.DeleteSecurityGroup", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = sp.client.Do(ctx, request, nil)
	return err
}

// DeleteUserFromSecurityGroup Removes user from a security group.
//
// Connection to a virtual server has access only to groups and users which located on this virtual server.
func (sp *SecurityPolicy3) DeleteUserFromSecurityGroup(ctx context.Context, lUserId, lGrpId int64) error {
	postData := []byte(fmt.Sprintf(`{"lUserId": %d, "lGrpId": %d}`, lUserId, lGrpId))
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy3.DeleteUserFromSecurityGroup", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = sp.client.Do(ctx, request, nil)
	return err
}

// MoveUserIntoOtherSecurityGroup Moves user from one security group into other security group.
func (sp *SecurityPolicy3) MoveUserIntoOtherSecurityGroup(ctx context.Context, lUserId, lGrpIdFrom, lGrpIdTo int64) error {
	postData := []byte(fmt.Sprintf(`{"lUserId": %d, "lGrpIdFrom": %d, "lGrpIdTo": %d}`, lUserId, lGrpIdFrom, lGrpIdTo))
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy3.MoveUserIntoOtherSecurityGroup", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = sp.client.Do(ctx, request, nil)
	return err
}

// SecurityGroupParams struct
type UpdateSecurityGroupParams struct {
	// lGrpId id of a created group.
	LGrpId int64 `json:"lGrpId,omitempty"`
	// PGrpParams parameters of a group.
	PGrpParams PGrpParams `json:"pGrpParams,omitempty"`
}

// UpdateSecurityGroup Security group update.
// If a group does not exist a error occurs.
// Connection to a virtual server has access only to groups which located on this virtual server.
func (sp *SecurityPolicy3) UpdateSecurityGroup(ctx context.Context, params UpdateSecurityGroupParams) error {
	postData, err := json.Marshal(params)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy3.UpdateSecurityGroup", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = sp.client.Do(ctx, request, nil)
	return err
}
