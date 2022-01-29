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

// UpdComps service provides means to manage updatable components (bases).
type UpdComps service

// UpdateParams struct of Update operation parameters
type UpdateParams struct {
	UPParams UPParams `json:"pParams"`
}

// UPParams struct
type UPParams struct {
	// ArrCategoryFilter Updatable component identifier list: each cell is ID of some updatable
	// component to download, e.g. "KSC"
	// (there exist many components on KL update servers, one must know, what is required)
	ArrCategoryFilter []string `json:"arrCategoryFilter"`
	// BDoUpdate Run update operation for given components (true) or retranslate them to KSC share (false)
	BDoUpdate bool `json:"bDoUpdate"`
}

// AsyncUpdate Requests asynchronous bases update or retranslation.
func (uc *UpdComps) AsyncUpdate(ctx context.Context, params UpdateParams) (*PxgValStr, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UpdComps.AsyncUpdate", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgValStr := new(PxgValStr)
	_, err = uc.client.Request(ctx, request, &pxgValStr)
	return pxgValStr, err
}

// Stop Stops operation by request.
func (uc *UpdComps) Stop(ctx context.Context, wsRequestId string, bWait bool) error {
	postData := []byte(fmt.Sprintf(`{"wsRequestId": "%s", "bWait": %v}`, wsRequestId, bWait))
	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UpdComps.Stop", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = uc.client.Request(ctx, request, nil)
	return err
}

// UpdateAsync Requests asynchronous bases update or retranslation.
func (uc *UpdComps) UpdateAsync(ctx context.Context, params UpdateParams) error {
	postData, err := json.Marshal(params)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UpdComps.UpdateAsync", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = uc.client.Request(ctx, request, nil)
	return err
}
