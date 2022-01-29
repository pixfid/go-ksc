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
	"log"
	"net/http"
)

// ExtTenant Manage external tenant info interface.
type ExtTenant service

// GetExternalTenantId Gets external tenant id.
func (et *ExtTenant) GetExternalTenantId(ctx context.Context, nVServerId int64) (*PxgValStr, error) {
	postData := []byte(fmt.Sprintf(`{"nVServerId": %d}`, nVServerId))
	request, err := http.NewRequest("POST", et.client.Server+"/api/v1.0/ExtTenant.GetExternalTenantId",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	externalTenantId := new(PxgValStr)
	raw, err := et.client.Request(ctx, request, &externalTenantId)

	if et.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return externalTenantId, err
}

type ExternalTenantIdparams struct {
	NVServerID   int64  `json:"nVServerId"`
	WstrTenantID string `json:"wstrTenantID"`
}

func (et *ExtTenant) SetExternalTenantId(ctx context.Context, params ExternalTenantIdparams) error {
	postData, err := json.Marshal(params)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", et.client.Server+"/api/v1.0/ExtTenant.SetExternalTenantId",
		bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	raw, err := et.client.Request(ctx, request, nil)

	if et.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}
