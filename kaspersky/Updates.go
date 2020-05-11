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

//	Updates Class Reference
//
//	Updates processing. More...
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
func (uda *Updates) GetAvailableUpdatesInfo(ctx context.Context, strLocalization string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strLocalization": "%s"}`, strLocalization))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/Updates.GetAvailableUpdatesInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uda.client.Do(ctx, request, nil)
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
func (uda *Updates) GetUpdatesInfo(ctx context.Context) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"pFilter" : ["KAS20EXCH"]}`))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/Updates.GetUpdatesInfo",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}
	raw, err := uda.client.Do(ctx, request, nil)
	return raw, err
}

//	Asynchronously remove updates.
//
//	Returns:
//	- (string) request id of asynchronous operation, to cancel call Updates.RemoveUpdatesCancel,
//	to get status call AsyncActionStateChecker.CheckActionState with returned request id as action guid
func (uda *Updates) RemoveUpdates(ctx context.Context) (*RequestID, []byte, error) {
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/Updates.RemoveUpdates",
		nil)
	if err != nil {
		return nil, nil, err
	}

	requestID := new(RequestID)
	raw, err := uda.client.Do(ctx, request, &requestID)
	return requestID, raw, err
}

//	Cancel asynchronous operation RemoveUpdates.
//
//	Parameters:
//	- strRequestId	(string) request id of asynchronous operation Updates.RemoveUpdates
func (uda *Updates) RemoveUpdatesCancel(ctx context.Context, strRequestId string) (*RequestID, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strRequestId": "%s"}`, strRequestId))
	request, err := http.NewRequest("POST", uda.client.Server+"/api/v1.0/Updates.RemoveUpdatesCancel",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	requestID := new(RequestID)
	raw, err := uda.client.Do(ctx, request, &requestID)
	return requestID, raw, err
}
