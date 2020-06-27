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

// Multitenancy product managing.
type Multitenancy service

// GetTenantId Retrieves tenant identity. Identity is unique for each tenant
func (m *Multitenancy) GetTenantId(ctx context.Context) (*PxgValStr, []byte, error) {
	request, err := http.NewRequest("POST", m.client.Server+"/api/v1.0/Multitenancy.GetTenantId", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := m.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// GetProducts Retrieves multitenancy products available for current tenant.
func (m *Multitenancy) GetProducts(ctx context.Context, strProdName, strProdVersion string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strProdName": "%s", "strProdVersion": "%s"}`, strProdName, strProdVersion))
	request, err := http.NewRequest("POST", m.client.Server+"/api/v1.0/Multitenancy.GetProducts", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := m.client.Do(ctx, request, nil)
	return raw, err
}

// GetAuthToken Get new binary token for current tennant
func (m *Multitenancy) GetAuthToken(ctx context.Context) (*PxgValStr, []byte, error) {
	request, err := http.NewRequest("POST", m.client.Server+"/api/v1.0/Multitenancy.GetAuthToken", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := m.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//VerifyTokenParam struct using in CheckAuthToken
type VerifyTokenParam struct {
	//tenant identity
	WstrTenantID string `json:"wstrTenantId,omitempty"`

	//binary token (see Multitenancy.GetAuthToken)
	BinToken string `json:"binToken,omitempty"`
}

// CheckAuthToken Verify token
func (m *Multitenancy) CheckAuthToken(ctx context.Context, params VerifyTokenParam) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", m.client.Server+"/api/v1.0/Multitenancy.CheckAuthToken",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := m.client.Do(ctx, request, nil)
	return raw, err
}
