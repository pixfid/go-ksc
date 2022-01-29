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

// GroupSync service for access to group synchronization objects.
type GroupSync service

// NSyncInfoParams struct using in GroupSync.GetSyncHostsInfo
type NSyncInfoParams struct {
	// NSync id of the group synchronization. Can be retrieved from policy attribute KLPOL_GSYN_ID
	NSync int64 `json:"nSync,omitempty"`

	// ArrFieldsToReturn array of attribute names to return.
	// See List of group synchronization host attributes for attribute names
	ArrFieldsToReturn []string `json:"arrFieldsToReturn"`

	// ArrFieldsToOrder array of containers each of them containing two attributes :
	// "Name" (string) name of List of group synchronization host attributes used for sorting
	// "Asc" (bool) ascending if true descending otherwise
	ArrFieldsToOrder []FieldsToOrder `json:"arrFieldsToOrder,omitempty"`

	// NLifeTime timeout in seconds to keep the result-set alive, zero means 'default value'
	NLifeTime int64 `json:"nLifeTime,omitempty"`
}

// GetSyncHostsInfo Acquire group synchronization state at target hosts.
// Returns forward iterator to access requested properties of the specified group synchronization at target hosts.
func (gs *GroupSync) GetSyncHostsInfo(ctx context.Context, params NSyncInfoParams) (*PxgValStr, error) {
	postData, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", gs.client.Server+"/api/v1.0/GroupSync.GetSyncHostsInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgValStr := new(PxgValStr)
	_, err = gs.client.Request(ctx, request, &pxgValStr)
	return pxgValStr, err
}

// GroupSyncInfoParams struct
type GroupSyncInfoParams struct {
	NSync             int64    `json:"nSync,omitempty"`
	ArrFieldsToReturn []string `json:"arrFieldsToReturn"`
}

// GroupSyncInfo struct
type GroupSyncInfo struct {
	SyncInfo *SyncInfo `json:"PxgRetVal,omitempty"`
}

type SyncInfo struct {
	GsynCntFailed    int64 `json:"gsyn_cnt_Failed,omitempty"`
	GsynCntFinished  int64 `json:"gsyn_cnt_Finished,omitempty"`
	GsynCntRunning   int64 `json:"gsyn_cnt_Running,omitempty"`
	GsynCntScheduled int64 `json:"gsyn_cnt_Scheduled,omitempty"`
	GsynCntFullCount int64 `json:"gsyn_cnt_FullCount,omitempty"`
}

// GetSyncInfo Acquire group synchronization properties.
//
// Returns requested properties of the specified group synchronization
func (gs *GroupSync) GetSyncInfo(ctx context.Context, params GroupSyncInfoParams) (*GroupSyncInfo, error) {
	postData, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", gs.client.Server+"/api/v1.0/GroupSync.GetSyncInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	groupSyncInfo := new(GroupSyncInfo)
	_, err = gs.client.Request(ctx, request, &groupSyncInfo)
	return groupSyncInfo, err
}

// GetSyncDeliveryTime Acquire group synchronization delivery time for the specified host.
//
// Returns UTC time when the specified synchronization has been delivered to the specified host
func (gs *GroupSync) GetSyncDeliveryTime(ctx context.Context, nSync int64, szwHostId string) (*PxgValInt,
	[]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nSync": %d, "szwHostId": "%s"}`, nSync, szwHostId))
	request, err := http.NewRequest("POST", gs.client.Server+"/api/v1.0/GroupSync.GetSyncDeliveryTime", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := gs.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}
