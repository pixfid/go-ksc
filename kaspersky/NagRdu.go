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

//	NagRdu Class service to remote diagnostics on host.
//	This service is implemented at Network Agent side, so use gateway connection to connect Network Agent and call interface methods.
type NagRdu service

// ChangeTraceParams Change trace-level for specific product, turns on/off tracing
func (nr *NagRdu) ChangeTraceParams(ctx context.Context, szwProductID string, nTraceLevel int64) (*CurrentHostState, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwProductID": "%s", "nTraceLevel": %d }`, szwProductID, nTraceLevel))
	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.ChangeTraceParams", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	currentHostState := new(CurrentHostState)
	raw, err := nr.client.Do(ctx, request, &currentHostState)
	return currentHostState, raw, err
}

// ChangeTraceRotatedParams Change rotated-trace-level for specific product, turns on/off tracing
func (nr *NagRdu) ChangeTraceRotatedParams(ctx context.Context, szwProductID string, nTraceLevel,
	nPartsCount, nMaxPartSize int64) (*CurrentHostState, []byte, error) {
	postData := []byte(fmt.Sprintf(`{ "szwProductID": "%s", "nTraceLevel": %d, "nPartsCount": %d, 
	"nMaxPartSize": %d }`, szwProductID, nTraceLevel, nPartsCount, nMaxPartSize))
	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.ChangeTraceRotatedParams", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	currentHostState := new(CurrentHostState)
	raw, err := nr.client.Do(ctx, request, &currentHostState)
	return currentHostState, raw, err
}

// ChangeXperfBaseParams Change XPerf trace-level for specific product, turns on/off XPerf tracing
func (nr *NagRdu) ChangeXperfBaseParams(ctx context.Context, szwProductID string, nTraceLevel, nXPerfMode int64) (*CurrentHostState,
	[]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"szwProductID": "%s", "nTraceLevel": %d, "nXPerfMode": %d }`, szwProductID,
		nTraceLevel,
		nXPerfMode))

	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.ChangeXperfBaseParams", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	currentHostState := new(CurrentHostState)
	raw, err := nr.client.Do(ctx, request, &currentHostState)
	return currentHostState, raw, err
}

// ChangeXperfRotatedParams Change XPerf rotated-trace-level for specific product, turns on/off XPerf tracing
func (nr *NagRdu) ChangeXperfRotatedParams(ctx context.Context, szwProductID string, nTraceLevel, nXPerfMode,
	nMaxPartSize int64) (*CurrentHostState, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwProductID": "%s", "nTraceLevel": %d, "nXPerfMode": %d, "nMaxPartSize": %d }`, szwProductID,
		nTraceLevel, nXPerfMode, nMaxPartSize))

	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.ChangeXperfRotatedParams", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	currentHostState := new(CurrentHostState)
	raw, err := nr.client.Do(ctx, request, &currentHostState)
	return currentHostState, raw, err
}

//CurrentHostState struct
type CurrentHostState struct {
	CHState *CHState `json:"PxgRetVal,omitempty"`
}

type CHState struct {
	EventLogs        []string          `json:"EventLogs"`
	HostDN           string            `json:"HostDN,omitempty"`
	InstallationLogs *InstallationLogs `json:"InstallationLogs,omitempty"`
	LastActionResult string            `json:"LastActionResult,omitempty"`
	Products         *Products         `json:"Products,omitempty"`
	WuaLogs          []string          `json:"WuaLogs"`
	WuaLogsWin10     []string          `json:"WuaLogs_Win10"`
}

type InstallationLogs struct {
	Type  string                 `json:"type,omitempty"`
	Value *InstallationLogsValue `json:"value,omitempty"`
}

type InstallationLogsValue struct {
	InstallationLogsMSI []interface{} `json:"InstallationLogs_MSI"`
	InstallationLogsRI  []string      `json:"InstallationLogs_RI"`
}

type Products struct {
	Type  string                    `json:"type,omitempty"`
	Value map[string]ProductsValues `json:"value,omitempty"`
}

type ProductsValues struct {
	Type  string         `json:"type,omitempty"`
	Value *ProductsValue `json:"value,omitempty"`
}

type ProductsValue struct {
	ComponentName          string        `json:"ComponentName,omitempty"`
	DiagLog                string        `json:"DiagLog,omitempty"`
	DiagTrace              string        `json:"DiagTrace,omitempty"`
	InstallPath            string        `json:"InstallPath,omitempty"`
	ProdDumps              []interface{} `json:"ProdDumps"`
	ProdProps              int64         `json:"ProdProps,omitempty"`
	ProductBuild           string        `json:"ProductBuild,omitempty"`
	ProductDN              string        `json:"ProductDN,omitempty"`
	ProductName            string        `json:"ProductName,omitempty"`
	ProductRunning         bool          `json:"ProductRunning,omitempty"`
	ProductVersion         string        `json:"ProductVersion,omitempty"`
	TraceFile              []string      `json:"TraceFile"`
	TraceLevel             int64         `json:"TraceLevel,omitempty"`
	TraceLimitDef          int64         `json:"TraceLimitDef,omitempty"`
	TraceLimitMax          int64         `json:"TraceLimitMax,omitempty"`
	TraceLimitMin          int64         `json:"TraceLimitMin,omitempty"`
	TraceRotatedEnable     int64         `json:"TraceRotatedEnable,omitempty"`
	TraceRotatedFileCount  int64         `json:"TraceRotatedFileCount,omitempty"`
	TraceRotatedFileSizeMB int64         `json:"TraceRotatedFileSizeMb,omitempty"`
	TraceRotatedLevel      int64         `json:"TraceRotatedLevel,omitempty"`
	XperfTraceFile         []interface{} `json:"XperfTraceFile"`
	XperfTraceLevel        int64         `json:"XperfTraceLevel,omitempty"`
	XperfTraceSize         int64         `json:"XperfTraceSize,omitempty"`
	XperfTraceState        int64         `json:"XperfTraceState,omitempty"`
	XperfTraceType         int64         `json:"XperfTraceType,omitempty"`
}

// CreateAndDownloadDumpAsync Asynchronously create and download dump for specific process
//
// Remarks:
// Check the operation state by calling AsyncActionStateChecker.CheckActionState periodically until it's finalized.
// If the operation succeeds then AsyncActionStateChecker.CheckActionState returns URL-path in pStateData.
// Otherwise, a call to AsyncActionStateChecker::CheckActionState returns error in pStateData.
func (nr *NagRdu) CreateAndDownloadDumpAsync(ctx context.Context, szwProcessName string) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwProcessName": "%s"}`, szwProcessName))
	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.CreateAndDownloadDumpAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := nr.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// DeleteFile Permanently delete specific file on host
func (nr *NagRdu) DeleteFile(ctx context.Context, szwRemoteFile string) (*CurrentHostState, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwRemoteFile": "%s"}`, szwRemoteFile))
	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.DeleteFile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	currentHostState := new(CurrentHostState)
	raw, err := nr.client.Do(ctx, request, &currentHostState)
	return currentHostState, raw, err
}

type RemoteFilesParams struct {
	PRemoteFiles []string `json:"pRemoteFiles"`
}

// DeleteFiles Permanently delete array of specific files on host
func (nr *NagRdu) DeleteFiles(ctx context.Context, params RemoteFilesParams) (*CurrentHostState, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, nil
	}

	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.DeleteFiles", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	currentHostState := new(CurrentHostState)
	raw, err := nr.client.Do(ctx, request, &currentHostState)
	return currentHostState, raw, err
}

// DownloadCommonDataAsync Asynchronously create archive with common-data (products local settings, policy, tasks, ...) and download it
//
//	Remarks:
//	Check the operation state by calling AsyncActionStateChecker.CheckActionState periodically until it's finalized.
//	If the operation succeeds then AsyncActionStateChecker.CheckActionState returns URL-path in pStateData.
//	Otherwise, a call to AsyncActionStateChecker::CheckActionState returns error in pStateData.
func (nr *NagRdu) DownloadCommonDataAsync(ctx context.Context) (*PxgValStr, []byte, error) {
	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.DownloadCommonDataAsync", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := nr.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// DownloadEventlogAsync Asynchronously download specific event-log
//
// Remarks:
// Check the operation state by calling AsyncActionStateChecker.CheckActionState periodically until it's finalized.
// If the operation succeeds then AsyncActionStateChecker.CheckActionState returns URL-path in pStateData.
// Otherwise, a call to AsyncActionStateChecker::CheckActionState returns error in pStateData.
func (nr *NagRdu) DownloadEventlogAsync(ctx context.Context, szwEventLog string) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwEventLog": "%s"}`, szwEventLog))
	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.DownloadEventlogAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := nr.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// ExecuteFileAsync Asynchronously run executable file, previously uploaded to host using GetUrlToUploadFileToHost.
// Uploaded file should be a zip-archive with executable-file (and, maybe, other files and folders) on 'utility'-folder
func (nr *NagRdu) ExecuteFileAsync(ctx context.Context, szwURL, szwShortExecName, szwParams string) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwURL": "%s", "szwShortExecName": "%s", "szwParams": "%s"}`, szwURL, szwShortExecName, szwParams))
	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.ExecuteFileAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := nr.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// ExecuteGsiAsync Asynchronously run GSI-utility
//
// Remarks:
// Check the operation state by calling AsyncActionStateChecker.CheckActionState periodically until it's finalized.
// If the operation succeeds then AsyncActionStateChecker.CheckActionState returns URL-path in pStateData.
// Otherwise, a call to AsyncActionStateChecker.CheckActionState returns error in pStateData.
func (nr *NagRdu) ExecuteGsiAsync(ctx context.Context) (*PxgValStr, []byte, error) {
	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.ExecuteGsiAsync", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := nr.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// GetCurrentHostState Acquire current host state
func (nr *NagRdu) GetCurrentHostState(ctx context.Context) (*CurrentHostState, []byte, error) {
	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.GetCurrentHostState", nil)
	if err != nil {
		return nil, nil, err
	}

	currentHostState := new(CurrentHostState)
	raw, err := nr.client.Do(ctx, request, &currentHostState)
	return currentHostState, raw, err
}

// GetUrlToDownloadFileFromHost Get URL-path for later download file from host
func (nr *NagRdu) GetUrlToDownloadFileFromHost(ctx context.Context, szwRemoteFile string) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwRemoteFile": "%s"}`, jsonEscape(szwRemoteFile)))
	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.GetUrlToDownloadFileFromHost", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := nr.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// GetUrlToUploadFileToHost Get URL-path for later upload file to host
func (nr *NagRdu) GetUrlToUploadFileToHost(ctx context.Context) (*PxgValStr, []byte, error) {
	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.GetUrlToUploadFileToHost", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := nr.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// RunKlnagchkAsync Asynchronously run diagnostic-utility (klnagchk.exe) for specific product
//
// Remarks:
// Check the operation state by calling AsyncActionStateChecker.CheckActionState periodically until it's finalized.
// If the operation succeeds then AsyncActionStateChecker.CheckActionState returns current host state in pStateData.
// Otherwise, a call to AsyncActionStateChecker::CheckActionState returns error in pStateData.
func (nr *NagRdu) RunKlnagchkAsync(ctx context.Context, szwProductID string) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwProductID": "%s"}`, szwProductID))
	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.RunKlnagchkAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := nr.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// SetProductStateAsync Asynchronously start, restart or stop specific product
//
// Remarks:
// Check the operation state by calling AsyncActionStateChecker.CheckActionState periodically until it's finalized.
// If the operation succeeds then AsyncActionStateChecker.CheckActionState returns current host state in pStateData.
// Otherwise, a call to AsyncActionStateChecker::CheckActionState returns error in pStateData.
func (nr *NagRdu) SetProductStateAsync(ctx context.Context, szwProductID string, nNewState int64) (*PxgValStr, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"szwProductID": "%s", "nNewState": %d }`, szwProductID, nNewState))
	request, err := http.NewRequest("POST", nr.client.Server+"/api/v1.0/NagRdu.SetProductStateAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := nr.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}
