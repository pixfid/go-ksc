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

//	QBTNetworkListApi Class Reference
//
//	Interface to working with Quarantine, Backup and TIF network lists.
//
//	To retrieve info for these lists use List of files from quarantine, backup or unprocessed network lists
//
//	List of all members.
type QBTNetworkListApi service

//NetworkListFileInfo struct
type NetworkListFileInfo struct {
	PPInfo *PPInfo `json:"PxgRetVal,omitempty"`
}

//PPInfo (params) container
type PPInfo struct {
	//Host name. (see "strHostName" attribute from Quarantine, Backup, TIF lists)
	StrHostName string `json:"strHostName,omitempty"`

	//Name of the network list.
	StrID string `json:"strId,omitempty"`

	//Id of an item on the host.
	//(see "strId" attribute from Quarantine, Backup, TIF lists)
	StrListName string `json:"strListName,omitempty"`
}

//	Acquire info about specified file from specified network list.
//
//	Parameters:
//	- itemId	(int64), identifier of the file. (see "nId" attribute from Quarantine, Backup, TIF lists)
//
//	Returns:
//	- ppInfo (params) contains following attributes:
//	|- "strHostName" (string), Host name. (see "strHostName" attribute from Quarantine, Backup, TIF lists)
//	|- "strListName" (string), name of the network list.
//	|- "strId" (string), id of an item on the host. (see "strId" attribute from Quarantine, Backup, TIF lists)
func (nc *QBTNetworkListApi) GetListItemInfo(ctx context.Context, itemId int64) (*NetworkListFileInfo, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"itemId": %d}`, itemId))
	request, err := http.NewRequest("POST", nc.client.Server+"/api/v1.0/QBTNetworkListApi.GetListItemInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	networkListFileInfo := new(NetworkListFileInfo)
	raw, err := nc.client.Do(ctx, request, &networkListFileInfo)
	return networkListFileInfo, raw, err
}

//QBTParams struct
type QBTParams struct {
	//Name of the network list.
	ListName string `json:"listName,omitempty"`

	//Identifier of the file. (see "nId" attribute from Quarantine, Backup, TIF lists)
	ItemID int64 `json:"itemId,omitempty"`

	//Name of the action. List of supported actions:
	//	╔══════════════════╦════════════════════════════════════════════════════╗
	//	║   Action name    ║                    Description                     ║
	//	╠══════════════════╬════════════════════════════════════════════════════╣
	//	║ "RestoreObj"     ║ Restore object from quarantine/backup              ║
	//	║ "DeleteObj"      ║ Delete object from quarantine/backup               ║
	//	║ "ScanQuarantine" ║ Scan object on quarantine. Desinfect object on tif ║
	//	╚══════════════════╩════════════════════════════════════════════════════╝
	TaskName string `json:"taskName,omitempty"`

	//Deprecated: May be empty but should be not null
	PTaskParams PTaskParams `json:"pTaskParams,omitempty"`
}

//PTaskParams struct
type PTaskParams struct{}

//	Initiate action under the specified file from specified network list.
//
//	Parameters:
//	- params	(QBTParams)
func (nc *QBTNetworkListApi) AddListItemTask(ctx context.Context, params QBTParams) ([]byte, error) {
	postData, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", nc.client.Server+"/api/v1.0/QBTNetworkListApi.AddListItemTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := nc.client.Do(ctx, request, nil)
	return raw, err
}

//NetworkTaskParam struct
type QBTsParam struct {
	//Name of the network list.
	ListName string `json:"listName,omitempty"`

	//Collection of the file identifiers.
	//Each entry of the array is (int) identifier of the file. (see "nId" attribute from Quarantine, Backup, TIF lists)
	PItemsIDS []int64 `json:"pItemsIds,omitempty"`

	//Name of the action. List of supported actions:
	//	╔══════════════════╦════════════════════════════════════════════════════╗
	//	║   Action name    ║                    Description                     ║
	//	╠══════════════════╬════════════════════════════════════════════════════╣
	//	║ "RestoreObj"     ║ Restore object from quarantine/backup              ║
	//	║ "DeleteObj"      ║ Delete object from quarantine/backup               ║
	//	║ "ScanQuarantine" ║ Scan object on quarantine. Desinfect object on tif ║
	//	╚══════════════════╩════════════════════════════════════════════════════╝
	TaskName string `json:"taskName,omitempty"`

	//Deprecated: May be empty but should be not null
	PTaskParams PTaskParams `json:"pTaskParams,omitempty"`
}

//Initiate action under the specified files from specified network list.
//
//	Parameters:
//	- params	(QBTsParam)
func (nc *QBTNetworkListApi) AddListItemsTask(ctx context.Context, params QBTsParam) ([]byte, error) {
	postData, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", nc.client.Server+"/api/v1.0/QBTNetworkListApi.AddListItemsTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := nc.client.Do(ctx, request, nil)
	return raw, err
}
