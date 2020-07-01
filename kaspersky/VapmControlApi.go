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

// VapmControlApi VAPM.
type VapmControlApi service

// PEulaIDParams struct using in VapmControlApi.AcceptEulas
type PEulaIDParams struct {
	PEulaIDs []int64 `json:"pEulaIDs"`
}

// AcceptEulas
// Accepts given EULAs.
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

// CancelDeleteFilesForUpdates
// Cancel the files cleanup process initiated by DeleteFilesForUpdates() call.
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

// CancelDownloadPatch
// Cancel the patch downloading started by DownloadPatchAsync().
func (vca *VapmControlApi) CancelDownloadPatch(ctx context.Context, wstrRequestId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrRequestId": "%s"}`, wstrRequestId))
	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.CancelDownloadPatch", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

type EulasIDSForUpdates struct {
	PUpdates []EulasIDSPUpdate `json:"pUpdates"`
	NLcid    int64             `json:"nLcid,omitempty"`
}

type EulasIDSPUpdate struct {
	Type  string        `json:"type,omitempty"`
	Value EulasIDSValue `json:"value,omitempty"`
}

type EulasIDSValue struct {
	NSource    int64 `json:"nSource,omitempty"`
	NPatchDBID int64 `json:"nPatchDbId,omitempty"`
}

// ChangeApproval
// Changes updates approval.
func (vca *VapmControlApi) ChangeApproval(ctx context.Context, params EulasIDSForUpdates) ([]byte, error) {
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

// ChangeVulnerabilityIgnorance
// Changes "ignore" state of a vulnerability.
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

// DeclineEulas
// Decline given EULAs.
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

type DeleteFilesForUpdatesParams struct {
	PUpdatesIDS   []PUpdatesID `json:"pUpdatesIds"`
	WstrRequestID string       `json:"wstrRequestId,omitempty"`
}

type PUpdatesID struct {
	Type  string        `json:"type,omitempty"`
	Value PUpdatesValue `json:"value,omitempty"`
}

type PUpdatesValue struct {
	NPatchDBID  int64 `json:"nPatchDbId,omitempty"`
	NRevisionID int64 `json:"nRevisionID,omitempty"`
}

// DeleteFilesForUpdates
// Cleanup all the files in all the server storages containing the bodies of the given patches.
func (vca *VapmControlApi) DeleteFilesForUpdates(ctx context.Context, params DeleteFilesForUpdatesParams) ([]byte, error) {
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

// DownloadPatchAsync
// Download 3-party patch to save locally.
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

// GetAttributesSetVersionNum
// Returns edition of supported attributes, EAttributesSetVersion.
func (vca *VapmControlApi) GetAttributesSetVersionNum(ctx context.Context) (*PxgValInt, error) {
	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetAttributesSetVersionNum", nil)
	if err != nil {
		return nil, err
	}

	pxgValInt := new(PxgValInt)
	_, err = vca.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, err
}

// GetDownloadPatchDataChunk
// Get the downloaded patch body chunk.
func (vca *VapmControlApi) GetDownloadPatchDataChunk(ctx context.Context, wstrRequestId string, nStartPos, nSizeMax int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrRequestId": "%s", "nStartPos": %d, "nSizeMax": %d}`, wstrRequestId, nStartPos, nSizeMax))

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetDownloadPatchDataChunk", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

// GetDownloadPatchResult
// Get the information on the patch download result.
func (vca *VapmControlApi) GetDownloadPatchResult(ctx context.Context, wstrRequestId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrRequestId": "%s"}`, wstrRequestId))

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetDownloadPatchResult", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

// PEULAParams struct using in VapmControlApi.GetEulaParams
type PEULAParams struct {
	PEULAParam *PEULAParam `json:"pEulaParams,omitempty"`
}

// PEULAParam struct
type PEULAParam struct {
	StrEULAURL string `json:"strEULAUrl,omitempty"`
	StrEULA    string `json:"strEULA,omitempty"`
}

// GetEulaParams Requests EULA params.
func (vca *VapmControlApi) GetEulaParams(ctx context.Context, nEulaId int64) (*PEULAParams, error) {
	postData := []byte(fmt.Sprintf(`{"nEulaId": %d}`, nEulaId))
	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetEulaParams", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pEulaParams := new(PEULAParams)
	_, err = vca.client.Do(ctx, request, &pEulaParams)
	return pEulaParams, err
}

// EulasIDSForPatchPrerequisitesParams struct
type EulasIDSForPatchPrerequisitesParams struct {
	// LlPatchGlobalID VAPM patch global identity ('nPatchGlbId' update attribute)
	LlPatchGlobalID int64 `json:"llPatchGlobalId,omitempty"`
	// NLCID LCID of the patch
	NLCID int64 `json:"nLCID,omitempty"`
}

// EulasIDS struct
type EulasIDS struct {
	PEulasIDS []int64 `json:"pEulasIds"`
}

// GetEulasIdsForPatchPrerequisites
// Requests the set of EULA ids for the distributives/patches which are required to install the given patch.
func (vca *VapmControlApi) GetEulasIdsForPatchPrerequisites(ctx context.Context, params EulasIDSForPatchPrerequisitesParams) (*EulasIDS, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetEulasIdsForPatchPrerequisites", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	eulasIDSForPatchPrerequisites := new(EulasIDS)
	_, err = vca.client.Do(ctx, request, &eulasIDSForPatchPrerequisites)
	return eulasIDSForPatchPrerequisites, err
}

// ApprovalParams struct
type EulasIdsForUpdatesParams struct {
	PUpdates []EulasIdsPUpdate `json:"pUpdates"`
	NLcid    int64             `json:"nLcid,omitempty"`
}

type EulasIdsPUpdate struct {
	Type  string        `json:"type,omitempty"`
	Value ApprovalValue `json:"value,omitempty"`
}

type ApprovalValue struct {
	NSource    int64 `json:"nSource,omitempty"`
	NPatchDbId int64 `json:"nPatchDbId,omitempty"`
}

type EulasID struct {
	PEulasID []int64 `json:"pEulaIds"`
}

// GetEulasIdsForUpdates
// Requests the set of EULA ids for the given set of updates.
func (vca *VapmControlApi) GetEulasIdsForUpdates(ctx context.Context, params EulasIdsForUpdatesParams) (*EulasID, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetEulasIdsForUpdates", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}
	eulasID := new(EulasID)
	_, err = vca.client.Do(ctx, request, &eulasID)
	return eulasID, err
}

type EulasIDSForVulnerabilitiesPatchesParams struct {
	PVulnerabilities []int64 `json:"pVulnerabilities"`
	NLCID            int64   `json:"nLCID,omitempty"`
}

// GetEulasIdsForVulnerabilitiesPatches Requests set of EULA ids for the given set of vulnerabilities.
// Remark: not implemented
func (vca *VapmControlApi) GetEulasIdsForVulnerabilitiesPatches(ctx context.Context, params EulasIDSForVulnerabilitiesPatchesParams) ([]byte, error) {
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

// PUpdates
type PUpdates struct {
	// PUpdates array of PUpdate to be approved/declined
	PUpdates []PUpdate `json:"pUpdates"`

	// NLcid preferred LCID
	NLcid int64 `json:"nLcid"`
}

type PUpdate struct {
	Type  string       `json:"type,omitempty"`
	Value PUpdateValue `json:"value,omitempty"`
}

type PUpdateValue struct {
	// NSource Type of update
	NSource int64 `json:"nSource,omitempty"`

	//NPatchDBID Update db id (Equals to 'nKlUpdateDbId', 'nRevisionID' or 'nPatchDbId' of v_vapm_update depending on 'nSource' value
	NPatchDBID int64 `json:"nPatchDbId,omitempty"`
}

type PEulasInfo struct {
	// PEulasInfo array of EULA params
	PEulasInfo []PEulasInfoElement `json:"pEulasInfo"`
}

type PEulasInfoElement struct {
	Type  *string          `json:"type,omitempty"`
	Value *PEulasInfoValue `json:"value,omitempty"`
}

// PEulasInfoValue
type PEulasInfoValue struct {
	// NEulaDbId EULA id
	NEulaDbId *int64 `json:"nEulaDbId,omitempty"`

	// StrEULAUrl EULA file URL
	StrEULAUrl *string `json:"strEULAUrl,omitempty"`

	// StrEULA EULA text
	StrEULA *string `json:"strEULA,omitempty"`
}

// GetEulasInfo Requests the set of EULA descriptors for the given set of updates.
func (vca *VapmControlApi) GetEulasInfo(ctx context.Context, params PUpdates) (*PEulasInfo, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetEulasInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pEulasInfo := new(PEulasInfo)
	_, err = vca.client.Do(ctx, request, &pEulasInfo)
	return pEulasInfo, err
}

// PendingRulesTasks contains task ids array
type PendingRulesTasks struct {
	// PTasksIDS Array of task ids
	PTasksIDS []int64 `json:"pTasksIds"`
}

// GetPendingRulesTasks Get identities of VAPM tasks which rules are still being processed.
func (vca *VapmControlApi) GetPendingRulesTasks(ctx context.Context) (*PendingRulesTasks, error) {
	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetPendingRulesTasks", nil)
	if err != nil {
		return nil, err
	}

	pendingRulesTasks := new(PendingRulesTasks)
	_, err = vca.client.Do(ctx, request, &pendingRulesTasks)
	return pendingRulesTasks, err
}

// SupportedLcids contains Lcids ids array
type SupportedLcids struct {
	// PLcids array of Lcids
	PLcids []int64 `json:"pLcids"`
}

// GetSupportedLcidsForPatchPrerequisites Gets all LCIDs supported by distributives/patches which are required to install the given patch.
func (vca *VapmControlApi) GetSupportedLcidsForPatchPrerequisites(ctx context.Context, llPatchGlobalId, nOriginalLcid int64) (*SupportedLcids, error) {
	postData := []byte(fmt.Sprintf(`{"llPatchGlobalId": %d, "nOriginalLcid": %d}`, llPatchGlobalId, nOriginalLcid))

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetSupportedLcidsForPatchPrerequisites", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	supportedLcids := new(SupportedLcids)
	_, err = vca.client.Do(ctx, request, *supportedLcids)
	return supportedLcids, err
}

// PSupportedLanguages supported languages for software update
type PSupportedLanguages struct {
	// PSupportedLanguages Sorted array of supported languages id
	PSupportedLanguages []int64 `json:"pSupportedLanguages"`
}

// GetUpdateSupportedLanguagesFilter Get filter of supported languages for software update.
func (vca *VapmControlApi) GetUpdateSupportedLanguagesFilter(ctx context.Context, nUpdateSource int64) (*PSupportedLanguages, error) {
	postData := []byte(fmt.Sprintf(`{"nUpdateSource": %d}`, nUpdateSource))

	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.GetUpdateSupportedLanguagesFilter", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pSupportedLanguages := new(PSupportedLanguages)
	_, err = vca.client.Do(ctx, request, &pSupportedLanguages)
	return pSupportedLanguages, err
}

// InitiateDownload Check if any updates have KLVAPM::DS_NEED_DOWNLOAD state, and start the download process if needed.
func (vca *VapmControlApi) InitiateDownload(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", vca.client.Server+"/api/v1.0/VapmControlApi.InitiateDownload",
		nil)
	if err != nil {
		return nil, err
	}

	raw, err := vca.client.Do(ctx, request, nil)
	return raw, err
}

// SetPackagesToFixVulnerability Set custom packages as patches for a vulnerability.
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
