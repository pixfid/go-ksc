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

//	AdmServerSettings interface.
//
//	Interface to manage server settings.
//
//	List of all members:
type AdmServerSettings service

//	Acquire shared folder.
//
//	Returns:
//	- (string) shared folder
func (as *AdmServerSettings) GetSharedFolder(ctx context.Context) (*PxgValStr, []byte, error) {
	request, err := http.NewRequest("POST", as.client.Server+"/api/v1.0/AdmServerSettings.GetSharedFolder", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	pxgValStr := new(PxgValStr)

	raw, err := as.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//	Change shared folder.
//
//	Parameters:
//	- wstrNetworkPath	(string) network path to shared folder
//
//	Returns:
//	- (string) id of asynchronous operation.
//
//Example:
//	"\\\\Server-ksc\\klshare\\"
//
//	Remarks:
//	Check the operation state by calling AsyncActionStateChecker::CheckActionState periodically until it's finalized.
//	If the operation success, then AsyncActionStateChecker::CheckActionState will return bFinalized=true and lStateCode=1.
//	Otherwise, a call to AsyncActionStateChecker::CheckActionState returns error in pStateData.
//
//	Exceptions:
//	- Throws	exception in case of error.
func (as *AdmServerSettings) ChangeSharedFolder(ctx context.Context, wstrNetworkPath string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrNetworkPath": "%s"}`, wstrNetworkPath))
	request, err := http.NewRequest("POST", as.client.Server+"/api/v1.0/AdmServerSettings.ChangeSharedFolder", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := as.client.Do(ctx, request, nil)
	return raw, err
}
