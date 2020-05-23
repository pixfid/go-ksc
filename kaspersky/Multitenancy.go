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

//	Multitenancy Class Reference
//
//	Multitenancy product managing.
//
//	List of all members.
type Multitenancy service

//	Retrieves tenant identity.
//
//	Identity is unique for each tenant
//
//	Returns:
//	- (string) tenant id
func (m *Multitenancy) GetTenantId(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", m.client.Server+"/api/v1.0/Multitenancy.GetTenantId", nil)
	if err != nil {
		return nil, err
	}

	raw, err := m.client.Do(ctx, request, nil)
	return raw, err
}

//	Retrieves multitenancy products available for current tenant.
//
//	Parameters:
//	- strProdName	(string) product name, if set then result will be filtered by this value
//	- strProdVersion	(string) product version, if set then result will be filtered by this value
//
//	Returns:
//	- (array) each element of array contains information about multitenancy product:
//
//	+---------------------------+--------+-------------------------+
//	|           Value           |  Type  |       Description       |
//	+---------------------------+--------+-------------------------+
//	| MTNC_PRODUCT_NAME         | string | product name            |
//	| MTNC_PRODUCT_VERSION      | string | product version         |
//	| MTNC_PRODUCT_DISP_NAME    | string | display product name    |
//	| MTNC_PRODUCT_DISP_VERSION | string | display product version |
//	+---------------------------+--------+-------------------------+
func (m *Multitenancy) GetProducts(ctx context.Context, strProdName, strProdVersion string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"strProdName": "%s",
	"strProdVersion": "%s"
	}`, strProdName, strProdVersion))
	request, err := http.NewRequest("POST", m.client.Server+"/api/v1.0/Multitenancy.GetProducts", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := m.client.Do(ctx, request, nil)
	return raw, err
}

//	Get new binary token for current tennant
//
//	Returns:
//	- new token for current tennant
//
//	NotWoking on KSC < 12
func (m *Multitenancy) GetAuthToken(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", m.client.Server+"/api/v1.0/Multitenancy.GetAuthToken", nil)
	if err != nil {
		return nil, err
	}

	raw, err := m.client.Do(ctx, request, nil)
	return raw, err
}

//TODO func (m *Multitenancy) CheckAuthToken(ctx context.Context, wstrTenantId string, binToken []byte) ([]byte, error)
