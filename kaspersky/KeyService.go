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
	"net/http"
)

// KeyService service for working with KeyService subsystem.
type KeyService service

// PEncryptedData struct
type PEncryptedData struct {
	PEncryptedData string `json:"pEncryptedData"`
}

// EncryptData Method creates crypto container.
func (ks *KeyService) EncryptData(ctx context.Context, pData string) (*PEncryptedData, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"pData": "%s"}`, pData))
	request, err := http.NewRequest("POST", ks.client.Server+"/api/v1.0/KeyService.EncryptData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pEncryptedData := new(PEncryptedData)
	raw, err := ks.client.Request(ctx, request, &pEncryptedData)
	return pEncryptedData, raw, err
}

// PDecryptedData struct
type PDecryptedData struct {
	PDecryptedData string `json:"pDecryptedData"`
}

// DecryptData Method unprotects crypto container created by EncryptData.
func (ks *KeyService) DecryptData(ctx context.Context, pEncryptedData, wstrProdName, wstrProdVersion string) (*PDecryptedData,
	[]byte, error) {
	postData := []byte(fmt.Sprintf(`{"pEncryptedData": "%s", "wstrProdName": "%s", "wstrProdVersion": "%s"}`,
		pEncryptedData, wstrProdName, wstrProdVersion))
	request, err := http.NewRequest("POST", ks.client.Server+"/api/v1.0/KeyService.DecryptData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pDecryptedData := new(PDecryptedData)
	raw, err := ks.client.Request(ctx, request, &pDecryptedData)
	return pDecryptedData, raw, err
}

// EncryptDataForHost Method creates a crypto container for chosen host. Data may be decrypted only locally on host.
func (ks *KeyService) EncryptDataForHost(ctx context.Context, wstrHostId, pData string) (*PEncryptedData, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"wstrHostId" : "%s", "pData": "%s"}`, wstrHostId, pData))
	request, err := http.NewRequest("POST", ks.client.Server+"/api/v1.0/KeyService.EncryptDataForHost", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pEncryptedData := new(PEncryptedData)
	raw, err := ks.client.Request(ctx, request, &pEncryptedData)
	return pEncryptedData, raw, err
}

// TransportCertificate struct using in KeyService.GenerateTransportCertificate
type TransportCertificate struct {
	// PPublic public part of certificate
	PPublic string `json:"pPublic"`

	// PPrivate private part of certificate
	PPrivate string `json:"pPrivate"`

	// WstrPass password for private part
	WstrPass string `json:"wstrPass"`
}

// GenerateTransportCertificate Method generates transport certificate.
func (ks *KeyService) GenerateTransportCertificate(ctx context.Context, wstrCommonName string) (*TransportCertificate, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrCommonName": "%s"}`, wstrCommonName))
	request, err := http.NewRequest("POST", ks.client.Server+"/api/v1.0/KeyService.GenerateTransportCertificate", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	transportCertificate := new(TransportCertificate)
	raw, err := ks.client.Request(ctx, request, &transportCertificate)
	return transportCertificate, raw, err
}
