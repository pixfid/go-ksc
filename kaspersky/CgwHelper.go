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

// CgwHelper (Connection Gateway) service to work with helper proxy.
type CgwHelper service

// GetSlaveServerLocation Retrieves Slave Server Location.
func (cp *CgwHelper) GetSlaveServerLocation(ctx context.Context, nSlaveServerId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nSlaveServerId": %d}`, nSlaveServerId))
	request, err := http.NewRequest("POST", cp.client.Server+"/api/v1.0/CgwHelper.GetSlaveServerLocation",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := cp.client.Do(ctx, request, nil)
	return raw, err
}

// NagentLocation struct
type NagentLocation struct {
	NagLocation *NagLocation `json:"PxgRetVal,omitempty"`
}

type NagLocation struct {
	GwLOCHostID          string `json:"GwLocHostId,omitempty"`
	GwLOCIndirect        bool   `json:"GwLocIndirect,omitempty"`
	GwLOCLocation        string `json:"GwLocLocation,omitempty"`
	GwLOCSignUDP         bool   `json:"GwLocSignUdp,omitempty"`
	GwLOCTargetComponent string `json:"GwLocTargetComponent,omitempty"`
	GwLOCUseCompression  bool   `json:"GwLocUseCompression,omitempty"`
}

// GetNagentLocation Retrieves Nagent Location by host name.
func (cp *CgwHelper) GetNagentLocation(ctx context.Context, wsHostName string) (*NagentLocation, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"wsHostName": "%s"}`, wsHostName))
	request, err := http.NewRequest("POST", cp.client.Server+"/api/v1.0/CgwHelper.GetNagentLocation",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	nagentLocation := new(NagentLocation)
	raw, err := cp.client.Do(ctx, request, &nagentLocation)
	return nagentLocation, raw, err
}
