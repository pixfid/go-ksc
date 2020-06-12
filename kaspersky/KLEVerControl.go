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

//	KLEVerControl Class Reference
//	Controls the possibility to download and automatically create installation packages.
//
//	List of all members.
type KLEVerControl service

//	Cancel asynchronous operation DownloadDistributiveAsync.
//
//	Parameters:
//	- wstrRequestId	(string) request id of asynchronous operation KLEVerControl.DownloadDistributiveAsync
func (kvc *KLEVerControl) CancelDownloadDistributive(ctx context.Context, wstrRequestId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrRequestId": "%s"}`, wstrRequestId))
	request, err := http.NewRequest("POST", kvc.client.Server+"/api/v1.0/KLEVerControl.CancelDownloadDistributive", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := kvc.client.Do(ctx, request, nil)
	return raw, err
}

//	Get result of asynchronous operation DownloadDistributiveAsync.
//
//	Parameters:
//	- wstrRequestId	(string) request id of asynchronous operation KLEVerControl.DownloadDistributiveAsync
//
//	Return:
//	- wstrDownloadPath	(string) path to download distributive from SC-server using HTTP GET method.
func (kvc *KLEVerControl) GetDownloadDistributiveResult(ctx context.Context, wstrRequestId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrRequestId": "%s"}`, wstrRequestId))
	request, err := http.NewRequest("POST", kvc.client.Server+"/api/v1.0/KLEVerControl.GetDownloadDistributiveResult", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := kvc.client.Do(ctx, request, nil)
	return raw, err
}

//	Initiate or cancel distributives downloading and installation
//	packages registration from KL public distributives storage.
//	The distributives are identified by "db_loc_id" from the appropriate
//	SrvView Kaspersky Lab corporate product distributives available for download.
//
//	Parameters:
//	- vecDistribLocIdsToCreate	(array) array of "Distributive localization database IDs" values (
//	"db_loc_id" from Kaspersky Lab corporate product distributives available for download) to download and create packages;
//
//	- vecDistribLocIdsNotToCreate	(array) array of "Distributive localization database IDs" values (
//	"db_loc_id" from Kaspersky Lab corporate product distributives available for download) to reset packages creation;
//
//	- parParams	(params) reserved;
func (kvc *KLEVerControl) ChangeCreatePackage(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", kvc.client.Server+"/api/v1.0/KLEVerControl.ChangeCreatePackage", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := kvc.client.Do(ctx, request, nil)
	return raw, err
}

//	Initiate downloading of the distributive by URL into SC-server.
//
//	Method is needed to download distributive by URL into SC-server.
//	After that the distributive will be available to downloading from SC-server.
//
//	Parameters:
//	- lDistribLocId	(int64) Distributive localization database Id that can be get from "db_loc_id" field
//	- pExtendedSettings	(params) additional parameters:
//	- "ExecutablePkg" - (bool) Download executable package
//
//	Returns:
//	- (string) request id of asynchronous operation,
//	to cancel call KLEVerControl.CancelDownloadDistributive,
//	to get status call AsyncActionStateChecker.CheckActionState with returned request id as action guid,
//	to get result after finishing call KLEVerControl.GetDownloadDistributiveResult
func (kvc *KLEVerControl) DownloadDistributiveAsync(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", kvc.client.Server+"/api/v1.0/KLEVerControl.DownloadDistributiveAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := kvc.client.Do(ctx, request, nil)
	return raw, err
}
