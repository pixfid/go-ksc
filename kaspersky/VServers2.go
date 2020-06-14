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

// VServers2 Class Reference
//
// Virtual servers processing..
//
// List of all members:
type VServers2 service

// VServerStatistic struct
type VServerStatistic struct {
	VSStatistic *VSStatistic `json:"PxgRetVal,omitempty"`
}

// VSStatistic struct
type VSStatistic struct {
	KlvsrvCreated   *Klvsrv       `json:"KLVSRV_CREATED,omitempty"`
	KlvsrvGroups    int64         `json:"KLVSRV_GROUPS,omitempty"`
	KlvsrvHosts     int64         `json:"KLVSRV_HOSTS,omitempty"`
	KlvsrvLicenses  []interface{} `json:"KLVSRV_LICENSES"`
	KlvsrvMdmios    int64         `json:"KLVSRV_MDMIOS,omitempty"`
	KlvsrvMobilies  int64         `json:"KLVSRV_MOBILIES,omitempty"`
	KlvsrvProducts  *Klvsrv       `json:"KLVSRV_PRODUCTS,omitempty"`
	KlvsrvProducts2 *Klvsrv       `json:"KLVSRV_PRODUCTS_2,omitempty"`
	KlvsrvUsers     int64         `json:"KLVSRV_USERS,omitempty"`
}

// Klvsrv struct
type Klvsrv struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value"`
}

// GetVServerStatistic
// Acquire info on virtual server.
// Returns info about the specified virtual server
//
// Parameters:
//	- lVsId	(int64) virtual server id
//
// Returns:
//	- (params) a container, see Virtual server statistic.
func (vs *VServers2) GetVServerStatistic(ctx context.Context, lVsId int) (*VServerStatistic, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"lVsId": %d}`, lVsId))
	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers2.GetVServerStatistic", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	vServerStatistic := new(VServerStatistic)
	raw, err := vs.client.Do(ctx, request, &vServerStatistic)
	return vServerStatistic, raw, err
}
