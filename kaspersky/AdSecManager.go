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

// AdSecManager service for Adaptive Security managing.
//
// Allows to approve or disprove detection results provided by Adaptive Security component.
type AdSecManager service

// DetectParams struct
type DetectParams struct {
	// PArrDetects detection results to disapprove
	PArrDetects []PArrDetect `json:"pArrDetects"`
}

type PArrDetect struct {
	Type  string       `json:"type,omitempty"`
	Value PDetectValue `json:"value,omitempty"`
}

type PDetectValue struct {
	// KlhstWksHostname host name
	KlhstWksHostname string `json:"KLHST_WKS_HOSTNAME,omitempty"`

	// KlhstWksProductName name of product
	KlhstWksProductName string `json:"KLHST_WKS_PRODUCT_NAME,omitempty"`

	// KlhstWksProductVersion version of product
	KlhstWksProductVersion string `json:"KLHST_WKS_PRODUCT_VERSION,omitempty"`

	// ListItemID id of item
	ListItemID string `json:"ListItemId,omitempty"`
}

// ApproveDetect Approves detection results provided by Adaptive Security component.
func (asm *AdSecManager) ApproveDetect(ctx context.Context, params DetectParams) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", asm.client.Server+"/api/v1.0/AdSecManager.ApproveDetect", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := asm.client.Do(ctx, request, nil)
	return raw, err
}

// DisproveDetect Disapprove detection results provided by Adaptive Security component.
func (asm *AdSecManager) DisproveDetect(ctx context.Context, params DetectParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", asm.client.Server+"/api/v1.0/AdSecManager.DisproveDetect", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := asm.client.Do(ctx, request, nil)
	return raw, err
}
