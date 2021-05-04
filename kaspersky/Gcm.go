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

type Gcm service

// CheckIfGcmServerSettingsPresent It checks if GCM server settings are present.
// true if settings are present; False otherwise.
func (gm *Gcm) CheckIfGcmServerSettingsPresent(ctx context.Context) (*PxgValBool, error) {
	request, err := http.NewRequest("POST", gm.client.Server+"/api/v1.0/Gcm.CheckIfGcmServerSettingsPresent", nil)
	if err != nil {
		return nil, err
	}

	result := new(PxgValBool)
	_, err = gm.client.Do(ctx, request, &result)
	return result, err
}

// CheckIfGcmServerSettingsShouldBeSet It checks if GCM server settings should be set.
// true if server settings should be set; False otherwise.
func (gm *Gcm) CheckIfGcmServerSettingsShouldBeSet(ctx context.Context) (*PxgValBool, error) {
	request, err := http.NewRequest("POST", gm.client.Server+"/api/v1.0/Gcm.CheckIfGcmServerSettingsShouldBeSet", nil)
	if err != nil {
		return nil, err
	}

	result := new(PxgValBool)
	_, err = gm.client.Do(ctx, request, &result)
	return result, err
}

// DeleteGcmServerSettings Deletes GCM (Google Cloud Messaging) server settings.
func (gm *Gcm) DeleteGcmServerSettings(ctx context.Context) (*PxgValBool, error) {
	request, err := http.NewRequest("POST", gm.client.Server+"/api/v1.0/Gcm.DeleteGcmServerSettings", nil)
	if err != nil {
		return nil, err
	}

	result := new(PxgValBool)
	_, err = gm.client.Do(ctx, request, &result)
	return result, err
}

type PropagationState struct {
	BPropagate2VS bool `json:"bPropagate2VS"`
}

// GetGcmPropagation2VS Retrieves GCM settings propagation option.
// GCM settings can be propagated to virtual server in one case only - if it is absent on virtual server.
func (gm *Gcm) GetGcmPropagation2VS(ctx context.Context) (*PropagationState, error) {
	request, err := http.NewRequest("POST", gm.client.Server+"/api/v1.0/Gcm.GetGcmPropagation2VS", nil)
	if err != nil {
		return nil, err
	}

	propagationState := new(PropagationState)
	_, err = gm.client.Do(ctx, request, &propagationState)
	return propagationState, err
}

// SetGcmPropagation2VS Sets possibility for GCM settings propagation from main server to virtual servers.
// GCM settings can be propagated to virtual server in one case only - if it is absent on virtual server.
func (gm *Gcm) SetGcmPropagation2VS(ctx context.Context, params PropagationState) (*PropagationState, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", gm.client.Server+"/api/v1.0/Gcm.SetGcmPropagation2VS",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	_, err = gm.client.Do(ctx, request, nil)
	return nil, err
}

// GCM Google Cloud Messaging
type GCM struct {
	// WstrSenderID SENDER_ID (project number), can be retrieved from Google Developers Console. Can't be empty.
	WstrSenderID string `json:"wstrSenderId"`
	// WstrAPIKey API key, can be retrieved from Google Developers Console. Can't be empty.
	WstrAPIKey string `json:"wstrApiKey"`
	PxgRetVal  bool   `json:"PxgRetVal,omitempty"`
}

// UpdateGcmServerSettings Update GCM (Google Cloud Messaging) server settings.
func (gm *Gcm) UpdateGcmServerSettings(ctx context.Context, params GCM) (*PxgValBool, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", gm.client.Server+"/api/v1.0/Gcm.UpdateGcmServerSettings",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	result := new(PxgValBool)
	_, err = gm.client.Do(ctx, request, &result)
	return result, err
}

// GetGcmServerSettings Retrieves GCM (Google Cloud Messaging) server settings.
func (gm *Gcm) GetGcmServerSettings(ctx context.Context) (*GCM, error) {
	request, err := http.NewRequest("POST", gm.client.Server+"/api/v1.0/Gcm.GetGcmServerSettings", nil)
	if err != nil {
		return nil, err
	}

	gcm := new(GCM)
	_, err = gm.client.Do(ctx, request, &gcm)
	return gcm, err
}
