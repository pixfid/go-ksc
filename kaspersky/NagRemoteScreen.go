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

// NagRemoteScreen Interface for remote screen session management..
type NagRemoteScreen service

// ExistingSessions struct
type ExistingSessions struct {
	ExistingSessions []ExistingSession `json:"PxgRetVal"`
}

// ExistingSession struct
type ExistingSession struct {
	Type         string       `json:"type"`
	SessionValue SessionValue `json:"value"`
}

// SessionValue struct
type SessionValue struct {
	KlnagRscrSessionDN string `json:"KLNAG_RSCR_SESSION_DN"`
	KlnagRscrSessionID string `json:"KLNAG_RSCR_SESSION_ID"`
}

// GetExistingSessions Returns existing remote screen sessions.
func (nrs *NagRemoteScreen) GetExistingSessions(ctx context.Context, nType int64) (*ExistingSessions, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nType": %d}`, nType))
	request, err := http.NewRequest("POST", nrs.client.Server+"/api/v1.0/NagRemoteScreen.GetExistingSessions",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	existingSessions := new(ExistingSessions)
	raw, err := nrs.client.Do(ctx, request, &existingSessions)
	return existingSessions, raw, err
}

// SessionHandle struct using in
type SessionHandle struct {
	PSharingHandle PSharingHandle `json:"PxgRetVal"`
}

// PSharingHandle struct
type PSharingHandle struct {
	KlnagRscrHandleID   int64  `json:"KLNAG_RSCR_HANDLE_ID"`
	KlnagRscrHandleType int64  `json:"KLNAG_RSCR_HANDLE_TYPE"`
	KlnagRscrHostname   string `json:"KLNAG_RSCR_HOSTNAME"`
	KlnagRscrPort       int64  `json:"KLNAG_RSCR_PORT"`
}

// OpenSession Shares the session, opens ports etc.
func (nrs *NagRemoteScreen) OpenSession(ctx context.Context, nType int64, szwID string) (*SessionHandle, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"nType": %d, "szwID": "%s"}`, nType, szwID))
	request, err := http.NewRequest("POST", nrs.client.Server+"/api/v1.0/NagRemoteScreen.OpenSession",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	sessionHandle := new(SessionHandle)
	raw, err := nrs.client.Do(ctx, request, &sessionHandle)
	return sessionHandle, raw, err
}

// SharingHandle struct using in NagRemoteScreen.CloseSession
type SharingHandle struct {
	//PSharingHandle value of the sharing handle returned by OpenSession
	PSharingHandle PSharingHandle `json:"pSharingHandle"`
}

// CloseSession Closes session.
func (nrs *NagRemoteScreen) CloseSession(ctx context.Context, params SharingHandle) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", nrs.client.Server+"/api/v1.0/NagRemoteScreen.CloseSession",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := nrs.client.Do(ctx, request, nil)
	return raw, err
}

// TunnelData using in NagRemoteScreen.GetDataForTunnel
type TunnelData struct {
	//NHostPortNumber nHostPortNumber
	NHostPortNumber int64 `json:"nHostPortNumber"`

	//WstrHostNameOrIPAddr wstrHostNameOrIpAddr
	WstrHostNameOrIPAddr string `json:"wstrHostNameOrIpAddr"`
}

// GetDataForTunnel Returns data to create an use tunnel
func (nrs *NagRemoteScreen) GetDataForTunnel(ctx context.Context, params SharingHandle) (*TunnelData, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", nrs.client.Server+"/api/v1.0/NagRemoteScreen.GetDataForTunnel",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	tunnelData := new(TunnelData)
	raw, err := nrs.client.Do(ctx, request, &tunnelData)
	return tunnelData, raw, err
}

// WdsDataParams using in NagRemoteScreen.GetWdsData
type WdsDataParams struct {
	PSharingHandle   PSharingHandle `json:"pSharingHandle"`
	NLocalPortNumber int64          `json:"nLocalPortNumber"`
}

// GetWdsData Returns data specific for Windows Desktop Sharing
func (nrs *NagRemoteScreen) GetWdsData(ctx context.Context, params WdsDataParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", nrs.client.Server+"/api/v1.0/NagRemoteScreen.GetWdsData",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := nrs.client.Do(ctx, request, nil)
	return raw, err
}
