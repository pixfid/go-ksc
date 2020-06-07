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

//	DataProtectionApi Class Reference
//
//	Allows to protect sensitive data in policies, tasks, and/or on specified host.
//
//	List of all members.
type DataProtectionApi service

//	Checks if Spl password policy compliance is enabled for the specified Administration Server
//	and checks the specified password for compliance to the password policy.
//
//	Password Policy is specified below
//
//	Characters allowed:
//
//	A – Z
//	a – z
//	0 – 9
//	@ # $ % ^ & * - _ ! + = [ ] { } | \ : ‘ , . ? / ` ~ “ ( ) ;
//
//	Characters disallowed:
//
//	- Unicode characters
//	- spaces
//	- Cannot contain a dot character '.' immediately preceding the '@' symbol
//
//	Password restrictions:
//
//	- 8 characters minimum and 16 characters maximum
//	- Must contain characters at least from any 3 of 4 groups mentioned in the section "Characters allowed"
//
//	Parameters:
//	- szwPassword	(string)	The password to check.
//
//	Exceptions:
//	- KLSTD.STDE_NOFUNC	the password does not comply with the password policy
func (dpa *DataProtectionApi) CheckPasswordSplPpc(ctx context.Context, szwPassword string) (*PxgValBool, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwPassword": "%s"}`, szwPassword))
	request, err := http.NewRequest("POST", dpa.client.Server+"/api/v1.0/DataProtectionApi.CheckPasswordSplPpc", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := dpa.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

type ToProtectData struct {
	SzwHostID string `json:"szwHostId,omitempty"`
	PData     string `json:"pData,omitempty"`
}

type ProtectedData struct {
	PDataProtected string `json:"pDataProtected,omitempty"`
}

//	Protects sensitive data to store in SettingsStorage or local task.
//
//	Parameters:
//	- szwHostId host name
//	- pData pointer to data
//
//	- pDataProtected pointer to protected data block.
//
//Exceptions:
//	- KLSTD.STDE_NOTPERM	host has no public key (
//	possibly it doesn't support data protection or nagent isn't installed
//	there or host belongs to other virtual server)
//	- KLSTD.STDE_NOFUNC	server doesn't support data protection
func (dpa *DataProtectionApi) ProtectDataForHost(ctx context.Context, params ToProtectData) (*ProtectedData, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", dpa.client.Server+"/api/v1.0/DataProtectionApi.ProtectDataForHost", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}
	protectedData := new(ProtectedData)
	raw, err := dpa.client.Do(ctx, request, &protectedData)
	return protectedData, raw, err
}

//	Protects sensitive data to store in policy or global/group task.
//
//	Parameters:
//	- pData pointer to data
//	- pDataProtected pointer to protected data block.
func (dpa *DataProtectionApi) ProtectDataGlobally(ctx context.Context, params ToProtectData) (*ProtectedData, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", dpa.client.Server+"/api/v1.0/DataProtectionApi.ProtectDataGlobally", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	protectedData := new(ProtectedData)
	raw, err := dpa.client.Do(ctx, request, &protectedData)
	return protectedData, raw, err
}

//	Protects sensitive data for the specified host (to store in its local settings or a local task)
//
//	Protects the specified text as UTF16 string encrypted with the key of the specified host.
//
//	Parameters:
//	- szwHostId host name
//	- szwPlainText plainText
//
//	Returns:
//	- Ciphertext
func (dpa *DataProtectionApi) ProtectUtf16StringForHost(ctx context.Context, szwHostId, szwPlainText string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"szwPassword" : "%s", "szwPlainText" : "%s"}`, szwHostId, szwPlainText))
	request, err := http.NewRequest("POST", dpa.client.Server+"/api/v1.0/DataProtectionApi.ProtectUtf16StringForHost", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := dpa.client.Do(ctx, request, nil)
	return raw, err
}

//	Protects sensitive data to store in policy, global/group task,
//	Administration Server settings.
//
//	Protects the specified text as UTF16 string encrypted with the key
//	of the Administration Server.
//
//	The same as Tasks.ProtectPassword
//
//	Parameters:
//	- szwPlainText plainText
//
//Returns:
//	- Ciphertext
func (dpa *DataProtectionApi) ProtectUtf16StringGlobally(ctx context.Context, szwPlainText string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"szwPlainText" : "%s"}`, szwPlainText))
	request, err := http.NewRequest("POST", dpa.client.Server+"/api/v1.0/DataProtectionApi.ProtectUtf16StringGlobally", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := dpa.client.Do(ctx, request, nil)
	return raw, err
}

//	Protects sensitive data for the specified host
//	(to store in its local settings or a local task)
//
//	Protects the specified text as UTF8 string encrypted with the key
//	of the specified host.
//
//	Parameters:
//	- szwHostId host name
//	- szwPlainText plainText
//
//Returns:
//	- Ciphertext
func (dpa *DataProtectionApi) ProtectUtf8StringForHost(ctx context.Context, szwHostId, szwPlainText string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"szwPassword" : "%s", "szwPlainText" : "%s"}`, szwHostId, szwPlainText))
	request, err := http.NewRequest("POST", dpa.client.Server+"/api/v1.0/DataProtectionApi.ProtectUtf8StringForHost", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := dpa.client.Do(ctx, request, nil)
	return raw, err
}

//	Protects sensitive data to store in policy, global/group task,
//	Administration Server settings.
//
//	Protects the specified text as UTF8 string encrypted with the key
//	of the Administration Server.
//
//	Parameters:
//	- szwPlainText plainText
//
//	Returns:
//	- Ciphertext
func (dpa *DataProtectionApi) ProtectUtf8StringGlobally(ctx context.Context, szwPlainText string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"szwPlainText" : "%s"}`, szwPlainText))
	request, err := http.NewRequest("POST", dpa.client.Server+"/api/v1.0/DataProtectionApi.ProtectUtf8StringGlobally", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := dpa.client.Do(ctx, request, nil)
	return raw, err
}
