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

// ChunkAccessor service working with host result-set, that is a server-side ordered collection of found hosts.
type ChunkAccessor service

// Release result-set. Releases the specified result-set and frees associated memory
func (ca *ChunkAccessor) Release(ctx context.Context, accessor string) bool {
	postData := []byte(fmt.Sprintf(`{"strAccessor": "%s"}`, accessor))
	request, err := http.NewRequest("POST", ca.client.Server+"/api/v1.0/ChunkAccessor.Release", bytes.NewBuffer(postData))
	if err != nil {
		return false
	}

	raw, err := ca.client.Request(ctx, request, nil)
	if raw != nil {
		return true
	}
	return false
}

// GetItemsCount Acquire count of result-set elements. Returns number of elements contained in the specified result-set.
func (ca *ChunkAccessor) GetItemsCount(ctx context.Context, accessor string) (*PxgValInt, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strAccessor": "%s"}`, accessor))
	request, err := http.NewRequest("POST", ca.client.Server+"/api/v1.0/ChunkAccessor.GetItemsCount", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := ca.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

type ItemsChunkParams struct {
	StrAccessor string `json:"strAccessor,omitempty"`
	NStart      int64  `json:"nStart,omitempty"`
	NCount      int64  `json:"nCount,omitempty"`
}

// GetItemsChunk Acquire subset of result-set elements by range.
// Returns specified nCount elements contained in the specified result-set beginning from position nStart.
func (ca *ChunkAccessor) GetItemsChunk(ctx context.Context, params ItemsChunkParams, result interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", ca.client.Server+"/api/v1.0/ChunkAccessor.GetItemsChunk", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := ca.client.Request(ctx, request, &result)
	return raw, err
}
