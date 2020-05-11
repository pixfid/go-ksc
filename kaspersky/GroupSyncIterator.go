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
	"fmt"
	"net/http"
)

//	GroupSyncIterator Class Reference
//
//	Access to the group synchronization forward iterator for the result-set. More...
//
//	List of all members.
type GroupSyncIterator service

//Releases the result-set.
//
//Releases the specified result-set and frees associated memory
//
//Parameters:
//	- szwIterator	(string) result-set ID
func (ca *GroupSyncIterator) ReleaseIterator(ctx context.Context, szwIterator string) bool {
	postData := []byte(fmt.Sprintf(`
	{
	"szwIterator": "%s"
	}`, szwIterator))
	request, err := http.NewRequest("POST", ca.client.Server+"/api/v1.0/GroupSyncIterator.ReleaseIterator", bytes.NewBuffer(postData))
	if err != nil {
		return false
	}

	raw, err := ca.client.Do(ctx, request, nil)
	if raw != nil {
		return true
	}
	return false
}

//	Acquire subset of elements contained in the result-set
//
//	Returns nCount elements contained in the specified result-set beginning from the current position and moves internal pointer to the new position.
//
//	Parameters:
//	- szwIterator	(string) forward iterator id
//	- nCount	(int64) number of elements to return
//
//Returns:
//	- int actual number of returned elements (less or equal to nCount)
//	- bEOF	(bool) returns false if the returned chunk is the last one,
//	and there's no need in further calls of this method
//	- pData	(params) container that has needed elements in the array with name "KLCSP_ITERATOR_ARRAY"
func (ca *GroupSyncIterator) GetNextItems(ctx context.Context, szwIterator string, nCount int64, v interface{}) (
	[]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
		"szwIterator": "%s",
		"nCount": %d
	}`, szwIterator, nCount))
	request, err := http.NewRequest("POST", ca.client.Server+"/api/v1.0/GroupSyncIterator.GetNextItems", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ca.client.Do(ctx, request, &v)
	return raw, err
}
