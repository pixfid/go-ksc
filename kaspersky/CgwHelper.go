/*
 *
 * 	Copyright (C) 2020  <Semchenko Aleksandr>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.If not, see <http://www.gnu.org/licenses/>.
 * /
 */

package kaspersky

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
)

//	CgwHelper Class Reference
//
//	CgwHelper (Connection Gateway) helper proxy class.
//
//	Public Member Functions
type CgwHelper struct {
	client *Client
}

//	Retrieves Slave Server Location.
//
//	Parameters:
//	- nSlaveServerId	Slave server ID.
//
//	Returns:
//	- (params) Location params (non-transparent for a user).
func (cp *CgwHelper) GetSlaveServerLocation(ctx context.Context, nSlaveServerId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nSlaveServerId": %d}`, nSlaveServerId))
	request, err := http.NewRequest("POST", cp.client.Server+"/api/v1.0/CgwHelper.GetSlaveServerLocation",
		bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := cp.client.Do(ctx, request, nil)
	return raw, err
}

//	Retrieves Nagent Location by host name.
//
//	Parameters:
//	- wsHostName	Host name.
//
//	Returns:
//	- (params) Location params (non-transparent for a user).
func (cp *CgwHelper) GetNagentLocation(ctx context.Context, wsHostName string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wsHostName": "%s"}`, wsHostName))
	request, err := http.NewRequest("POST", cp.client.Server+"/api/v1.0/CgwHelper.GetNagentLocation",
		bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := cp.client.Do(ctx, request, nil)
	return raw, err
}
