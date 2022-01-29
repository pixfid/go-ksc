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

// ServerHierarchy Server hierarchy management service.
type ServerHierarchy service

// DelServer Remove specified slave server.
//
// This action only removes slave server registration info from master server. To remove master server settings from slave server use HostGroup.SSWrite
// to overwrite master server connection settings section and set "KLSRV_MASTER_SRV_USE" to false.
func (sh *ServerHierarchy) DelServer(ctx context.Context, lServer int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lServer": %d}`, lServer))
	request, err := http.NewRequest("POST", sh.client.Server+"/api/v1.0/ServerHierarchy.DelServer", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, _ := sh.client.Request(ctx, request, nil)
	return raw, err
}

type ServerHierarchyParams struct {
	// LServer Slave server id

	LServer int `json:"lServer"`
	// PFields attributes to return. See list of slave server attributes for attributes list and description
	//
	// "KLSRVH_SRV_ID" Slave server id
	//
	// "KLSRVH_SRV_INST_ID" Slave server instance id
	//
	// "KLSRVH_SRV_ADDR" Slave server address and port
	//
	// "KLSRVH_SRV_DN" Display name
	//
	// "KLSRVH_SRV_GROUPID" Id of administration group where the slave server is located
	//
	// "KLSRVH_SRV_CERTIFICATE" Slave server certificate.
	//
	// "KLSRVH_SRV_PUBLIC_KEY_HASH" Slave server certificate MD5-hash
	//
	// "KLSRVH_SRV_STATUS" Slave server status: 0 means "Inactive", 1 means "Active".
	//
	// "KLSRVH_SRV_VERSION" Slave server version
	//
	// "KLSRVH_SRV_PASSIVE" Flag set if the slave is passive (does not connect to server, but accepts master connections instead)
	//
	// "KLSRVH_SRV_LAST_CONNECTED" Time when server was available last time
	//
	// "KLSRVH_SRV_MASTER_ADDR" Master server connection address, valid for non-passive slaves
	//
	// "KLSRVH_SRV_HOST_GUID" Slave server host identity
	PFields []string `json:"pFields"`
}

// GetServerInfo Acquire specified slave server attributes.
func (sh *ServerHierarchy) GetServerInfo(ctx context.Context, params ServerHierarchyParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sh.client.Server+"/api/v1.0/ServerHierarchy.GetServerInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sh.client.Request(ctx, request, nil)
	return raw, nil
}

// GetChildServers Enumerate slave servers for specified group.
func (sh *ServerHierarchy) GetChildServers(ctx context.Context, nGroupId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nGroupId": %d}`, nGroupId))
	request, err := http.NewRequest("POST", sh.client.Server+"/api/v1.0/ServerHierarchy.GetChildServers", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sh.client.Request(ctx, request, nil)
	if err != nil {
		return nil, err
	}
	return raw, err
}

// FindSlaveServers Searches for slave servers meeting specified criteria.
func (sh *ServerHierarchy) FindSlaveServers(ctx context.Context, params PFindParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sh.client.Server+"/api/v1.0/ServerHierarchy.FindSlaveServers", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sh.client.Request(ctx, request, nil)
	if err != nil {
		return nil, err
	}

	return raw, nil
}
