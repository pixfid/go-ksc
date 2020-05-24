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
	"fmt"
	"net/http"
)

//	GroupSyncIterator Class Reference
//
//	Access to the group synchronization forward iterator for the result-set..
//
//	List of all members.
type GroupSyncIterator service

//Releases the result-set.
//
//Releases the specified result-set and frees associated memory
//
//Parameters:
//	- szwIterator	(string) result-set ID
func (ca *GroupSyncIterator) ReleaseIterator(ctx context.Context, szwIterator string) error {
	postData := []byte(fmt.Sprintf(`
	{
	"szwIterator": "%s"
	}`, szwIterator))
	request, err := http.NewRequest("POST", ca.client.Server+"/api/v1.0/GroupSyncIterator.ReleaseIterator", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = ca.client.Do(ctx, request, nil)
	if err != nil {
		return err
	}
	return nil
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
