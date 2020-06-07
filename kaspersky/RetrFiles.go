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

//	RetrFiles Class Reference
//
//	Class provides means to get retranslated files info..
//
//	List of all members.
type RetrFiles service

//FilesRequest struct using in GetInfo
type FilesRequest struct {
	//Array of params, each cell (paramParams) contains request-info for one updatable file:
	FilesRequestElement []FilesRequestElement `json:"aRequest"`
}

type FilesRequestElement struct {
	Type              string             `json:"type,omitempty"`
	FilesRequestValue *FilesRequestValue `json:"value,omitempty"`
}

type FilesRequestValue struct {
	//primary index relative path in lowercase, e.g. "index/u1313g.xml";
	Index string `json:"Index,omitempty"`

	//updatable file component id in UPPERCASE, e.g. "KSC";
	CompID string `json:"CompId,omitempty"`

	//file name without path in lowercase, e.g. "kscdat.zip".
	FileName string `json:"FileName,omitempty"`
}

//	Synchronously requests information about some retranslated files.
//
//	Parameters:
//	- aRequest	paramArray(paramParams) Array of params,
//	each cell (paramParams) contains request-info for one updatable file:
//	|- KLUPD_RecentIndex ("Index"): primary index relative path in lowercase, e.g. "index/u1313g.xml";
//	|- KLUPD_RecentCompId ("CompId"): updatable file component id in UPPERCASE, e.g. "KSC";
//	|- KLUPD_RecentFileName ("FileName"): file name without path in lowercase, e.g. "kscdat.zip".
//
//	Returns:
//	- paramArray(paramArray(paramParams)) Array of found files info, correspondingly to incoming request-array,
//	cell-by-cell. Each ouput cell is paramArray(paramParams), that is list of matched files data, where every file-info is params:
//	KLUPD_RecentRelativeSrvPath ("RelativeSrvPath"): file's relative path inside retranslation folder, e.g. "updates/ksc/".
//	If nothing is found for given in-cell, then corresponding out-cell is NULL.
//	Null/empty in-params results in null out-array.
func (rf *RetrFiles) GetInfo(ctx context.Context, params FilesRequest) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", rf.client.Server+"/api/v1.0/RetrFiles.GetInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rf.client.Do(ctx, request, nil)
	return raw, err
}
