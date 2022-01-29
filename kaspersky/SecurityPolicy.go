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

// SecurityPolicy Allows to manage users and permissions.
type SecurityPolicy service

// PUserData struct
type PUserData struct {
	PUser PUser `json:"pUser"`
}

// PUser struct
type PUser struct {
	// KlsplUserName User name
	KlsplUserName string `json:"KLSPL_USER_NAME"`

	// KlsplUserPwdEncrypted Encrypted user password
	KlsplUserPwdEncrypted string `json:"KLSPL_USER_PWD_ENCRYPTED,omitempty"`

	// KlsplUserFullName User full name
	KlsplUserFullName string `json:"KLSPL_USER_FULL_NAME"`

	// KlsplUserDescription User description
	KlsplUserDescription string `json:"KLSPL_USER_DESCRIPTION"`

	// KlsplUserMail User mail
	KlsplUserMail string `json:"KLSPL_USER_MAIL"`

	// KlsplUserPhone User phone
	KlsplUserPhone string `json:"KLSPL_USER_PHONE"`

	// KlsplUserEnabled User account is enabled if true.
	KlsplUserEnabled bool `json:"KLSPL_USER_ENABLED"`
}

// AddUser Add new user.
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
	raw, err := sp.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// UpdateUser Modify existing user properties.
//
// Modifies properties of the specified user
func (sp *SecurityPolicy) UpdateUser(ctx context.Context, lUserId int, params PUserData) (*PxgValInt, []byte, error) {
	marshalledData, _ := json.Marshal(params.PUser)
	pUser := []byte(fmt.Sprintf(`{"pUser" : %v,"lUserId" : %d}`, string(marshalledData), lUserId))

	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy.UpdateUser", bytes.NewBuffer(pUser))
	if err != nil {
		log.Fatal(err.Error())
	}

	pxgValInt := new(PxgValInt)
	raw, err := sp.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//UserExInfo struct
type UserInfo struct {
	LUserID   int64 `json:"lUserId,omitempty"`
	PxgRetVal bool  `json:"PxgRetVal,omitempty"`
}

// UserInfoEx struct
type UserInfoEx struct {
	LUserID         int64       `json:"lUserId,omitempty"`
	NType           int64       `json:"nType,omitempty"`
	BinSystemID     interface{} `json:"binSystemId"`
	WstrDisplayName string      `json:"wstrDisplayName,omitempty"`
}

// GetCurrentUserId Acquire current internal user id.
func (sp *SecurityPolicy) GetCurrentUserId(ctx context.Context) (*UserInfo, []byte, error) {
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy.GetCurrentUserId",
		nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(UserInfo)
	raw, err := sp.client.Request(ctx, request, &user)
	return user, raw, err
}

// GetCurrentUserId2 Acquire current user id.
//
// For internal user: lUserId > 0;
//
// For non internal user: lUserId = -1; binSystemId - binary representation of user SID;
func (sp *SecurityPolicy) GetCurrentUserId2(ctx context.Context) (*UserInfoEx, []byte, error) {
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy.GetCurrentUserId2",
		nil)
	if err != nil {
		return nil, nil, err
	}

	userEx := new(UserInfoEx)
	raw, err := sp.client.Request(ctx, request, &userEx)
	return userEx, raw, err
}

//UsersInfo struct
type UsersInfo struct {
	UsersInfoArray []UsersInfoArray `json:"PxgRetVal"`
}

// UsersInfoArray struct
type UsersInfoArray struct {
	Type           string          `json:"type,omitempty"`
	UserProperties *UserProperties `json:"value,omitempty"`
}

// UserProperties struct
type UserProperties struct {
	// KlsplUserDescription User description
	KlsplUserDescription string `json:"KLSPL_USER_DESCRIPTION,omitempty"`

	// KlsplUserEnabled User account is enabled if true.
	KlsplUserEnabled bool `json:"KLSPL_USER_ENABLED,omitempty"`

	// KlsplUserFullName User full name
	KlsplUserFullName string `json:"KLSPL_USER_FULL_NAME,omitempty"`

	// KlsplUserID User id
	KlsplUserID int64 `json:"KLSPL_USER_ID,omitempty"`

	// KlsplUserMail User mail
	KlsplUserMail string `json:"KLSPL_USER_MAIL,omitempty"`

	// KlsplUserName User name
	KlsplUserName string `json:"KLSPL_USER_NAME,omitempty"`

	// KlsplUserPhone User phone
	KlsplUserPhone string `json:"KLSPL_USER_PHONE,omitempty"`

	// KlsplUserUis User account is based on UIS. For hosted server only
	KlsplUserUis bool `json:"KLSPL_USER_UIS,omitempty"`
}

// GetUsers Acquire existing user properties.
//
// Acquires properties of the specified user, or all users if lUserId==(-1);
func (sp *SecurityPolicy) GetUsers(ctx context.Context, lUserId, lVsId int64) (*UsersInfo, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"lUserId" : %d, "lVsId" : %d}`, lUserId, lVsId))
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy.GetUsers",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	userInf := new(UsersInfo)
	raw, err := sp.client.Request(ctx, request, &userInf)
	return userInf, raw, err
}

// LoadPerUserData Get current user personal data.
func (sp *SecurityPolicy) LoadPerUserData(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy.LoadPerUserData",
		nil)
	if err != nil {
		return nil, err
	}

	raw, err := sp.client.Request(ctx, request, nil)
	return raw, err
}

// SavePerUserData Save or replace current user personal data.
func (sp *SecurityPolicy) SavePerUserData(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy.SavePerUserData",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sp.client.Request(ctx, request, nil)
	return raw, err
}

// TrusteeParam struct
type TrusteeParam struct {
	LlTrusteeID  int64         `json:"llTrusteeId,omitempty"`
	PTrusteeData *PTrusteeData `json:"pUserData,omitempty"`
}

// PTrusteeData struct
type PTrusteeData struct {
	//KlsplUserID User id
	KlsplUserID int64 `json:"KLSPL_USER_ID,omitempty"`

	// KlsplUserName User name
	KlsplUserName string `json:"KLSPL_USER_NAME,omitempty"`

	// KlsplUserPwd Plain text user password. Obsolete
	KlsplUserPwd string `json:"KLSPL_USER_PWD,omitempty"`

	//KlsplUserPwdEncrypted Encrypted user password
	KlsplUserPwdEncrypted string `json:"KLSPL_USER_PWD_ENCRYPTED,omitempty"`

	// KlsplUserFullName User full name
	KlsplUserFullName string `json:"KLSPL_USER_FULL_NAME,omitempty"`

	// KlsplUserDescription User description
	KlsplUserDescription string `json:"KLSPL_USER_DESCRIPTION,omitempty"`

	// KlsplUserMail User mail
	KlsplUserMail string `json:"KLSPL_USER_MAIL,omitempty"`

	// KlsplUserPhone User phone
	KlsplUserPhone string `json:"KLSPL_USER_PHONE,omitempty"`

	// KlsplUserEnabled User account is enabled if true
	KlsplUserEnabled bool `json:"KLSPL_USER_ENABLED,omitempty"`

	// KlsplUserUis User account is based on UIS. For hosted server only
	KlsplUserUis bool `json:"KLSPL_USER_UIS,omitempty"`
}

// UpdateTrustee Modifies properties of the specified user (either internal user or user and group from Active Directory);
// for internal groups use SecurityPolicy3.UpdateSecurityGroup.
func (sp *SecurityPolicy) UpdateTrustee(ctx context.Context, params TrusteeParam) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", sp.client.Server+"/api/v1.0/SecurityPolicy.UpdateTrustee",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sp.client.Request(ctx, request, nil)
	return raw, err
}
