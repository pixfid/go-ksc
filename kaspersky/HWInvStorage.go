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
//	╔════╦═════════════╗
//	║ ID ║ Description ║
//	╠════╬═════════════╣
//	║  0 ║ XML         ║
//	║  1 ║ Excel (XML) ║
//	╚════╩═════════════╝
//
//	Returns:
//	- Id of asynchronous operation.
//	To get status use AsyncActionStateChecker.CheckActionState, lStateCode "0" means OK.
func (hw *HWInvStorage) ExportHWInvStorage2(ctx context.Context, eExportType int) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"eExportType": %d}`, eExportType))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.ExportHWInvStorage2", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := hw.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//	Cancel export of hardware inventory.
//
//	Parameters:
//	- wstrAsyncId	(string) Async id returned from HWInvStorage.ExportHWInvStorage2
func (hw *HWInvStorage) ExportHWInvStorageCancel(ctx context.Context, wstrAsyncId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrAsyncId": "%s"}`, wstrAsyncId))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.ExportHWInvStorageCancel", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}

//	Start import of hardware inventory.
//
//	Parameters:
//	- eImportType	[in64] (int64)
//
//	Import and export types:
//	╔════╦═════════════╗
//	║ ID ║ Description ║
//	╠════╬═════════════╣
//	║  0 ║ XML         ║
//	║  1 ║ Excel (XML) ║
//	╚════╩═════════════╝
//
//	Returns:
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
//	- arrDynColumnInfo	(array) Array of params. Format of element:
//	|- DynColId (string) - Column id
//	|- DynColName (string) - Column name
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

//Processing rules container
type ProcessingRules struct {
	PRules *PRules `json:"pRules,omitempty"`
}

//Processing rules
type PRules struct {
	//Array of Object type. See Object types. Type of each element is int64.
	AutoCorporativeByType []int64 `json:"AutoCorporativeByType"`
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
func (hw *HWInvStorage) GetProcessingRules(ctx context.Context) (*ProcessingRules, []byte, error) {
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.GetProcessingRules", nil)
	if err != nil {
		return nil, nil, err
	}

	processingRules := new(ProcessingRules)
	raw, err := hw.client.Do(ctx, request, &processingRules)
	return processingRules, raw, err
}

//	Set processing rules.
//
//	Parameters:
//	- pRules	(params) See Processing rules format
//
//	Exceptions:
//	- KLSTD::STDE_NOACCESS	- access denied
func (hw *HWInvStorage) SetProcessingRules(ctx context.Context, params ProcessingRules) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.SetProcessingRules", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}

//	Get hardware inventory object.
//
//	Parameters:
//	- nObjId 	(int64) Object id
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

//HWInvStorageResponse struct
type HWInvStorageResponse struct {
	//Data chunk
	PChunk string `json:"pChunk,omitempty"`

	//Actual size of retrieved data
	NGotDataSize int64 `json:"nGotDataSize,omitempty"`

	//Size of not retrieved data
	NDataSizeREST int64 `json:"nDataSizeRest,omitempty"`
}

//	Get exported data. Call this method until nDataSizeRest is not zero.
//
//	Parameters:
//	- wstrAsyncId	(string) Async id returned from HWInvStorage.ExportHWInvStorage2
//	- nGetDataSize	(int64) Max data size to retrieve
//
//	Return:
//	- response	HWInvStorageResponse
func (hw *HWInvStorage) ExportHWInvStorageGetData(ctx context.Context, wstrAsyncId string,
	nGetDataSize int64) (*HWInvStorageResponse, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrAsyncId": "%s","nGetDataSize": %d}`, wstrAsyncId, nGetDataSize))

	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.ExportHWInvStorageGetData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	hwInvStorageResponse := new(HWInvStorageResponse)
	raw, err := hw.client.Do(ctx, request, &hwInvStorageResponse)
	return hwInvStorageResponse, raw, err
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

//	Set corporative flag for array of devices.
//
//	Parameters:
//	- arrObjId	(array) Array of device ids. Max array size is 1000 elements.
//	For device id see Hardware inventory object format hardware object format.
//	- bState	(bool) New state
//
//	Exceptions:
//	- KLSTD::STDE_NOACCESS	- access denied
func (hw *HWInvStorage) SetCorpFlag2(ctx context.Context, arrObjId []int64, bState bool) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"arrObjId": %s, "bState" : %v}`, ToJson(arrObjId), bState))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.SetCorpFlag2", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}
