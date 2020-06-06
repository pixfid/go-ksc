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
	"net/http"
)

//	GatewayConnection Class Reference
//
//	Interface for creating gateway connections.
//
//	List of all members.
type GatewayConnection service

type GCParams struct {
	PLocations        []PLocation `json:"pLocations"`
	SzwTargetHostName *string     `json:"szwTargetHostName,omitempty"`
	NTargetPort       *int64      `json:"nTargetPort,omitempty"`
}

type PLocation struct {
	Type           *string      `json:"type,omitempty"`
	PLocationValue *NagLocation `json:"value,omitempty"`
}

type AuthKey struct {
	WstrAuthKey *string `json:"wstrAuthKey,omitempty"`
}

//	Create gateway connection.
//	See Creating gateway connections to know how to create gateway connections.
//	If any connection errors occur this method throws corresponding exception.
//
//	Parameters:
//	- GCParams
//		|- pLocations - (params) gateway nodes locations.
//
//	Return:
//	- wstrAuthKey	- one-time authentication key. Should be used in 'login' or Session.
//
//	StartSession method call in KSCGW authentication scheme. Valid only for 60 seconds.
func (gc *GatewayConnection) PrepareGatewayConnection(ctx context.Context, params GCParams) (*AuthKey, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", gc.client.Server+"/api/v1.0/GatewayConnection.PrepareGatewayConnection",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	authKey := new(AuthKey)
	raw, err := gc.client.Do(ctx, request, &authKey)
	return authKey, raw, err
}

//	Create tunnel connection.
//	See Creating gateway connections to know how to create tunnel connections.
//	If any connection errors occur this method throws corresponding exception.
//
//	Parameters:
//	- GCParams
//		|- pLocations - (params) gateway nodes locations.
//		|- szwTargetHostName (string)
//		|- nTargetPort (int64)
//
//	Return:
//	- wstrAuthKey	- one-time authentication key.
//
//	Should be used in 'login' method call in KSCGW authentication scheme. Valid only for 60 seconds.
func (gc *GatewayConnection) PrepareTunnelConnection(ctx context.Context, params GCParams) (*AuthKey, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", gc.client.Server+"/api/v1.0/GatewayConnection.PrepareTunnelConnection",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	authKey := new(AuthKey)
	raw, err := gc.client.Do(ctx, request, &authKey)
	return authKey, raw, err
}
