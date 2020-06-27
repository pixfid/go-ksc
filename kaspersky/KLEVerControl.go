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

//	KLEVerControl service to controls the possibility to download and automatically create installation packages.
type KLEVerControl service

// CancelDownloadDistributive Cancel asynchronous operation DownloadDistributiveAsync.
func (kvc *KLEVerControl) CancelDownloadDistributive(ctx context.Context, wstrRequestId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrRequestId": "%s"}`, wstrRequestId))
	request, err := http.NewRequest("POST", kvc.client.Server+"/api/v1.0/KLEVerControl.CancelDownloadDistributive", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := kvc.client.Do(ctx, request, nil)
	return raw, err
}

// GetDownloadDistributiveResult Get result of asynchronous operation DownloadDistributiveAsync.
func (kvc *KLEVerControl) GetDownloadDistributiveResult(ctx context.Context, wstrRequestId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrRequestId": "%s"}`, wstrRequestId))
	request, err := http.NewRequest("POST", kvc.client.Server+"/api/v1.0/KLEVerControl.GetDownloadDistributiveResult", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := kvc.client.Do(ctx, request, nil)
	return raw, err
}

// ChangeCreatePackage Initiate or cancel distributives downloading and installation packages registration from KL public distributives storage.
// The distributives are identified by "db_loc_id" from the appropriate SrvView Kaspersky Lab corporate product distributives available for download.
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

// DownloadDistributiveAsync Initiate downloading of the distributive by URL into SC-server.
// Method is needed to download distributive by URL into SC-server. After that the distributive will be available to downloading from SC-server.
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
