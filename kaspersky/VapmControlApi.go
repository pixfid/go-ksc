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

//	VapmControlApi Class Reference
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
	return pxgValInt, raw, nil
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
	return raw, nil
}

func (sc *VapmControlApi) InitiateDownload(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/VapmControlApi.InitiateDownload",
		nil)
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Do(ctx, request, nil)
	return raw, nil
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
