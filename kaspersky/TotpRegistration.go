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
	"log"
	"net/http"
)

// TotpRegistration TOTP registration.
type TotpRegistration service

type TotpSecretData struct {
	PSecret struct {
		TotpSecret   string `json:"TotpSecret"`
		TotpSecretQR struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"TotpSecretQR"`
		TotpSecretURI string `json:"TotpSecretUri"`
	} `json:"pSecret"`
	PxgRetVal string `json:"PxgRetVal"`
}

// GenerateSecret Generates TotpSecret with random key.
func (tr *TotpRegistration) GenerateSecret(ctx context.Context) (*TotpSecretData, error) {
	request, err := http.NewRequest("POST", tr.client.Server+"/api/v1.0/TotpRegistration.GenerateSecret", nil)
	if err != nil {
		return nil, err
	}

	totpSecretData := new(TotpSecretData)
	raw, err := tr.client.Do(ctx, request, &totpSecretData)

	if tr.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return totpSecretData, err
}

// IfCurrentUserMayClearSecret Checks whether current user is allowed to reset his TOTP secret.
func (tr *TotpRegistration) IfCurrentUserMayClearSecret(ctx context.Context) (*PxgValBool, error) {
	request, err := http.NewRequest("POST", tr.client.Server+"/api/v1.0/TotpRegistration.IfCurrentUserMayClearSecret", nil)
	if err != nil {
		return nil, err
	}

	result := new(PxgValBool)
	raw, err := tr.client.Do(ctx, request, &result)

	if tr.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return result, err
}

// SaveSecretForCurrentUser Set TOTP 2FA authentication for current user.
// If validation code does not pass verification, current user TOTP settings don't change.
//If validation code passes verification, secret will be used as second factor when users logs in.
//If user already has registered TOTP secret, it will be replaced with new one.
func (tr *TotpRegistration) SaveSecretForCurrentUser(ctx context.Context, wstrSecretId, wstrValidationCode string) error {
	postData := []byte(fmt.Sprintf(`{"wstrSecretId" : "%s", "wstrValidationCode" : "%s"}`, wstrSecretId, wstrValidationCode))

	request, err := http.NewRequest("POST", tr.client.Server+"/api/v1.0/TotpRegistration.SaveSecretForCurrentUser",
		bytes.NewBuffer(postData))

	if err != nil {
		return err
	}

	raw, err := tr.client.Do(ctx, request, nil)

	if tr.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}

// DeleteSecret Clears secret, identified by wstrSecretId.
func (tr *TotpRegistration) DeleteSecret(ctx context.Context, wstrSecretId string) error {
	postData := []byte(fmt.Sprintf(`{"wstrSecretId" : "%s"}`, wstrSecretId))

	request, err := http.NewRequest("POST", tr.client.Server+"/api/v1.0/TotpRegistration.DeleteSecret",
		bytes.NewBuffer(postData))

	if err != nil {
		return err
	}

	raw, err := tr.client.Do(ctx, request, nil)

	if tr.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}

// ClearSecretForCurrentUser Clears TOTP secret for current user. Function succeeds only
// if '2FA required for all' flag is not set, or if user is in exceptions list.
func (tr *TotpRegistration) ClearSecretForCurrentUser(ctx context.Context) (*PxgValBool, error) {
	request, err := http.NewRequest("POST", tr.client.Server+"/api/v1.0/TotpRegistration.ClearSecretForCurrentUser", nil)

	if err != nil {
		return nil, err
	}

	result := new(PxgValBool)
	raw, err := tr.client.Do(ctx, request, &result)

	if tr.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return result, err
}
