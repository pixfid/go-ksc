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

//	KeyService Class Reference
//	Interface for working with KeyService subsystem.
//
//	List of all members.
type KeyService service

type PEncryptedData struct {
	PEncryptedData string `json:"pEncryptedData"`
}

//	Method creates crypto container.
//
//	Parameters:
//	- pData	(binary base64 string) Data to encrypt (Max size is 512 KB).
//
//	Return:
//	- pEncryptedData	(binary base64 string) Encrypted data (Crypto container format).
func (ks *KeyService) EncryptData(ctx context.Context, pData string) (*PEncryptedData, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"pData": "%s"}`, pData))
	request, err := http.NewRequest("POST", ks.client.Server+"/api/v1.0/KeyService.EncryptData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pEncryptedData := new(PEncryptedData)
	raw, err := ks.client.Do(ctx, request, &pEncryptedData)
	return pEncryptedData, raw, err
}

type PDecryptedData struct {
	PDecryptedData string `json:"pDecryptedData"`
}

//	Method unprotects crypto container created by EncryptData.
//
//	Parameters:
//	- pEncryptedData	(binary base64 string) Encrypted data
//	- wstrProdName	[optional] (string) Product name
//	- wstrProdVersion	[optional] (string) Product version
//
//	Return:
//	- pDecryptedData	(binary base64 string) Decrypted data
func (ks *KeyService) DecryptData(ctx context.Context, pEncryptedData, wstrProdName, wstrProdVersion string) (*PDecryptedData,
	[]byte, error) {
	postData := []byte(fmt.Sprintf(`{"pEncryptedData": "%s", "wstrProdName": "%s", "wstrProdVersion": "%s"}`,
		pEncryptedData, wstrProdName, wstrProdVersion))
	request, err := http.NewRequest("POST", ks.client.Server+"/api/v1.0/KeyService.DecryptData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pDecryptedData := new(PDecryptedData)
	raw, err := ks.client.Do(ctx, request, &pDecryptedData)
	return pDecryptedData, raw, err
}

//	Method creates a crypto container for chosen host. Data may be decrypted only locally on host.
//
//	Parameters:
//	- wstrHostId	(string) Host id.
//	- pData	(binary base64 string) Data to encrypt (Max size is 512 KB).
//
//	Return:
//	- pEncryptedData	(binary base64 string) Encrypted data.
//
//	Exceptions:
//	KLSTD::STDE_NOTFOUND	May be thrown if managed host is not synced and key not found in db. You should wait a period of synchronization.
func (ks *KeyService) EncryptDataForHost(ctx context.Context, wstrHostId, pData string) (*PEncryptedData, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"wstrHostId" : "%s", "pData": "%s"}`, wstrHostId, pData))
	request, err := http.NewRequest("POST", ks.client.Server+"/api/v1.0/KeyService.EncryptDataForHost", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pEncryptedData := new(PEncryptedData)
	raw, err := ks.client.Do(ctx, request, &pEncryptedData)
	return pEncryptedData, raw, err
}

//TransportCertificate struct using in KeyService.GenerateTransportCertificate
type TransportCertificate struct {
	//Public part of certificate
	PPublic string `json:"pPublic"`

	//Private part of certificate
	PPrivate string `json:"pPrivate"`

	//Password for private part
	WstrPass string `json:"wstrPass"`
}

//	Method generates transport certificate.
//
//	Parameters:
//	- wstrCommonName	(string) Common name
//
//	Return:
//	- pPublic	(binary base64 string) Public part of certificate
//	- pPrivate	(binary base64 string) Private part of certificate
//	- wstrPass	(string) Password for private part
func (ks *KeyService) GenerateTransportCertificate(ctx context.Context, wstrCommonName string) (*TransportCertificate, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrCommonName": "%s"}`, wstrCommonName))
	request, err := http.NewRequest("POST", ks.client.Server+"/api/v1.0/KeyService.GenerateTransportCertificate", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	transportCertificate := new(TransportCertificate)
	raw, err := ks.client.Do(ctx, request, &transportCertificate)
	return transportCertificate, raw, err
}
