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

//	GroupSync Class Reference
//	Access to group synchronization objects. More...
//
//	List of all members.
type GroupSync service

//TODO GetSyncDeliveryTime
//TODO GetSyncHostsInfo

//GroupSyncInfoParams struct
type GroupSyncInfoParams struct {
	NSync             *int64   `json:"nSync,omitempty"`
	ArrFieldsToReturn []string `json:"arrFieldsToReturn"`
}

//GroupSyncInfo struct
type GroupSyncInfo struct {
	SyncInfo *SyncInfo `json:"PxgRetVal,omitempty"`
}

type SyncInfo struct {
	GsynCntFailed    *int64 `json:"gsyn_cnt_Failed,omitempty"`
	GsynCntFinished  *int64 `json:"gsyn_cnt_Finished,omitempty"`
	GsynCntRunning   *int64 `json:"gsyn_cnt_Running,omitempty"`
	GsynCntScheduled *int64 `json:"gsyn_cnt_Scheduled,omitempty"`
	GsynCntFullCount *int64 `json:"gsyn_cnt_FullCount,omitempty"`
}

//	Acquire group synchronization properties.
//
//	Returns requested properties of the specified group synchronization
//
//	- Parameters:
//	- (params) GroupSyncInfoParams container with values
//		|- nSync	(int) id of the group synchronization. Can be retrieved from policy attribute KLPOL_GSYN_ID
//		|- arrFieldsToReturn	(array) array of attribute names to return.
//	See List of group synchronization attributes for attribute names
//
//	Returns:
//	- (params) container with values of required attributes
func (gs *GroupSync) GetSyncInfo(ctx context.Context, params GroupSyncInfoParams) (*GroupSyncInfo, []byte, error) {
	postData, err := json.Marshal(&params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", gs.client.Server+"/api/v1.0/GroupSync.GetSyncInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	groupSyncInfo := new(GroupSyncInfo)
	raw, err := gs.client.Do(ctx, request, &groupSyncInfo)
	return groupSyncInfo, raw, err
}
