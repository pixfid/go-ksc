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

type HWInvStorage struct {
	client *Client
}

//Start export of hardware inventory.
//
//Parameters:
//	- eExportType	[in]
//Import and export types:
//	- ID | Description
//	- 0 |XML
//	- 1 |Excel (XML)
//Returns:
//	- Id of asynchronous operation. To get status use AsyncActionStateChecker.CheckActionState, lStateCode "0" means OK.
func (hw *HWInvStorage) ExportHWInvStorage2(ctx context.Context, eExportType int) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"eExportType": %d
	}`, eExportType))

	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.ExportHWInvStorage2", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := hw.client.Do(ctx, request, nil, false)
	return raw, err
}

//Start import of hardware inventory.
//
//Parameters:
//	- eImportType	[in64] (int64)
//
//Import and export types:
//	- ID | Description
//	- 0 |XML
//	- 1 |Excel (XML)
//
//Returns:
//	- Id of asynchronous operation. To get status use AsyncActionStateChecker.CheckActionState, lStateCode "0" means OK.
func (hw *HWInvStorage) ImportHWInvStorage2(ctx context.Context, eImportType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"eImportType": %d
	}`, eImportType))

	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.ImportHWInvStorage2", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := hw.client.Do(ctx, request,nil, false)
	return raw, err
}

func (hw *HWInvStorage) EnumDynColumns(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.EnumDynColumns", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := hw.client.Do(ctx, request, nil, false)
	return raw, err
}

func (hw *HWInvStorage) GetProcessingRules(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.GetProcessingRules", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := hw.client.Do(ctx, request, nil,false)
	return raw, err
}

func (hw *HWInvStorage) GetHWInvObject(ctx context.Context, nObjId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"nObjId": %d
	}`, nObjId))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.GetHWInvObject", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := hw.client.Do(ctx, request,nil,false)
	return raw, err
}

func (hw *HWInvStorage) ExportHWInvStorageGetData(ctx context.Context, wstrAsyncId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"wstrAsyncId": "%s",
	"nGetDataSize": %d
	}`, wstrAsyncId, 10000000))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.ExportHWInvStorageGetData", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := hw.client.Do(ctx, request, nil,false)
	return raw, err
}
