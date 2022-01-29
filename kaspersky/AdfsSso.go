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

// AdfsSso service for working with ADFS SSO. This service allow you to manage ADFS SSO settings
type AdfsSso service

// GetSettings Returns a ADFS SSO settings.
func (as *AdfsSso) GetSettings(ctx context.Context, bExtenedSettings bool) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"bExtenedSettings": %v}`, bExtenedSettings))
	request, err := http.NewRequest("POST", as.client.Server+"/api/v1.0/AdfsSso.GetSettings", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := as.client.Request(ctx, request, nil)
	return raw, err
}

// SetSettings Set a ADFS SSO settings.
func (as *AdfsSso) SetSettings(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", as.client.Server+"/api/v1.0/AdfsSso.SetSettings", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := as.client.Request(ctx, request, nil)
	return raw, err
}

// GetAdfsEnabled Get a ADFS SSO enabled/disabled.
func (as *AdfsSso) GetAdfsEnabled(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", as.client.Server+"/api/v1.0/AdfsSso.GetAdfsEnabled", nil)
	if err != nil {
		return nil, err
	}

	raw, err := as.client.Request(ctx, request, nil)
	return raw, err
}

// GetJwks Returns a ADFS JWKS.
func (as *AdfsSso) GetJwks(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", as.client.Server+"/api/v1.0/AdfsSso.GetJwks", nil)
	if err != nil {
		return nil, err
	}

	raw, err := as.client.Request(ctx, request, nil)
	return raw, err
}

// SetAdfsEnabled Set a ADFS SSO enabled/disabled.
func (as *AdfsSso) SetAdfsEnabled(ctx context.Context, bEnabled bool) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"bEnabled": %v}`, bEnabled))
	request, err := http.NewRequest("POST", as.client.Server+"/api/v1.0/AdfsSso.SetAdfsEnabled", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := as.client.Request(ctx, request, nil)
	return raw, err
}
