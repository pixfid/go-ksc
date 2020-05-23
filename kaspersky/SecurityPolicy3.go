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

//	SecurityPolicy3 Class Reference
//
//	Allows to manage security groups of internal users.
//	Use srvview SplUserGroupSrvViewName to get information about relationship between users and groups.
//	To use this class, a caller must has 'Change security rights'
//	permissions (access mask 0x400) in User Permissions functional area on Administration server.
//
//	List of all members.
type SecurityPolicy3 service

//	Creates a security group on a server.
//
//	If a group with such name exists a error occurs.
//
//	Parameters:
//	- pGrpParams	(params) parameters of a group. There are possible values:
//	+--------------------+------------------------+----------+
//	|       Values       |      Description       |   Type   |
//	+--------------------+------------------------+----------+
//	| KLSPL_SEC_GRP_NAME | name of a group        |  string  |
//	| KLSPL_SEC_GRP_DESC | description of a group |  string  |
//	+--------------------+------------------------+----------+
//	- lVsId	(int64) id of a virtual server, a value = 0 means main server.
//	- It is ignored in case of connection to virtual server.
//
//	Structure Example:
//	{
//	  "lVsId" : 0,
//	  "pGrpParams" : {
//	    "type" : "params",
//	    "value" : {
//	      "KLSPL_SEC_GRP_NAME" : "SECURITY_GROUP_NAME",
//	      "KLSPL_SEC_GRP_DESC" : "SECURITY_GROUP_DESCRIPTION"
//	    }
//	  }
//	}
//Returns:
//	- (int64) id of a created group.
func (sp *SecurityPolicy3) AddSecurityGroup(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy3.AddSecurityGroup",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sp.client.Do(ctx, request, nil)
	return raw, err
}

//	Adds user into a security group.
//
//	If a group or user does not exist a error occurs.
//	Connection to a virtual server has access only to groups and users which
//	located on this virtual server.
//
//	A user located on a virtual server can be added only into a group
//	located on this virtual server, otherwise a error occurs.
//
//	A user located on a main server can be added only into a group
//	located on this main server, otherwise a error occurs.
//
//	Parameters:
//	- lGrpId	(int64) id of a group.
//	- lUserId	(int64) id of a user.
func (sp *SecurityPolicy3) AddUserIntoSecurityGroup(ctx context.Context, lUserId, lGrpId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lUserId": %d, "lGrpId": %d}`, lUserId, lGrpId))
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy3.AddUserIntoSecurityGroup",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sp.client.Do(ctx, request, nil)
	return raw, err
}

//	Closes user connections.
//
//	Parameters:
//	- lUserId	(int64) id of a user.
func (sp *SecurityPolicy3) CloseUserConnections(ctx context.Context, lUserId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lUserId": %d }`, lUserId))
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy3.CloseUserConnections",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sp.client.Do(ctx, request, nil)
	return raw, err
}

//	Delete a security group.
//
//	Connection to a virtual server has access only to groups which located on this virtual server.
//
//	Parameters:
//	- lGrpId	(int64) id of a created group.
func (sp *SecurityPolicy3) DeleteSecurityGroup(ctx context.Context, lGrpId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lGrpId": %d}`, lGrpId))
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy3.DeleteSecurityGroup",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sp.client.Do(ctx, request, nil)
	return raw, err
}

//	Removes user from a security group.
//
//	Connection to a virtual server has access only to groups and users
//	which located on this virtual server.
//
//	Parameters:
//	- lGrpId	(int64) id of a group.
//	- lUserId	(int64) id of a user.
func (sp *SecurityPolicy3) DeleteUserFromSecurityGroup(ctx context.Context, lUserId, lGrpId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lUserId": %d, "lGrpId": %d}`, lUserId, lGrpId))
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy3.DeleteUserFromSecurityGroup",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sp.client.Do(ctx, request, nil)
	return raw, err
}

//	Moves user from one security group into other security group.
//
//	Parameters:
//	- lUserId		(int64) id of a user.
//	- lGrpIdFrom	(int64) id of a group.
//	- lGrpIdTo		(int64) id of a group.
func (sp *SecurityPolicy3) MoveUserIntoOtherSecurityGroup(ctx context.Context, lUserId, lGrpIdFrom, lGrpIdTo int64) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"lUserId": %d, "lGrpIdFrom": %d, "lGrpIdTo": %d}`, lUserId, lGrpIdFrom, lGrpIdTo))
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy3.MoveUserIntoOtherSecurityGroup",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sp.client.Do(ctx, request, nil)
	return raw, err
}

//	Security group update.
//
//	If a group does not exist a error occurs.
//	Connection to a virtual server has access only to groups which located on this virtual server.
//
//	Parameters:
//	- lGrpId	(int) id of a created group.
//	- pGrpParams	(params) parameters of a group. There are possible values:
//	+--------------------+------------------------+----------+
//	|       Values       |      Description       |   Type   |
//	+--------------------+------------------------+----------+
//	| KLSPL_SEC_GRP_NAME | name of a group        |  string  |
//	| KLSPL_SEC_GRP_DESC | description of a group |  string  |
//	+--------------------+------------------------+----------+
//
//	Structure Example:
//	{
//	  "lVsId" : 0,
//	  "pGrpParams" : {
//	    "type" : "params",
//	    "value" : {
//	      "KLSPL_SEC_GRP_NAME" : "SECURITY_GROUP_NAME",
//	      "KLSPL_SEC_GRP_DESC" : "SECURITY_GROUP_DESCRIPTION"
//	    }
//	  }
//	}
func (sp *SecurityPolicy3) UpdateSecurityGroup(ctx context.Context, params interface{}) ([]byte,
	error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy3.UpdateSecurityGroup",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sp.client.Do(ctx, request, nil)
	return raw, err
}
