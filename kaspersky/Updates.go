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

//	Updates Class Reference
//
//	Updates processing..
//
//	List of all members:
type Updates service

//	Get available updates info.
//
//	Parameters:
//	- strLocalization	(string) - localization to be used for compoments titles;
//	can be one of "ru", "en", "fr", "de"; for other values "en" localization will be used.
//	Return:
//	- pAvailableUpdateComps	(params) - see Well-known retranslated update components list
func (upd *Updates) GetAvailableUpdatesInfo(ctx context.Context, strLocalization string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strLocalization": "%s"}`, strLocalization))
	request, err := http.NewRequest("POST", upd.client.Server+"/api/v1.0/Updates.GetAvailableUpdatesInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := upd.client.Do(ctx, request, nil)
	return raw, err
}

//	Returns a list of retransmissions.
//
//	Parameters:
//	- pFilter	(array) array of strings, contains a list of bundle identities (
//	KLUPDSRV_BUNDLE_ID) for which you want to return data. If NULL then a list of all retranslated updates will be returned, but the only one returned attribute for all updates is KLUPDSRV_BUNDLE_ID.
//
//	Returns:
//(array) result, see List of attributes of a retranslated update component bundle
func (upd *Updates) GetUpdatesInfo(ctx context.Context) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"pFilter" : ["KAS20EXCH"]}`))
	request, err := http.NewRequest("POST", upd.client.Server+"/api/v1.0/Updates.GetUpdatesInfo",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}
	raw, err := upd.client.Do(ctx, request, nil)
	return raw, err
}

//	Asynchronously remove updates.
//
//	Returns:
//	- (string) request id of asynchronous operation, to cancel call Updates.RemoveUpdatesCancel,
//	to get status call AsyncActionStateChecker.CheckActionState with returned request id as action guid
func (upd *Updates) RemoveUpdates(ctx context.Context) (*RequestID, []byte, error) {
	request, err := http.NewRequest("POST", upd.client.Server+"/api/v1.0/Updates.RemoveUpdates",
		nil)
	if err != nil {
		return nil, nil, err
	}

	requestID := new(RequestID)
	raw, err := upd.client.Do(ctx, request, &requestID)
	return requestID, raw, err
}

//	Cancel asynchronous operation RemoveUpdates.
//
//	Parameters:
//	- strRequestId	(string) request id of asynchronous operation Updates.RemoveUpdates
func (upd *Updates) RemoveUpdatesCancel(ctx context.Context, strRequestId string) (*RequestID, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strRequestId": "%s"}`, strRequestId))
	request, err := http.NewRequest("POST", upd.client.Server+"/api/v1.0/Updates.RemoveUpdatesCancel",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	requestID := new(RequestID)
	raw, err := upd.client.Do(ctx, request, &requestID)
	return requestID, raw, err
}
