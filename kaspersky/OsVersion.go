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

// OsVersion Operating systems dictionary access.
//
// Provides access to the server-side operating systems dictionary
type OsVersion service

// OSIndices struct
type OSIndices struct {
	POSIndices []int64 `json:"pOsIndices"`
}

// OSAttributes struct
type OSAttributes struct {
	OSAttribute *[]OSAttribute `json:"PxgRetVal"`
}

// OSAttribute struct
type OSAttribute struct {
	Type             string            `json:"type,omitempty"`
	OSAttributeValue *OSAttributeValue `json:"value,omitempty"`
}

// OSAttributeValue struct
type OSAttributeValue struct {
	KlhstWksCtype         int64  `json:"KLHST_WKS_CTYPE,omitempty"`
	KlhstWksOSBuildNumber int64  `json:"KLHST_WKS_OS_BUILD_NUMBER,omitempty"`
	KlhstWksOSName        string `json:"KLHST_WKS_OS_NAME,omitempty"`
	KlhstWksOSVerMajor    int64  `json:"KLHST_WKS_OS_VER_MAJOR,omitempty"`
	KlhstWksOSVerMinor    int64  `json:"KLHST_WKS_OS_VER_MINOR,omitempty"`
	KlhstWksPtype         int64  `json:"KLHST_WKS_PTYPE,omitempty"`
	KlwnfOSIsServer       bool   `json:"KLWNF_OS_IS_SERVER,omitempty"`
}

// GetAttributesByOs Acquire attributes for specified operating systems.
func (ov *OsVersion) GetAttributesByOs(ctx context.Context, params OSIndices) (*OSAttributes, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ov.client.Server+"/api/v1.0/OsVersion.GetAttributesByOs", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	oSAttributes := new(OSAttributes)
	raw, err := ov.client.Request(ctx, request, &oSAttributes)
	return oSAttributes, raw, err
}

// OSRetValS struct
type OSRetValS struct {
	OSRetVal []OSRetVal `json:"PxgRetVal"`
}

// OSRetVal struct
type OSRetVal struct {
	Type    string   `json:"type,omitempty"`
	OSValue *OSValue `json:"value,omitempty"`
}

// OSValue struct
type OSValue struct {
	KlhstWksOSName  string `json:"KLHST_WKS_OS_NAME,omitempty"`
	KlwnfOSID       int64  `json:"KLWNF_OS_ID,omitempty"`
	KlwnfOSIsServer bool   `json:"KLWNF_OS_IS_SERVER,omitempty"`
}

// GetOsByAttributes Determine operating system by specified attributes.
//
// Example:
//
//	{
//  "pDataToResolve" : [
//    {
//      "type" : "params",
//      "value" : {
//        "KLHST_WKS_CTYPE" : 4194304,
//        "KLHST_WKS_OS_BUILD_NUMBER" : 0,
//        "KLHST_WKS_OS_NAME" : "Microsoft Windows 98",
//        "KLHST_WKS_OS_VER_MAJOR" : 4,
//        "KLHST_WKS_OS_VER_MINOR" : 10,
//        "KLHST_WKS_PTYPE" : 0,
//        "KLWNF_OS_IS_SERVER" : false
//      }
//    }
//  ]
//}
func (ov *OsVersion) GetOsByAttributes(ctx context.Context, params interface{}) (*OSRetValS, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ov.client.Server+"/api/v1.0/OsVersion.GetOsByAttributes", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	oSRetValS := new(OSRetValS)
	raw, err := ov.client.Request(ctx, request, &oSRetValS)
	return oSRetValS, raw, err
}
