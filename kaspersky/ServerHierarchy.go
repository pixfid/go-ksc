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

//	ServerHierarchy Class Reference
//
//	Server hierarchy management interface.
//
//	List of all members.
type ServerHierarchy service

//	Remove specified slave server.
//
//	This action only removes slave server registration info from master server.
//	To remove master server settings from slave server use HostGroup.SS_Write
//	to overwrite master server connection settings section and set "KLSRV_MASTER_SRV_USE" to false.
//	See Slave server registration for details.
//
//	Parameters:
//	- lServer	Slave server id
func (sh *ServerHierarchy) DelServer(ctx context.Context, lServer int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"lServer": %d
	}`, lServer))
	request, err := http.NewRequest("POST", sh.client.Server+"/api/v1.0/ServerHierarchy.DelServer", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, _ := sh.client.Do(ctx, request, nil)
	return raw, err
}

//	Acquire specified slave server attributes.
//
//	Parameters:
//	- lServer	Slave server id
//	- pFields	paramArray array of string attributes to return.
//	See list of slave server attributes for attributes list and description.
//
//	Returns:
//	- paramParams container with specified attributes of slave server
//
//	+----------------------------+----------+-------------------------------------------------------------------------------------------------------+
//	|         Attributes         |   Type   |                                              Description                                              |
//	+----------------------------+----------+-------------------------------------------------------------------------------------------------------+
//	| KLSRVH_SRV_ID              | int64    | Slave server id                                                                                       |
//	| KLSRVH_SRV_INST_ID         | string   | Slave server instance id                                                                              |
//	| KLSRVH_SRV_ADDR            | string   | Slave server address and port                                                                         |
//	| KLSRVH_SRV_DN              | string   | Display name                                                                                          |
//	| KLSRVH_SRV_GROUPID         | int64    | Id of administration group where the slave server is located                                          |
//	| KLSRVH_SRV_CERTIFICATE     | binary   | Slave server certificate.                                                                             |
//	| KLSRVH_SRV_PUBLIC_KEY_HASH | string   | Slave server certificate MD5-hash                                                                     |
//	| KLSRVH_SRV_STATUS          | int64    | Slave server status: 0 means "Inactive", 1 means "Active".                                            |
//	| KLSRVH_SRV_VERSION         | string   | Slave server version                                                                                  |
//	| KLSRVH_SRV_PASSIVE         | bool     | Flag set if the slave is passive (does not connect to server, but accepts master connections instead) |
//	| KLSRVH_SRV_LAST_CONNECTED  | DateTime | Time when server was available last time                                                              |
//	| KLSRVH_SRV_MASTER_ADDR     | string   | Master server connection address, valid for non-passive slaves                                        |
//	| KLSRVH_SRV_HOST_GUID       | string   | Slave server host identity                                                                            |
//	+----------------------------+----------+-------------------------------------------------------------------------------------------------------+
//
//	{
//		"lServer": 1, //Slave server ID
//		"pFields": //Attributes
//			[
//				"KLSRVH_SRV_ID",
//				"KLSRVH_SRV_INST_ID",
//				"KLSRVH_SRV_ADDR",
//				"... other"
//			]
//	}
func (sh *ServerHierarchy) GetServerInfo(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", sh.client.Server+"/api/v1.0/ServerHierarchy.GetServerInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sh.client.Do(ctx, request, nil)
	return raw, nil
}

//	Enumerate slave servers for specified group.
//
//	Parameters:
//	- nGroupId	administration group id where slave server located or -1 to acquire slave servers from all groups
//	Returns:
//	paramArray of paramParams containing following slave server attributes:
//
//	+----------------------------+----------+-------------------------------------------------------------------------------------------------------+
//	|         Attributes         |   Type   |                                              Description                                              |
//	+----------------------------+----------+-------------------------------------------------------------------------------------------------------+
//	| KLSRVH_SRV_ID              | int64    | Slave server id                                                                                       |
//	| KLSRVH_SRV_INST_ID         | string   | Slave server instance id                                                                              |
//	| KLSRVH_SRV_ADDR            | string   | Slave server address and port                                                                         |
//	| KLSRVH_SRV_DN              | string   | Display name                                                                                          |
//	| KLSRVH_SRV_GROUPID         | int64    | Id of administration group where the slave server is located                                          |
//	| KLSRVH_SRV_CERTIFICATE     | binary   | Slave server certificate.                                                                             |
//	| KLSRVH_SRV_PUBLIC_KEY_HASH | string   | Slave server certificate MD5-hash                                                                     |
//	| KLSRVH_SRV_STATUS          | int64    | Slave server status: 0 means "Inactive", 1 means "Active".                                            |
//	| KLSRVH_SRV_VERSION         | string   | Slave server version                                                                                  |
//	| KLSRVH_SRV_PASSIVE         | bool     | Flag set if the slave is passive (does not connect to server, but accepts master connections instead) |
//	| KLSRVH_SRV_LAST_CONNECTED  | DateTime | Time when server was available last time                                                              |
//	| KLSRVH_SRV_MASTER_ADDR     | string   | Master server connection address, valid for non-passive slaves                                        |
//	| KLSRVH_SRV_HOST_GUID       | string   | Slave server host identity                                                                            |
//	+----------------------------+----------+-------------------------------------------------------------------------------------------------------+
//	See list of slave server attributes for attributes description.
func (sh *ServerHierarchy) GetChildServers(ctx context.Context, nGroupId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"nGroupId": %d
	}`, nGroupId))
	request, err := http.NewRequest("POST", sh.client.Server+"/api/v1.0/ServerHierarchy.GetChildServers", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, _ := sh.client.Do(ctx, request, nil)
	return raw, err
}
