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

// VServers2 service to access Virtual servers processing.
type VServers2 service

// VServerStatistic struct
type VServerStatistic struct {
	VSStatistic *VSStatistic `json:"PxgRetVal,omitempty"`
}

// VSStatistic struct
type VSStatistic struct {
	KlvsrvCreated   *DateTime       `json:"KLVSRV_CREATED,omitempty"`
	KlvsrvGroups    *int64          `json:"KLVSRV_GROUPS,omitempty"`
	KlvsrvHosts     *int64          `json:"KLVSRV_HOSTS,omitempty"`
	KlvsrvLicenses  []KlvsrvLicense `json:"KLVSRV_LICENSES,omitempty"`
	KlvsrvMdmios    *int64          `json:"KLVSRV_MDMIOS,omitempty"`
	KlvsrvMobilies  *int64          `json:"KLVSRV_MOBILIES,omitempty"`
	KlvsrvProducts  *KlvsrvProducts `json:"KLVSRV_PRODUCTS,omitempty"`
	KlvsrvProducts2 *KlvsrvProducts `json:"KLVSRV_PRODUCTS_2,omitempty"`
	KlvsrvUsers     *int64          `json:"KLVSRV_USERS,omitempty"`
}

type KlvsrvLicense struct {
	Type               *string             `json:"type,omitempty"`
	KlvsrvLicenseValue *KlvsrvLicenseValue `json:"value,omitempty"`
}

type KlvsrvLicenseValue struct {
	KllicKeyType   *int64    `json:"KLLIC_KEY_TYPE,omitempty"`
	KllicLimitDate *DateTime `json:"KLLIC_LIMIT_DATE,omitempty"`
	KllicProdName  *string   `json:"KLLIC_PROD_NAME,omitempty"`
	KllicSerial    *string   `json:"KLLIC_SERIAL,omitempty"`
}

// Klvsrv struct
type KlvsrvProducts struct {
	Type  *string                 `json:"type,omitempty"`
	Value *map[string]KlvsrvProds `json:"value,omitempty"`
}

type KlvsrvProds struct {
	Type  *string                       `json:"type,omitempty"`
	Value *map[string]KlvsrvProdVersion `json:"value,omitempty"`
}

type KlvsrvProdVersion struct {
	Type            *string         `json:"type,omitempty"`
	KlvsrvProdValue KlvsrvProdValue `json:"value,omitempty"`
}

type KlvsrvProdValue struct {
	KlvsrvProdCnt    *int64             `json:"KLVSRV_PROD_CNT,omitempty"`
	KlvsrvProdDN     *string            `json:"KLVSRV_PROD_DN,omitempty"`
	KlvsrvProdDv     *string            `json:"KLVSRV_PROD_DV,omitempty"`
	KlvsrvProdBuilds []KlvsrvProdBuilds `json:"KLVSRV_PROD_BUILDS,omitempty"`
}

type KlvsrvProdBuilds struct {
	Type                 *string         `json:"type,omitempty"`
	KlvsrvProdBuildValue KlvsrvProdValue `json:"value,omitempty"`
}

// GetVServerStatistic Acquire info on virtual server. Returns info about the specified virtual server
func (vs *VServers2) GetVServerStatistic(ctx context.Context, lVsId int) (*VServerStatistic, error) {
	postData := []byte(fmt.Sprintf(`{"lVsId": %d}`, lVsId))
	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers2.GetVServerStatistic", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	vServerStatistic := new(VServerStatistic)
	_, err = vs.client.Do(ctx, request, &vServerStatistic)
	return vServerStatistic, err
}
