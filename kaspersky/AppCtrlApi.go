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
	"encoding/json"
	"net/http"
)

//	AppCtrlApi Class Reference
//
//	Interface to get info about execution files.
//
//	Public Member Functions
type AppCtrlApi service

type ExeFileInfoParams struct {
	SzwHostID  *string     `json:"szwHostId,omitempty"`
	LFileID    *int64      `json:"lFileId,omitempty"`
	ExePFilter *ExePFilter `json:"pFilter,omitempty"`
}

type ExePFilter struct {
	FileID   *string `json:"FILE_ID,omitempty"`
	FileName *string `json:"FILE_NAME,omitempty"`
}

//	Get data about instances of the execution file on the host.
//
//	Parameters:
//	- szwHostId	(string) host name, a unique server-generated string (see KLHST_WKS_HOSTNAME attribute).
//	It is NOT the same as computer network name (DNS-, FQDN-, NetBIOS-name) . If empty then will be returned only attributes from List of common attributes of execution file from AppControl
//	- lFileId	(long) The file identifier ( see FILE_ID attribute ). It is id from database,
//	so it is internal id and it is valid only for this SC-server )
//	- pFilter	(params) Specify set of the fields that should be returned.
//
//To do this for field 'FieldName' it is needed to add into pFilter the value of any type with name 'FieldName'
//If NULL than all possible fields will be returned.
//
//	Examples:
//	to get all possible fields: use NULL
//	to get fields "FILE_ID" and "FILE_NAME" use:
//                            +--"FILE_ID" = (string)""
//                            |
//                            +--"FILE_NAME" = (string)""
//
func (ac *AppCtrlApi) GetExeFileInfo(ctx context.Context, params ExeFileInfoParams) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ac.client.Server+"/api/v1.0/AppCtrlApi.GetExeFileInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ac.client.Do(ctx, request, nil)
	return raw, err
}
