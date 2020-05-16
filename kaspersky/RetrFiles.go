/*
 *
 * 	Copyright (C) 2020  <Semchenko Aleksandr>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.If not, see <http://www.gnu.org/licenses/>.
 * /
 */

package kaspersky

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

//	RetrFiles Class Reference
//	Class provides means to get retranslated files info. More...
//
//	List of all members.
type RetrFiles service

//FilesRequest struct using in GetInfo
type FilesRequest struct {
	FilesRequestElement []FilesRequestElement `json:"aRequest"`
}

type FilesRequestElement struct {
	Type              *string            `json:"type,omitempty"`
	FilesRequestValue *FilesRequestValue `json:"value,omitempty"`
}

type FilesRequestValue struct {
	Index    *string `json:"Index,omitempty"`
	CompID   *string `json:"CompId,omitempty"`
	FileName *string `json:"FileName,omitempty"`
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
