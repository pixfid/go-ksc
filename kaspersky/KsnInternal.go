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

// KsnInternal service for working with KsnProxy subsystem.
//
// This service allow you to set KPSN settings, check license and check connection.
type KsnInternal service

// CheckKsnConnection Check connection with KSN cloud (or KPSN)
func (sd *KsnInternal) CheckKsnConnection(ctx context.Context) (*PxgValBool, []byte, error) {
	request, err := http.NewRequest("POST", sd.client.Server+"/api/v1.0/KsnInternal.CheckKsnConnection", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := sd.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

// GetNKsnEulas Get all KPSN eula.
func (sd *KsnInternal) GetNKsnEulas(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", sd.client.Server+"/api/v1.0/KsnInternal.GetNKsnEulas", nil)
	if err != nil {
		return nil, err
	}

	raw, err := sd.client.Do(ctx, request, nil)
	return raw, err
}

// KsnSettings is returned by GetSettings.
type KsnSettings struct {
	PSettings *PSettings `json:"pSettings,omitempty"`
}

type PSettings struct {
	KsnproxyEnable                 bool  `json:"KSNPROXY_ENABLE,omitempty"`
	KsnproxyEnablePatchManStat     bool  `json:"KSNPROXY_ENABLE_PATCH_MAN_STAT,omitempty"`
	KsnproxyKsnType                int64 `json:"KSNPROXY_KSN_TYPE,omitempty"`
	KsnproxyTCPPort                int64 `json:"KSNPROXY_TCP_PORT,omitempty"`
	KsnproxyUDPPort                int64 `json:"KSNPROXY_UDP_PORT,omitempty"`
	KsnproxyUseMasterKsnproxyAsKsn bool  `json:"KSNPROXY_USE_MASTER_KSNPROXY_AS_KSN,omitempty"`
}

// GetSettings Returns settings of KsnProxy. May be used on virtual server.
func (sd *KsnInternal) GetSettings(ctx context.Context) (*KsnSettings, []byte, error) {
	request, err := http.NewRequest("POST", sd.client.Server+"/api/v1.0/KsnInternal.GetSettings", nil)
	if err != nil {
		return nil, nil, err
	}

	ksnSettings := new(KsnSettings)
	raw, err := sd.client.Do(ctx, request, &ksnSettings)
	return ksnSettings, raw, err
}

// NeedToSendStatistics Check possibility to send statistics.
func (sd *KsnInternal) NeedToSendStatistics(ctx context.Context) (*PxgValBool, []byte, error) {
	request, err := http.NewRequest("POST", sd.client.Server+"/api/v1.0/KsnInternal.NeedToSendStatistics", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := sd.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

// GetNKsnEula Get KPSN eula.
func (sd *KsnInternal) GetNKsnEula(ctx context.Context, wstrNKsnLoc string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrNKsnLoc": "%s"}`, wstrNKsnLoc))
	request, err := http.NewRequest("POST", sd.client.Server+"/api/v1.0/KsnInternal.GetNKsnEula", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sd.client.Do(ctx, request, nil)
	return raw, err
}
