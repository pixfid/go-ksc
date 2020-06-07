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

//	ExtAud Class Reference
//
//	Interface for working with ExtAudit subsystem.
//
//	This interface allow you to get a revision of an object and update description.
//
//	List of all members.
type ExtAud service

//
//	Get object revision.
//	╔════╦══════════════════════╗
//	║ ID ║     Description      ║
//	╠════╬══════════════════════╣
//	║  0 ║ None                 ║
//	║  1 ║ Policy               ║
//	║  2 ║ Task                 ║
//	║  3 ║ Package              ║
//	║  4 ║ Server               ║
//	║  5 ║ Virtual server       ║
//	║  6 ║ User                 ║
//	║  7 ║ Security Group       ║
//	║  8 ║ Administration Group ║
//	╚════╩══════════════════════╝
func (ea *ExtAud) GetRevision(ctx context.Context, nObjId, nObjType, nObjRevision int64, out interface{}) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"nObjId": %d,"nObjType": %d,"nObjRevision": %d}`, nObjId, nObjType, nObjRevision))
	request, err := http.NewRequest("POST", ea.client.Server+"/api/v1.0/ExtAud.GetRevision", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ea.client.Do(ctx, request, &out)
	return raw, err
}

//
//	Update revision description.
//	╔════╦══════════════════════╗
//	║ ID ║     Description      ║
//	╠════╬══════════════════════╣
//	║  0 ║ None                 ║
//	║  1 ║ Policy               ║
//	║  2 ║ Task                 ║
//	║  3 ║ Package              ║
//	║  4 ║ Server               ║
//	║  5 ║ Virtual server       ║
//	║  6 ║ User                 ║
//	║  7 ║ Security Group       ║
//	║  8 ║ Administration Group ║
//	╚════╩══════════════════════╝
func (ea *ExtAud) UpdateRevisionDesc(ctx context.Context, nObjId, nObjType, nObjRevision int64, wstrNewDescription string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"nObjId": %d,"nObjType": %d,"nObjRevision": %d, "wstrNewDescription": "%s"}`,
		nObjId,
		nObjType,
		nObjRevision, wstrNewDescription))
	request, err := http.NewRequest("POST", ea.client.Server+"/api/v1.0/ExtAud.UpdateRevisionDesc", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ea.client.Do(ctx, request, nil)
	return raw, err
}

type FinalDeleteParams struct {
	ArrObjects []ArrObject `json:"arrObjects"`
}

type ArrObject struct {
	Type  string   `json:"type,omitempty"`
	Value *FDValue `json:"value,omitempty"`
}

type FDValue struct {
	ObjID   int64 `json:"nObjId,omitempty"`
	ObjType int64 `json:"nObjType,omitempty"`
}

//	Final delete for deleted objects.
//
//	Parameters:
//	arrObjects	[interface{}] (array) Array of pairs ObjId-ObjType.
//	Max size of array is 100 elements.
func (ea *ExtAud) FinalDelete(ctx context.Context, params FinalDeleteParams) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ea.client.Server+"/api/v1.0/ExtAud.FinalDelete", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ea.client.Do(ctx, request, nil)
	return raw, err
}
