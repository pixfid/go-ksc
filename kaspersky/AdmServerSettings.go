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

//AdmServerSettings interface.
//
//Interface to manage server settings.
//
//List of all members:
type AdmServerSettings service

//GetSharedFolder
//Acquire shared folder.
//
//	Returns:
//	- (string) shared folder
func (as *AdmServerSettings) GetSharedFolder(ctx context.Context) (*PxgValStr, []byte, error) {
	request, err := http.NewRequest("POST", as.client.Server+"/api/v1.0/AdmServerSettings.GetSharedFolder", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := as.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//ChangeSharedFolder
//Change shared folder.
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
//	Check the operation state by calling AsyncActionStateChecker.CheckActionState periodically until it's finalized.
//	If the operation success, then AsyncActionStateChecker.CheckActionState will return bFinalized=true and lStateCode=1.
//	Otherwise, a call to AsyncActionStateChecker.CheckActionState returns error in pStateData.
//
//	Exceptions:
//	- Throws	exception in case of error.
func (as *AdmServerSettings) ChangeSharedFolder(ctx context.Context, wstrNetworkPath string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrNetworkPath": "%s"}`, wstrNetworkPath))
	request, err := http.NewRequest("POST", as.client.Server+"/api/v1.0/AdmServerSettings.ChangeSharedFolder", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := as.client.Do(ctx, request, nil)
	return raw, err
}
