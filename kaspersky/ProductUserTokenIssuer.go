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

// ProductUserTokenIssuer Product backend user token issuer.
//
// Allows to issue/revoke product backend tokens for current user
type ProductUserTokenIssuer service

// IssueUserToken Issue user token.
//
// Underlying backend command is "IssueUserToken" - get new user access/refresh token pair.
//All issued tokens for current user automatically revoked by KSC server when user rights changed for specified product
//or when KSC server restarts.
//Command result may be obtained using AsyncActionStateChecker.
//Command result format is defined by product backend, errors returned as Asynchronous action errors
func (srn *ProductUserTokenIssuer) IssueUserToken(ctx context.Context, wstrProductName, wstrProductVersion string) error {
	postData := []byte(fmt.Sprintf(`{"wstrProductName": "%s", "wstrProductVersion": "%s"}`, wstrProductName, wstrProductVersion))
	request, err := http.NewRequest("POST", srn.client.Server+"/api/v1.0/ProductUserTokenIssuer.IssueUserToken",
		bytes.NewBuffer(postData))

	if err != nil {
		return err
	}

	//TODO check response
	raw, err := srn.client.Do(ctx, request, nil)

	if srn.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}

// RevokeUserToken Revoke user token.
//
// Underlying backend command is "RevokeUserToken" - revoke all tokens for specific user.
//Command result may be obtained using AsyncActionStateChecker.
//Command result format is defined by product backend, errors returned as Asynchronous action errors
func (srn *ProductUserTokenIssuer) RevokeUserToken(ctx context.Context, wstrProductName, wstrProductVersion string) error {
	postData := []byte(fmt.Sprintf(`{"wstrProductName": "%s", "wstrProductVersion": "%s"}`, wstrProductName, wstrProductVersion))
	request, err := http.NewRequest("POST", srn.client.Server+"/api/v1.0/ProductUserTokenIssuer.RevokeUserToken",
		bytes.NewBuffer(postData))

	if err != nil {
		return err
	}

	//TODO check response
	raw, err := srn.client.Do(ctx, request, nil)

	if srn.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}
