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

//AppCtrlApi Class Reference
//
//Interface to get info about execution files.
//
//Public Member Functions
type AppCtrlApi service

type ExeFileInfoParams struct {
	SzwHostID  string      `json:"szwHostId,omitempty"`
	LFileID    int64       `json:"lFileId,omitempty"`
	ExePFilter *ExePFilter `json:"pFilter,omitempty"`
}

type ExePFilter struct {
	FileID   string `json:"FILE_ID,omitempty"`
	FileName string `json:"FILE_NAME,omitempty"`
}

//GetExeFileInfo
//Get data about instances of the execution file on the host.
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
