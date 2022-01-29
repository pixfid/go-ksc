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

// TotpUserSettings 2FA user settings
type TotpUserSettings service

// ClearUserSecret Clear TOTP secret for user (spl_users, ak_users only)
func (tgs *TotpUserSettings) ClearUserSecret(ctx context.Context, llTrusteeID int) error {
	postData := []byte(fmt.Sprintf(`{"llTrusteeId" : %d}`, llTrusteeID))
	request, err := http.NewRequest("POST", tgs.client.Server+"/api/v1.0/TotpUserSettings.ClearUserSecret",
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

// AddUserToTotpRequrementExceptions Add/remove user to/from exceptions
// list of "2FA is required for all users" requirement.
func (tgs *TotpUserSettings) AddUserToTotpRequrementExceptions(ctx context.Context, llTrusteeID int, bInExceptions bool) error {
	postData := []byte(fmt.Sprintf(`{"llTrusteeId" : %d, "bInExceptions" : %v}`, llTrusteeID, bInExceptions))

	request, err := http.NewRequest("POST", tgs.client.Server+"/api/v1.0/TotpUserSettings.AddUserToTotpRequrementExceptions",
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

// IfCanClearUser2FaSecret Checks whether calling user has enough access rights to clear 2FA secret.
func (tgs *TotpUserSettings) IfCanClearUser2FaSecret(ctx context.Context, llTrusteeId int) (*LoggedInUsing2FA, error) {
	postData := []byte(fmt.Sprintf(`{"llTrusteeId" : %d}`, llTrusteeId))
	request, err := http.NewRequest("POST", tgs.client.Server+"/api/v1.0/TotpUserSettings.IfCanClearUser2FaSecret",
		bytes.NewBuffer(postData))

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
