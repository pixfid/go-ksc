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

//	VapmControlApi Class Reference
//
//	VAPM.
//
//	List of all members.
type VapmControlApi service

//	Returns edition of supported attributes, KLVAPM::EAttributesSetVersion.
//
//	Returns:
//	- KLVAPM::EAttributesSetVersion value
func (sc *VapmControlApi) GetAttributesSetVersionNum(ctx context.Context, params Null) (*PxgValInt, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/VapmControlApi.GetAttributesSetVersionNum",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := sc.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//	Get identities of VAPM tasks which rules are still being processed.
//
//	Parameters:
//	- pTasksIds	[out] (array) Array of task ids
//	These tasks rules are still being processed and the information
//	on them could not be actual within appropriate SrvView results/
func (sc *VapmControlApi) GetPendingRulesTasks(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/VapmControlApi.GetPendingRulesTasks",
		nil)
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Do(ctx, request, nil)
	return raw, err
}

func (sc *VapmControlApi) InitiateDownload(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/VapmControlApi.InitiateDownload",
		nil)
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Do(ctx, request, nil)
	return raw, err
}

//TODO
/*
AcceptEulas
CancelDeleteFilesForUpdates
CancelDownloadPatch
ChangeApproval
ChangeVulnerabilityIgnorance
DeclineEulas
DeleteFilesForUpdates
DownloadPatchAsync
GetDownloadPatchDataChunk
GetDownloadPatchResult
GetEulaParams
GetEulasIdsForPatchPrerequisites
GetEulasIdsForUpdates
GetEulasIdsForVulnerabilitiesPatches
GetEulasInfo
GetSupportedLcidsForPatchPrerequisites
GetUpdateSupportedLanguagesFilter
SetPackagesToFixVulnerability
*/
