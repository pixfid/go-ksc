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

// HWInvStorage service for working with Hardware storage subsystem.
type HWInvStorage service

// AddDynColumn Add dynamic column.
func (hw *HWInvStorage) AddDynColumn(ctx context.Context, wstrColName string) (*PxgValStr, error) {
	postData := []byte(fmt.Sprintf(`{"wstrColName": "%s"}`, wstrColName))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.AddDynColumn", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgValStr := new(PxgValStr)
	_, err = hw.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, err
}

// ExportHWInvStorage2 Start export of hardware inventory.
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

// ExportHWInvStorageCancel Cancel export of hardware inventory.
func (hw *HWInvStorage) ExportHWInvStorageCancel(ctx context.Context, wstrAsyncId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrAsyncId": "%s"}`, wstrAsyncId))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.ExportHWInvStorageCancel", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}

// EnumDynColumns Start import of hardware inventory.
func (hw *HWInvStorage) ImportHWInvStorage2(ctx context.Context, eImportType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"eImportType": %d}`, eImportType))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.ImportHWInvStorage2", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}

// DynamicColumns struct
type DynamicColumns struct {
	// ArrDynColumnInfo Array of params.
	ArrDynColumnInfo []ArrDynColumnInfo `json:"arrDynColumnInfo"`
}

// ArrDynColumnInfo struct
type ArrDynColumnInfo struct {
	Type  *string             `json:"type,omitempty"`
	Value *DynColumnInfoValue `json:"value,omitempty"`
}

// DynColumnInfoValue struct
type DynColumnInfoValue struct {
	// DynColID Column id
	DynColID *string `json:"DynColId,omitempty"`

	// DynColName Column name
	DynColName *string `json:"DynColName,omitempty"`
}

// EnumDynColumns Return list of dynamic columns.
func (hw *HWInvStorage) EnumDynColumns(ctx context.Context) (*DynamicColumns, error) {
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.EnumDynColumns", nil)
	if err != nil {
		return nil, err
	}

	dynamicColumns := new(DynamicColumns)
	_, err = hw.client.Do(ctx, request, &dynamicColumns)
	return dynamicColumns, err
}

// ProcessingRules Processing rules container
type ProcessingRules struct {
	PRules *PRules `json:"pRules,omitempty"`
}

// Processing rules
type PRules struct {
	// AutoCorporativeByType Array of Object type. See Object types. Type of each element is int64.
	AutoCorporativeByType []int64 `json:"AutoCorporativeByType"`
}

// GetProcessingRules Get processing rules.
func (hw *HWInvStorage) GetProcessingRules(ctx context.Context) (*ProcessingRules, []byte, error) {
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.GetProcessingRules", nil)
	if err != nil {
		return nil, nil, err
	}

	processingRules := new(ProcessingRules)
	raw, err := hw.client.Do(ctx, request, &processingRules)
	return processingRules, raw, err
}

// SetProcessingRules Set processing rules.
func (hw *HWInvStorage) SetProcessingRules(ctx context.Context, params ProcessingRules) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.SetProcessingRules", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}

// GetHWInvObject Get hardware inventory object.
func (hw *HWInvStorage) GetHWInvObject(ctx context.Context, nObjId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nObjId": %d}`, nObjId))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.GetHWInvObject", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}

// HWInvStorageResponse struct
type HWInvStorageResponse struct {
	//Data chunk
	PChunk string `json:"pChunk,omitempty"`

	//Actual size of retrieved data
	NGotDataSize int64 `json:"nGotDataSize,omitempty"`

	//Size of not retrieved data
	NDataSizeREST int64 `json:"nDataSizeRest,omitempty"`
}

// ExportHWInvStorageGetData Get exported data. Call this method until nDataSizeRest is not zero.
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

// DelHWInvObject Delete hardware inventory object.
func (hw *HWInvStorage) DelHWInvObject(ctx context.Context, nObjId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nObjId": %d}`, nObjId))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.DelHWInvObject", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}

// DelHWInvObject2 Delete array of objects.
func (hw *HWInvStorage) DelHWInvObject2(ctx context.Context, arrObjId []int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"arrObjId": %s}`, ToJson(arrObjId)))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.DelHWInvObject2", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}

// SetCorpFlag2 Set corporative flag for array of devices.
func (hw *HWInvStorage) SetCorpFlag2(ctx context.Context, arrObjId []int64, bState bool) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"arrObjId": %s, "bState" : %v}`, ToJson(arrObjId), bState))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.SetCorpFlag2", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hw.client.Do(ctx, request, nil)
	return raw, err
}
