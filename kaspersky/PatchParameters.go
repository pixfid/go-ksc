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

//	PatchParameters Patch parameters processing.
type PatchParameters service

// GetTemplate Get template for command.
func (pp *PatchParameters) GetTemplate(ctx context.Context, patchID, locID int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"patchID": %d, "locID": %d}`, patchID, locID))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PatchParameters.GetTemplate", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Request(ctx, request, nil)
	return raw, err
}

// GetValues Get install command for patch.
func (pp *PatchParameters) GetValues(ctx context.Context, patchID, locID int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"patchID": %d, "locID": %d}`, patchID, locID))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PatchParameters.GetValues", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Request(ctx, request, nil)
	return raw, err
}

// GetValuesByPkg Get install command for patch.
func (pp *PatchParameters) GetValuesByPkg(ctx context.Context, packageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"packageId": %d}`, packageId))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PatchParameters.GetValuesByPkg", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Request(ctx, request, nil)
	return raw, err
}

// SetValues Set values for parameters of command.
func (pp *PatchParameters) SetValues(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PatchParameters.SetValues", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Request(ctx, request, nil)
	return raw, err
}

// SetValuesByPkg Set values for parameters of command.
func (pp *PatchParameters) SetValuesByPkg(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PatchParameters.SetValuesByPkg", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Request(ctx, request, nil)
	return raw, err
}
