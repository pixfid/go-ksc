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
	"net/http"
)

// SrvCloud service to acquire info about public clouds.
type SrvCloud service

// GetCloudsInfo Returns list of clouds of the current server.
//
// For main server also returns only clouds of main server.
func (sc *SrvCloud) GetCloudsInfo(ctx context.Context, params Null) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SrvCloud.GetCloudsInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Request(ctx, request, nil)
	return raw, err
}

// CloudHostInfoParams struct
type CloudHostInfoParams struct {
	// PCloudHostBinID Internal id of the cloud host (see KLHST_CLOUD_HOST_BINID).
	PCloudHostBinID string `json:"pCloudHostBinId,omitempty"`

	// PFields collection of cloud host attribute names that need to return see List of server cloud host attributes..
	PFields []string `json:"pFields"`

	// PParams For additional options. Reserved.
	PParams Null `json:"pParams"`
}

// GetCloudHostInfo Returns properties of the cloud host.
func (sc *SrvCloud) GetCloudHostInfo(ctx context.Context, params CloudHostInfoParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SrvCloud.GetCloudHostInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Request(ctx, request, nil)
	return raw, err
}
