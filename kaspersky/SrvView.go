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

// SrvView service to get plain-queries from SC-server.
type SrvView service

// SrvViewParams struct
type SrvViewParams struct {
	// WstrViewName name of srvview see List of supported srvviews.
	WstrViewName string `json:"wstrViewName"`

	// WstrFilter filter string, contains a condition over srvview attributes, see also Search filter syntax.
	WstrFilter string `json:"wstrFilter"`

	// VecFieldsToReturn array of srvview attribute names to return.
	VecFieldsToReturn []string `json:"vecFieldsToReturn"`

	// VecFieldsToOrder array of containers each of them containing two attributes:
	//	-	"Name" (paramString) name of attribute used for sorting
	//	-	"Asc" (paramBool) ascending if true descending otherwise
	VecFieldsToOrder []FieldsToOrder `json:"vecFieldsToOrder"`

	// PParams extra options. This parameter can contain additional options to configure query.
	//Some options are specific to the wstrViewName and are part of it specification.
	//And some options are common for all srvviews. List of common options:
	//"TOP_N" (int64) acquire only first N records
	PParams *ESrvViewParams `json:"pParams"`

	// LifetimeSEC max result-set lifetime in seconds
	LifetimeSEC int64 `json:"lifetimeSec"`
}

// ESrvViewParams struct
type ESrvViewParams struct {
	//acquire only first N records
	TopN int64 `json:"TOP_N,omitempty"`
}

// ResetIterator Find srvview data by filter string.
//
// Finds data records for srvview wstrViewName that satisfy conditions from filter string wstrFilter,
// and creates a server-side collection of found data records.
//
//	Parameters:
//	Example:
//	{
//		"wstrViewName":"HWInvStorageSrvViewName",
//		"wstrFilter":"(&(MotherBoard=\"*\"))",
//		"vecFieldsToReturn":["Id","Type","CPU"],
//		"vecFieldsToOrder":null,
//		"pParams":null,
//		"lifetimeSec":7200
//	}
// The result-set is destroyed and associated memory is freed in following cases:
// Passed lifetimeSec seconds after last access to the result-set (by methods GetRecordCount and GetRecordRange).
// Session to the Administration Server has been closed.
// ReleaseIterator has been called.
func (sv *SrvView) ResetIterator(ctx context.Context, params *SrvViewParams) (*WstrIteratorID, []byte, error) {
	postData, err := json.Marshal(&params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", sv.client.Server+"/api/v1.0/SrvView.ResetIterator", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	srvViewIter := new(WstrIteratorID)
	raw, err := sv.client.Request(ctx, request, &srvViewIter)
	return srvViewIter, raw, err
}

// GetRecordCount Acquire count of result-set elements.
//
// Returns number of elements contained in the specified result-set.
func (sv *SrvView) GetRecordCount(ctx context.Context, wstrIteratorId string) (*PxgValInt, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrIteratorId":"%s"}`, wstrIteratorId))
	request, err := http.NewRequest("POST", sv.client.Server+"/api/v1.0/SrvView.GetRecordCount", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := sv.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// ReleaseIterator Releases the specified result-set and frees associated memory
func (sv *SrvView) ReleaseIterator(ctx context.Context, wstrIteratorId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrIteratorId":"%s"}`, wstrIteratorId))
	request, err := http.NewRequest("POST", sv.client.Server+"/api/v1.0/SrvView.ReleaseIterator", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sv.client.Request(ctx, request, nil)
	return raw, err
}

// RecordRangeParams struct
type RecordRangeParams struct {
	WstrIteratorID string `json:"wstrIteratorId,omitempty"`
	NStart         int64  `json:"nStart,omitempty"`
	NEnd           int64  `json:"nEnd,omitempty"`
}

// GetRecordRange Acquire subset of result-set elements by range.
//
// Returns elements contained in the specified result-set in the diapason from position nStart to position nEnd.
func (sv *SrvView) GetRecordRange(ctx context.Context, params *RecordRangeParams, out interface{}) ([]byte, error) {
	postData, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sv.client.Server+"/api/v1.0/SrvView.GetRecordRange", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sv.client.Request(ctx, request, &out)
	return raw, err
}
