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
	"log"
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
		log.Fatal(err.Error())
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
		log.Fatal(err.Error())
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
		log.Fatal(err.Error())
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
//        +---EVP_EVENT_BLOCK_BEGIN_ID // paramLong, write here Event ID (event_db_id) to set start of block, see List of event attributes for attribute names.
//        +---EVP_EVENT_BLOCK_END_ID // paramLong, write here Event ID (event_db_id) to set end of block, see List of event attributes for attribute names.
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
//        +---EVP_EVENT_BLOCK_BEGIN_ID // paramLong, write here Event ID (event_db_id) to set start of block, see List of event attributes for attribute names.
//        +---EVP_EVENT_BLOCK_END_ID // paramLong, write here Event ID (event_db_id) to set end of block, see List of event attributes for attribute names.
//    +---1 // paramParams
//        ...
//    ...
//
func (ts *EventProcessing) CancelDelete(ctx context.Context, params interface{}) ([]byte, error) {
	return nil, nil
}
