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

// EventProcessing service implements the functionality for viewing and deleting events.
type EventProcessing service

// GetRecordCount Get record count in the result-set. Returns number of elements contained in the specified result-set.
func (ep *EventProcessing) GetRecordCount(ctx context.Context, strIteratorId string) (*PxgValInt, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strIteratorId": "%s"}`, strIteratorId))
	request, err := http.NewRequest("POST", ep.client.Server+"/api/v1.0/EventProcessing.GetRecordCount", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := ep.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// GetRecordRange Acquire subset of result-set elements by range.
// Returns elements contained in the specified result-set in the diapason from position nStart to position nEnd.
func (ep *EventProcessing) GetRecordRange(ctx context.Context, strIteratorId string, nStart, nEnd int64) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"strIteratorId": "%s", "nStart": %d, "nEnd": %d}`, strIteratorId, nStart, nEnd))
	request, err := http.NewRequest("POST", ep.client.Server+"/api/v1.0/EventProcessing.GetRecordRange", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ep.client.Request(ctx, request, nil)
	return raw, err
}

// ReleaseIterator Releases the specified result-set and frees associated memory.
func (ep *EventProcessing) ReleaseIterator(ctx context.Context, strIteratorId string) (*PxgValInt, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strIteratorId": "%s"}`, strIteratorId))
	request, err := http.NewRequest("POST", ep.client.Server+"/api/v1.0/EventProcessing.ReleaseIterator", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := ep.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// InitiateDelete Initiates mass delete of the events specified by pSettings in the result-set.
func (ep *EventProcessing) InitiateDelete(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", ep.client.Server+"/api/v1.0/EventProcessing.InitiateDelete",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ep.client.Request(ctx, request, nil)
	return raw, err
}

// CancelDelete Cancels mass delete of the events specified by pSettings in the result-set.
func (ep *EventProcessing) CancelDelete(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", ep.client.Server+"/api/v1.0/EventProcessing.CancelDelete",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ep.client.Request(ctx, request, nil)
	return raw, err
}
