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
	"fmt"
	"net/http"
)

//	ChunkAccessor Class Reference
//
//	Working with host result-set. More...
//
//	Working with a result-set, that is a server-side ordered collection of found hosts.
//
//	List of all members.
type ChunkAccessor service

//Release result-set.
//
//Releases the specified result-set and frees associated memory
//
//Parameters:
//	- strAccessor	(string) result-set ID, identifier of the server-side ordered collection of found hosts
func (ca *ChunkAccessor) Release(ctx context.Context, accessor string) bool {
	postData := []byte(fmt.Sprintf(`{"strAccessor": "%s"}`, accessor))
	request, err := http.NewRequest("POST", ca.client.Server+"/api/v1.0/ChunkAccessor.Release", bytes.NewBuffer(postData))
	if err != nil {
		return false
	}

	raw, err := ca.client.Do(ctx, request, nil)
	if raw != nil {
		return true
	}
	return false
}

//Acquire count of result-set elements.
//
//Returns number of elements contained in the specified result-set.
//
//Parameters:
//	- strAccessor	(string) result-set ID, identifier of the server-side ordered collection of found hosts
//Returns:
//	- (data.PxgValInt) number of elements contained in the specified result-set
func (ca *ChunkAccessor) GetItemsCount(ctx context.Context, accessor string) (*PxgValInt, []byte, error) {
	postData := []byte(fmt.Sprintf(`	{"strAccessor": "%s"}`, accessor))
	request, err := http.NewRequest("POST", ca.client.Server+"/api/v1.0/ChunkAccessor.GetItemsCount", bytes.NewBuffer(postData))
	if err != nil {
		panic(err)
	}

	pxgValInt := new(PxgValInt)
	raw, err := ca.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

type ItemsChunkParams struct {
	StrAccessor string `json:"strAccessor,omitempty"`
	NStart      int64  `json:"nStart,omitempty"`
	NCount      int64  `json:"nCount,omitempty"`
}

//Acquire subset of result-set elements by range.
//
//Returns specified nCount elements contained in the specified result-set beginning from position nStart.
//
//	Parameters:
//	- strAccessor	(string) result-set ID, identifier of the server-side ordered collection of found hosts
//	- nStart	(int64) zero-based start position
//	- nCount	(int64) number of elements to return
//	- [out]	pChunk	(params) container that has needed elements in the array with name "KLCSP_ITERATOR_ARRAY"
//	Returns:
//	- (int64) actual number of returned elements (less or equal to nCount)
func (ca *ChunkAccessor) GetItemsChunk(ctx context.Context, params ItemsChunkParams, result interface{}) ([]byte,
	error) {
	postData, _ := json.Marshal(params)

	request, err := http.NewRequest("POST", ca.client.Server+"/api/v1.0/ChunkAccessor.GetItemsChunk", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := ca.client.Do(ctx, request, &result)
	return raw, err
}
