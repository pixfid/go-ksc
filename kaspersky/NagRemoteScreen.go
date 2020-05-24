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
	"fmt"
	"net/http"
)

//	NagRemoteScreen Class Reference
//
//	Interface for remote screen session management..
//
//	List of all members.
type NagRemoteScreen service

//	Returns existing remote screen sessions.
//
//	Parameters:
//	- nType	(int64) type of remote screen (see Remote screen type)
//
//	╔═══════╦═════════════════════════╦══════════════════════════════════════════╗
//	║ Value ║      Mnemonic name      ║               Description                ║
//	╠═══════╬═════════════════════════╬══════════════════════════════════════════╣
//	║     1 ║ RST_WIN_RDP             ║ Remote desktop                           ║
//	║     2 ║ RST_WIN_DESKTOP_SHARING ║ Windows Desktop Sharing                  ║
//	║     4 ║ RST_VNC                 ║ Virtual Network Computing (VNC)          ║
//	║     8 ║ RST_VNC_HTTP            ║ Virtual Network Computing (VNC) via HTTP ║
//	╚═══════╩═════════════════════════╩══════════════════════════════════════════╝
//
//	Returns:
//	- (array) array of params, each contains KLNAG_RSCR_SESSION_* variables.
//TODO Call NagRemoteScreen.GetExistingSessions for the instance '' (listener '') does not exist (any more?)
func (nrs *NagRemoteScreen) GetExistingSessions(ctx context.Context, nType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nType": %d}`, nType))
	request, err := http.NewRequest("POST", nrs.client.Server+"/api/v1.0/NagRemoteScreen.GetExistingSessions",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := nrs.client.Do(ctx, request, nil)
	return raw, err
}

//	Shares the session, opens ports etc.
//
//	Parameters:
//	- nType	(int64) type of remote screen (see Remote screen type)
//	╔═══════╦═════════════════════════╦══════════════════════════════════════════╗
//	║ Value ║      Mnemonic name      ║               Description                ║
//	╠═══════╬═════════════════════════╬══════════════════════════════════════════╣
//	║     1 ║ RST_WIN_RDP             ║ Remote desktop                           ║
//	║     2 ║ RST_WIN_DESKTOP_SHARING ║ Windows Desktop Sharing                  ║
//	║     4 ║ RST_VNC                 ║ Virtual Network Computing (VNC)          ║
//	║     8 ║ RST_VNC_HTTP            ║ Virtual Network Computing (VNC) via HTTP ║
//	╚═══════╩═════════════════════════╩══════════════════════════════════════════╝
//	- szwID	(string) empty string for RDP, id of session for others
//
//	Returns:
//	- (params) sharing handle of the shared session
//TODO Call NagRemoteScreen.OpenSession for the instance '' (listener '') does not exist (any more?)
func (nrs *NagRemoteScreen) OpenSession(ctx context.Context, nType int64, szwID string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nType": %d, "szwID": "%s"}`, nType, szwID))
	request, err := http.NewRequest("POST", nrs.client.Server+"/api/v1.0/NagRemoteScreen.OpenSession",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := nrs.client.Do(ctx, request, nil)
	return raw, err
}

//TODO func (nrs *NagRemoteScreen) CloseSession
//TODO func (nrs *NagRemoteScreen) GetDataForTunnel
//TODO func (nrs *NagRemoteScreen) GetWdsData
