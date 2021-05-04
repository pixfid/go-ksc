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
	"context"
	"log"
	"net/http"
)

// MfaCacheInner MFA Cache provider (Only inside server)
type MfaCacheInner service

// GetMfaRequiredForAll Get "MFA required for all users" flag
func (mci *MfaCacheInner) GetMfaRequiredForAll(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", mci.client.Server+"/api/v1.0/MfaCacheInner.GetMfaRequiredForAll", nil)
	if err != nil {
		return nil, err
	}

	raw, err := mci.client.Do(ctx, request, nil)

	if mci.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return raw, err
}

func (mci *MfaCacheInner) GetMfaKeyIssuer(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", mci.client.Server+"/api/v1.0/MfaCacheInner.GetMfaKeyIssuer", nil)
	if err != nil {
		return nil, err
	}

	raw, err := mci.client.Do(ctx, request, nil)

	if mci.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return raw, err
}

func (mci *MfaCacheInner) GetTotpSecretKeySettings(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", mci.client.Server+"/api/v1.0/MfaCacheInner.GetTotpSecretKeySettings", nil)
	if err != nil {
		return nil, err
	}

	raw, err := mci.client.Do(ctx, request, nil)

	if mci.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return raw, err
}

func (mci *MfaCacheInner) GetTotpVerifySettings(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", mci.client.Server+"/api/v1.0/MfaCacheInner.GetTotpVerifySettings", nil)
	if err != nil {
		return nil, err
	}

	raw, err := mci.client.Do(ctx, request, nil)

	if mci.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return raw, err
}

func (mci *MfaCacheInner) IsCurrentUserExcludesMfa(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", mci.client.Server+"/api/v1.0/MfaCacheInner.IsCurrentUserExcludesMfa", nil)
	if err != nil {
		return nil, err
	}

	raw, err := mci.client.Do(ctx, request, nil)

	if mci.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return raw, err
}

//TODO AddUserToTotpRequrementExceptions
//TODO GetUserTotpSecret
//TODO GetUserTotpSecret
//TODO IsUserExcludesMfa
//TODO IsUserExcludesMfa
//TODO SetUserTotpSecret
//TODO SetUserTotpSecret
