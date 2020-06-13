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

//AdfsSso Class Reference
//
//Interface for working with ADFS SSO..
//
//This interface allow you to manage ADFS SSO settings
//
//List of all members.
type AdfsSso service

//AdfsSso.GetSettings
//Returns a ADFS SSO settings.
//
//	Returns:
//	- (params) ADFS SSO settings; See "ADFS SSO Settings".
func (as *AdfsSso) GetSettings(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", as.client.Server+"/api/v1.0/AdfsSso.GetSettings", nil)
	if err != nil {
		return nil, err
	}

	raw, err := as.client.Do(ctx, request, nil)
	return raw, err
}

//AdfsSso.SetSettings
//Set a ADFS SSO settings.
//
//	Parameters:
//	- pAdfsSettings	(params) ADFS SSO settings; "ADFS SSO Settings".
func (as *AdfsSso) SetSettings(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", as.client.Server+"/api/v1.0/AdfsSso.SetSettings", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := as.client.Do(ctx, request, nil)
	return raw, err
}
