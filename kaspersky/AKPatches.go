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
	"log"
	"net/http"
)

//	Interface to manage system of autoupdating by patch.exe patches.
//
//	List of all members.
type AKPatches struct {
	client *Client
}

//PatchParams struct
type PatchParams struct {
	SzwPatchID  string       `json:"szwPatchId"`
	PatchOption *PatchOption `json:"parOptions"`
}

type PatchOption struct {
	Type  string     `json:"type"`
	Value PatchValue `json:"value"`
}

type PatchValue struct {
	NeedBackupServer bool `json:"NeedBackupServer"`
}

//	Give approval to the patch on start installation.
//
//	This command is applicable to patches which have no right
//	to be installed without approval of the administrator.
//
//	Parameters:
//	- szwPatchId	(string). the patch identifier. see KLSTS_Par_PatchId.
//	- parOptions	(params) object containing additional parameters:
//
//	Example params request:
//	{
//	  "szwPatchId": "patchId",
//	  "parOptions": {
//	    "type": "params",
//	    "value": {
//	      "NeedBackupServer": true
//	    }
//	  }
//	}
//	|-"NeedBackupServer" (boolean) true if is required to do SC-server backup before installation of the patch,
//	optional
//
//	Returns:
//	- ppResult (params) contains following attributes:
//	- "CallResult" (boolean) true if success
//	- "ErrorDescription" (string) error description, empty if no error
func (akp *AKPatches) ApprovePatch(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)

	request, err := http.NewRequest("POST", akp.client.Server+"/api/v1.0/AKPatches.ApprovePatch",
		bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := akp.client.Do(ctx, request, nil)

	return raw, err
}

//	Forbid installation of the patch.
//
//	After this forbidden will be no more notifications that this patch expecting approval on installation
//
//	Parameters:
//	- szwPatchId	(string). the patch identifier. see KLSTS_Par_PatchId.
//	- parOptions	reserved. (params)
//	Example params request:
//	{
//	  "szwPatchId": "patchId",
//	  "parOptions": nil
//	}
//
//	Returns:
//	- ppResult (params) contains following attributes:
//	- "CallResult" (boolean) true if success
//	- "ErrorDescription" (string) error description, empty if no error
func (akp *AKPatches) ForbidPatch(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)

	request, err := http.NewRequest("POST", akp.client.Server+"/api/v1.0/AKPatches.ForbidPatch",
		bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := akp.client.Do(ctx, request, nil)

	return raw, err
}
