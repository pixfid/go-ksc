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
	postData := []byte(fmt.Sprintf(`{"eImportType": %d}`, eImportType))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.ImportHWInvStorage2", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}

//	Return list of dynamic columns.
//
//	Parameters:
//	- arrDynColumnInfo	[out] (array) Array of params. Format of element:
//	|- DynColId (wstring) - Column id
//	|- DynColName (wstring) - Column name
//
//	Exceptions:
//	- KLSTD::STDE_NOACCESS	- access denied
func (hw *HWInvStorage) EnumDynColumns(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.EnumDynColumns", nil)
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}

//	Get processing rules.
//
//	Parameters:
//
//	Return:
//	- pRules (params) See Processing rules format
//
//	Exceptions:
//	- KLSTD::STDE_NOACCESS	- access denied
func (hw *HWInvStorage) GetProcessingRules(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.GetProcessingRules", nil)
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}

//	Get hardware inventory object.
//
//	Parameters:
//	- nObjId	(int64) Object id
//
//	Return:
//	- pObj 	(params) Object body. See Object format
//
//	Exceptions:
//	- KLSTD::STDE_NOACCESS	- access denied
//	- KLSTD::STDE_NOTFOUND	- Object not found
func (hw *HWInvStorage) GetHWInvObject(ctx context.Context, nObjId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nObjId": %d}`, nObjId))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.GetHWInvObject", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}

//	Get exported data. Call this method until nDataSizeRest is not zero.
//
//	Parameters:
//	- wstrAsyncId	(string) Async id returned from HWInvStorage.ExportHWInvStorage2
//	- nGetDataSize	(int64) Max data size to retrieve
//
//	Return:
//	- pChunk	(binary) Data chunk
//	- nGotDataSize	(int64) Actual size of retrieved data
//	- nDataSizeRest	(int64) Size of not retrieved data
func (hw *HWInvStorage) ExportHWInvStorageGetData(ctx context.Context, wstrAsyncId string, nGetDataSize int64) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"wstrAsyncId": "%s","nGetDataSize": %d}`, wstrAsyncId, nGetDataSize))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.ExportHWInvStorageGetData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}

//	Delete hardware inventory object.
//
//	Parameters:
//	- nObjId	(int64) Object id
//
//	Exceptions:
//	- KLSTD::STDE_NOACCESS	- access denied
func (hw *HWInvStorage) DelHWInvObject(ctx context.Context, nObjId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nObjId": %d}`, nObjId))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.DelHWInvObject", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}

//	Delete array of objects.
//
//	Parameters:
//	- arrObjId	(array) Array of device ids. Max array size is 1000 elements.
//	For device id see Hardware inventory object format hardware object format.
//
//	Exceptions:
//	- KLSTD::STDE_NOACCESS	- access denied
func (hw *HWInvStorage) DelHWInvObject2(ctx context.Context, arrObjId []int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"arrObjId": %s}`, ToJson(arrObjId)))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.DelHWInvObject2", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}
