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

//	EventProcessing Class Reference
//
//	Interface implements the functionality for viewing and deleting events.
//
//	List of all members.
type EventProcessing service

//	Get record count in the result-set.
//
//	Returns number of elements contained in the specified result-set.
//
//	Parameters:
//	- strIteratorId	(string) result-set ID, identifier of the server-side ordered collection of found data records.
//
//	Returns:
//	- (int64) number of elements contained in the specified result-set.
func (ts *EventProcessing) GetRecordCount(ctx context.Context, strIteratorId string) (*PxgValInt, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strIteratorId": "%s"}`, strIteratorId))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/EventProcessing.GetRecordCount", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := ts.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//	Acquire subset of result-set elements by range.
//
//	Returns elements contained in the specified result-set in the diapason from position nStart to position nEnd.
//
//	Parameters:
//	- strIteratorId	(string) result-set ID, identifier of the server-side ordered collection of found data records.
//	- nStart	(int64) zero-based start position.
//	- nEnd	(int64) zero-based finish position.
//
//	Return:
//	pParamsEvents	(params) - container that has needed elements in the array with name "KLEVP_EVENT_RANGE_ARRAY".
func (ts *EventProcessing) GetRecordRange(ctx context.Context, strIteratorId string, nStart, nEnd int64) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"strIteratorId": "%s", "nStart": %d, "nEnd": %d}`, strIteratorId, nStart, nEnd))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/EventProcessing.GetRecordRange", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}

//	Release result-set.
//
//	Releases the specified result-set and frees associated memory.
//
//	Parameters:
//	- strIteratorId	(string) result-set ID,
//	identifier of the server-side ordered collection of found data records.
func (ts *EventProcessing) ReleaseIterator(ctx context.Context, strIteratorId string) (*PxgValInt, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strIteratorId": "%s"}`, strIteratorId))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/EventProcessing.ReleaseIterator", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := ts.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//	TODO Initiate delete events in the result-set.
//
//	Initiates mass delete of the events specified by pSettings in the result-set.
//
//	Parameters:
//	- strIteratorId	(string) result-set ID, identifier of the server-side ordered collection of found data records.
//	- pSettings	(params) contains range blocks to mass delete events.
//
//	Format of pSettings parameter:
//
//	EVP_EVENT_BLOCKS // paramArray
//    +---0 // paramParams
//        +---EVP_EVENT_BLOCK_BEGIN_ID // int64, write here Event ID (event_db_id) to set start of block, see List of event attributes for attribute names.
//        +---EVP_EVENT_BLOCK_END_ID // int64, write here Event ID (event_db_id) to set end of block, see List of event attributes for attribute names.
//    +---1 // paramParams
//        ...
//    ...
func (ts *EventProcessing) InitiateDelete(ctx context.Context, params interface{}) ([]byte, error) {
	return nil, nil
}

//	TODO Cancel delete events in the result-set.
//
//	Cancels mass delete of the events specified by pSettings in the result-set.
//
//	Parameters:
//	- strIteratorId	(string) result-set ID, identifier of the server-side ordered collection
//	of found data records.
//	- pSettings	(params) contains range blocks to mass delete events.
//
//	Format of pSettings parameter:
//
//	EVP_EVENT_BLOCKS // paramArray
//    +---0 // paramParams
//        +---EVP_EVENT_BLOCK_BEGIN_ID // int64, write here Event ID (event_db_id) to set start of block, see List of event attributes for attribute names.
//        +---EVP_EVENT_BLOCK_END_ID // int64, write here Event ID (event_db_id) to set end of block, see List of event attributes for attribute names.
//    +---1 // paramParams
//        ...
//    ...
//
func (ts *EventProcessing) CancelDelete(ctx context.Context, params interface{}) ([]byte, error) {
	return nil, nil
}
