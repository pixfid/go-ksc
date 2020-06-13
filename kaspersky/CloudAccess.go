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

//CloudAccess Class Reference
//
//Interface to check access of public clouds.
//
//List of all members.
type CloudAccess service

type Credentials struct {
	//	EnCloudType Type of the cloud infrastructure being accessed (a KLCLOUD_TYPE_*,
	//now only Microsoft Azure is supported)
	//	╔══════╦═════════════════════╗
	//	║ Type ║     Description     ║
	//	╠══════╬═════════════════════╣
	//	║    0 ║ Unknown             ║
	//	║    1 ║ Amazon Web Services ║
	//	║    2 ║ Microsoft Azure     ║
	//	║    3 ║ Non cloud host      ║
	//	╚══════╩═════════════════════╝
	EnCloudType int64 `json:"enCloudType,omitempty"`

	//PKeyPair Cloud authentication credentials. (paramParams, mandatory)
	//AWS: Not supported Azure:
	PKeyPair *PKeyPair `json:"pKeyPair,omitempty"`
}

type PKeyPair struct {
	//Subscription ID (paramString, mandatory if there is "AZURE_APP_ID")
	ClientID string `json:"CLIENT_ID,omitempty"`

	//Application ID (paramString, mandatory if there is "CLIENT_SECRET")
	AzureAppID string `json:"AZURE_APP_ID,omitempty"`

	//UTF-8 encoded Application authentication key string encrypted with KLCSPWD::ProtectDataGlobally
	//(or KLCSPWD::ProtectDataLocally), (paramBinary, optional)
	ClientSecret string `json:"CLIENT_SECRET,omitempty"`

	//Storage account name (paramString, mandatory if there is "AZURE_STORAGE_KEY")
	AzureStorageName string `json:"AZURE_STORAGE_NAME,omitempty"`

	//UTF-8 encoded Storage account key string encrypted with KLCSPWD::ProtectDataGlobally
	//(or KLCSPWD::ProtectDataLocally), (paramBinary, optional)
	AzureStorageKey string `json:"AZURE_STORAGE_KEY,omitempty"`
}

//VerifyCredentials
//Verify credentials.
//
//	Parameters:
//	- params Credentials
//
//	Returns:
//	- true if all credentials is valid, false otherwise.
func (ca *CloudAccess) VerifyCredentials(ctx context.Context, params Credentials) (*PxgValBool, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, nil
	}

	request, err := http.NewRequest("POST", ca.client.Server+"/api/v1.0/CloudAccess.VerifyCredentials", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, nil
	}

	pxgValBool := new(PxgValBool)
	raw, err := ca.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

//KeyPairAccess struct
type KeyPairAccess struct {
	//Access for pcloud device discovery
	BAllowScanning bool `json:"bAllowScanning,omitempty"`

	//Access for remote deployment
	BAllowDeployment bool `json:"bAllowDeployment,omitempty"`
}

//AcquireAccessForKeyPair
//Check key-pair access
//
//	Parameters:
//	- params	Credentials
//
//	Return:
//	KeyPairAccess
func (ca *CloudAccess) AcquireAccessForKeyPair(ctx context.Context, params Credentials) (*KeyPairAccess, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, nil
	}

	request, err := http.NewRequest("POST", ca.client.Server+"/api/v1.0/CloudAccess.AcquireAccessForKeyPair", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, nil
	}

	keyPairAccess := new(KeyPairAccess)
	raw, err := ca.client.Do(ctx, request, &keyPairAccess)
	return keyPairAccess, raw, err
}
