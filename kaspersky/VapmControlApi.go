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

//	VapmControlApi Class Reference
//
//	VAPM.
//
//	List of all members.
type VapmControlApi service

//PEulaIDParams struct using in VapmControlApi.AcceptEulas
type PEulaIDParams struct {
	PEulaIDs []int64 `json:"pEulaIDs"`
}

//	Accepts given EULAs.
//
//	Parameters:
//	- pEulaIDs	(array) set of EULA IDs to accept
func (vca *VapmControlApi) AcceptEulas(ctx context.Context, params PEulaIDParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.AcceptEulas",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//	Cancel the files cleanup process initiated by DeleteFilesForUpdates() call.
//
//	Parameters:
//	- wstrRequestId	(string) request ID, used to initiate the request by the DeleteFilesForUpdates() call
func (vca *VapmControlApi) CancelDeleteFilesForUpdates(ctx context.Context, wstrRequestId string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"wstrRequestId": "%s"}`, wstrRequestId))
	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.CancelDeleteFilesForUpdates", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//	Cancel the patch downloading started by DownloadPatchAsync().
//
//	Parameters:
//	- wstrRequestId	(string) request ID used to call DownloadPatchAsync().
func (vca *VapmControlApi) CancelDownloadPatch(ctx context.Context, wstrRequestId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrRequestId": "%s"}`, wstrRequestId))
	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.CancelDownloadPatch", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//TODO VapmControlApi::ChangeApproval

//	Changes "ignore" state of a vulnerability.
//
//	Parameters:
//	- wstrVulnerabilityUid	(string) Unique id of the vulnerability.
//	- wstrHostId	(string) Host identifier for the vulnerability to be ignored on; could be empty to change the
//	vulnerability state everywhere.
//	- bIgnore	(bool) Should the vulnerability be ignored or not.
func (vca *VapmControlApi) ChangeVulnerabilityIgnorance(ctx context.Context, wstrVulnerabilityUid, wstrHostId string, bIgnore bool) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrVulnerabilityUid": "%s", "wstrHostId": "%s", "bIgnore": %v}`,
		wstrVulnerabilityUid, wstrHostId, bIgnore))
	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.ChangeVulnerabilityIgnorance", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//	Decline given EULAs.
//
//	Parameters:
//	- pEulaIDs	(array) set of EULA IDs to decline
func (vca *VapmControlApi) DeclineEulas(ctx context.Context, params PEulaIDParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.DeclineEulas",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//TODO VapmControlApi::DeleteFilesForUpdates

//	Download 3-party patch to save locally.
//
//	Parameters:
//	- llPatchGlbId	(int64) - patch database ID
//	- nLcid	(int64) - patch Lcid
//	- wstrRequestId	(string) - request ID,
//	used to cancel the request by CancelDownloadPatch()
//	or to get the result by GetDownloadPatchResult().
func (vca *VapmControlApi) DownloadPatchAsync(ctx context.Context, llPatchGlbId, nLcid int64, wstrRequestId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"llPatchGlbId": %d, "nLcid": %d, "wstrRequestId": "%s"}`, llPatchGlbId, nLcid,
		wstrRequestId))
	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.DownloadPatchAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//Returns edition of supported attributes, EAttributesSetVersion.
//
//	Returns:
//	- EAttributesSetVersion value
func (vca *VapmControlApi) GetAttributesSetVersionNum(ctx context.Context) (*PxgValInt, []byte, error) {

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetAttributesSetVersionNum",
		nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := vca.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//TODO VapmControlApi::GetDownloadPatchDataChunk
//TODO VapmControlApi::GetDownloadPatchResult

//PEULAParams struct using in VapmControlApi.GetEulaParams
type PEULAParams struct {
	PEULAParam *PEULAParam `json:"pEulaParams,omitempty"`
}

type PEULAParam struct {
	StrEULAURL string `json:"strEULAUrl,omitempty"`
	StrEULA    string `json:"strEULA,omitempty"`
}

//	Requests EULA params.
//
//	Parameters:
//	- nEulaId	(int) - EULA id
//
//	Return:
//	- pEulaParams	(params) - EULA params, might contain 'strEULAUrl' (EULA file URL) or 'strEULA' (
//	EULA text); 'strEULAUrl' usage is preffered.
func (vca *VapmControlApi) GetEulaParams(ctx context.Context, nEulaId int64) (*PEULAParams, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nEulaId": %d}`, nEulaId))
	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetEulaParams", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pEulaParams := new(PEULAParams)
	raw, err := vca.client.Do(ctx, request, &pEulaParams)
	return pEulaParams, raw, err
}

//TODO VapmControlApi::GetEulasIdsForPatchPrerequisites
//TODO VapmControlApi::GetEulasIdsForUpdates
//TODO VapmControlApi::GetEulasIdsForVulnerabilitiesPatches
//TODO VapmControlApi::GetEulasInfo

//	Get identities of VAPM tasks which rules are still being processed.
//
//	Parameters:
//	- pTasksIds	(array) Array of task ids
//	These tasks rules are still being processed and the information
//	on them could not be actual within appropriate SrvView results/
func (vca *VapmControlApi) GetPendingRulesTasks(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetPendingRulesTasks",
		nil)
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//TODO VapmControlApi::GetSupportedLcidsForPatchPrerequisites
//TODO VapmControlApi::GetUpdateSupportedLanguagesFilter

//	Check if any updates have KLVAPM::DS_NEED_DOWNLOAD state, and start the download process if needed.
func (vca *VapmControlApi) InitiateDownload(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.InitiateDownload",
		nil)
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//TODO VapmControlApi::SetPackagesToFixVulnerability
