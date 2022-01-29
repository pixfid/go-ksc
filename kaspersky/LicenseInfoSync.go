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

// LicenseInfoSync service provided operating with licenses.
type LicenseInfoSync service

// AcquireKeysForProductOnHost Acquire serial numbers of the active and reserved licenses on the host.
func (lis *LicenseInfoSync) AcquireKeysForProductOnHost(ctx context.Context, szwHostName, szwProduct, szwVersion string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwHostName":"%s","szwProduct":"%s","szwVersion":"%s"}`, szwHostName, szwProduct, szwVersion))
	request, err := http.NewRequest("POST", lis.client.Server+"/api/v1.0/LicenseInfoSync.AcquireKeysForProductOnHost",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lis.client.Request(ctx, request, nil)
	return raw, err
}

// GetKeyDataForHost Get host-specific key data.
func (lis *LicenseInfoSync) GetKeyDataForHost(ctx context.Context, szwSerial, szwHostName, szwProduct, szwVersion string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwSerial":"%s", "szwHostName":"%s", "szwProduct":"%s", "szwVersion":"%s"}`,
		szwSerial, szwHostName, szwProduct, szwVersion))
	request, err := http.NewRequest("POST", lis.client.Server+"/api/v1.0/LicenseInfoSync.GetKeyDataForHost",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lis.client.Request(ctx, request, nil)
	return raw, err
}

// SaasKeyParam2 struct
type SaasKeyParam2 struct {
	// SaasPKeyInfo License attribute container, mandatory. See List of license key attributes for attribute names.
	SaasPKeyInfo *SaasPKeyInfo `json:"pKeyInfo,omitempty"`

	// BAsCurrent true to check it as an active one, false to check it as a reserved one.
	BAsCurrent bool `json:"bAsCurrent,omitempty"`
}

type SaasPKeyInfo struct {
	// KllicSerial License key serial number
	KllicSerial string `json:"KLLIC_SERIAL,omitempty"`
}

// IsLicForSaasValid2 Check if license is suitable for being used by the adm. server.
func (lis *LicenseInfoSync) IsLicForSaasValid2(ctx context.Context, params SaasKeyParam2) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", lis.client.Server+"/api/v1.0/LicenseInfoSync.IsLicForSaasValid2",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lis.client.Request(ctx, request, nil)
	return raw, err
}

// IsPCloudKey Check whether the key's product id belongs to the Public Cloud product ids list.
func (lis *LicenseInfoSync) IsPCloudKey(ctx context.Context, nProductId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nProductId": %d}`, nProductId))
	request, err := http.NewRequest("POST", lis.client.Server+"/api/v1.0/LicenseInfoSync.IsPCloudKey",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lis.client.Request(ctx, request, nil)
	return raw, err
}

// SynchronizeLicInfo2 Force synchronization of subscription licenses' metadata.
func (lis *LicenseInfoSync) SynchronizeLicInfo2(ctx context.Context) (*PxgValStr, []byte, error) {
	request, err := http.NewRequest("POST", lis.client.Server+"/api/v1.0/LicenseInfoSync.SynchronizeLicInfo2",
		nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := lis.client.Request(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// TryToInstallLicForSaas2 Install adm. server's license.
func (lis *LicenseInfoSync) TryToInstallLicForSaas2(ctx context.Context, params SaasKeyParam2) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", lis.client.Server+"/api/v1.0/LicenseInfoSync.TryToInstallLicForSaas2",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lis.client.Request(ctx, request, nil)
	return raw, err
}

// TryToUnistallLicense Uninstall adm. server's license.
func (lis *LicenseInfoSync) TryToUnistallLicense(ctx context.Context, bCurrent bool) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"bCurrent": %v}`, bCurrent))
	request, err := http.NewRequest("POST", lis.client.Server+"/api/v1.0/LicenseInfoSync.TryToUnistallLicense",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lis.client.Request(ctx, request, nil)
	return raw, err
}
