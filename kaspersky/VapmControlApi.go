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

//VapmControlApi.AcceptEulas
//Accepts given EULAs.
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

//VapmControlApi.CancelDeleteFilesForUpdates
//Cancel the files cleanup process initiated by DeleteFilesForUpdates() call.
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

//VapmControlApi.CancelDownloadPatch
//Cancel the patch downloading started by DownloadPatchAsync().
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

//VapmControlApi.ChangeApproval
//Changes updates approval.
//
//	Parameters:
//	- pUpdates	(array) array of update ids to be approved/declined; each entry is paramParams containing following
//	attributes:
//		|-'nSource' - Type of update, see Software updates source enum
//		|- 'nPatchDbId' - Update db id (Equals to 'nKlUpdateDbId',
//		'nRevisionID' or 'nPatchDbId' of v_vapm_update depending on 'nSource' value)
//	- nApprovementState	(int64) new approval state for the given updates; see Update approvement state enum
func (vca *VapmControlApi) ChangeApproval(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.ChangeApproval", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//VapmControlApi.ChangeVulnerabilityIgnorance
//Changes "ignore" state of a vulnerability.
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

//VapmControlApi.DeclineEulas
//Decline given EULAs.
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

//VapmControlApi.DeleteFilesForUpdates
//Cleanup all the files in all the server storages containing the bodies of the given patches.
//The operation progress is reported by 'KLEV_EventAsyncState' events.
//
//	Parameters:
//	- pUpdatesIds	(array) updates identities array; each entry is paramParams containing one of the following
//	attributes: 'nPatchDbId' or 'nRevisionID'
//
//	Return:
//	- wstrRequestId	(string) request ID, used to cancel the request by CancelDeleteFilesForUpdates(
//	) and to subscribe for the 'KLEV_EventAsyncState' events
func (vca *VapmControlApi) DeleteFilesForUpdates(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.DeleteFilesForUpdates", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//VapmControlApi.DownloadPatchAsync
//Download 3-party patch to save locally.
//
//	Parameters:
//	- llPatchGlbId	(int64) - patch database ID
//	- nLcid			(int64) - patch Lcid
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

//VapmControlApi.GetAttributesSetVersionNum
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

//VapmControlApi.GetDownloadPatchDataChunk
//Get the downloaded patch body chunk.
//
//	Parameters:
//	- wstrRequestId	(string) request ID used to call DownloadPatchAsync().
//	- nStartPos		(int64) requested chunk start position
//	- nSizeMax		(int64) maximum chunk size
//
//	Returns:
//	- (binary) data chunk
func (vca *VapmControlApi) GetDownloadPatchDataChunk(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetDownloadPatchDataChunk", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//VapmControlApi.GetDownloadPatchResult
//Get the information on the patch download result.
//
//	Parameters:
//	- wstrRequestId	(string) request ID used to call DownloadPatchAsync().
//
//	Return:
//	- wstrFileName	(string) patch file name
//	- nSize			(int64) patch file size
func (vca *VapmControlApi) GetDownloadPatchResult(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetDownloadPatchResult", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//PEULAParams struct using in VapmControlApi.GetEulaParams
type PEULAParams struct {
	PEULAParam *PEULAParam `json:"pEulaParams,omitempty"`
}

type PEULAParam struct {
	StrEULAURL string `json:"strEULAUrl,omitempty"`
	StrEULA    string `json:"strEULA,omitempty"`
}

//VapmControlApi.GetEulaParams
//Requests EULA params.
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

//VapmControlApi.GetEulasIdsForPatchPrerequisites
//Requests the set of EULA ids for the distributives/patches which are required to install the given patch.
//
//	Parameters:
//	- llPatchGlobalId	(int64) VAPM patch global identity ('nPatchGlbId' update attribute)
//	- nLCID				(int64) LCID of the patch
//
//	Return:
//	- pEulasIds			(array) vector of EULA ids
func (vca *VapmControlApi) GetEulasIdsForPatchPrerequisites(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetEulasIdsForPatchPrerequisites", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//VapmControlApi.GetEulasIdsForUpdates
//Requests the set of EULA ids for the given set of updates.
//
//	Parameters:
//	- pUpdates	(array) array of update ids to be approved/declined; each entry is paramParams containing following	attributes:
//		|- 'nSource' - Type of update, see Software updates source enum
//		|- 'nPatchDbId' - Update db id (Equals to 'nKlUpdateDbId', 'nRevisionID' or 'nPatchDbId' of v_vapm_update depending on 'nSource' value)
//	- nLcid	(int64) - preferred LCID
//
//	Return:
//	- pEulaIds	[out] (array) vector of EULA ids
func (vca *VapmControlApi) GetEulasIdsForUpdates(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetEulasIdsForUpdates", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//VapmControlApi.GetEulasIdsForVulnerabilitiesPatches
//Requests set of EULA ids for the given set of vulnerabilities.
//
//	Parameters:
//	- pVulnerabilities	(array) Vector of integer vulnerabilities ids.
//	- nLCID				(int64) Preferred LCID.
//	Returns:
//	- (array) Vector of EULA ids.
func (vca *VapmControlApi) GetEulasIdsForVulnerabilitiesPatches(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetEulasIdsForVulnerabilitiesPatches", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//VapmControlApi.GetEulasInfo
//Requests the set of EULA descriptors for the given set of updates.
//
//	Parameters:
//	- pUpdates	(array) array of update ids to be approved/declined; each entry is paramParams containing following
//	attributes:
//		|-'nSource' - Type of update, see Software updates source enum
//		|- 'nPatchDbId' - Update db id (Equals to 'nKlUpdateDbId',
//		'nRevisionID' or 'nPatchDbId' of v_vapm_update depending on 'nSource' value)
//	- nLcid	(int64) preferred LCID
//
//	Return:
//	- pEulasInfo	(array) array of EULA params: each entry is paramParams and might contain 'strEULAUrl' (
//	EULA file URL) or 'strEULA' (EULA text); 'strEULAUrl' usage is preffered.
func (vca *VapmControlApi) GetEulasInfo(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetEulasInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//VapmControlApi.GetPendingRulesTasks
//Get identities of VAPM tasks which rules are still being processed.
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

//VapmControlApi.GetSupportedLcidsForPatchPrerequisites
//Gets all LCIDs supported by distributives/patches which are required to install the given patch.
//
//	Parameters:
//	- llPatchGlobalId	(int64) VAPM patch global identity ('nPatchGlbId' update attribute).
//	- nOriginalLcid		(int64) LCID of the original patch
//
//	Return:
//	- pLcids	(array) array of Lcids
func (vca *VapmControlApi) GetSupportedLcidsForPatchPrerequisites(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetSupportedLcidsForPatchPrerequisites", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//VapmControlApi.GetUpdateSupportedLanguagesFilter
//Get filter of supported languages for software update.
//
//Parameters:
//	- nUpdateSource			(int64) Update source type KLVAPM::UpdateSource
//	- pSupportedLanguages	(array) Sorted array of supported languages id
//Use languages from pSupportedLanguages only if pSupportedLanguages not empty. Otherwise use all known languages (means empty filter)
func (vca *VapmControlApi) GetUpdateSupportedLanguagesFilter(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetUpdateSupportedLanguagesFilter", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//VapmControlApi.InitiateDownload
//Check if any updates have KLVAPM::DS_NEED_DOWNLOAD state, and start the download process if needed.
func (vca *VapmControlApi) InitiateDownload(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.InitiateDownload",
		nil)
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

//VapmControlApi.SetPackagesToFixVulnerability
//Set custom packages as patches for a vulnerability.
//
//	Parameters:
//	- wstrVulnerabilityUid	(string) Unique id of the vulnerability.
//	- pPackages				(array) Package ids to be used as patches (in the order of presence in the order).
//	- pParams				(params) Additional parameters describing the custom packages assigned to fix
//	the vulnerability; currently supported:
//	nVulnPatchPkgLCID (contains target installation LCID; 0 if suitable for any language of the product to be patched).
func (vca *VapmControlApi) SetPackagesToFixVulnerability(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.SetPackagesToFixVulnerability", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}
