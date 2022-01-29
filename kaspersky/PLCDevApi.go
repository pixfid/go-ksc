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

// PLCDevApi Interface allows to acquire and manage PLC devices registry.
//
// Administration server contains global list of the PLC devices.
//
// Every PLC device is identified by pPLCId - (binary) - binary data as array of 16 bytes.
type PLCDevApi service

// DeletePLCRemove PLC device.
func (pda *PLCDevApi) DeletePLC(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", pda.client.Server+"/api/v1.0/PLCDevApi.DeletePLC", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pda.client.Request(ctx, request, nil)
	return raw, err
}

// GetPLC Acquire attributes of specified PLC device.
func (pda *PLCDevApi) GetPLC(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", pda.client.Server+"/api/v1.0/PLCDevApi.GetPLC", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pda.client.Request(ctx, request, nil)
	return raw, err
}

// UpdatePLC Adds/Updates PLC device.
func (pda *PLCDevApi) UpdatePLC(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", pda.client.Server+"/api/v1.0/PLCDevApi.UpdatePLC", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pda.client.Request(ctx, request, nil)
	return raw, err
}
