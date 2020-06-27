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

// AppCtrlApi service to get info about execution files.
type AppCtrlApi service

type ExeFileInfoParams struct {
	SzwHostID  string      `json:"szwHostId,omitempty"`
	LFileID    int64       `json:"lFileId,omitempty"`
	ExePFilter *ExePFilter `json:"pFilter,omitempty"`
}

type ExePFilter struct {
	FileID   string `json:"FILE_ID,omitempty"`
	FileName string `json:"FILE_NAME,omitempty"`
}

// GetExeFileInfo Get data about instances of the execution file on the host.
//
// To do this for field 'FieldName' it is needed to add into pFilter the value of any type with name 'FieldName'
// If NULL than all possible fields will be returned.
func (ac *AppCtrlApi) GetExeFileInfo(ctx context.Context, params ExeFileInfoParams) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ac.client.Server+"/api/v1.0/AppCtrlApi.GetExeFileInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ac.client.Do(ctx, request, nil)
	return raw, err
}
