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

//	OsVersion Class Reference
//
//	Operating systems dictionary access.
//
//	Provides access to the server-side operating systems dictionary
//
//	List of all members.
type OsVersion service

//OSIndices struct
type OSIndices struct {
	POSIndices []int64 `json:"pOsIndices"`
}

//OSAttributes struct
type OSAttributes struct {
	OSAttribute *[]OSAttribute `json:"PxgRetVal"`
}

type OSAttribute struct {
	Type             string            `json:"type,omitempty"`
	OSAttributeValue *OSAttributeValue `json:"value,omitempty"`
}

type OSAttributeValue struct {
	KlhstWksCtype         int64  `json:"KLHST_WKS_CTYPE,omitempty"`
	KlhstWksOSBuildNumber int64  `json:"KLHST_WKS_OS_BUILD_NUMBER,omitempty"`
	KlhstWksOSName        string `json:"KLHST_WKS_OS_NAME,omitempty"`
	KlhstWksOSVerMajor    int64  `json:"KLHST_WKS_OS_VER_MAJOR,omitempty"`
	KlhstWksOSVerMinor    int64  `json:"KLHST_WKS_OS_VER_MINOR,omitempty"`
	KlhstWksPtype         int64  `json:"KLHST_WKS_PTYPE,omitempty"`
	KlwnfOSIsServer       bool   `json:"KLWNF_OS_IS_SERVER,omitempty"`
}

//	Acquire attributes for specified operating systems.
//
//	Returns values of attributes for specified operating systems.
//
//	Parameters:
//	- pOsIndices	paramArray, array of operating system IDs, each entry has type (paramInt) (see KLWNF_OS_ID)
//
//	Returns:
//	- paramArray of the same size as pOsIndices, each entry is a (paramParams) container.
//	The entry container is either empty (if bad ID was put into the corresponding entry of the pOsIndices array) or filled with following attributes:
//	|- KLHST_WKS_CTYPE
//	|- KLHST_WKS_PTYPE
//	|- KLHST_WKS_OS_VER_MAJOR
//	|- KLHST_WKS_OS_VER_MINOR
//	|- KLHST_WKS_OS_NAME
//	|- KLWNF_OS_IS_SERVER
//	|- KLHST_WKS_OS_BUILD_NUMBER
//
//	See also:
//	- Operating systems dictionary
//	- Operating system attributes
func (ov *OsVersion) GetAttributesByOs(ctx context.Context, params OSIndices) (*OSAttributes, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ov.client.Server+"/api/v1.0/OsVersion.GetAttributesByOs", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	oSAttributes := new(OSAttributes)
	raw, err := ov.client.Do(ctx, request, &oSAttributes)
	return oSAttributes, raw, err
}

//OSRetValS struct
type OSRetValS struct {
	OSRetVal []OSRetVal `json:"PxgRetVal"`
}

type OSRetVal struct {
	Type    string   `json:"type,omitempty"`
	OSValue *OSValue `json:"value,omitempty"`
}

type OSValue struct {
	KlhstWksOSName  string `json:"KLHST_WKS_OS_NAME,omitempty"`
	KlwnfOSID       int64  `json:"KLWNF_OS_ID,omitempty"`
	KlwnfOSIsServer bool   `json:"KLWNF_OS_IS_SERVER,omitempty"`
}

//	Determine operating system by specified attributes.
//
//	Determines operating system by specified attributes
//	KLHST_WKS_CTYPE,
//	KLHST_WKS_PTYPE,
//	KLHST_WKS_OS_VER_MAJOR,
//	KLHST_WKS_OS_VER_MINOR,
//	KLHST_WKS_OS_BUILD_NUMBER.
//
//	Parameters:
//	- pDataToResolve	paramArray, each entry is a (paramParams) container filled with following attributes:
//	|- KLHST_WKS_CTYPE
//	|- KLHST_WKS_PTYPE
//	|- KLHST_WKS_OS_VER_MAJOR
//	|- KLHST_WKS_OS_VER_MINOR
//	|- KLHST_WKS_OS_BUILD_NUMBER (optional)
//	|- KLHST_WKS_OS_NAME
//	|- KLWNF_OS_IS_SERVER
//
// 	Example params interface{}:
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
//
//	Returns:
//	- paramArray of the same size as pDataToResolve, each entry is a (paramParams) container.
//	The entry container is either empty (if failed to determine the operating system) or is filled with following attributes:
//	|- KLWNF_OS_ID
//	|- KLHST_WKS_OS_NAME
//	|- KLWNF_OS_IS_SERVER
//
//	See also:
//	- Operating systems dictionary
//	- Operating system attributes
func (ov *OsVersion) GetOsByAttributes(ctx context.Context, params interface{}) (*OSRetValS, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ov.client.Server+"/api/v1.0/OsVersion.GetOsByAttributes", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	oSRetValS := new(OSRetValS)
	raw, err := ov.client.Do(ctx, request, &oSRetValS)
	return oSRetValS, raw, err
}
