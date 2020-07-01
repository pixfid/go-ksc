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

type PpObj struct {
	PPObj HWInvObject `json:"pObj"`
}

// HWInvObject struct
type HWInvObject struct {
	Type         int64    `json:"Type,omitempty"`
	SubType      int64    `json:"SubType,omitempty"`
	Created      DateTime `json:"Created,omitempty"`
	LastVisible  DateTime `json:"LastVisible,omitempty"`
	IsWrittenOff bool     `json:"IsWrittenOff,omitempty"`
	WriteOffDate bool     `json:"WriteOffDate,omitempty"`
	InvNum       string   `json:"InvNum,omitempty"`
	UserName     string   `json:"UserName,omitempty"`
	Placement    string   `json:"Placement,omitempty"`
	Price        Long     `json:"Price,omitempty"`
	PurchaseDate DateTime `json:"PurchaseDate,omitempty"`
	Corporative  bool     `json:"Corporative,omitempty"`
	Name         string   `json:"Name,omitempty"`
	Description  string   `json:"Description,omitempty"`
	Manufacturer string   `json:"Manufacturer,omitempty"`
	SerialNumber string   `json:"SerialNumber,omitempty"`
	CPU          string   `json:"CPU,omitempty"`
	MemorySize   int64    `json:"MemorySize,omitempty"`
	DiskSize     int64    `json:"DiskSize,omitempty"`
	MotherBoard  string   `json:"MotherBoard,omitempty"`
	VidPID       string   `json:"VidPid,omitempty"`
	Capacity     int64    `json:"Capacity,omitempty"`
	MAC          string   `json:"Mac,omitempty"`
	StrMAC       string   `json:"StrMac,omitempty"`
	OS           string   `json:"OS,omitempty"`
	AdObject     string   `json:"AdObject,omitempty"`
	AdObjectDN   string   `json:"AdObjectDN,omitempty"`
	//DynColumns   []DynColumn `json:"DynColumns"`
	DynColID   string `json:"DynColId,omitempty"`
	DynColName string `json:"DynColName,omitempty"`
	DynColData string `json:"DynColData,omitempty"`
}

type DynColumn struct {
	Type  string         `json:"type,omitempty"`
	Value DynColumnValue `json:"value,omitempty"`
}

type DynColumnValue struct {
	DynColData string `json:"DynColData,omitempty"`
	DynColID   string `json:"DynColId,omitempty"`
	DynColName string `json:"DynColName,omitempty"`
}

func (hw *HWInvStorage) AddHWInvObject(ctx context.Context, params PpObj) (*PxgValInt, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.AddHWInvObject", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgValInt := new(PxgValInt)
	_, err = hw.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, err
}

// DelDynColumn Delete dynamic column.
func (hw *HWInvStorage) DelDynColumn(ctx context.Context, wstrColId string) error {
	postData := []byte(fmt.Sprintf(`{"wstrColId": "%s"}`, wstrColId))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.DelDynColumn", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = hw.client.Do(ctx, request, nil)
	return err
}

// DelHWInvObject Delete hardware inventory object.
func (hw *HWInvStorage) DelHWInvObject(ctx context.Context, nObjId int64) error {
	postData := []byte(fmt.Sprintf(`{"nObjId": %d}`, nObjId))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.DelHWInvObject", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = hw.client.Do(ctx, request, nil)
	return err
}

// DelHWInvObject2 Delete array of objects.
func (hw *HWInvStorage) DelHWInvObject2(ctx context.Context, arrObjId []int64) error {
	postData := []byte(fmt.Sprintf(`{"arrObjId": %s}`, ToJson(arrObjId)))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.DelHWInvObject2", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = hw.client.Do(ctx, request, nil)
	return err
}

// ExportHWInvStorage2 Start export of hardware inventory.
func (hw *HWInvStorage) ExportHWInvStorage2(ctx context.Context, eExportType int) (*PxgValStr, error) {
	postData := []byte(fmt.Sprintf(`{"eExportType": %d}`, eExportType))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.ExportHWInvStorage2", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgValStr := new(PxgValStr)
	_, err = hw.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, err
}

// ExportHWInvStorageCancel Cancel export of hardware inventory.
func (hw *HWInvStorage) ExportHWInvStorageCancel(ctx context.Context, wstrAsyncId string) error {
	postData := []byte(fmt.Sprintf(`{"wstrAsyncId": "%s"}`, wstrAsyncId))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.ExportHWInvStorageCancel", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = hw.client.Do(ctx, request, nil)
	return err
}

// EnumDynColumns Start import of hardware inventory.
func (hw *HWInvStorage) ImportHWInvStorage2(ctx context.Context, eImportType int64) (*PxgValStr, error) {
	postData := []byte(fmt.Sprintf(`{"eImportType": %d}`, eImportType))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.ImportHWInvStorage2", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgValStr := new(PxgValStr)
	_, err = hw.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, err
}

// ImportHWInvStorageCancel Cancel import of hardware inventory.
func (hw *HWInvStorage) ImportHWInvStorageCancel(ctx context.Context, params AsyncID) (*PxgValStr, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.ImportHWInvStorageCancel", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgValStr := new(PxgValStr)
	_, err = hw.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, err
}

// StorageSetData struct
type StorageSetData struct {
	WstrAsyncID string  `json:"wstrAsyncId,omitempty"`
	PChunk      *string `json:"pChunk,omitempty"`
}

// ImportHWInvStorageSetData Send chunk of importing data to server.
//	If pChunk is NULL then send data is finished and started data processing and importing to DB.
//	To get status use AsyncActionStateChecker.CheckActionState, lStateCode "0" means OK.
func (hw *HWInvStorage) ImportHWInvStorageSetData(ctx context.Context, params StorageSetData) error {
	postData, err := json.Marshal(params)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.ImportHWInvStorageSetData", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = hw.client.Do(ctx, request, nil)
	return err
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
func (hw *HWInvStorage) GetProcessingRules(ctx context.Context) (*ProcessingRules, error) {
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.GetProcessingRules", nil)
	if err != nil {
		return nil, err
	}

	processingRules := new(ProcessingRules)
	_, err = hw.client.Do(ctx, request, &processingRules)
	return processingRules, err
}

// SetProcessingRules Set processing rules.
func (hw *HWInvStorage) SetProcessingRules(ctx context.Context, params ProcessingRules) error {
	postData, err := json.Marshal(params)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.SetProcessingRules", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = hw.client.Do(ctx, request, nil)
	return err
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
	// PChunk Data chunk
	PChunk string `json:"pChunk,omitempty"`

	// NGotDataSize Actual size of retrieved data
	NGotDataSize int64 `json:"nGotDataSize,omitempty"`

	// NDataSizeREST Size of not retrieved data
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

// CorpFlagParams struct
type CorpFlagParams struct {
	// ArrObjId Array of device ids. Max array size is 1000 elements.
	ArrObjId []int64 `json:"arrObjId,omitempty"`

	// BState New state
	BState bool `json:"bState,omitempty"`
}

// SetCorpFlag2 Set corporative flag for array of devices.
func (hw *HWInvStorage) SetCorpFlag2(ctx context.Context, params CorpFlagParams) error {
	postData, err := json.Marshal(params)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.SetCorpFlag2", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = hw.client.Do(ctx, request, nil)
	return err
}

// HWInvObjectParams struct
type HWInvObjectParams struct {
	NObjID int64       `json:"nObjId,omitempty"`
	PObj   HWInvObject `json:"pObj,omitempty"`
}

// SetHWInvObject Set hardware inventory object.
func (hw *HWInvStorage) SetHWInvObject(ctx context.Context, params HWInvObjectParams) error {
	postData, err := json.Marshal(params)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.SetHWInvObject", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = hw.client.Do(ctx, request, nil)
	return err
}

// SetWriteOffFlag Set decommissioned flag.
func (hw *HWInvStorage) SetWriteOffFlag(ctx context.Context, nObjId int64, bFlag bool) error {
	postData := []byte(fmt.Sprintf(`{"nObjId": %d,"bFlag": %v}`, nObjId, bFlag))
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.SetWriteOffFlag", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = hw.client.Do(ctx, request, nil)
	return err
}

// WriteOffFlag struct
type WriteOffFlag struct {
	VecObjId []int64 `json:"vecObjId,omitempty"`
	BFlag    bool    `json:"bFlag,omitempty"`
}

// WriteOffFlag Set decommissioned flag for array of devices.
func (hw *HWInvStorage) SetWriteOffFlag2(ctx context.Context, params WriteOffFlag) error {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", hw.client.Server+"/api/v1.0/HWInvStorage.SetWriteOffFlag2", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = hw.client.Do(ctx, request, nil)
	return err
}
