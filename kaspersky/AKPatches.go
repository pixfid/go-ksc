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

// AKPatches service to manage system of autoupdating by patch.exe patches.
type AKPatches service

//PatchParams struct
type PatchParams struct {
	SzwPatchID  string       `json:"szwPatchId"`
	PatchOption *PatchOption `json:"parOptions"`
}

type PatchOption struct {
	Type  string     `json:"type"`
	Value PatchValue `json:"value"`
}

type PatchValue struct {
	NeedBackupServer bool `json:"NeedBackupServer"`
}

// ApprovePatch Give approval to the patch on start installation.
//
// This command is applicable to patches which have no right to be installed without approval of the administrator.
func (akp *AKPatches) ApprovePatch(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", akp.client.Server+"/api/v1.0/AKPatches.ApprovePatch",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := akp.client.Do(ctx, request, nil)
	return raw, err
}

// ForbidPatch Forbid installation of the patch.
//
// After this forbidden will be no more notifications that this patch expecting approval on installation
func (akp *AKPatches) ForbidPatch(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", akp.client.Server+"/api/v1.0/AKPatches.ForbidPatch",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := akp.client.Do(ctx, request, nil)
	return raw, err
}
