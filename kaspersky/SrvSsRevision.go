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
	"net/http"
)

//	SrvSsRevision Class Reference
//
//	Access to virtual server settings storage revisions. More...
//
//	List of all members.
type SrvSsRevision service

//	Open specified version of virtual server settings storage.
//
//	Parameters:
//	- nVServer	id of the virtual server, zero means the 'main server'
//	- nRevision	revision, zero means 'current version'
//	- szwType	only "SS_SETTINGS" is supported
//
//	Returns:
//	opened settings storage identifier, it must be closed via SrvSsRevision.SsRevision_Close method
func (ssr *SrvSsRevision) SsRevision_Open(ctx context.Context, nVServer, nRevision int64, szwType string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{
		"nVServer": %d, 
		"nRevision": %d, 
		"szwType": "%s"
	}`, nVServer, nRevision, szwType))
	request, err := http.NewRequest("POST", ssr.client.Server+"/api/v1.0/SrvSsRevision.SsRevision_Open", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ssr.client.Do(ctx, request, nil)
	return raw, err
}

//	Close settings storage opened by SrvSsRevision.SsRevision_Open
//
//	Parameters:
//	- szwId	settings storage identifier returned by SrvSsRevision.SsRevision_Open method
func (ssr *SrvSsRevision) SsRevision_Close(ctx context.Context, szwType string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{
		"szwType": "%s"
	}`, szwType))
	request, err := http.NewRequest("POST", ssr.client.Server+"/api/v1.0/SrvSsRevision.SsRevision_Close", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ssr.client.Do(ctx, request, nil)
	return raw, err
}
