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

//	IWebSrvSettings Class Reference
//
//	Web server settings proxy class.
//
//	List of all members.
type IWebSrvSettings service

//	GetCertificateInfo. Returns information about custom certificate.
//
//
//	If cert present then it return params with [["CERT_TYPE"] == 0 (PEM form)] and ["CERT_PUBLIC_PART"] fields.
//	In case if certificate not set, then it returns empty params with no any fields.
//
//	Returns:
//	Returned data format:
//	- "CERT_TYPE"         Certificate type (0 - PEM form, 1 - PKCS#12 form);
//	- "CERT_PUBLIC_PART"  Certificate's public part.
//
func (iws *IWebSrvSettings) GetCertificateInfo(ctx context.Context) (*PxgValCIFIL, []byte, error) {
	request, err := http.NewRequest("POST", iws.client.Server+"/api/v1.0/IWebSrvSettings.GetCertificateInfo", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValCIFIL := new(PxgValCIFIL)
	raw, err := iws.client.Do(ctx, request, &pxgValCIFIL)
	return pxgValCIFIL, raw, err
}

//	GetCustomPkgHttpFqdn. Returns custom HTTP FQDN.
//
//	Returns:
//	- Custom HTTP FQDN.
func (iws *IWebSrvSettings) GetCustomPkgHttpFqdn(ctx context.Context) (*PxgValStr, []byte, error) {
	request, err := http.NewRequest("POST", iws.client.Server+"/api/v1.0/IWebSrvSettings.GetCustomPkgHttpFqdn", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := iws.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//	SetCustomPkgHttpFqdn. Set's custom HTTP FQDN. It is useful for HTTP link generation.
//
//	Parameters:
//	- wsFqdn [string] Custom fqdn.
//
//	Exceptions:
//	- Throws	KLSTD::STDE_BADPARAM in case of incorrect FQDN.
func (iws *IWebSrvSettings) SetCustomPkgHttpFqdn(ctx context.Context, wsFqdn string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wsFqdn": "%s"}`, wsFqdn))
	request, err := http.NewRequest("POST", iws.client.Server+"/api/v1.0/IWebSrvSettings.SetCustomPkgHttpFqdn", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := iws.client.Do(ctx, request, nil)
	return raw, err
}

//	SetCustomCertificate. Sets custom certificate for Web Server's SSL listener.
//
//
//	FQDN name from certificate are used for HTTPS link generation.
//
//	Parameters:
//	- pCertData	[in] Params with certificate data (see Common format for certificate params).
func (iws *IWebSrvSettings) SetCustomCertificate(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", iws.client.Server+"/api/v1.0/IWebSrvSettings.SetCustomCertificate",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := iws.client.Do(ctx, request, nil)
	return raw, err
}
