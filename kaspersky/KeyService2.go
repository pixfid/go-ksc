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

//	KeyService2 Class Reference
//	Additional interface for working with KeyService subsystem.
//
//	List of all members.
type KeyService2 service

//	Imports encryption keys.
//
//	Parameters:
//	- pProtectedPass	(binary base64 string) Password protected with DataProtectionApi.ProtectUtf8StringGlobally
//
//	Return:
//	- pEncryptedKeysData	(binary base64 string) Earlier exported keys data
func (ks2 *KeyService2) ImportDpeKeys(ctx context.Context, pProtectedPass string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"pProtectedPass": "%s"}`, pProtectedPass))
	request, err := http.NewRequest("POST", ks2.client.Server+"/api/v1.0/KeyService2.ImportDpeKeys", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ks2.client.Do(ctx, request, nil)
	return raw, err
}

//	Exports own and all stored encryption keys.
//
//	Parameters:
//	- pProtectedPass	(binary base64 string) Password protected with DataProtectionApi::ProtectUtf8StringGlobally
//
//	Return:
//	- pEncryptedKeysData	(binary base64 string) Serialized binary data
func (ks2 *KeyService2) ExportDpeKeys(ctx context.Context, pProtectedPass string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"pProtectedPass": "%s"}`, pProtectedPass))
	request, err := http.NewRequest("POST", ks2.client.Server+"/api/v1.0/KeyService2.ExportDpeKeys", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ks2.client.Do(ctx, request, nil)
	return raw, err
}
