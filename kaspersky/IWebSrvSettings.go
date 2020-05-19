/*
 *
 * 	Copyright (C) 2020  <Semchenko Aleksandr>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.If not, see <http://www.gnu.org/licenses/>.
 * /
 */

package kaspersky

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

//	IWebSrvSettings Class Reference
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

//TODO func (iws *IWebSrvSettings) SetCustomCertificate(ctx context.Context, wsFqdn string) ([]byte, error) {}
