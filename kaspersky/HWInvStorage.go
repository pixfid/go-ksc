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

//	Interface for working with Hardware storage subsystem..
//
//	List of all members.
type HWInvStorage service

//	Start export of hardware inventory.
//
//	Parameters:
//	- eExportType	[in]
//
//	Import and export types:
//	+----+-------------+
//	| ID | Description |
//	+----+-------------+
//	|  0 | XML         |
//	|  1 | Excel (XML) |
//	+----+-------------+
//
//	Returns:
//	- Id of asynchronous operation.
//	To get status use AsyncActionStateChecker.CheckActionState, lStateCode "0" means OK.
func (hw *HWInvStorage) ExportHWInvStorage2(ctx context.Context, eExportType int) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"eExportType": %d
	}`, eExportType))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.ExportHWInvStorage2", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
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
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}

func (hw *HWInvStorage) EnumDynColumns(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.EnumDynColumns", nil)
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}

func (hw *HWInvStorage) GetProcessingRules(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.GetProcessingRules", nil)
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}

func (hw *HWInvStorage) GetHWInvObject(ctx context.Context, nObjId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"nObjId": %d
	}`, nObjId))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.GetHWInvObject", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
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
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}
