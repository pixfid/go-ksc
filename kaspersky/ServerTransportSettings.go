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

// ServerTransportSettings Server transport settings proxy class.
type ServerTransportSettings service

// GetNumberOfManagedDevicesAgentless Returns number of agentless managed devices.
//
// Note: It can be called from main server only !
func (sts *ServerTransportSettings) GetNumberOfManagedDevicesAgentless(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.GetNumberOfManagedDevicesAgentless", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := sts.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// GetNumberOfManagedDevicesKSM Returns number of managed devices for KSM (Kaspersky for Mobile).
//
// Note: It can be called from main server only !
func (sts *ServerTransportSettings) GetNumberOfManagedDevicesKSM(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.GetNumberOfManagedDevicesKSM", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := sts.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// IsFeatureActive Checks if feature is activated and certificate can be changed to some custom value.
func (sts *ServerTransportSettings) IsFeatureActive(ctx context.Context, szwCertType string) (*PxgValBool, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwCertType": "%s"}`, szwCertType))
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.IsFeatureActive", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := sts.client.Request(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

// SetFeatureActive Sets feature active.
func (sts *ServerTransportSettings) SetFeatureActive(ctx context.Context, szwCertType string, bFeatureActive bool) (*PxgValBool, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwCertType": "%s", "bFeatureActive" : %v}`, szwCertType, bFeatureActive))
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.SetFeatureActive", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := sts.client.Request(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

// CheckDefaultCertificateExists It checks if default certificate exists.
func (sts *ServerTransportSettings) CheckDefaultCertificateExists(ctx context.Context, szwCertType string) (*PxgValBool, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwCertType": "%s"}`, szwCertType))
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.CheckDefaultCertificateExists", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := sts.client.Request(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

//CurrentConnectionSettings struct
type CurrentConnectionSettings struct {
	CCSettings *CCSettings `json:"PxgRetVal,omitempty"`
}

// CCSettings struct
type CCSettings struct {
	// CERTPub current certificate's public key
	CERTPub *CERTPub `json:"CERT_PUB,omitempty"`

	// TrspSettingsFQDN actual endpoint FQDN (from certificate)
	TrspSettingsFQDN string `json:"TRSP_SETTINGS_FQDN,omitempty"`

	// TrspSettingsIsdefcertused is default certificate used ?
	TrspSettingsIsdefcertused bool `json:"TRSP_SETTINGS_ISDEFCERTUSED,omitempty"`

	// TrspSettingsOpenPort true if port should be opened, false otherwise.
	TrspSettingsOpenPort bool `json:"TRSP_SETTINGS_OPEN_PORT,omitempty"`

	// TrspSettingsPort actual enpoint port
	TrspSettingsPort int64 `json:"TRSP_SETTINGS_PORT,omitempty"`
}

// CERTPub struct
type CERTPub struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

// GetCurrentConnectionSettings Returns current connection settings.
func (sts *ServerTransportSettings) GetCurrentConnectionSettings(ctx context.Context, szwCertType string) (*CurrentConnectionSettings, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwCertType": "%s"}`, szwCertType))
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.GetCurrentConnectionSettings", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	cCSettings := new(CurrentConnectionSettings)
	raw, err := sts.client.Request(ctx, request, &cCSettings)
	return cCSettings, raw, err
}

// GetCustomSrvCertificateInfo Returns information about custom certificate.
func (sts *ServerTransportSettings) GetCustomSrvCertificateInfo(ctx context.Context, szwCertType string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwCertType": "%s"}`, szwCertType))
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.GetCustomSrvCertificateInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sts.client.Request(ctx, request, nil)
	return raw, err
}

// GetDefaultConnectionSettings Returns default connection settings.
func (sts *ServerTransportSettings) GetDefaultConnectionSettings(ctx context.Context, szwCertType string) (*CurrentConnectionSettings, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwCertType": "%s"}`, szwCertType))
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.GetDefaultConnectionSettings", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	cCSettings := new(CurrentConnectionSettings)
	raw, err := sts.client.Request(ctx, request, &cCSettings)
	return cCSettings, raw, err
}

// ResetCstmReserveCertificate Resets custom reserve certificate.
func (sts *ServerTransportSettings) ResetCstmReserveCertificate(ctx context.Context, szwCertType string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwCertType": "%s"}`, szwCertType))
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.ResetCstmReserveCertificate", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sts.client.Request(ctx, request, nil)
	return raw, err
}

// ResetDefaultReserveCertificate Resets default reserve certificate.
func (sts *ServerTransportSettings) ResetDefaultReserveCertificate(ctx context.Context, szwCertType string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwCertType": "%s"}`, szwCertType))
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.ResetDefaultReserveCertificate", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sts.client.Request(ctx, request, nil)
	return raw, err
}

// SetOrCreateDefaultCertificate. It sets or recreates default certificate.
func (sts *ServerTransportSettings) SetOrCreateDefaultCertificate(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.SetOrCreateDefaultCertificate",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sts.client.Request(ctx, request, nil)
	return raw, err
}

// SetCustomSrvCertificate Sets custom certificate for one of SC Server's SSL listener.
func (sts *ServerTransportSettings) SetCustomSrvCertificate(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.SetCustomSrvCertificate",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sts.client.Request(ctx, request, nil)
	return raw, err
}
