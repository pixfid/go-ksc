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
	"log"
	"net/http"
)

//	SecurityPolicy Class Reference
//
//	Detailed Description
//	Allows to manage users and permissions.
//
//	List of all members.
type SecurityPolicy service

type PUserData struct {
	PUser PUser `json:"pUser"`
}

type PUser struct {
	KlsplUserName        *string `json:"KLSPL_USER_NAME"`
	KlsplUserPwd         *string `json:"KLSPL_USER_PWD"`
	KlsplUserFullName    *string `json:"KLSPL_USER_FULL_NAME"`
	KlsplUserDescription *string `json:"KLSPL_USER_DESCRIPTION"`
	KlsplUserMail        *string `json:"KLSPL_USER_MAIL"`
	KlsplUserPhone       *string `json:"KLSPL_USER_PHONE"`
	KlsplUserEnabled     *bool   `json:"KLSPL_USER_ENABLED"`
}

//	Add new user.
//
//	Parameters:
//	- pUser	(params) user info, containing attributes "KLSPL_USER_*" (see List of user attributes).
//
//	Following attributes are required: -"KLSPL_USER_NAME" -"KLSPL_USER_PWD"
//
//	+----------------------------+-------------+------------------------------------------------------+
//	|           Param            |    Type     |                     Description                      |
//	+----------------------------+-------------+------------------------------------------------------+
//	| "KLSPL_USER_ID"            | Int64       | User id.                                             |
//	| "KLSPL_USER_NAME"          | string      | User name.                                           |
//	| "KLSPL_USER_PWD"           | string      | Plain text user password. Obsolete.                  |
//	| "KLSPL_USER_PWD_ENCRYPTED" | paramBinary | Encrypted user password.                             |
//	| "KLSPL_USER_FULL_NAME"     | string      | User full name.                                      |
//	| "KLSPL_USER_DESCRIPTION"   | string      | User description.                                    |
//	| "KLSPL_USER_MAIL"          | string      | User mail.                                           |
//	| "KLSPL_USER_PHONE"         | string      | User phone.                                          |
//	| "KLSPL_USER_ENABLED"       | bool        | User account is enabled if true.                     |
//	| "KLSPL_USER_UIS"           | bool        | User account is based on UIS. For hosted server only |
//	+----------------------------+-------------+------------------------------------------------------+
//Returns:
//	- (int64) user identifier
func (sp *SecurityPolicy) AddUser(ctx context.Context, params PUserData) (*PxgValInt, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy.AddUser", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := sp.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//	Modify existing user properties.
//
//	Modifies properties of the specified user
//
//	Parameters:
//	- lUserId	(int64) user id
//	- pUser	(params) user info, containing attributes "KLSPL_USER_*" (see List of user attributes).
func (sp *SecurityPolicy) UpdateUser(ctx context.Context, lUserId int, params PUserData) (*PxgValInt, []byte, error) {
	marshalledData, _ := json.Marshal(params.PUser)
	pUser := []byte(fmt.Sprintf(`{"pUser" : %v,"lUserId" : %d}`, string(marshalledData), lUserId))

	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy.UpdateUser", bytes.NewBuffer(pUser))
	if err != nil {
		log.Fatal(err.Error())
	}

	pxgValInt := new(PxgValInt)
	raw, err := sp.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//UserExInfo struct
type UserInfo struct {
	LUserID   *int64 `json:"lUserId,omitempty"`
	PxgRetVal *bool  `json:"PxgRetVal,omitempty"`
}

type UserInfoEx struct {
	LUserID         *int64      `json:"lUserId,omitempty"`
	NType           *int64      `json:"nType,omitempty"`
	BinSystemID     interface{} `json:"binSystemId"`
	WstrDisplayName *string     `json:"wstrDisplayName,omitempty"`
}

//	Acquire current internal user id.
//
//	Parameters:
//	lUserId	(int64) current user id if it is internal, -1 otherwise
//
//	Returns:
//	- (bool) true if current user is internal user
//
//
//	TODO: may be use input params???
func (sp *SecurityPolicy) GetCurrentUserId(ctx context.Context) (*UserInfo, []byte, error) {
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy.GetCurrentUserId",
		nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(UserInfo)
	raw, err := sp.client.Do(ctx, request, &user)
	return user, raw, err
}

//	Acquire current user id.
//
//	Parameters:
//	- nType	(int64) type of current user:
//	|- 0 for internal user
//	|- 0 1 for non internal user
//	- lUserId	(int64) current internal user id
//	- binSystemId	(binary) current user binary id
//	- wstrDisplayName	(string) current user display name
//
//	For internal user: lUserId > 0;
//	For non internal user: lUserId = -1; binSystemId - binary representation of user SID;
func (sp *SecurityPolicy) GetCurrentUserId2(ctx context.Context) (*UserInfoEx, []byte, error) {
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy.GetCurrentUserId2",
		nil)
	if err != nil {
		return nil, nil, err
	}

	userEx := new(UserInfoEx)
	raw, err := sp.client.Do(ctx, request, &userEx)
	return userEx, raw, err
}

//UsersInfo struct
type UsersInfo struct {
	UsersInfoArray []UsersInfoArray `json:"PxgRetVal"`
}

type UsersInfoArray struct {
	Type           *string         `json:"type,omitempty"`
	UserProperties *UserProperties `json:"value,omitempty"`
}

type UserProperties struct {
	KlsplUserDescription *string `json:"KLSPL_USER_DESCRIPTION,omitempty"`
	KlsplUserEnabled     *bool   `json:"KLSPL_USER_ENABLED,omitempty"`
	KlsplUserFullName    *string `json:"KLSPL_USER_FULL_NAME,omitempty"`
	KlsplUserID          *int64  `json:"KLSPL_USER_ID,omitempty"`
	KlsplUserMail        *string `json:"KLSPL_USER_MAIL,omitempty"`
	KlsplUserName        *string `json:"KLSPL_USER_NAME,omitempty"`
	KlsplUserPhone       *string `json:"KLSPL_USER_PHONE,omitempty"`
	KlsplUserUis         *bool   `json:"KLSPL_USER_UIS,omitempty"`
}

//	Acquire existing user properties.
//
//	Acquires properties of the specified user, or all users if lUserId==(-1);
//
//	Parameters:
//	- lUserId	(int64) user id
//	- lVsId	(int64) user id
//
//	Returns:
//	- (array) users info, an array of containers of attributes "KLSPL_USER_*" (see List of user attributes).
func (sp *SecurityPolicy) GetUsers(ctx context.Context, lUserId, lVsId int64) (*UsersInfo, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"lUserId" : %d, "lVsId" : %d}`, lUserId, lVsId))
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy.GetUsers",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	userInf := new(UsersInfo)
	raw, err := sp.client.Do(ctx, request, &userInf)
	return userInf, raw, err
}

//	Get current user personal data.
//
//	Returns:
//	- (params) personal current user data
func (sp *SecurityPolicy) LoadPerUserData(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy.LoadPerUserData",
		nil)
	if err != nil {
		return nil, err
	}

	raw, err := sp.client.Do(ctx, request, nil)
	return raw, err
}

//TODO SecurityPolicy.SavePerUserData
//TODO SecurityPolicy.UpdateTrustee
func (i *SecurityPolicy) UpdateTrustee(ctx context.Context, params interface{}) ([]byte, error) {
	return nil, nil
}
