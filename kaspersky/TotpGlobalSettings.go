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

// TotpGlobalSettings 2FA global settings
type TotpGlobalSettings service

// Get2FaRequiredForAll Read global flag "2FA is required for all users".
func (tgs *TotpGlobalSettings) Get2FaRequiredForAll(ctx context.Context) (*PxgValBool, error) {
	request, err := http.NewRequest("POST", tgs.client.Server+"/api/v1.0/TotpGlobalSettings.Get2FaRequiredForAll", nil)
	if err != nil {
		return nil, err
	}

	result := new(PxgValBool)
	raw, err := tgs.client.Request(ctx, request, &result)

	if tgs.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return result, err
}

type TOTPSettings struct {
	PSettings struct {
		TotpIssuer string `json:"TotpIssuer"`
	} `json:"pSettings"`
}

// GetTotpGlobalSettings Read global TOTP settings.
func (tgs *TotpGlobalSettings) GetTotpGlobalSettings(ctx context.Context) (*TOTPSettings, error) {
	request, err := http.NewRequest("POST", tgs.client.Server+"/api/v1.0/TotpGlobalSettings.GetTotpGlobalSettings", nil)
	if err != nil {
		return nil, err
	}

	result := new(TOTPSettings)
	raw, err := tgs.client.Request(ctx, request, &result)

	if tgs.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return result, err
}

type LoggedInUsing2FA struct {
	BLoggedInUsing2FA bool `json:"bLoggedInUsing2FA"`
	PxgRetVal         bool `json:"PxgRetVal"`
}

// IfCanConfigure2FaSettings Checks whether calling user has enough access rights to modify 2FA settings
// Used by console to decide whether to display configuration options and whether to start setting up 2FA for current user.
// 2FA settings modification is allowed only if this method returns true AND bLoggedInUsing2FA is true.
func (tgs *TotpGlobalSettings) IfCanConfigure2FaSettings(ctx context.Context) (*LoggedInUsing2FA, error) {
	request, err := http.NewRequest("POST", tgs.client.Server+"/api/v1.0/TotpGlobalSettings.IfCanConfigure2FaSettings", nil)
	if err != nil {
		return nil, err
	}

	result := new(LoggedInUsing2FA)
	raw, err := tgs.client.Request(ctx, request, &result)

	if tgs.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return result, err
}

// Set2FaRequiredForAll Set global flag "2FA is required for all users"
// Second factor must be configured for current user to be able to call this method.
func (tgs *TotpGlobalSettings) Set2FaRequiredForAll(ctx context.Context, bRequiredForAll bool) error {
	postData := []byte(fmt.Sprintf(`{"bRequiredForAll" : %v}`, bRequiredForAll))

	request, err := http.NewRequest("POST", tgs.client.Server+"/api/v1.0/TotpGlobalSettings.Set2FaRequiredForAll",
		bytes.NewBuffer(postData))

	if err != nil {
		return err
	}

	raw, err := tgs.client.Request(ctx, request, nil)

	if tgs.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}

//TODO SetTotpGlobalSettings
