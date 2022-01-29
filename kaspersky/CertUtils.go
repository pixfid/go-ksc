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
	"net/http"
)

// CertUtils Helpers for managing certificates.
type CertUtils service

type SelfSignedCERTParams struct {
	PParams SSCParams `json:"pParams"`
}

type SSCParams struct {
	KlsrvCertutilsCERTCommonName string `json:"KLSRV_CERTUTILS_CERT_COMMON_NAME"`
	KlsrvCertutilsCERTDNSNames   string `json:"KLSRV_CERTUTILS_CERT_DNS_NAMES"`
	KlsrvCertutilsCERTLifeTime   int64  `json:"KLSRV_CERTUTILS_CERT_LIFE_TIME"`
	KlsrvCertutilsCERTEkuServer  bool   `json:"KLSRV_CERTUTILS_CERT_EKU_SERVER"`
	KlsrvCertutilsCERTEkuClient  bool   `json:"KLSRV_CERTUTILS_CERT_EKU_CLIENT"`
}

type SelfSignedCERTResponse struct {
	PxgRetVal SSCRetVal `json:"PxgRetVal"`
}

type SSCRetVal struct {
	KlsrvCertutilsCERTHash        string                  `json:"KLSRV_CERTUTILS_CERT_HASH"`
	KlsrvCertutilsCERTPassword    string                  `json:"KLSRV_CERTUTILS_CERT_PASSWORD"`
	KlsrvCertutilsCERTPrivatePart KlsrvCertutilsCERTPPart `json:"KLSRV_CERTUTILS_CERT_PRIVATE_PART"`
	KlsrvCertutilsCERTPublicPart  KlsrvCertutilsCERTPPart `json:"KLSRV_CERTUTILS_CERT_PUBLIC_PART"`
}

type KlsrvCertutilsCERTPPart struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (cu *CertUtils) GenerateSelfSignedCertificate(ctx context.Context, params SelfSignedCERTParams) (*SelfSignedCERTResponse, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", cu.client.Server+"/api/v1.0/CertUtils.GenerateSelfSignedCertificate",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}
	selfSignedCERTResponse := new(SelfSignedCERTResponse)
	raw, err := cu.client.Request(ctx, request, &selfSignedCERTResponse)
	return selfSignedCERTResponse, raw, err
}
